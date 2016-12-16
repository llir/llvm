package ast

// VectorConst represents a vector constant.
type VectorConst struct {
	// Vector type.
	Type Type
	// Vector elements.
	Elems []Constant
}

// ArrayConst represents an array constant.
type ArrayConst struct {
	// Array type.
	Type Type
	// Array elements.
	Elems []Constant
	// Pretty-print as character array.
	CharArray bool
}

// StructConst represents a struct constant.
type StructConst struct {
	// Struct type.
	Type Type
	// Struct fields.
	Fields []Constant
}

// ZeroInitializerConst represents a zeroinitializer constant.
type ZeroInitializerConst struct {
	// Constant type.
	Type Type
}

// isConstant ensures that only constants can be assigned to the ast.Constant
// interface.
func (*VectorConst) isConstant()          {}
func (*ArrayConst) isConstant()           {}
func (*StructConst) isConstant()          {}
func (*ZeroInitializerConst) isConstant() {}
