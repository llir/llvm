package asm

import (
	"fmt"

	"github.com/llir/ll/ast"
	asmenum "github.com/llir/llvm/asm/enum"
	"github.com/llir/llvm/ir/constant"
	"github.com/llir/llvm/ir/types"
	"github.com/pkg/errors"
)

// === [ Translate AST to IR ] =================================================

// irConstantExpr translates the AST constant expression into an equivalent IR
// constant expression.
func (gen *generator) irConstantExpr(t types.Type, old ast.ConstantExpr) (constant.Expression, error) {
	switch old := old.(type) {
	// Unary expressions
	case *ast.FNegExpr:
		return gen.irFNegExpr(t, old)
	// Binary expressions
	case *ast.AddExpr:
		return gen.irAddExpr(t, old)
	case *ast.SubExpr:
		return gen.irSubExpr(t, old)
	case *ast.MulExpr:
		return gen.irMulExpr(t, old)
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

// --- [ Unary expressions ] ---------------------------------------------------

// ~~~ [ fneg ] ~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~

// irFNegExpr translates the AST fneg constant expression into an equivalent IR
// constant expression.
func (gen *generator) irFNegExpr(t types.Type, old *ast.FNegExpr) (*constant.ExprFNeg, error) {
	// X operand.
	x, err := gen.irTypeConst(old.X())
	if err != nil {
		return nil, errors.WithStack(err)
	}
	expr := constant.NewFNeg(x)
	if !t.Equal(expr.Typ) {
		return nil, errors.Errorf("constant expression type mismatch in `%v`; expected %q, got %q", expr, expr.Typ, t)
	}
	return expr, nil
}

// --- [ Binary expressions ] --------------------------------------------------

// ~~~ [ add ] ~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~

// irAddExpr translates the AST add constant expression into an equivalent IR
// constant expression.
func (gen *generator) irAddExpr(t types.Type, old *ast.AddExpr) (*constant.ExprAdd, error) {
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
	// (optional) Overflow flags.
	expr.OverflowFlags = irOverflowFlags(old.OverflowFlags())
	if !t.Equal(expr.Typ) {
		return nil, errors.Errorf("constant expression type mismatch in `%v`; expected %q, got %q", expr, expr.Typ, t)
	}
	return expr, nil
}

// ~~~ [ sub ] ~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~

// irSubExpr translates the AST sub constant expression into an equivalent IR
// constant expression.
func (gen *generator) irSubExpr(t types.Type, old *ast.SubExpr) (*constant.ExprSub, error) {
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
	// (optional) Overflow flags.
	expr.OverflowFlags = irOverflowFlags(old.OverflowFlags())
	if !t.Equal(expr.Typ) {
		return nil, errors.Errorf("constant expression type mismatch in `%v`; expected %q, got %q", expr, expr.Typ, t)
	}
	return expr, nil
}

// ~~~ [ mul ] ~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~

// irMulExpr translates the AST mul constant expression into an equivalent IR
// constant expression.
func (gen *generator) irMulExpr(t types.Type, old *ast.MulExpr) (*constant.ExprMul, error) {
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
	// (optional) Overflow flags.
	expr.OverflowFlags = irOverflowFlags(old.OverflowFlags())
	if !t.Equal(expr.Typ) {
		return nil, errors.Errorf("constant expression type mismatch in `%v`; expected %q, got %q", expr, expr.Typ, t)
	}
	return expr, nil
}

// --- [ Bitwise expressions ] -------------------------------------------------

// ~~~ [ shl ] ~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~

// irShlExpr translates the AST shl constant expression into an equivalent IR
// constant expression.
func (gen *generator) irShlExpr(t types.Type, old *ast.ShlExpr) (*constant.ExprShl, error) {
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
	// (optional) Overflow flags.
	expr.OverflowFlags = irOverflowFlags(old.OverflowFlags())
	if !t.Equal(expr.Typ) {
		return nil, errors.Errorf("constant expression type mismatch in `%v`; expected %q, got %q", expr, expr.Typ, t)
	}
	return expr, nil
}

// ~~~ [ lshr ] ~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~

// irLShrExpr translates the AST lshr constant expression into an equivalent IR
// constant expression.
func (gen *generator) irLShrExpr(t types.Type, old *ast.LShrExpr) (*constant.ExprLShr, error) {
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
	// (optional) Exact.
	_, expr.Exact = old.Exact()
	if !t.Equal(expr.Typ) {
		return nil, errors.Errorf("constant expression type mismatch in `%v`; expected %q, got %q", expr, expr.Typ, t)
	}
	return expr, nil
}

// ~~~ [ ashr ] ~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~

// irAShrExpr translates the AST ashr constant expression into an equivalent IR
// constant expression.
func (gen *generator) irAShrExpr(t types.Type, old *ast.AShrExpr) (*constant.ExprAShr, error) {
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
	// (optional) Exact.
	_, expr.Exact = old.Exact()
	if !t.Equal(expr.Typ) {
		return nil, errors.Errorf("constant expression type mismatch in `%v`; expected %q, got %q", expr, expr.Typ, t)
	}
	return expr, nil
}

// ~~~ [ and ] ~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~

// irAndExpr translates the AST and constant expression into an equivalent IR
// constant expression.
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
	if !t.Equal(expr.Typ) {
		return nil, errors.Errorf("constant expression type mismatch in `%v`; expected %q, got %q", expr, expr.Typ, t)
	}
	return expr, nil
}

// ~~~ [ or ] ~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~

// irOrExpr translates the AST or constant expression into an equivalent IR
// constant expression.
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
	if !t.Equal(expr.Typ) {
		return nil, errors.Errorf("constant expression type mismatch in `%v`; expected %q, got %q", expr, expr.Typ, t)
	}
	return expr, nil
}

