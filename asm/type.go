package asm

import (
	"fmt"
	"sort"
	"strconv"
	"strings"

	"github.com/llir/ll/ast"
	"github.com/llir/llvm/internal/enc"
	"github.com/llir/llvm/ir"
	"github.com/llir/llvm/ir/types"
	"github.com/pkg/errors"
)

// resolveTypeDefs resolves the type definitions of the given module. The
// returned value maps from type name (without '%' prefix) to the underlying
// type.
func (gen *generator) resolveTypeDefs() error {
	// Create corresponding named IR types (without bodies).
	gen.new.typeDefs = make(map[string]types.Type)
	for typeName, old := range gen.old.typeDefs {
		// track is used to identify self-referential named types.
		track := make(map[string]bool)
		t, err := newIRType(typeName, old.Typ(), gen.old.typeDefs, track)
		if err != nil {
			return errors.WithStack(err)
		}
		gen.new.typeDefs[typeName] = t
	}

	// Translate type defintions (including bodies).
	for alias, old := range gen.old.typeDefs {
		t := gen.new.typeDefs[alias]
		_, err := gen.astToIRTypeDef(t, old.Typ())
		if err != nil {
			return errors.WithStack(err)
		}
	}
	return nil
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
func newIRType(typeName string, old ast.LlvmNode, index map[string]*ast.TypeDef, track map[string]bool) (types.Type, error) {
	switch old := old.(type) {
	case *ast.OpaqueType:
		return &types.StructType{TypeName: typeName}, nil
	case *ast.ArrayType:
		return &types.ArrayType{TypeName: typeName}, nil
	case *ast.FloatType:
		return &types.FloatType{TypeName: typeName}, nil
	case *ast.FuncType:
		return &types.FuncType{TypeName: typeName}, nil
	case *ast.IntType:
		return &types.IntType{TypeName: typeName}, nil
	case *ast.LabelType:
		return &types.LabelType{TypeName: typeName}, nil
	case *ast.MMXType:
		return &types.MMXType{TypeName: typeName}, nil
	case *ast.MetadataType:
		return &types.MetadataType{TypeName: typeName}, nil
	case *ast.NamedType:
		if track[typeName] {
			var names []string
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
		return newIRType(newName, newTyp, index, track)
	case *ast.PointerType:
		return &types.PointerType{TypeName: typeName}, nil
	case *ast.StructType:
		return &types.StructType{TypeName: typeName}, nil
	case *ast.PackedStructType:
		return &types.StructType{TypeName: typeName}, nil
	case *ast.TokenType:
		return &types.TokenType{TypeName: typeName}, nil
	case *ast.VectorType:
		return &types.VectorType{TypeName: typeName}, nil
	case *ast.VoidType:
		return &types.VoidType{TypeName: typeName}, nil
	default:
		panic(fmt.Errorf("support for type %T not yet implemented", old))
	}
}

// === [ Types ] ===============================================================

// astToIRTypeDef translates the AST type into an equivalent IR type. A new IR
// type correspoding to the AST type is created if t is nil, otherwise the body
// of t is populated. Named types are resolved through ts.
func (gen *generator) astToIRTypeDef(t types.Type, old ast.LlvmNode) (types.Type, error) {
	switch old := old.(type) {
	case *ast.OpaqueType:
		return gen.astToIROpaqueType(t, old)
	case *ast.ArrayType:
		return gen.astToIRArrayType(t, old)
	case *ast.FloatType:
		return gen.astToIRFloatType(t, old)
	case *ast.FuncType:
		return gen.astToIRFuncType(t, old)
	case *ast.IntType:
		return gen.astToIRIntType(t, old)
	case *ast.LabelType:
		return gen.astToIRLabelType(t, old)
	case *ast.MMXType:
		return gen.astToIRMMXType(t, old)
	case *ast.MetadataType:
		return gen.astToIRMetadataType(t, old)
	case *ast.NamedType:
		return gen.astToIRNamedType(t, old)
	case *ast.PointerType:
		return gen.astToIRPointerType(t, old)
	case *ast.StructType:
		return gen.astToIRStructType(t, old)
	case *ast.PackedStructType:
		return gen.astToIRPackedStructType(t, old)
	case *ast.TokenType:
		return gen.astToIRTokenType(t, old)
	case *ast.VectorType:
		return gen.astToIRVectorType(t, old)
	case *ast.VoidType:
		return gen.astToIRVoidType(t, old)
	default:
		panic(fmt.Errorf("support for type %T not yet implemented", old))
	}
}

// --- [ Void Types ] ----------------------------------------------------------

func (gen *generator) astToIRVoidType(t types.Type, old *ast.VoidType) (types.Type, error) {
	typ, ok := t.(*types.VoidType)
	if t == nil {
		typ = &types.VoidType{}
	} else if !ok {
		// NOTE: Panic instead of returning error as this case should not be
		// possible, and would indicate a bug in the implementation.
		panic(fmt.Errorf("invalid IR type for AST void type; expected *types.VoidType, got %T", t))
	}
	// nothing to do.
	return typ, nil
}

// --- [ Function Types ] ------------------------------------------------------

func (gen *generator) astToIRFuncType(t types.Type, old *ast.FuncType) (types.Type, error) {
	typ, ok := t.(*types.FuncType)
	if t == nil {
		typ = &types.FuncType{}
	} else if !ok {
		// NOTE: Panic instead of returning error as this case should not be
		// possible, and would indicate a bug in the implementation.
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
	for _, p := range ps.Params() {
		param, err := gen.irType(p.Typ())
		if err != nil {
			return nil, errors.WithStack(err)
		}
		typ.Params = append(typ.Params, param)
	}
	// Variadic.
	_, variadic := ps.Variadic()
	typ.Variadic = variadic
	return typ, nil
}

// --- [ Integer Types ] -------------------------------------------------------

func (gen *generator) astToIRIntType(t types.Type, old *ast.IntType) (types.Type, error) {
	typ, ok := t.(*types.IntType)
	if t == nil {
		typ = &types.IntType{}
	} else if !ok {
		// NOTE: Panic instead of returning error as this case should not be
		// possible, and would indicate a bug in the implementation.
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

func (gen *generator) astToIRFloatType(t types.Type, old *ast.FloatType) (types.Type, error) {
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

func (gen *generator) astToIRMMXType(t types.Type, old *ast.MMXType) (types.Type, error) {
	typ, ok := t.(*types.MMXType)
	if t == nil {
		typ = &types.MMXType{}
	} else if !ok {
		// NOTE: Panic instead of returning error as this case should not be
		// possible, and would indicate a bug in the implementation.
		panic(fmt.Errorf("invalid IR type for AST MMX type; expected *types.MMXType, got %T", t))
	}
	// nothing to do.
	return typ, nil
}

// --- [ Pointer Types ] -------------------------------------------------------

func (gen *generator) astToIRPointerType(t types.Type, old *ast.PointerType) (types.Type, error) {
	typ, ok := t.(*types.PointerType)
	if t == nil {
		typ = &types.PointerType{}
	} else if !ok {
		// NOTE: Panic instead of returning error as this case should not be
		// possible, and would indicate a bug in the implementation.
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

// --- [ Vector Types ] --------------------------------------------------------

func (gen *generator) astToIRVectorType(t types.Type, old *ast.VectorType) (types.Type, error) {
	typ, ok := t.(*types.VectorType)
	if t == nil {
		typ = &types.VectorType{}
	} else if !ok {
		// NOTE: Panic instead of returning error as this case should not be
		// possible, and would indicate a bug in the implementation.
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

// --- [ Label Types ] ---------------------------------------------------------

func (gen *generator) astToIRLabelType(t types.Type, old *ast.LabelType) (types.Type, error) {
	typ, ok := t.(*types.LabelType)
	if t == nil {
		typ = &types.LabelType{}
	} else if !ok {
		// NOTE: Panic instead of returning error as this case should not be
		// possible, and would indicate a bug in the implementation.
		panic(fmt.Errorf("invalid IR type for AST label type; expected *types.LabelType, got %T", t))
	}
	// nothing to do.
	return typ, nil
}

// --- [ Token Types ] ---------------------------------------------------------

func (gen *generator) astToIRTokenType(t types.Type, old *ast.TokenType) (types.Type, error) {
	typ, ok := t.(*types.TokenType)
	if t == nil {
		typ = &types.TokenType{}
	} else if !ok {
		// NOTE: Panic instead of returning error as this case should not be
		// possible, and would indicate a bug in the implementation.
		panic(fmt.Errorf("invalid IR type for AST token type; expected *types.TokenType, got %T", t))
	}
	// nothing to do.
	return typ, nil
}

// --- [ Metadata Types ] ------------------------------------------------------

func (gen *generator) astToIRMetadataType(t types.Type, old *ast.MetadataType) (types.Type, error) {
	typ, ok := t.(*types.MetadataType)
	if t == nil {
		typ = &types.MetadataType{}
	} else if !ok {
		// NOTE: Panic instead of returning error as this case should not be
		// possible, and would indicate a bug in the implementation.
		panic(fmt.Errorf("invalid IR type for AST metadata type; expected *types.MetadataType, got %T", t))
	}
	// nothing to do.
	return typ, nil
}

// --- [ Array Types ] ---------------------------------------------------------

func (gen *generator) astToIRArrayType(t types.Type, old *ast.ArrayType) (types.Type, error) {
	typ, ok := t.(*types.ArrayType)
	if t == nil {
		typ = &types.ArrayType{}
	} else if !ok {
		// NOTE: Panic instead of returning error as this case should not be
		// possible, and would indicate a bug in the implementation.
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

func (gen *generator) astToIROpaqueType(t types.Type, old *ast.OpaqueType) (types.Type, error) {
	typ, ok := t.(*types.StructType)
	if t == nil {
		// NOTE: Panic instead of returning error as this case should not be
		// possible given the grammar.
		panic("invalid use of opaque type; only allowed in type definitions")
	} else if !ok {
		// NOTE: Panic instead of returning error as this case should not be
		// possible, and would indicate a bug in the implementation.
		panic(fmt.Errorf("invalid IR type for AST opaque type; expected *types.StructType, got %T", t))
	}
	// Opaque.
	typ.Opaque = true
	return typ, nil
}

func (gen *generator) astToIRStructType(t types.Type, old *ast.StructType) (types.Type, error) {
	typ, ok := t.(*types.StructType)
	if t == nil {
		typ = &types.StructType{}
	} else if !ok {
		// NOTE: Panic instead of returning error as this case should not be
		// possible, and would indicate a bug in the implementation.
		panic(fmt.Errorf("invalid IR type for AST struct type; expected *types.StructType, got %T", t))
	}
	// Packed.
	// Fields.
	for _, f := range old.Fields() {
		field, err := gen.irType(f)
		if err != nil {
			return nil, errors.WithStack(err)
		}
		typ.Fields = append(typ.Fields, field)
	}
	// struct body now present.
	typ.Opaque = false
	return typ, nil
}

func (gen *generator) astToIRPackedStructType(t types.Type, old *ast.PackedStructType) (types.Type, error) {
	typ, ok := t.(*types.StructType)
	if t == nil {
		typ = &types.StructType{}
	} else if !ok {
		// NOTE: Panic instead of returning error as this case should not be
		// possible, and would indicate a bug in the implementation.
		panic(fmt.Errorf("invalid IR type for AST struct type; expected *types.StructType, got %T", t))
	}
	// Packed.
	typ.Packed = true
	// Fields.
	for _, f := range old.Fields() {
		field, err := gen.irType(f)
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

func (gen *generator) astToIRNamedType(t types.Type, old *ast.NamedType) (types.Type, error) {
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

// TODO: rename irType to astToIRType?

// irType returns the IR type corresponding to the given AST type.
func (gen *generator) irType(old ast.LlvmNode) (types.Type, error) {
	return gen.astToIRTypeDef(nil, old)
}

// getTypeName returns the identifier (without '%' prefix) of the given global
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
