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
	return c.X.String()
}
