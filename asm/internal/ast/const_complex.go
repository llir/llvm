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
}

// CharArrayConst represents a character array constant.
type CharArrayConst struct {
	// Array type.
	Type Type
	// Array elements.
	Lit string
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

// isValue ensures that only values can be assigned to the ast.Value interface.
func (*VectorConst) isValue()          {}
func (*ArrayConst) isValue()           {}
func (*CharArrayConst) isValue()       {}
func (*StructConst) isValue()          {}
func (*ZeroInitializerConst) isValue() {}

// isConstant ensures that only constants can be assigned to the ast.Constant
// interface.
func (*VectorConst) isConstant()          {}
func (*ArrayConst) isConstant()           {}
func (*CharArrayConst) isConstant()       {}
func (*StructConst) isConstant()          {}
func (*ZeroInitializerConst) isConstant() {}
