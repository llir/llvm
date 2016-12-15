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
//    *ir.Global       (https://godoc.org/github.com/llir/llvm/ir#Global)
//    *ir.Function     (https://godoc.org/github.com/llir/llvm/ir#Function)
//    *ir.Param        (https://godoc.org/github.com/llir/llvm/ir#Param)
//    *ir.BasicBlock   (https://godoc.org/github.com/llir/llvm/ir#BasicBlock)
//    ir.Instruction   (https://godoc.org/github.com/llir/llvm/ir#Instruction)
type Named interface {
	Value
	// Name returns the name of the value.
	Name() string
	// SetName sets the name of the value.
	SetName(name string)
}

// Used represents a used value; e.g. a value used as an operand to an
// instruction.
type Used interface {
	Value
	// Uses returns the uses of the used value.
	Uses() []Use
	// AppendUse appends the given use to the used value.
	AppendUse(use Use)
	// SetUses sets the uses of the used value.
	SetUses(uses []Use)
}

// Use represents the use of a value; e.g. when used as an operand to an
// instruction.
type Use interface {
	// Replace replaces the used value with the given value.
	Replace(v Value)
	// User returns the user of the value.
	//
	// The returned user may have one of the following underlying types.
	//
	//    constant.Constant
	//    *ir.Global
	//    ir.Instruction
	//    ir.Terminator
	User() interface{}
}
