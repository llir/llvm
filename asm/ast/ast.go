// Package ast declares the types used to represent abstract syntax trees of
// LLVM IR assembly files.
package ast

import "fmt"

// === [ Modules ] =============================================================

// A Module represents an LLVM IR module.
type Module struct {
	// Top-level declarations.
	decls []TopLevelDecl
}

// TODO: Add underlying types of Node.

// A Node represents a node within the abstract syntax tree, and has one of the
// following underlying types.
//
//    *Module
type Node interface {
	fmt.Stringer
	// Start returns the start position of the node within the input stream.
	Start() int
}

// === [ Declarations ] ========================================================

// TODO: Add underlying types of TopLevelDecl.

// A TopLevelDecl node represents a top-level declaration, and has one of the
// following underlying types.
//
//    *TargetLayout
//    *TargetTriple
//    *TypeDef
//    *GlobalVarDecl
//    *FuncDecl
//    *AttrGroupDef
//    *MetadataDef
type TopLevelDecl interface {
	Node
	// isTopLevelDecl ensures that only top-level declaration nodes can be
	// assigned to the TopLevelDecl interface.
	isTopLevelDecl()
}

// Top-level declaration nodes.
type (
	// --- [ Target specifiers ] ------------------------------------------------

	// A TargetLayout node specifies the data layout of the target.
	//
	// Examples:
	//    target datalayout = "e-m:e-i64:64-f80:128-n8:16:32:64-S128"
	//
	// References:
	//    http://llvm.org/docs/LangRef.html#data-layout
	TargetLayout struct {
		// start position of the "target" keyword within the input stream.
		start int
		// Data layout of the target.
		layout string
	}

	// A TargetTriple node specifies the host architecture, operating system and
	// vendor of the target.
	//
	// Examples:
	//    target triple = "x86_64-unknown-linux-gnu"
	//
	// References:
	//    http://llvm.org/docs/LangRef.html#target-triple
	TargetTriple struct {
		// start position of the "target" keyword within the input stream.
		start int
		// Host architecture, operating system and vendor of the target.
		triple string
	}

	// --- [ Type definitions ] -------------------------------------------------

	// A TypeDef node represents a type definition.
	//
	// Examples:
	//    %point = type { i32, i32 }
	//
	// References:
	//    http://llvm.org/docs/LangRef.html#structure-type
	TypeDef struct {
		// Name of the identified type.
		name string
		// Underlying type.
		typ Type
	}
)
