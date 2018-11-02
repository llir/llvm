package ir

import (
	"fmt"
	"strings"

	"github.com/llir/l/internal/enc"
	"github.com/llir/l/ir/types"
)

// --- [ Array constants ] -----------------------------------------------------

// ConstArray is an LLVM IR array constant.
type ConstArray struct {
	// Array type.
	Typ *types.ArrayType
	// Array elements.
	Elems []Constant
}

// NewArray returns a new array constant based on the given array type and
// elements.
func NewArray(typ *types.ArrayType, elems ...Constant) *ConstArray {
	return &ConstArray{Typ: typ, Elems: elems}
}

// String returns the LLVM syntax representation of the constant as a type-value
// pair.
func (c *ConstArray) String() string {
	return fmt.Sprintf("%v %v", c.Type(), c.Ident())
}

// Type returns the type of the constant.
func (c *ConstArray) Type() types.Type {
	return c.Typ
}

// Ident returns the identifier associated with the constant.
func (c *ConstArray) Ident() string {
	// "[" TypeConsts "]"
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

// ConstCharArray is an LLVM IR character array constant.
type ConstCharArray struct {
	// Array type.
	Typ *types.ArrayType
	// Character array contents.
	X []byte
}

// NewCharArray returns a new character array constant based on the given
// character array contents.
func NewCharArray(x []byte) *ConstCharArray {
	typ := types.NewArray(int64(len(x)), types.I8)
	return &ConstCharArray{Typ: typ, X: x}
}

// NewCharArrayFromString returns a new character array constant based on the
// given UTF-8 string contents.
func NewCharArrayFromString(s string) *ConstCharArray {
	return NewCharArray([]byte(s))
}

// String returns the LLVM syntax representation of the constant as a type-value
// pair.
func (c *ConstCharArray) String() string {
	return fmt.Sprintf("%v %v", c.Type(), c.Ident())
}

// Type returns the type of the constant.
func (c *ConstCharArray) Type() types.Type {
	return c.Typ
}

// Ident returns the identifier associated with the constant.
func (c *ConstCharArray) Ident() string {
	// "c" StringLit
	return "c" + enc.Quote(c.X)
}
