package constant

import (
	"fmt"

	"github.com/llir/llvm/ir"
	"github.com/llir/llvm/ir/enum"
	"github.com/llir/llvm/ir/types"
)

// --- [ Other expressions ] ---------------------------------------------------

// ~~~ [ icmp ] ~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~

// ExprICmp is an LLVM IR icmp expression.
type ExprICmp struct {
	// Integer comparison predicate.
	Pred enum.IPred
	// Integer scalar or vector operands.
	X, Y ir.Constant
}

// NewICmp returns a new icmp expression based on the given integer comparison
// predicate and integer scalar or vector operands.
func NewICmp(pred enum.IPred, x, y ir.Constant) *ExprICmp {
	return &ExprICmp{Pred: pred, X: x, Y: y}
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
	// "icmp" IPred "(" Type Constant "," Type Constant ")"
	return fmt.Sprintf("icmp %v (%v, %v)", e.Pred, e.X, e.Y)
}

// Simplify returns an equivalent (and potentially simplified) constant to the
// constant expression.
func (e *ExprICmp) Simplify() ir.Constant {
	panic("not yet implemented")
}

// ~~~ [ fcmp ] ~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~

// ExprFCmp is an LLVM IR fcmp expression.
type ExprFCmp struct {
	// Floating-point comparison predicate.
	Pred enum.FPred
	// Floating-point scalar or vector operands.
	X, Y ir.Constant
}

// NewFCmp returns a new fcmp expression based on the given floating-point
// comparison predicate and floating-point scalar or vector operands.
func NewFCmp(pred enum.FPred, x, y ir.Constant) *ExprFCmp {
	return &ExprFCmp{Pred: pred, X: x, Y: y}
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
	// "fcmp" FPred "(" Type Constant "," Type Constant ")"
	return fmt.Sprintf("fcmp %v (%v, %v)", e.Pred, e.X, e.Y)
}

// Simplify returns an equivalent (and potentially simplified) constant to the
// constant expression.
func (e *ExprFCmp) Simplify() ir.Constant {
	panic("not yet implemented")
}

// ~~~ [ select ] ~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~

// ExprSelect is an LLVM IR select expression.
type ExprSelect struct {
	// Selection condition.
	Cond ir.Constant
	// Operands.
	X, Y ir.Constant
}

// NewSelect returns a new select expression based on the given selection
// condition and operands.
func NewSelect(cond, x, y ir.Constant) *ExprSelect {
	return &ExprSelect{Cond: cond, X: x, Y: x}
}

// String returns the LLVM syntax representation of the constant expression as a
// type-value pair.
func (e *ExprSelect) String() string {
	return fmt.Sprintf("%v %v", e.Type(), e.Ident())
}

// Type returns the type of the constant expression.
func (e *ExprSelect) Type() types.Type {
	return e.X.Type()
}

// Ident returns the identifier associated with the constant expression.
func (e *ExprSelect) Ident() string {
	// "select" "(" Type Constant "," Type Constant "," Type Constant ")"
	return fmt.Sprintf("select (%v, %v, %v)", e.Cond, e.X, e.Y)
}

// Simplify returns an equivalent (and potentially simplified) constant to the
// constant expression.
func (e *ExprSelect) Simplify() ir.Constant {
	panic("not yet implemented")
}
