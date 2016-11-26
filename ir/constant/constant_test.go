package constant_test

import (
	"github.com/llir/llvm/ir/constant"
)

// Valutate that the relevant types satisfy the constant.Constant interface.
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
)

// Valutate that the relevant types satisfy the constant.Expr interface.
var (
	// Constant expressions.
	// Binary instructions
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
	// Bitwise instructions
	_ constant.Expr = &constant.ExprShl{}
	_ constant.Expr = &constant.ExprLShr{}
	_ constant.Expr = &constant.ExprAShr{}
	_ constant.Expr = &constant.ExprAnd{}
	_ constant.Expr = &constant.ExprOr{}
	_ constant.Expr = &constant.ExprXor{}
	// Memory instructions
	_ constant.Expr = &constant.ExprGetElementPtr{}
	// Conversion instructions
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
	// Other instructions
	_ constant.Expr = &constant.ExprICmp{}
	_ constant.Expr = &constant.ExprFCmp{}
	_ constant.Expr = &constant.ExprSelect{}
)
