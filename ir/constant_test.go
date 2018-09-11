package ir

// Assert that each constant implements the ir.Constant interface.
var (
	_ Constant = (*Int)(nil)
	_ Constant = (*Float)(nil)
	_ Constant = (*Null)(nil)
	_ Constant = (*None)(nil)
	_ Constant = (*Struct)(nil)
	_ Constant = (*Array)(nil)
	_ Constant = (*Vector)(nil)
	_ Constant = (*ZeroInitializer)(nil)
	_ Constant = (*Global)(nil)
	_ Constant = (*Function)(nil)
	_ Constant = (*Undef)(nil)
	_ Constant = (*BlockAddress)(nil)
)
