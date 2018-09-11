package ir

import "github.com/llir/l/ir/types"

// Global is a global variable declaration or definition.
type Global struct {
}

// Type returns the type of the value.
func (g *Global) Type() types.Type {
	panic("not yet implemented")
}

// Ident returns the identifier associated with the value.
func (g *Global) Ident() string {
	panic("not yet implemented")
}

// IsConstant ensures that only constants can be assigned to the
// constant.Constant interface.
func (*Global) IsConstant() {}
