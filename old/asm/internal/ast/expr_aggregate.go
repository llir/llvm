// === [ Aggregate expressions ] ===============================================
//
// References:
//    http://llvm.org/docs/LangRef.html#aggregate-operations

package ast

// --- [ extractvalue ] ------------------------------------------------------

// ExprExtractValue represents an extractvalue expression.
//
// References:
//    http://llvm.org/docs/LangRef.html#extractvalue-instruction
type ExprExtractValue struct {
	// Type of the constant expression.
	Type Type
	// Aggregate constant.
	X Constant
	// Indices.
	Indices []int64
}

// --- [ insertvalue ] -------------------------------------------------------

// ExprInsertValue represents an insertvalue expression.
//
// References:
//    http://llvm.org/docs/LangRef.html#insertvalue-instruction
type ExprInsertValue struct {
	// Type of the constant expression.
	Type Type
	// Aggregate constant.
	X Constant
	// Element to insert.
	Elem Constant
	// Indices.
	Indices []int64
}

// isValue ensures that only values can be assigned to the ast.Value interface.
func (*ExprExtractValue) isValue() {}
func (*ExprInsertValue) isValue()  {}

// isConstant ensures that only constants can be assigned to the ast.Constant
// interface.
func (*ExprExtractValue) isConstant() {}
func (*ExprInsertValue) isConstant()  {}

// isConstExpr ensures that only constant expressions can be assigned to the
// ast.ConstExpr interface.
func (*ExprExtractValue) isConstExpr() {}
func (*ExprInsertValue) isConstExpr()  {}

// isMetadataNode ensures that only metadata nodes can be assigned to the
// ast.MetadataNode interface.
func (*ExprExtractValue) isMetadataNode() {}
func (*ExprInsertValue) isMetadataNode()  {}
