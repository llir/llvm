// Package parser implements syntactic analysis of LLVM IR assembly.
package parser

import (
	"io"
	"log"

	"github.com/mewlang/llvm/asm/lexer"
	"github.com/mewlang/llvm/asm/token"
	"github.com/mewlang/llvm/ir"
	"github.com/mewlang/llvm/types"
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

// ParseTokens parses the tokenized input into an in-memory representation of
// LLVM IR.
func ParseTokens(input []token.Token) (*ir.Module, error) {
	p := &parser{
		// filter input to a supported subset of the LLVM IR tokens.
		input: filter(input),
	}

	// Parse the tokenized input by repeatedly parsing top-level entities.
	module := new(ir.Module)
	for {
		err := p.parseTopLevelEntity(module)
		if err != nil {
			if err == io.EOF {
				// Terminate the parser at EOF.
				return module, nil
			}
			// TODO: Remove debug output.
			log.Printf("error at pos=%d (%q)\n", p.cur, p.input[p.cur:])
			return module, err
		}
	}
}

// A parser parses the tokenized input into an in-memory representation of LLVM
// IR.
type parser struct {
	// Tokenized input.
	input []token.Token
	// Current position in the input.
	cur int
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

// accept consumes the next token if it is of the given kind. It returns true if
// a token was consumed and false otherwise.
func (p *parser) accept(kind token.Kind) bool {
	tok := p.next()
	if kind == tok.Kind {
		return true
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
