package ir

import (
	"fmt"
	"strings"

	"github.com/llir/l/ir/types"
)

// --- [ Vector constants ] ----------------------------------------------------

// ConstVector is an LLVM IR vector constant.
type ConstVector struct {
	// Vector type.
	Typ *types.VectorType
	// Vector elements.
	Elems []Constant
}

// NewVector returns a new vector constant based on the given vector type and
// elements.
func NewVector(typ *types.VectorType, elems ...Constant) *ConstVector {
	return &ConstVector{Typ: typ, Elems: elems}
}

// String returns the LLVM syntax representation of the constant as a type-value
// pair.
func (c *ConstVector) String() string {
	return fmt.Sprintf("%v %v", c.Type(), c.Ident())
}

// Type returns the type of the constant.
func (c *ConstVector) Type() types.Type {
	return c.Typ
}

// Ident returns the identifier associated with the constant.
func (c *ConstVector) Ident() string {
	// "<" TypeConsts ">"
	buf := &strings.Builder{}
	buf.WriteString("<")
	for i, elem := range c.Elems {
		if i != 0 {
			buf.WriteString(", ")
		}
		buf.WriteString(elem.String())
	}
	buf.WriteString(">")
	return buf.String()
}
