// Package constant implements values representing immutable LLVM IR constants.
package constant

import (
	"github.com/llir/llvm/ir/types"
	"github.com/llir/llvm/ir/value"
)

// === [ Constants ] ===========================================================

// Convenience constants.
var (
	// None token constant.
	None = &NoneToken{} // none
	// Boolean constants.
	True  = NewInt(types.I1, 1) // true
	False = NewInt(types.I1, 0) // false
)

// TODO: include metadata node in doc comment list of "Complex constants"?

// Constant is an LLVM IR constant; a value that is immutable at runtime, such
// as an integer or floating-point literal, or the address of a function or
// global variable.
//
// A Constant has one of the following underlying types.
//
// # Simple constants
//
// https://llvm.org/docs/LangRef.html#simple-constants
//
//   - [*constant.Int]
//   - [*constant.Float]
//   - [*constant.Null]
//   - [*constant.NoneToken]
//
// # Complex constants
//
// https://llvm.org/docs/LangRef.html#complex-constants
//
//   - [*constant.Struct]
//   - [*constant.Array]
//   - [*constant.CharArray]
//   - [*constant.Vector]
//   - [*constant.ZeroInitializer]
//
// # Global variable and function addresses
//
// https://llvm.org/docs/LangRef.html#global-variable-and-function-addresses
//
//   - [*ir.Global]
//   - [*ir.Func]
//   - [*ir.Alias]
//   - [*ir.IFunc]
//
// # Undefined values
//
// https://llvm.org/docs/LangRef.html#undefined-values
//
//   - [*constant.Undef]
//
// # Poison values
//
// https://llvm.org/docs/LangRef.html#poison-values
//
//   - [*constant.Poison]
//
// # Addresses of basic blocks
//
// https://llvm.org/docs/LangRef.html#addresses-of-basic-blocks
//
//   - [*constant.BlockAddress]
//
// # Constant expressions
//
// https://llvm.org/docs/LangRef.html#constant-expressions
//
//   - [constant.Expression]
//
// [*ir.Global]: https://pkg.go.dev/github.com/llir/llvm/ir#Global
// [*ir.Func]: https://pkg.go.dev/github.com/llir/llvm/ir#Func
// [*ir.Alias]: https://pkg.go.dev/github.com/llir/llvm/ir#Alias
// [*ir.IFunc]: https://pkg.go.dev/github.com/llir/llvm/ir#IFunc
type Constant interface {
	value.Value
	// IsConstant ensures that only constants can be assigned to the
	// constant.Constant interface.
	IsConstant()
}

// NOTE: explicit links to pkg.go.dev are given for ir.Foo identifiers, as those
// were not recognized as links otherwise. Importing a dummy version of package
// ir (e.g. `var _ = &ir.Global{}`) would work and let godoc find the correct
// package, except for the fact that such an import creates a cyclic import.
