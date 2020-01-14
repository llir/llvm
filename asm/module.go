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
	id := int64(0)
	// 1. Index AST top-level entities.
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
					return errors.Errorf("type identifier %q already present; prev `%s`, new `%s`", enc.TypeName(name), text(prev), text(entity))
				}
			}
			gen.old.typeDefs[name] = entity
		case *ast.ComdatDef:
			name := comdatName(entity.Name())
			if prev, ok := gen.old.comdatDefs[name]; ok {
				return errors.Errorf("comdat name %q already present; prev `%s`, new `%s`", enc.ComdatName(name), text(prev), text(entity))
			}
			gen.old.comdatDefs[name] = entity
		case *ast.GlobalDecl:
			ident := giveUnnamedIdentID(globalIdent(entity.Name()), &id)
			if prev, ok := gen.old.globals[ident]; ok {
				return errors.Errorf("global identifier %q already present; prev `%s`, new `%s`", ident.Ident(), text(prev), text(entity))
			}
			gen.old.globals[ident] = entity
			gen.old.globalOrder = append(gen.old.globalOrder, ident)
		case *ast.IndirectSymbolDef:
			ident := giveUnnamedIdentID(globalIdent(entity.Name()), &id)
			if prev, ok := gen.old.globals[ident]; ok {
				return errors.Errorf("global identifier %q already present; prev `%s`, new `%s`", ident.Ident(), text(prev), text(entity))
			}
			gen.old.globals[ident] = entity
			gen.old.globalOrder = append(gen.old.globalOrder, ident)
		case *ast.FuncDecl:
			ident := giveUnnamedIdentID(globalIdent(entity.Header().Name()), &id)
			if prev, ok := gen.old.globals[ident]; ok {
				return errors.Errorf("global identifier %q already present; prev `%s`, new `%s`", ident.Ident(), text(prev), text(entity))
			}
			gen.old.globals[ident] = entity
			gen.old.globalOrder = append(gen.old.globalOrder, ident)
		case *ast.FuncDef:
			ident := giveUnnamedIdentID(globalIdent(entity.Header().Name()), &id)
			if prev, ok := gen.old.globals[ident]; ok {
				return errors.Errorf("global identifier %q already present; prev `%s`, new `%s`", ident.Ident(), text(prev), text(entity))
			}
			gen.old.globals[ident] = entity
			gen.old.globalOrder = append(gen.old.globalOrder, ident)
		case *ast.AttrGroupDef:
			id := attrGroupID(entity.ID())
			// Append attribute group definition, and merge at later stage if ID
			// maps to more than one attribute group definition.
			gen.old.attrGroupDefs[id] = append(gen.old.attrGroupDefs[id], entity)
		case *ast.NamedMetadataDef:
			name := metadataName(entity.Name())
			// Multiple named metadata definitions of the same name are allowed.
			// They are merged into a single named metadata definition with the
			// nodes of each definition appended.
			gen.old.namedMetadataDefs[name] = append(gen.old.namedMetadataDefs[name], entity)
		case *ast.MetadataDef:
			id := metadataID(entity.ID())
			if prev, ok := gen.old.metadataDefs[id]; ok {
				return errors.Errorf("metadata ID %q already present; prev `%s`, new `%s`", enc.MetadataID(id), text(prev), text(entity))
			}
			gen.old.metadataDefs[id] = entity
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

// giveUnnamedIdentID assigns an unused ID to the global identifier if unnamed.
func giveUnnamedIdentID(ident ir.GlobalIdent, id *int64) ir.GlobalIdent {
	if ident.IsUnnamed() {
		// Assign next unused ID to unnamed global identifier.
		ident.SetID(*id)
		*id++
	}
	return ident
}

// === [ Create and index IR ] =================================================

// createTopLevelEntities indexes IR top-level identifiers and creates
// scaffolding IR top-level declarations and definitions (without bodies but
// with types) of the given module.
func (gen *generator) createTopLevelEntities() error {
	// 4a. Index top-level identifiers and create scaffolding IR top-level
	//     declarations and definitions (without bodies but with types).
	//
	// Note: the substeps of 4a can be done concurrently.
	//
	// 4a1. Index global identifiers and create scaffolding IR global
	//      declarations and definitions, indirect symbol definitions (aliases
	//      and indirect functions), and function declarations and definitions
	//      (without bodies but with types).
	if err := gen.createGlobalEntities(); err != nil {
		return errors.WithStack(err)
	}
	// 4a2. Index attribute group IDs and create scaffolding IR attribute group
	//      definitions (without bodies).
	gen.createAttrGroupDefs()
	// 4a3. Index metadata names and create scaffolding IR named metadata
	//      definitions (without bodies).
	gen.createNamedMetadataDefs()
	// 4a4. Index metadata IDs and create scaffolding IR metadata definitions
	//      (without bodies).
	gen.createMetadataDefs()
	return nil
}

