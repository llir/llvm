package ast

// --- [ void ] ----------------------------------------------------------------

// VoidType represents a void type.
//
// References:
//    http://llvm.org/docs/LangRef.html#void-type
type VoidType struct {
}

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

// --- [ label ] ---------------------------------------------------------------

// LabelType represents a label type, which is used for basic block values.
//
// References:
//    http://llvm.org/docs/LangRef.html#label-type
type LabelType struct {
}

// --- [ metadata ] ------------------------------------------------------------

// MetadataType represents a metadata type.
//
// References:
//    http://llvm.org/docs/LangRef.html#metadata-type
type MetadataType struct {
}

// isValue ensures that only values can be assigned to the ast.Value interface.
func (*Param) isValue() {}

// isType ensures that only types can be assigned to the ast.Type interface.
func (*VoidType) isType()     {}
func (*FuncType) isType()     {}
func (*LabelType) isType()    {}
func (*MetadataType) isType() {}
