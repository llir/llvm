package constant

import (
	"fmt"

	"github.com/llir/llvm/ir/types"
)

// --- [ Token constants ] -----------------------------------------------------

// none is an LLVM IR none token constant.
type none struct {
}

// String returns the LLVM syntax representation of the constant as a type-value
// pair.
func (c *none) String() string {
	return fmt.Sprintf("%v %v", c.Type(), c.Ident())
}

// Type returns the type of the constant.
func (*none) Type() types.Type {
	return types.Token
}

// Ident returns the identifier associated with the constant.
func (*none) Ident() string {
	// "none"
	return "none"
}
