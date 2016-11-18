package instruction

import (
	"fmt"

	"github.com/llir/llvm/ir/types"
	"github.com/llir/llvm/ir/value"
)

// Add represents an addition instruction.
type Add struct {
	// Parent basic block.
	parent value.Value
	// Local variable name storing the result of the instruction.
	name string
	// Operands.
	x, y value.Value
}

// NewAdd returns a new add instruction based on the given operands.
func NewAdd(x, y value.Value) *Add {
	return &Add{x: x, y: y}
}

// Type returns the type of the instruction.
func (i *Add) Type() types.Type {
	return i.x.Type()
}

// Ident returns the identifier associated with the instruction.
func (i *Add) Ident() string {
	// TODO: Encode name if containing special characters.
	return "%" + i.name
}

// LLVMString returns the LLVM syntax representation of the instruction.
func (i *Add) LLVMString() string {
	return fmt.Sprintf("add %v %v, %v", i.Type(), i.x.Ident(), i.y.Ident())
}

// Parent returns the parent basic block of the instruction.
func (i *Add) Parent() value.Value {
	return i.parent
}

// SetParent sets the parent basic block of the instruction.
func (i *Add) SetParent(parent value.Value) {
	i.parent = parent
}

// SetName sets the name of the local variable storing the result of the
// instruction.
func (i *Add) SetName(name string) {
	i.name = name
}
