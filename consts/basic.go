package consts

import (
	"fmt"
	"strconv"

	"github.com/mewkiz/pkg/errutil"
	"github.com/mewlang/llvm/types"
)

// Int represents an integer constant.
//
// Examples:
//    42, -37, true, false, [us]0x[0-9A-Fa-f]+
//
// References:
//    http://llvm.org/docs/LangRef.html#simple-constants
type Int struct {
	typ *types.Int
	x   int64
}

// NewInt returns an integer constant based on the given integer type and string
// representation.
func NewInt(typ types.Type, s string) (*Int, error) {
	v := new(Int)
	var ok bool
	v.typ, ok = typ.(*types.Int)
	if !ok {
		return nil, errutil.Newf("invalid type %q; expected integer type", typ)
	}
	size := v.typ.Size()
	if size > 64 {
		// TODO: Add support for large integer constants (e.g. above 64-bit).
		err := fmt.Sprintf("not yet implemented; support for %q integer constants (e.g. bit width above 64).", typ)
		panic(err)
	}
	// TODO: Implement support for the following representations:
	//    true, false and [us]0x[0-9A-Fa-f]+
	var err error
	v.x, err = strconv.ParseInt(s, 10, size)
	if err != nil {
		return nil, errutil.Newf("invalid integer constant %q; %v", s, err)
	}
	return v, nil
}

// Trunc truncates the integer to a smaller or equally sized integer type.
func (v *Int) Trunc(to types.Type) (*Trunc, error) {
	expr := &Trunc{v: v}
	var ok bool
	expr.to, ok = to.(*types.Int)
	if !ok {
		return nil, errutil.Newf("invalid type %q; expected integer type", to)
	}
	if v.typ.Size() < expr.to.Size() {
		return nil, errutil.Newf("integer truncation from %q to %q; target type must be smaller or equally sized", v.typ, to)
	}
	return expr, nil
}

// ZExt zero extends the integer to a larger or equally sized integer type.
func (v *Int) ZExt(to types.Type) (*ZExt, error) {
	expr := &ZExt{v: v}
	var ok bool
	expr.to, ok = to.(*types.Int)
	if !ok {
		return nil, errutil.Newf("invalid type %q; expected integer type", to)
	}
	if v.typ.Size() > expr.to.Size() {
		return nil, errutil.Newf("zero extension from %q to %q; target type must be larger or equally sized", v.typ, to)
	}
	return expr, nil
}

// SExt sign extends the integer to a larger or equally sized integer type.
func (v *Int) SExt(to types.Type) (*SExt, error) {
	expr := &SExt{v: v}
	var ok bool
	expr.to, ok = to.(*types.Int)
	if !ok {
		return nil, errutil.Newf("invalid type %q; expected integer type", to)
	}
	if v.typ.Size() > expr.to.Size() {
		return nil, errutil.Newf("sign extension from %q to %q; target type must be larger or equally sized", v.typ, to)
	}
	return expr, nil
}

func (v *Int) String() string {
	// TODO: Add special case for true and false when the type is i1?
	return strconv.FormatInt(v.x, 10)
}

// Float represents a floating point constant.
//
// Examples:
//    123.45, 1.2345e+2, 0x[KLMH]?[0-9A-Fa-f]+
//
// References:
//    http://llvm.org/docs/LangRef.html#simple-constants
type Float struct {
	typ *types.Float
	x   float64
}

// NewFloat returns a floating point constant based on the given floating point
// type and string representation.
func NewFloat(typ types.Type, s string) (*Float, error) {
	v := new(Float)
	var ok bool
	v.typ, ok = typ.(*types.Float)
	if !ok {
		return nil, errutil.Newf("invalid type %q; expected floating point type", typ)
	}
	var nbits int
	switch v.typ.Kind() {
	case types.Float32:
		nbits = 32
	case types.Float64:
		nbits = 64
	default:
		// TODO: Add support for half, fp128, x86_fp80 and ppc_fp128.
		err := fmt.Sprintf("not yet implemented; support for %q floating point constants.\n", typ)
		panic(err)
	case types.Float16:
	}
	// TODO: Implement support for the following representation:
	//    0x[KLMH]?[0-9A-Fa-f]+
	var err error
	v.x, err = strconv.ParseFloat(s, nbits)
	if err != nil {
		return nil, errutil.Newf("invalid floating point constant %q; %v", s, err)
	}
	return v, nil
}

// TODO: Check if global names are used for anything except functions and global
// variables. If so, be more specific about @foo in the example below by
// providing a comment.

// Pointer represents a pointer constant.
//
// Examples:
//    null, @foo
//
// References:
//    http://llvm.org/docs/LangRef.html#simple-constants
type Pointer struct {
}

// TODO: Figure out how to represent pointer constants, and implement NewPointer
// afterwards.
