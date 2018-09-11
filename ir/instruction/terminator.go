package instruction

// === [ Terminators ] =========================================================

// --- [ ret ] -----------------------------------------------------------------

// Ret is an LLVM IR ret terminator.
type Ret struct {
}

// --- [ br ] ------------------------------------------------------------------

// Br is an unconditional LLVM IR br terminator.
type Br struct {
}

// CondBr is a conditional LLVM IR br terminator.
type CondBr struct {
}

// --- [ switch ] --------------------------------------------------------------

// Switch is an LLVM IR switch terminator.
type Switch struct {
}

// --- [ indirectbr ] ----------------------------------------------------------

// IndirectBr is an LLVM IR indirectbr terminator.
type IndirectBr struct {
}

// --- [ invoke ] --------------------------------------------------------------

// Invoke is an LLVM IR invoke terminator.
type Invoke struct {
}

// --- [ resume ] --------------------------------------------------------------

// Resume is an LLVM IR resume terminator.
type Resume struct {
}

// --- [ catchswitch ] ---------------------------------------------------------

// CatchSwitch is an LLVM IR catchswitch terminator.
type CatchSwitch struct {
}

// --- [ catchret ] ------------------------------------------------------------

// CatchRet is an LLVM IR catchret terminator.
type CatchRet struct {
}

// --- [ cleanupret ] ----------------------------------------------------------

// CleanupRet is an LLVM IR cleanupret terminator.
type CleanupRet struct {
}

// --- [ unreachable ] ---------------------------------------------------------

// Unreachable is an LLVM IR unreachable terminator.
type Unreachable struct {
}
