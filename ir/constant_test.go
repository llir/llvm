package ir

import "github.com/llir/l/ir/constant"

// Assert that each constant implements the constant.Constant interface.
var (
	_ constant.Constant = (*Global)(nil)
	_ constant.Constant = (*Function)(nil)
)
