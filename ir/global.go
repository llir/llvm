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
	// Underlying pointer type.
	underlying types.Type
	// Initial value; or nil if defined externally.
	init constant.Constant
	// Immutability of the global variable.
	immutable bool
	// Global variable type.
	typ *types.PointerType
}

// NewGlobal appends a new global variable to the module based on the given
// global variable name, type and optional initial value.
func NewGlobal(name string, underlying types.Type, init ...constant.Constant) *Global {
	typ := types.NewPointer(underlying)
	global := &Global{name: name, underlying: underlying, typ: typ}
	switch len(init) {
	case 0:
		// External global variable declaration with initial value.
	case 1:
		// Global variable definition with initial value.
		global.init = init[0]
	default:
		panic(fmt.Sprintf("invalid number of initializers; expected 0 or 1, got %d", len(init)))
	}
	return global
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

// Underlying returns the underlying type of the global variable.
func (g *Global) Underlying() types.Type {
	return g.underlying
}

// LLVMString returns the LLVM syntax representation of the global variable.
func (g *Global) LLVMString() string {
	imm := "global"
	if g.immutable {
		imm = "constant"
	}
	if g.init != nil {
		return fmt.Sprintf("%v = %v %v %v", g.Ident(), imm, g.underlying.LLVMString(), g.init.Ident())
	}
	return fmt.Sprintf("%v = external %v %v", g.Ident(), imm, g.underlying.LLVMString())
}
