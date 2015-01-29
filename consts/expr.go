// TODO: Add isConst, isExpr methods to each constant type. Implement the Value
// interface for each constant type.

package consts

import (
	"github.com/mewlang/llvm/types"
	"github.com/mewlang/llvm/values"
)

// TODO: Complete the list of expression implementations.

// An Expr represents a constant expression.
//
// Expr is one of the following concrete types:
//    *consts.Trunc
//    *consts.ZExt
//    *consts.SExt
//    *consts.FPTrunc
//    *consts.FPExt
//    *consts.FPToUint
//
// References:
//    http://llvm.org/docs/LangRef.html#constant-expressions
type Expr interface {
	Constant
	// isExpr ensures that only constant expressions can be assigned to the Expr
	// interface.
	isExpr()
}

// Trunc is a constant expression which truncates an integer constant to a
// smaller or equally sized integer type.
//
// Examples:
//    trunc(i32 15 to i3)   ; yields i3:7
//
// References:
//    http://llvm.org/docs/LangRef.html#constant-expressions
type Trunc struct {
	// Original constant.
	v *Int
	// New type.
	to *types.Int
}

// ZExt is a constant expression which zero extends an integer constant to a
// larger integer type.
//
// Examples:
//    zext(i1 1 to i5)   ; yields i5:1
//
// References:
//    http://llvm.org/docs/LangRef.html#constant-expressions
type ZExt struct {
	// Original constant.
	v *Int
	// New type.
	to *types.Int
}

// SExt is a constant expression which sign extends an integer constant to a
// larger or equally sized integer type.
//
// Examples:
//    sext(i1 1 to i5)   ; yields i5:31
//
// References:
//    http://llvm.org/docs/LangRef.html#constant-expressions
type SExt struct {
	// Original constant.
	v *Int
	// New type.
	to *types.Int
}

// FPTrunc is a constant expression which truncates a floating point constant to
// a smaller or equally sized floating point type.
//
// Examples:
//    fptrunc(double 4.0 to float)   ; yields float:4.0
//
// References:
//    http://llvm.org/docs/LangRef.html#constant-expressions
type FPTrunc struct {
	// Original constant.
	v *Float
	// New type.
	to *types.Float
}

// FPExt is a constant expression which extends a floating point constant to a
// larger or equally sized floating point type.
//
// Examples:
//    fpext(float 4.0 to double)   ; yields double:4.0
//
// References:
//    http://llvm.org/docs/LangRef.html#constant-expressions
type FPExt struct {
	// Original constant.
	v *Float
	// New type.
	to *types.Float
}

// FPToUint is a constant expression which converts a floating point constant to
// the corresponding unsigned integer constant.
//
// Examples:
//    fptoui(float 4.0 to i32)   ; yields i32:4
//
// References:
//    http://llvm.org/docs/LangRef.html#constant-expressions
type FPToUint struct {
	// Original floating point value or vector of floating point values.
	v values.Value
	// New type.
	to *types.Int
}
