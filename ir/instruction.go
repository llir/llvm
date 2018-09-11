package ir

type Instruction interface {
	// isInstruction ensures that only instructions can be assigned to the
	// instruction.Instruction interface.
	isInstruction()
}

// Binary instructions.
func (*Add) isInstruction()  {}
func (*FAdd) isInstruction() {}
func (*Sub) isInstruction()  {}
func (*FSub) isInstruction() {}
func (*Mul) isInstruction()  {}
func (*FMul) isInstruction() {}
func (*UDiv) isInstruction() {}
func (*SDiv) isInstruction() {}
func (*FDiv) isInstruction() {}
func (*URem) isInstruction() {}
func (*SRem) isInstruction() {}
func (*FRem) isInstruction() {}

// Bitwise instructions.
func (*Shl) isInstruction()  {}
func (*LShr) isInstruction() {}
func (*AShr) isInstruction() {}
func (*And) isInstruction()  {}
func (*Or) isInstruction()   {}
func (*Xor) isInstruction()  {}

// Vector instructions.
func (*ExtractElement) isInstruction() {}
func (*InsertElement) isInstruction()  {}
func (*ShuffleVector) isInstruction()  {}

// Aggregate instructions.
func (*ExtractValue) isInstruction() {}
func (*InsertValue) isInstruction()  {}

// Memory instructions.
func (*Alloca) isInstruction()        {}
func (*Load) isInstruction()          {}
func (*Store) isInstruction()         {}
func (*Fence) isInstruction()         {}
func (*CmpXchg) isInstruction()       {}
func (*AtomicRMW) isInstruction()     {}
func (*GetElementPtr) isInstruction() {}

// Conversion instructions.
func (*Trunc) isInstruction()         {}
func (*ZExt) isInstruction()          {}
func (*SExt) isInstruction()          {}
func (*FPTrunc) isInstruction()       {}
func (*FPExt) isInstruction()         {}
func (*FPToUI) isInstruction()        {}
func (*FPToSI) isInstruction()        {}
func (*UIToFP) isInstruction()        {}
func (*SIToFP) isInstruction()        {}
func (*PtrToInt) isInstruction()      {}
func (*IntToPtr) isInstruction()      {}
func (*BitCast) isInstruction()       {}
func (*AddrSpaceCast) isInstruction() {}

// Other instructions.
func (*ICmp) isInstruction()       {}
func (*FCmp) isInstruction()       {}
func (*Phi) isInstruction()        {}
func (*Select) isInstruction()     {}
func (*Call) isInstruction()       {}
func (*VAArg) isInstruction()      {}
func (*LandingPad) isInstruction() {}
func (*CatchPad) isInstruction()   {}
func (*CleanupPad) isInstruction() {}
