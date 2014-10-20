package ir

// An Instruction belongs to one of the following groups [1]:
//    * terminator instructions
//    * binary instructions
//    * bitwise binary instructions
//    * memory instructions
//    * other instructions
//
//    [1]: http://llvm.org/docs/LangRef.html#instruction-reference
type Instruction interface {
	isInst()
}

// =============================================================================
// Terminator Instructions
//
//    ref: http://llvm.org/docs/LangRef.html#terminators
// =============================================================================

// InstBranch represents an unconditional branch instruction [1].
//
//    [1]: http://llvm.org/docs/LangRef.html#i-br
type InstBranch struct {
	// Target branch.
	Target *BasicBlock
}

// InstCondBranch represents a conditional branch instruction [1].
//
//    [1]: http://llvm.org/docs/LangRef.html#i-br
type InstCondBranch struct {
	// Boolean branching condition.
	Cond Value
	// Target branch when the condition evaluates to true.
	TargetTrue *BasicBlock
	// Target branch when the condition evaluates to false.
	TargetFalse *BasicBlock
}

// =============================================================================
// Binary Operations
//
//    ref: http://llvm.org/docs/LangRef.html#binaryops
// =============================================================================

// TODO(u): Add binary operations.

// =============================================================================
// Bitwise Binary Operations
//
//    ref: http://llvm.org/docs/LangRef.html#bitwiseops
// =============================================================================

// TODO(u): Add bitwise binary operations.

// =============================================================================
// Memory Access and Addressing Operations
//
//    ref: http://llvm.org/docs/LangRef.html#memoryops
// =============================================================================

// TODO(u): Add memory access and addressing operations.

// =============================================================================
// Other Operations
//
//    ref: http://llvm.org/docs/LangRef.html#otherops
// =============================================================================

// TODO(u): Add other operations.

func (InstBranch) isInst()     {}
func (InstCondBranch) isInst() {}
