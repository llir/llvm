package ir

import "github.com/llir/l/ir/instruction"

type BasicBlock struct {
	// Instructions of the basic block.
	Insts []instruction.Instruction
	// Terminator of the basic block.
	Term instruction.Terminator
}
