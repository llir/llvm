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

// Type returns the type of the value produced by the instruction.
func (inst *ICmp) Type() types.Type {
	return types.I1
}

// String returns the string representation of the instruction.
func (inst *ICmp) String() string {
	return fmt.Sprintf("icmp %s %s %s, %s", inst.cond, inst.x.Type(), inst.x, inst.y)
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

func (*FCmp) Type() types.Type { panic("FCmp.Type: not yet implemented") }
func (*FCmp) String() string   { panic("FCmp.String: not yet implemented") }

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

// Type returns the type of the value produced by the instruction.
func (inst *PHI) Type() types.Type {
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
	// Label name of the predecessor basic block of the incoming value.
	pred string
}

// NewIncoming returns a new incoming value based on the given value and label
// name of the predecessor basic block.
func NewIncoming(val value.Value, pred string) (*Incoming, error) {
	return &Incoming{val: val, pred: pred}, nil
}

// String returns the string representation of the incoming value.
func (inc *Incoming) String() string {
	return fmt.Sprintf("[ %s, %s ]", inc.val, asm.EncLocal(inc.pred))
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

// Type returns the type of the value produced by the instruction.
func (inst *Select) Type() types.Type {
	return inst.x.Type()
}

// String returns the string representation of the instruction.
func (inst *Select) String() string {
	return fmt.Sprintf("select %s %s, %s %s, %s %s", inst.cond.Type(), inst.cond, inst.x.Type(), inst.x, inst.y.Type(), inst.y)
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

// Type returns the type of the value produced by the instruction.
func (inst *Call) Type() types.Type {
	return inst.result
}

// String returns the string representation of the instruction.
func (inst *Call) String() string {
	argsBuf := new(bytes.Buffer)
	for i, arg := range inst.args {
		if i > 0 {
			argsBuf.WriteString(", ")
		}
		fmt.Fprintf(argsBuf, "%s %s", arg.Type(), arg)
	}
	return fmt.Sprintf("call %s %s(%s)", inst.result, asm.EncGlobal(inst.fname), argsBuf)
}

type VAArg struct{}

func (*VAArg) Type() types.Type { panic("VAArg.Type: not yet implemented") }
func (*VAArg) String() string   { panic("VAArg.String: not yet implemented") }

type LandingPad struct{}

func (*LandingPad) Type() types.Type { panic("LandingPad.Type: not yet implemented") }
func (*LandingPad) String() string   { panic("LandingPad.String: not yet implemented") }

// isValueInst ensures that only instructions which return values can be
// assigned to the Value interface.
func (*ICmp) isValueInst()       {}
func (*FCmp) isValueInst()       {}
func (*PHI) isValueInst()        {}
func (*Select) isValueInst()     {}
func (*Call) isValueInst()       {}
func (*VAArg) isValueInst()      {}
func (*LandingPad) isValueInst() {}
