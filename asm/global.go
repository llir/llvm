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

// === [ Create and index IR ] =================================================

// createGlobalEntities indexes global identifiers and creates scaffolding IR
// global declarations and definitions, indirect symbol definitions (aliases and
// indirect functions), and function declarations and definitions (without
// bodies but with types) of the given module.
//
// post-condition: gen.new.globals maps from global identifier (without '@'
// prefix) to corresponding skeleton IR value.
func (gen *generator) createGlobalEntities() error {
	// 4a1. Index global identifiers and create scaffolding IR global
	//      declarations and definitions, indirect symbol definitions (aliases
	//      and indirect functions), and function declarations and definitions
	//      (without bodies but with types).
	for ident, old := range gen.old.globals {
		new, err := gen.newGlobalEntity(ident, old)
		if err != nil {
			return errors.WithStack(err)
		}
		gen.new.globals[ident] = new
	}
	return nil
}

// newGlobalEntity returns a new scaffolding IR value (without body but with
// type) based on the given AST global declaration or definition, indirect
// symbol definitions (aliases and indirect functions), or function declaration
// or definition.
func (gen *generator) newGlobalEntity(ident ir.GlobalIdent, old ast.LlvmNode) (constant.Constant, error) {
	switch old := old.(type) {
	case *ast.GlobalDecl:
		oldAddrSpace, _ := old.AddrSpace()
		return gen.newGlobal(ident, old.ContentType(), oldAddrSpace)
	case *ast.IndirectSymbolDef:
		return gen.newIndirectSymbol(ident, old)
	case *ast.FuncDecl:
		return gen.newFunc(ident, old.Header())
	case *ast.FuncDef:
		return gen.newFunc(ident, old.Header())
	default:
		panic(fmt.Errorf("support for global variable, indirect symbol or function %T not yet implemented", old))
	}
}

// newGlobal returns a new IR global variable declaration or definition (without
// body but with type) based on the given AST content type and optional address
// space.
func (gen *generator) newGlobal(ident ir.GlobalIdent, oldContentType ast.Type, oldAddrSpace ast.AddrSpace) (*ir.Global, error) {
	// Content type.
	contentType, err := gen.irType(oldContentType)
	if err != nil {
		return nil, errors.WithStack(err)
	}
	typ := types.NewPointer(contentType)
	// (optional) Address space.
	if oldAddrSpace.IsValid() {
		typ.AddrSpace = irAddrSpace(oldAddrSpace)
	}
	return &ir.Global{GlobalIdent: ident, ContentType: contentType, Typ: typ}, nil
}

// newIndirectSymbol returns a new IR indirect symbol definition (without body
// but with type) based on the given AST indirect symbol.
func (gen *generator) newIndirectSymbol(ident ir.GlobalIdent, old *ast.IndirectSymbolDef) (constant.Constant, error) {
	// Content type.
	contentType, err := gen.irType(old.ContentType())
	if err != nil {
		return nil, errors.WithStack(err)
	}
	typ := types.NewPointer(contentType)
	// Indirect symbol kind.
	kind := old.IndirectSymbolKind().Text()
	switch kind {
	case "alias":
		return &ir.Alias{GlobalIdent: ident, Typ: typ}, nil
	case "ifunc":
		return &ir.IFunc{GlobalIdent: ident, Typ: typ}, nil
	default:
		panic(fmt.Errorf("support for indirect symbol kind %q not yet implemented", kind))
	}
}

// newFunc returns a new IR function declaration or definition (without body but
// with type) based on the given AST function header.
func (gen *generator) newFunc(ident ir.GlobalIdent, hdr ast.FuncHeader) (*ir.Func, error) {
	// Function signature.
	sig, err := gen.irSigFromHeader(hdr)
	if err != nil {
		return nil, errors.WithStack(err)
	}
	typ := types.NewPointer(sig)
	// (optional) Address space.
	if n, ok := hdr.AddrSpace(); ok {
		typ.AddrSpace = irAddrSpace(n)
	}
	return &ir.Func{GlobalIdent: ident, Sig: sig, Typ: typ, Parent: gen.m}, nil
}

