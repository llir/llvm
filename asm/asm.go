// Package asm provides support for reading and writing the assembly language
// representation of LLVM IR.
package asm

import (
	"go/ast"
	"io"
	"io/ioutil"

	"github.com/llir/llvm/asm/internal/lexer"
	"github.com/llir/llvm/asm/internal/parser"
	"github.com/mewkiz/pkg/errutil"
)

// ParseFile parses the given LLVM IR assembly file into an LLVM IR module.
func ParseFile(path string) (*ast.File, error) {
	buf, err := ioutil.ReadFile(path)
	if err != nil {
		return nil, errutil.Err(err)
	}
	return ParseBytes(buf)
}

// Parse parses the given LLVM IR assembly file into an LLVM IR module, reading
// from r.
func Parse(r io.Reader) (*ast.File, error) {
	buf, err := ioutil.ReadAll(r)
	if err != nil {
		return nil, errutil.Err(err)
	}
	return ParseBytes(buf)
}

// ParseBytes parses the given LLVM IR assembly file into an LLVM IR module,
// reading from b.
func ParseBytes(b []byte) (*ast.File, error) {
	l := lexer.NewLexer(b)
	p := parser.NewParser()
	file, err := p.Parse(l)
	if err != nil {
		return nil, errutil.Err(err)
	}
	f, ok := file.(*ast.File)
	if !ok {
		return nil, errutil.Newf("invalid file type; expected *ast.File, got %T", file)
	}
	if err := check(f); err != nil {
		return nil, errutil.Err(err)
	}
	return f, nil
}

// ParseString parses the given LLVM IR assembly file into an LLVM IR module,
// reading from s.
func ParseString(s string) (*ast.File, error) {
	return ParseBytes([]byte(s))
}
