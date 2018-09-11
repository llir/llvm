package ir

import "github.com/llir/l/ir/types"

// ConstVector is a vector constant.
type ConstVector struct {
	// ConstVector elements.
	Elems []Constant
}

// NewVector returns a new vector constant based on the given vector elements.
func NewVector(elems ...Constant) *ConstVector {
	return &ConstVector{Elems: elems}
}

// Type returns the type of the constant.
func (c *ConstVector) Type() types.Type {
	panic("not yet implemented")
}

// Ident returns the identifier associated with the constant.
func (c *ConstVector) Ident() string {
	panic("not yet implemented")
}
