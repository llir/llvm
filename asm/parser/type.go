package parser

import (
	"strconv"
	"strings"

	"github.com/llir/llvm/asm"
	"github.com/llir/llvm/asm/token"
	"github.com/llir/llvm/types"
	"github.com/mewkiz/pkg/errutil"
)

// parseTypeDef parses a type definition or a type alias. The next token is
// either a LocalID or a LocalVar.
//
//    TypeDef  = TypeName "=" "type" ( StructType | AliasType ) .
//    TypeName = Local .
func (p *parser) parseTypeDef() error {
	name := p.next()
	if !p.accept(token.Equal) {
		return errutil.Newf(`expected "=" after type name %q, got %q token`, asm.EncLocal(name.Val), p.next())
	}
	if !p.accept(token.KwType) {
		return errutil.Newf(`expected "type" after type assignment, got %q token`, p.next())
	}
	typ, err := p.parseType()
	if err != nil {
		return errutil.Err(err)
	}
	switch typ := typ.(type) {
	case *types.Struct:
		// Identified structure.
		t, err := p.tctx.Struct(name.Val)
		if err != nil {
			return errutil.Err(err)
		}
		t.Struct = typ
		p.m.Types = append(p.m.Types, t)
	default:
		// Type alias.
		if err := p.tctx.SetAlias(name.Val, typ); err != nil {
			return errutil.Err(err)
		}
	}
	return nil
}

// parseType parses a type.
//
//    Type      = VoidType | IntType | FloatType | MMXType | LabelType |
//                MetadataType | FuncType | PointerType | VectorType |
//                ArrayType | StructType .
//    AliasType = VoidType | IntType | FloatType | MMXType | LabelType |
//                MetadataType | FuncType | PointerType | VectorType |
//                ArrayType .
//
//    VoidType        = "void" .
//    IntType         = "i" int_lit .
//    FloatType       = "half" | "float" | "double" | "fp128" | x86_fp80 |
//                      "ppc_fp128" .
//    MMXType         = "x86_mmx" .
//    LabelType       = "label" .
//    MetadataType    = "metadata" .
//    FuncType        = FuncResultType "(" ( FuncParamType { "," FuncParamType } ] [ "," "..." ]) | [ "..." ] ")" .
//    FuncResultType  = VoidType | IntType | FloatType | MMXType | PointerType |
//                      VectorType | ArrayType | StructType .
//    FuncParamType   = IntType | FloatType | MMXType | LabelType |
//                      MetadataType | PointerType | VectorType | ArrayType |
//                      StructType .
//    PointerType     = (IntType | FloatType | MMXType | FuncType |
//                      PointerType | VectorType | ArrayType | StructType) "*" .
//
//    IntsType   = ( IntType | IntVectorType ) .
//    FloatsType = ( FloatType | FloatVectorType ) .
func (p *parser) parseType() (typ types.Type, err error) {
	switch tok := p.next(); tok.Kind {
	// Basic type; e.g.
	//    i32
	//    float
	case token.Type:
		typ, err = basicTypeFromString(tok.Val)
		if err != nil {
			return nil, errutil.Err(err)
		}

	case token.Less:
		if p.accept(token.Lbrace) {
			// Packed array type; e.g.
			//    <{i32, i8}>
			typ, err = p.parseStructType(true)
			if err != nil {
				return nil, errutil.Err(err)
			}
		} else {
			// Vector type; e.g.
			//    <2 x i32>
			typ, err = p.parseVectorType()
			if err != nil {
				return nil, errutil.Err(err)
			}
		}

	// Array type; e.g.
	//    [2 x float]
	case token.Lbrack:
		typ, err = p.parseArrayType()
		if err != nil {
			return nil, errutil.Err(err)
		}

	// Structure type; e.g.
	//    {i32, float}
	case token.Lbrace:
		typ, err = p.parseStructType(false)
		if err != nil {
			return nil, errutil.Err(err)
		}

	// Identified structure or type alias; e.g.
	//    %42
	//    %foo
	case token.LocalID, token.LocalVar:
		name := tok.Val
		if alias, ok := p.tctx.Alias(name); ok {
			return alias, nil
		}
		typ, err = p.tctx.Struct(name)
		if err != nil {
			return nil, errutil.Err(err)
		}

	default:
		return nil, errutil.Newf("expected type; got %q token", tok)
	}

	for {
		// Pointer type; e.g.
		//    i32*
		//    [2 x float]*
		//    i8****
		for p.accept(token.Star) {
			elem := typ
			typ, err = types.NewPointer(elem)
			if err != nil {
				return nil, errutil.Err(err)
			}
		}

		// Function type; e.g.
		//    i32 (i8*, ...)
		//    [2 x float]* (i32)
		//    i32 (i32)* (i32)
		if p.accept(token.Lparen) {
			result := typ
			typ, err = p.parseFuncType(result)
			if err != nil {
				return nil, errutil.Err(err)
			}
		} else {
			break
		}
	}

	if typ == nil {
		return nil, errutil.New("expected type")
	}
	return typ, nil
}

