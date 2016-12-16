package ast

// A Constant represents an LLVM IR constant; a value that is immutable at
// runtime, such as an integer or a floating point literal.
//
// Constant may have one of the following underlying types.
//
// Simple constants
//
// http://llvm.org/docs/LangRef.html#simple-constants
//
//    *ast.IntConst
//    *ast.FloatConst
//    *ast.NullConst
//
// Complex constants
//
// http://llvm.org/docs/LangRef.html#complex-constants
//
//    *ast.VectorConst
//    *ast.ArrayConst
//    *ast.StructConst
//    *ast.ZeroInitializerConst
//
// Global variable and function addresses
//
//    *ast.Global
//    *ast.Function
//
// Constant expressions
//
// http://llvm.org/docs/LangRef.html#constant-expressions
//
//    ast.ConstExpr
type Constant interface {
	Value
	// isConstant ensures that only constants can be assigned to the ast.Constant
	// interface.
	isConstant()
}
