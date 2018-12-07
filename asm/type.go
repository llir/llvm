package asm

import (
	"fmt"
	"sort"
	"strconv"
	"strings"

	"github.com/llir/ll/ast"
	asmenum "github.com/llir/llvm/asm/enum"
	"github.com/llir/llvm/internal/enc"
	"github.com/llir/llvm/ir"
	"github.com/llir/llvm/ir/types"
	"github.com/pkg/errors"
)

// resolveTypeDefs resolves the type definitions of the given module.
func (gen *generator) resolveTypeDefs() error {
	// 2. Resolve IR type definitions.
	//
	// 2a. Index type identifiers and create scaffolding IR type definitions
	//     (without bodies).
	if err := gen.createTypeDefs(); err != nil {
		return errors.WithStack(err)
	}
	// 2b. Translate AST type definitions to IR.
	return gen.translateTypeDefs()
}

// === [ Create and index IR ] =================================================

// createTypeDefs indexes type identifiers and creates scaffolding IR type
// definitions (without bodies) of the given module.
//
// post-condition: gen.new.typeDefs maps from type identifier (without '%'
// prefix) to corresponding skeleton IR value.
func (gen *generator) createTypeDefs() error {
	// 2a. Index type identifiers and create scaffolding IR type definitions
	//     (without bodies).
	gen.new.typeDefs = make(map[string]types.Type)
	for typeName, old := range gen.old.typeDefs {
		// track is used to identify self-referential named types.
		track := make(map[string]bool)
		t, err := newType(typeName, old.Typ(), gen.old.typeDefs, track)
		if err != nil {
			return errors.WithStack(err)
		}
		gen.new.typeDefs[typeName] = t
	}
	return nil
}

// newType returns a new IR type (without body) based on the given AST type.
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
func newType(typeName string, old ast.LlvmNode, index map[string]*ast.TypeDef, track map[string]bool) (types.Type, error) {
	switch old := old.(type) {
	case *ast.VoidType:
		return &types.VoidType{TypeName: typeName}, nil
	case *ast.FuncType:
		return &types.FuncType{TypeName: typeName}, nil
	case *ast.IntType:
		return &types.IntType{TypeName: typeName}, nil
	case *ast.FloatType:
		return &types.FloatType{TypeName: typeName}, nil
	case *ast.MMXType:
		return &types.MMXType{TypeName: typeName}, nil
	case *ast.PointerType:
		return &types.PointerType{TypeName: typeName}, nil
	case *ast.VectorType:
		return &types.VectorType{TypeName: typeName}, nil
	case *ast.LabelType:
		return &types.LabelType{TypeName: typeName}, nil
	case *ast.TokenType:
		return &types.TokenType{TypeName: typeName}, nil
	case *ast.MetadataType:
		return &types.MetadataType{TypeName: typeName}, nil
	case *ast.ArrayType:
		return &types.ArrayType{TypeName: typeName}, nil
	case *ast.OpaqueType:
		return &types.StructType{TypeName: typeName}, nil
	case *ast.StructType:
		return &types.StructType{TypeName: typeName}, nil
	case *ast.PackedStructType:
		return &types.StructType{TypeName: typeName}, nil
	case *ast.NamedType:
		if track[typeName] {
			names := make([]string, 0, len(track))
			for name := range track {
				names = append(names, enc.Local(name))
			}
			sort.Strings(names)
			return nil, errors.Errorf("invalid named type; self-referential with type name(s) %s", strings.Join(names, ", "))
		}
		track[typeName] = true
		newIdent := localIdent(old.Name())
		newName := getTypeName(newIdent)
		newTyp := index[newName].Typ()
		return newType(newName, newTyp, index, track)
	default:
		panic(fmt.Errorf("support for type %T not yet implemented", old))
	}
}

// === [ Translate AST to IR ] =================================================

// translateTypeDefs translates the AST type definitions of the given module to
// IR.
func (gen *generator) translateTypeDefs() error {
	// 2b. Translate AST type definitions to IR.
	for typeName, old := range gen.old.typeDefs {
		t := gen.new.typeDefs[typeName]
		if _, err := gen.irTypeDef(t, old.Typ()); err != nil {
			return errors.WithStack(err)
		}
	}
	return nil
}