// ~~~ [ xor ] ~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~

// irXorExpr translates the AST xor constant expression into an equivalent IR
// constant expression.
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
	if !t.Equal(expr.Typ) {
		return nil, errors.Errorf("constant expression type mismatch in `%v`; expected %q, got %q", expr, expr.Typ, t)
	}
	return expr, nil
}

// --- [ Vector expressions ] --------------------------------------------------

// ~~~ [ extractelement ] ~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~

// irExtractElementExpr translates the AST extractelement constant expression
// into an equivalent IR constant expression.
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
	if !t.Equal(expr.Typ) {
		return nil, errors.Errorf("constant expression type mismatch in `%v`; expected %q, got %q", expr, expr.Typ, t)
	}
	return expr, nil
}

// ~~~ [ insertelement ] ~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~

// irInsertElementExpr translates the AST insertelement constant expression into
// an equivalent IR constant expression.
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
	if !t.Equal(expr.Typ) {
		return nil, errors.Errorf("constant expression type mismatch in `%v`; expected %q, got %q", expr, expr.Typ, t)
	}
	return expr, nil
}

// ~~~ [ shufflevector ] ~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~

// irShuffleVectorExpr translates the AST shufflevector constant expression into
// an equivalent IR constant expression.
func (gen *generator) irShuffleVectorExpr(t types.Type, old *ast.ShuffleVectorExpr) (*constant.ExprShuffleVector, error) {
	// X vector.
	x, err := gen.irTypeConst(old.X())
	if err != nil {
		return nil, errors.WithStack(err)
	}
	// Y vector.
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
	if !t.Equal(expr.Typ) {
		return nil, errors.Errorf("constant expression type mismatch in `%v`; expected %q, got %q", expr, expr.Typ, t)
	}
	return expr, nil
}

// --- [ Memory expressions ] --------------------------------------------------

// ~~~ [ getelementptr ] ~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~

// irGetElementPtrExpr translates the AST getelementptr constant expression into
// an equivalent IR constant expression.
func (gen *generator) irGetElementPtrExpr(t types.Type, old *ast.GetElementPtrExpr) (*constant.ExprGetElementPtr, error) {
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
	var indices []constant.Constant
	if oldIndices := old.Indices(); len(oldIndices) > 0 {
		indices = make([]constant.Constant, len(oldIndices))
		for i, oldIndex := range oldIndices {
			index, err := gen.irGEPIndex(oldIndex)
			if err != nil {
				return nil, errors.WithStack(err)
			}
			indices[i] = index
		}
	}
	expr := constant.NewGetElementPtr(elemType, src, indices...)
	// (optional) In-bounds.
	_, expr.InBounds = old.InBounds()
	if !elemType.Equal(expr.ElemType) {
		return nil, errors.Errorf("constant expression element type mismatch; expected %q, got %q", expr.ElemType, elemType)
	}
	if !t.Equal(expr.Typ) {
		return nil, errors.Errorf("constant expression type mismatch of `%v`; expected %q, got %q", expr, expr.Typ, t)
	}
	return expr, nil
}

// irGEPIndex translates the AST getelementptr index into an equivalent IR
// getelementptr index.
func (gen *generator) irGEPIndex(old ast.GEPIndex) (*constant.Index, error) {
	// Index.
	idx, err := gen.irTypeConst(old.Index())
	if err != nil {
		return nil, errors.WithStack(err)
	}
	index := constant.NewIndex(idx)
	// (optional) In-range.
	_, index.InRange = old.InRange()
	return index, nil
}

// --- [ Conversion expressions ] ----------------------------------------------

// ~~~ [ trunc ] ~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~

// irTruncExpr translates the AST trunc constant expression into an equivalent
// IR constant expression.
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
	if !t.Equal(expr.To) {
		return nil, errors.Errorf("constant expression type mismatch; expected %q, got %q", expr.To, t)
	}
	return expr, nil
}

