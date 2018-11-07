package asm

import (
	"fmt"

	"github.com/llir/ll/ast"
	asmenum "github.com/llir/llvm/asm/enum"
	"github.com/llir/llvm/internal/enc"
	"github.com/llir/llvm/ir"
	"github.com/llir/llvm/ir/constant"
	"github.com/llir/llvm/ir/types"
	"github.com/pkg/errors"
)

// resolveGlobals resolves the global variable and function declarations and
// defintions of the given module. The returned value maps from global
// identifier (without '@' prefix) to the corresponding IR value.
func (gen *generator) resolveGlobals(module *ast.Module) (map[string]constant.Constant, error) {
	// index maps from global identifier to underlying AST value.
	index := make(map[string]ast.LlvmNode)
	// Record order of global variable and function declarations and definitions.
	var globalOrder, funcOrder, aliasOrder, ifuncOrder []string
	// Index global variable and function declarations and definitions.
	for _, entity := range module.TopLevelEntities() {
		switch entity := entity.(type) {
		case *ast.GlobalDecl:
			name := global(entity.Name())
			globalOrder = append(globalOrder, name)
			if prev, ok := index[name]; ok {
				// TODO: don't report error if prev is a declaration (of same type)?
				return nil, errors.Errorf("AST global identifier %q already present; prev `%s`, new `%s`", enc.Global(name), text(prev), text(entity))
			}
			index[name] = entity
		case *ast.GlobalDef:
			name := global(entity.Name())
			globalOrder = append(globalOrder, name)
			if prev, ok := index[name]; ok {
				// TODO: don't report error if prev is a declaration (of same type)?
				return nil, errors.Errorf("AST global identifier %q already present; prev `%s`, new `%s`", enc.Global(name), text(prev), text(entity))
			}
			index[name] = entity
		case *ast.AliasDef:
			name := global(entity.Name())
			aliasOrder = append(aliasOrder, name)
			if prev, ok := index[name]; ok {
				// TODO: don't report error if prev is a declaration (of same type)?
				return nil, errors.Errorf("AST global identifier %q already present; prev `%s`, new `%s`", enc.Global(name), text(prev), text(entity))
			}
			index[name] = entity
		case *ast.IFuncDef:
			name := global(entity.Name())
			ifuncOrder = append(ifuncOrder, name)
			if prev, ok := index[name]; ok {
				// TODO: don't report error if prev is a declaration (of same type)?
				return nil, errors.Errorf("AST global identifier %q already present; prev `%s`, new `%s`", enc.Global(name), text(prev), text(entity))
			}
			index[name] = entity
		case *ast.FuncDecl:
			name := global(entity.Header().Name())
			funcOrder = append(funcOrder, name)
			if prev, ok := index[name]; ok {
				// TODO: don't report error if prev is a declaration (of same type)?
				return nil, errors.Errorf("AST global identifier %q already present; prev `%s`, new `%s`", enc.Global(name), text(prev), text(entity))
			}
			index[name] = entity
		case *ast.FuncDef:
			name := global(entity.Header().Name())
			funcOrder = append(funcOrder, name)
			if prev, ok := index[name]; ok {
				// TODO: don't report error if prev is a declaration (of same type)?
				return nil, errors.Errorf("AST global identifier %q already present; prev `%s`, new `%s`", enc.Global(name), text(prev), text(entity))
			}
			index[name] = entity
		}
	}

	// Create corresponding IR global variables and functions (without bodies but
	// with type).
	gen.gs = make(map[string]constant.Constant)
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
		_, err := gen.astToIRGlobal(g, old)
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

	// Add alias definitions to IR module in order of occurrence in input.
	for _, key := range aliasOrder {
		a, err := gen.alias(key)
		if err != nil {
			// NOTE: panic since this would indicate a bug in the implementation.
			panic(err)
		}
		gen.m.Aliases = append(gen.m.Aliases, a)
	}

	// Add IFunc definitions to IR module in order of occurrence in input.
	for _, key := range ifuncOrder {
		i, err := gen.ifunc(key)
		if err != nil {
			// NOTE: panic since this would indicate a bug in the implementation.
			panic(err)
		}
		gen.m.IFuncs = append(gen.m.IFuncs, i)
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

// newGlobal returns a new IR value (without body but with type) based on the
// given AST global variable or function.
func (gen *generator) newGlobal(name string, old ast.LlvmNode) (constant.Constant, error) {
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
	case *ast.AliasDef:
		alias := &ir.Alias{GlobalName: name}
		// Content type.
		contentType, err := gen.irType(old.ContentType())
		if err != nil {
			return nil, errors.WithStack(err)
		}
		alias.Typ = types.NewPointer(contentType)
		return alias, nil
	case *ast.IFuncDef:
		ifunc := &ir.IFunc{GlobalName: name}
		// Content type.
		contentType, err := gen.irType(old.ContentType())
		if err != nil {
			return nil, errors.WithStack(err)
		}
		ifunc.Typ = types.NewPointer(contentType)
		return ifunc, nil
	case *ast.FuncDecl:
		f := &ir.Function{GlobalName: name}
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
		sig.Variadic = ps.Variadic() != nil
		f.Sig = sig
		f.Typ = types.NewPointer(f.Sig)
		return f, nil
	case *ast.FuncDef:
		f := &ir.Function{GlobalName: name}
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
		sig.Variadic = ps.Variadic() != nil
		f.Sig = sig
		f.Typ = types.NewPointer(f.Sig)
		return f, nil
	default:
		panic(fmt.Errorf("support for global variable or function %T not yet implemented", old))
	}
}

// astToIRGlobal translates the AST global variable or function into an
// equivalent IR value.
func (gen *generator) astToIRGlobal(g constant.Constant, old ast.LlvmNode) (constant.Constant, error) {
	switch old := old.(type) {
	case *ast.GlobalDecl:
		return gen.astToIRGlobalDecl(g, old)
	case *ast.GlobalDef:
		return gen.astToIRGlobalDef(g, old)
	case *ast.AliasDef:
		return gen.astToIRAliasDef(g, old)
	case *ast.IFuncDef:
		return gen.astToIRIFuncDef(g, old)
	case *ast.FuncDecl:
		return gen.astToIRFuncDecl(g, old)
	case *ast.FuncDef:
		return gen.astToIRFuncDef(g, old)
	default:
		panic(fmt.Errorf("support for type %T not yet implemented", old))
	}
}

// ~~~ [ Global Variable Declaration ] ~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~

// astToIRGlobalDecl translates the given AST global declaration into an
// equivalent IR global declaration.
func (gen *generator) astToIRGlobalDecl(global constant.Constant, old *ast.GlobalDecl) (*ir.Global, error) {
	g, ok := global.(*ir.Global)
	if !ok {
		panic(fmt.Errorf("invalid IR type for AST global declaration; expected *ir.Global, got %T", global))
	}
	// (optional) Linkage.
	g.Linkage = asmenum.LinkageFromString(old.ExternLinkage().Text())
	// (optional) Preemption.
	if n := old.Preemption(); n != nil {
		g.Preemption = asmenum.PreemptionFromString(n.Text())
	}
	// (optional) Visibility.
	if n := old.Visibility(); n != nil {
		g.Visibility = asmenum.VisibilityFromString(n.Text())
	}
	// (optional) DLL storage class.
	if n := old.DLLStorageClass(); n != nil {
		g.DLLStorageClass = asmenum.DLLStorageClassFromString(n.Text())
	}
	// (optional) Thread local storage model.
	if n := old.ThreadLocal(); n != nil {
		g.TLSModel = irTLSModelFromThreadLocal(*n)
	}
	// (optional) Unnamed address.
	if n := old.UnnamedAddr(); n != nil {
		g.UnnamedAddr = asmenum.UnnamedAddrFromString(n.Text())
	}
	// (optional) Address space; stored in g.Typ.
	if n := old.AddrSpace(); n != nil {
		g.Typ.AddrSpace = irAddrSpace(*n)
	}
	// (optional) Externally initialized.
	g.ExternallyInitialized = old.ExternallyInitialized() != nil
	// Immutability of global variable (constant or global).
	g.Immutable = irImmutable(old.Immutable())
	// Content type: already stored during index.
	// ### [ Global attributes ] ###
	// TODO: handle GlobalAttrs.
	// (optional) Section name.
	// (optional) Comdat.
	if n := old.Comdat(); n != nil {
		comdatName := g.GlobalName
		if n := n.Name(); n != nil {
			comdatName = comdat(*n)
		}
		comdatDef, ok := gen.comdats[comdatName]
		if !ok {
			return nil, errors.Errorf("unable to locate Comdat identifier %q used in global declaration of %q", comdatName, g.GlobalName)
		}
		g.Comdat = comdatDef
	}
	// (optional) Alignment.
	// ### [/ Global attributes ] ###
	// (optional) Function attributes.
	for _, oldFuncAttr := range old.FuncAttrs() {
		funcAttr := irFuncAttribute(oldFuncAttr)
		g.FuncAttrs = append(g.FuncAttrs, funcAttr)
	}
	// (optional) Metadata.
	// TODO: handle metadata.
	return g, nil
}

// ~~~ [ Global Variable Definition ] ~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~

// astToIRGlobalDef translates the given AST global definition into an
// equivalent IR global definition.
func (gen *generator) astToIRGlobalDef(global constant.Constant, old *ast.GlobalDef) (*ir.Global, error) {
	g, ok := global.(*ir.Global)
	if !ok {
		panic(fmt.Errorf("invalid IR type for AST global definition; expected *ir.Global, got %T", global))
	}
	// (optional) Linkage.
	if n := old.Linkage(); n != nil {
		g.Linkage = asmenum.LinkageFromString(n.Text())
	}
	// (optional) Preemption.
	if n := old.Preemption(); n != nil {
		g.Preemption = asmenum.PreemptionFromString(n.Text())
	}
	// (optional) Visibility.
	if n := old.Visibility(); n != nil {
		g.Visibility = asmenum.VisibilityFromString(n.Text())
	}
	// (optional) DLL storage class.
	if n := old.DLLStorageClass(); n != nil {
		g.DLLStorageClass = asmenum.DLLStorageClassFromString(n.Text())
	}
	// (optional) Thread local storage model.
	if n := old.ThreadLocal(); n != nil {
		g.TLSModel = irTLSModelFromThreadLocal(*n)
	}
	// (optional) Unnamed address.
	if n := old.UnnamedAddr(); n != nil {
		g.UnnamedAddr = asmenum.UnnamedAddrFromString(n.Text())
	}
	// (optional) Address space; stored in g.Typ.
	if n := old.AddrSpace(); n != nil {
		g.Typ.AddrSpace = irAddrSpace(*n)
	}
	// (optional) Externally initialized.
	g.ExternallyInitialized = old.ExternallyInitialized() != nil
	// Immutability of global variable (constant or global).
	g.Immutable = irImmutable(old.Immutable())
	// Content type: already stored during index.
	// Initial value.
	init, err := gen.irConstant(g.ContentType, old.Init())
	if err != nil {
		return nil, errors.WithStack(err)
	}
	g.Init = init
	// ### [ Global attributes ] ###
	// TODO: handle GlobalAttrs.
	// (optional) Section name.
	// (optional) Comdat.
	if n := old.Comdat(); n != nil {
		comdatName := g.GlobalName
		if n := n.Name(); n != nil {
			comdatName = comdat(*n)
		}
		comdatDef, ok := gen.comdats[comdatName]
		if !ok {
			return nil, errors.Errorf("unable to locate Comdat identifier %q used in global definition of %q", comdatName, g.GlobalName)
		}
		g.Comdat = comdatDef
	}
	// (optional) Alignment.
	// ### [/ Global attributes ] ###
	// (optional) Function attributes.
	for _, oldFuncAttr := range old.FuncAttrs() {
		funcAttr := irFuncAttribute(oldFuncAttr)
		g.FuncAttrs = append(g.FuncAttrs, funcAttr)
	}
	// (optional) Metadata.
	// TODO: handle metadata.
	return g, nil
}

// ~~~ [ Alias Definition ] ~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~

// astToIRAliasDef translates the given AST alias definition into an equivalent
// IR alias definition.
func (gen *generator) astToIRAliasDef(alias constant.Constant, old *ast.AliasDef) (*ir.Alias, error) {
	a, ok := alias.(*ir.Alias)
	if !ok {
		panic(fmt.Errorf("invalid IR type for AST alias definition; expected *ir.Alias, got %T", alias))
	}
	// (optional) Linkage.
	// TODO: handle external linkage.
	if n := old.Linkage(); n != nil {
		a.Linkage = asmenum.LinkageFromString(n.Text())
	}
	// (optional) Preemption.
	if n := old.Preemption(); n != nil {
		a.Preemption = asmenum.PreemptionFromString(n.Text())
	}
	// (optional) Visibility.
	if n := old.Visibility(); n != nil {
		a.Visibility = asmenum.VisibilityFromString(n.Text())
	}
	// (optional) DLL storage class.
	if n := old.DLLStorageClass(); n != nil {
		a.DLLStorageClass = asmenum.DLLStorageClassFromString(n.Text())
	}
	// (optional) Thread local storage model.
	if n := old.ThreadLocal(); n != nil {
		a.TLSModel = irTLSModelFromThreadLocal(*n)
	}
	// (optional) Unnamed address.
	if n := old.UnnamedAddr(); n != nil {
		a.UnnamedAddr = asmenum.UnnamedAddrFromString(n.Text())
	}
	// Content type: already stored during index.
	// Aliasee.
	aliasee, err := gen.irTypeConst(old.Aliasee())
	if err != nil {
		return nil, errors.WithStack(err)
	}
	a.Aliasee = aliasee
	return a, nil
}

// ~~~ [ IFunc Definition ] ~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~

// astToIRIFuncDef translates the given AST IFunc definition into an equivalent
// IR IFunc definition.
func (gen *generator) astToIRIFuncDef(ifunc constant.Constant, old *ast.IFuncDef) (*ir.IFunc, error) {
	i, ok := ifunc.(*ir.IFunc)
	if !ok {
		panic(fmt.Errorf("invalid IR type for AST IFunc definition; expected *ir.IFunc, got %T", ifunc))
	}
	// (optional) Linkage.
	// TODO: handle external linkage.
	if n := old.Linkage(); n != nil {
		i.Linkage = asmenum.LinkageFromString(n.Text())
	}
	// (optional) Preemption.
	if n := old.Preemption(); n != nil {
		i.Preemption = asmenum.PreemptionFromString(n.Text())
	}
	// (optional) Visibility.
	if n := old.Visibility(); n != nil {
		i.Visibility = asmenum.VisibilityFromString(n.Text())
	}
	// (optional) DLL storage class.
	if n := old.DLLStorageClass(); n != nil {
		i.DLLStorageClass = asmenum.DLLStorageClassFromString(n.Text())
	}
	// (optional) Thread local storage model.
	if n := old.ThreadLocal(); n != nil {
		i.TLSModel = irTLSModelFromThreadLocal(*n)
	}
	// (optional) Unnamed address.
	if n := old.UnnamedAddr(); n != nil {
		i.UnnamedAddr = asmenum.UnnamedAddrFromString(n.Text())
	}
	// Content type: already stored during index.
	// Resolver.
	resolver, err := gen.irTypeConst(old.Resolver())
	if err != nil {
		return nil, errors.WithStack(err)
	}
	i.Resolver = resolver
	return i, nil
}

// ~~~ [ Function Declaration ] ~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~

// astToIRFuncDecl translates the given AST function declaration into an
// equivalent IR function declaration.
func (gen *generator) astToIRFuncDecl(fn constant.Constant, old *ast.FuncDecl) (*ir.Function, error) {
	f, ok := fn.(*ir.Function)
	if !ok {
		panic(fmt.Errorf("invalid IR type for AST function declaration; expected *ir.Function, got %T", fn))
	}
	// (optional) Metadata.
	f.Metadata = irMetadataAttachments(old.Metadata())
	// Function header.
	if err := gen.astToIRFuncHeader(f, old.Header()); err != nil {
		return nil, errors.WithStack(err)
	}
	return f, nil
}

// astToIRFuncHeader translates the given AST function header into an equivalent
// IR function header.
func (gen *generator) astToIRFuncHeader(f *ir.Function, old ast.FuncHeader) error {
	// (optional) Linkage.
	if n := old.Linkage(); n != nil {
		f.Linkage = asmenum.LinkageFromString(n.Text())
	}
	// (optional) Preemption.
	if n := old.Preemption(); n != nil {
		f.Preemption = asmenum.PreemptionFromString(n.Text())
	}
	// (optional) Visibility.
	if n := old.Visibility(); n != nil {
		f.Visibility = asmenum.VisibilityFromString(n.Text())
	}
	// (optional) DLL storage class.
	if n := old.DLLStorageClass(); n != nil {
		f.DLLStorageClass = asmenum.DLLStorageClassFromString(n.Text())
	}
	// (optional) Calling convention.
	if n := old.CallingConv(); n != nil {
		f.CallingConv = irCallingConv(n)
	}
	// (optional) Return attributes.
	// TODO: handle RetAttrs.
	// Return type: already stored during index.
	// Function parameters.
	ps := old.Params()
	for _, p := range ps.Params() {
		// Type.
		typ, err := gen.irType(p.Typ())
		if err != nil {
			return errors.WithStack(err)
		}
		// Name.
		name := optLocal(p.Name())
		param := ir.NewParam(typ, name)
		// (optional) Parameter attributes.
		for _, oldAttr := range p.Attrs() {
			attr := irParamAttribute(oldAttr)
			param.Attrs = append(param.Attrs, attr)
		}
		f.Params = append(f.Params, param)
	}
	// (optional) Unnamed address.
	if n := old.UnnamedAddr(); n != nil {
		f.UnnamedAddr = asmenum.UnnamedAddrFromString(n.Text())
	}
	// (optional) Address space; stored in f.Typ.
	if n := old.AddrSpace(); n != nil {
		f.Typ.AddrSpace = irAddrSpace(*n)
	}
	// (optional) Function attributes.
	for _, oldFuncAttr := range old.FuncAttrs() {
		funcAttr := irFuncAttribute(oldFuncAttr)
		f.FuncAttrs = append(f.FuncAttrs, funcAttr)
	}
	// (optional) Section.
	// TODO: handle section.
	// (optional) Comdat.
	if n := old.Comdat(); n != nil {
		comdatName := f.GlobalName
		if n := n.Name(); n != nil {
			comdatName = comdat(*n)
		}
		comdatDef, ok := gen.comdats[comdatName]
		if !ok {
			return errors.Errorf("unable to locate Comdat identifier %q used in function header of %q", comdatName, f.GlobalName)
		}
		f.Comdat = comdatDef
	}
	// (optional) Garbage collection.
	// TODO: handle gc.
	// (optional) Prefix.
	if n := old.Prefix(); n != nil {
		prefix, err := gen.irTypeConst(n.TypeConst())
		if err != nil {
			return errors.WithStack(err)
		}
		f.Prefix = prefix
	}
	// (optional) Prologue.
	if n := old.Prologue(); n != nil {
		prologue, err := gen.irTypeConst(n.TypeConst())
		if err != nil {
			return errors.WithStack(err)
		}
		f.Prologue = prologue
	}
	// (optional) Personality.
	if n := old.Personality(); n != nil {
		personality, err := gen.irTypeConst(n.TypeConst())
		if err != nil {
			return errors.WithStack(err)
		}
		f.Personality = personality
	}
	return nil
}

// ~~~ [ Function Definition ] ~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~

// astToIRFuncDef translates the given AST function definition into an
// equivalent IR function definition.
func (gen *generator) astToIRFuncDef(fn constant.Constant, old *ast.FuncDef) (*ir.Function, error) {
	f, ok := fn.(*ir.Function)
	if !ok {
		panic(fmt.Errorf("invalid IR type for AST function definition; expected *ir.Function, got %T", fn))
	}
	// Function header.
	if err := gen.astToIRFuncHeader(f, old.Header()); err != nil {
		return nil, errors.WithStack(err)
	}
	// (optional) Metadata.
	f.Metadata = irMetadataAttachments(old.Metadata())
	// Basic blocks.
	fgen := newFuncGen(gen, f)
	_, err := fgen.resolveLocals(old.Body())
	if err != nil {
		return nil, errors.WithStack(err)
	}
	// (optional) Use list orders.
	// TODO: translate use list orders.
	return f, nil
}

// ### [ Helper functions ] ####################################################

// text returns the text of the given node.
func text(n ast.LlvmNode) string {
	if n := n.LlvmNode(); n != nil {
		return n.Text()
	}
	return ""
}
