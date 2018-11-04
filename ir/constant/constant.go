// Package constant implements values representing immutable LLVM IR constants.
package constant

import (
	"github.com/llir/llvm/ir/types"
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

// IsConstant ensures that only constants can be assigned to the ir.Constant
// interface.
func (*Int) IsConstant() {}

// IsConstant ensures that only constants can be assigned to the ir.Constant
// interface.
func (*Float) IsConstant() {}

// IsConstant ensures that only constants can be assigned to the ir.Constant
// interface.
func (*Null) IsConstant() {}

// IsConstant ensures that only constants can be assigned to the ir.Constant
// interface.
func (*none) IsConstant() {}

// IsConstant ensures that only constants can be assigned to the ir.Constant
// interface.
func (*Struct) IsConstant() {}

// IsConstant ensures that only constants can be assigned to the ir.Constant
// interface.
func (*Array) IsConstant() {}

// IsConstant ensures that only constants can be assigned to the ir.Constant
// interface.
func (*CharArray) IsConstant() {}

// IsConstant ensures that only constants can be assigned to the ir.Constant
// interface.
func (*Vector) IsConstant() {}

// IsConstant ensures that only constants can be assigned to the ir.Constant
// interface.
func (*ZeroInitializer) IsConstant() {}

// IsConstant ensures that only constants can be assigned to the ir.Constant
// interface.
func (*Undef) IsConstant() {}

// IsConstant ensures that only constants can be assigned to the ir.Constant
// interface.
func (*BlockAddress) IsConstant() {}

// --- [ Binary expressions ] --------------------------------------------------

// IsConstant ensures that only constants can be assigned to the ir.Constant
// interface.
func (*ExprAdd) IsConstant() {}

// IsConstant ensures that only constants can be assigned to the ir.Constant
// interface.
func (*ExprFAdd) IsConstant() {}

// IsConstant ensures that only constants can be assigned to the ir.Constant
// interface.
func (*ExprSub) IsConstant() {}

// IsConstant ensures that only constants can be assigned to the ir.Constant
// interface.
func (*ExprFSub) IsConstant() {}

// IsConstant ensures that only constants can be assigned to the ir.Constant
// interface.
func (*ExprMul) IsConstant() {}

// IsConstant ensures that only constants can be assigned to the ir.Constant
// interface.
func (*ExprFMul) IsConstant() {}

// IsConstant ensures that only constants can be assigned to the ir.Constant
// interface.
func (*ExprUDiv) IsConstant() {}

// IsConstant ensures that only constants can be assigned to the ir.Constant
// interface.
func (*ExprSDiv) IsConstant() {}

// IsConstant ensures that only constants can be assigned to the ir.Constant
// interface.
func (*ExprFDiv) IsConstant() {}

// IsConstant ensures that only constants can be assigned to the ir.Constant
// interface.
func (*ExprURem) IsConstant() {}

// IsConstant ensures that only constants can be assigned to the ir.Constant
// interface.
func (*ExprSRem) IsConstant() {}

// IsConstant ensures that only constants can be assigned to the ir.Constant
// interface.
func (*ExprFRem) IsConstant() {}

// --- [ Bitwise expressions ] -------------------------------------------------

// IsConstant ensures that only constants can be assigned to the ir.Constant
// interface.
func (*ExprShl) IsConstant() {}

// IsConstant ensures that only constants can be assigned to the ir.Constant
// interface.
func (*ExprLShr) IsConstant() {}

// IsConstant ensures that only constants can be assigned to the ir.Constant
// interface.
func (*ExprAShr) IsConstant() {}

// IsConstant ensures that only constants can be assigned to the ir.Constant
// interface.
func (*ExprAnd) IsConstant() {}

// IsConstant ensures that only constants can be assigned to the ir.Constant
// interface.
func (*ExprOr) IsConstant() {}

// IsConstant ensures that only constants can be assigned to the ir.Constant
// interface.
func (*ExprXor) IsConstant() {}

// --- [ Vector expressions ] --------------------------------------------------

// IsConstant ensures that only constants can be assigned to the ir.Constant
// interface.
func (*ExprExtractElement) IsConstant() {}

// IsConstant ensures that only constants can be assigned to the ir.Constant
// interface.
func (*ExprInsertElement) IsConstant() {}

// IsConstant ensures that only constants can be assigned to the ir.Constant
// interface.
func (*ExprShuffleVector) IsConstant() {}

// --- [ Aggregate expressions ] -----------------------------------------------

// IsConstant ensures that only constants can be assigned to the ir.Constant
// interface.
func (*ExprExtractValue) IsConstant() {}

// IsConstant ensures that only constants can be assigned to the ir.Constant
// interface.
func (*ExprInsertValue) IsConstant() {}

// --- [ Memory expressions ] --------------------------------------------------

// IsConstant ensures that only constants can be assigned to the ir.Constant
// interface.
func (*ExprGetElementPtr) IsConstant() {}

// --- [ Conversion expressions ] ----------------------------------------------

// IsConstant ensures that only constants can be assigned to the ir.Constant
// interface.
func (*ExprTrunc) IsConstant() {}

// IsConstant ensures that only constants can be assigned to the ir.Constant
// interface.
func (*ExprZExt) IsConstant() {}

// IsConstant ensures that only constants can be assigned to the ir.Constant
// interface.
func (*ExprSExt) IsConstant() {}

// IsConstant ensures that only constants can be assigned to the ir.Constant
// interface.
func (*ExprFPTrunc) IsConstant() {}

// IsConstant ensures that only constants can be assigned to the ir.Constant
// interface.
func (*ExprFPExt) IsConstant() {}

// IsConstant ensures that only constants can be assigned to the ir.Constant
// interface.
func (*ExprFPToUI) IsConstant() {}

// IsConstant ensures that only constants can be assigned to the ir.Constant
// interface.
func (*ExprFPToSI) IsConstant() {}

// IsConstant ensures that only constants can be assigned to the ir.Constant
// interface.
func (*ExprUIToFP) IsConstant() {}

// IsConstant ensures that only constants can be assigned to the ir.Constant
// interface.
func (*ExprSIToFP) IsConstant() {}

// IsConstant ensures that only constants can be assigned to the ir.Constant
// interface.
func (*ExprPtrToInt) IsConstant() {}

// IsConstant ensures that only constants can be assigned to the ir.Constant
// interface.
func (*ExprIntToPtr) IsConstant() {}

// IsConstant ensures that only constants can be assigned to the ir.Constant
// interface.
func (*ExprBitCast) IsConstant() {}

// IsConstant ensures that only constants can be assigned to the ir.Constant
// interface.
func (*ExprAddrSpaceCast) IsConstant() {}

// --- [ Other expressions ] ---------------------------------------------------

// IsConstant ensures that only constants can be assigned to the ir.Constant
// interface.
func (*ExprICmp) IsConstant() {}

// IsConstant ensures that only constants can be assigned to the ir.Constant
// interface.
func (*ExprFCmp) IsConstant() {}

// IsConstant ensures that only constants can be assigned to the ir.Constant
// interface.
func (*ExprSelect) IsConstant() {}
