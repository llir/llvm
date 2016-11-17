package ir

import (
	"github.com/llir/llvm/ir"
	"github.com/llir/llvm/ir/types"
)

// A Function represents an LLVM IR function definition or external function
// declaration. The body of a function definition consists of a set of basic
// blocks, interconnected by control flow instructions.
//
// Functions may be referenced from terminator instructions (e.g. call), and are
// thus considered LLVM IR values of function type.
type Function struct {
	// Function name.
	name string
	// Function parameters.
	params []*Param
	// Function type.
	typ *types.Func
	// Basic blocks of the function.
	blocks []*BasicBlock
}

// NewFunction returns a new LLVM IR function based on the given name.
func NewFunction(name string) *ir.Function {
	return &ir.Function{name: name}
}

// A Param represents a function parameter.
type Param struct {
	// Parameter name.
	name string
	// Parameter type.
	typ types.Type
}
