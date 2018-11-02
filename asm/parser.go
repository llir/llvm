// Package asm implements a parser for LLVM IR assembly files.
package asm

import (
	"io/ioutil"

	"github.com/mewmew/l-tm/asm/ll/ast"
	"github.com/pkg/errors"
)

// ParseFile parses the given LLVM IR assembly file into an LLVM IR module.
func ParseFile(path string) (*ast.Module, error) {
	buf, err := ioutil.ReadFile(path)
	if err != nil {
		return nil, errors.WithStack(err)
	}
	content := string(buf)
	return Parse(path, content)
}

// Parse parses the given LLVM IR assembly file into an LLVM IR module, reading
// from content.
func Parse(path, content string) (*ast.Module, error) {
	tree, err := ast.Parse(path, content)
	if err != nil {
		return nil, errors.WithStack(err)
	}
	root := ast.ToLlvmNode(tree.Root())
	return root.(*ast.Module), nil
}
