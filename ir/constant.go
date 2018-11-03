package ir

import (
	"github.com/llir/llvm/ir/types"
	"github.com/llir/llvm/ir/value"
)

// === [ Constants ] ===========================================================

// Convenience constants.
var (
	// None token constant.
	None = &ConstNone{} // none
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
//    *ir.ConstInt     // https://godoc.org/github.com/llir/llvm/ir#ConstInt
//    *ir.ConstFloat   // https://godoc.org/github.com/llir/llvm/ir#ConstFloat
//    *ir.ConstNull    // https://godoc.org/github.com/llir/llvm/ir#ConstNull
//    *ir.ConstNone    // https://godoc.org/github.com/llir/llvm/ir#ConstNone
//
// Complex constants
//
// https://llvm.org/docs/LangRef.html#complex-constants
//
//    *ir.ConstStruct            // https://godoc.org/github.com/llir/llvm/ir#ConstStruct
//    *ir.ConstArray             // https://godoc.org/github.com/llir/llvm/ir#ConstArray
//    *ir.ConstCharArray         // https://godoc.org/github.com/llir/llvm/ir#ConstCharArray
//    *ir.ConstVector            // https://godoc.org/github.com/llir/llvm/ir#ConstVector
//    *ir.ConstZeroInitializer   // https://godoc.org/github.com/llir/llvm/ir#ConstZeroInitializer
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
//    *ir.ConstUndef   // https://godoc.org/github.com/llir/llvm/ir#ConstUndef
//
// Addresses of basic blocks
//
// https://llvm.org/docs/LangRef.html#addresses-of-basic-blocks
//
//    *ir.ConstBlockAddress   // https://godoc.org/github.com/llir/llvm/ir#ConstBlockAddress
//
// Constant expressions
//
// https://llvm.org/docs/LangRef.html#constant-expressions
//
//    ir.Expression   // https://godoc.org/github.com/llir/llvm/ir#Expression
type Constant interface {
	value.Value
	// isConstant ensures that only constants can be assigned to the ir.Constant
	// interface.
	isConstant()
}

// isConstant ensures that only constants can be assigned to the ir.Constant
// interface.
func (*ConstInt) isConstant()             {}
func (*ConstFloat) isConstant()           {}
func (*ConstNull) isConstant()            {}
func (*ConstNone) isConstant()            {}
func (*ConstStruct) isConstant()          {}
func (*ConstArray) isConstant()           {}
func (*ConstCharArray) isConstant()       {}
func (*ConstVector) isConstant()          {}
func (*ConstZeroInitializer) isConstant() {}
func (*Global) isConstant()               {}
func (*Function) isConstant()             {}
func (*ConstUndef) isConstant()           {}
func (*ConstBlockAddress) isConstant()    {}

// Binary expressions.
func (*ExprAdd) isConstant()  {}
func (*ExprFAdd) isConstant() {}
func (*ExprSub) isConstant()  {}
func (*ExprFSub) isConstant() {}
func (*ExprMul) isConstant()  {}
func (*ExprFMul) isConstant() {}
func (*ExprUDiv) isConstant() {}
func (*ExprSDiv) isConstant() {}
func (*ExprFDiv) isConstant() {}
func (*ExprURem) isConstant() {}
func (*ExprSRem) isConstant() {}
func (*ExprFRem) isConstant() {}

// Bitwise expressions.
func (*ExprShl) isConstant()  {}
func (*ExprLShr) isConstant() {}
func (*ExprAShr) isConstant() {}
func (*ExprAnd) isConstant()  {}
func (*ExprOr) isConstant()   {}
func (*ExprXor) isConstant()  {}

// Vector expressions.
func (*ExprExtractElement) isConstant() {}
func (*ExprInsertElement) isConstant()  {}
func (*ExprShuffleVector) isConstant()  {}

// Aggregate expressions.
func (*ExprExtractValue) isConstant() {}
func (*ExprInsertValue) isConstant()  {}

// Memory expressions.
func (*ExprGetElementPtr) isConstant() {}

// Conversion expressions.
func (*ExprTrunc) isConstant()         {}
func (*ExprZExt) isConstant()          {}
func (*ExprSExt) isConstant()          {}
func (*ExprFPTrunc) isConstant()       {}
func (*ExprFPExt) isConstant()         {}
func (*ExprFPToUI) isConstant()        {}
func (*ExprFPToSI) isConstant()        {}
func (*ExprUIToFP) isConstant()        {}
func (*ExprSIToFP) isConstant()        {}
func (*ExprPtrToInt) isConstant()      {}
func (*ExprIntToPtr) isConstant()      {}
func (*ExprBitCast) isConstant()       {}
func (*ExprAddrSpaceCast) isConstant() {}

// Other expressions.
func (*ExprICmp) isConstant()   {}
func (*ExprFCmp) isConstant()   {}
func (*ExprSelect) isConstant() {}
