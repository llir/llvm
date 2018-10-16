package asm

import (
	"fmt"
	"sort"
	"strconv"
	"strings"

	"github.com/llir/l/ir/types"
	"github.com/mewmew/l-tm/asm/ll/ast"
	"github.com/mewmew/l-tm/internal/enc"
	"github.com/pkg/errors"
)

// resolveTypeDefs resolves the type definitions of the given module. The
// returned value maps from type name (without '%' prefix) to the underlying
// type.
func (gen *generator) resolveTypeDefs(module *ast.Module) (map[string]types.Type, error) {
	// index maps from type name to underlying AST type.
	index := make(map[string]ast.LlvmNode)
	// Record order of type definitions.
	added := make(map[string]bool)
	var order []string
	// Index named AST types.
	for _, entity := range module.TopLevelEntities() {
		switch entity := entity.(type) {
		case *ast.TypeDef:
			alias := local(entity.Alias())
			if !added[alias] {
				// Only record the first type definition of each type name.
				//
				// Type definitions of opaque types may contain several type
				// definitions with the same type name.
				order = append(order, alias)
			}
			added[alias] = true
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

	// Create corresponding named IR types (without bodies).
	gen.ts = make(map[string]types.Type)
	for alias, typ := range index {
		// track is used to identify self-referential named types.
		track := make(map[string]bool)
		t, err := newIRType(alias, typ, index, track)
		if err != nil {
			return nil, errors.WithStack(err)
		}
		if prev, ok := gen.ts[alias]; ok {
			return nil, errors.Errorf("IR type definition with alias %q already present; prev `%s`, new `%s`", enc.Local(alias), prev.Def(), t.Def())
		}
		gen.ts[alias] = t
	}

	// Translate type defintions (including bodies).
	for alias, typ := range index {
		t := gen.ts[alias]
		_, err := gen.translateType(t, typ)
		if err != nil {
			return nil, errors.WithStack(err)
		}
	}

	// Add type definitions to IR module in order of occurrence in input.
	for _, key := range order {
		t := gen.ts[key]
		gen.m.TypeDefs = append(gen.m.TypeDefs, t)
	}
	return gen.ts, nil
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

// === [ Types ] ===============================================================

// translateType translates the AST type into an equivalent IR type. A new IR
// type correspoding to the AST type is created if t is nil, otherwise the body
// of t is populated. Named types are resolved through ts.
func (gen *generator) translateType(t types.Type, old ast.LlvmNode) (types.Type, error) {
	switch old := old.(type) {
	case *ast.OpaqueType:
		return gen.translateOpaqueType(t, old)
	case *ast.ArrayType:
		return gen.translateArrayType(t, old)
	case *ast.FloatType:
		return gen.translateFloatType(t, old)
	case *ast.FuncType:
		return gen.translateFuncType(t, old)
	case *ast.IntType:
		return gen.translateIntType(t, old)
	case *ast.LabelType:
		return gen.translateLabelType(t, old)
	case *ast.MMXType:
		return gen.translateMMXType(t, old)
	case *ast.MetadataType:
		return gen.translateMetadataType(t, old)
	case *ast.NamedType:
		return gen.translateNamedType(t, old)
	case *ast.PointerType:
		return gen.translatePointerType(t, old)
	case *ast.StructType:
		return gen.translateStructType(t, old)
	case *ast.TokenType:
		return gen.translateTokenType(t, old)
	case *ast.VectorType:
		return gen.translateVectorType(t, old)
	case *ast.VoidType:
		return gen.translateVoidType(t, old)
	default:
		panic(fmt.Errorf("support for type %T not yet implemented", old))
	}
}

// --- [ Void Types ] ----------------------------------------------------------

func (gen *generator) translateVoidType(t types.Type, old *ast.VoidType) (types.Type, error) {
	typ, ok := t.(*types.VoidType)
	if t == nil {
		typ = &types.VoidType{}
	} else if !ok {
		panic(fmt.Errorf("invalid IR type for AST void type; expected *types.VoidType, got %T", t))
	}
	// nothing to do.
	return typ, nil
}

// --- [ Function Types ] ------------------------------------------------------

func (gen *generator) translateFuncType(t types.Type, old *ast.FuncType) (types.Type, error) {
	typ, ok := t.(*types.FuncType)
	if t == nil {
		typ = &types.FuncType{}
	} else if !ok {
		panic(fmt.Errorf("invalid IR type for AST function type; expected *types.FuncType, got %T", t))
	}
	// Return type.
	retType, err := gen.translateType(nil, old.RetType())
	if err != nil {
		return nil, errors.WithStack(err)
	}
	typ.RetType = retType
	// Function parameters.
	ps := old.Params()
	for _, p := range ps.Params() {
		param, err := gen.translateType(nil, p.Typ())
		if err != nil {
			return nil, errors.WithStack(err)
		}
		typ.Params = append(typ.Params, param)
	}
	// Variadic.
	typ.Variadic = irVariadic(ps.Variadic())
	return typ, nil
}

// --- [ Integer Types ] -------------------------------------------------------

func (gen *generator) translateIntType(t types.Type, old *ast.IntType) (types.Type, error) {
	typ, ok := t.(*types.IntType)
	if t == nil {
		typ = &types.IntType{}
	} else if !ok {
		panic(fmt.Errorf("invalid IR type for AST integer type; expected *types.IntType, got %T", t))
	}
	// Bit size.
	typ.BitSize = irIntTypeBitSize(old)
	return typ, nil
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

// --- [ Floating-point Types ] ------------------------------------------------

func (gen *generator) translateFloatType(t types.Type, old *ast.FloatType) (types.Type, error) {
	typ, ok := t.(*types.FloatType)
	if t == nil {
		typ = &types.FloatType{}
	} else if !ok {
		panic(fmt.Errorf("invalid IR type for AST floating-point type; expected *types.FloatType, got %T", t))
	}
	// Floating-point kind.
	typ.Kind = irFloatKind(old.FloatKind())
	return typ, nil
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

// --- [ MMX Types ] -----------------------------------------------------------

func (gen *generator) translateMMXType(t types.Type, old *ast.MMXType) (types.Type, error) {
	typ, ok := t.(*types.MMXType)
	if t == nil {
		typ = &types.MMXType{}
	} else if !ok {
		panic(fmt.Errorf("invalid IR type for AST MMX type; expected *types.MMXType, got %T", t))
	}
	// nothing to do.
	return typ, nil
}

// --- [ Pointer Types ] -------------------------------------------------------

func (gen *generator) translatePointerType(t types.Type, old *ast.PointerType) (types.Type, error) {
	typ, ok := t.(*types.PointerType)
	if t == nil {
		typ = &types.PointerType{}
	} else if !ok {
		panic(fmt.Errorf("invalid IR type for AST pointer type; expected *types.PointerType, got %T", t))
	}
	// Element type.
	elemType, err := gen.translateType(nil, old.Elem())
	if err != nil {
		return nil, errors.WithStack(err)
	}
	typ.ElemType = elemType
	// Address space.
	typ.AddrSpace = irAddrSpace(old.AddrSpace())
	return typ, nil
}

// --- [ Vector Types ] --------------------------------------------------------

func (gen *generator) translateVectorType(t types.Type, old *ast.VectorType) (types.Type, error) {
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
	elem, err := gen.translateType(nil, old.Elem())
	if err != nil {
		return nil, errors.WithStack(err)
	}
	typ.ElemType = elem
	return typ, nil
}

// --- [ Label Types ] ---------------------------------------------------------

func (gen *generator) translateLabelType(t types.Type, old *ast.LabelType) (types.Type, error) {
	typ, ok := t.(*types.LabelType)
	if t == nil {
		typ = &types.LabelType{}
	} else if !ok {
		panic(fmt.Errorf("invalid IR type for AST label type; expected *types.LabelType, got %T", t))
	}
	// nothing to do.
	return typ, nil
}

// --- [ Token Types ] ---------------------------------------------------------

func (gen *generator) translateTokenType(t types.Type, old *ast.TokenType) (types.Type, error) {
	typ, ok := t.(*types.TokenType)
	if t == nil {
		typ = &types.TokenType{}
	} else if !ok {
		panic(fmt.Errorf("invalid IR type for AST token type; expected *types.TokenType, got %T", t))
	}
	// nothing to do.
	return typ, nil
}

// --- [ Metadata Types ] ------------------------------------------------------

func (gen *generator) translateMetadataType(t types.Type, old *ast.MetadataType) (types.Type, error) {
	typ, ok := t.(*types.MetadataType)
	if t == nil {
		typ = &types.MetadataType{}
	} else if !ok {
		panic(fmt.Errorf("invalid IR type for AST metadata type; expected *types.MetadataType, got %T", t))
	}
	// nothing to do.
	return typ, nil
}

// --- [ Array Types ] ---------------------------------------------------------

func (gen *generator) translateArrayType(t types.Type, old *ast.ArrayType) (types.Type, error) {
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
	elem, err := gen.translateType(nil, old.Elem())
	if err != nil {
		return nil, errors.WithStack(err)
	}
	typ.ElemType = elem
	return typ, nil
}

// --- [ Structure Types ] -----------------------------------------------------

func (gen *generator) translateOpaqueType(t types.Type, old *ast.OpaqueType) (types.Type, error) {
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
}

func (gen *generator) translateStructType(t types.Type, old *ast.StructType) (types.Type, error) {
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
		field, err := gen.translateType(nil, f)
		if err != nil {
			return nil, errors.WithStack(err)
		}
		typ.Fields = append(typ.Fields, field)
	}
	// struct body now present.
	typ.Opaque = false
	return typ, nil
}

// --- [ Named Types ] ---------------------------------------------------------

func (gen *generator) translateNamedType(t types.Type, old *ast.NamedType) (types.Type, error) {
	// Resolve named type.
	alias := local(old.Name())
	typ, ok := gen.ts[alias]
	if !ok {
		return nil, errors.Errorf("unable to locate type definition of named type %q", enc.Local(alias))
	}
	return typ, nil
}
