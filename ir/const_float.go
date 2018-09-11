package ir

import (
	"math/big"

	"github.com/llir/l/ir/types"
)

// Float is a floating-point constant.
type Float struct {
	// Floating-point type.
	Typ *types.FloatType
	// Floating-point constant.
	X *big.Float
	// NaN specifies whether the floating-point constant is Not-a-Number.
	NaN bool
}

// NewFloatFromFloat64 returns a new floating-point constant based on the given
// floating-point type and double precision floating-point value.
func NewFloatFromFloat64(typ *types.FloatType, x float64) *Float {
	panic("not yet implemented")
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
func NewFloatFromString(typ *types.FloatType, s string) *Float {
	panic("not yet implemented")
}

// Type returns the type of the constant.
func (c *Float) Type() types.Type {
	panic("not yet implemented")
}

// Ident returns the identifier associated with the constant.
func (c *Float) Ident() string {
	panic("not yet implemented")
}
