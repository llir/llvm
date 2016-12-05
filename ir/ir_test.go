package ir_test

import (
	"github.com/llir/llvm/ir"
	"github.com/llir/llvm/ir/constant"
	"github.com/llir/llvm/ir/value"
)

// Valutate that the relevant types satisfy the constant.Constant interface.
var (
	_ constant.Constant = &ir.Function{}
	_ constant.Constant = &ir.Global{}
)

// Valutate that the relevant types satisfy the ir.Instruction interface.
var (
	// Binary instructions
	_ ir.Instruction = &ir.InstAdd{}
	_ ir.Instruction = &ir.InstFAdd{}
	_ ir.Instruction = &ir.InstSub{}
	_ ir.Instruction = &ir.InstFSub{}
	_ ir.Instruction = &ir.InstMul{}
	_ ir.Instruction = &ir.InstFMul{}
	_ ir.Instruction = &ir.InstUDiv{}
	_ ir.Instruction = &ir.InstSDiv{}
	_ ir.Instruction = &ir.InstFDiv{}
	_ ir.Instruction = &ir.InstURem{}
	_ ir.Instruction = &ir.InstSRem{}
	_ ir.Instruction = &ir.InstFRem{}
	// Bitwise instructions
	_ ir.Instruction = &ir.InstShl{}
	_ ir.Instruction = &ir.InstLShr{}
	_ ir.Instruction = &ir.InstAShr{}
	_ ir.Instruction = &ir.InstAnd{}
	_ ir.Instruction = &ir.InstOr{}
	_ ir.Instruction = &ir.InstXor{}
	// Vector instructions
	// Aggregate instructions
	// Memory instructions
	_ ir.Instruction = &ir.InstAlloca{}
	_ ir.Instruction = &ir.InstLoad{}
	_ ir.Instruction = &ir.InstStore{}
	_ ir.Instruction = &ir.InstGetElementPtr{}
	// Conversion instructions
	_ ir.Instruction = &ir.InstTrunc{}
	_ ir.Instruction = &ir.InstZExt{}
	_ ir.Instruction = &ir.InstSExt{}
	_ ir.Instruction = &ir.InstFPTrunc{}
	_ ir.Instruction = &ir.InstFPExt{}
	_ ir.Instruction = &ir.InstFPToUI{}
	_ ir.Instruction = &ir.InstFPToSI{}
	_ ir.Instruction = &ir.InstUIToFP{}
	_ ir.Instruction = &ir.InstSIToFP{}
	_ ir.Instruction = &ir.InstPtrToInt{}
	_ ir.Instruction = &ir.InstIntToPtr{}
	_ ir.Instruction = &ir.InstBitCast{}
	_ ir.Instruction = &ir.InstAddrSpaceCast{}
	// Other instructions
	_ ir.Instruction = &ir.InstICmp{}
	_ ir.Instruction = &ir.InstFCmp{}
	_ ir.Instruction = &ir.InstPhi{}
	_ ir.Instruction = &ir.InstSelect{}
	_ ir.Instruction = &ir.InstCall{}
)

// Valutate that the relevant types satisfy the ir.Terminator interface.
var (
	// Terminators
	_ ir.Terminator = &ir.TermRet{}
	_ ir.Terminator = &ir.TermBr{}
	_ ir.Terminator = &ir.TermCondBr{}
	_ ir.Terminator = &ir.TermSwitch{}
	_ ir.Terminator = &ir.TermUnreachable{}
)

// Valutate that the relevant types satisfy the value.Value interface.
var (
	_ value.Value = &ir.BasicBlock{}
	// Binary instructions
	_ value.Value = &ir.InstAdd{}
	_ value.Value = &ir.InstFAdd{}
	_ value.Value = &ir.InstSub{}
	_ value.Value = &ir.InstFSub{}
	_ value.Value = &ir.InstMul{}
	_ value.Value = &ir.InstFMul{}
	_ value.Value = &ir.InstUDiv{}
	_ value.Value = &ir.InstSDiv{}
	_ value.Value = &ir.InstFDiv{}
	_ value.Value = &ir.InstURem{}
	_ value.Value = &ir.InstSRem{}
	_ value.Value = &ir.InstFRem{}
	// Bitwise instructions
	_ value.Value = &ir.InstShl{}
	_ value.Value = &ir.InstLShr{}
	_ value.Value = &ir.InstAShr{}
	_ value.Value = &ir.InstAnd{}
	_ value.Value = &ir.InstOr{}
	_ value.Value = &ir.InstXor{}
	// Vector instructions
	// Aggregate instructions
	// Memory instructions
	_ value.Value = &ir.InstAlloca{}
	_ value.Value = &ir.InstLoad{}
	_ value.Value = &ir.InstGetElementPtr{}
	// Conversion instructions
	_ value.Value = &ir.InstTrunc{}
	_ value.Value = &ir.InstZExt{}
	_ value.Value = &ir.InstSExt{}
	_ value.Value = &ir.InstFPTrunc{}
	_ value.Value = &ir.InstFPExt{}
	_ value.Value = &ir.InstFPToUI{}
	_ value.Value = &ir.InstFPToSI{}
	_ value.Value = &ir.InstUIToFP{}
	_ value.Value = &ir.InstSIToFP{}
	_ value.Value = &ir.InstPtrToInt{}
	_ value.Value = &ir.InstIntToPtr{}
	_ value.Value = &ir.InstBitCast{}
	_ value.Value = &ir.InstAddrSpaceCast{}
	// Other instructions
	_ value.Value = &ir.InstICmp{}
	_ value.Value = &ir.InstFCmp{}
	_ value.Value = &ir.InstPhi{}
	_ value.Value = &ir.InstSelect{}
	_ value.Value = &ir.InstCall{}
)
