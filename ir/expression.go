package ir

// Expression is an LLVM IR constant expression.
//
// An Expression has one of the following underlying types.
//
// Binary expressions
//
// https://llvm.org/docs/LangRef.html#constant-expressions
//
//    *ir.ExprAdd    // https://godoc.org/github.com/llir/l/ir#ExprAdd
//    *ir.ExprFAdd   // https://godoc.org/github.com/llir/l/ir#ExprFAdd
//    *ir.ExprSub    // https://godoc.org/github.com/llir/l/ir#ExprSub
//    *ir.ExprFSub   // https://godoc.org/github.com/llir/l/ir#ExprFSub
//    *ir.ExprMul    // https://godoc.org/github.com/llir/l/ir#ExprMul
//    *ir.ExprFMul   // https://godoc.org/github.com/llir/l/ir#ExprFMul
//    *ir.ExprUDiv   // https://godoc.org/github.com/llir/l/ir#ExprUDiv
//    *ir.ExprSDiv   // https://godoc.org/github.com/llir/l/ir#ExprSDiv
//    *ir.ExprFDiv   // https://godoc.org/github.com/llir/l/ir#ExprFDiv
//    *ir.ExprURem   // https://godoc.org/github.com/llir/l/ir#ExprURem
//    *ir.ExprSRem   // https://godoc.org/github.com/llir/l/ir#ExprSRem
//    *ir.ExprFRem   // https://godoc.org/github.com/llir/l/ir#ExprFRem
//
// Bitwise expressions
//
// https://llvm.org/docs/LangRef.html#constant-expressions
//
//    *ir.ExprShl    // https://godoc.org/github.com/llir/l/ir#ExprShl
//    *ir.ExprLShr   // https://godoc.org/github.com/llir/l/ir#ExprLShr
//    *ir.ExprAShr   // https://godoc.org/github.com/llir/l/ir#ExprAShr
//    *ir.ExprAnd    // https://godoc.org/github.com/llir/l/ir#ExprAnd
//    *ir.ExprOr     // https://godoc.org/github.com/llir/l/ir#ExprOr
//    *ir.ExprXor    // https://godoc.org/github.com/llir/l/ir#ExprXor
//
// Vector expressions
//
// https://llvm.org/docs/LangRef.html#constant-expressions
//
//    *ir.ExprExtractElement   // https://godoc.org/github.com/llir/l/ir#ExprExtractElement
//    *ir.ExprInsertElement    // https://godoc.org/github.com/llir/l/ir#ExprInsertElement
//    *ir.ExprShuffleVector    // https://godoc.org/github.com/llir/l/ir#ExprShuffleVector
//
// Aggregate expressions
//
// https://llvm.org/docs/LangRef.html#constant-expressions
//
//    *ir.ExprExtractValue   // https://godoc.org/github.com/llir/l/ir#ExprExtractValue
//    *ir.ExprInsertValue    // https://godoc.org/github.com/llir/l/ir#ExprInsertValue
//
// Memory expressions
//
// https://llvm.org/docs/LangRef.html#constant-expressions
//
//    *ir.ExprGetElementPtr   // https://godoc.org/github.com/llir/l/ir#ExprGetElementPtr
//
// Conversion expressions
//
// https://llvm.org/docs/LangRef.html#constant-expressions
//
//    *ir.ExprTrunc           // https://godoc.org/github.com/llir/l/ir#ExprTrunc
//    *ir.ExprZExt            // https://godoc.org/github.com/llir/l/ir#ExprZExt
//    *ir.ExprSExt            // https://godoc.org/github.com/llir/l/ir#ExprSExt
//    *ir.ExprFPTrunc         // https://godoc.org/github.com/llir/l/ir#ExprFPTrunc
//    *ir.ExprFPExt           // https://godoc.org/github.com/llir/l/ir#ExprFPExt
//    *ir.ExprFPToUI          // https://godoc.org/github.com/llir/l/ir#ExprFPToUI
//    *ir.ExprFPToSI          // https://godoc.org/github.com/llir/l/ir#ExprFPToSI
//    *ir.ExprUIToFP          // https://godoc.org/github.com/llir/l/ir#ExprUIToFP
//    *ir.ExprSIToFP          // https://godoc.org/github.com/llir/l/ir#ExprSIToFP
//    *ir.ExprPtrToInt        // https://godoc.org/github.com/llir/l/ir#ExprPtrToInt
//    *ir.ExprIntToPtr        // https://godoc.org/github.com/llir/l/ir#ExprIntToPtr
//    *ir.ExprBitCast         // https://godoc.org/github.com/llir/l/ir#ExprBitCast
//    *ir.ExprAddrSpaceCast   // https://godoc.org/github.com/llir/l/ir#ExprAddrSpaceCast
//
// Other expressions
//
// https://llvm.org/docs/LangRef.html#constant-expressions
//
//    *ir.ExprICmp     // https://godoc.org/github.com/llir/l/ir#ExprICmp
//    *ir.ExprFCmp     // https://godoc.org/github.com/llir/l/ir#ExprFCmp
//    *ir.ExprSelect   // https://godoc.org/github.com/llir/l/ir#ExprSelect
type Expression interface {
	Constant
	// Simplify returns an equivalent (and potentially simplified) constant of
	// the constant expression.
	Simplify() Constant
}