// --- [ Attribute group definitions ] -----------------------------------------

// createAttrGroupDefs indexes IR attribute group IDs and creates scaffolding IR
// attribute group definitions (without bodies) of the given module.
//
// post-condition: gen.new.attrGroupDefs maps from attribute group ID (without
// '#' prefix) to corresponding skeleton IR value.
func (gen *generator) createAttrGroupDefs() {
	// 4a2. Index attribute group IDs and create scaffolding IR attribute group
	//      definitions (without bodies).
	for id := range gen.old.attrGroupDefs {
		new := &ir.AttrGroupDef{ID: id}
		gen.new.attrGroupDefs[id] = new
	}
}

// --- [ Named metadata definitions ] ------------------------------------------

// createNamedMetadataDefs indexes IR metadata names and creates scaffolding IR
// named metadata definitions (without bodies) of the given module.
//
// post-condition: gen.new.namedMetadataDefs maps from metadata name (without
// '!' prefix) to corresponding skeleton IR value.
func (gen *generator) createNamedMetadataDefs() {
	// 4a3. Index metadata names and create scaffolding IR named metadata
	//      definitions (without bodies).
	for name := range gen.old.namedMetadataDefs {
		new := &metadata.NamedDef{Name: name}
		gen.new.namedMetadataDefs[name] = new
	}
}

// --- [ Metadata definitions ] ------------------------------------------------

// createMetadataDefs indexes IR metadata IDs and creates scaffolding IR
// metadata definitions (without bodies) of the given module.
//
// post-condition: gen.new.metadataDefs maps from metadata ID (without '!'
// prefix) to corresponding skeleton IR value.
func (gen *generator) createMetadataDefs() {
	// 4a4. Index metadata IDs and create scaffolding IR metadata definitions
	//      (without bodies).
	for id, md := range gen.old.metadataDefs {
		new := newMetadataDef(id, md)
		gen.new.metadataDefs[id] = new
	}
}

// newMetadataDef returns a new IR metadata definition (without body) based on
// the given AST metadata definition.
func newMetadataDef(id int64, old *ast.MetadataDef) metadata.Definition {
	switch oldNode := old.MDNode().(type) {
	case *ast.MDTuple:
		new := &metadata.Tuple{}
		new.SetID(id)
		return new
	case ast.SpecializedMDNode:
		new := newSpecializedMDNode(oldNode)
		new.SetID(id)
		return new
	default:
		panic(fmt.Errorf("support for metadata node %T not yet implemented", oldNode))
	}
}

// newSpecializedMDNode returns a new IR specialized metadata node (without
// body) based on the given AST specialized metadata node.
func newSpecializedMDNode(old ast.SpecializedMDNode) metadata.SpecializedNode {
	switch old := old.(type) {
	case *ast.DIBasicType:
		return &metadata.DIBasicType{}
	case *ast.DICommonBlock:
		return &metadata.DICommonBlock{}
	case *ast.DICompileUnit:
		return &metadata.DICompileUnit{}
	case *ast.DICompositeType:
		return &metadata.DICompositeType{}
	case *ast.DIDerivedType:
		return &metadata.DIDerivedType{}
	case *ast.DIEnumerator:
		return &metadata.DIEnumerator{}
	case *ast.DIExpression:
		return &metadata.DIExpression{}
	case *ast.DIFile:
		return &metadata.DIFile{}
	case *ast.DIGlobalVariable:
		return &metadata.DIGlobalVariable{}
	case *ast.DIGlobalVariableExpression:
		return &metadata.DIGlobalVariableExpression{}
	case *ast.DIImportedEntity:
		return &metadata.DIImportedEntity{}
	case *ast.DILabel:
		return &metadata.DILabel{}
	case *ast.DILexicalBlock:
		return &metadata.DILexicalBlock{}
	case *ast.DILexicalBlockFile:
		return &metadata.DILexicalBlockFile{}
	case *ast.DILocalVariable:
		return &metadata.DILocalVariable{}
	case *ast.DILocation:
		return &metadata.DILocation{}
	case *ast.DIMacro:
		return &metadata.DIMacro{}
	case *ast.DIMacroFile:
		return &metadata.DIMacroFile{}
	case *ast.DIModule:
		return &metadata.DIModule{}
	case *ast.DINamespace:
		return &metadata.DINamespace{}
	case *ast.DIObjCProperty:
		return &metadata.DIObjCProperty{}
	case *ast.DISubprogram:
		return &metadata.DISubprogram{}
	case *ast.DISubrange:
		return &metadata.DISubrange{}
	case *ast.DISubroutineType:
		return &metadata.DISubroutineType{}
	case *ast.DITemplateTypeParameter:
		return &metadata.DITemplateTypeParameter{}
	case *ast.DITemplateValueParameter:
		return &metadata.DITemplateValueParameter{}
	case *ast.GenericDINode:
		return &metadata.GenericDINode{}
	default:
		panic(fmt.Errorf("support for %T not yet implemented", old))
	}
}