// ### [ Helper functions ] ####################################################

// irSigFromHeader translates the AST function signature to an equivalent IR
// function type.
func (gen *generator) irSigFromHeader(old ast.FuncHeader) (*types.FuncType, error) {
	// Return type.
	sig := &types.FuncType{}
	retType, err := gen.irType(old.RetType())
	if err != nil {
		return nil, errors.WithStack(err)
	}
	sig.RetType = retType
	// Function parameters.
	ps := old.Params()
	if oldParams := ps.Params(); len(oldParams) > 0 {
		sig.Params = make([]types.Type, len(oldParams))
		for i, oldParam := range oldParams {
			param, err := gen.irType(oldParam.Typ())
			if err != nil {
				return nil, errors.WithStack(err)
			}
			sig.Params[i] = param
		}
	}
	// Variadic.
	_, sig.Variadic = ps.Variadic()
	return sig, nil
}

// === [ Translate AST to IR ] =================================================

// translateGlobalEntities translate AST global declarations and definitions,
// indirect symbol definitions, and function declarations and definitions to IR.
func (gen *generator) translateGlobalEntities() error {
	// TODO: make concurrent and benchmark difference in walltime.

	// 4b1. Translate AST global declarations and definitions, indirect symbol
	//      definitions, and function declarations and definitions to IR.
	for ident, old := range gen.old.globals {
		v, ok := gen.new.globals[ident]
		if !ok {
			panic(fmt.Errorf("unable to locate global identifier %q", ident.Ident()))
		}
		switch old := old.(type) {
		case *ast.GlobalDecl:
			new, ok := v.(*ir.Global)
			if !ok {
				panic(fmt.Errorf("invalid global declaration type; expected *ir.Global, got %T", v))
			}
			if err := gen.irGlobal(new, old); err != nil {
				return errors.WithStack(err)
			}
		case *ast.IndirectSymbolDef:
			kind := old.IndirectSymbolKind().Text()
			switch kind {
			case "alias":
				new, ok := v.(*ir.Alias)
				if !ok {
					panic(fmt.Errorf("invalid alias definition type; expected *ir.Alias, got %T", v))
				}
				if err := gen.irAlias(new, old); err != nil {
					return errors.WithStack(err)
				}
			case "ifunc":
				new, ok := v.(*ir.IFunc)
				if !ok {
					panic(fmt.Errorf("invalid IFunc definition type; expected *ir.IFunc, got %T", v))
				}
				if err := gen.irIFunc(new, old); err != nil {
					return errors.WithStack(err)
				}
			default:
				panic(fmt.Errorf("support for indirect symbol kind %q not yet implemented", kind))
			}
		case *ast.FuncDecl:
			new, ok := v.(*ir.Func)
			if !ok {
				panic(fmt.Errorf("invalid function declaration type; expected *ir.Func, got %T", v))
			}
			if err := gen.irFuncDecl(new, old); err != nil {
				return errors.WithStack(err)
			}
		case *ast.FuncDef:
			new, ok := v.(*ir.Func)
			if !ok {
				panic(fmt.Errorf("invalid function definition type; expected *ir.Func, got %T", v))
			}
			if err := gen.irFuncDef(new, old); err != nil {
				return errors.WithStack(err)
			}
		default:
			panic(fmt.Errorf("support for global variable, indirect symbol or function %T not yet implemented", old))
		}
	}
	return nil
}

// --- [ Global declarations ] -------------------------------------------------

