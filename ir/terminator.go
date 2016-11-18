package ir

// A Terminator represents an LLVM IR terminator.
//
// Terminator may have one of the following underlying types.
//
//    TODO
type Terminator interface {
	// Parent returns the parent basic block of the instruction.
	Parent() *BasicBlock
	// LLVMString returns the LLVM syntax representation of the terminator.
	LLVMString() string
	// Successors returns the successor basic blocks of the terminator.
	Successors() []*BasicBlock
}
