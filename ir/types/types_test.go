package types_test

import (
	"github.com/llir/llvm/ir/types"
	"github.com/llir/llvm/ir/value"
)

// Validates that the value.Value interface is implemented by the relevant
// types.
var (
	_ value.Value = &types.Param{}
)
