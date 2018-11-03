package ir

import (
	"fmt"
	"strings"

	"github.com/llir/llvm/ir/types"
)

// --- [ Struct constants ] ----------------------------------------------------

// ConstStruct is an LLVM IR struct constant.
type ConstStruct struct {
	// Struct type.
	Typ *types.StructType
	// Struct fields.
	Fields []Constant
}

// NewStruct returns a new struct constant based on the given struct type and
// fields.
func NewStruct(typ *types.StructType, fields ...Constant) *ConstStruct {
	return &ConstStruct{Typ: typ, Fields: fields}
}

// String returns the LLVM syntax representation of the constant as a type-value
// pair.
func (c *ConstStruct) String() string {
	return fmt.Sprintf("%v %v", c.Type(), c.Ident())
}

// Type returns the type of the constant.
func (c *ConstStruct) Type() types.Type {
	return c.Typ
}

// Ident returns the identifier associated with the constant.
func (c *ConstStruct) Ident() string {
	// "{" Elems "}"
	// "<" "{" Elems "}" ">"
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
