package asm

import (
	"fmt"

	"github.com/llir/l/ir"
	"github.com/llir/l/ir/types"
	"github.com/mewmew/l-tm/asm/ll/ast"
	"github.com/mewmew/l-tm/internal/enc"
	"github.com/pkg/errors"
)

// resolveGlobals resolves the global variable and function declarations and
// defintions of the given module. The returned value maps from global
// identifier (without '@' prefix) to the corresponding IR value.
func (gen *generator) resolveGlobals(module *ast.Module) (map[string]ir.Constant, error) {
	// index maps from global identifier to underlying AST value.
	index := make(map[string]ast.LlvmNode)
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

	// Create corresponding IR global variables and functions (without bodies but
	// with a type).
	gen.gs = make(map[string]ir.Constant)
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
		g, err := gen.global(key)
		if err != nil {
			// NOTE: panic since this would indicate a bug in the implementation.
			panic(err)
		}
		gen.m.Globals = append(gen.m.Globals, g)
	}

	// Add function declarations and definitions to IR module in order of
	// occurrence in input.
	for _, key := range funcOrder {
		f, err := gen.function(key)
		if err != nil {
			// NOTE: panic since this would indicate a bug in the implementation.
			panic(err)
		}
		gen.m.Funcs = append(gen.m.Funcs, f)
	}

	return gen.gs, nil
}

// newGlobal returns a new IR value (without body but with a type) based on the
// given AST global variable or function.
func (gen *generator) newGlobal(name string, old ast.LlvmNode) (ir.Constant, error) {
	switch old := old.(type) {
	case *ast.GlobalDecl:
		g := &ir.Global{GlobalName: name}
		// Content type.
		contentType, err := gen.irType(old.ContentType())
		if err != nil {
			return nil, errors.WithStack(err)
		}
		g.ContentType = contentType
		g.Typ = types.NewPointer(g.ContentType)
		return g, nil
	case *ast.GlobalDef:
		g := &ir.Global{GlobalName: name}
		// Content type.
		contentType, err := gen.irType(old.ContentType())
		if err != nil {
			return nil, errors.WithStack(err)
		}
		g.ContentType = contentType
		g.Typ = types.NewPointer(g.ContentType)
		return g, nil
	case *ast.FuncDecl:
		f := &ir.Function{FuncName: name}
		hdr := old.Header()
		sig := &types.FuncType{}
		// Return type.
		retType, err := gen.irType(hdr.RetType())
		if err != nil {
			return nil, errors.WithStack(err)
		}
		sig.RetType = retType
		// Function parameters.
		ps := hdr.Params()
		for _, p := range ps.Params() {
			param, err := gen.irType(p.Typ())
			if err != nil {
				return nil, errors.WithStack(err)
			}
			sig.Params = append(sig.Params, param)
		}
		// Variadic.
		sig.Variadic = irVariadic(ps.Variadic())
		f.Sig = sig
		f.Typ = types.NewPointer(f.Sig)
		return f, nil
	case *ast.FuncDef:
		f := &ir.Function{FuncName: name}
		sig := &types.FuncType{}
		hdr := old.Header()
		// Return type.
		retType, err := gen.irType(hdr.RetType())
		if err != nil {
			return nil, errors.WithStack(err)
		}
		sig.RetType = retType
		// Function parameters.
		ps := hdr.Params()
		for _, p := range ps.Params() {
			param, err := gen.irType(p.Typ())
			if err != nil {
				return nil, errors.WithStack(err)
			}
			sig.Params = append(sig.Params, param)
		}
		// Variadic.
		sig.Variadic = irVariadic(ps.Variadic())
		f.Sig = sig
		f.Typ = types.NewPointer(f.Sig)
		return f, nil
	default:
		panic(fmt.Errorf("support for global variable or function %T not yet implemented", old))
	}
}

