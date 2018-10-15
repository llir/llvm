package asm

import (
	"fmt"
	"sort"
	"strings"

	"github.com/kr/pretty"
	"github.com/llir/l/ir"
	"github.com/llir/l/ir/types"
	"github.com/mewmew/l-tm/asm/ll/ast"
	"github.com/mewmew/l-tm/internal/enc"
	"github.com/pkg/errors"
)

// Translate translates the AST of the given module to an equivalent LLVM IR
// module.
func Translate(module *ast.Module) (*ir.Module, error) {
	m := &ir.Module{}
	ts, err := resolveTypeDefs(module)
	if err != nil {
		return nil, errors.WithStack(err)
	}
	pretty.Println(ts)
	return m, nil
}

// resolveTypeDefs resolves the type definitions of the given module. The
// returned value maps from type name (without '%' prefix) to the underlying
// type.
func resolveTypeDefs(module *ast.Module) (map[string]types.Type, error) {
	// index maps from type name to underlying AST type.
	index := make(map[string]ast.LlvmNode)
	// Index named AST types.
	for _, entity := range module.TopLevelEntities() {
		switch entity := entity.(type) {
		case *ast.TypeDef:
			fmt.Printf("entity: %T\n", entity)
			alias := local(entity.Alias())
			fmt.Println("   alias:", alias)
			typ := entity.Typ()
			switch typ.(type) {
			case *ast.OpaqueType:
			case ast.Type:
			default:
				panic(fmt.Errorf("support for type %T not yet implemented", typ))
			}
			fmt.Println("   typ:", typ.Text())
			if prev, ok := index[alias]; ok {
				return nil, errors.Errorf("AST type definition with alias %q already present; prev `%s`, new `%s`", enc.Local(alias), prev.Text(), typ.Text())
			}
			index[alias] = typ
		}
	}

	// ts maps from type name to underlying IR type.
	ts := make(map[string]types.Type)
	// Create corresponding named IR types; without bodies.
	for alias, typ := range index {
		track := make(map[string]bool)
		t, err := newIRType(alias, typ, index, track)
		if err != nil {
			return nil, errors.WithStack(err)
		}
		if prev, ok := ts[alias]; ok {
			return nil, errors.Errorf("IR type definition with alias %q already present; prev `%s`, new `%s`", enc.Local(alias), prev.Def(), t.Def())
		}
		ts[alias] = t
	}

	// Translate type definitions.

	return ts, nil
}

// newIRType returns a new IR type (without body) based on the given AST type.
// Named types are resolved to their underlying type through lookup in index. An
// error is returned for (potentially recursive) self-referential name types.
//
// For instance, the following is disallowed.
//
//    ; self-referential named type.
//    %a = type %a
//
//    ; recursively self-referential named types.
//    %b = type %c
//    %c = type %b
//
// The following is allowed, however.
//
//    ; struct type containing pointer to itself.
//    %d = type { %d* }
func newIRType(alias string, typ ast.LlvmNode, index map[string]ast.LlvmNode, track map[string]bool) (types.Type, error) {
	switch typ := typ.(type) {
	case *ast.OpaqueType:
		return &types.StructType{Alias: alias, Opaque: true}, nil
	case *ast.ArrayType:
		return &types.ArrayType{Alias: alias}, nil
	case *ast.FloatType:
		return &types.FloatType{Alias: alias}, nil
	case *ast.FuncType:
		return &types.FuncType{Alias: alias}, nil
	case *ast.IntType:
		return &types.IntType{Alias: alias}, nil
	case *ast.LabelType:
		return &types.LabelType{Alias: alias}, nil
	case *ast.MMXType:
		return &types.MMXType{Alias: alias}, nil
	case *ast.MetadataType:
		return &types.MetadataType{Alias: alias}, nil
	case *ast.NamedType:
		if track[alias] {
			var names []string
			for name := range track {
				names = append(names, enc.Local(name))
			}
			sort.Strings(names)
			return nil, errors.Errorf("self-referential named type with type name(s) %s", strings.Join(names, ", "))
		}
		track[alias] = true
		newAlias := local(typ.Name())
		newTyp := index[newAlias]
		return newIRType(newAlias, newTyp, index, track)
	case *ast.PointerType:
		return &types.PointerType{Alias: alias}, nil
	case *ast.StructType:
		return &types.StructType{Alias: alias}, nil
	case *ast.TokenType:
		return &types.TokenType{Alias: alias}, nil
	case *ast.VectorType:
		return &types.VectorType{Alias: alias}, nil
	case *ast.VoidType:
		return &types.VoidType{Alias: alias}, nil
	default:
		panic(fmt.Errorf("support for type %T not yet implemented", typ))
	}
}

// local returns the name of the local identifier.
func local(l ast.LocalIdent) string {
	text := l.Text()
	const prefix = "%"
	if !strings.HasPrefix(text, prefix) {
		// NOTE: Panic instead of returning error as this case should not be
		// possible given the grammar.
		panic(fmt.Errorf("invalid local identifier %q; missing '%s' prefix", text, prefix))
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
