package ir

// Assert that each constant implements the ir.Constant interface.
var (
	// Constants.
	_ Constant = (*Global)(nil)
	_ Constant = (*Function)(nil)
)
