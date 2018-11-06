package asm

import (
	"github.com/llir/ll/ast"
	asmenum "github.com/llir/llvm/asm/enum"
	"github.com/llir/llvm/internal/enc"
	"github.com/llir/llvm/ir"
	"github.com/pkg/errors"
)

// resolveComdats resolves the global Comdat definitions of the given module.
// The returned value maps from Comdat identifier (without '$' prefix) to the
// corresponding Comdat defintion.
func (gen *generator) resolveComdats(module *ast.Module) (map[string]*ir.ComdatDef, error) {
	// Index and translate Comdat definitions.
	gen.comdats = make(map[string]*ir.ComdatDef)
	for _, entity := range module.TopLevelEntities() {
		switch entity := entity.(type) {
		case *ast.ComdatDef:
			name := comdat(entity.Name())
			comdatDef := &ir.ComdatDef{
				Name: name,
				Kind: asmenum.SelectionKindFromString(entity.Kind().Text()),
			}
			if prev, ok := gen.comdats[name]; ok {
				return nil, errors.Errorf("AST Comdat identifier %q already present; prev `%s`, new `%s`", enc.Comdat(name), prev.Def(), comdatDef.Def())
			}
			gen.comdats[name] = comdatDef
			gen.m.ComdatDefs = append(gen.m.ComdatDefs, comdatDef)
		}
	}
	return gen.comdats, nil
}
