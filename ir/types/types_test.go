package types

// Assert that each type implements the types.Type interface.
var (
	_ Type = (*VoidType)(nil)
	_ Type = (*FuncType)(nil)
	_ Type = (*IntType)(nil)
	_ Type = (*FloatType)(nil)
	_ Type = (*MMXType)(nil)
	_ Type = (*PointerType)(nil)
	_ Type = (*VectorType)(nil)
	_ Type = (*LabelType)(nil)
	_ Type = (*TokenType)(nil)
	_ Type = (*MetadataType)(nil)
	_ Type = (*ArrayType)(nil)
	_ Type = (*StructType)(nil)
)
