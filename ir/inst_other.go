// === [ Other instructions ] ==================================================
//
// References:
//    http://llvm.org/docs/LangRef.html#other-operations

package ir

import (
	"bytes"
	"fmt"

	"github.com/llir/llvm/ir/internal/enc"
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
	ident string
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
func (inst *InstICmp) Type() types.Type {
	return inst.typ
}

// Ident returns the identifier associated with the instruction.
func (inst *InstICmp) Ident() string {
	return enc.Local(inst.ident)
}

// SetIdent sets the identifier associated with the instruction.
func (inst *InstICmp) SetIdent(ident string) {
	inst.ident = ident
}

// String returns the LLVM syntax representation of the instruction.
func (inst *InstICmp) String() string {
	x, y := inst.X(), inst.Y()
	return fmt.Sprintf("%s = icmp %s %s %s, %s %s",
		inst.Ident(),
		inst.Cond(),
		x.Type(),
		x.Ident(),
		y.Type(),
		y.Ident())
}

// Parent returns the parent basic block of the instruction.
func (inst *InstICmp) Parent() *BasicBlock {
	return inst.parent
}

// SetParent sets the parent basic block of the instruction.
func (inst *InstICmp) SetParent(parent *BasicBlock) {
	inst.parent = parent
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
	IntEQ  IntPred = iota + 1 // eq: equal
	IntNE                     // ne: not equal
	IntUGT                    // ugt: unsigned greater than
	IntUGE                    // uge: unsigned greater than or equal
	IntULT                    // ult: unsigned less than
	IntULE                    // ule: unsigned less than or equal
	IntSGT                    // sgt: signed greater than
	IntSGE                    // sge: signed greater than or equal
	IntSLT                    // slt: signed less than
	IntSLE                    // sle: signed less than or equal
)

// String returns the LLVM syntax representation of the integer condition
// code.
func (cond IntPred) String() string {
	m := map[IntPred]string{
		IntEQ:  "eq",
		IntNE:  "ne",
		IntUGT: "ugt",
		IntUGE: "uge",
		IntULT: "ult",
		IntULE: "ule",
		IntSGT: "sgt",
		IntSGE: "sge",
		IntSLT: "slt",
		IntSLE: "sle",
	}
	if s, ok := m[cond]; ok {
		return s
	}
	return fmt.Sprintf("<unknown integer condition code %d>", int(cond))
}

// --- [ fcmp ] ----------------------------------------------------------------

// InstFCmp represents an fcmp instruction.
//
// References:
//    http://llvm.org/docs/LangRef.html#fcmp-instruction
type InstFCmp struct {
	// Parent basic block.
	parent *BasicBlock
	// Identifier associated with the instruction.
	ident string
	// Floating-point condition code.
	cond FloatPred
	// Operands.
	x, y value.Value
	// Type of the instruction.
	typ types.Type
}

// NewFCmp returns a new fcmp instruction based on the given floating-point
// condition code and operands.
func NewFCmp(cond FloatPred, x, y value.Value) *InstFCmp {
	var typ types.Type = types.I1
	if t, ok := x.Type().(*types.VectorType); ok {
		typ = types.NewVector(types.I1, t.Len())
	}
	return &InstFCmp{cond: cond, x: x, y: y, typ: typ}
}

// Type returns the type of the instruction.
func (inst *InstFCmp) Type() types.Type {
	return inst.typ
}

// Ident returns the identifier associated with the instruction.
func (inst *InstFCmp) Ident() string {
	return enc.Local(inst.ident)
}

// SetIdent sets the identifier associated with the instruction.
func (inst *InstFCmp) SetIdent(ident string) {
	inst.ident = ident
}

// String returns the LLVM syntax representation of the instruction.
func (inst *InstFCmp) String() string {
	x, y := inst.X(), inst.Y()
	return fmt.Sprintf("%s = fcmp %s %s %s, %s %s",
		inst.Ident(),
		inst.Cond(),
		x.Type(),
		x.Ident(),
		y.Type(),
		y.Ident())
}

// Parent returns the parent basic block of the instruction.
func (inst *InstFCmp) Parent() *BasicBlock {
	return inst.parent
}

// SetParent sets the parent basic block of the instruction.
func (inst *InstFCmp) SetParent(parent *BasicBlock) {
	inst.parent = parent
}

// Cond returns the floating-point condition code of the fcmp instruction.
func (inst *InstFCmp) Cond() FloatPred {
	return inst.cond
}

// X returns the x operand of the fcmp instruction.
func (inst *InstFCmp) X() value.Value {
	return inst.x
}

// Y returns the y operand of the fcmp instruction.
func (inst *InstFCmp) Y() value.Value {
	return inst.y
}

// FloatPred represents the set of condition codes of the fcmp instruction.
type FloatPred int

// Floating-point condition codes.
const (
	FloatFalse FloatPred = iota + 1 // false: no comparison, always returns false
	FloatOEQ                        // oeq: ordered and equal
	FloatOGT                        // ogt: ordered and greater than
	FloatOGE                        // oge: ordered and greater than or equal
	FloatOLT                        // olt: ordered and less than
	FloatOLE                        // ole: ordered and less than or equal
	FloatONE                        // one: ordered and not equal
	FloatORD                        // ord: ordered (no nans)
	FloatUEQ                        // ueq: unordered or equal
	FloatUGT                        // ugt: unordered or greater than
	FloatUGE                        // uge: unordered or greater than or equal
	FloatULT                        // ult: unordered or less than
	FloatULE                        // ule: unordered or less than or equal
	FloatUNE                        // une: unordered or not equal
	FloatUNO                        // uno: unordered (either nans)
	FloatTrue                       // true: no comparison, always returns true
)

// String returns the LLVM syntax representation of the floating-point
// condition code.
func (cond FloatPred) String() string {
	m := map[FloatPred]string{
		FloatFalse: "false",
		FloatOEQ:   "oeq",
		FloatOGT:   "ogt",
		FloatOGE:   "oge",
		FloatOLT:   "olt",
		FloatOLE:   "ole",
		FloatONE:   "one",
		FloatORD:   "ord",
		FloatUEQ:   "ueq",
		FloatUGT:   "ugt",
		FloatUGE:   "uge",
		FloatULT:   "ult",
		FloatULE:   "ule",
		FloatUNE:   "une",
		FloatUNO:   "uno",
		FloatTrue:  "true",
	}
	if s, ok := m[cond]; ok {
		return s
	}
	return fmt.Sprintf("<unknown floating-point condition code %d>", int(cond))
}

// --- [ phi ] -----------------------------------------------------------------

// InstPhi represents a phi instruction.
//
// References:
//    http://llvm.org/docs/LangRef.html#phi-instruction
type InstPhi struct {
	// Parent basic block.
	parent *BasicBlock
	// Identifier associated with the instruction.
	ident string
	// Incoming values.
	incs []*Incoming
	// Type of the instruction.
	typ types.Type
}

// NewPhi returns a new phi instruction based on the given incoming values.
func NewPhi(incs ...*Incoming) *InstPhi {
	if len(incs) < 1 {
		panic(fmt.Sprintf("invalid number of incoming values; expected > 0, got %d", len(incs)))
	}
	typ := incs[0].x.Type()
	return &InstPhi{incs: incs, typ: typ}
}

// Type returns the type of the instruction.
func (inst *InstPhi) Type() types.Type {
	return inst.typ
}

// Ident returns the identifier associated with the instruction.
func (inst *InstPhi) Ident() string {
	return enc.Local(inst.ident)
}

// SetIdent sets the identifier associated with the instruction.
func (inst *InstPhi) SetIdent(ident string) {
	inst.ident = ident
}

// String returns the LLVM syntax representation of the instruction.
func (inst *InstPhi) String() string {
	buf := &bytes.Buffer{}
	fmt.Fprintf(buf, "%s = phi %s ",
		inst.Ident(),
		inst.Type())
	for j, inc := range inst.Incs() {
		if j != 0 {
			buf.WriteString(", ")
		}
		fmt.Fprintf(buf, "[ %s, %s ]",
			inc.X().Ident(),
			inc.Pred().Ident())
	}
	return buf.String()
}

// Parent returns the parent basic block of the instruction.
func (inst *InstPhi) Parent() *BasicBlock {
	return inst.parent
}

// SetParent sets the parent basic block of the instruction.
func (inst *InstPhi) SetParent(parent *BasicBlock) {
	inst.parent = parent
}

// Incs returns the incoming values of the phi instruction.
func (inst *InstPhi) Incs() []*Incoming {
	return inst.incs
}

// Incoming represents an incoming value of a phi instruction.
type Incoming struct {
	// Incoming value.
	x value.Value
	// Predecessor basic block of the incoming value.
	pred *BasicBlock
}

// NewIncoming returns a new incoming value based on the given value and
// predecessor basic block.
func NewIncoming(x value.Value, pred *BasicBlock) *Incoming {
	return &Incoming{x: x, pred: pred}
}

// X returns the incoming value.
func (inc *Incoming) X() value.Value {
	return inc.x
}

// Pred returns the predecessor basic block of the incoming value.
func (inc *Incoming) Pred() *BasicBlock {
	return inc.pred
}

// --- [ select ] --------------------------------------------------------------

// InstSelect represents a select instruction.
//
// References:
//    http://llvm.org/docs/LangRef.html#select-instruction
type InstSelect struct {
	// Parent basic block.
	parent *BasicBlock
	// Identifier associated with the instruction.
	ident string
	// Selection condition.
	cond value.Value
	// Operands.
	x, y value.Value
}

// NewSelect returns a new select instruction based on the given selection
// condition and operands.
func NewSelect(cond, x, y value.Value) *InstSelect {
	return &InstSelect{cond: cond, x: x, y: y}
}

// Type returns the type of the instruction.
func (inst *InstSelect) Type() types.Type {
	return inst.x.Type()
}

// Ident returns the identifier associated with the instruction.
func (inst *InstSelect) Ident() string {
	return enc.Local(inst.ident)
}

// SetIdent sets the identifier associated with the instruction.
func (inst *InstSelect) SetIdent(ident string) {
	inst.ident = ident
}

// String returns the LLVM syntax representation of the instruction.
func (inst *InstSelect) String() string {
	cond, x, y := inst.Cond(), inst.X(), inst.Y()
	return fmt.Sprintf("%s = select %s %s, %s %s, %s %s",
		inst.Ident(),
		cond.Type(),
		cond.Ident(),
		x.Type(),
		x.Ident(),
		y.Type(),
		y.Ident())
}

// Parent returns the parent basic block of the instruction.
func (inst *InstSelect) Parent() *BasicBlock {
	return inst.parent
}

// SetParent sets the parent basic block of the instruction.
func (inst *InstSelect) SetParent(parent *BasicBlock) {
	inst.parent = parent
}

// Cond returns the selection condition of the select instruction.
func (inst *InstSelect) Cond() value.Value {
	return inst.cond
}

// X returns the x operand of the select instruction.
func (inst *InstSelect) X() value.Value {
	return inst.x
}

// Y returns the y operand of the select instruction.
func (inst *InstSelect) Y() value.Value {
	return inst.y
}

// --- [ call ] ----------------------------------------------------------------

// InstCall represents a call instruction.
//
// References:
//    http://llvm.org/docs/LangRef.html#call-instruction
type InstCall struct {
	// Parent basic block.
	parent *BasicBlock
	// Identifier associated with the instruction.
	ident string
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
func (inst *InstCall) Type() types.Type {
	return inst.callee.sig.RetType()
}

// Ident returns the identifier associated with the instruction.
func (inst *InstCall) Ident() string {
	return enc.Local(inst.ident)
}

// SetIdent sets the identifier associated with the instruction.
func (inst *InstCall) SetIdent(ident string) {
	inst.ident = ident
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
		inst.Callee().Ident())
	for i, arg := range inst.Args() {
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
func (inst *InstCall) Parent() *BasicBlock {
	return inst.parent
}

// SetParent sets the parent basic block of the instruction.
func (inst *InstCall) SetParent(parent *BasicBlock) {
	inst.parent = parent
}

// Callee returns the callee of the call instruction.
func (inst *InstCall) Callee() *Function {
	return inst.callee
}

// Args returns the function arguments of the call instruction.
func (inst *InstCall) Args() []value.Value {
	return inst.args
}

// --- [ va_arg ] --------------------------------------------------------------

// --- [ landingpad ] ----------------------------------------------------------

// --- [ catchpad ] ------------------------------------------------------------

// --- [ cleanuppad ] ----------------------------------------------------------
