package asm

import (
	"fmt"

	"github.com/llir/l/ir"
	"github.com/mewmew/l-tm/asm/ll/ast"
)

// --- [ Vector instructions ] -------------------------------------------------

// ~~~ [ extractelement ] ~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~

func (fgen *funcGen) astToIRInstExtractElement(inst ir.Instruction, old *ast.ExtractElementInst) (*ir.InstExtractElement, error) {
	i, ok := inst.(*ir.InstExtractElement)
	if !ok {
		// NOTE: panic since this would indicate a bug in the implementation.
		panic(fmt.Errorf("invalid IR instruction for AST instruction; expected *ir.InstExtractElement, got %T", inst))
	}
	// TODO: implement
	return i, nil
}

// ~~~ [ insertelement ] ~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~

func (fgen *funcGen) astToIRInstInsertElement(inst ir.Instruction, old *ast.InsertElementInst) (*ir.InstInsertElement, error) {
	i, ok := inst.(*ir.InstInsertElement)
	if !ok {
		// NOTE: panic since this would indicate a bug in the implementation.
		panic(fmt.Errorf("invalid IR instruction for AST instruction; expected *ir.InstInsertElement, got %T", inst))
	}
	// TODO: implement
	return i, nil
}

// ~~~ [ shufflevector ] ~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~

func (fgen *funcGen) astToIRInstShuffleVector(inst ir.Instruction, old *ast.ShuffleVectorInst) (*ir.InstShuffleVector, error) {
	i, ok := inst.(*ir.InstShuffleVector)
	if !ok {
		// NOTE: panic since this would indicate a bug in the implementation.
		panic(fmt.Errorf("invalid IR instruction for AST instruction; expected *ir.InstShuffleVector, got %T", inst))
	}
	// TODO: implement
	return i, nil
}
