// TODO: Consider removing the Type method from terminator instructions, as it
// should have become unnecessary when ValueInst was introduced. Also consider
// removing the Type method from `store` and `fence`, the two non-terminator
// instructions which are not value instructions.

package instruction

import (
	"fmt"

	"github.com/llir/llvm/asm"
	"github.com/llir/llvm/ir/types"
	"github.com/llir/llvm/ir/value"
	"github.com/mewkiz/pkg/errutil"
)

// A Terminator is a control flow instruction (e.g. br, ret, …) which terminates
// a basic block.
//
// References:
//    http://llvm.org/docs/LangRef.html#terminator-instructions
type Terminator interface {
	value.Value
	// isTerm ensures that only terminator instructions can be assigned to the
	// Terminator interface.
	isTerm()
}

// Make sure that each terminator instruction implements the Terminator
// interface.
var (
	_ Terminator = &Ret{}
	_ Terminator = &Br{}
	_ Terminator = &Switch{}
	_ Terminator = &IndirectBr{}
	_ Terminator = &Invoke{}
	_ Terminator = &Resume{}
	_ Terminator = &Unreachable{}
)

// A Ret instruction returns control flow (and optionally a value) from a callee
// back to its caller.
//
// Syntax:
//    ret <Type> <Val>
//    ret void
//
// Semantics:
//    return Val;
//    return;
//
// Reference:
//    http://llvm.org/docs/LangRef.html#ret-instruction
type Ret struct {
	// Return type.
	typ types.Type
	// Return value; or nil in case of a void return.
	val value.Value
}

// NewRet returns a new ret instruction based on the given return type and
// value. A nil value indicates a "void" return instruction.
func NewRet(typ types.Type, val value.Value) (*Ret, error) {
	// Sanity check.
	switch {
	case typ.Equal(types.NewVoid()):
		// Void return.
		if val != nil {
			return nil, errutil.Newf(`expected no return value for return type "void"; got %q`, val)
		}
	default:
		// Value return.
		if val == nil {
			return nil, errutil.Newf(`expected return value for return type %q; got nil`, typ)
		}
		if valTyp := val.Type(); !typ.Equal(valTyp) {
			return nil, errutil.Newf("type mismatch between return type %q and return value %q", typ, valTyp)
		}
	}

	return &Ret{typ: typ, val: val}, nil
}

// Type returns the type of the value produced by the instruction.
func (inst *Ret) Type() types.Type {
	return inst.typ
}

// String returns the string representation of the instruction.
func (inst *Ret) String() string {
	if inst.val == nil {
		return fmt.Sprintf("ret %v", inst.typ)
	}
	return fmt.Sprintf("ret %v %v", inst.val.Type(), inst.val)
}

// TODO: Add support for the remaining terminator instructions:
//    http://llvm.org/docs/LangRef.html#switch-instruction
//    http://llvm.org/docs/LangRef.html#indirectbr-instruction
//    http://llvm.org/docs/LangRef.html#invoke-instruction
//    http://llvm.org/docs/LangRef.html#resume-instruction
//    http://llvm.org/docs/LangRef.html#unreachable-instruction

// Jmp represents an unconditional branch instruction.
type Jmp struct {
	// Basic block label name of the target branch.
	target string
}

// NewJmp returns a new unconditional branch instruction based on the given
// target branch.
func NewJmp(target string) (*Jmp, error) {
	return &Jmp{target: target}, nil
}

// Target returns the basic block label name of the target branch.
func (inst *Jmp) Target() string {
	return inst.target
}

// Type returns the type of the value produced by the instruction.
func (*Jmp) Type() types.Type {
	return types.NewVoid()
}

// String returns the string representation of the instruction.
func (inst *Jmp) String() string {
	return fmt.Sprintf("br label %s", asm.EncLocal(inst.target))
}

// Br represents a conditional branch instruction.
type Br struct {
	// Branching condition.
	cond value.Value
	// Basic block label name of the true target branch.
	trueBranch string
	// Basic block label name of the false target branch.
	falseBranch string
}

// NewBr returns a new conditional branch instruction based on the given
// branching condition, and the true and false target branches.
func NewBr(cond value.Value, trueBranch, falseBranch string) (*Br, error) {
	if !types.Equal(cond.Type(), types.I1) {
		return nil, errutil.Newf("conditional type mismatch; expected i1, got %v", cond.Type())
	}
	return &Br{cond: cond, trueBranch: trueBranch, falseBranch: falseBranch}, nil
}

// Cond returns the branching condition of the instruction
func (inst *Br) Cond() value.Value {
	return inst.cond
}

// TrueBranch returns the basic block label name of the true target branch.
func (inst *Br) TrueBranch() string {
	return inst.trueBranch
}

// FalseBranch returns the basic block label name of the false target branch.
func (inst *Br) FalseBranch() string {
	return inst.falseBranch
}

// Type returns the type of the value produced by the instruction.
func (*Br) Type() types.Type {
	return types.NewVoid()
}

// String returns the string representation of the instruction.
func (inst *Br) String() string {
	return fmt.Sprintf("br %s %s, label %s, label %s", inst.cond.Type(), inst.cond, asm.EncLocal(inst.trueBranch), asm.EncLocal(inst.falseBranch))
}

type Switch struct{}

func (*Switch) Type() types.Type { panic("Switch.Type: not yet implemented") }
func (*Switch) String() string   { panic("Switch.String: not yet implemented") }

type IndirectBr struct{}

func (*IndirectBr) Type() types.Type { panic("IndirectBr.Type: not yet implemented") }
func (*IndirectBr) String() string   { panic("IndirectBr.String: not yet implemented") }

type Invoke struct{}

func (*Invoke) Type() types.Type { panic("Invoke.Type: not yet implemented") }
func (*Invoke) String() string   { panic("Invoke.String: not yet implemented") }

type Resume struct{}

func (*Resume) Type() types.Type { panic("Resume.Type: not yet implemented") }
func (*Resume) String() string   { panic("Resume.String: not yet implemented") }

type Unreachable struct{}

func (*Unreachable) Type() types.Type { panic("Unreachable.Type: not yet implemented") }
func (*Unreachable) String() string   { panic("Unreachable.String: not yet implemented") }

// isTerm ensures that only terminator instructions can be assigned to the
// Terminator interface.
func (*Ret) isTerm()         {}
func (*Br) isTerm()          {}
func (*Switch) isTerm()      {}
func (*IndirectBr) isTerm()  {}
func (*Invoke) isTerm()      {}
func (*Resume) isTerm()      {}
func (*Unreachable) isTerm() {}
