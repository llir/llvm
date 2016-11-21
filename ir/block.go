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
// Basic blocks may be referenced from terminators (e.g. br), and are thus
// considered LLVM IR values of label type.
type BasicBlock struct {
	// Parent function of the basic block.
	parent *Function
	// Label name of the basic block; or empty if anonymous basic block.
	name string
	// Non-branching instructions of the basic block.
	insts []Instruction
	// Terminator of the basic block.
	term Terminator
}

// NewBlock returns a new basic block based on the given label name. An empty
// label name indicates an anonymous basic block.
func NewBlock(name string) *BasicBlock {
	return &BasicBlock{name: name}
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
	buf := &bytes.Buffer{}
	// TODO: Encode name if containing special characters.
	fmt.Fprintf(buf, "%s:\n", b.name)
	for _, i := range b.Insts() {
		fmt.Fprintf(buf, "\t%s\n", i.LLVMString())
	}
	fmt.Fprintf(buf, "\t%s", b.Term().LLVMString())
	return buf.String()
}

// Insts returns the non-branching instructions of the basic block.
func (b *BasicBlock) Insts() []Instruction {
	return b.insts
}

// Term returns the terminator of the basic block.
func (b *BasicBlock) Term() Terminator {
	return b.term
}

// SetTerm sets the terminator of the basic block.
func (b *BasicBlock) SetTerm(t Terminator) {
	if t, ok := t.(parentSetter); ok {
		t.SetParent(b)
	}
	b.term = t
}

// Parent returns the parent function of the basic block.
func (b *BasicBlock) Parent() *Function {
	return b.parent
}

// SetParent sets the parent function of the basic block.
func (b *BasicBlock) SetParent(parent *Function) {
	b.parent = parent
}

// AppendInst appends the given instruction to the basic block.
func (b *BasicBlock) AppendInst(i Instruction) {
	if i, ok := i.(parentSetter); ok {
		i.SetParent(b)
	}
	b.insts = append(b.insts, i)
}

// --- [ Binary instructions ] -------------------------------------------------

// NewAdd appends a new add instruction to the basic block based on the given
// operands.
func (b *BasicBlock) NewAdd(x, y value.Value) *InstAdd {
	i := NewAdd(x, y)
	b.AppendInst(i)
	return i
}

// NewMul appends a new mul instruction to the basic block based on the given
// operands.
func (b *BasicBlock) NewMul(x, y value.Value) *InstMul {
	i := NewMul(x, y)
	b.AppendInst(i)
	return i
}

// --- [ Bitwise instructions ] ------------------------------------------------

// --- [ Vector instructions ] -------------------------------------------------

// --- [ Aggregate instructions ] ----------------------------------------------

// --- [ Memory instructions ] -------------------------------------------------

// NewLoad appends a new load instruction to the basic block based on the given
// source address.
func (b *BasicBlock) NewLoad(src value.Value) *InstLoad {
	i := NewLoad(src)
	b.AppendInst(i)
	return i
}

// --- [ Conversion instructions ] ---------------------------------------------

// --- [ Other instructions ] --------------------------------------------------

// NewCall appends a new call instruction to the basic block based on the given
// callee and function arguments.
func (b *BasicBlock) NewCall(callee *Function, args ...value.Value) *InstCall {
	i := NewCall(callee, args...)
	b.AppendInst(i)
	return i
}

// --- [ Terminators ] ---------------------------------------------------------

// NewRet sets the terminator of the basic block to a new ret terminator based
// on the given return value. A nil return value indicates a "void" return.
func (b *BasicBlock) NewRet(x value.Value) *TermRet {
	t := NewRet(x)
	b.SetTerm(t)
	return t
}
