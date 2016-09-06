package irx

import (
	"fmt"

	"github.com/llir/llvm/ir"
	"github.com/llir/llvm/ir/constant"
	"github.com/llir/llvm/ir/instruction"
	"github.com/llir/llvm/ir/value"
	"github.com/mewkiz/pkg/errutil"
)

// === [ Modules ] =============================================================

// dummyMap maps local variable names to their corresponding values.
type dummyMap map[string]value.Value

// fixModule replaces dummy values within the given module with their
// corresponding local variables.
func fixModule(module *ir.Module) *ir.Module {
	for i, oldFunc := range module.Funcs {
		// Allocate function specific mapping from local variables names to their
		// corresponding values.
		m := make(dummyMap)
		f := m.fixFunc(oldFunc)
		module.Funcs[i] = f
	}
	return module
}

// === [ Functions ] ===========================================================

// fixFunc replaces dummy values within the given function with their
// corresponding local variables.
func (m dummyMap) fixFunc(oldFunc *ir.Function) *ir.Function {
	f := ir.NewFunction(oldFunc.Name(), oldFunc.Sig())
	for _, oldBlock := range oldFunc.Blocks() {
		block := m.fixBlock(oldBlock)
		f.AppendBlock(block)
	}
	if err := f.AssignIDs(); err != nil {
		panic(errutil.Err(err))
	}
	return f
}

// === [ Basic blocks ] ========================================================

// fixBlock replaces dummy values within the given basic block with their
// corresponding local variables.
func (m dummyMap) fixBlock(oldBlock *ir.BasicBlock) *ir.BasicBlock {
	block := ir.NewBasicBlock(oldBlock.Name())
	for _, oldInst := range oldBlock.Insts() {
		inst := m.fixInst(oldInst)
		block.AppendInst(inst)
	}
	term := m.fixTerm(oldBlock.Term())
	block.SetTerm(term)
	return block
}

// === [ Instructions ] ========================================================

// fixInst replaces dummy values within the given instruction with their
// corresponding local variables.
func (m dummyMap) fixInst(oldInst instruction.Instruction) instruction.Instruction {
	switch oldInst := oldInst.(type) {
	case *instruction.LocalVarDef:
		return m.fixLocalVarDefInst(oldInst)
	case *instruction.Store:
		return m.fixStoreInst(oldInst)
	case *instruction.Fence:
		return m.fixFenceInst(oldInst)
	default:
		panic(fmt.Sprintf("support for instruction type %T not yet implemented", oldInst))
	}
}

// fixLocalVarDefInst replaces dummy values within the given LocalVarDef
// instruction with their corresponding local variables.
func (m dummyMap) fixLocalVarDefInst(oldInst *instruction.LocalVarDef) *instruction.LocalVarDef {
	name := oldInst.Name()
	oldValInst := oldInst.ValInst()
	valInst := m.fixValueInst(oldValInst)
	inst, err := instruction.NewLocalVarDef(name, valInst)
	if err != nil {
		panic(errutil.Err(err))
	}
	return inst
}

// fixStoreInst replaces dummy values within the given Store instruction with
// their corresponding local variables.
func (m dummyMap) fixStoreInst(oldInst *instruction.Store) *instruction.Store {
	oldVal := oldInst.Val()
	val := m.fixValue(oldVal)
	oldAddr := oldInst.Addr()
	addr := m.fixValue(oldAddr)
	inst, err := instruction.NewStore(val, addr)
	if err != nil {
		panic(errutil.Err(err))
	}
	return inst
}

// fixFenceInst replaces dummy values within the given Fence instruction with
// their corresponding local variables.
func (m dummyMap) fixFenceInst(oldInst *instruction.Fence) *instruction.Fence {
	panic("irx.dummyMap.fixFenceInst: not yet implemented")
}

// === [ Value instructions ] ==================================================

// fixValueInst replaces dummy values within the given value instruction with
// their corresponding local variables.
func (m dummyMap) fixValueInst(oldValInst instruction.ValueInst) instruction.ValueInst {
	panic("irx.dummyMap.fixValueInst: not yet implemented")
}

// --- [ Binary Operations ] ---------------------------------------------------

//    *Add
//    *FAdd
//    *Sub
//    *FSub
//    *Mul
//    *FMul
//    *UDiv
//    *SDiv
//    *FDiv
//    *URem
//    *SRem
//    *FRem

// --- [ Bitwise Binary Operations ] -------------------------------------------

//    *ShL
//    *LShR
//    *AShR
//    *And
//    *Or
//    *Xor

// --- [ Vector Operations ] ---------------------------------------------------

//    *ExtractElement
//    *InsertElement
//    *ShuffleVector

// --- [ Aggregate Operations ] ------------------------------------------------

//    *ExtractValue
//    *InsertValue

// --- [ Memory Access and Addressing Operations ] -----------------------------

//    *Alloca
//    *Load
//    *CmpXchg
//    *AtomicRMW
//    *GetElementPtr

// --- [ Conversion Operations ] -----------------------------------------------

//    *Trunc
//    *ZExt
//    *SExt
//    *FPTrunc
//    *FPExt
//    *FPToUI
//    *FPToSI
//    *UIToFP
//    *SIToFP
//    *PtrToInt
//    *IntToPtr
//    *BitCast
//    *AddrSpaceCast

// --- [ Other Operations ] ----------------------------------------------------

//    *ICmp
//    *FCmp
//    *PHI
//    *Select
//    *Call
//    *VAArg
//    *LandingPad

// === [ Terminators ] =========================================================

//    *Ret
//    *Jmp
//    *Br
//    *Switch
//    *IndirectBr
//    *Invoke
//    *Resume
//    *Unreachable

// fixTerm replaces dummy values within the given terminator with their
// corresponding local variables.
func (m dummyMap) fixTerm(oldTerm instruction.Terminator) instruction.Terminator {
	switch oldTerm := oldTerm.(type) {
	case *instruction.Ret:
		oldVal := oldTerm.Value()
		var val value.Value
		if oldVal != nil {
			val = m.fixValue(oldVal)
		}
		term, err := instruction.NewRet(val)
		if err != nil {
			panic(errutil.Err(err))
		}
		return term
	case *instruction.Jmp:
		panic("irx.dummyMap.fixTerm: Jmp not yet implemented")
	case *instruction.Br:
		panic("irx.dummyMap.fixTerm: Br not yet implemented")
	case *instruction.Switch:
		panic("irx.dummyMap.fixTerm: Switch not yet implemented")
	case *instruction.IndirectBr:
		panic("irx.dummyMap.fixTerm: IndirectBr not yet implemented")
	case *instruction.Invoke:
		panic("irx.dummyMap.fixTerm: Invoke not yet implemented")
	case *instruction.Resume:
		panic("irx.dummyMap.fixTerm: Resume not yet implemented")
	case *instruction.Unreachable:
		panic("irx.dummyMap.fixTerm: Unreachable not yet implemented")
	default:
		panic(fmt.Sprintf("support for terminator type %T not yet implemented", oldTerm))
	}
}

// === [ Values ] ==============================================================

// fixValue replaces dummy values within the given value with their
// corresponding local variables.
func (m dummyMap) fixValue(oldVal value.Value) value.Value {
	switch oldVal := oldVal.(type) {
	case constant.Constant:
		return oldVal
	default:
		panic(fmt.Sprintf("support for value type %T not yet implemented", oldVal))
	}
}
