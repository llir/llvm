package consts

import (
	"fmt"
	"strconv"

	"github.com/mewkiz/pkg/errutil"
	"github.com/mewlang/llvm/types"
	"github.com/mewlang/llvm/values"
)

// TODO: Track the upstream removal of HexIntConstant (ref: discussion with
// Sean on llvm-dev).

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
	// Verify integer type.
	v := new(Int)
	var ok bool
	v.typ, ok = typ.(*types.Int)
	if !ok {
		return nil, errutil.Newf("invalid type %q for integer constant", typ)
	}
	size := v.typ.Size()
	if size > 64 {
		// TODO: Add support for large integer constants (e.g. above 64-bits).
		err := fmt.Sprintf("not yet implemented; support for %q integer constants (e.g. above 64-bits)", typ)
		panic(err)
	}

	// Parse boolean constant.
	if size == 1 {
		switch s {
		case "1", "true":
			v.x = 1
		case "0", "false":
			v.x = 0
		}
		return v, nil
	} else if s == "true" || s == "false" {
		return nil, errutil.Newf("integer constant %q type mismatch; expected i1, got %v", s, typ)
	}

	// TODO: Implement support for the HexIntConstant representation:
	//    [us]0x[0-9A-Fa-f]+

	// Parse integer constant.
	var err error
	v.x, err = strconv.ParseInt(s, 10, size)
	if err != nil {
		return nil, errutil.Newf("invalid integer constant %q; %v", s, err)
	}

	return v, nil
}

// Type returns the type of the value.
func (v *Int) Type() types.Type {
	return v.typ
}

// UseList returns a list of all values which uses the value.
func (v *Int) UseList() []values.Value {
	panic("not yet implemented.")
}

// ReplaceAll replaces all uses of the value with new.
func (v *Int) ReplaceAll(new values.Value) error {
	panic("not yet implemented.")
}

// String returns the string representation of v, either as a signed integer
// (e.g. 42, -13) or as a boolean (e.g. true, false) depending on the type.
func (v *Int) String() string {
	if v.typ.Size() == 1 {
		if v.x == 1 {
			return "true"
		}
		return "false"
	}
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
	// Verify floating point type.
	v := new(Float)
	var ok bool
	v.typ, ok = typ.(*types.Float)
	if !ok {
		return nil, errutil.Newf("invalid type %q for floating point constant", typ)
	}
	var size int
	switch v.typ.Kind() {
	case types.Float32:
		size = 32
	case types.Float64:
		size = 64
	default:
		// TODO: Add support for half, fp128, x86_fp80 and ppc_fp128.
		err := fmt.Sprintf("not yet implemented; support for %q floating point constants", typ)
		panic(err)
	}

	// TODO: Implement support for the following representation:
	//    0x[KLMH]?[0-9A-Fa-f]+

	// TODO: Verify that the input string can be precisely represented by the
	// floating point value. For instance, a precise representation of Pi, e.g.
	// 3.14159265358979323846264338327950288419716939937510 should not be
	// truncated to 3.141592653589793 (as is the case for float64) without
	// generating an error.

	// Parse floating point constant.
	var err error
	v.x, err = strconv.ParseFloat(s, size)
	if err != nil {
		return nil, errutil.Newf("invalid floating point constant %q; %v", s, err)
	}

	return v, nil
}

// Type returns the type of the value.
func (v *Float) Type() types.Type {
	return v.typ
}

// UseList returns a list of all values which uses the value.
func (v *Float) UseList() []values.Value {
	panic("not yet implemented.")
}

// ReplaceAll replaces all uses of the value with new.
func (v *Float) ReplaceAll(new values.Value) error {
	panic("not yet implemented.")
}

// String returns the string representation of v using scientific notation (e.g.
// -2.5e+10) for large exponents and regular floating point representation
// otherwise (e.g. 3.14).
func (v *Float) String() string {
	var size int
	switch v.typ.Kind() {
	case types.Float32:
		size = 32
	case types.Float64:
		size = 64
	default:
		// TODO: Add support for half, fp128, x86_fp80 and ppc_fp128.
		err := fmt.Sprintf("not yet implemented; support for %q floating point constants", v.typ)
		panic(err)
	}
	return strconv.FormatFloat(v.x, 'g', -1, size)
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
	typ *types.Pointer
}

// TODO: Figure out how to represent pointer constants. Add the necessary fields
// to the Pointer struct and implement the NewPointer constructor afterwards.

// Type returns the type of the value.
func (v *Pointer) Type() types.Type {
	return v.typ
}

// UseList returns a list of all values which uses the value.
func (v *Pointer) UseList() []values.Value {
	panic("not yet implemented.")
}

// ReplaceAll replaces all uses of the value with new.
func (v *Pointer) ReplaceAll(new values.Value) error {
	panic("not yet implemented.")
}

// isConst ensures that only constant values can be assigned to the Constant
// interface.
func (*Int) isConst()     {}
func (*Float) isConst()   {}
func (*Pointer) isConst() {}
