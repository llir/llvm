package ir

// Assert that each instruction implements the instruction.Instruction
// interface.
var (
	// Binary instructions.
	_ Instruction = (*Add)(nil)
	_ Instruction = (*FAdd)(nil)
	_ Instruction = (*Sub)(nil)
	_ Instruction = (*FSub)(nil)
	_ Instruction = (*Mul)(nil)
	_ Instruction = (*FMul)(nil)
	_ Instruction = (*UDiv)(nil)
	_ Instruction = (*SDiv)(nil)
	_ Instruction = (*FDiv)(nil)
	_ Instruction = (*URem)(nil)
	_ Instruction = (*SRem)(nil)
	_ Instruction = (*FRem)(nil)
	// Bitwise instructions.
	_ Instruction = (*Shl)(nil)
	_ Instruction = (*LShr)(nil)
	_ Instruction = (*AShr)(nil)
	_ Instruction = (*And)(nil)
	_ Instruction = (*Or)(nil)
	_ Instruction = (*Xor)(nil)
	// Vector instructions.
	_ Instruction = (*ExtractElement)(nil)
	_ Instruction = (*InsertElement)(nil)
	_ Instruction = (*ShuffleVector)(nil)
	// Aggregate instructions.
	_ Instruction = (*ExtractValue)(nil)
	_ Instruction = (*InsertValue)(nil)
	// Memory instructions.
	_ Instruction = (*Alloca)(nil)
	_ Instruction = (*Load)(nil)
	_ Instruction = (*Store)(nil)
	_ Instruction = (*Fence)(nil)
	_ Instruction = (*CmpXchg)(nil)
	_ Instruction = (*AtomicRMW)(nil)
	_ Instruction = (*GetElementPtr)(nil)
	// Conversion instructions.
	_ Instruction = (*Trunc)(nil)
	_ Instruction = (*ZExt)(nil)
	_ Instruction = (*SExt)(nil)
	_ Instruction = (*FPTrunc)(nil)
	_ Instruction = (*FPExt)(nil)
	_ Instruction = (*FPToUI)(nil)
	_ Instruction = (*FPToSI)(nil)
	_ Instruction = (*UIToFP)(nil)
	_ Instruction = (*SIToFP)(nil)
	_ Instruction = (*PtrToInt)(nil)
	_ Instruction = (*IntToPtr)(nil)
	_ Instruction = (*BitCast)(nil)
	_ Instruction = (*AddrSpaceCast)(nil)
	// Other instructions.
	_ Instruction = (*ICmp)(nil)
	_ Instruction = (*FCmp)(nil)
	_ Instruction = (*Phi)(nil)
	_ Instruction = (*Select)(nil)
	_ Instruction = (*Call)(nil)
	_ Instruction = (*VAArg)(nil)
	_ Instruction = (*LandingPad)(nil)
	_ Instruction = (*CatchPad)(nil)
	_ Instruction = (*CleanupPad)(nil)
)

// Assert that each terminator implements the instruction.Terminator interface.
var (
	// Terminators.
	_ Terminator = (*Ret)(nil)
	_ Terminator = (*Br)(nil)
	_ Terminator = (*CondBr)(nil)
	_ Terminator = (*Switch)(nil)
	_ Terminator = (*IndirectBr)(nil)
	_ Terminator = (*Invoke)(nil)
	_ Terminator = (*Resume)(nil)
	_ Terminator = (*CatchSwitch)(nil)
	_ Terminator = (*CatchRet)(nil)
	_ Terminator = (*CleanupRet)(nil)
	_ Terminator = (*Unreachable)(nil)
)
