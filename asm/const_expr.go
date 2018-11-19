package asm

import (
	"fmt"

	"github.com/llir/ll/ast"
	asmenum "github.com/llir/llvm/asm/enum"
	"github.com/llir/llvm/ir/constant"
	"github.com/llir/llvm/ir/types"
	"github.com/pkg/errors"
)

// === [ Constant expressions ] ================================================

// irConstantExpr translates the given AST constant expression into an
// equivalent IR constant expression.
func (gen *generator) irConstantExpr(t types.Type, old ast.ConstantExpr) (constant.Expression, error) {
	switch old := old.(type) {
	// Binary expressions
	case *ast.AddExpr:
		return gen.irAddExpr(t, old)
	case *ast.FAddExpr:
		return gen.irFAddExpr(t, old)
	case *ast.SubExpr:
		return gen.irSubExpr(t, old)
	case *ast.FSubExpr:
		return gen.irFSubExpr(t, old)
	case *ast.MulExpr:
		return gen.irMulExpr(t, old)
	case *ast.FMulExpr:
		return gen.irFMulExpr(t, old)
	case *ast.UDivExpr:
		return gen.irUDivExpr(t, old)
	case *ast.SDivExpr:
		return gen.irSDivExpr(t, old)
	case *ast.FDivExpr:
		return gen.irFDivExpr(t, old)
	case *ast.URemExpr:
		return gen.irURemExpr(t, old)
	case *ast.SRemExpr:
		return gen.irSRemExpr(t, old)
	case *ast.FRemExpr:
		return gen.irFRemExpr(t, old)
	// Bitwise expressions
	case *ast.ShlExpr:
		return gen.irShlExpr(t, old)
	case *ast.LShrExpr:
		return gen.irLShrExpr(t, old)
	case *ast.AShrExpr:
		return gen.irAShrExpr(t, old)
	case *ast.AndExpr:
		return gen.irAndExpr(t, old)
	case *ast.OrExpr:
		return gen.irOrExpr(t, old)
	case *ast.XorExpr:
		return gen.irXorExpr(t, old)
	// Vector expressions
	case *ast.ExtractElementExpr:
		return gen.irExtractElementExpr(t, old)
	case *ast.InsertElementExpr:
		return gen.irInsertElementExpr(t, old)
	case *ast.ShuffleVectorExpr:
		return gen.irShuffleVectorExpr(t, old)
	// Aggregate expressions
	case *ast.ExtractValueExpr:
		return gen.irExtractValueExpr(t, old)
	case *ast.InsertValueExpr:
		return gen.irInsertValueExpr(t, old)
	// Memory expressions
	case *ast.GetElementPtrExpr:
		return gen.irGetElementPtrExpr(t, old)
	// Conversion expressions
	case *ast.TruncExpr:
		return gen.irTruncExpr(t, old)
	case *ast.ZExtExpr:
		return gen.irZExtExpr(t, old)
	case *ast.SExtExpr:
		return gen.irSExtExpr(t, old)
	case *ast.FPTruncExpr:
		return gen.irFPTruncExpr(t, old)
	case *ast.FPExtExpr:
		return gen.irFPExtExpr(t, old)
	case *ast.FPToUIExpr:
		return gen.irFPToUIExpr(t, old)
	case *ast.FPToSIExpr:
		return gen.irFPToSIExpr(t, old)
	case *ast.UIToFPExpr:
		return gen.irUIToFPExpr(t, old)
	case *ast.SIToFPExpr:
		return gen.irSIToFPExpr(t, old)
	case *ast.PtrToIntExpr:
		return gen.irPtrToIntExpr(t, old)
	case *ast.IntToPtrExpr:
		return gen.irIntToPtrExpr(t, old)
	case *ast.BitCastExpr:
		return gen.irBitCastExpr(t, old)
	case *ast.AddrSpaceCastExpr:
		return gen.irAddrSpaceCastExpr(t, old)
	// Other expressions
	case *ast.ICmpExpr:
		return gen.irICmpExpr(t, old)
	case *ast.FCmpExpr:
		return gen.irFCmpExpr(t, old)
	case *ast.SelectExpr:
		return gen.irSelectExpr(t, old)
	default:
		panic(fmt.Errorf("support for AST constant expression %T not yet implemented", old))
	}
}

