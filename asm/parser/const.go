package parser

import (
	"strconv"

	"github.com/llir/llvm/asm/token"
	"github.com/llir/llvm/constant"
	"github.com/llir/llvm/types"
	"github.com/mewkiz/pkg/errutil"
)

// TODO: Add EBNF for ConstExpr.

// parseConst parses a constant value.
//
// Syntax:
//    Const = IntConst | FloatConst | PointerConst | VectorConst | ArrayConst |
//            StructConst | ConstExpr .
//
// References:
//    http://llvm.org/docs/LangRef.html#constants
func (p *parser) parseConst() (constant.Constant, error) {
	typ, err := p.parseType()
	if err != nil {
		return nil, errutil.Err(err)
	}
	switch typ := typ.(type) {
	case *types.Int:
		return p.parseIntConst(typ)
	case *types.Float:
		return p.parseFloatConst(typ)
	case *types.Pointer:
		return p.parsePointerConst(typ)
	case *types.Vector:
		return p.parseVectorConst(typ)
	case *types.Array:
		return p.parseArrayConst(typ)
	case *types.Struct:
		return p.parseStructConst(typ)
	default:
		return nil, errutil.Newf("support for constants of type %q not yet implemented", typ)
	}
}

// parseIntConst parses an integer constant. An integer type as already been
// consumed.
//
// Syntax:
//    IntConst = IntType IntValue .
//    IntValue = int_lit .
//
// Examples:
//    i32 42
//
// References:
//    http://llvm.org/docs/LangRef.html#simple-constants
func (p *parser) parseIntConst(typ *types.Int) (*constant.Int, error) {
	s, err := p.expect(token.Int)
	if err != nil {
		return nil, errutil.Err(err)
	}
	return constant.NewInt(typ, s)
}

// parseFloatConst parses a floating-point constant. A floating-point type as
// already been consumed.
//
// Syntax:
//    FloatConst = FloatType FloatValue .
//    FloatValue = float_lit .
//
// Examples:
//    double 3.14
//
// References:
//    http://llvm.org/docs/LangRef.html#simple-constants
func (p *parser) parseFloatConst(typ *types.Float) (*constant.Float, error) {
	s, err := p.expect(token.Float)
	if err != nil {
		return nil, errutil.Err(err)
	}
	return constant.NewFloat(typ, s)
}

// parsePointerConst parses a pointer constant. A pointer type as already been
// consumed.
//
// Syntax:
//    PointerConst = PointerType PointerValue .
//    PointerValue = "null" | Global .
//
// Examples:
//    i8* null
//    i32(i8*, ...)* @printf
//
// References:
//    http://llvm.org/docs/LangRef.html#simple-constants
//    http://llvm.org/docs/LangRef.html#global-variable-and-function-addresses
func (p *parser) parsePointerConst(typ *types.Pointer) (*constant.Pointer, error) {
	panic("parser.parsePointerConst: not yet implemented")
}

// parseVectorConst parses a vector constant. A vector type as already been
// consumed.
//
// Syntax:
//    VectorConst       = VectorType VectorValue .
//    VectorValue       = "<" VectorElementList ">" .
//    VectorElementList = VectorElement { "," VectorElement } .
//    VectorElement     = IntConst | FloatConst | PointerConst .
//
// Examples:
//    <2 x i32> <i32 10, i32 20>
//
// References:
//    http://llvm.org/docs/LangRef.html#complex-constants
func (p *parser) parseVectorConst(typ *types.Vector) (*constant.Vector, error) {
	// Vector constant; e.g.
	//    <2 x i32> <i32 10, i32 20>
	if _, err := p.expect(token.Less); err != nil {
		return nil, errutil.Err(err)
	}
	elems, err := p.parseElems(typ.Len(), typ.Elem())
	if err != nil {
		return nil, errutil.Err(err)
	}
	if _, err := p.expect(token.Greater); err != nil {
		return nil, errutil.Err(err)
	}
	return constant.NewVector(typ, elems)
}

