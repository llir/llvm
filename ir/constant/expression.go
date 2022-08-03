package constant

// === [ Expressions ] =========================================================

// Expression is an LLVM IR constant expression.
//
// An Expression has one of the following underlying types.
//
// # Unary expressions
//
// https://llvm.org/docs/LangRef.html#constant-expressions
//
//   - [*constant.ExprFNeg]
//
// # Binary expressions
//
// https://llvm.org/docs/LangRef.html#constant-expressions
//
//   - [*constant.ExprAdd]
//   - [*constant.ExprSub]
//   - [*constant.ExprMul]
//
// # Bitwise expressions
//
// https://llvm.org/docs/LangRef.html#constant-expressions
//
//   - [*constant.ExprShl]
//   - [*constant.ExprLShr]
//   - [*constant.ExprAShr]
//   - [*constant.ExprAnd]
//   - [*constant.ExprOr]
//   - [*constant.ExprXor]
//
// # Vector expressions
//
// https://llvm.org/docs/LangRef.html#constant-expressions
//
//   - [*constant.ExprExtractElement]
//   - [*constant.ExprInsertElement]
//   - [*constant.ExprShuffleVector]
//
// # Memory expressions
//
// https://llvm.org/docs/LangRef.html#constant-expressions
//
//   - [*constant.ExprGetElementPtr]
//
// # Conversion expressions
//
// https://llvm.org/docs/LangRef.html#constant-expressions
//
//   - [*constant.ExprTrunc]
//   - [*constant.ExprZExt]
//   - [*constant.ExprSExt]
//   - [*constant.ExprFPTrunc]
//   - [*constant.ExprFPExt]
//   - [*constant.ExprFPToUI]
//   - [*constant.ExprFPToSI]
//   - [*constant.ExprUIToFP]
//   - [*constant.ExprSIToFP]
//   - [*constant.ExprPtrToInt]
//   - [*constant.ExprIntToPtr]
//   - [*constant.ExprBitCast]
//   - [*constant.ExprAddrSpaceCast]
//
// # Other expressions
//
// https://llvm.org/docs/LangRef.html#constant-expressions
//
//   - [*constant.ExprICmp]
//   - [*constant.ExprFCmp]
//   - [*constant.ExprSelect]
type Expression interface {
	Constant
	// IsExpression ensures that only constants expressions can be assigned to
	// the constant.Expression interface.
	IsExpression()
}
