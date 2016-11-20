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
	// Functions of the module.
	funcs []*Function
	// Global variables of the module.
	globals []*Global
}

// NewModule returns a new LLVM IR module.
func NewModule() *Module {
	return &Module{}
}

// LLVMString returns the LLVM syntax representation of the module.
func (m *Module) LLVMString() string {
	buf := &bytes.Buffer{}
	for _, f := range m.Funcs() {
		fmt.Fprintln(buf, f.LLVMString())
	}
	for _, g := range m.Globals() {
		fmt.Fprintln(buf, g.LLVMString())
	}
	return buf.String()
}

// Funcs returns the functions of the module.
func (m *Module) Funcs() []*Function {
	return m.funcs
}

// Globals returns the global variables of the module.
func (m *Module) Globals() []*Global {
	return m.globals
}

// AppendFunction appends the given function to the module.
func (m *Module) AppendFunction(f *Function) {
	f.SetParent(m)
	m.funcs = append(m.funcs, f)
}

// AppendGlobal appends the given global variable to the module.
func (m *Module) AppendGlobal(g *Global) {
	m.globals = append(m.globals, g)
}

// NewFunction appends a new function to the module based on the given function
// name, return type and parameters.
func (m *Module) NewFunction(name string, ret types.Type, params ...*types.Param) *Function {
	f := NewFunction(name, ret, params...)
	m.AppendFunction(f)
	return f
}

// NewGlobalDecl appends a new external global variable declaration to the
// module based on the given global variable name and content type.
func (m *Module) NewGlobalDecl(name string, content types.Type) *Global {
	g := NewGlobalDecl(name, content)
	m.AppendGlobal(g)
	return g
}

// NewGlobalDef appends a new global variable definition to the module based on
// the given global variable name and initial value.
func (m *Module) NewGlobalDef(name string, init constant.Constant) *Global {
	g := NewGlobalDef(name, init)
	m.AppendGlobal(g)
	return g
}