// parseArrayConst parses an array constant. An array type as already been
// consumed.
//
// Syntax:
//    ArrayConst       = ArrayType ArrayValue .
//    ArrayValue       = ( "[" ArrayElementList "]" ) | ( "c" string_lit ) .
//    ArrayElementList = ArrayElement { "," ArrayElement } .
//    ArrayElement     = IntConst | FloatConst | PointerConst | VectorConst |
//                       ArrayConst | StructConst .
//
// Examples:
//    [2 x i32] [i32 10, i32 20]
//
// References:
//    http://llvm.org/docs/LangRef.html#complex-constants
func (p *parser) parseArrayConst(typ *types.Array) (*constant.Array, error) {
	// Character array constant; e.g.
	//    c"hello world\0A\00"
	if p.accept(token.KwC) {
		elemTyp, ok := typ.Elem().(*types.Int)
		if !ok || elemTyp.Size() != 8 {
			return nil, errutil.Newf(`invalid element type of character array constant; expected "i8", got %q`, typ.Elem())
		}
		s, err := p.expect(token.String)
		if err != nil {
			return nil, errutil.Err(err)
		}
		if len(s) != typ.Len() {
			return nil, errutil.Newf("constant array length mismatch; expected %d, got %d", typ.Len(), len(s))
		}
		var elems []constant.Constant
		for i := 0; i < len(s); i++ {
			b := s[i]
			// TODO: Find a clean way of representing constant vectors, arrays and
			// structures which requires less memory.

			// HACK: Using itoa as the constant API requires strings. This will be
			// fixed at the same time as the above TODO which locates a more
			// compact representation of constant vectors, arrays and structures.
			elem, err := constant.NewInt(elemTyp, strconv.Itoa(int(b)))
			if err != nil {
				return nil, errutil.Err(err)
			}
			elems = append(elems, elem)
		}
		return constant.NewArray(typ, elems)
	}

	// Array constant; e.g.
	//    [2 x i32] [i32 10, i32 20]
	if _, err := p.expect(token.Lbrack); err != nil {
		return nil, errutil.Err(err)
	}
	elems, err := p.parseElems(typ.Len(), typ.Elem())
	if err != nil {
		return nil, errutil.Err(err)
	}
	if _, err := p.expect(token.Rbrack); err != nil {
		return nil, errutil.Err(err)
	}
	return constant.NewArray(typ, elems)
}

// parseElems parses a comma separated list containing n constants of the given
// type.
func (p *parser) parseElems(n int, elemTyp types.Type) (elems []constant.Constant, err error) {
	for i := 0; i < n; i++ {
		if i > 0 {
			if _, err := p.expect(token.Comma); err != nil {
				return nil, errutil.Err(err)
			}
		}
		elem, err := p.parseConst()
		if err != nil {
			return nil, errutil.Err(err)
		}
		if !elemTyp.Equal(elem.Type()) {
			return nil, errutil.Newf("constant element type mismatch; expected %q, got %q", elemTyp, elem.Type())
		}
		elems = append(elems, elem)
	}
	return elems, nil
}

// parseStructConst parses a structure constant. A structure type as already
// been consumed.
//
// Syntax:
//    StructConst       = StructType StructValue .
//    StructValue       = "{" StructElementList "}" |
//                        "<" "{" StructElementList "}" ">" .
//    StructElementList = StructElement { "," StructElement } .
//    StructElement     = IntConst | FloatConst | PointerConst | VectorConst |
//                        ArrayConst | StructConst .
//
// Examples:
//    {i8, i32} {i8 10, i32 20}
//
// References:
//    http://llvm.org/docs/LangRef.html#complex-constants
func (p *parser) parseStructConst(typ *types.Struct) (*constant.Struct, error) {
	// Array constant; e.g.
	//    {i8, i32} {i8 10, i32 20}
	if _, err := p.expect(token.Lbrace); err != nil {
		return nil, errutil.Err(err)
	}
	var fields []constant.Constant
	for i, fieldTyp := range typ.Fields() {
		if i > 0 {
			if _, err := p.expect(token.Comma); err != nil {
				return nil, errutil.Err(err)
			}
		}
		field, err := p.parseConst()
		if err != nil {
			return nil, errutil.Err(err)
		}
		if !fieldTyp.Equal(field.Type()) {
			return nil, errutil.Newf("constant field type mismatch; expected %q, got %q", fieldTyp, field.Type())
		}
		fields = append(fields, field)
	}
	if _, err := p.expect(token.Rbrace); err != nil {
		return nil, errutil.Err(err)
	}
	return constant.NewStruct(typ, fields)
}
