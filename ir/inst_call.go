package ir

import (
	"bytes"
	"fmt"

	"github.com/llir/llvm/ir/types"
	"github.com/llir/llvm/ir/value"
)

// InstCall represents a call instruction.
type InstCall struct {
	// Parent basic block.
	parent *BasicBlock
	// Local variable name storing the result of the instruction.
	name string
	// Callee.
	callee *Function
	// Function arguments.
	args []value.Value
	// Result type of the call instruction.
	result types.Type
}

// NewCall returns a new call instruction based on the given callee and function
// arguments.
func NewCall(callee *Function, args ...value.Value) *InstCall {
	return &InstCall{callee: callee, args: args, result: callee.typ.RetType()}
}

// Type returns the type of the instruction.
func (i *InstCall) Type() types.Type {
	return i.result
}

// Ident returns the identifier associated with the instruction.
func (i *InstCall) Ident() string {
	// TODO: Encode name if containing special characters.
	return "%" + i.name
}

// LLVMString returns the LLVM syntax representation of the instruction.
func (i *InstCall) LLVMString() string {
	buf := &bytes.Buffer{}
	if !i.result.Equal(types.Void) {
		fmt.Fprintf(buf, "%v = ", i.Ident())
	}
	fmt.Fprintf(buf, "call %v %v(", i.result.LLVMString(), i.callee.Ident())
	for i, arg := range i.args {
		if i != 0 {
			buf.WriteString(", ")
		}
		fmt.Fprintf(buf, "%v %v", arg.Type().LLVMString(), arg.Ident())
	}
	buf.WriteString(")")
	return buf.String()
}

// Parent returns the parent basic block of the instruction.
func (i *InstCall) Parent() *BasicBlock {
	return i.parent
}

// SetParent sets the parent basic block of the instruction.
func (i *InstCall) SetParent(parent *BasicBlock) {
	i.parent = parent
}

// SetName sets the name of the local variable storing the result of the
// instruction.
func (i *InstCall) SetName(name string) {
	i.name = name
}
