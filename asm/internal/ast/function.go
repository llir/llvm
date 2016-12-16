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

// isValue ensures that only values can be assigned to the ast.Value interface.
func (*Function) isValue() {}

// isNamedValue ensures that only named values can be assigned to the
// ast.NamedValue interface.
func (*Function) isNamedValue() {}

// isConstant ensures that only constants can be assigned to the ast.Constant
// interface.
func (*Function) isConstant() {}
