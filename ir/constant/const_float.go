package constant

import (
	"fmt"
	"log"
	"math"
	"math/big"
	"strconv"
	"strings"

	"github.com/mewmew/floats/float80x86"

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
		switch {
		case strings.HasPrefix(s, "0xK"):
			hex := s[len("0xK"):]
			part1 := hex[:4]
			part2 := hex[4:]
			se, err := strconv.ParseUint(part1, 16, 16)
			if err != nil {
				return nil, errors.WithStack(err)
			}
			m, err := strconv.ParseUint(part2, 16, 64)
			if err != nil {
				return nil, errors.WithStack(err)
			}
			f := float80x86.NewFromBits(uint16(se), m)
			x, nan := f.Big()
			return &Float{
				Typ: typ,
				X:   x,
				NaN: nan,
			}, nil
		case strings.HasPrefix(s, "0xL"):
			//s = s[len("0xL"):]
		case strings.HasPrefix(s, "0xM"):
			//s = s[len("0xM"):]
		case strings.HasPrefix(s, "0xH"):
			//s = s[len("0xH"):]
		default:
			hex := s[len("0x"):]
			x, err := strconv.ParseUint(hex, 16, 64)
			if err != nil {
				return nil, errors.WithStack(err)
			}
			switch typ.Kind {
			case types.FloatKindFloat:
				// ref: https://groups.google.com/d/msg/llvm-dev/IlqV3TbSk6M/27dAggZOMb0J
				//
				// The exact bit representation of the float is laid out with the
				// corresponding bitwise representation of a double:  the sign bit is
				// copied over, the exponent is encoded in the larger width, and the 23
				// bits of significand fills in the top 23 bits of significand in the
				// double.  A double has 52 bits of significand, so this means that the
				// last 29 bits of significand will always be ignored.  As an
				// error-detection measure, the IR parser requires them to be zero.
				f := math.Float64frombits(x)
				// precision=24 is handled during read-out. use precision=53 for
				// internal storage.
				c := big.NewFloat(f)
				const precision = 24
				c.SetPrec(precision)
				return &Float{
					Typ: typ,
					X:   c,
				}, nil
			case types.FloatKindDouble:
				f := math.Float64frombits(x)
				return &Float{
					Typ: typ,
					X:   big.NewFloat(f),
				}, nil
			default:
				panic(fmt.Errorf("support for hexadecimal floating-point literal %q of kind %v not yet implemented", s, typ.Kind))
			}
		}
		log.Printf("constant.NewFloatFromString(%q): not yet implemented", s)
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
	return fmt.Sprintf("%s %s", c.Type(), c.Ident())
}

// Type returns the type of the constant.
func (c *Float) Type() types.Type {
	return c.Typ
}

// Ident returns the identifier associated with the constant.
func (c *Float) Ident() string {
	// FloatLit
	// TODO: add support for hexadecimal format.
	// TODO: add support for NaN, +-Inf.

	switch c.Typ.Kind {
	case types.FloatKindX86FP80:
		// TODO: handle NaN.
		f, acc := float80x86.NewFromBig(c.X)
		// TODO: check acc.
		_ = acc
		se, m := f.Bits()
		return fmt.Sprintf("0xK%04X%016X", se, m)
	case types.FloatKindFloat:
		// ref: https://groups.google.com/d/msg/llvm-dev/IlqV3TbSk6M/27dAggZOMb0J
		//
		// The exact bit representation of the float is laid out with the
		// corresponding bitwise representation of a double:  the sign bit is
		// copied over, the exponent is encoded in the larger width, and the 23
		// bits of significand fills in the top 23 bits of significand in the
		// double.  A double has 52 bits of significand, so this means that the
		// last 29 bits of significand will always be ignored.  As an
		// error-detection measure, the IR parser requires them to be zero.
		if c.NaN || c.X.IsInf() || !exact32(c.X) {
			// Single precision.
			//
			//     1 bit:  sign
			//     8 bits: exponent
			//    23 bits: mantissa
			//
			//    bias: 127
			f, _ := c.X.Float32()
			bits32 := math.Float32bits(f)
			// 0b10000000000000000000000000000000
			sign := uint64(bits32 & 0x80000000 >> 31)
			// 0b01111111100000000000000000000000
			const bias32 = 127
			exp := uint64((bits32 & 0x7F800000 >> 23) - bias32)
			// 0b00000000011111111111111111111111
			mant := uint64(bits32 & 0x7FFFFF)
			// Double precision.
			//
			//     1 bit:  sign
			//    11 bits: exponent
			//    52 bits: mantissa
			//
			//    bias: 1023
			var bits64 uint64
			bits64 |= sign << 53
			const bias64 = 1023
			bits64 |= (exp + bias64) << 52
			bits64 |= mant << (52 - 23)
			return fmt.Sprintf("0x%016X", bits64)
		}
	case types.FloatKindDouble:
		if c.NaN || c.X.IsInf() || !exact64(c.X) {
			f, _ := c.X.Float64()
			x := math.Float64bits(f)
			// Note, to match Clang output we do not zero-pad the hexadecimal
			// output.
			return fmt.Sprintf("0x%X", x)
			//return fmt.Sprintf("0x%016X", x)
		}
	}

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
	return s
}

// exact32 reports whether x may be represented exactly as a float32,
func exact32(x *big.Float) bool {
	f, acc := x.Float32()
	if acc != big.Exact {
		return false
	}
	s1 := strconv.FormatFloat(float64(f), 'e', -1, 32)
	s2 := trimZeros(x.Text('e', 100))
	return s1 == s2
}

// exact64 reports whether x may be represented exactly as a float64,
func exact64(x *big.Float) bool {
	f, acc := x.Float64()
	if acc != big.Exact {
		return false
	}
	s1 := strconv.FormatFloat(f, 'e', -1, 64)
	s2 := trimZeros(x.Text('e', 100))
	return s1 == s2
}

// trimZeros trims trailing zeroes after the decimal point in the given
// floating-point value (represented in scientific notation). If all digits
// after the decimal point are trimmed this way, the decimal point is also
// trimmed.
func trimZeros(s string) string {
	epos := strings.Index(s, "e")
	if epos == -1 {
		panic(fmt.Errorf("unable to locate position of exponent (e.g. e+02) in %q", s))
	}
	pos := epos - 1
	for ; pos > 0; pos-- {
		if s[pos] != '0' {
			break
		}
	}
	if s[pos] != '.' {
		pos++
	}
	return fmt.Sprintf("%s%s", s[:pos], s[epos:])
}
