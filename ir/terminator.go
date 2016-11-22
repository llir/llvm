package ir

import (
	"fmt"

	"github.com/llir/llvm/ir/value"
)

// A Terminator represents an LLVM IR terminator.
//
// Terminator may have one of the following underlying types.
//
//    *ir.TermRet
//    *ir.TermBr
//    *ir.TermCondBr
//    TODO
type Terminator interface {
	// LLVMString returns the LLVM syntax representation of the terminator.
	LLVMString() string
	// Parent returns the parent basic block of the terminator.
	Parent() *BasicBlock
	// Successors returns the successor basic blocks of the terminator.
	Successors() []*BasicBlock
}

// --- [ ret ] -----------------------------------------------------------------

// TermRet represents a ret terminator.
//
// References:
//    http://llvm.org/docs/LangRef.html#ret-instruction
type TermRet struct {
	// Parent basic block.
	parent *BasicBlock
	// Return value; or nil if "void" return.
	x value.Value
}

// NewRet returns a new ret terminator based on the given return value. A nil
// return value indicates a "void" return.
func NewRet(x value.Value) *TermRet {
	return &TermRet{x: x}
}

// LLVMString returns the LLVM syntax representation of the terminator.
func (t *TermRet) LLVMString() string {
	if x, ok := t.X(); ok {
		return fmt.Sprintf("ret %s %s",
			x.Type().LLVMString(),
			x.Ident())
	}
	return "ret void"
}

// Parent returns the parent basic block of the terminator.
func (t *TermRet) Parent() *BasicBlock {
	return t.parent
}

// SetParent sets the parent basic block of the terminator.
func (t *TermRet) SetParent(parent *BasicBlock) {
	t.parent = parent
}

// Successors returns the successor basic blocks of the terminator.
func (t *TermRet) Successors() []*BasicBlock {
	// ret terminators have no successors.
	return nil
}

// X returns the return value of the ret terminator and a boolean indicating if
// a return value was present.
func (t *TermRet) X() (value.Value, bool) {
	if t.x != nil {
		return t.x, true
	}
	return nil, false
}

// --- [ br ] ------------------------------------------------------------------

// TermBr represents an unconditional br terminator.
//
// References:
//    http://llvm.org/docs/LangRef.html#br-instruction
type TermBr struct {
	// Parent basic block.
	parent *BasicBlock
	// Target branch.
	target *BasicBlock
}

// NewBr returns a new unconditional br terminator based on the given target
// branch.
func NewBr(target *BasicBlock) *TermBr {
	return &TermBr{target: target}
}

// LLVMString returns the LLVM syntax representation of the terminator.
func (t *TermBr) LLVMString() string {
	return fmt.Sprintf("br label %s", t.Target().Ident())
}

// Parent returns the parent basic block of the terminator.
func (t *TermBr) Parent() *BasicBlock {
	return t.parent
}

// SetParent sets the parent basic block of the terminator.
func (t *TermBr) SetParent(parent *BasicBlock) {
	t.parent = parent
}

// Successors returns the successor basic blocks of the terminator.
func (t *TermBr) Successors() []*BasicBlock {
	return []*BasicBlock{t.target}
}

// Target returns the target branch of the br terminator.
func (t *TermBr) Target() *BasicBlock {
	return t.target
}

// --- [ conditional br ] ------------------------------------------------------

// TermCondBr represents a conditional br terminator.
//
// References:
//    http://llvm.org/docs/LangRef.html#br-instruction
type TermCondBr struct {
	// Parent basic block.
	parent *BasicBlock
	// Branching condition.
	cond value.Value
	// Target branch when condition is true.
	targetTrue *BasicBlock
	// Target branch when condition is false.
	targetFalse *BasicBlock
}

// NewCondBr returns a new conditional br terminator based on the given
// branching condition and conditional target branches.
func NewCondBr(cond value.Value, targetTrue, targetFalse *BasicBlock) *TermCondBr {
	return &TermCondBr{cond: cond, targetTrue: targetTrue, targetFalse: targetFalse}
}

// LLVMString returns the LLVM syntax representation of the terminator.
func (t *TermCondBr) LLVMString() string {
	return fmt.Sprintf("br i1 %s, label %s, label %s",
		t.Cond().Ident(),
		t.TargetTrue().Ident(),
		t.TargetFalse().Ident())
}

// Parent returns the parent basic block of the terminator.
func (t *TermCondBr) Parent() *BasicBlock {
	return t.parent
}

// SetParent sets the parent basic block of the terminator.
func (t *TermCondBr) SetParent(parent *BasicBlock) {
	t.parent = parent
}

// Successors returns the successor basic blocks of the terminator.
func (t *TermCondBr) Successors() []*BasicBlock {
	return []*BasicBlock{t.targetTrue, t.targetFalse}
}

// Cond returns the branching condition of the br terminator.
func (t *TermCondBr) Cond() value.Value {
	return t.cond
}

// TargetTrue returns the target branch when condition is true of the br
// terminator.
func (t *TermCondBr) TargetTrue() *BasicBlock {
	return t.targetTrue
}

// TargetFalse returns the target branch when condition is false of the br
// terminator.
func (t *TermCondBr) TargetFalse() *BasicBlock {
	return t.targetFalse
}

// --- [ switch ] --------------------------------------------------------------

// TODO: Add support for switch.

// --- [ indirectbr ] ----------------------------------------------------------

// --- [ invoke ] --------------------------------------------------------------

// --- [ resume ] --------------------------------------------------------------

// --- [ catchswitch ] ---------------------------------------------------------

// --- [ catchret ] ------------------------------------------------------------

// --- [ cleanupret ] ----------------------------------------------------------

// --- [ unreachable ] ---------------------------------------------------------

// TODO: Add support for unreachable.
