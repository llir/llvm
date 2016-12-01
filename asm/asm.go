// Package asm implements a parser for LLVM IR assembly.
package asm

import (
	"io"
	"io/ioutil"

	"github.com/llir/llvm/asm/internal/lexer"
	"github.com/llir/llvm/asm/internal/parser"
	"github.com/llir/llvm/ir"
	"github.com/pkg/errors"
)

// ParseFile parses the given LLVM IR assembly file into an LLVM IR module.
func ParseFile(path string) (*ir.Module, error) {
	buf, err := ioutil.ReadFile(path)
	if err != nil {
		return nil, errors.WithStack(err)
	}
	return ParseBytes(buf)
}

// Parse parses the given LLVM IR assembly file into an LLVM IR module, reading
// from r.
func Parse(r io.Reader) (*ir.Module, error) {
	buf, err := ioutil.ReadAll(r)
	if err != nil {
		return nil, errors.WithStack(err)
	}
	return ParseBytes(buf)
}

// ParseBytes parses the given LLVM IR assembly file into an LLVM IR module,
// reading from b.
func ParseBytes(b []byte) (*ir.Module, error) {
	l := lexer.NewLexer(b)
	p := parser.NewParser()
	module, err := p.Parse(l)
	if err != nil {
		return nil, errors.WithStack(err)
	}
	m, ok := module.(*ir.Module)
	if !ok {
		return nil, errors.Errorf("invalid module type; expected *ir.Module, got %T", module)
	}
	return m, nil
}

// ParseString parses the given LLVM IR assembly file into an LLVM IR module,
// reading from s.
func ParseString(s string) (*ir.Module, error) {
	return ParseBytes([]byte(s))
}
