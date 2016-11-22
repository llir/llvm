package ir

import (
	"bytes"
	"fmt"

	"github.com/llir/llvm/ir/types"
	"github.com/llir/llvm/ir/value"
)

// --- [ icmp ] ----------------------------------------------------------------

// --- [ fcmp ] ----------------------------------------------------------------

// --- [ phi ] -----------------------------------------------------------------

// --- [ select ] --------------------------------------------------------------

// --- [ call ] ----------------------------------------------------------------

// InstCall represents a call instruction.
//
// References:
//    http://llvm.org/docs/LangRef.html#call-instruction
type InstCall struct {
	// Parent basic block.
	parent *BasicBlock
	// Identifier associated with the instruction.
	id string
	// Callee.
	callee *Function
	// Function arguments.
	args []value.Value
}

// NewCall returns a new call instruction based on the given callee and function
// arguments.
func NewCall(callee *Function, args ...value.Value) *InstCall {
	return &InstCall{callee: callee, args: args}
}

// Type returns the type of the instruction.
func (i *InstCall) Type() types.Type {
	return i.callee.typ.RetType()
}

// Ident returns the identifier associated with the instruction.
func (i *InstCall) Ident() string {
	return local(i.id)
}

// SetIdent sets the identifier associated with the instruction.
func (i *InstCall) SetIdent(id string) {
	i.id = id
}

// LLVMString returns the LLVM syntax representation of the instruction.
func (i *InstCall) LLVMString() string {
	buf := &bytes.Buffer{}
	typ := i.Type()
	if !typ.Equal(types.Void) {
		fmt.Fprintf(buf, "%s = ", i.Ident())
	}
	fmt.Fprintf(buf, "call %s %s(",
		typ.LLVMString(),
		i.Callee().Ident())
	for i, arg := range i.Args() {
		if i != 0 {
			buf.WriteString(", ")
		}
		fmt.Fprintf(buf, "%s %s",
			arg.Type().LLVMString(),
			arg.Ident())
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

// Callee returns the callee of the call instruction.
func (i *InstCall) Callee() *Function {
	return i.callee
}

// Args returns the function arguments of the call instruction.
func (i *InstCall) Args() []value.Value {
	return i.args
}

// --- [ va_arg ] --------------------------------------------------------------

// --- [ landingpad ] ----------------------------------------------------------

// --- [ catchpad ] ------------------------------------------------------------

// --- [ cleanuppad ] ----------------------------------------------------------
