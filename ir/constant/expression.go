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
//	*constant.ExprFNeg   // https://pkg.go.dev/github.com/llir/llvm/ir/constant#ExprFNeg
//
// # Binary expressions
//
// https://llvm.org/docs/LangRef.html#constant-expressions
//
//	*constant.ExprAdd    // https://pkg.go.dev/github.com/llir/llvm/ir/constant#ExprAdd
//	*constant.ExprSub    // https://pkg.go.dev/github.com/llir/llvm/ir/constant#ExprSub
//	*constant.ExprMul    // https://pkg.go.dev/github.com/llir/llvm/ir/constant#ExprMul
//
// # Bitwise expressions
//
// https://llvm.org/docs/LangRef.html#constant-expressions
//
//	*constant.ExprShl    // https://pkg.go.dev/github.com/llir/llvm/ir/constant#ExprShl
//	*constant.ExprLShr   // https://pkg.go.dev/github.com/llir/llvm/ir/constant#ExprLShr
//	*constant.ExprAShr   // https://pkg.go.dev/github.com/llir/llvm/ir/constant#ExprAShr
//	*constant.ExprAnd    // https://pkg.go.dev/github.com/llir/llvm/ir/constant#ExprAnd
//	*constant.ExprOr     // https://pkg.go.dev/github.com/llir/llvm/ir/constant#ExprOr
//	*constant.ExprXor    // https://pkg.go.dev/github.com/llir/llvm/ir/constant#ExprXor
//
// # Vector expressions
//
// https://llvm.org/docs/LangRef.html#constant-expressions
//
//	*constant.ExprExtractElement   // https://pkg.go.dev/github.com/llir/llvm/ir/constant#ExprExtractElement
//	*constant.ExprInsertElement    // https://pkg.go.dev/github.com/llir/llvm/ir/constant#ExprInsertElement
//	*constant.ExprShuffleVector    // https://pkg.go.dev/github.com/llir/llvm/ir/constant#ExprShuffleVector
//
// # Memory expressions
//
// https://llvm.org/docs/LangRef.html#constant-expressions
//
//	*constant.ExprGetElementPtr   // https://pkg.go.dev/github.com/llir/llvm/ir/constant#ExprGetElementPtr
//
// # Conversion expressions
//
// https://llvm.org/docs/LangRef.html#constant-expressions
//
//	*constant.ExprTrunc           // https://pkg.go.dev/github.com/llir/llvm/ir/constant#ExprTrunc
//	*constant.ExprZExt            // https://pkg.go.dev/github.com/llir/llvm/ir/constant#ExprZExt
//	*constant.ExprSExt            // https://pkg.go.dev/github.com/llir/llvm/ir/constant#ExprSExt
//	*constant.ExprFPTrunc         // https://pkg.go.dev/github.com/llir/llvm/ir/constant#ExprFPTrunc
//	*constant.ExprFPExt           // https://pkg.go.dev/github.com/llir/llvm/ir/constant#ExprFPExt
//	*constant.ExprFPToUI          // https://pkg.go.dev/github.com/llir/llvm/ir/constant#ExprFPToUI
//	*constant.ExprFPToSI          // https://pkg.go.dev/github.com/llir/llvm/ir/constant#ExprFPToSI
//	*constant.ExprUIToFP          // https://pkg.go.dev/github.com/llir/llvm/ir/constant#ExprUIToFP
//	*constant.ExprSIToFP          // https://pkg.go.dev/github.com/llir/llvm/ir/constant#ExprSIToFP
//	*constant.ExprPtrToInt        // https://pkg.go.dev/github.com/llir/llvm/ir/constant#ExprPtrToInt
//	*constant.ExprIntToPtr        // https://pkg.go.dev/github.com/llir/llvm/ir/constant#ExprIntToPtr
//	*constant.ExprBitCast         // https://pkg.go.dev/github.com/llir/llvm/ir/constant#ExprBitCast
//	*constant.ExprAddrSpaceCast   // https://pkg.go.dev/github.com/llir/llvm/ir/constant#ExprAddrSpaceCast
//
// # Other expressions
//
// https://llvm.org/docs/LangRef.html#constant-expressions
//
//	*constant.ExprICmp     // https://pkg.go.dev/github.com/llir/llvm/ir/constant#ExprICmp
//	*constant.ExprFCmp     // https://pkg.go.dev/github.com/llir/llvm/ir/constant#ExprFCmp
//	*constant.ExprSelect   // https://pkg.go.dev/github.com/llir/llvm/ir/constant#ExprSelect
type Expression interface {
	Constant
	// IsExpression ensures that only constants expressions can be assigned to
	// the constant.Expression interface.
	IsExpression()
}
