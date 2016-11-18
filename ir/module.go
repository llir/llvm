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
	for _, f := range m.funcs {
		fmt.Fprintln(buf, f.LLVMString())
	}
	for _, g := range m.globals {
		fmt.Fprintln(buf, g.LLVMString())
	}
	return buf.String()
}

// NewFunction appends a new function to the module based on the given function
// name, return type and parameters.
func (m *Module) NewFunction(name string, ret types.Type, params ...*Param) *Function {
	f := NewFunction(name, ret, params...)
	f.SetParent(m)
	m.funcs = append(m.funcs, f)
	return f
}

// NewGlobal appends a new global variable to the module based on the given
// global variable name, underlying type and optional initial value.
func (m *Module) NewGlobal(name string, underlying types.Type, init ...constant.Constant) *Global {
	global := NewGlobal(name, underlying, init...)
	m.globals = append(m.globals, global)
	return global
}
