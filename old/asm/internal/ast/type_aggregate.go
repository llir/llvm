package ast

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
	// Opaque struct type.
	//
	// References:
	//    http://llvm.org/docs/LangRef.html#opaque-structure-types
	Opaque bool
}

// isType ensures that only types can be assigned to the ast.Type interface.
func (*ArrayType) isType()  {}
func (*StructType) isType() {}
