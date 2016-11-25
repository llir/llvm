// === [ Complex constants ] ===================================================
//
// References:
//    http://llvm.org/docs/LangRef.html#complex-constants

package constant

import "github.com/llir/llvm/ir/types"

// --- [ vector ] --------------------------------------------------------------

type Vector struct {
}

// --- [ array ] ---------------------------------------------------------------

type Array struct {
}

// --- [ struct ] --------------------------------------------------------------

type Struct struct {
}

// --- [ zeroinitializer ] -----------------------------------------------------

// ZeroInitializer represents a zeroinitializer constant.
type ZeroInitializer struct {
	// Constant type.
	typ types.Type
}

// NewZeroInitializer returns a new zeroinitializer constant based on the given
// type.
func NewZeroInitializer(typ types.Type) *ZeroInitializer {
	return &ZeroInitializer{typ: typ}
}

// Type returns the type of the constant.
func (c *ZeroInitializer) Type() types.Type {
	return c.typ
}

// Ident returns the value of the constant.
func (c *ZeroInitializer) Ident() string {
	return "zeroinitializer"
}
