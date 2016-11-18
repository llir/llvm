package instruction_test

import (
	"github.com/llir/llvm/ir/instruction"
)

// Validates that the instruction.Instruction interface is implemented by the
// relevant types.
var (
	_ instruction.Instruction = &instruction.Add{}
	_ instruction.Instruction = &instruction.Call{}
	_ instruction.Instruction = &instruction.Load{}
	_ instruction.Instruction = &instruction.Mul{}
)

// Validates that the instruction.Terminator interface is implemented by the
// relevant types.
var (
	_ instruction.Terminator = &instruction.Ret{}
)
