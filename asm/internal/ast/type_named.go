package ast

// NamedType represents the type definition of a type alias or an identified
// struct type.
type NamedType struct {
	// Type name.
	Name string
	// Type definition.
	Def Type
}

// isType ensures that only types can be assigned to the ast.Type interface.
func (*NamedType) isType() {}
