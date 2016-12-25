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

// GetName returns the name of the value.
func (global *Global) GetName() string {
	return global.Name
}

// SetName sets the name of the value.
func (global *Global) SetName(name string) {
	global.Name = name
}

// isValue ensures that only values can be assigned to the ast.Value interface.
func (*Global) isValue() {}

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

// GetName returns the name of the value.
func (global *GlobalDummy) GetName() string {
	return global.Name
}

// SetName sets the name of the value.
func (global *GlobalDummy) SetName(name string) {
	global.Name = name
}

// isValue ensures that only values can be assigned to the ast.Value interface.
func (*GlobalDummy) isValue() {}

// isConstant ensures that only constants can be assigned to the ast.Constant
// interface.
func (*GlobalDummy) isConstant() {}
