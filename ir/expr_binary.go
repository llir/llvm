package ir

import (
	"fmt"

	"github.com/llir/l/ir/types"
)

// --- [ Binary expressions ] --------------------------------------------------

// ~~~ [ add ] ~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~

// ExprAdd is an LLVM IR add expression.
type ExprAdd struct {
	// Operands.
	X, Y Constant // integer scalar or vector constants
}

// NewAddExpr returns a new add expression based on the given operands.
func NewAddExpr(x, y Constant) *ExprAdd {
	return &ExprAdd{X: x, Y: y}
}

// String returns the LLVM syntax representation of the constant expression as a
// type-value pair.
func (e *ExprAdd) String() string {
	return fmt.Sprintf("%v %v", e.Type(), e.Ident())
}

// Type returns the type of the constant expression.
func (e *ExprAdd) Type() types.Type {
	return e.X.Type()
}

// Ident returns the identifier associated with the constant expression.
func (e *ExprAdd) Ident() string {
	panic("not yet implemented")
}

// Simplify returns an equivalent (and potentially simplified) constant to the
// constant expression.
func (e *ExprAdd) Simplify() Constant {
	panic("not yet implemented")
}

// ~~~ [ fadd ] ~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~

// ExprFAdd is an LLVM IR fadd expression.
type ExprFAdd struct {
	// Operands.
	X, Y Constant // floating-point scalar or vector constants
}

// NewFAddExpr returns a new fadd expression based on the given operands.
func NewFAddExpr(x, y Constant) *ExprFAdd {
	return &ExprFAdd{X: x, Y: y}
}

// String returns the LLVM syntax representation of the constant expression as a
// type-value pair.
func (e *ExprFAdd) String() string {
	return fmt.Sprintf("%v %v", e.Type(), e.Ident())
}

// Type returns the type of the constant expression.
func (e *ExprFAdd) Type() types.Type {
	return e.X.Type()
}

// Ident returns the identifier associated with the constant expression.
func (e *ExprFAdd) Ident() string {
	panic("not yet implemented")
}

// Simplify returns an equivalent (and potentially simplified) constant to the
// constant expression.
func (e *ExprFAdd) Simplify() Constant {
	panic("not yet implemented")
}

// ~~~ [ sub ] ~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~

// ExprSub is an LLVM IR sub expression.
type ExprSub struct {
	// Operands.
	X, Y Constant // integer scalar or vector constants
}

// NewSubExpr returns a new sub expression based on the given operands.
func NewSubExpr(x, y Constant) *ExprSub {
	return &ExprSub{X: x, Y: y}
}

// String returns the LLVM syntax representation of the constant expression as a
// type-value pair.
func (e *ExprSub) String() string {
	return fmt.Sprintf("%v %v", e.Type(), e.Ident())
}

// Type returns the type of the constant expression.
func (e *ExprSub) Type() types.Type {
	return e.X.Type()
}

// Ident returns the identifier associated with the constant expression.
func (e *ExprSub) Ident() string {
	panic("not yet implemented")
}

// Simplify returns an equivalent (and potentially simplified) constant to the
// constant expression.
func (e *ExprSub) Simplify() Constant {
	panic("not yet implemented")
}

// ~~~ [ fsub ] ~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~

// ExprFSub is an LLVM IR fsub expression.
type ExprFSub struct {
	// Operands.
	X, Y Constant // floating-point scalar or vector constants
}

// NewFSubExpr returns a new fsub expression based on the given operands.
func NewFSubExpr(x, y Constant) *ExprFSub {
	return &ExprFSub{X: x, Y: y}
}

// String returns the LLVM syntax representation of the constant expression as a
// type-value pair.
func (e *ExprFSub) String() string {
	return fmt.Sprintf("%v %v", e.Type(), e.Ident())
}

// Type returns the type of the constant expression.
func (e *ExprFSub) Type() types.Type {
	return e.X.Type()
}

// Ident returns the identifier associated with the constant expression.
func (e *ExprFSub) Ident() string {
	panic("not yet implemented")
}

