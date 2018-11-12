package ir

import "github.com/llir/llvm/ir/types"

// --- [ Functions ] -----------------------------------------------------------

// NewFunction appends a new function to the module based on the given function
// name, return type and function parameters.
func (m *Module) NewFunction(name string, retType types.Type, params ...*Param) *Function {
	f := NewFunction(name, retType, params...)
	m.Funcs = append(m.Funcs, f)
	return f
}
