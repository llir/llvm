package instruction

import (
	"fmt"

	"github.com/llir/llvm/asm"
	"github.com/llir/llvm/ir/types"
)

// A LocalVarDef represents a local variable definition.
//
// Examples:
//    %foo = add i32 13, 29
type LocalVarDef struct {
	// Name of the local variable.
	name string
	// Initial value.
	val ValueInst
}

// NewLocalVarDef returns a new local variable definition based on the given
// name and initial value instruction.
func NewLocalVarDef(name string, val ValueInst) (*LocalVarDef, error) {
	return &LocalVarDef{name: name, val: val}, nil
}

// Name returns the name of the defined local variable.
func (def *LocalVarDef) Name() string {
	return def.name
}

// Value returns the instruction defining the initial value of the local
// variable definition.
func (def *LocalVarDef) Value() ValueInst {
	return def.val
}

// Type returns the type of the value.
func (def *LocalVarDef) Type() types.Type {
	return def.val.Type()
}

// TODO: Add support for printing unnamed local identifiers; e.g. %3.

// String returns the string representation of the local variable definition.
func (def *LocalVarDef) String() string {
	if types.IsVoid(def.Type()) {
		return def.Value().String()
	}
	return fmt.Sprintf("%s = %s", asm.EncLocal(def.Name()), def.Value())
}

// isInst ensures that only non-branching instructions can be assigned to the
// Instruction interface.
func (def *LocalVarDef) isInst() {}

// TODO: Consider implementing Local using a pointer to the corresponding
// LocalVarDecl, and defining Local.String in terms of LocalVarDef.Name(), thus
// keeping the names in sync.

// TODO: Consider moving Local to the ir package, as it may be used for basic
// blocks as well.

// A Local represents a local variable.
type Local struct {
	// Name of the local variable.
	name string
	// Variable type.
	typ types.Type
}

// NewLocal returns a new local variable based on the given name and type.
func NewLocal(name string, typ types.Type) (*Local, error) {
	return &Local{name: name, typ: typ}, nil
}

// Type returns the type of the value.
func (l *Local) Type() types.Type {
	return l.typ
}

// String returns the string representation of the local variable.
func (l *Local) String() string {
	return asm.EncLocal(l.name)
}
