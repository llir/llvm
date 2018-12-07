package asm

import (
	"fmt"

	"github.com/llir/ll/ast"
	"github.com/llir/llvm/ir/types"
	"github.com/llir/llvm/ir/value"
	"github.com/pkg/errors"
)

// irValue translates the AST value into an equivalent IR value.
func (fgen *funcGen) irValue(typ types.Type, old ast.Value) (value.Value, error) {
	switch old := old.(type) {
	case *ast.GlobalIdent:
		ident := globalIdent(*old)
		v, ok := fgen.gen.new.globals[ident]
		if !ok {
			return nil, errors.Errorf("unable to locate global identifier %q", ident.Ident())
		}
		return v, nil
	case *ast.LocalIdent:
		ident := localIdent(*old)
		v, ok := fgen.locals[ident]
		if !ok {
			return nil, errors.Errorf("unable to locate local identifier %q", ident.Ident())
		}
		return v, nil
	case *ast.InlineAsm:
		return irInlineAsm(typ, old), nil
	case ast.Constant:
		return fgen.gen.irConstant(typ, old)
	default:
		panic(fmt.Errorf("support for AST value %T not yet implemented", old))
	}
}

// irTypeConst translates the AST type-value pair into an equivalent IR value.
func (fgen *funcGen) irTypeValue(old ast.TypeValue) (value.Value, error) {
	// Type.
	typ, err := fgen.gen.irType(old.Typ())
	if err != nil {
		return nil, errors.WithStack(err)
	}
	// Value.
	return fgen.irValue(typ, old.Val())
}
