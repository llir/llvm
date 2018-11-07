package ir

// IsConstant ensures that only constants can be assigned to the
// constant.Constant interface.
func (*Global) IsConstant() {}

// IsConstant ensures that only constants can be assigned to the
// constant.Constant interface.
func (*Function) IsConstant() {}

// IsConstant ensures that only constants can be assigned to the
// constant.Constant interface.
func (*Alias) IsConstant() {}

// IsConstant ensures that only constants can be assigned to the
// constant.Constant interface.
func (*IFunc) IsConstant() {}
