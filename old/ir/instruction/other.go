// TODO: Enforce stricter check of x and y types. Should be of integer or
// integer vector type. Do this for every binary and bitwise instruction as
// well.

// === [ Other instructions ] ==================================================
//
// References:
//    http://llvm.org/docs/LangRef.html#other-operations

package instruction

import (
	"bytes"
	"fmt"

	"github.com/llir/llvm/ir/internal/enc"
	"github.com/llir/llvm/ir/types"
	"github.com/llir/llvm/ir/value"
	"github.com/mewkiz/pkg/errutil"
)

// --- [ icmp ] ----------------------------------------------------------------

// ICmp represents an icmp instruction, which is used for integer comparison
// operations.
//
// References:
//    http://llvm.org/docs/LangRef.html#icmp-instruction
type ICmp struct {
	// Condition.
	cond ICond
	// Operands.
	x, y value.Value
}

// NewICmp returns a new icmp instruction based on the given condition and
// operands.
//
// Pre-condition: x and y are of identical types. x and y are of integer,
// integer vector, pointer or pointer vector type.
func NewICmp(cond ICond, x, y value.Value) (*ICmp, error) {
	// Validate that x and y are of identical types.
	if !types.Equal(x.Type(), y.Type()) {
		return nil, errutil.Newf("type mismatch between x (%v) and y (%v)", x.Type(), y.Type())
	}
	// Validate that x and y are of integer, integer vector, pointer or pointer
	// vector type.
	if !types.IsInts(x.Type()) && !types.IsPointers(x.Type()) {
		return nil, errutil.Newf("invalid x operand type; expected integer, integer vector, pointer or pointer vector, got %v", x.Type())
	}
	if !types.IsInts(y.Type()) && !types.IsPointers(y.Type()) {
		return nil, errutil.Newf("invalid y operand type; expected integer, integer vector, pointer or pointer vector, got %v", y.Type())
	}
	return &ICmp{cond: cond, x: x, y: y}, nil
}

// Cond returns the integer comparison condition of the icmp instruction.
func (inst *ICmp) Cond() ICond {
	return inst.cond
}

// X returns the x operand of the icmp instruction.
func (inst *ICmp) X() value.Value {
	return inst.x
}

// Y returns the y operand of the icmp instruction.
func (inst *ICmp) Y() value.Value {
	return inst.y
}

// RetType returns the type of the value produced by the instruction.
func (inst *ICmp) RetType() types.Type {
	if typ, ok := inst.x.Type().(*types.Vector); ok {
		// TODO: Move to constructor?
		t, _ := types.NewVector(types.I1, typ.Len())
		return t
	}
	return types.I1
}

// String returns the string representation of the instruction.
func (inst *ICmp) String() string {
	x, y := inst.X(), inst.Y()
	return fmt.Sprintf("icmp %s %s %s, %s", inst.Cond(), x.Type(), x.ValueString(), y.ValueString())
}

// ICond represents an integer comparison condition.
type ICond int

// Integer comparison conditions.
const (
	ICondEQ  ICond = iota // equal
	ICondNE               // not equal
	ICondUGT              // unsigned greater than
	ICondUGE              // unsigned greater or equal
	ICondULT              // unsigned less than
	ICondULE              // unsigned less or equal
	ICondSGT              // signed greater than
	ICondSGE              // signed greater or equal
	ICondSLT              // signed less than
	ICondSLE              // signed less or equal
)

// String returns the string representation of the integer comparison condition.
func (cond ICond) String() string {
	m := map[ICond]string{
		ICondEQ:  "eq",
		ICondNE:  "ne",
		ICondUGT: "ugt",
		ICondUGE: "uge",
		ICondULT: "ult",
		ICondULE: "ule",
		ICondSGT: "sgt",
		ICondSGE: "sge",
		ICondSLT: "slt",
		ICondSLE: "sle",
	}
	if s, ok := m[cond]; ok {
		return s
	}
	return fmt.Sprintf("ICond(%d)", int(cond))
}

// --- [ fcmp ] ----------------------------------------------------------------

// TODO: Implement FCmp.

// FCmp represents a fcmp instruction, which is used for floating-point
// comparison operations.
//
// References:
//    http://llvm.org/docs/LangRef.html#fcmp-instruction
type FCmp struct{}

// RetType returns the type of the value produced by the instruction.
func (*FCmp) RetType() types.Type { panic("FCmp.RetType: not yet implemented") }
func (*FCmp) String() string      { panic("FCmp.String: not yet implemented") }

