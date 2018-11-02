package ast

// LocalDummy represents a dummy local identifier.
type LocalDummy struct {
	// Local name.
	Name string
	// Type associated with the localIdent.
	Type Type
}

// GetName returns the name of the value.
func (local *LocalDummy) GetName() string {
	return local.Name
}

// SetName sets the name of the value.
func (local *LocalDummy) SetName(name string) {
	local.Name = name
}

// isValue ensures that only values can be assigned to the ast.Value interface.
func (*LocalDummy) isValue() {}
