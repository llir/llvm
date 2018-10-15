package asm

import (
	"fmt"
	"sort"
	"strconv"
	"strings"
	"time"

	"github.com/llir/l/ir"
	"github.com/llir/l/ir/types"
	"github.com/mewmew/l-tm/asm/ll/ast"
	"github.com/mewmew/l-tm/internal/enc"
	"github.com/pkg/errors"
)

// TODO: remove flag after we reach our performance goals.

// DoTypeResolution enables type resolution of type defintions.
var DoTypeResolution = true

// Translate translates the AST of the given module to an equivalent LLVM IR
// module.
func Translate(module *ast.Module) (*ir.Module, error) {
	gen := newGenerator()
	if DoTypeResolution {
		typeResolutionStart := time.Now()
		_, err := gen.resolveTypeDefs(module)
		if err != nil {
			return nil, errors.WithStack(err)
		}
		fmt.Println("type resolution of type definitions took:", time.Since(typeResolutionStart))
		fmt.Println()
	}
	return gen.m, nil
}

// generator keeps track of global and local identifiers when translating values
// and types from AST to IR representation.
type generator struct {
	// LLVM IR module being generated.
	m *ir.Module
}

// newGenerator returns a new generator for translating an LLVM IR module from
// AST to IR representation.
func newGenerator() *generator {
	return &generator{
		m: &ir.Module{},
	}
}

// resolveTypeDefs resolves the type definitions of the given module. The
// returned value maps from type name (without '%' prefix) to the underlying
// type.
func (gen *generator) resolveTypeDefs(module *ast.Module) (map[string]types.Type, error) {
	// index maps from type name to underlying AST type.
	index := make(map[string]ast.LlvmNode)
	// Index named AST types.
	var order []string
	for _, entity := range module.TopLevelEntities() {
		switch entity := entity.(type) {
		case *ast.TypeDef:
			alias := local(entity.Alias())
			order = append(order, alias)
			typ := entity.Typ()
			switch typ.(type) {
			case *ast.OpaqueType:
			case ast.Type:
			default:
				panic(fmt.Errorf("support for type %T not yet implemented", typ))
			}
			if prev, ok := index[alias]; ok {
				if _, ok := prev.(*ast.OpaqueType); !ok {
					return nil, errors.Errorf("AST type definition with alias %q already present; prev `%s`, new `%s`", enc.Local(alias), prev.Text(), typ.Text())
				}
			}
			index[alias] = typ
		}
	}

	// ts maps from type name to underlying IR type.
	ts := make(map[string]types.Type)
	// Create corresponding named IR types (without bodies).
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

	// Translate type defintions (including bodies).
	for alias, typ := range index {
		t := ts[alias]
		_, err := translateIRType(t, typ, ts)
		if err != nil {
			return nil, errors.WithStack(err)
		}
	}

	// Add type definitions to the IR module, in the same order of occurrance as
	// the input.
	added := make(map[string]bool)
	for _, key := range order {
		if added[key] {
			// Add only the first type definition of each type name.
			//
			// Type definitions of opaque types may contain several type
			// definitions with the same type name.
			continue
		}
		added[key] = true
		t := ts[key]
		gen.m.TypeDefs = append(gen.m.TypeDefs, t)
	}

	return ts, nil
}

