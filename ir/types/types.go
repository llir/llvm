// Package types declares the data types of LLVM IR.
package types

// A Type represents an LLVM IR type.
//
// Type may have one of the following underlying types.
//
//   TODO
type Type interface {
	// Equal reports whether t and u are of equal type.
	Equal(u Type) bool
}

// Func represents a function type.
type Func struct {
}
