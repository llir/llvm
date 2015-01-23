package ir

// A Function declaration specifies the name and type of a function. A function
// definition contains a set of basic blocks, interconnected by control flow
// instructions (such as br), which forms the nodes in a Control Flow Graph of
// the function [1,2].
//
//    [1]: http://llvm.org/docs/LangRef.html#functions
//    [2]: http://llvm.org/docs/LangRef.html#terminators
type Function struct {
	// Function name.
	Name string
	// Function type.
	Type FuncType
	// Basic blocks of the function (or nil if function declaration).
	Blocks []*BasicBlock
}

// FuncType representes a function type.
type FuncType struct {
	// Function return argument type.
	Ret Type
	// Function argument types.
	Args []Type
}
