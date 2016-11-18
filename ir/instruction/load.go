package instruction

import (
	"fmt"

	"github.com/llir/llvm/ir/types"
	"github.com/llir/llvm/ir/value"
)

// Load represents a load instruction.
type Load struct {
	// Parent basic block.
	parent value.Value
	// Local variable name storing the result of the instruction.
	name string
	// Source address.
	src value.Value
	// Result type produced by the instruction.
	typ types.Type
}

// NewLoad returns a new load instruction based on the given source address.
func NewLoad(src value.Value) *Load {
	if typ, ok := src.Type().(*types.PointerType); ok {
		return &Load{src: src, typ: typ.Elem()}
	}
	panic(fmt.Sprintf("invalid source address type; expected *types.PointerType, got %T", src.Type()))
}

// Type returns the type of the instruction.
func (i *Load) Type() types.Type {
	return i.typ
}

// Ident returns the identifier associated with the instruction.
func (i *Load) Ident() string {
	// TODO: Encode name if containing special characters.
	return "%" + i.name
}

// LLVMString returns the LLVM syntax representation of the instruction.
func (i *Load) LLVMString() string {
	return fmt.Sprintf("%v = load %v, %v %v", i.Ident(), i.Type().LLVMString(), i.src.Type().LLVMString(), i.src.Ident())
}

// Parent returns the parent basic block of the instruction.
func (i *Load) Parent() value.Value {
	return i.parent
}

// SetParent sets the parent basic block of the instruction.
func (i *Load) SetParent(parent value.Value) {
	i.parent = parent
}

// SetName sets the name of the local variable storing the result of the
// instruction.
func (i *Load) SetName(name string) {
	i.name = name
}
