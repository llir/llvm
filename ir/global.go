package ir

import (
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
	// TODO: Figure out how to represent the pointer type and the
	// actual/underlying type of a global variable.

	// Global variable type.
	typ types.Type
	// Initial value; or nil if defined externally.
	init constant.Constant
	// Immutability of the global variable.
	immutable bool
}
