package ir

// A BasicBlock is a sequence of non-branching instructions, terminated by a
// control flow instruction (such as br or ret) [1].
//
//    [1]: http://llvm.org/docs/LangRef.html#terminators
type BasicBlock struct {
	// Basic block label name.
	Name string
	// Parent function of the basic block.
	Parent *Function
	// Non-terminator instructions of the basic block.
	Insts []Instruction
	// Terminator instruction of the basic block.
	Term Terminator
}
