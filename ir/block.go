package ir

import (
	"bytes"
	"fmt"

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
	insts []Instruction
	// Terminator instruction of the basic block.
	term Terminator
}

// Type returns the type of the basic block.
func (b *BasicBlock) Type() types.Type {
	return types.Label
}

// Ident returns the identifier associated with the basic block.
func (b *BasicBlock) Ident() string {
	// TODO: Encode name if containing special characters.
	return "%" + b.name
}

// LLVMString returns the LLVM syntax representation of the basic block.
func (b *BasicBlock) LLVMString() string {
	// TODO: Encode name if containing special characters.
	buf := &bytes.Buffer{}
	fmt.Fprintf(buf, "%s:\n", b.name)
	for _, inst := range b.insts {
		fmt.Fprintf(buf, "\t%v\n", inst.LLVMString())
	}
	fmt.Fprintf(buf, "\t%v\n", b.term.LLVMString())
	return buf.String()
}

// NewAdd appends a new add instruction to the basic block based on the given
// operands.
func (b *BasicBlock) NewAdd(x, y value.Value) *InstAdd {
	inst := NewAdd(x, y)
	inst.SetParent(b)
	b.insts = append(b.insts, inst)
	return inst
}

// NewCall appends a new call instruction to the basic block based on the given
// callee and function arguments.
func (b *BasicBlock) NewCall(callee *Function, args ...value.Value) *InstCall {
	inst := NewCall(callee, args...)
	inst.SetParent(b)
	b.insts = append(b.insts, inst)
	return inst
}

// NewMul appends a new mul instruction to the basic block based on the given
// operands.
func (b *BasicBlock) NewMul(x, y value.Value) *InstMul {
	inst := NewMul(x, y)
	inst.SetParent(b)
	b.insts = append(b.insts, inst)
	return inst
}

// NewLoad appends a new load instruction to the basic block based on the given
// source address.
func (b *BasicBlock) NewLoad(src value.Value) *InstLoad {
	inst := NewLoad(src)
	inst.SetParent(b)
	b.insts = append(b.insts, inst)
	return inst
}

// NewRet appends a new ret instruction to the basic block based on the given
// return value. A nil return value indicates a "void" return instruction.
func (b *BasicBlock) NewRet(x value.Value) *TermRet {
	term := NewRet(x)
	term.SetParent(b)
	b.term = term
	return term
}

// Parent returns the parent function of the basic block.
func (b *BasicBlock) Parent() *Function {
	return b.parent
}

// SetParent sets the parent function of the basic block.
func (b *BasicBlock) SetParent(parent *Function) {
	b.parent = parent
}
