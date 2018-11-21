package asm

import (
	"fmt"

	"github.com/llir/ll/ast"
	asmenum "github.com/llir/llvm/asm/enum"
	"github.com/llir/llvm/internal/enc"
	"github.com/llir/llvm/ir"
	"github.com/llir/llvm/ir/metadata"
	"github.com/pkg/errors"
)

// === [ Index AST ] ===========================================================

// indexTopLevelEntities indexes the AST top-level entities of the given module.
func (gen *generator) indexTopLevelEntities(old *ast.Module) error {
	// Index AST top-level entities.
	// track added type definitions.
	addedTypeDef := make(map[string]bool)
	for _, entity := range old.TopLevelEntities() {
		switch entity := entity.(type) {
		case *ast.SourceFilename:
			gen.m.SourceFilename = unquote(entity.Name().Text())
		case *ast.TargetDataLayout:
			gen.m.DataLayout = unquote(entity.DataLayout().Text())
		case *ast.TargetTriple:
			gen.m.TargetTriple = unquote(entity.TargetTriple().Text())
		case *ast.ModuleAsm:
			asm := unquote(entity.Asm().Text())
			gen.m.ModuleAsms = append(gen.m.ModuleAsms, asm)
		case *ast.TypeDef:
			ident := localIdent(entity.Name())
			name := getTypeName(ident)
			if prev, ok := gen.old.typeDefs[name]; ok {
				if _, ok := prev.Typ().(*ast.OpaqueType); !ok {
					return errors.Errorf("type identifier %q already present; prev `%s`, new `%s`", enc.Local(name), text(prev), text(entity))
				}
			}
			gen.old.typeDefs[name] = entity
			if !addedTypeDef[name] {
				// Only record the first type definition of each type name.
				//
				// Type definitions of opaque types may contain several type
				// definitions with the same type name.
				gen.old.typeDefOrder = append(gen.old.typeDefOrder, name)
			}
			addedTypeDef[name] = true
		case *ast.ComdatDef:
			name := comdatName(entity.Name())
			if prev, ok := gen.old.comdatDefs[name]; ok {
				return errors.Errorf("comdat name %q already present; prev `%s`, new `%s`", enc.Comdat(name), text(prev), text(entity))
			}
			gen.old.comdatDefs[name] = entity
			gen.old.comdatDefOrder = append(gen.old.comdatDefOrder, name)
		case *ast.GlobalDecl:
			ident := globalIdent(entity.Name())
			if prev, ok := gen.old.globals[ident]; ok {
				return errors.Errorf("global identifier %q already present; prev `%s`, new `%s`", enc.Global(ident), text(prev), text(entity))
			}
			gen.old.globals[ident] = entity
			gen.old.globalOrder = append(gen.old.globalOrder, ident)
		case *ast.GlobalDef:
			ident := globalIdent(entity.Name())
			if prev, ok := gen.old.globals[ident]; ok {
				return errors.Errorf("global identifier %q already present; prev `%s`, new `%s`", enc.Global(ident), text(prev), text(entity))
			}
			gen.old.globals[ident] = entity
			gen.old.globalOrder = append(gen.old.globalOrder, ident)
		case *ast.IndirectSymbolDef:
			ident := globalIdent(entity.Name())
			if prev, ok := gen.old.globals[ident]; ok {
				return errors.Errorf("global identifier %q already present; prev `%s`, new `%s`", enc.Global(ident), text(prev), text(entity))
			}
			gen.old.globals[ident] = entity
			gen.old.indirectSymbolDefOrder = append(gen.old.indirectSymbolDefOrder, ident)
		case *ast.FuncDecl:
			ident := globalIdent(entity.Header().Name())
			if prev, ok := gen.old.globals[ident]; ok {
				return errors.Errorf("global identifier %q already present; prev `%s`, new `%s`", enc.Global(ident), text(prev), text(entity))
			}
			gen.old.globals[ident] = entity
			gen.old.funcOrder = append(gen.old.funcOrder, ident)
		case *ast.FuncDef:
			ident := globalIdent(entity.Header().Name())
			if prev, ok := gen.old.globals[ident]; ok {
				return errors.Errorf("global identifier %q already present; prev `%s`, new `%s`", enc.Global(ident), text(prev), text(entity))
			}
			gen.old.globals[ident] = entity
			gen.old.funcOrder = append(gen.old.funcOrder, ident)
		case *ast.AttrGroupDef:
			id := attrGroupID(entity.ID())
			if prev, ok := gen.old.attrGroupDefs[id]; ok {
				return errors.Errorf("attribute group ID %q already present; prev `%s`, new `%s`", enc.AttrGroupID(id), text(prev), text(entity))
			}
			gen.old.attrGroupDefs[id] = entity
			gen.old.attrGroupDefOrder = append(gen.old.attrGroupDefOrder, id)
		case *ast.NamedMetadataDef:
			name := metadataName(entity.Name())
			if prev, ok := gen.old.namedMetadataDefs[name]; ok {
				return errors.Errorf("metadata name %q already present; prev `%s`, new `%s`", enc.Metadata(name), text(prev), text(entity))
			}
			gen.old.namedMetadataDefs[name] = entity
			gen.old.namedMetadataDefOrder = append(gen.old.namedMetadataDefOrder, name)
		case *ast.MetadataDef:
			id := metadataID(entity.ID())
			if prev, ok := gen.old.metadataDefs[id]; ok {
				return errors.Errorf("metadata ID %q already present; prev `%s`, new `%s`", enc.Metadata(id), text(prev), text(entity))
			}
			gen.old.metadataDefs[id] = entity
			gen.old.metadataDefOrder = append(gen.old.metadataDefOrder, id)
		case *ast.UseListOrder:
			gen.old.useListOrders = append(gen.old.useListOrders, entity)
		case *ast.UseListOrderBB:
			gen.old.useListOrderBBs = append(gen.old.useListOrderBBs, entity)
		default:
			panic(fmt.Errorf("support for AST top-level entity %T not yet implemented", entity))
		}
	}
	return nil
}

