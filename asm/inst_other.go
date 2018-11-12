package asm

import (
	"fmt"

	"github.com/llir/ll/ast"
	asmenum "github.com/llir/llvm/asm/enum"
	"github.com/llir/llvm/ir"
	"github.com/llir/llvm/ir/types"
	"github.com/pkg/errors"
)

// --- [ Other instructions ] --------------------------------------------------

// ~~~ [ icmp ] ~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~

// astToIRInstICmp translates the given AST icmp instruction into an equivalent
// IR instruction.
func (fgen *funcGen) astToIRInstICmp(inst ir.Instruction, old *ast.ICmpInst) (*ir.InstICmp, error) {
	i, ok := inst.(*ir.InstICmp)
	if !ok {
		panic(fmt.Errorf("invalid IR instruction for AST instruction; expected *ir.InstICmp, got %T", inst))
	}
	// Integer comparison predicate.
	i.Pred = asmenum.IPredFromString(old.Pred().Text())
	// X operand.
	x, err := fgen.astToIRTypeValue(old.X())
	if err != nil {
		return nil, errors.WithStack(err)
	}
	i.X = x
	// Y operand.
	y, err := fgen.astToIRValue(x.Type(), old.Y())
	if err != nil {
		return nil, errors.WithStack(err)
	}
	i.Y = y
	// (optional) Metadata.
	md, err := fgen.gen.irMetadataAttachments(old.Metadata())
	if err != nil {
		return nil, errors.WithStack(err)
	}
	i.Metadata = md
	return i, nil
}

// ~~~ [ fcmp ] ~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~

// astToIRInstFCmp translates the given AST fcmp instruction into an equivalent
// IR instruction.
func (fgen *funcGen) astToIRInstFCmp(inst ir.Instruction, old *ast.FCmpInst) (*ir.InstFCmp, error) {
	i, ok := inst.(*ir.InstFCmp)
	if !ok {
		panic(fmt.Errorf("invalid IR instruction for AST instruction; expected *ir.InstFCmp, got %T", inst))
	}
	// (optional) Fast math flags.
	i.FastMathFlags = irFastMathFlags(old.FastMathFlags())
	// Floating-point comparison predicate.
	i.Pred = asmenum.FPredFromString(old.Pred().Text())
	// X operand.
	x, err := fgen.astToIRTypeValue(old.X())
	if err != nil {
		return nil, errors.WithStack(err)
	}
	i.X = x
	// Y operand.
	y, err := fgen.astToIRValue(x.Type(), old.Y())
	if err != nil {
		return nil, errors.WithStack(err)
	}
	i.Y = y
	// (optional) Metadata.
	md, err := fgen.gen.irMetadataAttachments(old.Metadata())
	if err != nil {
		return nil, errors.WithStack(err)
	}
	i.Metadata = md
	return i, nil
}

// ~~~ [ phi ] ~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~

// astToIRInstPhi translates the given AST phi instruction into an equivalent IR
// instruction.
func (fgen *funcGen) astToIRInstPhi(inst ir.Instruction, old *ast.PhiInst) (*ir.InstPhi, error) {
	i, ok := inst.(*ir.InstPhi)
	if !ok {
		panic(fmt.Errorf("invalid IR instruction for AST instruction; expected *ir.InstPhi, got %T", inst))
	}
	typ, err := fgen.gen.irType(old.Typ())
	if err != nil {
		return nil, errors.WithStack(err)
	}
	// Incoming values.
	for _, oldInc := range old.Incs() {
		inc, err := fgen.irIncoming(typ, oldInc.X(), oldInc.Pred())
		if err != nil {
			return nil, errors.WithStack(err)
		}
		i.Incs = append(i.Incs, inc)
	}
	// (optional) Metadata.
	md, err := fgen.gen.irMetadataAttachments(old.Metadata())
	if err != nil {
		return nil, errors.WithStack(err)
	}
	i.Metadata = md
	return i, nil
}

// ~~~ [ select ] ~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~

