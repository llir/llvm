// Package value provides a definition of LLVM IR values.
package value

import "github.com/llir/llvm/ir/types"

// A Value represents an LLVM IR value, which may be used as an operand of other
// values.
//
// Value may have one of the following underlying types.
//
//    *ir.BasicBlock
//    *ir.Function
//    *ir.Global
//    *types.Param
//    constant.Constant
//    ir.Instruction
type Value interface {
	// Type returns the type of the value.
	Type() types.Type
	// Ident returns the identifier associated with the value.
	Ident() string
}
