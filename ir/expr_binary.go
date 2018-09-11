package ir

import "github.com/llir/l/ir/types"

// --- [ Binary expressions ] --------------------------------------------------

// ~~~ [ add ] ~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~

// AddExpr is an LLVM IR add expression.
type AddExpr struct {
	// Operands.
	X, Y Constant // integer scalar or vector constants
}

// NewAddExpr returns a new add expression based on the given operands.
func NewAddExpr(x, y Constant) *AddExpr {
	return &AddExpr{X: x, Y: y}
}

// Type returns the type of the constant expression.
func (e *AddExpr) Type() types.Type {
	return e.X.Type()
}

// Ident returns the identifier associated with the constant expression.
func (e *AddExpr) Ident() string {
	panic("not yet implemented")
}

// Simplify returns an equivalent (and potentially simplified) constant of the
// constant expression.
func (e *AddExpr) Simplify() Constant {
	panic("not yet implemented")
}

// ~~~ [ fadd ] ~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~

// FAddExpr is an LLVM IR fadd expression.
type FAddExpr struct {
	// Operands.
	X, Y Constant // floating-point scalar or vector constants
}

// NewFAddExpr returns a new fadd expression based on the given operands.
func NewFAddExpr(x, y Constant) *FAddExpr {
	return &FAddExpr{X: x, Y: y}
}

// Type returns the type of the constant expression.
func (e *FAddExpr) Type() types.Type {
	return e.X.Type()
}

// Ident returns the identifier associated with the constant expression.
func (e *FAddExpr) Ident() string {
	panic("not yet implemented")
}

// Simplify returns an equivalent (and potentially simplified) constant of the
// constant expression.
func (e *FAddExpr) Simplify() Constant {
	panic("not yet implemented")
}

// ~~~ [ sub ] ~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~

// SubExpr is an LLVM IR sub expression.
type SubExpr struct {
	// Operands.
	X, Y Constant // integer scalar or vector constants
}

// NewSubExpr returns a new sub expression based on the given operands.
func NewSubExpr(x, y Constant) *SubExpr {
	return &SubExpr{X: x, Y: y}
}

// Type returns the type of the constant expression.
func (e *SubExpr) Type() types.Type {
	return e.X.Type()
}

// Ident returns the identifier associated with the constant expression.
func (e *SubExpr) Ident() string {
	panic("not yet implemented")
}

// Simplify returns an equivalent (and potentially simplified) constant of the
// constant expression.
func (e *SubExpr) Simplify() Constant {
	panic("not yet implemented")
}

// ~~~ [ fsub ] ~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~

// FSubExpr is an LLVM IR fsub expression.
type FSubExpr struct {
	// Operands.
	X, Y Constant // floating-point scalar or vector constants
}

// NewFSubExpr returns a new fsub expression based on the given operands.
func NewFSubExpr(x, y Constant) *FSubExpr {
	return &FSubExpr{X: x, Y: y}
}

// Type returns the type of the constant expression.
func (e *FSubExpr) Type() types.Type {
	return e.X.Type()
}

// Ident returns the identifier associated with the constant expression.
func (e *FSubExpr) Ident() string {
	panic("not yet implemented")
}

// Simplify returns an equivalent (and potentially simplified) constant of the
// constant expression.
func (e *FSubExpr) Simplify() Constant {
	panic("not yet implemented")
}

// ~~~ [ mul ] ~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~

// MulExpr is an LLVM IR mul expression.
type MulExpr struct {
	// Operands.
	X, Y Constant // integer scalar or vector constants
}

// NewMulExpr returns a new mul expression based on the given operands.
func NewMulExpr(x, y Constant) *MulExpr {
	return &MulExpr{X: x, Y: y}
}

// Type returns the type of the constant expression.
func (e *MulExpr) Type() types.Type {
	return e.X.Type()
}

// Ident returns the identifier associated with the constant expression.
func (e *MulExpr) Ident() string {
	panic("not yet implemented")
}

// Simplify returns an equivalent (and potentially simplified) constant of the
// constant expression.
func (e *MulExpr) Simplify() Constant {
	panic("not yet implemented")
}

// ~~~ [ fmul ] ~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~

// FMulExpr is an LLVM IR fmul expression.
type FMulExpr struct {
	// Operands.
	X, Y Constant // floating-point scalar or vector constants
}

// NewFMulExpr returns a new fmul expression based on the given operands.
func NewFMulExpr(x, y Constant) *FMulExpr {
	return &FMulExpr{X: x, Y: y}
}

// Type returns the type of the constant expression.
func (e *FMulExpr) Type() types.Type {
	return e.X.Type()
}

// Ident returns the identifier associated with the constant expression.
func (e *FMulExpr) Ident() string {
	panic("not yet implemented")
}

// Simplify returns an equivalent (and potentially simplified) constant of the
// constant expression.
func (e *FMulExpr) Simplify() Constant {
	panic("not yet implemented")
}

