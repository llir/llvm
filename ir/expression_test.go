package ir

// Assert that each constant expression implements the ir.Expression interface.
var (
	// Binary expressions.
	_ Expression = (*AddExpr)(nil)
	_ Expression = (*FAddExpr)(nil)
	_ Expression = (*SubExpr)(nil)
	_ Expression = (*FSubExpr)(nil)
	_ Expression = (*MulExpr)(nil)
	_ Expression = (*FMulExpr)(nil)
	_ Expression = (*UDivExpr)(nil)
	_ Expression = (*SDivExpr)(nil)
	_ Expression = (*FDivExpr)(nil)
	_ Expression = (*URemExpr)(nil)
	_ Expression = (*SRemExpr)(nil)
	_ Expression = (*FRemExpr)(nil)
	// Bitwise expressions.
	_ Expression = (*ShlExpr)(nil)
	_ Expression = (*LShrExpr)(nil)
	_ Expression = (*AShrExpr)(nil)
	_ Expression = (*AndExpr)(nil)
	_ Expression = (*OrExpr)(nil)
	_ Expression = (*XorExpr)(nil)
	// Vector expressions.
	_ Expression = (*ExtractElementExpr)(nil)
	_ Expression = (*InsertElementExpr)(nil)
	_ Expression = (*ShuffleVectorExpr)(nil)
	// Aggregate expressions.
	_ Expression = (*ExtractValueExpr)(nil)
	_ Expression = (*InsertValueExpr)(nil)
	// Memory expressions.
	_ Expression = (*GetElementPtrExpr)(nil)
	// Conversion expressions.
	_ Expression = (*TruncExpr)(nil)
	_ Expression = (*ZExtExpr)(nil)
	_ Expression = (*SExtExpr)(nil)
	_ Expression = (*FPTruncExpr)(nil)
	_ Expression = (*FPExtExpr)(nil)
	_ Expression = (*FPToUIExpr)(nil)
	_ Expression = (*FPToSIExpr)(nil)
	_ Expression = (*UIToFPExpr)(nil)
	_ Expression = (*SIToFPExpr)(nil)
	_ Expression = (*PtrToIntExpr)(nil)
	_ Expression = (*IntToPtrExpr)(nil)
	_ Expression = (*BitCastExpr)(nil)
	_ Expression = (*AddrSpaceCastExpr)(nil)
	// Other expressions.
	_ Expression = (*ICmpExpr)(nil)
	_ Expression = (*FCmpExpr)(nil)
	_ Expression = (*SelectExpr)(nil)
)
