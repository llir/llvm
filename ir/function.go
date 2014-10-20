package ir

// A Function definition contains a set of basic blocks, interconnected by
// control flow instructions (such as br), which forms the nodes in a Control
// Flow Graph of the function [1,2].
//
//    [1]: http://llvm.org/docs/LangRef.html#functions
//    [2]: http://llvm.org/docs/LangRef.html#terminators
type Function struct {
	// Basic blocks of the function.
	Blocks []*BasicBlock

	// Linkage type.
	Linkage Linkage
	// Visibility style.
	Visibility Visibility
	// Calling convention.
	CC CallingConv
}
