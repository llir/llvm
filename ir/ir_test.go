package ir_test

import (
	"fmt"

	"github.com/llir/llvm/ir"
	"github.com/llir/llvm/ir/value"
)

// Validates that the value.Value interface is implemented by the relevant
// types.
var (
	_ value.Value = &ir.BasicBlock{}
	_ value.Value = &ir.Function{}
	_ value.Value = &ir.Global{}
	// TODO: Check that all instructions producing results implement the
	// value.Value interface.
)

// Validates that the fmt.Stringer interface is implemented by the relevant
// types.
var (
	_ fmt.Stringer = &ir.BasicBlock{}
	_ fmt.Stringer = &ir.Function{}
	_ fmt.Stringer = &ir.Global{}
	_ fmt.Stringer = &ir.Module{}
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
	_ ir.Instruction = &ir.InstShl{}
	_ ir.Instruction = &ir.InstLShr{}
	_ ir.Instruction = &ir.InstAShr{}
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
