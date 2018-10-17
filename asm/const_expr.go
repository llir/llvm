package asm

import (
	"fmt"

	"github.com/llir/l/ir"
	"github.com/llir/l/ir/types"
	"github.com/mewmew/l-tm/asm/ll/ast"
)

// === [ Constant expressions ] ================================================

func (gen *generator) irConstantExpr(t types.Type, old ast.ConstantExpr) (ir.Expression, error) {
	switch old := old.(type) {
	// Binary expressions
	case ast.AddExpr:
		return gen.irAddExpr(t, old)
	case ast.FAddExpr:
		return gen.irFAddExpr(t, old)
	case ast.SubExpr:
		return gen.irSubExpr(t, old)
	case ast.FSubExpr:
		return gen.irFSubExpr(t, old)
	case ast.MulExpr:
		return gen.irMulExpr(t, old)
	case ast.FMulExpr:
		return gen.irFMulExpr(t, old)
	case ast.UDivExpr:
		return gen.irUDivExpr(t, old)
	case ast.SDivExpr:
		return gen.irSDivExpr(t, old)
	case ast.FDivExpr:
		return gen.irFDivExpr(t, old)
	case ast.URemExpr:
		return gen.irURemExpr(t, old)
	case ast.SRemExpr:
		return gen.irSRemExpr(t, old)
	case ast.FRemExpr:
		return gen.irFRemExpr(t, old)
	// Bitwise expressions
	case ast.ShlExpr:
		return gen.irShlExpr(t, old)
	case ast.LShrExpr:
		return gen.irLShrExpr(t, old)
	case ast.AShrExpr:
		return gen.irAShrExpr(t, old)
	case ast.AndExpr:
		return gen.irAndExpr(t, old)
	case ast.OrExpr:
		return gen.irOrExpr(t, old)
	case ast.XorExpr:
		return gen.irXorExpr(t, old)
	// Vector expressions
	case ast.ExtractElementExpr:
		return gen.irExtractElementExpr(t, old)
	case ast.InsertElementExpr:
		return gen.irInsertElementExpr(t, old)
	case ast.ShuffleVectorExpr:
		return gen.irShuffleVectorExpr(t, old)
	// Aggregate expressions
	case ast.ExtractValueExpr:
		return gen.irExtractValueExpr(t, old)
	case ast.InsertValueExpr:
		return gen.irInsertValueExpr(t, old)
	// Memory expressions
	case ast.GetElementPtrExpr:
		return gen.irGetElementPtrExpr(t, old)
	// Conversion expressions
	case ast.TruncExpr:
		return gen.irTruncExpr(t, old)
	case ast.ZExtExpr:
		return gen.irZExtExpr(t, old)
	case ast.SExtExpr:
		return gen.irSExtExpr(t, old)
	case ast.FPTruncExpr:
		return gen.irFPTruncExpr(t, old)
	case ast.FPExtExpr:
		return gen.irFPExtExpr(t, old)
	case ast.FPToUIExpr:
		return gen.irFPToUIExpr(t, old)
	case ast.FPToSIExpr:
		return gen.irFPToSIExpr(t, old)
	case ast.UIToFPExpr:
		return gen.irUIToFPExpr(t, old)
	case ast.SIToFPExpr:
		return gen.irSIToFPExpr(t, old)
	case ast.PtrToIntExpr:
		return gen.irPtrToIntExpr(t, old)
	case ast.IntToPtrExpr:
		return gen.irIntToPtrExpr(t, old)
	case ast.BitCastExpr:
		return gen.irBitCastExpr(t, old)
	case ast.AddrSpaceCastExpr:
		return gen.irAddrSpaceCastExpr(t, old)
	// Other expressions
	case ast.ICmpExpr:
		return gen.irICmpExpr(t, old)
	case ast.FCmpExpr:
		return gen.irFCmpExpr(t, old)
	case ast.SelectExpr:
		return gen.irSelectExpr(t, old)
	default:
		panic(fmt.Errorf("support for AST constant expression %T not yet implemented", old))
	}
}

// --- [ Binary expressions ] -------------------------------------------------

// ~~~ [ add ] ~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~

func (gen *generator) irAddExpr(t types.Type, old ast.AddExpr) (*ir.ExprAdd, error) {
	panic("not yet implemented")
}

