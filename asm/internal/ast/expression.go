package ast

// A ConstExpr represents an LLVM IR constant expression.
//
// ConstExpr may have one of the following underlying types.
//
// Binary expressions
//
// http://llvm.org/docs/LangRef.html#binary-operations
//
//    *ast.ExprAdd
//    *ast.ExprFAdd
//    *ast.ExprSub
//    *ast.ExprFSub
//    *ast.ExprMul
//    *ast.ExprFMul
//    *ast.ExprUDiv
//    *ast.ExprSDiv
//    *ast.ExprFDiv
//    *ast.ExprURem
//    *ast.ExprSRem
//    *ast.ExprFRem
//
// Bitwise expressions
//
// http://llvm.org/docs/LangRef.html#bitwise-binary-operations
//
//    *ast.ExprShl
//    *ast.ExprLShr
//    *ast.ExprAShr
//    *ast.ExprAnd
//    *ast.ExprOr
//    *ast.ExprXor
//
// Vector expressions
//
// http://llvm.org/docs/LangRef.html#vector-operations
//
//    *ast.ExprExtractElement
//    *ast.ExprInsertElement
//    *ast.ExprShuffleVector
//
// Aggregate expressions
//
// http://llvm.org/docs/LangRef.html#aggregate-operations
//
//    *ast.ExprExtractValue
//    *ast.ExprInsertValue
//
// Memory expressions
//
// http://llvm.org/docs/LangRef.html#memory-access-and-addressing-operations
//
//    *ast.ExprGetElementPtr
//
// Conversion expressions
//
// http://llvm.org/docs/LangRef.html#conversion-operations
//
//    *ast.ExprTrunc
//    *ast.ExprZExt
//    *ast.ExprSExt
//    *ast.ExprFPTrunc
//    *ast.ExprFPExt
//    *ast.ExprFPToUI
//    *ast.ExprFPToSI
//    *ast.ExprUIToFP
//    *ast.ExprSIToFP
//    *ast.ExprPtrToInt
//    *ast.ExprIntToPtr
//    *ast.ExprBitCast
//    *ast.ExprAddrSpaceCast
//
// Other expressions
//
// http://llvm.org/docs/LangRef.html#other-operations
//
//    *ast.ExprICmp
//    *ast.ExprFCmp
//    *ast.ExprSelect
type ConstExpr interface {
	Constant
	// isConstExpr ensures that only constant expressions can be assigned to the
	// ast.ConstExpr interface.
	isConstExpr()
}
