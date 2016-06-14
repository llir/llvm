package instruction

import (
	"bytes"
	"fmt"

	"github.com/llir/llvm/asm"
	"github.com/llir/llvm/ir/types"
	"github.com/llir/llvm/ir/value"
	"github.com/mewkiz/pkg/errutil"
)

// References:
//    http://llvm.org/docs/LangRef.html#other-operations

// TODO: Add support for the remaining other operations:
//    http://llvm.org/docs/LangRef.html#fcmp-instruction
//    http://llvm.org/docs/LangRef.html#phi-instruction
//    http://llvm.org/docs/LangRef.html#select-instruction
//    http://llvm.org/docs/LangRef.html#call-instruction
//    http://llvm.org/docs/LangRef.html#va-arg-instruction
//    http://llvm.org/docs/LangRef.html#landingpad-instruction

// ICmp represents an integer comparison instruction.
type ICmp struct {
	// Condition.
	cond ICond
	// Operands.
	x, y value.Value
}

func NewICmp(cond ICond, x, y value.Value) (*ICmp, error) {
	if !types.Equal(x.Type(), y.Type()) {
		return nil, errutil.Newf("type mismatch between type of x (%v) and y (%v)", x.Type(), y.Type())
	}
	return &ICmp{cond: cond, x: x, y: y}, nil
}

// RetType returns the type of the value produced by the instruction.
func (inst *ICmp) RetType() types.Type {
	return types.I1
}

// String returns the string representation of the instruction.
func (inst *ICmp) String() string {
	return fmt.Sprintf("icmp %s %s %s, %s", inst.cond, inst.x.Type(), inst.x.ValueString(), inst.y.ValueString())
}

// ICond represents an integer comparison condition.
type ICond int

// Integer comparison conditions.
const (
	ICondEq  ICond = iota // equal
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

func (cond ICond) String() string {
	m := map[ICond]string{
		ICondEq:  "eq",
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

type FCmp struct{}

// RetType returns the type of the value produced by the instruction.
func (*FCmp) RetType() types.Type { panic("FCmp.RetType: not yet implemented") }
func (*FCmp) String() string      { panic("FCmp.String: not yet implemented") }

// PHI represents a phi instruction, which is used to implement the φ node in
// the SSA graph representation of the function.
type PHI struct {
	// Type of the incoming values.
	typ types.Type
	// Incoming values.
	incs []*Incoming
}

// NewPHI returns a new phi instruction based on the given type and incoming
// values.
func NewPHI(typ types.Type, incs []*Incoming) (*PHI, error) {
	return &PHI{typ: typ, incs: incs}, nil
}

// RetType returns the type of the value produced by the instruction.
func (inst *PHI) RetType() types.Type {
	return inst.typ
}

// String returns the string representation of the instruction.
func (inst *PHI) String() string {
	incsBuf := new(bytes.Buffer)
	for i, inc := range inst.incs {
		if i > 0 {
			incsBuf.WriteString(", ")
		}
		incsBuf.WriteString(inc.String())
	}
	return fmt.Sprintf("phi %s %s", inst.typ, incsBuf)
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

// TODO: Consider returning *ir.BasicBlock from Pred. The problem is that this
// would create a circular dependency.

// Value returns the incoming value of a given predecessor basic block.
func (inc *Incoming) Value() value.Value {
	return inc.val
}

// Pred returns the predecessor basic block of the incoming value.
func (inc *Incoming) Pred() value.NamedValue {
	return inc.pred
}

// String returns the string representation of the incoming value.
func (inc *Incoming) String() string {
	return fmt.Sprintf("[ %s, %s ]", inc.val.ValueString(), asm.EncLocal(inc.pred.Name()))
}

// Select represents a select instruction.
type Select struct {
	// Selection condition.
	cond value.Value
	// Operands.
	x, y value.Value
}

// NewSelect returns a new select instruction based on the given selection
// condition, and operands.
func NewSelect(cond, x, y value.Value) (*Select, error) {
	// TODO: Add support for boolean vector selection condition type.
	if !types.Equal(cond.Type(), types.I1) {
		return nil, errutil.Newf("invalid selection condition type; expected i1, got %v", cond.Type())
	}
	if !types.Equal(x.Type(), y.Type()) {
		return nil, errutil.Newf("mismatch between operand type x %v and operand type y %v", x.Type(), y.Type())
	}
	return &Select{cond: cond, x: x, y: y}, nil
}

// Cond returns the selection condition of the instruction.
func (inst *Select) Cond() value.Value {
	return inst.cond
}

// X returns the x operand of the instruction.
func (inst *Select) X() value.Value {
	return inst.x
}

// Y returns the y operand of the instruction.
func (inst *Select) Y() value.Value {
	return inst.y
}

// RetType returns the type of the value produced by the instruction.
func (inst *Select) RetType() types.Type {
	return inst.x.Type()
}

// String returns the string representation of the instruction.
func (inst *Select) String() string {
	return fmt.Sprintf("select %s %s, %s %s, %s %s", inst.cond.Type(), inst.cond.ValueString(), inst.x.Type(), inst.x.ValueString(), inst.y.Type(), inst.y.ValueString())
}

// Call represents a call instruction.
type Call struct {
	// Function return type.
	result types.Type
	// Function name.
	fname string
	// Function arguments.
	args []value.Value
}

// NewCall returns a new call instruction based on the given function return
// type, function name, and function arguments.
func NewCall(result types.Type, fname string, args []value.Value) (*Call, error) {
	return &Call{result: result, fname: fname, args: args}, nil
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
	return fmt.Sprintf("call %s %s(%s)", inst.result, asm.EncGlobal(inst.fname), argsBuf)
}

type VAArg struct{}

// RetType returns the type of the value produced by the instruction.
func (*VAArg) RetType() types.Type { panic("VAArg.RetType: not yet implemented") }
func (*VAArg) String() string      { panic("VAArg.String: not yet implemented") }

type LandingPad struct{}

// RetType returns the type of the value produced by the instruction.
func (*LandingPad) RetType() types.Type { panic("LandingPad.RetType: not yet implemented") }
func (*LandingPad) String() string      { panic("LandingPad.String: not yet implemented") }
