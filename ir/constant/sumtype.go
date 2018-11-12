package constant

// === [ constant.Constant ] ===================================================

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
func (*NoneToken) IsConstant() {}

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
