package constant

import (
	"fmt"
	"math/big"

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
//         [-]? [0-9]+
//    * hexadecimal integer literal
//         TODO: add support for hexadecimal integer literal notation.
func NewIntFromString(typ *types.IntType, s string) (*Int, error) {
	// TODO: handle boolean literals and hexadecimal integer literals.
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
	// int_lit
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
