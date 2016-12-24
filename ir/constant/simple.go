// === [ Simple constants ] ====================================================
//
// References:
//    http://llvm.org/docs/LangRef.html#simple-constants

package constant

import (
	"fmt"
	"math/big"
	"strings"

	"github.com/llir/llvm/ir/types"
)

// --- [ int ] -----------------------------------------------------------------

// Int represents an integer constant.
type Int struct {
	// Integer type.
	Typ *types.IntType
	// Integer value.
	X *big.Int
}

// NewInt returns a new integer constant based on the given integer value and
// type.
func NewInt(x int64, typ types.Type) *Int {
	t, ok := typ.(*types.IntType)
	if !ok {
		panic(fmt.Errorf("invalid integer constant type; expected *types.IntType, got %T", typ))
	}
	return &Int{Typ: t, X: big.NewInt(x)}
}

// NewIntFromString returns a new integer constant based on the given integer
// string and type.
func NewIntFromString(s string, typ types.Type) *Int {
	// Parse boolean integer constants.
	c := NewInt(0, typ)
	if types.IsBool(c.Typ) {
		switch s {
		case "0", "false":
			c.X.SetInt64(0)
		case "1", "true":
			c.X.SetInt64(1)
		default:
			panic(fmt.Errorf("invalid integer constant %q for type i1", s))
		}
		return c
	} else if s == "true" || s == "false" {
		panic(fmt.Errorf("invalid integer constant %q for type %s", s, typ))
	}

	// Parse decimal or hexadecimal integer constant.
	base := 10
	// Hexadecimal integer format
	//
	//    [us]0x[0-9A-Fa-f]+
	switch {
	case strings.HasPrefix(s, "u0x"), strings.HasPrefix(s, "s0x"):
		s = s[len("u0x"):]
		base = 16
	case strings.HasPrefix(s, "0x"):
		s = s[len("0x"):]
		base = 16
	}
	if _, ok := c.X.SetString(s, base); !ok {
		panic(fmt.Errorf("unable to parse constant %q of type %s", s, typ))
	}
	return c
}

// Type returns the type of the constant.
func (c *Int) Type() types.Type {
	return c.Typ
}

// Ident returns the string representation of the constant.
func (c *Int) Ident() string {
	if c.Typ.Size == 1 {
		switch c.Int64() {
		case 0:
			return "false"
		case 1:
			return "true"
		}
	}
	return c.X.String()
}

// Immutable ensures that only constants can be assigned to the
// constant.Constant interface.
func (*Int) Immutable() {}

// Int64 returns the int64 representation of the integer constant.
func (c *Int) Int64() int64 {
	return c.X.Int64()
}

// --- [ float ] ---------------------------------------------------------------

// Float represents a floating-point constant.
type Float struct {
	// Floating-point type.
	Typ *types.FloatType
	// Floating-point value.
	X *big.Float
}

// NewFloat returns a new floating-point constant based on the given
// floating-point value and type.
func NewFloat(x float64, typ types.Type) *Float {
	t, ok := typ.(*types.FloatType)
	if !ok {
		panic(fmt.Errorf("invalid floating-point constant type; expected *types.FloatType, got %T", typ))
	}
	return &Float{Typ: t, X: big.NewFloat(x)}
}

// NewFloatFromString returns a new floating-point constant based on the given
// floating-point string and type.
func NewFloatFromString(s string, typ types.Type) *Float {
	// Parse floating-point constant.
	t, ok := typ.(*types.FloatType)
	if !ok {
		panic(fmt.Errorf("invalid floating-point constant type; expected *types.FloatType, got %T", typ))
	}
	c := &Float{X: &big.Float{}, Typ: t}
	if _, ok := c.X.SetString(s); !ok {
		panic(fmt.Errorf("unable to parse floating-point constant %q", s))
	}

	// TODO: Implement support for the following floating-point representation:
	//    0x[KLMH]?[0-9A-Fa-f]+

	// Verify that there was no precision loss.
	switch t.Kind {
	case types.FloatKindIEEE_32:
		if x, acc := c.X.Float32(); acc != big.Exact {
			panic(fmt.Errorf(`invalid floating-point constant %q for type %s; precision loss ("%g")`, s, typ, x))
		}
	case types.FloatKindIEEE_64:
		if x, acc := c.X.Float64(); acc != big.Exact {
			panic(fmt.Errorf(`invalid floating-point constant %q for type %s; precision loss ("%g")`, s, typ, x))
		}
	}
	return c
}

// Type returns the type of the constant.
func (c *Float) Type() types.Type {
	return c.Typ
}

// Ident returns the string representation of the constant.
func (c *Float) Ident() string {
	// Insert decimal point if not present.
	//    3e4 -> 3.0e4
	//    42  -> 42.0
	s := c.X.Text('g', -1)
	if !strings.ContainsRune(s, '.') {
		if pos := strings.IndexByte(s, 'e'); pos != -1 {
			s = s[:pos] + ".0" + s[pos:]
		} else {
			s += ".0"
		}
	}

	// Drop explicit plus sign in exponents.
	//    3.0e+4 -> 3.0e4
	return strings.Replace(s, "e+", "e", -1)
}

// Immutable ensures that only constants can be assigned to the
// constant.Constant interface.
func (*Float) Immutable() {}

// Float64 returns the float64 representation of the floating-point constant.
func (c *Float) Float64() float64 {
	x, _ := c.X.Float64()
	return x
}

// --- [ null ] ----------------------------------------------------------------

// Null represents a null pointer constant.
type Null struct {
	// Pointer type.
	Typ *types.PointerType
}

// NewNull returns a new null pointer constant based on the given pointer type.
func NewNull(typ types.Type) *Null {
	t, ok := typ.(*types.PointerType)
	if !ok {
		panic(fmt.Errorf("invalid null pointer constant type; expected *types.PointerType, got %T", typ))
	}
	return &Null{Typ: t}
}

// Type returns the type of the constant.
func (c *Null) Type() types.Type {
	return c.Typ
}

// Ident returns the string representation of the constant.
func (c *Null) Ident() string {
	return "null"
}

// Immutable ensures that only constants can be assigned to the
// constant.Constant interface.
func (*Null) Immutable() {}
