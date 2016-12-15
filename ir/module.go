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
	types []*types.NamedType
	// Global variables of the module.
	globals []*Global
	// Functions of the module.
	funcs []*Function
}

// NewModule returns a new LLVM IR module.
func NewModule() *Module {
	return &Module{}
}

// String returns the LLVM syntax representation of the module.
func (m *Module) String() string {
	buf := &bytes.Buffer{}
	for _, typ := range m.Types() {
		def, ok := typ.Def()
		if !ok {
			panic(fmt.Sprintf("invalid type definition %q; expected underlying type definition, got nil", typ))
		}
		if def, ok := def.(*types.StructType); ok {
			fmt.Fprintf(buf, "%s = type { ", typ)
			for i, field := range def.Fields() {
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
	for _, global := range m.Globals() {
		fmt.Fprintln(buf, global)
	}
	for _, f := range m.Funcs() {
		fmt.Fprintln(buf, f)
	}
	return buf.String()
}

// Types returns the type definitions of the module.
func (m *Module) Types() []*types.NamedType {
	return m.types
}

// Globals returns the global variables of the module.
func (m *Module) Globals() []*Global {
	return m.globals
}

// Funcs returns the functions of the module.
func (m *Module) Funcs() []*Function {
	return m.funcs
}

// AppendType appends the given type definition to the module.
func (m *Module) AppendType(typ *types.NamedType) {
	m.types = append(m.types, typ)
}

// AppendGlobal appends the given global variable to the module.
func (m *Module) AppendGlobal(global *Global) {
	m.globals = append(m.globals, global)
}

// AppendFunction appends the given function to the module.
func (m *Module) AppendFunction(f *Function) {
	f.SetParent(m)
	m.funcs = append(m.funcs, f)
}

// NewType appends a new type definition to the module based on the given type
// name and underlying type definition.
//
// A nil underlying type definition may be used to specify an opaque struct
// type, the body of which may later be specified using the SetDef method.
func (m *Module) NewType(name string, def types.Type) *types.NamedType {
	typ := types.NewNamed(name, def)
	m.AppendType(typ)
	return typ
}

// NewGlobalDecl appends a new external global variable declaration to the
// module based on the given global variable name and content type.
func (m *Module) NewGlobalDecl(name string, content types.Type) *Global {
	global := NewGlobalDecl(name, content)
	m.AppendGlobal(global)
	return global
}

// NewGlobalDef appends a new global variable definition to the module based on
// the given global variable name and initial value.
func (m *Module) NewGlobalDef(name string, init constant.Constant) *Global {
	global := NewGlobalDef(name, init)
	m.AppendGlobal(global)
	return global
}

// NewFunction appends a new function to the module based on the given function
// name, return type and parameters.
func (m *Module) NewFunction(name string, ret types.Type, params ...*Param) *Function {
	f := NewFunction(name, ret, params...)
	m.AppendFunction(f)
	return f
}
