package ir

// A Module contains top-level function definitions, external function
// declarations, global variables, type definitions and metadata [1].
//
//    [1]: http://llvm.org/docs/LangRef.html#module-structure
type Module struct {

	// TODO(u): Add external function declarations, or let *Function handle
	// function definitions without bodies. If *Function handles both, update the
	// doc comment to:
	//    Function definitions and external function declarations (Blocks is nil).

	// Function definitions.
	Funcs []*Function

	// Global variables.
	Globals []Value
	// Type definitions.
	Types []Type
	// Metadata.
	Metadata []*Metadata
}
