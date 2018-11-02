package ir

// Assert that each constant implements the ir.Constant interface.
var (
	// Constant expressions.
	// Checked in expression_test.go as ir.Expression embeds ir.Constant.
	_ Constant = Expression(nil)

	// Constants.
	_ Constant = (*ConstInt)(nil)
	_ Constant = (*ConstFloat)(nil)
	_ Constant = (*ConstNull)(nil)
	_ Constant = (*ConstNone)(nil)
	_ Constant = (*ConstStruct)(nil)
	_ Constant = (*ConstArray)(nil)
	_ Constant = (*ConstCharArray)(nil)
	_ Constant = (*ConstVector)(nil)
	_ Constant = (*ConstZeroInitializer)(nil)
	_ Constant = (*Global)(nil)
	_ Constant = (*Function)(nil)
	_ Constant = (*ConstUndef)(nil)
	_ Constant = (*ConstBlockAddress)(nil)
)
