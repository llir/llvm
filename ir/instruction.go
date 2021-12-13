package ir

// === [ Instructions ] ========================================================

// Instruction is an LLVM IR instruction. All instructions (except store and
// fence) implement the value.Named interface and may thus be used directly as
// values.
//
// An Instruction has one of the following underlying types.
//
// Unary instructions
//
// https://llvm.org/docs/LangRef.html#unary-operations
//
//    *ir.InstFNeg   // https://pkg.go.dev/github.com/llir/llvm/ir#InstFNeg
//
// Binary instructions
//
// https://llvm.org/docs/LangRef.html#binary-operations
//
//    *ir.InstAdd    // https://pkg.go.dev/github.com/llir/llvm/ir#InstAdd
//    *ir.InstFAdd   // https://pkg.go.dev/github.com/llir/llvm/ir#InstFAdd
//    *ir.InstSub    // https://pkg.go.dev/github.com/llir/llvm/ir#InstSub
//    *ir.InstFSub   // https://pkg.go.dev/github.com/llir/llvm/ir#InstFSub
//    *ir.InstMul    // https://pkg.go.dev/github.com/llir/llvm/ir#InstMul
//    *ir.InstFMul   // https://pkg.go.dev/github.com/llir/llvm/ir#InstFMul
//    *ir.InstUDiv   // https://pkg.go.dev/github.com/llir/llvm/ir#InstUDiv
//    *ir.InstSDiv   // https://pkg.go.dev/github.com/llir/llvm/ir#InstSDiv
//    *ir.InstFDiv   // https://pkg.go.dev/github.com/llir/llvm/ir#InstFDiv
//    *ir.InstURem   // https://pkg.go.dev/github.com/llir/llvm/ir#InstURem
//    *ir.InstSRem   // https://pkg.go.dev/github.com/llir/llvm/ir#InstSRem
//    *ir.InstFRem   // https://pkg.go.dev/github.com/llir/llvm/ir#InstFRem
//
// Bitwise instructions
//
// https://llvm.org/docs/LangRef.html#bitwise-binary-operations
//
//    *ir.InstShl    // https://pkg.go.dev/github.com/llir/llvm/ir#InstShl
//    *ir.InstLShr   // https://pkg.go.dev/github.com/llir/llvm/ir#InstLShr
//    *ir.InstAShr   // https://pkg.go.dev/github.com/llir/llvm/ir#InstAShr
//    *ir.InstAnd    // https://pkg.go.dev/github.com/llir/llvm/ir#InstAnd
//    *ir.InstOr     // https://pkg.go.dev/github.com/llir/llvm/ir#InstOr
//    *ir.InstXor    // https://pkg.go.dev/github.com/llir/llvm/ir#InstXor
//
// Vector instructions
//
// https://llvm.org/docs/LangRef.html#vector-operations
//
//    *ir.InstExtractElement   // https://pkg.go.dev/github.com/llir/llvm/ir#InstExtractElement
//    *ir.InstInsertElement    // https://pkg.go.dev/github.com/llir/llvm/ir#InstInsertElement
//    *ir.InstShuffleVector    // https://pkg.go.dev/github.com/llir/llvm/ir#InstShuffleVector
//
// Aggregate instructions
//
// https://llvm.org/docs/LangRef.html#aggregate-operations
//
//    *ir.InstExtractValue   // https://pkg.go.dev/github.com/llir/llvm/ir#InstExtractValue
//    *ir.InstInsertValue    // https://pkg.go.dev/github.com/llir/llvm/ir#InstInsertValue
//
// Memory instructions
//
// https://llvm.org/docs/LangRef.html#memory-access-and-addressing-operations
//
//    *ir.InstAlloca          // https://pkg.go.dev/github.com/llir/llvm/ir#InstAlloca
//    *ir.InstLoad            // https://pkg.go.dev/github.com/llir/llvm/ir#InstLoad
//    *ir.InstStore           // https://pkg.go.dev/github.com/llir/llvm/ir#InstStore
//    *ir.InstFence           // https://pkg.go.dev/github.com/llir/llvm/ir#InstFence
//    *ir.InstCmpXchg         // https://pkg.go.dev/github.com/llir/llvm/ir#InstCmpXchg
//    *ir.InstAtomicRMW       // https://pkg.go.dev/github.com/llir/llvm/ir#InstAtomicRMW
//    *ir.InstGetElementPtr   // https://pkg.go.dev/github.com/llir/llvm/ir#InstGetElementPtr
//
// Conversion instructions
//
// https://llvm.org/docs/LangRef.html#conversion-operations
//
//    *ir.InstTrunc           // https://pkg.go.dev/github.com/llir/llvm/ir#InstTrunc
//    *ir.InstZExt            // https://pkg.go.dev/github.com/llir/llvm/ir#InstZExt
//    *ir.InstSExt            // https://pkg.go.dev/github.com/llir/llvm/ir#InstSExt
//    *ir.InstFPTrunc         // https://pkg.go.dev/github.com/llir/llvm/ir#InstFPTrunc
//    *ir.InstFPExt           // https://pkg.go.dev/github.com/llir/llvm/ir#InstFPExt
//    *ir.InstFPToUI          // https://pkg.go.dev/github.com/llir/llvm/ir#InstFPToUI
//    *ir.InstFPToSI          // https://pkg.go.dev/github.com/llir/llvm/ir#InstFPToSI
//    *ir.InstUIToFP          // https://pkg.go.dev/github.com/llir/llvm/ir#InstUIToFP
//    *ir.InstSIToFP          // https://pkg.go.dev/github.com/llir/llvm/ir#InstSIToFP
//    *ir.InstPtrToInt        // https://pkg.go.dev/github.com/llir/llvm/ir#InstPtrToInt
//    *ir.InstIntToPtr        // https://pkg.go.dev/github.com/llir/llvm/ir#InstIntToPtr
//    *ir.InstBitCast         // https://pkg.go.dev/github.com/llir/llvm/ir#InstBitCast
//    *ir.InstAddrSpaceCast   // https://pkg.go.dev/github.com/llir/llvm/ir#InstAddrSpaceCast
//
// Other instructions
//
// https://llvm.org/docs/LangRef.html#other-operations
//
//    *ir.InstICmp         // https://pkg.go.dev/github.com/llir/llvm/ir#InstICmp
//    *ir.InstFCmp         // https://pkg.go.dev/github.com/llir/llvm/ir#InstFCmp
//    *ir.InstPhi          // https://pkg.go.dev/github.com/llir/llvm/ir#InstPhi
//    *ir.InstSelect       // https://pkg.go.dev/github.com/llir/llvm/ir#InstSelect
//    *ir.InstFreeze       // https://pkg.go.dev/github.com/llir/llvm/ir#InstFreeze
//    *ir.InstCall         // https://pkg.go.dev/github.com/llir/llvm/ir#InstCall
//    *ir.InstVAArg        // https://pkg.go.dev/github.com/llir/llvm/ir#InstVAArg
//    *ir.InstLandingPad   // https://pkg.go.dev/github.com/llir/llvm/ir#InstLandingPad
//    *ir.InstCatchPad     // https://pkg.go.dev/github.com/llir/llvm/ir#InstCatchPad
//    *ir.InstCleanupPad   // https://pkg.go.dev/github.com/llir/llvm/ir#InstCleanupPad
type Instruction interface {
	LLStringer
	// isInstruction ensures that only instructions can be assigned to the
	// instruction.Instruction interface.
	isInstruction()
}