// irTypeDef translates the AST type into an equivalent IR type. A new IR type
// correspoding to the AST type is created if t is nil, otherwise the body of t
// is populated. Named types are resolved through gen.new.typeDefs.
func (gen *generator) irTypeDef(t types.Type, old ast.LlvmNode) (types.Type, error) {
	switch old := old.(type) {
	case *ast.VoidType:
		return gen.irVoidType(t, old)
	case *ast.FuncType:
		return gen.irFuncType(t, old)
	case *ast.IntType:
		return gen.irIntType(t, old)
	case *ast.FloatType:
		return gen.irFloatType(t, old)
	case *ast.MMXType:
		return gen.irMMXType(t, old)
	case *ast.PointerType:
		return gen.irPointerType(t, old)
	case *ast.VectorType:
		return gen.irVectorType(t, old)
	case *ast.LabelType:
		return gen.irLabelType(t, old)
	case *ast.TokenType:
		return gen.irTokenType(t, old)
	case *ast.MetadataType:
		return gen.irMetadataType(t, old)
	case *ast.ArrayType:
		return gen.irArrayType(t, old)
	case *ast.OpaqueType:
		return gen.irOpaqueType(t, old)
	case *ast.StructType:
		return gen.irStructType(t, old)
	case *ast.PackedStructType:
		return gen.irPackedStructType(t, old)
	case *ast.NamedType:
		return gen.irNamedType(t, old)
	default:
		panic(fmt.Errorf("support for type %T not yet implemented", old))
	}
}

// --- [ Void types ] ----------------------------------------------------------

// irVoidType translates the AST void type into an equivalent IR type. A new IR
// type correspoding to the AST type is created if t is nil, otherwise the body
// of t is populated.
func (gen *generator) irVoidType(t types.Type, old *ast.VoidType) (types.Type, error) {
	typ, ok := t.(*types.VoidType)
	if t == nil {
		typ = &types.VoidType{}
	} else if !ok {
		panic(fmt.Errorf("invalid IR type for AST void type; expected *types.VoidType, got %T", t))
	}
	// nothing to do.
	return typ, nil
}

// --- [ Function types ] ------------------------------------------------------

// irFuncType translates the AST function type into an equivalent IR type. A new
// IR type correspoding to the AST type is created if t is nil, otherwise the
// body of t is populated.
func (gen *generator) irFuncType(t types.Type, old *ast.FuncType) (types.Type, error) {
	typ, ok := t.(*types.FuncType)
	if t == nil {
		typ = &types.FuncType{}
	} else if !ok {
		panic(fmt.Errorf("invalid IR type for AST function type; expected *types.FuncType, got %T", t))
	}
	// Return type.
	retType, err := gen.irType(old.RetType())
	if err != nil {
		return nil, errors.WithStack(err)
	}
	typ.RetType = retType
	// Function parameters.
	ps := old.Params()
	if oldParams := ps.Params(); len(oldParams) > 0 {
		typ.Params = make([]types.Type, len(oldParams))
		for i, oldParam := range oldParams {
			param, err := gen.irType(oldParam.Typ())
			if err != nil {
				return nil, errors.WithStack(err)
			}
			typ.Params[i] = param
		}
	}
	// Variadic.
	_, typ.Variadic = ps.Variadic()
	return typ, nil
}

// --- [ Integer types ] -------------------------------------------------------

// irIntType translates the AST integer type into an equivalent IR type. A new
// IR type correspoding to the AST type is created if t is nil, otherwise the
// body of t is populated.
func (gen *generator) irIntType(t types.Type, old *ast.IntType) (types.Type, error) {
	typ, ok := t.(*types.IntType)
	if t == nil {
		typ = &types.IntType{}
	} else if !ok {
		panic(fmt.Errorf("invalid IR type for AST integer type; expected *types.IntType, got %T", t))
	}
	// Bit size.
	typ.BitSize = irBitSize(old)
	return typ, nil
}

// irBitSize returns the bit size of the given AST integer type.
func irBitSize(n *ast.IntType) uint64 {
	text := n.Text()
	const prefix = "i"
	if !strings.HasPrefix(text, prefix) {
		panic(fmt.Errorf("invalid integer type %q; missing '%s' prefix", text, prefix))
	}
	text = text[len(prefix):]
	x, err := strconv.ParseUint(text, 10, 64)
	if err != nil {
		panic(fmt.Errorf("unable to parse bit size %q; %v", text, err))
	}
	return x
}

// --- [ Floating-point types ] ------------------------------------------------

