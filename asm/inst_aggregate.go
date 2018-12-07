package asm

import (
	"fmt"

	"github.com/llir/ll/ast"
	"github.com/llir/llvm/ir"
	"github.com/llir/llvm/ir/types"
	"github.com/pkg/errors"
)

// === [ Create IR ] ===========================================================

// newExtractValueInst returns a new IR extractvalue instruction (without body
// but with type) based on the given AST extractvalue instruction.
func (fgen *funcGen) newExtractValueInst(ident ir.LocalIdent, old *ast.ExtractValueInst) (*ir.InstExtractValue, error) {
	xType, err := fgen.gen.irType(old.X().Typ())
	if err != nil {
		return nil, errors.WithStack(err)
	}
	indices := uintSlice(old.Indices())
	typ := aggregateElemType(xType, indices)
	return &ir.InstExtractValue{LocalIdent: ident, Typ: typ}, nil
}

// newInsertValueInst returns a new IR insertvalue instruction (without body but
// with type) based on the given AST insertvalue instruction.
func (fgen *funcGen) newInsertValueInst(ident ir.LocalIdent, old *ast.InsertValueInst) (*ir.InstInsertValue, error) {
	typ, err := fgen.gen.irType(old.X().Typ())
	if err != nil {
		return nil, errors.WithStack(err)
	}
	return &ir.InstInsertValue{LocalIdent: ident, Typ: typ}, nil
}

// === [ Translate AST to IR ] =================================================

// --- [ extractvalue ] --------------------------------------------------------

// irExtractValueInst translates the given AST extractvalue instruction into an
// equivalent IR instruction.
func (fgen *funcGen) irExtractValueInst(new ir.Instruction, old *ast.ExtractValueInst) error {
	inst, ok := new.(*ir.InstExtractValue)
	if !ok {
		panic(fmt.Errorf("invalid IR instruction for AST instruction; expected *ir.InstExtractValue, got %T", new))
	}
	// Aggregate value.
	x, err := fgen.irTypeValue(old.X())
	if err != nil {
		return errors.WithStack(err)
	}
	inst.X = x
	// Element indices.
	inst.Indices = uintSlice(old.Indices())
	// (optional) Metadata.
	md, err := fgen.gen.irMetadataAttachments(old.Metadata())
	if err != nil {
		return errors.WithStack(err)
	}
	inst.Metadata = md
	return nil
}

// --- [ insertvalue ] ---------------------------------------------------------

// irInsertValueInst translates the given AST insertvalue instruction into an
// equivalent IR instruction.
func (fgen *funcGen) irInsertValueInst(new ir.Instruction, old *ast.InsertValueInst) error {
	inst, ok := new.(*ir.InstInsertValue)
	if !ok {
		panic(fmt.Errorf("invalid IR instruction for AST instruction; expected *ir.InstInsertValue, got %T", new))
	}
	// Aggregate value.
	x, err := fgen.irTypeValue(old.X())
	if err != nil {
		return errors.WithStack(err)
	}
	inst.X = x
	// Element to insert.
	elem, err := fgen.irTypeValue(old.Elem())
	if err != nil {
		return errors.WithStack(err)
	}
	inst.Elem = elem
	// Element indices.
	inst.Indices = uintSlice(old.Indices())
	// (optional) Metadata.
	md, err := fgen.gen.irMetadataAttachments(old.Metadata())
	if err != nil {
		return errors.WithStack(err)
	}
	inst.Metadata = md
	return nil
}

// ### [ Helper functions ] ####################################################

// aggregateElemType returns the element type at the position in the aggregate
// type specified by the given indices.
func aggregateElemType(t types.Type, indices []uint64) types.Type {
	// Base case.
	if len(indices) == 0 {
		return t
	}
	switch t := t.(type) {
	case *types.ArrayType:
		return aggregateElemType(t.ElemType, indices[1:])
	case *types.StructType:
		return aggregateElemType(t.Fields[indices[0]], indices[1:])
	default:
		panic(fmt.Errorf("support for aggregate type %T not yet implemented", t))
	}
}
