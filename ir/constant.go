package ir

import (
	"github.com/llir/llvm/ir/value"
)

// === [ Constants ] ===========================================================

// Constant is an LLVM IR constant; a value that is immutable at runtime, such
// as an integer or floating-point literal, or the address of a function or
// global variable.
//
// A Constant has one of the following underlying types.
//
// Simple constants
//
// https://llvm.org/docs/LangRef.html#simple-constants
//
//    *constant.Int     // https://godoc.org/github.com/llir/llvm/ir/constant#Int
//    *constant.Float   // https://godoc.org/github.com/llir/llvm/ir/constant#Float
//    *constant.Null    // https://godoc.org/github.com/llir/llvm/ir/constant#Null
//    *constant.None    // https://godoc.org/github.com/llir/llvm/ir/constant#None
//
// Complex constants
//
// https://llvm.org/docs/LangRef.html#complex-constants
//
//    *constant.Struct            // https://godoc.org/github.com/llir/llvm/ir/constant#Struct
//    *constant.Array             // https://godoc.org/github.com/llir/llvm/ir/constant#Array
//    *constant.CharArray         // https://godoc.org/github.com/llir/llvm/ir/constant#CharArray
//    *constant.Vector            // https://godoc.org/github.com/llir/llvm/ir/constant#Vector
//    *constant.ZeroInitializer   // https://godoc.org/github.com/llir/llvm/ir/constant#ZeroInitializer
//    TODO: include metadata node?
//
// Global variable and function addresses
//
// https://llvm.org/docs/LangRef.html#global-variable-and-function-addresses
//
//    *ir.Global     // https://godoc.org/github.com/llir/llvm/ir#Global
//    *ir.Function   // https://godoc.org/github.com/llir/llvm/ir#Function
//
// Undefined values
//
// https://llvm.org/docs/LangRef.html#undefined-values
//
//    *constant.Undef   // https://godoc.org/github.com/llir/llvm/ir/constant#Undef
//
// Addresses of basic blocks
//
// https://llvm.org/docs/LangRef.html#addresses-of-basic-blocks
//
//    *constant.BlockAddress   // https://godoc.org/github.com/llir/llvm/ir/constant#BlockAddress
//
// Constant expressions
//
// https://llvm.org/docs/LangRef.html#constant-expressions
//
//    constant.Expression   // https://godoc.org/github.com/llir/llvm/ir/constant#Expression
type Constant interface {
	value.Value
	// IsConstant ensures that only constants can be assigned to the ir.Constant
	// interface.
	IsConstant()
}

// IsConstant ensures that only constants can be assigned to the ir.Constant
// interface.
func (*Global) IsConstant() {}

// IsConstant ensures that only constants can be assigned to the ir.Constant
// interface.
func (*Function) IsConstant() {}
