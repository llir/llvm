package instruction_test

import (
	"github.com/llir/llvm/instruction"
)

// Validates that the instruction.Instruction interface is implemented by the
// relevant types.
var (
	_ instruction.Instruction = &instruction.Add{}
)

// Validates that the instruction.Terminator interface is implemented by the
// relevant types.
var (
	_ instruction.Terminator = &instruction.Ret{}
)