// parseVectorType parses a vector type. A "<" token has already been consumed.
//
// Syntax:
//    VectorType      = IntVectorType | FloatVectorType |
//                      "<" VectorLen "x" PointerType ">" .
//    IntVectorType   = "<" VectorLen "x" IntType ">" .
//    FloatVectorType = "<" VectorLen "x" FloatType ">" .
//    VectorLen       = int_lit .
//
// Example:
//    <2 x i32>
func (p *parser) parseVectorType() (*types.Vector, error) {
	// Vector length.
	s, ok := p.try(token.Int)
	if !ok {
		return nil, errutil.New("expected vector length")
	}
	n, err := strconv.Atoi(s)
	if err != nil {
		return nil, errutil.Newf("invalid vector length (%v); %v", s, err)
	}
	if n < 1 {
		return nil, errutil.Newf("invalid vector length (%d); expected >= 1", n)
	}

	// "x" token.
	if !p.accept(token.KwX) {
		return nil, errutil.New("expected 'x' after vector length")
	}

	// Element type.
	elem, err := p.parseType()
	if err != nil {
		return nil, errutil.Err(err)
	}

	// End of vector.
	if !p.accept(token.Greater) {
		return nil, errutil.New("expected '>' at end of vector")
	}

	return types.NewVector(elem, n)
}

// parseArrayType parses a array type. A "[" token has already been consumed.
//
// Syntax:
//    ArrayType = "[" ArrayLen "x" ElemType "]" .
//    ElemType  = IntType | FloatType | MMXType | PointerType | VectorType |
//                ArrayType | StructType .
//    ArrayLen  = int_lit .
//
// Example:
//    [5 x float]
func (p *parser) parseArrayType() (*types.Array, error) {
	// Array length.
	s, ok := p.try(token.Int)
	if !ok {
		return nil, errutil.New("expected array length")
	}
	n, err := strconv.Atoi(s)
	if err != nil {
		return nil, errutil.Newf("invalid array length (%v); %v", s, err)
	}
	if n < 0 {
		return nil, errutil.Newf("invalid array length (%d); expected >= 0", n)
	}

	// "x" token.
	if !p.accept(token.KwX) {
		return nil, errutil.New("expected 'x' after array length")
	}

	// Element type.
	elem, err := p.parseType()
	if err != nil {
		return nil, errutil.Err(err)
	}

	// End of array.
	if !p.accept(token.Rbrack) {
		return nil, errutil.New("expected ']' at end of array")
	}

	return types.NewArray(elem, n)
}

// parseStructType parses a structure type. The structure is 1 byte aligned if
// packed is true. A "{" token has already been consumed, unless the structure
// is packed in which case a "<" and a "{" token have already been consumed.
//
// Syntax:
//    StructType = "{" [ FieldType { "," FieldType } ] "}" | "<" "{" [ FieldType { "," FieldType } ] "}" ">" .
//    FieldType  = IntType | FloatType | MMXType | PointerType | VectorType |
//                 ArrayType | StructType .
//
// Example:
//    {i32, float}
//    <{i32, i8}>
func (p *parser) parseStructType(packed bool) (*types.Struct, error) {
	// Early return for empty structure.
	if p.accept(token.Rbrace) {
		if packed && !p.accept(token.Greater) {
			return nil, errutil.New("expected '>' at end of packed structure")
		}
		return types.NewStruct(nil, packed)
	}

	// Structure fields
	var fields []types.Type
	for i := 0; ; i++ {
		if i > 0 && !p.accept(token.Comma) {
			break
		}
		if field, ok := p.tryType(); ok {
			fields = append(fields, field)
		}
	}

	// End of structure.
	if !p.accept(token.Rbrace) {
		return nil, errutil.New("expected '}' at end of structure")
	}
	if packed && !p.accept(token.Greater) {
		return nil, errutil.New("expected '>' at end of packed structure")
	}

	return types.NewStruct(fields, packed)
}

// parseFuncType parses a function type. A result type, an optional function
// name and a "(" token has already been consumed.
//
//    FuncType   = FuncResult "(" FuncParams ")" .
//    FuncResult = Type .
//    FuncParams = [ FuncParam { "," FuncParam } [ "," "..." ] ] | "..." .
//    FuncParam  = Type .
func (p *parser) parseFuncType(result types.Type) (typ *types.Func, err error) {
	// Early return for empty parameter list.
	if p.accept(token.Rparen) {
		return types.NewFunc(result, nil, false)
	}

	// Function parameters.
	var params []types.Type
	variadic := false
	for i := 0; ; i++ {
		if i > 0 && !p.accept(token.Comma) {
			break
		}
		if param, ok := p.tryType(); ok {
			params = append(params, param)
			if _, ok := p.tryLocal(); ok {
				return nil, errutil.New("argument name invalid in function type")
			}
		} else if p.accept(token.Ellipsis) {
			variadic = true
			break
		} else {
			return nil, errutil.New("expected type")
		}
	}
	if !p.accept(token.Rparen) {
		return nil, errutil.New("expected ')' at end of argument list")
	}

	return types.NewFunc(result, params, variadic)
}

// basicTypeFromString returns the basic type corresponding to s.
func basicTypeFromString(s string) (types.Type, error) {
	switch s {
	case "void":
		return types.NewVoid(), nil
	case "half":
		return types.NewFloat(types.Float16)
	case "float":
		return types.NewFloat(types.Float32)
	case "double":
		return types.NewFloat(types.Float64)
	case "fp128":
		return types.NewFloat(types.Float128)
	case "x86_fp80":
		return types.NewFloat(types.Float80_x86)
	case "ppc_fp128":
		return types.NewFloat(types.Float128_PPC)
	case "x86_mmx":
		return types.NewMMX(), nil
	case "label":
		return types.NewLabel(), nil
	case "metadata":
		return types.NewMetadata(), nil
	}

	// Integer type (e.g. i32).
	if !strings.HasPrefix(s, "i") {
		return nil, errutil.Newf("unknown basic type %q", s)
	}
	n, err := strconv.Atoi(s[1:]) // skip leading "i".
	if err != nil {
		return nil, errutil.Newf("unknown basic type %q", s)
	}
	return types.NewInt(n)
}
