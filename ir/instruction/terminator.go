package instruction

import (
	"github.com/llir/llvm/ir/types"
	"github.com/llir/llvm/ir/value"
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

func (*Ret) Type() types.Type { panic("Ret.Type: not yet implemented") }
func (*Ret) String() string   { panic("Ret.String: not yet implemented") }

// TODO: Add support for the remaining terminator instructions:
//    http://llvm.org/docs/LangRef.html#br-instruction
//    http://llvm.org/docs/LangRef.html#switch-instruction
//    http://llvm.org/docs/LangRef.html#indirectbr-instruction
//    http://llvm.org/docs/LangRef.html#invoke-instruction
//    http://llvm.org/docs/LangRef.html#resume-instruction
//    http://llvm.org/docs/LangRef.html#unreachable-instruction

type Br struct{}

func (*Br) Type() types.Type { panic("Br.Type: not yet implemented") }
func (*Br) String() string   { panic("Br.String: not yet implemented") }

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
