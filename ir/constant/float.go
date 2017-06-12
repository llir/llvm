// === [ Simple constants ] ====================================================
//
// References:
//    http://llvm.org/docs/LangRef.html#simple-constants

package constant

import (
	"fmt"
	"math/big"
	"strings"

	"github.com/llir/llvm/internal/floats"
	"github.com/llir/llvm/ir/types"
)

// --- [ floating-point ] ------------------------------------------------------

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
	c := &Float{Typ: t}

	// Parse floating-point literal in hexadecimal format.
	switch {
	case strings.HasPrefix(s, "0xK"):
		//   HexFP80Constant   0xK[0-9A-Fa-f]+    // 20 hex digits

		// TODO: Implement support for the 0xK floating-point representation format.
		c.X = &big.Float{}
		c.X.SetString("0.0") // TODO: Remove placehold zero value.
		//panic(fmt.Errorf("support for floating-point constant 0xK representation not yet implemented; unable to parse floating-point constant %q", s))
		s = s[len("0xK"):]
		return c
	case strings.HasPrefix(s, "0xL"):
		//   HexFP128Constant  0xL[0-9A-Fa-f]+    // 32 hex digits

		// TODO: Implement support for the 0xL floating-point representation format.
		c.X = &big.Float{}
		c.X.SetString("0.0") // TODO: Remove placehold zero value.
		//panic(fmt.Errorf("support for floating-point constant 0xL representation not yet implemented; unable to parse floating-point constant %q", s))
		s = s[len("0xL"):]
		return c
	case strings.HasPrefix(s, "0xM"):
		//   HexPPC128Constant 0xM[0-9A-Fa-f]+    // 32 hex digits

		// TODO: Implement support for the 0xM floating-point representation format.
		c.X = &big.Float{}
		c.X.SetString("0.0") // TODO: Remove placehold zero value.
		//panic(fmt.Errorf("support for floating-point constant 0xM representation not yet implemented; unable to parse floating-point constant %q", s))
		s = s[len("0xM"):]
		return c
	case strings.HasPrefix(s, "0xH"):
		//   HexHalfConstant   0xH[0-9A-Fa-f]+    // 4 hex digits

		str := s[len("0xH"):]
		x := floats.NewFloat16FromString(str)
		c.X = big.NewFloat(x.Float64())
		return c
	case strings.HasPrefix(s, "0x"):
		//   HexFPConstant     0x[0-9A-Fa-f]+     // 16 hex digits

		// TODO: Implement support for the 0x floating-point representation format.
		c.X = &big.Float{}
		c.X.SetString("0.0") // TODO: Remove placehold zero value.
		//panic(fmt.Errorf("support for floating-point constant 0x representation not yet implemented; unable to parse floating-point constant %q", s))
		s = s[len("0x"):]
		return c
	}

	// Parse floating-point literal.
	//
	//   FPConstant        [-+]?[0-9]+[.][0-9]*([eE][-+]?[0-9]+)?
	c.X = &big.Float{}
	if _, ok := c.X.SetString(s); !ok {
		panic(fmt.Errorf("unable to parse floating-point constant %q", s))
	}
	return c
}

// Type returns the type of the constant.
func (c *Float) Type() types.Type {
	return c.Typ
}

// Ident returns the string representation of the constant.
func (c *Float) Ident() string {
	// Hexadecimal format is always used for long double, and there are three
	// forms of long double.
	kind := c.Typ.Kind
	switch kind {
	case types.FloatKindIEEE_128:
		// The IEEE 128-bit format is represented by 0xL followed by 32
		// hexadecimal digits.

		// TODO: Implement.
		panic(fmt.Errorf("support for floating-point kind %v not yet implemented", kind))
	case types.FloatKindDoubleExtended_80:
		// The 80-bit format used by x86 is represented as 0xK followed by 20
		// hexadecimal digits.

		// TODO: Represent precise float80 using c.X.MantExp and
		// floats.NewFloat80FromBits.
		f := floats.NewFloat80FromFloat64(c.Float64())
		return "0xK" + f.String()
	case types.FloatKindDoubleDouble_128:
		// The 128-bit format used by PowerPC (two adjacent doubles) is
		// represented by 0xM followed by 32 hexadecimal digits.

		// TODO: Implement.
		panic(fmt.Errorf("support for floating-point kind %v not yet implemented", kind))
	}

	// TODO: Handle NaN, special values (e.g. Pi). Print those in hexadecimal
	// representation.

	// Use hexadecimal representation for +Inf and -Inf.
	printHex := c.X.IsInf()
	if printHex {
		switch kind {
		case types.FloatKindIEEE_16:
			// TODO: Implement.
			x, _ := c.X.Float64()
			f16, _ := floats.NewFloat16FromFloat64(x)
			return "0xH" + f16.String()
		case types.FloatKindIEEE_32:
			// TODO: Implement.
			panic(fmt.Errorf("support for floating-point kind %v not yet implemented", kind))
		case types.FloatKindIEEE_64:
			// TODO: Implement.
			panic(fmt.Errorf("support for floating-point kind %v not yet implemented", kind))
		default:
			panic(fmt.Errorf("support for floating-point kind %v not yet implemented", kind))
		}
	}

	// Insert decimal point if not present.
	//    3e4 -> 3.0e4
	//    42  -> 42.0
	s := c.X.Text('f', -1)
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

// MetadataNode ensures that only metadata nodes can be assigned to the
// ir.MetadataNode interface.
func (*Float) MetadataNode() {}

// Float64 returns the float64 representation of the floating-point constant.
func (c *Float) Float64() float64 {
	x, _ := c.X.Float64()
	return x
}
