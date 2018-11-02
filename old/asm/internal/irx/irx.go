// Package irx implements utility functions for translating ASTs of LLVM IR
// assembly to equivalent LLVM IR modules.
package irx

import (
	"fmt"

	"github.com/llir/llvm/internal/enc"
	"github.com/llir/llvm/ir"
	"github.com/llir/llvm/ir/metadata"
	"github.com/llir/llvm/ir/types"
	"github.com/llir/llvm/ir/value"
)

// A Module represents an LLVM IR module generator.
type Module struct {
	// Module being generated.
	*ir.Module

	// Per module.

	// types maps from type identifiers to their corresponding LLVM IR types.
	types map[string]types.Type
	// globals maps global identifiers to their corresponding LLVM IR values.
	globals map[string]value.Named
	// metadata maps metadata IDs to their corresponding LLVM IR metadata.
	metadata map[string]*metadata.Metadata

	// Per function.

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
		Module:   m,
		types:    make(map[string]types.Type),
		globals:  make(map[string]value.Named),
		metadata: make(map[string]*metadata.Metadata),
	}
}

// getType returns the type of the given type name.
func (m *Module) getType(name string) types.Type {
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

// getMetadata returns the metadata of the given metadata ID.
func (m *Module) getMetadata(id string) *metadata.Metadata {
	metadata, ok := m.metadata[id]
	if !ok {
		panic(fmt.Errorf("unable to locate metadata ID %q", enc.Metadata(id)))
	}
	return metadata
}

// getLocal returns the local value of the given local identifier.
func (m *Module) getLocal(name string) value.Named {
	local, ok := m.locals[name]
	if !ok {
		panic(fmt.Errorf("unable to locate local identifier %q", name))
	}
	return local
}
