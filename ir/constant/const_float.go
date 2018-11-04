package constant

import (
	"fmt"
	"log"
	"math"
	"math/big"
	"strings"

	"github.com/llir/llvm/ir/types"
	"github.com/pkg/errors"
)

// --- [ Floating-point constants ] --------------------------------------------

// Float is an LLVM IR floating-point constant.
type Float struct {
	// Floating-point type.
	Typ *types.FloatType
	// Floating-point constant.
	X *big.Float
	// NaN specifies whether the floating-point constant is Not-a-Number.
	NaN bool
}

// NewFloat returns a new floating-point constant based on the given
// floating-point type and double precision floating-point value.
func NewFloat(typ *types.FloatType, x float64) *Float {
	if math.IsNaN(x) {
		// TODO: store sign of NaN?
		return &Float{Typ: typ, NaN: true}
	}
	return &Float{Typ: typ, X: big.NewFloat(x)}
}

// NewFloatFromString returns a new floating-point constant based on the given
// floating-point type and floating-point string.
//
// The floating-point string may be expressed in one of the following forms.
//
//    * fraction floating-point literal
//         [+-]? [0-9]+ [.] [0-9]*
//    * scientific notation floating-point literal
//         [+-]? [0-9]+ [.] [0-9]* [eE] [+-]? [0-9]+
//    * hexadecimal floating-point literal
//         0x[0-9A-Fa-f]{16}  // HexFP
//         0xK[0-9A-Fa-f]{20} // HexFP80
//         0xL[0-9A-Fa-f]{32} // HexFP128
//         0xM[0-9A-Fa-f]{32} // HexPPC128
//         0xH[0-9A-Fa-f]{4}  // HexHalf
func NewFloatFromString(typ *types.FloatType, s string) (*Float, error) {
	// TODO: implement NewFloatFromString. return 0 for now.
	if strings.HasPrefix(s, "0x") {
		log.Printf("ir.NewFloatFromString(%q): not yet implemented", s)
		return NewFloat(typ, 0), nil
		//panic(fmt.Errorf("support for hexadecimal floating-point constant %q not yet implemented", s))
	}
	switch typ.Kind {
	case types.FloatKindFloat:
		const precision = 24
		x, _, err := big.ParseFloat(s, 10, precision, big.ToNearestEven)
		if err != nil {
			return nil, errors.WithStack(err)
		}
		c := &Float{
			Typ: typ,
			X:   x,
		}
		return c, nil
	case types.FloatKindDouble:
		const precision = 53
		x, _, err := big.ParseFloat(s, 10, precision, big.ToNearestEven)
		if err != nil {
			return nil, errors.WithStack(err)
		}
		c := &Float{
			Typ: typ,
			X:   x,
		}
		return c, nil
	default:
		panic(fmt.Errorf("support for floating-point kind %v not yet implemented", typ.Kind))
	}
}

// String returns the LLVM syntax representation of the constant as a type-value
// pair.
func (c *Float) String() string {
	return fmt.Sprintf("%v %v", c.Type(), c.Ident())
}

// Type returns the type of the constant.
func (c *Float) Type() types.Type {
	return c.Typ
}

// Ident returns the identifier associated with the constant.
func (c *Float) Ident() string {
	// float_lit
	// TODO: add support for hexadecimal format.
	// TODO: add support for NaN.

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
