package asm

import (
	"fmt"

	"github.com/llir/ll/ast"
	"github.com/llir/llvm/ir"
	"github.com/pkg/errors"
)

// === [ Create IR ] ===========================================================

// newInst returns a new IR instruction (without body but with type) based on
// the given AST instruction.
func (fgen *funcGen) newInst(old ast.Instruction) (ir.Instruction, error) {
	switch old := old.(type) {
	// Value instructions.
	case *ast.LocalDefInst:
		ident := localIdent(old.Name())
		return fgen.newValueInst(ident, old.Inst())
	case ast.ValueInstruction:
		unnamed := ir.LocalIdent{}
		return fgen.newValueInst(unnamed, old)
	// Non-value instructions.
	case *ast.StoreInst:
		return &ir.InstStore{}, nil
	case *ast.FenceInst:
		return &ir.InstFence{}, nil
	default:
		panic(fmt.Errorf("support for AST instruction type %T not yet implemented", old))
	}
}

// newValueInst returns a new IR value instruction (without body but with type)
// based on the given AST value instruction.
func (fgen *funcGen) newValueInst(ident ir.LocalIdent, old ast.ValueInstruction) (ir.Instruction, error) {
	switch old := old.(type) {
	// Unary instructions
	case *ast.FNegInst:
		return fgen.newFNegInst(ident, old)
	// Binary instructions
	case *ast.AddInst:
		return fgen.newAddInst(ident, old)
	case *ast.FAddInst:
		return fgen.newFAddInst(ident, old)
	case *ast.SubInst:
		return fgen.newSubInst(ident, old)
	case *ast.FSubInst:
		return fgen.newFSubInst(ident, old)
	case *ast.MulInst:
		return fgen.newMulInst(ident, old)
	case *ast.FMulInst:
		return fgen.newFMulInst(ident, old)
	case *ast.UDivInst:
		return fgen.newUDivInst(ident, old)
	case *ast.SDivInst:
		return fgen.newSDivInst(ident, old)
	case *ast.FDivInst:
		return fgen.newFDivInst(ident, old)
	case *ast.URemInst:
		return fgen.newURemInst(ident, old)
	case *ast.SRemInst:
		return fgen.newSRemInst(ident, old)
	case *ast.FRemInst:
		return fgen.newFRemInst(ident, old)
	// Bitwise instructions
	case *ast.ShlInst:
		return fgen.newShlInst(ident, old)
	case *ast.LShrInst:
		return fgen.newLShrInst(ident, old)
	case *ast.AShrInst:
		return fgen.newAShrInst(ident, old)
	case *ast.AndInst:
		return fgen.newAndInst(ident, old)
	case *ast.OrInst:
		return fgen.newOrInst(ident, old)
	case *ast.XorInst:
		return fgen.newXorInst(ident, old)
	// Vector instructions
	case *ast.ExtractElementInst:
		return fgen.newExtractElementInst(ident, old)
	case *ast.InsertElementInst:
		return fgen.newInsertElementInst(ident, old)
	case *ast.ShuffleVectorInst:
		return fgen.newShuffleVectorInst(ident, old)
	// Aggregate instructions
	case *ast.ExtractValueInst:
		return fgen.newExtractValueInst(ident, old)
	case *ast.InsertValueInst:
		return fgen.newInsertValueInst(ident, old)
	// Memory instructions
	case *ast.AllocaInst:
		return fgen.newAllocaInst(ident, old)
	case *ast.LoadInst:
		return fgen.newLoadInst(ident, old)
	case *ast.CmpXchgInst:
		return fgen.newCmpXchgInst(ident, old)
	case *ast.AtomicRMWInst:
		return fgen.newAtomicRMWInst(ident, old)
	case *ast.GetElementPtrInst:
		return fgen.newGetElementPtrInst(ident, old)
	// Conversion instructions
	case *ast.TruncInst:
		return fgen.newTruncInst(ident, old)
	case *ast.ZExtInst:
		return fgen.newZExtInst(ident, old)
	case *ast.SExtInst:
		return fgen.newSExtInst(ident, old)
	case *ast.FPTruncInst:
		return fgen.newFPTruncInst(ident, old)
	case *ast.FPExtInst:
		return fgen.newFPExtInst(ident, old)
	case *ast.FPToUIInst:
		return fgen.newFPToUIInst(ident, old)
	case *ast.FPToSIInst:
		return fgen.newFPToSIInst(ident, old)
	case *ast.UIToFPInst:
		return fgen.newUIToFPInst(ident, old)
	case *ast.SIToFPInst:
		return fgen.newSIToFPInst(ident, old)
	case *ast.PtrToIntInst:
		return fgen.newPtrToIntInst(ident, old)
	case *ast.IntToPtrInst:
		return fgen.newIntToPtrInst(ident, old)
	case *ast.BitCastInst:
		return fgen.newBitCastInst(ident, old)
	case *ast.AddrSpaceCastInst:
		return fgen.newAddrSpaceCastInst(ident, old)
	// Other instructions
	case *ast.ICmpInst:
		return fgen.newICmpInst(ident, old)
	case *ast.FCmpInst:
		return fgen.newFCmpInst(ident, old)
	case *ast.PhiInst:
		return fgen.newPhiInst(ident, old)
	case *ast.SelectInst:
		return fgen.newSelectInst(ident, old)
	case *ast.CallInst:
		return fgen.newCallInst(ident, old)
	case *ast.VAArgInst:
		return fgen.newVAArgInst(ident, old)
	case *ast.LandingPadInst:
		return fgen.newLandingPadInst(ident, old)
	case *ast.CatchPadInst:
		// Result type is always token.
		return &ir.InstCatchPad{LocalIdent: ident}, nil
	case *ast.CleanupPadInst:
		// Result type is always token.
		return &ir.InstCleanupPad{LocalIdent: ident}, nil
	default:
		panic(fmt.Errorf("support for AST value instruction type %T not yet implemented", old))
	}
}

