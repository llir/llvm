//go:generate go run gen.go

// Package ast declares the types used to represent abstract syntax trees of
// LLVM IR modules.
package ast

// A Module represents an LLVM IR module, which consists of top-level type
// definitions, global variables, functions, and metadata.
type Module struct {
	// Type definitions.
	Types []*NamedType
	// Global variables of the module.
	Globals []*Global
	// Functions of the module.
	Funcs []*Function
	// Named metadata of the module.
	NamedMetadata []*NamedMetadata
	// Metadata of the module.
	Metadata []*Metadata
}
