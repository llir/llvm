package constant

import "github.com/llir/l/ir/types"

// ZeroInitializer is a zeroinitializer constant.
type ZeroInitializer struct {
	// Constant type.
	Typ types.Type
}

// NewZeroInitializer returns a new zeroinitializer constant based on the given
// type.
func NewZeroInitializer(typ types.Type) *ZeroInitializer {
	return &ZeroInitializer{Typ: typ}
}

// Type returns the type of the constant.
func (c *ZeroInitializer) Type() types.Type {
	return c.Typ
}

// Ident returns the identifier associated with the constant.
func (c *ZeroInitializer) Ident() string {
	return "zeroinitializer"
}
