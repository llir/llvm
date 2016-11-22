package ir

import (
	"bytes"
	"fmt"

	"github.com/llir/llvm/ir/types"
	"github.com/llir/llvm/ir/value"
)

// --- [ icmp ] ----------------------------------------------------------------

// InstICmp represents an icmp instruction.
//
// References:
//    http://llvm.org/docs/LangRef.html#icmp-instruction
type InstICmp struct {
	// Parent basic block.
	parent *BasicBlock
	// Identifier associated with the instruction.
	id string
	// Integer condition code.
	cond IntPred
	// Operands.
	x, y value.Value
	// Type of the instruction.
	typ types.Type
}

// NewICmp returns a new icmp instruction based on the given integer condition
// code and operands.
func NewICmp(cond IntPred, x, y value.Value) *InstICmp {
	var typ types.Type = types.I1
	if t, ok := x.Type().(*types.VectorType); ok {
		typ = types.NewVector(types.I1, t.Len())
	}
	return &InstICmp{cond: cond, x: x, y: y, typ: typ}
}

// Type returns the type of the instruction.
func (i *InstICmp) Type() types.Type {
	return i.typ
}

// Ident returns the identifier associated with the instruction.
func (i *InstICmp) Ident() string {
	return local(i.id)
}

// SetIdent sets the identifier associated with the instruction.
func (i *InstICmp) SetIdent(id string) {
	i.id = id
}

// LLVMString returns the LLVM syntax representation of the instruction.
func (i *InstICmp) LLVMString() string {
	x, y := i.X(), i.Y()
	return fmt.Sprintf("%s = icmp %s %s %s, %s %s",
		i.Ident(),
		i.Cond().LLVMString(),
		x.Type().LLVMString(),
		x.Ident(),
		y.Type().LLVMString(),
		y.Ident())
}

// Parent returns the parent basic block of the instruction.
func (i *InstICmp) Parent() *BasicBlock {
	return i.parent
}

// SetParent sets the parent basic block of the instruction.
func (i *InstICmp) SetParent(parent *BasicBlock) {
	i.parent = parent
}

// Cond returns the integer condition code of the icmp instruction.
func (inst *InstICmp) Cond() IntPred {
	return inst.cond
}

// X returns the x operand of the icmp instruction.
func (inst *InstICmp) X() value.Value {
	return inst.x
}

// Y returns the y operand of the icmp instruction.
func (inst *InstICmp) Y() value.Value {
	return inst.y
}

// IntPred represents the set of condition codes of the icmp instruction.
type IntPred int

// Integer condition codes.
const (
	IntPredEQ  IntPred = iota + 1 // eq: equal
	IntPredNE                     // ne: not equal
	IntPredUGT                    // ugt: unsigned greater than
	IntPredUGE                    // uge: unsigned greater than or equal
	IntPredULT                    // ult: unsigned less than
	IntPredULE                    // ule: unsigned less than or equal
	IntPredSGT                    // sgt: signed greater than
	IntPredSGE                    // sge: signed greater than or equal
	IntPredSLT                    // slt: signed less than
	IntPredSLE                    // sle: signed less than or equal
)

// LLVMString returns the LLVM syntax representation of the integer condition
// code.
func (cond IntPred) LLVMString() string {
	m := map[IntPred]string{
		IntPredEQ:  "eq",
		IntPredNE:  "ne",
		IntPredUGT: "ugt",
		IntPredUGE: "uge",
		IntPredULT: "ult",
		IntPredULE: "ule",
		IntPredSGT: "sgt",
		IntPredSGE: "sge",
		IntPredSLT: "slt",
		IntPredSLE: "sle",
	}
	if s, ok := m[cond]; ok {
		return s
	}
	return fmt.Sprintf("<unknown integer condition code %d>", int(cond))
}

// --- [ fcmp ] ----------------------------------------------------------------

// TODO: Add support for fcmp.

// --- [ phi ] -----------------------------------------------------------------

// TODO: Add support for phi.

// --- [ select ] --------------------------------------------------------------

// TODO: Add support for select.

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
