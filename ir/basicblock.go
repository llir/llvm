package ir

// A BasicBlock is a sequence of non-branching instructions, terminated by a
// control flow instruction (such as br or ret) [1].
//
//    [1]: http://llvm.org/docs/LangRef.html#terminators
type BasicBlock struct {
	// Parent function of the basic block.
	Parent *Function
	// Instructions.
	Insts []Instruction
}