// === [ Translate AST to IR ] =================================================

// translateInsts translates the AST instructions of the given function to IR.
func (fgen *funcGen) translateInsts(oldBlocks []ast.BasicBlock) error {
	for i, oldBlock := range oldBlocks {
		block := fgen.f.Blocks[i]
		for j, old := range oldBlock.Insts() {
			new := block.Insts[j]
			if err := fgen.irInst(new, old); err != nil {
				return errors.WithStack(err)
			}
		}
	}
	return nil
}

// irInst translates the AST instruction into an equivalent IR instruction.
func (fgen *funcGen) irInst(new ir.Instruction, old ast.Instruction) error {
	switch old := old.(type) {
	// Value instructions.
	case *ast.LocalDefInst:
		return fgen.irValueInst(new, old.Inst())
	case ast.ValueInstruction:
		return fgen.irValueInst(new, old)
	// Non-value instructions.
	case *ast.StoreInst:
		return fgen.irStoreInst(new, old)
	case *ast.FenceInst:
		return fgen.irFenceInst(new, old)
	default:
		panic(fmt.Errorf("support for AST instruction type %T not yet implemented", old))
	}
}

// irValueInst translates the AST value instruction into an equivalent IR value
// instruction.
func (fgen *funcGen) irValueInst(new ir.Instruction, old ast.ValueInstruction) error {
	switch old := old.(type) {
	// Unary instructions
	case *ast.FNegInst:
		return fgen.irFNegInst(new, old)
	// Binary instructions
	case *ast.AddInst:
		return fgen.irAddInst(new, old)
	case *ast.FAddInst:
		return fgen.irFAddInst(new, old)
	case *ast.SubInst:
		return fgen.irSubInst(new, old)
	case *ast.FSubInst:
		return fgen.irFSubInst(new, old)
	case *ast.MulInst:
		return fgen.irMulInst(new, old)
	case *ast.FMulInst:
		return fgen.irFMulInst(new, old)
	case *ast.UDivInst:
		return fgen.irUDivInst(new, old)
	case *ast.SDivInst:
		return fgen.irSDivInst(new, old)
	case *ast.FDivInst:
		return fgen.irFDivInst(new, old)
	case *ast.URemInst:
		return fgen.irURemInst(new, old)
	case *ast.SRemInst:
		return fgen.irSRemInst(new, old)
	case *ast.FRemInst:
		return fgen.irFRemInst(new, old)
	// Bitwise instructions
	case *ast.ShlInst:
		return fgen.irShlInst(new, old)
	case *ast.LShrInst:
		return fgen.irLShrInst(new, old)
	case *ast.AShrInst:
		return fgen.irAShrInst(new, old)
	case *ast.AndInst:
		return fgen.irAndInst(new, old)
	case *ast.OrInst:
		return fgen.irOrInst(new, old)
	case *ast.XorInst:
		return fgen.irXorInst(new, old)
	// Vector instructions
	case *ast.ExtractElementInst:
		return fgen.irExtractElementInst(new, old)
	case *ast.InsertElementInst:
		return fgen.irInsertElementInst(new, old)
	case *ast.ShuffleVectorInst:
		return fgen.irShuffleVectorInst(new, old)
	// Aggregate instructions
	case *ast.ExtractValueInst:
		return fgen.irExtractValueInst(new, old)
	case *ast.InsertValueInst:
		return fgen.irInsertValueInst(new, old)
	// Memory instructions
	case *ast.AllocaInst:
		return fgen.irAllocaInst(new, old)
	case *ast.LoadInst:
		return fgen.irLoadInst(new, old)
	case *ast.CmpXchgInst:
		return fgen.irCmpXchgInst(new, old)
	case *ast.AtomicRMWInst:
		return fgen.irAtomicRMWInst(new, old)
	case *ast.GetElementPtrInst:
		return fgen.irGetElementPtrInst(new, old)
	// Conversion instructions
	case *ast.TruncInst:
		return fgen.irTruncInst(new, old)
	case *ast.ZExtInst:
		return fgen.irZExtInst(new, old)
	case *ast.SExtInst:
		return fgen.irSExtInst(new, old)
	case *ast.FPTruncInst:
		return fgen.irFPTruncInst(new, old)
	case *ast.FPExtInst:
		return fgen.irFPExtInst(new, old)
	case *ast.FPToUIInst:
		return fgen.irFPToUIInst(new, old)
	case *ast.FPToSIInst:
		return fgen.irFPToSIInst(new, old)
	case *ast.UIToFPInst:
		return fgen.irUIToFPInst(new, old)
	case *ast.SIToFPInst:
		return fgen.irSIToFPInst(new, old)
	case *ast.PtrToIntInst:
		return fgen.irPtrToIntInst(new, old)
	case *ast.IntToPtrInst:
		return fgen.irIntToPtrInst(new, old)
	case *ast.BitCastInst:
		return fgen.irBitCastInst(new, old)
	case *ast.AddrSpaceCastInst:
		return fgen.irAddrSpaceCastInst(new, old)
	// Other instructions
	case *ast.ICmpInst:
		return fgen.irICmpInst(new, old)
	case *ast.FCmpInst:
		return fgen.irFCmpInst(new, old)
	case *ast.PhiInst:
		return fgen.irPhiInst(new, old)
	case *ast.SelectInst:
		return fgen.irSelectInst(new, old)
	case *ast.CallInst:
		return fgen.irCallInst(new, old)
	case *ast.VAArgInst:
		return fgen.irVAArgInst(new, old)
	case *ast.LandingPadInst:
		return fgen.irLandingPadInst(new, old)
	case *ast.CatchPadInst:
		return fgen.irCatchPadInst(new, old)
	case *ast.CleanupPadInst:
		return fgen.irCleanupPadInst(new, old)
	default:
		panic(fmt.Errorf("support for AST value instruction type %T not yet implemented", old))
	}
}
