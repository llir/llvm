package asm

import (
	"fmt"

	"github.com/llir/l/ir"
	"github.com/llir/l/ir/types"
	"github.com/mewmew/l-tm/asm/ll/ast"
	"github.com/mewmew/l-tm/internal/enc"
	"github.com/pkg/errors"
)

// === [ Constants ] ===========================================================

func (gen *generator) irConstant(t types.Type, old ast.Constant) (ir.Constant, error) {
	switch old := old.(type) {
	case *ast.BoolConst:
		return gen.irBoolConst(t, old)
	case *ast.IntConst:
		return gen.irIntConst(t, old)
	case *ast.FloatConst:
		return gen.irFloatConst(t, old)
	case *ast.NullConst:
		return gen.irNullConst(t, old)
	case *ast.NoneConst:
		return gen.irNoneConst(t, old)
	case *ast.StructConst:
		return gen.irStructConst(t, old)
	case *ast.ArrayConst:
		return gen.irArrayConst(t, old)
	case *ast.CharArrayConst:
		return gen.irCharArrayConst(t, old)
	case *ast.VectorConst:
		return gen.irVectorConst(t, old)
	case *ast.ZeroInitializerConst:
		return gen.irZeroInitializerConst(t, old)
	case *ast.UndefConst:
		return gen.irUndefConst(t, old)
	case *ast.BlockAddressConst:
		return gen.irBlockAddressConst(t, old)
	case *ast.GlobalIdent:
		name := global(*old)
		g, ok := gen.gs[name]
		if !ok {
			return nil, errors.Errorf("unable to locate global identifier %q", enc.Global(name))
		}
		return g, nil
	case ast.ConstantExpr:
		return gen.irConstantExpr(t, old)
	default:
		panic(fmt.Errorf("support for AST constant %T not yet implemented", old))
	}
}

func (gen *generator) irTypeConst(old ast.TypeConst) (ir.Constant, error) {
	typ, err := gen.irType(old.Typ())
	if err != nil {
		return nil, errors.WithStack(err)
	}
	return gen.irConstant(typ, old.Val())
}

// --- [ Boolean Constants ] ---------------------------------------------------

func (gen *generator) irBoolConst(t types.Type, old *ast.BoolConst) (*ir.ConstInt, error) {
	typ, ok := t.(*types.IntType)
	if !ok {
		return nil, errors.Errorf("invalid type of boolean constant; expected *types.IntType, got %T", t)
	}
	if typ.BitSize != 1 {
		return nil, errors.Errorf("invalid integer type bit size of boolean constant; expected 1, got %d", typ.BitSize)
	}
	v := boolLit(old.BoolLit())
	if v {
		return ir.True, nil
	}
	return ir.False, nil
}

// --- [ Integer Constants ] ---------------------------------------------------

func (gen *generator) irIntConst(t types.Type, old *ast.IntConst) (*ir.ConstInt, error) {
	typ, ok := t.(*types.IntType)
	if !ok {
		return nil, errors.Errorf("invalid type of integer constant; expected *types.IntType, got %T", t)
	}
	s := old.IntLit().Text()
	return ir.NewIntFromString(typ, s)
}

// --- [ Floating-point Constants ] --------------------------------------------

func (gen *generator) irFloatConst(t types.Type, old *ast.FloatConst) (*ir.ConstFloat, error) {
	typ, ok := t.(*types.FloatType)
	if !ok {
		return nil, errors.Errorf("invalid type of floating-point constant; expected *types.FloatType, got %T", t)
	}
	s := old.FloatLit().Text()
	return ir.NewFloatFromString(typ, s)
}

// --- [ Null Pointer Constants ] ----------------------------------------------

func (gen *generator) irNullConst(t types.Type, old *ast.NullConst) (*ir.ConstNull, error) {
	typ, ok := t.(*types.PointerType)
	if !ok {
		return nil, errors.Errorf("invalid type of null constant; expected *types.PointerType, got %T", t)
	}
	return ir.NewNull(typ), nil
}

// --- [ Token Constants ] -----------------------------------------------------

func (gen *generator) irNoneConst(t types.Type, old *ast.NoneConst) (*ir.ConstNone, error) {
	panic("not yet implemented")
}

// --- [ Structure Constants ] -------------------------------------------------

func (gen *generator) irStructConst(t types.Type, old *ast.StructConst) (*ir.ConstStruct, error) {
	typ, ok := t.(*types.StructType)
	if !ok {
		return nil, errors.Errorf("invalid type of struct constant; expected *types.StructType, got %T", t)
	}
	var fields []ir.Constant
	for _, f := range old.Fields() {
		field, err := gen.irTypeConst(f)
		if err != nil {
			return nil, errors.WithStack(err)
		}
		fields = append(fields, field)
	}
	return ir.NewStruct(typ, fields...), nil
}

// --- [ Array Constants ] -----------------------------------------------------

func (gen *generator) irArrayConst(t types.Type, old *ast.ArrayConst) (*ir.ConstArray, error) {
	typ, ok := t.(*types.ArrayType)
	if !ok {
		return nil, errors.Errorf("invalid type of array constant; expected *types.ArrayType, got %T", t)
	}
	var elems []ir.Constant
	for _, e := range old.Elems() {
		elem, err := gen.irTypeConst(e)
		if err != nil {
			return nil, errors.WithStack(err)
		}
		elems = append(elems, elem)
	}
	return ir.NewArray(typ, elems...), nil
}

func (gen *generator) irCharArrayConst(t types.Type, old *ast.CharArrayConst) (*ir.ConstCharArray, error) {
	data := stringLitBytes(old.Val())
	// TODO: validate that t and type of newly created character array constant
	// match.

	// TODO: also decide whether to update ir.NewCharArray to include a type as
	// its first parameter, thus making it consistent with ir.NewArray.
	return ir.NewCharArray(data), nil
}

// --- [ Vector Constants ] ----------------------------------------------------

func (gen *generator) irVectorConst(t types.Type, old *ast.VectorConst) (*ir.ConstVector, error) {
	panic("not yet implemented")
}

// --- [ Zero Initialization Constants ] ---------------------------------------

func (gen *generator) irZeroInitializerConst(t types.Type, old *ast.ZeroInitializerConst) (*ir.ConstZeroInitializer, error) {
	return ir.NewZeroInitializer(t), nil
}

// --- [ Undefined Values ] ----------------------------------------------------

func (gen *generator) irUndefConst(t types.Type, old *ast.UndefConst) (*ir.ConstUndef, error) {
	panic("not yet implemented")
}

// --- [ Addresses of Basic Blocks ] -------------------------------------------

func (gen *generator) irBlockAddressConst(t types.Type, old *ast.BlockAddressConst) (*ir.ConstBlockAddress, error) {
	panic("not yet implemented")
}
