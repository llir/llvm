package instruction

import (
	"fmt"

	"github.com/llir/llvm/ir/internal/enc"
	"github.com/llir/llvm/ir/types"
	"github.com/llir/llvm/ir/value"
	"github.com/mewkiz/pkg/errutil"
)

// A Ret instruction returns control flow, and optionally a value, from a callee
// back to its caller.
//
// Syntax:
//    ret <Type> <Val>
//    ret void
//
// Semantics:
//    return val;
//    return;
//
// Reference:
//    http://llvm.org/docs/LangRef.html#ret-instruction
type Ret struct {
	// Return value; or nil if "void" return.
	val value.Value
}

// NewRet returns a new ret instruction based on the given return value. A nil
// return value indicates a "void" return instruction.
func NewRet(val value.Value) (*Ret, error) {
	if val != nil && types.IsVoid(val.Type()) {
		return nil, errutil.Newf(`expected no return value for return type "void"; got %q`, val)
	}
	return &Ret{val: val}, nil
}

// Value returns the return value of the ret instruction. A nil return value
// indicates a "void" return instruction.
func (term *Ret) Value() value.Value {
	return term.val
}

// String returns the string representation of the instruction.
func (term *Ret) String() string {
	if term.val == nil {
		return "ret void"
	}
	val := term.Value()
	return fmt.Sprintf("ret %v %v", val.Type(), val.ValueString())
}

// TODO: Add support for the remaining terminator instructions:
//    http://llvm.org/docs/LangRef.html#switch-instruction
//    http://llvm.org/docs/LangRef.html#indirectbr-instruction
//    http://llvm.org/docs/LangRef.html#invoke-instruction
//    http://llvm.org/docs/LangRef.html#resume-instruction

// Jmp represents an unconditional branch instruction.
type Jmp struct {
	// Basic block label name of the target branch.
	target value.NamedValue
}

// NewJmp returns a new unconditional branch instruction based on the given
// target branch.
func NewJmp(target value.NamedValue) (*Jmp, error) {
	// TODO: Validate that target is an *ir.BasicBlock. Better yet, chance the
	// signature of NewJmp to enforce this. Another approach, is to simply check
	// that the type of target is "label".
	return &Jmp{target: target}, nil
}

// TODO: Consider returning *ir.BasicBlock from Target. The problem is that this
// would create a circular dependency.

// Target returns the basic block label name of the target branch.
func (term *Jmp) Target() value.NamedValue {
	return term.target
}

// String returns the string representation of the instruction.
func (term *Jmp) String() string {
	return fmt.Sprintf("br label %s", enc.Local(term.target.Name()))
}

// Br represents a conditional branch instruction.
type Br struct {
	// Branching condition.
	cond value.Value
	// Basic block label name of the true target branch.
	trueBranch value.NamedValue
	// Basic block label name of the false target branch.
	falseBranch value.NamedValue
}

// NewBr returns a new conditional branch instruction based on the given
// branching condition, and the true and false target branches.
func NewBr(cond value.Value, trueBranch, falseBranch value.NamedValue) (*Br, error) {
	// TODO: Validate that trueBranch and falseBranch are of type *ir.BasicBlock.
	// Better yet, chance the signature of NewBr to enforce this. Another
	// approach, is to simply check that the type of trueBranch and falseBranch
	// are both "label".
	if !types.Equal(cond.Type(), types.I1) {
		return nil, errutil.Newf("conditional type mismatch; expected i1, got %v", cond.Type())
	}
	return &Br{cond: cond, trueBranch: trueBranch, falseBranch: falseBranch}, nil
}

// Cond returns the branching condition of the instruction
func (term *Br) Cond() value.Value {
	return term.cond
}

// TODO: Consider returning *ir.BasicBlock from TrueBranch and FalseBranch. The
// problem is that this would create a circular dependency.

// TrueBranch returns the basic block label name of the true target branch.
func (term *Br) TrueBranch() value.NamedValue {
	return term.trueBranch
}

// FalseBranch returns the basic block label name of the false target branch.
func (term *Br) FalseBranch() value.NamedValue {
	return term.falseBranch
}

// String returns the string representation of the instruction.
func (term *Br) String() string {
	return fmt.Sprintf("br %s %s, label %s, label %s", term.cond.Type(), term.cond.ValueString(), enc.Local(term.trueBranch.Name()), enc.Local(term.falseBranch.Name()))
}

type Switch struct{}

func (*Switch) String() string { panic("Switch.String: not yet implemented") }

type IndirectBr struct{}

func (*IndirectBr) String() string { panic("IndirectBr.String: not yet implemented") }

type Invoke struct{}

func (*Invoke) String() string { panic("Invoke.String: not yet implemented") }

type Resume struct{}

func (*Resume) String() string { panic("Resume.String: not yet implemented") }

// Unreachable represents an unreachable instruction.
//
// References:
//    http://llvm.org/docs/LangRef.html#unreachable-instruction
type Unreachable struct{}

// NewUnreachable returns a new unreachable instruction.
func NewUnreachable() (*Unreachable, error) {
	return &Unreachable{}, nil
}

// String returns the string representation of the instruction.
func (term *Unreachable) String() string {
	return "unreachable"
}

// isTerm ensures that only terminator instructions can be assigned to the
// Terminator interface.
func (*Ret) isTerm()         {}
func (*Jmp) isTerm()         {}
func (*Br) isTerm()          {}
func (*Switch) isTerm()      {}
func (*IndirectBr) isTerm()  {}
func (*Invoke) isTerm()      {}
func (*Resume) isTerm()      {}
func (*Unreachable) isTerm() {}
