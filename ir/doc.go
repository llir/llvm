// Package ir declares the types used to represent the LLVM IR language [1].
//
// LLVM code is organized into modules containing top-level definitions, such as
// functions and global variables. A function definition contains a set of basic
// blocks, which forms the nodes in a Control Flow Graph of the function. Each
// basic block consists of a sequence of non-branching instructions, terminated
// by a control flow instruction (such as br or ret).
//
//    [1]: http://llvm.org/docs/LangRef.html
package ir
