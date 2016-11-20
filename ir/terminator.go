package ir

// A Terminator represents an LLVM IR terminator.
//
// Terminator may have one of the following underlying types.
//
//    TODO
type Terminator interface {
	// LLVMString returns the LLVM syntax representation of the terminator.
	LLVMString() string
	// Parent returns the parent basic block of the terminator.
	Parent() *BasicBlock
	// Successors returns the successor basic blocks of the terminator.
	Successors() []*BasicBlock
}
