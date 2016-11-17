// Package instruction declares the instructions of LLVM IR.
package instruction

import "github.com/llir/llvm/ir/value"

// An Instruction represents a non-branching LLVM IR instruction.
type Instruction interface {
	value.Named
	// Block returns the parent basic block of the instruction.
	Block() value.Value
}
