package ir

import (
	"github.com/mewlang/llvm/types"
	"github.com/mewlang/llvm/values"
)

// TODO: Use map from Global/Local to *Function, Value, types.Type and *Metadata
// instead of slice.

// A Module contains top-level function definitions, external function
// declarations, global variables, type definitions and metadata [1].
//
//    [1]: http://llvm.org/docs/LangRef.html#module-structure
type Module struct {
	// Function definitions and external function declarations (Blocks is nil).
	Funcs []*Function
	// Global variables.
	Globals []values.Value
	// Type definitions.
	Types []types.Type
	// Metadata.
	Metadata []*Metadata
}
