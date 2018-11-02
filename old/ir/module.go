//go:generate go run gen.go

// === [ Modules ] =============================================================
//
// References:
//    http://llvm.org/docs/LangRef.html#module-structure

// Package ir declares the types used to represent LLVM IR modules.
package ir

import (
	"bytes"
	"fmt"

	"github.com/llir/llvm/internal/enc"
	"github.com/llir/llvm/ir/constant"
	"github.com/llir/llvm/ir/metadata"
	"github.com/llir/llvm/ir/types"
)

// A Module represents an LLVM IR module, which consists of top-level type
// definitions, global variables, functions, and metadata.
type Module struct {
	// Data layout.
	DataLayout string
	// Target triple.
	TargetTriple string
	// Type definitions.
	Types []types.Type
	// Global variables of the module.
	Globals []*Global
	// Functions of the module.
	Funcs []*Function
	// Named metadata of the module.
	NamedMetadata []*metadata.Named
	// Metadata of the module.
	Metadata []*metadata.Metadata
}

// NewModule returns a new LLVM IR module.
func NewModule() *Module {
	return &Module{}
}

// String returns the LLVM syntax representation of the module.
func (m *Module) String() string {
	buf := &bytes.Buffer{}
	if len(m.DataLayout) > 0 {
		fmt.Fprintf(buf, "target datalayout = %q\n", m.DataLayout)
	}
	if len(m.TargetTriple) > 0 {
		fmt.Fprintf(buf, "target triple = %q\n", m.TargetTriple)
	}
	for _, typ := range m.Types {
		if len(buf.Bytes()) > 0 {
			buf.WriteString("\n")
		}
		name := enc.Local(typ.GetName())
		fmt.Fprintf(buf, "%s = type %s\n", name, typ.Def())
	}
	for _, global := range m.Globals {
		if len(buf.Bytes()) > 0 {
			buf.WriteString("\n")
		}
		fmt.Fprintln(buf, global)
	}
	for _, f := range m.Funcs {
		if len(buf.Bytes()) > 0 {
			buf.WriteString("\n")
		}
		fmt.Fprintln(buf, f)
	}
	for _, md := range m.NamedMetadata {
		if len(buf.Bytes()) > 0 {
			buf.WriteString("\n")
		}
		name := enc.Metadata(md.Name)
		fmt.Fprintf(buf, "%s = %s\n", name, md.Def())
	}
	for _, md := range m.Metadata {
		if len(buf.Bytes()) > 0 {
			buf.WriteString("\n")
		}
		id := enc.Metadata(md.ID)
		fmt.Fprintf(buf, "%s = %s\n", id, md.Def())
	}
	return buf.String()
}

// AppendFunction appends the given function to the module.
func (m *Module) AppendFunction(f *Function) {
	f.Parent = m
	m.Funcs = append(m.Funcs, f)
}

// NewType appends a new type definition to the module based on the given type
// name and underlying type definition.
func (m *Module) NewType(name string, typ types.Type) types.Type {
	typ.SetName(name)
	m.Types = append(m.Types, typ)
	return typ
}

// NewGlobalDecl appends a new external global variable declaration to the
// module based on the given global variable name and content type.
func (m *Module) NewGlobalDecl(name string, content types.Type) *Global {
	global := NewGlobalDecl(name, content)
	m.Globals = append(m.Globals, global)
	return global
}

// NewGlobalDef appends a new global variable definition to the module based on
// the given global variable name and initial value.
func (m *Module) NewGlobalDef(name string, init constant.Constant) *Global {
	global := NewGlobalDef(name, init)
	m.Globals = append(m.Globals, global)
	return global
}

// NewFunction appends a new function to the module based on the given function
// name, return type and parameters.
func (m *Module) NewFunction(name string, ret types.Type, params ...*types.Param) *Function {
	f := NewFunction(name, ret, params...)
	m.AppendFunction(f)
	return f
}
