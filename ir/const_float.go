package ir

import (
	"fmt"
	"log"
	"math"
	"math/big"

	"github.com/llir/llvm/ir/types"
)

// --- [ Floating-point constants ] --------------------------------------------

// ConstFloat is an LLVM IR floating-point constant.
type ConstFloat struct {
	// Floating-point type.
	Typ *types.FloatType
	// Floating-point constant.
	X *big.Float
	// NaN specifies whether the floating-point constant is Not-a-Number.
	NaN bool
}

// NewFloat returns a new floating-point constant based on the given
// floating-point type and double precision floating-point value.
func NewFloat(typ *types.FloatType, x float64) *ConstFloat {
	if math.IsNaN(x) {
		// TODO: store sign of NaN?
		return &ConstFloat{Typ: typ, NaN: true}
	}
	return &ConstFloat{Typ: typ, X: big.NewFloat(x)}
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
func NewFloatFromString(typ *types.FloatType, s string) (*ConstFloat, error) {
	log.Printf("ir.NewFloatFromString(%q): not yet implemented", s)
	// TODO: implement NewFloatFromString. return 0 for now.
	return NewFloat(typ, 0), nil
}

// String returns the LLVM syntax representation of the constant as a type-value
// pair.
func (c *ConstFloat) String() string {
	return fmt.Sprintf("%v %v", c.Type(), c.Ident())
}

// Type returns the type of the constant.
func (c *ConstFloat) Type() types.Type {
	return c.Typ
}

// Ident returns the identifier associated with the constant.
func (c *ConstFloat) Ident() string {
	// float_lit
	panic("not yet implemented")
}
