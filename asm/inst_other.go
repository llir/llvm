package asm

import (
	"fmt"

	"github.com/llir/ll/ast"
	asmenum "github.com/llir/llvm/asm/enum"
	"github.com/llir/llvm/ir"
	"github.com/llir/llvm/ir/types"
	"github.com/llir/llvm/ir/value"
	"github.com/pkg/errors"
)

// === [ Create IR ] ===========================================================

// newICmpInst returns a new IR icmp instruction (without body but with type)
// based on the given AST icmp instruction.
func (fgen *funcGen) newICmpInst(ident ir.LocalIdent, old *ast.ICmpInst) (*ir.InstICmp, error) {
	xType, err := fgen.gen.irType(old.X().Typ())
	if err != nil {
		return nil, errors.WithStack(err)
	}
	var typ types.Type
	switch xType := xType.(type) {
	case *types.IntType, *types.PointerType:
		typ = types.I1
	case *types.VectorType:
		typ = types.NewVector(xType.Len, types.I1)
	default:
		panic(fmt.Errorf("invalid icmp operand type; expected *types.IntType, *types.PointerType or *types.VectorType, got %T", xType))
	}
	return &ir.InstICmp{LocalIdent: ident, Typ: typ}, nil
}

// newFCmpInst returns a new IR fcmp instruction (without body but with type)
// based on the given AST fcmp instruction.
func (fgen *funcGen) newFCmpInst(ident ir.LocalIdent, old *ast.FCmpInst) (*ir.InstFCmp, error) {
	xType, err := fgen.gen.irType(old.X().Typ())
	if err != nil {
		return nil, errors.WithStack(err)
	}
	var typ types.Type
	switch xType := xType.(type) {
	case *types.FloatType:
		typ = types.I1
	case *types.VectorType:
		typ = types.NewVector(xType.Len, types.I1)
	default:
		panic(fmt.Errorf("invalid fcmp operand type; expected *types.FloatType or *types.VectorType, got %T", xType))
	}
	return &ir.InstFCmp{LocalIdent: ident, Typ: typ}, nil
}

// newPhiInst returns a new IR phi instruction (without body but with type)
// based on the given AST phi instruction.
func (fgen *funcGen) newPhiInst(ident ir.LocalIdent, old *ast.PhiInst) (*ir.InstPhi, error) {
	typ, err := fgen.gen.irType(old.Typ())
	if err != nil {
		return nil, errors.WithStack(err)
	}
	return &ir.InstPhi{LocalIdent: ident, Typ: typ}, nil
}

// newSelectInst returns a new IR select instruction (without body but with
// type) based on the given AST select instruction.
func (fgen *funcGen) newSelectInst(ident ir.LocalIdent, old *ast.SelectInst) (*ir.InstSelect, error) {
	typ, err := fgen.gen.irType(old.X().Typ())
	if err != nil {
		return nil, errors.WithStack(err)
	}
	return &ir.InstSelect{LocalIdent: ident, Typ: typ}, nil
}

// newCallInst returns a new IR call instruction (without body but with type)
// based on the given AST call instruction.
func (fgen *funcGen) newCallInst(ident ir.LocalIdent, old *ast.CallInst) (*ir.InstCall, error) {
	// Note: the type of call instructions must be determined before assigning
	// local IDs, as they may be values or non-values based on return type.
	typ, err := fgen.gen.irType(old.Typ())
	if err != nil {
		return nil, errors.WithStack(err)
	}
	return &ir.InstCall{LocalIdent: ident, Typ: typ}, nil
}

// newVAArgInst returns a new IR vaarg instruction (without body but with type)
// based on the given AST vaarg instruction.
func (fgen *funcGen) newVAArgInst(ident ir.LocalIdent, old *ast.VAArgInst) (*ir.InstVAArg, error) {
	argType, err := fgen.gen.irType(old.ArgType())
	if err != nil {
		return nil, errors.WithStack(err)
	}
	return &ir.InstVAArg{LocalIdent: ident, ArgType: argType}, nil
}

// newLandingPadInst returns a new IR landingpad instruction (without body but
// with type) based on the given AST landingpad instruction.
func (fgen *funcGen) newLandingPadInst(ident ir.LocalIdent, old *ast.LandingPadInst) (*ir.InstLandingPad, error) {
	resultType, err := fgen.gen.irType(old.ResultType())
	if err != nil {
		return nil, errors.WithStack(err)
	}
	return &ir.InstLandingPad{LocalIdent: ident, ResultType: resultType}, nil
}

