package ast

// A Global represents an LLVM IR global variable definition or external global
// variable declaration.
type Global struct {
	// Global variable name.
	Name string
	// Content type.
	Content Type
	// Initial value; or nil if defined externally.
	Init Constant
	// Immutability of the global variable.
	Immutable bool
}

// isValue ensures that only values can be assigned to the ast.Value interface.
func (*Global) isValue() {}

// isNamedValue ensures that only named values can be assigned to the
// ast.NamedValue interface.
func (*Global) isNamedValue() {}

// isConstant ensures that only constants can be assigned to the ast.Constant
// interface.
func (*Global) isConstant() {}

// GlobalDummy represents a dummy global identifier.
type GlobalDummy struct {
	// Global name.
	Name string
	// Type associated with the global.
	Type Type
}

// isValue ensures that only values can be assigned to the ast.Value interface.
func (*GlobalDummy) isValue() {}

// isNamedValue ensures that only named values can be assigned to the
// ast.NamedValue interface.
func (*GlobalDummy) isNamedValue() {}

// isConstant ensures that only constants can be assigned to the ast.Constant
// interface.
func (*GlobalDummy) isConstant() {}
