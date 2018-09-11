package ir

import "github.com/llir/l/ir/types"

// --- [ Vector expressions ] --------------------------------------------------

// ~~~ [ extractelement ] ~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~

// ExtractElementExpr is an LLVM IR extractelement expression.
type ExtractElementExpr struct {
	// Vector.
	X Constant
	// Element index.
	Index Constant
}

// NewExtractElementExpr returns a new extractelement expression based on the
// given vector and element index.
func NewExtractElementExpr(x, index Constant) *ExtractElementExpr {
	return &ExtractElementExpr{X: x, Index: index}
}

// Type returns the type of the constant expression.
func (e *ExtractElementExpr) Type() types.Type {
	panic("not yet implemented")
}

// Ident returns the identifier associated with the constant expression.
func (e *ExtractElementExpr) Ident() string {
	panic("not yet implemented")
}

// Simplify returns an equivalent (and potentially simplified) constant of the
// constant expression.
func (e *ExtractElementExpr) Simplify() Constant {
	panic("not yet implemented")
}

// ~~~ [ insertelement ] ~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~

// InsertElementExpr is an LLVM IR insertelement expression.
type InsertElementExpr struct {
	// Vector.
	X Constant
	// Element to insert.
	Elem Constant
	// Element index.
	Index Constant
}

// NewInsertElementExpr returns a new insertelement expression based on the
// given vector, element and element index.
func NewInsertElementExpr(x, elem, index Constant) *InsertElementExpr {
	return &InsertElementExpr{X: x, Elem: elem, Index: index}
}

// Type returns the type of the constant expression.
func (e *InsertElementExpr) Type() types.Type {
	panic("not yet implemented")
}

// Ident returns the identifier associated with the constant expression.
func (e *InsertElementExpr) Ident() string {
	panic("not yet implemented")
}

// Simplify returns an equivalent (and potentially simplified) constant of the
// constant expression.
func (e *InsertElementExpr) Simplify() Constant {
	panic("not yet implemented")
}

// ~~~ [ shufflevector ] ~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~

// ShuffleVectorExpr is an LLVM IR shufflevector expression.
type ShuffleVectorExpr struct {
	// Vectors.
	X, Y Constant
	// Shuffle mask.
	Mask Constant
}

// NewShuffleVectorExpr returns a new shufflevector expression based on the
// given vectors and shuffle mask.
func NewShuffleVectorExpr(x, y, mask Constant) *ShuffleVectorExpr {
	return &ShuffleVectorExpr{X: x, Y: y, Mask: mask}
}

// Type returns the type of the constant expression.
func (e *ShuffleVectorExpr) Type() types.Type {
	panic("not yet implemented")
}

// Ident returns the identifier associated with the constant expression.
func (e *ShuffleVectorExpr) Ident() string {
	panic("not yet implemented")
}

// Simplify returns an equivalent (and potentially simplified) constant of the
// constant expression.
func (e *ShuffleVectorExpr) Simplify() Constant {
	panic("not yet implemented")
}
