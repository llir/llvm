// Order of translation.
//
// Note: step 3 and the substeps of 4a can be done concurrently.
// Note: the substeps of 4a can be done concurrently.
// Note: the substeps of 4b can be done concurrently.
// Note: steps 5-7 can be done concurrently.
// Note: the substeps of 8 can be done concurrently.
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
// Note: step 3 and the substeps of 4a can be done concurrently.
//
// 4. Resolve remaining IR top-level entities.
//
//    a) Index top-level identifiers and create scaffolding IR top-level
//       declarations and definitions (without bodies but with types).
//
//       Note: the substeps of 4a can be done concurrently.
//
//       1. Index global identifiers and create scaffolding IR global
//          declarations and definitions, indirect symbol definitions, and
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
//       Note: the substeps of 4b can be done concurrently.
//
//       1. Translate AST global declarations and definitions, indirect symbol
//          definitions, and function declarations and definitions to IR.
//
//       2. Translate AST attribute group definitions to IR.
//
//       3. Translate AST named metadata definitions to IR.
//
//       4. Translate AST metadata definitions to IR.
//
// Note: steps 5-7 can be done concurrenty.
//
// 5. Translate use-list orders.
//
// 6. Translate basic block specific use-list orders.
//
// 7. Fix basic block references in blockaddress constants.
//
// 8. Add IR top-level declarations and definitions to the IR module in order of
//    occurrence in the input.
//
//    Note: the substeps of 8 can be done concurrently.
//
//    a) Add IR type definitions to the IR module in natural sorting order.
//
//    b) Add IR comdat definitions to the IR module in natural sorting order.
//
//    c) Add IR global variable declarations and definitions, indirect symbol
//       definitions, and function declarations and definitions to the IR module
//       in order of occurrence in the input.
//
//    d) Add IR attribute group definitions to the IR module in numeric order.
//
//    e) Add IR named metadata definitions to the IR module.
//
//    f) Add IR metadata definitions to the IR module in numeric order.

package asm

import (
	"fmt"
	"sort"
	"time"

	"github.com/llir/ll/ast"
	"github.com/llir/llvm/internal/enc"
	"github.com/llir/llvm/internal/natsort"
	"github.com/llir/llvm/ir"
	"github.com/llir/llvm/ir/constant"
	"github.com/llir/llvm/ir/metadata"
	"github.com/llir/llvm/ir/types"
	"github.com/pkg/errors"
)

// translate translates the given AST module into an equivalent IR module.
func translate(old *ast.Module) (*ir.Module, error) {
	gen := newGenerator()
	// 1. Index AST top-level entities.
	indexStart := time.Now()
	if err := gen.translateTargetDefs(old); err != nil {
		return nil, errors.WithStack(err)
	}
	if err := gen.indexTopLevelEntities(old); err != nil {
		return nil, errors.WithStack(err)
	}
	dbg.Println("index AST top-level entities took:", time.Since(indexStart))
	// 2. Resolve IR type definitions.
	typeStart := time.Now()
	if err := gen.resolveTypeDefs(); err != nil {
		return nil, errors.WithStack(err)
	}
	dbg.Println("type resolution took:", time.Since(typeStart))
	// 3. Translate AST comdat definitions to IR.
	//
	// Note: step 3 and the substeps of 4a can be done concurrently.
	gen.translateComdatDefs()
	// 4. Resolve remaining IR top-level entities.
	//
	// 4a. Index top-level identifiers and create scaffolding IR top-level
	//     declarations and definitions (without bodies but with types).
	createStart := time.Now()
	if err := gen.createTopLevelEntities(); err != nil {
		return nil, errors.WithStack(err)
	}
	dbg.Println("create IR top-level entities took:", time.Since(createStart))
	// 4b. Translate AST top-level declarations and definitions to IR.
	//
	// Note: the substeps of 4b can be done concurrently.
	translateStart := time.Now()
	if err := gen.translateTopLevelEntities(); err != nil {
		return nil, errors.WithStack(err)
	}
	dbg.Println("translate AST to IR took:", time.Since(translateStart))
	// Note: step 5-7 can be done concurrenty.
	//
	// 5. Translate use-list orders.
	if err := gen.translateUseListOrders(); err != nil {
		return nil, errors.WithStack(err)
	}
	// 6. Translate basic block specific use-list orders.
	if err := gen.translateUseListOrderBBs(); err != nil {
		return nil, errors.WithStack(err)
	}
	// 7. Fix basic block references in blockaddress constants.
	for _, c := range gen.todo {
		if err := fixBlockAddressConst(c); err != nil {
			return nil, errors.WithStack(err)
		}
	}
	// 8. Add IR top-level declarations and definitions to the IR module in order
	//    of occurrence in the input.
	//
	// Note: the substeps of 8 can be done concurrently.
	addStart := time.Now()
	gen.addDefsToModule()
	dbg.Println("add IR definitions to IR module took:", time.Since(addStart))
	return gen.m, nil
}

