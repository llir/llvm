package ir

import (
	"github.com/llir/l/ir/ll"
	"github.com/llir/l/ir/types"
	"github.com/llir/l/ir/value"
)

// === [ Terminators ] =========================================================

// Terminator is an LLVM IR terminator instruction (a control flow instruction).
//
// A Terminator has one of the following underlying types.
//
// Terminators
//
// https://llvm.org/docs/LangRef.html#terminator-instructions
//
//    *ir.TermRet           // https://godoc.org/github.com/llir/l/ir#TermRet
//    *ir.TermBr            // https://godoc.org/github.com/llir/l/ir#TermBr
//    *ir.TermCondBr        // https://godoc.org/github.com/llir/l/ir#TermCondBr
//    *ir.TermSwitch        // https://godoc.org/github.com/llir/l/ir#TermSwitch
//    *ir.TermIndirectBr    // https://godoc.org/github.com/llir/l/ir#TermIndirectBr
//    *ir.TermInvoke        // https://godoc.org/github.com/llir/l/ir#TermInvoke
//    *ir.TermResume        // https://godoc.org/github.com/llir/l/ir#TermResume
//    *ir.TermCatchSwitch   // https://godoc.org/github.com/llir/l/ir#TermCatchSwitch
//    *ir.TermCatchRet      // https://godoc.org/github.com/llir/l/ir#TermCatchRet
//    *ir.TermCleanupRet    // https://godoc.org/github.com/llir/l/ir#TermCleanupRet
//    *ir.TermUnreachable   // https://godoc.org/github.com/llir/l/ir#TermUnreachable
type Terminator interface {
	// Succs returns the successor basic blocks of the terminator.
	Succs() []*BasicBlock
}

// --- [ ret ] -----------------------------------------------------------------

// TermRet is an LLVM IR ret terminator.
type TermRet struct {
	// Return value; or nil if void return.
	X value.Value
}

// NewRet returns a new ret terminator based on the given return value. A nil
// return value indicates a void return.
func NewRet(x value.Value) *TermRet {
	return &TermRet{X: x}
}

// Succs returns the successor basic blocks of the terminator.
func (*TermRet) Succs() []*BasicBlock {
	// no successors.
	return nil
}

// --- [ br ] ------------------------------------------------------------------

// TermBr is an unconditional LLVM IR br terminator.
type TermBr struct {
	// Target basic block.
	Target *BasicBlock
}

// NewBr returns a new unconditional br terminator based on the given target
// basic block.
func NewBr(target *BasicBlock) *TermBr {
	return &TermBr{Target: target}
}

// Succs returns the successor basic blocks of the terminator.
func (term *TermBr) Succs() []*BasicBlock {
	return []*BasicBlock{term.Target}
}

// --- [ conditional br ] ------------------------------------------------------

// TermCondBr is a conditional LLVM IR br terminator.
type TermCondBr struct {
	// Branching condition.
	Cond value.Value
	// True condition target basic block.
	TargetTrue *BasicBlock
	// False condition target basic block.
	TargetFalse *BasicBlock
}

// NewCondBr returns a new conditional br terminator based on the given
// branching condition and conditional target basic blocks.
func NewCondBr(cond value.Value, targetTrue, targetFalse *BasicBlock) *TermCondBr {
	return &TermCondBr{Cond: cond, TargetTrue: targetTrue, TargetFalse: targetFalse}
}

// Succs returns the successor basic blocks of the terminator.
func (term *TermCondBr) Succs() []*BasicBlock {
	return []*BasicBlock{term.TargetTrue, term.TargetFalse}
}

// --- [ switch ] --------------------------------------------------------------

// TermSwitch is an LLVM IR switch terminator.
type TermSwitch struct {
	// Control variable.
	X value.Value
	// Default target basic block.
	TargetDefault *BasicBlock
	// TermSwitch cases.
	Cases []*Case
}

// NewSwitch returns a new switch terminator based on the given control
// variable, default target basic block and switch cases.
func NewSwitch(x value.Value, targetDefault *BasicBlock, cases ...*Case) *TermSwitch {
	return &TermSwitch{X: x, TargetDefault: targetDefault, Cases: cases}
}

