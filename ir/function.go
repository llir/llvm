package ir

import "github.com/llir/llvm/ir/types"

// A Function represents an LLVM IR function definition or external function
// declaration. The body of a function definition consists of a set of basic
// blocks, interconnected by control flow instructions.
//
// Functions may be referenced from terminator instructions (e.g. call), and are
// thus considered LLVM IR values of function type.
type Function struct {
	// Parent module of the function.
	parent *Module
	// Function name.
	name string
	// Function parameters.
	params []*Param
	// Function type.
	typ *types.FuncType
	// Basic blocks of the function.
	blocks []*BasicBlock
}

// NewFunction returns a new LLVM IR function based on the given name.
func NewFunction(name string, ret types.Type, params ...Param) *Function {
	return &Function{name: name}
}

// Type returns the type of the function.
func (f *Function) Type() types.Type {
	return f.typ
}

// Ident returns the identifier associated with the function.
func (f *Function) Ident() string {
	// TODO: Encode name if containing special characters.
	return "@" + f.name
}

// LLVMString returns the LLVM syntax representation of the function.
func (f *Function) LLVMString() string {
	panic("not yet implemented")
}

// A Param represents a function parameter.
type Param struct {
	// Parameter name.
	name string
	// Parameter type.
	typ types.Type
}
