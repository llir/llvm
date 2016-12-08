package irx

import (
	"github.com/llir/llvm/ir"
	"github.com/llir/llvm/ir/value"
)

// Valutate that the relevant types satisfy the ir.Instruction interface.
var (
	// Other instructions
	_ ir.Instruction = &instPhiDummy{}
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
	// Other instructions
	_ value.Named = &instPhiDummy{}
	_ value.Named = &instCallDummy{}
)
