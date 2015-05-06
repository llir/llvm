package types

import (
	"fmt"

	"github.com/llir/llvm/asm"
)

// Context represents the type context of a LLVM IR module. It is used during
// the creation of named types. The only types that can be named are identified
// structures.
//
// References:
//    http://blog.llvm.org/2011/11/llvm-30-type-system-rewrite.html
type Context struct {
	// structs tracks the identified structures of the LLVM IR module.
	structs map[string]*NamedStruct
}

// NewContext returns a new type context for a LLVM IR module.
func NewContext() *Context {
	return &Context{structs: make(map[string]*NamedStruct)}
}

// Get returns the identified structure of the given name. To enable recursive
// and forward references, identified structures are initially created without
// bodies. It is the caller's responsibility to populate the body of the
// identified structure.
func (ctx *Context) Get(name string) *NamedStruct {
	t, ok := ctx.structs[name]
	if !ok {
		t = &NamedStruct{name: name}
		ctx.structs[name] = t
	}
	return t
}

// Validate validates the type context by verifying that all identified
// structures have been assigned bodies.
func (ctx *Context) Validate() error {
	for _, t := range ctx.structs {
		if t.Struct == nil {
			return fmt.Errorf("empty body of identified structure %q", asm.EncLocal(t.Name()))
		}
	}
	return nil
}

// NamedStruct represents an identified structure.
type NamedStruct struct {
	// name specifies the name of the identified structure.
	name string
	// Structure body.
	*Struct
}

// Name returns the name of the identified structure.
func (t *NamedStruct) Name() string {
	// TODO: Handle unnamed identified structures (generate anonymous names).
	return t.name
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
	return asm.EncLocal(t.Name())
}
