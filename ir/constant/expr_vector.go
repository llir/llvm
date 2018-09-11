package constant

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
