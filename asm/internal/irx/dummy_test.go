package irx

import (
	"github.com/llir/llvm/ir"
	"github.com/llir/llvm/ir/constant"
	"github.com/llir/llvm/ir/value"
)

// Valutate that the relevant types satisfy the constant.Constant interface.
var (
	_ constant.Constant = &globalDummy{}
)

// Valutate that the relevant types satisfy the ir.Instruction interface.
var (
	// Other instructions
	_ ir.Instruction = &instPhiDummy{}
	_ ir.Instruction = &instCallDummy{}
)

// Valutate that the relevant types satisfy the ir.Terminator interface.
var (
	// Terminators
	_ ir.Terminator = &termBrDummy{}
	_ ir.Terminator = &termCondBrDummy{}
	_ ir.Terminator = &termSwitchDummy{}
)

// Valutate that the relevant types satisfy the value.Named interface.
var (
	_ value.Named = &globalDummy{}
	_ value.Named = &localDummy{}
	// Other instructions
	_ value.Named = &instPhiDummy{}
	_ value.Named = &instCallDummy{}
)