// irGlobal translates the AST global declaration into an equivalent IR
// global declaration.
func (gen *generator) irGlobal(new *ir.Global, old *ast.GlobalDecl) error {
	// (optional) Linkage.
	if n, ok := old.Linkage(); ok {
		new.Linkage = asmenum.LinkageFromString(n.LlvmNode().Text())
	}
	// (optional) Preemption.
	if n, ok := old.Preemption(); ok {
		new.Preemption = asmenum.PreemptionFromString(n.Text())
	}
	// (optional) Visibility.
	if n, ok := old.Visibility(); ok {
		new.Visibility = asmenum.VisibilityFromString(n.Text())
	}
	// (optional) DLL storage class.
	if n, ok := old.DLLStorageClass(); ok {
		new.DLLStorageClass = asmenum.DLLStorageClassFromString(n.Text())
	}
	// (optional) Thread local storage model.
	if n, ok := old.ThreadLocal(); ok {
		new.TLSModel = irTLSModelFromThreadLocal(n)
	}
	// (optional) Unnamed address.
	if n, ok := old.UnnamedAddr(); ok {
		new.UnnamedAddr = asmenum.UnnamedAddrFromString(n.Text())
	}
	// (optional) Address space: handled in newGlobalEntity.
	// (optional) Externally initialized.
	_, new.ExternallyInitialized = old.ExternallyInitialized()
	// Immutability of global variable (constant or global).
	new.Immutable = irImmutable(old.Immutable())
	// Content type: handled in newGlobalEntity.
	// Initial value (only used in global variable definitions).
	if n, ok := old.Init(); ok {
		init, err := gen.irConstant(new.ContentType, n)
		if err != nil {
			return errors.WithStack(err)
		}
		new.Init = init
	}
	// (optional) Section name.
	if n, ok := old.Section(); ok {
		new.Section = stringLit(n.Name())
	}
	// (optional) Comdat.
	if n, ok := old.Comdat(); ok {
		// When comdat name is omitted, the global name is used as an implicit
		// comdat name.
		name := new.Name()
		if n, ok := n.Name(); ok {
			name = comdatName(n)
		}
		def, ok := gen.new.comdatDefs[name]
		if !ok {
			return errors.Errorf("unable to locate comdat identifier %q used in global declaration of %q", enc.Comdat(name), new.Ident())
		}
		new.Comdat = def
	}
	// (optional) Alignment.
	if n, ok := old.Align(); ok {
		new.Align = irAlign(n)
	}
	// (optional) Metadata.
	md, err := gen.irMetadataAttachments(old.Metadata())
	if err != nil {
		return errors.WithStack(err)
	}
	new.Metadata = md
	// (optional) Function attributes.
	if oldFuncAttrs := old.FuncAttrs(); len(oldFuncAttrs) > 0 {
		new.FuncAttrs = make([]ir.FuncAttribute, len(oldFuncAttrs))
		for i, oldFuncAttr := range oldFuncAttrs {
			funcAttr := gen.irFuncAttribute(oldFuncAttr)
			new.FuncAttrs[i] = funcAttr
		}
	}
	return nil
}

// --- [ Alias definitions ] ---------------------------------------------------

// irAlias translates the AST indirect symbol definition (alias) into an
// equivalent IR alias definition.
func (gen *generator) irAlias(new *ir.Alias, old *ast.IndirectSymbolDef) error {
	// (optional) Linkage.
	if n, ok := old.Linkage(); ok {
		new.Linkage = asmenum.LinkageFromString(n.Text())
	}
	if n, ok := old.ExternLinkage(); ok {
		new.Linkage = asmenum.LinkageFromString(n.Text())
	}
	// (optional) Preemption.
	if n, ok := old.Preemption(); ok {
		new.Preemption = asmenum.PreemptionFromString(n.Text())
	}
	// (optional) Visibility.
	if n, ok := old.Visibility(); ok {
		new.Visibility = asmenum.VisibilityFromString(n.Text())
	}
	// (optional) DLL storage class.
	if n, ok := old.DLLStorageClass(); ok {
		new.DLLStorageClass = asmenum.DLLStorageClassFromString(n.Text())
	}
	// (optional) Thread local storage model.
	if n, ok := old.ThreadLocal(); ok {
		new.TLSModel = irTLSModelFromThreadLocal(n)
	}
	// (optional) Unnamed address.
	if n, ok := old.UnnamedAddr(); ok {
		new.UnnamedAddr = asmenum.UnnamedAddrFromString(n.Text())
	}
	// Content type: handled in newGlobalEntity.
	// Aliasee.
	aliasee, err := gen.irIndirectSymbol(new.Typ, old.IndirectSymbol())
	if err != nil {
		return errors.WithStack(err)
	}
	new.Aliasee = aliasee
	return nil
}

