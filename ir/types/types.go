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

// FuncType represents a function type.
type FuncType struct {
}

// LabelType represents a label type, which is used for basic block values.
type LabelType struct {
}

// Equal reports whether t and u are of equal type.
func (t LabelType) Equal(u Type) bool {
	_, ok := u.(LabelType)
	return ok
}

var Label = &LabelType{}
