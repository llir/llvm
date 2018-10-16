package asm

import (
	"fmt"
	"time"

	"github.com/llir/l/ir"
	"github.com/llir/l/ir/types"
	"github.com/mewmew/l-tm/asm/ll/ast"
	"github.com/pkg/errors"
)

// TODO: remove flag after we reach our performance goals.
var (
	// DoTypeResolution enables type resolution of type defintions.
	DoTypeResolution = true
)

// Translate translates the AST of the given module to an equivalent LLVM IR
// module.
func Translate(module *ast.Module) (*ir.Module, error) {
	gen := newGenerator()
	if DoTypeResolution {
		typeResolutionStart := time.Now()
		_, err := gen.resolveTypeDefs(module)
		if err != nil {
			return nil, errors.WithStack(err)
		}
		fmt.Println("type resolution of type definitions took:", time.Since(typeResolutionStart))
		fmt.Println()
	}
	return gen.m, nil
}

// generator keeps track of global and local identifiers when translating values
// and types from AST to IR representation.
type generator struct {
	// LLVM IR module being generated.
	m *ir.Module

	// ts maps from type name to underlying IR type.
	ts map[string]types.Type
}

// newGenerator returns a new generator for translating an LLVM IR module from
// AST to IR representation.
func newGenerator() *generator {
	return &generator{
		m: &ir.Module{},
	}
}
