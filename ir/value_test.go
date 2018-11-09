package ir

import (
	"github.com/llir/llvm/ir/constant"
	"github.com/llir/llvm/ir/metadata"
	"github.com/llir/llvm/ir/value"
)

// Assert that each value implements the value.Value interface.
var (
	// Constants.
	// Checked in constant_test.go as ir.Constant embeds value.Value.
	_ value.Value = constant.Constant(nil)
	// Named values.
	// Checked in value_test.go as value.Named embeds value.Value.
	_ value.Value = value.Named(nil)
	// Inline assembler expressions.
	_ value.Value = (*InlineAsm)(nil)
	// Metadata values.
	_ value.Value = (*metadata.Value)(nil)
)

// Assert that each named value implements the value.Named interface.
var (
	// Other values.
	_ value.Named = (*Global)(nil)
	_ value.Named = (*Function)(nil)
	_ value.Named = (*Param)(nil)
	_ value.Named = (*BasicBlock)(nil)

	// Instructions.
	// Binary instructions.
	_ value.Named = (*InstAdd)(nil)
	_ value.Named = (*InstFAdd)(nil)
	_ value.Named = (*InstSub)(nil)
	_ value.Named = (*InstFSub)(nil)
	_ value.Named = (*InstMul)(nil)
	_ value.Named = (*InstFMul)(nil)
	_ value.Named = (*InstUDiv)(nil)
	_ value.Named = (*InstSDiv)(nil)
	_ value.Named = (*InstFDiv)(nil)
	_ value.Named = (*InstURem)(nil)
	_ value.Named = (*InstSRem)(nil)
	_ value.Named = (*InstFRem)(nil)
	// Bitwise instructions.
	_ value.Named = (*InstShl)(nil)
	_ value.Named = (*InstLShr)(nil)
	_ value.Named = (*InstAShr)(nil)
	_ value.Named = (*InstAnd)(nil)
	_ value.Named = (*InstOr)(nil)
	_ value.Named = (*InstXor)(nil)
	// Vector instructions.
	_ value.Named = (*InstExtractElement)(nil)
	_ value.Named = (*InstInsertElement)(nil)
	_ value.Named = (*InstShuffleVector)(nil)
	// Aggregate instructions.
	_ value.Named = (*InstExtractValue)(nil)
	_ value.Named = (*InstInsertValue)(nil)
	// Memory instructions.
	_ value.Named = (*InstAlloca)(nil)
	_ value.Named = (*InstLoad)(nil)
	_ value.Named = (*InstCmpXchg)(nil)
	_ value.Named = (*InstAtomicRMW)(nil)
	_ value.Named = (*InstGetElementPtr)(nil)
	// Conversion instructions.
	_ value.Named = (*InstTrunc)(nil)
	_ value.Named = (*InstZExt)(nil)
	_ value.Named = (*InstSExt)(nil)
	_ value.Named = (*InstFPTrunc)(nil)
	_ value.Named = (*InstFPExt)(nil)
	_ value.Named = (*InstFPToUI)(nil)
	_ value.Named = (*InstFPToSI)(nil)
	_ value.Named = (*InstUIToFP)(nil)
	_ value.Named = (*InstSIToFP)(nil)
	_ value.Named = (*InstPtrToInt)(nil)
	_ value.Named = (*InstIntToPtr)(nil)
	_ value.Named = (*InstBitCast)(nil)
	_ value.Named = (*InstAddrSpaceCast)(nil)
	// Other instructions.
	_ value.Named = (*InstICmp)(nil)
	_ value.Named = (*InstFCmp)(nil)
	_ value.Named = (*InstPhi)(nil)
	_ value.Named = (*InstSelect)(nil)
	_ value.Named = (*InstCall)(nil)
	_ value.Named = (*InstVAArg)(nil)
	_ value.Named = (*InstLandingPad)(nil)
	_ value.Named = (*InstCatchPad)(nil)
	_ value.Named = (*InstCleanupPad)(nil)

	// Terminators.
	_ value.Named = (*TermInvoke)(nil)
	_ value.Named = (*TermCatchSwitch)(nil) // token result used by catchpad
)