// ~~~ [ fadd ] ~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~

func (gen *generator) irFAddExpr(t types.Type, old ast.FAddExpr) (*ir.ExprFAdd, error) {
	panic("not yet implemented")
}

// ~~~ [ sub ] ~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~

func (gen *generator) irSubExpr(t types.Type, old ast.SubExpr) (*ir.ExprSub, error) {
	panic("not yet implemented")
}

// ~~~ [ fsub ] ~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~

func (gen *generator) irFSubExpr(t types.Type, old ast.FSubExpr) (*ir.ExprFSub, error) {
	panic("not yet implemented")
}

// ~~~ [ mul ] ~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~

func (gen *generator) irMulExpr(t types.Type, old ast.MulExpr) (*ir.ExprMul, error) {
	panic("not yet implemented")
}

// ~~~ [ fmul ] ~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~

func (gen *generator) irFMulExpr(t types.Type, old ast.FMulExpr) (*ir.ExprFMul, error) {
	panic("not yet implemented")
}

// ~~~ [ udiv ] ~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~

func (gen *generator) irUDivExpr(t types.Type, old ast.UDivExpr) (*ir.ExprUDiv, error) {
	panic("not yet implemented")
}

// ~~~ [ sdiv ] ~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~

func (gen *generator) irSDivExpr(t types.Type, old ast.SDivExpr) (*ir.ExprSDiv, error) {
	panic("not yet implemented")
}

// ~~~ [ fdiv ] ~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~

func (gen *generator) irFDivExpr(t types.Type, old ast.FDivExpr) (*ir.ExprFDiv, error) {
	panic("not yet implemented")
}

// ~~~ [ urem ] ~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~

func (gen *generator) irURemExpr(t types.Type, old ast.URemExpr) (*ir.ExprURem, error) {
	panic("not yet implemented")
}

// ~~~ [ srem ] ~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~

func (gen *generator) irSRemExpr(t types.Type, old ast.SRemExpr) (*ir.ExprSRem, error) {
	panic("not yet implemented")
}

// ~~~ [ frem ] ~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~

func (gen *generator) irFRemExpr(t types.Type, old ast.FRemExpr) (*ir.ExprFRem, error) {
	panic("not yet implemented")
}

// --- [ Bitwise expressions ] -------------------------------------------------

// ~~~ [ shl ] ~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~

func (gen *generator) irShlExpr(t types.Type, old ast.ShlExpr) (*ir.ExprShl, error) {
	panic("not yet implemented")
}

// ~~~ [ lshr ] ~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~

func (gen *generator) irLShrExpr(t types.Type, old ast.LShrExpr) (*ir.ExprLShr, error) {
	panic("not yet implemented")
}

// ~~~ [ ashr ] ~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~

func (gen *generator) irAShrExpr(t types.Type, old ast.AShrExpr) (*ir.ExprAShr, error) {
	panic("not yet implemented")
}

// ~~~ [ and ] ~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~

func (gen *generator) irAndExpr(t types.Type, old ast.AndExpr) (*ir.ExprAnd, error) {
	panic("not yet implemented")
}

// ~~~ [ or ] ~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~

func (gen *generator) irOrExpr(t types.Type, old ast.OrExpr) (*ir.ExprOr, error) {
	panic("not yet implemented")
}

// ~~~ [ xor ] ~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~

func (gen *generator) irXorExpr(t types.Type, old ast.XorExpr) (*ir.ExprXor, error) {
	panic("not yet implemented")
}

// --- [ Vector expressions ] --------------------------------------------------

// ~~~ [ extractelement ] ~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~

func (gen *generator) irExtractElementExpr(t types.Type, old ast.ExtractElementExpr) (*ir.ExprExtractElement, error) {
	panic("not yet implemented")
}

// ~~~ [ insertelement ] ~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~

func (gen *generator) irInsertElementExpr(t types.Type, old ast.InsertElementExpr) (*ir.ExprInsertElement, error) {
	panic("not yet implemented")
}

// ~~~ [ shufflevector ] ~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~

func (gen *generator) irShuffleVectorExpr(t types.Type, old ast.ShuffleVectorExpr) (*ir.ExprShuffleVector, error) {
	panic("not yet implemented")
}

// --- [ Aggregate expressions ] -----------------------------------------------

// ~~~ [ extractvalue ] ~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~