// --- [ Binary expressions ] -------------------------------------------------

// ~~~ [ add ] ~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~

// irAddExpr translates the given AST add constant expression into an equivalent
// IR constant expression.
func (gen *generator) irAddExpr(t types.Type, old *ast.AddExpr) (*constant.ExprAdd, error) {
	// (optional) Overflow flags.
	overflowFlags := irOverflowFlags(old.OverflowFlags())
	// X operand.
	x, err := gen.irTypeConst(old.X())
	if err != nil {
		return nil, errors.WithStack(err)
	}
	// Y operand.
	y, err := gen.irTypeConst(old.Y())
	if err != nil {
		return nil, errors.WithStack(err)
	}
	expr := constant.NewAdd(x, y)
	expr.OverflowFlags = overflowFlags
	return expr, nil
}

// ~~~ [ fadd ] ~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~

// irFAddExpr translates the given AST fadd constant expression into an
// equivalent IR constant expression.
func (gen *generator) irFAddExpr(t types.Type, old *ast.FAddExpr) (*constant.ExprFAdd, error) {
	// X operand.
	x, err := gen.irTypeConst(old.X())
	if err != nil {
		return nil, errors.WithStack(err)
	}
	// Y operand.
	y, err := gen.irTypeConst(old.Y())
	if err != nil {
		return nil, errors.WithStack(err)
	}
	expr := constant.NewFAdd(x, y)
	return expr, nil
}

// ~~~ [ sub ] ~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~

// irSubExpr translates the given AST sub constant expression into an equivalent
// IR constant expression.
func (gen *generator) irSubExpr(t types.Type, old *ast.SubExpr) (*constant.ExprSub, error) {
	// (optional) Overflow flags.
	overflowFlags := irOverflowFlags(old.OverflowFlags())
	// X operand.
	x, err := gen.irTypeConst(old.X())
	if err != nil {
		return nil, errors.WithStack(err)
	}
	// Y operand.
	y, err := gen.irTypeConst(old.Y())
	if err != nil {
		return nil, errors.WithStack(err)
	}
	expr := constant.NewSub(x, y)
	expr.OverflowFlags = overflowFlags
	return expr, nil
}

// ~~~ [ fsub ] ~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~

// irFSubExpr translates the given AST fsub constant expression into an
// equivalent IR constant expression.
func (gen *generator) irFSubExpr(t types.Type, old *ast.FSubExpr) (*constant.ExprFSub, error) {
	// X operand.
	x, err := gen.irTypeConst(old.X())
	if err != nil {
		return nil, errors.WithStack(err)
	}
	// Y operand.
	y, err := gen.irTypeConst(old.Y())
	if err != nil {
		return nil, errors.WithStack(err)
	}
	expr := constant.NewFSub(x, y)
	return expr, nil
}

// ~~~ [ mul ] ~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~

// irMulExpr translates the given AST mul constant expression into an equivalent
// IR constant expression.
func (gen *generator) irMulExpr(t types.Type, old *ast.MulExpr) (*constant.ExprMul, error) {
	// (optional) Overflow flags.
	overflowFlags := irOverflowFlags(old.OverflowFlags())
	// X operand.
	x, err := gen.irTypeConst(old.X())
	if err != nil {
		return nil, errors.WithStack(err)
	}
	// Y operand.
	y, err := gen.irTypeConst(old.Y())
	if err != nil {
		return nil, errors.WithStack(err)
	}
	expr := constant.NewMul(x, y)
	expr.OverflowFlags = overflowFlags
	return expr, nil
}

// ~~~ [ fmul ] ~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~

// irFMulExpr translates the given AST fmul constant expression into an
// equivalent IR constant expression.
func (gen *generator) irFMulExpr(t types.Type, old *ast.FMulExpr) (*constant.ExprFMul, error) {
	// X operand.
	x, err := gen.irTypeConst(old.X())
	if err != nil {
		return nil, errors.WithStack(err)
	}
	// Y operand.
	y, err := gen.irTypeConst(old.Y())
	if err != nil {
		return nil, errors.WithStack(err)
	}
	expr := constant.NewFMul(x, y)
	return expr, nil
}

