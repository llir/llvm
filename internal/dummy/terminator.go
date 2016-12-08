package dummy

import (
	"fmt"

	"github.com/llir/llvm/ir"
	"github.com/llir/llvm/ir/value"
)

// --- [ br ] ------------------------------------------------------------------

// TermBr represents a dummy unconditional br terminator.
type TermBr struct {
	// Parent basic block.
	parent *ir.BasicBlock
	// Target branch.
	target string
}

// NewBr returns a new dummy unconditional br terminator based on the given
// target branch.
func NewBr(target string) *TermBr {
	return &TermBr{target: target}
}

// String returns the LLVM syntax representation of the terminator.
func (term *TermBr) String() string {
	return fmt.Sprintf("br label %s", term.Target())
}

// Parent returns the parent basic block of the terminator.
func (term *TermBr) Parent() *ir.BasicBlock {
	return term.parent
}

// SetParent sets the parent basic block of the terminator.
func (term *TermBr) SetParent(parent *ir.BasicBlock) {
	term.parent = parent
}

// Successors returns the successor basic blocks of the terminator.
func (term *TermBr) Successors() []*ir.BasicBlock {
	panic("dummy implementation")
}

// Target returns the target branch of the br terminator.
func (term *TermBr) Target() string {
	return term.target
}

// --- [ conditional br ] ------------------------------------------------------

// TermCondBr represents a dummy conditional br terminator.
type TermCondBr struct {
	// Parent basic block.
	parent *ir.BasicBlock
	// Branching condition.
	cond value.Value
	// Target branch when condition is true.
	targetTrue string
	// Target branch when condition is false.
	targetFalse string
}

// NewCondBr returns a new dummy conditional br terminator based on the given
// branching condition and conditional target branches.
func NewCondBr(cond value.Value, targetTrue, targetFalse string) *TermCondBr {
	return &TermCondBr{cond: cond, targetTrue: targetTrue, targetFalse: targetFalse}
}

// String returns the LLVM syntax representation of the terminator.
func (term *TermCondBr) String() string {
	return fmt.Sprintf("br i1 %s, label %s, label %s",
		term.Cond().Ident(),
		term.TargetTrue(),
		term.TargetFalse())
}

// Parent returns the parent basic block of the terminator.
func (term *TermCondBr) Parent() *ir.BasicBlock {
	return term.parent
}

// SetParent sets the parent basic block of the terminator.
func (term *TermCondBr) SetParent(parent *ir.BasicBlock) {
	term.parent = parent
}

// Successors returns the successor basic blocks of the terminator.
func (term *TermCondBr) Successors() []*ir.BasicBlock {
	panic("dummy implementation")
}

// Cond returns the branching condition of the br terminator.
func (term *TermCondBr) Cond() value.Value {
	return term.cond
}

// TargetTrue returns the target branch when condition is true of the br
// terminator.
func (term *TermCondBr) TargetTrue() string {
	return term.targetTrue
}

// TargetFalse returns the target branch when condition is false of the br
// terminator.
func (term *TermCondBr) TargetFalse() string {
	return term.targetFalse
}