// ~~~ [ zext ] ~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~

// irZExtExpr translates the AST zext constant expression into an equivalent IR
// constant expression.
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
	if !t.Equal(expr.To) {
		return nil, errors.Errorf("constant expression type mismatch; expected %q, got %q", expr.To, t)
	}
	return expr, nil
}

// ~~~ [ sext ] ~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~

// irSExtExpr translates the AST sext constant expression into an equivalent IR
// constant expression.
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
	if !t.Equal(expr.To) {
		return nil, errors.Errorf("constant expression type mismatch; expected %q, got %q", expr.To, t)
	}
	return expr, nil
}

// ~~~ [ fptrunc ] ~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~

// irFPTruncExpr translates the AST fptrunc constant expression into an
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
	if !t.Equal(expr.To) {
		return nil, errors.Errorf("constant expression type mismatch; expected %q, got %q", expr.To, t)
	}
	return expr, nil
}

// ~~~ [ fpext ] ~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~

// irFPExtExpr translates the AST fpext constant expression into an equivalent
// IR constant expression.
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
	if !t.Equal(expr.To) {
		return nil, errors.Errorf("constant expression type mismatch; expected %q, got %q", expr.To, t)
	}
	return expr, nil
}

// ~~~ [ fptoui ] ~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~

// irFPToUIExpr translates the AST fptoui constant expression into an equivalent
// IR constant expression.
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
	if !t.Equal(expr.To) {
		return nil, errors.Errorf("constant expression type mismatch; expected %q, got %q", expr.To, t)
	}
	return expr, nil
}

// ~~~ [ fptosi ] ~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~

// irFPToSIExpr translates the AST fptosi constant expression into an equivalent
// IR constant expression.
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
	if !t.Equal(expr.To) {
		return nil, errors.Errorf("constant expression type mismatch; expected %q, got %q", expr.To, t)
	}
	return expr, nil
}

// ~~~ [ uitofp ] ~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~

// irUIToFPExpr translates the AST uitofp constant expression into an equivalent
// IR constant expression.
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
	if !t.Equal(expr.To) {
		return nil, errors.Errorf("constant expression type mismatch; expected %q, got %q", expr.To, t)
	}
	return expr, nil
}

// ~~~ [ sitofp ] ~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~

// irSIToFPExpr translates the AST sitofp constant expression into an equivalent
// IR constant expression.
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
	if !t.Equal(expr.To) {
		return nil, errors.Errorf("constant expression type mismatch; expected %q, got %q", expr.To, t)
	}
	return expr, nil
}

// ~~~ [ ptrtoint ] ~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~

// irPtrToIntExpr translates the AST ptrtoint constant expression into an
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
	if !t.Equal(expr.To) {
		return nil, errors.Errorf("constant expression type mismatch; expected %q, got %q", expr.To, t)
	}
	return expr, nil
}

// ~~~ [ inttoptr ] ~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~

// irIntToPtrExpr translates the AST inttoptr constant expression into an
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
	if !t.Equal(expr.To) {
		return nil, errors.Errorf("constant expression type mismatch; expected %q, got %q", expr.To, t)
	}
	return expr, nil
}

// ~~~ [ bitcast ] ~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~

// irBitCastExpr translates the AST bitcast constant expression into an
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
	if !t.Equal(expr.To) {
		return nil, errors.Errorf("constant expression type mismatch; expected %q, got %q", expr.To, t)
	}
	return expr, nil
}

// ~~~ [ addrspacecast ] ~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~

// irAddrSpaceCastExpr translates the AST addrspacecast constant expression into
// an equivalent IR constant expression.
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
	if !t.Equal(expr.To) {
		return nil, errors.Errorf("constant expression type mismatch; expected %q, got %q", expr.To, t)
	}
	return expr, nil
}

// --- [ Other expressions ] ---------------------------------------------------

// ~~~ [ icmp ] ~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~

// irICmpExpr translates the AST icmp constant expression into an equivalent IR
// constant expression.
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
	if !t.Equal(expr.Typ) {
		return nil, errors.Errorf("constant expression type mismatch in `%v`; expected %q, got %q", expr, expr.Typ, t)
	}
	return expr, nil
}

// ~~~ [ fcmp ] ~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~

// irFCmpExpr translates the AST fcmp constant expression into an equivalent IR
// constant expression.
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
	if !t.Equal(expr.Typ) {
		return nil, errors.Errorf("constant expression type mismatch in `%v`; expected %q, got %q", expr, expr.Typ, t)
	}
	return expr, nil
}

// ~~~ [ select ] ~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~

// irSelectExpr translates the AST select constant expression into an equivalent
// IR constant expression.
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
	if !t.Equal(expr.Typ) {
		return nil, errors.Errorf("constant expression type mismatch in `%v`; expected %q, got %q", expr, expr.Typ, t)
	}
	return expr, nil
}
