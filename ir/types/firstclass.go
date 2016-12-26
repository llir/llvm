// === [ First class types ] ===================================================
//
// References:
//    http://llvm.org/docs/LangRef.html#first-class-types

package types

import (
	"bytes"
	"fmt"
)

// --- [ integer ] -------------------------------------------------------------

// IntType represents an integer type.
//
// References:
//    http://llvm.org/docs/LangRef.html#integer-type
type IntType struct {
	// Bit size.
	Size int
}

// NewInt returns a new integer type based on the given bit size.
func NewInt(size int) *IntType {
	return &IntType{Size: size}
}

// String returns the LLVM syntax representation of the type.
func (t *IntType) String() string {
	return fmt.Sprintf("i%d", t.Size)
}

// Equal reports whether t and u are of equal type.
func (t *IntType) Equal(u Type) bool {
	if u, ok := u.(*IntType); ok {
		return t.Size == u.Size
	}
	return false
}

// --- [ floating-point ] ------------------------------------------------------

// FloatType represents a floating-point type.
//
// References:
//    http://llvm.org/docs/LangRef.html#floating-point-types
type FloatType struct {
	// Floating-point kind.
	Kind FloatKind
}

// String returns the LLVM syntax representation of the type.
func (t *FloatType) String() string {
	return t.Kind.String()
}

// Equal reports whether t and u are of equal type.
func (t *FloatType) Equal(u Type) bool {
	if u, ok := u.(*FloatType); ok {
		return t.Kind == u.Kind
	}
	return false
}

// FloatKind represents the set of floating-point kinds.
type FloatKind int

// Floating point kinds.
const (
	FloatKindIEEE_16           FloatKind = iota // half:      16-bit floating point type
	FloatKindIEEE_32                            // float:     32-bit floating point type
	FloatKindIEEE_64                            // double:    64-bit floating point type
	FloatKindIEEE_128                           // fp128:     128-bit floating point type (112-bit mantissa)
	FloatKindDoubleExtended_80                  // x86_fp80:  80-bit floating point type (x87)
	FloatKindDoubleDouble_128                   // ppc_fp128: 128-bit floating point type (two 64-bits, PowerPC)
)

// String returns the LLVM syntax representation of the floating-point kind.
func (kind FloatKind) String() string {
	switch kind {
	case FloatKindIEEE_16:
		return "half"
	case FloatKindIEEE_32:
		return "float"
	case FloatKindIEEE_64:
		return "double"
	case FloatKindIEEE_128:
		return "fp128"
	case FloatKindDoubleExtended_80:
		return "x86_fp80"
	case FloatKindDoubleDouble_128:
		return "ppc_fp128"
	}
	return fmt.Sprintf("<unknown floating-point kind %d>", int(kind))
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

// --- [ label ] ---------------------------------------------------------------

// LabelType represents a label type, which is used for basic block values.
//
// References:
//    http://llvm.org/docs/LangRef.html#label-type
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

// --- [ metadata ] ------------------------------------------------------------

// MetadataType represents a metadata type.
//
// References:
//    http://llvm.org/docs/LangRef.html#metadata-type
type MetadataType struct {
}

// String returns the LLVM syntax representation of the type.
func (t *MetadataType) String() string {
	return "metadata"
}

// Equal reports whether t and u are of equal type.
func (t *MetadataType) Equal(u Type) bool {
	_, ok := u.(*MetadataType)
	return ok
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
