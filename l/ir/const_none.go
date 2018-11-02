package ir

import (
	"fmt"

	"github.com/llir/l/ir/types"
)

// --- [ Token constants ] -----------------------------------------------------

// ConstNone is an LLVM IR none token constant.
type ConstNone struct {
}

// String returns the LLVM syntax representation of the constant as a type-value
// pair.
func (c *ConstNone) String() string {
	return fmt.Sprintf("%v %v", c.Type(), c.Ident())
}

// Type returns the type of the constant.
func (*ConstNone) Type() types.Type {
	return types.Token
}

// Ident returns the identifier associated with the constant.
func (*ConstNone) Ident() string {
	// "none"
	return "none"
}
