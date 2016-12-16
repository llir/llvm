package ast

// LocalDummy represents a dummy local identifier.
type LocalDummy struct {
	// Local name.
	Name string
	// Type associated with the localIdent.
	Type Type
}

// isValue ensures that only values can be assigned to the ast.Value interface.
func (*LocalDummy) isValue() {}

// isNamedValue ensures that only named values can be assigned to the
// ast.NamedValue interface.
func (*LocalDummy) isNamedValue() {}
