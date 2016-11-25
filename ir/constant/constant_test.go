package constant_test

import (
	"github.com/llir/llvm/ir/constant"
)

// Valutate that the relevant types satisfy the constant.Constant interface.
var (
	// Simple constants.
	_ constant.Constant = &constant.Int{}
	_ constant.Constant = &constant.Float{}
	_ constant.Constant = &constant.Null{}
	// Complex constants.
	_ constant.Constant = &constant.Vector{}
	_ constant.Constant = &constant.Array{}
	_ constant.Constant = &constant.Struct{}
	_ constant.Constant = &constant.ZeroInitializer{}
	// Constant expressions.
)
