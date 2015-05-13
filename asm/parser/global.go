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
//    GlobalDecl = GlobalName "=" DeclLinkage [ Visibility ] [ DLLStorage ] [ ThreadLocal ] ( "global" | "constant" ) Type .
//    GlobalDef  = GlobalName "=" [ DefLinkage ] [ Visibility ] [ DLLStorage ] [ ThreadLocal ] ( "global" | "constant" ) Type InitValue .
//    GlobalName = Global .
//    InitValue  = Const .
//
// Examples:
//    @x = global i32 7
//    @y = external global i32
//    @hello = constant [13 x i8] c"hello world\0A\00"
//
// References:
//   http://llvm.org/docs/LangRef.html#global-variables
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

	// Consume optional thread local storage model.
	if err := p.tryThreadLocal(); err != nil {
		return errutil.Err(err)
	}

	immutable := false
	switch tok := p.next(); tok.Kind {
	case token.KwGlobal:
	case token.KwConstant:
		immutable = true
	default:
		decl := "definition"
		if extern {
			decl = "declaration"
		}
		return errutil.Newf(`invalid global variable %s for %q; expected "global" or "constant", got %q`, decl, asm.EncGlobal(name.Val), tok)
	}
	var global *ir.GlobalDecl
	if extern {
		typ, err := p.parseType()
		if err != nil {
			return errutil.Err(err)
		}
		global = ir.NewGlobalDecl(name.Val, typ, immutable)
	} else {
		val, err := p.parseConst()
		if err != nil {
			return errutil.Err(err)
		}
		global = ir.NewGlobalDef(name.Val, val, immutable)
	}
	p.m.Globals = append(p.m.Globals, global)
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
