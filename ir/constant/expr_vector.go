// === [ Vector expressions ] ==================================================
//
// References:
//    http://llvm.org/docs/LangRef.html#vector-operations

package constant

import (
	"fmt"

	"github.com/llir/llvm/ir/types"
)

// --- [ extractelement ] ------------------------------------------------------

// ExprExtractElement represents an extractelement expression.
//
// References:
//    http://llvm.org/docs/LangRef.html#extractelement-instruction
type ExprExtractElement struct {
	// Type of the constant expression.
	Typ types.Type
	// Vector.
	X Constant
	// Index.
	Index Constant
}

// NewExtractElement returns a new extractelement expression based on the given
// vector and index.
func NewExtractElement(x, index Constant) *ExprExtractElement {
	t, ok := x.Type().(*types.VectorType)
	if !ok {
		panic(fmt.Errorf("invalid vector type; expected *types.VectorType, got %T", x.Type()))
	}
	return &ExprExtractElement{
		Typ:   t.Elem,
		X:     x,
		Index: index,
	}
}

// Type returns the type of the constant expression.
func (expr *ExprExtractElement) Type() types.Type {
	return expr.Typ
}

// Ident returns the string representation of the constant expression.
func (expr *ExprExtractElement) Ident() string {
	return fmt.Sprintf("extractelement (%s %s, %s %s)",
		expr.X.Type(),
		expr.X.Ident(),
		expr.Index.Type(),
		expr.Index.Ident())
}

// Immutable ensures that only constants can be assigned to the
// constant.Constant interface.
func (*ExprExtractElement) Immutable() {}

// Simplify returns a simplified version of the constant expression.
func (expr *ExprExtractElement) Simplify() Constant {
	panic("not yet implemented")
}

// MetadataNode ensures that only metadata nodes can be assigned to the
// ir.MetadataNode interface.
func (*ExprExtractElement) MetadataNode() {}

// --- [ insertelement ] -------------------------------------------------------

// ExprInsertElement represents an insertelement expression.
//
// References:
//    http://llvm.org/docs/LangRef.html#insertelement-instruction
type ExprInsertElement struct {
	// Vector.
	X Constant
	// Element to insert.
	Elem Constant
	// Index.
	Index Constant
}

// NewInsertElement returns a new insertelement expression based on the given
// vector, element and index.
func NewInsertElement(x, elem, index Constant) *ExprInsertElement {
	return &ExprInsertElement{
		X:     x,
		Elem:  elem,
		Index: index,
	}
}

// Type returns the type of the constant expression.
func (expr *ExprInsertElement) Type() types.Type {
	return expr.X.Type()
}

// Ident returns the string representation of the constant expression.
func (expr *ExprInsertElement) Ident() string {
	return fmt.Sprintf("insertelement (%s %s, %s %s, %s %s)",
		expr.X.Type(),
		expr.X.Ident(),
		expr.Elem.Type(),
		expr.Elem.Ident(),
		expr.Index.Type(),
		expr.Index.Ident())
}

// Immutable ensures that only constants can be assigned to the
// constant.Constant interface.
func (*ExprInsertElement) Immutable() {}

// Simplify returns a simplified version of the constant expression.
func (expr *ExprInsertElement) Simplify() Constant {
	panic("not yet implemented")
}

// MetadataNode ensures that only metadata nodes can be assigned to the
// ir.MetadataNode interface.
func (*ExprInsertElement) MetadataNode() {}

// --- [ shufflevector ] ------------------------------------------------------

// ExprShuffleVector represents an shufflevector expression.
//
// References:
//    http://llvm.org/docs/LangRef.html#shufflevector-instruction
type ExprShuffleVector struct {
	// Vector 1.
	X Constant
	// Vector 2.
	Y Constant
	// Shuffle mask.
	Mask Constant
}

// NewShuffleVector returns a new shufflevector expression based on the given
// vectors and shuffle mask.
func NewShuffleVector(x, y, mask Constant) *ExprShuffleVector {
	return &ExprShuffleVector{
		X:    x,
		Y:    y,
		Mask: mask,
	}
}

// Type returns the type of the constant expression.
func (expr *ExprShuffleVector) Type() types.Type {
	return expr.Mask.Type()
}

// Ident returns the string representation of the constant expression.
func (expr *ExprShuffleVector) Ident() string {
	return fmt.Sprintf("shufflevector (%s %s, %s %s, %s %s)",
		expr.X.Type(),
		expr.X.Ident(),
		expr.Y.Type(),
		expr.Y.Ident(),
		expr.Mask.Type(),
		expr.Mask.Ident())
}

// Immutable ensures that only constants can be assigned to the
// constant.Constant interface.
func (*ExprShuffleVector) Immutable() {}

// Simplify returns a simplified version of the constant expression.
func (expr *ExprShuffleVector) Simplify() Constant {
	panic("not yet implemented")
}

// MetadataNode ensures that only metadata nodes can be assigned to the
// ir.MetadataNode interface.
func (*ExprShuffleVector) MetadataNode() {}
