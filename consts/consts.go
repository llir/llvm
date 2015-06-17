// TODO: Figure out where metadata belongs. Is it a Constant? It feels like
// metadata belongs elsewhere, even if it is immutable. One reason being that
// derived constants cannot be created using metadata (verify!).

// Package consts implements values representing immutable LLVM IR constants.
package consts

import "github.com/llir/llvm/value"

// A Constant represents a value that is immutable at runtime, such as an
// integer or a floating point literal. Pointers known to be immutable at
// runtime are also constants (e.g. null). Functions and global variables are
// therefore constants as their addresses are immutable.
//
// Derived types (e.g. vectors, arrays and structures) containing only constant
// values are also constants. Lastly, constants may be used in constant
// expressions to create new constants.
//
// Constant is one of the following types:
//    *consts.Int
//    *consts.Float
//    *consts.Pointer
//    *consts.Vector
//    *consts.Array
//    *consts.Struct
//    consts.Expr
//
// References:
//    http://llvm.org/docs/LangRef.html#constants
type Constant interface {
	value.Value
	// isConst ensures that only constant values can be assigned to the Constant
	// interface.
	isConst()
}

// Make sure that each constant implements the Constant interface.
var (
	_ Constant = &Int{}
	_ Constant = &Float{}
	_ Constant = &Pointer{}
	_ Constant = &Vector{}
	_ Constant = &Array{}
	_ Constant = &Struct{}
	_ Constant = Expr(nil)
)
