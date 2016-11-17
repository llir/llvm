// Package value provides a definition of LLVM IR values.
package value

import "github.com/llir/llvm/ir/types"

// A Value represents an LLVM IR value, which may be used as an operand of other
// values.
//
// Value may have one of the following underlying types.
//
//    constant.Constant
//    Named
type Value interface {
	// Type returns the type of the value.
	Type() types.Type
	// LLVMString returns the LLVM syntax representation of the value.
	LLVMString() string
}

// A Named value represents an LLVM IR value which may be referred to by an
// associated identifier.
//
// Named may have one of the following underlying types.
//
//    *ir.BasicBlock
//    *ir.Function
//    *ir.Global
//    *ir.Param
//    instruction.Instruction
type Named interface {
	Value
	// Ident returns the identifier associated with the value.
	Ident() string
}
