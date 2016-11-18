package instruction

import (
	"fmt"

	"github.com/llir/llvm/ir/types"
	"github.com/llir/llvm/ir/value"
)

// Ret represents a return instruction.
type Ret struct {
	// Parent basic block.
	parent value.Value
	// Local variable name storing the result of the instruction.
	name string
	// Return value; or nil if "void" return.
	x value.Value
}

// TODO: Consider changing the signature of NewRet to
//
//    func NewRet(x ...value.Value) *Ret
//
// thus making the value optional.
//
// If taking this route, evalute other aspects of the API which may use the same
// method for optional arguments (e.g. new global variable with or without
// initializer).

// NewRet returns a new ret instruction based on the given return value. A nil
// return value indicates a "void" return instruction.
func NewRet(x value.Value) *Ret {
	return &Ret{x: x}
}

// Type returns the type of the instruction.
func (i *Ret) Type() types.Type {
	if i.x != nil {
		return i.x.Type()
	}
	return types.Void
}

// Ident returns the identifier associated with the instruction.
func (i *Ret) Ident() string {
	// TODO: Encode name if containing special characters.
	return "%" + i.name
}

// LLVMString returns the LLVM syntax representation of the instruction.
func (i *Ret) LLVMString() string {
	if i.x != nil {
		return fmt.Sprintf("ret %v %v", i.x.Type().LLVMString(), i.x.Ident())
	}
	return "ret void"
}

// Successors returns the successor basic blocks of the terminator.
func (i *Ret) Successors() []value.Value {
	// Return instructions have no successors.
	return nil
}

// Parent returns the parent basic block of the instruction.
func (i *Ret) Parent() value.Value {
	return i.parent
}

// SetParent sets the parent basic block of the instruction.
func (i *Ret) SetParent(parent value.Value) {
	i.parent = parent
}

// SetName sets the name of the local variable storing the result of the
// instruction.
func (i *Ret) SetName(name string) {
	i.name = name
}
