package ir

import "github.com/llir/llvm/ir/instruction"

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
	name string
	// Parent function of the basic block.
	parent *Function
	// Non-terminator instructions of the basic block.
	insts []instruction.Instruction
	// Terminator instruction of the basic block.
	term instruction.Terminator
}

// NewBasicBlock returns a new basic block based on the given name, non-
// terminating instructions and terminator.
func NewBasicBlock(name string, insts []instruction.Instruction, term instruction.Terminator) *BasicBlock {
	return &BasicBlock{name: name, insts: insts, term: term}
}

// Name returns the name of the basic block.
func (block *BasicBlock) Name() string {
	return block.name
}

// Parent returns the parent function of the basic block.
func (block *BasicBlock) Parent() *Function {
	return block.parent
}

// TODO: Add SetParent method to BasicBlock?

// Insts returns the non-terminating instructions of the basic block.
func (block *BasicBlock) Insts() []instruction.Instruction {
	return block.insts
}

// TODO: Add AppendInst and SetInsts methods to BasicBlock? Analogously defined
// as for Function.AppendBlock and Function.SetBlocks.

// Term returns the terminator of the basic block.
func (block *BasicBlock) Term() instruction.Terminator {
	return block.term
}
