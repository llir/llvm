package asm

import (
	"fmt"

	"github.com/llir/ll/ast"
	"github.com/llir/llvm/ir"
)

// --- [ Other instructions ] --------------------------------------------------

// ~~~ [ icmp ] ~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~

// astToIRInstICmp translates the given AST icmp instruction into an equivalent
// IR instruction.
func (fgen *funcGen) astToIRInstICmp(inst ir.Instruction, old *ast.ICmpInst) (*ir.InstICmp, error) {
	i, ok := inst.(*ir.InstICmp)
	if !ok {
		panic(fmt.Errorf("invalid IR instruction for AST instruction; expected *ir.InstICmp, got %T", inst))
	}
	// TODO: implement
	// (optional) Metadata.
	i.Metadata = irMetadataAttachments(old.Metadata())
	return i, nil
}

// ~~~ [ fcmp ] ~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~

// astToIRInstFCmp translates the given AST fcmp instruction into an equivalent
// IR instruction.
func (fgen *funcGen) astToIRInstFCmp(inst ir.Instruction, old *ast.FCmpInst) (*ir.InstFCmp, error) {
	i, ok := inst.(*ir.InstFCmp)
	if !ok {
		panic(fmt.Errorf("invalid IR instruction for AST instruction; expected *ir.InstFCmp, got %T", inst))
	}
	// Fast math flags.
	i.FastMathFlags = irFastMathFlags(old.FastMathFlags())
	// TODO: implement
	// (optional) Metadata.
	i.Metadata = irMetadataAttachments(old.Metadata())
	return i, nil
}

// ~~~ [ phi ] ~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~

// astToIRInstPhi translates the given AST phi instruction into an equivalent IR
// instruction.
func (fgen *funcGen) astToIRInstPhi(inst ir.Instruction, old *ast.PhiInst) (*ir.InstPhi, error) {
	i, ok := inst.(*ir.InstPhi)
	if !ok {
		panic(fmt.Errorf("invalid IR instruction for AST instruction; expected *ir.InstPhi, got %T", inst))
	}
	// TODO: implement
	// (optional) Metadata.
	i.Metadata = irMetadataAttachments(old.Metadata())
	return i, nil
}

// ~~~ [ select ] ~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~

// astToIRInstSelect translates the given AST select instruction into an
// equivalent IR instruction.
func (fgen *funcGen) astToIRInstSelect(inst ir.Instruction, old *ast.SelectInst) (*ir.InstSelect, error) {
	i, ok := inst.(*ir.InstSelect)
	if !ok {
		panic(fmt.Errorf("invalid IR instruction for AST instruction; expected *ir.InstSelect, got %T", inst))
	}
	// TODO: implement
	// (optional) Metadata.
	i.Metadata = irMetadataAttachments(old.Metadata())
	return i, nil
}

// ~~~ [ call ] ~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~

// astToIRInstCall translates the given AST call instruction into an equivalent
// IR instruction.
func (fgen *funcGen) astToIRInstCall(inst ir.Instruction, old *ast.CallInst) (*ir.InstCall, error) {
	i, ok := inst.(*ir.InstCall)
	if !ok {
		panic(fmt.Errorf("invalid IR instruction for AST instruction; expected *ir.InstCall, got %T", inst))
	}
	// Fast math flags.
	i.FastMathFlags = irFastMathFlags(old.FastMathFlags())
	// Calling convention.
	i.CallingConv = irOptCallingConv(old.CallingConv())
	return i, nil
}

// ~~~ [ va_arg ] ~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~

// astToIRInstVAArg translates the given AST vaarg instruction into an
// equivalent IR instruction.
func (fgen *funcGen) astToIRInstVAArg(inst ir.Instruction, old *ast.VAArgInst) (*ir.InstVAArg, error) {
	i, ok := inst.(*ir.InstVAArg)
	if !ok {
		panic(fmt.Errorf("invalid IR instruction for AST instruction; expected *ir.InstVAArg, got %T", inst))
	}
	// TODO: implement
	// (optional) Metadata.
	i.Metadata = irMetadataAttachments(old.Metadata())
	return i, nil
}

// ~~~ [ landingpad ] ~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~

// astToIRInstLandingPad translates the given AST landingpad instruction into an
// equivalent IR instruction.
func (fgen *funcGen) astToIRInstLandingPad(inst ir.Instruction, old *ast.LandingPadInst) (*ir.InstLandingPad, error) {
	i, ok := inst.(*ir.InstLandingPad)
	if !ok {
		panic(fmt.Errorf("invalid IR instruction for AST instruction; expected *ir.InstLandingPad, got %T", inst))
	}
	// TODO: implement
	// (optional) Metadata.
	i.Metadata = irMetadataAttachments(old.Metadata())
	return i, nil
}

// ~~~ [ catchpad ] ~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~

// astToIRInstCatchPad translates the given AST catchpad instruction into an
// equivalent IR instruction.
func (fgen *funcGen) astToIRInstCatchPad(inst ir.Instruction, old *ast.CatchPadInst) (*ir.InstCatchPad, error) {
	i, ok := inst.(*ir.InstCatchPad)
	if !ok {
		panic(fmt.Errorf("invalid IR instruction for AST instruction; expected *ir.InstCatchPad, got %T", inst))
	}
	// TODO: implement
	// (optional) Metadata.
	i.Metadata = irMetadataAttachments(old.Metadata())
	return i, nil
}

// ~~~ [ cleanuppad ] ~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~

// astToIRInstCleanupPad translates the given AST cleanuppad instruction into an
// equivalent IR instruction.
func (fgen *funcGen) astToIRInstCleanupPad(inst ir.Instruction, old *ast.CleanupPadInst) (*ir.InstCleanupPad, error) {
	i, ok := inst.(*ir.InstCleanupPad)
	if !ok {
		panic(fmt.Errorf("invalid IR instruction for AST instruction; expected *ir.InstCleanupPad, got %T", inst))
	}
	// TODO: implement
	// (optional) Metadata.
	i.Metadata = irMetadataAttachments(old.Metadata())
	return i, nil
}
