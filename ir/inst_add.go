package ir

import (
	"fmt"

	"github.com/llir/llvm/ir/types"
	"github.com/llir/llvm/ir/value"
)

// InstAdd represents an add instruction.
type InstAdd struct {
	// Parent basic block.
	parent *BasicBlock
	// Local variable name storing the result of the instruction.
	name string
	// Operands.
	x, y value.Value
}

// NewAdd returns a new add instruction based on the given operands.
func NewAdd(x, y value.Value) *InstAdd {
	return &InstAdd{x: x, y: y}
}

// Type returns the type of the instruction.
func (i *InstAdd) Type() types.Type {
	return i.x.Type()
}

// Ident returns the identifier associated with the instruction.
func (i *InstAdd) Ident() string {
	// TODO: Encode name if containing special characters.
	return "%" + i.name
}

// LLVMString returns the LLVM syntax representation of the instruction.
func (i *InstAdd) LLVMString() string {
	return fmt.Sprintf("%v = add %v %v, %v", i.Ident(), i.Type().LLVMString(), i.x.Ident(), i.y.Ident())
}

// Parent returns the parent basic block of the instruction.
func (i *InstAdd) Parent() *BasicBlock {
	return i.parent
}

// SetParent sets the parent basic block of the instruction.
func (i *InstAdd) SetParent(parent *BasicBlock) {
	i.parent = parent
}

// SetName sets the name of the local variable storing the result of the
// instruction.
func (i *InstAdd) SetName(name string) {
	i.name = name
}
