package asm

import (
	"fmt"

	"github.com/llir/l/ir"
	"github.com/mewmew/l-tm/asm/ll/ast"
	"github.com/pkg/errors"
)

// newIRTerm returns a new IR terminator (with type and without body) based on
// the given AST terminator.
func (fgen *funcGen) newIRTerm(old ast.Terminator) (ir.Terminator, error) {
	switch old := old.(type) {
	// Value terminators.
	case *ast.LocalDefTerm:
		name := local(old.Name())
		return fgen.newIRValueTerm(name, old.Term())
	case ast.ValueTerminator:
		return fgen.newIRValueTerm("", old)
	// Non-value instructions.
	case *ast.RetTerm:
		return &ir.TermRet{}, nil
	case *ast.BrTerm:
		return &ir.TermBr{}, nil
	case *ast.CondBrTerm:
		return &ir.TermCondBr{}, nil
	case *ast.SwitchTerm:
		return &ir.TermSwitch{}, nil
	case *ast.IndirectBrTerm:
		return &ir.TermIndirectBr{}, nil
	case *ast.ResumeTerm:
		return &ir.TermResume{}, nil
	case *ast.CatchRetTerm:
		return &ir.TermCatchRet{}, nil
	case *ast.CleanupRetTerm:
		return &ir.TermCleanupRet{}, nil
	case *ast.UnreachableTerm:
		return &ir.TermUnreachable{}, nil
	default:
		panic(fmt.Errorf("support for terminator %T not yet implemented", old))
	}
}

// newIRValueTerm returns a new IR value terminator (with type and without body)
// based on the given AST value terminator.
func (fgen *funcGen) newIRValueTerm(name string, old ast.ValueTerminator) (ir.Terminator, error) {
	switch old := old.(type) {
	case *ast.InvokeTerm:
		typ, err := fgen.gen.irType(old.Typ())
		if err != nil {
			return nil, errors.WithStack(err)
		}
		return &ir.TermInvoke{LocalName: name, Typ: typ}, nil
	case *ast.CatchSwitchTerm:
		return &ir.TermCatchSwitch{LocalName: name}, nil
	default:
		panic(fmt.Errorf("support for value terminator %T not yet implemented", old))
	}
}
