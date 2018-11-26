package asm

import (
	"fmt"

	"github.com/llir/ll/ast"
	asmenum "github.com/llir/llvm/asm/enum"
	"github.com/llir/llvm/internal/enc"
	"github.com/llir/llvm/ir"
	"github.com/llir/llvm/ir/constant"
	"github.com/llir/llvm/ir/types"
	"github.com/llir/llvm/ir/value"
	"github.com/pkg/errors"
)

// === [ Create and index IR ] =================================================

// createGlobals indexes global identifiers and creates scaffolding IR global
// declarations and definitions, alias and IFunc definitions, and function
// declarations and definitions (without bodies but with types) of the given
// module.
func (gen *generator) createGlobals() error {
	for ident, old := range gen.old.globals {
		new, err := gen.newGlobal(ident, old)
		if err != nil {
			return errors.WithStack(err)
		}
		gen.new.globals[ident] = new
	}
	return nil
}

// newGlobal returns a new scaffolding IR value (without body but with type)
// based on the given AST global declaration or definition, alias or IFunc
// definition, or function declaration or definition.
func (gen *generator) newGlobal(ident ir.GlobalIdent, old ast.LlvmNode) (constant.Constant, error) {
	switch old := old.(type) {
	case *ast.GlobalDecl:
		new := &ir.Global{}
		setGlobalIdent(new, ident)
		// Content type.
		contentType, err := gen.irType(old.ContentType())
		if err != nil {
			return nil, errors.WithStack(err)
		}
		new.ContentType = contentType
		new.Typ = types.NewPointer(new.ContentType)
		// (optional) Address space.
		if n := old.AddrSpace(); n.IsValid() {
			new.Typ.AddrSpace = irAddrSpace(n)
		}
		return new, nil
	case *ast.GlobalDef:
		new := &ir.Global{}
		setGlobalIdent(new, ident)
		// Content type.
		contentType, err := gen.irType(old.ContentType())
		if err != nil {
			return nil, errors.WithStack(err)
		}
		new.ContentType = contentType
		new.Typ = types.NewPointer(new.ContentType)
		// (optional) Address space.
		if n := old.AddrSpace(); n.IsValid() {
			new.Typ.AddrSpace = irAddrSpace(n)
		}
		return new, nil
	case *ast.IndirectSymbolDef:
		// Content type.
		contentType, err := gen.irType(old.ContentType())
		if err != nil {
			return nil, errors.WithStack(err)
		}
		// Indirect symbol kind.
		kind := old.IndirectSymbolKind().Text()
		switch kind {
		case "alias":
			new := &ir.Alias{Typ: types.NewPointer(contentType)}
			setGlobalIdent(new, ident)
			return new, nil
		case "ifunc":
			new := &ir.IFunc{Typ: types.NewPointer(contentType)}
			setGlobalIdent(new, ident)
			return new, nil
		default:
			panic(fmt.Errorf("support for indirect symbol kind %q not yet implemented", kind))
		}
	case *ast.FuncDecl:
		new := &ir.Function{}
		setGlobalIdent(new, ident)
		// Function signature.
		sig, err := gen.sigFromHeader(old.Header())
		if err != nil {
			return nil, errors.WithStack(err)
		}
		new.Sig = sig
		new.Typ = types.NewPointer(new.Sig)
		// (optional) Address space.
		if n := old.Header().AddrSpace(); n.IsValid() {
			new.Typ.AddrSpace = irAddrSpace(n)
		}
		return new, nil
	case *ast.FuncDef:
		new := &ir.Function{}
		setGlobalIdent(new, ident)
		// Function signature.
		sig, err := gen.sigFromHeader(old.Header())
		if err != nil {
			return nil, errors.WithStack(err)
		}
		new.Sig = sig
		new.Typ = types.NewPointer(new.Sig)
		// (optional) Address space.
		if n := old.Header().AddrSpace(); n.IsValid() {
			new.Typ.AddrSpace = irAddrSpace(n)
		}
		return new, nil
	default:
		panic(fmt.Errorf("support for global variable, indirect symbol or function %T not yet implemented", old))
	}
}

// === [ Translate AST to IR ] =================================================

