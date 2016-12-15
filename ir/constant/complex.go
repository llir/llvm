// === [ Complex constants ] ===================================================
//
// References:
//    http://llvm.org/docs/LangRef.html#complex-constants

package constant

import (
	"bytes"
	"fmt"

	"github.com/llir/llvm/internal/enc"
	"github.com/llir/llvm/ir/types"
)

// --- [ vector ] --------------------------------------------------------------

// Vector represents a vector constant.
type Vector struct {
	// Vector type.
	typ *types.VectorType
	// Vector elements.
	elems []Constant
	// Track uses of the value.
	used
}

// NewVector returns a new vector constant based on the given elements.
func NewVector(elems ...Constant) *Vector {
	if len(elems) < 1 {
		panic(fmt.Sprintf("invalid number of vector elements; expected > 0, got %d", len(elems)))
	}
	typ := types.NewVector(elems[0].Type(), int64(len(elems)))
	return &Vector{typ: typ, elems: elems}
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

// Immutable ensures that only constants can be assigned to the
// constant.Constant interface.
func (*Vector) Immutable() {}

// Elems returns the elements of the vector constant.
func (c *Vector) Elems() []Constant {
	return c.elems
}

// --- [ array ] ---------------------------------------------------------------

// Array represents an array constant.
type Array struct {
	// Array type.
	typ *types.ArrayType
	// Array elements.
	elems []Constant
	// Pretty-print as character array.
	charArray bool
	// Track uses of the value.
	used
}

// NewArray returns a new array constant based on the given elements.
func NewArray(elems ...Constant) *Array {
	if len(elems) < 1 {
		panic(fmt.Sprintf("invalid number of array elements; expected > 0, got %d", len(elems)))
	}
	typ := types.NewArray(elems[0].Type(), int64(len(elems)))
	return &Array{typ: typ, elems: elems}
}

// Type returns the type of the constant.
func (c *Array) Type() types.Type {
	return c.typ
}

// Ident returns the string representation of the constant.
func (c *Array) Ident() string {
	// Pretty-print character arrays.
	if c.CharArray() {
		buf := make([]byte, 0, len(c.elems))
		for _, elem := range c.Elems() {
			e, ok := elem.(*Int)
			if !ok {
				panic(fmt.Sprintf("invalid array element type; expected *constant.Int, got %T", elem))
			}
			b := byte(e.Int64())
			buf = append(buf, b)
		}
		return fmt.Sprintf(`c"%s"`, enc.Escape(string(buf)))
	}
	// Print regular arrays.
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

// Immutable ensures that only constants can be assigned to the
// constant.Constant interface.
func (*Array) Immutable() {}

// Elems returns the elements of the array constant.
func (c *Array) Elems() []Constant {
	return c.elems
}

// CharArray reports whether the given array is a character array.
func (c *Array) CharArray() bool {
	return c.charArray
}

// SetCharArray sets the given array to a character array.
func (c *Array) SetCharArray(charArray bool) {
	c.charArray = charArray
}

// --- [ struct ] --------------------------------------------------------------

// Struct represents a struct constant.
type Struct struct {
	// Struct type.
	typ *types.StructType
	// Struct fields.
	fields []Constant
	// Track uses of the value.
	used
}

// NewStruct returns a new struct constant based on the given struct fields.
func NewStruct(fields ...Constant) *Struct {
	var fieldTypes []types.Type
	for _, field := range fields {
		fieldTypes = append(fieldTypes, field.Type())
	}
	typ := types.NewStruct(fieldTypes...)
	return &Struct{typ: typ, fields: fields}
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

// Immutable ensures that only constants can be assigned to the
// constant.Constant interface.
func (*Struct) Immutable() {}

// Fields returns the struct fields of the struct constant.
func (c *Struct) Fields() []Constant {
	return c.fields
}

// --- [ zeroinitializer ] -----------------------------------------------------

// ZeroInitializer represents a zeroinitializer constant.
type ZeroInitializer struct {
	// Constant type.
	typ types.Type
	// Track uses of the value.
	used
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

// Immutable ensures that only constants can be assigned to the
// constant.Constant interface.
func (*ZeroInitializer) Immutable() {}
