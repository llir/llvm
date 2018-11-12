// Package asm implements a parser for LLVM IR assembly files.
package asm

import (
	"io"
	"io/ioutil"
	"log"
	"os"
	"time"

	"github.com/llir/ll/ast"
	"github.com/llir/llvm/ir"
	"github.com/mewkiz/pkg/term"
	"github.com/pkg/errors"
)

var (
	// dbg is a logger which logs debug messages with "asm:" prefix to standard error.
	dbg = log.New(os.Stderr, term.MagentaBold("asm:")+" ", 0)
)

// ParseFile parses the given LLVM IR assembly file into an LLVM IR module.
func ParseFile(path string) (*ir.Module, error) {
	buf, err := ioutil.ReadFile(path)
	if err != nil {
		return nil, errors.WithStack(err)
	}
	content := string(buf)
	return ParseString(path, content)
}

// Parse parses the given LLVM IR assembly file into an LLVM IR module, reading
// from r. An optional path to the source file may be specified for error
// reporting.
func Parse(path string, r io.Reader) (*ir.Module, error) {
	buf, err := ioutil.ReadAll(r)
	if err != nil {
		return nil, errors.WithStack(err)
	}
	content := string(buf)
	return ParseString(path, content)
}

// ParseBytes parses the given LLVM IR assembly file into an LLVM IR module,
// reading from b. An optional path to the source file may be specified for
// error reporting.
func ParseBytes(path string, b []byte) (*ir.Module, error) {
	content := string(b)
	return ParseString(path, content)
}

// ParseString parses the given LLVM IR assembly file into an LLVM IR module,
// reading from content. An optional path to the source file may be specified
// for error reporting.
func ParseString(path, content string) (*ir.Module, error) {
	parseStart := time.Now()
	tree, err := ast.Parse(path, content)
	if err != nil {
		return nil, errors.Wrapf(err, "unable to parse %q into AST", path)
	}
	root := ast.ToLlvmNode(tree.Root())
	dbg.Println("parsing into AST took:", time.Since(parseStart))
	return translate(root.(*ast.Module))
}
