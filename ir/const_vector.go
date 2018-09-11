package ir

import "github.com/llir/l/ir/types"

// Vector is a vector constant.
type Vector struct {
	// Vector elements.
	Elems []Constant
}

// NewVector returns a new vector constant based on the given vector elements.
func NewVector(elems ...Constant) *Vector {
	return &Vector{Elems: elems}
}

// Type returns the type of the constant.
func (c *Vector) Type() types.Type {
	panic("not yet implemented")
}

// Ident returns the identifier associated with the constant.
func (c *Vector) Ident() string {
	panic("not yet implemented")
}
