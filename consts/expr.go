package consts

import (
	"fmt"

	"github.com/llir/llvm/types"
	"github.com/llir/llvm/value"
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
	// Calc calculates and returns a constant which is equivalent to the constant
	// expression.
	Calc() Constant
}

// Make sure that each constant expression implements the Expr interface.
var (
	_ Expr = &IntTrunc{}
	_ Expr = &IntZeroExt{}
	_ Expr = &IntSignExt{}
	_ Expr = &FloatTrunc{}
	_ Expr = &FloatExt{}
	_ Expr = &FloatToUint{}
	_ Expr = &FloatToInt{}
	_ Expr = &UintToFloat{}
	_ Expr = &IntToFloat{}
)

// IntTrunc is a constant expression which truncates an integer constant to a
// smaller or equally sized integer type.
//
// Examples:
//    trunc(i32 15 to i3)   ; yields i3:7
//
// References:
//    http://llvm.org/docs/LangRef.html#constant-expressions
type IntTrunc struct {
	// Original integer constant.
	orig *Int
	// New integer type.
	to *types.Int
}

// NewIntTrunc returns a constant expression which truncates the integer
// constant orig to a smaller or equally sized integer type.
func NewIntTrunc(orig Constant, to types.Type) (*IntTrunc, error) {
	// Verify type of original integer constant.
	exp := new(IntTrunc)
	var ok bool
	exp.orig, ok = orig.(*Int)
	if !ok {
		return nil, fmt.Errorf("invalid integer truncation; expected integer constant for orig, got %q", orig.Type())
	}

	// Verify target type.
	exp.to, ok = to.(*types.Int)
	if !ok {
		return nil, fmt.Errorf("invalid integer truncation; expected integer target type, got %q", to)
	}
	newSize, origSize := exp.to.Size(), exp.orig.typ.Size()
	if newSize > origSize {
		return nil, fmt.Errorf("invalid integer truncation; target size (%d) larger than original size (%d)", newSize, origSize)
	}

	return exp, nil
}

// Type returns the type of the value.
func (exp *IntTrunc) Type() types.Type {
	return exp.to
}

// Calc calculates and returns a constant which is equivalent to the constant
// expression.
func (exp *IntTrunc) Calc() Constant {
	panic("IntTrunc.Calc: not yet implemented.")
}

// String returns a string representation of the integer truncation expression.
// The expression string representation is preceded by the type of the constant,
// e.g.
//
//    i3 trunc(i32 15 to i3)
func (exp *IntTrunc) String() string {
	return fmt.Sprintf("%s trunc(%s to %s)", exp.Type(), exp.orig, exp.to)
}

// IntZeroExt is a constant expression which zero extends an integer constant to
// a larger or equally sized integer type.
//
// Examples:
//    zext(i1 true to i5)   ; yields i5:1
//
// References:
//    http://llvm.org/docs/LangRef.html#constant-expressions
type IntZeroExt struct {
	// Original integer constant.
	orig *Int
	// New integer type.
	to *types.Int
}

// NewIntZeroExt returns a constant expression which zero extends the integer
// constant orig to a larger or equally sized integer type.
func NewIntZeroExt(orig Constant, to types.Type) (*IntZeroExt, error) {
	// Verify type of original integer constant.
	exp := new(IntZeroExt)
	var ok bool
	exp.orig, ok = orig.(*Int)
	if !ok {
		return nil, fmt.Errorf("invalid integer zero extension; expected integer constant for orig, got %q", orig.Type())
	}

	// Verify target type.
	exp.to, ok = to.(*types.Int)
	if !ok {
		return nil, fmt.Errorf("invalid integer zero extension; expected integer target type, got %q", to)
	}
	newSize, origSize := exp.to.Size(), exp.orig.typ.Size()
	if newSize < origSize {
		return nil, fmt.Errorf("invalid integer zero extension; target size (%d) smaller than original size (%d)", newSize, origSize)
	}

	return exp, nil
}

// Type returns the type of the value.
func (exp *IntZeroExt) Type() types.Type {
	return exp.to
}

// Calc calculates and returns a constant which is equivalent to the constant
// expression.
func (exp *IntZeroExt) Calc() Constant {
	panic("IntZeroExt.Calc: not yet implemented.")
}

