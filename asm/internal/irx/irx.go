// Package irx implements utility functions for translating ASTs of LLVM IR
// assembly to equivalent LLVM IR modules.
package irx

import (
	"fmt"

	"github.com/llir/llvm/ir"
	"github.com/llir/llvm/ir/types"
	"github.com/llir/llvm/ir/value"
)

// A Module represents an LLVM IR module generator.
type Module struct {
	// Module being generated.
	*ir.Module
	// types maps from type identifiers to their corresponding LLVM IR types.
	types map[string]*types.NamedType
	// globals maps global identifiers to their corresponding LLVM IR values.
	globals map[string]value.Named
	// locals maps local identifiers to their corresponding LLVM IR values; reset
	// once per function definition.
	locals map[string]value.Named
	// List of errors encountered during translation.
	errs []error
}

// NewModule returns a new module generator.
func NewModule() *Module {
	m := ir.NewModule()
	return &Module{
		Module:  m,
		types:   make(map[string]*types.NamedType),
		globals: make(map[string]value.Named),
	}
}

// getType returns the type of the given type name.
func (m *Module) getType(name string) *types.NamedType {
	typ, ok := m.types[name]
	if !ok {
		panic(fmt.Errorf("unable to locate type name %q", name))
	}
	return typ
}

// getGlobal returns the global value of the given global identifier.
func (m *Module) getGlobal(name string) value.Named {
	global, ok := m.globals[name]
	if !ok {
		panic(fmt.Errorf("unable to locate global identifier %q", name))
	}
	return global
}

// getLocal returns the local value of the given local identifier.
func (m *Module) getLocal(name string) value.Named {
	local, ok := m.locals[name]
	if !ok {
		panic(fmt.Errorf("unable to locate local identifier %q", name))
	}
	return local
}
