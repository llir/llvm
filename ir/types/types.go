// Package types declares the data types of LLVM IR.
package types

import (
	"bytes"
	"fmt"
)

// A Type represents an LLVM IR type.
//
// Type may have one of the following underlying types.
//
//   TODO
type Type interface {
	fmt.Stringer
	// Equal reports whether t and u are of equal type.
	Equal(u Type) bool
}

// FuncType represents a function type.
type FuncType struct {
	ret      Type
	params   []*Param
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
	fmt.Fprintf(buf, "%s (", t.ret)
	for i, param := range t.params {
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
		return true
	}
	return false
}

// RetType returns the return type of the function type.
func (t *FuncType) RetType() Type {
	return t.ret
}

// Params returns the parameter types of the function type.
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

// LabelType represents a label type, which is used for basic block values.
type LabelType struct {
}

// String returns the LLVM syntax representation of the type.
func (t *LabelType) String() string {
	return "label"
}

// Equal reports whether t and u are of equal type.
func (t *LabelType) Equal(u Type) bool {
	_, ok := u.(*LabelType)
	return ok
}

// IntType represents an integer type.
type IntType struct {
	// Bit size.
	bits int
}

// NewInt returns a new integer type based on the given bit size.
func NewInt(bits int) *IntType {
	return &IntType{bits: bits}
}

// String returns the LLVM syntax representation of the type.
func (t *IntType) String() string {
	return fmt.Sprintf("i%d", t.bits)
}

// Equal reports whether t and u are of equal type.
func (t *IntType) Equal(u Type) bool {
	if u, ok := u.(*IntType); ok {
		return t.bits == u.bits
	}
	return false
}

// Bits returns the bit size of the integer type.
func (t *IntType) Bits() int {
	return t.bits
}

// PointerType represents a pointer type.
type PointerType struct {
	// Element type.
	elem Type
}

// NewPointer returns a new pointer type based on the given element type.
func NewPointer(elem Type) *PointerType {
	return &PointerType{elem: elem}
}

// String returns the LLVM syntax representation of the type.
func (t *PointerType) String() string {
	return fmt.Sprintf("%s*", t.elem)
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

// VoidType represents a void type.
type VoidType struct {
}

// String returns the LLVM syntax representation of the type.
func (t *VoidType) String() string {
	return "void"
}

// Equal reports whether t and u are of equal type.
func (t *VoidType) Equal(u Type) bool {
	_, ok := u.(*VoidType)
	return ok
}

// VectorType represents a vector type.
type VectorType struct {
	elem Type
	len  int
}

// NewVector returns a new vector type based on the given element type and
// vector length.
func NewVector(elem Type, len int) *VectorType {
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
func (t *VectorType) Len() int {
	return t.len
}

// ArrayType represents an array type.
type ArrayType struct {
	elem Type
	len  int
}

// NewArray returns a new array type based on the given element type and array
// length.
func NewArray(elem Type, len int) *ArrayType {
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
func (t *ArrayType) Len() int {
	return t.len
}

// StructType represents a struct type.
type StructType struct {
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
	for i, field := range t.fields {
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

var (
	Label = &LabelType{}
	Void  = &VoidType{}
	I1    = NewInt(1)
	I8    = NewInt(8)
	I16   = NewInt(16)
	I32   = NewInt(32)
	I64   = NewInt(64)
)
