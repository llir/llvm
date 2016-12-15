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
	ret Type
	// Function parameters.
	params []*Param
	// Variadicity of the function type.
	variadic bool
}

// NewFunc returns a new function type based on the given return type and
// parameters.
func NewFunc(ret Type, params ...*Param) *FuncType {
	return &FuncType{ret: ret, params: params}
}

// String returns the LLVM syntax representation of the type.
func (t *FuncType) String() string {
	buf := &bytes.Buffer{}
	fmt.Fprintf(buf, "%s (", t.RetType())
	for i, param := range t.Params() {
		if i != 0 {
			buf.WriteString(", ")
		}
		buf.WriteString(param.Type().String())
	}
	if t.Variadic() {
		if len(t.params) > 0 {
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
		if !t.ret.Equal(u.ret) {
			return false
		}
		if len(t.params) != len(u.params) {
			return false
		}
		for i := range t.params {
			if !t.params[i].Type().Equal(u.params[i].Type()) {
				return false
			}
		}
		return t.variadic == u.variadic
	}
	return false
}

// RetType returns the return type of the function type.
func (t *FuncType) RetType() Type {
	return t.ret
}

// Params returns the function parameters of the function type.
func (t *FuncType) Params() []*Param {
	return t.params
}

// Variadic reports whether the function type is variadic.
func (t *FuncType) Variadic() bool {
	return t.variadic
}

// SetVariadic sets the variadicity of the function type.
func (t *FuncType) SetVariadic(variadic bool) {
	t.variadic = variadic
}

// AppendParam appends the given function parameter to the function type.
func (t *FuncType) AppendParam(param *Param) {
	t.params = append(t.params, param)
}

// NewParam appends a new function parameter to the function type based on the
// given parameter name and type.
func (t *FuncType) NewParam(name string, typ Type) *Param {
	param := NewParam(name, typ)
	t.AppendParam(param)
	return param
}

// A Param represents an LLVM IR function parameter.
type Param struct {
	// Parameter name.
	name string
	// Parameter type.
	typ Type
}

// NewParam returns a new function parameter based on the given parameter name
// and type.
func NewParam(name string, typ Type) *Param {
	return &Param{name: name, typ: typ}
}

// Type returns the type of the function parameter.
func (param *Param) Type() Type {
	return param.typ
}

// Name returns the name of the function parameter.
func (param *Param) Name() string {
	return param.name
}

// SetName sets the name of the function parameter.
func (param *Param) SetName(name string) {
	param.name = name
}

// --- [ pointer ] -------------------------------------------------------------

// PointerType represents a pointer type.
//
// References:
//    http://llvm.org/docs/LangRef.html#pointer-type
type PointerType struct {
	// Element type.
	elem Type
	// Address space.
	space int64
}

// NewPointer returns a new pointer type based on the given element type.
func NewPointer(elem Type) *PointerType {
	return &PointerType{elem: elem}
}

// String returns the LLVM syntax representation of the type.
func (t *PointerType) String() string {
	if t.space != 0 {
		return fmt.Sprintf("%s addrspace(%d)*", t.Elem(), t.AddrSpace())
	}
	return fmt.Sprintf("%s*", t.Elem())
}

// Equal reports whether t and u are of equal type.
func (t *PointerType) Equal(u Type) bool {
	if u, ok := u.(*PointerType); ok {
		return t.elem.Equal(u.elem)
	}
	return false
}

// Elem returns the element type of the pointer type.
func (t *PointerType) Elem() Type {
	return t.elem
}

// AddrSpace returns the address space of the pointer type.
func (t *PointerType) AddrSpace() int64 {
	return t.space
}

// SetAddrSpace sets the address space of the pointer type.
func (t *PointerType) SetAddrSpace(space int64) {
	t.space = space
}

// --- [ vector ] --------------------------------------------------------------

// VectorType represents a vector type.
//
// References:
//    http://llvm.org/docs/LangRef.html#vector-type
type VectorType struct {
	// Element type.
	elem Type
	// Vector length.
	len int64
}

// NewVector returns a new vector type based on the given element type and
// vector length.
func NewVector(elem Type, len int64) *VectorType {
	return &VectorType{elem: elem, len: len}
}

// String returns the LLVM syntax representation of the type.
func (t *VectorType) String() string {
	return fmt.Sprintf("<%d x %s>",
		t.Len(),
		t.Elem())
}

// Equal reports whether t and u are of equal type.
func (t *VectorType) Equal(u Type) bool {
	if u, ok := u.(*VectorType); ok {
		if !t.elem.Equal(u.elem) {
			return false
		}
		return t.len == u.len
	}
	return false
}

// Elem returns the element type of the vector type.
func (t *VectorType) Elem() Type {
	return t.elem
}

// Len returns the length of the vector type.
func (t *VectorType) Len() int64 {
	return t.len
}

// --- [ array ] ---------------------------------------------------------------

// ArrayType represents an array type.
//
// References:
//    http://llvm.org/docs/LangRef.html#array-type
type ArrayType struct {
	// Element type.
	elem Type
	// Array length.
	len int64
}

// NewArray returns a new array type based on the given element type and array
// length.
func NewArray(elem Type, len int64) *ArrayType {
	return &ArrayType{elem: elem, len: len}
}

// String returns the LLVM syntax representation of the type.
func (t *ArrayType) String() string {
	return fmt.Sprintf("[%d x %s]",
		t.Len(),
		t.Elem())
}

// Equal reports whether t and u are of equal type.
func (t *ArrayType) Equal(u Type) bool {
	if u, ok := u.(*ArrayType); ok {
		if !t.elem.Equal(u.elem) {
			return false
		}
		return t.len == u.len
	}
	return false
}

// Elem returns the element type of the array type.
func (t *ArrayType) Elem() Type {
	return t.elem
}

// Len returns the length of the array type.
func (t *ArrayType) Len() int64 {
	return t.len
}

// --- [ struct ] --------------------------------------------------------------

// StructType represents a struct type.
//
// References:
//    http://llvm.org/docs/LangRef.html#structure-type
type StructType struct {
	// Struct fields.
	fields []Type
}

// NewStruct returns a new struct type based on the given struct fields.
func NewStruct(fields ...Type) *StructType {
	return &StructType{fields: fields}
}

// String returns the LLVM syntax representation of the type.
func (t *StructType) String() string {
	buf := &bytes.Buffer{}
	buf.WriteString("{")
	for i, field := range t.Fields() {
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
		if len(t.fields) != len(u.fields) {
			return false
		}
		for i := range t.fields {
			if !t.fields[i].Equal(u.fields[i]) {
				return false
			}
		}
		return true
	}
	return false
}

// Fields returns the struct fields of the struct type.
func (t *StructType) Fields() []Type {
	return t.fields
}
