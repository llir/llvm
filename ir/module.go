package ir

// A Module contains top-level function definitions, external function
// declarations, global variables, type definitions and metadata [1].
//
//    [1]: http://llvm.org/docs/LangRef.html#module-structure
type Module struct {
	// Function definitions and external function declarations (Blocks is nil).
	Funcs []*Function
	// Global variables.
	Globals []Value
	// Type definitions.
	Types []Type
	// Metadata.
	Metadata []*Metadata
}