// translateIRType translates the AST type into an equivalent IR type. A new IR
// type correspoding to the AST type is created if t is nil, otherwise the body
// of t is populated. Named types are resolved through ts.
func translateIRType(t types.Type, old ast.LlvmNode, ts map[string]types.Type) (types.Type, error) {
	switch old := old.(type) {
	case *ast.OpaqueType:
		typ, ok := t.(*types.StructType)
		if t == nil {
			// NOTE: Panic instead of returning error as this case should not be
			// possible given the grammar.
			panic("invalid use of opaque type; only allowed in type definitions")
		} else if !ok {
			panic(fmt.Errorf("invalid IR type for AST opaque type; expected *types.StructType, got %T", t))
		}
		// nothing to do.
		return typ, nil
	case *ast.ArrayType:
		typ, ok := t.(*types.ArrayType)
		if t == nil {
			typ = &types.ArrayType{}
		} else if !ok {
			panic(fmt.Errorf("invalid IR type for AST array type; expected *types.ArrayType, got %T", t))
		}
		// Array length.
		len := uintLit(old.Len())
		typ.Len = int64(len)
		// Element type.
		elem, err := translateIRType(nil, old.Elem(), ts)
		if err != nil {
			return nil, errors.WithStack(err)
		}
		typ.ElemType = elem
		return typ, nil
	case *ast.FloatType:
		typ, ok := t.(*types.FloatType)
		if t == nil {
			typ = &types.FloatType{}
		} else if !ok {
			panic(fmt.Errorf("invalid IR type for AST floating-point type; expected *types.FloatType, got %T", t))
		}
		// Floating-point kind.
		typ.Kind = irFloatKind(old.FloatKind())
		return typ, nil
	case *ast.FuncType:
		typ, ok := t.(*types.FuncType)
		if t == nil {
			typ = &types.FuncType{}
		} else if !ok {
			panic(fmt.Errorf("invalid IR type for AST function type; expected *types.FuncType, got %T", t))
		}
		// Return type.
		retType, err := translateIRType(nil, old.RetType(), ts)
		if err != nil {
			return nil, errors.WithStack(err)
		}
		typ.RetType = retType
		// Function parameters.
		ps := old.Params()
		for _, p := range ps.Params() {
			param, err := translateIRType(nil, p.Typ(), ts)
			if err != nil {
				return nil, errors.WithStack(err)
			}
			typ.Params = append(typ.Params, param)
		}
		// Variadic.
		typ.Variadic = irOptVariadic(ps.Variadic())
		return typ, nil
	case *ast.IntType:
		typ, ok := t.(*types.IntType)
		if t == nil {
			typ = &types.IntType{}
		} else if !ok {
			panic(fmt.Errorf("invalid IR type for AST integer type; expected *types.IntType, got %T", t))
		}
		// Bit size.
		typ.BitSize = irIntTypeBitSize(old)
		return typ, nil
	case *ast.LabelType:
		typ, ok := t.(*types.LabelType)
		if t == nil {
			typ = &types.LabelType{}
		} else if !ok {
			panic(fmt.Errorf("invalid IR type for AST label type; expected *types.LabelType, got %T", t))
		}
		// nothing to do.
		return typ, nil
	case *ast.MMXType:
		typ, ok := t.(*types.MMXType)
		if t == nil {
			typ = &types.MMXType{}
		} else if !ok {
			panic(fmt.Errorf("invalid IR type for AST MMX type; expected *types.MMXType, got %T", t))
		}
		// nothing to do.
		return typ, nil
	case *ast.MetadataType:
		typ, ok := t.(*types.MetadataType)
		if t == nil {
			typ = &types.MetadataType{}
		} else if !ok {
			panic(fmt.Errorf("invalid IR type for AST metadata type; expected *types.MetadataType, got %T", t))
		}
		// nothing to do.
		return typ, nil
	case *ast.NamedType:
		// Resolve named type.
		alias := local(old.Name())
		if _, ok := ts[alias]; !ok {
			return nil, errors.Errorf("unable to locate type definition of named type %q", enc.Local(alias))
		}
		return ts[alias], nil
	case *ast.PointerType:
		typ, ok := t.(*types.PointerType)
		if t == nil {
			typ = &types.PointerType{}
		} else if !ok {
			panic(fmt.Errorf("invalid IR type for AST pointer type; expected *types.PointerType, got %T", t))
		}
		// Element type.
		elemType, err := translateIRType(nil, old.Elem(), ts)
		if err != nil {
			return nil, errors.WithStack(err)
		}
		typ.ElemType = elemType
		// Address space.
		typ.AddrSpace = irOptAddrSpace(old.AddrSpace())
		return typ, nil
	case *ast.StructType:
		typ, ok := t.(*types.StructType)
		if t == nil {
			typ = &types.StructType{}
		} else if !ok {
			panic(fmt.Errorf("invalid IR type for AST struct type; expected *types.StructType, got %T", t))
		}
		// Packed.
		// TODO: Figure out how to represent packed in grammar.
		// Fields.
		for _, f := range old.Fields() {
			field, err := translateIRType(nil, f, ts)
			if err != nil {
				return nil, errors.WithStack(err)
			}
			typ.Fields = append(typ.Fields, field)
		}
		// struct body now present.
		typ.Opaque = false
		return typ, nil
	case *ast.TokenType:
		typ, ok := t.(*types.TokenType)
		if t == nil {
			typ = &types.TokenType{}
		} else if !ok {
			panic(fmt.Errorf("invalid IR type for AST token type; expected *types.TokenType, got %T", t))
		}
		// nothing to do.
		return typ, nil
	case *ast.VectorType:
		typ, ok := t.(*types.VectorType)
		if t == nil {
			typ = &types.VectorType{}
		} else if !ok {
			panic(fmt.Errorf("invalid IR type for AST vector type; expected *types.VectorType, got %T", t))
		}
		// Vector length.
		len := uintLit(old.Len())
		typ.Len = int64(len)
		// Element type.
		elem, err := translateIRType(nil, old.Elem(), ts)
		if err != nil {
			return nil, errors.WithStack(err)
		}
		typ.ElemType = elem
		return typ, nil
	case *ast.VoidType:
		typ, ok := t.(*types.VoidType)
		if t == nil {
			typ = &types.VoidType{}
		} else if !ok {
			panic(fmt.Errorf("invalid IR type for AST void type; expected *types.VoidType, got %T", t))
		}
		// nothing to do.
		return typ, nil
	default:
		panic(fmt.Errorf("support for type %T not yet implemented", old))
	}
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
			return nil, errors.Errorf("invalid named type; self-referential with type name(s) %s", strings.Join(names, ", "))
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

// local returns the name (without '%' prefix) of the given local identifier.
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

// uintLit returns the unsigned integer value corresponding to the given
// unsigned integer literal.
func uintLit(l ast.UintLit) uint64 {
	text := l.Text()
	x, err := strconv.ParseUint(text, 10, 64)
	if err != nil {
		// NOTE: Panic instead of returning error as this case should not be
		// possible given the grammar.

		// TODO: figure out how to update the grammar for UintLit to remove the
		// optional sign.
		panic(fmt.Errorf("unable to parse unsigned integer literal %q; %v", text, err))
	}
	return x
}

// unquote returns the unquoted version of s if quoted, and the original string
// otherwise.
func unquote(s string) string {
	if len(s) >= 2 && strings.HasPrefix(s, `"`) && strings.HasSuffix(s, `"`) {
		return string(enc.Unquote(s))
	}
	return s
}

// irFloatKind returns the IR floating-point kind corresponding to the given AST
// floating-point kind.
func irFloatKind(kind ast.FloatKind) types.FloatKind {
	text := kind.Text()
	switch text {
	case "half":
		return types.FloatKindHalf
	case "float":
		return types.FloatKindFloat
	case "double":
		return types.FloatKindDouble
	case "x86_fp80":
		return types.FloatKindX86FP80
	case "fp128":
		return types.FloatKindFP128
	case "ppc_fp128":
		return types.FloatKindPPCFP128
	default:
		panic(fmt.Errorf("support for floating-point kind %q not yet implemented", text))
	}
}

// irOptVariadic returns the variadic boolean corresponding to the given AST
// ellipsis.
func irOptVariadic(n *ast.Ellipsis) bool {
	// TODO: check why Variadic is non-nil for `Variadic=Ellipsisopt`, regardless
	// of whether the input is (...) or ().
	//
	// It seems that the Variadic.Text simply returns empty string when
	// Ellipsisopt reduces to \empty.
	//
	// Using `n.Text() == "..."` for now, would like to use `n != nil`.
	return n.Text() == "..."
}

// irIntTypeBitSize returns the integer type bit size corresponding to the given
// AST integer type.
func irIntTypeBitSize(n *ast.IntType) int64 {
	text := n.Text()
	const prefix = "i"
	if !strings.HasPrefix(text, prefix) {
		// NOTE: Panic instead of returning error as this case should not be
		// possible given the grammar.
		panic(fmt.Errorf("invalid integer type %q; missing '%s' prefix", text, prefix))
	}
	text = text[len(prefix):]
	x, err := strconv.ParseInt(text, 10, 64)
	if err != nil {
		// NOTE: Panic instead of returning error as this case should not be
		// possible given the grammar.
		panic(fmt.Errorf("unable to parse integer type bit size %q; %v", text, err))
	}
	return x
}

// irOptAddrSpace returns the IR address space corresponding to the given AST
// address space.
func irOptAddrSpace(n *ast.AddrSpace) types.AddrSpace {
	// \empty is used when address space not present.
	if n.Text() == "" {
		return 0
	}
	x := uintLit(n.N())
	return types.AddrSpace(x)
}
