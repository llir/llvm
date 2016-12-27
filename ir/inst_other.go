// === [ Other instructions ] ==================================================
//
// References:
//    http://llvm.org/docs/LangRef.html#other-operations

package ir

import (
	"bytes"
	"fmt"

	"github.com/llir/llvm/internal/enc"
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
	Parent *BasicBlock
	// Name of the local variable associated with the instruction.
	Name string
	// Type of the instruction.
	Typ types.Type
	// Integer condition code.
	Cond IntPred
	// Operands.
	X, Y value.Value
}

// NewICmp returns a new icmp instruction based on the given integer condition
// code and operands.
func NewICmp(cond IntPred, x, y value.Value) *InstICmp {
	var typ types.Type = types.I1
	if t, ok := x.Type().(*types.VectorType); ok {
		typ = types.NewVector(types.I1, t.Len)
	}
	return &InstICmp{
		Typ:  typ,
		Cond: cond,
		X:    x,
		Y:    y,
	}
}

// Type returns the type of the instruction.
func (inst *InstICmp) Type() types.Type {
	return inst.Typ
}

// Ident returns the identifier associated with the instruction.
func (inst *InstICmp) Ident() string {
	return enc.Local(inst.Name)
}

// GetName returns the name of the local variable associated with the
// instruction.
func (inst *InstICmp) GetName() string {
	return inst.Name
}

// SetName sets the name of the local variable associated with the instruction.
func (inst *InstICmp) SetName(name string) {
	inst.Name = name
}

// String returns the LLVM syntax representation of the instruction.
func (inst *InstICmp) String() string {
	return fmt.Sprintf("%s = icmp %s %s %s, %s",
		inst.Ident(),
		inst.Cond,
		inst.X.Type(),
		inst.X.Ident(),
		inst.Y.Ident())
}

// GetParent returns the parent basic block of the instruction.
func (inst *InstICmp) GetParent() *BasicBlock {
	return inst.Parent
}

// SetParent sets the parent basic block of the instruction.
func (inst *InstICmp) SetParent(parent *BasicBlock) {
	inst.Parent = parent
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
	Parent *BasicBlock
	// Name of the local variable associated with the instruction.
	Name string
	// Type of the instruction.
	Typ types.Type
	// Floating-point condition code.
	Cond FloatPred
	// Operands.
	X, Y value.Value
}

// NewFCmp returns a new fcmp instruction based on the given floating-point
// condition code and operands.
func NewFCmp(cond FloatPred, x, y value.Value) *InstFCmp {
	var typ types.Type = types.I1
	if t, ok := x.Type().(*types.VectorType); ok {
		typ = types.NewVector(types.I1, t.Len)
	}
	return &InstFCmp{
		Typ:  typ,
		Cond: cond,
		X:    x,
		Y:    y,
	}
}

// Type returns the type of the instruction.
func (inst *InstFCmp) Type() types.Type {
	return inst.Typ
}

// Ident returns the identifier associated with the instruction.
func (inst *InstFCmp) Ident() string {
	return enc.Local(inst.Name)
}

// GetName returns the name of the local variable associated with the
// instruction.
func (inst *InstFCmp) GetName() string {
	return inst.Name
}

// SetName sets the name of the local variable associated with the instruction.
func (inst *InstFCmp) SetName(name string) {
	inst.Name = name
}

// String returns the LLVM syntax representation of the instruction.
func (inst *InstFCmp) String() string {
	return fmt.Sprintf("%s = fcmp %s %s %s, %s",
		inst.Ident(),
		inst.Cond,
		inst.X.Type(),
		inst.X.Ident(),
		inst.Y.Ident())
}

// GetParent returns the parent basic block of the instruction.
func (inst *InstFCmp) GetParent() *BasicBlock {
	return inst.Parent
}

// SetParent sets the parent basic block of the instruction.
func (inst *InstFCmp) SetParent(parent *BasicBlock) {
	inst.Parent = parent
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
	Parent *BasicBlock
	// Name of the local variable associated with the instruction.
	Name string
	// Type of the instruction.
	Typ types.Type
	// Incoming values.
	Incs []*Incoming
}

// NewPhi returns a new phi instruction based on the given incoming values.
func NewPhi(incs ...*Incoming) *InstPhi {
	if len(incs) < 1 {
		panic(fmt.Errorf("invalid number of incoming values; expected > 0, got %d", len(incs)))
	}
	typ := incs[0].X.Type()
	return &InstPhi{
		Typ:  typ,
		Incs: incs,
	}
}

// Type returns the type of the instruction.
func (inst *InstPhi) Type() types.Type {
	return inst.Typ
}

// Ident returns the identifier associated with the instruction.
func (inst *InstPhi) Ident() string {
	return enc.Local(inst.Name)
}

// GetName returns the name of the local variable associated with the
// instruction.
func (inst *InstPhi) GetName() string {
	return inst.Name
}

// SetName sets the name of the local variable associated with the instruction.
func (inst *InstPhi) SetName(name string) {
	inst.Name = name
}

// String returns the LLVM syntax representation of the instruction.
func (inst *InstPhi) String() string {
	buf := &bytes.Buffer{}
	fmt.Fprintf(buf, "%s = phi %s ",
		inst.Ident(),
		inst.Type())
	for i, inc := range inst.Incs {
		if i != 0 {
			buf.WriteString(", ")
		}
		fmt.Fprintf(buf, "[ %s, %s ]",
			inc.X.Ident(),
			inc.Pred.Ident())
	}
	return buf.String()
}