// Succs returns the successor basic blocks of the terminator.
func (term *TermSwitch) Succs() []*BasicBlock {
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

// TermIndirectBr is an LLVM IR indirectbr terminator.
type TermIndirectBr struct {
	// Target address.
	Addr *ConstBlockAddress
	// Set of valid target basic blocks.
	ValidTargets []*BasicBlock
}

// NewIndirectBr returns a new indirectbr terminator based on the given target
// address (derived from a blockaddress constant) and set of valid target basic
// blocks.
func NewIndirectBr(addr *ConstBlockAddress, validTargets ...*BasicBlock) *TermIndirectBr {
	return &TermIndirectBr{Addr: addr, ValidTargets: validTargets}
}

// Succs returns the successor basic blocks of the terminator.
func (term *TermIndirectBr) Succs() []*BasicBlock {
	return term.ValidTargets
}

// --- [ invoke ] --------------------------------------------------------------

// TermInvoke is an LLVM IR invoke terminator.
type TermInvoke struct {
	// Name of local variable associated with the result.
	LocalName string
	// Invokee (callee function).
	// TODO: specify the set of underlying types of Invokee.
	Invokee value.Value
	// Function arguments.
	Args []ll.Arg
	// Normal control flow return point.
	Normal *BasicBlock
	// Exception control flow return point.
	Exception *BasicBlock
}

// NewInvoke returns a new invoke terminator based on the given invokee, function
// arguments and control flow return points for normal and exceptional
// execution.
func NewInvoke(invokee value.Value, args []ll.Arg, normal, exception *BasicBlock) *TermInvoke {
	return &TermInvoke{Invokee: invokee, Args: args, Normal: normal, Exception: exception}
}

// Type returns the type of the terminator.
func (term *TermInvoke) Type() types.Type {
	panic("not yet implemented")
}

// Ident returns the identifier associated with the terminator.
func (term *TermInvoke) Ident() string {
	panic("not yet implemented")
}

// Name returns the name of the terminator.
func (term *TermInvoke) Name() string {
	return term.LocalName
}

// SetName sets the name of the terminator.
func (term *TermInvoke) SetName(name string) {
	term.LocalName = name
}

// Succs returns the successor basic blocks of the terminator.
func (term *TermInvoke) Succs() []*BasicBlock {
	return []*BasicBlock{term.Normal, term.Exception}
}

// --- [ resume ] --------------------------------------------------------------

// TermResume is an LLVM IR resume terminator.
type TermResume struct {
	// Exception argument to propagate.
	X value.Value
}

// NewResume returns a new resume terminator based on the given exception
// argument to propagate.
func NewResume(x value.Value) *TermResume {
	return &TermResume{X: x}
}

// Succs returns the successor basic blocks of the terminator.
func (term *TermResume) Succs() []*BasicBlock {
	// no successors.
	return nil
}

// --- [ catchswitch ] ---------------------------------------------------------

// TermCatchSwitch is an LLVM IR catchswitch terminator.
type TermCatchSwitch struct {
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
func NewCatchSwitch(scope ll.ExceptionScope, handlers []*BasicBlock, unwindTarget ll.UnwindTarget) *TermCatchSwitch {
	return &TermCatchSwitch{Scope: scope, Handlers: handlers, UnwindTarget: unwindTarget}
}

// Type returns the type of the terminator.
func (term *TermCatchSwitch) Type() types.Type {
	return types.Token
}

// Ident returns the identifier associated with the terminator.
func (term *TermCatchSwitch) Ident() string {
	panic("not yet implemented")
}

// Name returns the name of the terminator.
func (term *TermCatchSwitch) Name() string {
	return term.LocalName
}

// SetName sets the name of the terminator.
func (term *TermCatchSwitch) SetName(name string) {
	term.LocalName = name
}

// Succs returns the successor basic blocks of the terminator.
func (term *TermCatchSwitch) Succs() []*BasicBlock {
	if unwindTarget, ok := term.UnwindTarget.(*BasicBlock); ok {
		return append(term.Handlers, unwindTarget)
	}
	return term.Handlers
}

// --- [ catchret ] ------------------------------------------------------------

// TermCatchRet is an LLVM IR catchret terminator.
type TermCatchRet struct {
	// Exit catchpad.
	From *InstCatchPad
	// Target basic block to transfer control flow to.
	To *BasicBlock
}

// NewCatchRet returns a new catchret terminator based on the given exit
// catchpad and target basic block.
func NewCatchRet(from *InstCatchPad, to *BasicBlock) *TermCatchRet {
	return &TermCatchRet{From: from, To: to}
}

// Succs returns the successor basic blocks of the terminator.
func (term *TermCatchRet) Succs() []*BasicBlock {
	return []*BasicBlock{term.To}
}

// --- [ cleanupret ] ----------------------------------------------------------

// TermCleanupRet is an LLVM IR cleanupret terminator.
type TermCleanupRet struct {
	// Exit cleanuppad.
	From *InstCleanupPad
	// Unwind target; basic block or caller function.
	To ll.UnwindTarget
}

// NewCleanupRet returns a new cleanupret terminator based on the given exit
// cleanuppad and unwind target.
func NewCleanupRet(from *InstCleanupPad, to ll.UnwindTarget) *TermCleanupRet {
	return &TermCleanupRet{From: from, To: to}
}

// Succs returns the successor basic blocks of the terminator.
func (term *TermCleanupRet) Succs() []*BasicBlock {
	if unwindTarget, ok := term.To.(*BasicBlock); ok {
		return []*BasicBlock{unwindTarget}
	}
	return nil
}

// --- [ unreachable ] ---------------------------------------------------------

// TermUnreachable is an LLVM IR unreachable terminator.
type TermUnreachable struct {
}

// NewUnreachable returns a new unreachable terminator.
func NewUnreachable() *TermUnreachable {
	return &TermUnreachable{}
}

// Succs returns the successor basic blocks of the terminator.
func (term *TermUnreachable) Succs() []*BasicBlock {
	// no successors.
	return nil
}
