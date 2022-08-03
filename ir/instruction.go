package ir

import "github.com/llir/llvm/ir/value"

// === [ Instructions ] ========================================================

// Instruction is an LLVM IR instruction. All instructions (except store and
// fence) implement the value.Named interface and may thus be used directly as
// values.
//
// An Instruction has one of the following underlying types.
//
// # Unary instructions
//
// https://llvm.org/docs/LangRef.html#unary-operations
//
//   - [*ir.InstFNeg]
//
// # Binary instructions
//
// https://llvm.org/docs/LangRef.html#binary-operations
//
//   - [*ir.InstAdd]
//   - [*ir.InstFAdd]
//   - [*ir.InstSub]
//   - [*ir.InstFSub]
//   - [*ir.InstMul]
//   - [*ir.InstFMul]
//   - [*ir.InstUDiv]
//   - [*ir.InstSDiv]
//   - [*ir.InstFDiv]
//   - [*ir.InstURem]
//   - [*ir.InstSRem]
//   - [*ir.InstFRem]
//
// # Bitwise instructions
//
// https://llvm.org/docs/LangRef.html#bitwise-binary-operations
//
//   - [*ir.InstShl]
//   - [*ir.InstLShr]
//   - [*ir.InstAShr]
//   - [*ir.InstAnd]
//   - [*ir.InstOr]
//   - [*ir.InstXor]
//
// # Vector instructions
//
// https://llvm.org/docs/LangRef.html#vector-operations
//
//   - [*ir.InstExtractElement]
//   - [*ir.InstInsertElement]
//   - [*ir.InstShuffleVector]
//
// # Aggregate instructions
//
// https://llvm.org/docs/LangRef.html#aggregate-operations
//
//   - [*ir.InstExtractValue]
//   - [*ir.InstInsertValue]
//
// # Memory instructions
//
// https://llvm.org/docs/LangRef.html#memory-access-and-addressing-operations
//
//   - [*ir.InstAlloca]
//   - [*ir.InstLoad]
//   - [*ir.InstStore]
//   - [*ir.InstFence]
//   - [*ir.InstCmpXchg]
//   - [*ir.InstAtomicRMW]
//   - [*ir.InstGetElementPtr]
//
// # Conversion instructions
//
// https://llvm.org/docs/LangRef.html#conversion-operations
//
//   - [*ir.InstTrunc]
//   - [*ir.InstZExt]
//   - [*ir.InstSExt]
//   - [*ir.InstFPTrunc]
//   - [*ir.InstFPExt]
//   - [*ir.InstFPToUI]
//   - [*ir.InstFPToSI]
//   - [*ir.InstUIToFP]
//   - [*ir.InstSIToFP]
//   - [*ir.InstPtrToInt]
//   - [*ir.InstIntToPtr]
//   - [*ir.InstBitCast]
//   - [*ir.InstAddrSpaceCast]
//
// # Other instructions
//
// https://llvm.org/docs/LangRef.html#other-operations
//
//   - [*ir.InstICmp]
//   - [*ir.InstFCmp]
//   - [*ir.InstPhi]
//   - [*ir.InstSelect]
//   - [*ir.InstFreeze]
//   - [*ir.InstCall]
//   - [*ir.InstVAArg]
//   - [*ir.InstLandingPad]
//   - [*ir.InstCatchPad]
//   - [*ir.InstCleanupPad]
type Instruction interface {
	LLStringer
	// isInstruction ensures that only instructions can be assigned to the
	// instruction.Instruction interface.
	isInstruction()
	// Instruction implements the value.User interface.
	value.User
}
