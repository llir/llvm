package ir

import (
	"fmt"
	"strings"

	"github.com/llir/llvm/ir/types"
)

// --- [ Aggregate expressions ] -----------------------------------------------

// ~~~ [ extractvalue ] ~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~

// ExprExtractValue is an LLVM IR extractvalue expression.
type ExprExtractValue struct {
	// Aggregate value.
	X Constant
	// Element indices.
	Indices []int64
}

// NewExtractValueExpr returns a new extractvalue expression based on the given
// aggregate value and indicies.
func NewExtractValueExpr(x Constant, indices ...int64) *ExprExtractValue {
	return &ExprExtractValue{X: x, Indices: indices}
}

// String returns the LLVM syntax representation of the constant expression as a
// type-value pair.
func (e *ExprExtractValue) String() string {
	return fmt.Sprintf("%v %v", e.Type(), e.Ident())
}

// Type returns the type of the constant expression.
func (e *ExprExtractValue) Type() types.Type {
	panic("not yet implemented")
}

// Ident returns the identifier associated with the constant expression.
func (e *ExprExtractValue) Ident() string {
	// "extractvalue" "(" Type Constant Indices ")"
	buf := &strings.Builder{}
	fmt.Fprintf(buf, "extractvalue (%v", e.X)
	for _, index := range e.Indices {
		fmt.Fprintf(buf, ", %v", index)
	}
	buf.WriteString(")")
	return buf.String()
}

// Simplify returns an equivalent (and potentially simplified) constant to the
// constant expression.
func (e *ExprExtractValue) Simplify() Constant {
	panic("not yet implemented")
}

// ~~~ [ insertvalue ] ~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~

// ExprInsertValue is an LLVM IR insertvalue expression.
type ExprInsertValue struct {
	// Aggregate value.
	X Constant
	// Element to insert.
	Elem Constant
	// Element indices.
	Indices []int64
}

// NewInsertValueExpr returns a new insertvalue expression based on the given
// aggregate value, element and indicies.
func NewInsertValueExpr(x, elem Constant, indices ...int64) *ExprInsertValue {
	return &ExprInsertValue{X: x, Elem: elem, Indices: indices}
}

// String returns the LLVM syntax representation of the constant expression as a
// type-value pair.
func (e *ExprInsertValue) String() string {
	return fmt.Sprintf("%v %v", e.Type(), e.Ident())
}

// Type returns the type of the constant expression.
func (e *ExprInsertValue) Type() types.Type {
	panic("not yet implemented")
}

// Ident returns the identifier associated with the constant expression.
func (e *ExprInsertValue) Ident() string {
	// "insertvalue" "(" Type Constant "," Type Constant Indices ")"
	buf := &strings.Builder{}
	fmt.Fprintf(buf, "insertvalue (%v, %v", e.X, e.Elem)
	for _, index := range e.Indices {
		fmt.Fprintf(buf, ", %v", index)
	}
	buf.WriteString(")")
	return buf.String()
}

// Simplify returns an equivalent (and potentially simplified) constant to the
// constant expression.
func (e *ExprInsertValue) Simplify() Constant {
	panic("not yet implemented")
}
