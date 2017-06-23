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
//    *constant.Int     (https://godoc.org/github.com/llir/llvm/ir/constant#Int)
//    *constant.Float   (https://godoc.org/github.com/llir/llvm/ir/constant#Float)
//    *constant.Null    (https://godoc.org/github.com/llir/llvm/ir/constant#Null)
//
// Complex constants
//
// http://llvm.org/docs/LangRef.html#complex-constants
//
//    *constant.Vector            (https://godoc.org/github.com/llir/llvm/ir/constant#Vector)
//    *constant.Array             (https://godoc.org/github.com/llir/llvm/ir/constant#Array)
//    *constant.Struct            (https://godoc.org/github.com/llir/llvm/ir/constant#Struct)
//    *constant.ZeroInitializer   (https://godoc.org/github.com/llir/llvm/ir/constant#ZeroInitializer)
//
// Global variable and function addresses
//
//    *ir.Global     (https://godoc.org/github.com/llir/llvm/ir#Global)
//    *ir.Function   (https://godoc.org/github.com/llir/llvm/ir#Function)
//
// Undefined value constants
//
//    *constant.Undef   (https://godoc.org/github.com/llir/llvm/ir/constant#Undef)
//
// Constant expressions
//
// http://llvm.org/docs/LangRef.html#constant-expressions
//
//    constant.Expr   (https://godoc.org/github.com/llir/llvm/ir/constant#Expr)
type Constant interface {
	value.Value
	// Immutable ensures that only constants can be assigned to the
	// constant.Constant interface.
	Immutable()
}

// Convenience constants.
var (
	// True represents the `true` constant.
	True = NewInt(1, types.I1)
	// False represents the `false` constant.
	False = NewInt(0, types.I1)
)
