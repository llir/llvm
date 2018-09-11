package ir

// Assert that each constant implements the ir.Constant interface.
var (
	_ Constant = (*ConstInt)(nil)
	_ Constant = (*ConstFloat)(nil)
	_ Constant = (*ConstNull)(nil)
	_ Constant = (*ConstNone)(nil)
	_ Constant = (*ConstStruct)(nil)
	_ Constant = (*ConstArray)(nil)
	_ Constant = (*ConstVector)(nil)
	_ Constant = (*ConstZeroInitializer)(nil)
	_ Constant = (*Global)(nil)
	_ Constant = (*Function)(nil)
	_ Constant = (*ConstUndef)(nil)
	_ Constant = (*ConstBlockAddress)(nil)
)
