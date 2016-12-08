// === [ Other instructions ] ==================================================

package dummy

import (
	"bytes"
	"fmt"

	"github.com/llir/llvm/internal/enc"
	"github.com/llir/llvm/ir"
	"github.com/llir/llvm/ir/types"
	"github.com/llir/llvm/ir/value"
)

// --- [ call ] ----------------------------------------------------------------

// InstCall represents a dummy call instruction.
type InstCall struct {
	// Parent basic block.
	parent *ir.BasicBlock
	// Name of the local variable associated with the instruction.
	name string
	// Return type.
	ret types.Type
	// Callee.
	callee string
	// Function arguments.
	args []value.Value
}

// NewCall returns a new dummy call instruction based on the given return type,
// callee and function arguments.
func NewCall(ret types.Type, callee string, args ...value.Value) *InstCall {
	return &InstCall{ret: ret, callee: callee, args: args}
}

// Type returns the type of the instruction.
func (inst *InstCall) Type() types.Type {
	return inst.ret
}

// Ident returns the identifier associated with the instruction.
func (inst *InstCall) Ident() string {
	return enc.Local(inst.name)
}

// Name returns the name of the local variable associated with the instruction.
func (inst *InstCall) Name() string {
	return inst.name
}

// SetName sets the name of the local variable associated with the instruction.
func (inst *InstCall) SetName(name string) {
	inst.name = name
}

// String returns the LLVM syntax representation of the instruction.
func (inst *InstCall) String() string {
	buf := &bytes.Buffer{}
	typ := inst.Type()
	if !typ.Equal(types.Void) {
		fmt.Fprintf(buf, "%s = ", inst.Ident())
	}
	fmt.Fprintf(buf, "call %s %s(",
		typ,
		inst.callee)
	for i, arg := range inst.args {
		if i != 0 {
			buf.WriteString(", ")
		}
		fmt.Fprintf(buf, "%s %s",
			arg.Type(),
			arg.Ident())
	}
	buf.WriteString(")")
	return buf.String()
}

// Parent returns the parent basic block of the instruction.
func (inst *InstCall) Parent() *ir.BasicBlock {
	return inst.parent
}

// SetParent sets the parent basic block of the instruction.
func (inst *InstCall) SetParent(parent *ir.BasicBlock) {
	inst.parent = parent
}

// Callee returns the callee of the call instruction.
func (inst *InstCall) Callee() string {
	return inst.callee
}

// Args returns the function arguments of the call instruction.
func (inst *InstCall) Args() []value.Value {
	return inst.args
}
