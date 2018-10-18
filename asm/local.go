// TODO: make concurrent :)

package asm

import (
	"fmt"

	"github.com/llir/l/ir"
	"github.com/llir/l/ir/value"
	"github.com/mewmew/l-tm/asm/ll/ast"
	"github.com/pkg/errors"
)

type funcGen struct {
	// Module generator.
	*generator

	// LLVM IR function being generated.
	f *ir.Function

	// ls maps from local identifier (without '%' prefix) to corresponding IR
	// value.
	ls map[string]value.Value
}

func newFuncGen(gen *generator, f *ir.Function) *funcGen {
	return &funcGen{
		generator: gen,
		f:         f,
		ls:        make(map[string]value.Value),
	}
}

// resolveLocals resolves the local variables, basic blocks and function
// parameters of the given function body. The returned value maps from local
// identifier (without '%' prefix) to the corresponding IR value.
func (fgen *funcGen) resolveLocals(body ast.FuncBody) (map[string]value.Value, error) {
	// Create instructions (without bodies), in preparation for index.
	f := fgen.f
	bbs := body.Blocks()
	for _, b := range bbs {
		blockName := label(*b.Name())
		block := ir.NewBlock(blockName)
		for _, i := range b.Insts() {
			inst, err := fgen.newIRInst(i)
			if err != nil {
				return nil, errors.WithStack(err)
			}
			block.Insts = append(block.Insts, inst)
		}
		f.Blocks = append(f.Blocks, block)
	}
	// Assign local IDs.
	if err := f.AssignIDs(); err != nil {
		return nil, errors.WithStack(err)
	}
	// Index local identifiers.
	// TODO: index local identifiers.

	// Translate instructions.
	for i, block := range f.Blocks {
		insts := bbs[i].Insts()
		for j, inst := range block.Insts {
			old := insts[j]
			if _, err := fgen.translateInst(inst, old); err != nil {
				return nil, errors.WithStack(err)
			}
		}
	}
	// TODO: implement.
	panic("not yet implemented")
	// Translate terminators.
	// TODO: implement.
}

// newIRInst returns a new IR instruction (without body) based on the given AST
// instruction.
func (fgen *funcGen) newIRInst(old ast.Instruction) (ir.Instruction, error) {
	switch old := old.(type) {
	// Value instruction.
	case *ast.LocalDef:
		name := local(old.Name())
		return fgen.newIRValueInst(name, old.Inst())
	case ast.ValueInstruction:
		return fgen.newIRValueInst("", old)
	// Non-value instructions.
	case *ast.StoreInst:
		return &ir.InstStore{}, nil
	case *ast.FenceInst:
		return &ir.InstFence{}, nil
	case *ast.CmpXchgInst:
		return &ir.InstCmpXchg{}, nil
	case *ast.AtomicRMWInst:
		return &ir.InstAtomicRMW{}, nil
	default:
		panic(fmt.Errorf("support for AST instruction type %T not yet implemented", old))
	}
}

