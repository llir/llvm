package asm

import (
	"fmt"

	"github.com/llir/l/ir"
	"github.com/mewmew/l-tm/asm/ll/ast"
	"github.com/mewmew/l-tm/internal/enc"
	"github.com/pkg/errors"
)

// === [ Terminators ] =========================================================

// +++ [ Index ] +++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++

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
	// Non-value terminators.
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
		// Invokee type.
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

// +++ [ Translate ] +++++++++++++++++++++++++++++++++++++++++++++++++++++++++++

// astToIRTerm translates the given AST terminator into an equivalent IR
// terminator.
func (fgen *funcGen) astToIRTerm(term ir.Terminator, old ast.Terminator) error {
	switch old := old.(type) {
	// Value terminators.
	case *ast.LocalDefTerm:
		name := local(old.Name())
		v, ok := fgen.ls[name]
		if !ok {
			return errors.Errorf("unable to locate local variable %q", enc.Local(name))
		}
		t, ok := v.(ir.Terminator)
		if !ok {
			return errors.Errorf("invalid terminator type of %q; expected ir.Terminator, got %T", enc.Local(name), v)
		}
		return fgen.astToIRValueTerm(t, old.Term())
	case ast.ValueTerminator:
		return fgen.astToIRValueTerm(term, old)
	// Non-value terminators.
	case *ast.RetTerm:
		return fgen.astToIRTermRet(term, old)
	case *ast.BrTerm:
		return fgen.astToIRTermBr(term, old)
	case *ast.CondBrTerm:
		return fgen.astToIRTermCondBr(term, old)
	case *ast.SwitchTerm:
		return fgen.astToIRTermSwitch(term, old)
	case *ast.IndirectBrTerm:
		return fgen.astToIRTermIndirectBr(term, old)
	case *ast.ResumeTerm:
		return fgen.astToIRTermResume(term, old)
	case *ast.CatchRetTerm:
		return fgen.astToIRTermCatchRet(term, old)
	case *ast.CleanupRetTerm:
		return fgen.astToIRTermCleanupRet(term, old)
	case *ast.UnreachableTerm:
		return fgen.astToIRTermUnreachable(term, old)
	default:
		panic(fmt.Errorf("support for terminator %T not yet implemented", old))
	}
}

// astToIRValueTerm translates the given AST value terminator into an equivalent
// IR terminator.
func (fgen *funcGen) astToIRValueTerm(term ir.Terminator, old ast.ValueTerminator) error {
	switch old := old.(type) {
	case *ast.InvokeTerm:
		return fgen.astToIRTermInvoke(term, old)
	case *ast.CatchSwitchTerm:
		return fgen.astToIRTermCatchSwitch(term, old)
	default:
		panic(fmt.Errorf("support for value terminator %T not yet implemented", old))
	}
}

// --- [ ret ] -----------------------------------------------------------------

// astToIRTermRet translates the given AST ret terminator into an equivalent IR
// terminator.
func (fgen *funcGen) astToIRTermRet(term ir.Terminator, old *ast.RetTerm) error {
	t, ok := term.(*ir.TermRet)
	if !ok {
		panic(fmt.Errorf("invalid IR terminator for AST terminator; expected *ir.TermRet, got %T", term))
	}
	// TODO: implement.
	_ = t
	return nil
}

// --- [ br ] ------------------------------------------------------------------

// astToIRTermBr translates the given AST br terminator into an equivalent IR
// terminator.
func (fgen *funcGen) astToIRTermBr(term ir.Terminator, old *ast.BrTerm) error {
	t, ok := term.(*ir.TermBr)
	if !ok {
		panic(fmt.Errorf("invalid IR terminator for AST terminator; expected *ir.TermBr, got %T", term))
	}
	// TODO: implement.
	_ = t
	return nil
}

// --- [ condbr ] --------------------------------------------------------------

// astToIRTermCondBr translates the given AST condbr terminator into an
// equivalent IR terminator.
func (fgen *funcGen) astToIRTermCondBr(term ir.Terminator, old *ast.CondBrTerm) error {
	t, ok := term.(*ir.TermCondBr)
	if !ok {
		panic(fmt.Errorf("invalid IR terminator for AST terminator; expected *ir.TermCondBr, got %T", term))
	}
	// TODO: implement.
	_ = t
	return nil
}

