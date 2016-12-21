// Package irx implements utility functions for translating ASTs of LLVM IR
// assembly to equivalent LLVM IR modules.
package irx

import (
	"github.com/llir/llvm/asm/internal/ast"
	"github.com/llir/llvm/ir"
)

// Translate translates the AST of the given module to an equivalent LLVM IR
// module.
func Translate(module *ast.Module) (*ir.Module, error) {
	m := ir.NewModule()
	return m, nil
}
