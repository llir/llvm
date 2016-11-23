package types_test

import (
	"github.com/llir/llvm/ir/types"
	"github.com/llir/llvm/ir/value"
)

// Valutate that the relevant types satisfy the value.Value interface.
var (
	_ value.Value = &types.Param{}
)

// Valutate that the relevant types satisfy the types.Type interface.
var (
	_ types.Type = &types.VoidType{}
	_ types.Type = &types.LabelType{}
	_ types.Type = &types.IntType{}
	_ types.Type = &types.FloatType{}
	_ types.Type = &types.FuncType{}
	_ types.Type = &types.PointerType{}
	_ types.Type = &types.VectorType{}
	_ types.Type = &types.ArrayType{}
	_ types.Type = &types.StructType{}
)