// newIRValueInst returns a new IR value instruction (without body) based on the
// given AST value instruction.
func (fgen *funcGen) newIRValueInst(name string, old ast.ValueInstruction) (ir.Instruction, error) {
	switch old := old.(type) {
	// Binary instructions
	case *ast.AddInst:
		return &ir.InstAdd{LocalName: name}, nil
	case *ast.FAddInst:
		return &ir.InstFAdd{LocalName: name}, nil
	case *ast.SubInst:
		return &ir.InstSub{LocalName: name}, nil
	case *ast.FSubInst:
		return &ir.InstFSub{LocalName: name}, nil
	case *ast.MulInst:
		return &ir.InstMul{LocalName: name}, nil
	case *ast.FMulInst:
		return &ir.InstFMul{LocalName: name}, nil
	case *ast.UDivInst:
		return &ir.InstUDiv{LocalName: name}, nil
	case *ast.SDivInst:
		return &ir.InstSDiv{LocalName: name}, nil
	case *ast.FDivInst:
		return &ir.InstFDiv{LocalName: name}, nil
	case *ast.URemInst:
		return &ir.InstURem{LocalName: name}, nil
	case *ast.SRemInst:
		return &ir.InstSRem{LocalName: name}, nil
	case *ast.FRemInst:
		return &ir.InstFRem{LocalName: name}, nil
	// Bitwise instructions
	case *ast.ShlInst:
		return &ir.InstShl{LocalName: name}, nil
	case *ast.LShrInst:
		return &ir.InstLShr{LocalName: name}, nil
	case *ast.AShrInst:
		return &ir.InstAShr{LocalName: name}, nil
	case *ast.AndInst:
		return &ir.InstAnd{LocalName: name}, nil
	case *ast.OrInst:
		return &ir.InstOr{LocalName: name}, nil
	case *ast.XorInst:
		return &ir.InstXor{LocalName: name}, nil
	// Vector instructions
	case *ast.ExtractElementInst:
		return &ir.InstExtractElement{LocalName: name}, nil
	case *ast.InsertElementInst:
		return &ir.InstInsertElement{LocalName: name}, nil
	case *ast.ShuffleVectorInst:
		return &ir.InstShuffleVector{LocalName: name}, nil
	// Aggregate instructions
	case *ast.ExtractValueInst:
		return &ir.InstExtractValue{LocalName: name}, nil
	case *ast.InsertValueInst:
		return &ir.InstInsertValue{LocalName: name}, nil
	// Memory instructions
	case *ast.AllocaInst:
		return &ir.InstAlloca{LocalName: name}, nil
	case *ast.LoadInst:
		return &ir.InstLoad{LocalName: name}, nil
	case *ast.GetElementPtrInst:
		return &ir.InstGetElementPtr{LocalName: name}, nil
	// Conversion instructions
	case *ast.TruncInst:
		return &ir.InstTrunc{LocalName: name}, nil
	case *ast.ZExtInst:
		return &ir.InstZExt{LocalName: name}, nil
	case *ast.SExtInst:
		return &ir.InstSExt{LocalName: name}, nil
	case *ast.FPTruncInst:
		return &ir.InstFPTrunc{LocalName: name}, nil
	case *ast.FPExtInst:
		return &ir.InstFPExt{LocalName: name}, nil
	case *ast.FPToUIInst:
		return &ir.InstFPToUI{LocalName: name}, nil
	case *ast.FPToSIInst:
		return &ir.InstFPToSI{LocalName: name}, nil
	case *ast.UIToFPInst:
		return &ir.InstUIToFP{LocalName: name}, nil
	case *ast.SIToFPInst:
		return &ir.InstSIToFP{LocalName: name}, nil
	case *ast.PtrToIntInst:
		return &ir.InstPtrToInt{LocalName: name}, nil
	case *ast.IntToPtrInst:
		return &ir.InstIntToPtr{LocalName: name}, nil
	case *ast.BitCastInst:
		return &ir.InstBitCast{LocalName: name}, nil
	case *ast.AddrSpaceCastInst:
		return &ir.InstAddrSpaceCast{LocalName: name}, nil
	// Other instructions
	case *ast.ICmpInst:
		return &ir.InstICmp{LocalName: name}, nil
	case *ast.FCmpInst:
		return &ir.InstFCmp{LocalName: name}, nil
	case *ast.PhiInst:
		return &ir.InstPhi{LocalName: name}, nil
	case *ast.SelectInst:
		return &ir.InstSelect{LocalName: name}, nil
	case *ast.CallInst:
		// NOTE: We need to store the type of call instructions before invoking
		// f.AssignIDs, since call instructions may be value instructions or
		// non-value instructions based on return type.
		typ, err := fgen.irType(old.Typ())
		if err != nil {
			return nil, errors.WithStack(err)
		}
		return &ir.InstCall{LocalName: name, Typ: typ}, nil
	case *ast.VAArgInst:
		return &ir.InstVAArg{LocalName: name}, nil
	case *ast.LandingPadInst:
		return &ir.InstLandingPad{LocalName: name}, nil
	case *ast.CatchPadInst:
		return &ir.InstCatchPad{LocalName: name}, nil
	case *ast.CleanupPadInst:
		return &ir.InstCleanupPad{LocalName: name}, nil
	default:
		panic(fmt.Errorf("support for AST value instruction type %T not yet implemented", old))
	}
}

