package constant

import "github.com/llir/llvm/ir"

// Assert that each constant expression implements the ir.Expression interface.
var (
	// Binary expressions.
	_ ir.Expression = (*ExprAdd)(nil)
	_ ir.Expression = (*ExprFAdd)(nil)
	_ ir.Expression = (*ExprSub)(nil)
	_ ir.Expression = (*ExprFSub)(nil)
	_ ir.Expression = (*ExprMul)(nil)
	_ ir.Expression = (*ExprFMul)(nil)
	_ ir.Expression = (*ExprUDiv)(nil)
	_ ir.Expression = (*ExprSDiv)(nil)
	_ ir.Expression = (*ExprFDiv)(nil)
	_ ir.Expression = (*ExprURem)(nil)
	_ ir.Expression = (*ExprSRem)(nil)
	_ ir.Expression = (*ExprFRem)(nil)
	// Bitwise expressions.
	_ ir.Expression = (*ExprShl)(nil)
	_ ir.Expression = (*ExprLShr)(nil)
	_ ir.Expression = (*ExprAShr)(nil)
	_ ir.Expression = (*ExprAnd)(nil)
	_ ir.Expression = (*ExprOr)(nil)
	_ ir.Expression = (*ExprXor)(nil)
	// Vector expressions.
	_ ir.Expression = (*ExprExtractElement)(nil)
	_ ir.Expression = (*ExprInsertElement)(nil)
	_ ir.Expression = (*ExprShuffleVector)(nil)
	// Aggregate expressions.
	_ ir.Expression = (*ExprExtractValue)(nil)
	_ ir.Expression = (*ExprInsertValue)(nil)
	// Memory expressions.
	_ ir.Expression = (*ExprGetElementPtr)(nil)
	// Conversion expressions.
	_ ir.Expression = (*ExprTrunc)(nil)
	_ ir.Expression = (*ExprZExt)(nil)
	_ ir.Expression = (*ExprSExt)(nil)
	_ ir.Expression = (*ExprFPTrunc)(nil)
	_ ir.Expression = (*ExprFPExt)(nil)
	_ ir.Expression = (*ExprFPToUI)(nil)
	_ ir.Expression = (*ExprFPToSI)(nil)
	_ ir.Expression = (*ExprUIToFP)(nil)
	_ ir.Expression = (*ExprSIToFP)(nil)
	_ ir.Expression = (*ExprPtrToInt)(nil)
	_ ir.Expression = (*ExprIntToPtr)(nil)
	_ ir.Expression = (*ExprBitCast)(nil)
	_ ir.Expression = (*ExprAddrSpaceCast)(nil)
	// Other expressions.
	_ ir.Expression = (*ExprICmp)(nil)
	_ ir.Expression = (*ExprFCmp)(nil)
	_ ir.Expression = (*ExprSelect)(nil)
)