// --- [ phi ] -----------------------------------------------------------------

// PHI represents a phi instruction, which is used to implement the φ node in
// the SSA graph of a function.
//
// References:
//    http://llvm.org/docs/LangRef.html#phi-instruction
type PHI struct {
	// Incoming values.
	incs []*Incoming
}

// NewPHI returns a new phi instruction based on the given incoming values.
//
// Pre-condition: incs is non-empty and each incoming value is of identical
// type.
func NewPHI(incs []*Incoming) (*PHI, error) {
	// Validate that incs is non-empty.
	if len(incs) < 1 {
		return nil, errutil.Newf("invalid number of incoming values; expected > 0, got %d", len(incs))
	}
	// Validate that each incoming value is of identical type.
	a := incs[0].val.Type()
	for i := 1; i < len(incs); i++ {
		b := incs[i].val.Type()
		if !types.Equal(a, b) {
			return nil, errutil.Newf("type mismatch between incoming value 0 (%v) and %d (%v)", a, i, b)
		}
	}
	return &PHI{incs: incs}, nil
}

// Incs returns the incomming values of the phi instruction.
func (inst *PHI) Incs() []*Incoming {
	return inst.incs
}

// RetType returns the type of the value produced by the instruction.
func (inst *PHI) RetType() types.Type {
	return inst.incs[0].val.Type()
}

// String returns the string representation of the instruction.
func (inst *PHI) String() string {
	incsBuf := new(bytes.Buffer)
	for i, inc := range inst.Incs() {
		if i > 0 {
			incsBuf.WriteString(", ")
		}
		incsBuf.WriteString(inc.String())
	}
	return fmt.Sprintf("phi %s %s", inst.RetType(), incsBuf)
}

// Incoming represents an incoming value from a predecessor basic block, as
// specified by PHI instructions.
type Incoming struct {
	// Incoming value.
	val value.Value
	// Predecessor basic block of the incoming value.
	pred value.NamedValue
}

// NewIncoming returns a new incoming value based on the given value and
// predecessor basic block.
func NewIncoming(val value.Value, pred value.NamedValue) (*Incoming, error) {
	// TODO: Validate that pred is of type *ir.BasicBlock. Better yet, chance the
	// signature of NewIncoming to enforce this. Another approach, is to simply
	// check that the type of pred is "label".
	return &Incoming{val: val, pred: pred}, nil
}

// Value returns the incoming value of the predecessor basic block.
func (inc *Incoming) Value() value.Value {
	return inc.val
}

// TODO: Consider returning *ir.BasicBlock from Pred. The problem is that this
// would create a circular dependency.

// Pred returns the predecessor basic block of the incoming value.
func (inc *Incoming) Pred() value.NamedValue {
	return inc.pred
}

// String returns the string representation of the incoming value.
func (inc *Incoming) String() string {
	return fmt.Sprintf("[ %s, %s ]", inc.Value().ValueString(), inc.Pred().ValueString())
}

// --- [ select ] --------------------------------------------------------------

// Select represents a select instruction, which is used to select one of two
// values, without branching, based on a boolean selection condition.
//
// References:
//    http://llvm.org/docs/LangRef.html#select-instructionn
type Select struct {
	// Selection condition.
	cond value.Value
	// Operands.
	x, y value.Value
}

// NewSelect returns a new select instruction based on the given selection
// condition, and operands.
//
// Pre-condition: cond is of boolean or boolean vector type. x and y are of
// identical types.
func NewSelect(cond, x, y value.Value) (*Select, error) {
	// Validate that cond is of boolean or boolean vector type.
	if !types.IsBools(cond.Type()) {
		return nil, errutil.Newf("invalid selection condition type; expected boolean or boolean vector, got %v", cond.Type())
	}
	// Validate that x and y are of identical types.
	if !types.Equal(x.Type(), y.Type()) {
		return nil, errutil.Newf("type mismatch between x (%v) and y (%v)", x.Type(), y.Type())
	}
	return &Select{cond: cond, x: x, y: y}, nil
}

// Cond returns the selection condition of the select instruction.
func (inst *Select) Cond() value.Value {
	return inst.cond
}

// X returns the x operand of the select instruction.
func (inst *Select) X() value.Value {
	return inst.x
}

// Y returns the y operand of the select instruction.
func (inst *Select) Y() value.Value {
	return inst.y
}

