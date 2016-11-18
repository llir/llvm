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
	_ value.Value = &ir.Param{}
)

// Validates that the LLVMStringer interface is implemented by the relevant
// types.
var (
	_ LLVMStringer = &ir.Module{}
)
