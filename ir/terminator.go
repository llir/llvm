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
	// Succs returns the successor basic blocks of the terminator.
	Succs() []*BasicBlock
}

// --- [ ret ] -----------------------------------------------------------------

// TermRet represents a ret terminator.
//
// References:
//    http://llvm.org/docs/LangRef.html#ret-instruction
type TermRet struct {
	// Parent basic block.
	Parent *BasicBlock
	// Return value; or nil if "void" return.
	X value.Value
}

// NewRet returns a new ret terminator based on the given return value. A nil
// return value indicates a "void" return.
func NewRet(x value.Value) *TermRet {
	term := &TermRet{
		X: x,
	}
	if term.X != nil {
		trackValue(&term.X, term)
	}
	return term
}

// String returns the LLVM syntax representation of the terminator.
func (term *TermRet) String() string {
	if term.X != nil {
		return fmt.Sprintf("ret %s %s",
			term.X.Type(),
			term.X.Ident())
	}
	return "ret void"
}

// GetParent returns the parent basic block of the terminator.
func (term *TermRet) GetParent() *BasicBlock {
	return term.Parent
}

// SetParent sets the parent basic block of the terminator.
func (term *TermRet) SetParent(parent *BasicBlock) {
	term.Parent = parent
}

// Succs returns the successor basic blocks of the terminator.
func (term *TermRet) Succs() []*BasicBlock {
	// ret terminators have no successors.
	return nil
}

// --- [ br ] ------------------------------------------------------------------

// TermBr represents an unconditional br terminator.
//
// References:
//    http://llvm.org/docs/LangRef.html#br-instruction
type TermBr struct {
	// Parent basic block.
	Parent *BasicBlock
	// Target branch.
	Target *BasicBlock
	// Successors basic blocks.
	Successors []*BasicBlock
}

// NewBr returns a new unconditional br terminator based on the given target
// branch.
func NewBr(target *BasicBlock) *TermBr {
	successors := []*BasicBlock{target}
	term := &TermBr{
		Target:     target,
		Successors: successors,
	}
	trackBlock(&term.Target, term)
	return term
}

// String returns the LLVM syntax representation of the terminator.
func (term *TermBr) String() string {
	return fmt.Sprintf("br label %s", term.Target.Ident())
}

// GetParent returns the parent basic block of the terminator.
func (term *TermBr) GetParent() *BasicBlock {
	return term.Parent
}

// SetParent sets the parent basic block of the terminator.
func (term *TermBr) SetParent(parent *BasicBlock) {
	term.Parent = parent
}

// Succs returns the successor basic blocks of the terminator.
func (term *TermBr) Succs() []*BasicBlock {
	return term.Successors
}

// --- [ conditional br ] ------------------------------------------------------

// TermCondBr represents a conditional br terminator.
//
// References:
//    http://llvm.org/docs/LangRef.html#br-instruction
type TermCondBr struct {
	// Parent basic block.
	Parent *BasicBlock
	// Branching condition.
	Cond value.Value
	// Target branch when condition is true.
	TargetTrue *BasicBlock
	// Target branch when condition is false.
	TargetFalse *BasicBlock
	// Successors basic blocks.
	Successors []*BasicBlock
}

// NewCondBr returns a new conditional br terminator based on the given
// branching condition and conditional target branches.
func NewCondBr(cond value.Value, targetTrue, targetFalse *BasicBlock) *TermCondBr {
	successors := []*BasicBlock{targetTrue, targetFalse}
	term := &TermCondBr{
		Cond:        cond,
		TargetTrue:  targetTrue,
		TargetFalse: targetFalse,
		Successors:  successors,
	}
	trackValue(&term.Cond, term)
	trackBlock(&term.TargetTrue, term)
	trackBlock(&term.TargetFalse, term)
	return term
}

