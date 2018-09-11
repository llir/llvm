package types

// VectorType is an LLVM IR vector type.
type VectorType struct {
}

// Equal reports whether t and u are of equal type.
func (t *VectorType) Equal(u Type) bool {
	panic("not yet implemented")
}
