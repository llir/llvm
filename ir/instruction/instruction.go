// Package instruction declares the instructions of LLVM IR.
package instruction

import "github.com/llir/llvm/ir/value"

// An Instruction represents a non-branching LLVM IR instruction.
//
// Instruction may have one of the following underlying types.
//
//    TODO
type Instruction interface {
	value.Value
	// Block returns the parent basic block of the instruction.
	Block() value.Value
}
