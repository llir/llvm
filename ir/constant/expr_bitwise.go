package constant

import (
	"fmt"
	"strings"

	"github.com/llir/llvm/ir/enum"
	"github.com/llir/llvm/ir/types"
)

// --- [ Bitwise expressions ] -------------------------------------------------

// ~~~ [ shl ] ~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~

// ExprShl is an LLVM IR shl expression.
type ExprShl struct {
	// Operands.
	X, Y Constant // integer scalars or vectors

	// extra.

	// (optional) Integer overflow flags.
	OverflowFlags []enum.OverflowFlag
}

// NewShl returns a new shl expression based on the given operands.
func NewShl(x, y Constant) *ExprShl {
	return &ExprShl{X: x, Y: y}
}

// String returns the LLVM syntax representation of the constant expression as a
// type-value pair.
func (e *ExprShl) String() string {
	return fmt.Sprintf("%v %v", e.Type(), e.Ident())
}

// Type returns the type of the constant expression.
func (e *ExprShl) Type() types.Type {
	return e.X.Type()
}

// Ident returns the identifier associated with the constant expression.
func (e *ExprShl) Ident() string {
	// "shl" OverflowFlags "(" Type Constant "," Type Constant ")"
	buf := &strings.Builder{}
	buf.WriteString("shl")
	for _, flag := range e.OverflowFlags {
		fmt.Fprintf(buf, " %v", flag)
	}
	fmt.Fprintf(buf, " (%v, %v)", e.X, e.Y)
	return buf.String()
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

	// extra.

	// (optional) The result is a poison value if any of the bits shifted out are
	// non-zero.
	Exact bool
}

// NewLShr returns a new lshr expression based on the given operands.
func NewLShr(x, y Constant) *ExprLShr {
	return &ExprLShr{X: x, Y: y}
}

// String returns the LLVM syntax representation of the constant expression as a
// type-value pair.
func (e *ExprLShr) String() string {
	return fmt.Sprintf("%v %v", e.Type(), e.Ident())
}

// Type returns the type of the constant expression.
func (e *ExprLShr) Type() types.Type {
	return e.X.Type()
}

// Ident returns the identifier associated with the constant expression.
func (e *ExprLShr) Ident() string {
	// "lshr" OptExact "(" Type Constant "," Type Constant ")"
	buf := &strings.Builder{}
	buf.WriteString("lshr")
	if e.Exact {
		buf.WriteString(" exact")
	}
	fmt.Fprintf(buf, " (%v, %v)", e.X, e.Y)
	return buf.String()
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

	// extra.

	// (optional) The result is a poison value if any of the bits shifted out are
	// non-zero.
	Exact bool
}

// NewAShr returns a new ashr expression based on the given operands.
func NewAShr(x, y Constant) *ExprAShr {
	return &ExprAShr{X: x, Y: y}
}

// String returns the LLVM syntax representation of the constant expression as a
// type-value pair.
func (e *ExprAShr) String() string {
	return fmt.Sprintf("%v %v", e.Type(), e.Ident())
}

// Type returns the type of the constant expression.
func (e *ExprAShr) Type() types.Type {
	return e.X.Type()
}

// Ident returns the identifier associated with the constant expression.
func (e *ExprAShr) Ident() string {
	// "ashr" OptExact "(" Type Constant "," Type Constant ")"
	buf := &strings.Builder{}
	buf.WriteString("ashr")
	if e.Exact {
		buf.WriteString(" exact")
	}
	fmt.Fprintf(buf, " (%v, %v)", e.X, e.Y)
	return buf.String()
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

// NewAnd returns a new and expression based on the given operands.
func NewAnd(x, y Constant) *ExprAnd {
	return &ExprAnd{X: x, Y: y}
}

// String returns the LLVM syntax representation of the constant expression as a
// type-value pair.
func (e *ExprAnd) String() string {
	return fmt.Sprintf("%v %v", e.Type(), e.Ident())
}

// Type returns the type of the constant expression.
func (e *ExprAnd) Type() types.Type {
	return e.X.Type()
}

// Ident returns the identifier associated with the constant expression.
func (e *ExprAnd) Ident() string {
	// "and" "(" Type Constant "," Type Constant ")"
	return fmt.Sprintf("and (%v, %v)", e.X, e.Y)
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

// NewOr returns a new or expression based on the given operands.
func NewOr(x, y Constant) *ExprOr {
	return &ExprOr{X: x, Y: y}
}

// String returns the LLVM syntax representation of the constant expression as a
// type-value pair.
func (e *ExprOr) String() string {
	return fmt.Sprintf("%v %v", e.Type(), e.Ident())
}

// Type returns the type of the constant expression.
func (e *ExprOr) Type() types.Type {
	return e.X.Type()
}

// Ident returns the identifier associated with the constant expression.
func (e *ExprOr) Ident() string {
	// "or" "(" Type Constant "," Type Constant ")"
	return fmt.Sprintf("or (%v, %v)", e.X, e.Y)
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

// NewXor returns a new xor expression based on the given operands.
func NewXor(x, y Constant) *ExprXor {
	return &ExprXor{X: x, Y: y}
}

// String returns the LLVM syntax representation of the constant expression as a
// type-value pair.
func (e *ExprXor) String() string {
	return fmt.Sprintf("%v %v", e.Type(), e.Ident())
}

// Type returns the type of the constant expression.
func (e *ExprXor) Type() types.Type {
	return e.X.Type()
}

// Ident returns the identifier associated with the constant expression.
func (e *ExprXor) Ident() string {
	// "xor" "(" Type Constant "," Type Constant ")"
	return fmt.Sprintf("xor (%v, %v)", e.X, e.Y)
}

// Simplify returns an equivalent (and potentially simplified) constant to the
// constant expression.
func (e *ExprXor) Simplify() Constant {
	panic("not yet implemented")
}
