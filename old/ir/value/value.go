// Package value provides a definition of LLVM IR values.
package value

import (
	"fmt"

	"github.com/llir/llvm/ir/types"
)

// A Value represents a computed value that may be used as an operand of other
// values.
//
// Value may have one of the following underlying types.
//
//    *ir.Function
//    *ir.BasicBlock
//    *ir.GlobalDecl
//    *instruction.LocalVarDef
//    constant.Constant
//    *types.Param
type Value interface {
	fmt.Stringer
	// ValueString returns a string representation of the value.
	ValueString() string
	// Type returns the type of the value.
	Type() types.Type
}

// A NamedValue represents a named value; e.g. the global identifier "foo" of a
// global variable declaration, the label "bar" of a basic block, or the local
// identifier "baz" of a local variable definition.
//
//    @foo = global i32 0
//
//    bar:
//       %baz = add i32 3, 5
//       store i32 42, i32* @foo
//
// NamedValue may have one of the following underlying types.
//
//    *ir.Function
//    *ir.BasicBlock
//    *instruction.LocalVarDef
type NamedValue interface {
	Value
	// Name returns the name of the value.
	Name() string
}