// String returns a string representation of the integer zero extension
// expression. The expression string representation is preceded by the type of
// the constant, e.g.
//
//    i5 zext(i1 true to i5)
func (exp *IntZeroExt) String() string {
	return fmt.Sprintf("%s zext(%s to %s)", exp.Type(), exp.orig, exp.to)
}

// IntSignExt is a constant expression which sign extends an integer constant to
// a larger or equally sized integer type.
//
// Examples:
//    sext(i1 true to i5)   ; yields i5:31
//
// References:
//    http://llvm.org/docs/LangRef.html#constant-expressions
type IntSignExt struct {
	// Original integer constant.
	orig *Int
	// New integer type.
	to *types.Int
}

// NewIntSignExt returns a constant expression which sign extends the integer
// constant orig to a larger or equally sized integer type.
func NewIntSignExt(orig Constant, to types.Type) (*IntSignExt, error) {
	// Verify type of original integer constant.
	exp := new(IntSignExt)
	var ok bool
	exp.orig, ok = orig.(*Int)
	if !ok {
		return nil, fmt.Errorf("invalid integer sign extension; expected integer constant for orig, got %q", orig.Type())
	}

	// Verify target type.
	exp.to, ok = to.(*types.Int)
	if !ok {
		return nil, fmt.Errorf("invalid integer sign extension; expected integer target type, got %q", to)
	}
	newSize, origSize := exp.to.Size(), exp.orig.typ.Size()
	if newSize < origSize {
		return nil, fmt.Errorf("invalid integer sign extension; target size (%d) smaller than original size (%d)", newSize, origSize)
	}

	return exp, nil
}

// Type returns the type of the value.
func (exp *IntSignExt) Type() types.Type {
	return exp.to
}

// Calc calculates and returns a constant which is equivalent to the constant
// expression.
func (exp *IntSignExt) Calc() Constant {
	panic("IntSignExt.Calc: not yet implemented.")
}

// String returns a string representation of the integer sign extension
// expression. The expression string representation is preceded by the type of
// the constant, e.g.
//
//    i5 sext(i1 true to i5)
func (exp *IntSignExt) String() string {
	return fmt.Sprintf("%s sext(%s to %s)", exp.Type(), exp.orig, exp.to)
}

// FloatTrunc is a constant expression which truncates a floating point constant
// to a smaller floating point type or one of the same kind.
//
// Examples:
//    fptrunc(double 4.0 to float)   ; yields float:4.0
//
// References:
//    http://llvm.org/docs/LangRef.html#constant-expressions
type FloatTrunc struct {
	// Original floating point constant.
	orig *Float
	// New floating point type.
	to *types.Float
}

// NewFloatTrunc returns a constant expression which truncates the floating
// point constant orig to a smaller floating point type or one of the same kind.
func NewFloatTrunc(orig Constant, to types.Type) (*FloatTrunc, error) {
	// Verify type of original floating point constant.
	exp := new(FloatTrunc)
	var ok bool
	exp.orig, ok = orig.(*Float)
	if !ok {
		return nil, fmt.Errorf("invalid floating point truncation; expected floating point constant for orig, got %q", orig.Type())
	}

	// Verify target type.
	exp.to, ok = to.(*types.Float)
	if !ok {
		return nil, fmt.Errorf("invalid floating point truncation; expected floating point target type, got %q", to)
	}
	newSize, origSize := exp.to.Size(), exp.orig.typ.Size()
	newKind, origKind := exp.to.Kind(), exp.orig.typ.Kind()
	if newSize > origSize {
		return nil, fmt.Errorf("invalid floating point truncation; target size (%d) larger than original size (%d)", newSize, origSize)
	} else if newSize == origSize && newKind != origKind {
		return nil, fmt.Errorf("invalid floating point truncation; cannot convert from %q to %q", exp.orig.typ, exp.to)
	}

	return exp, nil
}

// Type returns the type of the value.
func (exp *FloatTrunc) Type() types.Type {
	return exp.to
}

// Calc calculates and returns a constant which is equivalent to the constant
// expression.
func (exp *FloatTrunc) Calc() Constant {
	panic("FloatTrunc.Calc: not yet implemented.")
}