// ~~~ [ udiv ] ~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~

// UDivExpr is an LLVM IR udiv expression.
type UDivExpr struct {
	// Operands.
	X, Y Constant // integer scalar or vector constants
}

// NewUDivExpr returns a new udiv expression based on the given operands.
func NewUDivExpr(x, y Constant) *UDivExpr {
	return &UDivExpr{X: x, Y: y}
}

// Type returns the type of the constant expression.
func (e *UDivExpr) Type() types.Type {
	return e.X.Type()
}

// Ident returns the identifier associated with the constant expression.
func (e *UDivExpr) Ident() string {
	panic("not yet implemented")
}

// Simplify returns an equivalent (and potentially simplified) constant of the
// constant expression.
func (e *UDivExpr) Simplify() Constant {
	panic("not yet implemented")
}

// ~~~ [ sdiv ] ~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~

// SDivExpr is an LLVM IR sdiv expression.
type SDivExpr struct {
	// Operands.
	X, Y Constant // integer scalar or vector constants
}

// NewSDivExpr returns a new sdiv expression based on the given operands.
func NewSDivExpr(x, y Constant) *SDivExpr {
	return &SDivExpr{X: x, Y: y}
}

// Type returns the type of the constant expression.
func (e *SDivExpr) Type() types.Type {
	return e.X.Type()
}

// Ident returns the identifier associated with the constant expression.
func (e *SDivExpr) Ident() string {
	panic("not yet implemented")
}

// Simplify returns an equivalent (and potentially simplified) constant of the
// constant expression.
func (e *SDivExpr) Simplify() Constant {
	panic("not yet implemented")
}

// ~~~ [ fdiv ] ~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~

// FDivExpr is an LLVM IR fdiv expression.
type FDivExpr struct {
	// Operands.
	X, Y Constant // floating-point scalar or vector constants
}

// NewFDivExpr returns a new fdiv expression based on the given operands.
func NewFDivExpr(x, y Constant) *FDivExpr {
	return &FDivExpr{X: x, Y: y}
}

// Type returns the type of the constant expression.
func (e *FDivExpr) Type() types.Type {
	return e.X.Type()
}

// Ident returns the identifier associated with the constant expression.
func (e *FDivExpr) Ident() string {
	panic("not yet implemented")
}

// Simplify returns an equivalent (and potentially simplified) constant of the
// constant expression.
func (e *FDivExpr) Simplify() Constant {
	panic("not yet implemented")
}

// ~~~ [ urem ] ~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~

// URemExpr is an LLVM IR urem expression.
type URemExpr struct {
	// Operands.
	X, Y Constant // integer scalar or vector constants
}

// NewURemExpr returns a new urem expression based on the given operands.
func NewURemExpr(x, y Constant) *URemExpr {
	return &URemExpr{X: x, Y: y}
}

// Type returns the type of the constant expression.
func (e *URemExpr) Type() types.Type {
	return e.X.Type()
}

// Ident returns the identifier associated with the constant expression.
func (e *URemExpr) Ident() string {
	panic("not yet implemented")
}

// Simplify returns an equivalent (and potentially simplified) constant of the
// constant expression.
func (e *URemExpr) Simplify() Constant {
	panic("not yet implemented")
}

// ~~~ [ srem ] ~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~

// SRemExpr is an LLVM IR srem expression.
type SRemExpr struct {
	// Operands.
	X, Y Constant // integer scalar or vector constants
}

// NewSRemExpr returns a new srem expression based on the given operands.
func NewSRemExpr(x, y Constant) *SRemExpr {
	return &SRemExpr{X: x, Y: y}
}

// Type returns the type of the constant expression.
func (e *SRemExpr) Type() types.Type {
	return e.X.Type()
}

// Ident returns the identifier associated with the constant expression.
func (e *SRemExpr) Ident() string {
	panic("not yet implemented")
}

// Simplify returns an equivalent (and potentially simplified) constant of the
// constant expression.
func (e *SRemExpr) Simplify() Constant {
	panic("not yet implemented")
}

// ~~~ [ frem ] ~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~

// FRemExpr is an LLVM IR frem expression.
type FRemExpr struct {
	// Operands.
	X, Y Constant // floating-point scalar or vector constants
}

// NewFRemExpr returns a new frem expression based on the given operands.
func NewFRemExpr(x, y Constant) *FRemExpr {
	return &FRemExpr{X: x, Y: y}
}

// Type returns the type of the constant expression.
func (e *FRemExpr) Type() types.Type {
	return e.X.Type()
}

// Ident returns the identifier associated with the constant expression.
func (e *FRemExpr) Ident() string {
	panic("not yet implemented")
}

// Simplify returns an equivalent (and potentially simplified) constant of the
// constant expression.
func (e *FRemExpr) Simplify() Constant {
	panic("not yet implemented")
}
