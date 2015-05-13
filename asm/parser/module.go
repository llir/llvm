package parser

// Characters
//    newline      = /* the Unicode code point U+000A */ .
//    unicode_char = /* an arbitrary Unicode code point except newline */ .
//
// Letters and digits
//    letter        = ( "-" | "a" … "z" | "A" … "Z" | "$" | "." | "_" ) .
//    decimal_digit = "0" … "9" .
//    hex_digit     = "0" … "9" | "A" … "F" | "a" … "f" .
//
// Identifiers
//    Global    = GlobalID | GlobalVar .
//    Local     = LocalID | LocalVar .
//    GlobalID  = "@" ID .
//    GlobalVar = "@" Var .
//    LocalID   = "%" ID .
//    LocalVar  = "%" Var .
//    ID        = int_lit .
//    Var       = letter { letter | decimal_digit } | string_lit .
//
// Integer literals
//    int_lit = decimal_digit { decimal_digit } .
//
// Floating-point literals
//    float_lit = [ "+" | "-" ] decimals "." [ decimals ] [ exponent ] .
//    decimals  = decimal_digit { decimal_digit } .
//    exponent  = ( "e" | "E" ) [ "+" | "-" ] decimals .
//
// String literals
//    string_lit     =  `"` { unicode_value | newline } `"` .
//    unicode_value  = unicode_char | hex_byte_value .
//    hex_byte_value = `\` hex_digit hex_digit .

import (
	"io"

	"github.com/llir/llvm/asm/token"
	"github.com/mewkiz/pkg/errutil"
)

// TODO: Complete TopLevelEntity EBNF definition.

// parseTopLevelEntity parses a top-level entity and stores it in the module.
//
// Syntax:
//    TopLevelEntity = TargetSpec | TypeDef | FuncDecl | FuncDef .
func (p *parser) parseTopLevelEntity() error {
	switch tok := p.next(); tok.Kind {
	case token.Error:
		return errutil.New(tok.Val)

	case token.EOF:
		// Terminate the parser at EOF.
		return io.EOF

	// Target specification; e.g.
	//    target datalayout = "foo"
	//    target triple = "foo"
	case token.KwTarget:
		return p.parseTargetSpec()

	// Type definition (named types and type aliases); e.g.
	//    %x     = type i32
	//    %y     = type i32
	//    %point = type {%x, %y}
	case token.LocalID, token.LocalVar:
		p.backup()
		return p.parseTypeDef()

	// Global variable definition or external global variable declaration; e.g.
	//    @x = global i32 42
	//    @y = external global i32
	case token.GlobalID, token.GlobalVar:
		p.backup()
		return p.parseGlobalDecl()

	// External function declaration; e.g.
	//    declare i32 @printf(i8*, ...)
	case token.KwDeclare:
		return p.parseFuncDecl()

	// Function definition; e.g.
	//    define i32 @main() {
	//       ret i32 42
	//    }
	case token.KwDefine:
		return p.parseFuncDef()

	default:
		return errutil.Newf("invalid token kind %q (%q); expected top-level entity", tok.Kind, tok.Val)
	}
}

// parseTargetSpec parses a target specification and stores it in the module. A
// "target" token has already been consumed.
//
// Syntax:
//    TargetSpec   = "target" ( DataLayout | TargetTriple ) .
//    DataLayout   = "datalayout" "=" string_lit .
//    TargetTriple = "triple" "=" string_lit .
//
// Examples:
//    target datalayout = "e-m:e-i64:64-f80:128-n8:16:32:64-S128"
//    target triple = "x86_64-unknown-linux-gnu"
//
// References:
//    http://llvm.org/docs/LangRef.html#data-layout
//    http://llvm.org/docs/LangRef.html#target-triple
func (p *parser) parseTargetSpec() error {
	property := p.next()
	switch property.Kind {
	case token.KwDatalayout, token.KwTriple:
		// valid.
	default:
		return errutil.Newf("unknown target property %q", property)
	}
	if !p.accept(token.Equal) {
		return errutil.Newf(`expected "=" after target %s, got %q token`, property, p.next())
	}
	s, err := p.expect(token.String)
	if err != nil {
		return errutil.Err(err)
	}
	switch property.Kind {
	case token.KwDatalayout:
		p.m.Layout = s
	case token.KwTriple:
		p.m.Target = s
	}
	return nil
}
