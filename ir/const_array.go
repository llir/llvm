package ir

import "github.com/llir/l/ir/types"

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

// TODO: define CharArray struct, or add Char bool to ConstArray type?
