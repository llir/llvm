package irx

import (
	"fmt"

	"github.com/llir/llvm/asm/internal/ast"
	"github.com/llir/llvm/ir"
)

// irIntPred returns the corresponding LLVM IR integer predicate of the given
// integer predicate.
func irIntPred(cond ast.IntPred) ir.IntPred {
	switch cond {
	case ast.IntEQ:
		return ir.IntEQ
	case ast.IntNE:
		return ir.IntNE
	case ast.IntUGT:
		return ir.IntUGT
	case ast.IntUGE:
		return ir.IntUGE
	case ast.IntULT:
		return ir.IntULT
	case ast.IntULE:
		return ir.IntULE
	case ast.IntSGT:
		return ir.IntSGT
	case ast.IntSGE:
		return ir.IntSGE
	case ast.IntSLT:
		return ir.IntSLT
	case ast.IntSLE:
		return ir.IntSLE
	}
	panic(fmt.Errorf("support for integer predicate %v not yet implemented", cond))
}

// irFloatPred returns the corresponding LLVM IR floating-point predicate of the
// given floating-point predicate.
func irFloatPred(cond ast.FloatPred) ir.FloatPred {
	switch cond {
	case ast.FloatFalse:
		return ir.FloatFalse
	case ast.FloatOEQ:
		return ir.FloatOEQ
	case ast.FloatOGT:
		return ir.FloatOGT
	case ast.FloatOGE:
		return ir.FloatOGE
	case ast.FloatOLT:
		return ir.FloatOLT
	case ast.FloatOLE:
		return ir.FloatOLE
	case ast.FloatONE:
		return ir.FloatONE
	case ast.FloatORD:
		return ir.FloatORD
	case ast.FloatUEQ:
		return ir.FloatUEQ
	case ast.FloatUGT:
		return ir.FloatUGT
	case ast.FloatUGE:
		return ir.FloatUGE
	case ast.FloatULT:
		return ir.FloatULT
	case ast.FloatULE:
		return ir.FloatULE
	case ast.FloatUNE:
		return ir.FloatUNE
	case ast.FloatUNO:
		return ir.FloatUNO
	case ast.FloatTrue:
		return ir.FloatTrue
	}
	panic(fmt.Errorf("support for floating-point predicate %v not yet implemented", cond))
}
