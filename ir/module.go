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

	"github.com/llir/llvm/ir/constant"
	"github.com/llir/llvm/ir/types"
)

// A Module represents an LLVM IR module, which consists of top-level type
// definitions, global variables, functions, and metadata.
type Module struct {
	// Type definitions.
	Types []*types.NamedType
	// Global variables of the module.
	Globals []*Global
	// Functions of the module.
	Funcs []*Function
}

// NewModule returns a new LLVM IR module.
func NewModule() *Module {
	return &Module{}
}

// String returns the LLVM syntax representation of the module.
func (m *Module) String() string {
	buf := &bytes.Buffer{}
	for _, typ := range m.Types {
		if typ.Def == nil {
			panic(fmt.Errorf("invalid type definition %q; expected underlying type definition, got nil", typ))
		}
		if def, ok := typ.Def.(*types.StructType); ok {
			fmt.Fprintf(buf, "%s = type { ", typ)
			for i, field := range def.Fields {
				if i != 0 {
					buf.WriteString(", ")
				}
				buf.WriteString(field.String())
			}
			buf.WriteString(" }\n")
		} else {
			fmt.Fprintf(buf, "%s = type %s\n", typ, def)
		}
	}
	for _, global := range m.Globals {
		fmt.Fprintln(buf, global)
	}
	for _, f := range m.Funcs {
		fmt.Fprintln(buf, f)
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
//
// A nil underlying type definition may be used to specify an opaque struct
// type, the body of which may later be specified using the SetDef method.
func (m *Module) NewType(name string, def types.Type) *types.NamedType {
	typ := types.NewNamed(name, def)
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
