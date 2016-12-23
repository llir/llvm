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

// A Function represents an LLVM IR function generator.
type Function struct {
	// Function being generated.
	*ir.Function
	// Current basic block being generated.
	curBlock *BasicBlock
	// locals maps local identifiers to their corresponding LLVM IR values.
	locals map[string]value.Named
}

// NewFunction returns a new function generator based on the given function name
// and signature.
//
// The caller is responsible for initializing basic blocks.
func NewFunction(name string, ret types.Type, params ...*ir.Param) *Function {
	f := ir.NewFunction(name, ret, params...)
	return &Function{
		Function: f,
		locals:   make(map[string]value.Named),
	}
}

// NewBlock returns a new basic block generator based on the given name and
// parent function.
func (f *Function) NewBlock(name string) *BasicBlock {
	block := ir.NewBlock(name)
	return &BasicBlock{
		BasicBlock: block,
		parent:     f,
	}
}

// A BasicBlock represents an LLVM IR basic block generator.
type BasicBlock struct {
	// Basic block being generated.
	*ir.BasicBlock
	// Parent function of the basic block.
	parent *Function
}

// SetTerm sets the terminator of the basic block.
func (block *BasicBlock) SetTerm(term ir.Terminator) {
	if block.Term != nil {
		panic(fmt.Errorf("terminator instruction already set for basic block; old term (%v), new term (%v), basic block (%v)", term, block.Term, block))
	}
	block.BasicBlock.Term = term
	block.parent.AppendBlock(block.BasicBlock)
}
