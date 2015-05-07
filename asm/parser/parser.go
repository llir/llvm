// Package parser implements syntactic analysis of LLVM IR assembly.
package parser

import (
	"io"
	"log"

	"github.com/llir/llvm/asm/lexer"
	"github.com/llir/llvm/asm/token"
	"github.com/llir/llvm/ir"
	"github.com/llir/llvm/types"
	"github.com/mewkiz/pkg/errutil"
)

// Parse parses the input read from r into an in-memory representation of LLVM
// IR.
func Parse(r io.Reader) (*ir.Module, error) {
	input, err := lexer.Parse(r)
	if err != nil {
		return nil, err
	}
	return ParseTokens(input)
}

// ParseFile parses the input read from path into an in-memory representation of
// LLVM IR.
func ParseFile(path string) (*ir.Module, error) {
	input, err := lexer.ParseFile(path)
	if err != nil {
		return nil, err
	}
	return ParseTokens(input)
}

// ParseString parses the input string into an in-memory representation of LLVM
// IR.
func ParseString(s string) (*ir.Module, error) {
	input := lexer.ParseString(s)
	return ParseTokens(input)
}

// ParseTokens parses the tokenized input into an in-memory representation of
// LLVM IR.
func ParseTokens(input []token.Token) (*ir.Module, error) {
	// Parse the tokenized input by repeatedly parsing top-level entities.
	p := &parser{
		input: filter(input),
		m:     new(ir.Module),
		tctx:  types.NewContext(),
	}
	for {
		err := p.parseTopLevelEntity()
		if err != nil {
			if err == io.EOF {
				break
			}
			// TODO: Remove debug output.
			log.Printf("error at pos=%d (%q)", p.cur, p.input[p.cur:])
			return p.m, errutil.Err(err)
		}
	}
	// Validate that the identified structure types have defined type bodies.
	if err := p.tctx.Validate(); err != nil {
		return p.m, errutil.Err(err)
	}
	return p.m, nil
}

// A parser parses the tokenized input into an in-memory representation of LLVM
// IR.
type parser struct {
	// Tokenized input.
	input []token.Token
	// Start position of the current entity.
	start int
	// Current position in the input.
	cur int
	// An in-memory representation of the parsed LLVM IR module.
	m *ir.Module
	// Type context.
	tctx *types.Context
}

// next consumes and returns the next token of the input.
func (p *parser) next() token.Token {
	if p.cur >= len(p.input) {
		panic("invalid call to next; end of input")
	}
	tok := p.input[p.cur]
	p.cur++
	return tok
}

// backup backs up one token in the input. It can only be called once per call
// to next.
func (p *parser) backup() {
	p.cur--
}

// accept consumes the next token if it is from the valid set. It returns true
// if a token was consumed and false otherwise.
func (p *parser) accept(valid ...token.Kind) bool {
	tok := p.next()
	for _, kind := range valid {
		if kind == tok.Kind {
			return true
		}
	}
	p.backup()
	return false
}

// try tries to consume a token from the valid set and return its value. The
// value of ok is true if a token was consumed this way, and false otherwise.
func (p *parser) try(valid ...token.Kind) (val string, ok bool) {
	tok := p.next()
	for _, kind := range valid {
		if kind == tok.Kind {
			return tok.Val, true
		}
	}
	p.backup()
	return "", false
}

// tryGlobal tries to consume a global (GlobalVar or GlobalID) and return its
// name. The value of ok is true if a token was consumed this way, and false
// otherwise.
func (p *parser) tryGlobal() (name string, ok bool) {
	return p.try(token.GlobalVar, token.GlobalID)
}

// tryLocal tries to consume a local (LocalVar or LocalID) and return its name.
// The value of ok is true if a token was consumed this way, and false
// otherwise.
func (p *parser) tryLocal() (name string, ok bool) {
	return p.try(token.LocalVar, token.LocalID)
}

// tryType tries to consume a type. The value of ok is true if one or more
// tokens were consumed this way, and false otherwise.
func (p *parser) tryType() (typ types.Type, ok bool) {
	cur := p.cur
	typ, err := p.parseType()
	if err != nil {
		p.cur = cur
		return nil, false
	}
	return typ, true
}

// TODO: Check which support methods are actually required and remove the rest.
//    * try?
//    * tryLocal?
//    * tryGlobal?
//    * tryType?
//    * expected?

// expect consumes the next token and returns its value after validating the
// expected token kind.
func (p *parser) expect(kind token.Kind) (s string, err error) {
	tok := p.next()
	if tok.Kind != kind {
		return "", errutil.Newf("expected %v, got %v (%q)", pretty(kind), pretty(tok.Kind), tok)
	}
	return tok.Val, nil
}

// pretty returns a pretty-printed version of the given token kind.
func pretty(kind token.Kind) string {
	var m = map[token.Kind]string{
		token.Equal:  `"="`,
		token.Int:    "integer literal",
		token.Float:  "floating point literal",
		token.String: "string literal",
	}
	if s, ok := m[kind]; ok {
		return s
	}
	log.Printf("not yet implemented; pretty-printing of token kind %v", kind)
	return kind.String()
}