// === [ Create and index IR ] =================================================

// createTopLevelEntities indexes top-level identifiers and create scaffolding
// IR top-level declarations and definitions (without bodies but with types) of
// the given module.
func (gen *generator) createTopLevelEntities() error {
	// 1. Index global identifiers and create scaffolding IR global declarations
	// and definitions, alias and IFunc definitions, and function declarations
	// and definitions (without bodies but with types).
	if err := gen.createGlobals(); err != nil {
		return errors.WithStack(err)
	}
	// 2. Index attribute group IDs and create scaffolding IR attribute group
	//    definitions (without bodies).
	if err := gen.createAttrGroupDefs(); err != nil {
		return errors.WithStack(err)
	}
	// 3. Index metadata names and create scaffolding IR named metadata definitions
	//    (without bodies).
	if err := gen.createNamedMetadataDefs(); err != nil {
		return errors.WithStack(err)
	}
	// 4. Index metadata IDs and create scaffolding IR metadata definitions (without
	//    bodies).
	if err := gen.createMetadataDefs(); err != nil {
		return errors.WithStack(err)
	}
	return nil
}

// --- [ Attribute group definitions ] -----------------------------------------

// createAttrGroupDefs indexes attribute group IDs and creates scaffolding IR
// attribute group definitions (without bodies) of the given module.
func (gen *generator) createAttrGroupDefs() error {
	for id := range gen.old.attrGroupDefs {
		new := &ir.AttrGroupDef{ID: id}
		gen.new.attrGroupDefs[id] = new
	}
	return nil
}

// --- [ Named metadata definitions ] ------------------------------------------

// createNamedMetadataDefs indexes metadata names and creates scaffolding IR
// named metadata definitions (without bodies) of the given module.
func (gen *generator) createNamedMetadataDefs() error {
	for name := range gen.old.namedMetadataDefs {
		new := &metadata.NamedMetadataDef{Name: name}
		gen.new.namedMetadataDefs[name] = new
	}
	return nil
}

// --- [ Metadata definitions ] ------------------------------------------------

// createMetadataDefs indexes metadata IDs and creates scaffolding IR metadata
// definitions (without bodies) of the given module.
func (gen *generator) createMetadataDefs() error {
	for id := range gen.old.metadataDefs {
		new := &metadata.MetadataDef{ID: id}
		gen.new.metadataDefs[id] = new
	}
	return nil
}

// === [ Translate AST to IR ] =================================================

