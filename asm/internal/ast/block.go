package ast

// A BasicBlock represents an LLVM IR basic block, which consists of a sequence
// of non-branching instructions, terminated by a control flow instruction (e.g.
// br or ret).
type BasicBlock struct {
	// Label name of the basic block; or empty if unnamed basic block.
	Name string
	// Non-branching instructions of the basic block.
	Insts []Instruction
	// Terminator of the basic block.
	Term Terminator
}

// isValue ensures that only values can be assigned to the ast.Value interface.
func (*BasicBlock) isValue() {}

// isNamedValue ensures that only named values can be assigned to the
// ast.NamedValue interface.
func (*BasicBlock) isNamedValue() {}
