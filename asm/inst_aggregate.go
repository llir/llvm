package asm

import (
	"fmt"

	"github.com/llir/llvm/ir"
	"github.com/llir/ll/ast"
)

// --- [ Aggregate instructions ] ----------------------------------------------

// ~~~ [ extractvalue ] ~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~

func (fgen *funcGen) astToIRInstExtractValue(inst ir.Instruction, old *ast.ExtractValueInst) (*ir.InstExtractValue, error) {
	i, ok := inst.(*ir.InstExtractValue)
	if !ok {
		// NOTE: panic since this would indicate a bug in the implementation.
		panic(fmt.Errorf("invalid IR instruction for AST instruction; expected *ir.InstExtractValue, got %T", inst))
	}
	// TODO: implement
	return i, nil
}

// ~~~ [ insertvalue ] ~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~

func (fgen *funcGen) astToIRInstInsertValue(inst ir.Instruction, old *ast.InsertValueInst) (*ir.InstInsertValue, error) {
	i, ok := inst.(*ir.InstInsertValue)
	if !ok {
		// NOTE: panic since this would indicate a bug in the implementation.
		panic(fmt.Errorf("invalid IR instruction for AST instruction; expected *ir.InstInsertValue, got %T", inst))
	}
	// TODO: implement
	return i, nil
}
