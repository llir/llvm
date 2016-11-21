package ir

import (
	"fmt"

	"github.com/llir/llvm/ir/value"
)

// TermRet represents a ret terminator.
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
