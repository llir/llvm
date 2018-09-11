package types

// ArrayType is an LLVM IR array type.
type ArrayType struct {
}

// Equal reports whether t and u are of equal type.
func (t *ArrayType) Equal(u Type) bool {
	panic("not yet implemented")
}