// RetType returns the type of the value produced by the instruction.
func (inst *Select) RetType() types.Type {
	return inst.x.Type()
}

// String returns the string representation of the instruction.
func (inst *Select) String() string {
	cond := inst.Cond()
	x, y := inst.X(), inst.Y()
	return fmt.Sprintf("select %s %s, %s %s, %s %s", cond.Type(), cond.ValueString(), x.Type(), x.ValueString(), y.Type(), y.ValueString())
}

// --- [ call ] ----------------------------------------------------------------

// Call represents a call instruction, which is used for function calls.
//
// References:
//    http://llvm.org/docs/LangRef.html#call-instruction
type Call struct {
	// Function return type.
	result types.Type
	// Callee function name.
	callee string
	// Function arguments.
	args []value.Value
}

// TODO: Change type of callee to value.NamedValue or *ir.Function?

// NewCall returns a new call instruction based on the given function return
// type, callee function name, and function arguments.
func NewCall(result types.Type, callee string, args []value.Value) (*Call, error) {
	// TODO: Validate argument types based on function signature.
	return &Call{result: result, callee: callee, args: args}, nil
}

// TODO: Rename Callee to FuncName or Target or ...?

// TODO: Update the return type of Callee once the type of callee is updated.

// Callee returns the callee of the call instruction.
func (inst *Call) Callee() string {
	return inst.callee
}

// Args returns the arguments of the call instruction.
func (inst *Call) Args() []value.Value {
	return inst.args
}

// RetType returns the type of the value produced by the instruction.
func (inst *Call) RetType() types.Type {
	return inst.result
}

// String returns the string representation of the instruction.
func (inst *Call) String() string {
	argsBuf := new(bytes.Buffer)
	for i, arg := range inst.args {
		if i > 0 {
			argsBuf.WriteString(", ")
		}
		fmt.Fprintf(argsBuf, "%s %s", arg.Type(), arg.ValueString())
	}
	// TODO: Use inst.callee.ValueString() rather than enc.Global(inst.callee),
	// once the type of callee is changed to NamedValue or *ir.Function.
	return fmt.Sprintf("call %s %s(%s)", inst.result, enc.Global(inst.callee), argsBuf)
}

// --- [ vaarg ] ---------------------------------------------------------------

// TODO: Implement VAArg.

// VAArg represents a va_arg instruction, which is used to access the variable
// arguments of a function call.
//
// References:
//    http://llvm.org/docs/LangRef.html#va-arg-instruction
type VAArg struct{}

// RetType returns the type of the value produced by the instruction.
func (*VAArg) RetType() types.Type { panic("VAArg.RetType: not yet implemented") }
func (*VAArg) String() string      { panic("VAArg.String: not yet implemented") }

// --- [ landingpad ] ----------------------------------------------------------

// TODO: Implement LandingPad.

// LandingPad represents a landingpad instruction, which is used to specify a
// landing pad basic block where caught exceptions land.
//
// References:
//    http://llvm.org/docs/LangRef.html#landingpad-instruction
type LandingPad struct{}

// RetType returns the type of the value produced by the instruction.
func (*LandingPad) RetType() types.Type { panic("LandingPad.RetType: not yet implemented") }
func (*LandingPad) String() string      { panic("LandingPad.String: not yet implemented") }

// --- [ catchpad ] ----------------------------------------------------------

// TODO: Implement CatchPad.

// CatchPad represents a catchpad instruction, which is used to specify a catch
// handler basic block that attempts to transfer control to catch an exception.
//
// References:
//    http://llvm.org/docs/LangRef.html#catchpad-instruction
type CatchPad struct{}

// RetType returns the type of the value produced by the instruction.
func (*CatchPad) RetType() types.Type { panic("CatchPad.RetType: not yet implemented") }
func (*CatchPad) String() string      { panic("CatchPad.String: not yet implemented") }

// --- [ cleanuppad ] ----------------------------------------------------------

// TODO: Implement CleanupPad.

// CleanupPad represents a cleanuppad instruction, which is used to specify a
// cleanup handler basic block that attempts to transfer control to run cleanup
// actions.
//
// References:
//    http://llvm.org/docs/LangRef.html#cleanuppad-instruction
type CleanupPad struct{}

// RetType returns the type of the value produced by the instruction.
func (*CleanupPad) RetType() types.Type { panic("CleanupPad.RetType: not yet implemented") }
func (*CleanupPad) String() string      { panic("CleanupPad.String: not yet implemented") }
