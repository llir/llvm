package ir

type Instruction interface {
	// isInstruction ensures that only instructions can be assigned to the
	// instruction.Instruction interface.
	isInstruction()
}

// Binary instructions.
func (*InstAdd) isInstruction()  {}
func (*InstFAdd) isInstruction() {}
func (*InstSub) isInstruction()  {}
func (*InstFSub) isInstruction() {}
func (*InstMul) isInstruction()  {}
func (*InstFMul) isInstruction() {}
func (*InstUDiv) isInstruction() {}
func (*InstSDiv) isInstruction() {}
func (*InstFDiv) isInstruction() {}
func (*InstURem) isInstruction() {}
func (*InstSRem) isInstruction() {}
func (*InstFRem) isInstruction() {}

// Bitwise instructions.
func (*InstShl) isInstruction()  {}
func (*InstLShr) isInstruction() {}
func (*InstAShr) isInstruction() {}
func (*InstAnd) isInstruction()  {}
func (*InstOr) isInstruction()   {}
func (*InstXor) isInstruction()  {}

// Vector instructions.
func (*InstExtractElement) isInstruction() {}
func (*InstInsertElement) isInstruction()  {}
func (*InstShuffleVector) isInstruction()  {}

// Aggregate instructions.
func (*InstExtractValue) isInstruction() {}
func (*InstInsertValue) isInstruction()  {}

// Memory instructions.
func (*InstAlloca) isInstruction()        {}
func (*InstLoad) isInstruction()          {}
func (*InstStore) isInstruction()         {}
func (*InstFence) isInstruction()         {}
func (*InstCmpXchg) isInstruction()       {}
func (*InstAtomicRMW) isInstruction()     {}
func (*InstGetElementPtr) isInstruction() {}

// Conversion instructions.
func (*InstTrunc) isInstruction()         {}
func (*InstZExt) isInstruction()          {}
func (*InstSExt) isInstruction()          {}
func (*InstFPTrunc) isInstruction()       {}
func (*InstFPExt) isInstruction()         {}
func (*InstFPToUI) isInstruction()        {}
func (*InstFPToSI) isInstruction()        {}
func (*InstUIToFP) isInstruction()        {}
func (*InstSIToFP) isInstruction()        {}
func (*InstPtrToInt) isInstruction()      {}
func (*InstIntToPtr) isInstruction()      {}
func (*InstBitCast) isInstruction()       {}
func (*InstAddrSpaceCast) isInstruction() {}

// Other instructions.
func (*InstICmp) isInstruction()       {}
func (*InstFCmp) isInstruction()       {}
func (*InstPhi) isInstruction()        {}
func (*InstSelect) isInstruction()     {}
func (*InstCall) isInstruction()       {}
func (*InstVAArg) isInstruction()      {}
func (*InstLandingPad) isInstruction() {}
func (*InstCatchPad) isInstruction()   {}
func (*InstCleanupPad) isInstruction() {}
