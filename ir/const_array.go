package ir

import "github.com/llir/l/ir/types"

// --- [ Array constants ] -----------------------------------------------------

// ConstArray is an array constant.
type ConstArray struct {
	// ConstArray elements.
	Elems []Constant
}

// NewArray returns a new array constant based on the given array elements.
func NewArray(elems ...Constant) *ConstArray {
	return &ConstArray{Elems: elems}
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

// ConstCharArray is a character array constant.
type ConstCharArray struct {
	// Character array contents.
	X string // TODO: check if LLVM IR strings are UTF-8 encoded. If not, change type to []byte.
}

// NewCharArray returns a new character array constant based on the given string
// contents.
func NewCharArray(x string) *ConstCharArray {
	return &ConstCharArray{X: x}
}

// Type returns the type of the constant.
func (c *ConstCharArray) Type() types.Type {
	panic("not yet implemented")
}

// Ident returns the identifier associated with the constant.
func (c *ConstCharArray) Ident() string {
	panic("not yet implemented")
}