// ~~~ [ udiv ] ~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~

// irUDivExpr translates the given AST udiv constant expression into an
// equivalent IR constant expression.
func (gen *generator) irUDivExpr(t types.Type, old *ast.UDivExpr) (*constant.ExprUDiv, error) {
	// (optional) Exact.
	exact := old.Exact().IsValid()
	// X operand.
	x, err := gen.irTypeConst(old.X())
	if err != nil {
		return nil, errors.WithStack(err)
	}
	// Y operand.
	y, err := gen.irTypeConst(old.Y())
	if err != nil {
		return nil, errors.WithStack(err)
	}
	expr := constant.NewUDiv(x, y)
	expr.Exact = exact
	return expr, nil
}

// ~~~ [ sdiv ] ~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~

// irSDivExpr translates the given AST sdiv constant expression into an
// equivalent IR constant expression.
func (gen *generator) irSDivExpr(t types.Type, old *ast.SDivExpr) (*constant.ExprSDiv, error) {
	// (optional) Exact.
	exact := old.Exact().IsValid()
	// X operand.
	x, err := gen.irTypeConst(old.X())
	if err != nil {
		return nil, errors.WithStack(err)
	}
	// Y operand.
	y, err := gen.irTypeConst(old.Y())
	if err != nil {
		return nil, errors.WithStack(err)
	}
	expr := constant.NewSDiv(x, y)
	expr.Exact = exact
	return expr, nil
}

// ~~~ [ fdiv ] ~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~

// irFDivExpr translates the given AST fdiv constant expression into an
// equivalent IR constant expression.
func (gen *generator) irFDivExpr(t types.Type, old *ast.FDivExpr) (*constant.ExprFDiv, error) {
	// X operand.
	x, err := gen.irTypeConst(old.X())
	if err != nil {
		return nil, errors.WithStack(err)
	}
	// Y operand.
	y, err := gen.irTypeConst(old.Y())
	if err != nil {
		return nil, errors.WithStack(err)
	}
	expr := constant.NewFDiv(x, y)
	return expr, nil
}

// ~~~ [ urem ] ~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~

// irURemExpr translates the given AST urem constant expression into an
// equivalent IR constant expression.
func (gen *generator) irURemExpr(t types.Type, old *ast.URemExpr) (*constant.ExprURem, error) {
	// X operand.
	x, err := gen.irTypeConst(old.X())
	if err != nil {
		return nil, errors.WithStack(err)
	}
	// Y operand.
	y, err := gen.irTypeConst(old.Y())
	if err != nil {
		return nil, errors.WithStack(err)
	}
	expr := constant.NewURem(x, y)
	return expr, nil
}

// ~~~ [ srem ] ~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~

// irSRemExpr translates the given AST srem constant expression into an
// equivalent IR constant expression.
func (gen *generator) irSRemExpr(t types.Type, old *ast.SRemExpr) (*constant.ExprSRem, error) {
	// X operand.
	x, err := gen.irTypeConst(old.X())
	if err != nil {
		return nil, errors.WithStack(err)
	}
	// Y operand.
	y, err := gen.irTypeConst(old.Y())
	if err != nil {
		return nil, errors.WithStack(err)
	}
	expr := constant.NewSRem(x, y)
	return expr, nil
}

// ~~~ [ frem ] ~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~

// irFRemExpr translates the given AST frem constant expression into an
// equivalent IR constant expression.
func (gen *generator) irFRemExpr(t types.Type, old *ast.FRemExpr) (*constant.ExprFRem, error) {
	// X operand.
	x, err := gen.irTypeConst(old.X())
	if err != nil {
		return nil, errors.WithStack(err)
	}
	// Y operand.
	y, err := gen.irTypeConst(old.Y())
	if err != nil {
		return nil, errors.WithStack(err)
	}
	expr := constant.NewFRem(x, y)
	return expr, nil
}