// translateTopLevelEntities translates the AST top-level declarations and
// definitions of the given module to IR.
func (gen *generator) translateTopLevelEntities() error {
	// TODO: make concurrent and benchmark difference in walltime.

	// 1. Translate AST global declarations and definitions, alias and IFunc
	//    definitions, and function declarations and definitions to IR.
	if err := gen.translateGlobals(); err != nil {
		return errors.WithStack(err)
	}
	// 2. Translate AST attribute group definitions to IR.
	if err := gen.translateAttrGroupDefs(); err != nil {
		return errors.WithStack(err)
	}
	// 3. Translate AST named metadata definitions to IR.
	if err := gen.translateNamedMetadataDefs(); err != nil {
		return errors.WithStack(err)
	}
	// 4. Translate AST metadata definitions to IR.
	if err := gen.translateMetadataDefs(); err != nil {
		return errors.WithStack(err)
	}
	return nil
}

// --- [ Comdat definitions ] --------------------------------------------------

// translateComdatDefs translates the AST comdat definitions of the given module
// to IR.
func (gen *generator) translateComdatDefs() error {
	for name, old := range gen.old.comdatDefs {
		new := &ir.ComdatDef{
			Name: name,
			Kind: asmenum.SelectionKindFromString(old.Kind().Text()),
		}
		gen.new.comdatDefs[name] = new
	}
	return nil
}

// --- [ Attribute group definitions ] -----------------------------------------

// translateAttrGroupDefs translates the AST attribute group definitions of the
// given module to IR.
func (gen *generator) translateAttrGroupDefs() error {
	// TODO: make concurrent and benchmark.
	for id, old := range gen.old.attrGroupDefs {
		new, ok := gen.new.attrGroupDefs[id]
		if !ok {
			panic(fmt.Errorf("unable to locate attribute group ID %q", enc.AttrGroupID(id)))
		}
		if err := gen.translateAttrGroupDef(new, old); err != nil {
			return errors.WithStack(err)
		}
	}
	return nil
}

// translateAttrGroupDef translates the given AST attribute group definition to
// IR.
func (gen *generator) translateAttrGroupDef(new *ir.AttrGroupDef, old *ast.AttrGroupDef) error {
	for _, oldFuncAttr := range old.Attrs() {
		funcAttr := gen.irFuncAttribute(oldFuncAttr)
		new.FuncAttrs = append(new.FuncAttrs, funcAttr)
	}
	return nil
}

// --- [ Named metadata definitions ] ------------------------------------------

// translateNamedMetadataDefs translates the AST named metadata definitions of
// the given module to IR.
func (gen *generator) translateNamedMetadataDefs() error {
	// TODO: make concurrent and benchmark.
	for name, old := range gen.old.namedMetadataDefs {
		new, ok := gen.new.namedMetadataDefs[name]
		if !ok {
			panic(fmt.Errorf("unable to locate metadata name %q", enc.Metadata(name)))
		}
		if err := gen.translateNamedMetadataDef(new, old); err != nil {
			return errors.WithStack(err)
		}
	}
	return nil
}

// translateNamedMetadataDef translates the given AST named metadata definition
// to IR.
func (gen *generator) translateNamedMetadataDef(new *metadata.NamedMetadataDef, old *ast.NamedMetadataDef) error {
	// Nodes.
	for _, oldNode := range old.MDNodes() {
		node, err := gen.irMetadataNode(oldNode)
		if err != nil {
			return errors.WithStack(err)
		}
		new.Nodes = append(new.Nodes, node)
	}
	return nil
}

// --- [ Metadata definitions ] ------------------------------------------------

// translateMetadataDefs translates the AST metadata definitions of the given
// module to IR.
func (gen *generator) translateMetadataDefs() error {
	// TODO: make concurrent and benchmark.
	for id, old := range gen.old.metadataDefs {
		new, ok := gen.new.metadataDefs[id]
		if !ok {
			panic(fmt.Errorf("unable to locate metadata ID %q", enc.Metadata(id)))
		}
		if err := gen.translateMetadataDef(new, old); err != nil {
			return errors.WithStack(err)
		}
	}
	return nil
}

// translateMetadataDef translates the given AST metadata definition to IR.
func (gen *generator) translateMetadataDef(new *metadata.MetadataDef, old *ast.MetadataDef) error {
	// (optional) Distinct.
	new.Distinct = old.Distinct().IsValid()
	// Node.
	switch oldNode := old.MDNode().(type) {
	case *ast.MDTuple:
		node, err := gen.irMDTuple(oldNode)
		if err != nil {
			return errors.WithStack(err)
		}
		new.Node = node
	case ast.SpecializedMDNode:
		node, err := gen.irSpecializedMDNode(oldNode)
		if err != nil {
			return errors.WithStack(err)
		}
		new.Node = node
	default:
		panic(fmt.Errorf("support for metadata node %T not yet implemented", old))
	}
	return nil
}