// astToIRInstSelect translates the given AST select instruction into an
// equivalent IR instruction.
func (fgen *funcGen) astToIRInstSelect(inst ir.Instruction, old *ast.SelectInst) (*ir.InstSelect, error) {
	i, ok := inst.(*ir.InstSelect)
	if !ok {
		panic(fmt.Errorf("invalid IR instruction for AST instruction; expected *ir.InstSelect, got %T", inst))
	}
	// Selection condition.
	cond, err := fgen.astToIRTypeValue(old.Cond())
	if err != nil {
		return nil, errors.WithStack(err)
	}
	i.Cond = cond
	// X operand.
	x, err := fgen.astToIRTypeValue(old.X())
	if err != nil {
		return nil, errors.WithStack(err)
	}
	i.X = x
	// Y operand.
	y, err := fgen.astToIRTypeValue(old.Y())
	if err != nil {
		return nil, errors.WithStack(err)
	}
	i.Y = y
	// (optional) Metadata.
	md, err := fgen.gen.irMetadataAttachments(old.Metadata())
	if err != nil {
		return nil, errors.WithStack(err)
	}
	i.Metadata = md
	return i, nil
}

// ~~~ [ call ] ~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~

// astToIRInstCall translates the given AST call instruction into an equivalent
// IR instruction.
func (fgen *funcGen) astToIRInstCall(inst ir.Instruction, old *ast.CallInst) (*ir.InstCall, error) {
	i, ok := inst.(*ir.InstCall)
	if !ok {
		panic(fmt.Errorf("invalid IR instruction for AST instruction; expected *ir.InstCall, got %T", inst))
	}
	// (optional) Tail.
	if n := old.Tail(); n.IsValid() {
		i.Tail = asmenum.TailFromString(n.Text())
	}
	// (optional) Fast math flags.
	i.FastMathFlags = irFastMathFlags(old.FastMathFlags())
	// (optional) Calling convention.
	if n := old.CallingConv(); n.LlvmNode().IsValid() {
		i.CallingConv = irCallingConv(n)
	}
	// (optional) Return attributes.
	for _, oldRetAttr := range old.ReturnAttrs() {
		retAttr := irReturnAttribute(oldRetAttr)
		i.ReturnAttrs = append(i.ReturnAttrs, retAttr)
	}
	// (optional) Address space.
	if n := old.AddrSpace(); n.IsValid() {
		i.AddrSpace = irAddrSpace(n)
	}
	// Callee.
	typ, err := fgen.gen.irType(old.Typ())
	if err != nil {
		return nil, errors.WithStack(err)
	}
	sig, ok := typ.(*types.FuncType)
	if !ok {
		// Preliminary function signature. Only used by astToIRValue for inline
		// assembly callees and constrant expressions.
		sig = types.NewFunc(typ)
		// TODO: add parameters to sig.
	}
	callee, err := fgen.astToIRValue(sig, old.Callee())
	if err != nil {
		return nil, errors.WithStack(err)
	}
	i.Callee = callee
	// Function arguments.
	for _, oldArg := range old.Args().Args() {
		arg, err := fgen.irArg(oldArg)
		if err != nil {
			return nil, errors.WithStack(err)
		}
		i.Args = append(i.Args, arg)
	}
	// (optional) Function attributes.
	for _, oldFuncAttr := range old.FuncAttrs() {
		funcAttr := fgen.gen.irFuncAttribute(oldFuncAttr)
		i.FuncAttrs = append(i.FuncAttrs, funcAttr)
	}
	// (optional) Operand bundles.
	for _, oldOperandBundle := range old.OperandBundles() {
		operandBundle, err := fgen.irOperandBundle(oldOperandBundle)
		if err != nil {
			return nil, errors.WithStack(err)
		}
		i.OperandBundles = append(i.OperandBundles, operandBundle)
	}
	// (optional) Metadata.
	md, err := fgen.gen.irMetadataAttachments(old.Metadata())
	if err != nil {
		return nil, errors.WithStack(err)
	}
	i.Metadata = md
	return i, nil
}

// ~~~ [ va_arg ] ~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~

// astToIRInstVAArg translates the given AST vaarg instruction into an
// equivalent IR instruction.
func (fgen *funcGen) astToIRInstVAArg(inst ir.Instruction, old *ast.VAArgInst) (*ir.InstVAArg, error) {
	i, ok := inst.(*ir.InstVAArg)
	if !ok {
		panic(fmt.Errorf("invalid IR instruction for AST instruction; expected *ir.InstVAArg, got %T", inst))
	}
	// Variable argument list.
	argList, err := fgen.astToIRTypeValue(old.ArgList())
	if err != nil {
		return nil, errors.WithStack(err)
	}
	i.ArgList = argList
	// Argument type.
	argType, err := fgen.gen.irType(old.ArgType())
	if err != nil {
		return nil, errors.WithStack(err)
	}
	i.ArgType = argType
	// (optional) Metadata.
	md, err := fgen.gen.irMetadataAttachments(old.Metadata())
	if err != nil {
		return nil, errors.WithStack(err)
	}
	i.Metadata = md
	return i, nil
}

