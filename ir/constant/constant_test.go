package constant_test

import (
	"github.com/llir/llvm/ir/constant"
)

// Valutate that the relevant types satisfy the constant.Constant interface.
var (
	_ constant.Constant = &constant.Int{}
	_ constant.Constant = &constant.Float{}
)
