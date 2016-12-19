//+build ignore

// TODO: Print value in generic error string.

package irx

import (
	"github.com/llir/llvm/ir"
	"github.com/llir/llvm/ir/value"
	"github.com/pkg/errors"
)

// === [ Modules ] =============================================================

// === [ Type definitions ] ====================================================

// === [ Global variables ] ====================================================

// === [ Functions ] ===========================================================

// === [ Identifiers ] =========================================================

// === [ Types ] ===============================================================

// === [ Values ] ==============================================================

// === [ Constants ] ===========================================================

// === [ Basic blocks ] ========================================================

// === [ Instructions ] ========================================================

// NewNamedInstruction returns a named instruction based on the given local
// variable name and instruction.
func NewNamedInstruction(name, inst interface{}) (ir.Instruction, error) {
	// namedInstruction represents a namedInstruction instruction.
	type namedInstruction interface {
		ir.Instruction
		value.Named
	}
	n, ok := name.(*LocalIdent)
	if !ok {
		return nil, errors.Errorf("invalid local variable name type; expected *irx.LocalIdent, got %T", name)
	}
	i, ok := inst.(namedInstruction)
	if !ok {
		return nil, errors.Errorf("invalid instruction type; expected namedInstruction, got %T", inst)
	}
	i.SetName(n.name)
	return i, nil
}

// === [ Terminators ] =========================================================