// --- [ Bitwise expressions ] -------------------------------------------------

// ~~~ [ shl ] ~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~

// irShlExpr translates the given AST shl constant expression into an equivalent
// IR constant expression.
func (gen *generator) irShlExpr(t types.Type, old *ast.ShlExpr) (*constant.ExprShl, error) {
	// (optional) Overflow flags.
	overflowFlags := irOverflowFlags(old.OverflowFlags())
	// X operand.
	x, err := gen.irTypeConst(old.X())
	if err != nil {
		return nil, errors.WithStack(err)
	}
	// Y operand.
	y, err := gen.irTypeConst(old.Y())
	if err != nil {
		return nil, errors.WithStack(err)
	}
	expr := constant.NewShl(x, y)
	expr.OverflowFlags = overflowFlags
	return expr, nil
}

// ~~~ [ lshr ] ~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~

// irLShrExpr translates the given AST lshr constant expression into an
// equivalent IR constant expression.
func (gen *generator) irLShrExpr(t types.Type, old *ast.LShrExpr) (*constant.ExprLShr, error) {
	// (optional) Exact.
	exact := old.Exact().IsValid()
	// X operand.
	x, err := gen.irTypeConst(old.X())
	if err != nil {
		return nil, errors.WithStack(err)
	}
	// Y operand.
	y, err := gen.irTypeConst(old.Y())
	if err != nil {
		return nil, errors.WithStack(err)
	}
	expr := constant.NewLShr(x, y)
	expr.Exact = exact
	return expr, nil
}

// ~~~ [ ashr ] ~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~

// irAShrExpr translates the given AST ashr constant expression into an
// equivalent IR constant expression.
func (gen *generator) irAShrExpr(t types.Type, old *ast.AShrExpr) (*constant.ExprAShr, error) {
	// (optional) Exact.
	exact := old.Exact().IsValid()
	// X operand.
	x, err := gen.irTypeConst(old.X())
	if err != nil {
		return nil, errors.WithStack(err)
	}
	// Y operand.
	y, err := gen.irTypeConst(old.Y())
	if err != nil {
		return nil, errors.WithStack(err)
	}
	expr := constant.NewAShr(x, y)
	expr.Exact = exact
	return expr, nil
}

// ~~~ [ and ] ~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~

// irAndExpr translates the given AST and constant expression into an equivalent
// IR constant expression.
func (gen *generator) irAndExpr(t types.Type, old *ast.AndExpr) (*constant.ExprAnd, error) {
	// X operand.
	x, err := gen.irTypeConst(old.X())
	if err != nil {
		return nil, errors.WithStack(err)
	}
	// Y operand.
	y, err := gen.irTypeConst(old.Y())
	if err != nil {
		return nil, errors.WithStack(err)
	}
	expr := constant.NewAnd(x, y)
	return expr, nil
}

// ~~~ [ or ] ~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~

// irOrExpr translates the given AST or constant expression into an equivalent
// IR constant expression.
func (gen *generator) irOrExpr(t types.Type, old *ast.OrExpr) (*constant.ExprOr, error) {
	// X operand.
	x, err := gen.irTypeConst(old.X())
	if err != nil {
		return nil, errors.WithStack(err)
	}
	// Y operand.
	y, err := gen.irTypeConst(old.Y())
	if err != nil {
		return nil, errors.WithStack(err)
	}
	expr := constant.NewOr(x, y)
	return expr, nil
}

// ~~~ [ xor ] ~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~

// irXorExpr translates the given AST xor constant expression into an equivalent
// IR constant expression.
func (gen *generator) irXorExpr(t types.Type, old *ast.XorExpr) (*constant.ExprXor, error) {
	// X operand.
	x, err := gen.irTypeConst(old.X())
	if err != nil {
		return nil, errors.WithStack(err)
	}
	// Y operand.
	y, err := gen.irTypeConst(old.Y())
	if err != nil {
		return nil, errors.WithStack(err)
	}
	expr := constant.NewXor(x, y)
	return expr, nil
}

// --- [ Vector expressions ] --------------------------------------------------

// ~~~ [ extractelement ] ~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~

