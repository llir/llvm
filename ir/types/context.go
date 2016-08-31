package types

import (
	"fmt"

	"github.com/llir/llvm/ir/internal/enc"
)

// Context represents the type context of a LLVM IR module. It is used during
// the creation of named types (i.e. identified structures) and type aliases.
// The type aliases are translated to concrete types during parsing.
//
// References:
//    http://blog.llvm.org/2011/11/llvm-30-type-system-rewrite.html
type Context struct {
	// structs tracks the identified structures of the LLVM IR module; e.g.
	//    %x = type { %y, i16, i32 }
	structs map[string]*NamedStruct
	// alias tracks the type aliases of the LLVM IR module; e.g.
	//    %y = type i8
	alias map[string]Type
}

// NewContext returns a new type context for a LLVM IR module.
func NewContext() *Context {
	ctx := &Context{
		structs: make(map[string]*NamedStruct),
		alias:   make(map[string]Type),
	}
	return ctx
}

// Struct returns the identified structure of the given name. If no such
// structure exists, a new identified structure is created. To enable recursive
// and forward references, identified structures are initially created without
// bodies. It is the caller's responsibility to populate the body of the
// identified structure.
func (ctx *Context) Struct(name string) (*NamedStruct, error) {
	if old, ok := ctx.alias[name]; ok {
		return nil, fmt.Errorf("redefinition of type alias %q (%q) as identified structure", enc.Local(name), old)
	}
	t, ok := ctx.structs[name]
	if !ok {
		t = &NamedStruct{name: name}
		ctx.structs[name] = t
	}
	return t, nil
}

// Alias returns the underlying type of the provided type alias. The boolean
// return value reports whether such a type alias exists.
func (ctx *Context) Alias(name string) (Type, bool) {
	typ, ok := ctx.alias[name]
	return typ, ok
}

// SetAlias creates a type alias between the given name and type.
func (ctx *Context) SetAlias(name string, typ Type) error {
	if _, ok := ctx.structs[name]; ok {
		return fmt.Errorf("redefinition of identified structure %q as type alias (%q)", enc.Local(name), typ)
	}
	if old, ok := ctx.alias[name]; ok {
		return fmt.Errorf("redefinition of type alias %q; old mapping %q, new mapping %q", enc.Local(name), old, typ)
	}
	ctx.alias[name] = typ
	return nil
}

// Validate validates the type context by verifying that all identified
// structures have been assigned bodies.
func (ctx *Context) Validate() error {
	for _, t := range ctx.structs {
		if t.typ == nil {
			return fmt.Errorf("empty body of identified structure %q", enc.Local(t.Name()))
		}
	}
	return nil
}

// NamedStruct represents an identified structure.
type NamedStruct struct {
	// name specifies the name of the identified structure.
	name string
	// Structure body.
	typ *Struct
}

// Name returns the name of the identified structure.
func (t *NamedStruct) Name() string {
	// TODO: Handle unnamed identified structures (generate anonymous names).
	return t.name
}

// Struct returns the underlying structure type of the identified structure.
func (t *NamedStruct) Struct() *Struct {
	return t.typ
}

// SetStruct sets the underlying structure type definition of the identified
// structure.
func (t *NamedStruct) SetStruct(typ *Struct) error {
	if typ == nil {
		t.typ = typ
	} else if !t.typ.Equal(typ) {
		return fmt.Errorf("redefinition of identified structure %q; old definition %v, new definition %v", enc.Local(t.Name()), t.typ, typ)
	}
	return nil
}

// Equal returns true if the given types are equal, and false otherwise.
//
// Identified structure types are uniqued by type names, not by structural
// identity.
func (t *NamedStruct) Equal(u Type) bool {
	if u, ok := u.(*NamedStruct); ok {
		return t.Name() == u.Name()
	}
	return false
}

// String returns a string representation of the identified structure type.
//
// As identified structure types may include recursive references, they are
// always printed by their type names.
func (t *NamedStruct) String() string {
	// e.g. "%regset"
	return enc.Local(t.Name())
}
