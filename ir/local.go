package ir

import (
	"github.com/llir/llvm/asm"
	"github.com/llir/llvm/ir/types"
)

// A Local represents the local identifier of a basic block or a local variable
// definition.
type Local struct {
	// Name of the local identifier.
	name string
	// Value type.
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
