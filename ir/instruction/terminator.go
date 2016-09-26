package instruction

import (
	"fmt"

	"github.com/llir/llvm/ir/internal/enc"
	"github.com/llir/llvm/ir/types"
	"github.com/llir/llvm/ir/value"
	"github.com/mewkiz/pkg/errutil"
)

// --- [ ret ] -----------------------------------------------------------------

// Ret represents a ret instruction, which returns control flow, and optionally
// a value, from a callee back to its caller.
//
// References:
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
	val := term.Value()
	if val == nil {
		return "ret void"
	}
	return fmt.Sprintf("ret %v %v", val.Type(), val.ValueString())
}

// --- [ jmp ] -----------------------------------------------------------------

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

	// TODO: Re-enable type checking. Figure out how to handle dummy types.
	//if !types.IsLabel(target.Type()) {
	//	return nil, errutil.Newf("invalid target type; expected *types.Label, got %T", target.Type())
	//}
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

// --- [ br ] ------------------------------------------------------------------

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
	cond := term.Cond()
	trueBranch, falseBranch := term.TrueBranch(), term.FalseBranch()
	// TODO: Make use of ValueString rather than enc.Local for trueBranch and
	// falseBranch.
	return fmt.Sprintf("br %s %s, label %s, label %s", cond.Type(), cond.ValueString(), enc.Local(trueBranch.Name()), enc.Local(falseBranch.Name()))
}

// --- [ switch ] --------------------------------------------------------------

// TODO: Implement Switch.

// Switch represents a switch instruction.
//
// References:
//    http://llvm.org/docs/LangRef.html#switch-instruction
type Switch struct{}

// String returns the string representation of the instruction.
func (term *Switch) String() string { panic("Switch.String: not yet implemented") }

// --- [ indirectbr ] ----------------------------------------------------------

// TODO: Implement IndirectBr.

// IndirectBr represents an indirectbr instruction.
//
// References:
//    http://llvm.org/docs/LangRef.html#indirectbr-instruction
type IndirectBr struct{}

// String returns the string representation of the instruction.
func (term *IndirectBr) String() string { panic("IndirectBr.String: not yet implemented") }

// --- [ invoke ] --------------------------------------------------------------

// TODO: Implement Invoke.

// Invoke represents an invoke instruction.
//
// References:
//    http://llvm.org/docs/LangRef.html#invoke-instruction
type Invoke struct{}

// String returns the string representation of the instruction.
func (term *Invoke) String() string { panic("Invoke.String: not yet implemented") }

// --- [ resume ] --------------------------------------------------------------

// TODO: Implement Resume.

// Resume represents a resume instruction.
//
// References:
//    http://llvm.org/docs/LangRef.html#resume-instruction
type Resume struct{}

// String returns the string representation of the instruction.
func (term *Resume) String() string { panic("Resume.String: not yet implemented") }

// --- [ catchswitch ] ---------------------------------------------------------

// TODO: Implement CatchSwitch.

// CatchSwitch represents a catchswitch instruction.
//
// References:
//    http://llvm.org/docs/LangRef.html#catchswitch-instruction
type CatchSwitch struct{}

// String returns the string representation of the instruction.
func (term *CatchSwitch) String() string { panic("CatchSwitch.String: not yet implemented") }

// --- [ catchret ] ------------------------------------------------------------

// TODO: Implement CatchRet.

// CatchRet represents a catchret instruction.
//
// References:
//    http://llvm.org/docs/LangRef.html#catchret-instruction
type CatchRet struct{}

// String returns the string representation of the instruction.
func (term *CatchRet) String() string { panic("CatchRet.String: not yet implemented") }

// --- [ cleanupret ] ----------------------------------------------------------

// TODO: Implement CleanupRet.

// CleanupRet represents a cleanupret instruction.
//
// References:
//    http://llvm.org/docs/LangRef.html#cleanupret-instruction
type CleanupRet struct{}

// String returns the string representation of the instruction.
func (term *CleanupRet) String() string { panic("CleanupRet.String: not yet implemented") }

// --- [ unreachable ] ---------------------------------------------------------

// TODO: Implement Unreachable.

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
func (*CatchSwitch) isTerm() {}
func (*CatchRet) isTerm()    {}
func (*CleanupRet) isTerm()  {}
func (*Unreachable) isTerm() {}
