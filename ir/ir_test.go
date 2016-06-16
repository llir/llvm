package ir_test

import (
	"github.com/llir/llvm/ir"
	"github.com/llir/llvm/ir/value"
)

// Ensure that each value implements the Value interface.
var (
	_ value.Value = &ir.GlobalDecl{}
	_ value.Value = &ir.Function{}
	_ value.Value = &ir.BasicBlock{}
)
