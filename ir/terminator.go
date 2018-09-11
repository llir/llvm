package ir

import (
	"github.com/llir/l/ir/ll"
	"github.com/llir/l/ir/value"
)

// === [ Terminators ] =========================================================

type Terminator interface {
	// Succs returns the successor basic blocks of the terminator.
	Succs() []*BasicBlock
}

// --- [ ret ] -----------------------------------------------------------------

// Ret is an LLVM IR ret terminator.
type Ret struct {
	// Return value; or nil if void return.
	X value.Value
}

// NewRet returns a new ret terminator based on the given return value. A nil
// return value indicates a void return.
func NewRet(x value.Value) *Ret {
	return &Ret{X: x}
}

// Succs returns the successor basic blocks of the terminator.
func (*Ret) Succs() []*BasicBlock {
	// no successors.
	return nil
}

// --- [ br ] ------------------------------------------------------------------

// Br is an unconditional LLVM IR br terminator.
type Br struct {
	// Target basic block.
	Target *BasicBlock
}

// NewBr returns a new unconditional br terminator based on the given target
// basic block.
func NewBr(target *BasicBlock) *Br {
	return &Br{Target: target}
}

// Succs returns the successor basic blocks of the terminator.
func (term *Br) Succs() []*BasicBlock {
	return []*BasicBlock{term.Target}
}

// --- [ conditional br ] ------------------------------------------------------

// CondBr is a conditional LLVM IR br terminator.
type CondBr struct {
	// Branching condition.
	Cond value.Value
	// True condition target basic block.
	TargetTrue *BasicBlock
	// False condition target basic block.
	TargetFalse *BasicBlock
}

// NewCondBr returns a new conditional br terminator based on the given
// branching condition and conditional target basic blocks.
func NewCondBr(cond value.Value, targetTrue, targetFalse *BasicBlock) *CondBr {
	return &CondBr{Cond: cond, TargetTrue: targetTrue, TargetFalse: targetFalse}
}

// Succs returns the successor basic blocks of the terminator.
func (term *CondBr) Succs() []*BasicBlock {
	return []*BasicBlock{term.TargetTrue, term.TargetFalse}
}

// --- [ switch ] --------------------------------------------------------------

// Switch is an LLVM IR switch terminator.
type Switch struct {
	// Control variable.
	X value.Value
	// Default target basic block.
	TargetDefault *BasicBlock
	// Switch cases.
	Cases []*Case
}

// NewSwitch returns a new switch terminator based on the given control
// variable, default target basic block and switch cases.
func NewSwitch(x value.Value, targetDefault *BasicBlock, cases ...*Case) *Switch {
	return &Switch{X: x, TargetDefault: targetDefault, Cases: cases}
}

// Succs returns the successor basic blocks of the terminator.
func (term *Switch) Succs() []*BasicBlock {
	succs := make([]*BasicBlock, 0, 1+len(term.Cases))
	succs = append(succs, term.TargetDefault)
	for _, c := range term.Cases {
		succs = append(succs, c.Target)
	}
	return succs
}

// Case is a switch case.
type Case struct {
	// Case comparand.
	X Constant // integer constant or interger constant expression
	// Case target basic block.
	Target *BasicBlock
}

// NewCase returns a new switch case based on the given case comparand and
// target basic block.
func NewCase(x Constant, target *BasicBlock) *Case {
	return &Case{X: x, Target: target}
}

// --- [ indirectbr ] ----------------------------------------------------------

// IndirectBr is an LLVM IR indirectbr terminator.
type IndirectBr struct {
	// Target address.
	Addr *BlockAddress
	// Set of valid target basic blocks.
	ValidTargets []*BasicBlock
}

// NewIndirectBr returns a new indirectbr terminator based on the given target
// address (derived from a blockaddress constant) and set of valid target basic
// blocks.
func NewIndirectBr(addr *BlockAddress, validTargets ...*BasicBlock) *IndirectBr {
	return &IndirectBr{Addr: addr, ValidTargets: validTargets}
}

// Succs returns the successor basic blocks of the terminator.
func (term *IndirectBr) Succs() []*BasicBlock {
	return term.ValidTargets
}

// --- [ invoke ] --------------------------------------------------------------

