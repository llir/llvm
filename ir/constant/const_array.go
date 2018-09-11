package constant

import "github.com/llir/l/ir/types"

// Array is an array constant.
type Array struct {
	// Array elements.
	Elems []Constant
}

// NewArray returns a new array constant based on the given array elements.
func NewArray(elems ...Constant) *Array {
	return &Array{Elems: elems}
}

// Type returns the type of the constant.
func (c *Array) Type() types.Type {
	panic("not yet implemented")
}

// Ident returns the identifier associated with the constant.
func (c *Array) Ident() string {
	panic("not yet implemented")
}

// TODO: define CharArray struct, or add Char bool to Array type?