// irExtractElementExpr translates the given AST extractelement constant
// expression into an equivalent IR constant expression.
func (gen *generator) irExtractElementExpr(t types.Type, old *ast.ExtractElementExpr) (*constant.ExprExtractElement, error) {
	// Vector.
	x, err := gen.irTypeConst(old.X())
	if err != nil {
		return nil, errors.WithStack(err)
	}
	// Index.
	index, err := gen.irTypeConst(old.Index())
	if err != nil {
		return nil, errors.WithStack(err)
	}
	expr := constant.NewExtractElement(x, index)
	return expr, nil
}

// ~~~ [ insertelement ] ~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~

// irInsertElementExpr translates the given AST insertelement constant
// expression into an equivalent IR constant expression.
func (gen *generator) irInsertElementExpr(t types.Type, old *ast.InsertElementExpr) (*constant.ExprInsertElement, error) {
	// Vector.
	x, err := gen.irTypeConst(old.X())
	if err != nil {
		return nil, errors.WithStack(err)
	}
	// Element.
	elem, err := gen.irTypeConst(old.Elem())
	if err != nil {
		return nil, errors.WithStack(err)
	}
	// Index.
	index, err := gen.irTypeConst(old.Index())
	if err != nil {
		return nil, errors.WithStack(err)
	}
	expr := constant.NewInsertElement(x, elem, index)
	return expr, nil
}

// ~~~ [ shufflevector ] ~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~

// irShuffleVectorExpr translates the given AST shufflevector constant
// expression into an equivalent IR constant expression.
func (gen *generator) irShuffleVectorExpr(t types.Type, old *ast.ShuffleVectorExpr) (*constant.ExprShuffleVector, error) {
	// X vector.
	x, err := gen.irTypeConst(old.X())
	if err != nil {
		return nil, errors.WithStack(err)
	}
	// X vector.
	y, err := gen.irTypeConst(old.Y())
	if err != nil {
		return nil, errors.WithStack(err)
	}
	// Shuffle mask.
	mask, err := gen.irTypeConst(old.Mask())
	if err != nil {
		return nil, errors.WithStack(err)
	}
	expr := constant.NewShuffleVector(x, y, mask)
	return expr, nil
}

// --- [ Aggregate expressions ] -----------------------------------------------

// ~~~ [ extractvalue ] ~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~

// irExtractValueExpr translates the given AST extractvalue constant expression
// into an equivalent IR constant expression.
func (gen *generator) irExtractValueExpr(t types.Type, old *ast.ExtractValueExpr) (*constant.ExprExtractValue, error) {
	// Aggregate value.
	x, err := gen.irTypeConst(old.X())
	if err != nil {
		return nil, errors.WithStack(err)
	}
	// Element indices.
	var indices []int64
	for _, index := range uintSlice(old.Indices()) {
		indices = append(indices, int64(index))
	}
	expr := constant.NewExtractValue(x, indices...)
	return expr, nil
}

// ~~~ [ insertvalue ] ~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~

// irInsertValueExpr translates the given AST insertvalue constant expression
// into an equivalent IR constant expression.
func (gen *generator) irInsertValueExpr(t types.Type, old *ast.InsertValueExpr) (*constant.ExprInsertValue, error) {
	// Aggregate value.
	x, err := gen.irTypeConst(old.X())
	if err != nil {
		return nil, errors.WithStack(err)
	}
	// Element.
	elem, err := gen.irTypeConst(old.Elem())
	if err != nil {
		return nil, errors.WithStack(err)
	}
	// Element indices.
	var indices []int64
	for _, index := range uintSlice(old.Indices()) {
		indices = append(indices, int64(index))
	}
	expr := constant.NewInsertValue(x, elem, indices...)
	return expr, nil
}

// --- [ Memory expressions ] --------------------------------------------------

// ~~~ [ getelementptr ] ~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~

