package constant

import "github.com/llir/l/ir/value"

// Constant is an LLVM IR constant.
// TODO: document the underlying types of Constant.
type Constant interface {
	value.Value
	// IsConstant ensures that only constants can be assigned to the
	// constant.Constant interface.
	IsConstant()
}

// IsConstant ensures that only constants can be assigned to the
// constant.Constant interface.
func (*Int) IsConstant()             {}
func (*Float) IsConstant()           {}
func (*Null) IsConstant()            {}
func (*None) IsConstant()            {}
func (*Struct) IsConstant()          {}
func (*Array) IsConstant()           {}
func (*Vector) IsConstant()          {}
func (*ZeroInitializer) IsConstant() {}
func (*Undef) IsConstant()           {}
func (*BlockAddress) IsConstant()    {}

// Binary expressions.
func (*Add) IsConstant()  {}
func (*FAdd) IsConstant() {}
func (*Sub) IsConstant()  {}
func (*FSub) IsConstant() {}
func (*Mul) IsConstant()  {}
func (*FMul) IsConstant() {}
func (*UDiv) IsConstant() {}
func (*SDiv) IsConstant() {}
func (*FDiv) IsConstant() {}
func (*URem) IsConstant() {}
func (*SRem) IsConstant() {}
func (*FRem) IsConstant() {}

// Bitwise expressions.
func (*Shl) IsConstant()  {}
func (*LShr) IsConstant() {}
func (*AShr) IsConstant() {}
func (*And) IsConstant()  {}
func (*Or) IsConstant()   {}
func (*Xor) IsConstant()  {}

// Vector expressions.
func (*ExtractElement) IsConstant() {}
func (*InsertElement) IsConstant()  {}
func (*ShuffleVector) IsConstant()  {}

// Aggregate expressions.
func (*ExtractValue) IsConstant() {}
func (*InsertValue) IsConstant()  {}

// Memory expressions.
func (*GetElementPtr) IsConstant() {}

// Conversion expressions.
func (*Trunc) IsConstant()         {}
func (*ZExt) IsConstant()          {}
func (*SExt) IsConstant()          {}
func (*FPTrunc) IsConstant()       {}
func (*FPExt) IsConstant()         {}
func (*FPToUI) IsConstant()        {}
func (*FPToSI) IsConstant()        {}
func (*UIToFP) IsConstant()        {}
func (*SIToFP) IsConstant()        {}
func (*PtrToInt) IsConstant()      {}
func (*IntToPtr) IsConstant()      {}
func (*BitCast) IsConstant()       {}
func (*AddrSpaceCast) IsConstant() {}

// Other expressions.
func (*ICmp) IsConstant()   {}
func (*FCmp) IsConstant()   {}
func (*Select) IsConstant() {}
