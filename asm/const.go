package asm

import (
	"fmt"

	"github.com/llir/ll/ast"
	"github.com/llir/llvm/internal/enc"
	"github.com/llir/llvm/ir"
	"github.com/llir/llvm/ir/constant"
	"github.com/llir/llvm/ir/types"
	"github.com/pkg/errors"
)

// === [ Translate AST to IR ] =================================================

// irConstant translates the AST constant into an equivalent IR constant.
func (gen *generator) irConstant(t types.Type, old ast.Constant) (constant.Constant, error) {
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
		return constant.NewZeroInitializer(t), nil
	case *ast.UndefConst:
		return constant.NewUndef(t), nil
	case *ast.BlockAddressConst:
		return gen.irBlockAddressConst(t, old)
	case *ast.GlobalIdent:
		ident := globalIdent(*old)
		c, ok := gen.new.globals[ident]
		if !ok {
			return nil, errors.Errorf("unable to locate global identifier %q", ident.Ident())
		}
		return c, nil
	case ast.ConstantExpr:
		return gen.irConstantExpr(t, old)
	default:
		panic(fmt.Errorf("support for AST constant %T not yet implemented", old))
	}
}

// irTypeConst translates the AST type-constant pair into an equivalent IR
// constant.
func (gen *generator) irTypeConst(old ast.TypeConst) (constant.Constant, error) {
	// Type.
	typ, err := gen.irType(old.Typ())
	if err != nil {
		return nil, errors.WithStack(err)
	}
	// Constant.
	return gen.irConstant(typ, old.Val())
}

// --- [ Boolean constants ] ---------------------------------------------------

// irBoolConst translates the AST boolean constant into an equivalent IR integer
// constant.
func (gen *generator) irBoolConst(t types.Type, old *ast.BoolConst) (*constant.Int, error) {
	typ, ok := t.(*types.IntType)
	if !ok {
		return nil, errors.Errorf("invalid type of boolean constant; expected *types.IntType, got %T", t)
	}
	if !typ.Equal(types.I1) {
		return nil, errors.Errorf("boolean type mismatch; expected %q, got %q", types.I1, typ)
	}
	return constant.NewBool(boolLit(old.BoolLit())), nil
}

// --- [ Integer constants ] ---------------------------------------------------

// irIntConst translates the AST integer constant into an equivalent IR integer
// constant.
func (gen *generator) irIntConst(t types.Type, old *ast.IntConst) (*constant.Int, error) {
	typ, ok := t.(*types.IntType)
	if !ok {
		line, col := old.LineColumn()
		return nil, errors.Errorf("%d:%d: invalid type of integer constant; expected *types.IntType, got %T", line, col, t)
	}
	s := old.IntLit().Text()
	return constant.NewIntFromString(typ, s)
}

// --- [ Floating-point constants ] --------------------------------------------

// irFloatConst translates the AST floating-point constant into an equivalent IR
// floating-point constant.
func (gen *generator) irFloatConst(t types.Type, old *ast.FloatConst) (*constant.Float, error) {
	typ, ok := t.(*types.FloatType)
	if !ok {
		return nil, errors.Errorf("invalid type of floating-point constant; expected *types.FloatType, got %T", t)
	}
	s := old.FloatLit().Text()
	return constant.NewFloatFromString(typ, s)
}

// --- [ Null pointer constants ] ----------------------------------------------

// irNullConst translates the AST null pointer constant into an equivalent IR
// null pointer constant.
func (gen *generator) irNullConst(t types.Type, old *ast.NullConst) (*constant.Null, error) {
	typ, ok := t.(*types.PointerType)
	if !ok {
		return nil, errors.Errorf("invalid type of null pointer constant; expected *types.PointerType, got %T", t)
	}
	return constant.NewNull(typ), nil
}

// --- [ Token constants ] -----------------------------------------------------

// irNoneConst translates the AST none token constant into an equivalent IR none
// token constant.
func (gen *generator) irNoneConst(t types.Type, old *ast.NoneConst) (constant.Constant, error) {
	if !t.Equal(types.Token) {
		return nil, errors.Errorf("invalid type of none token constant; expected %q, got %q", types.Token, t)
	}
	return constant.None, nil
}

// --- [ Struct constants ] ----------------------------------------------------

