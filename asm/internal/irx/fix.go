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

// set maps name to the given value.
func (m dummyMap) set(name string, val value.Value) {
	if old, ok := m[name]; ok {
		panic(fmt.Sprintf("mapping for %q already present; old value %v, new value %v", name, old, val))
	}
	m[name] = val
}

// get returns the value for the given name.
func (m dummyMap) get(name string) value.Value {
	val, ok := m[name]
	if !ok {
		panic(fmt.Sprintf("unable to locate mapping for %q", name))
	}
	return val
}

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

	// Add mapping for basic blocks so they may be forward-referenced from
	// instructions.
	for _, oldBlock := range oldFunc.Blocks() {
		name := oldBlock.Name()
		block := ir.NewBasicBlock(name)
		m.set(name, block)
	}

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
	name := oldBlock.Name()
	block := m.get(name)
	b, ok := block.(*ir.BasicBlock)
	if !ok {
		panic(fmt.Sprintf("invalid basic block type; expected *ir.BasicBlock, got %T", block))
	}
	for _, oldInst := range oldBlock.Insts() {
		inst := m.fixInst(oldInst)
		b.AppendInst(inst)
	}
	term := m.fixTerm(oldBlock.Term())
	b.SetTerm(term)
	return b
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
	m.set(name, inst)
	return inst
}