// irGetElementPtrExpr translates the given AST getelementptr constant
// expression into an equivalent IR constant expression.
func (gen *generator) irGetElementPtrExpr(t types.Type, old *ast.GetElementPtrExpr) (*constant.ExprGetElementPtr, error) {
	// (optional) In-bounds.
	inBounds := old.InBounds().IsValid()
	// Element type.
	elemType, err := gen.irType(old.ElemType())
	if err != nil {
		return nil, errors.WithStack(err)
	}
	// Source.
	src, err := gen.irTypeConst(old.Src())
	if err != nil {
		return nil, errors.WithStack(err)
	}
	// Indices.
	var indices []*constant.Index
	for _, idx := range old.Indices() {
		index, err := gen.irGEPIndex(idx)
		if err != nil {
			return nil, errors.WithStack(err)
		}
		indices = append(indices, index)
	}
	expr := constant.NewGetElementPtr(src, indices...)
	_ = elemType
	// TODO: validate type elemType against expr.ElemType.
	// TODO: validate type t against expr.Typ.
	// (optional) In-bounds.
	expr.InBounds = inBounds
	return expr, nil
}

// irGEPIndex translates the given AST getelementptr index into an equivalent IR
// getelementptr index.
func (gen *generator) irGEPIndex(old ast.GEPIndex) (*constant.Index, error) {
	// (optional) In-range.
	inRange := old.InRange().IsValid()
	// Index.
	idx, err := gen.irTypeConst(old.Index())
	if err != nil {
		return nil, errors.WithStack(err)
	}
	index := constant.NewIndex(idx)
	index.InRange = inRange
	return index, nil
}

// --- [ Conversion expressions ] ----------------------------------------------

// ~~~ [ trunc ] ~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~

// irTruncExpr translates the given AST trunc constant expression into an
// equivalent IR constant expression.
func (gen *generator) irTruncExpr(t types.Type, old *ast.TruncExpr) (*constant.ExprTrunc, error) {
	// From.
	from, err := gen.irTypeConst(old.From())
	if err != nil {
		return nil, errors.WithStack(err)
	}
	// To.
	to, err := gen.irType(old.To())
	if err != nil {
		return nil, errors.WithStack(err)
	}
	expr := constant.NewTrunc(from, to)
	// TODO: validate type t against expr.Typ.
	return expr, nil
}

// ~~~ [ zext ] ~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~

// irZExtExpr translates the given AST zext constant expression into an
// equivalent IR constant expression.
func (gen *generator) irZExtExpr(t types.Type, old *ast.ZExtExpr) (*constant.ExprZExt, error) {
	// From.
	from, err := gen.irTypeConst(old.From())
	if err != nil {
		return nil, errors.WithStack(err)
	}
	// To.
	to, err := gen.irType(old.To())
	if err != nil {
		return nil, errors.WithStack(err)
	}
	expr := constant.NewZExt(from, to)
	// TODO: validate type t against expr.Typ.
	return expr, nil
}

// ~~~ [ sext ] ~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~

// irSExtExpr translates the given AST sext constant expression into an
// equivalent IR constant expression.
func (gen *generator) irSExtExpr(t types.Type, old *ast.SExtExpr) (*constant.ExprSExt, error) {
	// From.
	from, err := gen.irTypeConst(old.From())
	if err != nil {
		return nil, errors.WithStack(err)
	}
	// To.
	to, err := gen.irType(old.To())
	if err != nil {
		return nil, errors.WithStack(err)
	}
	expr := constant.NewSExt(from, to)
	// TODO: validate type t against expr.Typ.
	return expr, nil
}

// ~~~ [ fptrunc ] ~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~

// irFPTruncExpr translates the given AST fptrunc constant expression into an
// equivalent IR constant expression.
func (gen *generator) irFPTruncExpr(t types.Type, old *ast.FPTruncExpr) (*constant.ExprFPTrunc, error) {
	// From.
	from, err := gen.irTypeConst(old.From())
	if err != nil {
		return nil, errors.WithStack(err)
	}
	// To.
	to, err := gen.irType(old.To())
	if err != nil {
		return nil, errors.WithStack(err)
	}
	expr := constant.NewFPTrunc(from, to)
	// TODO: validate type t against expr.Typ.
	return expr, nil
}

// ~~~ [ fpext ] ~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~

