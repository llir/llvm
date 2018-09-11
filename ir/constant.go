package ir

import "github.com/llir/l/ir/value"

// Constant is an LLVM IR constant.
// TODO: document the underlying types of Constant.
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
func (*Global) isConstant()          {}
func (*Function) isConstant()        {}
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
