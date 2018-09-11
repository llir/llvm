package ir

// BasicBlock is an LLVM IR basic block.
type BasicBlock struct {
	// Instructions of the basic block.
	Insts []Instruction
	// Terminator of the basic block.
	Term Terminator
}
