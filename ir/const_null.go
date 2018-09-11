package ir

import "github.com/llir/l/ir/types"

// Null is a null pointer constant.
type Null struct {
	// Pointer type.
	Typ *types.PointerType
}

// NewNull returns a new null pointer constant based on the given pointer type.
func NewNull(typ *types.PointerType) *Null {
	return &Null{Typ: typ}
}

// Type returns the type of the constant.
func (c *Null) Type() types.Type {
	return c.Typ
}

// Ident returns the identifier associated with the constant.
func (*Null) Ident() string {
	return "null"
}
