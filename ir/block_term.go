package ir

import (
	"github.com/llir/l/ir/instruction"
	"github.com/llir/l/ir/ll"
	"github.com/llir/l/ir/value"
)

// === [ Terminators ] =========================================================

// --- [ ret ] -----------------------------------------------------------------

// NewRet returns a new ret terminator based on the given return value. A nil
// return value indicates a void return.
func (block *BasicBlock) NewRet(x value.Value) *instruction.Ret {
	term := instruction.NewRet(x)
	block.Term = term
	return term
}

// --- [ br ] ------------------------------------------------------------------

// NewBr returns a new unconditional br terminator based on the given target
// basic block.
func (block *BasicBlock) NewBr(target value.Value) *instruction.Br {
	term := instruction.NewBr(target)
	block.Term = term
	return term
}

// --- [ conditional br ] ------------------------------------------------------

// NewCondBr returns a new conditional br terminator based on the given
// branching condition and conditional target basic blocks.
func (block *BasicBlock) NewCondBr(cond, targetTrue, targetFalse value.Value) *instruction.CondBr {
	term := instruction.NewCondBr(cond, targetTrue, targetFalse)
	block.Term = term
	return term
}

// --- [ switch ] --------------------------------------------------------------

// NewSwitch returns a new switch terminator based on the given control
// variable, default target basic block and switch cases.
func (block *BasicBlock) NewSwitch(x, targetDefault value.Value, cases ...*ll.Case) *instruction.Switch {
	term := instruction.NewSwitch(x, targetDefault, cases...)
	block.Term = term
	return term
}

// --- [ indirectbr ] ----------------------------------------------------------

// NewIndirectBr returns a new indirectbr terminator based on the given target
// address (derived from a blockaddress constant) and set of valid target basic
// blocks.
func (block *BasicBlock) NewIndirectBr(addr value.Value, validTargets ...value.Value) *instruction.IndirectBr {
	term := instruction.NewIndirectBr(addr, validTargets...)
	block.Term = term
	return term
}

// --- [ invoke ] --------------------------------------------------------------

// NewInvoke returns a new invoke terminator based on the given callee, function
// arguments and control flow return points for normal and exceptional
// execution.
func (block *BasicBlock) NewInvoke(callee value.Value, args []ll.Arg, normal, exception value.Value) *instruction.Invoke {
	term := instruction.NewInvoke(callee, args, normal, exception)
	block.Term = term
	return term
}

// --- [ resume ] --------------------------------------------------------------

// NewResume returns a new resume terminator based on the given exception
// argument to propagate.
func (block *BasicBlock) NewResume(x value.Value) *instruction.Resume {
	term := instruction.NewResume(x)
	block.Term = term
	return term
}

// --- [ catchswitch ] ---------------------------------------------------------

// NewCatchSwitch returns a new catchswitch terminator based on the given
// exception scope, exception handlers and unwind target.
func (block *BasicBlock) NewCatchSwitch(scope ll.ExceptionScope, handlers []value.Value, unwindTarget ll.UnwindTarget) *instruction.CatchSwitch {
	term := instruction.NewCatchSwitch(scope, handlers, unwindTarget)
	block.Term = term
	return term
}

// --- [ catchret ] ------------------------------------------------------------

// NewCatchRet returns a new catchret terminator based on the given exit
// catchpad and target basic block.
func (block *BasicBlock) NewCatchRet(from, to value.Value) *instruction.CatchRet {
	term := instruction.NewCatchRet(from, to)
	block.Term = term
	return term
}

// --- [ cleanupret ] ----------------------------------------------------------

// NewCleanupRet returns a new cleanupret terminator based on the given exit
// cleanuppad and unwind target.
func (block *BasicBlock) NewCleanupRet(from value.Value, to ll.UnwindTarget) *instruction.CleanupRet {
	term := instruction.NewCleanupRet(from, to)
	block.Term = term
	return term
}

// --- [ unreachable ] ---------------------------------------------------------

// NewUnreachable returns a new unreachable terminator.
func (block *BasicBlock) NewUnreachable() *instruction.Unreachable {
	term := instruction.NewUnreachable()
	block.Term = term
	return term
}
