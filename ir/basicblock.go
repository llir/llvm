package ir

import (
	"bytes"
	"fmt"

	"github.com/llir/llvm/ir/instruction"
)

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
func NewBasicBlock(name string, insts []instruction.Instruction, term instruction.Terminator) (*BasicBlock, error) {
	// TODO: Verify that name is not a local ID. Unnamed basic blocks should be
	// assigned a local ID implicitly by the internal localID counter of the
	// given function rather than explicitly assigned.
	return &BasicBlock{name: name, insts: insts, term: term}, nil
}

// Name returns the name of the basic block.
func (block *BasicBlock) Name() string {
	return block.name
}

// TODO: Add note to SetName not set local IDs explicitly, as these are assigned
// implicitly by the internal localID counter.

// SetName sets the name of the basic block.
func (block *BasicBlock) SetName(name string) {
	block.name = name
}

// Parent returns the parent function of the basic block.
func (block *BasicBlock) Parent() *Function {
	return block.parent
}

// SetParent sets the parent function of the basic block.
func (block *BasicBlock) SetParent(parent *Function) {
	block.parent = parent
}

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

// String returns the string representation of the basic block.
func (block *BasicBlock) String() string {
	buf := new(bytes.Buffer)
	if len(block.Name()) > 0 {
		fmt.Fprintf(buf, "%s:\n", block.Name())
	}
	for _, inst := range block.Insts() {
		fmt.Fprintf(buf, "\t%s\n", inst)
	}
	fmt.Fprintf(buf, "\t%s", block.Term())
	return buf.String()
}
