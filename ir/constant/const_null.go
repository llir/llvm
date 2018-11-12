package constant

import (
	"fmt"

	"github.com/llir/llvm/ir/types"
)

// --- [ Null pointer constants ] ----------------------------------------------

// Null is an LLVM IR null pointer constant.
type Null struct {
	// Pointer type.
	Typ *types.PointerType
}

// NewNull returns a new null pointer constant based on the given pointer type.
func NewNull(typ *types.PointerType) *Null {
	return &Null{Typ: typ}
}

// String returns the LLVM syntax representation of the constant as a type-value
// pair.
func (c *Null) String() string {
	return fmt.Sprintf("%v %v", c.Type(), c.Ident())
}

// Type returns the type of the constant.
func (c *Null) Type() types.Type {
	return c.Typ
}

// Ident returns the identifier associated with the constant.
func (*Null) Ident() string {
	// NullLit
	return "null"
}
