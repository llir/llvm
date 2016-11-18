// Package ir declares the types used to represent LLVM IR modules.
package ir

import (
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

// AppendFunc appends a function to the module.
func (m *Module) AppendFunc(name string, ret types.Type, params ...Param) *Function {
	f := NewFunction(name, ret, params)
	m.funcs = append(m.funcs, f)
	return f
}
