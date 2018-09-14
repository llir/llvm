package ir

import "github.com/llir/l/ir/types"

// --- [ zeroinitializer constants ] -------------------------------------------

// ConstZeroInitializer is an LLVM IR zeroinitializer constant.
type ConstZeroInitializer struct {
	// Constant type.
	Typ types.Type
}

// NewZeroInitializer returns a new zeroinitializer constant based on the given
// type.
func NewZeroInitializer(typ types.Type) *ConstZeroInitializer {
	return &ConstZeroInitializer{Typ: typ}
}

// Type returns the type of the constant.
func (c *ConstZeroInitializer) Type() types.Type {
	return c.Typ
}

// Ident returns the identifier associated with the constant.
func (c *ConstZeroInitializer) Ident() string {
	return "zeroinitializer"
}
