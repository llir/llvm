// Package values provides a definition of LLVM IR values.
package value

import (
	"fmt"

	"github.com/llir/llvm/ir/types"
)

// TODO: Complete the list of value implementations.
//
//    *ir.BasicBlock

// A Value represents a computed value that may be used as an operand of other
// values. Some values can have a name and they belong to a function or a
// module.
//
// Value is one of the following types:
//
//    constant.Constant
//    *instruction.Local
type Value interface {
	fmt.Stringer
	// Type returns the type of the value.
	Type() types.Type
}

// TODO: Evaluate if NamedValue is needed.

// A NamedValue represents a named value; e.g. the global identifier "foo" of a
// global variable declaration, the label "bar" of a basic block, or the local
// identifier "baz" of a local variable definition, .
//
//    @foo = global i32 0
//
//    bar:
//       %baz = add i32 3, 5
//       store i32 42, i32* @foo
type NamedValue interface {
	Value
	// Name returns the name of the value.
	Name() string
}