// irFloatType translates the AST floating-point type into an equivalent IR
// type. A new IR type correspoding to the AST type is created if t is nil,
// otherwise the body of t is populated.
func (gen *generator) irFloatType(t types.Type, old *ast.FloatType) (types.Type, error) {
	typ, ok := t.(*types.FloatType)
	if t == nil {
		typ = &types.FloatType{}
	} else if !ok {
		panic(fmt.Errorf("invalid IR type for AST floating-point type; expected *types.FloatType, got %T", t))
	}
	// Floating-point kind.
	typ.Kind = asmenum.FloatKindFromString(old.FloatKind().Text())
	return typ, nil
}

// --- [ MMX types ] -----------------------------------------------------------

// irMMXType translates the AST MMX type into an equivalent IR type. A new IR
// type correspoding to the AST type is created if t is nil, otherwise the body
// of t is populated.
func (gen *generator) irMMXType(t types.Type, old *ast.MMXType) (types.Type, error) {
	typ, ok := t.(*types.MMXType)
	if t == nil {
		typ = &types.MMXType{}
	} else if !ok {
		panic(fmt.Errorf("invalid IR type for AST MMX type; expected *types.MMXType, got %T", t))
	}
	// nothing to do.
	return typ, nil
}

// --- [ Pointer types ] -------------------------------------------------------

// irPointerType translates the AST pointer type into an equivalent IR type. A
// new IR type correspoding to the AST type is created if t is nil, otherwise
// the body of t is populated.
func (gen *generator) irPointerType(t types.Type, old *ast.PointerType) (types.Type, error) {
	typ, ok := t.(*types.PointerType)
	if t == nil {
		typ = &types.PointerType{}
	} else if !ok {
		panic(fmt.Errorf("invalid IR type for AST pointer type; expected *types.PointerType, got %T", t))
	}
	// Element type.
	elemType, err := gen.irType(old.Elem())
	if err != nil {
		return nil, errors.WithStack(err)
	}
	typ.ElemType = elemType
	// Address space.
	if n, ok := old.AddrSpace(); ok {
		typ.AddrSpace = irAddrSpace(n)
	}
	return typ, nil
}

// --- [ Vector types ] --------------------------------------------------------

// irVectorType translates the AST vector type into an equivalent IR type. A new
// IR type correspoding to the AST type is created if t is nil, otherwise the
// body of t is populated.
func (gen *generator) irVectorType(t types.Type, old *ast.VectorType) (types.Type, error) {
	typ, ok := t.(*types.VectorType)
	if t == nil {
		typ = &types.VectorType{}
	} else if !ok {
		panic(fmt.Errorf("invalid IR type for AST vector type; expected *types.VectorType, got %T", t))
	}
	// Vector length.
	typ.Len = uintLit(old.Len())
	// Element type.
	elem, err := gen.irType(old.Elem())
	if err != nil {
		return nil, errors.WithStack(err)
	}
	typ.ElemType = elem
	return typ, nil
}

// --- [ Label types ] ---------------------------------------------------------

// irLabelType translates the AST label type into an equivalent IR type. A new
// IR type correspoding to the AST type is created if t is nil, otherwise the
// body of t is populated.
func (gen *generator) irLabelType(t types.Type, old *ast.LabelType) (types.Type, error) {
	typ, ok := t.(*types.LabelType)
	if t == nil {
		typ = &types.LabelType{}
	} else if !ok {
		panic(fmt.Errorf("invalid IR type for AST label type; expected *types.LabelType, got %T", t))
	}
	// nothing to do.
	return typ, nil
}

// --- [ Token types ] ---------------------------------------------------------

// irTokenType translates the AST token type into an equivalent IR type. A new
// IR type correspoding to the AST type is created if t is nil, otherwise the
// body of t is populated.
func (gen *generator) irTokenType(t types.Type, old *ast.TokenType) (types.Type, error) {
	typ, ok := t.(*types.TokenType)
	if t == nil {
		typ = &types.TokenType{}
	} else if !ok {
		panic(fmt.Errorf("invalid IR type for AST token type; expected *types.TokenType, got %T", t))
	}
	// nothing to do.
	return typ, nil
}

// --- [ Metadata types ] ------------------------------------------------------

