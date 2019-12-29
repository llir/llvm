package asm

import (
	"fmt"

	"github.com/llir/ll/ast"
	"github.com/llir/llvm/ir"
	"github.com/llir/llvm/ir/types"
	"github.com/llir/llvm/ir/value"
	"github.com/pkg/errors"
)

// === [ Create IR ] ===========================================================

// newTerm returns a new IR terminator (without body but with type) based on the
// given AST terminator.
func (fgen *funcGen) newTerm(old ast.Terminator) (ir.Terminator, error) {
	switch old := old.(type) {
	// Value terminators.
	case *ast.LocalDefTerm:
		ident := localIdent(old.Name())
		return fgen.newValueTerm(ident, old.Term())
	case ast.ValueTerminator:
		unnamed := ir.LocalIdent{}
		return fgen.newValueTerm(unnamed, old)
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

// newValueTerm returns a new IR value terminator (without body but with type)
// based on the given AST value terminator.
func (fgen *funcGen) newValueTerm(ident ir.LocalIdent, old ast.ValueTerminator) (ir.Terminator, error) {
	switch old := old.(type) {
	case *ast.InvokeTerm:
		return fgen.newInvokeTerm(ident, old)
	case *ast.CallBrTerm:
		return fgen.newCallBrTerm(ident, old)
	case *ast.CatchSwitchTerm:
		// Result type is always token.
		return &ir.TermCatchSwitch{LocalIdent: ident}, nil
	default:
		panic(fmt.Errorf("support for value terminator %T not yet implemented", old))
	}
}

// newInvokeTerm returns a new IR invoke terminator (without body but with type)
// based on the given AST invoke terminator.
func (fgen *funcGen) newInvokeTerm(ident ir.LocalIdent, old *ast.InvokeTerm) (*ir.TermInvoke, error) {
	typ, err := fgen.gen.irType(old.Typ())
	if err != nil {
		return nil, errors.WithStack(err)
	}
	// Resolve return type of variadic functions.
	if funcType, ok := typ.(*types.FuncType); ok {
		typ = funcType.RetType
	}
	return &ir.TermInvoke{LocalIdent: ident, Typ: typ}, nil
}

// newCallBrTerm returns a new IR callbr terminator (without body but with type)
// based on the given AST callbr terminator.
func (fgen *funcGen) newCallBrTerm(ident ir.LocalIdent, old *ast.CallBrTerm) (*ir.TermCallBr, error) {
	typ, err := fgen.gen.irType(old.Typ())
	if err != nil {
		return nil, errors.WithStack(err)
	}
	// Resolve return type of variadic functions.
	if funcType, ok := typ.(*types.FuncType); ok {
		typ = funcType.RetType
	}
	return &ir.TermCallBr{LocalIdent: ident, Typ: typ}, nil
}

// === [ Translate AST to IR ] =================================================

// translateTerms translates the AST terminators of the given function to IR.
func (fgen *funcGen) translateTerms(oldBlocks []ast.BasicBlock) error {
	for i, oldBlock := range oldBlocks {
		block := fgen.f.Blocks[i]
		old := oldBlock.Term()
		if err := fgen.irTerm(block.Term, old); err != nil {
			return errors.WithStack(err)
		}
	}
	return nil
}

// irTerm translates the AST terminator into an equivalent IR terminator.
func (fgen *funcGen) irTerm(new ir.Terminator, old ast.Terminator) error {
	switch old := old.(type) {
	// Value terminators.
	case *ast.LocalDefTerm:
		return fgen.irValueTerm(new, old.Term())
	case ast.ValueTerminator:
		return fgen.irValueTerm(new, old)
	// Non-value terminators.
	case *ast.RetTerm:
		return fgen.irRetTerm(new, old)
	case *ast.BrTerm:
		return fgen.irBrTerm(new, old)
	case *ast.CondBrTerm:
		return fgen.irCondBrTerm(new, old)
	case *ast.SwitchTerm:
		return fgen.irSwitchTerm(new, old)
	case *ast.IndirectBrTerm:
		return fgen.irIndirectBrTerm(new, old)
	case *ast.ResumeTerm:
		return fgen.irResumeTerm(new, old)
	case *ast.CatchRetTerm:
		return fgen.irCatchRetTerm(new, old)
	case *ast.CleanupRetTerm:
		return fgen.irCleanupRetTerm(new, old)
	case *ast.UnreachableTerm:
		return fgen.irUnreachableTerm(new, old)
	default:
		panic(fmt.Errorf("support for terminator %T not yet implemented", old))
	}
}

// irValueTerm translates the AST value terminator into an equivalent IR value
// terminator.
func (fgen *funcGen) irValueTerm(new ir.Terminator, old ast.ValueTerminator) error {
	switch old := old.(type) {
	case *ast.InvokeTerm:
		return fgen.irInvokeTerm(new, old)
	case *ast.CallBrTerm:
		return fgen.irCallBrTerm(new, old)
	case *ast.CatchSwitchTerm:
		return fgen.irCatchSwitchTerm(new, old)
	default:
		panic(fmt.Errorf("support for value terminator %T not yet implemented", old))
	}
}

// --- [ ret ] -----------------------------------------------------------------

// irRetTerm translates the AST ret terminator into an equivalent IR terminator.
func (fgen *funcGen) irRetTerm(new ir.Terminator, old *ast.RetTerm) error {
	term, ok := new.(*ir.TermRet)
	if !ok {
		panic(fmt.Errorf("invalid IR terminator for AST terminator; expected *ir.TermRet, got %T", new))
	}
	// Return type.
	typ, err := fgen.gen.irType(old.XTyp())
	if err != nil {
		return errors.WithStack(err)
	}
	// Check if non-void return.
	if n, ok := old.X(); ok {
		// Return value.
		x, err := fgen.irValue(typ, n)
		if err != nil {
			return errors.WithStack(err)
		}
		term.X = x
	}
	// (optional) Metadata.
	md, err := fgen.gen.irMetadataAttachments(old.Metadata())
	if err != nil {
		return errors.WithStack(err)
	}
	term.Metadata = md
	return nil
}

// --- [ br ] ------------------------------------------------------------------

// irBrTerm translates the AST br terminator into an equivalent IR terminator.
func (fgen *funcGen) irBrTerm(new ir.Terminator, old *ast.BrTerm) error {
	term, ok := new.(*ir.TermBr)
	if !ok {
		panic(fmt.Errorf("invalid IR terminator for AST terminator; expected *ir.TermBr, got %T", new))
	}
	// Target.
	target, err := fgen.irBlock(old.Target())
	if err != nil {
		return errors.WithStack(err)
	}
	term.Target = target
	// (optional) Metadata.
	md, err := fgen.gen.irMetadataAttachments(old.Metadata())
	if err != nil {
		return errors.WithStack(err)
	}
	term.Metadata = md
	return nil
}

// --- [ condbr ] --------------------------------------------------------------

// irCondBrTerm translates the AST condbr terminator into an equivalent IR
// terminator.
func (fgen *funcGen) irCondBrTerm(new ir.Terminator, old *ast.CondBrTerm) error {
	term, ok := new.(*ir.TermCondBr)
	if !ok {
		panic(fmt.Errorf("invalid IR terminator for AST terminator; expected *ir.TermCondBr, got %T", new))
	}
	// Branching condition.
	ct := old.CondTyp()
	condType, err := fgen.gen.irType(&ct)
	if err != nil {
		return errors.WithStack(err)
	}
	cond, err := fgen.irValue(condType, old.Cond())
	if err != nil {
		return errors.WithStack(err)
	}
	term.Cond = cond
	// Target true.
	targetTrue, err := fgen.irBlock(old.TargetTrue())
	if err != nil {
		return errors.WithStack(err)
	}
	term.TargetTrue = targetTrue
	// Target false.
	targetFalse, err := fgen.irBlock(old.TargetFalse())
	if err != nil {
		return errors.WithStack(err)
	}
	term.TargetFalse = targetFalse
	// (optional) Metadata.
	md, err := fgen.gen.irMetadataAttachments(old.Metadata())
	if err != nil {
		return errors.WithStack(err)
	}
	term.Metadata = md
	return nil
}

// --- [ switch ] --------------------------------------------------------------

// irSwitchTerm translates the AST switch terminator into an equivalent IR
// terminator.
func (fgen *funcGen) irSwitchTerm(new ir.Terminator, old *ast.SwitchTerm) error {
	term, ok := new.(*ir.TermSwitch)
	if !ok {
		panic(fmt.Errorf("invalid IR terminator for AST terminator; expected *ir.TermSwitch, got %T", new))
	}
	// Control variable.
	x, err := fgen.irTypeValue(old.X())
	if err != nil {
		return errors.WithStack(err)
	}
	term.X = x
	// Default target.
	targetDefault, err := fgen.irBlock(old.Default())
	if err != nil {
		return errors.WithStack(err)
	}
	term.TargetDefault = targetDefault
	// Switch cases.
	if oldCases := old.Cases(); len(oldCases) > 0 {
		term.Cases = make([]*ir.Case, len(oldCases))
		for i, oldCase := range oldCases {
			c, err := fgen.irCase(oldCase)
			if err != nil {
				return errors.WithStack(err)
			}
			term.Cases[i] = c
		}
	}
	// (optional) Metadata.
	md, err := fgen.gen.irMetadataAttachments(old.Metadata())
	if err != nil {
		return errors.WithStack(err)
	}
	term.Metadata = md
	return nil
}

// --- [ indirectbr ] ----------------------------------------------------------

// irIndirectBrTerm translates the AST indirectbr terminator into an equivalent
// IR terminator.
func (fgen *funcGen) irIndirectBrTerm(new ir.Terminator, old *ast.IndirectBrTerm) error {
	term, ok := new.(*ir.TermIndirectBr)
	if !ok {
		panic(fmt.Errorf("invalid IR terminator for AST terminator; expected *ir.TermIndirectBr, got %T", new))
	}
	// Target address.
	addr, err := fgen.irTypeValue(old.Addr())
	if err != nil {
		return errors.WithStack(err)
	}
	term.Addr = addr
	// Valid targets.
	if oldValidTargets := old.ValidTargets(); len(oldValidTargets) > 0 {
		term.ValidTargets = make([]value.Value, len(oldValidTargets))
		for i, oldValidTarget := range oldValidTargets {
			validTarget, err := fgen.irBlock(oldValidTarget)
			if err != nil {
				return errors.WithStack(err)
			}
			term.ValidTargets[i] = validTarget
		}
	}
	// (optional) Metadata.
	md, err := fgen.gen.irMetadataAttachments(old.Metadata())
	if err != nil {
		return errors.WithStack(err)
	}
	term.Metadata = md
	return nil
}

// --- [ invoke ] --------------------------------------------------------------

// irInvokeTerm translates the AST invoke terminator into an equivalent IR
// terminator.
func (fgen *funcGen) irInvokeTerm(new ir.Terminator, old *ast.InvokeTerm) error {
	term, ok := new.(*ir.TermInvoke)
	if !ok {
		panic(fmt.Errorf("invalid IR terminator for AST terminator; expected *ir.TermInvoke, got %T", new))
	}
	// Function arguments.
	if oldArgs := old.Args().Args(); len(oldArgs) > 0 {
		term.Args = make([]value.Value, len(oldArgs))
		for i, oldArg := range oldArgs {
			arg, err := fgen.irArg(oldArg)
			if err != nil {
				return errors.WithStack(err)
			}
			term.Args[i] = arg
		}
	}
	// Invokee.
	typ, err := fgen.gen.irType(old.Typ())
	if err != nil {
		return errors.WithStack(err)
	}
	sig, ok := typ.(*types.FuncType)
	if !ok {
		// Preliminary function signature. Only used by fgen.irValue for inline
		// assembly invokees and constrant expressions.
		var paramTypes []types.Type
		if len(term.Args) > 0 {
			paramTypes = make([]types.Type, len(term.Args))
			for i, arg := range term.Args {
				paramTypes[i] = arg.Type()
			}
		}
		sig = types.NewFunc(typ, paramTypes...)
	}
	// The invokee type is always pointer to function type.
	ptrToSig := types.NewPointer(sig)
	invokee, err := fgen.irValue(ptrToSig, old.Invokee())
	if err != nil {
		return errors.WithStack(err)
	}
	term.Invokee = invokee
	// Normal control flow return point.
	normalRetTarget, err := fgen.irBlock(old.NormalRetTarget())
	if err != nil {
		return errors.WithStack(err)
	}
	term.NormalRetTarget = normalRetTarget
	// Exception control flow return point.
	exceptionRetTarget, err := fgen.irBlock(old.ExceptionRetTarget())
	if err != nil {
		return errors.WithStack(err)
	}
	term.ExceptionRetTarget = exceptionRetTarget
	// (optional) Calling convention.
	if n, ok := old.CallingConv(); ok {
		term.CallingConv = irCallingConv(n)
	}
	// (optional) Return attributes.
	if oldReturnAttrs := old.ReturnAttrs(); len(oldReturnAttrs) > 0 {
		term.ReturnAttrs = make([]ir.ReturnAttribute, len(oldReturnAttrs))
		for i, oldRetAttr := range oldReturnAttrs {
			retAttr := irReturnAttribute(oldRetAttr)
			term.ReturnAttrs[i] = retAttr
		}
	}
	// (optional) Address space.
	if n, ok := old.AddrSpace(); ok {
		term.AddrSpace = irAddrSpace(n)
	}
	// (optional) Function attributes.
	if oldFuncAttrs := old.FuncAttrs(); len(oldFuncAttrs) > 0 {
		term.FuncAttrs = make([]ir.FuncAttribute, len(oldFuncAttrs))
		for i, oldFuncAttr := range oldFuncAttrs {
			funcAttr := fgen.gen.irFuncAttribute(oldFuncAttr)
			term.FuncAttrs[i] = funcAttr
		}
	}
	// (optional) Operand bundles.
	if oldOperandBundles := old.OperandBundles(); len(oldOperandBundles) > 0 {
		term.OperandBundles = make([]*ir.OperandBundle, len(oldOperandBundles))
		for i, oldOperandBundle := range oldOperandBundles {
			operandBundle, err := fgen.irOperandBundle(oldOperandBundle)
			if err != nil {
				return errors.WithStack(err)
			}
			term.OperandBundles[i] = operandBundle
		}
	}
	// (optional) Metadata.
	md, err := fgen.gen.irMetadataAttachments(old.Metadata())
	if err != nil {
		return errors.WithStack(err)
	}
	term.Metadata = md
	return nil
}

// --- [ callbr ] --------------------------------------------------------------

// irCallBrTerm translates the AST callbr terminator into an equivalent IR
// terminator.
func (fgen *funcGen) irCallBrTerm(new ir.Terminator, old *ast.CallBrTerm) error {
	term, ok := new.(*ir.TermCallBr)
	if !ok {
		panic(fmt.Errorf("invalid IR terminator for AST terminator; expected *ir.TermCallBr, got %T", new))
	}
	// Function arguments.
	if oldArgs := old.Args().Args(); len(oldArgs) > 0 {
		term.Args = make([]value.Value, len(oldArgs))
		for i, oldArg := range oldArgs {
			arg, err := fgen.irArg(oldArg)
			if err != nil {
				return errors.WithStack(err)
			}
			term.Args[i] = arg
		}
	}
	// Callee.
	typ, err := fgen.gen.irType(old.Typ())
	if err != nil {
		return errors.WithStack(err)
	}
	sig, ok := typ.(*types.FuncType)
	if !ok {
		// Preliminary function signature. Only used by fgen.irValue for inline
		// assembly callees and constrant expressions.
		var paramTypes []types.Type
		if len(term.Args) > 0 {
			paramTypes = make([]types.Type, len(term.Args))
			for i, arg := range term.Args {
				paramTypes[i] = arg.Type()
			}
		}
		sig = types.NewFunc(typ, paramTypes...)
	}
	// The callee type is always pointer to function type.
	ptrToSig := types.NewPointer(sig)
	callee, err := fgen.irValue(ptrToSig, old.Callee())
	if err != nil {
		return errors.WithStack(err)
	}
	term.Callee = callee
	// Normal control flow return point.
	normalRetTarget, err := fgen.irBlock(old.NormalRetTarget())
	if err != nil {
		return errors.WithStack(err)
	}
	term.NormalRetTarget = normalRetTarget
	// Exception control flow return point.
	for _, oldOtherRetTarget := range old.OtherRetTargets() {
		otherRetTarget, err := fgen.irBlock(oldOtherRetTarget)
		if err != nil {
			return errors.WithStack(err)
		}
		term.OtherRetTargets = append(term.OtherRetTargets, otherRetTarget)
	}
	// (optional) Calling convention.
	if n, ok := old.CallingConv(); ok {
		term.CallingConv = irCallingConv(n)
	}
	// (optional) Return attributes.
	if oldReturnAttrs := old.ReturnAttrs(); len(oldReturnAttrs) > 0 {
		term.ReturnAttrs = make([]ir.ReturnAttribute, len(oldReturnAttrs))
		for i, oldRetAttr := range oldReturnAttrs {
			retAttr := irReturnAttribute(oldRetAttr)
			term.ReturnAttrs[i] = retAttr
		}
	}
	// (optional) Address space.
	if n, ok := old.AddrSpace(); ok {
		term.AddrSpace = irAddrSpace(n)
	}
	// (optional) Function attributes.
	if oldFuncAttrs := old.FuncAttrs(); len(oldFuncAttrs) > 0 {
		term.FuncAttrs = make([]ir.FuncAttribute, len(oldFuncAttrs))
		for i, oldFuncAttr := range oldFuncAttrs {
			funcAttr := fgen.gen.irFuncAttribute(oldFuncAttr)
			term.FuncAttrs[i] = funcAttr
		}
	}
	// (optional) Operand bundles.
	if oldOperandBundles := old.OperandBundles(); len(oldOperandBundles) > 0 {
		term.OperandBundles = make([]*ir.OperandBundle, len(oldOperandBundles))
		for i, oldOperandBundle := range oldOperandBundles {
			operandBundle, err := fgen.irOperandBundle(oldOperandBundle)
			if err != nil {
				return errors.WithStack(err)
			}
			term.OperandBundles[i] = operandBundle
		}
	}
	// (optional) Metadata.
	md, err := fgen.gen.irMetadataAttachments(old.Metadata())
	if err != nil {
		return errors.WithStack(err)
	}
	term.Metadata = md
	return nil
}

// --- [ resume ] --------------------------------------------------------------

// irResumeTerm translates the AST resume terminator into an equivalent IR
// terminator.
func (fgen *funcGen) irResumeTerm(new ir.Terminator, old *ast.ResumeTerm) error {
	term, ok := new.(*ir.TermResume)
	if !ok {
		panic(fmt.Errorf("invalid IR terminator for AST terminator; expected *ir.TermResume, got %T", new))
	}
	// Exception argument to propagate.
	x, err := fgen.irTypeValue(old.X())
	if err != nil {
		return errors.WithStack(err)
	}
	term.X = x
	// (optional) Metadata.
	md, err := fgen.gen.irMetadataAttachments(old.Metadata())
	if err != nil {
		return errors.WithStack(err)
	}
	term.Metadata = md
	return nil
}

// --- [ catchswitch ] ---------------------------------------------------------

// irCatchSwitchTerm translates the AST catchswitch terminator into an
// equivalent IR terminator.
func (fgen *funcGen) irCatchSwitchTerm(new ir.Terminator, old *ast.CatchSwitchTerm) error {
	term, ok := new.(*ir.TermCatchSwitch)
	if !ok {
		panic(fmt.Errorf("invalid IR terminator for AST terminator; expected *ir.TermCatchSwitch, got %T", new))
	}
	// Parent exception pad.
	parentPad, err := fgen.irExceptionPad(old.ParentPad())
	if err != nil {
		return errors.WithStack(err)
	}
	term.ParentPad = parentPad
	// Exception handlers.
	if oldHandlers := old.Handlers().Labels(); len(oldHandlers) > 0 {
		term.Handlers = make([]value.Value, len(oldHandlers))
		for i, oldHandler := range oldHandlers {
			handler, err := fgen.irBlock(oldHandler)
			if err != nil {
				return errors.WithStack(err)
			}
			term.Handlers[i] = handler
		}
	}
	// Optional default unwind target basic block; if nil unwind to caller.
	defaultUnwindTarget, err := fgen.irUnwindTarget(old.DefaultUnwindTarget())
	if err != nil {
		return errors.WithStack(err)
	}
	if defaultUnwindTarget != nil {
		// Note: since DefaultUnwindTarget is an interface we have to be careful
		// with typed nil values (e.g. `(*ir.Block)(nil)`). This is to ensure that
		// DefaultUnwindTarget is nil and not `{Type: ir.Block, Value: nil}`.
		//
		// ref: https://golang.org/doc/faq#nil_error
		term.DefaultUnwindTarget = defaultUnwindTarget
	}
	// (optional) Metadata.
	md, err := fgen.gen.irMetadataAttachments(old.Metadata())
	if err != nil {
		return errors.WithStack(err)
	}
	term.Metadata = md
	return nil
}

// --- [ catchret ] ------------------------------------------------------------

// irCatchRetTerm translates the AST catchret terminator into an equivalent IR
// terminator.
func (fgen *funcGen) irCatchRetTerm(new ir.Terminator, old *ast.CatchRetTerm) error {
	term, ok := new.(*ir.TermCatchRet)
	if !ok {
		panic(fmt.Errorf("invalid IR terminator for AST terminator; expected *ir.TermCatchRet, got %T", new))
	}
	// Exit catchpad.
	v, err := fgen.irValue(types.Token, old.CatchPad())
	if err != nil {
		return errors.WithStack(err)
	}
	catchpad, ok := v.(*ir.InstCatchPad)
	if !ok {
		return errors.Errorf("invalid catchpad type; expected *ir.InstCatchPad, got %T", v)
	}
	term.CatchPad = catchpad
	// Target basic block to transfer control flow to.
	target, err := fgen.irBlock(old.Target())
	if err != nil {
		return errors.WithStack(err)
	}
	term.Target = target
	// (optional) Metadata.
	md, err := fgen.gen.irMetadataAttachments(old.Metadata())
	if err != nil {
		return errors.WithStack(err)
	}
	term.Metadata = md
	return nil
}

// --- [ cleanupret ] ----------------------------------------------------------

// irCleanupRetTerm translates the AST cleanupret terminator into an equivalent
// IR terminator.
func (fgen *funcGen) irCleanupRetTerm(new ir.Terminator, old *ast.CleanupRetTerm) error {
	term, ok := new.(*ir.TermCleanupRet)
	if !ok {
		panic(fmt.Errorf("invalid IR terminator for AST terminator; expected *ir.TermCleanupRet, got %T", new))
	}
	// Exit cleanuppad.
	v, err := fgen.irValue(types.Token, old.CleanupPad())
	if err != nil {
		return errors.WithStack(err)
	}
	cleanuppad, ok := v.(*ir.InstCleanupPad)
	if !ok {
		return errors.Errorf("invalid cleanuppad type; expected *ir.InstCleanupPad, got %T", v)
	}
	term.CleanupPad = cleanuppad
	// Optional unwind target basic block; if nil unwind to caller.
	unwindTarget, err := fgen.irUnwindTarget(old.UnwindTarget())
	if err != nil {
		return errors.WithStack(err)
	}
	if unwindTarget != nil {
		// Note: since UnwindTarget is an interface we have to be careful
		// with typed nil values (e.g. `(*ir.Block)(nil)`). This is to ensure that
		// UnwindTarget is nil and not `{Type: ir.Block, Value: nil}`.
		//
		// ref: https://golang.org/doc/faq#nil_error
		term.UnwindTarget = unwindTarget
	}
	// (optional) Metadata.
	md, err := fgen.gen.irMetadataAttachments(old.Metadata())
	if err != nil {
		return errors.WithStack(err)
	}
	term.Metadata = md
	return nil
}

// --- [ unreachable ] ---------------------------------------------------------

// irUnreachableTerm translates the AST unreachable terminator into an
// equivalent IR terminator.
func (fgen *funcGen) irUnreachableTerm(new ir.Terminator, old *ast.UnreachableTerm) error {
	term, ok := new.(*ir.TermUnreachable)
	if !ok {
		panic(fmt.Errorf("invalid IR terminator for AST terminator; expected *ir.TermUnreachable, got %T", new))
	}
	// (optional) Metadata.
	md, err := fgen.gen.irMetadataAttachments(old.Metadata())
	if err != nil {
		return errors.WithStack(err)
	}
	term.Metadata = md
	return nil
}