// translateGlobals translates the AST global declarations and definitions of
// the given module to IR.
func (gen *generator) translateGlobals() error {
	// TODO: make concurrent and benchmark. Each gen.translateGlobalDecl,
	// gen.translateGlobalDef, etc can be run in a Go-routine.
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
			if err := gen.translateGlobalDecl(new, old); err != nil {
				return errors.WithStack(err)
			}
		case *ast.GlobalDef:
			new, ok := v.(*ir.Global)
			if !ok {
				panic(fmt.Errorf("invalid global definition type; expected *ir.Global, got %T", v))
			}
			if err := gen.translateGlobalDef(new, old); err != nil {
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
				if err := gen.translateAliasDef(new, old); err != nil {
					return errors.WithStack(err)
				}
			case "ifunc":
				new, ok := v.(*ir.IFunc)
				if !ok {
					panic(fmt.Errorf("invalid IFunc definition type; expected *ir.IFunc, got %T", v))
				}
				if err := gen.translateIFuncDef(new, old); err != nil {
					return errors.WithStack(err)
				}
			default:
				panic(fmt.Errorf("support for indirect symbol kind %q not yet implemented", kind))
			}
		case *ast.FuncDecl:
			new, ok := v.(*ir.Function)
			if !ok {
				panic(fmt.Errorf("invalid function declaration type; expected *ir.Function, got %T", v))
			}
			if err := gen.translateFuncDecl(new, old); err != nil {
				return errors.WithStack(err)
			}
		case *ast.FuncDef:
			new, ok := v.(*ir.Function)
			if !ok {
				panic(fmt.Errorf("invalid function definition type; expected *ir.Function, got %T", v))
			}
			if err := gen.translateFuncDef(new, old); err != nil {
				return errors.WithStack(err)
			}
		default:
			panic(fmt.Errorf("support for global variable, indirect symbol or function %T not yet implemented", old))
		}
	}
	return nil
}

// --- [ Global declarations ] -------------------------------------------------

// translateGlobalDecl translates the given AST global declarations to IR.
func (gen *generator) translateGlobalDecl(new *ir.Global, old *ast.GlobalDecl) error {
	// (optional) Linkage.
	new.Linkage = asmenum.LinkageFromString(old.ExternLinkage().Text())
	// (optional) Preemption.
	if n := old.Preemption(); n.IsValid() {
		new.Preemption = asmenum.PreemptionFromString(n.Text())
	}
	// (optional) Visibility.
	if n := old.Visibility(); n.IsValid() {
		new.Visibility = asmenum.VisibilityFromString(n.Text())
	}
	// (optional) DLL storage class.
	if n := old.DLLStorageClass(); n.IsValid() {
		new.DLLStorageClass = asmenum.DLLStorageClassFromString(n.Text())
	}
	// (optional) Thread local storage model.
	if n := old.ThreadLocal(); n.IsValid() {
		new.TLSModel = irTLSModelFromThreadLocal(n)
	}
	// (optional) Unnamed address.
	if n := old.UnnamedAddr(); n.IsValid() {
		new.UnnamedAddr = asmenum.UnnamedAddrFromString(n.Text())
	}
	// (optional) Address space: handled in newGlobal.
	// (optional) Externally initialized.
	new.ExternallyInitialized = old.ExternallyInitialized().IsValid()
	// Immutability of global variable (constant or global).
	new.Immutable = irImmutable(old.Immutable())
	// Content type: handled in newGlobal.
	// (optional) Section name.
	if n := old.Section(); n.IsValid() {
		new.Section = stringLit(n.Name())
	}
	// (optional) Comdat.
	if n := old.Comdat(); n.IsValid() {
		// When comdat name is omitted, the global name is used as an implicit
		// comdat name.
		name := new.Name()
		if n := n.Name(); n.IsValid() {
			name = comdatName(n)
		}
		def, ok := gen.new.comdatDefs[name]
		if !ok {
			return errors.Errorf("unable to locate comdat identifier %q used in global declaration of %q", enc.Comdat(name), new.Ident())
		}
		new.Comdat = def
	}
	// (optional) Alignment.
	if n := old.Align(); n.IsValid() {
		new.Align = irAlign(n)
	}
	// (optional) Metadata.
	md, err := gen.irMetadataAttachments(old.Metadata())
	if err != nil {
		return errors.WithStack(err)
	}
	new.Metadata = md
	// (optional) Function attributes.
	for _, oldFuncAttr := range old.FuncAttrs() {
		funcAttr := gen.irFuncAttribute(oldFuncAttr)
		new.FuncAttrs = append(new.FuncAttrs, funcAttr)
	}
	return nil
}

