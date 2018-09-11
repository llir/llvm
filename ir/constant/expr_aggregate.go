package constant

import "github.com/llir/l/ir/types"

// --- [ Aggregate expressions ] -----------------------------------------------

// ~~~ [ extractvalue ] ~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~

// ExtractValue is an LLVM IR extractvalue expression.
type ExtractValue struct {
	// Aggregate value.
	X Constant
	// Element indices.
	Indices []int64
}

// NewExtractValue returns a new extractvalue expression based on the given
// aggregate value and indicies.
func NewExtractValue(x Constant, indices ...int64) *ExtractValue {
	return &ExtractValue{X: x, Indices: indices}
}

// Type returns the type of the constant expression.
func (e *ExtractValue) Type() types.Type {
	panic("not yet implemented")
}

// Ident returns the identifier associated with the constant expression.
func (e *ExtractValue) Ident() string {
	panic("not yet implemented")
}

// Simplify returns an equivalent (and potentially simplified) constant of the
// constant expression.
func (e *ExtractValue) Simplify() Constant {
	panic("not yet implemented")
}

// ~~~ [ insertvalue ] ~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~

// InsertValue is an LLVM IR insertvalue expression.
type InsertValue struct {
	// Aggregate value.
	X Constant
	// Element to insert.
	Elem Constant
	// Element indices.
	Indices []int64
}

// NewInsertValue returns a new insertvalue expression based on the given
// aggregate value, element and indicies.
func NewInsertValue(x, elem Constant, indices ...int64) *InsertValue {
	return &InsertValue{X: x, Elem: elem, Indices: indices}
}

// Type returns the type of the constant expression.
func (e *InsertValue) Type() types.Type {
	panic("not yet implemented")
}

// Ident returns the identifier associated with the constant expression.
func (e *InsertValue) Ident() string {
	panic("not yet implemented")
}

// Simplify returns an equivalent (and potentially simplified) constant of the
// constant expression.
func (e *InsertValue) Simplify() Constant {
	panic("not yet implemented")
}
