package asm

import (
	"fmt"

	"github.com/llir/ll/ast"
	"github.com/llir/llvm/ir"
	"github.com/llir/llvm/ir/types"
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
		ident := localIdent(old.Name())
		return fgen.newIRValueTerm(ident, old.Term())
	case ast.ValueTerminator:
		return fgen.newIRValueTerm(ir.LocalIdent{}, old)
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
func (fgen *funcGen) newIRValueTerm(ident ir.LocalIdent, old ast.ValueTerminator) (ir.Terminator, error) {
	switch old := old.(type) {
	case *ast.InvokeTerm:
		// Invokee type.
		typ, err := fgen.gen.irType(old.Typ())
		if err != nil {
			return nil, errors.WithStack(err)
		}
		return &ir.TermInvoke{LocalIdent: ident, Typ: typ}, nil
	case *ast.CatchSwitchTerm:
		// Result type is always token.
		return &ir.TermCatchSwitch{LocalIdent: ident}, nil
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
		ident := localIdent(old.Name())
		v, ok := fgen.ls[ident]
		if !ok {
			return errors.Errorf("unable to locate local variable %q", ident.Ident())
		}
		t, ok := v.(ir.Terminator)
		if !ok {
			return errors.Errorf("invalid terminator type of %q; expected ir.Terminator, got %T", ident.Ident(), v)
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
	// Return type.
	typ, err := fgen.gen.irType(old.XTyp())
	if err != nil {
		return errors.WithStack(err)
	}
	// Check if not void return.
	if !typ.Equal(types.Void) {
		// Return value.
		x, err := fgen.astToIRValue(typ, old.X())
		if err != nil {
			return errors.WithStack(err)
		}
		t.X = x
	}
	// (optional) Metadata.
	md, err := fgen.gen.irMetadataAttachments(old.Metadata())
	if err != nil {
		return errors.WithStack(err)
	}
	t.Metadata = md
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
	// Target.
	target, err := fgen.irBasicBlock(old.Target())
	if err != nil {
		return errors.WithStack(err)
	}
	t.Target = target
	// (optional) Metadata.
	md, err := fgen.gen.irMetadataAttachments(old.Metadata())
	if err != nil {
		return errors.WithStack(err)
	}
	t.Metadata = md
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
	// Branching condition.
	ct := old.CondTyp()
	condType, err := fgen.gen.irType(&ct)
	if err != nil {
		return errors.WithStack(err)
	}
	cond, err := fgen.astToIRValue(condType, old.Cond())
	if err != nil {
		return errors.WithStack(err)
	}
	t.Cond = cond
	// Target true.
	targetTrue, err := fgen.irBasicBlock(old.TargetTrue())
	if err != nil {
		return errors.WithStack(err)
	}
	t.TargetTrue = targetTrue
	// Target false.
	targetFalse, err := fgen.irBasicBlock(old.TargetFalse())
	if err != nil {
		return errors.WithStack(err)
	}
	t.TargetFalse = targetFalse
	// (optional) Metadata.
	md, err := fgen.gen.irMetadataAttachments(old.Metadata())
	if err != nil {
		return errors.WithStack(err)
	}
	t.Metadata = md
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
	// Control variable.
	x, err := fgen.astToIRTypeValue(old.X())
	if err != nil {
		return errors.WithStack(err)
	}
	t.X = x
	// Default target.
	targetDefault, err := fgen.irBasicBlock(old.Default())
	if err != nil {
		return errors.WithStack(err)
	}
	t.TargetDefault = targetDefault
	// Switch cases.
	for _, oldCase := range old.Cases() {
		c, err := fgen.irCase(oldCase)
		if err != nil {
			return errors.WithStack(err)
		}
		t.Cases = append(t.Cases, c)
	}
	// (optional) Metadata.
	md, err := fgen.gen.irMetadataAttachments(old.Metadata())
	if err != nil {
		return errors.WithStack(err)
	}
	t.Metadata = md
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
	// Target address.
	addr, err := fgen.astToIRTypeValue(old.Addr())
	if err != nil {
		return errors.WithStack(err)
	}
	t.Addr = addr
	// Valid targets.
	for _, oldValidTarget := range old.ValidTargets() {
		validTarget, err := fgen.irBasicBlock(oldValidTarget)
		if err != nil {
			return errors.WithStack(err)
		}
		t.ValidTargets = append(t.ValidTargets, validTarget)
	}
	// (optional) Metadata.
	md, err := fgen.gen.irMetadataAttachments(old.Metadata())
	if err != nil {
		return errors.WithStack(err)
	}
	t.Metadata = md
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
	// (optional) Calling convention.
	if n := old.CallingConv(); n.LlvmNode().IsValid() {
		t.CallingConv = irCallingConv(n)
	}
	// (optional) Return attributes.
	for _, oldRetAttr := range old.ReturnAttrs() {
		retAttr := irReturnAttribute(oldRetAttr)
		t.ReturnAttrs = append(t.ReturnAttrs, retAttr)
	}
	// (optional) Address space.
	if n, ok := old.AddrSpace(); ok {
		t.AddrSpace = irAddrSpace(n)
	}
	// Invokee.
	typ, err := fgen.gen.irType(old.Typ())
	if err != nil {
		return errors.WithStack(err)
	}
	sig, ok := typ.(*types.FuncType)
	if !ok {
		// Preliminary function signature. Only used by astToIRValue for inline
		// assembly invokees and constrant expressions.
		sig = types.NewFunc(typ)
		// TODO: add parameters to sig.
	}
	invokee, err := fgen.astToIRValue(sig, old.Invokee())
	if err != nil {
		return errors.WithStack(err)
	}
	t.Invokee = invokee
	// Function arguments.
	for _, oldArg := range old.Args().Args() {
		arg, err := fgen.irArg(oldArg)
		if err != nil {
			return errors.WithStack(err)
		}
		t.Args = append(t.Args, arg)
	}
	// (optional) Function attributes.
	for _, oldFuncAttr := range old.FuncAttrs() {
		funcAttr := fgen.gen.irFuncAttribute(oldFuncAttr)
		t.FuncAttrs = append(t.FuncAttrs, funcAttr)
	}
	// (optional) Operand bundles.
	for _, oldOperandBundle := range old.OperandBundles() {
		operandBundle, err := fgen.irOperandBundle(oldOperandBundle)
		if err != nil {
			return errors.WithStack(err)
		}
		t.OperandBundles = append(t.OperandBundles, operandBundle)
	}
	// Normal control flow return point.
	normal, err := fgen.irBasicBlock(old.Normal())
	if err != nil {
		return errors.WithStack(err)
	}
	t.Normal = normal
	// Exception control flow return point.
	exception, err := fgen.irBasicBlock(old.Exception())
	if err != nil {
		return errors.WithStack(err)
	}
	t.Exception = exception
	// (optional) Metadata.
	md, err := fgen.gen.irMetadataAttachments(old.Metadata())
	if err != nil {
		return errors.WithStack(err)
	}
	t.Metadata = md
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
	// Exception argument to propagate.
	x, err := fgen.astToIRTypeValue(old.X())
	if err != nil {
		return errors.WithStack(err)
	}
	t.X = x
	// (optional) Metadata.
	md, err := fgen.gen.irMetadataAttachments(old.Metadata())
	if err != nil {
		return errors.WithStack(err)
	}
	t.Metadata = md
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
	// Exception scope.
	scope, err := fgen.irExceptionScope(old.Scope())
	if err != nil {
		return errors.WithStack(err)
	}
	t.Scope = scope
	// Exception handlers.
	for _, oldHandler := range old.Handlers().Labels() {
		handler, err := fgen.irBasicBlock(oldHandler)
		if err != nil {
			return errors.WithStack(err)
		}
		t.Handlers = append(t.Handlers, handler)
	}
	// Unwind target.
	unwindTarget, err := fgen.irUnwindTarget(old.UnwindTarget())
	if err != nil {
		return errors.WithStack(err)
	}
	t.UnwindTarget = unwindTarget
	// (optional) Metadata.
	md, err := fgen.gen.irMetadataAttachments(old.Metadata())
	if err != nil {
		return errors.WithStack(err)
	}
	t.Metadata = md
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
	// Exit catchpad.
	v, err := fgen.astToIRValue(types.Token, old.From())
	if err != nil {
		return errors.WithStack(err)
	}
	catchpad, ok := v.(*ir.InstCatchPad)
	if !ok {
		return errors.Errorf("invalid catchpad type; expected *ir.InstCatchPad, got %T", v)
	}
	t.From = catchpad
	// Target basic block to transfer control flow to.
	to, err := fgen.irBasicBlock(old.To())
	if err != nil {
		return errors.WithStack(err)
	}
	t.To = to
	// (optional) Metadata.
	md, err := fgen.gen.irMetadataAttachments(old.Metadata())
	if err != nil {
		return errors.WithStack(err)
	}
	t.Metadata = md
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
	// Exit cleanuppad.
	v, err := fgen.astToIRValue(types.Token, old.From())
	if err != nil {
		return errors.WithStack(err)
	}
	cleanuppad, ok := v.(*ir.InstCleanupPad)
	if !ok {
		return errors.Errorf("invalid cleanuppad type; expected *ir.InstCleanupPad, got %T", v)
	}
	t.From = cleanuppad
	// Unwind target.
	unwindTarget, err := fgen.irUnwindTarget(old.UnwindTarget())
	if err != nil {
		return errors.WithStack(err)
	}
	t.UnwindTarget = unwindTarget
	// (optional) Metadata.
	md, err := fgen.gen.irMetadataAttachments(old.Metadata())
	if err != nil {
		return errors.WithStack(err)
	}
	t.Metadata = md
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
	// (optional) Metadata.
	md, err := fgen.gen.irMetadataAttachments(old.Metadata())
	if err != nil {
		return errors.WithStack(err)
	}
	t.Metadata = md
	return nil
}