// === [ Translate AST to IR ] =================================================

// --- [ icmp ] ----------------------------------------------------------------

// irICmpInst translates the given AST icmp instruction into an equivalent IR
// instruction.
func (fgen *funcGen) irICmpInst(new ir.Instruction, old *ast.ICmpInst) error {
	inst, ok := new.(*ir.InstICmp)
	if !ok {
		panic(fmt.Errorf("invalid IR instruction for AST instruction; expected *ir.InstICmp, got %T", new))
	}
	// Integer comparison predicate.
	inst.Pred = asmenum.IPredFromString(old.Pred().Text())
	// X operand.
	x, err := fgen.irTypeValue(old.X())
	if err != nil {
		return errors.WithStack(err)
	}
	inst.X = x
	// Y operand.
	y, err := fgen.irValue(x.Type(), old.Y())
	if err != nil {
		return errors.WithStack(err)
	}
	inst.Y = y
	// (optional) Metadata.
	md, err := fgen.gen.irMetadataAttachments(old.Metadata())
	if err != nil {
		return errors.WithStack(err)
	}
	inst.Metadata = md
	return nil
}

// --- [ fcmp ] ----------------------------------------------------------------

// irFCmpInst translates the given AST fcmp instruction into an equivalent IR
// instruction.
func (fgen *funcGen) irFCmpInst(new ir.Instruction, old *ast.FCmpInst) error {
	inst, ok := new.(*ir.InstFCmp)
	if !ok {
		panic(fmt.Errorf("invalid IR instruction for AST instruction; expected *ir.InstFCmp, got %T", new))
	}
	// Floating-point comparison predicate.
	inst.Pred = asmenum.FPredFromString(old.Pred().Text())
	// X operand.
	x, err := fgen.irTypeValue(old.X())
	if err != nil {
		return errors.WithStack(err)
	}
	inst.X = x
	// Y operand.
	y, err := fgen.irValue(x.Type(), old.Y())
	if err != nil {
		return errors.WithStack(err)
	}
	inst.Y = y
	// (optional) Fast math flags.
	inst.FastMathFlags = irFastMathFlags(old.FastMathFlags())
	// (optional) Metadata.
	md, err := fgen.gen.irMetadataAttachments(old.Metadata())
	if err != nil {
		return errors.WithStack(err)
	}
	inst.Metadata = md
	return nil
}

// --- [ phi ] -----------------------------------------------------------------

// irPhiInst translates the given AST phi instruction into an equivalent IR
// instruction.
func (fgen *funcGen) irPhiInst(new ir.Instruction, old *ast.PhiInst) error {
	inst, ok := new.(*ir.InstPhi)
	if !ok {
		panic(fmt.Errorf("invalid IR instruction for AST instruction; expected *ir.InstPhi, got %T", new))
	}
	// Type of incoming values.
	typ, err := fgen.gen.irType(old.Typ())
	if err != nil {
		return errors.WithStack(err)
	}
	// Incoming values.
	if oldIncs := old.Incs(); len(oldIncs) > 0 {
		inst.Incs = make([]*ir.Incoming, len(oldIncs))
		for i, oldInc := range oldIncs {
			inc, err := fgen.irIncoming(typ, oldInc.X(), oldInc.Pred())
			if err != nil {
				return errors.WithStack(err)
			}
			inst.Incs[i] = inc
		}
	}
	// (optional) Metadata.
	md, err := fgen.gen.irMetadataAttachments(old.Metadata())
	if err != nil {
		return errors.WithStack(err)
	}
	inst.Metadata = md
	return nil
}

// --- [ select ] --------------------------------------------------------------

// irSelectInst translates the given AST select instruction into an equivalent
// IR instruction.
func (fgen *funcGen) irSelectInst(new ir.Instruction, old *ast.SelectInst) error {
	inst, ok := new.(*ir.InstSelect)
	if !ok {
		panic(fmt.Errorf("invalid IR instruction for AST instruction; expected *ir.InstSelect, got %T", new))
	}
	// Selection condition.
	cond, err := fgen.irTypeValue(old.Cond())
	if err != nil {
		return errors.WithStack(err)
	}
	inst.Cond = cond
	// X operand.
	x, err := fgen.irTypeValue(old.X())
	if err != nil {
		return errors.WithStack(err)
	}
	inst.X = x
	// Y operand.
	y, err := fgen.irTypeValue(old.Y())
	if err != nil {
		return errors.WithStack(err)
	}
	inst.Y = y
	// (optional) Fast math flags.
	inst.FastMathFlags = irFastMathFlags(old.FastMathFlags())
	// (optional) Metadata.
	md, err := fgen.gen.irMetadataAttachments(old.Metadata())
	if err != nil {
		return errors.WithStack(err)
	}
	inst.Metadata = md
	return nil
}

