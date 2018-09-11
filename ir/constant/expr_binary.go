package constant

import "github.com/llir/l/ir/types"

// --- [ Binary expressions ] --------------------------------------------------

// ~~~ [ add ] ~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~

// Add is an LLVM IR add expression.
type Add struct {
	// Operands.
	X, Y Constant // integer scalar or vector constants
}

// NewAdd returns a new add expression based on the given operands.
func NewAdd(x, y Constant) *Add {
	return &Add{X: x, Y: y}
}

// Type returns the type of the constant expression.
func (e *Add) Type() types.Type {
	return e.X.Type()
}

// Ident returns the identifier associated with the constant expression.
func (e *Add) Ident() string {
	panic("not yet implemented")
}

// Simplify returns an equivalent (and potentially simplified) constant of the
// constant expression.
func (e *Add) Simplify() Constant {
	panic("not yet implemented")
}

// ~~~ [ fadd ] ~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~

// FAdd is an LLVM IR fadd expression.
type FAdd struct {
	// Operands.
	X, Y Constant // floating-point scalar or vector constants
}

// NewFAdd returns a new fadd expression based on the given operands.
func NewFAdd(x, y Constant) *FAdd {
	return &FAdd{X: x, Y: y}
}

// Type returns the type of the constant expression.
func (e *FAdd) Type() types.Type {
	return e.X.Type()
}

// Ident returns the identifier associated with the constant expression.
func (e *FAdd) Ident() string {
	panic("not yet implemented")
}

// Simplify returns an equivalent (and potentially simplified) constant of the
// constant expression.
func (e *FAdd) Simplify() Constant {
	panic("not yet implemented")
}

// ~~~ [ sub ] ~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~

// Sub is an LLVM IR sub expression.
type Sub struct {
	// Operands.
	X, Y Constant // integer scalar or vector constants
}

// NewSub returns a new sub expression based on the given operands.
func NewSub(x, y Constant) *Sub {
	return &Sub{X: x, Y: y}
}

// Type returns the type of the constant expression.
func (e *Sub) Type() types.Type {
	return e.X.Type()
}

// Ident returns the identifier associated with the constant expression.
func (e *Sub) Ident() string {
	panic("not yet implemented")
}

// Simplify returns an equivalent (and potentially simplified) constant of the
// constant expression.
func (e *Sub) Simplify() Constant {
	panic("not yet implemented")
}

// ~~~ [ fsub ] ~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~

// FSub is an LLVM IR fsub expression.
type FSub struct {
	// Operands.
	X, Y Constant // floating-point scalar or vector constants
}

// NewFSub returns a new fsub expression based on the given operands.
func NewFSub(x, y Constant) *FSub {
	return &FSub{X: x, Y: y}
}

// Type returns the type of the constant expression.
func (e *FSub) Type() types.Type {
	return e.X.Type()
}

// Ident returns the identifier associated with the constant expression.
func (e *FSub) Ident() string {
	panic("not yet implemented")
}

// Simplify returns an equivalent (and potentially simplified) constant of the
// constant expression.
func (e *FSub) Simplify() Constant {
	panic("not yet implemented")
}

// ~~~ [ mul ] ~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~

// Mul is an LLVM IR mul expression.
type Mul struct {
	// Operands.
	X, Y Constant // integer scalar or vector constants
}

// NewMul returns a new mul expression based on the given operands.
func NewMul(x, y Constant) *Mul {
	return &Mul{X: x, Y: y}
}

// Type returns the type of the constant expression.
func (e *Mul) Type() types.Type {
	return e.X.Type()
}

// Ident returns the identifier associated with the constant expression.
func (e *Mul) Ident() string {
	panic("not yet implemented")
}

// Simplify returns an equivalent (and potentially simplified) constant of the
// constant expression.
func (e *Mul) Simplify() Constant {
	panic("not yet implemented")
}

// ~~~ [ fmul ] ~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~

// FMul is an LLVM IR fmul expression.
type FMul struct {
	// Operands.
	X, Y Constant // floating-point scalar or vector constants
}

// NewFMul returns a new fmul expression based on the given operands.
func NewFMul(x, y Constant) *FMul {
	return &FMul{X: x, Y: y}
}

// Type returns the type of the constant expression.
func (e *FMul) Type() types.Type {
	return e.X.Type()
}

// Ident returns the identifier associated with the constant expression.
func (e *FMul) Ident() string {
	panic("not yet implemented")
}

// Simplify returns an equivalent (and potentially simplified) constant of the
// constant expression.
func (e *FMul) Simplify() Constant {
	panic("not yet implemented")
}

