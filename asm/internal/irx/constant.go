package irx

import (
	"fmt"

	"github.com/llir/llvm/asm/internal/ast"
	"github.com/llir/llvm/ir"
	"github.com/llir/llvm/ir/constant"
	"github.com/pkg/errors"
)

// irConstant returns the corresponding LLVM IR constant of the given constant.
func (m *Module) irConstant(old ast.Constant) constant.Constant {
	switch old := old.(type) {
	// Simple constants
	case *ast.IntConst:
		return constant.NewIntFromString(old.Lit, m.irType(old.Type))
	case *ast.FloatConst:
		return constant.NewFloatFromString(old.Lit, m.irType(old.Type))
	case *ast.NullConst:
		return constant.NewNull(m.irType(old.Type))

	// Complex constants
	case *ast.VectorConst:
		var elems []constant.Constant
		for _, oldElem := range old.Elems {
			elems = append(elems, m.irConstant(oldElem))
		}
		c := constant.NewVector(elems...)
		if got, want := c.Type(), m.irType(old.Type); !got.Equal(want) {
			m.errs = append(m.errs, errors.Errorf("vector type mismatch; expected `%v`, got `%v`", want, got))
		}
		return c
	case *ast.ArrayConst:
		var elems []constant.Constant
		for _, oldElem := range old.Elems {
			elems = append(elems, m.irConstant(oldElem))
		}
		c := constant.NewArray(elems...)
		if got, want := c.Type(), m.irType(old.Type); !got.Equal(want) {
			m.errs = append(m.errs, errors.Errorf("array type mismatch; expected `%v`, got `%v`", want, got))
		}
		c.CharArray = old.CharArray
		return c
	case *ast.StructConst:
		var fields []constant.Constant
		for _, oldField := range old.Fields {
			fields = append(fields, m.irConstant(oldField))
		}
		c := constant.NewStruct(fields...)
		if got, want := c.Type(), m.irType(old.Type); !got.Equal(want) {
			m.errs = append(m.errs, errors.Errorf("struct type mismatch; expected `%v`, got `%v`", want, got))
		}
		return c
	case *ast.ZeroInitializerConst:
		return constant.NewZeroInitializer(m.irType(old.Type))

	// Global variable and function addresses
	case *ast.Global:
		// TODO: Validate old.Type against type of resolved global?
		// Not possible currently, as globals have already been resolved by astx.
		// Consider postponing global resolution until irx, so that
		// *ast.GlobalDummy.Type may be compared against global.Type.
		v := m.getGlobal(old.Name)
		global, ok := v.(*ir.Global)
		if !ok {
			panic(fmt.Errorf("invalid global type; expected *ir.Global, got %T", v))
		}
		return global
	case *ast.Function:
		// TODO: Validate old.Type against type of resolved function?
		// Not possible currently, as globals have already been resolved by astx.
		// Consider postponing global resolution until irx, so that
		// *ast.GlobalDummy.Type may be compared against f.Type.
		v := m.getGlobal(old.Name)
		f, ok := v.(*ir.Function)
		if !ok {
			panic(fmt.Errorf("invalid function type; expected *ir.Function, got %T", v))
		}
		return f

	// Binary expressions
	case *ast.ExprAdd:
		panic("not yet implemented")
	case *ast.ExprFAdd:
		panic("not yet implemented")
	case *ast.ExprSub:
		panic("not yet implemented")
	case *ast.ExprFSub:
		panic("not yet implemented")
	case *ast.ExprMul:
		panic("not yet implemented")
	case *ast.ExprFMul:
		panic("not yet implemented")
	case *ast.ExprUDiv:
		panic("not yet implemented")
	case *ast.ExprSDiv:
		panic("not yet implemented")
	case *ast.ExprFDiv:
		panic("not yet implemented")
	case *ast.ExprURem:
		panic("not yet implemented")
	case *ast.ExprSRem:
		panic("not yet implemented")
	case *ast.ExprFRem:
		panic("not yet implemented")
	// Bitwise expressions
	case *ast.ExprShl:
		panic("not yet implemented")
	case *ast.ExprLShr:
		panic("not yet implemented")
	case *ast.ExprAShr:
		panic("not yet implemented")
	case *ast.ExprAnd:
		panic("not yet implemented")
	case *ast.ExprOr:
		panic("not yet implemented")
	case *ast.ExprXor:
		panic("not yet implemented")

	// Memory expressions
	case *ast.ExprGetElementPtr:
		panic("not yet implemented")

	// Conversion expressions
	case *ast.ExprTrunc:
		panic("not yet implemented")
	case *ast.ExprZExt:
		panic("not yet implemented")
	case *ast.ExprSExt:
		panic("not yet implemented")
	case *ast.ExprFPTrunc:
		panic("not yet implemented")
	case *ast.ExprFPExt:
		panic("not yet implemented")
	case *ast.ExprFPToUI:
		panic("not yet implemented")
	case *ast.ExprFPToSI:
		panic("not yet implemented")
	case *ast.ExprUIToFP:
		panic("not yet implemented")
	case *ast.ExprSIToFP:
		panic("not yet implemented")
	case *ast.ExprPtrToInt:
		panic("not yet implemented")
	case *ast.ExprIntToPtr:
		panic("not yet implemented")
	case *ast.ExprBitCast:
		panic("not yet implemented")
	case *ast.ExprAddrSpaceCast:
		panic("not yet implemented")

	// Other expressions
	case *ast.ExprICmp:
		panic("not yet implemented")
	case *ast.ExprFCmp:
		panic("not yet implemented")
	case *ast.ExprSelect:
		panic("not yet implemented")

	default:
		panic(fmt.Errorf("support for %T not yet implemented", old))
	}
}
