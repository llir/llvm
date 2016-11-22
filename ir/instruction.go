package ir

// TODO: Add remaining underlying instruction types.

// An Instruction represents a non-branching LLVM IR instruction.
//
// Instructions which produce results may be referenced from other instructions,
// and are thus considered LLVM IR values. Note, not all instructions produce
// results (e.g. store).
//
// Instruction may have one of the following underlying types.
//
//    TODO
type Instruction interface {
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
