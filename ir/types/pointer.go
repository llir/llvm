package types

// PointerType is an LLVM IR pointer type.
type PointerType struct {
}

// Equal reports whether t and u are of equal type.
func (t *PointerType) Equal(u Type) bool {
	panic("not yet implemented")
}