// --- [ IFunc definitions ] ---------------------------------------------------

// irIFunc translates the AST indirect symbol definition (IFunc) into an
// equivalent IR indirect function definition.
func (gen *generator) irIFunc(new *ir.IFunc, old *ast.IndirectSymbolDef) error {
	// (optional) Linkage.
	if n, ok := old.Linkage(); ok {
		new.Linkage = asmenum.LinkageFromString(n.Text())
	}
	if n, ok := old.ExternLinkage(); ok {
		new.Linkage = asmenum.LinkageFromString(n.Text())
	}
	// (optional) Preemption.
	if n, ok := old.Preemption(); ok {
		new.Preemption = asmenum.PreemptionFromString(n.Text())
	}
	// (optional) Visibility.
	if n, ok := old.Visibility(); ok {
		new.Visibility = asmenum.VisibilityFromString(n.Text())
	}
	// (optional) DLL storage class.
	if n, ok := old.DLLStorageClass(); ok {
		new.DLLStorageClass = asmenum.DLLStorageClassFromString(n.Text())
	}
	// (optional) Thread local storage model.
	if n, ok := old.ThreadLocal(); ok {
		new.TLSModel = irTLSModelFromThreadLocal(n)
	}
	// (optional) Unnamed address.
	if n, ok := old.UnnamedAddr(); ok {
		new.UnnamedAddr = asmenum.UnnamedAddrFromString(n.Text())
	}
	// Content type: handled in newGlobalEntity.
	// Resolver.
	resolver, err := gen.irIndirectSymbol(new.Typ, old.IndirectSymbol())
	if err != nil {
		return errors.WithStack(err)
	}
	new.Resolver = resolver
	return nil
}

// --- [ Function declarations ] -----------------------------------------------

// irFuncDecl translates the AST function declaration into an equivalent IR
// function declaration.
func (gen *generator) irFuncDecl(new *ir.Func, old *ast.FuncDecl) error {
	// (optional) Metadata.
	md, err := gen.irMetadataAttachments(old.Metadata())
	if err != nil {
		return errors.WithStack(err)
	}
	new.Metadata = md
	// Function header.
	return gen.irFuncHeader(new, old.Header())
}

// --- [ Function definitions ] ------------------------------------------------

// irFuncDef translates the AST function definition into an equivalent IR
// function definition.
func (gen *generator) irFuncDef(new *ir.Func, old *ast.FuncDef) error {
	// Function header.
	if err := gen.irFuncHeader(new, old.Header()); err != nil {
		return errors.WithStack(err)
	}
	// (optional) Metadata.
	md, err := gen.irMetadataAttachments(old.Metadata())
	if err != nil {
		return errors.WithStack(err)
	}
	new.Metadata = md
	// Basic blocks.
	fgen := newFuncGen(gen, new)
	oldBody := old.Body()
	if err := fgen.resolveLocals(oldBody); err != nil {
		return errors.WithStack(err)
	}
	// (optional) Use list orders.
	if oldUseListOrders := oldBody.UseListOrders(); len(oldUseListOrders) > 0 {
		new.UseListOrders = make([]*ir.UseListOrder, len(oldUseListOrders))
		for i, oldUseListOrder := range oldUseListOrders {
			useListOrder, err := fgen.irUseListOrder(oldUseListOrder)
			if err != nil {
				return errors.WithStack(err)
			}
			new.UseListOrders[i] = useListOrder
		}
	}
	return nil
}

// ~~~ [ Function headers ] ~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~

