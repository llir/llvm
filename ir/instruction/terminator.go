package instruction

import "github.com/llir/llvm/ir/value"

// A Terminator represents an LLVM IR terminator.
type Terminator interface {
	// Block returns the parent basic block of the terminator.
	Block() value.Value
	// LLVMString returns the LLVM syntax representation of the terminator.
	LLVMString() string
}
