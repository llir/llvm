package ir

import "github.com/llir/l/ir/types"

// --- [ Bitwise expressions ] -------------------------------------------------

// ~~~ [ shl ] ~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~

// ShlExpr is an LLVM IR shl expression.
type ShlExpr struct {
	// Operands.
	X, Y Constant // integer scalars or vectors
}

// NewShlExpr returns a new shl expression based on the given operands.
func NewShlExpr(x, y Constant) *ShlExpr {
	return &ShlExpr{X: x, Y: y}
}

// Type returns the type of the constant expression.
func (e *ShlExpr) Type() types.Type {
	return e.X.Type()
}

// Ident returns the identifier associated with the constant expression.
func (e *ShlExpr) Ident() string {
	panic("not yet implemented")
}

// Simplify returns an equivalent (and potentially simplified) constant of the
// constant expression.
func (e *ShlExpr) Simplify() Constant {
	panic("not yet implemented")
}

// ~~~ [ lshr ] ~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~

// LShrExpr is an LLVM IR lshr expression.
type LShrExpr struct {
	// Operands.
	X, Y Constant // integer scalars or vectors
}

// NewLShrExpr returns a new lshr expression based on the given operands.
func NewLShrExpr(x, y Constant) *LShrExpr {
	return &LShrExpr{X: x, Y: y}
}

// Type returns the type of the constant expression.
func (e *LShrExpr) Type() types.Type {
	return e.X.Type()
}

// Ident returns the identifier associated with the constant expression.
func (e *LShrExpr) Ident() string {
	panic("not yet implemented")
}

// Simplify returns an equivalent (and potentially simplified) constant of the
// constant expression.
func (e *LShrExpr) Simplify() Constant {
	panic("not yet implemented")
}

// ~~~ [ ashr ] ~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~

// AShrExpr is an LLVM IR ashr expression.
type AShrExpr struct {
	// Operands.
	X, Y Constant // integer scalars or vectors
}

// NewAShrExpr returns a new ashr expression based on the given operands.
func NewAShrExpr(x, y Constant) *AShrExpr {
	return &AShrExpr{X: x, Y: y}
}

// Type returns the type of the constant expression.
func (e *AShrExpr) Type() types.Type {
	return e.X.Type()
}

// Ident returns the identifier associated with the constant expression.
func (e *AShrExpr) Ident() string {
	panic("not yet implemented")
}

// Simplify returns an equivalent (and potentially simplified) constant of the
// constant expression.
func (e *AShrExpr) Simplify() Constant {
	panic("not yet implemented")
}

// ~~~ [ and ] ~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~

// AndExpr is an LLVM IR and expression.
type AndExpr struct {
	// Operands.
	X, Y Constant // integer scalars or vectors
}

// NewAndExpr returns a new and expression based on the given operands.
func NewAndExpr(x, y Constant) *AndExpr {
	return &AndExpr{X: x, Y: y}
}

// Type returns the type of the constant expression.
func (e *AndExpr) Type() types.Type {
	return e.X.Type()
}

// Ident returns the identifier associated with the constant expression.
func (e *AndExpr) Ident() string {
	panic("not yet implemented")
}

// Simplify returns an equivalent (and potentially simplified) constant of the
// constant expression.
func (e *AndExpr) Simplify() Constant {
	panic("not yet implemented")
}

// ~~~ [ or ] ~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~

// OrExpr is an LLVM IR or expression.
type OrExpr struct {
	// Operands.
	X, Y Constant // integer scalars or vectors
}

// NewOrExpr returns a new or expression based on the given operands.
func NewOrExpr(x, y Constant) *OrExpr {
	return &OrExpr{X: x, Y: y}
}

// Type returns the type of the constant expression.
func (e *OrExpr) Type() types.Type {
	return e.X.Type()
}

// Ident returns the identifier associated with the constant expression.
func (e *OrExpr) Ident() string {
	panic("not yet implemented")
}

// Simplify returns an equivalent (and potentially simplified) constant of the
// constant expression.
func (e *OrExpr) Simplify() Constant {
	panic("not yet implemented")
}

// ~~~ [ xor ] ~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~

// XorExpr is an LLVM IR xor expression.
type XorExpr struct {
	// Operands.
	X, Y Constant // integer scalars or vectors
}

// NewXorExpr returns a new xor expression based on the given operands.
func NewXorExpr(x, y Constant) *XorExpr {
	return &XorExpr{X: x, Y: y}
}

// Type returns the type of the constant expression.
func (e *XorExpr) Type() types.Type {
	return e.X.Type()
}

// Ident returns the identifier associated with the constant expression.
func (e *XorExpr) Ident() string {
	panic("not yet implemented")
}

// Simplify returns an equivalent (and potentially simplified) constant of the
// constant expression.
func (e *XorExpr) Simplify() Constant {
	panic("not yet implemented")
}