// ~~~ [ udiv ] ~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~

// UDiv is an LLVM IR udiv expression.
type UDiv struct {
	// Operands.
	X, Y Constant // integer scalar or vector constants
}

// NewUDiv returns a new udiv expression based on the given operands.
func NewUDiv(x, y Constant) *UDiv {
	return &UDiv{X: x, Y: y}
}

// Type returns the type of the constant expression.
func (e *UDiv) Type() types.Type {
	return e.X.Type()
}

// Ident returns the identifier associated with the constant expression.
func (e *UDiv) Ident() string {
	panic("not yet implemented")
}

// Simplify returns an equivalent (and potentially simplified) constant of the
// constant expression.
func (e *UDiv) Simplify() Constant {
	panic("not yet implemented")
}

// ~~~ [ sdiv ] ~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~

// SDiv is an LLVM IR sdiv expression.
type SDiv struct {
	// Operands.
	X, Y Constant // integer scalar or vector constants
}

// NewSDiv returns a new sdiv expression based on the given operands.
func NewSDiv(x, y Constant) *SDiv {
	return &SDiv{X: x, Y: y}
}

// Type returns the type of the constant expression.
func (e *SDiv) Type() types.Type {
	return e.X.Type()
}

// Ident returns the identifier associated with the constant expression.
func (e *SDiv) Ident() string {
	panic("not yet implemented")
}

// Simplify returns an equivalent (and potentially simplified) constant of the
// constant expression.
func (e *SDiv) Simplify() Constant {
	panic("not yet implemented")
}

// ~~~ [ fdiv ] ~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~

// FDiv is an LLVM IR fdiv expression.
type FDiv struct {
	// Operands.
	X, Y Constant // floating-point scalar or vector constants
}

// NewFDiv returns a new fdiv expression based on the given operands.
func NewFDiv(x, y Constant) *FDiv {
	return &FDiv{X: x, Y: y}
}

// Type returns the type of the constant expression.
func (e *FDiv) Type() types.Type {
	return e.X.Type()
}

// Ident returns the identifier associated with the constant expression.
func (e *FDiv) Ident() string {
	panic("not yet implemented")
}

// Simplify returns an equivalent (and potentially simplified) constant of the
// constant expression.
func (e *FDiv) Simplify() Constant {
	panic("not yet implemented")
}

// ~~~ [ urem ] ~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~

// URem is an LLVM IR urem expression.
type URem struct {
	// Operands.
	X, Y Constant // integer scalar or vector constants
}

// NewURem returns a new urem expression based on the given operands.
func NewURem(x, y Constant) *URem {
	return &URem{X: x, Y: y}
}

// Type returns the type of the constant expression.
func (e *URem) Type() types.Type {
	return e.X.Type()
}

// Ident returns the identifier associated with the constant expression.
func (e *URem) Ident() string {
	panic("not yet implemented")
}

// Simplify returns an equivalent (and potentially simplified) constant of the
// constant expression.
func (e *URem) Simplify() Constant {
	panic("not yet implemented")
}

// ~~~ [ srem ] ~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~

// SRem is an LLVM IR srem expression.
type SRem struct {
	// Operands.
	X, Y Constant // integer scalar or vector constants
}

// NewSRem returns a new srem expression based on the given operands.
func NewSRem(x, y Constant) *SRem {
	return &SRem{X: x, Y: y}
}

// Type returns the type of the constant expression.
func (e *SRem) Type() types.Type {
	return e.X.Type()
}

// Ident returns the identifier associated with the constant expression.
func (e *SRem) Ident() string {
	panic("not yet implemented")
}

// Simplify returns an equivalent (and potentially simplified) constant of the
// constant expression.
func (e *SRem) Simplify() Constant {
	panic("not yet implemented")
}

// ~~~ [ frem ] ~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~

// FRem is an LLVM IR frem expression.
type FRem struct {
	// Operands.
	X, Y Constant // floating-point scalar or vector constants
}

// NewFRem returns a new frem expression based on the given operands.
func NewFRem(x, y Constant) *FRem {
	return &FRem{X: x, Y: y}
}

// Type returns the type of the constant expression.
func (e *FRem) Type() types.Type {
	return e.X.Type()
}

// Ident returns the identifier associated with the constant expression.
func (e *FRem) Ident() string {
	panic("not yet implemented")
}

// Simplify returns an equivalent (and potentially simplified) constant of the
// constant expression.
func (e *FRem) Simplify() Constant {
	panic("not yet implemented")
}
