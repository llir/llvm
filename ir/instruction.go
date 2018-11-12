package ir

// === [ Instructions ] ========================================================

// Instruction is an LLVM IR instruction. All instructions (except store and
// fence) implement the value.Named interface and may thus be used directly as
// values.
//
// An Instruction has one of the following underlying types.
//
// Binary instructions
//
// https://llvm.org/docs/LangRef.html#binary-operations
//
//    *ir.InstAdd    // https://godoc.org/github.com/llir/llvm/ir#InstAdd
//    *ir.InstFAdd   // https://godoc.org/github.com/llir/llvm/ir#InstFAdd
//    *ir.InstSub    // https://godoc.org/github.com/llir/llvm/ir#InstSub
//    *ir.InstFSub   // https://godoc.org/github.com/llir/llvm/ir#InstFSub
//    *ir.InstMul    // https://godoc.org/github.com/llir/llvm/ir#InstMul
//    *ir.InstFMul   // https://godoc.org/github.com/llir/llvm/ir#InstFMul
//    *ir.InstUDiv   // https://godoc.org/github.com/llir/llvm/ir#InstUDiv
//    *ir.InstSDiv   // https://godoc.org/github.com/llir/llvm/ir#InstSDiv
//    *ir.InstFDiv   // https://godoc.org/github.com/llir/llvm/ir#InstFDiv
//    *ir.InstURem   // https://godoc.org/github.com/llir/llvm/ir#InstURem
//    *ir.InstSRem   // https://godoc.org/github.com/llir/llvm/ir#InstSRem
//    *ir.InstFRem   // https://godoc.org/github.com/llir/llvm/ir#InstFRem
//
// Bitwise instructions
//
// https://llvm.org/docs/LangRef.html#bitwise-binary-operations
//
//    *ir.InstShl    // https://godoc.org/github.com/llir/llvm/ir#InstShl
//    *ir.InstLShr   // https://godoc.org/github.com/llir/llvm/ir#InstLShr
//    *ir.InstAShr   // https://godoc.org/github.com/llir/llvm/ir#InstAShr
//    *ir.InstAnd    // https://godoc.org/github.com/llir/llvm/ir#InstAnd
//    *ir.InstOr     // https://godoc.org/github.com/llir/llvm/ir#InstOr
//    *ir.InstXor    // https://godoc.org/github.com/llir/llvm/ir#InstXor
//
// Vector instructions
//
// https://llvm.org/docs/LangRef.html#vector-operations
//
//    *ir.InstExtractElement   // https://godoc.org/github.com/llir/llvm/ir#InstExtractElement
//    *ir.InstInsertElement    // https://godoc.org/github.com/llir/llvm/ir#InstInsertElement
//    *ir.InstShuffleVector    // https://godoc.org/github.com/llir/llvm/ir#InstShuffleVector
//
// Aggregate instructions
//
// https://llvm.org/docs/LangRef.html#aggregate-operations
//
//    *ir.InstExtractValue   // https://godoc.org/github.com/llir/llvm/ir#InstExtractValue
//    *ir.InstInsertValue    // https://godoc.org/github.com/llir/llvm/ir#InstInsertValue
//
// Memory instructions
//
// https://llvm.org/docs/LangRef.html#memory-access-and-addressing-operations
//
//    *ir.InstAlloca          // https://godoc.org/github.com/llir/llvm/ir#InstAlloca
//    *ir.InstLoad            // https://godoc.org/github.com/llir/llvm/ir#InstLoad
//    *ir.InstStore           // https://godoc.org/github.com/llir/llvm/ir#InstStore
//    *ir.InstFence           // https://godoc.org/github.com/llir/llvm/ir#InstFence
//    *ir.InstCmpXchg         // https://godoc.org/github.com/llir/llvm/ir#InstCmpXchg
//    *ir.InstAtomicRMW       // https://godoc.org/github.com/llir/llvm/ir#InstAtomicRMW
//    *ir.InstGetElementPtr   // https://godoc.org/github.com/llir/llvm/ir#InstGetElementPtr
//
// Conversion instructions
//
// https://llvm.org/docs/LangRef.html#conversion-operations
//
//    *ir.InstTrunc           // https://godoc.org/github.com/llir/llvm/ir#InstTrunc
//    *ir.InstZExt            // https://godoc.org/github.com/llir/llvm/ir#InstZExt
//    *ir.InstSExt            // https://godoc.org/github.com/llir/llvm/ir#InstSExt
//    *ir.InstFPTrunc         // https://godoc.org/github.com/llir/llvm/ir#InstFPTrunc
//    *ir.InstFPExt           // https://godoc.org/github.com/llir/llvm/ir#InstFPExt
//    *ir.InstFPToUI          // https://godoc.org/github.com/llir/llvm/ir#InstFPToUI
//    *ir.InstFPToSI          // https://godoc.org/github.com/llir/llvm/ir#InstFPToSI
//    *ir.InstUIToFP          // https://godoc.org/github.com/llir/llvm/ir#InstUIToFP
//    *ir.InstSIToFP          // https://godoc.org/github.com/llir/llvm/ir#InstSIToFP
//    *ir.InstPtrToInt        // https://godoc.org/github.com/llir/llvm/ir#InstPtrToInt
//    *ir.InstIntToPtr        // https://godoc.org/github.com/llir/llvm/ir#InstIntToPtr
//    *ir.InstBitCast         // https://godoc.org/github.com/llir/llvm/ir#InstBitCast
//    *ir.InstAddrSpaceCast   // https://godoc.org/github.com/llir/llvm/ir#InstAddrSpaceCast
//
// Other instructions
//
// https://llvm.org/docs/LangRef.html#other-operations
//
//    *ir.InstICmp         // https://godoc.org/github.com/llir/llvm/ir#InstICmp
//    *ir.InstFCmp         // https://godoc.org/github.com/llir/llvm/ir#InstFCmp
//    *ir.InstPhi          // https://godoc.org/github.com/llir/llvm/ir#InstPhi
//    *ir.InstSelect       // https://godoc.org/github.com/llir/llvm/ir#InstSelect
//    *ir.InstCall         // https://godoc.org/github.com/llir/llvm/ir#InstCall
//    *ir.InstVAArg        // https://godoc.org/github.com/llir/llvm/ir#InstVAArg
//    *ir.InstLandingPad   // https://godoc.org/github.com/llir/llvm/ir#InstLandingPad
//    *ir.InstCatchPad     // https://godoc.org/github.com/llir/llvm/ir#InstCatchPad
//    *ir.InstCleanupPad   // https://godoc.org/github.com/llir/llvm/ir#InstCleanupPad
type Instruction interface {
	// Def returns the LLVM syntax representation of the instruction.
	Def() string
	// isInstruction ensures that only instructions can be assigned to the
	// instruction.Instruction interface.
	isInstruction()
}
