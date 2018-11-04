package constant

import (
	"fmt"
	"strings"

	"github.com/llir/llvm/ir"
	"github.com/llir/llvm/ir/types"
)

// --- [ Struct constants ] ----------------------------------------------------

// Struct is an LLVM IR struct constant.
type Struct struct {
	// Struct type.
	Typ *types.StructType
	// Struct fields.
	Fields []ir.Constant
}

// NewStruct returns a new struct constant based on the given struct type and
// fields.
func NewStruct(typ *types.StructType, fields ...ir.Constant) *Struct {
	return &Struct{Typ: typ, Fields: fields}
}

// String returns the LLVM syntax representation of the constant as a type-value
// pair.
func (c *Struct) String() string {
	return fmt.Sprintf("%v %v", c.Type(), c.Ident())
}

// Type returns the type of the constant.
func (c *Struct) Type() types.Type {
	return c.Typ
}

// Ident returns the identifier associated with the constant.
func (c *Struct) Ident() string {
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
