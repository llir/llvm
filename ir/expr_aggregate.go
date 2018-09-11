package ir

import "github.com/llir/l/ir/types"

// --- [ Aggregate expressions ] -----------------------------------------------

// ~~~ [ extractvalue ] ~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~

// ExtractValueExpr is an LLVM IR extractvalue expression.
type ExtractValueExpr struct {
	// Aggregate value.
	X Constant
	// Element indices.
	Indices []int64
}

// NewExtractValueExpr returns a new extractvalue expression based on the given
// aggregate value and indicies.
func NewExtractValueExpr(x Constant, indices ...int64) *ExtractValueExpr {
	return &ExtractValueExpr{X: x, Indices: indices}
}

// Type returns the type of the constant expression.
func (e *ExtractValueExpr) Type() types.Type {
	panic("not yet implemented")
}

// Ident returns the identifier associated with the constant expression.
func (e *ExtractValueExpr) Ident() string {
	panic("not yet implemented")
}

// Simplify returns an equivalent (and potentially simplified) constant of the
// constant expression.
func (e *ExtractValueExpr) Simplify() Constant {
	panic("not yet implemented")
}

// ~~~ [ insertvalue ] ~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~

// InsertValueExpr is an LLVM IR insertvalue expression.
type InsertValueExpr struct {
	// Aggregate value.
	X Constant
	// Element to insert.
	Elem Constant
	// Element indices.
	Indices []int64
}

// NewInsertValueExpr returns a new insertvalue expression based on the given
// aggregate value, element and indicies.
func NewInsertValueExpr(x, elem Constant, indices ...int64) *InsertValueExpr {
	return &InsertValueExpr{X: x, Elem: elem, Indices: indices}
}

// Type returns the type of the constant expression.
func (e *InsertValueExpr) Type() types.Type {
	panic("not yet implemented")
}

// Ident returns the identifier associated with the constant expression.
func (e *InsertValueExpr) Ident() string {
	panic("not yet implemented")
}

// Simplify returns an equivalent (and potentially simplified) constant of the
// constant expression.
func (e *InsertValueExpr) Simplify() Constant {
	panic("not yet implemented")
}