// irFPExtExpr translates the given AST fpext constant expression into an
// equivalent IR constant expression.
func (gen *generator) irFPExtExpr(t types.Type, old *ast.FPExtExpr) (*constant.ExprFPExt, error) {
	// From.
	from, err := gen.irTypeConst(old.From())
	if err != nil {
		return nil, errors.WithStack(err)
	}
	// To.
	to, err := gen.irType(old.To())
	if err != nil {
		return nil, errors.WithStack(err)
	}
	expr := constant.NewFPExt(from, to)
	// TODO: validate type t against expr.Typ.
	return expr, nil
}

// ~~~ [ fptoui ] ~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~

// irFPToUIExpr translates the given AST fptoui constant expression into an
// equivalent IR constant expression.
func (gen *generator) irFPToUIExpr(t types.Type, old *ast.FPToUIExpr) (*constant.ExprFPToUI, error) {
	// From.
	from, err := gen.irTypeConst(old.From())
	if err != nil {
		return nil, errors.WithStack(err)
	}
	// To.
	to, err := gen.irType(old.To())
	if err != nil {
		return nil, errors.WithStack(err)
	}
	expr := constant.NewFPToUI(from, to)
	// TODO: validate type t against expr.Typ.
	return expr, nil
}

// ~~~ [ fptosi ] ~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~

// irFPToSIExpr translates the given AST fptosi constant expression into an
// equivalent IR constant expression.
func (gen *generator) irFPToSIExpr(t types.Type, old *ast.FPToSIExpr) (*constant.ExprFPToSI, error) {
	// From.
	from, err := gen.irTypeConst(old.From())
	if err != nil {
		return nil, errors.WithStack(err)
	}
	// To.
	to, err := gen.irType(old.To())
	if err != nil {
		return nil, errors.WithStack(err)
	}
	expr := constant.NewFPToSI(from, to)
	// TODO: validate type t against expr.Typ.
	return expr, nil
}

// ~~~ [ uitofp ] ~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~

// irUIToFPExpr translates the given AST uitofp constant expression into an
// equivalent IR constant expression.
func (gen *generator) irUIToFPExpr(t types.Type, old *ast.UIToFPExpr) (*constant.ExprUIToFP, error) {
	// From.
	from, err := gen.irTypeConst(old.From())
	if err != nil {
		return nil, errors.WithStack(err)
	}
	// To.
	to, err := gen.irType(old.To())
	if err != nil {
		return nil, errors.WithStack(err)
	}
	expr := constant.NewUIToFP(from, to)
	// TODO: validate type t against expr.Typ.
	return expr, nil
}

// ~~~ [ sitofp ] ~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~

// irSIToFPExpr translates the given AST sitofp constant expression into an
// equivalent IR constant expression.
func (gen *generator) irSIToFPExpr(t types.Type, old *ast.SIToFPExpr) (*constant.ExprSIToFP, error) {
	// From.
	from, err := gen.irTypeConst(old.From())
	if err != nil {
		return nil, errors.WithStack(err)
	}
	// To.
	to, err := gen.irType(old.To())
	if err != nil {
		return nil, errors.WithStack(err)
	}
	expr := constant.NewSIToFP(from, to)
	// TODO: validate type t against expr.Typ.
	return expr, nil
}

// ~~~ [ ptrtoint ] ~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~

// irPtrToIntExpr translates the given AST ptrtoint constant expression into an
// equivalent IR constant expression.
func (gen *generator) irPtrToIntExpr(t types.Type, old *ast.PtrToIntExpr) (*constant.ExprPtrToInt, error) {
	// From.
	from, err := gen.irTypeConst(old.From())
	if err != nil {
		return nil, errors.WithStack(err)
	}
	// To.
	to, err := gen.irType(old.To())
	if err != nil {
		return nil, errors.WithStack(err)
	}
	expr := constant.NewPtrToInt(from, to)
	// TODO: validate type t against expr.Typ.
	return expr, nil
}

// ~~~ [ inttoptr ] ~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~

