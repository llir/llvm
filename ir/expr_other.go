package ir

import (
	"github.com/llir/l/ir/ll"
	"github.com/llir/l/ir/types"
)

// --- [ Other expressions ] ---------------------------------------------------

// ~~~ [ icmp ] ~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~

// ICmpExpr is an LLVM IR icmp expression.
type ICmpExpr struct {
	// Integer comparison condition.
	Cond ll.ICond
	// Integer scalar or vector operands.
	X, Y Constant
}

// NewICmpExpr returns a new icmp expression based on the given integer
// comparison condition and integer scalar or vector operands.
func NewICmpExpr(cond ll.ICond, x, y Constant) *ICmpExpr {
	return &ICmpExpr{Cond: cond, X: x, Y: y}
}

// Type returns the type of the constant expression.
func (e *ICmpExpr) Type() types.Type {
	panic("not yet implemented")
}

// Ident returns the identifier associated with the constant expression.
func (e *ICmpExpr) Ident() string {
	panic("not yet implemented")
}

// Simplify returns an equivalent (and potentially simplified) constant of the
// constant expression.
func (e *ICmpExpr) Simplify() Constant {
	panic("not yet implemented")
}

// ~~~ [ fcmp ] ~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~

// FCmpExpr is an LLVM IR fcmp expression.
type FCmpExpr struct {
	// Floating-point comparison condition.
	Cond ll.FCond
	// Floating-point scalar or vector operands.
	X, Y Constant
}

// NewFCmpExpr returns a new fcmp expression based on the given floating-point
// comparison condition and floating-point scalar or vector operands.
func NewFCmpExpr(cond ll.FCond, x, y Constant) *FCmpExpr {
	return &FCmpExpr{Cond: cond, X: x, Y: y}
}

// Type returns the type of the constant expression.
func (e *FCmpExpr) Type() types.Type {
	panic("not yet implemented")
}

// Ident returns the identifier associated with the constant expression.
func (e *FCmpExpr) Ident() string {
	panic("not yet implemented")
}

// Simplify returns an equivalent (and potentially simplified) constant of the
// constant expression.
func (e *FCmpExpr) Simplify() Constant {
	panic("not yet implemented")
}

// ~~~ [ select ] ~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~

// SelectExpr is an LLVM IR select expression.
type SelectExpr struct {
	// Selection condition.
	Cond Constant
	// Operands.
	X, Y Constant
}

// NewSelectExpr returns a new select expression based on the given selection
// condition and operands.
func NewSelectExpr(cond, x, y Constant) *SelectExpr {
	return &SelectExpr{Cond: cond, X: x, Y: x}
}

// Type returns the type of the constant expression.
func (e *SelectExpr) Type() types.Type {
	panic("not yet implemented")
}

// Ident returns the identifier associated with the constant expression.
func (e *SelectExpr) Ident() string {
	panic("not yet implemented")
}

// Simplify returns an equivalent (and potentially simplified) constant of the
// constant expression.
func (e *SelectExpr) Simplify() Constant {
	panic("not yet implemented")
}
