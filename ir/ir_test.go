package ir_test

import (
	"github.com/llir/llvm/ir"
	"github.com/llir/llvm/ir/constant"
	"github.com/llir/llvm/ir/value"
)

// Valutate that the relevant types satisfy the constant.Constant interface.
var (
	_ constant.Constant = &ir.Global{}
	_ constant.Constant = &ir.Function{}
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

// Valutate that the relevant types satisfy the value.Named interface.
var (
	_ value.Named = &ir.Global{}
	_ value.Named = &ir.Function{}
	_ value.Named = &ir.Param{}
	_ value.Named = &ir.BasicBlock{}
	// Binary instructions
	_ value.Named = &ir.InstAdd{}
	_ value.Named = &ir.InstFAdd{}
	_ value.Named = &ir.InstSub{}
	_ value.Named = &ir.InstFSub{}
	_ value.Named = &ir.InstMul{}
	_ value.Named = &ir.InstFMul{}
	_ value.Named = &ir.InstUDiv{}
	_ value.Named = &ir.InstSDiv{}
	_ value.Named = &ir.InstFDiv{}
	_ value.Named = &ir.InstURem{}
	_ value.Named = &ir.InstSRem{}
	_ value.Named = &ir.InstFRem{}
	// Bitwise instructions
	_ value.Named = &ir.InstShl{}
	_ value.Named = &ir.InstLShr{}
	_ value.Named = &ir.InstAShr{}
	_ value.Named = &ir.InstAnd{}
	_ value.Named = &ir.InstOr{}
	_ value.Named = &ir.InstXor{}
	// Vector instructions
	// Aggregate instructions
	// Memory instructions
	_ value.Named = &ir.InstAlloca{}
	_ value.Named = &ir.InstLoad{}
	_ value.Named = &ir.InstGetElementPtr{}
	// Conversion instructions
	_ value.Named = &ir.InstTrunc{}
	_ value.Named = &ir.InstZExt{}
	_ value.Named = &ir.InstSExt{}
	_ value.Named = &ir.InstFPTrunc{}
	_ value.Named = &ir.InstFPExt{}
	_ value.Named = &ir.InstFPToUI{}
	_ value.Named = &ir.InstFPToSI{}
	_ value.Named = &ir.InstUIToFP{}
	_ value.Named = &ir.InstSIToFP{}
	_ value.Named = &ir.InstPtrToInt{}
	_ value.Named = &ir.InstIntToPtr{}
	_ value.Named = &ir.InstBitCast{}
	_ value.Named = &ir.InstAddrSpaceCast{}
	// Other instructions
	_ value.Named = &ir.InstICmp{}
	_ value.Named = &ir.InstFCmp{}
	_ value.Named = &ir.InstPhi{}
	_ value.Named = &ir.InstSelect{}
	_ value.Named = &ir.InstCall{}
)
