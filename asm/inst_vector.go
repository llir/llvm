package asm

import (
	"fmt"

	"github.com/llir/ll/ast"
	"github.com/llir/llvm/ir"
	"github.com/pkg/errors"
)

// --- [ Vector instructions ] -------------------------------------------------

// ~~~ [ extractelement ] ~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~

// astToIRInstExtractElement translates the given AST extractelement instruction
// into an equivalent IR instruction.
func (fgen *funcGen) astToIRInstExtractElement(inst ir.Instruction, old *ast.ExtractElementInst) (*ir.InstExtractElement, error) {
	i, ok := inst.(*ir.InstExtractElement)
	if !ok {
		panic(fmt.Errorf("invalid IR instruction for AST instruction; expected *ir.InstExtractElement, got %T", inst))
	}
	// Vector.
	x, err := fgen.astToIRTypeValue(old.X())
	if err != nil {
		return nil, errors.WithStack(err)
	}
	i.X = x
	// Element index.
	index, err := fgen.astToIRTypeValue(old.Index())
	if err != nil {
		return nil, errors.WithStack(err)
	}
	i.Index = index
	// (optional) Metadata.
	md, err := fgen.gen.irMetadataAttachments(old.Metadata())
	if err != nil {
		return nil, errors.WithStack(err)
	}
	i.Metadata = md
	return i, nil
}

// ~~~ [ insertelement ] ~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~

// astToIRInstInsertElement translates the given AST insertelement instruction
// into an equivalent IR instruction.
func (fgen *funcGen) astToIRInstInsertElement(inst ir.Instruction, old *ast.InsertElementInst) (*ir.InstInsertElement, error) {
	i, ok := inst.(*ir.InstInsertElement)
	if !ok {
		panic(fmt.Errorf("invalid IR instruction for AST instruction; expected *ir.InstInsertElement, got %T", inst))
	}
	// Vector.
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
	// Element index.
	index, err := fgen.astToIRTypeValue(old.Index())
	if err != nil {
		return nil, errors.WithStack(err)
	}
	i.Index = index
	// (optional) Metadata.
	md, err := fgen.gen.irMetadataAttachments(old.Metadata())
	if err != nil {
		return nil, errors.WithStack(err)
	}
	i.Metadata = md
	return i, nil
}

// ~~~ [ shufflevector ] ~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~

// astToIRInstShuffleVector translates the given AST shufflevector instruction
// into an equivalent IR instruction.
func (fgen *funcGen) astToIRInstShuffleVector(inst ir.Instruction, old *ast.ShuffleVectorInst) (*ir.InstShuffleVector, error) {
	i, ok := inst.(*ir.InstShuffleVector)
	if !ok {
		panic(fmt.Errorf("invalid IR instruction for AST instruction; expected *ir.InstShuffleVector, got %T", inst))
	}
	// X vector.
	x, err := fgen.astToIRTypeValue(old.X())
	if err != nil {
		return nil, errors.WithStack(err)
	}
	i.X = x
	// Y vector.
	y, err := fgen.astToIRTypeValue(old.Y())
	if err != nil {
		return nil, errors.WithStack(err)
	}
	i.Y = y
	// Shuffle mask.
	mask, err := fgen.astToIRTypeValue(old.Mask())
	if err != nil {
		return nil, errors.WithStack(err)
	}
	i.Mask = mask
	// (optional) Metadata.
	md, err := fgen.gen.irMetadataAttachments(old.Metadata())
	if err != nil {
		return nil, errors.WithStack(err)
	}
	i.Metadata = md
	return i, nil
}
