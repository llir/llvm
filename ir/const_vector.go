package ir

import (
	"fmt"

	"github.com/llir/l/ir/types"
)

// --- [ Vector constants ] ----------------------------------------------------

// ConstVector is an LLVM IR vector constant.
type ConstVector struct {
	// Vector elements.
	Elems []Constant
}

// NewVector returns a new vector constant based on the given vector elements.
func NewVector(elems ...Constant) *ConstVector {
	return &ConstVector{Elems: elems}
}

// String returns the LLVM syntax representation of the constant as a type-value
// pair.
func (c *ConstVector) String() string {
	return fmt.Sprintf("%v %v", c.Type(), c.Ident())
}

// Type returns the type of the constant.
func (c *ConstVector) Type() types.Type {
	panic("not yet implemented")
}

// Ident returns the identifier associated with the constant.
func (c *ConstVector) Ident() string {
	panic("not yet implemented")
}
