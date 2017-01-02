// === [ Aggregate types ] =====================================================
//
// References:
//    http://llvm.org/docs/LangRef.html#aggregate-types

package types

import (
	"bytes"
	"fmt"

	"github.com/llir/llvm/internal/enc"
)

// --- [ array ] ---------------------------------------------------------------

// ArrayType represents an array type.
//
// References:
//    http://llvm.org/docs/LangRef.html#array-type
type ArrayType struct {
	// Type name alias.
	Name string
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
	if len(t.Name) > 0 {
		return enc.Local(t.Name)
	}
	return t.Def()
}

// Def returns the LLVM syntax representation of the definition of the type.
func (t *ArrayType) Def() string {
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

// GetName returns the name of the type.
func (t *ArrayType) GetName() string {
	return t.Name
}

// SetName sets the name of the type.
func (t *ArrayType) SetName(name string) {
	t.Name = name
}

// --- [ struct ] --------------------------------------------------------------

// StructType represents a struct type. Identified struct types are uniqued by
// type names, not by structural identity.
//
// References:
//    http://llvm.org/docs/LangRef.html#structure-type
type StructType struct {
	// Type name of identified struct type; or empty if struct type literal.
	Name string
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
	if t.Identified() {
		return enc.Local(t.Name)
	}
	return t.Def()
}

// Def returns the LLVM syntax representation of the definition of the type.
func (t *StructType) Def() string {
	if t.Opaque {
		return "opaque"
	}
	buf := &bytes.Buffer{}
	buf.WriteString("{")
	if len(t.Fields) > 0 {
		// Use same output format as Clang.
		buf.WriteString(" ")
	}
	for i, field := range t.Fields {
		if i != 0 {
			buf.WriteString(", ")
		}
		buf.WriteString(field.String())
	}
	if len(t.Fields) > 0 {
		// Use same output format as Clang.
		buf.WriteString(" ")
	}
	buf.WriteString("}")
	return buf.String()
}

// Equal reports whether t and u are of equal type.
func (t *StructType) Equal(u Type) bool {
	if u, ok := u.(*StructType); ok {
		// Identified struct types are uniqued by type names, not by structural
		// identity.
		if t.Identified() || u.Identified() {
			return t.Name == u.Name
		}
		// Literal struct types are uniqued by structural identity.
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

// GetName returns the name of the type.
func (t *StructType) GetName() string {
	return t.Name
}

// SetName sets the name of the type.
func (t *StructType) SetName(name string) {
	t.Name = name
}

// Identified reports whether t is an identified struct type.
func (t *StructType) Identified() bool {
	return len(t.Name) > 0
}