func (gen *generator) irExtractValueExpr(t types.Type, old ast.ExtractValueExpr) (*ir.ExprExtractValue, error) {
	panic("not yet implemented")
}

// ~~~ [ insertvalue ] ~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~

func (gen *generator) irInsertValueExpr(t types.Type, old ast.InsertValueExpr) (*ir.ExprInsertValue, error) {
	panic("not yet implemented")
}

// --- [ Memory expressions ] --------------------------------------------------

// ~~~ [ getelementptr ] ~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~

func (gen *generator) irGetElementPtrExpr(t types.Type, old ast.GetElementPtrExpr) (*ir.ExprGetElementPtr, error) {
	panic("not yet implemented")
}

// --- [ Conversion expressions ] ----------------------------------------------

// ~~~ [ trunc ] ~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~

func (gen *generator) irTruncExpr(t types.Type, old ast.TruncExpr) (*ir.ExprTrunc, error) {
	panic("not yet implemented")
}

// ~~~ [ zext ] ~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~

func (gen *generator) irZExtExpr(t types.Type, old ast.ZExtExpr) (*ir.ExprZExt, error) {
	panic("not yet implemented")
}

// ~~~ [ sext ] ~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~

func (gen *generator) irSExtExpr(t types.Type, old ast.SExtExpr) (*ir.ExprSExt, error) {
	panic("not yet implemented")
}

// ~~~ [ fptrunc ] ~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~

func (gen *generator) irFPTruncExpr(t types.Type, old ast.FPTruncExpr) (*ir.ExprFPTrunc, error) {
	panic("not yet implemented")
}

// ~~~ [ fpext ] ~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~

func (gen *generator) irFPExtExpr(t types.Type, old ast.FPExtExpr) (*ir.ExprFPExt, error) {
	panic("not yet implemented")
}

// ~~~ [ fptoui ] ~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~

func (gen *generator) irFPToUIExpr(t types.Type, old ast.FPToUIExpr) (*ir.ExprFPToUI, error) {
	panic("not yet implemented")
}

// ~~~ [ fptosi ] ~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~

func (gen *generator) irFPToSIExpr(t types.Type, old ast.FPToSIExpr) (*ir.ExprFPToSI, error) {
	panic("not yet implemented")
}

// ~~~ [ uitofp ] ~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~

func (gen *generator) irUIToFPExpr(t types.Type, old ast.UIToFPExpr) (*ir.ExprUIToFP, error) {
	panic("not yet implemented")
}

// ~~~ [ sitofp ] ~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~

func (gen *generator) irSIToFPExpr(t types.Type, old ast.SIToFPExpr) (*ir.ExprSIToFP, error) {
	panic("not yet implemented")
}

// ~~~ [ ptrtoint ] ~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~

func (gen *generator) irPtrToIntExpr(t types.Type, old ast.PtrToIntExpr) (*ir.ExprPtrToInt, error) {
	panic("not yet implemented")
}

// ~~~ [ inttoptr ] ~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~

func (gen *generator) irIntToPtrExpr(t types.Type, old ast.IntToPtrExpr) (*ir.ExprIntToPtr, error) {
	panic("not yet implemented")
}

// ~~~ [ bitcast ] ~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~

func (gen *generator) irBitCastExpr(t types.Type, old ast.BitCastExpr) (*ir.ExprBitCast, error) {
	panic("not yet implemented")
}

// ~~~ [ addrspacecast ] ~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~

func (gen *generator) irAddrSpaceCastExpr(t types.Type, old ast.AddrSpaceCastExpr) (*ir.ExprAddrSpaceCast, error) {
	panic("not yet implemented")
}

// --- [ Other expressions ] ---------------------------------------------------

// ~~~ [ icmp ] ~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~

func (gen *generator) irICmpExpr(t types.Type, old ast.ICmpExpr) (*ir.ExprICmp, error) {
	panic("not yet implemented")
}

// ~~~ [ fcmp ] ~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~

func (gen *generator) irFCmpExpr(t types.Type, old ast.FCmpExpr) (*ir.ExprFCmp, error) {
	panic("not yet implemented")
}

// ~~~ [ select ] ~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~

func (gen *generator) irSelectExpr(t types.Type, old ast.SelectExpr) (*ir.ExprSelect, error) {
	panic("not yet implemented")
}
