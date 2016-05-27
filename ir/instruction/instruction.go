//go:generate go run gen.go

// Package instruction declares the instructions of LLVM IR.
package instruction

import "github.com/llir/llvm/ir/value"

// An Instruction performs a non-branching operation and belongs to one of the
// following groups:
//
//    * Binary Operations
//    * Bitwise Binary Operations
//    * Vector Operations
//    * Aggregate Operations
//    * Memory Access and Addressing Operations
//    * Conversion Operations
//    * Other Operations
type Instruction interface {
	value.Value
	// isInst ensures that only non-branching instructions can be assigned to the
	// Instruction interface.
	isInst()
}

// Make sure that each instruction implements the Instruction interface.
var (
	// Local variable declarations.
	_ Instruction = &LocalVarDef{}

	// Memory Access and Addressing Operations
	_ Instruction = &Store{}
	_ Instruction = &Fence{}
)
