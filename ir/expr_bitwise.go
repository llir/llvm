package ir

import "github.com/llir/l/ir/types"

// --- [ Bitwise expressions ] -------------------------------------------------

// ~~~ [ shl ] ~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~

// ExprShl is an LLVM IR shl expression.
type ExprShl struct {
	// Operands.
	X, Y Constant // integer scalars or vectors
}

// NewShlExpr returns a new shl expression based on the given operands.
func NewShlExpr(x, y Constant) *ExprShl {
	return &ExprShl{X: x, Y: y}
}

// Type returns the type of the constant expression.
func (e *ExprShl) Type() types.Type {
	return e.X.Type()
}

// Ident returns the identifier associated with the constant expression.
func (e *ExprShl) Ident() string {
	panic("not yet implemented")
}

// Simplify returns an equivalent (and potentially simplified) constant to the
// constant expression.
func (e *ExprShl) Simplify() Constant {
	panic("not yet implemented")
}

// ~~~ [ lshr ] ~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~

// ExprLShr is an LLVM IR lshr expression.
type ExprLShr struct {
	// Operands.
	X, Y Constant // integer scalars or vectors
}

// NewLShrExpr returns a new lshr expression based on the given operands.
func NewLShrExpr(x, y Constant) *ExprLShr {
	return &ExprLShr{X: x, Y: y}
}

// Type returns the type of the constant expression.
func (e *ExprLShr) Type() types.Type {
	return e.X.Type()
}

// Ident returns the identifier associated with the constant expression.
func (e *ExprLShr) Ident() string {
	panic("not yet implemented")
}

// Simplify returns an equivalent (and potentially simplified) constant to the
// constant expression.
func (e *ExprLShr) Simplify() Constant {
	panic("not yet implemented")
}

// ~~~ [ ashr ] ~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~

// ExprAShr is an LLVM IR ashr expression.
type ExprAShr struct {
	// Operands.
	X, Y Constant // integer scalars or vectors
}

// NewAShrExpr returns a new ashr expression based on the given operands.
func NewAShrExpr(x, y Constant) *ExprAShr {
	return &ExprAShr{X: x, Y: y}
}

// Type returns the type of the constant expression.
func (e *ExprAShr) Type() types.Type {
	return e.X.Type()
}

// Ident returns the identifier associated with the constant expression.
func (e *ExprAShr) Ident() string {
	panic("not yet implemented")
}

// Simplify returns an equivalent (and potentially simplified) constant to the
// constant expression.
func (e *ExprAShr) Simplify() Constant {
	panic("not yet implemented")
}

// ~~~ [ and ] ~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~

// ExprAnd is an LLVM IR and expression.
type ExprAnd struct {
	// Operands.
	X, Y Constant // integer scalars or vectors
}

// NewAndExpr returns a new and expression based on the given operands.
func NewAndExpr(x, y Constant) *ExprAnd {
	return &ExprAnd{X: x, Y: y}
}

// Type returns the type of the constant expression.
func (e *ExprAnd) Type() types.Type {
	return e.X.Type()
}

// Ident returns the identifier associated with the constant expression.
func (e *ExprAnd) Ident() string {
	panic("not yet implemented")
}

// Simplify returns an equivalent (and potentially simplified) constant to the
// constant expression.
func (e *ExprAnd) Simplify() Constant {
	panic("not yet implemented")
}

// ~~~ [ or ] ~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~

// ExprOr is an LLVM IR or expression.
type ExprOr struct {
	// Operands.
	X, Y Constant // integer scalars or vectors
}

// NewOrExpr returns a new or expression based on the given operands.
func NewOrExpr(x, y Constant) *ExprOr {
	return &ExprOr{X: x, Y: y}
}

// Type returns the type of the constant expression.
func (e *ExprOr) Type() types.Type {
	return e.X.Type()
}

// Ident returns the identifier associated with the constant expression.
func (e *ExprOr) Ident() string {
	panic("not yet implemented")
}

// Simplify returns an equivalent (and potentially simplified) constant to the
// constant expression.
func (e *ExprOr) Simplify() Constant {
	panic("not yet implemented")
}

// ~~~ [ xor ] ~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~

// ExprXor is an LLVM IR xor expression.
type ExprXor struct {
	// Operands.
	X, Y Constant // integer scalars or vectors
}

// NewXorExpr returns a new xor expression based on the given operands.
func NewXorExpr(x, y Constant) *ExprXor {
	return &ExprXor{X: x, Y: y}
}

// Type returns the type of the constant expression.
func (e *ExprXor) Type() types.Type {
	return e.X.Type()
}

// Ident returns the identifier associated with the constant expression.
func (e *ExprXor) Ident() string {
	panic("not yet implemented")
}

// Simplify returns an equivalent (and potentially simplified) constant to the
// constant expression.
func (e *ExprXor) Simplify() Constant {
	panic("not yet implemented")
}
