package constant

import (
	"fmt"

	"github.com/llir/llvm/ir/value"
)

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

// constantTracker tracks the use of a value.
type constantTracker struct {
	// Original value.
	orig *Constant
	// User of the value.
	//
	// User may have one of the following underlying types.
	//
	//    constant.Constant
	//    *ir.Global
	//    ir.Instruction
	//    ir.Terminator
	user interface{}
}

// trackConstant tracks the use of the given constant.
func trackConstant(orig *Constant, user interface{}) {
	use := &constantTracker{orig: orig, user: user}
	used, ok := (*orig).(value.Used)
	if !ok {
		panic(fmt.Sprintf("invalid used value type; expected value.Used, got %T", *orig))
	}
	used.AppendUse(use)
}

// Replace replaces the used value with the given value.
func (use *constantTracker) Replace(v value.Value) {
	c, ok := v.(Constant)
	if !ok {
		panic(fmt.Sprintf("invalid constant type; expected constant.Constant, got %T", v))
	}
	*use.orig = c
}

// User returns the user of the value.
//
// The returned user may have one of the following underlying types.
//
//    constant.Constant
//    *ir.Global
//    ir.Instruction
//    ir.Terminator
func (use *constantTracker) User() interface{} {
	return use.user
}
