package ir

import (
	"fmt"

	"github.com/llir/llvm/asm"
	"github.com/llir/llvm/types"
	"github.com/llir/llvm/values"
)

// A GlobalDecl represents a global variable definition or an external global
// variable declaration.
//
// Examples:
//    @x = global i32 42
//    @s = constant [13 x i8] c"hello world\0A\00"
//    @y = external global i32
//
// References:
//    http://www.llvm.org/docs/LangRef.html#globalvars
type GlobalDecl struct {
	// Variable name.
	name string
	// Variable type.
	typ types.Type
	// Initial value, or nil if defined externally.
	val values.Value
	// Specifies whether the global variable is immutable.
	immutable bool
}

// NewGlobalDecl returns a new global variable definition of the given name,
// type and initial value. The global variable is defined externally if val is
// nil. By default, global variables are mutable. Invoke SetImmutable to create
// read-only variables.
func NewGlobalDecl(name string, typ types.Type, val values.Value) (*GlobalDecl, error) {
	if val != nil && !typ.Equal(val.Type()) {
		return nil, fmt.Errorf("invalid global variable definition; type mismatch between variable declaration (%q) and initial value (%q)", typ, val.Type())
	}
	return &GlobalDecl{name: name, typ: typ, val: val}, nil
}

// Name returns the name of the value.
func (d *GlobalDecl) Name() string {
	return d.name
}

// Type returns the type of the value.
func (d *GlobalDecl) Type() types.Type {
	return d.typ
}

// Value returns the initial value of the variable definition, or nil if defined
// externally.
func (d *GlobalDecl) Value() values.Value {
	return d.val
}

// Immutable reports whether the global variable is immutable.
func (d *GlobalDecl) Immutable() bool {
	return d.immutable
}

// SetImmutable controls whether the global variable is immutable.
func (d *GlobalDecl) SetImmutable(immutable bool) {
	d.immutable = immutable
}

// String returns the string representation of the global variable declaration.
func (d *GlobalDecl) String() string {
	decl := "global"
	if d.Immutable() {
		decl = "constant"
	}
	if d.val != nil {
		// Global variable definition; e.g.
		//     @x = global i32 42
		//     @s = constant [13 x i8] c"hello world\0A\00"
		return fmt.Sprintf("%s = %s %s", asm.EncGlobal(d.Name()), decl, d.Value())
	}
	// External global variable declaration; e.g.
	//    @y = external global i32
	return fmt.Sprintf("%s = external %s %s", asm.EncGlobal(d.Name()), decl, d.Type())
}
