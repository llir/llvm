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
//    *ir.InstAdd    // https://godoc.org/github.com/llir/l/ir#InstAdd
//    *ir.InstFAdd   // https://godoc.org/github.com/llir/l/ir#InstFAdd
//    *ir.InstSub    // https://godoc.org/github.com/llir/l/ir#InstSub
//    *ir.InstFSub   // https://godoc.org/github.com/llir/l/ir#InstFSub
//    *ir.InstMul    // https://godoc.org/github.com/llir/l/ir#InstMul
//    *ir.InstFMul   // https://godoc.org/github.com/llir/l/ir#InstFMul
//    *ir.InstUDiv   // https://godoc.org/github.com/llir/l/ir#InstUDiv
//    *ir.InstSDiv   // https://godoc.org/github.com/llir/l/ir#InstSDiv
//    *ir.InstFDiv   // https://godoc.org/github.com/llir/l/ir#InstFDiv
//    *ir.InstURem   // https://godoc.org/github.com/llir/l/ir#InstURem
//    *ir.InstSRem   // https://godoc.org/github.com/llir/l/ir#InstSRem
//    *ir.InstFRem   // https://godoc.org/github.com/llir/l/ir#InstFRem
//
// Bitwise instructions
//
// https://llvm.org/docs/LangRef.html#bitwise-binary-operations
//
//    *ir.InstShl    // https://godoc.org/github.com/llir/l/ir#InstShl
//    *ir.InstLShr   // https://godoc.org/github.com/llir/l/ir#InstLShr
//    *ir.InstAShr   // https://godoc.org/github.com/llir/l/ir#InstAShr
//    *ir.InstAnd    // https://godoc.org/github.com/llir/l/ir#InstAnd
//    *ir.InstOr     // https://godoc.org/github.com/llir/l/ir#InstOr
//    *ir.InstXor    // https://godoc.org/github.com/llir/l/ir#InstXor
//
// Vector instructions
//
// https://llvm.org/docs/LangRef.html#vector-operations
//
//    *ir.InstExtractElement   // https://godoc.org/github.com/llir/l/ir#InstExtractElement
//    *ir.InstInsertElement    // https://godoc.org/github.com/llir/l/ir#InstInsertElement
//    *ir.InstShuffleVector    // https://godoc.org/github.com/llir/l/ir#InstShuffleVector
//
// Aggregate instructions
//
// https://llvm.org/docs/LangRef.html#aggregate-operations
//
//    *ir.InstExtractValue   // https://godoc.org/github.com/llir/l/ir#InstExtractValue
//    *ir.InstInsertValue    // https://godoc.org/github.com/llir/l/ir#InstInsertValue
//
// Memory instructions
//
// https://llvm.org/docs/LangRef.html#memory-access-and-addressing-operations
//
//    *ir.InstAlloca          // https://godoc.org/github.com/llir/l/ir#InstAlloca
//    *ir.InstLoad            // https://godoc.org/github.com/llir/l/ir#InstLoad
//    *ir.InstStore           // https://godoc.org/github.com/llir/l/ir#InstStore
//    *ir.InstFence           // https://godoc.org/github.com/llir/l/ir#InstFence
//    *ir.InstCmpXchg         // https://godoc.org/github.com/llir/l/ir#InstCmpXchg
//    *ir.InstAtomicRMW       // https://godoc.org/github.com/llir/l/ir#InstAtomicRMW
//    *ir.InstGetElementPtr   // https://godoc.org/github.com/llir/l/ir#InstGetElementPtr
//
// Conversion instructions
//
// https://llvm.org/docs/LangRef.html#conversion-operations
//
//    *ir.InstTrunc           // https://godoc.org/github.com/llir/l/ir#InstTrunc
//    *ir.InstZExt            // https://godoc.org/github.com/llir/l/ir#InstZExt
//    *ir.InstSExt            // https://godoc.org/github.com/llir/l/ir#InstSExt
//    *ir.InstFPTrunc         // https://godoc.org/github.com/llir/l/ir#InstFPTrunc
//    *ir.InstFPExt           // https://godoc.org/github.com/llir/l/ir#InstFPExt
//    *ir.InstFPToUI          // https://godoc.org/github.com/llir/l/ir#InstFPToUI
//    *ir.InstFPToSI          // https://godoc.org/github.com/llir/l/ir#InstFPToSI
//    *ir.InstUIToFP          // https://godoc.org/github.com/llir/l/ir#InstUIToFP
//    *ir.InstSIToFP          // https://godoc.org/github.com/llir/l/ir#InstSIToFP
//    *ir.InstPtrToInt        // https://godoc.org/github.com/llir/l/ir#InstPtrToInt
//    *ir.InstIntToPtr        // https://godoc.org/github.com/llir/l/ir#InstIntToPtr
//    *ir.InstBitCast         // https://godoc.org/github.com/llir/l/ir#InstBitCast
//    *ir.InstAddrSpaceCast   // https://godoc.org/github.com/llir/l/ir#InstAddrSpaceCast
//
// Other instructions
//
// https://llvm.org/docs/LangRef.html#other-operations
//
//    *ir.InstICmp         // https://godoc.org/github.com/llir/l/ir#InstICmp
//    *ir.InstFCmp         // https://godoc.org/github.com/llir/l/ir#InstFCmp
//    *ir.InstPhi          // https://godoc.org/github.com/llir/l/ir#InstPhi
//    *ir.InstSelect       // https://godoc.org/github.com/llir/l/ir#InstSelect
//    *ir.InstCall         // https://godoc.org/github.com/llir/l/ir#InstCall
//    *ir.InstVAArg        // https://godoc.org/github.com/llir/l/ir#InstVAArg
//    *ir.InstLandingPad   // https://godoc.org/github.com/llir/l/ir#InstLandingPad
//    *ir.InstCatchPad     // https://godoc.org/github.com/llir/l/ir#InstCatchPad
//    *ir.InstCleanupPad   // https://godoc.org/github.com/llir/l/ir#InstCleanupPad
type Instruction interface {
	// isInstruction ensures that only instructions can be assigned to the
	// instruction.Instruction interface.
	isInstruction()
}

