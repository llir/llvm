package asm

import (
	"fmt"
	"time"

	"github.com/llir/ll/ast"
	"github.com/llir/llvm/ir"
	"github.com/llir/llvm/ir/types"
	"github.com/pkg/errors"
)

// TODO: remove flag after we reach our performance goals.
var (
	// DoTypeResolution enables type resolution of type defintions.
	DoTypeResolution = true
	// DoGlobalResolution enables global resolution of global variable and
	// function delcarations and defintions.
	DoGlobalResolution = true
)

// translate translates the AST of the given module to an equivalent LLVM IR
// module.
func translate(module *ast.Module) (*ir.Module, error) {
	gen := newGenerator()
	// Resolve types.
	if DoTypeResolution {
		typeResolutionStart := time.Now()
		_, err := gen.resolveTypeDefs(module)
		if err != nil {
			return nil, errors.WithStack(err)
		}
		fmt.Println("type resolution of type definitions took:", time.Since(typeResolutionStart))
		fmt.Println()
	}
	// Resolve globals.
	if DoGlobalResolution {
		globalResolutionStart := time.Now()
		_, err := gen.resolveGlobals(module)
		if err != nil {
			return nil, errors.WithStack(err)
		}
		fmt.Println("global resolution of global variable and function declarations and definitions took:", time.Since(globalResolutionStart))
		fmt.Println()
	}
	// Resolve functions.
	// TODO: implement.
	// Fix dummy values.
	for _, c := range gen.todo {
		if err := fixBlockAddressConst(c); err != nil {
			return nil, errors.WithStack(err)
		}
	}
	return gen.m, nil
}

// generator keeps track of global and local identifiers when translating values
// and types from AST to IR representation.
type generator struct {
	// LLVM IR module being generated.
	m *ir.Module

	// ts maps from type name (without '%' prefix) to underlying IR type.
	ts map[string]types.Type

	// gs maps from global identifier (without '@' prefix) to corresponding
	// IR value.
	gs map[string]ir.Constant

	// TOOD: add rw mutex to gen.todo for access to blockaddress constant.

	// Fix dummy basic blocks after translation of function bodies and assignment
	// of local IDs.
	todo []*ir.ConstBlockAddress
}

// newGenerator returns a new generator for translating an LLVM IR module from
// AST to IR representation.
func newGenerator() *generator {
	return &generator{
		m: &ir.Module{},
	}
}

// global returns the IR global variable of the given name.
func (gen *generator) global(name string) (*ir.Global, error) {
	v, ok := gen.gs[name]
	if !ok {
		return nil, errors.Errorf("unable to locate global variable %q", name)
	}
	g, ok := v.(*ir.Global)
	if !ok {
		return nil, errors.Errorf("invalid global variable type of %q; expected *ir.Global, got %T", name, v)
	}
	return g, nil
}

// function returns the IR function of the given name.
func (gen *generator) function(name string) (*ir.Function, error) {
	v, ok := gen.gs[name]
	if !ok {
		return nil, errors.Errorf("unable to locate function %q", name)
	}
	f, ok := v.(*ir.Function)
	if !ok {
		return nil, errors.Errorf("invalid function type of %q; expected *ir.Function, got %T", name, v)
	}
	return f, nil
}
