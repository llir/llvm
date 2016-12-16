package ast

// --- [ getelementptr ] -------------------------------------------------------

// ExprGetElementPtr represents a getelementptr expression.
//
// References:
//    http://llvm.org/docs/LangRef.html#getelementptr-instruction
type ExprGetElementPtr struct {
	// Source address element type.
	Elem Type
	// Source address.
	Src Constant
	// Element indices.
	Indices []Constant
}
