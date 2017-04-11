package ir_test

import (
	"github.com/llir/llvm/ir"
	"github.com/llir/llvm/ir/constant"
	"github.com/llir/llvm/ir/value"
)

// Validate that the relevant types satisfy the constant.Constant interface.
var (
	_ constant.Constant = &ir.Global{}
	_ constant.Constant = &ir.Function{}
)

// Validate that the relevant types satisfy the ir.Instruction interface.
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
	_ ir.Instruction = &ir.InstExtractElement{}
	_ ir.Instruction = &ir.InstInsertElement{}
	_ ir.Instruction = &ir.InstShuffleVector{}
	// Aggregate instructions
	_ ir.Instruction = &ir.InstExtractValue{}
	_ ir.Instruction = &ir.InstInsertValue{}
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

// Validate that the relevant types satisfy the ir.Terminator interface.
var (
	// Terminators
	_ ir.Terminator = &ir.TermRet{}
	_ ir.Terminator = &ir.TermBr{}
	_ ir.Terminator = &ir.TermCondBr{}
	_ ir.Terminator = &ir.TermSwitch{}
	_ ir.Terminator = &ir.TermUnreachable{}
)

// Validate that the relevant types satisfy the value.Named interface.
var (
	_ value.Named = &ir.Global{}
	_ value.Named = &ir.Function{}
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
	_ value.Named = &ir.InstExtractElement{}
	_ value.Named = &ir.InstInsertElement{}
	_ value.Named = &ir.InstShuffleVector{}
	// Aggregate instructions
	_ value.Named = &ir.InstExtractValue{}
	_ value.Named = &ir.InstInsertValue{}
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

// Validate that the relevant types satisfy the ir.MetadataNode interface.
var (
	_ ir.MetadataNode = &ir.Metadata{}
	_ ir.MetadataNode = &ir.MetadataString{}
	// Simple constants.
	_ ir.MetadataNode = &constant.Int{}
	_ ir.MetadataNode = &constant.Float{}
	_ ir.MetadataNode = &constant.Null{}
	// Complex constants.
	_ ir.MetadataNode = &constant.Vector{}
	_ ir.MetadataNode = &constant.Array{}
	_ ir.MetadataNode = &constant.Struct{}
	_ ir.MetadataNode = &constant.ZeroInitializer{}
	// Constant expressions.
	// Binary instructions
	_ ir.MetadataNode = &constant.ExprAdd{}
	_ ir.MetadataNode = &constant.ExprFAdd{}
	_ ir.MetadataNode = &constant.ExprSub{}
	_ ir.MetadataNode = &constant.ExprFSub{}
	_ ir.MetadataNode = &constant.ExprMul{}
	_ ir.MetadataNode = &constant.ExprFMul{}
	_ ir.MetadataNode = &constant.ExprUDiv{}
	_ ir.MetadataNode = &constant.ExprSDiv{}
	_ ir.MetadataNode = &constant.ExprFDiv{}
	_ ir.MetadataNode = &constant.ExprURem{}
	_ ir.MetadataNode = &constant.ExprSRem{}
	_ ir.MetadataNode = &constant.ExprFRem{}
	// Bitwise instructions
	_ ir.MetadataNode = &constant.ExprShl{}
	_ ir.MetadataNode = &constant.ExprLShr{}
	_ ir.MetadataNode = &constant.ExprAShr{}
	_ ir.MetadataNode = &constant.ExprAnd{}
	_ ir.MetadataNode = &constant.ExprOr{}
	_ ir.MetadataNode = &constant.ExprXor{}
	// Memory instructions
	_ ir.MetadataNode = &constant.ExprGetElementPtr{}
	// Conversion instructions
	_ ir.MetadataNode = &constant.ExprTrunc{}
	_ ir.MetadataNode = &constant.ExprZExt{}
	_ ir.MetadataNode = &constant.ExprSExt{}
	_ ir.MetadataNode = &constant.ExprFPTrunc{}
	_ ir.MetadataNode = &constant.ExprFPExt{}
	_ ir.MetadataNode = &constant.ExprFPToUI{}
	_ ir.MetadataNode = &constant.ExprFPToSI{}
	_ ir.MetadataNode = &constant.ExprUIToFP{}
	_ ir.MetadataNode = &constant.ExprSIToFP{}
	_ ir.MetadataNode = &constant.ExprPtrToInt{}
	_ ir.MetadataNode = &constant.ExprIntToPtr{}
	_ ir.MetadataNode = &constant.ExprBitCast{}
	_ ir.MetadataNode = &constant.ExprAddrSpaceCast{}
	// Other instructions
	_ ir.MetadataNode = &constant.ExprICmp{}
	_ ir.MetadataNode = &constant.ExprFCmp{}
	_ ir.MetadataNode = &constant.ExprSelect{}
)
