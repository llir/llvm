package ir

import (
	"github.com/llir/llvm/ir/constant"
	"github.com/llir/llvm/ir/types"
)

// --- [ Global variables ] ----------------------------------------------------

// NewGlobalDecl appends a new global variable declaration to the module based
// on the given global variable name and content type.
func (m *Module) NewGlobalDecl(name string, contentType types.Type) *Global {
	g := NewGlobalDecl(name, contentType)
	m.Globals = append(m.Globals, g)
	return g
}

// NewGlobalDef appends a new global variable definition to the module based on
// the given global variable name and initial value.
func (m *Module) NewGlobalDef(name string, init constant.Constant) *Global {
	g := NewGlobalDef(name, init)
	m.Globals = append(m.Globals, g)
	return g
}
