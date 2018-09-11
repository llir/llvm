package instruction

import (
	"github.com/llir/l/ir/value"
)

// --- [ Vector instructions ] -------------------------------------------------

// ~~~ [ extractelement ] ~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~

// ExtractElement is an LLVM IR extractelement instruction.
type ExtractElement struct {
	// Name of local variable associated with the result.
	LocalName string
	// Vector.
	X value.Value
	// Element index.
	Index value.Value
}

// NewExtractElement returns a new extractelement instruction based on the given
// vector and element index.
func NewExtractElement(x, index value.Value) *ExtractElement {
	return &ExtractElement{X: x, Index: index}
}

// ~~~ [ insertelement ] ~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~

// InsertElement is an LLVM IR insertelement instruction.
type InsertElement struct {
	// Name of local variable associated with the result.
	LocalName string
	// Vector.
	X value.Value
	// Element.
	Elem value.Value
	// Element index.
	Index value.Value
}

// NewInsertElement returns a new insertelement instruction based on the given
// vector, element and element index.
func NewInsertElement(x, elem, index value.Value) *InsertElement {
	return &InsertElement{X: x, Elem: elem, Index: index}
}

// ~~~ [ shufflevector ] ~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~

// ShuffleVector is an LLVM IR shufflevector instruction.
type ShuffleVector struct {
	// Vectors.
	X, Y value.Value
	// Shuffle mask.
	Mask value.Value
}

// NewShuffleVector returns a new shufflevector instruction based on the given
// vectors and shuffle mask.
func NewShuffleVector(x, y, mask value.Value) *ShuffleVector {
	return &ShuffleVector{X: x, Y: y, Mask: mask}
}
