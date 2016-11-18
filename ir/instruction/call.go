package instruction

import (
	"bytes"
	"fmt"

	"github.com/llir/llvm/ir/types"
	"github.com/llir/llvm/ir/value"
)

// Call represents a call instruction.
type Call struct {
	// Parent basic block.
	parent value.Value
	// Local variable name storing the result of the instruction.
	name string
	// TODO: Figure out how to represent f as *ir.Function rather than
	// value.Value.

	// Callee.
	callee value.Value
	// Function arguments.
	args []value.Value
	// Result type of the call instruction.
	result types.Type
}

// NewCall returns a new call instruction based on the given callee and function
// arguments.
func NewCall(callee value.Value, args ...value.Value) *Call {
	if typ, ok := callee.Type().(*types.FuncType); ok {
		result := typ.Ret()
		return &Call{callee: callee, args: args, result: result}
	}
	panic(fmt.Sprintf("invalid callee type; expected *types.FuncType, got %T", callee))
}

// Type returns the type of the instruction.
func (i *Call) Type() types.Type {
	return i.result
}

// Ident returns the identifier associated with the instruction.
func (i *Call) Ident() string {
	// TODO: Encode name if containing special characters.
	return "%" + i.name
}

// LLVMString returns the LLVM syntax representation of the instruction.
func (i *Call) LLVMString() string {
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
func (i *Call) Parent() value.Value {
	return i.parent
}

// SetParent sets the parent basic block of the instruction.
func (i *Call) SetParent(parent value.Value) {
	i.parent = parent
}

// SetName sets the name of the local variable storing the result of the
// instruction.
func (i *Call) SetName(name string) {
	i.name = name
}
