package constant

import "github.com/llir/l/ir/types"

// --- [ Vector expressions ] --------------------------------------------------

// ~~~ [ extractelement ] ~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~

// ExtractElement is an LLVM IR extractelement expression.
type ExtractElement struct {
	// Vector.
	X Constant
	// Element index.
	Index Constant
}

// NewExtractElement returns a new extractelement expression based on the given
// vector and element index.
func NewExtractElement(x, index Constant) *ExtractElement {
	return &ExtractElement{X: x, Index: index}
}

// Type returns the type of the constant expression.
func (e *ExtractElement) Type() types.Type {
	panic("not yet implemented")
}

// Ident returns the identifier associated with the constant expression.
func (e *ExtractElement) Ident() string {
	panic("not yet implemented")
}

// Simplify returns an equivalent (and potentially simplified) constant of the
// constant expression.
func (e *ExtractElement) Simplify() Constant {
	panic("not yet implemented")
}

// ~~~ [ insertelement ] ~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~

// InsertElement is an LLVM IR insertelement expression.
type InsertElement struct {
	// Vector.
	X Constant
	// Element to insert.
	Elem Constant
	// Element index.
	Index Constant
}

// NewInsertElement returns a new insertelement expression based on the given
// vector, element and element index.
func NewInsertElement(x, elem, index Constant) *InsertElement {
	return &InsertElement{X: x, Elem: elem, Index: index}
}

// Type returns the type of the constant expression.
func (e *InsertElement) Type() types.Type {
	panic("not yet implemented")
}

// Ident returns the identifier associated with the constant expression.
func (e *InsertElement) Ident() string {
	panic("not yet implemented")
}

// Simplify returns an equivalent (and potentially simplified) constant of the
// constant expression.
func (e *InsertElement) Simplify() Constant {
	panic("not yet implemented")
}

// ~~~ [ shufflevector ] ~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~

// ShuffleVector is an LLVM IR shufflevector expression.
type ShuffleVector struct {
	// Vectors.
	X, Y Constant
	// Shuffle mask.
	Mask Constant
}

// NewShuffleVector returns a new shufflevector expression based on the given
// vectors and shuffle mask.
func NewShuffleVector(x, y, mask Constant) *ShuffleVector {
	return &ShuffleVector{X: x, Y: y, Mask: mask}
}

// Type returns the type of the constant expression.
func (e *ShuffleVector) Type() types.Type {
	panic("not yet implemented")
}

// Ident returns the identifier associated with the constant expression.
func (e *ShuffleVector) Ident() string {
	panic("not yet implemented")
}

// Simplify returns an equivalent (and potentially simplified) constant of the
// constant expression.
func (e *ShuffleVector) Simplify() Constant {
	panic("not yet implemented")
}
