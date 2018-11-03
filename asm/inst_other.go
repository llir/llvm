package asm

import (
	"fmt"

	"github.com/llir/llvm/ir"
	"github.com/llir/ll/ast"
)

// --- [ Other instructions ] --------------------------------------------------

// ~~~ [ icmp ] ~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~

func (fgen *funcGen) astToIRInstICmp(inst ir.Instruction, old *ast.ICmpInst) (*ir.InstICmp, error) {
	i, ok := inst.(*ir.InstICmp)
	if !ok {
		// NOTE: panic since this would indicate a bug in the implementation.
		panic(fmt.Errorf("invalid IR instruction for AST instruction; expected *ir.InstICmp, got %T", inst))
	}
	// TODO: implement
	return i, nil
}

// ~~~ [ fcmp ] ~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~

func (fgen *funcGen) astToIRInstFCmp(inst ir.Instruction, old *ast.FCmpInst) (*ir.InstFCmp, error) {
	i, ok := inst.(*ir.InstFCmp)
	if !ok {
		// NOTE: panic since this would indicate a bug in the implementation.
		panic(fmt.Errorf("invalid IR instruction for AST instruction; expected *ir.InstFCmp, got %T", inst))
	}
	// Fast math flags.
	i.FastMathFlags = irFastMathFlags(old.FastMathFlags())
	// TODO: implement
	return i, nil
}

// ~~~ [ phi ] ~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~

func (fgen *funcGen) astToIRInstPhi(inst ir.Instruction, old *ast.PhiInst) (*ir.InstPhi, error) {
	i, ok := inst.(*ir.InstPhi)
	if !ok {
		// NOTE: panic since this would indicate a bug in the implementation.
		panic(fmt.Errorf("invalid IR instruction for AST instruction; expected *ir.InstPhi, got %T", inst))
	}
	// TODO: implement
	return i, nil
}

// ~~~ [ select ] ~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~

func (fgen *funcGen) astToIRInstSelect(inst ir.Instruction, old *ast.SelectInst) (*ir.InstSelect, error) {
	i, ok := inst.(*ir.InstSelect)
	if !ok {
		// NOTE: panic since this would indicate a bug in the implementation.
		panic(fmt.Errorf("invalid IR instruction for AST instruction; expected *ir.InstSelect, got %T", inst))
	}
	// TODO: implement
	return i, nil
}

// ~~~ [ call ] ~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~

func (fgen *funcGen) astToIRInstCall(inst ir.Instruction, old *ast.CallInst) (*ir.InstCall, error) {
	i, ok := inst.(*ir.InstCall)
	if !ok {
		// NOTE: panic since this would indicate a bug in the implementation.
		panic(fmt.Errorf("invalid IR instruction for AST instruction; expected *ir.InstCall, got %T", inst))
	}
	// Fast math flags.
	i.FastMathFlags = irFastMathFlags(old.FastMathFlags())
	// Calling convention.
	i.CallingConv = irOptCallingConv(old.CallingConv())
	return i, nil
}

// ~~~ [ va_arg ] ~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~

func (fgen *funcGen) astToIRInstVAArg(inst ir.Instruction, old *ast.VAArgInst) (*ir.InstVAArg, error) {
	i, ok := inst.(*ir.InstVAArg)
	if !ok {
		// NOTE: panic since this would indicate a bug in the implementation.
		panic(fmt.Errorf("invalid IR instruction for AST instruction; expected *ir.InstVAArg, got %T", inst))
	}
	// TODO: implement
	return i, nil
}

// ~~~ [ landingpad ] ~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~

func (fgen *funcGen) astToIRInstLandingPad(inst ir.Instruction, old *ast.LandingPadInst) (*ir.InstLandingPad, error) {
	i, ok := inst.(*ir.InstLandingPad)
	if !ok {
		// NOTE: panic since this would indicate a bug in the implementation.
		panic(fmt.Errorf("invalid IR instruction for AST instruction; expected *ir.InstLandingPad, got %T", inst))
	}
	// TODO: implement
	return i, nil
}

// ~~~ [ catchpad ] ~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~

func (fgen *funcGen) astToIRInstCatchPad(inst ir.Instruction, old *ast.CatchPadInst) (*ir.InstCatchPad, error) {
	i, ok := inst.(*ir.InstCatchPad)
	if !ok {
		// NOTE: panic since this would indicate a bug in the implementation.
		panic(fmt.Errorf("invalid IR instruction for AST instruction; expected *ir.InstCatchPad, got %T", inst))
	}
	// TODO: implement
	return i, nil
}

// ~~~ [ cleanuppad ] ~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~

func (fgen *funcGen) astToIRInstCleanupPad(inst ir.Instruction, old *ast.CleanupPadInst) (*ir.InstCleanupPad, error) {
	i, ok := inst.(*ir.InstCleanupPad)
	if !ok {
		// NOTE: panic since this would indicate a bug in the implementation.
		panic(fmt.Errorf("invalid IR instruction for AST instruction; expected *ir.InstCleanupPad, got %T", inst))
	}
	// TODO: implement
	return i, nil
}
