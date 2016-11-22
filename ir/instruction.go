package ir

import "github.com/llir/llvm/ir/value"

// TODO: Add remaining underlying instruction types.

// An Instruction represents a non-branching LLVM IR instruction.
//
// Instruction may have one of the following underlying types.
//
//    TODO
type Instruction interface {
	value.Value
	// LLVMString returns the LLVM syntax representation of the instruction.
	LLVMString() string
	// Parent returns the parent basic block of the instruction.
	Parent() *BasicBlock
}

// parentSetter is the interface that wraps the SetParent method of
// instructions and terminators.
type parentSetter interface {
	// SetParent sets the parent basic block of the instruction.
	SetParent(b *BasicBlock)
}
