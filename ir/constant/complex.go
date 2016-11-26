// === [ Complex constants ] ===================================================
//
// References:
//    http://llvm.org/docs/LangRef.html#complex-constants

package constant

import (
	"bytes"
	"fmt"

	"github.com/llir/llvm/ir/types"
)

// --- [ vector ] --------------------------------------------------------------

// Vector represents a vector constant.
type Vector struct {
	// Vector elements.
	elems []Constant
	// Vector type.
	typ *types.VectorType
}

// NewVector returns a new vector constant based on the given elements.
func NewVector(elems ...Constant) *Vector {
	if len(elems) < 1 {
		panic(fmt.Sprintf("invalid number of vector elements; expected > 0, got %d", len(elems)))
	}
	typ := types.NewVector(elems[0].Type(), len(elems))
	return &Vector{elems: elems, typ: typ}
}

// Type returns the type of the constant.
func (c *Vector) Type() types.Type {
	return c.typ
}

// Ident returns the string representation of the constant.
func (c *Vector) Ident() string {
	buf := &bytes.Buffer{}
	buf.WriteString("<")
	for i, elem := range c.Elems() {
		if i != 0 {
			buf.WriteString(", ")
		}
		fmt.Fprintf(buf, "%s %s",
			elem.Type(),
			elem.Ident())
	}
	buf.WriteString(">")
	return buf.String()
}

// Elems returns the elements of the vector constant.
func (c *Vector) Elems() []Constant {
	return c.elems
}

// --- [ array ] ---------------------------------------------------------------

// TODO: Add support for character arrays.

// Array represents an array constant.
type Array struct {
	// Array elements.
	elems []Constant
	// Array type.
	typ *types.ArrayType
}

// NewArray returns a new array constant based on the given elements.
func NewArray(elems ...Constant) *Array {
	if len(elems) < 1 {
		panic(fmt.Sprintf("invalid number of array elements; expected > 0, got %d", len(elems)))
	}
	typ := types.NewArray(elems[0].Type(), len(elems))
	return &Array{elems: elems, typ: typ}
}

// Type returns the type of the constant.
func (c *Array) Type() types.Type {
	return c.typ
}

// Ident returns the string representation of the constant.
func (c *Array) Ident() string {
	buf := &bytes.Buffer{}
	buf.WriteString("[")
	for i, elem := range c.Elems() {
		if i != 0 {
			buf.WriteString(", ")
		}
		fmt.Fprintf(buf, "%s %s",
			elem.Type(),
			elem.Ident())
	}
	buf.WriteString("]")
	return buf.String()
}

// Elems returns the elements of the array constant.
func (c *Array) Elems() []Constant {
	return c.elems
}

// --- [ struct ] --------------------------------------------------------------

// Struct represents a struct constant.
type Struct struct {
	// Struct fields.
	fields []Constant
	// Struct type.
	typ *types.StructType
}

// NewStruct returns a new struct constant based on the given struct fields.
func NewStruct(fields ...Constant) *Struct {
	var fieldTypes []types.Type
	for _, field := range fields {
		fieldTypes = append(fieldTypes, field.Type())
	}
	typ := types.NewStruct(fieldTypes...)
	return &Struct{fields: fields, typ: typ}
}

// Type returns the type of the constant.
func (c *Struct) Type() types.Type {
	return c.typ
}

// Ident returns the string representation of the constant.
func (c *Struct) Ident() string {
	buf := &bytes.Buffer{}
	buf.WriteString("{")
	for i, field := range c.Fields() {
		if i != 0 {
			buf.WriteString(", ")
		}
		fmt.Fprintf(buf, "%s %s",
			field.Type(),
			field.Ident())
	}
	buf.WriteString("}")
	return buf.String()
}

// Fields returns the struct fields of the struct constant.
func (c *Struct) Fields() []Constant {
	return c.fields
}

// --- [ zeroinitializer ] -----------------------------------------------------

// ZeroInitializer represents a zeroinitializer constant.
type ZeroInitializer struct {
	// Constant type.
	typ types.Type
}

// NewZeroInitializer returns a new zeroinitializer constant based on the given
// type.
func NewZeroInitializer(typ types.Type) *ZeroInitializer {
	return &ZeroInitializer{typ: typ}
}

// Type returns the type of the constant.
func (c *ZeroInitializer) Type() types.Type {
	return c.typ
}

// Ident returns the string representation of the constant.
func (c *ZeroInitializer) Ident() string {
	return "zeroinitializer"
}
