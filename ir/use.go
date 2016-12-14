package ir

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

// valueTracker tracks the use of a value.
type valueTracker struct {
	// Original value.
	orig *value.Value
}

// trackValue tracks the use of the given value.
func trackValue(orig *value.Value) {
	use := &valueTracker{orig: orig}
	used, ok := (*orig).(value.Used)
	if !ok {
		panic(fmt.Sprintf("invalid used value type; expected value.Used, got %T", *orig))
	}
	used.AppendUse(use)
}

// Replace replaces the used value with the given value.
func (use *valueTracker) Replace(v value.Value) {
	*use.orig = v
}

// blockTracker tracks the use of a basic block.
type blockTracker struct {
	// Original basic block.
	orig **BasicBlock
}

func trackBlock(orig **BasicBlock) {
	use := &blockTracker{orig: orig}
	(*orig).AppendUse(use)
}

// Replace replaces the used value with the given value.
func (use *blockTracker) Replace(v value.Value) {
	block, ok := v.(*BasicBlock)
	if !ok {
		panic(fmt.Sprintf("invalid basic block type; expected *ir.BasicBlock, got %T", v))
	}
	*use.orig = block
}
