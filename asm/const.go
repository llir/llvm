package asm

import (
	"fmt"

	"github.com/llir/ll/ast"
	"github.com/llir/llvm/ir"
	"github.com/llir/llvm/ir/constant"
	"github.com/llir/llvm/ir/types"
	"github.com/pkg/errors"
)

// === [ Constants ] ===========================================================

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
		return gen.irZeroInitializerConst(t, old)
	case *ast.UndefConst:
		return gen.irUndefConst(t, old)
	case *ast.BlockAddressConst:
		return gen.irBlockAddressConst(t, old)
	case *ast.GlobalIdent:
		ident := globalIdent(*old)
		v, ok := gen.new.globals[ident]
		if !ok {
			return nil, errors.Errorf("unable to locate global identifier %q", ident)
		}
		return v, nil
	case ast.ConstantExpr:
		return gen.irConstantExpr(t, old)
	default:
		panic(fmt.Errorf("support for AST constant %T not yet implemented", old))
	}
}

func (gen *generator) irTypeConst(old ast.TypeConst) (constant.Constant, error) {
	// Type.
	typ, err := gen.irType(old.Typ())
	if err != nil {
		return nil, errors.WithStack(err)
	}
	// Constant.
	return gen.irConstant(typ, old.Val())
}

// --- [ Boolean Constants ] ---------------------------------------------------

func (gen *generator) irBoolConst(t types.Type, old *ast.BoolConst) (*constant.Int, error) {
	typ, ok := t.(*types.IntType)
	if !ok {
		return nil, errors.Errorf("invalid type of boolean constant; expected *types.IntType, got %T", t)
	}
	if typ.BitSize != 1 {
		return nil, errors.Errorf("invalid integer type bit size of boolean constant; expected 1, got %d", typ.BitSize)
	}
	v := boolLit(old.BoolLit())
	if v {
		return constant.True, nil
	}
	return constant.False, nil
}

// --- [ Integer Constants ] ---------------------------------------------------

func (gen *generator) irIntConst(t types.Type, old *ast.IntConst) (*constant.Int, error) {
	typ, ok := t.(*types.IntType)
	if !ok {
		return nil, errors.Errorf("invalid type of integer constant; expected *types.IntType, got %T", t)
	}
	s := old.IntLit().Text()
	return constant.NewIntFromString(typ, s)
}

// --- [ Floating-point Constants ] --------------------------------------------

func (gen *generator) irFloatConst(t types.Type, old *ast.FloatConst) (*constant.Float, error) {
	typ, ok := t.(*types.FloatType)
	if !ok {
		return nil, errors.Errorf("invalid type of floating-point constant; expected *types.FloatType, got %T", t)
	}
	s := old.FloatLit().Text()
	return constant.NewFloatFromString(typ, s)
}

// --- [ Null Pointer Constants ] ----------------------------------------------

func (gen *generator) irNullConst(t types.Type, old *ast.NullConst) (*constant.Null, error) {
	typ, ok := t.(*types.PointerType)
	if !ok {
		return nil, errors.Errorf("invalid type of null constant; expected *types.PointerType, got %T", t)
	}
	return constant.NewNull(typ), nil
}

// --- [ Token Constants ] -----------------------------------------------------

func (gen *generator) irNoneConst(t types.Type, old *ast.NoneConst) (constant.Constant, error) {
	// TODO: validate type t.
	return constant.None, nil
}

// --- [ Structure Constants ] -------------------------------------------------

func (gen *generator) irStructConst(t types.Type, old *ast.StructConst) (*constant.Struct, error) {
	typ, ok := t.(*types.StructType)
	if !ok {
		return nil, errors.Errorf("invalid type of struct constant; expected *types.StructType, got %T", t)
	}
	var fields []constant.Constant
	for _, f := range old.Fields() {
		field, err := gen.irTypeConst(f)
		if err != nil {
			return nil, errors.WithStack(err)
		}
		fields = append(fields, field)
	}
	c := constant.NewStruct(fields...)
	c.Typ.TypeName = typ.TypeName
	c.Typ.Packed = typ.Packed
	return c, nil
}

