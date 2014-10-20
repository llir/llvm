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
	// isInst ensures that only instructions can be assigned to the Instruction
	// interface.
	isInst()
}

// =============================================================================
// Terminator Instructions
//
//    ref: http://llvm.org/docs/LangRef.html#terminators
// =============================================================================

// InstReturn represents a return instruction [1] in one of the following forms:
//    ret <Type> <Val>
//    ret void
//
//    [1]: http://llvm.org/docs/LangRef.html#i-ret
type InstReturn struct {
	// Return type.
	Type Type
	// Return value; or nil in case of a void return.
	Val Value
}

// InstBranch represents an unconditional branch instruction in the form of [1]:
//    br label <Target>
//
//    [1]: http://llvm.org/docs/LangRef.html#i-br
type InstBranch struct {
	// Target branch.
	Target *BasicBlock
}

// InstCondBranch represents a conditional branch instruction in the form
// of [1]:
//    br i1 <Cond>, label <TargetTrue>, label <TargetFalse>
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

// isInst ensures that only instructions can be assigned to the Instruction
// interface.
func (InstBranch) isInst()     {}
func (InstCondBranch) isInst() {}
