package types

// Type is an LLVM IR type.
type Type interface {
	// Equal reports whether t and u are of equal type.
	Equal(u Type) bool
}

// Convenience types.
var (
	Label = &LabelType{}
	Token = &TokenType{}
)
