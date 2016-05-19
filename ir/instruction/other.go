package instruction

import (
	"fmt"

	"github.com/llir/llvm/ir/types"
	"github.com/llir/llvm/ir/value"
	"github.com/mewkiz/pkg/errutil"
)

// References:
//    http://llvm.org/docs/LangRef.html#other-operations

// TODO: Add support for the remaining other operations:
//    http://llvm.org/docs/LangRef.html#icmp-instruction
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
	return &ICmp{cond: cond, x: x, y: x}, nil
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
	ICondEq  = iota // equal
	ICondNE         // not equal
	ICondUGT        // unsigned greater than
	ICondUGE        // unsigned greater or equal
	ICondULT        // unsigned less than
	ICondULE        // unsigned less or equal
	ICondSGT        // signed greater than
	ICondSGE        // signed greater or equal
	ICondSLT        // signed less than
	ICondSLE        // signed less or equal
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

type PHI struct{}

func (*PHI) Type() types.Type { panic("PHI.Type: not yet implemented") }
func (*PHI) String() string   { panic("PHI.String: not yet implemented") }

type Select struct{}

func (*Select) Type() types.Type { panic("Select.Type: not yet implemented") }
func (*Select) String() string   { panic("Select.String: not yet implemented") }

type Call struct{}

func (*Call) Type() types.Type { panic("Call.Type: not yet implemented") }
func (*Call) String() string   { panic("Call.String: not yet implemented") }

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