// irFuncHeader translates the AST function header into an equivalent IR
// function.
func (gen *generator) irFuncHeader(new *ir.Func, old ast.FuncHeader) error {
	// (optional) Linkage.
	if n, ok := old.Linkage(); ok {
		new.Linkage = asmenum.LinkageFromString(n.Text())
	}
	if n, ok := old.ExternLinkage(); ok {
		new.Linkage = asmenum.LinkageFromString(n.Text())
	}
	// (optional) Preemption.
	if n, ok := old.Preemption(); ok {
		new.Preemption = asmenum.PreemptionFromString(n.Text())
	}
	// (optional) Visibility.
	if n, ok := old.Visibility(); ok {
		new.Visibility = asmenum.VisibilityFromString(n.Text())
	}
	// (optional) DLL storage class.
	if n, ok := old.DLLStorageClass(); ok {
		new.DLLStorageClass = asmenum.DLLStorageClassFromString(n.Text())
	}
	// (optional) Calling convention.
	if n, ok := old.CallingConv(); ok {
		new.CallingConv = irCallingConv(n)
	}
	// (optional) Return attributes.
	if oldReturnAttrs := old.ReturnAttrs(); len(oldReturnAttrs) > 0 {
		new.ReturnAttrs = make([]ir.ReturnAttribute, len(oldReturnAttrs))
		for i, oldRetAttr := range oldReturnAttrs {
			retAttr := irReturnAttribute(oldRetAttr)
			new.ReturnAttrs[i] = retAttr
		}
	}
	// Return type: handled in newGlobalEntity.
	// Function parameters.
	ps := old.Params()
	if oldParams := ps.Params(); len(oldParams) > 0 {
		new.Params = make([]*ir.Param, len(oldParams))
		for i, oldParam := range oldParams {
			// Type.
			typ, err := gen.irType(oldParam.Typ())
			if err != nil {
				return errors.WithStack(err)
			}
			// Name.
			param := ir.NewParam("", typ)
			if n, ok := oldParam.Name(); ok {
				ident := localIdent(n)
				param.LocalIdent = ident
			}
			// (optional) Parameter attributes.
			if oldParamAttrs := oldParam.Attrs(); len(oldParamAttrs) > 0 {
				param.Attrs = make([]ir.ParamAttribute, len(oldParamAttrs))
				for j, oldParamAttr := range oldParamAttrs {
					paramAttr, err := gen.irParamAttribute(oldParamAttr)
					if err != nil {
						return errors.WithStack(err)
					}
					param.Attrs[j] = paramAttr
				}
			}
			new.Params[i] = param
		}
	}
	// (optional) Unnamed address.
	if n, ok := old.UnnamedAddr(); ok {
		new.UnnamedAddr = asmenum.UnnamedAddrFromString(n.Text())
	}
	// (optional) Address space: handled in newGlobalEntity.
	// (optional) Function attributes.
	if oldFuncAttrs := old.FuncAttrs(); len(oldFuncAttrs) > 0 {
		new.FuncAttrs = make([]ir.FuncAttribute, len(oldFuncAttrs))
		for i, oldFuncAttr := range oldFuncAttrs {
			funcAttr := gen.irFuncAttribute(oldFuncAttr)
			new.FuncAttrs[i] = funcAttr
		}
	}

	// (optional) Section name.
	if n, ok := old.Section(); ok {
		new.Section = stringLit(n.Name())
	}
	// (optional) Comdat.
	if n, ok := old.Comdat(); ok {
		// When comdat name is omitted, the function name is used as an implicit
		// comdat name.
		name := new.Name()
		if n, ok := n.Name(); ok {
			name = comdatName(n)
		}
		def, ok := gen.new.comdatDefs[name]
		if !ok {
			return errors.Errorf("unable to locate comdat identifier %q used in function header of %q", enc.Comdat(name), new.Ident())
		}
		new.Comdat = def
	}
	// (optional) Garbage collection.
	if n, ok := old.GCNode(); ok {
		new.GC = stringLit(n.Name())
	}
	// (optional) Prefix.
	if n, ok := old.Prefix(); ok {
		prefix, err := gen.irTypeConst(n.TypeConst())
		if err != nil {
			return errors.WithStack(err)
		}
		new.Prefix = prefix
	}
	// (optional) Prologue.
	if n, ok := old.Prologue(); ok {
		prologue, err := gen.irTypeConst(n.TypeConst())
		if err != nil {
			return errors.WithStack(err)
		}
		new.Prologue = prologue
	}
	// (optional) Personality.
	if n, ok := old.Personality(); ok {
		personality, err := gen.irTypeConst(n.TypeConst())
		if err != nil {
			return errors.WithStack(err)
		}
		new.Personality = personality
	}
	return nil
}
