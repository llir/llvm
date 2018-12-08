package constant

import (
	"fmt"
	"strings"

	"github.com/llir/llvm/ir/types"
)

// --- [ Struct constants ] ----------------------------------------------------

// Struct is an LLVM IR struct constant.
type Struct struct {
	// Struct type.
	Typ *types.StructType
	// Struct fields.
	Fields []Constant

	// Struct implements constant.Constant.
	Constant
}

// NewStruct returns a new struct constant based on the given struct type and
// fields.
func NewStruct(fields ...Constant) *Struct {
	c := &Struct{Fields: fields}
	// Compute type.
	c.Type()
	return c
}

// String returns the LLVM syntax representation of the constant as a type-value
// pair.
func (c *Struct) String() string {
	return fmt.Sprintf("%s %s", c.Type(), c.Ident())
}

// Type returns the type of the constant.
func (c *Struct) Type() types.Type {
	// Cache type if not present.
	if c.Typ == nil {
		var fieldTypes []types.Type
		for _, field := range c.Fields {
			fieldTypes = append(fieldTypes, field.Type())
		}
		c.Typ = types.NewStruct(fieldTypes...)
	}
	return c.Typ
}

// Ident returns the identifier associated with the constant.
func (c *Struct) Ident() string {
	// Struct constant.
	//
	//    '{' Fields=(TypeConst separator ',')+? '}'
	//
	// Packed struct constant.
	//
	//    '<' '{' Fields=(TypeConst separator ',')+? '}' '>'
	if len(c.Fields) == 0 {
		if c.Typ.Packed {
			return "<{}>"
		}
		return "{}"
	}
	buf := &strings.Builder{}
	if c.Typ.Packed {
		buf.WriteString("<")
	}
	buf.WriteString("{ ")
	for i, field := range c.Fields {
		if i != 0 {
			buf.WriteString(", ")
		}
		buf.WriteString(field.String())
	}
	buf.WriteString(" }")
	if c.Typ.Packed {
		buf.WriteString(">")
	}
	return buf.String()
}
