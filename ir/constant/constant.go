// Package constant implements values representing immutable LLVM IR constants.
package constant

import "github.com/llir/llvm/ir/value"

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
