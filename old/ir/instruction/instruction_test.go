package instruction_test

import (
	"github.com/llir/llvm/ir/instruction"
	"github.com/llir/llvm/ir/value"
)

// Ensure that each value implements the Value interface.
var (
	_ value.Value = &instruction.LocalVarDef{}
)

// Ensure that each instruction implements the Instruction interface.
var (
	// Local variable declarations.
	_ instruction.Instruction = &instruction.LocalVarDef{}

	// Memory Access and Addressing Operations
	_ instruction.Instruction = &instruction.Store{}
	_ instruction.Instruction = &instruction.Fence{}
)

// Ensure that each terminator instruction implements the Terminator interface.
var (
	_ instruction.Terminator = &instruction.Ret{}
	_ instruction.Terminator = &instruction.Jmp{}
	_ instruction.Terminator = &instruction.Br{}
	_ instruction.Terminator = &instruction.Switch{}
	_ instruction.Terminator = &instruction.IndirectBr{}
	_ instruction.Terminator = &instruction.Invoke{}
	_ instruction.Terminator = &instruction.Resume{}
	_ instruction.Terminator = &instruction.Unreachable{}
)

// Ensure that each value instruction implements the ValueInst interface.
var (
	// Binary Operations
	_ instruction.ValueInst = &instruction.Add{}
	_ instruction.ValueInst = &instruction.FAdd{}
	_ instruction.ValueInst = &instruction.Sub{}
	_ instruction.ValueInst = &instruction.FSub{}
	_ instruction.ValueInst = &instruction.Mul{}
	_ instruction.ValueInst = &instruction.FMul{}
	_ instruction.ValueInst = &instruction.UDiv{}
	_ instruction.ValueInst = &instruction.SDiv{}
	_ instruction.ValueInst = &instruction.FDiv{}
	_ instruction.ValueInst = &instruction.URem{}
	_ instruction.ValueInst = &instruction.SRem{}
	_ instruction.ValueInst = &instruction.FRem{}

	// Bitwise Binary Operations
	_ instruction.ValueInst = &instruction.ShL{}
	_ instruction.ValueInst = &instruction.LShR{}
	_ instruction.ValueInst = &instruction.AShR{}
	_ instruction.ValueInst = &instruction.And{}
	_ instruction.ValueInst = &instruction.Or{}
	_ instruction.ValueInst = &instruction.Xor{}

	// Vector Operations
	_ instruction.ValueInst = &instruction.ExtractElement{}
	_ instruction.ValueInst = &instruction.InsertElement{}
	_ instruction.ValueInst = &instruction.ShuffleVector{}

	// Aggregate Operations
	_ instruction.ValueInst = &instruction.ExtractValue{}
	_ instruction.ValueInst = &instruction.InsertValue{}

	// Memory Access and Addressing Operations
	_ instruction.ValueInst = &instruction.Alloca{}
	_ instruction.ValueInst = &instruction.Load{}
	_ instruction.ValueInst = &instruction.CmpXchg{}
	_ instruction.ValueInst = &instruction.AtomicRMW{}
	_ instruction.ValueInst = &instruction.GetElementPtr{}

	// Conversion Operations
	_ instruction.ValueInst = &instruction.Trunc{}
	_ instruction.ValueInst = &instruction.ZExt{}
	_ instruction.ValueInst = &instruction.SExt{}
	_ instruction.ValueInst = &instruction.FPTrunc{}
	_ instruction.ValueInst = &instruction.FPExt{}
	_ instruction.ValueInst = &instruction.FPToUI{}
	_ instruction.ValueInst = &instruction.FPToSI{}
	_ instruction.ValueInst = &instruction.UIToFP{}
	_ instruction.ValueInst = &instruction.SIToFP{}
	_ instruction.ValueInst = &instruction.PtrToInt{}
	_ instruction.ValueInst = &instruction.IntToPtr{}
	_ instruction.ValueInst = &instruction.BitCast{}
	_ instruction.ValueInst = &instruction.AddrSpaceCast{}

	// Other Operations
	_ instruction.ValueInst = &instruction.ICmp{}
	_ instruction.ValueInst = &instruction.FCmp{}
	_ instruction.ValueInst = &instruction.PHI{}
	_ instruction.ValueInst = &instruction.Select{}
	_ instruction.ValueInst = &instruction.Call{}
	_ instruction.ValueInst = &instruction.VAArg{}
	_ instruction.ValueInst = &instruction.LandingPad{}
)
