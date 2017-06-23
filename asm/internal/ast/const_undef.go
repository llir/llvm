package ast

// UndefConst represents an undef constant.
type UndefConst struct {
	// Constant type.
	Type Type
}

// isValue ensures that only values can be assigned to the ast.Value interface.
func (*UndefConst) isValue() {}

// isConstant ensures that only constants can be assigned to the ast.Constant
// interface.
func (*UndefConst) isConstant() {}

// isMetadataNode ensures that only metadata nodes can be assigned to the
// ast.MetadataNode interface.
func (*UndefConst) isMetadataNode() {}
