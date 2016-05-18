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

// Name returns the name of the defined local variable.
func (def *LocalVarDef) Name() string {
	return def.name
}

// Value returns the instruction defining the initial value of the local
// variable definition.
func (def *LocalVarDef) Value() ValueInst {
	return def.val
}

// // Type returns the type of the value.
func (def *LocalVarDef) Type() types.Type {
	return def.val.Type()
}

// String returns the string representation of the local variable definition.
func (def *LocalVarDef) String() string {
	return fmt.Sprintf("%s = %s", asm.EncLocal(def.Name()), def.Value())
}

// NewLocalVarDef returns a new local variable definition based on the given
// name and initial value instruction.
func NewLocalVarDef(name string, val ValueInst) *LocalVarDef {
	return &LocalVarDef{
		name: name,
		val:  val,
	}
}

// isInst ensures that only non-branching instructions can be assigned to the
// Instruction interface.
func (def *LocalVarDef) isInst() {}
