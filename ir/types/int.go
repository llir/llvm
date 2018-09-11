package types

// IntType is an LLVM IR integer type.
type IntType struct {
}

// Equal reports whether t and u are of equal type.
func (t *IntType) Equal(u Type) bool {
	panic("not yet implemented")
}