// --- [ switch ] --------------------------------------------------------------

// astToIRTermSwitch translates the given AST switch terminator into an
// equivalent IR terminator.
func (fgen *funcGen) astToIRTermSwitch(term ir.Terminator, old *ast.SwitchTerm) error {
	t, ok := term.(*ir.TermSwitch)
	if !ok {
		panic(fmt.Errorf("invalid IR terminator for AST terminator; expected *ir.TermSwitch, got %T", term))
	}
	// TODO: implement.
	_ = t
	return nil
}

// --- [ indirectbr ] ----------------------------------------------------------

// astToIRTermIndirectBr translates the given AST indirectbr terminator into an
// equivalent IR terminator.
func (fgen *funcGen) astToIRTermIndirectBr(term ir.Terminator, old *ast.IndirectBrTerm) error {
	t, ok := term.(*ir.TermIndirectBr)
	if !ok {
		panic(fmt.Errorf("invalid IR terminator for AST terminator; expected *ir.TermIndirectBr, got %T", term))
	}
	// TODO: implement.
	_ = t
	return nil
}

// --- [ invoke ] --------------------------------------------------------------

// astToIRTermInvoke translates the given AST invoke terminator into an
// equivalent IR terminator.
func (fgen *funcGen) astToIRTermInvoke(term ir.Terminator, old *ast.InvokeTerm) error {
	t, ok := term.(*ir.TermInvoke)
	if !ok {
		panic(fmt.Errorf("invalid IR terminator for AST terminator; expected *ir.TermInvoke, got %T", term))
	}
	// TODO: implement.
	_ = t
	return nil
}

// --- [ resume ] --------------------------------------------------------------

// astToIRTermResume translates the given AST resume terminator into an
// equivalent IR terminator.
func (fgen *funcGen) astToIRTermResume(term ir.Terminator, old *ast.ResumeTerm) error {
	t, ok := term.(*ir.TermResume)
	if !ok {
		panic(fmt.Errorf("invalid IR terminator for AST terminator; expected *ir.TermResume, got %T", term))
	}
	// TODO: implement.
	_ = t
	return nil
}

// --- [ catchswitch ] ---------------------------------------------------------

// astToIRTermCatchSwitch translates the given AST catchswitch terminator into
// an equivalent IR terminator.
func (fgen *funcGen) astToIRTermCatchSwitch(term ir.Terminator, old *ast.CatchSwitchTerm) error {
	t, ok := term.(*ir.TermCatchSwitch)
	if !ok {
		panic(fmt.Errorf("invalid IR terminator for AST terminator; expected *ir.TermCatchSwitch, got %T", term))
	}
	// TODO: implement.
	_ = t
	return nil
}

// --- [ catchret ] ------------------------------------------------------------

// astToIRTermCatchRet translates the given AST catchret terminator into an
// equivalent IR terminator.
func (fgen *funcGen) astToIRTermCatchRet(term ir.Terminator, old *ast.CatchRetTerm) error {
	t, ok := term.(*ir.TermCatchRet)
	if !ok {
		panic(fmt.Errorf("invalid IR terminator for AST terminator; expected *ir.TermCatchRet, got %T", term))
	}
	// TODO: implement.
	_ = t
	return nil
}

// --- [ cleanupret ] ----------------------------------------------------------

// astToIRTermCleanupRet translates the given AST cleanupret terminator into an
// equivalent IR terminator.
func (fgen *funcGen) astToIRTermCleanupRet(term ir.Terminator, old *ast.CleanupRetTerm) error {
	t, ok := term.(*ir.TermCleanupRet)
	if !ok {
		panic(fmt.Errorf("invalid IR terminator for AST terminator; expected *ir.TermCleanupRet, got %T", term))
	}
	// TODO: implement.
	_ = t
	return nil
}

// --- [ unreachable ] ---------------------------------------------------------

// astToIRTermUnreachable translates the given AST unreachable terminator into
// an equivalent IR terminator.
func (fgen *funcGen) astToIRTermUnreachable(term ir.Terminator, old *ast.UnreachableTerm) error {
	t, ok := term.(*ir.TermUnreachable)
	if !ok {
		panic(fmt.Errorf("invalid IR terminator for AST terminator; expected *ir.TermUnreachable, got %T", term))
	}
	// TODO: implement.
	_ = t
	return nil
}
