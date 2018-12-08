package constant

// === [ constant.Constant ] ===================================================

// constantSumtype implements constant.Constant.
type constantSumtype struct{}

// IsConstant ensures that only constants can be assigned to the
// constant.Constant interface.
func (constantSumtype) IsConstant() {}
