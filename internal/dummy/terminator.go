package dummy

import (
	"bytes"
	"fmt"

	"github.com/llir/llvm/internal/enc"
	"github.com/llir/llvm/ir"
	"github.com/llir/llvm/ir/constant"
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
	return fmt.Sprintf("br label %s", enc.Local(term.Target()))
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
		enc.Local(term.TargetTrue()),
		enc.Local(term.TargetFalse()))
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

// --- [ switch ] --------------------------------------------------------------

// TermSwitch represents a dummy switch terminator.
type TermSwitch struct {
	// Parent basic block.
	parent *ir.BasicBlock
	// Control variable.
	x value.Value
	// Default target branch.
	targetDefault string
	// Switch cases.
	cases []*Case
}

// TODO: Consider renaming x to control to avoid confusion between term.X() and
// case.X().

// NewSwitch returns a new dummy switch terminator based on the given control
// variable, default target branch and switch cases.
func NewSwitch(x value.Value, targetDefault string, cases ...*Case) *TermSwitch {
	return &TermSwitch{x: x, targetDefault: targetDefault, cases: cases}
}

// String returns the LLVM syntax representation of the terminator.
func (term *TermSwitch) String() string {
	buf := &bytes.Buffer{}
	x := term.X()
	fmt.Fprintf(buf, "switch %s %s, label %s [\n",
		x.Type(),
		x.Ident(),
		enc.Local(term.TargetDefault()))
	for _, c := range term.Cases() {
		x := c.X()
		fmt.Fprintf(buf, "\t\t%s %s, label %s\n",
			x.Type(),
			x.Ident(),
			enc.Local(c.Target()))
	}
	buf.WriteString("\t]")
	return buf.String()
}

// Parent returns the parent basic block of the terminator.
func (term *TermSwitch) Parent() *ir.BasicBlock {
	return term.parent
}

// SetParent sets the parent basic block of the terminator.
func (term *TermSwitch) SetParent(parent *ir.BasicBlock) {
	term.parent = parent
}

// Successors returns the successor basic blocks of the terminator.
func (term *TermSwitch) Successors() []*ir.BasicBlock {
	panic("dummy implementation")
}

// X returns the control variable of the switch terminator.
func (term *TermSwitch) X() value.Value {
	return term.x
}

// TargetDefault returns the default target branch of the switch terminator.
func (term *TermSwitch) TargetDefault() string {
	return term.targetDefault
}

// Cases returns the switch cases of the switch terminator.
func (term *TermSwitch) Cases() []*Case {
	return term.cases
}

// Case represents a case of a switch terminator.
type Case struct {
	// Case comparand.
	x *constant.Int
	// Case target branch.
	target string
}

// NewCase returns a new switch case based on the given case comparand and
// target branch.
func NewCase(x *constant.Int, target string) *Case {
	return &Case{x: x, target: target}
}

// X returns the case comparand.
func (c *Case) X() *constant.Int {
	return c.x
}

// Target returns the case target branch.
func (c *Case) Target() string {
	return c.target
}
