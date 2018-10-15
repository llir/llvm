package asm

import (
	"fmt"

	"github.com/llir/l/ir"
	"github.com/mewmew/l-tm/asm/ll/ast"
)

// Translate translates the AST of the given module to an equivalent LLVM IR
// module.
func Translate(module *ast.Module) *ir.Module {
	m := &ir.Module{}
	for _, entity := range module.TopLevelEntities() {
		fmt.Printf("entity: %T\n", entity)
	}
	return m
}
