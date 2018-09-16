package ir

import (
	"fmt"

	"github.com/llir/l/ir/types"
)

// --- [ Null pointer constants ] ----------------------------------------------

// ConstNull is an LLVM IR null pointer constant.
type ConstNull struct {
	// Pointer type.
	Typ *types.PointerType
}

// NewNull returns a new null pointer constant based on the given pointer type.
func NewNull(typ *types.PointerType) *ConstNull {
	return &ConstNull{Typ: typ}
}

// String returns the LLVM syntax representation of the constant as a type-value
// pair.
func (c *ConstNull) String() string {
	return fmt.Sprintf("%v %v", c.Type(), c.Ident())
}

// Type returns the type of the constant.
func (c *ConstNull) Type() types.Type {
	return c.Typ
}

// Ident returns the identifier associated with the constant.
func (*ConstNull) Ident() string {
	// "null"
	return "null"
}