// String returns a string representation of the floating point truncation
// expression. The expression string representation is preceded by the type of
// the constant, e.g.
//
//    float fptrunc(double 4.0 to float)
func (exp *FloatTrunc) String() string {
	return fmt.Sprintf("%s fptrunc(%s to %s)", exp.Type(), exp.orig, exp.to)
}

// FloatExt is a constant expression which extends a floating point constant to
// a larger floating point type or one of the same kind.
//
// Examples:
//    fpext(float 4.0 to double)   ; yields double:4.0
//
// References:
//    http://llvm.org/docs/LangRef.html#constant-expressions
type FloatExt struct {
	// Original floating point constant.
	orig *Float
	// New floating point type.
	to *types.Float
}

// NewFloatExt returns a constant expression which extends the floating point
// constant orig to a larger floating point type or one of the same kind.
func NewFloatExt(orig Constant, to types.Type) (*FloatExt, error) {
	// Verify type of original floating point constant.
	exp := new(FloatExt)
	var ok bool
	exp.orig, ok = orig.(*Float)
	if !ok {
		return nil, fmt.Errorf("invalid floating point extension; expected floating point constant for orig, got %q", orig.Type())
	}

	// Verify target type.
	exp.to, ok = to.(*types.Float)
	if !ok {
		return nil, fmt.Errorf("invalid floating point extension; expected floating point target type, got %q", to)
	}
	newSize, origSize := exp.to.Size(), exp.orig.typ.Size()
	newKind, origKind := exp.to.Kind(), exp.orig.typ.Kind()
	if newSize < origSize {
		return nil, fmt.Errorf("invalid floating point extension; target size (%d) smaller than original size (%d)", newSize, origSize)
	} else if newSize == origSize && newKind != origKind {
		return nil, fmt.Errorf("invalid floating point extension; cannot convert from %q to %q", exp.orig.typ, exp.to)
	}

	return exp, nil
}

// Type returns the type of the value.
func (exp *FloatExt) Type() types.Type {
	return exp.to
}

// Calc calculates and returns a constant which is equivalent to the constant
// expression.
func (exp *FloatExt) Calc() Constant {
	panic("FloatExt.Calc: not yet implemented.")
}

// String returns a string representation of the floating point extension
// expression. The expression string representation is preceded by the type of
// the constant, e.g.
//
//    double fpext(float 4.0 to double)
func (exp *FloatExt) String() string {
	return fmt.Sprintf("%s fpext(%s to %s)", exp.Type(), exp.orig, exp.to)
}

// FloatToUint is a constant expression which converts a floating point constant
// (or constant vector) to the corresponding unsigned integer constant (or
// constant vector).
//
// Examples:
//    fptoui(float 4.0 to i32)                       ; yields i32:4
//    fptoui(<1 x float> <float 3.0> to <1 x i32>)   ; yields <1 x i32>:<i32 3>
//
// References:
//    http://llvm.org/docs/LangRef.html#constant-expressions
type FloatToUint struct {
	// Original floating point value (or vector).
	orig value.Value
	// New integer type (or vector).
	to types.Type
}

// NewFloatToUint returns a constant expression which converts the floating
// point constant (or constant vector) orig to the corresponding unsigned
// integer constant (or constant vector).
func NewFloatToUint(orig Constant, to types.Type) (*FloatToUint, error) {
	// Verify type of original floating point constant (or constant vector).
	if !types.IsFloats(orig.Type()) {
		return nil, fmt.Errorf("invalid floating point to unsigned integer conversion; expected floating point constant (or constant vector) for orig, got %q", orig.Type())
	}

	// Verify target type.
	if !types.IsInts(to) {
		return nil, fmt.Errorf("invalid floating point to unsigned integer conversion; expected integer (or integer vector) target type, got %q", to)
	}

	// Verify that both are either basic types or vectors of the same length.
	if !types.SameLength(orig.Type(), to) {
		return nil, fmt.Errorf("invalid floating point to unsigned integer conversion; cannot convert from %q to %q", orig.Type(), to)
	}

	return &FloatToUint{orig: orig, to: to}, nil
}

// Type returns the type of the value.
func (exp *FloatToUint) Type() types.Type {
	return exp.to
}

