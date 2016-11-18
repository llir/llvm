// Package constant implements values representing immutable LLVM IR constants.
package constant

import (
	"fmt"

	"github.com/llir/llvm/ir/types"
	"github.com/llir/llvm/ir/value"
)

// A Constant represents an LLVM IR constant; a value that is immutable at
// runtime, such as an integer or a floating point literal.
//
// Pointers known to be immutable at runtime are also constants (e.g. null).
// Functions and global variables are therefore considered constants as their
// addresses are immutable.
//
// Derived types (e.g. vectors, arrays and structures) containing only constant
// values are also constants.
//
// Lastly, constants may be used in constant expressions to create new
// constants.
//
// Constant may have one of the following underlying types.
//
//    TODO
type Constant interface {
	value.Value
}

// Int represents an integer constant.
type Int struct {
	// Constant value.
	x int64
	// Constant type.
	typ *types.IntType
}

// NewInt returns a new integer constant of the given value and type.
func NewInt(x int64, typ types.Type) *Int {
	if typ, ok := typ.(*types.IntType); ok {
		return &Int{x: x, typ: typ}
	}
	panic(fmt.Sprintf("invalid integer constant type; expected *types.IntType, got %T", typ))
}

// Type returns the type of the integer constant.
func (i *Int) Type() types.Type {
	return i.typ
}

// Ident returns the value of the integer constant.
func (i *Int) Ident() string {
	if i.typ.Bits() == 1 {
		switch i.x {
		case 0:
			return "false"
		case 1:
			return "true"
		default:
			panic(fmt.Sprintf("invalid integer constant value; expected 0 or 1, got %d", i.x))
		}
	}
	return fmt.Sprintf("%d", i.x)
}

// LLVMString returns the LLVM syntax representation of the integer constant.
func (i *Int) LLVMString() string {
	return fmt.Sprintf("%v %v", i.typ, i.Ident())
}
