package constant

import "github.com/llir/llvm/ir"

// Assert that each constant implements the ir.Constant interface.
var (
	// Constant expressions.
	// Checked in expression_test.go as ir.Expression embeds ir.Constant.
	_ ir.Constant = ir.Expression(nil)

	// Constants.
	_ ir.Constant = (*Int)(nil)
	_ ir.Constant = (*Float)(nil)
	_ ir.Constant = (*Null)(nil)
	_ ir.Constant = (*none)(nil)
	_ ir.Constant = (*Struct)(nil)
	_ ir.Constant = (*Array)(nil)
	_ ir.Constant = (*CharArray)(nil)
	_ ir.Constant = (*Vector)(nil)
	_ ir.Constant = (*ZeroInitializer)(nil)
	_ ir.Constant = (*Undef)(nil)
	_ ir.Constant = (*BlockAddress)(nil)
)
