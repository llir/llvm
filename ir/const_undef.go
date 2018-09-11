package ir

import "github.com/llir/l/ir/types"

// Undef is an undefined constant.
type Undef struct {
	// Constant type.
	Typ types.Type
}

// NewUndef returns a new undefined constant based on the given type.
func NewUndef(typ types.Type) *Undef {
	return &Undef{Typ: typ}
}

// Type returns the type of the constant.
func (c *Undef) Type() types.Type {
	return c.Typ
}

// Ident returns the identifier associated with the constant.
func (*Undef) Ident() string {
	return "undef"
}
