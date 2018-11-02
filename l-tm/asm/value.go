package asm

import (
	"fmt"

	"github.com/llir/l/ir/types"
	"github.com/llir/l/ir/value"
	"github.com/mewmew/l-tm/asm/ll/ast"
	"github.com/pkg/errors"
)

func (fgen *funcGen) astToIRValue(typ types.Type, old ast.Value) (value.Value, error) {
	switch old := old.(type) {
	case *ast.GlobalIdent:
		name := global(*old)
		v, ok := fgen.gen.gs[name]
		if !ok {
			return nil, errors.Errorf("unable to locate global identifier %q", name)
		}
		return v, nil
	case *ast.LocalIdent:
		name := local(*old)
		v, ok := fgen.ls[name]
		if !ok {
			return nil, errors.Errorf("unable to locate local identifier %q", name)
		}
		return v, nil
	case *ast.InlineAsm:
		// TODO: implement
		panic("not yet implemented")
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
