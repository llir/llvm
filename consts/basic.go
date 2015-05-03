package consts

import (
	"fmt"
	"math/big"
	"strconv"
	"strings"

	"github.com/llir/llvm/types"
)

// Int represents an integer constant.
//
// Examples:
//    42, -37, true, false, [us]0x[0-9A-Fa-f]+
//
// References:
//    http://llvm.org/docs/LangRef.html#simple-constants
type Int struct {
	// Integer type.
	typ *types.Int
	// Integer value.
	x *big.Int
}

// NewInt returns an integer constant based on the given integer type and string
// representation.
func NewInt(typ types.Type, s string) (*Int, error) {
	// Verify integer type.
	v := &Int{
		x: new(big.Int),
	}
	var ok bool
	v.typ, ok = typ.(*types.Int)
	if !ok {
		return nil, fmt.Errorf("invalid type %q for integer constant", typ)
	}

	// Parse boolean constant.
	if v.typ.Size() == 1 {
		switch s {
		case "1", "true":
			v.x.SetInt64(1)
		case "0", "false":
			v.x.SetInt64(0)
		default:
			return nil, fmt.Errorf("invalid integer constant %q for boolean type", s)
		}
		return v, nil
	} else if s == "true" || s == "false" {
		return nil, fmt.Errorf("integer constant %q type mismatch; expected i1, got %v", s, typ)
	}

	// TODO: Implement support for the HexIntConstant representation:
	//    [us]0x[0-9A-Fa-f]+
	// TODO: Track the upstream removal of HexIntConstant
	//    ref: http://lists.cs.uiuc.edu/pipermail/llvmdev/2015-February/081621.html

	// Parse integer constant.
	if _, ok = v.x.SetString(s, 10); !ok {
		return nil, fmt.Errorf("unable to parse integer constant %q", s)
	}

	return v, nil
}

// Type returns the type of the value.
func (v *Int) Type() types.Type {
	return v.typ
}

// String returns a string representation of the integer, either as a signed
// integer (e.g. 42, -13) or as a boolean (e.g. true, false) depending on the
// type. The integer string representation is preceded by the type of the
// constant, e.g.
//
//    i1 true
//    i32 -13
//    i64 42
func (v *Int) String() string {
	var s string
	if v.typ.Size() == 1 {
		switch v.x.Int64() {
		case 1:
			s = "true"
		default:
			s = "false"
		}
	} else {
		s = v.x.String()
	}

	return fmt.Sprintf("%s %s", v.Type(), s)
}

// Float represents a floating point constant.
//
// Examples:
//    123.45, 1.2345e2, 0x[KLMH]?[0-9A-Fa-f]+
//
// References:
//    http://llvm.org/docs/LangRef.html#simple-constants
type Float struct {
	// Floating point type.
	typ *types.Float
	// Floating point value.
	x *big.Float
}

// NewFloat returns a floating point constant based on the given floating point
// type and string representation.
func NewFloat(typ types.Type, s string) (*Float, error) {
	// Verify floating point type.
	v := &Float{
		x: new(big.Float),
	}
	var ok bool
	v.typ, ok = typ.(*types.Float)
	if !ok {
		return nil, fmt.Errorf("invalid type %q for floating point constant", typ)
	}

	// TODO: Implement support for the following representation:
	//    0x[KLMH]?[0-9A-Fa-f]+

	// Parse floating point constant.
	if _, ok := v.x.SetString(s); !ok {
		return nil, fmt.Errorf("unable to parse floating point constant %q", s)
	}

	// Verify that there was no precision loss.
	size := v.typ.Size()
	switch size {
	case 32:
		if x, acc := v.x.Float32(); acc != big.Exact {
			return nil, fmt.Errorf(`invalid floating point constant %q for type %q; precision loss ("%g")`, s, v.typ, x)
		}
	case 64:
		if x, acc := v.x.Float64(); acc != big.Exact {
			return nil, fmt.Errorf(`invalid floating point constant %q for type %q; precision loss ("%g")`, s, v.typ, x)
		}
	}

	return v, nil
}

// Type returns the type of the value.
func (v *Float) Type() types.Type {
	return v.typ
}

// String returns a string representation of the floating point constant using
// scientific notation (e.g. -2.5e10) for large exponents and regular floating
// point representation otherwise (e.g. 3.14). The floating point string
// representation is preceded by the type of the constant, e.g.
//
//    float 2.0
//    double 3.14
//    double -2.5e10
func (v *Float) String() string {
	// TODO: Replace the code between the "START" and "END" comments with
	//
	//     s := v.x.Format('g', -1)
	//
	// when big.Float.bigFtoa has been implemented. Right now bigFtoa contains
	// the following comment "TODO(gri): complete this". Or better yet, replace
	// the code with
	//
	//     s := v.x.String()
	//
	// when big.Float.String() is precise as tracked by the following comment
	// "BUG(gri): Float.String uses x.Format('g', 10) rather than x.Format('g', -1).".

	// START
	var s string
	size := v.typ.Size()
	switch {
	case size <= 64:
		x, _ := v.x.Float64()
		s = strconv.FormatFloat(x, 'g', -1, size)
	default:
		s = v.x.String()
	}
	// END

	// Insert decimal point if not present.
	//    3e4 -> 3.0e4
	//    42  -> 42.0
	if !strings.ContainsRune(s, '.') {
		pos := strings.IndexByte(s, 'e')
		if pos != -1 {
			s = s[:pos] + ".0" + s[pos:]
		} else {
			s = s + ".0"
		}
	}

	// Drop explicit plus sign in exponents.
	//    3.0e+4 -> 3.0e4
	s = strings.Replace(s, "e+", "e", -1)

	return fmt.Sprintf("%s %s", v.Type(), s)
}

// Pointer represents a pointer constant.
//
// Examples:
//    null, @foo
//
// References:
//    http://llvm.org/docs/LangRef.html#simple-constants
type Pointer struct {
	// Pointer type.
	typ *types.Pointer
}

// TODO: Figure out how to represent pointer constants. Add the necessary fields
// to the Pointer struct and implement the NewPointer constructor afterwards.

// Type returns the type of the value.
func (v *Pointer) Type() types.Type {
	return v.typ
}

// isConst ensures that only constant values can be assigned to the Constant
// interface.
func (*Int) isConst()     {}
func (*Float) isConst()   {}
func (*Pointer) isConst() {}
