package instruction

import "github.com/llir/llvm/ir/value"

// TODO: Consider removing the instruction.Value and use value.Value throughout.
//
// Cons:
//    * Would make it valid to assign `%foo = 3`

// A ValueInst instruction is a non-terminating instruction which returns a
// value.
type ValueInst interface {
	value.Value
	// isValueInst ensures that only instructions which return values can be
	// assigned to the Value interface.
	isValueInst()
}

var (
	// Binary Operations
	_ ValueInst = &Add{}
	_ ValueInst = &FAdd{}
	_ ValueInst = &Sub{}
	_ ValueInst = &FSub{}
	_ ValueInst = &Mul{}
	_ ValueInst = &FMul{}
	_ ValueInst = &UDiv{}
	_ ValueInst = &SDiv{}
	_ ValueInst = &FDiv{}
	_ ValueInst = &URem{}
	_ ValueInst = &SRem{}
	_ ValueInst = &FRem{}

	// Bitwise Binary Operations
	_ ValueInst = &ShL{}
	_ ValueInst = &LShR{}
	_ ValueInst = &AShR{}
	_ ValueInst = &And{}
	_ ValueInst = &Or{}
	_ ValueInst = &Xor{}

	// Vector Operations
	_ ValueInst = &ExtractElement{}
	_ ValueInst = &InsertElement{}
	_ ValueInst = &ShuffleVector{}

	// Aggregate Operations
	_ ValueInst = &ExtractValue{}
	_ ValueInst = &InsertValue{}

	// Memory Access and Addressing Operations
	_ ValueInst = &Alloca{}
	_ ValueInst = &Load{}
	_ ValueInst = &CmpXchg{}
	_ ValueInst = &AtomicRMW{}
	_ ValueInst = &GetElementPtr{}

	// Conversion Operations
	_ ValueInst = &Trunc{}
	_ ValueInst = &ZExt{}
	_ ValueInst = &SExt{}
	_ ValueInst = &FPTrunc{}
	_ ValueInst = &FPExt{}
	_ ValueInst = &FPToUI{}
	_ ValueInst = &FPToSI{}
	_ ValueInst = &UIToFP{}
	_ ValueInst = &SIToFP{}
	_ ValueInst = &PtrToInt{}
	_ ValueInst = &IntToPtr{}
	_ ValueInst = &BitCast{}
	_ ValueInst = &AddrSpaceCast{}

	// Other Operations
	_ ValueInst = &ICmp{}
	_ ValueInst = &FCmp{}
	_ ValueInst = &PHI{}
	_ ValueInst = &Select{}
	_ ValueInst = &Call{}
	_ ValueInst = &VAArg{}
	_ ValueInst = &LandingPad{}
)
