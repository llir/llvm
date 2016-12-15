package constant

import "github.com/llir/llvm/ir/value"

// used represents a used value; e.g. a value used as an operand to an
// instruction.
type used struct {
	// Uses of the used value.
	uses []value.Use
}

// Uses returns the uses of the used value.
func (u *used) Uses() []value.Use {
	return u.uses
}

// AppendUse appends the given use to the used value.
func (u *used) AppendUse(use value.Use) {
	u.uses = append(u.uses, use)
}

// SetUses sets the uses of the used value.
func (u *used) SetUses(uses []value.Use) {
	u.uses = uses
}
