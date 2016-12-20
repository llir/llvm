package ast

// A Function represents an LLVM IR function definition or external function
// declaration. The body of a function definition consists of a set of basic
// blocks, interconnected by control flow instructions.
type Function struct {
	// Function name.
	Name string
	// Function signature.
	Sig *FuncType
	// Basic blocks of the function; or nil if defined externally.
	Blocks []*BasicBlock
}

// GetName returns the name of the value.
func (f *Function) GetName() string {
	return f.Name
}

// SetName sets the name of the value.
func (f *Function) SetName(name string) {
	f.Name = name
}

// isValue ensures that only values can be assigned to the ast.Value interface.
func (*Function) isValue() {}

// isConstant ensures that only constants can be assigned to the ast.Constant
// interface.
func (*Function) isConstant() {}
