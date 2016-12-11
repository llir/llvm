package ir

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

// use represents the use of a value; e.g. when used as an argument to an
// instruction.
type use struct {
	// replace replaces the used value with the given value.
	replace func(v value.Value)
}

// newUse returns a new use of a value based on the given value replacement
// function.
func newUse(replace func(v value.Value)) value.Use {
	return &use{replace: replace}
}

// Replace replaces the used value with the given value.
func (use *use) Replace(v value.Value) {
	use.replace(v)
}
