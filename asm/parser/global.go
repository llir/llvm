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
//    GlobalDecl = GlobalName "=" DeclLinkage ( "global" | "constant" ) Type .
//    GlobalDef  = GlobalName "=" [ DefLinkage ] ( "global" | "constant" ) Type InitValue .
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
	name := p.next()
	if !p.accept(token.Equal) {
		return errutil.Newf(`expected "=" after global variable name %q, got %q token`, asm.EncGlobal(name.Val), p.next())
	}
	extern := p.tryLinkage()
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

// tryLinkage parses the optional linkage type of a global variable or function,
// and reports whether it is externally visible.
//
// Syntax:
//    Linkage = DeclLinkage | DefLinkage .
//
//    DeclLinkage = "extern_weak" |
//                  "external" .
//    DefLinkage  = "appending" |
//                  "available_externally" |
//                  "common" |
//                  "internal" |
//                  "linkonce" |
//                  "linkonce_odr" |
//                  "private" |
//                  "weak" |
//                  "weak_odr" .
//
// References:
//    http://llvm.org/docs/LangRef.html#linkage
func (p *parser) tryLinkage() (extern bool) {
	// NOTE: Only a subset of the linkage information is retained, namely the
	// external visibility.
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
