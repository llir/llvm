package constant

import (
	"fmt"
	"log"
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
		return NewInt(typ, 1), nil
	case "false":
		if !typ.Equal(types.I1) {
			return nil, errors.Errorf("invalid boolean type; expected i1, got %T", typ)
		}
		return NewInt(typ, 0), nil
	}
	// Hexadecimal integer literal.
	switch {
	case strings.HasPrefix(s, "u0x"):
		s = s[len("u0x"):]
		x, _ := (&big.Int{}).SetString(s, 16)
		if x == nil {
			return nil, errors.Errorf("unable to parse integer constant %q", s)
		}
		return &Int{Typ: typ, X: x}, nil
	case strings.HasPrefix(s, "s0x"):
		// TODO: figure out how to handle negative values. Use typ.BitSize.
		// e.g. what value should s0x0012312 represent?
		log.Printf("support for signed hexadecimal integers (%q) not yet implemented", s)
		s = s[len("s0x"):]
		x, _ := (&big.Int{}).SetString(s, 16)
		if x == nil {
			return nil, errors.Errorf("unable to parse integer constant %q", s)
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