// --- [ Global definitions ] --------------------------------------------------

// translateGlobalDef translates the given AST global definition to IR.
func (gen *generator) translateGlobalDef(new *ir.Global, old *ast.GlobalDef) error {
	// (optional) Linkage.
	if n := old.Linkage(); n.IsValid() {
		new.Linkage = asmenum.LinkageFromString(n.Text())
	}
	// (optional) Preemption.
	if n := old.Preemption(); n.IsValid() {
		new.Preemption = asmenum.PreemptionFromString(n.Text())
	}
	// (optional) Visibility.
	if n := old.Visibility(); n.IsValid() {
		new.Visibility = asmenum.VisibilityFromString(n.Text())
	}
	// (optional) DLL storage class.
	if n := old.DLLStorageClass(); n.IsValid() {
		new.DLLStorageClass = asmenum.DLLStorageClassFromString(n.Text())
	}
	// (optional) Thread local storage model.
	if n := old.ThreadLocal(); n.IsValid() {
		new.TLSModel = irTLSModelFromThreadLocal(n)
	}
	// (optional) Unnamed address.
	if n := old.UnnamedAddr(); n.IsValid() {
		new.UnnamedAddr = asmenum.UnnamedAddrFromString(n.Text())
	}
	// (optional) Address space: handled in newGlobal.
	// (optional) Externally initialized.
	new.ExternallyInitialized = old.ExternallyInitialized().IsValid()
	// Immutability of global variable (constant or global).
	new.Immutable = irImmutable(old.Immutable())
	// Content type: handled in newGlobal.
	// Initial value.
	init, err := gen.irConstant(new.ContentType, old.Init())
	if err != nil {
		return errors.WithStack(err)
	}
	new.Init = init
	// (optional) Section name.
	if n := old.Section(); n.IsValid() {
		new.Section = stringLit(n.Name())
	}
	// (optional) Comdat.
	if n := old.Comdat(); n.IsValid() {
		// When comdat name is omitted, the global name is used as an implicit
		// comdat name.
		name := new.Name()
		if n := n.Name(); n.IsValid() {
			name = comdatName(n)
		}
		def, ok := gen.new.comdatDefs[name]
		if !ok {
			return errors.Errorf("unable to locate comdat identifier %q used in global declaration of %q", enc.Comdat(name), new.Ident())
		}
		new.Comdat = def
	}
	// (optional) Alignment.
	if n := old.Align(); n.IsValid() {
		new.Align = irAlign(n)
	}
	// (optional) Metadata.
	md, err := gen.irMetadataAttachments(old.Metadata())
	if err != nil {
		return errors.WithStack(err)
	}
	new.Metadata = md
	// (optional) Function attributes.
	for _, oldFuncAttr := range old.FuncAttrs() {
		funcAttr := gen.irFuncAttribute(oldFuncAttr)
		new.FuncAttrs = append(new.FuncAttrs, funcAttr)
	}
	return nil
}

// --- [ Alias definitions ] ---------------------------------------------------

// translateAliasDef translates the given AST alias definition to IR.
func (gen *generator) translateAliasDef(new *ir.Alias, old *ast.IndirectSymbolDef) error {
	// (optional) Linkage.
	// TODO: check that linkage is handled correctly.
	if n := old.Linkage(); n.IsValid() {
		new.Linkage = asmenum.LinkageFromString(n.Text())
	}
	if n := old.ExternLinkage(); n.IsValid() {
		new.Linkage = asmenum.LinkageFromString(n.Text())
	}
	// (optional) Preemption.
	if n := old.Preemption(); n.IsValid() {
		new.Preemption = asmenum.PreemptionFromString(n.Text())
	}
	// (optional) Visibility.
	if n := old.Visibility(); n.IsValid() {
		new.Visibility = asmenum.VisibilityFromString(n.Text())
	}
	// (optional) DLL storage class.
	if n := old.DLLStorageClass(); n.IsValid() {
		new.DLLStorageClass = asmenum.DLLStorageClassFromString(n.Text())
	}
	// (optional) Thread local storage model.
	if n := old.ThreadLocal(); n.IsValid() {
		new.TLSModel = irTLSModelFromThreadLocal(n)
	}
	// (optional) Unnamed address.
	if n := old.UnnamedAddr(); n.IsValid() {
		new.UnnamedAddr = asmenum.UnnamedAddrFromString(n.Text())
	}
	// Content type: handled in newGlobal.
	// Aliasee.
	aliasee, err := gen.irIndirectSymbol(new.Typ, old.IndirectSymbol())
	if err != nil {
		return errors.WithStack(err)
	}
	new.Aliasee = aliasee
	return nil
}

