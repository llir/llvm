package constant

import (
	"fmt"

	"github.com/llir/llvm/ir/types"
)

// --- [ dso_local_equivalent constants ] --------------------------------------

// DSOLocalEquivalent is an LLVM IR dso_local_equivalent constant; a constant
// representing a function which is functionally equivalent to a given
// function, but is always defined in the current linkage unit.
//
// ref: https://llvm.org/docs/LangRef.html#dso-local-equivalent
type DSOLocalEquivalent struct {
	// Underlying function.
	Func Constant // *ir.Func
}

// NewDSOLocalEquivalent returns a new dso_local_equivalent constant based on
// the given function.
func NewDSOLocalEquivalent(f Constant) *DSOLocalEquivalent {
	return &DSOLocalEquivalent{Func: f}
}

// String returns the LLVM syntax representation of the constant as a type-value
// pair.
func (c *DSOLocalEquivalent) String() string {
	return fmt.Sprintf("%s %s", c.Type(), c.Ident())
}

// Type returns the type of the constant.
func (c *DSOLocalEquivalent) Type() types.Type {
	return c.Func.Type()
}

// Ident returns the identifier associated with the constant.
func (c *DSOLocalEquivalent) Ident() string {
	// 'dso_local_equivalent' Func=GlobalIdent
	return fmt.Sprintf("dso_local_equivalent %s", c.Func.Ident())
}