// translateGlobal translates the AST global variable or function into an
// equivalent IR value.
func (gen *generator) translateGlobal(g ir.Constant, old ast.LlvmNode) (ir.Constant, error) {
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

func (gen *generator) translateGlobalDecl(g ir.Constant, old *ast.GlobalDecl) (ir.Constant, error) {
	global, ok := g.(*ir.Global)
	if !ok {
		panic(fmt.Errorf("invalid IR type for AST global declaration; expected *ir.Global, got %T", g))
	}
	// Linkage.
	global.Linkage = irLinkage(old.ExternLinkage().Text())
	// Preemption.
	global.Preemption = irPreemption(old.Preemption())
	// Visibility.
	global.Visibility = irVisibility(old.Visibility())
	// DLL storage class.
	global.DLLStorageClass = irDLLStorageClass(old.DLLStorageClass())
	// Thread local storage model.
	global.TLSModel = irTLSModelFromThreadLocal(old.ThreadLocal())
	// Unnamed address.
	global.UnnamedAddr = irUnnamedAddr(old.UnnamedAddr())
	// Address space.
	global.Typ.AddrSpace = irAddrSpace(old.AddrSpace())
	// Externally initialized.
	global.ExternallyInitialized = irExternallyInitialized(old.ExternallyInitialized())
	// Immutable (constant or global).
	global.Immutable = irImmutable(old.Immutable())
	// Content type already stored during index.
	// TODO: handle GlobalAttrs.
	// TODO: handle FuncAttrs.
	return global, nil
}

// ~~~ [ Global Variable Definition ] ~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~

func (gen *generator) translateGlobalDef(g ir.Constant, old *ast.GlobalDef) (ir.Constant, error) {
	global, ok := g.(*ir.Global)
	if !ok {
		panic(fmt.Errorf("invalid IR type for AST global definition; expected *ir.Global, got %T", g))
	}
	// Linkage.
	global.Linkage = irLinkage(old.Linkage().Text())
	// Preemption.
	global.Preemption = irPreemption(old.Preemption())
	// Visibility.
	global.Visibility = irVisibility(old.Visibility())
	// DLL storage class.
	global.DLLStorageClass = irDLLStorageClass(old.DLLStorageClass())
	// Thread local storage model.
	global.TLSModel = irTLSModelFromThreadLocal(old.ThreadLocal())
	// Unnamed address.
	global.UnnamedAddr = irUnnamedAddr(old.UnnamedAddr())
	// Address space.
	global.Typ.AddrSpace = irAddrSpace(old.AddrSpace())
	// Externally initialized.
	global.ExternallyInitialized = irExternallyInitialized(old.ExternallyInitialized())
	// Immutable (constant or global).
	global.Immutable = irImmutable(old.Immutable())
	// Content type already stored during index.
	init, err := gen.irConstant(global.ContentType, old.Init())
	if err != nil {
		return nil, errors.WithStack(err)
	}
	global.Init = init
	// TODO: handle GlobalAttrs.
	// TODO: handle FuncAttrs.
	return global, nil
}

// ~~~ [ Indirect Symbol Definition ] ~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~

// TODO: add alias definition and IFuncs.

// ~~~ [ Function Declaration ] ~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~

func (gen *generator) translateFuncDecl(g ir.Constant, old *ast.FuncDecl) (ir.Constant, error) {
	f, ok := g.(*ir.Function)
	if !ok {
		panic(fmt.Errorf("invalid IR type for AST function declaration; expected *ir.Function, got %T", g))
	}
	// Metadata.
	// TODO: translate function metadata.
	if err := gen.translateFuncHeader(f, old.Header()); err != nil {
		return nil, errors.WithStack(err)
	}
	return f, nil
}

func (gen *generator) translateFuncHeader(f *ir.Function, hdr ast.FuncHeader) error {
	// Linkage.
	f.Linkage = irLinkage(hdr.ExternLinkage().Text())
	// Preemption.
	f.Preemption = irPreemption(hdr.Preemption())
	// Visibility.
	f.Visibility = irVisibility(hdr.Visibility())
	// DLL storage class.
	f.DLLStorageClass = irDLLStorageClass(hdr.DLLStorageClass())
	// Calling convention.
	// TODO: translate CallingConv.
	// Return attributes.
	// TODO: handle ReturnAttrs.
	// Return type; already handled.
	// Function name; already handled.
	// Function parameters.
	ps := hdr.Params()
	for _, p := range ps.Params() {
		// Type.
		typ, err := gen.irType(p.Typ())
		if err != nil {
			return errors.WithStack(err)
		}
		// Parameter attributes.
		// TODO: handle Attrs.
		name := local(*p.Name())
		param := ir.NewParam(typ, name)
		f.Params = append(f.Params, param)
	}

	// Unnamed address.
	f.UnnamedAddr = irUnnamedAddr(hdr.UnnamedAddr())
	// Address space.
	f.Typ.AddrSpace = irAddrSpace(hdr.AddrSpace())
	// Function attributes.
	// TODO: handle FuncAttrs.
	// Section.
	// TODO: handle Section.
	// Comdat.
	// TODO: handle Comdat.
	// GC.
	// TODO: handle GC.
	// Prefix.
	// TODO: handle Prefix.
	// Prologue.
	// TODO: handle Prologue.
	// Personality.
	// TODO: handle Personality.
	return nil
}

// ~~~ [ Function Definition ] ~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~

func (gen *generator) translateFuncDef(g ir.Constant, old *ast.FuncDef) (ir.Constant, error) {
	f, ok := g.(*ir.Function)
	if !ok {
		panic(fmt.Errorf("invalid IR type for AST function definition; expected *ir.Function, got %T", g))
	}
	if err := gen.translateFuncHeader(f, old.Header()); err != nil {
		return nil, errors.WithStack(err)
	}
	// Metadata.
	// TODO: translate function metadata.
	// Basic blocks.
	fgen := newFuncGen(gen, f)
	_, err := fgen.resolveLocals(old.Body())
	if err != nil {
		return nil, errors.WithStack(err)
	}
	// Use list orders.
	// TODO: translate use list orders.
	return f, nil
}
