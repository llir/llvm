package instruction

import "github.com/llir/llvm/ir/value"

// TODO: Figure out how to return a *ir.BasicBlock instead of a value.Value.

// A Terminator represents an LLVM IR terminator.
//
// Terminator may have one of the following underlying types.
//
//    TODO
type Terminator interface {
	// Block returns the parent basic block of the terminator.
	Block() value.Value
	// LLVMString returns the LLVM syntax representation of the terminator.
	LLVMString() string
	// Successors returns the successor basic blocks of the terminator.
	Successors() []value.Value
}