// fixStoreInst replaces dummy values within the given Store instruction with
// their corresponding local variables.
func (m dummyMap) fixStoreInst(oldInst *instruction.Store) *instruction.Store {
	src := m.fixValue(oldInst.Src())
	dstAddr := m.fixValue(oldInst.DstAddr())
	inst, err := instruction.NewStore(src, dstAddr)
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
	switch oldValInst := oldValInst.(type) {
	// Binary Operations
	case *instruction.Add:
		x := m.fixValue(oldValInst.X())
		y := m.fixValue(oldValInst.Y())
		inst, err := instruction.NewAdd(x, y)
		if err != nil {
			panic(errutil.Err(err))
		}
		return inst
	case *instruction.FAdd:
		x := m.fixValue(oldValInst.X())
		y := m.fixValue(oldValInst.Y())
		inst, err := instruction.NewFAdd(x, y)
		if err != nil {
			panic(errutil.Err(err))
		}
		return inst
	case *instruction.Sub:
		x := m.fixValue(oldValInst.X())
		y := m.fixValue(oldValInst.Y())
		inst, err := instruction.NewSub(x, y)
		if err != nil {
			panic(errutil.Err(err))
		}
		return inst
	case *instruction.FSub:
		x := m.fixValue(oldValInst.X())
		y := m.fixValue(oldValInst.Y())
		inst, err := instruction.NewFSub(x, y)
		if err != nil {
			panic(errutil.Err(err))
		}
		return inst
	case *instruction.Mul:
		x := m.fixValue(oldValInst.X())
		y := m.fixValue(oldValInst.Y())
		inst, err := instruction.NewMul(x, y)
		if err != nil {
			panic(errutil.Err(err))
		}
		return inst
	case *instruction.FMul:
		x := m.fixValue(oldValInst.X())
		y := m.fixValue(oldValInst.Y())
		inst, err := instruction.NewFMul(x, y)
		if err != nil {
			panic(errutil.Err(err))
		}
		return inst
	case *instruction.UDiv:
		x := m.fixValue(oldValInst.X())
		y := m.fixValue(oldValInst.Y())
		inst, err := instruction.NewUDiv(x, y)
		if err != nil {
			panic(errutil.Err(err))
		}
		return inst
	case *instruction.SDiv:
		x := m.fixValue(oldValInst.X())
		y := m.fixValue(oldValInst.Y())
		inst, err := instruction.NewSDiv(x, y)
		if err != nil {
			panic(errutil.Err(err))
		}
		return inst
	case *instruction.FDiv:
		x := m.fixValue(oldValInst.X())
		y := m.fixValue(oldValInst.Y())
		inst, err := instruction.NewFDiv(x, y)
		if err != nil {
			panic(errutil.Err(err))
		}
		return inst
	case *instruction.URem:
		x := m.fixValue(oldValInst.X())
		y := m.fixValue(oldValInst.Y())
		inst, err := instruction.NewURem(x, y)
		if err != nil {
			panic(errutil.Err(err))
		}
		return inst
	case *instruction.SRem:
		x := m.fixValue(oldValInst.X())
		y := m.fixValue(oldValInst.Y())
		inst, err := instruction.NewSRem(x, y)
		if err != nil {
			panic(errutil.Err(err))
		}
		return inst
	case *instruction.FRem:
		x := m.fixValue(oldValInst.X())
		y := m.fixValue(oldValInst.Y())
		inst, err := instruction.NewFRem(x, y)
		if err != nil {
			panic(errutil.Err(err))
		}
		return inst

	// Bitwise Binary Operations
	case *instruction.ShL:
		x := m.fixValue(oldValInst.X())
		y := m.fixValue(oldValInst.Y())
		inst, err := instruction.NewShL(x, y)
		if err != nil {
			panic(errutil.Err(err))
		}
		return inst
	case *instruction.LShR:
		x := m.fixValue(oldValInst.X())
		y := m.fixValue(oldValInst.Y())
		inst, err := instruction.NewLShR(x, y)
		if err != nil {
			panic(errutil.Err(err))
		}
		return inst
	case *instruction.AShR:
		x := m.fixValue(oldValInst.X())
		y := m.fixValue(oldValInst.Y())
		inst, err := instruction.NewAShR(x, y)
		if err != nil {
			panic(errutil.Err(err))
		}
		return inst
	case *instruction.And:
		x := m.fixValue(oldValInst.X())
		y := m.fixValue(oldValInst.Y())
		inst, err := instruction.NewAnd(x, y)
		if err != nil {
			panic(errutil.Err(err))
		}
		return inst
	case *instruction.Or:
		x := m.fixValue(oldValInst.X())
		y := m.fixValue(oldValInst.Y())
		inst, err := instruction.NewOr(x, y)
		if err != nil {
			panic(errutil.Err(err))
		}
		return inst
	case *instruction.Xor:
		x := m.fixValue(oldValInst.X())
		y := m.fixValue(oldValInst.Y())
		inst, err := instruction.NewXor(x, y)
		if err != nil {
			panic(errutil.Err(err))
		}
		return inst

	// Vector Operations
	case *instruction.ExtractElement:
		panic("irx.dummyMap.fixValueInst: ExtractElement not yet implemented")
	case *instruction.InsertElement:
		panic("irx.dummyMap.fixValueInst: InsertElement not yet implemented")
	case *instruction.ShuffleVector:
		panic("irx.dummyMap.fixValueInst: ShuffleVector not yet implemented")

	// Aggregate Operations
	case *instruction.ExtractValue:
		panic("irx.dummyMap.fixValueInst: ExtractValue not yet implemented")
	case *instruction.InsertValue:
		panic("irx.dummyMap.fixValueInst: InsertValue not yet implemented")

	// Memory Access and Addressing Operations
	case *instruction.Alloca:
		// Nothing to do; alloca contains no dummy values.
		return oldValInst
	case *instruction.Load:
		srcAddr := m.fixValue(oldValInst.SrcAddr())
		inst, err := instruction.NewLoad(srcAddr)
		if err != nil {
			panic(errutil.Err(err))
		}
		return inst
	case *instruction.CmpXchg:
		panic("irx.dummyMap.fixValueInst: CmpXchg not yet implemented")
	case *instruction.AtomicRMW:
		panic("irx.dummyMap.fixValueInst: AtomicRMW not yet implemented")
	case *instruction.GetElementPtr:
		srcAddr := m.fixValue(oldValInst.SrcAddr())
		var indices []value.Value
		for _, oldIndex := range oldValInst.Indices() {
			index := m.fixValue(oldIndex)
			indices = append(indices, index)
		}
		inst, err := instruction.NewGetElementPtr(srcAddr, indices)
		if err != nil {
			panic(errutil.Err(err))
		}
		return inst

	// Conversion Operations
	case *instruction.Trunc:
		from := m.fixValue(oldValInst.From())
		to := oldValInst.RetType()
		inst, err := instruction.NewTrunc(from, to)
		if err != nil {
			panic(errutil.Err(err))
		}
		return inst
	case *instruction.ZExt:
		from := m.fixValue(oldValInst.From())
		to := oldValInst.RetType()
		inst, err := instruction.NewZExt(from, to)
		if err != nil {
			panic(errutil.Err(err))
		}
		return inst
	case *instruction.SExt:
		from := m.fixValue(oldValInst.From())
		to := oldValInst.RetType()
		inst, err := instruction.NewSExt(from, to)
		if err != nil {
			panic(errutil.Err(err))
		}
		return inst
	case *instruction.FPTrunc:
		from := m.fixValue(oldValInst.From())
		to := oldValInst.RetType()
		inst, err := instruction.NewFPTrunc(from, to)
		if err != nil {
			panic(errutil.Err(err))
		}
		return inst
	case *instruction.FPExt:
		from := m.fixValue(oldValInst.From())
		to := oldValInst.RetType()
		inst, err := instruction.NewFPExt(from, to)
		if err != nil {
			panic(errutil.Err(err))
		}
		return inst
	case *instruction.FPToUI:
		from := m.fixValue(oldValInst.From())
		to := oldValInst.RetType()
		inst, err := instruction.NewFPToUI(from, to)
		if err != nil {
			panic(errutil.Err(err))
		}
		return inst
	case *instruction.FPToSI:
		from := m.fixValue(oldValInst.From())
		to := oldValInst.RetType()
		inst, err := instruction.NewFPToSI(from, to)
		if err != nil {
			panic(errutil.Err(err))
		}
		return inst
	case *instruction.UIToFP:
		from := m.fixValue(oldValInst.From())
		to := oldValInst.RetType()
		inst, err := instruction.NewUIToFP(from, to)
		if err != nil {
			panic(errutil.Err(err))
		}
		return inst
	case *instruction.SIToFP:
		from := m.fixValue(oldValInst.From())
		to := oldValInst.RetType()
		inst, err := instruction.NewSIToFP(from, to)
		if err != nil {
			panic(errutil.Err(err))
		}
		return inst
	case *instruction.PtrToInt:
		from := m.fixValue(oldValInst.From())
		to := oldValInst.RetType()
		inst, err := instruction.NewPtrToInt(from, to)
		if err != nil {
			panic(errutil.Err(err))
		}
		return inst
	case *instruction.IntToPtr:
		from := m.fixValue(oldValInst.From())
		to := oldValInst.RetType()
		inst, err := instruction.NewIntToPtr(from, to)
		if err != nil {
			panic(errutil.Err(err))
		}
		return inst
	case *instruction.BitCast:
		from := m.fixValue(oldValInst.From())
		to := oldValInst.RetType()
		inst, err := instruction.NewBitCast(from, to)
		if err != nil {
			panic(errutil.Err(err))
		}
		return inst
	case *instruction.AddrSpaceCast:
		from := m.fixValue(oldValInst.From())
		to := oldValInst.RetType()
		inst, err := instruction.NewAddrSpaceCast(from, to)
		if err != nil {
			panic(errutil.Err(err))
		}
		return inst

	// Other Operations
	case *instruction.ICmp:
		cond := oldValInst.Cond()
		x := m.fixValue(oldValInst.X())
		y := m.fixValue(oldValInst.Y())
		inst, err := instruction.NewICmp(cond, x, y)
		if err != nil {
			panic(errutil.Err(err))
		}
		return inst
	case *instruction.FCmp:
		panic("irx.dummyMap.fixValueInst: FCmp not yet implemented")
	case *instruction.PHI:
		oldIncs := oldValInst.Incs()
		var incs []*instruction.Incoming
		for _, oldInc := range oldIncs {
			val := m.fixValue(oldInc.Value())
			pred := m.fixNamedValue(oldInc.Pred())
			inc, err := instruction.NewIncoming(val, pred)
			if err != nil {
				panic(errutil.Err(err))
			}
			incs = append(incs, inc)
		}
		inst, err := instruction.NewPHI(incs)
		if err != nil {
			panic(errutil.Err(err))
		}
		return inst
	case *instruction.Select:
		cond := m.fixValue(oldValInst.Cond())
		x := m.fixValue(oldValInst.X())
		y := m.fixValue(oldValInst.Y())
		inst, err := instruction.NewSelect(cond, x, y)
		if err != nil {
			panic(errutil.Err(err))
		}
		return inst
	case *instruction.Call:
		result := oldValInst.RetType()
		// TODO: Fix value of callee if the type of Callee changes from string to
		// value.Value.
		callee := oldValInst.Callee()
		var args []value.Value
		for _, oldArg := range oldValInst.Args() {
			arg := m.fixValue(oldArg)
			args = append(args, arg)
		}
		inst, err := instruction.NewCall(result, callee, args)
		if err != nil {
			panic(errutil.Err(err))
		}
		return inst
	case *instruction.VAArg:
		panic("irx.dummyMap.fixValueInst: VAArg not yet implemented")
	case *instruction.LandingPad:
		panic("irx.dummyMap.fixValueInst: LandingPad not yet implemented")
	case *instruction.CatchPad:
		panic("irx.dummyMap.fixValueInst: CatchPad not yet implemented")
	case *instruction.CleanupPad:
		panic("irx.dummyMap.fixValueInst: CleanupPad not yet implemented")
	default:
		panic("irx.dummyMap.fixValueInst: not yet implemented")
	}
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
		target := m.fixNamedValue(oldTerm.Target())
		term, err := instruction.NewJmp(target)
		if err != nil {
			panic(errutil.Err(err))
		}
		return term
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
	case *LocalDummy:
		return m.get(oldVal.name)
	case constant.Constant:
		return oldVal
	default:
		panic(fmt.Sprintf("support for value type %T not yet implemented", oldVal))
	}
}

// fixNamedValue replaces named dummy values within the given value with their
// corresponding local variables.
func (m dummyMap) fixNamedValue(oldVal value.NamedValue) value.NamedValue {
	switch oldVal := oldVal.(type) {
	case *LocalDummy:
		val := m.get(oldVal.name)
		v, ok := val.(value.NamedValue)
		if !ok {
			panic(fmt.Sprintf("invalid type of named value; expected value.NamedValue, got %T", val))
		}
		return v
	default:
		panic(fmt.Sprintf("support for named value type %T not yet implemented", oldVal))
	}
}