// GetParent returns the parent basic block of the instruction.
func (inst *InstPhi) GetParent() *BasicBlock {
	return inst.Parent
}

// SetParent sets the parent basic block of the instruction.
func (inst *InstPhi) SetParent(parent *BasicBlock) {
	inst.Parent = parent
}

// Incoming represents an incoming value of a phi instruction.
type Incoming struct {
	// Incoming value.
	X value.Value
	// Predecessor basic block of the incoming value.
	Pred *BasicBlock
}

// NewIncoming returns a new incoming value based on the given value and
// predecessor basic block.
func NewIncoming(x value.Value, pred *BasicBlock) *Incoming {
	return &Incoming{
		X:    x,
		Pred: pred,
	}
}

// --- [ select ] --------------------------------------------------------------

// InstSelect represents a select instruction.
//
// References:
//    http://llvm.org/docs/LangRef.html#select-instruction
type InstSelect struct {
	// Parent basic block.
	Parent *BasicBlock
	// Name of the local variable associated with the instruction.
	Name string
	// Selection condition.
	Cond value.Value
	// Operands.
	X, Y value.Value
}

// NewSelect returns a new select instruction based on the given selection
// condition and operands.
func NewSelect(cond, x, y value.Value) *InstSelect {
	return &InstSelect{
		Cond: cond,
		X:    x,
		Y:    y,
	}
}

// Type returns the type of the instruction.
func (inst *InstSelect) Type() types.Type {
	return inst.X.Type()
}

// Ident returns the identifier associated with the instruction.
func (inst *InstSelect) Ident() string {
	return enc.Local(inst.Name)
}

// GetName returns the name of the local variable associated with the
// instruction.
func (inst *InstSelect) GetName() string {
	return inst.Name
}

// SetName sets the name of the local variable associated with the instruction.
func (inst *InstSelect) SetName(name string) {
	inst.Name = name
}

// String returns the LLVM syntax representation of the instruction.
func (inst *InstSelect) String() string {
	return fmt.Sprintf("%s = select %s %s, %s %s, %s %s",
		inst.Ident(),
		inst.Cond.Type(),
		inst.Cond.Ident(),
		inst.X.Type(),
		inst.X.Ident(),
		inst.Y.Type(),
		inst.Y.Ident())
}

// GetParent returns the parent basic block of the instruction.
func (inst *InstSelect) GetParent() *BasicBlock {
	return inst.Parent
}

// SetParent sets the parent basic block of the instruction.
func (inst *InstSelect) SetParent(parent *BasicBlock) {
	inst.Parent = parent
}

// --- [ call ] ----------------------------------------------------------------

// InstCall represents a call instruction.
//
// References:
//    http://llvm.org/docs/LangRef.html#call-instruction
type InstCall struct {
	// Parent basic block.
	Parent *BasicBlock
	// Name of the local variable associated with the instruction.
	Name string
	// Callee.
	//
	// Callee may have one of the following underlying types.
	//
	//    *ir.Function
	//    *types.Param
	Callee value.Named
	// Callee signature.
	Sig *types.FuncType
	// Function arguments.
	Args []value.Value
}

// NewCall returns a new call instruction based on the given callee and function
// arguments.
//
// The callee value may have one of the following underlying types.
//
//    *ir.Function
//    *types.Param
func NewCall(callee value.Named, args ...value.Value) *InstCall {
	typ, ok := callee.Type().(*types.PointerType)
	if !ok {
		panic(fmt.Errorf("invalid callee type, expected *types.PointerType, got %T", callee.Type()))
	}
	sig, ok := typ.Elem.(*types.FuncType)
	if !ok {
		panic(fmt.Errorf("invalid callee signature type, expected *types.FuncType, got %T", typ.Elem))
	}
	return &InstCall{
		Callee: callee,
		Sig:    sig,
		Args:   args,
	}
}

// Type returns the type of the instruction.
func (inst *InstCall) Type() types.Type {
	return inst.Sig.Ret
}

// Ident returns the identifier associated with the instruction.
func (inst *InstCall) Ident() string {
	return enc.Local(inst.Name)
}

// GetName returns the name of the local variable associated with the
// instruction.
func (inst *InstCall) GetName() string {
	return inst.Name
}

// SetName sets the name of the local variable associated with the instruction.
func (inst *InstCall) SetName(name string) {
	inst.Name = name
}

// String returns the LLVM syntax representation of the instruction.
func (inst *InstCall) String() string {
	buf := &bytes.Buffer{}
	if !inst.Type().Equal(types.Void) {
		fmt.Fprintf(buf, "%s = ", inst.Ident())
	}
	// Print callee signature instead of return type for variadic callees.
	sig := inst.Sig
	ret := sig.Ret.String()
	if sig.Variadic {
		ret = sig.String()
	}
	fmt.Fprintf(buf, "call %s %s(",
		ret,
		inst.Callee.Ident())
	for i, arg := range inst.Args {
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

// GetParent returns the parent basic block of the instruction.
func (inst *InstCall) GetParent() *BasicBlock {
	return inst.Parent
}

// SetParent sets the parent basic block of the instruction.
func (inst *InstCall) SetParent(parent *BasicBlock) {
	inst.Parent = parent
}

// --- [ va_arg ] --------------------------------------------------------------

// --- [ landingpad ] ----------------------------------------------------------

// --- [ catchpad ] ------------------------------------------------------------

// --- [ cleanuppad ] ----------------------------------------------------------
