package instruction

import (
	"fmt"

	"github.com/llir/llvm/ir/types"
	"github.com/llir/llvm/ir/value"
)

// Mul represents a multiplication instruction.
type Mul struct {
	// Parent basic block.
	parent value.Value
	// Local variable name storing the result of the instruction.
	name string
	// Operands.
	x, y value.Value
}

// NewMul returns a new mul instruction based on the given operands.
func NewMul(x, y value.Value) *Mul {
	return &Mul{x: x, y: y}
}

// Type returns the type of the instruction.
func (i *Mul) Type() types.Type {
	return i.x.Type()
}

// Ident returns the identifier associated with the instruction.
func (i *Mul) Ident() string {
	// TODO: Encode name if containing special characters.
	return "%" + i.name
}

// LLVMString returns the LLVM syntax representation of the instruction.
func (i *Mul) LLVMString() string {
	return fmt.Sprintf("%v = mul %v %v, %v", i.Ident(), i.Type().LLVMString(), i.x.Ident(), i.y.Ident())
}

// Parent returns the parent basic block of the instruction.
func (i *Mul) Parent() value.Value {
	return i.parent
}

// SetParent sets the parent basic block of the instruction.
func (i *Mul) SetParent(parent value.Value) {
	i.parent = parent
}

// SetName sets the name of the local variable storing the result of the
// instruction.
func (i *Mul) SetName(name string) {
	i.name = name
}
