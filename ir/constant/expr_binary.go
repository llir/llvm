package constant

import (
	"fmt"
	"strings"

	"github.com/llir/llvm/ir/enum"
	"github.com/llir/llvm/ir/types"
)

// --- [ Binary expressions ] --------------------------------------------------

// ~~~ [ add ] ~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~

// ExprAdd is an LLVM IR add expression.
type ExprAdd struct {
	// Operands.
	X, Y Constant // integer scalar or vector constants

	// extra.

	// Type of result produced by the constant expression.
	Typ types.Type
	// (optional) Integer overflow flags.
	OverflowFlags []enum.OverflowFlag
}

// NewAdd returns a new add expression based on the given operands.
func NewAdd(x, y Constant) *ExprAdd {
	e := &ExprAdd{X: x, Y: y}
	// Compute type.
	e.Type()
	return e
}

// String returns the LLVM syntax representation of the constant expression as a
// type-value pair.
func (e *ExprAdd) String() string {
	return fmt.Sprintf("%s %s", e.Type(), e.Ident())
}

// Type returns the type of the constant expression.
func (e *ExprAdd) Type() types.Type {
	// Cache type if not present.
	if e.Typ == nil {
		e.Typ = e.X.Type()
	}
	return e.Typ
}

// Ident returns the identifier associated with the constant expression.
func (e *ExprAdd) Ident() string {
	// 'add' OverflowFlags=OverflowFlag* '(' X=TypeConst ',' Y=TypeConst ')'
	buf := &strings.Builder{}
	buf.WriteString("add")
	for _, flag := range e.OverflowFlags {
		fmt.Fprintf(buf, " %s", flag)
	}
	fmt.Fprintf(buf, " (%s, %s)", e.X, e.Y)
	return buf.String()
}

// ~~~ [ sub ] ~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~

// ExprSub is an LLVM IR sub expression.
type ExprSub struct {
	// Operands.
	X, Y Constant // integer scalar or vector constants

	// extra.

	// Type of result produced by the constant expression.
	Typ types.Type
	// (optional) Integer overflow flags.
	OverflowFlags []enum.OverflowFlag
}

// NewSub returns a new sub expression based on the given operands.
func NewSub(x, y Constant) *ExprSub {
	e := &ExprSub{X: x, Y: y}
	// Compute type.
	e.Type()
	return e
}

// String returns the LLVM syntax representation of the constant expression as a
// type-value pair.
func (e *ExprSub) String() string {
	return fmt.Sprintf("%s %s", e.Type(), e.Ident())
}

// Type returns the type of the constant expression.
func (e *ExprSub) Type() types.Type {
	// Cache type if not present.
	if e.Typ == nil {
		e.Typ = e.X.Type()
	}
	return e.Typ
}

// Ident returns the identifier associated with the constant expression.
func (e *ExprSub) Ident() string {
	// 'sub' OverflowFlags=OverflowFlag* '(' X=TypeConst ',' Y=TypeConst ')'
	buf := &strings.Builder{}
	buf.WriteString("sub")
	for _, flag := range e.OverflowFlags {
		fmt.Fprintf(buf, " %s", flag)
	}
	fmt.Fprintf(buf, " (%s, %s)", e.X, e.Y)
	return buf.String()
}

// ~~~ [ mul ] ~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~

// ExprMul is an LLVM IR mul expression.
type ExprMul struct {
	// Operands.
	X, Y Constant // integer scalar or vector constants

	// extra.

	// Type of result produced by the constant expression.
	Typ types.Type
	// (optional) Integer overflow flags.
	OverflowFlags []enum.OverflowFlag
}

// NewMul returns a new mul expression based on the given operands.
func NewMul(x, y Constant) *ExprMul {
	e := &ExprMul{X: x, Y: y}
	// Compute type.
	e.Type()
	return e
}

// String returns the LLVM syntax representation of the constant expression as a
// type-value pair.
func (e *ExprMul) String() string {
	return fmt.Sprintf("%s %s", e.Type(), e.Ident())
}

// Type returns the type of the constant expression.
func (e *ExprMul) Type() types.Type {
	// Cache type if not present.
	if e.Typ == nil {
		e.Typ = e.X.Type()
	}
	return e.Typ
}

// Ident returns the identifier associated with the constant expression.
func (e *ExprMul) Ident() string {
	// 'mul' OverflowFlags=OverflowFlag* '(' X=TypeConst ',' Y=TypeConst ')'
	buf := &strings.Builder{}
	buf.WriteString("mul")
	for _, flag := range e.OverflowFlags {
		fmt.Fprintf(buf, " %s", flag)
	}
	fmt.Fprintf(buf, " (%s, %s)", e.X, e.Y)
	return buf.String()
}
