package ir

import (
	"fmt"

	"github.com/llir/llvm/ir/types"
	"github.com/llir/llvm/ir/value"
)

// InstMul represents an mul instruction.
type InstMul struct {
	// Parent basic block.
	parent *BasicBlock
	// Identifier associated with the instruction.
	id string
	// Operands.
	x, y value.Value
}

// NewMul returns a new mul instruction based on the given operands.
func NewMul(x, y value.Value) *InstMul {
	return &InstMul{x: x, y: y}
}

// Type returns the type of the instruction.
func (i *InstMul) Type() types.Type {
	return i.x.Type()
}

// Ident returns the identifier associated with the instruction.
func (i *InstMul) Ident() string {
	return local(i.id)
}

// SetIdent sets the identifier associated with the instruction.
func (i *InstMul) SetIdent(id string) {
	i.id = id
}

// LLVMString returns the LLVM syntax representation of the instruction.
func (i *InstMul) LLVMString() string {
	return fmt.Sprintf("%s = mul %s %s, %s",
		i.Ident(),
		i.Type().LLVMString(),
		i.x.Ident(),
		i.y.Ident())
}

// Parent returns the parent basic block of the instruction.
func (i *InstMul) Parent() *BasicBlock {
	return i.parent
}

// SetParent sets the parent basic block of the instruction.
func (i *InstMul) SetParent(parent *BasicBlock) {
	i.parent = parent
}