// --- [ IFunc definitions ] ---------------------------------------------------

// translateIFuncDef translates the given AST IFunc definition to IR.
func (gen *generator) translateIFuncDef(new *ir.IFunc, old *ast.IndirectSymbolDef) error {
	// (optional) Linkage.
	// TODO: check that linkage is handled correctly.
	if n := old.Linkage(); n.IsValid() {
		new.Linkage = asmenum.LinkageFromString(n.Text())
	}
	if n := old.ExternLinkage(); n.IsValid() {
		new.Linkage = asmenum.LinkageFromString(n.Text())
	}
	// (optional) Preemption.
	if n := old.Preemption(); n.IsValid() {
		new.Preemption = asmenum.PreemptionFromString(n.Text())
	}
	// (optional) Visibility.
	if n := old.Visibility(); n.IsValid() {
		new.Visibility = asmenum.VisibilityFromString(n.Text())
	}
	// (optional) DLL storage class.
	if n := old.DLLStorageClass(); n.IsValid() {
		new.DLLStorageClass = asmenum.DLLStorageClassFromString(n.Text())
	}
	// (optional) Thread local storage model.
	if n := old.ThreadLocal(); n.IsValid() {
		new.TLSModel = irTLSModelFromThreadLocal(n)
	}
	// (optional) Unnamed address.
	if n := old.UnnamedAddr(); n.IsValid() {
		new.UnnamedAddr = asmenum.UnnamedAddrFromString(n.Text())
	}
	// Content type: handled in newGlobal.
	// Resolver.
	resolver, err := gen.irIndirectSymbol(new.Typ, old.IndirectSymbol())
	if err != nil {
		return errors.WithStack(err)
	}
	new.Resolver = resolver
	return nil
}

// --- [ Function declarations ] -----------------------------------------------

// translateFuncDecl translates the given AST function declaration to IR.
func (gen *generator) translateFuncDecl(new *ir.Function, old *ast.FuncDecl) error {
	// (optional) Metadata.
	md, err := gen.irMetadataAttachments(old.Metadata())
	if err != nil {
		return errors.WithStack(err)
	}
	new.Metadata = md
	// Function header.
	if err := gen.translateFuncHeader(new, old.Header()); err != nil {
		return errors.WithStack(err)
	}
	return nil
}

// --- [ Function definitions ] ------------------------------------------------

