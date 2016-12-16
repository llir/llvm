package ast

// --- [ getelementptr ] -------------------------------------------------------

// ExprGetElementPtr represents a getelementptr expression.
//
// References:
//    http://llvm.org/docs/LangRef.html#getelementptr-instruction
type ExprGetElementPtr struct {
	// Type of the constant expression.
	Type Type
	// Source address element type.
	Elem Type
	// Source address.
	Src Constant
	// Element indices.
	Indices []Constant
}

// isValue ensures that only values can be assigned to the ast.Value interface.
func (*ExprGetElementPtr) isValue() {}

// isConstant ensures that only constants can be assigned to the ast.Constant
// interface.
func (*ExprGetElementPtr) isConstant() {}

// isConstExpr ensures that only constant expressions can be assigned to the
// ast.ConstExpr interface.
func (*ExprGetElementPtr) isConstExpr() {}