// irMetadataType translates the AST metadata type into an equivalent IR type. A
// new IR type correspoding to the AST type is created if t is nil, otherwise
// the body of t is populated.
func (gen *generator) irMetadataType(t types.Type, old *ast.MetadataType) (types.Type, error) {
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

// irArrayType translates the AST array type into an equivalent IR type. A new
// IR type correspoding to the AST type is created if t is nil, otherwise the
// body of t is populated.
func (gen *generator) irArrayType(t types.Type, old *ast.ArrayType) (types.Type, error) {
	typ, ok := t.(*types.ArrayType)
	if t == nil {
		typ = &types.ArrayType{}
	} else if !ok {
		panic(fmt.Errorf("invalid IR type for AST array type; expected *types.ArrayType, got %T", t))
	}
	// Array length.
	typ.Len = uintLit(old.Len())
	// Element type.
	elem, err := gen.irType(old.Elem())
	if err != nil {
		return nil, errors.WithStack(err)
	}
	typ.ElemType = elem
	return typ, nil
}

// --- [ Structure Types ] -----------------------------------------------------

// irOpaqueType translates the AST opaque struct type into an equivalent IR
// type.
func (gen *generator) irOpaqueType(t types.Type, old *ast.OpaqueType) (types.Type, error) {
	typ, ok := t.(*types.StructType)
	if t == nil {
		// Panic as this case should not be reachable by the grammar.
		panic("invalid use of opaque type; only allowed in type definitions")
	} else if !ok {
		panic(fmt.Errorf("invalid IR type for AST opaque type; expected *types.StructType, got %T", t))
	}
	// Opaque.
	typ.Opaque = true
	return typ, nil
}

// irStructType translates the AST struct type into an equivalent IR type. A new
// IR type correspoding to the AST type is created if t is nil, otherwise the
// body of t is populated.
func (gen *generator) irStructType(t types.Type, old *ast.StructType) (types.Type, error) {
	typ, ok := t.(*types.StructType)
	if t == nil {
		typ = &types.StructType{}
	} else if !ok {
		panic(fmt.Errorf("invalid IR type for AST struct type; expected *types.StructType, got %T", t))
	}
	// Packed (not present).
	// Fields.
	if oldFields := old.Fields(); len(oldFields) > 0 {
		typ.Fields = make([]types.Type, len(oldFields))
		for i, oldField := range oldFields {
			field, err := gen.irType(oldField)
			if err != nil {
				return nil, errors.WithStack(err)
			}
			typ.Fields[i] = field
		}
	}
	// struct body now present.
	typ.Opaque = false
	return typ, nil
}

// irPackedStructType translates the AST packed struct type into an equivalent
// IR type. A new IR type correspoding to the AST type is created if t is nil,
// otherwise the body of t is populated.
func (gen *generator) irPackedStructType(t types.Type, old *ast.PackedStructType) (types.Type, error) {
	typ, ok := t.(*types.StructType)
	if t == nil {
		typ = &types.StructType{}
	} else if !ok {
		panic(fmt.Errorf("invalid IR type for AST struct type; expected *types.StructType, got %T", t))
	}
	// Packed.
	typ.Packed = true
	// Fields.
	if oldFields := old.Fields(); len(oldFields) > 0 {
		typ.Fields = make([]types.Type, len(oldFields))
		for i, oldField := range oldFields {
			field, err := gen.irType(oldField)
			if err != nil {
				return nil, errors.WithStack(err)
			}
			typ.Fields[i] = field
		}
	}
	// struct body now present.
	typ.Opaque = false
	return typ, nil
}

// --- [ Named Types ] ---------------------------------------------------------

// irNamedType translates the AST named type into an equivalent IR type.
func (gen *generator) irNamedType(t types.Type, old *ast.NamedType) (types.Type, error) {
	// TODO: make use of t?
	// Resolve named type.
	ident := localIdent(old.Name())
	name := getTypeName(ident)
	typ, ok := gen.new.typeDefs[name]
	if !ok {
		return nil, errors.Errorf("unable to locate type definition of named type %q", enc.Local(name))
	}
	return typ, nil
}

// ### [ Helpers ] #############################################################

// irType returns the IR type corresponding to the given AST type.
func (gen *generator) irType(old ast.LlvmNode) (types.Type, error) {
	return gen.irTypeDef(nil, old)
}

// getTypeName returns the identifier (without '%' prefix) of the given type
// identifier.
func getTypeName(ident ir.LocalIdent) string {
	if ident.IsUnnamed() {
		return strconv.FormatInt(ident.LocalID, 10)
	}
	if x, err := strconv.ParseInt(ident.LocalName, 10, 64); err == nil {
		// Print LocalName with quotes if it is a number; e.g. %"42".
		return fmt.Sprintf(`"%d"`, x)
	}
	return ident.LocalName
}