// translateFuncDef translates the given AST function definition to IR.
func (gen *generator) translateFuncDef(new *ir.Function, old *ast.FuncDef) error {
	// Function header.
	if err := gen.translateFuncHeader(new, old.Header()); err != nil {
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
	for _, oldUseListOrder := range oldBody.UseListOrders() {
		useListOrder, err := fgen.irUseListOrder(oldUseListOrder)
		if err != nil {
			return errors.WithStack(err)
		}
		new.UseListOrders = append(new.UseListOrders, useListOrder)
	}
	return nil
}

// ~~~ [ Function headers ] ~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~

// translateFuncHeader translates the given AST function header to IR.
func (gen *generator) translateFuncHeader(new *ir.Function, old ast.FuncHeader) error {
	// (optional) Linkage.
	// TODO: check that linkage is handled correctly.
	if n := old.Linkage(); n.IsValid() {
		new.Linkage = asmenum.LinkageFromString(n.Text())
	}
	if n := old.ExternLinkage(); n.IsValid() {
		new.Linkage = asmenum.LinkageFromString(n.Text())
	}
	// (optional) Preemption.
	if n := old.Preemption(); n.IsValid() {
		new.Preemption = asmenum.PreemptionFromString(n.Text())
	}
	// (optional) Visibility.
	if n := old.Visibility(); n.IsValid() {
		new.Visibility = asmenum.VisibilityFromString(n.Text())
	}
	// (optional) DLL storage class.
	if n := old.DLLStorageClass(); n.IsValid() {
		new.DLLStorageClass = asmenum.DLLStorageClassFromString(n.Text())
	}
	// (optional) Calling convention.
	if n := old.CallingConv(); n.LlvmNode().IsValid() {
		new.CallingConv = irCallingConv(n)
	}
	// (optional) Return attributes.
	for _, oldRetAttr := range old.ReturnAttrs() {
		retAttr := irReturnAttribute(oldRetAttr)
		new.ReturnAttrs = append(new.ReturnAttrs, retAttr)
	}
	// Return type: handled in newGlobal.
	// Function parameters.
	ps := old.Params()
	for _, p := range ps.Params() {
		// Type.
		typ, err := gen.irType(p.Typ())
		if err != nil {
			return errors.WithStack(err)
		}
		// Name.
		param := ir.NewParam("", typ)
		if p.Name().IsValid() {
			ident := localIdent(p.Name())
			param.LocalIdent = ident
		}
		// (optional) Parameter attributes.
		for _, oldParamAttr := range p.Attrs() {
			paramAttr := irParamAttribute(oldParamAttr)
			param.Attrs = append(param.Attrs, paramAttr)
		}
		new.Params = append(new.Params, param)
	}
	// (optional) Unnamed address.
	if n := old.UnnamedAddr(); n.IsValid() {
		new.UnnamedAddr = asmenum.UnnamedAddrFromString(n.Text())
	}
	// (optional) Address space: handled in newGlobal.
	// (optional) Function attributes.
	for _, oldFuncAttr := range old.FuncAttrs() {
		funcAttr := gen.irFuncAttribute(oldFuncAttr)
		new.FuncAttrs = append(new.FuncAttrs, funcAttr)
	}

	// (optional) Section name.
	if n := old.Section(); n.IsValid() {
		new.Section = stringLit(n.Name())
	}
	// (optional) Comdat.
	if n := old.Comdat(); n.IsValid() {
		// When comdat name is omitted, the function name is used as an implicit
		// comdat name.
		name := new.Name()
		if n := n.Name(); n.IsValid() {
			name = comdatName(n)
		}
		def, ok := gen.new.comdatDefs[name]
		if !ok {
			return errors.Errorf("unable to locate comdat identifier %q used in function header of %q", enc.Comdat(name), new.Ident())
		}
		new.Comdat = def
	}
	// (optional) Garbage collection.
	if n := old.GCNode(); n.IsValid() {
		new.GC = stringLit(n.Name())
	}
	// (optional) Prefix.
	if n := old.Prefix(); n.IsValid() {
		prefix, err := gen.irTypeConst(n.TypeConst())
		if err != nil {
			return errors.WithStack(err)
		}
		new.Prefix = prefix
	}
	// (optional) Prologue.
	if n := old.Prologue(); n.IsValid() {
		prologue, err := gen.irTypeConst(n.TypeConst())
		if err != nil {
			return errors.WithStack(err)
		}
		new.Prologue = prologue
	}
	// (optional) Personality.
	if n := old.Personality(); n.IsValid() {
		personality, err := gen.irTypeConst(n.TypeConst())
		if err != nil {
			return errors.WithStack(err)
		}
		new.Personality = personality
	}
	return nil
}

// ### [ Helper functions ] ####################################################

// sigFromHeader returns the function signature of the given function header.
func (gen *generator) sigFromHeader(old ast.FuncHeader) (*types.FuncType, error) {
	sig := &types.FuncType{}
	// Return type.
	retType, err := gen.irType(old.RetType())
	if err != nil {
		return nil, errors.WithStack(err)
	}
	sig.RetType = retType
	// Function parameters.
	ps := old.Params()
	for _, p := range ps.Params() {
		param, err := gen.irType(p.Typ())
		if err != nil {
			return nil, errors.WithStack(err)
		}
		sig.Params = append(sig.Params, param)
	}
	// Variadic.
	sig.Variadic = ps.Variadic().IsValid()
	return sig, nil
}

// text returns the text of the given node.
func text(n ast.LlvmNode) string {
	if n := n.LlvmNode(); n != nil {
		return n.Text()
	}
	return ""
}

// global is a global variable.
type global interface {
	value.Named
	// ID returns the ID of the global identifier.
	ID() int64
	// SetID sets the ID of the global identifier.
	SetID(id int64)
}

// setGlobalIdent sets the identifier of the given global variable.
func setGlobalIdent(g global, ident ir.GlobalIdent) {
	if ident.IsUnnamed() {
		g.SetID(ident.GlobalID)
	} else {
		g.SetName(ident.GlobalName)
	}
}
