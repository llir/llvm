package ir

import (
	"fmt"

	"github.com/llir/l/ir/types"
)

// --- [ Memory expressions ] --------------------------------------------------

// ~~~ [ getelementptr ] ~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~

// ExprGetElementPtr is an LLVM IR getelementptr expression.
type ExprGetElementPtr struct {
	// Source address.
	Src Constant
	// Element indicies.
	Indices []Constant
}

// NewGetElementPtrExpr returns a new getelementptr expression based on the
// given source address and element indices.
func NewGetElementPtrExpr(src Constant, indices ...Constant) *ExprGetElementPtr {
	return &ExprGetElementPtr{Src: src, Indices: indices}
}

// String returns the LLVM syntax representation of the constant expression as a
// type-value pair.
func (e *ExprGetElementPtr) String() string {
	return fmt.Sprintf("%v %v", e.Type(), e.Ident())
}

// Type returns the type of the constant expression.
func (e *ExprGetElementPtr) Type() types.Type {
	panic("not yet implemented")
}

// Ident returns the identifier associated with the constant expression.
func (e *ExprGetElementPtr) Ident() string {
	panic("not yet implemented")
}

// Simplify returns an equivalent (and potentially simplified) constant to the
// constant expression.
func (e *ExprGetElementPtr) Simplify() Constant {
	panic("not yet implemented")
}
