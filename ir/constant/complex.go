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
	Typ *types.VectorType
	// Vector elements.
	Elems []Constant
	// Track uses of the value.
	used
}

// NewVector returns a new vector constant based on the given elements.
func NewVector(elems ...Constant) *Vector {
	if len(elems) < 1 {
		panic(fmt.Sprintf("invalid number of vector elements; expected > 0, got %d", len(elems)))
	}
	typ := types.NewVector(elems[0].Type(), int64(len(elems)))
	return &Vector{Typ: typ, Elems: elems}
}

// Type returns the type of the constant.
func (c *Vector) Type() types.Type {
	return c.Typ
}

// Ident returns the string representation of the constant.
func (c *Vector) Ident() string {
	buf := &bytes.Buffer{}
	buf.WriteString("<")
	for i, elem := range c.Elems {
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

// --- [ array ] ---------------------------------------------------------------

// Array represents an array constant.
type Array struct {
	// Array type.
	Typ *types.ArrayType
	// Array elements.
	Elems []Constant
	// Pretty-print as character array.
	CharArray bool
	// Track uses of the value.
	used
}

// NewArray returns a new array constant based on the given elements.
func NewArray(elems ...Constant) *Array {
	if len(elems) < 1 {
		panic(fmt.Sprintf("invalid number of array elements; expected > 0, got %d", len(elems)))
	}
	typ := types.NewArray(elems[0].Type(), int64(len(elems)))
	return &Array{Typ: typ, Elems: elems}
}

// Type returns the type of the constant.
func (c *Array) Type() types.Type {
	return c.Typ
}

// Ident returns the string representation of the constant.
func (c *Array) Ident() string {
	// Pretty-print character arrays.
	if c.CharArray {
		buf := make([]byte, 0, len(c.Elems))
		for _, elem := range c.Elems {
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
	for i, elem := range c.Elems {
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

// --- [ struct ] --------------------------------------------------------------

// Struct represents a struct constant.
type Struct struct {
	// Struct type.
	Typ *types.StructType
	// Struct fields.
	Fields []Constant
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
	return &Struct{Typ: typ, Fields: fields}
}

// Type returns the type of the constant.
func (c *Struct) Type() types.Type {
	return c.Typ
}

// Ident returns the string representation of the constant.
func (c *Struct) Ident() string {
	buf := &bytes.Buffer{}
	buf.WriteString("{")
	for i, field := range c.Fields {
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

// --- [ zeroinitializer ] -----------------------------------------------------

// ZeroInitializer represents a zeroinitializer constant.
type ZeroInitializer struct {
	// Constant type.
	Typ types.Type
	// Track uses of the value.
	used
}

// NewZeroInitializer returns a new zeroinitializer constant based on the given
// type.
func NewZeroInitializer(typ types.Type) *ZeroInitializer {
	return &ZeroInitializer{Typ: typ}
}

// Type returns the type of the constant.
func (c *ZeroInitializer) Type() types.Type {
	return c.Typ
}

// Ident returns the string representation of the constant.
func (c *ZeroInitializer) Ident() string {
	return "zeroinitializer"
}

// Immutable ensures that only constants can be assigned to the
// constant.Constant interface.
func (*ZeroInitializer) Immutable() {}
