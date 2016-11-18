package ir

import (
	"fmt"

	"github.com/llir/llvm/ir/types"
	"github.com/llir/llvm/ir/value"
)

// TermRet represents a return terminator instruction.
type TermRet struct {
	// Parent basic block.
	parent *BasicBlock
	// Local variable name storing the result of the instruction.
	name string
	// Return value; or nil if "void" return.
	x value.Value
}

// TODO: Consider changing the signature of NewRet to
//
//    func NewRet(x ...value.Value) *TermRet
//
// thus making the value optional.
//
// If taking this route, evalute other aspects of the API which may use the same
// method for optional arguments (e.g. new global variable with or without
// initializer).

// NewRet returns a new ret instruction based on the given return value. A nil
// return value indicates a "void" return instruction.
func NewRet(x value.Value) *TermRet {
	return &TermRet{x: x}
}

// Type returns the type of the instruction.
func (i *TermRet) Type() types.Type {
	if i.x != nil {
		return i.x.Type()
	}
	return types.Void
}

// Ident returns the identifier associated with the instruction.
func (i *TermRet) Ident() string {
	// TODO: Encode name if containing special characters.
	return "%" + i.name
}

// LLVMString returns the LLVM syntax representation of the instruction.
func (i *TermRet) LLVMString() string {
	if i.x != nil {
		return fmt.Sprintf("ret %v %v", i.x.Type().LLVMString(), i.x.Ident())
	}
	return "ret void"
}

// Successors returns the successor basic blocks of the terminator.
func (i *TermRet) Successors() []*BasicBlock {
	// Return instructions have no successors.
	return nil
}

// Parent returns the parent basic block of the instruction.
func (i *TermRet) Parent() *BasicBlock {
	return i.parent
}

// SetParent sets the parent basic block of the instruction.
func (i *TermRet) SetParent(parent *BasicBlock) {
	i.parent = parent
}

// SetName sets the name of the local variable storing the result of the
// instruction.
func (i *TermRet) SetName(name string) {
	i.name = name
}
