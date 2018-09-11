package constant

// Assert that each constant expression implements the constant.Expression
// interface.
var (
	// Binary expressions.
	_ Expression = (*Add)(nil)
	_ Expression = (*FAdd)(nil)
	_ Expression = (*Sub)(nil)
	_ Expression = (*FSub)(nil)
	_ Expression = (*Mul)(nil)
	_ Expression = (*FMul)(nil)
	_ Expression = (*UDiv)(nil)
	_ Expression = (*SDiv)(nil)
	_ Expression = (*FDiv)(nil)
	_ Expression = (*URem)(nil)
	_ Expression = (*SRem)(nil)
	_ Expression = (*FRem)(nil)
	// Bitwise expressions.
	_ Expression = (*Shl)(nil)
	_ Expression = (*LShr)(nil)
	_ Expression = (*AShr)(nil)
	_ Expression = (*And)(nil)
	_ Expression = (*Or)(nil)
	_ Expression = (*Xor)(nil)
	// Vector expressions.
	_ Expression = (*ExtractElement)(nil)
	_ Expression = (*InsertElement)(nil)
	_ Expression = (*ShuffleVector)(nil)
	// Aggregate expressions.
	_ Expression = (*ExtractValue)(nil)
	_ Expression = (*InsertValue)(nil)
	// Memory expressions.
	_ Expression = (*GetElementPtr)(nil)
	// Conversion expressions.
	_ Expression = (*Trunc)(nil)
	_ Expression = (*ZExt)(nil)
	_ Expression = (*SExt)(nil)
	_ Expression = (*FPTrunc)(nil)
	_ Expression = (*FPExt)(nil)
	_ Expression = (*FPToUI)(nil)
	_ Expression = (*FPToSI)(nil)
	_ Expression = (*UIToFP)(nil)
	_ Expression = (*SIToFP)(nil)
	_ Expression = (*PtrToInt)(nil)
	_ Expression = (*IntToPtr)(nil)
	_ Expression = (*BitCast)(nil)
	_ Expression = (*AddrSpaceCast)(nil)
	// Other expressions.
	_ Expression = (*ICmp)(nil)
	_ Expression = (*FCmp)(nil)
	_ Expression = (*Select)(nil)
)
