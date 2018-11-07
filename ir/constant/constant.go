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
	None = &none{} // none
	// Boolean constants.
	True  = NewInt(types.I1, 1) // true
	False = NewInt(types.I1, 0) // false
)

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
//    *ir.Alias      // https://godoc.org/github.com/llir/llvm/ir#Alias
//    *ir.IFunc      // https://godoc.org/github.com/llir/llvm/ir#IFunc
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
	// IsConstant ensures that only constants can be assigned to the
	// constant.Constant interface.
	IsConstant()
}

// IsConstant ensures that only constants can be assigned to the
// constant.Constant interface.
func (*Int) IsConstant() {}

// IsConstant ensures that only constants can be assigned to the
// constant.Constant interface.
func (*Float) IsConstant() {}

// IsConstant ensures that only constants can be assigned to the
// constant.Constant interface.
func (*Null) IsConstant() {}

// IsConstant ensures that only constants can be assigned to the
// constant.Constant interface.
func (*none) IsConstant() {}

// IsConstant ensures that only constants can be assigned to the
// constant.Constant interface.
func (*Struct) IsConstant() {}

// IsConstant ensures that only constants can be assigned to the
// constant.Constant interface.
func (*Array) IsConstant() {}

// IsConstant ensures that only constants can be assigned to the
// constant.Constant interface.
func (*CharArray) IsConstant() {}

// IsConstant ensures that only constants can be assigned to the
// constant.Constant interface.
func (*Vector) IsConstant() {}

// IsConstant ensures that only constants can be assigned to the
// constant.Constant interface.
func (*ZeroInitializer) IsConstant() {}

// IsConstant ensures that only constants can be assigned to the
// constant.Constant interface.
func (*Undef) IsConstant() {}

// IsConstant ensures that only constants can be assigned to the
// constant.Constant interface.
func (*BlockAddress) IsConstant() {}

// --- [ Binary expressions ] --------------------------------------------------

// IsConstant ensures that only constants can be assigned to the
// constant.Constant interface.
func (*ExprAdd) IsConstant() {}

// IsConstant ensures that only constants can be assigned to the
// constant.Constant interface.
func (*ExprFAdd) IsConstant() {}

// IsConstant ensures that only constants can be assigned to the
// constant.Constant interface.
func (*ExprSub) IsConstant() {}

// IsConstant ensures that only constants can be assigned to the
// constant.Constant interface.
func (*ExprFSub) IsConstant() {}

// IsConstant ensures that only constants can be assigned to the
// constant.Constant interface.
func (*ExprMul) IsConstant() {}

// IsConstant ensures that only constants can be assigned to the
// constant.Constant interface.
func (*ExprFMul) IsConstant() {}

// IsConstant ensures that only constants can be assigned to the
// constant.Constant interface.
func (*ExprUDiv) IsConstant() {}

// IsConstant ensures that only constants can be assigned to the
// constant.Constant interface.
func (*ExprSDiv) IsConstant() {}

// IsConstant ensures that only constants can be assigned to the
// constant.Constant interface.
func (*ExprFDiv) IsConstant() {}

// IsConstant ensures that only constants can be assigned to the
// constant.Constant interface.
func (*ExprURem) IsConstant() {}

// IsConstant ensures that only constants can be assigned to the
// constant.Constant interface.
func (*ExprSRem) IsConstant() {}

// IsConstant ensures that only constants can be assigned to the
// constant.Constant interface.
func (*ExprFRem) IsConstant() {}

// --- [ Bitwise expressions ] -------------------------------------------------

// IsConstant ensures that only constants can be assigned to the
// constant.Constant interface.
func (*ExprShl) IsConstant() {}

// IsConstant ensures that only constants can be assigned to the
// constant.Constant interface.
func (*ExprLShr) IsConstant() {}

// IsConstant ensures that only constants can be assigned to the
// constant.Constant interface.
func (*ExprAShr) IsConstant() {}

// IsConstant ensures that only constants can be assigned to the
// constant.Constant interface.
func (*ExprAnd) IsConstant() {}

// IsConstant ensures that only constants can be assigned to the
// constant.Constant interface.
func (*ExprOr) IsConstant() {}

// IsConstant ensures that only constants can be assigned to the
// constant.Constant interface.
func (*ExprXor) IsConstant() {}

// --- [ Vector expressions ] --------------------------------------------------

// IsConstant ensures that only constants can be assigned to the
// constant.Constant interface.
func (*ExprExtractElement) IsConstant() {}

// IsConstant ensures that only constants can be assigned to the
// constant.Constant interface.
func (*ExprInsertElement) IsConstant() {}

// IsConstant ensures that only constants can be assigned to the
// constant.Constant interface.
func (*ExprShuffleVector) IsConstant() {}

// --- [ Aggregate expressions ] -----------------------------------------------

// IsConstant ensures that only constants can be assigned to the
// constant.Constant interface.
func (*ExprExtractValue) IsConstant() {}

// IsConstant ensures that only constants can be assigned to the
// constant.Constant interface.
func (*ExprInsertValue) IsConstant() {}

// --- [ Memory expressions ] --------------------------------------------------

// IsConstant ensures that only constants can be assigned to the
// constant.Constant interface.
func (*ExprGetElementPtr) IsConstant() {}

// --- [ Conversion expressions ] ----------------------------------------------

// IsConstant ensures that only constants can be assigned to the
// constant.Constant interface.
func (*ExprTrunc) IsConstant() {}

// IsConstant ensures that only constants can be assigned to the
// constant.Constant interface.
func (*ExprZExt) IsConstant() {}

// IsConstant ensures that only constants can be assigned to the
// constant.Constant interface.
func (*ExprSExt) IsConstant() {}

// IsConstant ensures that only constants can be assigned to the
// constant.Constant interface.
func (*ExprFPTrunc) IsConstant() {}

// IsConstant ensures that only constants can be assigned to the
// constant.Constant interface.
func (*ExprFPExt) IsConstant() {}

// IsConstant ensures that only constants can be assigned to the
// constant.Constant interface.
func (*ExprFPToUI) IsConstant() {}

// IsConstant ensures that only constants can be assigned to the
// constant.Constant interface.
func (*ExprFPToSI) IsConstant() {}

// IsConstant ensures that only constants can be assigned to the
// constant.Constant interface.
func (*ExprUIToFP) IsConstant() {}

// IsConstant ensures that only constants can be assigned to the
// constant.Constant interface.
func (*ExprSIToFP) IsConstant() {}

// IsConstant ensures that only constants can be assigned to the
// constant.Constant interface.
func (*ExprPtrToInt) IsConstant() {}

// IsConstant ensures that only constants can be assigned to the
// constant.Constant interface.
func (*ExprIntToPtr) IsConstant() {}

// IsConstant ensures that only constants can be assigned to the
// constant.Constant interface.
func (*ExprBitCast) IsConstant() {}

// IsConstant ensures that only constants can be assigned to the
// constant.Constant interface.
func (*ExprAddrSpaceCast) IsConstant() {}

// --- [ Other expressions ] ---------------------------------------------------

// IsConstant ensures that only constants can be assigned to the
// constant.Constant interface.
func (*ExprICmp) IsConstant() {}

// IsConstant ensures that only constants can be assigned to the
// constant.Constant interface.
func (*ExprFCmp) IsConstant() {}

// IsConstant ensures that only constants can be assigned to the
// constant.Constant interface.
func (*ExprSelect) IsConstant() {}
