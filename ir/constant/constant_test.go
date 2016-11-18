package constant_test

import (
	"github.com/llir/llvm/ir/constant"
)

// Validates that the constant.Constant interface is implemented by the relevant
// types.
var (
	_ constant.Constant = &constant.Int{}
)
