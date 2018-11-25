package ir

import "github.com/llir/llvm/ir/types"

// --- [ Functions ] -----------------------------------------------------------

// NewFunc appends a new function to the module based on the given function
// name, return type and function parameters.
func (m *Module) NewFunc(name string, retType types.Type, params ...*Param) *Function {
	f := NewFunc(name, retType, params...)
	m.Funcs = append(m.Funcs, f)
	return f
}
