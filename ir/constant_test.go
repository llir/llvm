package ir

import "github.com/llir/llvm/ir/constant"

// Assert that each constant implements the constant.Constant interface.
var (
	// Constants.
	_ constant.Constant = (*Global)(nil)
	_ constant.Constant = (*Function)(nil)
	_ constant.Constant = (*Alias)(nil)
	_ constant.Constant = (*IFunc)(nil)
)
