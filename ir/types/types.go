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
	// Equal reports whether t and u are of equal type.
	Equal(u Type) bool
	// LLVMString returns the LLVM syntax representation of the type.
	LLVMString() string
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

// LLVMString returns the LLVM syntax representation of the type.
func (t *FuncType) LLVMString() string {
	buf := &bytes.Buffer{}
	fmt.Fprintf(buf, "%s (", t.ret.LLVMString())
	for i, p := range t.params {
		if i != 0 {
			buf.WriteString(", ")
		}
		fmt.Fprintf(buf, "%s", p.Type().LLVMString())
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
func (t *FuncType) AppendParam(p *Param) {
	t.params = append(t.params, p)
}

// NewParam appends a new function parameter to the function type based on the
// given parameter name and type.
func (t *FuncType) NewParam(name string, typ Type) *Param {
	p := NewParam(name, typ)
	t.AppendParam(p)
	return p
}

// LabelType represents a label type, which is used for basic block values.
type LabelType struct {
}

// LLVMString returns the LLVM syntax representation of the type.
func (t *LabelType) LLVMString() string {
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

// LLVMString returns the LLVM syntax representation of the type.
func (t *IntType) LLVMString() string {
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

// LLVMString returns the LLVM syntax representation of the type.
func (t *PointerType) LLVMString() string {
	return fmt.Sprintf("*%s", t.elem.LLVMString())
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

// LLVMString returns the LLVM syntax representation of the type.
func (t *VoidType) LLVMString() string {
	return "void"
}

// Equal reports whether t and u are of equal type.
func (t *VoidType) Equal(u Type) bool {
	_, ok := u.(*VoidType)
	return ok
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
