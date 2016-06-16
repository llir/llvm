package ir

import (
	"fmt"

	"github.com/llir/llvm/asm"
	"github.com/llir/llvm/ir/types"
	"github.com/llir/llvm/ir/value"
	"github.com/mewkiz/pkg/errutil"
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
	// Global variable type.
	typ *types.Pointer
	// Underlying type of global variable.
	underlying types.Type
	// Initial value, or nil if defined externally.
	val value.Value
	// Specifies whether the global variable is immutable.
	immutable bool
}

// NewGlobalDef returns a new global variable definition of the given name and
// initial value. The variable is read-only if immutable is true.
func NewGlobalDef(name string, val value.Value, immutable bool) (*GlobalDecl, error) {
	d, err := NewGlobalDecl(name, val.Type(), immutable)
	if err != nil {
		return nil, errutil.Err(err)
	}
	d.val = val
	return d, nil
}

// NewGlobalDecl returns a new external global variable declaration of the given
// name and type. The variable is read-only if immutable is true.
func NewGlobalDecl(name string, underlying types.Type, immutable bool) (*GlobalDecl, error) {
	typ, err := types.NewPointer(underlying)
	if err != nil {
		return nil, errutil.Err(err)
	}
	return &GlobalDecl{name: name, typ: typ, underlying: underlying, immutable: immutable}, nil
}

// Name returns the name of the value.
func (d *GlobalDecl) Name() string {
	return d.name
}

// Type returns the type of the value.
func (d *GlobalDecl) Type() types.Type {
	return d.typ
}

// Underlying returns the underlying type of the global variable.
func (d *GlobalDecl) Underlying() types.Type {
	return d.underlying
}

// Value returns the initial value of the variable definition, or nil if defined
// externally.
func (d *GlobalDecl) Value() value.Value {
	return d.val
}

// Immutable reports whether the global variable is immutable.
func (d *GlobalDecl) Immutable() bool {
	return d.immutable
}

// String returns the string representation of the global variable declaration.
func (d *GlobalDecl) String() string {
	immutableSpec := "global"
	if d.Immutable() {
		immutableSpec = "constant"
	}

	// External global variable declaration; e.g.
	//    @y = external global i32
	if d.val == nil {
		return fmt.Sprintf("%s = external %s %s", asm.EncGlobal(d.Name()), immutableSpec, d.Underlying())
	}

	// Global variable definition; e.g.
	//     @x = global i32 42
	//     @s = constant [13 x i8] c"hello world\0A\00"
	return fmt.Sprintf("%s = %s %s %s", asm.EncGlobal(d.Name()), immutableSpec, d.Underlying(), d.Value().ValueString())
}

// ValueString returns a string representation of the value.
func (d *GlobalDecl) ValueString() string {
	return asm.EncGlobal(d.Name())
}
