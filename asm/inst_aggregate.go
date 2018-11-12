package asm

import (
	"fmt"

	"github.com/llir/ll/ast"
	"github.com/llir/llvm/ir"
	"github.com/pkg/errors"
)

// --- [ Aggregate instructions ] ----------------------------------------------

// ~~~ [ extractvalue ] ~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~

// astToIRInstExtractValue translates the given AST extractvalue instruction
// into an equivalent IR instruction.
func (fgen *funcGen) astToIRInstExtractValue(inst ir.Instruction, old *ast.ExtractValueInst) (*ir.InstExtractValue, error) {
	i, ok := inst.(*ir.InstExtractValue)
	if !ok {
		panic(fmt.Errorf("invalid IR instruction for AST instruction; expected *ir.InstExtractValue, got %T", inst))
	}
	// Aggregate value.
	x, err := fgen.astToIRTypeValue(old.X())
	if err != nil {
		return nil, errors.WithStack(err)
	}
	i.X = x
	// Element indices.
	indices := uintSlice(old.Indices())
	for _, index := range indices {
		i.Indices = append(i.Indices, int64(index))
	}
	// (optional) Metadata.
	md, err := fgen.gen.irMetadataAttachments(old.Metadata())
	if err != nil {
		return nil, errors.WithStack(err)
	}
	i.Metadata = md
	return i, nil
}

// ~~~ [ insertvalue ] ~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~

// astToIRInstInsertValue translates the given AST insertvalue instruction into
// an equivalent IR instruction.
func (fgen *funcGen) astToIRInstInsertValue(inst ir.Instruction, old *ast.InsertValueInst) (*ir.InstInsertValue, error) {
	i, ok := inst.(*ir.InstInsertValue)
	if !ok {
		panic(fmt.Errorf("invalid IR instruction for AST instruction; expected *ir.InstInsertValue, got %T", inst))
	}
	// Aggregate value.
	x, err := fgen.astToIRTypeValue(old.X())
	if err != nil {
		return nil, errors.WithStack(err)
	}
	i.X = x
	// Element to insert.
	elem, err := fgen.astToIRTypeValue(old.Elem())
	if err != nil {
		return nil, errors.WithStack(err)
	}
	i.Elem = elem
	// Element indices.
	indices := uintSlice(old.Indices())
	for _, index := range indices {
		i.Indices = append(i.Indices, int64(index))
	}
	// (optional) Metadata.
	md, err := fgen.gen.irMetadataAttachments(old.Metadata())
	if err != nil {
		return nil, errors.WithStack(err)
	}
	i.Metadata = md
	return i, nil
}