// ~~~ [ landingpad ] ~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~

// astToIRInstLandingPad translates the given AST landingpad instruction into an
// equivalent IR instruction.
func (fgen *funcGen) astToIRInstLandingPad(inst ir.Instruction, old *ast.LandingPadInst) (*ir.InstLandingPad, error) {
	i, ok := inst.(*ir.InstLandingPad)
	if !ok {
		panic(fmt.Errorf("invalid IR instruction for AST instruction; expected *ir.InstLandingPad, got %T", inst))
	}
	// Result type.
	resultType, err := fgen.gen.irType(old.ResultType())
	if err != nil {
		return nil, errors.WithStack(err)
	}
	i.ResultType = resultType
	// (optional) Cleanup landing pad.
	i.Cleanup = old.Cleanup().IsValid()
	// Filter and catch clauses.
	for _, oldClause := range old.Clauses() {
		clause, err := fgen.irClause(oldClause)
		if err != nil {
			return nil, errors.WithStack(err)
		}
		i.Clauses = append(i.Clauses, clause)
	}
	// (optional) Metadata.
	md, err := fgen.gen.irMetadataAttachments(old.Metadata())
	if err != nil {
		return nil, errors.WithStack(err)
	}
	i.Metadata = md
	return i, nil
}

// ~~~ [ catchpad ] ~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~

// astToIRInstCatchPad translates the given AST catchpad instruction into an
// equivalent IR instruction.
func (fgen *funcGen) astToIRInstCatchPad(inst ir.Instruction, old *ast.CatchPadInst) (*ir.InstCatchPad, error) {
	i, ok := inst.(*ir.InstCatchPad)
	if !ok {
		panic(fmt.Errorf("invalid IR instruction for AST instruction; expected *ir.InstCatchPad, got %T", inst))
	}
	// Exception scope.
	ident := localIdent(old.Scope())
	v, ok := fgen.ls[ident]
	if !ok {
		return nil, errors.Errorf("unable to locate local identifier %q", ident)
	}
	scope, ok := v.(*ir.TermCatchSwitch)
	if !ok {
		return nil, errors.Errorf("invalid scope type; expected *ir.TermCatchSwitch, got %T", v)
	}
	i.Scope = scope
	// Exception arguments.
	for _, oldArg := range old.Args() {
		arg, err := fgen.irExceptionArg(oldArg)
		if err != nil {
			return nil, errors.WithStack(err)
		}
		i.Args = append(i.Args, arg)
	}
	// (optional) Metadata.
	md, err := fgen.gen.irMetadataAttachments(old.Metadata())
	if err != nil {
		return nil, errors.WithStack(err)
	}
	i.Metadata = md
	return i, nil
}

// ~~~ [ cleanuppad ] ~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~

// astToIRInstCleanupPad translates the given AST cleanuppad instruction into an
// equivalent IR instruction.
func (fgen *funcGen) astToIRInstCleanupPad(inst ir.Instruction, old *ast.CleanupPadInst) (*ir.InstCleanupPad, error) {
	i, ok := inst.(*ir.InstCleanupPad)
	if !ok {
		panic(fmt.Errorf("invalid IR instruction for AST instruction; expected *ir.InstCleanupPad, got %T", inst))
	}
	// Exception scope.
	scope, err := fgen.irExceptionScope(old.Scope())
	if err != nil {
		return nil, errors.WithStack(err)
	}
	i.Scope = scope
	// Exception arguments.
	for _, oldArg := range old.Args() {
		arg, err := fgen.irExceptionArg(oldArg)
		if err != nil {
			return nil, errors.WithStack(err)
		}
		i.Args = append(i.Args, arg)
	}
	// (optional) Metadata.
	md, err := fgen.gen.irMetadataAttachments(old.Metadata())
	if err != nil {
		return nil, errors.WithStack(err)
	}
	i.Metadata = md
	return i, nil
}
