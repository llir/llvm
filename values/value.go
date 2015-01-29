// Package values provides a definition of LLVM IR values.
package values

import "github.com/mewlang/llvm/types"

// TODO: Complete the list of value implementations.

// A Value represents a computed value that may be used as an operand of other
// values. Some values can have a name and they belong to a module.
//
// Value is one of the following types:
//
//    *ir.BasicBlock
//    ir.Instruction
//    ir.Terminator
type Value interface {
	// UseList returns a list of all values which uses the value.
	UseList() []Value
	// Type returns the type of the value.
	Type() types.Type
	// ReplaceAll replaces all uses of the value with new.
	ReplaceAll(new Value) error
}
