package ir

import "github.com/llir/l/ir/value"

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
//    *ir.ConstInt     // https://godoc.org/github.com/llir/l/ir#ConstInt
//    *ir.ConstFloat   // https://godoc.org/github.com/llir/l/ir#ConstFloat
//    *ir.ConstNull    // https://godoc.org/github.com/llir/l/ir#ConstNull
//    *ir.ConstNone    // https://godoc.org/github.com/llir/l/ir#ConstNone
//
// Complex constants
//
// https://llvm.org/docs/LangRef.html#complex-constants
//
//    *ir.ConstStruct            // https://godoc.org/github.com/llir/l/ir#ConstStruct
//    *ir.ConstArray             // https://godoc.org/github.com/llir/l/ir#ConstArray
//    *ir.ConstVector            // https://godoc.org/github.com/llir/l/ir#ConstVector
//    *ir.ConstZeroInitializer   // https://godoc.org/github.com/llir/l/ir#ConstZeroInitializer
//    TODO: include metadata node?
//
// Global variable and function addresses
//
// https://llvm.org/docs/LangRef.html#global-variable-and-function-addresses
//
//    *ir.Global     // https://godoc.org/github.com/llir/l/ir#Global
//    *ir.Function   // https://godoc.org/github.com/llir/l/ir#Function
//
// Undefined values
//
// https://llvm.org/docs/LangRef.html#undefined-values
//
//    *ir.ConstUndef   // https://godoc.org/github.com/llir/l/ir#ConstUndef
//
// Addresses of basic blocks
//
// https://llvm.org/docs/LangRef.html#addresses-of-basic-blocks
//
//    *ir.ConstBlockAddress   // https://godoc.org/github.com/llir/l/ir#ConstBlockAddress
//
// Constant expressions
//
// https://llvm.org/docs/LangRef.html#constant-expressions
//
//    ir.Expression   // https://godoc.org/github.com/llir/l/ir#Expression
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
