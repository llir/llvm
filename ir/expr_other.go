package ir

import (
	"fmt"

	"github.com/llir/l/ir/ll"
	"github.com/llir/l/ir/types"
)

// --- [ Other expressions ] ---------------------------------------------------

// ~~~ [ icmp ] ~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~

// ExprICmp is an LLVM IR icmp expression.
type ExprICmp struct {
	// Integer comparison condition.
	Cond ll.ICond
	// Integer scalar or vector operands.
	X, Y Constant
}

// NewICmpExpr returns a new icmp expression based on the given integer
// comparison condition and integer scalar or vector operands.
func NewICmpExpr(cond ll.ICond, x, y Constant) *ExprICmp {
	return &ExprICmp{Cond: cond, X: x, Y: y}
}

// String returns the LLVM syntax representation of the constant expression as a
// type-value pair.
func (e *ExprICmp) String() string {
	return fmt.Sprintf("%v %v", e.Type(), e.Ident())
}

// Type returns the type of the constant expression.
func (e *ExprICmp) Type() types.Type {
	panic("not yet implemented")
}

// Ident returns the identifier associated with the constant expression.
func (e *ExprICmp) Ident() string {
	panic("not yet implemented")
}

// Simplify returns an equivalent (and potentially simplified) constant to the
// constant expression.
func (e *ExprICmp) Simplify() Constant {
	panic("not yet implemented")
}

// ~~~ [ fcmp ] ~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~

// ExprFCmp is an LLVM IR fcmp expression.
type ExprFCmp struct {
	// Floating-point comparison condition.
	Cond ll.FCond
	// Floating-point scalar or vector operands.
	X, Y Constant
}

// NewFCmpExpr returns a new fcmp expression based on the given floating-point
// comparison condition and floating-point scalar or vector operands.
func NewFCmpExpr(cond ll.FCond, x, y Constant) *ExprFCmp {
	return &ExprFCmp{Cond: cond, X: x, Y: y}
}

// String returns the LLVM syntax representation of the constant expression as a
// type-value pair.
func (e *ExprFCmp) String() string {
	return fmt.Sprintf("%v %v", e.Type(), e.Ident())
}

// Type returns the type of the constant expression.
func (e *ExprFCmp) Type() types.Type {
	panic("not yet implemented")
}

// Ident returns the identifier associated with the constant expression.
func (e *ExprFCmp) Ident() string {
	panic("not yet implemented")
}

// Simplify returns an equivalent (and potentially simplified) constant to the
// constant expression.
func (e *ExprFCmp) Simplify() Constant {
	panic("not yet implemented")
}

// ~~~ [ select ] ~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~

// ExprSelect is an LLVM IR select expression.
type ExprSelect struct {
	// Selection condition.
	Cond Constant
	// Operands.
	X, Y Constant
}

// NewSelectExpr returns a new select expression based on the given selection
// condition and operands.
func NewSelectExpr(cond, x, y Constant) *ExprSelect {
	return &ExprSelect{Cond: cond, X: x, Y: x}
}

// String returns the LLVM syntax representation of the constant expression as a
// type-value pair.
func (e *ExprSelect) String() string {
	return fmt.Sprintf("%v %v", e.Type(), e.Ident())
}

// Type returns the type of the constant expression.
func (e *ExprSelect) Type() types.Type {
	panic("not yet implemented")
}

// Ident returns the identifier associated with the constant expression.
func (e *ExprSelect) Ident() string {
	panic("not yet implemented")
}

// Simplify returns an equivalent (and potentially simplified) constant to the
// constant expression.
func (e *ExprSelect) Simplify() Constant {
	panic("not yet implemented")
}
