package constant_test

import (
	"github.com/llir/llvm/ir/constant"
	"github.com/llir/llvm/ir/metadata"
)

// Validate that the relevant types satisfy the constant.Constant interface.
var (
	// Simple constants.
	_ constant.Constant = &constant.Int{}
	_ constant.Constant = &constant.Float{}
	_ constant.Constant = &constant.Null{}
	// Complex constants.
	_ constant.Constant = &constant.Vector{}
	_ constant.Constant = &constant.Array{}
	_ constant.Constant = &constant.Struct{}
	_ constant.Constant = &constant.ZeroInitializer{}
	_ constant.Constant = &constant.Undef{}
)

// Validate that the relevant types satisfy the constant.Expr interface.
var (
	// Binary expressions.
	_ constant.Expr = &constant.ExprAdd{}
	_ constant.Expr = &constant.ExprFAdd{}
	_ constant.Expr = &constant.ExprSub{}
	_ constant.Expr = &constant.ExprFSub{}
	_ constant.Expr = &constant.ExprMul{}
	_ constant.Expr = &constant.ExprFMul{}
	_ constant.Expr = &constant.ExprUDiv{}
	_ constant.Expr = &constant.ExprSDiv{}
	_ constant.Expr = &constant.ExprFDiv{}
	_ constant.Expr = &constant.ExprURem{}
	_ constant.Expr = &constant.ExprSRem{}
	_ constant.Expr = &constant.ExprFRem{}
	// Bitwise expressions.
	_ constant.Expr = &constant.ExprShl{}
	_ constant.Expr = &constant.ExprLShr{}
	_ constant.Expr = &constant.ExprAShr{}
	_ constant.Expr = &constant.ExprAnd{}
	_ constant.Expr = &constant.ExprOr{}
	_ constant.Expr = &constant.ExprXor{}
	// Vector expressions.
	_ constant.Expr = &constant.ExprExtractElement{}
	_ constant.Expr = &constant.ExprInsertElement{}
	_ constant.Expr = &constant.ExprShuffleVector{}
	// Aggregate expressions.
	_ constant.Expr = &constant.ExprExtractValue{}
	_ constant.Expr = &constant.ExprInsertValue{}
	// Memory expressions.
	_ constant.Expr = &constant.ExprGetElementPtr{}
	// Conversion expressions.
	_ constant.Expr = &constant.ExprTrunc{}
	_ constant.Expr = &constant.ExprZExt{}
	_ constant.Expr = &constant.ExprSExt{}
	_ constant.Expr = &constant.ExprFPTrunc{}
	_ constant.Expr = &constant.ExprFPExt{}
	_ constant.Expr = &constant.ExprFPToUI{}
	_ constant.Expr = &constant.ExprFPToSI{}
	_ constant.Expr = &constant.ExprUIToFP{}
	_ constant.Expr = &constant.ExprSIToFP{}
	_ constant.Expr = &constant.ExprPtrToInt{}
	_ constant.Expr = &constant.ExprIntToPtr{}
	_ constant.Expr = &constant.ExprBitCast{}
	_ constant.Expr = &constant.ExprAddrSpaceCast{}
	// Other expressions.
	_ constant.Expr = &constant.ExprICmp{}
	_ constant.Expr = &constant.ExprFCmp{}
	_ constant.Expr = &constant.ExprSelect{}
)

// Validate that the relevant types satisfy the metadata.Node interface.
var (
	// Simple constants.
	_ metadata.Node = &constant.Int{}
	_ metadata.Node = &constant.Float{}
	_ metadata.Node = &constant.Null{}
	// Complex constants.
	_ metadata.Node = &constant.Vector{}
	_ metadata.Node = &constant.Array{}
	_ metadata.Node = &constant.Struct{}
	_ metadata.Node = &constant.ZeroInitializer{}
	_ metadata.Node = &constant.Undef{}
	// Binary expressions.
	_ metadata.Node = &constant.ExprAdd{}
	_ metadata.Node = &constant.ExprFAdd{}
	_ metadata.Node = &constant.ExprSub{}
	_ metadata.Node = &constant.ExprFSub{}
	_ metadata.Node = &constant.ExprMul{}
	_ metadata.Node = &constant.ExprFMul{}
	_ metadata.Node = &constant.ExprUDiv{}
	_ metadata.Node = &constant.ExprSDiv{}
	_ metadata.Node = &constant.ExprFDiv{}
	_ metadata.Node = &constant.ExprURem{}
	_ metadata.Node = &constant.ExprSRem{}
	_ metadata.Node = &constant.ExprFRem{}
	// Bitwise expressions.
	_ metadata.Node = &constant.ExprShl{}
	_ metadata.Node = &constant.ExprLShr{}
	_ metadata.Node = &constant.ExprAShr{}
	_ metadata.Node = &constant.ExprAnd{}
	_ metadata.Node = &constant.ExprOr{}
	_ metadata.Node = &constant.ExprXor{}
	// Vector expressions.
	_ metadata.Node = &constant.ExprExtractElement{}
	_ metadata.Node = &constant.ExprInsertElement{}
	_ metadata.Node = &constant.ExprShuffleVector{}
	// Aggregate expressions.
	_ metadata.Node = &constant.ExprExtractValue{}
	_ metadata.Node = &constant.ExprInsertValue{}
	// Memory expressions.
	_ metadata.Node = &constant.ExprGetElementPtr{}
	// Conversion expressions.
	_ metadata.Node = &constant.ExprTrunc{}
	_ metadata.Node = &constant.ExprZExt{}
	_ metadata.Node = &constant.ExprSExt{}
	_ metadata.Node = &constant.ExprFPTrunc{}
	_ metadata.Node = &constant.ExprFPExt{}
	_ metadata.Node = &constant.ExprFPToUI{}
	_ metadata.Node = &constant.ExprFPToSI{}
	_ metadata.Node = &constant.ExprUIToFP{}
	_ metadata.Node = &constant.ExprSIToFP{}
	_ metadata.Node = &constant.ExprPtrToInt{}
	_ metadata.Node = &constant.ExprIntToPtr{}
	_ metadata.Node = &constant.ExprBitCast{}
	_ metadata.Node = &constant.ExprAddrSpaceCast{}
	// Other expressions.
	_ metadata.Node = &constant.ExprICmp{}
	_ metadata.Node = &constant.ExprFCmp{}
	_ metadata.Node = &constant.ExprSelect{}
)
