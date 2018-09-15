package ir

import (
	"fmt"
	"math/big"

	"github.com/llir/l/ir/types"
)

// --- [ Integer constants ] ---------------------------------------------------

// ConstInt is an LLVM IR integer constant.
type ConstInt struct {
	// Integer type.
	Typ *types.IntType
	// Integer constant.
	X *big.Int
}

// NewInt returns a new integer constant based on the given integer type and
// 64-bit interger value.
func NewInt(typ *types.IntType, x int64) *ConstInt {
	return &ConstInt{Typ: types.I1, X: big.NewInt(x)}
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
func NewIntFromString(typ *types.IntType, s string) *ConstInt {
	panic("not yet implemented")
}

// String returns the LLVM syntax representation of the constant as a type-value
// pair.
func (c *ConstInt) String() string {
	return fmt.Sprintf("%v %v", c.Type(), c.Ident())
}

// Type returns the type of the constant.
func (c *ConstInt) Type() types.Type {
	return c.Typ
}

// Ident returns the identifier associated with the constant.
func (c *ConstInt) Ident() string {
	panic("not yet implemented")
}
