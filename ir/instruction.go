package ir

import "fmt"

// An Instruction represents a non-branching LLVM IR instruction.
//
// Instructions which produce results may be referenced from other instructions,
// and are thus considered LLVM IR values. Note, not all instructions produce
// results (e.g. store).
//
// Instruction may have one of the following underlying types.
//
// Binary instructions
//
// http://llvm.org/docs/LangRef.html#binary-operations
//
//    *ir.InstAdd
//    *ir.InstFAdd
//    *ir.InstSub
//    *ir.InstFSub
//    *ir.InstMul
//    *ir.InstFMul
//    *ir.InstUDiv
//    *ir.InstSDiv
//    *ir.InstFDiv
//    *ir.InstURem
//    *ir.InstSRem
//    *ir.InstFRem
//
// Bitwise instructions
//
// http://llvm.org/docs/LangRef.html#bitwise-binary-operations
//
//    *ir.InstShl
//    *ir.InstLShr
//    *ir.InstAShr
//    *ir.InstAnd
//    *ir.InstOr
//    *ir.InstXor
//
// Memory instructions
//
// http://llvm.org/docs/LangRef.html#memory-access-and-addressing-operations
//
//    *ir.InstAlloca
//    *ir.InstLoad
//    *ir.InstStore
//    *ir.InstGetElementPtr
//
// Conversion instructions
//
// http://llvm.org/docs/LangRef.html#conversion-operations
//
//    *ir.InstTrunc
//    *ir.InstZExt
//    *ir.InstSExt
//    *ir.InstFPTrunc
//    *ir.InstFPExt
//    *ir.InstFPToUI
//    *ir.InstFPToSI
//    *ir.InstUIToFP
//    *ir.InstSIToFP
//    *ir.InstPtrToInt
//    *ir.InstIntToPtr
//    *ir.InstBitCast
//    *ir.InstAddrSpaceCast
//
// Other instructions
//
// http://llvm.org/docs/LangRef.html#other-operations
//
//    *ir.InstICmp
//    *ir.InstFCmp
//    *ir.InstPhi
//    *ir.InstSelect
//    *ir.InstCall
type Instruction interface {
	fmt.Stringer
	// Parent returns the parent basic block of the instruction.
	Parent() *BasicBlock
}
