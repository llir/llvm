// References:
//    http://llvm.org/docs/LangRef.html#constant-expressions

package constant

import (
	"bytes"
	"fmt"

	"github.com/llir/llvm/ir/types"
	"github.com/llir/llvm/ir/value"
	"github.com/mewkiz/pkg/errutil"
)

// TODO: Complete the list of expression implementations.

// An Expr represents a constant expression.
//
// Expr is one of the following concrete types:
//    *constant.Trunc
//    *constant.ZExt
//    *constant.SExt
//    *constant.FPTrunc
//    *constant.FPExt
//    *constant.FPToUI
//    *constant.FPToSI
//    *constant.UIToFP
//    *constant.SIToFP
//    *constant.GetElementPtr
//
// References:
//    http://llvm.org/docs/LangRef.html#constant-expressions
type Expr interface {
	Constant
	// Calc calculates and returns a constant which is equivalent to the constant
	// expression.
	Calc() Constant
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
	// Original integer constant.
	orig *Int
	// New integer type.
	to *types.Int
}

// NewTrunc returns a constant expression which truncates the integer constant
// orig to a smaller or equally sized integer type.
func NewTrunc(orig Constant, to types.Type) (*Trunc, error) {
	// Verify type of original integer constant.
	exp := new(Trunc)
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
func (exp *Trunc) Type() types.Type {
	return exp.to
}

// Calc calculates and returns a constant which is equivalent to the constant
// expression.
func (exp *Trunc) Calc() Constant {
	panic("Trunc.Calc: not yet implemented.")
}

// String returns a string representation of the integer truncation expression;
// e.g.
//
//    trunc(i32 15 to i3)
func (exp *Trunc) String() string {
	return fmt.Sprintf("trunc (%s %s to %s)", exp.orig.Type(), exp.orig, exp.to)
}

// ZExt is a constant expression which zero extends an integer constant to a
// larger or equally sized integer type.
//
// Examples:
//    zext(i1 true to i5)   ; yields i5:1
//
// References:
//    http://llvm.org/docs/LangRef.html#constant-expressions
type ZExt struct {
	// Original integer constant.
	orig *Int
	// New integer type.
	to *types.Int
}

// NewZExt returns a constant expression which zero extends the integer constant
// orig to a larger or equally sized integer type.
func NewZExt(orig Constant, to types.Type) (*ZExt, error) {
	// Verify type of original integer constant.
	exp := new(ZExt)
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
func (exp *ZExt) Type() types.Type {
	return exp.to
}

// Calc calculates and returns a constant which is equivalent to the constant
// expression.
func (exp *ZExt) Calc() Constant {
	panic("ZExt.Calc: not yet implemented.")
}

// String returns a string representation of the integer zero extension
// expression; e.g.
//
//    zext(i1 true to i5)
func (exp *ZExt) String() string {
	return fmt.Sprintf("zext (%s %s to %s)", exp.orig.Type(), exp.orig, exp.to)
}

// SExt is a constant expression which sign extends an integer constant to a
// larger or equally sized integer type.
//
// Examples:
//    sext(i1 true to i5)   ; yields i5:31
//
// References:
//    http://llvm.org/docs/LangRef.html#constant-expressions
type SExt struct {
	// Original integer constant.
	orig *Int
	// New integer type.
	to *types.Int
}

// NewSExt returns a constant expression which sign extends the integer constant
// orig to a larger or equally sized integer type.
func NewSExt(orig Constant, to types.Type) (*SExt, error) {
	// Verify type of original integer constant.
	exp := new(SExt)
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
func (exp *SExt) Type() types.Type {
	return exp.to
}

// Calc calculates and returns a constant which is equivalent to the constant
// expression.
func (exp *SExt) Calc() Constant {
	panic("SExt.Calc: not yet implemented.")
}

// String returns a string representation of the integer sign extension
// expression; e.g.
//
//    sext(i1 true to i5)
func (exp *SExt) String() string {
	return fmt.Sprintf("sext (%s %s to %s)", exp.orig.Type(), exp.orig, exp.to)
}

// FPTrunc is a constant expression which truncates a floating point constant to
// a smaller floating point type or one of the same kind.
//
// Examples:
//    fptrunc(double 4.0 to float)   ; yields float:4.0
//
// References:
//    http://llvm.org/docs/LangRef.html#constant-expressions
type FPTrunc struct {
	// Original floating point constant.
	orig *Float
	// New floating point type.
	to *types.Float
}

// NewFPTrunc returns a constant expression which truncates the floating point
// constant orig to a smaller floating point type or one of the same kind.
func NewFPTrunc(orig Constant, to types.Type) (*FPTrunc, error) {
	// Verify type of original floating point constant.
	exp := new(FPTrunc)
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
func (exp *FPTrunc) Type() types.Type {
	return exp.to
}

// Calc calculates and returns a constant which is equivalent to the constant
// expression.
func (exp *FPTrunc) Calc() Constant {
	panic("FPTrunc.Calc: not yet implemented.")
}

// String returns a string representation of the floating point truncation
// expression; e.g.
//
//    float fptrunc(double 4.0 to float)
func (exp *FPTrunc) String() string {
	return fmt.Sprintf("fptrunc (%s %s to %s)", exp.orig.Type(), exp.orig, exp.to)
}

// FPExt is a constant expression which extends a floating point constant to a
// larger floating point type or one of the same kind.
//
// Examples:
//    fpext(float 4.0 to double)   ; yields double:4.0
//
// References:
//    http://llvm.org/docs/LangRef.html#constant-expressions
type FPExt struct {
	// Original floating point constant.
	orig *Float
	// New floating point type.
	to *types.Float
}

// NewFPExt returns a constant expression which extends the floating point
// constant orig to a larger floating point type or one of the same kind.
func NewFPExt(orig Constant, to types.Type) (*FPExt, error) {
	// Verify type of original floating point constant.
	exp := new(FPExt)
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
func (exp *FPExt) Type() types.Type {
	return exp.to
}

// Calc calculates and returns a constant which is equivalent to the constant
// expression.
func (exp *FPExt) Calc() Constant {
	panic("FPExt.Calc: not yet implemented.")
}

// String returns a string representation of the floating point extension
// expression; e.g.
//
//    fpext(float 4.0 to double)
func (exp *FPExt) String() string {
	return fmt.Sprintf("fpext (%s %s to %s)", exp.orig.Type(), exp.orig, exp.to)
}

// FPToUI is a constant expression which converts a floating point constant (or
// constant vector) to the corresponding unsigned integer constant (or constant
// vector).
//
// Examples:
//    fptoui(float 4.0 to i32)                       ; yields i32:4
//    fptoui(<1 x float> <float 3.0> to <1 x i32>)   ; yields <1 x i32>:<i32 3>
//
// References:
//    http://llvm.org/docs/LangRef.html#constant-expressions
type FPToUI struct {
	// Original floating point value (or vector).
	orig value.Value
	// New integer type (or vector).
	to types.Type
}

// NewFPToUI returns a constant expression which converts the floating point
// constant (or constant vector) orig to the corresponding unsigned integer
// constant (or constant vector).
func NewFPToUI(orig Constant, to types.Type) (*FPToUI, error) {
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

	return &FPToUI{orig: orig, to: to}, nil
}

// Type returns the type of the value.
func (exp *FPToUI) Type() types.Type {
	return exp.to
}

// Calc calculates and returns a constant which is equivalent to the constant
// expression.
func (exp *FPToUI) Calc() Constant {
	panic("FPToUI.Calc: not yet implemented.")
}

// String returns a string representation of the constant expression which
// converts a floating point constant (or constant vector) to the corresponding
// unsigned integer constant (or constant vector); e.g.
//
//    fptoui(float 4.0 to i32)
//    fptoui(<2 x float> <float 3.0, float 4.0> to <2 x i32>)
func (exp *FPToUI) String() string {
	return fmt.Sprintf("fptoui (%s %s to %s)", exp.orig.Type(), exp.orig, exp.to)
}

// FPToSI is a constant expression which converts a floating point constant (or
// constant vector) to the corresponding signed integer constant (or constant
// vector).
//
// Examples:
//    fptosi(float -4.0 to i32)                       ; yields i32:-4
//    fptosi(<1 x float> <float -3.0> to <1 x i32>)   ; yields <1 x i32>:<i32 -3>
//
// References:
//    http://llvm.org/docs/LangRef.html#constant-expressions
type FPToSI struct {
	// Original floating point value (or vector).
	orig value.Value
	// New type (or vector).
	to types.Type
}

// NewFPToSI returns a constant expression which converts the floating point
// constant (or constant vector) orig to the corresponding signed integer
// constant (or constant vector).
func NewFPToSI(orig Constant, to types.Type) (*FPToSI, error) {
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

	return &FPToSI{orig: orig, to: to}, nil
}

// Type returns the type of the value.
func (exp *FPToSI) Type() types.Type {
	return exp.to
}

// Calc calculates and returns a constant which is equivalent to the constant
// expression.
func (exp *FPToSI) Calc() Constant {
	panic("FPToSI.Calc: not yet implemented.")
}

// String returns a string representation of the constant expression which
// converts a floating point constant (or constant vector) to the corresponding
// signed integer constant (or constant vector); e.g.
//
//    fptosi(float -4.0 to i32)
//    fptosi(<2 x float> <float -3.0, float 4.0> to <2 x i32>)
func (exp *FPToSI) String() string {
	return fmt.Sprintf("fptosi (%s %s to %s)", exp.orig.Type(), exp.orig, exp.to)
}

// UIToFP is a constant expression which converts an unsigned integer constant
// (or constant vector) to the corresponding floating point constant (or
// constant vector).
//
// Examples:
//    uitofp(i32 4 to float)                     ; yields float:4.0
//    uitofp(<1 x i32> <i32 3> to <1 x float>)   ; yields <1 x float>:<float 3.0>
//
// References:
//    http://llvm.org/docs/LangRef.html#constant-expressions
type UIToFP struct {
	// Original unsigned integer value (or vector).
	orig value.Value
	// New floating point type (or vector).
	to types.Type
}

// NewUIToFP returns a constant expression which converts the unsigned integer
// constant (or constant vector) orig to the corresponding floating point
// constant (or constant vector).
func NewUIToFP(orig Constant, to types.Type) (*UIToFP, error) {
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

	return &UIToFP{orig: orig, to: to}, nil
}

// Type returns the type of the value.
func (exp *UIToFP) Type() types.Type {
	return exp.to
}

// Calc calculates and returns a constant which is equivalent to the constant
// expression.
func (exp *UIToFP) Calc() Constant {
	panic("UIToFP.Calc: not yet implemented.")
}

// String returns a string representation of the constant expression which
// converts an unsigned integer constant (or constant vector) to the
// corresponding floating point constant (or constant vector); e.g.
//
//    uitofp(i32 4 to float)
//    uitofp(<2 x i32> <i32 3, i32 42> to <2 x float>)
func (exp *UIToFP) String() string {
	return fmt.Sprintf("uitofp (%s %s to %s)", exp.orig.Type(), exp.orig, exp.to)
}

// SIToFP is a constant expression which converts a signed integer constant (or
// constant vector) to the corresponding floating point constant (or constant
// vector).
//
// Examples:
//    sitofp(i32 -4 to float)                     ; yields float:-4.0
//    sitofp(<1 x i32> <i32 -3> to <1 x float>)   ; yields <1 x float>:<float -3.0>
//
// References:
//    http://llvm.org/docs/LangRef.html#constant-expressions
type SIToFP struct {
	// Original signed integer value (or vector).
	orig value.Value
	// New floating point type (or vector).
	to types.Type
}

// NewSIToFP returns a constant expression which converts the signed integer
// constant (or constant vector) orig to the corresponding floating point
// constant (or constant vector).
func NewSIToFP(orig Constant, to types.Type) (*SIToFP, error) {
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

	return &SIToFP{orig: orig, to: to}, nil
}

// Type returns the type of the value.
func (exp *SIToFP) Type() types.Type {
	return exp.to
}

// Calc calculates and returns a constant which is equivalent to the constant
// expression.
func (exp *SIToFP) Calc() Constant {
	panic("SIToFP.Calc: not yet implemented.")
}

// String returns a string representation of the constant expression which
// converts a signed integer constant (or constant vector) to the corresponding
// floating point constant (or constant vector); e.g.
//
//    sitofp(i32 -4 to float)
//    sitofp(<2 x i32> <i32 -3, i32 15> to <2 x float>)
func (exp *SIToFP) String() string {
	return fmt.Sprintf("sitofp (%s %s to %s)", exp.orig.Type(), exp.orig, exp.to)
}

// TODO: Add support for the following constant expressions:
//    - ptrtoint
//    - inttoptr
//    - bitcast
//    - addrspacecast

// GetElementPtr represents a getelementptr expression.
type GetElementPtr struct {
	// Value type.
	typ types.Type
	// Element type.
	elem types.Type
	// Memory address of the element.
	addr value.Value
	// Element indices.
	indices []value.Value
}

// NewGetElementPtr returns a new getelementptr expression based on the given
// type, element type, address and element indices.
func NewGetElementPtr(typ, elem types.Type, addr value.Value, indices []value.Value) (*GetElementPtr, error) {
	// Sanity checks.
	switch addrType := addr.Type().(type) {
	case *types.Pointer:
		if !types.Equal(elem, addrType.Elem()) {
			return nil, errutil.Newf("type mismatch between %v and %v", elem, addrType.Elem())
		}
	default:
		return nil, errutil.Newf("invalid pointer type; expected *types.Pointer, got %T", addrType)
	}
	return &GetElementPtr{typ: typ, elem: elem, addr: addr, indices: indices}, nil
}

// Type returns the type of the value produced by the expression.
func (exp *GetElementPtr) Type() types.Type {
	return exp.typ
}

// Calc calculates and returns a constant which is equivalent to the constant
// expression.
func (exp *GetElementPtr) Calc() Constant {
	panic("GetElementPtr.Calc: not yet implemented.")
}

// String returns the string representation of the expression.
func (exp *GetElementPtr) String() string {
	indicesBuf := new(bytes.Buffer)
	for _, index := range exp.indices {
		fmt.Fprintf(indicesBuf, ", %s %s", index.Type(), index)
	}
	return fmt.Sprintf("getelementptr (%s, %s %s%s)", exp.elem, exp.addr.Type(), exp.addr, indicesBuf)
}

// TODO: Add support for the following constant expressions:
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
func (*Trunc) isConst()         {}
func (*ZExt) isConst()          {}
func (*SExt) isConst()          {}
func (*FPTrunc) isConst()       {}
func (*FPExt) isConst()         {}
func (*FPToUI) isConst()        {}
func (*FPToSI) isConst()        {}
func (*UIToFP) isConst()        {}
func (*SIToFP) isConst()        {}
func (*GetElementPtr) isConst() {}
