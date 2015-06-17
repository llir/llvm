package ir

import (
	"github.com/llir/llvm/consts"
	"github.com/llir/llvm/types"
	"github.com/llir/llvm/value"
)

// A Terminator is a control flow instruction (e.g. br, ret, …) which terminates
// a basic block.
//
// References:
//    http://llvm.org/docs/LangRef.html#terminator-instructions
type Terminator interface {
	// isTerm ensures that only terminator instructions can be assigned to the
	// Terminator interface.
	isTerm()
}

// =============================================================================
// Terminator Instructions
//
//    ref: http://llvm.org/docs/LangRef.html#terminators
// =============================================================================

// The ReturnInst returns control flow (and optionally a value) from a function
// back to the caller.
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
//    http://llvm.org/docs/LangRef.html#i-ret
type ReturnInst struct {
	// Return type.
	Type types.Type
	// Return value; or nil in case of a void return.
	Val value.Value
}

// The CondBranchInst transfers control flow to one of two basic blocks in the
// current function based on a boolean branching condition.
//
// Syntax:
//    br i1 <Cond>, label <TargetTrue>, label <TargetFalse>
//
// Semantics:
//    if (Cond) { goto TargetTrue } else { goto TargetFalse }
//
// References:
//    http://llvm.org/docs/LangRef.html#i-br
type CondBranchInst struct {
	// Boolean branching condition.
	Cond value.Value
	// Target branch when the condition evaluates to true.
	True *BasicBlock
	// Target branch when the condition evaluates to false.
	False *BasicBlock
}

// The BranchInst transfers control flow to a basic block in the current
// function.
//
// Syntax:
//    br label <Target>
//
// Semantics:
//    goto Target;
//
// References:
//    http://llvm.org/docs/LangRef.html#i-br
type BranchInst struct {
	// Target branch.
	Target *BasicBlock
}

// The SwitchInst transfers control flow to one of several basic blocks in the
// current function.
//
// Syntax:
//    switch <IntType> <Val>, label <TargetDefault> [ <IntType> <Const1>, label <Target1> ... ]
//
// Semantics:
//    switch (Val) {
//       case Const1:
//          // Target1
//       default:
//          // TargetDefault
//    }
//
// References:
//    http://llvm.org/docs/LangRef.html#i-switch
type SwitchInst struct {
	// TODO(u): Restrict Type to IntType, Value to IntValue and Constant to IntConstant.

	// Comparasion type.
	Type types.Type
	// Comparasion value.
	Val value.Value
	// Default target.
	Default *BasicBlock
	// Switch cases.
	Cases []struct {
		// Case value.
		Val consts.Constant
		// Case target.
		Target *BasicBlock
	}
}

// TODO(u): Add the following terminator instructions:
//    - indirectbr
//    - invoke
//    - resume
//    - unreachable

// The UnreachableInst indicates that a particular portion of the code is not
// reachable (e.g. code after a no-return function).
//
// Syntax:
//    unreachable
//
// Semantics:
//    // No defined semantics.
//
// References:
//    http://llvm.org/docs/LangRef.html#i-unreachable
type UnreachableInst struct {
}

// isTerm ensures that only terminator instructions can be assigned to the
// Terminator interface.
func (ReturnInst) isTerm()     {}
func (CondBranchInst) isTerm() {}
func (BranchInst) isTerm()     {}
func (SwitchInst) isTerm()     {}