// Calc calculates and returns a constant which is equivalent to the constant
// expression.
func (exp *FloatToUint) Calc() Constant {
	panic("FloatToUint.Calc: not yet implemented.")
}

// String returns a string representation of the constant expression which
// converts a floating point constant (or constant vector) to the corresponding
// unsigned integer constant (or constant vector). The expression string
// representation is preceded by the type of the constant, e.g.
//
//    i32 fptoui(float 4.0 to i32)
//    <2 x i32> fptoui(<2 x float> <float 3.0, float 4.0> to <2 x i32>)
func (exp *FloatToUint) String() string {
	return fmt.Sprintf("%s fptoui(%s to %s)", exp.Type(), exp.orig, exp.to)
}

// FloatToInt is a constant expression which converts a floating point constant
// (or constant vector) to the corresponding signed integer constant (or
// constant vector).
//
// Examples:
//    fptosi(float -4.0 to i32)                       ; yields i32:-4
//    fptosi(<1 x float> <float -3.0> to <1 x i32>)   ; yields <1 x i32>:<i32 -3>
//
// References:
//    http://llvm.org/docs/LangRef.html#constant-expressions
type FloatToInt struct {
	// Original floating point value (or vector).
	orig value.Value
	// New type (or vector).
	to types.Type
}

// NewFloatToInt returns a constant expression which converts the floating point
// constant (or constant vector) orig to the corresponding signed integer
// constant (or constant vector).
func NewFloatToInt(orig Constant, to types.Type) (*FloatToInt, error) {
	// Verify type of original floating point constant (or constant vector).
	if !types.IsFloats(orig.Type()) {
		return nil, fmt.Errorf("invalid floating point to signed integer conversion; expected floating point constant (or constant vector) for orig, got %q", orig.Type())
	}

	// Verify target type.
	if !types.IsInts(to) {
		return nil, fmt.Errorf("invalid floating point to signed integer conversion; expected integer (or integer vector) target type, got %q", to)
	}

	// Verify that both are either basic types or vectors of the same length.
	if !types.SameLength(orig.Type(), to) {
		return nil, fmt.Errorf("invalid floating point to signed integer conversion; cannot convert from %q to %q", orig.Type(), to)
	}

	return &FloatToInt{orig: orig, to: to}, nil
}

// Type returns the type of the value.
func (exp *FloatToInt) Type() types.Type {
	return exp.to
}

// Calc calculates and returns a constant which is equivalent to the constant
// expression.
func (exp *FloatToInt) Calc() Constant {
	panic("FloatToInt.Calc: not yet implemented.")
}

// String returns a string representation of the constant expression which
// converts a floating point constant (or constant vector) to the corresponding
// signed integer constant (or constant vector). The expression string
// representation is preceded by the type of the constant, e.g.
//
//    i32 fptosi(float -4.0 to i32)
//    <2 x i32> fptosi(<2 x float> <float -3.0, float 4.0> to <2 x i32>)
func (exp *FloatToInt) String() string {
	return fmt.Sprintf("%s fptosi(%s to %s)", exp.Type(), exp.orig, exp.to)
}

// UintToFloat is a constant expression which converts an unsigned integer
// constant (or constant vector) to the corresponding floating point constant
// (or constant vector).
//
// Examples:
//    uitofp(i32 4 to float)                     ; yields float:4.0
//    uitofp(<1 x i32> <i32 3> to <1 x float>)   ; yields <1 x float>:<float 3.0>
//
// References:
//    http://llvm.org/docs/LangRef.html#constant-expressions
type UintToFloat struct {
	// Original unsigned integer value (or vector).
	orig value.Value
	// New floating point type (or vector).
	to types.Type
}

