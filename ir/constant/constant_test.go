package constant

// Assert that each constant implements the constant.Constant interface.
var (
	// Constant expressions. Checked in expression_test.go as constant.Expression
	// embeds constant.Constant.
	_ Constant = Expression(nil)

	// Constants.
	_ Constant = (*Int)(nil)
	_ Constant = (*Float)(nil)
	_ Constant = (*Null)(nil)
	_ Constant = (*none)(nil)
	_ Constant = (*Struct)(nil)
	_ Constant = (*Array)(nil)
	_ Constant = (*CharArray)(nil)
	_ Constant = (*Vector)(nil)
	_ Constant = (*ZeroInitializer)(nil)
	_ Constant = (*Undef)(nil)
	_ Constant = (*BlockAddress)(nil)
)