// addDefsToModule adds IR top-level declarations and definitions to the IR
// module in order of occurrence in the input.
func (gen *generator) addDefsToModule() {
	// 8. Add IR top-level declarations and definitions to the IR module in order
	//    of occurrence in the input.
	//
	// Note: the substeps of 8 can be done concurrently.
	//
	// 8a. Add IR type definitions to the IR module in natural sorting order.
	gen.addTypeDefsToModule()
	// 8b. Add IR comdat definitions to the IR module in natural sorting order.
	gen.addComdatDefsToModule()
	// 8c. Add IR global variable declarations and definitions, indirect symbol
	//     definitions, and function declarations and definitions to the IR
	//     module in order of occurrence in the input.
	gen.addGlobalEntitiesToModule()
	// 8d. Add IR attribute group definitions to the IR module in numeric order.
	gen.addAttrGroupDefsToModule()
	// 8e. Add IR named metadata definitions to the IR module.
	gen.addNamedMetadataDefsToModule()
	// 8f. Add IR metadata definitions to the IR module in numeric order.
	gen.addMetadataDefsToModule()
}

// addTypeDefsToModule adds IR type definitions to the IR module in natural
// sorting order.
func (gen *generator) addTypeDefsToModule() {
	// 8a. Add IR type definitions to the IR module in natural sorting order.
	typeNames := make([]string, 0, len(gen.old.typeDefs))
	for name := range gen.old.typeDefs {
		typeNames = append(typeNames, name)
	}
	natsort.Strings(typeNames)
	if len(typeNames) > 0 {
		gen.m.TypeDefs = make([]types.Type, len(typeNames))
		for i, name := range typeNames {
			def, ok := gen.new.typeDefs[name]
			if !ok {
				panic(fmt.Errorf("unable to locate type identifier %q", enc.TypeName(name)))
			}
			gen.m.TypeDefs[i] = def
		}
	}
}

// addComdatDefsToModule adds IR comdat definitions to the IR module in natural
// sorting order.
func (gen *generator) addComdatDefsToModule() {
	// 8b. Add IR comdat definitions to the IR module in natural sorting order.
	comdatNames := make([]string, 0, len(gen.old.comdatDefs))
	for name := range gen.old.comdatDefs {
		comdatNames = append(comdatNames, name)
	}
	natsort.Strings(comdatNames)
	if len(comdatNames) > 0 {
		gen.m.ComdatDefs = make([]*ir.ComdatDef, len(comdatNames))
		for i, name := range comdatNames {
			def, ok := gen.new.comdatDefs[name]
			if !ok {
				panic(fmt.Errorf("unable to locate comdat name %q", enc.ComdatName(name)))
			}
			gen.m.ComdatDefs[i] = def
		}
	}
}