// irIntToPtrExpr translates the given AST inttoptr constant expression into an
// equivalent IR constant expression.
func (gen *generator) irIntToPtrExpr(t types.Type, old *ast.IntToPtrExpr) (*constant.ExprIntToPtr, error) {
	// From.
	from, err := gen.irTypeConst(old.From())
	if err != nil {
		return nil, errors.WithStack(err)
	}
	// To.
	to, err := gen.irType(old.To())
	if err != nil {
		return nil, errors.WithStack(err)
	}
	expr := constant.NewIntToPtr(from, to)
	// TODO: validate type t against expr.Typ.
	return expr, nil
}

// ~~~ [ bitcast ] ~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~

// irBitCastExpr translates the given AST bitcast constant expression into an
// equivalent IR constant expression.
func (gen *generator) irBitCastExpr(t types.Type, old *ast.BitCastExpr) (*constant.ExprBitCast, error) {
	// From.
	from, err := gen.irTypeConst(old.From())
	if err != nil {
		return nil, errors.WithStack(err)
	}
	// To.
	to, err := gen.irType(old.To())
	if err != nil {
		return nil, errors.WithStack(err)
	}
	expr := constant.NewBitCast(from, to)
	// TODO: validate type t against expr.Typ.
	return expr, nil
}

// ~~~ [ addrspacecast ] ~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~

// irAddrSpaceCastExpr translates the given AST addrspacecast constant
// expression into an equivalent IR constant expression.
func (gen *generator) irAddrSpaceCastExpr(t types.Type, old *ast.AddrSpaceCastExpr) (*constant.ExprAddrSpaceCast, error) {
	// From.
	from, err := gen.irTypeConst(old.From())
	if err != nil {
		return nil, errors.WithStack(err)
	}
	// To.
	to, err := gen.irType(old.To())
	if err != nil {
		return nil, errors.WithStack(err)
	}
	expr := constant.NewAddrSpaceCast(from, to)
	// TODO: validate type t against expr.Typ.
	return expr, nil
}

// --- [ Other expressions ] ---------------------------------------------------

// ~~~ [ icmp ] ~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~

// irICmpExpr translates the given AST icmp constant expression into an
// equivalent IR constant expression.
func (gen *generator) irICmpExpr(t types.Type, old *ast.ICmpExpr) (*constant.ExprICmp, error) {
	// Integer comparison predicate.
	pred := asmenum.IPredFromString(old.Pred().Text())
	// X operand.
	x, err := gen.irTypeConst(old.X())
	if err != nil {
		return nil, errors.WithStack(err)
	}
	// Y operand.
	y, err := gen.irTypeConst(old.Y())
	if err != nil {
		return nil, errors.WithStack(err)
	}
	expr := constant.NewICmp(pred, x, y)
	return expr, nil
}

// ~~~ [ fcmp ] ~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~

// irFCmpExpr translates the given AST fcmp constant expression into an
// equivalent IR constant expression.
func (gen *generator) irFCmpExpr(t types.Type, old *ast.FCmpExpr) (*constant.ExprFCmp, error) {
	// Floating-point comparison predicate.
	pred := asmenum.FPredFromString(old.Pred().Text())
	// X operand.
	x, err := gen.irTypeConst(old.X())
	if err != nil {
		return nil, errors.WithStack(err)
	}
	// Y operand.
	y, err := gen.irTypeConst(old.Y())
	if err != nil {
		return nil, errors.WithStack(err)
	}
	expr := constant.NewFCmp(pred, x, y)
	return expr, nil
}

// ~~~ [ select ] ~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~

// irSelectExpr translates the given AST select constant expression into an
// equivalent IR constant expression.
func (gen *generator) irSelectExpr(t types.Type, old *ast.SelectExpr) (*constant.ExprSelect, error) {
	// Selection condition.
	cond, err := gen.irTypeConst(old.Cond())
	if err != nil {
		return nil, errors.WithStack(err)
	}
	// X operand.
	x, err := gen.irTypeConst(old.X())
	if err != nil {
		return nil, errors.WithStack(err)
	}
	// Y operand.
	y, err := gen.irTypeConst(old.Y())
	if err != nil {
		return nil, errors.WithStack(err)
	}
	expr := constant.NewSelect(cond, x, y)
	return expr, nil
}
