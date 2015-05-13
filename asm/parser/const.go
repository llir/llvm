package parser

import (
	"github.com/llir/llvm/consts"
	"github.com/llir/llvm/types"
	"github.com/mewkiz/pkg/errutil"
)

// TODO: Add EBNF for ConstExpr.

// parseConst parses a constant value.
//
//    Const = IntConst | FloatConst | PointerConst | VectorConst | ArrayConst |
//            StructConst | ConstExpr .
//
// References:
//    http://llvm.org/docs/LangRef.html#constants
func (p *parser) parseConst() (consts.Constant, error) {
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
func (p *parser) parseIntConst(typ *types.Int) (*consts.Int, error) {
	panic("not yet implemented")
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
func (p *parser) parseFloatConst(typ *types.Float) (*consts.Float, error) {
	panic("not yet implemented")
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
func (p *parser) parsePointerConst(typ *types.Pointer) (*consts.Pointer, error) {
	panic("not yet implemented")
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
func (p *parser) parseVectorConst(typ *types.Vector) (*consts.Vector, error) {
	panic("not yet implemented")
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
func (p *parser) parseArrayConst(typ *types.Array) (*consts.Array, error) {
	panic("not yet implemented")
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
func (p *parser) parseStructConst(typ *types.Struct) (*consts.Struct, error) {
	panic("not yet implemented")
}
