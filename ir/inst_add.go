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
	// Identifier associated with the instruction.
	id string
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
	return local(i.id)
}

// SetIdent sets the identifier associated with the instruction.
func (i *InstAdd) SetIdent(id string) {
	i.id = id
}

// LLVMString returns the LLVM syntax representation of the instruction.
func (i *InstAdd) LLVMString() string {
	return fmt.Sprintf("%s = add %s %s, %s",
		i.Ident(),
		i.Type().LLVMString(),
		i.x.Ident(),
		i.y.Ident())
}

// Parent returns the parent basic block of the instruction.
func (i *InstAdd) Parent() *BasicBlock {
	return i.parent
}

// SetParent sets the parent basic block of the instruction.
func (i *InstAdd) SetParent(parent *BasicBlock) {
	i.parent = parent
}