// NewUintToFloat returns a constant expression which converts the unsigned
// integer constant (or constant vector) orig to the corresponding floating
// point constant (or constant vector).
func NewUintToFloat(orig Constant, to types.Type) (*UintToFloat, error) {
	// Verify type of original integer constant (or constant vector).
	if !types.IsInts(orig.Type()) {
		return nil, fmt.Errorf("invalid unsigned integer to floating point conversion; expected integer constant (or constant vector) for orig, got %q", orig.Type())
	}

	// Verify target type.
	if !types.IsFloats(to) {
		return nil, fmt.Errorf("invalid unsigned integer to floating point conversion; expected floating point (or floating point vector) target type, got %q", to)
	}

	// Verify that both are either basic types or vectors of the same length.
	if !types.SameLength(orig.Type(), to) {
		return nil, fmt.Errorf("invalid unsigned integer to floating point conversion; cannot convert from %q to %q", orig.Type(), to)
	}

	return &UintToFloat{orig: orig, to: to}, nil
}

// Type returns the type of the value.
func (exp *UintToFloat) Type() types.Type {
	return exp.to
}

// Calc calculates and returns a constant which is equivalent to the constant
// expression.
func (exp *UintToFloat) Calc() Constant {
	panic("UintToFloat.Calc: not yet implemented.")
}

// String returns a string representation of the constant expression which
// converts an unsigned integer constant (or constant vector) to the
// corresponding floating point constant (or constant vector). The expression
// string representation is preceded by the type of the constant, e.g.
//
//    float uitofp(i32 4 to float)
//    <2 x float> uitofp(<2 x i32> <i32 3, i32 42> to <2 x float>)
func (exp *UintToFloat) String() string {
	return fmt.Sprintf("%s uitofp(%s to %s)", exp.Type(), exp.orig, exp.to)
}

// IntToFloat is a constant expression which converts a signed integer constant
// (or constant vector) to the corresponding floating point constant (or
// constant vector).
//
// Examples:
//    sitofp(i32 -4 to float)                     ; yields float:-4.0
//    sitofp(<1 x i32> <i32 -3> to <1 x float>)   ; yields <1 x float>:<float -3.0>
//
// References:
//    http://llvm.org/docs/LangRef.html#constant-expressions
type IntToFloat struct {
	// Original signed integer value (or vector).
	orig value.Value
	// New floating point type (or vector).
	to types.Type
}

// NewIntToFloat returns a constant expression which converts the signed integer
// constant (or constant vector) orig to the corresponding floating point
// constant (or constant vector).
func NewIntToFloat(orig Constant, to types.Type) (*IntToFloat, error) {
	// Verify type of original integer constant (or constant vector).
	if !types.IsInts(orig.Type()) {
		return nil, fmt.Errorf("invalid signed integer to floating point conversion; expected integer constant (or constant vector) for orig, got %q", orig.Type())
	}

	// Verify target type.
	if !types.IsFloats(to) {
		return nil, fmt.Errorf("invalid signed integer to floating point conversion; expected floating point (or floating point vector) target type, got %q", to)
	}

	// Verify that both are either basic types or vectors of the same length.
	if !types.SameLength(orig.Type(), to) {
		return nil, fmt.Errorf("invalid signed integer to floating point conversion; cannot convert from %q to %q", orig.Type(), to)
	}

	return &IntToFloat{orig: orig, to: to}, nil
}

// Type returns the type of the value.
func (exp *IntToFloat) Type() types.Type {
	return exp.to
}

// Calc calculates and returns a constant which is equivalent to the constant
// expression.
func (exp *IntToFloat) Calc() Constant {
	panic("IntToFloat.Calc: not yet implemented.")
}

// String returns a string representation of the constant expression which
// converts a signed integer constant (or constant vector) to the corresponding
// floating point constant (or constant vector). The expression string
// representation is preceded by the type of the constant, e.g.
//
//    float sitofp(i32 -4 to float)
//    <2 x float> sitofp(<2 x i32> <i32 -3, i32 15> to <2 x float>)
func (exp *IntToFloat) String() string {
	return fmt.Sprintf("%s sitofp(%s to %s)", exp.Type(), exp.orig, exp.to)
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

// isConst ensures that only constant values can be assigned to the Constant
// interface.
func (*IntTrunc) isConst()    {}
func (*IntZeroExt) isConst()  {}
func (*IntSignExt) isConst()  {}
func (*FloatTrunc) isConst()  {}
func (*FloatExt) isConst()    {}
func (*FloatToUint) isConst() {}
func (*FloatToInt) isConst()  {}
func (*UintToFloat) isConst() {}
func (*IntToFloat) isConst()  {}
