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
	// TODO: Verify that name is not a local ID. Unnamed local variable
	// definitions should be assigned a local ID implicitly by the internal
	// localID counter of the given function rather than explicitly assigned.
	return &LocalVarDef{name: name, val: val}, nil
}

// Name returns the name of the defined local variable.
func (def *LocalVarDef) Name() string {
	return def.name
}

// TODO: Add note to SetName not set local IDs explicitly, as these are assigned
// implicitly by the internal localID counter.

// SetName sets the name of the defined local variable.
func (def *LocalVarDef) SetName(name string) {
	def.name = name
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
	if types.IsVoid(def.Type()) || len(def.Name()) == 0 {
		return def.Value().String()
	}
	return fmt.Sprintf("%s = %s", asm.EncLocal(def.Name()), def.Value())
}

// isInst ensures that only non-branching instructions can be assigned to the
// Instruction interface.
func (def *LocalVarDef) isInst() {}
