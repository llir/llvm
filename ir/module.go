package ir

import (
	"bytes"
	"fmt"

	"github.com/llir/llvm/asm"
	"github.com/llir/llvm/types"
)

// A Module contains top-level type definitions, global variables, function
// definitions, external function declarations, and metadata.
//
// References:
//    http://llvm.org/docs/LangRef.html#module-structure
type Module struct {
	// Data layout.
	layout string
	// Target host.
	target string
	// Type definitions.
	typeDecls []types.NamedStruct
	// Global variables.
	globalDecls []*GlobalDecl
	// Function definitions and external function declarations (Blocks is nil).
	funcDecls []*Function
	// Metadata nodes.
	metaDecls []*Metadata
}

// Layout returns the data layout of the module.
//
// Examples:
//    target datalayout = "e-m:e-i64:64-f80:128-n8:16:32:64-S128"
//
// References:
//    http://llvm.org/docs/LangRef.html#data-layout
func (m *Module) Layout() string {
	return m.layout
}

// Target returns the target host of the module.
//
// Examples:
//    x86_64-unknown-linux-gnu
//
// References:
//    http://llvm.org/docs/LangRef.html#target-triple
func (m *Module) Target() string {
	return m.target
}

// String returns a string representation of the module and its top-level
// declarations and definitions of types, global variables, functions and
// metadata notes.
func (m *Module) String() string {
	// Data layout; e.g.
	//    target datalayout = "e-m:e-i64:64-f80:128-n8:16:32:64-S128"
	buf := new(bytes.Buffer)
	if len(m.layout) > 0 {
		fmt.Fprintf(buf, "target datalayout = %q\n", m.layout)
	}

	// Target triple; e.g.
	//    target triple = "x86_64-unknown-linux-gnu"
	if len(m.target) > 0 {
		fmt.Fprintf(buf, "target triple = %q\n", m.target)
	}

	// Type definitions; e.g.
	//    %foo = type {i32}
	for _, typeDecl := range m.typeDecls {
		fmt.Fprintln(buf, "%s = type %v\n", asm.EncLocal(typeDecl.Name()), typeDecl.Struct)
	}

	// Global variables; e.g.
	//    @x = global i32 42
	for _, globalDecl := range m.globalDecls {
		fmt.Fprintln(buf, globalDecl)
	}

	// TODO: Print functions.
	// TODO: Print named metadata.
	// TODO: Print metadata.
	panic("not yet implemented.")
}