// Binary instructions.
func (*InstAdd) isInstruction()  {}
func (*InstFAdd) isInstruction() {}
func (*InstSub) isInstruction()  {}
func (*InstFSub) isInstruction() {}
func (*InstMul) isInstruction()  {}
func (*InstFMul) isInstruction() {}
func (*InstUDiv) isInstruction() {}
func (*InstSDiv) isInstruction() {}
func (*InstFDiv) isInstruction() {}
func (*InstURem) isInstruction() {}
func (*InstSRem) isInstruction() {}
func (*InstFRem) isInstruction() {}

// Bitwise instructions.
func (*InstShl) isInstruction()  {}
func (*InstLShr) isInstruction() {}
func (*InstAShr) isInstruction() {}
func (*InstAnd) isInstruction()  {}
func (*InstOr) isInstruction()   {}
func (*InstXor) isInstruction()  {}

// Vector instructions.
func (*InstExtractElement) isInstruction() {}
func (*InstInsertElement) isInstruction()  {}
func (*InstShuffleVector) isInstruction()  {}

// Aggregate instructions.
func (*InstExtractValue) isInstruction() {}
func (*InstInsertValue) isInstruction()  {}

// Memory instructions.
func (*InstAlloca) isInstruction()        {}
func (*InstLoad) isInstruction()          {}
func (*InstStore) isInstruction()         {}
func (*InstFence) isInstruction()         {}
func (*InstCmpXchg) isInstruction()       {}
func (*InstAtomicRMW) isInstruction()     {}
func (*InstGetElementPtr) isInstruction() {}

// Conversion instructions.
func (*InstTrunc) isInstruction()         {}
func (*InstZExt) isInstruction()          {}
func (*InstSExt) isInstruction()          {}
func (*InstFPTrunc) isInstruction()       {}
func (*InstFPExt) isInstruction()         {}
func (*InstFPToUI) isInstruction()        {}
func (*InstFPToSI) isInstruction()        {}
func (*InstUIToFP) isInstruction()        {}
func (*InstSIToFP) isInstruction()        {}
func (*InstPtrToInt) isInstruction()      {}
func (*InstIntToPtr) isInstruction()      {}
func (*InstBitCast) isInstruction()       {}
func (*InstAddrSpaceCast) isInstruction() {}

// Other instructions.
func (*InstICmp) isInstruction()       {}
func (*InstFCmp) isInstruction()       {}
func (*InstPhi) isInstruction()        {}
func (*InstSelect) isInstruction()     {}
func (*InstCall) isInstruction()       {}
func (*InstVAArg) isInstruction()      {}
func (*InstLandingPad) isInstruction() {}
func (*InstCatchPad) isInstruction()   {}
func (*InstCleanupPad) isInstruction() {}