// Invoke is an LLVM IR invoke terminator.
type Invoke struct {
	// Name of local variable associated with the result.
	LocalName string
	// Callee.
	// TODO: specify the set of underlying types of Callee.
	Callee value.Value
	// Function arguments.
	Args []ll.Arg
	// Normal control flow return point.
	Normal *BasicBlock
	// Exception control flow return point.
	Exception *BasicBlock
}

// NewInvoke returns a new invoke terminator based on the given callee, function
// arguments and control flow return points for normal and exceptional
// execution.
func NewInvoke(callee value.Value, args []ll.Arg, normal, exception *BasicBlock) *Invoke {
	return &Invoke{Callee: callee, Args: args, Normal: normal, Exception: exception}
}

// Succs returns the successor basic blocks of the terminator.
func (term *Invoke) Succs() []*BasicBlock {
	return []*BasicBlock{term.Normal, term.Exception}
}

// --- [ resume ] --------------------------------------------------------------

// Resume is an LLVM IR resume terminator.
type Resume struct {
	// Exception argument to propagate.
	X value.Value
}

// NewResume returns a new resume terminator based on the given exception
// argument to propagate.
func NewResume(x value.Value) *Resume {
	return &Resume{X: x}
}

// Succs returns the successor basic blocks of the terminator.
func (term *Resume) Succs() []*BasicBlock {
	// no successors.
	return nil
}

// --- [ catchswitch ] ---------------------------------------------------------

// CatchSwitch is an LLVM IR catchswitch terminator.
type CatchSwitch struct {
	// Name of local variable associated with the result.
	LocalName string
	// Exception scope.
	Scope ll.ExceptionScope // TODO: rename to Parent? rename to From?
	// Exception handlers.
	Handlers []*BasicBlock
	// Unwind target; basic block or caller function.
	UnwindTarget ll.UnwindTarget // TODO: rename to To?
}

// NewCatchSwitch returns a new catchswitch terminator based on the given
// exception scope, exception handlers and unwind target.
func NewCatchSwitch(scope ll.ExceptionScope, handlers []*BasicBlock, unwindTarget ll.UnwindTarget) *CatchSwitch {
	return &CatchSwitch{Scope: scope, Handlers: handlers, UnwindTarget: unwindTarget}
}

// Succs returns the successor basic blocks of the terminator.
func (term *CatchSwitch) Succs() []*BasicBlock {
	if unwindTarget, ok := term.UnwindTarget.(*BasicBlock); ok {
		return append(term.Handlers, unwindTarget)
	}
	return term.Handlers
}

// --- [ catchret ] ------------------------------------------------------------

// CatchRet is an LLVM IR catchret terminator.
type CatchRet struct {
	// Exit catchpad.
	From *CatchPad
	// Target basic block to transfer control flow to.
	To *BasicBlock
}

// NewCatchRet returns a new catchret terminator based on the given exit
// catchpad and target basic block.
func NewCatchRet(from *CatchPad, to *BasicBlock) *CatchRet {
	return &CatchRet{From: from, To: to}
}

// Succs returns the successor basic blocks of the terminator.
func (term *CatchRet) Succs() []*BasicBlock {
	return []*BasicBlock{term.To}
}

// --- [ cleanupret ] ----------------------------------------------------------

// CleanupRet is an LLVM IR cleanupret terminator.
type CleanupRet struct {
	// Exit cleanuppad.
	From *CleanupPad
	// Unwind target; basic block or caller function.
	To ll.UnwindTarget
}

// NewCleanupRet returns a new cleanupret terminator based on the given exit
// cleanuppad and unwind target.
func NewCleanupRet(from *CleanupPad, to ll.UnwindTarget) *CleanupRet {
	return &CleanupRet{From: from, To: to}
}

// Succs returns the successor basic blocks of the terminator.
func (term *CleanupRet) Succs() []*BasicBlock {
	if unwindTarget, ok := term.To.(*BasicBlock); ok {
		return []*BasicBlock{unwindTarget}
	}
	return nil
}

// --- [ unreachable ] ---------------------------------------------------------

// Unreachable is an LLVM IR unreachable terminator.
type Unreachable struct {
}

// NewUnreachable returns a new unreachable terminator.
func NewUnreachable() *Unreachable {
	return &Unreachable{}
}

// Succs returns the successor basic blocks of the terminator.
func (term *Unreachable) Succs() []*BasicBlock {
	// no successors.
	return nil
}