// === [ Translate AST to IR ] =================================================

// translateTopLevelEntities translates the AST top-level declarations and
// definitions of the given module to IR.
func (gen *generator) translateTopLevelEntities() error {
	// TODO: make concurrent and benchmark difference in walltime.

	// 4b. Translate AST top-level declarations and definitions to IR.
	//
	// Note: the substeps of 4b can be done concurrently.
	//
	// 4b1. Translate AST global declarations and definitions, alias and IFunc
	//      definitions, and function declarations and definitions to IR.
	if err := gen.translateGlobalEntities(); err != nil {
		return errors.WithStack(err)
	}
	// 4b2. Translate AST attribute group definitions to IR.
	gen.translateAttrGroupDefs()
	// 4b3. Translate AST named metadata definitions to IR.
	if err := gen.translateNamedMetadataDefs(); err != nil {
		return errors.WithStack(err)
	}
	// 4b4. Translate AST metadata definitions to IR.
	return gen.translateMetadataDefs()
}

// --- [ Comdat definitions ] --------------------------------------------------

// translateComdatDefs translates the AST comdat definitions of the given module
// to IR.
func (gen *generator) translateComdatDefs() {
	// 3. Translate AST comdat definitions to IR.
	//
	// Note: step 3 and the substeps of 4a can be done concurrently.
	for name, old := range gen.old.comdatDefs {
		new := &ir.ComdatDef{
			Name: name,
			Kind: asmenum.SelectionKindFromString(old.Kind().Text()),
		}
		gen.new.comdatDefs[name] = new
	}
}

// --- [ Attribute group definitions ] -----------------------------------------

// translateAttrGroupDefs translates the AST attribute group definitions of the
// given module to IR.
func (gen *generator) translateAttrGroupDefs() {
	// 4b2. Translate AST attribute group definitions to IR.
	for id, old := range gen.old.attrGroupDefs {
		new, ok := gen.new.attrGroupDefs[id]
		if !ok {
			panic(fmt.Errorf("unable to locate attribute group ID %q", enc.AttrGroupID(id)))
		}
		gen.irAttrGroupDef(new, old)
	}
}

// irAttrGroupDef translates the AST attribute group definitions (one or more)
// to an equivalent IR attribute group definition. Mulriple definitions of the
// same ID are merged into a single attribute group definition.
func (gen *generator) irAttrGroupDef(new *ir.AttrGroupDef, oldDefs []*ast.AttrGroupDef) {
	// present is used to prevent duplicate attributes when merging multiple
	// attribute group definitions.
	present := make(map[string]bool)
	for _, oldDef := range oldDefs {
		for _, oldFuncAttr := range oldDef.FuncAttrs() {
			lit := oldFuncAttr.LlvmNode().Text()
			if present[lit] {
				// skip duplicate attribute.
				continue
			}
			funcAttr := gen.irFuncAttribute(oldFuncAttr)
			new.FuncAttrs = append(new.FuncAttrs, funcAttr)
			present[lit] = true
		}
	}
}

// --- [ Named metadata definitions ] ------------------------------------------

// translateNamedMetadataDefs translates the AST named metadata definitions of
// the given module to IR.
func (gen *generator) translateNamedMetadataDefs() error {
	// 4b3. Translate AST named metadata definitions to IR.
	for name, old := range gen.old.namedMetadataDefs {
		new, ok := gen.new.namedMetadataDefs[name]
		if !ok {
			panic(fmt.Errorf("unable to locate metadata name %q", enc.MetadataName(name)))
		}
		for _, oldDef := range old {
			if err := gen.irNamedMetadataDef(new, oldDef); err != nil {
				return errors.WithStack(err)
			}
		}
	}
	return nil
}

