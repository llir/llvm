package ast

// A Terminator represents an LLVM IR terminator.
//
// Terminator may have one of the following underlying types.
//
// Terminators
//
// http://llvm.org/docs/LangRef.html#terminator-instructions
//
//    *ast.TermRet
//    *ast.TermBr
//    *ast.TermCondBr
//    *ast.TermSwitch
//    *ast.TermUnreachable
type Terminator interface {
	// isTerm ensures that only terminators can be assigned to the ast.Terminator
	// interface.
	isTerm()
}

// --- [ ret ] -----------------------------------------------------------------

// TermRet represents a ret terminator.
//
// References:
//    http://llvm.org/docs/LangRef.html#ret-instruction
type TermRet struct {
	// Return value; or nil if "void" return.
	X Value
}

// --- [ br ] ------------------------------------------------------------------

// TermBr represents an unconditional br terminator.
//
// References:
//    http://llvm.org/docs/LangRef.html#br-instruction
type TermBr struct {
	// Target branch.
	Target *BasicBlock
}

// --- [ conditional br ] ------------------------------------------------------

// TermCondBr represents a conditional br terminator.
//
// References:
//    http://llvm.org/docs/LangRef.html#br-instruction
type TermCondBr struct {
	// Branching condition.
	Cond Value
	// Target branch when condition is true.
	TargetTrue *BasicBlock
	// Target branch when condition is false.
	TargetFalse *BasicBlock
}

// --- [ switch ] --------------------------------------------------------------

// TermSwitch represents a switch terminator.
//
// References:
//    http://llvm.org/docs/LangRef.html#switch-instruction
type TermSwitch struct {
	// Control variable.
	X Value
	// Default target branch.
	TargetDefault *BasicBlock
	// Switch cases.
	Cases []*Case
}

// Case represents a case of a switch terminator.
type Case struct {
	// Case comparand.
	X *IntConst
	// Case target branch.
	Target *BasicBlock
}

// --- [ indirectbr ] ----------------------------------------------------------

// --- [ invoke ] --------------------------------------------------------------

// --- [ resume ] --------------------------------------------------------------

// --- [ catchswitch ] ---------------------------------------------------------

// --- [ catchret ] ------------------------------------------------------------

// --- [ cleanupret ] ----------------------------------------------------------

// --- [ unreachable ] ---------------------------------------------------------

// TermUnreachable represents an unreachable terminator.
//
// References:
//    http://llvm.org/docs/LangRef.html#unreachable-instruction
type TermUnreachable struct {
}

// isTerm ensures that only terminators can be assigned to the ast.Terminator
// interface.
func (*TermRet) isTerm()         {}
func (*TermBr) isTerm()          {}
func (*TermCondBr) isTerm()      {}
func (*TermSwitch) isTerm()      {}
func (*TermUnreachable) isTerm() {}