// --- [ call ] ----------------------------------------------------------------

// irCallInst translates the given AST call instruction into an equivalent IR
// instruction.
func (fgen *funcGen) irCallInst(new ir.Instruction, old *ast.CallInst) error {
	inst, ok := new.(*ir.InstCall)
	if !ok {
		panic(fmt.Errorf("invalid IR instruction for AST instruction; expected *ir.InstCall, got %T", new))
	}
	// Function arguments.
	if oldArgs := old.Args().Args(); len(oldArgs) > 0 {
		inst.Args = make([]value.Value, len(oldArgs))
		for i, oldArg := range oldArgs {
			arg, err := fgen.irArg(oldArg)
			if err != nil {
				return errors.WithStack(err)
			}
			inst.Args[i] = arg
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
		if len(inst.Args) > 0 {
			paramTypes = make([]types.Type, len(inst.Args))
			for i, arg := range inst.Args {
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
	inst.Callee = callee
	// (optional) Tail.
	if n, ok := old.Tail(); ok {
		inst.Tail = asmenum.TailFromString(n.Text())
	}
	// (optional) Fast math flags.
	inst.FastMathFlags = irFastMathFlags(old.FastMathFlags())
	// (optional) Calling convention.
	if n, ok := old.CallingConv(); ok {
		inst.CallingConv = irCallingConv(n)
	}
	// (optional) Return attributes.
	if oldReturnAttrs := old.ReturnAttrs(); len(oldReturnAttrs) > 0 {
		inst.ReturnAttrs = make([]ir.ReturnAttribute, len(oldReturnAttrs))
		for i, oldRetAttr := range oldReturnAttrs {
			retAttr := irReturnAttribute(oldRetAttr)
			inst.ReturnAttrs[i] = retAttr
		}
	}
	// (optional) Address space.
	if n, ok := old.AddrSpace(); ok {
		inst.AddrSpace = irAddrSpace(n)
	}
	// (optional) Function attributes.
	if oldFuncAttrs := old.FuncAttrs(); len(oldFuncAttrs) > 0 {
		inst.FuncAttrs = make([]ir.FuncAttribute, len(oldFuncAttrs))
		for i, oldFuncAttr := range oldFuncAttrs {
			funcAttr := fgen.gen.irFuncAttribute(oldFuncAttr)
			inst.FuncAttrs[i] = funcAttr
		}
	}
	// (optional) Operand bundles.
	if oldOperandBundles := old.OperandBundles(); len(oldOperandBundles) > 0 {
		inst.OperandBundles = make([]*ir.OperandBundle, len(oldOperandBundles))
		for i, oldOperandBundle := range oldOperandBundles {
			operandBundle, err := fgen.irOperandBundle(oldOperandBundle)
			if err != nil {
				return errors.WithStack(err)
			}
			inst.OperandBundles[i] = operandBundle
		}
	}
	// (optional) Metadata.
	md, err := fgen.gen.irMetadataAttachments(old.Metadata())
	if err != nil {
		return errors.WithStack(err)
	}
	inst.Metadata = md
	return nil
}

// --- [ va_arg ] --------------------------------------------------------------

// irVAArgInst translates the given AST vaarg instruction into an equivalent IR
// instruction.
func (fgen *funcGen) irVAArgInst(new ir.Instruction, old *ast.VAArgInst) error {
	inst, ok := new.(*ir.InstVAArg)
	if !ok {
		panic(fmt.Errorf("invalid IR instruction for AST instruction; expected *ir.InstVAArg, got %T", new))
	}
	// Variable argument list.
	argList, err := fgen.irTypeValue(old.ArgList())
	if err != nil {
		return errors.WithStack(err)
	}
	inst.ArgList = argList
	// Argument type.
	argType, err := fgen.gen.irType(old.ArgType())
	if err != nil {
		return errors.WithStack(err)
	}
	inst.ArgType = argType
	// (optional) Metadata.
	md, err := fgen.gen.irMetadataAttachments(old.Metadata())
	if err != nil {
		return errors.WithStack(err)
	}
	inst.Metadata = md
	return nil
}

// --- [ landingpad ] ----------------------------------------------------------

// irLandingPadInst translates the given AST landingpad instruction into an
// equivalent IR instruction.
func (fgen *funcGen) irLandingPadInst(new ir.Instruction, old *ast.LandingPadInst) error {
	inst, ok := new.(*ir.InstLandingPad)
	if !ok {
		panic(fmt.Errorf("invalid IR instruction for AST instruction; expected *ir.InstLandingPad, got %T", new))
	}
	// Result type.
	resultType, err := fgen.gen.irType(old.ResultType())
	if err != nil {
		return errors.WithStack(err)
	}
	inst.ResultType = resultType
	// (optional) Cleanup landing pad.
	_, inst.Cleanup = old.Cleanup()
	// Filter and catch clauses.
	if oldClauses := old.Clauses(); len(oldClauses) > 0 {
		inst.Clauses = make([]*ir.Clause, len(oldClauses))
		for i, oldClause := range oldClauses {
			clause, err := fgen.irClause(oldClause)
			if err != nil {
				return errors.WithStack(err)
			}
			inst.Clauses[i] = clause
		}
	}
	// (optional) Metadata.
	md, err := fgen.gen.irMetadataAttachments(old.Metadata())
	if err != nil {
		return errors.WithStack(err)
	}
	inst.Metadata = md
	return nil
}

// --- [ catchpad ] ------------------------------------------------------------

// irCatchPadInst translates the given AST catchpad instruction into an
// equivalent IR instruction.
func (fgen *funcGen) irCatchPadInst(new ir.Instruction, old *ast.CatchPadInst) error {
	inst, ok := new.(*ir.InstCatchPad)
	if !ok {
		panic(fmt.Errorf("invalid IR instruction for AST instruction; expected *ir.InstCatchPad, got %T", new))
	}
	// Exception scope.
	ident := localIdent(old.Scope())
	v, ok := fgen.locals[ident]
	if !ok {
		return errors.Errorf("unable to locate local identifier %q", ident.Ident())
	}
	scope, ok := v.(*ir.TermCatchSwitch)
	if !ok {
		return errors.Errorf("invalid scope type; expected *ir.TermCatchSwitch, got %T", v)
	}
	inst.Scope = scope
	// Exception arguments.
	if oldArgs := old.Args(); len(oldArgs) > 0 {
		inst.Args = make([]value.Value, len(oldArgs))
		for i, oldArg := range oldArgs {
			arg, err := fgen.irExceptionArg(oldArg)
			if err != nil {
				return errors.WithStack(err)
			}
			inst.Args[i] = arg
		}
	}
	// (optional) Metadata.
	md, err := fgen.gen.irMetadataAttachments(old.Metadata())
	if err != nil {
		return errors.WithStack(err)
	}
	inst.Metadata = md
	return nil
}

// --- [ cleanuppad ] ----------------------------------------------------------

// irCleanupPadInst translates the given AST cleanuppad instruction into an
// equivalent IR instruction.
func (fgen *funcGen) irCleanupPadInst(new ir.Instruction, old *ast.CleanupPadInst) error {
	inst, ok := new.(*ir.InstCleanupPad)
	if !ok {
		panic(fmt.Errorf("invalid IR instruction for AST instruction; expected *ir.InstCleanupPad, got %T", new))
	}
	// Exception scope.
	scope, err := fgen.irExceptionScope(old.Scope())
	if err != nil {
		return errors.WithStack(err)
	}
	inst.Scope = scope
	// Exception arguments.
	if oldArgs := old.Args(); len(oldArgs) > 0 {
		inst.Args = make([]value.Value, len(oldArgs))
		for i, oldArg := range oldArgs {
			arg, err := fgen.irExceptionArg(oldArg)
			if err != nil {
				return errors.WithStack(err)
			}
			inst.Args[i] = arg
		}
	}
	// (optional) Metadata.
	md, err := fgen.gen.irMetadataAttachments(old.Metadata())
	if err != nil {
		return errors.WithStack(err)
	}
	inst.Metadata = md
	return nil
}
