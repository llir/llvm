package ir

import (
	"bytes"
	"fmt"

	"github.com/mewlang/llvm/types"
	"github.com/mewlang/llvm/values"
)

// TODO: Use map from Global/Local to *Function, Value, types.Type and *Metadata
// instead of slice.

// A Module contains top-level function definitions, external function
// declarations, global variables, type definitions and metadata.
//
// References:
//    http://llvm.org/docs/LangRef.html#module-structure
type Module struct {
	// layout specifies how data is laid out in memory as a list of
	// specifications separated by the minus sign character (-). When
	// constructing the data layout for a given target, LLVM starts with a
	// default set of specifications which are then overridden by the
	// specifications of layout.
	//
	// Examples:
	//    target datalayout = "e-m:e-i64:64-f80:128-n8:16:32:64-S128"
	//
	// References:
	//    http://llvm.org/docs/LangRef.html#data-layout
	layout string
	// target describes the target host as a series of identifiers delimited by
	// the minus sign character (-). The canonical forms for target triple
	// strings are:
	//    ARCHITECTURE-VENDOR-OPERATING_SYSTEM
	//    ARCHITECTURE-VENDOR-OPERATING_SYSTEM-ENVIRONMENT
	//
	// Examples:
	//    x86_64-unknown-linux-gnu
	//
	// References:
	//    http://llvm.org/docs/LangRef.html#target-triple
	target string
	// Function definitions and external function declarations (Blocks is nil).
	funcs []*Function
	// Global variables.
	globals []values.Value
	// Type definitions.
	types []types.Type
	// Metadata.
	metadata []*Metadata
}

func (module *Module) String() string {
	buf := new(bytes.Buffer)
	// Data layout.
	if len(module.layout) > 0 {
		// target datalayout = "e-m:e-i64:64-f80:128-n8:16:32:64-S128"
		fmt.Fprintf(buf, "target datalayout = %q\n", module.layout)
	}
	// Target triple.
	if len(module.target) > 0 {
		// target triple = "x86_64-unknown-linux-gnu"
		fmt.Fprintf(buf, "target triple = %q\n", module.target)
	}
	panic("not yet implemented.")
}