// --- [ Array Constants ] -----------------------------------------------------

func (gen *generator) irArrayConst(t types.Type, old *ast.ArrayConst) (*constant.Array, error) {
	typ, ok := t.(*types.ArrayType)
	if !ok {
		return nil, errors.Errorf("invalid type of array constant; expected *types.ArrayType, got %T", t)
	}
	var elems []constant.Constant
	for _, e := range old.Elems() {
		elem, err := gen.irTypeConst(e)
		if err != nil {
			return nil, errors.WithStack(err)
		}
		elems = append(elems, elem)
	}
	c := constant.NewArray(elems...)
	// TODO: check that typ is equal to c.Typ.
	_ = typ
	return c, nil
}

func (gen *generator) irCharArrayConst(t types.Type, old *ast.CharArrayConst) (*constant.CharArray, error) {
	data := stringLitBytes(old.Val())
	expr := constant.NewCharArray(data)
	// TODO: validate t against expr.Typ.
	return expr, nil
}

// --- [ Vector Constants ] ----------------------------------------------------

func (gen *generator) irVectorConst(t types.Type, old *ast.VectorConst) (*constant.Vector, error) {
	typ, ok := t.(*types.VectorType)
	if !ok {
		return nil, errors.Errorf("invalid type of vector constant; expected *types.VectorType, got %T", t)
	}
	var elems []constant.Constant
	for _, e := range old.Elems() {
		elem, err := gen.irTypeConst(e)
		if err != nil {
			return nil, errors.WithStack(err)
		}
		elems = append(elems, elem)
	}
	c := constant.NewVector(elems...)
	// TODO: check that typ is equal to c.Typ.
	_ = typ
	return c, nil
}

// --- [ Zero Initialization Constants ] ---------------------------------------

func (gen *generator) irZeroInitializerConst(t types.Type, old *ast.ZeroInitializerConst) (*constant.ZeroInitializer, error) {
	return constant.NewZeroInitializer(t), nil
}

// --- [ Undefined Values ] ----------------------------------------------------

func (gen *generator) irUndefConst(t types.Type, old *ast.UndefConst) (*constant.Undef, error) {
	return constant.NewUndef(t), nil
}

// --- [ Addresses of Basic Blocks ] -------------------------------------------

func (gen *generator) irBlockAddressConst(t types.Type, old *ast.BlockAddressConst) (*constant.BlockAddress, error) {
	// Function.
	funcName := globalIdent(old.Func())
	v, ok := gen.new.globals[funcName]
	if !ok {
		return nil, errors.Errorf("unable to locate global identifier %q", funcName)
	}
	f, ok := v.(*ir.Function)
	if !ok {
		return nil, errors.Errorf("invalid function type; expected *ir.Function, got %T", v)
	}
	// Basic block.
	blockIdent := localIdent(old.Block())
	// Add dummy basic block to track the name recorded by the AST. Resolve the
	// proper basic block after translation of function bodies and assignment of
	// local IDs.
	block := &ir.BasicBlock{
		LocalIdent: blockIdent,
	}
	expr := constant.NewBlockAddress(f, block)
	gen.todo = append(gen.todo, expr)
	// TODO: validate type t against expr.Typ. Store t in todo?
	return expr, nil
}

// Pre-condition: translate function body and assign local IDs of c.Func.
func fixBlockAddressConst(c *constant.BlockAddress) error {
	f, ok := c.Func.(*ir.Function)
	if !ok {
		panic(fmt.Errorf("invalid function type in blockaddress constant; expected *ir.Function, got %T", c.Func))
	}
	bb, ok := c.Block.(*ir.BasicBlock)
	if !ok {
		panic(fmt.Errorf("invalid basic block type in blockaddress constant; expected *ir.BasicBlock, got %T", c.Block))
	}
	for _, block := range f.Blocks {
		if block.LocalIdent == bb.LocalIdent {
			c.Block = block
			return nil
		}
	}
	return errors.Errorf("unable to locate basic block %q in function %q", c.Block.Ident(), f.Ident())
}
