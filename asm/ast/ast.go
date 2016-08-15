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
	// isTopLevelDecl ensures that only top-level declaration nodes can be
	// assigned to the TopLevelDecl interface.
	isTopLevelDecl()
}

// --- [ Target specifiers ] ------------------------------------------------

// TargetLayout specifies the data layout of the target.
//
// Examples:
//    target datalayout = "e-m:e-i64:64-f80:128-n8:16:32:64-S128"
//
// References:
//    http://llvm.org/docs/LangRef.html#data-layout
type TargetLayout struct {
	// Data layout of the target.
	Layout string
}

// TargetTriple specifies the host architecture, operating system and vendor of
// the target.
//
// Examples:
//    target triple = "x86_64-unknown-linux-gnu"
//
// References:
//    http://llvm.org/docs/LangRef.html#target-triple
type TargetTriple struct {
	// Host architecture, operating system and vendor of the target.
	Triple string
}

// --- [ Type definitions ] -------------------------------------------------

// A TypeDef represents a type definition.
//
// Examples:
//    %point = type { i32, i32 }
//
// References:
//    http://llvm.org/docs/LangRef.html#structure-type
type TypeDef struct {
	// Name of the identified type.
	Name string
	// Underlying type.
	Type Type
}

// A GlobalVarDecl represents a global variable definition or an external global
// variable declaration.
//
// Examples:
//    @x = global i32 42
//    @y = external global i32
//    @s = constant [13 x i8] c"hello world\0A\00"
//
// References:
//    http://llvm.org/docs/LangRef.html#global-variables
type GlobalVarDecl struct {
	// Global variable name.
	Name string
	// Linkage type of the global variable.
	Linkage Linkage
	// The address of the global variable is not significant, only its content.
	UnnamedAddr bool
	// Immutability of the global variable.
	Immutable bool
	// Global variable type.
	Type Type
	// Initial value, or nil if defined externally.
	Val Value
	// Alignment; or 0 if not aligned.
	Align int
}

// Linkage specifies the linkage type of a global variable or function.
//
// Examples:
//    linkonce_odr
//
// References:
//    http://llvm.org/docs/LangRef.html#linkage
type Linkage int

// Linkage types.
const (
	LinkageNone                Linkage = iota
	LinkageAppending                   // Valid for global variable definition.
	LinkageAvailableExternally         // Valid for global variable definition or function definition.
	LinkageCommon                      // Valid for global variable definition.
	LinkageExternal                    // Valid for external global variable declaration, function declaration or function defintion.
	LinkageExternWeak                  // Valid for external global variable declaration or function declaration.
	LinkageInternal                    // Valid for global variable definition or function definition.
	LinkageLinkonce                    // Valid for global variable definition or function definition.
	LinkageLinkonceODR                 // Valid for global variable definition or function definition.
	LinkagePrivate                     // Valid for global variable definition or function definition.
	LinkageWeak                        // Valid for global variable definition or function definition.
	LinkageWeakODR                     // Valid for global variable definition or function definition.
)

// String returns the string representation of the linkage type.
func (l Linkage) String() string {
	switch l {
	case LinkageNone:
		return "none"
	case LinkageAppending:
		return "appending"
	case LinkageAvailableExternally:
		return "available_externally"
	case LinkageCommon:
		return "common"
	case LinkageExternal:
		return "external"
	case LinkageExternWeak:
		return "extern_weak"
	case LinkageInternal:
		return "internal"
	case LinkageLinkonce:
		return "linkonce"
	case LinkageLinkonceODR:
		return "linkonce_odr"
	case LinkagePrivate:
		return "private"
	case LinkageWeak:
		return "weak"
	case LinkageWeakODR:
		return "weak_odr"
	}
	panic(fmt.Sprintf("support for linkage type %d not yet implemented", int(l)))
}

// --- [ Function declarations ] -----------------------------------------------

// TODO: Continue here (last edit 2016-08-15). Figure out how to represent
// function results and parameters, as a FuncType signature, or as individual
// Params? Paramters may have additional parameter attributes attached, thus
// making FuncType a suboptimal representation.

// A FuncDecl represents a function declaration.
//
// Examples:
//    declare i32 @printf(i32*, ...)
//    define i32 @add(i32 %x, i32 %y) { … }
//
// References:
//    http://llvm.org/docs/LangRef.html#functions
type FuncDecl struct {
	// Linkage type of the function declaration.
	Linkage Linkage
	// Function name.
	Name string
	// Function signature.
	Sig *FuncType
	// Function body; or nil if function declaration.
	Body FuncBody
}

// A FuncType represents a function type.
type FuncType struct {
}
