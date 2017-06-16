// === [ Constant expressions ] ================================================
//
// References:
//    http://llvm.org/docs/LangRef.html#constant-expressions

package constant

// An Expr represents an LLVM IR constant expression.
//
// Expr may have one of the following underlying types.
//
// Binary expressions
//
// http://llvm.org/docs/LangRef.html#binary-operations
//
//    *constant.ExprAdd    (https://godoc.org/github.com/llir/llvm/ir/constant#ExprAdd)
//    *constant.ExprFAdd   (https://godoc.org/github.com/llir/llvm/ir/constant#ExprFAdd)
//    *constant.ExprSub    (https://godoc.org/github.com/llir/llvm/ir/constant#ExprSub)
//    *constant.ExprFSub   (https://godoc.org/github.com/llir/llvm/ir/constant#ExprFSub)
//    *constant.ExprMul    (https://godoc.org/github.com/llir/llvm/ir/constant#ExprMul)
//    *constant.ExprFMul   (https://godoc.org/github.com/llir/llvm/ir/constant#ExprFMul)
//    *constant.ExprUDiv   (https://godoc.org/github.com/llir/llvm/ir/constant#ExprUDiv)
//    *constant.ExprSDiv   (https://godoc.org/github.com/llir/llvm/ir/constant#ExprSDiv)
//    *constant.ExprFDiv   (https://godoc.org/github.com/llir/llvm/ir/constant#ExprFDiv)
//    *constant.ExprURem   (https://godoc.org/github.com/llir/llvm/ir/constant#ExprURem)
//    *constant.ExprSRem   (https://godoc.org/github.com/llir/llvm/ir/constant#ExprSRem)
//    *constant.ExprFRem   (https://godoc.org/github.com/llir/llvm/ir/constant#ExprFRem)
//
// Bitwise expressions
//
// http://llvm.org/docs/LangRef.html#bitwise-binary-operations
//
//    *constant.ExprShl    (https://godoc.org/github.com/llir/llvm/ir/constant#ExprShl)
//    *constant.ExprLShr   (https://godoc.org/github.com/llir/llvm/ir/constant#ExprLShr)
//    *constant.ExprAShr   (https://godoc.org/github.com/llir/llvm/ir/constant#ExprAShr)
//    *constant.ExprAnd    (https://godoc.org/github.com/llir/llvm/ir/constant#ExprAnd)
//    *constant.ExprOr     (https://godoc.org/github.com/llir/llvm/ir/constant#ExprOr)
//    *constant.ExprXor    (https://godoc.org/github.com/llir/llvm/ir/constant#ExprXor)
//
// Vector expressions
//
// http://llvm.org/docs/LangRef.html#vector-operations
//
//    *constant.ExprExtractElement   (https://godoc.org/github.com/llir/llvm/ir/constant#ExprExtractElement)
//    *constant.ExprInsertElement    (https://godoc.org/github.com/llir/llvm/ir/constant#ExprInsertElement)
//    *constant.ExprShuffleVector    (https://godoc.org/github.com/llir/llvm/ir/constant#ExprShuffleVector)
//
// Aggregate expressions
//
// http://llvm.org/docs/LangRef.html#aggregate-operations
//
//    *constant.ExprExtractValue   (https://godoc.org/github.com/llir/llvm/ir/constant#ExprExtractValue)
//    *constant.ExprInsertValue    (https://godoc.org/github.com/llir/llvm/ir/constant#ExprInsertValue)
//
// Memory expressions
//
// http://llvm.org/docs/LangRef.html#memory-access-and-addressing-operations
//
//    *constant.ExprGetElementPtr   (https://godoc.org/github.com/llir/llvm/ir/constant#ExprGetElementPtr)
//
// Conversion expressions
//
// http://llvm.org/docs/LangRef.html#conversion-operations
//
//    *constant.ExprTrunc           (https://godoc.org/github.com/llir/llvm/ir/constant#ExprTrunc)
//    *constant.ExprZExt            (https://godoc.org/github.com/llir/llvm/ir/constant#ExprZExt)
//    *constant.ExprSExt            (https://godoc.org/github.com/llir/llvm/ir/constant#ExprSExt)
//    *constant.ExprFPTrunc         (https://godoc.org/github.com/llir/llvm/ir/constant#ExprFPTrunc)
//    *constant.ExprFPExt           (https://godoc.org/github.com/llir/llvm/ir/constant#ExprFPExt)
//    *constant.ExprFPToUI          (https://godoc.org/github.com/llir/llvm/ir/constant#ExprFPToUI)
//    *constant.ExprFPToSI          (https://godoc.org/github.com/llir/llvm/ir/constant#ExprFPToSI)
//    *constant.ExprUIToFP          (https://godoc.org/github.com/llir/llvm/ir/constant#ExprUIToFP)
//    *constant.ExprSIToFP          (https://godoc.org/github.com/llir/llvm/ir/constant#ExprSIToFP)
//    *constant.ExprPtrToInt        (https://godoc.org/github.com/llir/llvm/ir/constant#ExprPtrToInt)
//    *constant.ExprIntToPtr        (https://godoc.org/github.com/llir/llvm/ir/constant#ExprIntToPtr)
//    *constant.ExprBitCast         (https://godoc.org/github.com/llir/llvm/ir/constant#ExprBitCast)
//    *constant.ExprAddrSpaceCast   (https://godoc.org/github.com/llir/llvm/ir/constant#ExprAddrSpaceCast)
//
// Other expressions
//
// http://llvm.org/docs/LangRef.html#other-operations
//
//    *constant.ExprICmp     (https://godoc.org/github.com/llir/llvm/ir/constant#ExprICmp)
//    *constant.ExprFCmp     (https://godoc.org/github.com/llir/llvm/ir/constant#ExprFCmp)
//    *constant.ExprSelect   (https://godoc.org/github.com/llir/llvm/ir/constant#ExprSelect)
type Expr interface {
	Constant
	// Simplify returns a simplified version of the constant expression.
	Simplify() Constant
}
