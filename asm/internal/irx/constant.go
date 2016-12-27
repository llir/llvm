package irx

import (
	"fmt"

	"github.com/llir/llvm/asm/internal/ast"
	"github.com/llir/llvm/ir"
	"github.com/llir/llvm/ir/constant"
	"github.com/llir/llvm/ir/types"
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
		return c
	case *ast.CharArrayConst:
		var elems []constant.Constant
		for i := 0; i < len(old.Lit); i++ {
			b := int64(old.Lit[i])
			elem := constant.NewInt(b, types.I8)
			elems = append(elems, elem)
		}
		c := constant.NewArray(elems...)
		if got, want := c.Type(), m.irType(old.Type); !got.Equal(want) {
			m.errs = append(m.errs, errors.Errorf("character array type mismatch; expected `%v`, got `%v`", want, got))
		}
		c.CharArray = true
		return c
	case *ast.StructConst:
		var fields []constant.Constant
		for _, oldField := range old.Fields {
			fields = append(fields, m.irConstant(oldField))
		}
		c := constant.NewStruct(fields...)
		if got, want := c.Type(), m.irType(old.Type); !got.Equal(want) {
			err := errors.Errorf("struct type mismatch; expected `%v`, got `%v`", want, got)
			m.errs = append(m.errs, err)
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
		x, y := m.irConstant(old.X), m.irConstant(old.Y)
		c := constant.NewAdd(x, y)
		if got, want := c.Type(), m.irType(old.Type); !got.Equal(want) {
			m.errs = append(m.errs, errors.Errorf("add expression type mismatch; expected `%v`, got `%v`", want, got))
		}
		return c
	case *ast.ExprFAdd:
		x, y := m.irConstant(old.X), m.irConstant(old.Y)
		c := constant.NewFAdd(x, y)
		if got, want := c.Type(), m.irType(old.Type); !got.Equal(want) {
			m.errs = append(m.errs, errors.Errorf("fadd expression type mismatch; expected `%v`, got `%v`", want, got))
		}
		return c
	case *ast.ExprSub:
		x, y := m.irConstant(old.X), m.irConstant(old.Y)
		c := constant.NewSub(x, y)
		if got, want := c.Type(), m.irType(old.Type); !got.Equal(want) {
			m.errs = append(m.errs, errors.Errorf("sub expression type mismatch; expected `%v`, got `%v`", want, got))
		}
		return c
	case *ast.ExprFSub:
		x, y := m.irConstant(old.X), m.irConstant(old.Y)
		c := constant.NewFSub(x, y)
		if got, want := c.Type(), m.irType(old.Type); !got.Equal(want) {
			m.errs = append(m.errs, errors.Errorf("fsub expression type mismatch; expected `%v`, got `%v`", want, got))
		}
		return c
	case *ast.ExprMul:
		x, y := m.irConstant(old.X), m.irConstant(old.Y)
		c := constant.NewMul(x, y)
		if got, want := c.Type(), m.irType(old.Type); !got.Equal(want) {
			m.errs = append(m.errs, errors.Errorf("mul expression type mismatch; expected `%v`, got `%v`", want, got))
		}
		return c
	case *ast.ExprFMul:
		x, y := m.irConstant(old.X), m.irConstant(old.Y)
		c := constant.NewFMul(x, y)
		if got, want := c.Type(), m.irType(old.Type); !got.Equal(want) {
			m.errs = append(m.errs, errors.Errorf("fmul expression type mismatch; expected `%v`, got `%v`", want, got))
		}
		return c
	case *ast.ExprUDiv:
		x, y := m.irConstant(old.X), m.irConstant(old.Y)
		c := constant.NewUDiv(x, y)
		if got, want := c.Type(), m.irType(old.Type); !got.Equal(want) {
			m.errs = append(m.errs, errors.Errorf("udiv expression type mismatch; expected `%v`, got `%v`", want, got))
		}
		return c
	case *ast.ExprSDiv:
		x, y := m.irConstant(old.X), m.irConstant(old.Y)
		c := constant.NewSDiv(x, y)
		if got, want := c.Type(), m.irType(old.Type); !got.Equal(want) {
			m.errs = append(m.errs, errors.Errorf("sdiv expression type mismatch; expected `%v`, got `%v`", want, got))
		}
		return c
	case *ast.ExprFDiv:
		x, y := m.irConstant(old.X), m.irConstant(old.Y)
		c := constant.NewFDiv(x, y)
		if got, want := c.Type(), m.irType(old.Type); !got.Equal(want) {
			m.errs = append(m.errs, errors.Errorf("fdiv expression type mismatch; expected `%v`, got `%v`", want, got))
		}
		return c
	case *ast.ExprURem:
		x, y := m.irConstant(old.X), m.irConstant(old.Y)
		c := constant.NewURem(x, y)
		if got, want := c.Type(), m.irType(old.Type); !got.Equal(want) {
			m.errs = append(m.errs, errors.Errorf("urem expression type mismatch; expected `%v`, got `%v`", want, got))
		}
		return c
	case *ast.ExprSRem:
		x, y := m.irConstant(old.X), m.irConstant(old.Y)
		c := constant.NewSRem(x, y)
		if got, want := c.Type(), m.irType(old.Type); !got.Equal(want) {
			m.errs = append(m.errs, errors.Errorf("srem expression type mismatch; expected `%v`, got `%v`", want, got))
		}
		return c
	case *ast.ExprFRem:
		x, y := m.irConstant(old.X), m.irConstant(old.Y)
		c := constant.NewFRem(x, y)
		if got, want := c.Type(), m.irType(old.Type); !got.Equal(want) {
			m.errs = append(m.errs, errors.Errorf("frem expression type mismatch; expected `%v`, got `%v`", want, got))
		}
		return c

	// Bitwise expressions
	case *ast.ExprShl:
		x, y := m.irConstant(old.X), m.irConstant(old.Y)
		c := constant.NewShl(x, y)
		if got, want := c.Type(), m.irType(old.Type); !got.Equal(want) {
			m.errs = append(m.errs, errors.Errorf("shl expression type mismatch; expected `%v`, got `%v`", want, got))
		}
		return c
	case *ast.ExprLShr:
		x, y := m.irConstant(old.X), m.irConstant(old.Y)
		c := constant.NewLShr(x, y)
		if got, want := c.Type(), m.irType(old.Type); !got.Equal(want) {
			m.errs = append(m.errs, errors.Errorf("lshr expression type mismatch; expected `%v`, got `%v`", want, got))
		}
		return c
	case *ast.ExprAShr:
		x, y := m.irConstant(old.X), m.irConstant(old.Y)
		c := constant.NewAShr(x, y)
		if got, want := c.Type(), m.irType(old.Type); !got.Equal(want) {
			m.errs = append(m.errs, errors.Errorf("ashr expression type mismatch; expected `%v`, got `%v`", want, got))
		}
		return c
	case *ast.ExprAnd:
		x, y := m.irConstant(old.X), m.irConstant(old.Y)
		c := constant.NewAnd(x, y)
		if got, want := c.Type(), m.irType(old.Type); !got.Equal(want) {
			m.errs = append(m.errs, errors.Errorf("and expression type mismatch; expected `%v`, got `%v`", want, got))
		}
		return c
	case *ast.ExprOr:
		x, y := m.irConstant(old.X), m.irConstant(old.Y)
		c := constant.NewOr(x, y)
		if got, want := c.Type(), m.irType(old.Type); !got.Equal(want) {
			m.errs = append(m.errs, errors.Errorf("or expression type mismatch; expected `%v`, got `%v`", want, got))
		}
		return c
	case *ast.ExprXor:
		x, y := m.irConstant(old.X), m.irConstant(old.Y)
		c := constant.NewXor(x, y)
		if got, want := c.Type(), m.irType(old.Type); !got.Equal(want) {
			m.errs = append(m.errs, errors.Errorf("xor expression type mismatch; expected `%v`, got `%v`", want, got))
		}
		return c

	// Memory expressions
	case *ast.ExprGetElementPtr:
		src := m.irConstant(old.Src)
		if srcType, ok := src.Type().(*types.PointerType); !ok {
			m.errs = append(m.errs, errors.Errorf("invalid source type; expected *types.PointerType, got %T", src.Type()))
		} else if got, want := srcType.Elem, m.irType(old.Elem); !got.Equal(want) {
			m.errs = append(m.errs, errors.Errorf("source element type mismatch; expected `%v`, got `%v`", want, got))
		}
		var indices []constant.Constant
		for _, oldIndex := range old.Indices {
			index := m.irConstant(oldIndex)
			indices = append(indices, index)
		}
		c := constant.NewGetElementPtr(src, indices...)
		if got, want := c.Type(), m.irType(old.Type); !got.Equal(want) {
			m.errs = append(m.errs, errors.Errorf("getelementptr expression type mismatch; expected `%v`, got `%v`", want, got))
		}
		return c

	// Conversion expressions
	case *ast.ExprTrunc:
		from := m.irConstant(old.From)
		to := m.irType(old.To)
		c := constant.NewTrunc(from, to)
		if got, want := c.Type(), m.irType(old.Type); !got.Equal(want) {
			m.errs = append(m.errs, errors.Errorf("trunc expression type mismatch; expected `%v`, got `%v`", want, got))
		}
		return c
	case *ast.ExprZExt:
		from := m.irConstant(old.From)
		to := m.irType(old.To)
		c := constant.NewZExt(from, to)
		if got, want := c.Type(), m.irType(old.Type); !got.Equal(want) {
			m.errs = append(m.errs, errors.Errorf("zext expression type mismatch; expected `%v`, got `%v`", want, got))
		}
		return c
	case *ast.ExprSExt:
		from := m.irConstant(old.From)
		to := m.irType(old.To)
		c := constant.NewSExt(from, to)
		if got, want := c.Type(), m.irType(old.Type); !got.Equal(want) {
			m.errs = append(m.errs, errors.Errorf("sext expression type mismatch; expected `%v`, got `%v`", want, got))
		}
		return c
	case *ast.ExprFPTrunc:
		from := m.irConstant(old.From)
		to := m.irType(old.To)
		c := constant.NewFPTrunc(from, to)
		if got, want := c.Type(), m.irType(old.Type); !got.Equal(want) {
			m.errs = append(m.errs, errors.Errorf("fptrunc expression type mismatch; expected `%v`, got `%v`", want, got))
		}
		return c
	case *ast.ExprFPExt:
		from := m.irConstant(old.From)
		to := m.irType(old.To)
		c := constant.NewFPExt(from, to)
		if got, want := c.Type(), m.irType(old.Type); !got.Equal(want) {
			m.errs = append(m.errs, errors.Errorf("fpext expression type mismatch; expected `%v`, got `%v`", want, got))
		}
		return c
	case *ast.ExprFPToUI:
		from := m.irConstant(old.From)
		to := m.irType(old.To)
		c := constant.NewFPToUI(from, to)
		if got, want := c.Type(), m.irType(old.Type); !got.Equal(want) {
			m.errs = append(m.errs, errors.Errorf("fptoui expression type mismatch; expected `%v`, got `%v`", want, got))
		}
		return c
	case *ast.ExprFPToSI:
		from := m.irConstant(old.From)
		to := m.irType(old.To)
		c := constant.NewFPToSI(from, to)
		if got, want := c.Type(), m.irType(old.Type); !got.Equal(want) {
			m.errs = append(m.errs, errors.Errorf("fptosi expression type mismatch; expected `%v`, got `%v`", want, got))
		}
		return c
	case *ast.ExprUIToFP:
		from := m.irConstant(old.From)
		to := m.irType(old.To)
		c := constant.NewUIToFP(from, to)
		if got, want := c.Type(), m.irType(old.Type); !got.Equal(want) {
			m.errs = append(m.errs, errors.Errorf("uitofp expression type mismatch; expected `%v`, got `%v`", want, got))
		}
		return c
	case *ast.ExprSIToFP:
		from := m.irConstant(old.From)
		to := m.irType(old.To)
		c := constant.NewSIToFP(from, to)
		if got, want := c.Type(), m.irType(old.Type); !got.Equal(want) {
			m.errs = append(m.errs, errors.Errorf("sitofp expression type mismatch; expected `%v`, got `%v`", want, got))
		}
		return c
	case *ast.ExprPtrToInt:
		from := m.irConstant(old.From)
		to := m.irType(old.To)
		c := constant.NewPtrToInt(from, to)
		if got, want := c.Type(), m.irType(old.Type); !got.Equal(want) {
			m.errs = append(m.errs, errors.Errorf("ptrtoint expression type mismatch; expected `%v`, got `%v`", want, got))
		}
		return c
	case *ast.ExprIntToPtr:
		from := m.irConstant(old.From)
		to := m.irType(old.To)
		c := constant.NewIntToPtr(from, to)
		if got, want := c.Type(), m.irType(old.Type); !got.Equal(want) {
			m.errs = append(m.errs, errors.Errorf("inttoptr expression type mismatch; expected `%v`, got `%v`", want, got))
		}
		return c
	case *ast.ExprBitCast:
		from := m.irConstant(old.From)
		to := m.irType(old.To)
		c := constant.NewBitCast(from, to)
		if got, want := c.Type(), m.irType(old.Type); !got.Equal(want) {
			m.errs = append(m.errs, errors.Errorf("bitcast expression type mismatch; expected `%v`, got `%v`", want, got))
		}
		return c
	case *ast.ExprAddrSpaceCast:
		from := m.irConstant(old.From)
		to := m.irType(old.To)
		c := constant.NewAddrSpaceCast(from, to)
		if got, want := c.Type(), m.irType(old.Type); !got.Equal(want) {
			m.errs = append(m.errs, errors.Errorf("addrspacecast expression type mismatch; expected `%v`, got `%v`", want, got))
		}
		return c

	// Other expressions
	case *ast.ExprICmp:
		cond := constant.IntPred(irIntPred(old.Cond))
		x, y := m.irConstant(old.X), m.irConstant(old.Y)
		c := constant.NewICmp(cond, x, y)
		if got, want := c.Type(), m.irType(old.Type); !got.Equal(want) {
			m.errs = append(m.errs, errors.Errorf("icmp expression type mismatch; expected `%v`, got `%v`", want, got))
		}
		return c
	case *ast.ExprFCmp:
		cond := constant.FloatPred(irFloatPred(old.Cond))
		x, y := m.irConstant(old.X), m.irConstant(old.Y)
		c := constant.NewFCmp(cond, x, y)
		if got, want := c.Type(), m.irType(old.Type); !got.Equal(want) {
			m.errs = append(m.errs, errors.Errorf("fcmp expression type mismatch; expected `%v`, got `%v`", want, got))
		}
		return c
	case *ast.ExprSelect:
		cond := m.irConstant(old.Cond)
		x, y := m.irConstant(old.X), m.irConstant(old.Y)
		c := constant.NewSelect(cond, x, y)
		if got, want := c.Type(), m.irType(old.Type); !got.Equal(want) {
			m.errs = append(m.errs, errors.Errorf("select expression type mismatch; expected `%v`, got `%v`", want, got))
		}
		return c

	default:
		panic(fmt.Errorf("support for constant %T not yet implemented", old))
	}
}
