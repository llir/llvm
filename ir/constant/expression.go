package constant

// === [ Expressions ] =========================================================

// Expression is an LLVM IR constant expression.
//
// An Expression has one of the following underlying types.
//
// Binary expressions
//
// https://llvm.org/docs/LangRef.html#constant-expressions
//
//    *constant.ExprAdd    // https://godoc.org/github.com/llir/llvm/ir/constant#ExprAdd
//    *constant.ExprFAdd   // https://godoc.org/github.com/llir/llvm/ir/constant#ExprFAdd
//    *constant.ExprSub    // https://godoc.org/github.com/llir/llvm/ir/constant#ExprSub
//    *constant.ExprFSub   // https://godoc.org/github.com/llir/llvm/ir/constant#ExprFSub
//    *constant.ExprMul    // https://godoc.org/github.com/llir/llvm/ir/constant#ExprMul
//    *constant.ExprFMul   // https://godoc.org/github.com/llir/llvm/ir/constant#ExprFMul
//    *constant.ExprUDiv   // https://godoc.org/github.com/llir/llvm/ir/constant#ExprUDiv
//    *constant.ExprSDiv   // https://godoc.org/github.com/llir/llvm/ir/constant#ExprSDiv
//    *constant.ExprFDiv   // https://godoc.org/github.com/llir/llvm/ir/constant#ExprFDiv
//    *constant.ExprURem   // https://godoc.org/github.com/llir/llvm/ir/constant#ExprURem
//    *constant.ExprSRem   // https://godoc.org/github.com/llir/llvm/ir/constant#ExprSRem
//    *constant.ExprFRem   // https://godoc.org/github.com/llir/llvm/ir/constant#ExprFRem
//
// Bitwise expressions
//
// https://llvm.org/docs/LangRef.html#constant-expressions
//
//    *constant.ExprShl    // https://godoc.org/github.com/llir/llvm/ir/constant#ExprShl
//    *constant.ExprLShr   // https://godoc.org/github.com/llir/llvm/ir/constant#ExprLShr
//    *constant.ExprAShr   // https://godoc.org/github.com/llir/llvm/ir/constant#ExprAShr
//    *constant.ExprAnd    // https://godoc.org/github.com/llir/llvm/ir/constant#ExprAnd
//    *constant.ExprOr     // https://godoc.org/github.com/llir/llvm/ir/constant#ExprOr
//    *constant.ExprXor    // https://godoc.org/github.com/llir/llvm/ir/constant#ExprXor
//
// Vector expressions
//
// https://llvm.org/docs/LangRef.html#constant-expressions
//
//    *constant.ExprExtractElement   // https://godoc.org/github.com/llir/llvm/ir/constant#ExprExtractElement
//    *constant.ExprInsertElement    // https://godoc.org/github.com/llir/llvm/ir/constant#ExprInsertElement
//    *constant.ExprShuffleVector    // https://godoc.org/github.com/llir/llvm/ir/constant#ExprShuffleVector
//
// Aggregate expressions
//
// https://llvm.org/docs/LangRef.html#constant-expressions
//
//    *constant.ExprExtractValue   // https://godoc.org/github.com/llir/llvm/ir/constant#ExprExtractValue
//    *constant.ExprInsertValue    // https://godoc.org/github.com/llir/llvm/ir/constant#ExprInsertValue
//
// Memory expressions
//
// https://llvm.org/docs/LangRef.html#constant-expressions
//
//    *constant.ExprGetElementPtr   // https://godoc.org/github.com/llir/llvm/ir/constant#ExprGetElementPtr
//
// Conversion expressions
//
// https://llvm.org/docs/LangRef.html#constant-expressions
//
//    *constant.ExprTrunc           // https://godoc.org/github.com/llir/llvm/ir/constant#ExprTrunc
//    *constant.ExprZExt            // https://godoc.org/github.com/llir/llvm/ir/constant#ExprZExt
//    *constant.ExprSExt            // https://godoc.org/github.com/llir/llvm/ir/constant#ExprSExt
//    *constant.ExprFPTrunc         // https://godoc.org/github.com/llir/llvm/ir/constant#ExprFPTrunc
//    *constant.ExprFPExt           // https://godoc.org/github.com/llir/llvm/ir/constant#ExprFPExt
//    *constant.ExprFPToUI          // https://godoc.org/github.com/llir/llvm/ir/constant#ExprFPToUI
//    *constant.ExprFPToSI          // https://godoc.org/github.com/llir/llvm/ir/constant#ExprFPToSI
//    *constant.ExprUIToFP          // https://godoc.org/github.com/llir/llvm/ir/constant#ExprUIToFP
//    *constant.ExprSIToFP          // https://godoc.org/github.com/llir/llvm/ir/constant#ExprSIToFP
//    *constant.ExprPtrToInt        // https://godoc.org/github.com/llir/llvm/ir/constant#ExprPtrToInt
//    *constant.ExprIntToPtr        // https://godoc.org/github.com/llir/llvm/ir/constant#ExprIntToPtr
//    *constant.ExprBitCast         // https://godoc.org/github.com/llir/llvm/ir/constant#ExprBitCast
//    *constant.ExprAddrSpaceCast   // https://godoc.org/github.com/llir/llvm/ir/constant#ExprAddrSpaceCast
//
// Other expressions
//
// https://llvm.org/docs/LangRef.html#constant-expressions
//
//    *constant.ExprICmp     // https://godoc.org/github.com/llir/llvm/ir/constant#ExprICmp
//    *constant.ExprFCmp     // https://godoc.org/github.com/llir/llvm/ir/constant#ExprFCmp
//    *constant.ExprSelect   // https://godoc.org/github.com/llir/llvm/ir/constant#ExprSelect
type Expression interface {
	Constant
	// Simplify returns an equivalent (and potentially simplified) constant to
	// the constant expression.
	Simplify() Constant
}
