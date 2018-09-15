package ir

import (
	"fmt"

	"github.com/llir/l/ir/types"
)

// --- [ Array constants ] -----------------------------------------------------

// ConstArray is an LLVM IR array constant.
type ConstArray struct {
	// Array elements.
	Elems []Constant
}

// NewArray returns a new array constant based on the given array elements.
func NewArray(elems ...Constant) *ConstArray {
	return &ConstArray{Elems: elems}
}

// String returns the LLVM syntax representation of the constant as a type-value
// pair.
func (c *ConstArray) String() string {
	return fmt.Sprintf("%v %v", c.Type(), c.Ident())
}

// Type returns the type of the constant.
func (c *ConstArray) Type() types.Type {
	panic("not yet implemented")
}

// Ident returns the identifier associated with the constant.
func (c *ConstArray) Ident() string {
	panic("not yet implemented")
}

// --- [ Character array constants ] -------------------------------------------

// ConstCharArray is an LLVM IR character array constant.
type ConstCharArray struct {
	// Character array contents.
	X string // TODO: check if LLVM IR strings are UTF-8 encoded. If not, change type to []byte.
}

// NewCharArray returns a new character array constant based on the given string
// contents.
func NewCharArray(x string) *ConstCharArray {
	return &ConstCharArray{X: x}
}

// String returns the LLVM syntax representation of the constant as a type-value
// pair.
func (c *ConstCharArray) String() string {
	return fmt.Sprintf("%v %v", c.Type(), c.Ident())
}

// Type returns the type of the constant.
func (c *ConstCharArray) Type() types.Type {
	panic("not yet implemented")
}

// Ident returns the identifier associated with the constant.
func (c *ConstCharArray) Ident() string {
	panic("not yet implemented")
}
