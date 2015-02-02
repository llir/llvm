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
//    *consts.IntTrunc
//    *consts.IntZeroExt
//    *consts.IntSignExt
//    *consts.FloatTrunc
//    *consts.FloatExt
//    *consts.FloatToUint
//    *consts.FloatToInt
//    *consts.UintToFloat
//    *consts.IntToFloat
//
// References:
//    http://llvm.org/docs/LangRef.html#constant-expressions
type Expr interface {
	Constant
	// isExpr ensures that only constant expressions can be assigned to the Expr
	// interface.
	isExpr()
}

// IntTrunc is a constant expression which truncates an integer constant to a
// smaller or equally sized integer type.
//
// Examples:
//    trunc(i32 15 to i3)   ; yields i3:7
//
// References:
//    http://llvm.org/docs/LangRef.html#constant-expressions
type IntTrunc struct {
	// Original constant.
	v *Int
	// New type.
	to *types.Int
}

// IntZeroExt is a constant expression which zero extends an integer constant to
// a larger or equally sized integer type.
//
// Examples:
//    zext(i1 1 to i5)   ; yields i5:1
//
// References:
//    http://llvm.org/docs/LangRef.html#constant-expressions
type IntZeroExt struct {
	// Original integer constant.
	v *Int
	// New integer type.
	to *types.Int
}

// IntSignExt is a constant expression which sign extends an integer constant to
// a larger or equally sized integer type.
//
// Examples:
//    sext(i1 1 to i5)   ; yields i5:31
//
// References:
//    http://llvm.org/docs/LangRef.html#constant-expressions
type IntSignExt struct {
	// Original integer constant.
	v *Int
	// New integer type.
	to *types.Int
}

// FloatTrunc is a constant expression which truncates a floating point constant
// to a smaller or equally sized floating point type.
//
// Examples:
//    fptrunc(double 4.0 to float)   ; yields float:4.0
//
// References:
//    http://llvm.org/docs/LangRef.html#constant-expressions
type FloatTrunc struct {
	// Original floating point constant.
	v *Float
	// New floating point type.
	to *types.Float
}

// FloatExt is a constant expression which extends a floating point constant to
// a larger or equally sized floating point type.
//
// Examples:
//    fpext(float 4.0 to double)   ; yields double:4.0
//
// References:
//    http://llvm.org/docs/LangRef.html#constant-expressions
type FloatExt struct {
	// Original floating point constant.
	v *Float
	// New floating point type.
	to *types.Float
}

// FloatToUint is a constant expression which converts a floating point constant
// (or constant vector) to the corresponding unsigned integer constant (or
// constant vector).
//
// Examples:
//    fptoui(float 4.0 to i32)                        ; yields i32:4
//    fptoui(<1 x float> <float 3.0> to <1 x i32>))   ; yields <1 x i32>:<i32 3>
//
// References:
//    http://llvm.org/docs/LangRef.html#constant-expressions
type FloatToUint struct {
	// Original floating point value (or vector).
	v values.Value
	// New integer type.
	to *types.Int
}

// FloatToInt is a constant expression which converts a floating point constant
// (or constant vector) to the corresponding signed integer constant (or
// constant vector).
//
// Examples:
//    fptosi(float -4.0 to i32)                        ; yields i32:-4
//    fptosi(<1 x float> <float -3.0> to <1 x i32>))   ; yields <1 x i32>:<i32 -3>
//
// References:
//    http://llvm.org/docs/LangRef.html#constant-expressions
type FloatToInt struct {
	// Original floating point value (or vector).
	v values.Value
	// New type.
	to *types.Int
}

// UintToFloat is a constant expression which converts an unsigned integer
// constant (or constant vector) to the corresponding floating point constant
// (or constant vector).
//
// Examples:
//    uitofp(i32 4 to float)                      ; yields float:4.0
//    uitofp(<1 x i32> <i32 3> to <1 x float>))   ; yields <1 x float>:<float 3.0>
//
// References:
//    http://llvm.org/docs/LangRef.html#constant-expressions
type UintToFloat struct {
	// Original unsigned integer value (or vector).
	v values.Value
	// New floating point type.
	to *types.Float
}

// IntToFloat is a constant expression which converts a signed integer constant
// (or constant vector) to the corresponding floating point constant (or
// constant vector).
//
// Examples:
//    sitofp(i32 -4 to float)                      ; yields float:-4.0
//    sitofp(<1 x i32> <i32 -3> to <1 x float>))   ; yields <1 x float>:<float -3.0>
//
// References:
//    http://llvm.org/docs/LangRef.html#constant-expressions
type IntToFloat struct {
	// Original signed integer value (or vector).
	v values.Value
	// New floating point type.
	to *types.Float
}

// TODO: Add support for the following constant expressions:
//    - ptrtoint
//    - inttoptr
//    - bitcast
//    - addrspacecast
//    - getelementptr
//    - select
//    - icmp
//    - fcmp
//    - extractelement
//    - insertelement
//    - shufflevector
//    - extractvalue
//    - insertvalue
//    - OPCODE (LHS, RHS)
//         * OPCODE may be any of the binary or bitwise binary operations.
