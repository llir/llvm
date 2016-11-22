package ir

import (
	"bytes"
	"fmt"

	"github.com/llir/llvm/ir/constant"
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
	// Successors basic blocks.
	successors []*BasicBlock
}

// NewBr returns a new unconditional br terminator based on the given target
// branch.
func NewBr(target *BasicBlock) *TermBr {
	successors := []*BasicBlock{target}
	return &TermBr{target: target, successors: successors}
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
	return t.successors
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
	// Successors basic blocks.
	successors []*BasicBlock
}

// NewCondBr returns a new conditional br terminator based on the given
// branching condition and conditional target branches.
func NewCondBr(cond value.Value, targetTrue, targetFalse *BasicBlock) *TermCondBr {
	successors := []*BasicBlock{targetTrue, targetFalse}
	return &TermCondBr{cond: cond, targetTrue: targetTrue, targetFalse: targetFalse, successors: successors}
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
	return t.successors
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

// TermSwitch represents a switch terminator.
//
// References:
//    http://llvm.org/docs/LangRef.html#switch-instruction
type TermSwitch struct {
	// Parent basic block.
	parent *BasicBlock
	// Control variable.
	x value.Value
	// Default target branch.
	targetDefault *BasicBlock
	// Switch cases.
	cases []*Case
	// Successors basic blocks.
	successors []*BasicBlock
}

// NewSwitch returns a new switch terminator based on the given control
// variable, default target branch and switch cases.
func NewSwitch(x value.Value, targetDefault *BasicBlock, cases ...*Case) *TermSwitch {
	successors := []*BasicBlock{targetDefault}
	for _, c := range cases {
		successors = append(successors, c.target)
	}
	return &TermSwitch{x: x, targetDefault: targetDefault, cases: cases, successors: successors}
}

// LLVMString returns the LLVM syntax representation of the terminator.
func (t *TermSwitch) LLVMString() string {
	buf := &bytes.Buffer{}
	x := t.X()
	fmt.Fprintf(buf, "switch %s %s, label %s [ ",
		x.Type().LLVMString(),
		x.Ident(),
		t.TargetDefault().Ident())
	for i, c := range t.Cases() {
		if i != 0 {
			buf.WriteString("\n\t\t")
		}
		x := c.X()
		fmt.Fprintf(buf, "%s %s, label %s",
			x.Type(),
			x.Ident(),
			c.Target().Ident())
	}
	buf.WriteString(" ]")
	return buf.String()
}

// Parent returns the parent basic block of the terminator.
func (t *TermSwitch) Parent() *BasicBlock {
	return t.parent
}

// SetParent sets the parent basic block of the terminator.
func (t *TermSwitch) SetParent(parent *BasicBlock) {
	t.parent = parent
}

// Successors returns the successor basic blocks of the terminator.
func (t *TermSwitch) Successors() []*BasicBlock {
	return t.successors
}

// X returns the control variable of the switch terminator.
func (t *TermSwitch) X() value.Value {
	return t.x
}

// TargetDefault returns the default target branch of the switch terminator.
func (t *TermSwitch) TargetDefault() *BasicBlock {
	return t.targetDefault
}

// Cases returns the switch cases of the switch terminator.
func (t *TermSwitch) Cases() []*Case {
	return t.cases
}

// A Case represents a case of a switch terminator.
type Case struct {
	// Case comparand.
	x *constant.Int
	// Case target branch.
	target *BasicBlock
}

// NewCase returns a new switch case based on the given case comparand and
// target branch.
func NewCase(x *constant.Int, target *BasicBlock) *Case {
	return &Case{x: x, target: target}
}

// X returns the case comparand.
func (c *Case) X() *constant.Int {
	return c.x
}

// Target returns the case target branch.
func (c *Case) Target() *BasicBlock {
	return c.target
}

// --- [ indirectbr ] ----------------------------------------------------------

// --- [ invoke ] --------------------------------------------------------------

// --- [ resume ] --------------------------------------------------------------

// --- [ catchswitch ] ---------------------------------------------------------

// --- [ catchret ] ------------------------------------------------------------

// --- [ cleanupret ] ----------------------------------------------------------

// --- [ unreachable ] ---------------------------------------------------------

// TODO: Add support for unreachable.
