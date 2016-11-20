package ir

import (
	"fmt"

	"github.com/llir/llvm/ir/constant"
	"github.com/llir/llvm/ir/types"
)

// A Global represents an LLVM IR global variable definition or external global
// variable declaration.
//
// Global variables always define a pointer to their "content" type because they
// describe a region of memory, and all memory objects in LLVM are accessed
// through pointers.
//
// Global variables may be referenced from instructions (e.g. load), and are
// thus considered LLVM IR values of pointer type.
type Global struct {
	// Global variable name.
	name string
	// Content type.
	content types.Type
	// Initial value; or nil if defined externally.
	init constant.Constant
	// Global variable type.
	typ *types.PointerType
	// Immutability of the global variable.
	immutable bool
}

// NewGlobalDecl returns a new external global variable declaration based on the
// given global variable name and content type.
func NewGlobalDecl(name string, content types.Type) *Global {
	typ := types.NewPointer(content)
	return &Global{name: name, content: content, typ: typ}
}

// NewGlobalDef returns a new global variable definition based on the given
// global variable name and initial value.
func NewGlobalDef(name string, init constant.Constant) *Global {
	content := init.Type()
	typ := types.NewPointer(content)
	return &Global{name: name, content: content, init: init, typ: typ}
}

// Type returns the type of the global variable.
func (g *Global) Type() types.Type {
	return g.typ
}

// Ident returns the identifier associated with the global variable.
func (g *Global) Ident() string {
	// TODO: Encode name if containing special characters.
	return "@" + g.name
}

// LLVMString returns the LLVM syntax representation of the global variable.
func (g *Global) LLVMString() string {
	imm := "global"
	if g.Immutable() {
		imm = "constant"
	}
	content := g.ContentType()
	if init, ok := g.Init(); ok {
		// Global variable definition.
		return fmt.Sprintf("%s = %s %s %s",
			g.Ident(),
			imm,
			content.LLVMString(),
			init.Ident())
	}
	// External global variable declaration.
	return fmt.Sprintf("%s = external %s %s",
		g.Ident(),
		imm,
		content.LLVMString())
}

// ContentType returns the content type of the global variable.
func (g *Global) ContentType() types.Type {
	return g.content
}

// Init returns the initial value of the global variable and a boolean
// indicating if an initializer was present.
func (g *Global) Init() (constant.Constant, bool) {
	if g.init != nil {
		return g.init, true
	}
	return nil, false
}

// Immutable reports whether the global variable is immutable.
func (g *Global) Immutable() bool {
	return g.immutable
}

// SetImmutable sets the immutability of the global variable.
func (g *Global) SetImmutable(immutable bool) {
	g.immutable = immutable
}
