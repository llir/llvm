package ast

// IntConst represents an integer constant.
type IntConst struct {
	// Integer type.
	Type Type
	// Constant literal value.
	Lit string
}

// FloatConst represents a floating-point constant.
type FloatConst struct {
	// Floating-point type.
	Type Type
	// Constant literal value.
	Lit string
}

// NullConst represents a null pointer constant.
type NullConst struct {
	// Pointer type.
	Type Type
}

// isValue ensures that only values can be assigned to the ast.Value interface.
func (*IntConst) isValue()   {}
func (*FloatConst) isValue() {}
func (*NullConst) isValue()  {}

// isConstant ensures that only constants can be assigned to the ast.Constant
// interface.
func (*IntConst) isConstant()   {}
func (*FloatConst) isConstant() {}
func (*NullConst) isConstant()  {}
