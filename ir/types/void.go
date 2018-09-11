package types

// VoidType is an LLVM IR void type.
type VoidType struct {
}

// Equal reports whether t and u are of equal type.
func (t *VoidType) Equal(u Type) bool {
	panic("not yet implemented")
}
