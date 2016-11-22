package ir_test

import (
	"github.com/llir/llvm/ir"
	"github.com/llir/llvm/ir/value"
)

// LLVMStringer is implemented by any value that has a LLVMString method, which
// defines the LLVM syntax of that value.
type LLVMStringer interface {
	// LLVMString returns the LLVM syntax representation of the terminator.
	LLVMString() string
}

// Validates that the value.Value interface is implemented by the relevant
// types.
var (
	_ value.Value = &ir.BasicBlock{}
	_ value.Value = &ir.Function{}
	_ value.Value = &ir.Global{}
	// TODO: Check that all instructions producing results implement the
	// value.Value interface.
)

// Validates that the LLVMStringer interface is implemented by the relevant
// types.
var (
	_ LLVMStringer = &ir.BasicBlock{}
	_ LLVMStringer = &ir.Function{}
	_ LLVMStringer = &ir.Global{}
	_ LLVMStringer = &ir.Module{}
)

// Validates that the ir.Instruction interface is implemented by the relevant
// types.
var (
	// Binary instructions
	_ ir.Instruction = &ir.InstAdd{}
	_ ir.Instruction = &ir.InstFAdd{}
	_ ir.Instruction = &ir.InstSub{}
	_ ir.Instruction = &ir.InstFSub{}
	_ ir.Instruction = &ir.InstMul{}
	_ ir.Instruction = &ir.InstFMul{}
	_ ir.Instruction = &ir.InstUDiv{}
	_ ir.Instruction = &ir.InstSDiv{}
	_ ir.Instruction = &ir.InstFDiv{}
	_ ir.Instruction = &ir.InstURem{}
	_ ir.Instruction = &ir.InstSRem{}
	_ ir.Instruction = &ir.InstFRem{}

	// Bitwise instructions
	_ ir.Instruction = &ir.InstShL{}
	_ ir.Instruction = &ir.InstLShR{}
	_ ir.Instruction = &ir.InstAShR{}
	_ ir.Instruction = &ir.InstAnd{}
	_ ir.Instruction = &ir.InstOr{}
	_ ir.Instruction = &ir.InstXor{}

	// Vector instructions

	// Aggregate instructions

	// Memory instructions
	_ ir.Instruction = &ir.InstLoad{}

	// Conversion instructions

	// Other instructions
	_ ir.Instruction = &ir.InstCall{}
)

// Validates that the ir.Terminator interface is implemented by the relevant
// types.
var (
	_ ir.Terminator = &ir.TermRet{}
)
