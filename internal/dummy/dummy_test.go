package dummy_test

import (
	"github.com/llir/llvm/internal/dummy"
	"github.com/llir/llvm/ir"
	"github.com/llir/llvm/ir/constant"
	"github.com/llir/llvm/ir/value"
)

// Valutate that the relevant types satisfy the constant.Constant interface.
var (
	_ constant.Constant = &dummy.Global{}
)

// Valutate that the relevant types satisfy the ir.Instruction interface.
var (
	// Other instructions
	_ ir.Instruction = &dummy.InstCall{}
)

// Valutate that the relevant types satisfy the value.Named interface.
var (
	_ value.Named = &dummy.Global{}
	_ value.Named = &dummy.Local{}
	// Other instructions
	_ value.Named = &dummy.InstCall{}
)
