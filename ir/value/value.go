// Package value provides a definition of LLVM IR values.
package value

import "github.com/llir/llvm/ir/types"

// A Value represents an LLVM IR value, which may be used as an operand of
// instructions and terminators.
//
// Value may have one of the following underlying types.
//
//    constant.Constant
//    value.Named
type Value interface {
	// Type returns the type of the value.
	Type() types.Type
	// Ident returns the identifier associated with the value.
	Ident() string
}

// Named represents a named LLVM IR value.
//
// Named may have one of the following underlying types.
//
//    *ir.BasicBlock
//    *ir.Function
//    *ir.Global
//    *types.Param
//    ir.Instruction
type Named interface {
	Value
	// Name returns the name of the value.
	Name() string
	// SetName sets the name of the value.
	SetName(name string)
}
