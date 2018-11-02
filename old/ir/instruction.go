// === [ Instructions ] ========================================================
//
// References:
//    http://llvm.org/docs/LangRef.html#instruction-reference

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
//    *ir.InstAdd    (https://godoc.org/github.com/llir/llvm/ir#InstAdd)
//    *ir.InstFAdd   (https://godoc.org/github.com/llir/llvm/ir#InstFAdd)
//    *ir.InstSub    (https://godoc.org/github.com/llir/llvm/ir#InstSub)
//    *ir.InstFSub   (https://godoc.org/github.com/llir/llvm/ir#InstFSub)
//    *ir.InstMul    (https://godoc.org/github.com/llir/llvm/ir#InstMul)
//    *ir.InstFMul   (https://godoc.org/github.com/llir/llvm/ir#InstFMul)
//    *ir.InstUDiv   (https://godoc.org/github.com/llir/llvm/ir#InstUDiv)
//    *ir.InstSDiv   (https://godoc.org/github.com/llir/llvm/ir#InstSDiv)
//    *ir.InstFDiv   (https://godoc.org/github.com/llir/llvm/ir#InstFDiv)
//    *ir.InstURem   (https://godoc.org/github.com/llir/llvm/ir#InstURem)
//    *ir.InstSRem   (https://godoc.org/github.com/llir/llvm/ir#InstSRem)
//    *ir.InstFRem   (https://godoc.org/github.com/llir/llvm/ir#InstFRem)
//
// Bitwise instructions
//
// http://llvm.org/docs/LangRef.html#bitwise-binary-operations
//
//    *ir.InstShl    (https://godoc.org/github.com/llir/llvm/ir#InstShl)
//    *ir.InstLShr   (https://godoc.org/github.com/llir/llvm/ir#InstLShr)
//    *ir.InstAShr   (https://godoc.org/github.com/llir/llvm/ir#InstAShr)
//    *ir.InstAnd    (https://godoc.org/github.com/llir/llvm/ir#InstAnd)
//    *ir.InstOr     (https://godoc.org/github.com/llir/llvm/ir#InstOr)
//    *ir.InstXor    (https://godoc.org/github.com/llir/llvm/ir#InstXor)
//
// Vector instructions
//
// http://llvm.org/docs/LangRef.html#vector-operations
//
//    *ir.InstExtractElement   (https://godoc.org/github.com/llir/llvm/ir#InstExtractElement)
//    *ir.InstInsertElement    (https://godoc.org/github.com/llir/llvm/ir#InstInsertElement)
//    *ir.InstShuffleVector    (https://godoc.org/github.com/llir/llvm/ir#InstShuffleVector)
//
// Aggregate instructions
//
// http://llvm.org/docs/LangRef.html#aggregate-operations
//
//    *ir.InstExtractValue   (https://godoc.org/github.com/llir/llvm/ir#InstExtractValue)
//    *ir.InstInsertValue    (https://godoc.org/github.com/llir/llvm/ir#InstInsertValue)
//
// Memory instructions
//
// http://llvm.org/docs/LangRef.html#memory-access-and-addressing-operations
//
//    *ir.InstAlloca          (https://godoc.org/github.com/llir/llvm/ir#InstAlloca)
//    *ir.InstLoad            (https://godoc.org/github.com/llir/llvm/ir#InstLoad)
//    *ir.InstStore           (https://godoc.org/github.com/llir/llvm/ir#InstStore)
//    *ir.InstGetElementPtr   (https://godoc.org/github.com/llir/llvm/ir#InstGetElementPtr)
//
// Conversion instructions
//
// http://llvm.org/docs/LangRef.html#conversion-operations
//
//    *ir.InstTrunc           (https://godoc.org/github.com/llir/llvm/ir#InstTrunc)
//    *ir.InstZExt            (https://godoc.org/github.com/llir/llvm/ir#InstZExt)
//    *ir.InstSExt            (https://godoc.org/github.com/llir/llvm/ir#InstSExt)
//    *ir.InstFPTrunc         (https://godoc.org/github.com/llir/llvm/ir#InstFPTrunc)
//    *ir.InstFPExt           (https://godoc.org/github.com/llir/llvm/ir#InstFPExt)
//    *ir.InstFPToUI          (https://godoc.org/github.com/llir/llvm/ir#InstFPToUI)
//    *ir.InstFPToSI          (https://godoc.org/github.com/llir/llvm/ir#InstFPToSI)
//    *ir.InstUIToFP          (https://godoc.org/github.com/llir/llvm/ir#InstUIToFP)
//    *ir.InstSIToFP          (https://godoc.org/github.com/llir/llvm/ir#InstSIToFP)
//    *ir.InstPtrToInt        (https://godoc.org/github.com/llir/llvm/ir#InstPtrToInt)
//    *ir.InstIntToPtr        (https://godoc.org/github.com/llir/llvm/ir#InstIntToPtr)
//    *ir.InstBitCast         (https://godoc.org/github.com/llir/llvm/ir#InstBitCast)
//    *ir.InstAddrSpaceCast   (https://godoc.org/github.com/llir/llvm/ir#InstAddrSpaceCast)
//
// Other instructions
//
// http://llvm.org/docs/LangRef.html#other-operations
//
//    *ir.InstICmp     (https://godoc.org/github.com/llir/llvm/ir#InstICmp)
//    *ir.InstFCmp     (https://godoc.org/github.com/llir/llvm/ir#InstFCmp)
//    *ir.InstPhi      (https://godoc.org/github.com/llir/llvm/ir#InstPhi)
//    *ir.InstSelect   (https://godoc.org/github.com/llir/llvm/ir#InstSelect)
//    *ir.InstCall     (https://godoc.org/github.com/llir/llvm/ir#InstCall)
type Instruction interface {
	fmt.Stringer
	// GetParent returns the parent basic block of the instruction.
	GetParent() *BasicBlock
	// SetParent sets the parent basic block of the instruction.
	SetParent(parent *BasicBlock)
}