// addGlobalEntitiesToModule adds IR global variable declarations and
// definitions, indirect symbol definitions, and function declarations and
// definitions to the IR module in order of occurrence in the input.
func (gen *generator) addGlobalEntitiesToModule() {
	// 8c. Add IR global variable declarations and definitions, indirect symbol
	//     definitions, and function declarations and definitions to the IR
	//     module in order of occurrence in the input.
	for _, ident := range gen.old.globalOrder {
		v, ok := gen.new.globals[ident]
		if !ok {
			panic(fmt.Errorf("unable to locate global identifier %q", ident.Ident()))
		}
		switch def := v.(type) {
		case *ir.Global:
			gen.m.Globals = append(gen.m.Globals, def)
		case *ir.Alias:
			gen.m.Aliases = append(gen.m.Aliases, def)
		case *ir.IFunc:
			gen.m.IFuncs = append(gen.m.IFuncs, def)
		case *ir.Func:
			gen.m.Funcs = append(gen.m.Funcs, def)
		default:
			panic(fmt.Errorf("support for global %T not yet implemented", v))
		}
	}
}

// addAttrGroupDefsToModule adds IR attribute group definitions to the IR module
// in numeric order.
func (gen *generator) addAttrGroupDefsToModule() {
	// 8d. Add IR attribute group definitions to the IR module in numeric order.
	attrGroupIDs := make([]int64, 0, len(gen.old.attrGroupDefs))
	for id := range gen.old.attrGroupDefs {
		attrGroupIDs = append(attrGroupIDs, id)
	}
	less := func(i, j int) bool {
		return attrGroupIDs[i] < attrGroupIDs[j]
	}
	sort.Slice(attrGroupIDs, less)
	if len(attrGroupIDs) > 0 {
		gen.m.AttrGroupDefs = make([]*ir.AttrGroupDef, len(attrGroupIDs))
		for i, id := range attrGroupIDs {
			def, ok := gen.new.attrGroupDefs[id]
			if !ok {
				panic(fmt.Errorf("unable to locate attribute group ID %q", enc.AttrGroupID(id)))
			}
			gen.m.AttrGroupDefs[i] = def
		}
	}
}

// addNamedMetadataDefsToModule adds IR named metadata definitions to the IR
// module.
func (gen *generator) addNamedMetadataDefsToModule() {
	// 8e. Add IR named metadata definitions to the IR module.
	for name, def := range gen.new.namedMetadataDefs {
		gen.m.NamedMetadataDefs[name] = def
	}
}

// addMetadataDefsToModule adds IR metadata definitions to the IR module in
// numeric order.
func (gen *generator) addMetadataDefsToModule() {
	metadataIDs := make([]int64, 0, len(gen.old.metadataDefs))
	for id := range gen.old.metadataDefs {
		metadataIDs = append(metadataIDs, id)
	}
	less := func(i, j int) bool {
		return metadataIDs[i] < metadataIDs[j]
	}
	sort.Slice(metadataIDs, less)
	if len(metadataIDs) > 0 {
		gen.m.MetadataDefs = make([]metadata.Definition, len(metadataIDs))
		for i, id := range metadataIDs {
			def, ok := gen.new.metadataDefs[id]
			if !ok {
				panic(fmt.Errorf("unable to locate metadata ID %q", enc.MetadataID(id)))
			}
			gen.m.MetadataDefs[i] = def
		}
	}
}

// ### [ Helper functions ] ####################################################

// fixBlockAddressConst fixes the basic block of the given blockaddress
// constant. During translation of constants, blockaddress constants are
// assigned dummy basic blocks since function bodies have yet to be translated.
//
// pre-condition: translated function body and assigned local IDs of c.Func.
func fixBlockAddressConst(c *constant.BlockAddress) error {
	f, ok := c.Func.(*ir.Func)
	if !ok {
		panic(fmt.Errorf("invalid function type in blockaddress constant; expected *ir.Func, got %T", c.Func))
	}
	bb, ok := c.Block.(*ir.Block)
	if !ok {
		panic(fmt.Errorf("invalid basic block type in blockaddress constant; expected *ir.Block, got %T", c.Block))
	}
	block, err := findBlock(f, bb.LocalIdent)
	if err != nil {
		return errors.WithStack(err)
	}
	c.Block = block
	return nil
}
