package constant

import (
	"fmt"
	"strings"

	"github.com/llir/llvm/internal/enc"
	"github.com/llir/llvm/ir/types"
)

// --- [ Array constants ] -----------------------------------------------------

// Array is an LLVM IR array constant.
type Array struct {
	// Array type.
	Typ *types.ArrayType
	// Array elements.
	Elems []Constant
}

// NewArray returns a new array constant based on the given array type and
// elements.
func NewArray(elems ...Constant) *Array {
	c := &Array{Elems: elems}
	// Compute type.
	c.Type()
	return c
}

// String returns the LLVM syntax representation of the constant as a type-value
// pair.
func (c *Array) String() string {
	return fmt.Sprintf("%s %s", c.Type(), c.Ident())
}

// Type returns the type of the constant.
func (c *Array) Type() types.Type {
	// Cache type if not present.
	if c.Typ == nil {
		elemType := c.Elems[0].Type()
		c.Typ = types.NewArray(int64(len(c.Elems)), elemType)
	}
	return c.Typ
}

// Ident returns the identifier associated with the constant.
func (c *Array) Ident() string {
	// '[' Elems=(TypeConst separator ',')* ']'
	buf := &strings.Builder{}
	buf.WriteString("[")
	for i, elem := range c.Elems {
		if i != 0 {
			buf.WriteString(", ")
		}
		buf.WriteString(elem.String())
	}
	buf.WriteString("]")
	return buf.String()
}

// --- [ Character array constants ] -------------------------------------------

// CharArray is an LLVM IR character array constant.
type CharArray struct {
	// Array type.
	Typ *types.ArrayType
	// Character array contents.
	X []byte
}

// NewCharArray returns a new character array constant based on the given
// character array contents.
func NewCharArray(x []byte) *CharArray {
	typ := types.NewArray(int64(len(x)), types.I8)
	return &CharArray{Typ: typ, X: x}
}

// NewCharArrayFromString returns a new character array constant based on the
// given UTF-8 string contents.
func NewCharArrayFromString(s string) *CharArray {
	return NewCharArray([]byte(s))
}

// String returns the LLVM syntax representation of the constant as a type-value
// pair.
func (c *CharArray) String() string {
	return fmt.Sprintf("%s %s", c.Type(), c.Ident())
}

// Type returns the type of the constant.
func (c *CharArray) Type() types.Type {
	return c.Typ
}

// Ident returns the identifier associated with the constant.
func (c *CharArray) Ident() string {
	// 'c' Val=StringLit
	return "c" + enc.Quote(c.X)
}
