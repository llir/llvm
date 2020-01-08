package constant

import (
	"fmt"
	"math/big"
	"strings"

	"github.com/llir/llvm/ir/types"
	"github.com/pkg/errors"
)

// --- [ Integer constants ] ---------------------------------------------------

// Int is an LLVM IR integer constant.
type Int struct {
	// Integer type.
	Typ *types.IntType
	// Integer constant.
	X *big.Int
}

// NewInt returns a new integer constant based on the given integer type and
// 64-bit interger value.
func NewInt(typ *types.IntType, x int64) *Int {
	return &Int{Typ: typ, X: big.NewInt(x)}
}

// NewBool returns a new boolean constant based on the given boolean value.
func NewBool(x bool) *Int {
	if x {
		return True
	}
	return False
}

// NewIntFromString returns a new integer constant based on the given integer
// type and string.
//
// The integer string may be expressed in one of the following forms.
//
//    * boolean literal
//         true | false
//    * integer literal
//         [-]?[0-9]+
//    * hexadecimal integer literal
//         [us]0x[0-9A-Fa-f]+
func NewIntFromString(typ *types.IntType, s string) (*Int, error) {
	// Boolean literal.
	switch s {
	case "true":
		if !typ.Equal(types.I1) {
			return nil, errors.Errorf("invalid boolean type; expected i1, got %T", typ)
		}
		return True, nil
	case "false":
		if !typ.Equal(types.I1) {
			return nil, errors.Errorf("invalid boolean type; expected i1, got %T", typ)
		}
		return False, nil
	}
	// Hexadecimal integer literal.
	switch {
	// unsigned hexadecimal integer literal
	case strings.HasPrefix(s, "u0x"):
		s = s[len("u0x"):]
		const base = 16
		x, _ := (&big.Int{}).SetString(s, base)
		if x == nil {
			return nil, errors.Errorf("unable to parse integer constant %q", s)
		}
		return &Int{Typ: typ, X: x}, nil
	// signed hexadecimal integer literal
	case strings.HasPrefix(s, "s0x"):
		// Parse signed hexadecimal integer literal in two's complement notation.
		// First parse as unsigned hex, then check if sign bit is set.
		s = s[len("s0x"):]
		const base = 16
		x, _ := (&big.Int{}).SetString(s, base)
		if x == nil {
			return nil, errors.Errorf("unable to parse integer constant %q", s)
		}
		// Check if signed.
		if x.Bit(int(typ.BitSize)-1) == 1 {
			// Compute actual negative value from two's complement.
			//
			// If x is 0xFFFF with type i16, then the actual negative value is
			// `x - 0x10000`, in other words `x - 2^n`.
			n := int64(typ.BitSize)
			// n^2
			maxPlus1 := new(big.Int).Exp(big.NewInt(2), big.NewInt(n), nil)
			x = new(big.Int).Sub(x, maxPlus1)

		}
		return &Int{Typ: typ, X: x}, nil
	}
	// Integer literal.
	x, _ := (&big.Int{}).SetString(s, 10)
	if x == nil {
		return nil, errors.Errorf("unable to parse integer constant %q", s)
	}
	return &Int{Typ: typ, X: x}, nil
}

// String returns the LLVM syntax representation of the constant as a type-value
// pair.
func (c *Int) String() string {
	return fmt.Sprintf("%v %v", c.Type(), c.Ident())
}

// Type returns the type of the constant.
func (c *Int) Type() types.Type {
	return c.Typ
}

// Ident returns the identifier associated with the constant.
func (c *Int) Ident() string {
	// IntLit
	if c.Typ.BitSize == 1 {
		// "true"
		// "false"
		switch x := c.X.Int64(); x {
		case 0:
			return "false"
		case 1:
			return "true"
		default:
			panic(fmt.Errorf("invalid integer value of boolean type; expected 0 or 1, got %d", x))
		}
	}
	// Output x in hexadecimal notation if x is positive, greater than or equal
	// to 0x1000 and has a significantly lower entropy than decimal notation.
	const minEntropyDiff = 0.3
	threshold := big.NewInt(0x1000) // 4096
	// Check entropy if x is >= 0x1000.
	if c.X.Cmp(threshold) >= 0 {
		if decimalEntropy(c.X) >= hexEntropy(c.X)+minEntropyDiff {
			return "u0x" + strings.ToUpper(c.X.Text(16))
		}
	}
	return c.X.String()
}

