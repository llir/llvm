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
	name string
	// Type definition.
	def Type
}

// NewNamed returns a new type definition based on the given type name and
// underlying type definition.
//
// A nil underlying type definition may be used to specify an opaque struct
// type, the body of which may later be specified using the SetDef method.
func NewNamed(name string, def Type) *NamedType {
	return &NamedType{name: name, def: def}
}

// Equal reports whether t and u are of equal type.
func (t *NamedType) Equal(u Type) bool {
	// Identified struct types are uniqued by type names, not by structural
	// identity.
	if _, ok := t.def.(*StructType); ok {
		if u, ok := u.(*NamedType); ok {
			if _, ok := u.def.(*StructType); ok {
				// t and u are both identified struct types; check equality by type
				// name.
				return t.name == u.name
			}
			// t is an identified struct type and u is a type alias (i.e. non-
			// struct type); thus never equal.
			return false
		}
		// t is an identified struct and u is an unnamed type; thus never equal.
		return false
	}
	// t is a type alias; check using regular type equality.
	return t.def.Equal(u)
}

// String returns the LLVM syntax representation of the type.
func (t *NamedType) String() string {
	return enc.Local(t.name)
}

// Name returns the name of the named type.
func (t *NamedType) Name() string {
	return t.name
}

// SetName sets the name of the named type.
func (t *NamedType) SetName(name string) {
	t.name = name
}

// Def returns the definition of the named type and a boolean indicating if a
// type definition was present.
func (t *NamedType) Def() (Type, bool) {
	if t.def != nil {
		return t.def, true
	}
	return nil, false
}

// SetDef sets the definition of the named type.
func (t *NamedType) SetDef(def Type) {
	t.def = def
}
