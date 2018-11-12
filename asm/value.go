package asm

import (
	"fmt"

	"github.com/llir/ll/ast"
	"github.com/llir/llvm/ir"
	"github.com/llir/llvm/ir/types"
	"github.com/llir/llvm/ir/value"
	"github.com/pkg/errors"
)

func (fgen *funcGen) astToIRValue(typ types.Type, old ast.Value) (value.Value, error) {
	switch old := old.(type) {
	case *ast.GlobalIdent:
		ident := globalIdent(*old)
		v, ok := fgen.gen.new.globals[ident]
		if !ok {
			return nil, errors.Errorf("unable to locate global identifier %q", ident)
		}
		return v, nil
	case *ast.LocalIdent:
		ident := localIdent(*old)
		v, ok := fgen.ls[ident]
		if !ok {
			return nil, errors.Errorf("unable to locate local identifier %q", ident)
		}
		return v, nil
	case *ast.InlineAsm:
		asm := &ir.InlineAsm{}
		// (optional) Side effect.
		asm.SideEffect = old.SideEffect().IsValid()
		// (optional) Stack alignment.
		asm.AlignStack = old.AlignStackTok().IsValid()
		// (optional) Intel dialect.
		asm.IntelDialect = old.IntelDialect().IsValid()
		// Assembly instructions.
		asm.Asm = stringLit(old.Asm())
		// Constraints.
		asm.Constraint = stringLit(old.Constraints())
		return asm, nil
	case ast.Constant:
		return fgen.gen.irConstant(typ, old)
	default:
		panic(fmt.Errorf("support for AST value %T not yet implemented", old))
	}
}

func (fgen *funcGen) astToIRTypeValue(old ast.TypeValue) (value.Value, error) {
	// Type.
	typ, err := fgen.gen.irType(old.Typ())
	if err != nil {
		return nil, errors.WithStack(err)
	}
	// Value.
	return fgen.astToIRValue(typ, old.Val())
}
