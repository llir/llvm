package ir

import (
	"bytes"
	"fmt"

	"github.com/llir/llvm/ir/instruction"
	"github.com/llir/llvm/ir/types"
	"github.com/llir/llvm/ir/value"
)

// A BasicBlock represents an LLVM IR basic block, which consists of a sequence
// of non-branching instructions, terminated by a control flow instruction (e.g.
// br or ret).
//
// Basic blocks may be referenced from terminator instructions (e.g. br), and
// are thus considered LLVM IR values of label type.
type BasicBlock struct {
	// Parent function of the basic block.
	parent *Function
	// Label name of the basic block.
	name string
	// Non-branching instructions of the basic block.
	insts []instruction.Instruction
	// Terminator instruction of the basic block.
	term instruction.Terminator
}

// Type returns the type of the basic block.
func (block *BasicBlock) Type() types.Type {
	return types.Label
}

// Ident returns the identifier associated with the basic block.
func (block *BasicBlock) Ident() string {
	// TODO: Encode name if containing special characters.
	return "%" + block.name
}

// LLVMString returns the LLVM syntax representation of the basic block.
func (block *BasicBlock) LLVMString() string {
	// TODO: Encode name if containing special characters.
	buf := &bytes.Buffer{}
	fmt.Fprintf(buf, "%s:\n", block.name)
	for _, inst := range block.insts {
		fmt.Fprintf(buf, "\t%v\n", inst)
	}
	fmt.Fprintf(buf, "\t%v\n", block.term)
	return buf.String()
}

// Append appends the given instruction to the basic block.
func (block *BasicBlock) Append(inst instruction.Instruction) {
	block.insts = append(block.insts, inst)
}

// AppendAdd appends an add instruction to the basic block.
func (block *BasicBlock) AppendAdd(x, y value.Value) *instruction.Add {
	add := instruction.NewAdd(x, y)
	add.SetParent(block)
	block.Append(add)
	return add
}

// SetParent sets the parent function of the basic block.
func (block *BasicBlock) SetParent(parent value.Value) {
	block.parent = parent
}
