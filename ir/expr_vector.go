package ir

import "github.com/llir/l/ir/types"

// --- [ Vector expressions ] --------------------------------------------------

// ~~~ [ extractelement ] ~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~

// ExprExtractElement is an LLVM IR extractelement expression.
type ExprExtractElement struct {
	// Vector.
	X Constant
	// Element index.
	Index Constant
}

// NewExtractElementExpr returns a new extractelement expression based on the
// given vector and element index.
func NewExtractElementExpr(x, index Constant) *ExprExtractElement {
	return &ExprExtractElement{X: x, Index: index}
}

// Type returns the type of the constant expression.
func (e *ExprExtractElement) Type() types.Type {
	panic("not yet implemented")
}

// Ident returns the identifier associated with the constant expression.
func (e *ExprExtractElement) Ident() string {
	panic("not yet implemented")
}

// Simplify returns an equivalent (and potentially simplified) constant of the
// constant expression.
func (e *ExprExtractElement) Simplify() Constant {
	panic("not yet implemented")
}

// ~~~ [ insertelement ] ~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~

// ExprInsertElement is an LLVM IR insertelement expression.
type ExprInsertElement struct {
	// Vector.
	X Constant
	// Element to insert.
	Elem Constant
	// Element index.
	Index Constant
}

// NewInsertElementExpr returns a new insertelement expression based on the
// given vector, element and element index.
func NewInsertElementExpr(x, elem, index Constant) *ExprInsertElement {
	return &ExprInsertElement{X: x, Elem: elem, Index: index}
}

// Type returns the type of the constant expression.
func (e *ExprInsertElement) Type() types.Type {
	panic("not yet implemented")
}

// Ident returns the identifier associated with the constant expression.
func (e *ExprInsertElement) Ident() string {
	panic("not yet implemented")
}

// Simplify returns an equivalent (and potentially simplified) constant of the
// constant expression.
func (e *ExprInsertElement) Simplify() Constant {
	panic("not yet implemented")
}

// ~~~ [ shufflevector ] ~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~

// ExprShuffleVector is an LLVM IR shufflevector expression.
type ExprShuffleVector struct {
	// Vectors.
	X, Y Constant
	// Shuffle mask.
	Mask Constant
}

// NewShuffleVectorExpr returns a new shufflevector expression based on the
// given vectors and shuffle mask.
func NewShuffleVectorExpr(x, y, mask Constant) *ExprShuffleVector {
	return &ExprShuffleVector{X: x, Y: y, Mask: mask}
}

// Type returns the type of the constant expression.
func (e *ExprShuffleVector) Type() types.Type {
	panic("not yet implemented")
}

// Ident returns the identifier associated with the constant expression.
func (e *ExprShuffleVector) Ident() string {
	panic("not yet implemented")
}

// Simplify returns an equivalent (and potentially simplified) constant of the
// constant expression.
func (e *ExprShuffleVector) Simplify() Constant {
	panic("not yet implemented")
}
