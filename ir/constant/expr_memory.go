package constant

import "github.com/llir/l/ir/types"

// --- [ Memory expressions ] --------------------------------------------------

// ~~~ [ getelementptr ] ~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~

// GetElementPtr is an LLVM IR getelementptr expression.
type GetElementPtr struct {
	// Source address.
	Src Constant
	// Element indicies.
	Indices []Constant
}

// NewGetElementPtr returns a new getelementptr expression based on the given
// source address and element indices.
func NewGetElementPtr(src Constant, indices ...Constant) *GetElementPtr {
	return &GetElementPtr{Src: src, Indices: indices}
}

// Type returns the type of the constant expression.
func (e *GetElementPtr) Type() types.Type {
	panic("not yet implemented")
}

// Ident returns the identifier associated with the constant expression.
func (e *GetElementPtr) Ident() string {
	panic("not yet implemented")
}

// Simplify returns an equivalent (and potentially simplified) constant of the
// constant expression.
func (e *GetElementPtr) Simplify() Constant {
	panic("not yet implemented")
}
