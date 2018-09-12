package ir

import "github.com/llir/l/ir/types"

// Global is a global variable declaration or definition.
type Global struct {
	// Global variable name.
	GlobalName string
}

// Type returns the type of the global variable.
func (g *Global) Type() types.Type {
	panic("not yet implemented")
}

// Ident returns the identifier associated with the global variable.
func (g *Global) Ident() string {
	panic("not yet implemented")
}

// Name returns the name of the global variable.
func (g *Global) Name() string {
	return g.GlobalName
}

// SetName sets the name of the global variable.
func (g *Global) SetName(name string) {
	g.GlobalName = name
}