// String returns the LLVM syntax representation of the terminator.
func (term *TermCondBr) String() string {
	return fmt.Sprintf("br i1 %s, label %s, label %s",
		term.Cond.Ident(),
		term.TargetTrue.Ident(),
		term.TargetFalse.Ident())
}

// GetParent returns the parent basic block of the terminator.
func (term *TermCondBr) GetParent() *BasicBlock {
	return term.Parent
}

// SetParent sets the parent basic block of the terminator.
func (term *TermCondBr) SetParent(parent *BasicBlock) {
	term.Parent = parent
}

// Succs returns the successor basic blocks of the terminator.
func (term *TermCondBr) Succs() []*BasicBlock {
	return term.Successors
}

// --- [ switch ] --------------------------------------------------------------

// TermSwitch represents a switch terminator.
//
// References:
//    http://llvm.org/docs/LangRef.html#switch-instruction
type TermSwitch struct {
	// Parent basic block.
	Parent *BasicBlock
	// Control variable.
	X value.Value
	// Default target branch.
	TargetDefault *BasicBlock
	// Switch cases.
	Cases []*Case
	// Successors basic blocks.
	Successors []*BasicBlock
}

// TODO: Consider renaming x to control to avoid confusion between term.X() and
// case.X().

// NewSwitch returns a new switch terminator based on the given control
// variable, default target branch and switch cases.
func NewSwitch(x value.Value, targetDefault *BasicBlock, cases ...*Case) *TermSwitch {
	successors := []*BasicBlock{targetDefault}
	for _, c := range cases {
		successors = append(successors, c.Target)
	}
	term := &TermSwitch{
		X:             x,
		TargetDefault: targetDefault,
		Cases:         cases,
		Successors:    successors,
	}
	trackValue(&term.X, term)
	trackBlock(&term.TargetDefault, term)
	for _, c := range term.Cases {
		trackIntConst(&c.X, term)
		trackBlock(&c.Target, term)
	}
	return term
}

// String returns the LLVM syntax representation of the terminator.
func (term *TermSwitch) String() string {
	buf := &bytes.Buffer{}
	fmt.Fprintf(buf, "switch %s %s, label %s [\n",
		term.X.Type(),
		term.X.Ident(),
		term.TargetDefault.Ident())
	for _, c := range term.Cases {
		fmt.Fprintf(buf, "\t\t%s %s, label %s\n",
			c.X.Type(),
			c.X.Ident(),
			c.Target.Ident())
	}
	buf.WriteString("\t]")
	return buf.String()
}

// GetParent returns the parent basic block of the terminator.
func (term *TermSwitch) GetParent() *BasicBlock {
	return term.Parent
}

// SetParent sets the parent basic block of the terminator.
func (term *TermSwitch) SetParent(parent *BasicBlock) {
	term.Parent = parent
}

// Succs returns the successor basic blocks of the terminator.
func (term *TermSwitch) Succs() []*BasicBlock {
	return term.Successors
}

// Case represents a case of a switch terminator.
type Case struct {
	// Case comparand.
	X *constant.Int
	// Case target branch.
	Target *BasicBlock
}

// NewCase returns a new switch case based on the given case comparand and
// target branch.
func NewCase(x *constant.Int, target *BasicBlock) *Case {
	return &Case{X: x, Target: target}
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
	Parent *BasicBlock
}

// NewUnreachable returns a new unreachable terminator.
func NewUnreachable() *TermUnreachable {
	return &TermUnreachable{}
}

// String returns the LLVM syntax representation of the terminator.
func (term *TermUnreachable) String() string {
	return "unreachable"
}

// GetParent returns the parent basic block of the terminator.
func (term *TermUnreachable) GetParent() *BasicBlock {
	return term.Parent
}

// SetParent sets the parent basic block of the terminator.
func (term *TermUnreachable) SetParent(parent *BasicBlock) {
	term.Parent = parent
}

// Succs returns the successor basic blocks of the terminator.
func (term *TermUnreachable) Succs() []*BasicBlock {
	// unreachable terminators have no successors.
	return nil
}
