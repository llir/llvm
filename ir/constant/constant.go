// === [ Constants ] ===========================================================
//
// References:
//    http://llvm.org/docs/LangRef.html#constants

// Package constant implements values representing immutable LLVM IR constants.
package constant

import (
	"github.com/llir/llvm/ir/types"
	"github.com/llir/llvm/ir/value"
)

// A Constant represents an LLVM IR constant; a value that is immutable at
// runtime, such as an integer or a floating point literal.
//
// Constant may have one of the following underlying types.
//
// Simple constants
//
// http://llvm.org/docs/LangRef.html#simple-constants
//
//    *constant.Int
//    *constant.Float
//    *constant.Null
//
// Complex constants
//
// http://llvm.org/docs/LangRef.html#complex-constants
//
//    *constant.Vector
//    *constant.Array
//    *constant.Struct
//    *constant.ZeroInitializer
//
// Constant expressions
//
// http://llvm.org/docs/LangRef.html#constant-expressions
//
//    constant.Expr
type Constant interface {
	value.Value
}

// Convenience constants.
var (
	// True represents the `true` constant.
	True = NewInt(1, types.I1)
	// False represents the `false` constant.
	False = NewInt(0, types.I1)
)
