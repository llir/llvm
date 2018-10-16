package asm

import (
	"fmt"

	"github.com/llir/l/ir"
	"github.com/llir/l/ir/value"
	"github.com/mewmew/l-tm/asm/ll/ast"
	"github.com/mewmew/l-tm/internal/enc"
	"github.com/pkg/errors"
)

// resolveGlobals resolves the global variable and function declarations and
// defintions of the given module. The returned value maps from global
// identifier (without '@' prefix) to the corresponding IR value.
func (gen *generator) resolveGlobals(module *ast.Module) (map[string]value.Value, error) {
	// index maps from global identifier to underlying AST value.
	index := make(map[string]ast.LlvmNode) // TODO: see if we can make the type stronger. ast.Value, ast.Constant?
	// Record order of global variable and function declarations and definitions.
	var globalOrder, funcOrder []string
	// Index global variable and function declarations and definitions.
	for _, entity := range module.TopLevelEntities() {
		switch entity := entity.(type) {
		case *ast.GlobalDecl:
			name := global(entity.Name())
			globalOrder = append(globalOrder, name)
			if prev, ok := index[name]; ok {
				// TODO: don't report error if prev is a declaration (of same type)?
				return nil, errors.Errorf("AST global identifier %q already present; prev `%s`, new `%s`", enc.Global(name), prev.Text(), entity.Text())
			}
			index[name] = entity
		case *ast.GlobalDef:
			name := global(entity.Name())
			globalOrder = append(globalOrder, name)
			if prev, ok := index[name]; ok {
				// TODO: don't report error if prev is a declaration (of same type)?
				return nil, errors.Errorf("AST global identifier %q already present; prev `%s`, new `%s`", enc.Global(name), prev.Text(), entity.Text())
			}
			index[name] = entity
		case *ast.FuncDecl:
			name := global(entity.Header().Name())
			funcOrder = append(funcOrder, name)
			if prev, ok := index[name]; ok {
				// TODO: don't report error if prev is a declaration (of same type)?
				return nil, errors.Errorf("AST global identifier %q already present; prev `%s`, new `%s`", enc.Global(name), prev.Text(), entity.Text())
			}
			index[name] = entity
		case *ast.FuncDef:
			name := global(entity.Header().Name())
			funcOrder = append(funcOrder, name)
			if prev, ok := index[name]; ok {
				// TODO: don't report error if prev is a declaration (of same type)?
				return nil, errors.Errorf("AST global identifier %q already present; prev `%s`, new `%s`", enc.Global(name), prev.Text(), entity.Text())
			}
			index[name] = entity
			// TODO: handle alias definitions and IFuncs.
			//case *ast.AliasDef:
			//case *ast.IFuncDef:
		}
	}

	// Create corresponding IR global variables and functions (without bodies).
	gen.gs = make(map[string]value.Value)
	for name, old := range index {
		g, err := gen.newGlobal(name, old)
		if err != nil {
			return nil, errors.WithStack(err)
		}
		gen.gs[name] = g
	}

	// Translate global variables and functions (including bodies).
	for name, old := range index {
		g := gen.gs[name]
		_, err := gen.translateGlobal(g, old)
		if err != nil {
			return nil, errors.WithStack(err)
		}
	}

	// Add global variable declarations and definitions to IR module in order of
	// occurrence in input.
	for _, key := range globalOrder {
		g, ok := gen.gs[key].(*ir.Global)
		if !ok {
			// NOTE: Panic instead of returning error as this case should not be
			// possible, and would indicate a bug in the implementation.
			panic(fmt.Errorf("invalid IR type for AST global variable declaration or definition; expected *ir.Global, got %T", g))
		}
		gen.m.Globals = append(gen.m.Globals, g)
	}

	// Add function declarations and definitions to IR module in order of
	// occurrence in input.
	for _, key := range funcOrder {
		f, ok := gen.gs[key].(*ir.Function)
		if !ok {
			// NOTE: Panic instead of returning error as this case should not be
			// possible, and would indicate a bug in the implementation.
			panic(fmt.Errorf("invalid IR type for AST function declaration or definition; expected *ir.Function, got %T", f))
		}
		gen.m.Funcs = append(gen.m.Funcs, f)
	}

	return gen.gs, nil
}

// newGlobal returns a new IR value (without body but with a type) based on the
// given AST global variable or function.
func (gen *generator) newGlobal(name string, old ast.LlvmNode) (value.Value, error) {
	switch old := old.(type) {
	case *ast.GlobalDecl:
		// TODO: Add type.
		return &ir.Global{GlobalName: name}, nil
	case *ast.GlobalDef:
		// TODO: Add type.
		return &ir.Global{GlobalName: name}, nil
	case *ast.FuncDecl:
		// TODO: Add type.
		return &ir.Function{FuncName: name}, nil
	case *ast.FuncDef:
		// TODO: Add type.
		return &ir.Function{FuncName: name}, nil
	default:
		panic(fmt.Errorf("support for global variable or function %T not yet implemented", old))
	}
}

// translateGlobal translates the AST global variable or function into an
// equivalent IR value.
func (gen *generator) translateGlobal(g value.Value, old ast.LlvmNode) (value.Value, error) {
	switch old := old.(type) {
	case *ast.GlobalDecl:
		return gen.translateGlobalDecl(g, old)
	case *ast.GlobalDef:
		return gen.translateGlobalDef(g, old)
	case *ast.FuncDecl:
		return gen.translateFuncDecl(g, old)
	case *ast.FuncDef:
		return gen.translateFuncDef(g, old)
	default:
		panic(fmt.Errorf("support for type %T not yet implemented", old))
	}
}

// ~~~ [ Global Variable Declaration ] ~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~

func (gen *generator) translateGlobalDecl(g value.Value, old *ast.GlobalDecl) (value.Value, error) {
	global, ok := g.(*ir.Global)
	if !ok {
		panic(fmt.Errorf("invalid IR type for AST global declaration; expected *ir.Global, got %T", g))
	}
	// TODO: implement
	return global, nil
}

// ~~~ [ Global Variable Definition ] ~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~

func (gen *generator) translateGlobalDef(g value.Value, old *ast.GlobalDef) (value.Value, error) {
	global, ok := g.(*ir.Global)
	if !ok {
		panic(fmt.Errorf("invalid IR type for AST global definition; expected *ir.Global, got %T", g))
	}
	// TODO: implement
	return global, nil
}

// ~~~ [ Indirect Symbol Definition ] ~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~

// TODO: add alias definition and IFuncs.

// ~~~ [ Function Declaration ] ~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~

func (gen *generator) translateFuncDecl(g value.Value, old *ast.FuncDecl) (value.Value, error) {
	f, ok := g.(*ir.Function)
	if !ok {
		panic(fmt.Errorf("invalid IR type for AST function declaration; expected *ir.Function, got %T", g))
	}
	// TODO: implement
	return f, nil
}

// ~~~ [ Function Definition ] ~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~

func (gen *generator) translateFuncDef(g value.Value, old *ast.FuncDef) (value.Value, error) {
	f, ok := g.(*ir.Function)
	if !ok {
		panic(fmt.Errorf("invalid IR type for AST function definition; expected *ir.Function, got %T", g))
	}
	// TODO: implement
	return f, nil
}
