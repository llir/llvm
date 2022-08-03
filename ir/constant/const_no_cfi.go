package constant

import (
	"fmt"

	"github.com/llir/llvm/ir/types"
)

// --- [ no_cfi constants ] ----------------------------------------------------

// NoCFI is an LLVM IR no_cfi constant; a constant representing a function which
// does not get replaced with a reference to the CFI jump table
// (control-flow integrity).
//
// refs:
//
//   - https://llvm.org/docs/LangRef.html#no-cfi
type NoCFI struct {
	// Underlying function.
	Func Constant // *ir.Func
}

// NewNoCFI returns a new no_cfi constant based on the given function.
func NewNoCFI(f Constant) *NoCFI {
	return &NoCFI{Func: f}
}

// String returns the LLVM syntax representation of the constant as a type-value
// pair.
func (c *NoCFI) String() string {
	return fmt.Sprintf("%s %s", c.Type(), c.Ident())
}

// Type returns the type of the constant.
func (c *NoCFI) Type() types.Type {
	return c.Func.Type()
}

// Ident returns the identifier associated with the constant.
func (c *NoCFI) Ident() string {
	// 'no_cfi' Func=GlobalIdent
	return fmt.Sprintf("no_cfi %s", c.Func.Ident())
}
