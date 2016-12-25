package ast

// --- [ function ] ------------------------------------------------------------

// FuncType represents a function type.
//
// References:
//    http://llvm.org/docs/LangRef.html#function-type
type FuncType struct {
	// Return type.
	Ret Type
	// Function parameters.
	Params []*Param
	// Variadicity of the function type.
	Variadic bool
}

// A Param represents an LLVM IR function parameter.
type Param struct {
	// Parameter name.
	Name string
	// Parameter type.
	Type Type
}

// GetName returns the name of the value.
func (param *Param) GetName() string {
	return param.Name
}

// SetName sets the name of the value.
func (param *Param) SetName(name string) {
	param.Name = name
}

// isValue ensures that only values can be assigned to the ast.Value interface.
func (*Param) isValue() {}

// --- [ pointer ] -------------------------------------------------------------

// PointerType represents a pointer type.
//
// References:
//    http://llvm.org/docs/LangRef.html#pointer-type
type PointerType struct {
	// Element type.
	Elem Type
	// Address space.
	AddrSpace int64
}

// --- [ vector ] --------------------------------------------------------------

// VectorType represents a vector type.
//
// References:
//    http://llvm.org/docs/LangRef.html#vector-type
type VectorType struct {
	// Element type.
	Elem Type
	// Vector length.
	Len int64
}

// --- [ array ] ---------------------------------------------------------------

// ArrayType represents an array type.
//
// References:
//    http://llvm.org/docs/LangRef.html#array-type
type ArrayType struct {
	// Element type.
	Elem Type
	// Array length.
	Len int64
}

// --- [ struct ] --------------------------------------------------------------

// StructType represents a struct type.
//
// References:
//    http://llvm.org/docs/LangRef.html#structure-type
type StructType struct {
	// Struct fields.
	Fields []Type
}

// isType ensures that only types can be assigned to the ast.Type interface.
func (*FuncType) isType()    {}
func (*PointerType) isType() {}
func (*VectorType) isType()  {}
func (*ArrayType) isType()   {}
func (*StructType) isType()  {}
