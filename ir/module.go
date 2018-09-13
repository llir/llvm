// Package ir declares the types used to represent LLVM IR modules.
package ir

import "github.com/llir/l/ir/types"

// Module is an LLVM IR module.
type Module struct {
	// Type definitions.
	TypeDefs []types.Type
	// Global variable declarations and definitions.
	Globals []*Global
	// Function declarations and definitions.
	Funcs []*Function
}
