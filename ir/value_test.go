package ir

import "github.com/llir/l/ir/value"

// Assert that each value implements the value.Value interface.
var (
	// Binary instructions.
	_ value.Value = (*InstAdd)(nil)
	_ value.Value = (*InstFAdd)(nil)
	_ value.Value = (*InstSub)(nil)
	_ value.Value = (*InstFSub)(nil)
	_ value.Value = (*InstMul)(nil)
	_ value.Value = (*InstFMul)(nil)
	_ value.Value = (*InstUDiv)(nil)
	_ value.Value = (*InstSDiv)(nil)
	_ value.Value = (*InstFDiv)(nil)
	_ value.Value = (*InstURem)(nil)
	_ value.Value = (*InstSRem)(nil)
	_ value.Value = (*InstFRem)(nil)
	// Bitwise instructions.
	_ value.Value = (*InstShl)(nil)
	_ value.Value = (*InstLShr)(nil)
	_ value.Value = (*InstAShr)(nil)
	_ value.Value = (*InstAnd)(nil)
	_ value.Value = (*InstOr)(nil)
	_ value.Value = (*InstXor)(nil)
	// Vector instructions.
	_ value.Value = (*InstExtractElement)(nil)
	_ value.Value = (*InstInsertElement)(nil)
	_ value.Value = (*InstShuffleVector)(nil)
	// Aggregate instructions.
	_ value.Value = (*InstExtractValue)(nil)
	_ value.Value = (*InstInsertValue)(nil)
	// Memory instructions.
	_ value.Value = (*InstAlloca)(nil)
	_ value.Value = (*InstLoad)(nil)
	//_ value.Value = (*InstStore)(nil)
	//_ value.Value = (*InstFence)(nil)
	_ value.Value = (*InstCmpXchg)(nil)
	_ value.Value = (*InstAtomicRMW)(nil)
	_ value.Value = (*InstGetElementPtr)(nil)
	// Conversion instructions.
	_ value.Value = (*InstTrunc)(nil)
	_ value.Value = (*InstZExt)(nil)
	_ value.Value = (*InstSExt)(nil)
	_ value.Value = (*InstFPTrunc)(nil)
	_ value.Value = (*InstFPExt)(nil)
	_ value.Value = (*InstFPToUI)(nil)
	_ value.Value = (*InstFPToSI)(nil)
	_ value.Value = (*InstUIToFP)(nil)
	_ value.Value = (*InstSIToFP)(nil)
	_ value.Value = (*InstPtrToInt)(nil)
	_ value.Value = (*InstIntToPtr)(nil)
	_ value.Value = (*InstBitCast)(nil)
	_ value.Value = (*InstAddrSpaceCast)(nil)
	// Other instructions.
	_ value.Value = (*InstICmp)(nil)
	_ value.Value = (*InstFCmp)(nil)
	_ value.Value = (*InstPhi)(nil)
	_ value.Value = (*InstSelect)(nil)
	_ value.Value = (*InstCall)(nil)
	_ value.Value = (*InstVAArg)(nil)
	_ value.Value = (*InstLandingPad)(nil)
	_ value.Value = (*InstCatchPad)(nil)
	_ value.Value = (*InstCleanupPad)(nil)
	// Terminators.
	_ value.Value = (*TermInvoke)(nil)
	_ value.Value = (*TermCatchSwitch)(nil)
)
