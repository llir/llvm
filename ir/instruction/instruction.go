// Package instruction declares the instructions of LLVM IR.
package instruction

import "github.com/llir/llvm/ir/value"

// An Instruction performs a non-branching operation and belongs to one of the
// following groups:
//
//    * Binary Operations
//    * Bitwise Binary Operations
//    * Vector Operations
//    * Aggregate Operations
//    * Memory Access and Addressing Operations
//    * Conversion Operations
//    * Other Operations
type Instruction interface {
	value.Value
	// isInst ensures that only non-branching instructions can be assigned to the
	// Instruction interface.
	isInst()
}

// Make sure that each instruction implements the Instruction interface.
var (
	// Binary Operations
	_ Instruction = &Add{}
	_ Instruction = &FAdd{}
	_ Instruction = &Sub{}
	_ Instruction = &FSub{}
	_ Instruction = &Mul{}
	_ Instruction = &FMul{}
	_ Instruction = &UDiv{}
	_ Instruction = &SDiv{}
	_ Instruction = &FDiv{}
	_ Instruction = &URem{}
	_ Instruction = &SRem{}
	_ Instruction = &FRem{}
	// Bitwise Binary Operations
	_ Instruction = &Shl{}
	_ Instruction = &LShr{}
	_ Instruction = &AShr{}
	_ Instruction = &And{}
	_ Instruction = &Or{}
	_ Instruction = &Xor{}
	// Vector Operations
	_ Instruction = &ExtractElement{}
	_ Instruction = &InsertElement{}
	_ Instruction = &ShuffleVector{}
	// Aggregate Operations
	_ Instruction = &ExtractValue{}
	_ Instruction = &InsertValue{}
	// Memory Access and Addressing Operations
	_ Instruction = &Alloca{}
	_ Instruction = &Load{}
	_ Instruction = &Store{}
	_ Instruction = &Fence{}
	_ Instruction = &CmpXchg{}
	_ Instruction = &AtomicRMW{}
	_ Instruction = &GetElementPtr{}
	// Conversion Operations
	_ Instruction = &Trunc{}
	_ Instruction = &ZExt{}
	_ Instruction = &SExt{}
	_ Instruction = &FPTrunc{}
	_ Instruction = &FPExt{}
	_ Instruction = &FPToUI{}
	_ Instruction = &FPToSI{}
	_ Instruction = &UIToFP{}
	_ Instruction = &SIToFP{}
	_ Instruction = &PtrToInt{}
	_ Instruction = &IntToPtr{}
	_ Instruction = &BitCast{}
	_ Instruction = &AddrSpaceCast{}
	// Other Operations
	_ Instruction = &ICmp{}
	_ Instruction = &FCmp{}
	_ Instruction = &PHI{}
	_ Instruction = &Select{}
	_ Instruction = &Call{}
	_ Instruction = &VAArg{}
	_ Instruction = &LandingPad{}
)
