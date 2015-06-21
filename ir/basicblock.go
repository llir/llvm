package ir

import "github.com/llir/llvm/instruction"

// A BasicBlock is a sequence of non-branching instructions, terminated by a
// control flow instruction (such as br or ret).
//
// Basic blocks are values since they can be referenced from instructions (such
// as br and switch). The type of a basic block is label.
//
// References:
//    http://llvm.org/docs/LangRef.html#terminators
type BasicBlock struct {
	// Basic block label name.
	Name string
	// Parent function of the basic block.
	Parent *Function
	// Non-terminator instructions of the basic block.
	Insts []instruction.Instruction
	// Terminator instruction of the basic block.
	Term instruction.Terminator
}
