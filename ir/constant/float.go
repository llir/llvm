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
	c := &Float{X: &big.Float{}, Typ: t}

	// Parse floating-point literal based on its representation format.
	//
	//   HexFP80Constant   0xK[0-9A-Fa-f]+    // 20 hex digits
	//   HexFP128Constant  0xL[0-9A-Fa-f]+    // 32 hex digits
	//   HexPPC128Constant 0xM[0-9A-Fa-f]+    // 32 hex digits
	//   HexHalfConstant   0xH[0-9A-Fa-f]+    // 4 hex digits
	//   HexFPConstant     0x[0-9A-Fa-f]+     // 16 hex digits
	//   FPConstant        [-+]?[0-9]+[.][0-9]*([eE][-+]?[0-9]+)?
	switch {
	case strings.HasPrefix(s, "0xK"):
		// TODO: Implement support for the 0xK floating-point representation format.
		c.X.SetString("0.0") // TODO: Remove placehold zero value.
		//panic(fmt.Errorf("support for floating-point constant 0xK representation not yet implemented; unable to parse floating-point constant %q", s))
		s = s[len("0xK"):]
	case strings.HasPrefix(s, "0xL"):
		// TODO: Implement support for the 0xL floating-point representation format.
		c.X.SetString("0.0") // TODO: Remove placehold zero value.
		//panic(fmt.Errorf("support for floating-point constant 0xL representation not yet implemented; unable to parse floating-point constant %q", s))
		s = s[len("0xL"):]
	case strings.HasPrefix(s, "0xM"):
		// TODO: Implement support for the 0xM floating-point representation format.
		c.X.SetString("0.0") // TODO: Remove placehold zero value.
		//panic(fmt.Errorf("support for floating-point constant 0xM representation not yet implemented; unable to parse floating-point constant %q", s))
		s = s[len("0xM"):]
	case strings.HasPrefix(s, "0xH"):
		// TODO: Implement support for the 0xH floating-point representation format.
		c.X.SetString("0.0") // TODO: Remove placehold zero value.
		//panic(fmt.Errorf("support for floating-point constant 0xH representation not yet implemented; unable to parse floating-point constant %q", s))
		s = s[len("0xH"):]
	case strings.HasPrefix(s, "0x"):
		// TODO: Implement support for the 0x floating-point representation format.
		c.X.SetString("0.0") // TODO: Remove placehold zero value.
		//panic(fmt.Errorf("support for floating-point constant 0x representation not yet implemented; unable to parse floating-point constant %q", s))
		s = s[len("0x"):]
	default:
		if _, ok := c.X.SetString(s); !ok {
			panic(fmt.Errorf("unable to parse floating-point constant %q", s))
		}
	}

	// Verify that there was no precision loss.

	// TODO: Re-enable precision loss validation. Disabled for now to allow for
	// parsing of as part of sqlite.ll.
	//
	//    %797 = fmul double %796, 1.000000e-01

	//switch t.Kind {
	//case types.FloatKindIEEE_32:
	//	if x, acc := c.X.Float32(); acc != big.Exact {
	//		panic(fmt.Errorf(`invalid floating-point constant %q for type %s; precision loss ("%g")`, s, typ, x))
	//	}
	//case types.FloatKindIEEE_64:
	//	if x, acc := c.X.Float64(); acc != big.Exact {
	//		panic(fmt.Errorf(`invalid floating-point constant %q for type %s; precision loss ("%g")`, s, typ, x))
	//	}
	//}
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
