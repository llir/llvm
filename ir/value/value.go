// Package value provides a definition of LLVM IR values.
package value

import "github.com/llir/llvm/ir/types"

// A Value represents an LLVM IR value, which may be used as an operand of
// instructions and terminators.
//
// Value may have one of the following underlying types.
//
//    constant.Constant   (https://godoc.org/github.com/llir/llvm/ir/constant#Constant)
//    value.Named         (https://godoc.org/github.com/llir/llvm/ir/value#Named)
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
//    *ir.BasicBlock   (https://godoc.org/github.com/llir/llvm/ir#BasicBlock)
//    *ir.Function     (https://godoc.org/github.com/llir/llvm/ir#Function)
//    *ir.Global       (https://godoc.org/github.com/llir/llvm/ir#Global)
//    *types.Param     (https://godoc.org/github.com/llir/llvm/ir/types#Param)
//    ir.Instruction   (https://godoc.org/github.com/llir/llvm/ir#Instruction)
type Named interface {
	Value
	// Name returns the name of the value.
	Name() string
	// SetName sets the name of the value.
	SetName(name string)
}
