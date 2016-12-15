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

// --- [ phi ] -----------------------------------------------------------------

// InstPhi represents a dummy phi instruction.
type InstPhi struct {
	// Parent basic block.
	parent *ir.BasicBlock
	// Name of the local variable associated with the instruction.
	name string
	// Type of the instruction.
	typ types.Type
	// Incoming values.
	incs []*Incoming
	// Track uses of the value.
	used
}

// NewPhi returns a new dummy phi instruction based on the given type and
// incoming values.
func NewPhi(typ types.Type, incs ...*Incoming) *InstPhi {
	return &InstPhi{typ: typ, incs: incs}
}

// Type returns the type of the instruction.
func (inst *InstPhi) Type() types.Type {
	return inst.typ
}

// Ident returns the identifier associated with the instruction.
func (inst *InstPhi) Ident() string {
	return enc.Local(inst.name)
}

// Name returns the name of the local variable associated with the instruction.
func (inst *InstPhi) Name() string {
	return inst.name
}

// SetName sets the name of the local variable associated with the instruction.
func (inst *InstPhi) SetName(name string) {
	inst.name = name
}

// String returns the LLVM syntax representation of the instruction.
func (inst *InstPhi) String() string {
	buf := &bytes.Buffer{}
	fmt.Fprintf(buf, "%s = phi %s ",
		inst.Ident(),
		inst.Type())
	for i, inc := range inst.Incs() {
		if i != 0 {
			buf.WriteString(", ")
		}
		x, ok := inc.X().(value.Value)
		if !ok {
			panic(fmt.Sprintf("invalid x type; expected value.Value, got %T", inc.X()))
		}
		fmt.Fprintf(buf, "[ %s, %s ]",
			x,
			enc.Local(inc.Pred()))
	}
	return buf.String()
}

// Parent returns the parent basic block of the instruction.
func (inst *InstPhi) Parent() *ir.BasicBlock {
	return inst.parent
}

// SetParent sets the parent basic block of the instruction.
func (inst *InstPhi) SetParent(parent *ir.BasicBlock) {
	inst.parent = parent
}

// Incs returns the incoming values of the phi instruction.
func (inst *InstPhi) Incs() []*Incoming {
	return inst.incs
}

// Incoming represents a dummy incoming value of a phi instruction.
type Incoming struct {
	// Incoming value.
	//
	// Initially holds *irx.IntLit, *irx.LocalIdent, ... when created from
	// irx.NewIncoming since the type is not yet known. The irx.NewPhiInst later
	// replaces this with a value (e.g. *constant.Int, *dummy.Local, ...).
	x interface{}
	// Predecessor basic block of the incoming value.
	pred string
}

// NewIncoming returns a new dummy incoming value based on the given value and
// predecessor basic block.
func NewIncoming(x interface{}, pred string) *Incoming {
	return &Incoming{x: x, pred: pred}
}

// X returns the incoming value.
func (inc *Incoming) X() interface{} {
	return inc.x
}

// SetX sets the incoming value.
func (inc *Incoming) SetX(x value.Value) {
	inc.x = x
}

// Pred returns the predecessor basic block of the incoming value.
func (inc *Incoming) Pred() string {
	return inc.pred
}

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
	// Specifies whether the callee is a local identifier.
	calleeLocal bool
	// Function arguments.
	args []value.Value
	// Track uses of the value.
	used
}

// NewCall returns a new dummy call instruction based on the given return type,
// callee and function arguments.
func NewCall(ret types.Type, callee string, calleeLocal bool, args ...value.Value) *InstCall {
	return &InstCall{ret: ret, callee: callee, calleeLocal: calleeLocal, args: args}
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
	callee := enc.Global(inst.Callee())
	if inst.calleeLocal {
		callee = enc.Local(inst.Callee())
	}
	fmt.Fprintf(buf, "call %s %s(",
		typ,
		callee)
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

// CalleeLocal reports whether the callee is a local identifier.
func (inst *InstCall) CalleeLocal() bool {
	return inst.calleeLocal
}

// Args returns the function arguments of the call instruction.
func (inst *InstCall) Args() []value.Value {
	return inst.args
}