// ### [ Helper functions ] ####################################################

// hexEntropy returns the entropy of x when encoded in hexadecimal notation. The
// entropy is in range (0.0, 1.0] and is determined by the number of unique hex
// digits required to represent x in hexadecimal notation divided by the total
// number of hex digits ignoring prefix (capped by base 16).
//
// For instance, the hexadecimal value 0x80000000 (2147483648 in decimal)
// requires two unique hex digits to be represented in hexadecimal notation,
// ignoring prefix; namely '0' and '8'.
//
// Hex digits of 0x80000000:
//    0 0 0 0 0 0 0
//    8
//
// The total number of hex digits in 0x80000000 is 8. Thus, the entropy of
// 0x80000000 in hexadecimal notation is
//
//    unique_digits/total_digits
//    = 2/8
//    = 0.25
func hexEntropy(x *big.Int) float64 {
	const base = 16
	return intEntropy(x, base)
}

// decimalEntropy returns the entropy of x when encoded in decimal notation. The
// entropy is in range (0.0, 1.0] and is determined by the number of unique
// decimal digits required to represent x in decimal notation divided by the
// total number of digits (capped by base 10).
//
// For instance, the decimal value 2147483648 (0x80000000 in hex) requires seven
// unique decimal digits to be represented in decimal notation; namely '1', '2',
// '3', '4', '6', '7' and '8'.
//
// Decimal digits of 2147483648:
//    1
//    2
//    3
//    4 4 4
//    6
//    7
//    8 8
//
// The total number of decimal digits in 2147483648 is 10. Thus, the entropy of
// 2147483648 in decimal notation is
//
//    unique_digits/total_digits
//    = 7/10
//    = 0.7
func decimalEntropy(x *big.Int) float64 {
	const base = 10
	return intEntropy(x, base)
}

// intEntropy returns the entropy of x when encoded in base notation. Base must
// be between 2 and 62, inclusive. The entropy is in range (0.0, 1.0] and is
// determined by the number of unique digits required to represent x in base
// notation divided by the total number of digits (capped by base).
func intEntropy(x *big.Int, base int) float64 {
	if base < 2 || base > 62 {
		panic(fmt.Errorf("invalid base; expected 2 <= base <= 62, got %d", base))
	}
	const maxBase = 62
	var digits [maxBase]bool
	s := x.Text(base)
	// Locate unique digits.
	for i := 0; i < len(s); i++ {
		b := s[i]
		if b == '-' {
			// skip sign.
			continue
		}
		d := digitValue(b)
		digits[d] = true
	}
	// Count unique digits.
	uniqueDigits := 0
	for i := 0; i < base; i++ {
		if digits[i] {
			uniqueDigits++
		}
	}
	length := len(s)
	if length > base {
		length = base
	}
	return float64(uniqueDigits) / float64(length)
}

// digitValue returns the integer value of the given digit byte. As defined by
// *big.Int.Text, the digit uses the lower-case letters 'a' to 'z' for digit
// values 10 to 35, and the upper-case letters 'A' to 'Z' for digit values 36 to
// 61.
func digitValue(b byte) int {
	switch {
	case '0' <= b && b <= '9':
		return 0 + int(b-'0')
	case 'a' <= b && b <= 'z':
		return 10 + int(b-'a')
	case 'A' <= b && b <= 'Z':
		return 36 + int(b-'A')
	default:
		panic(fmt.Errorf("invalid digit byte; expected [0-9a-zA-Z], got %#U", b))
	}
}