// Simplify returns an equivalent (and potentially simplified) constant to the
// constant expression.
func (e *ExprFSub) Simplify() Constant {
	panic("not yet implemented")
}

// ~~~ [ mul ] ~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~

// ExprMul is an LLVM IR mul expression.
type ExprMul struct {
	// Operands.
	X, Y Constant // integer scalar or vector constants
}

// NewMulExpr returns a new mul expression based on the given operands.
func NewMulExpr(x, y Constant) *ExprMul {
	return &ExprMul{X: x, Y: y}
}

// String returns the LLVM syntax representation of the constant expression as a
// type-value pair.
func (e *ExprMul) String() string {
	return fmt.Sprintf("%v %v", e.Type(), e.Ident())
}

// Type returns the type of the constant expression.
func (e *ExprMul) Type() types.Type {
	return e.X.Type()
}

// Ident returns the identifier associated with the constant expression.
func (e *ExprMul) Ident() string {
	panic("not yet implemented")
}

// Simplify returns an equivalent (and potentially simplified) constant to the
// constant expression.
func (e *ExprMul) Simplify() Constant {
	panic("not yet implemented")
}

// ~~~ [ fmul ] ~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~

// ExprFMul is an LLVM IR fmul expression.
type ExprFMul struct {
	// Operands.
	X, Y Constant // floating-point scalar or vector constants
}

// NewFMulExpr returns a new fmul expression based on the given operands.
func NewFMulExpr(x, y Constant) *ExprFMul {
	return &ExprFMul{X: x, Y: y}
}

// String returns the LLVM syntax representation of the constant expression as a
// type-value pair.
func (e *ExprFMul) String() string {
	return fmt.Sprintf("%v %v", e.Type(), e.Ident())
}

// Type returns the type of the constant expression.
func (e *ExprFMul) Type() types.Type {
	return e.X.Type()
}

// Ident returns the identifier associated with the constant expression.
func (e *ExprFMul) Ident() string {
	panic("not yet implemented")
}

// Simplify returns an equivalent (and potentially simplified) constant to the
// constant expression.
func (e *ExprFMul) Simplify() Constant {
	panic("not yet implemented")
}

// ~~~ [ udiv ] ~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~

// ExprUDiv is an LLVM IR udiv expression.
type ExprUDiv struct {
	// Operands.
	X, Y Constant // integer scalar or vector constants
}

// NewUDivExpr returns a new udiv expression based on the given operands.
func NewUDivExpr(x, y Constant) *ExprUDiv {
	return &ExprUDiv{X: x, Y: y}
}

// String returns the LLVM syntax representation of the constant expression as a
// type-value pair.
func (e *ExprUDiv) String() string {
	return fmt.Sprintf("%v %v", e.Type(), e.Ident())
}

// Type returns the type of the constant expression.
func (e *ExprUDiv) Type() types.Type {
	return e.X.Type()
}

// Ident returns the identifier associated with the constant expression.
func (e *ExprUDiv) Ident() string {
	panic("not yet implemented")
}

// Simplify returns an equivalent (and potentially simplified) constant to the
// constant expression.
func (e *ExprUDiv) Simplify() Constant {
	panic("not yet implemented")
}

// ~~~ [ sdiv ] ~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~

// ExprSDiv is an LLVM IR sdiv expression.
type ExprSDiv struct {
	// Operands.
	X, Y Constant // integer scalar or vector constants
}

// NewSDivExpr returns a new sdiv expression based on the given operands.
func NewSDivExpr(x, y Constant) *ExprSDiv {
	return &ExprSDiv{X: x, Y: y}
}

// String returns the LLVM syntax representation of the constant expression as a
// type-value pair.
func (e *ExprSDiv) String() string {
	return fmt.Sprintf("%v %v", e.Type(), e.Ident())
}

// Type returns the type of the constant expression.
func (e *ExprSDiv) Type() types.Type {
	return e.X.Type()
}

// Ident returns the identifier associated with the constant expression.
func (e *ExprSDiv) Ident() string {
	panic("not yet implemented")
}

// Simplify returns an equivalent (and potentially simplified) constant to the
// constant expression.
func (e *ExprSDiv) Simplify() Constant {
	panic("not yet implemented")
}

