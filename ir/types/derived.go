// === [ Derived types ] =======================================================

package types

import (
	"bytes"
	"fmt"
)

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

// NewFunc returns a new function type based on the given return type and
// parameters.
func NewFunc(ret Type, params ...*Param) *FuncType {
	return &FuncType{Ret: ret, Params: params}
}

// String returns the LLVM syntax representation of the type.
func (t *FuncType) String() string {
	buf := &bytes.Buffer{}
	fmt.Fprintf(buf, "%s (", t.Ret)
	for i, param := range t.Params {
		if i != 0 {
			buf.WriteString(", ")
		}
		buf.WriteString(param.Typ.String())
	}
	if t.Variadic {
		if len(t.Params) > 0 {
			buf.WriteString(", ")
		}
		buf.WriteString("...")
	}
	buf.WriteString(")")
	return buf.String()
}

// Equal reports whether t and u are of equal type.
func (t *FuncType) Equal(u Type) bool {
	if u, ok := u.(*FuncType); ok {
		if !t.Ret.Equal(u.Ret) {
			return false
		}
		if len(t.Params) != len(u.Params) {
			return false
		}
		for i, tp := range t.Params {
			up := u.Params[i]
			if !tp.Typ.Equal(up.Typ) {
				return false
			}
		}
		return t.Variadic == u.Variadic
	}
	return false
}

// NewParam appends a new function parameter to the function type based on the
// given parameter name and type.
func (t *FuncType) NewParam(name string, typ Type) *Param {
	param := NewParam(name, typ)
	t.Params = append(t.Params, param)
	return param
}

// A Param represents an LLVM IR function parameter.
type Param struct {
	// Parameter name.
	Name string
	// Parameter type.
	Typ Type
}

// NewParam returns a new function parameter based on the given parameter name
// and type.
func NewParam(name string, typ Type) *Param {
	return &Param{Name: name, Typ: typ}
}

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

// NewPointer returns a new pointer type based on the given element type.
func NewPointer(elem Type) *PointerType {
	return &PointerType{Elem: elem}
}

// String returns the LLVM syntax representation of the type.
func (t *PointerType) String() string {
	if t.AddrSpace != 0 {
		return fmt.Sprintf("%s addrspace(%d)*", t.Elem, t.AddrSpace)
	}
	return fmt.Sprintf("%s*", t.Elem)
}

// Equal reports whether t and u are of equal type.
func (t *PointerType) Equal(u Type) bool {
	if u, ok := u.(*PointerType); ok {
		return t.Elem.Equal(u.Elem) && t.AddrSpace == u.AddrSpace
	}
	return false
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

// NewVector returns a new vector type based on the given element type and
// vector length.
func NewVector(elem Type, len int64) *VectorType {
	return &VectorType{Elem: elem, Len: len}
}

// String returns the LLVM syntax representation of the type.
func (t *VectorType) String() string {
	return fmt.Sprintf("<%d x %s>",
		t.Len,
		t.Elem)
}

// Equal reports whether t and u are of equal type.
func (t *VectorType) Equal(u Type) bool {
	if u, ok := u.(*VectorType); ok {
		return t.Elem.Equal(u.Elem) && t.Len == u.Len
	}
	return false
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

// NewArray returns a new array type based on the given element type and array
// length.
func NewArray(elem Type, len int64) *ArrayType {
	return &ArrayType{Elem: elem, Len: len}
}

// String returns the LLVM syntax representation of the type.
func (t *ArrayType) String() string {
	return fmt.Sprintf("[%d x %s]",
		t.Len,
		t.Elem)
}

// Equal reports whether t and u are of equal type.
func (t *ArrayType) Equal(u Type) bool {
	if u, ok := u.(*ArrayType); ok {
		return t.Elem.Equal(u.Elem) && t.Len == u.Len
	}
	return false
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

// NewStruct returns a new struct type based on the given struct fields.
func NewStruct(fields ...Type) *StructType {
	return &StructType{Fields: fields}
}

// String returns the LLVM syntax representation of the type.
func (t *StructType) String() string {
	buf := &bytes.Buffer{}
	buf.WriteString("{")
	for i, field := range t.Fields {
		if i != 0 {
			buf.WriteString(", ")
		}
		buf.WriteString(field.String())
	}
	buf.WriteString("}")
	return buf.String()
}

// Equal reports whether t and u are of equal type.
func (t *StructType) Equal(u Type) bool {
	if u, ok := u.(*StructType); ok {
		if len(t.Fields) != len(u.Fields) {
			return false
		}
		for i, tf := range t.Fields {
			uf := u.Fields[i]
			if !tf.Equal(uf) {
				return false
			}
		}
		return true
	}
	return false
}
