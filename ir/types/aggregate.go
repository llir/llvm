// === [ Aggregate types ] =====================================================
//
// References:
//    http://llvm.org/docs/LangRef.html#aggregate-types

package types

import (
	"bytes"
	"fmt"
)

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
	switch u := u.(type) {
	case *ArrayType:
		return t.Elem.Equal(u.Elem) && t.Len == u.Len
	case *NamedType:
		return t.Equal(u.Def)
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
	// Opaque struct type.
	//
	// References:
	//    http://llvm.org/docs/LangRef.html#opaque-structure-types
	Opaque bool
}

// NewStruct returns a new struct type based on the given struct fields.
func NewStruct(fields ...Type) *StructType {
	return &StructType{Fields: fields}
}

// String returns the LLVM syntax representation of the type.
func (t *StructType) String() string {
	if t.Opaque {
		return "opaque"
	}
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
	switch u := u.(type) {
	case *StructType:
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
	case *NamedType:
		// t is an an unnamed type and u is a type alias or an identified struct;
		// thus never equal.
		return false
	}
	return false
}
