package parser

import (
	"github.com/llir/llvm/asm"
	"github.com/llir/llvm/asm/token"
	"github.com/llir/llvm/ir"
	"github.com/mewkiz/pkg/errutil"
)

// parseGlobalDecl parses a global variable definition or an external global
// variable declaration. The next token is either a GlobalID or a GlobalVar.
//
// Syntax:
//    GlobalDecl = GlobalName "=" DeclLinkage GlobalProperties Type GlobalAttrList .
//    GlobalDef  = GlobalName "=" [ DefLinkage ] GlobalProperties InitConst GlobalAttrList .
//    GlobalName = Global .
//    InitConst  = Const .
//
//    GlobalProperties = [ Visibility ] [ DLLStorage ] [ ThreadLocal ]
//                       [ "unnamed_addr" ] [ AddrSpace ]
//                       [ "externally_initialized" ]
//                       ( "constant" | "global" ) .
//
//    GlobalAttrList = { "," GlobalAttr } .
//    GlobalAttr     = AlignAttr | ComdatAttr | SectionAttr .
//
// Examples:
//    @x = global i32 7
//    @y = external global i32
//    @hello = constant [13 x i8] c"hello world\0A\00"
//
// References:
//    http://llvm.org/docs/LangRef.html#global-variables
func (p *parser) parseGlobalDecl() error {
	// Consume global name.
	name := p.next()

	// Consume "=" token.
	if !p.accept(token.Equal) {
		return errutil.Newf(`expected "=" after global variable name %q, got %q token`, asm.EncGlobal(name.Val), p.next())
	}

	// Consume optional linkage type.
	extern := p.tryLinkage()

	// Consume optional visibility style.
	p.tryVisibility()

	// Consume optional DLL storage class.
	p.tryDLLStorage()

	// Consume optional TLS model.
	if err := p.tryThreadLocal(); err != nil {
		return errutil.Err(err)
	}

	// Consume optional "unnamed_addr" token.
	p.accept(token.KwUnnamedAddr)

	// Consume optional address space.
	if _, err := p.tryAddrSpace(); err != nil {
		return errutil.Err(err)
	}

	// Consume optional "externally_initialized" token.
	p.accept(token.KwExternallyInitialized)

	// Consume "constant" or "global" token.
	immutable := false
	switch tok := p.next(); tok.Kind {
	case token.KwConstant:
		immutable = true
	case token.KwGlobal:
	default:
		decl := "definition"
		if extern {
			decl = "declaration"
		}
		return errutil.Newf(`invalid global variable %s for %q; expected "global" or "constant", got %q`, decl, asm.EncGlobal(name.Val), tok)
	}

	var global *ir.GlobalDecl
	if extern {
		// Consume type.
		typ, err := p.parseType()
		if err != nil {
			return errutil.Err(err)
		}
		global = ir.NewGlobalDecl(name.Val, typ, immutable)
	} else {
		// Consume type and initial value.
		val, err := p.parseConst()
		if err != nil {
			return errutil.Err(err)
		}
		global = ir.NewGlobalDef(name.Val, val, immutable)
	}
	p.m.Globals = append(p.m.Globals, global)

	// Consume optional global attributes; e.g.
	//    align 8
	//    comdat($foo)
	//    section "foo"
	for p.accept(token.Comma) {
		switch tok := p.next(); tok.Kind {
		case token.KwAlign:
			if _, err := p.parseAlignAttr(); err != nil {
				return errutil.Err(err)
			}
		case token.KwComdat:
			if _, err := p.parseComdatAttr(); err != nil {
				return errutil.Err(err)
			}
		case token.KwSection:
			if _, err := p.parseSectionAttr(); err != nil {
				return errutil.Err(err)
			}
		default:
			return errutil.Newf(`expected "align", "comdat" or "section", got %q token`, p.next())
		}
	}

	return nil
}

// tryLinkage tries to consume the optional linkage type of a global variable or
// function, and reports whether it is externally visible.
//
// Syntax:
//    Linkage = DeclLinkage | DefLinkage .
//
//    DeclLinkage = "extern_weak" | "external" .
//    DefLinkage  = "appending" | "available_externally" | "common" |
//                  "internal" | "linkonce" | "linkonce_odr" | "private" |
//                  "weak" | "weak_odr" .
//
// References:
//    http://llvm.org/docs/LangRef.html#linkage
func (p *parser) tryLinkage() (extern bool) {
	// NOTE: Currently, a subset of the information is retained about the linkage
	// type, namely the external visibility.

	switch tok := p.next(); tok.Kind {
	case token.KwAppending:
	case token.KwAvailableExternally:
	case token.KwCommon:
	case token.KwExternWeak:
		return true
	case token.KwExternal:
		return true
	case token.KwInternal:
	case token.KwLinkonce:
	case token.KwLinkonceOdr:
	case token.KwPrivate:
	case token.KwWeak:
	case token.KwWeakOdr:
	default:
		// The consumed token is not a linkage type, backup.
		p.backup()
	}
	return false
}

