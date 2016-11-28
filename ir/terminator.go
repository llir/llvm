// === [ Terminators ] =========================================================
//
// References:
//    http://llvm.org/docs/LangRef.html#terminator-instructions

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
// Terminators
//
// http://llvm.org/docs/LangRef.html#terminator-instructions
//
//    *ir.TermRet           (https://godoc.org/github.com/llir/llvm/ir#TermRet)
//    *ir.TermBr            (https://godoc.org/github.com/llir/llvm/ir#TermBr)
//    *ir.TermCondBr        (https://godoc.org/github.com/llir/llvm/ir#TermCondBr)
//    *ir.TermSwitch        (https://godoc.org/github.com/llir/llvm/ir#TermSwitch)
//    *ir.TermUnreachable   (https://godoc.org/github.com/llir/llvm/ir#TermUnreachable)
type Terminator interface {
	Instruction
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

// String returns the LLVM syntax representation of the terminator.
func (term *TermRet) String() string {
	if x, ok := term.X(); ok {
		return fmt.Sprintf("ret %s %s",
			x.Type(),
			x.Ident())
	}
	return "ret void"
}

// Parent returns the parent basic block of the terminator.
func (term *TermRet) Parent() *BasicBlock {
	return term.parent
}

// SetParent sets the parent basic block of the terminator.
func (term *TermRet) SetParent(parent *BasicBlock) {
	term.parent = parent
}

// Successors returns the successor basic blocks of the terminator.
func (term *TermRet) Successors() []*BasicBlock {
	// ret terminators have no successors.
	return nil
}

// X returns the return value of the ret terminator and a boolean indicating if
// a return value was present.
func (term *TermRet) X() (value.Value, bool) {
	if term.x != nil {
		return term.x, true
	}
	return nil, false
}

// SetX sets the return value of the ret terminator.
func (term *TermRet) SetX(x value.Value) {
	term.x = x
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

// String returns the LLVM syntax representation of the terminator.
func (term *TermBr) String() string {
	return fmt.Sprintf("br label %s", term.Target().Ident())
}

// Parent returns the parent basic block of the terminator.
func (term *TermBr) Parent() *BasicBlock {
	return term.parent
}

// SetParent sets the parent basic block of the terminator.
func (term *TermBr) SetParent(parent *BasicBlock) {
	term.parent = parent
}

// Successors returns the successor basic blocks of the terminator.
func (term *TermBr) Successors() []*BasicBlock {
	return term.successors
}

// Target returns the target branch of the br terminator.
func (term *TermBr) Target() *BasicBlock {
	return term.target
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

// String returns the LLVM syntax representation of the terminator.
func (term *TermCondBr) String() string {
	return fmt.Sprintf("br i1 %s, label %s, label %s",
		term.Cond().Ident(),
		term.TargetTrue().Ident(),
		term.TargetFalse().Ident())
}

// Parent returns the parent basic block of the terminator.
func (term *TermCondBr) Parent() *BasicBlock {
	return term.parent
}

// SetParent sets the parent basic block of the terminator.
func (term *TermCondBr) SetParent(parent *BasicBlock) {
	term.parent = parent
}

// Successors returns the successor basic blocks of the terminator.
func (term *TermCondBr) Successors() []*BasicBlock {
	return term.successors
}

// Cond returns the branching condition of the br terminator.
func (term *TermCondBr) Cond() value.Value {
	return term.cond
}

// TargetTrue returns the target branch when condition is true of the br
// terminator.
func (term *TermCondBr) TargetTrue() *BasicBlock {
	return term.targetTrue
}

// TargetFalse returns the target branch when condition is false of the br
// terminator.
func (term *TermCondBr) TargetFalse() *BasicBlock {
	return term.targetFalse
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

// String returns the LLVM syntax representation of the terminator.
func (term *TermSwitch) String() string {
	buf := &bytes.Buffer{}
	x := term.X()
	fmt.Fprintf(buf, "switch %s %s, label %s [ ",
		x.Type(),
		x.Ident(),
		term.TargetDefault().Ident())
	for i, c := range term.Cases() {
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
func (term *TermSwitch) Parent() *BasicBlock {
	return term.parent
}

// SetParent sets the parent basic block of the terminator.
func (term *TermSwitch) SetParent(parent *BasicBlock) {
	term.parent = parent
}

// Successors returns the successor basic blocks of the terminator.
func (term *TermSwitch) Successors() []*BasicBlock {
	return term.successors
}

// X returns the control variable of the switch terminator.
func (term *TermSwitch) X() value.Value {
	return term.x
}

// TargetDefault returns the default target branch of the switch terminator.
func (term *TermSwitch) TargetDefault() *BasicBlock {
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

// TermUnreachable represents an unreachable terminator.
//
// References:
//    http://llvm.org/docs/LangRef.html#unreachable-instruction
type TermUnreachable struct {
	// Parent basic block.
	parent *BasicBlock
}

// NewUnreachable returns a new unreachable terminator.
func NewUnreachable() *TermUnreachable {
	return &TermUnreachable{}
}

// String returns the LLVM syntax representation of the terminator.
func (term *TermUnreachable) String() string {
	return "unreachable"
}

// Parent returns the parent basic block of the terminator.
func (term *TermUnreachable) Parent() *BasicBlock {
	return term.parent
}

// SetParent sets the parent basic block of the terminator.
func (term *TermUnreachable) SetParent(parent *BasicBlock) {
	term.parent = parent
}

// Successors returns the successor basic blocks of the terminator.
func (term *TermUnreachable) Successors() []*BasicBlock {
	// unreachable terminators have no successors.
	return nil
}
