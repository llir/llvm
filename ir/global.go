package ir

import (
	"fmt"

	"github.com/llir/l/internal/enc"
	"github.com/llir/l/ir/types"
)

// === [ Global variables ] ====================================================

// Global is a global variable declaration or definition.
type Global struct {
	// Global variable name.
	GlobalName string
	// Immutability of global variable (constant or global).
	Immutable bool
	// Content type.
	ContentType types.Type
	// Initial value; or nil if declaration.
	Init Constant
}

// NewGlobalDecl returns a new global variable declaration based on the given
// global variable name and content type.
func NewGlobalDecl(name string, contentType types.Type) *Global {
	return &Global{GlobalName: name, ContentType: contentType}
}

// NewGlobalDef returns a new global variable definition based on the given
// global variable name and initial value.
func NewGlobalDef(name string, init Constant) *Global {
	return &Global{GlobalName: name, ContentType: init.Type(), Init: init}
}

// String returns the LLVM syntax representation of the global variable as a
// type-value pair.
func (g *Global) String() string {
	return fmt.Sprintf("%v %v", g.Type(), g.Ident())
}

// Type returns the type of the global variable.
func (g *Global) Type() types.Type {
	// TODO: cache type?
	return types.NewPointer(g.ContentType)
}

// Ident returns the identifier associated with the global variable.
func (g *Global) Ident() string {
	return enc.Global(g.GlobalName)
}

// Name returns the name of the global variable.
func (g *Global) Name() string {
	return g.GlobalName
}

// SetName sets the name of the global variable.
func (g *Global) SetName(name string) {
	g.GlobalName = name
}

// Def returns the LLVM syntax representation of the global variable definition
// or declaration.
func (g *Global) Def() string {
	panic("not yet implemented")
}