// irNamedMetadataDef translates the given AST named metadata definition to an
// equivalent IR named metadata definition.
func (gen *generator) irNamedMetadataDef(new *metadata.NamedDef, old *ast.NamedMetadataDef) error {
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
	// 4b4. Translate AST metadata definitions to IR.
	for id, old := range gen.old.metadataDefs {
		new, ok := gen.new.metadataDefs[id]
		if !ok {
			panic(fmt.Errorf("unable to locate metadata ID %q", enc.MetadataID(id)))
		}
		if err := gen.irMetadataDef(new, old); err != nil {
			return errors.WithStack(err)
		}
	}
	return nil
}

// irMetadataDef translates the given AST metadata definition to an equivalent
// IR metadata definition.
func (gen *generator) irMetadataDef(new metadata.Definition, old *ast.MetadataDef) error {
	// (optional) Distinct.
	if _, ok := old.Distinct(); ok {
		new.SetDistinct(true)
	}
	switch oldNode := old.MDNode().(type) {
	case *ast.MDTuple:
		_, err := gen.irMDTuple(new, oldNode)
		if err != nil {
			return errors.WithStack(err)
		}
	case ast.SpecializedMDNode:
		_, err := gen.irSpecializedMDNode(new, oldNode)
		if err != nil {
			return errors.WithStack(err)
		}
	default:
		panic(fmt.Errorf("support for metadata node %T not yet implemented", old))
	}
	return nil
}

// --- [ Use-list orders ] -----------------------------------------------------

// translateUseListOrders translates the AST use-list orders of the given
// module to IR.
func (gen *generator) translateUseListOrders() error {
	// 5. Translate use-list orders.
	if len(gen.old.useListOrders) > 0 {
		gen.m.UseListOrders = make([]*ir.UseListOrder, len(gen.old.useListOrders))
		for i, oldUseListOrder := range gen.old.useListOrders {
			useListOrder, err := gen.irUseListOrder(oldUseListOrder)
			if err != nil {
				return errors.WithStack(err)
			}
			gen.m.UseListOrders[i] = useListOrder
		}
	}
	return nil
}

// irUseListOrder returns the IR use-list order corresponding to the given AST
// use-list order.
func (gen *generator) irUseListOrder(old *ast.UseListOrder) (*ir.UseListOrder, error) {
	// Value.
	oldVal := old.Val()
	typ, err := gen.irType(oldVal.Typ())
	if err != nil {
		return nil, errors.WithStack(err)
	}
	oldConst, ok := oldVal.Val().(ast.Constant)
	if !ok {
		panic(fmt.Errorf("support for use-list order value %T not yet implemented", oldVal.Val()))
	}
	c, err := gen.irConstant(typ, oldConst)
	if err != nil {
		return nil, errors.WithStack(err)
	}
	// Indices.
	indices := uintSlice(old.Indices())
	useListOrder := &ir.UseListOrder{
		Value:   c,
		Indices: indices,
	}
	return useListOrder, nil
}

// --- [ Basic block specific use-list orders ] --------------------------------

// translateUseListOrderBBs translates the AST basic block specific use-list
// orders of the given module to IR.
func (gen *generator) translateUseListOrderBBs() error {
	// 6. Translate basic block specific use-list orders.
	if len(gen.old.useListOrderBBs) > 0 {
		gen.m.UseListOrderBBs = make([]*ir.UseListOrderBB, len(gen.old.useListOrderBBs))
		for i, oldUseListOrderBB := range gen.old.useListOrderBBs {
			useListOrderBB, err := gen.irUseListOrderBB(oldUseListOrderBB)
			if err != nil {
				return errors.WithStack(err)
			}
			gen.m.UseListOrderBBs[i] = useListOrderBB
		}
	}
	return nil
}

// irUseListOrderBB translates the given AST basic block specific use-list order
// to an equivalent IR basic block specific use-list order.
func (gen *generator) irUseListOrderBB(old *ast.UseListOrderBB) (*ir.UseListOrderBB, error) {
	// Function.
	funcIdent := globalIdent(old.Func())
	v, ok := gen.new.globals[funcIdent]
	if !ok {
		return nil, errors.Errorf("unable to locate global identifier %q", funcIdent.Ident())
	}
	f, ok := v.(*ir.Func)
	if !ok {
		return nil, errors.Errorf("invalid function type of %q; expected *ir.Func, got %T", funcIdent.Ident(), v)
	}
	// Basic block.
	blockIdent := localIdent(old.Block())
	block, err := findBlock(f, blockIdent)
	if err != nil {
		return nil, errors.WithStack(err)
	}
	// Indices.
	indices := uintSlice(old.Indices())
	useListOrderBB := &ir.UseListOrderBB{
		Func:    f,
		Block:   block,
		Indices: indices,
	}
	return useListOrderBB, nil
}
