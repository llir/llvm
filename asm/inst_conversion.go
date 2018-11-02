package asm

import (
	"fmt"

	"github.com/llir/l/ir"
	"github.com/mewmew/l-tm/asm/ll/ast"
)

// --- [ Conversion instructions ] ---------------------------------------------

// ~~~ [ trunc ] ~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~

func (fgen *funcGen) astToIRInstTrunc(inst ir.Instruction, old *ast.TruncInst) (*ir.InstTrunc, error) {
	i, ok := inst.(*ir.InstTrunc)
	if !ok {
		// NOTE: panic since this would indicate a bug in the implementation.
		panic(fmt.Errorf("invalid IR instruction for AST instruction; expected *ir.InstTrunc, got %T", inst))
	}
	// TODO: implement
	return i, nil
}

// ~~~ [ zext ] ~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~

func (fgen *funcGen) astToIRInstZExt(inst ir.Instruction, old *ast.ZExtInst) (*ir.InstZExt, error) {
	i, ok := inst.(*ir.InstZExt)
	if !ok {
		// NOTE: panic since this would indicate a bug in the implementation.
		panic(fmt.Errorf("invalid IR instruction for AST instruction; expected *ir.InstZExt, got %T", inst))
	}
	// TODO: implement
	return i, nil
}

// ~~~ [ sext ] ~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~

func (fgen *funcGen) astToIRInstSExt(inst ir.Instruction, old *ast.SExtInst) (*ir.InstSExt, error) {
	i, ok := inst.(*ir.InstSExt)
	if !ok {
		// NOTE: panic since this would indicate a bug in the implementation.
		panic(fmt.Errorf("invalid IR instruction for AST instruction; expected *ir.InstSExt, got %T", inst))
	}
	// TODO: implement
	return i, nil
}

// ~~~ [ fptrunc ] ~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~

func (fgen *funcGen) astToIRInstFPTrunc(inst ir.Instruction, old *ast.FPTruncInst) (*ir.InstFPTrunc, error) {
	i, ok := inst.(*ir.InstFPTrunc)
	if !ok {
		// NOTE: panic since this would indicate a bug in the implementation.
		panic(fmt.Errorf("invalid IR instruction for AST instruction; expected *ir.InstFPTrunc, got %T", inst))
	}
	// TODO: implement
	return i, nil
}

// ~~~ [ fpext ] ~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~

func (fgen *funcGen) astToIRInstFPExt(inst ir.Instruction, old *ast.FPExtInst) (*ir.InstFPExt, error) {
	i, ok := inst.(*ir.InstFPExt)
	if !ok {
		// NOTE: panic since this would indicate a bug in the implementation.
		panic(fmt.Errorf("invalid IR instruction for AST instruction; expected *ir.InstFPExt, got %T", inst))
	}
	// TODO: implement
	return i, nil
}

// ~~~ [ fptoui ] ~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~

func (fgen *funcGen) astToIRInstFPToUI(inst ir.Instruction, old *ast.FPToUIInst) (*ir.InstFPToUI, error) {
	i, ok := inst.(*ir.InstFPToUI)
	if !ok {
		// NOTE: panic since this would indicate a bug in the implementation.
		panic(fmt.Errorf("invalid IR instruction for AST instruction; expected *ir.InstFPToUI, got %T", inst))
	}
	// TODO: implement
	return i, nil
}

// ~~~ [ fptosi ] ~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~

func (fgen *funcGen) astToIRInstFPToSI(inst ir.Instruction, old *ast.FPToSIInst) (*ir.InstFPToSI, error) {
	i, ok := inst.(*ir.InstFPToSI)
	if !ok {
		// NOTE: panic since this would indicate a bug in the implementation.
		panic(fmt.Errorf("invalid IR instruction for AST instruction; expected *ir.InstFPToSI, got %T", inst))
	}
	// TODO: implement
	return i, nil
}

// ~~~ [ uitofp ] ~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~

func (fgen *funcGen) astToIRInstUIToFP(inst ir.Instruction, old *ast.UIToFPInst) (*ir.InstUIToFP, error) {
	i, ok := inst.(*ir.InstUIToFP)
	if !ok {
		// NOTE: panic since this would indicate a bug in the implementation.
		panic(fmt.Errorf("invalid IR instruction for AST instruction; expected *ir.InstUIToFP, got %T", inst))
	}
	// TODO: implement
	return i, nil
}

// ~~~ [ sitofp ] ~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~

func (fgen *funcGen) astToIRInstSIToFP(inst ir.Instruction, old *ast.SIToFPInst) (*ir.InstSIToFP, error) {
	i, ok := inst.(*ir.InstSIToFP)
	if !ok {
		// NOTE: panic since this would indicate a bug in the implementation.
		panic(fmt.Errorf("invalid IR instruction for AST instruction; expected *ir.InstSIToFP, got %T", inst))
	}
	// TODO: implement
	return i, nil
}

// ~~~ [ ptrtoint ] ~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~

func (fgen *funcGen) astToIRInstPtrToInt(inst ir.Instruction, old *ast.PtrToIntInst) (*ir.InstPtrToInt, error) {
	i, ok := inst.(*ir.InstPtrToInt)
	if !ok {
		// NOTE: panic since this would indicate a bug in the implementation.
		panic(fmt.Errorf("invalid IR instruction for AST instruction; expected *ir.InstPtrToInt, got %T", inst))
	}
	// TODO: implement
	return i, nil
}

// ~~~ [ inttoptr ] ~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~

func (fgen *funcGen) astToIRInstIntToPtr(inst ir.Instruction, old *ast.IntToPtrInst) (*ir.InstIntToPtr, error) {
	i, ok := inst.(*ir.InstIntToPtr)
	if !ok {
		// NOTE: panic since this would indicate a bug in the implementation.
		panic(fmt.Errorf("invalid IR instruction for AST instruction; expected *ir.InstIntToPtr, got %T", inst))
	}
	// TODO: implement
	return i, nil
}

// ~~~ [ bitcast ] ~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~

func (fgen *funcGen) astToIRInstBitCast(inst ir.Instruction, old *ast.BitCastInst) (*ir.InstBitCast, error) {
	i, ok := inst.(*ir.InstBitCast)
	if !ok {
		// NOTE: panic since this would indicate a bug in the implementation.
		panic(fmt.Errorf("invalid IR instruction for AST instruction; expected *ir.InstBitCast, got %T", inst))
	}
	// TODO: implement
	return i, nil
}

// ~~~ [ addrspacecast ] ~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~

func (fgen *funcGen) astToIRInstAddrSpaceCast(inst ir.Instruction, old *ast.AddrSpaceCastInst) (*ir.InstAddrSpaceCast, error) {
	i, ok := inst.(*ir.InstAddrSpaceCast)
	if !ok {
		// NOTE: panic since this would indicate a bug in the implementation.
		panic(fmt.Errorf("invalid IR instruction for AST instruction; expected *ir.InstAddrSpaceCast, got %T", inst))
	}
	// TODO: implement
	return i, nil
}
