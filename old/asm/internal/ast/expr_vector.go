// === [ Vector expressions ] ==================================================
//
// References:
//    http://llvm.org/docs/LangRef.html#vector-operations

package ast

// --- [ extractelement ] ------------------------------------------------------

// ExprExtractElement represents an extractelement expression.
//
// References:
//    http://llvm.org/docs/LangRef.html#extractelement-instruction
type ExprExtractElement struct {
	// Type of the constant expression.
	Type Type
	// Vector.
	X Constant
	// Index.
	Index Constant
}

// --- [ insertelement ] -------------------------------------------------------

// ExprInsertElement represents an insertelement expression.
//
// References:
//    http://llvm.org/docs/LangRef.html#insertelement-instruction
type ExprInsertElement struct {
	// Type of the constant expression.
	Type Type
	// Vector.
	X Constant
	// Element to insert.
	Elem Constant
	// Index.
	Index Constant
}

// --- [ shufflevector ] -------------------------------------------------------

// ExprShuffleVector represents a shufflevector expression.
//
// References:
//    http://llvm.org/docs/LangRef.html#shufflevector-instruction
type ExprShuffleVector struct {
	// Type of the constant expression.
	Type Type
	// Vector 1.
	X Constant
	// Vector 2.
	Y Constant
	// Shuffle mask.
	Mask Constant
}

// isValue ensures that only values can be assigned to the ast.Value interface.
func (*ExprExtractElement) isValue() {}
func (*ExprInsertElement) isValue()  {}
func (*ExprShuffleVector) isValue()  {}

// isConstant ensures that only constants can be assigned to the ast.Constant
// interface.
func (*ExprExtractElement) isConstant() {}
func (*ExprInsertElement) isConstant()  {}
func (*ExprShuffleVector) isConstant()  {}

// isConstExpr ensures that only constant expressions can be assigned to the
// ast.ConstExpr interface.
func (*ExprExtractElement) isConstExpr() {}
func (*ExprInsertElement) isConstExpr()  {}
func (*ExprShuffleVector) isConstExpr()  {}

// isMetadataNode ensures that only metadata nodes can be assigned to the
// ast.MetadataNode interface.
func (*ExprExtractElement) isMetadataNode() {}
func (*ExprInsertElement) isMetadataNode()  {}
func (*ExprShuffleVector) isMetadataNode()  {}
