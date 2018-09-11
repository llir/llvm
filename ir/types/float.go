package types

// FloatType is an LLVM IR floating-point type.
type FloatType struct {
}

// Equal reports whether t and u are of equal type.
func (t *FloatType) Equal(u Type) bool {
	panic("not yet implemented")
}
