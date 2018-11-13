// Order of translation.
//
// NOTE: the substeps of 4b can be done concurrently. NOTE: step 5-7 can be done
// concurrently.
//
// 1. Index AST top-level entities.
//
// 2. Resolve IR type definitions.
//
//    a) Index type identifiers and create scaffolding IR type definitions
//       (without bodies).
//
//    b) Translate AST type definitions to IR.
//
// 3. Translate AST comdat definitions to IR.
//
// 4. Resolve remaining IR top-level entities.
//
//    a) Index top-level identifiers and create scaffolding IR top-level
//       declarations and definitions (without bodies but with types).
//
//       1. Index global identifiers and create scaffolding IR global
//          declarations and definitions, alias and IFunc definitions, and
//          function declarations and definitions (without bodies but with
//          types).
//
//       2. Index attribute group IDs and create scaffolding IR attribute group
//          definitions (without bodies).
//
//       3. Index metadata names and create scaffolding IR named metadata
//          definitions (without bodies).
//
//       4. Index metadata IDs and create scaffolding IR metadata definitions
//          (without bodies).
//
//    b) Translate AST top-level declarations and definitions to IR.
//
//       NOTE: the substeps of 4b can be done concurrently.
//
//       1. Translate AST global declarations and definitions, alias and IFunc
//          definitions, and function declarations and definitions to IR.
//
//       2. Translate AST attribute group definitions to IR.
//
//       3. Translate AST named metadata definitions to IR.
//
//       4. Translate AST metadata definitions to IR.
//
// NOTE: step 5-7 can be done concurrenty.
//
// 5. Translate use-list orders.
//
// 6. Translate basic block specific use-list orders.
//
// 7. Fix basic block references in blockaddress constants.
//
// 8. Add IR top-level declarations and definitions to the module in order of
//    occurrence in the input.

package asm

import (
	"fmt"

	"github.com/llir/ll/ast"
	"github.com/llir/llvm/internal/enc"
	"github.com/llir/llvm/ir"
	"github.com/pkg/errors"
)

// TODO: remove flag after we reach our performance goals.
var (
	// DoTypeResolution enables type resolution of type defintions.
	DoTypeResolution = true
	// DoGlobalResolution enables global resolution of global variable and
	// function delcarations and defintions.
	DoGlobalResolution = true
)

// translate translates the given AST module into an equivalent IR module.
func translate(old *ast.Module) (*ir.Module, error) {
	gen := newGenerator()
	// 1. Index AST top-level entities.
	if err := gen.indexTopLevelEntities(old); err != nil {
		return nil, errors.WithStack(err)
	}
	// 2. Resolve IR type definitions.
	if err := gen.resolveTypeDefs(); err != nil {
		return nil, errors.WithStack(err)
	}
	// 3. Translate AST comdat definitions to IR.
	if err := gen.translateComdatDefs(); err != nil {
		return nil, errors.WithStack(err)
	}
	// 4a. Index top-level identifiers and create scaffolding IR top-level
	//     declarations and definitions (without bodies but with types).
	if err := gen.createTopLevelEntities(); err != nil {
		return nil, errors.WithStack(err)
	}
	// 4b. Translate AST top-level declarations and definitions to IR.
	// NOTE: the substeps of 4b can be done concurrently.
	if err := gen.translateTopLevelEntities(); err != nil {
		return nil, errors.WithStack(err)
	}
	// NOTE: step 5-7 can be done concurrenty.
	// 5. Translate use-list orders.
	for _, oldUseListOrder := range gen.old.useListOrders {
		useListOrder, err := gen.irUseListOrder(*oldUseListOrder)
		if err != nil {
			return nil, errors.WithStack(err)
		}
		gen.m.UseListOrders = append(gen.m.UseListOrders, useListOrder)
	}
	// 6. Translate basic block specific use-list orders.
	for _, oldUseListOrderBB := range gen.old.useListOrderBBs {
		useListOrderBB, err := gen.irUseListOrderBB(*oldUseListOrderBB)
		if err != nil {
			return nil, errors.WithStack(err)
		}
		gen.m.UseListOrderBBs = append(gen.m.UseListOrderBBs, useListOrderBB)
	}
	// 7. Fix basic block references in blockaddress constants.
	for _, c := range gen.todo {
		if err := fixBlockAddressConst(c); err != nil {
			return nil, errors.WithStack(err)
		}
	}
	// 8. Add IR top-level declarations and definitions to the module in order of
	//    occurrence in the input.
	gen.addDefsToModule()
	return gen.m, nil
}

