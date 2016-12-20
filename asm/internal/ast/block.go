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

// GetName returns the name of the value.
func (block *BasicBlock) GetName() string {
	return block.Name
}

// SetName sets the name of the value.
func (block *BasicBlock) SetName(name string) {
	block.Name = name
}

// isValue ensures that only values can be assigned to the ast.Value interface.
func (*BasicBlock) isValue() {}