// irStructConst translates the AST struct constant into an equivalent IR struct
// constant.
func (gen *generator) irStructConst(t types.Type, old *ast.StructConst) (*constant.Struct, error) {
	typ, ok := t.(*types.StructType)
	if !ok {
		return nil, errors.Errorf("invalid type of struct constant; expected *types.StructType, got %T", t)
	}
	var fields []constant.Constant
	if oldFields := old.Fields(); len(oldFields) > 0 {
		fields = make([]constant.Constant, len(oldFields))
		for i, oldField := range oldFields {
			field, err := gen.irTypeConst(oldField)
			if err != nil {
				return nil, errors.WithStack(err)
			}
			fields[i] = field
		}
	}
	c := constant.NewStruct(typ, fields...)
	return c, nil
}

// --- [ Array constants ] -----------------------------------------------------

// irArrayConst translates the AST array constant into an equivalent IR array
// constant.
func (gen *generator) irArrayConst(t types.Type, old *ast.ArrayConst) (*constant.Array, error) {
	typ, ok := t.(*types.ArrayType)
	if !ok {
		return nil, errors.Errorf("invalid type of array constant; expected *types.ArrayType, got %T", t)
	}
	oldElems := old.Elems()
	if len(oldElems) == 0 {
		typ := types.NewArray(0, typ.ElemType)
		if !t.Equal(typ) {
			return nil, errors.Errorf("array type mismatch; expected %q, got %q", typ, t)
		}
		return &constant.Array{Typ: typ}, nil
	}
	elems := make([]constant.Constant, len(oldElems))
	for i, oldElem := range oldElems {
		elem, err := gen.irTypeConst(oldElem)
		if err != nil {
			return nil, errors.WithStack(err)
		}
		elems[i] = elem
	}
	c := constant.NewArray(typ, elems...)
	return c, nil
}

// irCharArrayConst translates the AST character array constant into an
// equivalent IR character array constant.
func (gen *generator) irCharArrayConst(t types.Type, old *ast.CharArrayConst) (*constant.CharArray, error) {
	data := enc.Unquote(old.Val().Text())
	c := constant.NewCharArray(data)
	if !t.Equal(c.Typ) {
		return nil, errors.Errorf("character array type mismatch; expected %q, got %q", c.Typ, t)
	}
	return c, nil
}

// --- [ Vector constants ] ----------------------------------------------------

// irVectorConst translates the AST vector constant into an equivalent IR vector
// constant.
func (gen *generator) irVectorConst(t types.Type, old *ast.VectorConst) (*constant.Vector, error) {
	typ, ok := t.(*types.VectorType)
	if !ok {
		return nil, errors.Errorf("invalid type of vector constant; expected *types.VectorType, got %T", t)
	}
	oldElems := old.Elems()
	if len(oldElems) == 0 {
		return nil, errors.New("zero element vector is illegal")
	}
	elems := make([]constant.Constant, len(oldElems))
	for i, oldElem := range oldElems {
		elem, err := gen.irTypeConst(oldElem)
		if err != nil {
			return nil, errors.WithStack(err)
		}
		elems[i] = elem
	}
	c := constant.NewVector(typ, elems...)
	return c, nil
}

// --- [ Addresses of basic blocks ] -------------------------------------------

// irBlockAddressConst translates the AST blockaddress constant into an
// equivalent IR blockaddress constant.
func (gen *generator) irBlockAddressConst(t types.Type, old *ast.BlockAddressConst) (*constant.BlockAddress, error) {
	// Function.
	funcName := globalIdent(old.Func())
	v, ok := gen.new.globals[funcName]
	if !ok {
		return nil, errors.Errorf("unable to locate global identifier %q", funcName)
	}
	f, ok := v.(*ir.Func)
	if !ok {
		return nil, errors.Errorf("invalid function type; expected *ir.Func, got %T", v)
	}
	// Basic block.
	blockIdent := localIdent(old.Block())
	// Add dummy basic block to track the name recorded by the AST. Resolve the
	// proper basic block after translation of function bodies and assignment of
	// local IDs.
	block := &ir.Block{
		LocalIdent: blockIdent,
	}
	c := constant.NewBlockAddress(f, block)
	gen.todo = append(gen.todo, c)
	if typ := c.Type(); !t.Equal(typ) {
		return nil, errors.Errorf("blockaddress constant type mismatch; expected %q, got %q", typ, t)
	}
	return c, nil
}
