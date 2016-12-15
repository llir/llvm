package ir

import (
	"fmt"

	"github.com/llir/llvm/ir/constant"
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
	// User of the value.
	//
	// User may have one of the following underlying types.
	//
	//    ir.Instruction
	//    ir.Terminator
	user Instruction
}

// newUse returns a new use of a value based on the given value replacement
// function.
func newUse(replace func(v value.Value), user Instruction) value.Use {
	return &use{replace: replace, user: user}
}

// Replace replaces the used value with the given value.
func (use *use) Replace(v value.Value) {
	use.replace(v)
}

// User returns the user of the value.
//
// The returned user may have one of the following underlying types.
//
//    ir.Instruction
//    ir.Terminator
func (use *use) User() interface{} {
	return use.user
}

// valueTracker tracks the use of a value.
type valueTracker struct {
	// Original value.
	orig *value.Value
	// User of the value.
	//
	// User may have one of the following underlying types.
	//
	//    ir.Instruction
	//    ir.Terminator
	user Instruction
}

// trackValue tracks the use of the given value.
func trackValue(orig *value.Value, user Instruction) {
	use := &valueTracker{orig: orig, user: user}
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

// User returns the user of the value.
//
// The returned user may have one of the following underlying types.
//
//    ir.Instruction
//    ir.Terminator
func (use *valueTracker) User() interface{} {
	return use.user
}

// blockTracker tracks the use of a basic block.
type blockTracker struct {
	// Original basic block.
	orig **BasicBlock
	// User of the value.
	//
	// User may have one of the following underlying types.
	//
	//    ir.Instruction
	//    ir.Terminator
	user Instruction
}

// trackBlock tracks the use of the given basic block.
func trackBlock(orig **BasicBlock, user Instruction) {
	use := &blockTracker{orig: orig, user: user}
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

// User returns the user of the value.
//
// The returned user may have one of the following underlying types.
//
//    ir.Instruction
//    ir.Terminator
func (use *blockTracker) User() interface{} {
	return use.user
}

// namedTracker tracks the use of a named value.
type namedTracker struct {
	// Original named value.
	orig *value.Named
	// User of the value.
	//
	// User may have one of the following underlying types.
	//
	//    ir.Instruction
	//    ir.Terminator
	user Instruction
}

// trackNamed tracks the use of the given named value.
func trackNamed(orig *value.Named, user Instruction) {
	use := &namedTracker{orig: orig, user: user}
	used, ok := (*orig).(value.Used)
	if !ok {
		panic(fmt.Sprintf("invalid used value type; expected value.Used, got %T", *orig))
	}
	used.AppendUse(use)
}

// Replace replaces the used value with the given value.
func (use *namedTracker) Replace(v value.Value) {
	n, ok := v.(value.Named)
	if !ok {
		panic(fmt.Sprintf("invalid named value type; expected value.Named, got %T", v))
	}
	*use.orig = n
}

// User returns the user of the value.
//
// The returned user may have one of the following underlying types.
//
//    ir.Instruction
//    ir.Terminator
func (use *namedTracker) User() interface{} {
	return use.user
}

// constantTracker tracks the use of a value.
type constantTracker struct {
	// Original value.
	orig *constant.Constant
	// User of the value.
	//
	// User may have one of the following underlying types.
	//
	//    ir.Instruction
	//    ir.Terminator
	user Instruction
}

// trackConstant tracks the use of the given constant.
func trackConstant(orig *constant.Constant, user Instruction) {
	use := &constantTracker{orig: orig, user: user}
	used, ok := (*orig).(value.Used)
	if !ok {
		panic(fmt.Sprintf("invalid used value type; expected value.Used, got %T", *orig))
	}
	used.AppendUse(use)
}

// Replace replaces the used value with the given value.
func (use *constantTracker) Replace(v value.Value) {
	c, ok := v.(constant.Constant)
	if !ok {
		panic(fmt.Sprintf("invalid constant type; expected constant.Constant, got %T", v))
	}
	*use.orig = c
}

// User returns the user of the value.
//
// The returned user may have one of the following underlying types.
//
//    ir.Instruction
//    ir.Terminator
func (use *constantTracker) User() interface{} {
	return use.user
}

// intConstTracker tracks the use of a value.
type intConstTracker struct {
	// Original value.
	orig **constant.Int
	// User of the value.
	//
	// User may have one of the following underlying types.
	//
	//    ir.Instruction
	//    ir.Terminator
	user Instruction
}

// trackIntConst tracks the use of the given integer constant.
func trackIntConst(orig **constant.Int, user Instruction) {
	use := &intConstTracker{orig: orig, user: user}
	(*orig).AppendUse(use)
}

// Replace replaces the used value with the given value.
func (use *intConstTracker) Replace(v value.Value) {
	c, ok := v.(*constant.Int)
	if !ok {
		panic(fmt.Sprintf("invalid integer constant type; expected *constant.Int, got %T", v))
	}
	*use.orig = c
}

// User returns the user of the value.
//
// The returned user may have one of the following underlying types.
//
//    ir.Instruction
//    ir.Terminator
func (use *intConstTracker) User() interface{} {
	return use.user
}
