package constant

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/llir/llvm/asm"
	"github.com/llir/llvm/internal/math/big"
	"github.com/llir/llvm/ir/types"
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

	// Parse hexadecimal integer constant.
	if s, ok := stripHexPrefix(s); ok {
		if _, ok := v.x.SetString(s, 16); !ok {
			return nil, fmt.Errorf("unable to parse hexadecimal integer constant %q", s)
		}
		return v, nil
	}

	// Parse integer constant.
	if _, ok := v.x.SetString(s, 10); !ok {
		return nil, fmt.Errorf("unable to parse integer constant %q", s)
	}

	return v, nil
}

// stripHexPrefix attempts to strip a hexadecimal prefix from the given string.
// The returned boolean value reports whether a hexadecimal prefix was located.
//
// Hexadecimal integer format.
//
//    [us]0x[0-9A-Fa-f]+
func stripHexPrefix(s string) (string, bool) {
	switch {
	case strings.HasPrefix(s, "u0x"), strings.HasPrefix(s, "s0x"):
		return s[len("u0x"):], true
	case strings.HasPrefix(s, "0x"):
		return s[len("0x"):], true
	}
	return s, false
}

// Type returns the type of the value.
func (v *Int) Type() types.Type {
	return v.typ
}

// String returns a string representation of the integer, either as a signed
// integer (e.g. 42, -13) or as a boolean (e.g. true, false) depending on the
// type; e.g.
//
//    true
//    -13
//    42
func (v *Int) String() string {
	if v.typ.Size() == 1 {
		switch v.x.Int64() {
		case 1:
			return "true"
		default:
			return "false"
		}
	}
	return v.x.String()
}

// ValueString returns a string representation of the value.
func (v *Int) ValueString() string {
	return v.String()
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
// point representation otherwise (e.g. 3.14); e.g.
//
//    2.0
//    3.14
//    -2.5e10
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
	return strings.Replace(s, "e+", "e", -1)
}

// ValueString returns a string representation of the value.
func (v *Float) ValueString() string {
	return v.String()
}

// TODO: Consider merging Pointer with NullPointer.

// Pointer represents a pointer constant.
//
// Examples:
//    @foo
//
// References:
//    http://llvm.org/docs/LangRef.html#global-variable-and-function-addresses
type Pointer struct {
	// Pointer type.
	typ *types.Pointer
	// Global variable name.
	name string
}

// NewPointer returns a new pointer constant based on the given type and global
// identifier name.
func NewPointer(typ *types.Pointer, name string) (*Pointer, error) {
	return &Pointer{typ: typ, name: name}, nil
}

// Type returns the type of the value.
func (v *Pointer) Type() types.Type {
	return v.typ
}

// String returns a string representation of the pointer constant; e.g.
//
//    @printf
func (v *Pointer) String() string {
	return asm.EncGlobal(v.name)
}

// ValueString returns a string representation of the value.
func (v *Pointer) ValueString() string {
	return v.String()
}

// NullPointer represents a null pointer constant.
//
// Examples:
//    null
//
// References:
//    http://llvm.org/docs/LangRef.html#simple-constants
type NullPointer struct {
	// Pointer type.
	typ *types.Pointer
}

// NewNullPointer returns a new null pointer constant based on the given type.
func NewNullPointer(typ *types.Pointer) (*NullPointer, error) {
	return &NullPointer{typ: typ}, nil
}

// Type returns the type of the value.
func (v *NullPointer) Type() types.Type {
	return v.typ
}

// String returns a string representation of the null pointer constant; e.g.
//
//    null
func (v *NullPointer) String() string {
	return "null"
}

// ValueString returns a string representation of the value.
func (v *NullPointer) ValueString() string {
	return v.String()
}

// ZeroInitializer represents a zero initializer.
//
// Examples:
//    zeroinitializer
//
// References:
//    http://llvm.org/docs/LangRef.html#complex-constants
type ZeroInitializer struct {
	// Value type.
	typ types.Type
}

// NewZeroInitializer returns a new zero initializer based on the given type.
func NewZeroInitializer(typ types.Type) *ZeroInitializer {
	return &ZeroInitializer{typ: typ}
}

// Type returns the type of the value.
func (v *ZeroInitializer) Type() types.Type {
	return v.typ
}

// String returns a string representation of the zero initializer; e.g.
//
//    zeroinitializer
func (v *ZeroInitializer) String() string {
	return "zeroinitializer"
}

// ValueString returns a string representation of the value.
func (v *ZeroInitializer) ValueString() string {
	return v.String()
}

// isConst ensures that only constant values can be assigned to the Constant
// interface.
func (*Int) isConst()             {}
func (*Float) isConst()           {}
func (*Pointer) isConst()         {}
func (*NullPointer) isConst()     {}
func (*ZeroInitializer) isConst() {}
