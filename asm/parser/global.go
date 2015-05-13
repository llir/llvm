package parser

import (
	"github.com/llir/llvm/asm"
	"github.com/llir/llvm/asm/token"
	"github.com/llir/llvm/ir"
	"github.com/mewkiz/pkg/errutil"
)

// parseGlobalDecl parses a global variable definition or declaration. The next
// token is either a GlobalID or a GlobalVar.
//
//    GlobalDef  = GlobalName "=" ( "global" | "constant" ) Type InitValue .
//    GlobalDecl = GlobalName "=" "external" ( "global" | "constant" ) Type .
//    GlobalName = Global .
//    InitValue  = Const .
func (p *parser) parseGlobalDecl() error {
	name := p.next()
	if !p.accept(token.Equal) {
		return errutil.Newf(`expected "=" after global variable name %q, got %q token`, asm.EncGlobal(name.Val), p.next())
	}
	extern := p.accept(token.KwExternal)
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
