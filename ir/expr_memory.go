package ir

import "github.com/llir/l/ir/types"

// --- [ Memory expressions ] --------------------------------------------------

// ~~~ [ getelementptr ] ~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~

// GetElementPtrExpr is an LLVM IR getelementptr expression.
type GetElementPtrExpr struct {
	// Source address.
	Src Constant
	// Element indicies.
	Indices []Constant
}

// NewGetElementPtrExpr returns a new getelementptr expression based on the
// given source address and element indices.
func NewGetElementPtrExpr(src Constant, indices ...Constant) *GetElementPtrExpr {
	return &GetElementPtrExpr{Src: src, Indices: indices}
}

// Type returns the type of the constant expression.
func (e *GetElementPtrExpr) Type() types.Type {
	panic("not yet implemented")
}

// Ident returns the identifier associated with the constant expression.
func (e *GetElementPtrExpr) Ident() string {
	panic("not yet implemented")
}

// Simplify returns an equivalent (and potentially simplified) constant of the
// constant expression.
func (e *GetElementPtrExpr) Simplify() Constant {
	panic("not yet implemented")
}