// addDefsToModule adds IR top-level declarations and definitions to the module
// in order of occurrence in the input.
func (gen *generator) addDefsToModule() {
	for _, ident := range gen.old.typeDefOrder {
		def, ok := gen.new.typeDefs[ident]
		if !ok {
			panic(fmt.Errorf("unable to locate type identifier %q", enc.Local(ident)))
		}
		gen.m.TypeDefs = append(gen.m.TypeDefs, def)
	}
	for _, name := range gen.old.comdatDefOrder {
		def, ok := gen.new.comdatDefs[name]
		if !ok {
			panic(fmt.Errorf("unable to locate comdat name %q", enc.Comdat(name)))
		}
		gen.m.ComdatDefs = append(gen.m.ComdatDefs, def)
	}
	for _, ident := range gen.old.globalOrder {
		v, ok := gen.new.globals[ident]
		if !ok {
			panic(fmt.Errorf("unable to locate global identifier %q", enc.Global(ident)))
		}
		def, ok := v.(*ir.Global)
		if !ok {
			panic(fmt.Errorf("invalid global declaration or definition type; expected *ir.Global, got %T", v))
		}
		gen.m.Globals = append(gen.m.Globals, def)
	}
	for _, ident := range gen.old.indirectSymbolDefOrder {
		v, ok := gen.new.globals[ident]
		if !ok {
			panic(fmt.Errorf("unable to locate global identifier %q", enc.Global(ident)))
		}
		switch v := v.(type) {
		case *ir.Alias:
			gen.m.Aliases = append(gen.m.Aliases, v)
		case *ir.IFunc:
			gen.m.IFuncs = append(gen.m.IFuncs, v)
		default:
			panic(fmt.Errorf("invalid indirect symbol definition type; expected *ir.Alias or *ir.IFunc, got %T", v))

		}
	}
	for _, ident := range gen.old.funcOrder {
		v, ok := gen.new.globals[ident]
		if !ok {
			panic(fmt.Errorf("unable to locate global identifier %q", enc.Global(ident)))
		}
		def, ok := v.(*ir.Function)
		if !ok {
			panic(fmt.Errorf("invalid function declaration or definition type; expected *ir.Function, got %T", v))
		}
		gen.m.Funcs = append(gen.m.Funcs, def)
	}
	for _, id := range gen.old.attrGroupDefOrder {
		def, ok := gen.new.attrGroupDefs[id]
		if !ok {
			panic(fmt.Errorf("unable to locate attribute group ID %q", enc.AttrGroupID(id)))
		}
		gen.m.AttrGroupDefs = append(gen.m.AttrGroupDefs, def)
	}
	for _, name := range gen.old.namedMetadataDefOrder {
		def, ok := gen.new.namedMetadataDefs[name]
		if !ok {
			panic(fmt.Errorf("unable to locate metadata name %q", enc.Metadata(name)))
		}
		gen.m.NamedMetadataDefs = append(gen.m.NamedMetadataDefs, def)
	}
	for _, id := range gen.old.metadataDefOrder {
		def, ok := gen.new.metadataDefs[id]
		if !ok {
			panic(fmt.Errorf("unable to locate metadata ID %q", enc.Metadata(id)))
		}
		gen.m.MetadataDefs = append(gen.m.MetadataDefs, def)
	}
}
