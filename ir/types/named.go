// === [ Named types ] =========================================================

package types

import "github.com/llir/llvm/internal/enc"

// NamedType represents the type definition of a type alias or an identified
// struct type.
//
// A type alias represents an alias between a type name and a non-struct type. A
// type alias and its aliased type are of equal type.
//
// Identified struct types are uniqued by type names, not by structural
// identity.
type NamedType struct {
	// Type name.
	Name string
	// Type definition.
	Def Type
}

// NewNamed returns a new type definition based on the given type name and
// underlying type definition.
//
// A nil underlying type definition may be used to specify an opaque struct
// type, the body of which may later be specified using the SetDef method.
func NewNamed(name string, def Type) *NamedType {
	return &NamedType{Name: name, Def: def}
}

// Equal reports whether t and u are of equal type.
func (t *NamedType) Equal(u Type) bool {
	// Identified struct types are uniqued by type names, not by structural
	// identity.
	if _, ok := t.Def.(*StructType); ok {
		if u, ok := u.(*NamedType); ok {
			if _, ok := u.Def.(*StructType); ok {
				// t and u are both identified struct types; check equality by type
				// name.
				return t.Name == u.Name
			}
			// t is an identified struct type and u is a type alias (i.e. non-
			// struct type); thus never equal.
			return false
		}
		// t is an identified struct and u is an unnamed type; thus never equal.
		return false
	}
	// t is a type alias; check using regular type equality.
	return t.Def.Equal(u)
}

// String returns the LLVM syntax representation of the type.
func (t *NamedType) String() string {
	return enc.Local(t.Name)
}
