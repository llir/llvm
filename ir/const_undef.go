package ir

import (
	"fmt"

	"github.com/llir/l/ir/types"
)

// --- [ Undefined values ] ----------------------------------------------------

// ConstUndef is an LLVM IR undefined value.
type ConstUndef struct {
	// Constant type.
	Typ types.Type
}

// NewUndef returns a new undefined value based on the given type.
func NewUndef(typ types.Type) *ConstUndef {
	return &ConstUndef{Typ: typ}
}

// String returns the LLVM syntax representation of the constant as a type-value
// pair.
func (c *ConstUndef) String() string {
	return fmt.Sprintf("%v %v", c.Type(), c.Ident())
}

// Type returns the type of the constant.
func (c *ConstUndef) Type() types.Type {
	return c.Typ
}

// Ident returns the identifier associated with the constant.
func (*ConstUndef) Ident() string {
	return "undef"
}
