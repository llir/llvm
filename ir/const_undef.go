package ir

import "github.com/llir/l/ir/types"

// ConstUndef is an undefined constant.
type ConstUndef struct {
	// Constant type.
	Typ types.Type
}

// NewUndef returns a new undefined constant based on the given type.
func NewUndef(typ types.Type) *ConstUndef {
	return &ConstUndef{Typ: typ}
}

// Type returns the type of the constant.
func (c *ConstUndef) Type() types.Type {
	return c.Typ
}

// Ident returns the identifier associated with the constant.
func (*ConstUndef) Ident() string {
	return "undef"
}
