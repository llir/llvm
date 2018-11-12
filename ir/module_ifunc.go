package ir

import "github.com/llir/llvm/ir/constant"

// NewIFunc appends a new indirect function to the module based on the given
// IFunc name and resolver.
func (m *Module) NewIFunc(name string, resolver constant.Constant) *IFunc {
	ifunc := NewIFunc(name, resolver)
	m.IFuncs = append(m.IFuncs, ifunc)
	return ifunc
}
