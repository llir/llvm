// Package parser implements syntactic analysis of LLVM IR assembly.
package parser

import (
	"io"

	"github.com/mewlang/llvm/asm/lexer"
	"github.com/mewlang/llvm/asm/token"
	"github.com/mewlang/llvm/ir"
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