package ir

import "github.com/llir/l/ir/value"

// Constant is an LLVM IR constant.
// TODO: document the underlying types of Constant.
type Constant interface {
	value.Value
	// isConstant ensures that only constants can be assigned to the ir.Constant
	// interface.
	isConstant()
}

// isConstant ensures that only constants can be assigned to the ir.Constant
// interface.
func (*Int) isConstant()             {}
func (*Float) isConstant()           {}
func (*Null) isConstant()            {}
func (*None) isConstant()            {}
func (*Struct) isConstant()          {}
func (*Array) isConstant()           {}
func (*Vector) isConstant()          {}
func (*ZeroInitializer) isConstant() {}
func (*Global) isConstant()          {}
func (*Function) isConstant()        {}
func (*Undef) isConstant()           {}
func (*BlockAddress) isConstant()    {}

// Binary expressions.
func (*AddExpr) isConstant()  {}
func (*FAddExpr) isConstant() {}
func (*SubExpr) isConstant()  {}
func (*FSubExpr) isConstant() {}
func (*MulExpr) isConstant()  {}
func (*FMulExpr) isConstant() {}
func (*UDivExpr) isConstant() {}
func (*SDivExpr) isConstant() {}
func (*FDivExpr) isConstant() {}
func (*URemExpr) isConstant() {}
func (*SRemExpr) isConstant() {}
func (*FRemExpr) isConstant() {}

// Bitwise expressions.
func (*ShlExpr) isConstant()  {}
func (*LShrExpr) isConstant() {}
func (*AShrExpr) isConstant() {}
func (*AndExpr) isConstant()  {}
func (*OrExpr) isConstant()   {}
func (*XorExpr) isConstant()  {}

// Vector expressions.
func (*ExtractElementExpr) isConstant() {}
func (*InsertElementExpr) isConstant()  {}
func (*ShuffleVectorExpr) isConstant()  {}

// Aggregate expressions.
func (*ExtractValueExpr) isConstant() {}
func (*InsertValueExpr) isConstant()  {}

// Memory expressions.
func (*GetElementPtrExpr) isConstant() {}

// Conversion expressions.
func (*TruncExpr) isConstant()         {}
func (*ZExtExpr) isConstant()          {}
func (*SExtExpr) isConstant()          {}
func (*FPTruncExpr) isConstant()       {}
func (*FPExtExpr) isConstant()         {}
func (*FPToUIExpr) isConstant()        {}
func (*FPToSIExpr) isConstant()        {}
func (*UIToFPExpr) isConstant()        {}
func (*SIToFPExpr) isConstant()        {}
func (*PtrToIntExpr) isConstant()      {}
func (*IntToPtrExpr) isConstant()      {}
func (*BitCastExpr) isConstant()       {}
func (*AddrSpaceCastExpr) isConstant() {}

// Other expressions.
func (*ICmpExpr) isConstant()   {}
func (*FCmpExpr) isConstant()   {}
func (*SelectExpr) isConstant() {}
