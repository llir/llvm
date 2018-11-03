package ir

// === [ Expressions ] =========================================================

// Expression is an LLVM IR constant expression.
//
// An Expression has one of the following underlying types.
//
// Binary expressions
//
// https://llvm.org/docs/LangRef.html#constant-expressions
//
//    *ir.ExprAdd    // https://godoc.org/github.com/llir/llvm/ir#ExprAdd
//    *ir.ExprFAdd   // https://godoc.org/github.com/llir/llvm/ir#ExprFAdd
//    *ir.ExprSub    // https://godoc.org/github.com/llir/llvm/ir#ExprSub
//    *ir.ExprFSub   // https://godoc.org/github.com/llir/llvm/ir#ExprFSub
//    *ir.ExprMul    // https://godoc.org/github.com/llir/llvm/ir#ExprMul
//    *ir.ExprFMul   // https://godoc.org/github.com/llir/llvm/ir#ExprFMul
//    *ir.ExprUDiv   // https://godoc.org/github.com/llir/llvm/ir#ExprUDiv
//    *ir.ExprSDiv   // https://godoc.org/github.com/llir/llvm/ir#ExprSDiv
//    *ir.ExprFDiv   // https://godoc.org/github.com/llir/llvm/ir#ExprFDiv
//    *ir.ExprURem   // https://godoc.org/github.com/llir/llvm/ir#ExprURem
//    *ir.ExprSRem   // https://godoc.org/github.com/llir/llvm/ir#ExprSRem
//    *ir.ExprFRem   // https://godoc.org/github.com/llir/llvm/ir#ExprFRem
//
// Bitwise expressions
//
// https://llvm.org/docs/LangRef.html#constant-expressions
//
//    *ir.ExprShl    // https://godoc.org/github.com/llir/llvm/ir#ExprShl
//    *ir.ExprLShr   // https://godoc.org/github.com/llir/llvm/ir#ExprLShr
//    *ir.ExprAShr   // https://godoc.org/github.com/llir/llvm/ir#ExprAShr
//    *ir.ExprAnd    // https://godoc.org/github.com/llir/llvm/ir#ExprAnd
//    *ir.ExprOr     // https://godoc.org/github.com/llir/llvm/ir#ExprOr
//    *ir.ExprXor    // https://godoc.org/github.com/llir/llvm/ir#ExprXor
//
// Vector expressions
//
// https://llvm.org/docs/LangRef.html#constant-expressions
//
//    *ir.ExprExtractElement   // https://godoc.org/github.com/llir/llvm/ir#ExprExtractElement
//    *ir.ExprInsertElement    // https://godoc.org/github.com/llir/llvm/ir#ExprInsertElement
//    *ir.ExprShuffleVector    // https://godoc.org/github.com/llir/llvm/ir#ExprShuffleVector
//
// Aggregate expressions
//
// https://llvm.org/docs/LangRef.html#constant-expressions
//
//    *ir.ExprExtractValue   // https://godoc.org/github.com/llir/llvm/ir#ExprExtractValue
//    *ir.ExprInsertValue    // https://godoc.org/github.com/llir/llvm/ir#ExprInsertValue
//
// Memory expressions
//
// https://llvm.org/docs/LangRef.html#constant-expressions
//
//    *ir.ExprGetElementPtr   // https://godoc.org/github.com/llir/llvm/ir#ExprGetElementPtr
//
// Conversion expressions
//
// https://llvm.org/docs/LangRef.html#constant-expressions
//
//    *ir.ExprTrunc           // https://godoc.org/github.com/llir/llvm/ir#ExprTrunc
//    *ir.ExprZExt            // https://godoc.org/github.com/llir/llvm/ir#ExprZExt
//    *ir.ExprSExt            // https://godoc.org/github.com/llir/llvm/ir#ExprSExt
//    *ir.ExprFPTrunc         // https://godoc.org/github.com/llir/llvm/ir#ExprFPTrunc
//    *ir.ExprFPExt           // https://godoc.org/github.com/llir/llvm/ir#ExprFPExt
//    *ir.ExprFPToUI          // https://godoc.org/github.com/llir/llvm/ir#ExprFPToUI
//    *ir.ExprFPToSI          // https://godoc.org/github.com/llir/llvm/ir#ExprFPToSI
//    *ir.ExprUIToFP          // https://godoc.org/github.com/llir/llvm/ir#ExprUIToFP
//    *ir.ExprSIToFP          // https://godoc.org/github.com/llir/llvm/ir#ExprSIToFP
//    *ir.ExprPtrToInt        // https://godoc.org/github.com/llir/llvm/ir#ExprPtrToInt
//    *ir.ExprIntToPtr        // https://godoc.org/github.com/llir/llvm/ir#ExprIntToPtr
//    *ir.ExprBitCast         // https://godoc.org/github.com/llir/llvm/ir#ExprBitCast
//    *ir.ExprAddrSpaceCast   // https://godoc.org/github.com/llir/llvm/ir#ExprAddrSpaceCast
//
// Other expressions
//
// https://llvm.org/docs/LangRef.html#constant-expressions
//
//    *ir.ExprICmp     // https://godoc.org/github.com/llir/llvm/ir#ExprICmp
//    *ir.ExprFCmp     // https://godoc.org/github.com/llir/llvm/ir#ExprFCmp
//    *ir.ExprSelect   // https://godoc.org/github.com/llir/llvm/ir#ExprSelect
type Expression interface {
	Constant
	// Simplify returns an equivalent (and potentially simplified) constant to
	// the constant expression.
	Simplify() Constant
}
