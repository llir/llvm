package asm

import (
	"fmt"
	"strings"

	"github.com/llir/l/ir"
	"github.com/mewmew/l-tm/asm/ll/ast"
	"github.com/mewmew/l-tm/internal/enc"
)

// Translate translates the AST of the given module to an equivalent LLVM IR
// module.
func Translate(module *ast.Module) *ir.Module {
	m := &ir.Module{}
	// Translate type definitions.
	for _, entity := range module.TopLevelEntities() {
		switch entity := entity.(type) {
		case *ast.TypeDef:
			fmt.Println("alias:", local(entity.Alias()))
		default:
			fmt.Printf("entity: %T\n", entity)
		}
	}
	return m
}

// local returns the name of the local identifier.
func local(l ast.LocalIdent) string {
	text := l.Text()
	const prefix = "%"
	if !strings.HasPrefix(text, prefix) {
		panic(fmt.Errorf("invalid local identifier `%s`; missing '%s' prefix", text, prefix))
	}
	text = text[len(prefix):]
	return unquote(text)
}

// unquote returns the unquoted version of s if quoted, and the original string
// otherwise.
func unquote(s string) string {
	if len(s) >= 2 && strings.HasPrefix(s, `"`) && strings.HasSuffix(s, `"`) {
		return string(enc.Unquote(s))
	}
	return s
}
