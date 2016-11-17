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

// Validates that the LLVMStringer and value.Named interface are implemented by
// the relevant types.
var (
	_ LLVMStringer = &ir.Module{}
	_ value.Named  = &ir.BasicBlock{}
	_ value.Named  = &ir.Function{}
	_ value.Named  = &ir.Global{}
	_ value.Named  = &ir.Param{}
)
