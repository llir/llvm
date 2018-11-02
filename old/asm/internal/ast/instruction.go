package ast

// An Instruction represents a non-branching LLVM IR instruction.
//
// Instruction may have one of the following underlying types.
//
// Binary instructions
//
// http://llvm.org/docs/LangRef.html#binary-operations
//
//    *ast.InstAdd
//    *ast.InstFAdd
//    *ast.InstSub
//    *ast.InstFSub
//    *ast.InstMul
//    *ast.InstFMul
//    *ast.InstUDiv
//    *ast.InstSDiv
//    *ast.InstFDiv
//    *ast.InstURem
//    *ast.InstSRem
//    *ast.InstFRem
//
// Bitwise instructions
//
// http://llvm.org/docs/LangRef.html#bitwise-binary-operations
//
//    *ast.InstShl
//    *ast.InstLShr
//    *ast.InstAShr
//    *ast.InstAnd
//    *ast.InstOr
//    *ast.InstXor
//
// Vector instructions
//
// http://llvm.org/docs/LangRef.html#vector-operations
//
//    *ast.InstExtractElement
//    *ast.InstInsertElement
//    *ast.InstShuffleVector
//
// Aggregate instructions
//
// http://llvm.org/docs/LangRef.html#aggregate-operations
//
//    *ast.InstExtractValue
//    *ast.InstInsertValue
//
// Memory instructions
//
// http://llvm.org/docs/LangRef.html#memory-access-and-addressing-operations
//
//    *ast.InstAlloca
//    *ast.InstLoad
//    *ast.InstStore
//    *ast.InstGetElementPtr
//
// Conversion instructions
//
// http://llvm.org/docs/LangRef.html#conversion-operations
//
//    *ast.InstTrunc
//    *ast.InstZExt
//    *ast.InstSExt
//    *ast.InstFPTrunc
//    *ast.InstFPExt
//    *ast.InstFPToUI
//    *ast.InstFPToSI
//    *ast.InstUIToFP
//    *ast.InstSIToFP
//    *ast.InstPtrToInt
//    *ast.InstIntToPtr
//    *ast.InstBitCast
//    *ast.InstAddrSpaceCast
//
// Other instructions
//
// http://llvm.org/docs/LangRef.html#other-operations
//
//    *ast.InstICmp
//    *ast.InstFCmp
//    *ast.InstPhi
//    *ast.InstSelect
//    *ast.InstCall
type Instruction interface {
	// isInst ensures that only instructions can be assigned to the
	// ast.Instruction interface.
	isInst()
}
