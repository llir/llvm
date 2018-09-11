package constant

import (
	"github.com/llir/l/ir/ll"
	"github.com/llir/l/ir/types"
)

// --- [ Other expressions ] ---------------------------------------------------

// ~~~ [ icmp ] ~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~

// ICmp is an LLVM IR icmp expression.
type ICmp struct {
	// Integer comparison condition.
	Cond ll.ICond
	// Integer scalar or vector operands.
	X, Y Constant
}

// NewICmp returns a new icmp expression based on the given integer comparison
// condition and integer scalar or vector operands.
func NewICmp(cond ll.ICond, x, y Constant) *ICmp {
	return &ICmp{Cond: cond, X: x, Y: y}
}

// Type returns the type of the constant expression.
func (e *ICmp) Type() types.Type {
	panic("not yet implemented")
}

// Ident returns the identifier associated with the constant expression.
func (e *ICmp) Ident() string {
	panic("not yet implemented")
}

// Simplify returns an equivalent (and potentially simplified) constant of the
// constant expression.
func (e *ICmp) Simplify() Constant {
	panic("not yet implemented")
}

// ~~~ [ fcmp ] ~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~

// FCmp is an LLVM IR fcmp expression.
type FCmp struct {
	// Floating-point comparison condition.
	Cond ll.FCond
	// Floating-point scalar or vector operands.
	X, Y Constant
}

// NewFCmp returns a new fcmp expression based on the given floating-point
// comparison condition and floating-point scalar or vector operands.
func NewFCmp(cond ll.FCond, x, y Constant) *FCmp {
	return &FCmp{Cond: cond, X: x, Y: y}
}

// Type returns the type of the constant expression.
func (e *FCmp) Type() types.Type {
	panic("not yet implemented")
}

// Ident returns the identifier associated with the constant expression.
func (e *FCmp) Ident() string {
	panic("not yet implemented")
}

// Simplify returns an equivalent (and potentially simplified) constant of the
// constant expression.
func (e *FCmp) Simplify() Constant {
	panic("not yet implemented")
}

// ~~~ [ select ] ~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~

// Select is an LLVM IR select expression.
type Select struct {
	// Selection condition.
	Cond Constant
	// Operands.
	X, Y Constant
}

// NewSelect returns a new select expression based on the given selection
// condition and operands.
func NewSelect(cond, x, y Constant) *Select {
	return &Select{Cond: cond, X: x, Y: x}
}

// Type returns the type of the constant expression.
func (e *Select) Type() types.Type {
	panic("not yet implemented")
}

// Ident returns the identifier associated with the constant expression.
func (e *Select) Ident() string {
	panic("not yet implemented")
}

// Simplify returns an equivalent (and potentially simplified) constant of the
// constant expression.
func (e *Select) Simplify() Constant {
	panic("not yet implemented")
}
