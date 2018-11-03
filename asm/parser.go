// Package asm implements a parser for LLVM IR assembly files.
package asm

import (
	"fmt"
	"io/ioutil"
	"time"

	"github.com/llir/ll/ast"
	"github.com/llir/llvm/ir"
	"github.com/pkg/errors"
)

// ParseFile parses the given LLVM IR assembly file into an LLVM IR module.
func ParseFile(path string) (*ir.Module, error) {
	buf, err := ioutil.ReadFile(path)
	if err != nil {
		return nil, errors.WithStack(err)
	}
	content := string(buf)
	return Parse(path, content)
}

// Parse parses the given LLVM IR assembly file into an LLVM IR module, reading
// from content.
func Parse(path, content string) (*ir.Module, error) {
	parseStart := time.Now()
	tree, err := ast.Parse(path, content)
	if err != nil {
		return nil, errors.WithStack(err)
	}
	root := ast.ToLlvmNode(tree.Root())
	fmt.Println("parsing into AST took:", time.Since(parseStart))
	fmt.Println()
	return translate(root.(*ast.Module))
}