// ~~~ [ fdiv ] ~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~

// ExprFDiv is an LLVM IR fdiv expression.
type ExprFDiv struct {
	// Operands.
	X, Y Constant // floating-point scalar or vector constants
}

// NewFDivExpr returns a new fdiv expression based on the given operands.
func NewFDivExpr(x, y Constant) *ExprFDiv {
	return &ExprFDiv{X: x, Y: y}
}

// String returns the LLVM syntax representation of the constant expression as a
// type-value pair.
func (e *ExprFDiv) String() string {
	return fmt.Sprintf("%v %v", e.Type(), e.Ident())
}

// Type returns the type of the constant expression.
func (e *ExprFDiv) Type() types.Type {
	return e.X.Type()
}

// Ident returns the identifier associated with the constant expression.
func (e *ExprFDiv) Ident() string {
	panic("not yet implemented")
}

// Simplify returns an equivalent (and potentially simplified) constant to the
// constant expression.
func (e *ExprFDiv) Simplify() Constant {
	panic("not yet implemented")
}

// ~~~ [ urem ] ~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~

// ExprURem is an LLVM IR urem expression.
type ExprURem struct {
	// Operands.
	X, Y Constant // integer scalar or vector constants
}

// NewURemExpr returns a new urem expression based on the given operands.
func NewURemExpr(x, y Constant) *ExprURem {
	return &ExprURem{X: x, Y: y}
}

// String returns the LLVM syntax representation of the constant expression as a
// type-value pair.
func (e *ExprURem) String() string {
	return fmt.Sprintf("%v %v", e.Type(), e.Ident())
}

// Type returns the type of the constant expression.
func (e *ExprURem) Type() types.Type {
	return e.X.Type()
}

// Ident returns the identifier associated with the constant expression.
func (e *ExprURem) Ident() string {
	panic("not yet implemented")
}

// Simplify returns an equivalent (and potentially simplified) constant to the
// constant expression.
func (e *ExprURem) Simplify() Constant {
	panic("not yet implemented")
}

// ~~~ [ srem ] ~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~

// ExprSRem is an LLVM IR srem expression.
type ExprSRem struct {
	// Operands.
	X, Y Constant // integer scalar or vector constants
}

// NewSRemExpr returns a new srem expression based on the given operands.
func NewSRemExpr(x, y Constant) *ExprSRem {
	return &ExprSRem{X: x, Y: y}
}

// String returns the LLVM syntax representation of the constant expression as a
// type-value pair.
func (e *ExprSRem) String() string {
	return fmt.Sprintf("%v %v", e.Type(), e.Ident())
}

// Type returns the type of the constant expression.
func (e *ExprSRem) Type() types.Type {
	return e.X.Type()
}

// Ident returns the identifier associated with the constant expression.
func (e *ExprSRem) Ident() string {
	panic("not yet implemented")
}

// Simplify returns an equivalent (and potentially simplified) constant to the
// constant expression.
func (e *ExprSRem) Simplify() Constant {
	panic("not yet implemented")
}

// ~~~ [ frem ] ~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~

// ExprFRem is an LLVM IR frem expression.
type ExprFRem struct {
	// Operands.
	X, Y Constant // floating-point scalar or vector constants
}

// NewFRemExpr returns a new frem expression based on the given operands.
func NewFRemExpr(x, y Constant) *ExprFRem {
	return &ExprFRem{X: x, Y: y}
}

// String returns the LLVM syntax representation of the constant expression as a
// type-value pair.
func (e *ExprFRem) String() string {
	return fmt.Sprintf("%v %v", e.Type(), e.Ident())
}

// Type returns the type of the constant expression.
func (e *ExprFRem) Type() types.Type {
	return e.X.Type()
}

// Ident returns the identifier associated with the constant expression.
func (e *ExprFRem) Ident() string {
	panic("not yet implemented")
}

// Simplify returns an equivalent (and potentially simplified) constant to the
// constant expression.
func (e *ExprFRem) Simplify() Constant {
	panic("not yet implemented")
}