// tryVisibility tries to consume the optional visibility style of a global
// variable or function.
//
// Syntax:
//    Visibility = "default" | "hidden" | "protected" .
//
// References:
//    http://llvm.org/docs/LangRef.html#visibility-styles
func (p *parser) tryVisibility() {
	// NOTE: Currently, no information is retained about the visibility style.

	switch tok := p.next(); tok.Kind {
	case token.KwDefault:
	case token.KwHidden:
	case token.KwProtected:
	default:
		// The consumed token is not a visibility style, backup.
		p.backup()
	}
}

// tryDLLStorage tries to consume the optional DLL storage class of a global
// variable or function.
//
// Syntax:
//    DLLStorage = "dllexport" | "dllimport" .
//
// References:
//    http://llvm.org/docs/LangRef.html#dllstorageclass
func (p *parser) tryDLLStorage() {
	// NOTE: Currently, no information is retained about the DLL storage class.
	switch tok := p.next(); tok.Kind {
	case token.KwDllexport:
	case token.KwDllimport:
	default:
		// The consumed token is not a DLL storage class, backup.
		p.backup()
	}
}

// tryThreadLocal tries to consume the optional thread local storage model of a
// global variable.
//
// Syntax:
//    ThreadLocal = "thread_local" [ "(" TLSModel ")" ] .
//    TLSModel    = "initialexec" | "localdynamic" | "localexec" .
//
// References:
//    http://llvm.org/docs/LangRef.html#thread-local-storage-models
func (p *parser) tryThreadLocal() error {
	// NOTE: Currently, no information is retained about the thread local storage
	// model.

	// Consume "thread_local" token.
	if !p.accept(token.KwThreadLocal) {
		return nil
	}

	// Consume TLS model.
	if p.accept(token.Lparen) {
		switch tok := p.next(); tok.Kind {
		case token.KwInitialexec:
		case token.KwLocaldynamic:
		case token.KwLocalexec:
		default:
			return errutil.Newf(`expected "initialexec", "localdynamic" or "localexec", got %q token`, tok)
		}
		if !p.accept(token.Rparen) {
			return errutil.Newf(`expected ")" after TLS model, got %q token`, p.next())
		}
	}

	return nil
}

// tryAddrSpace tries to consume the optional address space of a global
// variable.
//
// Syntax:
//    AddrSpace = "addrspace" "(" int_lit ")" .
//
// References:
//    http://llvm.org/docs/LangRef.html#global-variables
func (p *parser) tryAddrSpace() (addrspace int, err error) {
	// NOTE: Currently, no information is retained about the address space.

	// Consume "addrspace" token.
	if !p.accept(token.KwAddrspace) {
		return 0, nil
	}

	// Consume "(" token.
	if !p.accept(token.Lparen) {
		return 0, errutil.Newf(`expected "(" in address space, got %q token`, p.next())
	}

	// Consume integer literal.
	addrspace, err = p.parseInt()
	if err != nil {
		return 0, errutil.Err(err)
	}

	// Consume ")" token.
	if !p.accept(token.Rparen) {
		return 0, errutil.Newf(`expected ")" in address space, got %q token`, p.next())
	}

	return addrspace, nil
}

// parseAlignAttr parses an alignment attribute of a global variable or
// function. An "align" token has already been consumed.
//
// Syntax:
//    AlignAttr = "align" int_lit .
//
// References:
//    http://llvm.org/docs/LangRef.html#global-variables
func (p *parser) parseAlignAttr() (n int, err error) {
	n, err = p.parseInt()
	if err != nil {
		return 0, errutil.Err(err)
	}

	// Verify that n is a power of 2.
	if !isPow2(n) {
		return 0, errutil.Newf("invalid alignment; expected power of 2, got %d", n)
	}
	return n, nil
}

// isPow2 returns true if x is a power of 2.
func isPow2(x int) bool {
	return x > 0 && (x&(x-1)) == 0
}

// parseComdatAttr parses a COMDAT attribute of a global variable or function. A
// "comdat" token has already been consumed.
//
// Syntax:
//    ComdatAttr = "comdat" [ "(" ComdatName ")" ] .
//    ComdatName = ComdatVar .
//
// References:
//    http://llvm.org/docs/LangRef.html#comdats
func (p *parser) parseComdatAttr() (name string, err error) {
	if !p.accept(token.Lparen) {
		return "", nil
	}
	name, err = p.expect(token.ComdatVar)
	if err != nil {
		return "", errutil.Err(err)
	}
	if len(name) < 1 {
		return "", errutil.New("empty COMDAT variable name")
	}
	if !p.accept(token.Rparen) {
		return "", errutil.Newf(`expected ")" after COMDAT variable, got %q token`, p.next())
	}
	return name, nil
}

// parseSectionAttr parses a section attribute of a global variable or function.
// A "section" token has already been consumed.
//
// Syntax:
//    SectionAttr = "section" SectionName .
//    SectionName = string_lit .
//
// References:
//    http://llvm.org/docs/LangRef.html#global-variables
func (p *parser) parseSectionAttr() (name string, err error) {
	name, ok := p.try(token.String)
	if !ok {
		return "", errutil.Newf("expected section name, got %q token", p.next())
	}
	return name, nil
}