// === [ Instructions ] ========================================================

// translateInst translates the AST instruction into an equivalent IR
// instruction.
func (fgen *funcGen) translateInst(inst ir.Instruction, old ast.Instruction) (ir.Instruction, error) {
	switch old := old.(type) {
	case *ast.AddInst:
		return fgen.translateAddInst(inst, old)
	default:
		panic(fmt.Errorf("support for instruction type %T not yet implemented", old))
	}
}

// --- [ Binary instructions ] -------------------------------------------------

// ~~~ [ add ] ~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~

func (fgen *funcGen) translateAddInst(inst ir.Instruction, old *ast.AddInst) (*ir.InstAdd, error) {
	i, ok := inst.(*ir.InstAdd)
	if !ok {
		// NOTE: panic since this would indicate a bug in the implementation.
		panic(fmt.Errorf("invalid IR instruction for AST instruction; expected *ir.InstAdd, got %T", inst))
	}
	// TODO: implement
	return i, nil
}

// ~~~ [ fadd ] ~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~

// ~~~ [ sub ] ~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~

// ~~~ [ fsub ] ~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~

// ~~~ [ mul ] ~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~

// ~~~ [ fmul ] ~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~

// ~~~ [ udiv ] ~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~

// ~~~ [ sdiv ] ~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~

// ~~~ [ fdiv ] ~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~

// ~~~ [ urem ] ~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~

// ~~~ [ srem ] ~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~

// ~~~ [ frem ] ~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~

// --- [ Bitwise instructions ] ------------------------------------------------

// ~~~ [ shl ] ~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~

// ~~~ [ lshr ] ~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~

// ~~~ [ ashr ] ~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~

// ~~~ [ and ] ~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~

// ~~~ [ or ] ~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~

// ~~~ [ xor ] ~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~

// --- [ Vector instructions ] -------------------------------------------------

// ~~~ [ extractelement ] ~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~

// ~~~ [ insertelement ] ~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~

// ~~~ [ shufflevector ] ~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~

// --- [ Aggregate instructions ] ----------------------------------------------

// ~~~ [ extractvalue ] ~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~

// ~~~ [ insertvalue ] ~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~

// --- [ Memory instructions ] -------------------------------------------------

// ~~~ [ alloca ] ~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~

// ~~~ [ load ] ~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~

// ~~~ [ store ] ~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~

// ~~~ [ fence ] ~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~

// ~~~ [ cmpxchg ] ~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~

// ~~~ [ atomicrmw ] ~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~

// ~~~ [ getelementptr ] ~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~

// --- [ Conversion instructions ] ---------------------------------------------

// ~~~ [ trunc ] ~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~

// ~~~ [ zext ] ~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~

// ~~~ [ sext ] ~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~

// ~~~ [ fptrunc ] ~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~

// ~~~ [ fpext ] ~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~

// ~~~ [ fptoui ] ~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~

// ~~~ [ fptosi ] ~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~

// ~~~ [ uitofp ] ~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~

// ~~~ [ sitofp ] ~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~

// ~~~ [ ptrtoint ] ~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~

// ~~~ [ inttoptr ] ~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~

// ~~~ [ bitcast ] ~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~

// ~~~ [ addrspacecast ] ~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~

// --- [ Other instructions ] --------------------------------------------------

// ~~~ [ icmp ] ~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~

// ~~~ [ fcmp ] ~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~

// ~~~ [ phi ] ~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~

// ~~~ [ select ] ~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~

// ~~~ [ call ] ~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~

// ~~~ [ va_arg ] ~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~

// ~~~ [ landingpad ] ~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~

// --- [ catchpad ] ------------------------------------------------------------

// --- [ cleanuppad ] ----------------------------------------------------------
