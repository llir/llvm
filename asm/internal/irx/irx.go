//+build ignore

// TODO: Print value in generic error string.

package irx

import (
	"github.com/llir/llvm/internal/dummy"
	"github.com/llir/llvm/ir"
	"github.com/llir/llvm/ir/constant"
	"github.com/llir/llvm/ir/types"
	"github.com/llir/llvm/ir/value"
	"github.com/pkg/errors"
)

// === [ Modules ] =============================================================

// === [ Type definitions ] ====================================================

// === [ Global variables ] ====================================================

// === [ Functions ] ===========================================================

// NewFunctionDecl returns a new function declaration based on the given
// return type, function name and parameters.
func NewFunctionDecl(ret, name, params interface{}) (*ir.Function, error) {
	r, ok := ret.(types.Type)
	if !ok {
		return nil, errors.Errorf("invalid function return type; expected types.Type, got %T", ret)
	}
	n, ok := name.(*GlobalIdent)
	if !ok {
		return nil, errors.Errorf("invalid function name type; expected *irx.GlobalIdent, got %T", name)
	}
	f := ir.NewFunction(n.name, r)
	switch ps := params.(type) {
	case *Params:
		for _, param := range ps.params {
			f.NewParam(param.Name(), param.Type())
		}
		f.SetVariadic(ps.variadic)
	case nil:
		// no parameters.
	default:
		return nil, errors.Errorf("invalid function parameters type; expected *irx.Params or nil, got %T", params)
	}
	return f, nil
}

// NewFunctionDef returns a new function definition based on the given function
// header and body.
func NewFunctionDef(header, body interface{}) (*ir.Function, error) {
	f, ok := header.(*ir.Function)
	if !ok {
		return nil, errors.Errorf("invalid function header type; expected *ir.Function, got %T", header)
	}
	blocks, ok := body.([]*ir.BasicBlock)
	if !ok {
		return nil, errors.Errorf("invalid function body type; expected []*ir.BasicBlock, got %T", body)
	}
	for _, block := range blocks {
		f.AppendBlock(block)
	}
	return f, nil
}

// === [ Identifiers ] =========================================================

// === [ Types ] ===============================================================

// === [ Values ] ==============================================================

// === [ Constants ] ===========================================================

// === [ Basic blocks ] ========================================================

// NewBasicBlockList returns a new basic block list based on the given basic
// block.
func NewBasicBlockList(block interface{}) ([]*ir.BasicBlock, error) {
	b, ok := block.(*ir.BasicBlock)
	if !ok {
		return nil, errors.Errorf("invalid basic block type; expected *ir.BasicBlock, got %T", block)
	}
	return []*ir.BasicBlock{b}, nil
}

// AppendBasicBlock appends the given basic block to the basic block list.
func AppendBasicBlock(blocks, block interface{}) ([]*ir.BasicBlock, error) {
	bs, ok := blocks.([]*ir.BasicBlock)
	if !ok {
		return nil, errors.Errorf("invalid basic block list type; expected []*ir.BasicBlock, got %T", blocks)
	}
	b, ok := block.(*ir.BasicBlock)
	if !ok {
		return nil, errors.Errorf("invalid basic block type; expected *ir.BasicBlock, got %T", block)
	}
	return append(bs, b), nil
}

// NewBasicBlock returns a new basic block based on the given label name, non-
// branching instructions and terminator.
func NewBasicBlock(name, insts, term interface{}) (*ir.BasicBlock, error) {
	block := ir.NewBlock("")
	switch name := name.(type) {
	case *LabelIdent:
		block.SetName(name.name)
	case nil:
		// unnamed basic block.
	default:
		return nil, errors.Errorf("invalid label name type; expected *irx.LabelIdent or nil, got %T", name)
	}
	var is []ir.Instruction
	switch insts := insts.(type) {
	case []ir.Instruction:
		is = insts
	case nil:
		// no instructions.
	default:
		return nil, errors.Errorf("invalid instruction list type; expected []ir.Instruction, got %T", insts)
	}
	t, ok := term.(ir.Terminator)
	if !ok {
		return nil, errors.Errorf("invalid terminator type; expected ir.Terminator, got %T", term)
	}
	for _, inst := range is {
		block.AppendInst(inst)
	}
	block.SetTerm(t)
	return block, nil
}

// === [ Instructions ] ========================================================

// NewInstructionList returns a new instruction list based on the given
// instruction.
func NewInstructionList(inst interface{}) ([]ir.Instruction, error) {
	i, ok := inst.(ir.Instruction)
	if !ok {
		return nil, errors.Errorf("invalid instruction type; expected ir.Instruction, got %T", inst)
	}
	return []ir.Instruction{i}, nil
}

// AppendInstruction appends the given instruction to the instruction list.
func AppendInstruction(insts, inst interface{}) ([]ir.Instruction, error) {
	is, ok := insts.([]ir.Instruction)
	if !ok {
		return nil, errors.Errorf("invalid instruction list type; expected []ir.Instruction, got %T", insts)
	}
	i, ok := inst.(ir.Instruction)
	if !ok {
		return nil, errors.Errorf("invalid instruction type; expected ir.Instruction, got %T", inst)
	}
	return append(is, i), nil
}

// NewNamedInstruction returns a named instruction based on the given local
// variable name and instruction.
func NewNamedInstruction(name, inst interface{}) (ir.Instruction, error) {
	// namedInstruction represents a namedInstruction instruction.
	type namedInstruction interface {
		ir.Instruction
		value.Named
	}
	n, ok := name.(*LocalIdent)
	if !ok {
		return nil, errors.Errorf("invalid local variable name type; expected *irx.LocalIdent, got %T", name)
	}
	i, ok := inst.(namedInstruction)
	if !ok {
		return nil, errors.Errorf("invalid instruction type; expected namedInstruction, got %T", inst)
	}
	i.SetName(n.name)
	return i, nil
}

// --- [ Binary instructions ] -------------------------------------------------

// NewAddInst returns a new add instruction based on the given type and
// operands.
func NewAddInst(typ, xVal, yVal interface{}) (*ir.InstAdd, error) {
	x, err := NewValue(typ, xVal)
	if err != nil {
		return nil, errors.WithStack(err)
	}
	y, err := NewValue(typ, yVal)
	if err != nil {
		return nil, errors.WithStack(err)
	}
	return ir.NewAdd(x, y), nil
}

// NewFAddInst returns a new fadd instruction based on the given type and
// operands.
func NewFAddInst(typ, xVal, yVal interface{}) (*ir.InstFAdd, error) {
	x, err := NewValue(typ, xVal)
	if err != nil {
		return nil, errors.WithStack(err)
	}
	y, err := NewValue(typ, yVal)
	if err != nil {
		return nil, errors.WithStack(err)
	}
	return ir.NewFAdd(x, y), nil
}

// NewSubInst returns a new sub instruction based on the given type and
// operands.
func NewSubInst(typ, xVal, yVal interface{}) (*ir.InstSub, error) {
	x, err := NewValue(typ, xVal)
	if err != nil {
		return nil, errors.WithStack(err)
	}
	y, err := NewValue(typ, yVal)
	if err != nil {
		return nil, errors.WithStack(err)
	}
	return ir.NewSub(x, y), nil
}

// NewFSubInst returns a new fsub instruction based on the given type and
// operands.
func NewFSubInst(typ, xVal, yVal interface{}) (*ir.InstFSub, error) {
	x, err := NewValue(typ, xVal)
	if err != nil {
		return nil, errors.WithStack(err)
	}
	y, err := NewValue(typ, yVal)
	if err != nil {
		return nil, errors.WithStack(err)
	}
	return ir.NewFSub(x, y), nil
}

// NewMulInst returns a new mul instruction based on the given type and
// operands.
func NewMulInst(typ, xVal, yVal interface{}) (*ir.InstMul, error) {
	x, err := NewValue(typ, xVal)
	if err != nil {
		return nil, errors.WithStack(err)
	}
	y, err := NewValue(typ, yVal)
	if err != nil {
		return nil, errors.WithStack(err)
	}
	return ir.NewMul(x, y), nil
}

// NewFMulInst returns a new fmul instruction based on the given type and
// operands.
func NewFMulInst(typ, xVal, yVal interface{}) (*ir.InstFMul, error) {
	x, err := NewValue(typ, xVal)
	if err != nil {
		return nil, errors.WithStack(err)
	}
	y, err := NewValue(typ, yVal)
	if err != nil {
		return nil, errors.WithStack(err)
	}
	return ir.NewFMul(x, y), nil
}

// NewUDivInst returns a new udiv instruction based on the given type and
// operands.
func NewUDivInst(typ, xVal, yVal interface{}) (*ir.InstUDiv, error) {
	x, err := NewValue(typ, xVal)
	if err != nil {
		return nil, errors.WithStack(err)
	}
	y, err := NewValue(typ, yVal)
	if err != nil {
		return nil, errors.WithStack(err)
	}
	return ir.NewUDiv(x, y), nil
}

// NewSDivInst returns a new sdiv instruction based on the given type and
// operands.
func NewSDivInst(typ, xVal, yVal interface{}) (*ir.InstSDiv, error) {
	x, err := NewValue(typ, xVal)
	if err != nil {
		return nil, errors.WithStack(err)
	}
	y, err := NewValue(typ, yVal)
	if err != nil {
		return nil, errors.WithStack(err)
	}
	return ir.NewSDiv(x, y), nil
}

// NewFDivInst returns a new fdiv instruction based on the given type and
// operands.
func NewFDivInst(typ, xVal, yVal interface{}) (*ir.InstFDiv, error) {
	x, err := NewValue(typ, xVal)
	if err != nil {
		return nil, errors.WithStack(err)
	}
	y, err := NewValue(typ, yVal)
	if err != nil {
		return nil, errors.WithStack(err)
	}
	return ir.NewFDiv(x, y), nil
}

// NewURemInst returns a new urem instruction based on the given type and
// operands.
func NewURemInst(typ, xVal, yVal interface{}) (*ir.InstURem, error) {
	x, err := NewValue(typ, xVal)
	if err != nil {
		return nil, errors.WithStack(err)
	}
	y, err := NewValue(typ, yVal)
	if err != nil {
		return nil, errors.WithStack(err)
	}
	return ir.NewURem(x, y), nil
}

// NewSRemInst returns a new srem instruction based on the given type and
// operands.
func NewSRemInst(typ, xVal, yVal interface{}) (*ir.InstSRem, error) {
	x, err := NewValue(typ, xVal)
	if err != nil {
		return nil, errors.WithStack(err)
	}
	y, err := NewValue(typ, yVal)
	if err != nil {
		return nil, errors.WithStack(err)
	}
	return ir.NewSRem(x, y), nil
}

// NewFRemInst returns a new frem instruction based on the given type and
// operands.
func NewFRemInst(typ, xVal, yVal interface{}) (*ir.InstFRem, error) {
	x, err := NewValue(typ, xVal)
	if err != nil {
		return nil, errors.WithStack(err)
	}
	y, err := NewValue(typ, yVal)
	if err != nil {
		return nil, errors.WithStack(err)
	}
	return ir.NewFRem(x, y), nil
}

// --- [ Bitwise instructions ] ------------------------------------------------

// NewShlInst returns a new shl instruction based on the given type and
// operands.
func NewShlInst(typ, xVal, yVal interface{}) (*ir.InstShl, error) {
	x, err := NewValue(typ, xVal)
	if err != nil {
		return nil, errors.WithStack(err)
	}
	y, err := NewValue(typ, yVal)
	if err != nil {
		return nil, errors.WithStack(err)
	}
	return ir.NewShl(x, y), nil
}

// NewLShrInst returns a new lshr instruction based on the given type and
// operands.
func NewLShrInst(typ, xVal, yVal interface{}) (*ir.InstLShr, error) {
	x, err := NewValue(typ, xVal)
	if err != nil {
		return nil, errors.WithStack(err)
	}
	y, err := NewValue(typ, yVal)
	if err != nil {
		return nil, errors.WithStack(err)
	}
	return ir.NewLShr(x, y), nil
}

// NewAShrInst returns a new ashr instruction based on the given type and
// operands.
func NewAShrInst(typ, xVal, yVal interface{}) (*ir.InstAShr, error) {
	x, err := NewValue(typ, xVal)
	if err != nil {
		return nil, errors.WithStack(err)
	}
	y, err := NewValue(typ, yVal)
	if err != nil {
		return nil, errors.WithStack(err)
	}
	return ir.NewAShr(x, y), nil
}

// NewAndInst returns a new and instruction based on the given type and
// operands.
func NewAndInst(typ, xVal, yVal interface{}) (*ir.InstAnd, error) {
	x, err := NewValue(typ, xVal)
	if err != nil {
		return nil, errors.WithStack(err)
	}
	y, err := NewValue(typ, yVal)
	if err != nil {
		return nil, errors.WithStack(err)
	}
	return ir.NewAnd(x, y), nil
}

// NewOrInst returns a new or instruction based on the given type and
// operands.
func NewOrInst(typ, xVal, yVal interface{}) (*ir.InstOr, error) {
	x, err := NewValue(typ, xVal)
	if err != nil {
		return nil, errors.WithStack(err)
	}
	y, err := NewValue(typ, yVal)
	if err != nil {
		return nil, errors.WithStack(err)
	}
	return ir.NewOr(x, y), nil
}

// NewXorInst returns a new xor instruction based on the given type and
// operands.
func NewXorInst(typ, xVal, yVal interface{}) (*ir.InstXor, error) {
	x, err := NewValue(typ, xVal)
	if err != nil {
		return nil, errors.WithStack(err)
	}
	y, err := NewValue(typ, yVal)
	if err != nil {
		return nil, errors.WithStack(err)
	}
	return ir.NewXor(x, y), nil
}

// --- [ Memory instructions ] -------------------------------------------------

// NewAllocaInst returns a new alloca instruction based on the given element
// type and number of elements.
func NewAllocaInst(elem, nelems interface{}) (*ir.InstAlloca, error) {
	e, ok := elem.(types.Type)
	if !ok {
		return nil, errors.Errorf("invalid element type; expected types.Type, got %T", elem)
	}
	inst := ir.NewAlloca(e)
	switch nelems := nelems.(type) {
	case value.Value:
		inst.SetNElems(nelems)
	case nil:
		// no nelems.
	default:
		return nil, errors.Errorf("invalid number of elements type; expected value.Value or nil, got %T", nelems)
	}
	return inst, nil
}

// NewLoadInst returns a new load instruction based on the given element type,
// source address type and value.
func NewLoadInst(elem, srcTyp, src interface{}) (*dummy.InstLoad, error) {
	e, ok := elem.(types.Type)
	if !ok {
		return nil, errors.Errorf("invalid element type; expected types.Type, got %T", elem)
	}
	s, err := NewValue(srcTyp, src)
	if err != nil {
		return nil, errors.WithStack(err)
	}
	// Store e in *dummy.InstLoad, so that it may be evaluated against
	// inst.Type() after type resolution.
	return dummy.NewLoad(e, s), nil
}

// NewStoreInst returns a new store instruction based on the given element type,
// source address type and value.
func NewStoreInst(srcTyp, srcVal, dstTyp, dstVal interface{}) (*ir.InstStore, error) {
	src, err := NewValue(srcTyp, srcVal)
	if err != nil {
		return nil, errors.WithStack(err)
	}
	dst, err := NewValue(dstTyp, dstVal)
	if err != nil {
		return nil, errors.WithStack(err)
	}
	return ir.NewStore(src, dst), nil
}

// NewGetElementPtrInst returns a new getelementptr instruction based on the
// given element type, source address type and value, and element indices.
func NewGetElementPtrInst(elem, srcTyp, srcVal, indices interface{}) (*dummy.InstGetElementPtr, error) {
	e, ok := elem.(types.Type)
	if !ok {
		return nil, errors.Errorf("invalid element type; expected types.Type, got %T", elem)
	}
	st, ok := srcTyp.(*types.PointerType)
	if !ok {
		return nil, errors.Errorf("invalid source type; expected *types.Pointer, got %T", srcTyp)
	}
	src, err := NewValue(st, srcVal)
	if err != nil {
		return nil, errors.WithStack(err)
	}
	var is []value.Value
	switch indices := indices.(type) {
	case []value.Value:
		is = indices
	case nil:
		// no indices.
	default:
		return nil, errors.Errorf("invalid indices type; expected []value.Value or nil, got %T", indices)
	}
	// Store e in *dummy.InstGetElementPtr, so that it may be evaluated against
	// st.Elem() after type resolution.
	return dummy.NewGetElementPtr(e, src, is...), nil
}

// --- [ Conversion instructions ] ---------------------------------------------

// NewTruncInst returns a new trunc instruction based on the given source value
// and target type.
func NewTruncInst(fromTyp, fromVal, to interface{}) (*ir.InstTrunc, error) {
	from, err := NewValue(fromTyp, fromVal)
	if err != nil {
		return nil, errors.WithStack(err)
	}
	t, ok := to.(types.Type)
	if !ok {
		return nil, errors.Errorf("invalid type; expected types.Type, got %T", to)
	}
	return ir.NewTrunc(from, t), nil
}

// NewZExtInst returns a new zext instruction based on the given source value
// and target type.
func NewZExtInst(fromTyp, fromVal, to interface{}) (*ir.InstZExt, error) {
	from, err := NewValue(fromTyp, fromVal)
	if err != nil {
		return nil, errors.WithStack(err)
	}
	t, ok := to.(types.Type)
	if !ok {
		return nil, errors.Errorf("invalid type; expected types.Type, got %T", to)
	}
	return ir.NewZExt(from, t), nil
}

// NewSExtInst returns a new sext instruction based on the given source value
// and target type.
func NewSExtInst(fromTyp, fromVal, to interface{}) (*ir.InstSExt, error) {
	from, err := NewValue(fromTyp, fromVal)
	if err != nil {
		return nil, errors.WithStack(err)
	}
	t, ok := to.(types.Type)
	if !ok {
		return nil, errors.Errorf("invalid type; expected types.Type, got %T", to)
	}
	return ir.NewSExt(from, t), nil
}

// NewFPTruncInst returns a new fptrunc instruction based on the given source value
// and target type.
func NewFPTruncInst(fromTyp, fromVal, to interface{}) (*ir.InstFPTrunc, error) {
	from, err := NewValue(fromTyp, fromVal)
	if err != nil {
		return nil, errors.WithStack(err)
	}
	t, ok := to.(types.Type)
	if !ok {
		return nil, errors.Errorf("invalid type; expected types.Type, got %T", to)
	}
	return ir.NewFPTrunc(from, t), nil
}

// NewFPExtInst returns a new fpext instruction based on the given source value
// and target type.
func NewFPExtInst(fromTyp, fromVal, to interface{}) (*ir.InstFPExt, error) {
	from, err := NewValue(fromTyp, fromVal)
	if err != nil {
		return nil, errors.WithStack(err)
	}
	t, ok := to.(types.Type)
	if !ok {
		return nil, errors.Errorf("invalid type; expected types.Type, got %T", to)
	}
	return ir.NewFPExt(from, t), nil
}

// NewFPToUIInst returns a new fptoui instruction based on the given source value
// and target type.
func NewFPToUIInst(fromTyp, fromVal, to interface{}) (*ir.InstFPToUI, error) {
	from, err := NewValue(fromTyp, fromVal)
	if err != nil {
		return nil, errors.WithStack(err)
	}
	t, ok := to.(types.Type)
	if !ok {
		return nil, errors.Errorf("invalid type; expected types.Type, got %T", to)
	}
	return ir.NewFPToUI(from, t), nil
}

// NewFPToSIInst returns a new fptosi instruction based on the given source value
// and target type.
func NewFPToSIInst(fromTyp, fromVal, to interface{}) (*ir.InstFPToSI, error) {
	from, err := NewValue(fromTyp, fromVal)
	if err != nil {
		return nil, errors.WithStack(err)
	}
	t, ok := to.(types.Type)
	if !ok {
		return nil, errors.Errorf("invalid type; expected types.Type, got %T", to)
	}
	return ir.NewFPToSI(from, t), nil
}

// NewUIToFPInst returns a new uitofp instruction based on the given source value
// and target type.
func NewUIToFPInst(fromTyp, fromVal, to interface{}) (*ir.InstUIToFP, error) {
	from, err := NewValue(fromTyp, fromVal)
	if err != nil {
		return nil, errors.WithStack(err)
	}
	t, ok := to.(types.Type)
	if !ok {
		return nil, errors.Errorf("invalid type; expected types.Type, got %T", to)
	}
	return ir.NewUIToFP(from, t), nil
}

// NewSIToFPInst returns a new sitofp instruction based on the given source value
// and target type.
func NewSIToFPInst(fromTyp, fromVal, to interface{}) (*ir.InstSIToFP, error) {
	from, err := NewValue(fromTyp, fromVal)
	if err != nil {
		return nil, errors.WithStack(err)
	}
	t, ok := to.(types.Type)
	if !ok {
		return nil, errors.Errorf("invalid type; expected types.Type, got %T", to)
	}
	return ir.NewSIToFP(from, t), nil
}

// NewPtrToIntInst returns a new ptrtoint instruction based on the given source value
// and target type.
func NewPtrToIntInst(fromTyp, fromVal, to interface{}) (*ir.InstPtrToInt, error) {
	from, err := NewValue(fromTyp, fromVal)
	if err != nil {
		return nil, errors.WithStack(err)
	}
	t, ok := to.(types.Type)
	if !ok {
		return nil, errors.Errorf("invalid type; expected types.Type, got %T", to)
	}
	return ir.NewPtrToInt(from, t), nil
}

// NewIntToPtrInst returns a new inttoptr instruction based on the given source value
// and target type.
func NewIntToPtrInst(fromTyp, fromVal, to interface{}) (*ir.InstIntToPtr, error) {
	from, err := NewValue(fromTyp, fromVal)
	if err != nil {
		return nil, errors.WithStack(err)
	}
	t, ok := to.(types.Type)
	if !ok {
		return nil, errors.Errorf("invalid type; expected types.Type, got %T", to)
	}
	return ir.NewIntToPtr(from, t), nil
}

// NewBitCastInst returns a new bitcast instruction based on the given source value
// and target type.
func NewBitCastInst(fromTyp, fromVal, to interface{}) (*ir.InstBitCast, error) {
	from, err := NewValue(fromTyp, fromVal)
	if err != nil {
		return nil, errors.WithStack(err)
	}
	t, ok := to.(types.Type)
	if !ok {
		return nil, errors.Errorf("invalid type; expected types.Type, got %T", to)
	}
	return ir.NewBitCast(from, t), nil
}

// NewAddrSpaceCastInst returns a new addrspacecast instruction based on the given source value
// and target type.
func NewAddrSpaceCastInst(fromTyp, fromVal, to interface{}) (*ir.InstAddrSpaceCast, error) {
	from, err := NewValue(fromTyp, fromVal)
	if err != nil {
		return nil, errors.WithStack(err)
	}
	t, ok := to.(types.Type)
	if !ok {
		return nil, errors.Errorf("invalid type; expected types.Type, got %T", to)
	}
	return ir.NewAddrSpaceCast(from, t), nil
}

// --- [ Other instructions ] --------------------------------------------------

// NewICmpInst returns a new icmp instruction based on the given integer
// condition code, type and operands.
func NewICmpInst(cond, typ, xVal, yVal interface{}) (*ir.InstICmp, error) {
	c, ok := cond.(ir.IntPred)
	if !ok {
		return nil, errors.Errorf("invalid integer predicate type; expected ir.IntPred, got %T", cond)
	}
	x, err := NewValue(typ, xVal)
	if err != nil {
		return nil, errors.WithStack(err)
	}
	y, err := NewValue(typ, yVal)
	if err != nil {
		return nil, errors.WithStack(err)
	}
	return ir.NewICmp(c, x, y), nil
}

// NewFCmpInst returns a new fcmp instruction based on the given floating-point
// condition code, type and operands.
func NewFCmpInst(cond, typ, xVal, yVal interface{}) (*ir.InstFCmp, error) {
	c, ok := cond.(ir.FloatPred)
	if !ok {
		return nil, errors.Errorf("invalid floating-point predicate type; expected ir.FloatPred, got %T", cond)
	}
	x, err := NewValue(typ, xVal)
	if err != nil {
		return nil, errors.WithStack(err)
	}
	y, err := NewValue(typ, yVal)
	if err != nil {
		return nil, errors.WithStack(err)
	}
	return ir.NewFCmp(c, x, y), nil
}

// NewPhiInst returns a new phi instruction based on the given incoming values.
func NewPhiInst(typ, incs interface{}) (*dummy.InstPhi, error) {
	t, ok := typ.(types.Type)
	if !ok {
		return nil, errors.Errorf("invalid type; expected types.Type, got %T", typ)
	}
	is, ok := incs.([]*dummy.Incoming)
	if !ok {
		return nil, errors.Errorf("invalid incoming value list type; expected []*dummy.Incoming, got %T", incs)
	}
	for _, inc := range is {
		x, err := NewValue(typ, inc.X())
		if err != nil {
			return nil, errors.WithStack(err)
		}
		inc.SetX(x)
	}
	return dummy.NewPhi(t, is...), nil
}

// NewIncomingList returns a new incoming value list based on the given incoming
// value.
func NewIncomingList(inc interface{}) ([]*dummy.Incoming, error) {
	i, ok := inc.(*dummy.Incoming)
	if !ok {
		return nil, errors.Errorf("invalid incoming value type; expected *dummy.Incoming, got %T", inc)
	}
	return []*dummy.Incoming{i}, nil
}

// AppendIncoming appends the given incoming value to the incoming value list.
func AppendIncoming(incs, inc interface{}) ([]*dummy.Incoming, error) {
	is, ok := incs.([]*dummy.Incoming)
	if !ok {
		return nil, errors.Errorf("invalid incoming value list type; expected []*dummy.Incoming, got %T", incs)
	}
	i, ok := inc.(*dummy.Incoming)
	if !ok {
		return nil, errors.Errorf("invalid incoming value type; expected *dummy.Incoming, got %T", inc)
	}
	return append(is, i), nil
}

// NewIncoming returns a new incoming value based on the given value and
// predecessor basic block.
func NewIncoming(x, pred interface{}) (*dummy.Incoming, error) {
	p, ok := pred.(*LocalIdent)
	if !ok {
		return nil, errors.Errorf("invalid predecessor type; expected *irx.LocalIdent, got %T", pred)
	}
	return dummy.NewIncoming(x, p.name), nil
}

// NewSelect returns a new select instruction based on the given selection
// condition type and value, and operands.
func NewSelectInst(condTyp, condVal, xTyp, xVal, yTyp, yVal interface{}) (*ir.InstSelect, error) {
	cond, err := NewValue(condTyp, condVal)
	if err != nil {
		return nil, errors.WithStack(err)
	}
	x, err := NewValue(xTyp, xVal)
	if err != nil {
		return nil, errors.WithStack(err)
	}
	y, err := NewValue(yTyp, yVal)
	if err != nil {
		return nil, errors.WithStack(err)
	}
	return ir.NewSelect(cond, x, y), nil
}

// NewCallInst returns a new call instruction based on the given return type,
// callee name, and function arguments.
func NewCallInst(retTyp, callee, args interface{}) (*dummy.InstCall, error) {
	r, ok := retTyp.(types.Type)
	if !ok {
		return nil, errors.Errorf("invalid return type; expected types.Type, got %T", retTyp)
	}
	var name string
	var calleeLocal bool
	switch callee := callee.(type) {
	case *GlobalIdent:
		name = callee.name
	case *LocalIdent:
		name = callee.name
		calleeLocal = true
	default:
		return nil, errors.Errorf("invalid callee type; expected *irx.GlobalIdent, got %T", callee)
	}
	var as []value.Value
	switch args := args.(type) {
	case []value.Value:
		as = args
	case nil:
		// no arguments.
	default:
		return nil, errors.Errorf("invalid function arguments type; expected []value.Value or nil, got %T", args)
	}
	return dummy.NewCall(r, name, calleeLocal, as...), nil
}

// === [ Terminators ] =========================================================

// NewRetTerm returns a new ret terminator based on the given return type and
// value.
func NewRetTerm(typ, val interface{}) (*ir.TermRet, error) {
	v, err := NewValue(typ, val)
	if err != nil {
		return nil, errors.WithStack(err)
	}
	return ir.NewRet(v), nil
}

// NewBrTerm returns a new unconditional br terminator based on the given target
// branch.
func NewBrTerm(target interface{}) (*dummy.TermBr, error) {
	t, ok := target.(*LocalIdent)
	if !ok {
		return nil, errors.Errorf("invalid target branch type; expected *irx.LocalIdent, got %T", target)
	}
	return dummy.NewBr(t.name), nil
}

// NewCondBrTerm returns a new conditional br terminator based on the given
// branching condition type and value, and conditional target branches.
func NewCondBrTerm(condTyp, condVal, targetTrue, targetFalse interface{}) (*dummy.TermCondBr, error) {
	cond, err := NewValue(condTyp, condVal)
	if err != nil {
		return nil, errors.WithStack(err)
	}
	tTrue, ok := targetTrue.(*LocalIdent)
	if !ok {
		return nil, errors.Errorf("invalid true target branch type; expected *irx.LocalIdent, got %T", targetTrue)
	}
	tFalse, ok := targetFalse.(*LocalIdent)
	if !ok {
		return nil, errors.Errorf("invalid true target branch type; expected *irx.LocalIdent, got %T", targetFalse)
	}
	return dummy.NewCondBr(cond, tTrue.name, tFalse.name), nil
}

// NewSwitchTerm returns a new switch terminator based on the given control
// variable type and value, default target branch and switch cases.
func NewSwitchTerm(xTyp, xVal, targetDefault, cases interface{}) (*dummy.TermSwitch, error) {
	x, err := NewValue(xTyp, xVal)
	if err != nil {
		return nil, errors.WithStack(err)
	}
	tDefault, ok := targetDefault.(*LocalIdent)
	if !ok {
		return nil, errors.Errorf("invalid default target branch type; expected *irx.LocalIdent, got %T", targetDefault)
	}
	var cs []*dummy.Case
	switch cases := cases.(type) {
	case []*dummy.Case:
		cs = cases
	case nil:
		// no cases.
	default:
		return nil, errors.Errorf("invalid switch cases type; expected []*dummy.Case or nil, got %T", cases)
	}
	return dummy.NewSwitch(x, tDefault.name, cs...), nil
}

// NewCaseList returns a new switch case list based on the given case.
func NewCaseList(switchCase interface{}) ([]*dummy.Case, error) {
	c, ok := switchCase.(*dummy.Case)
	if !ok {
		return nil, errors.Errorf("invalid switch case type; expected *dummy.Case, got %T", switchCase)
	}
	return []*dummy.Case{c}, nil
}

// AppendCase appends the given case to the switch case list.
func AppendCase(cases, switchCase interface{}) ([]*dummy.Case, error) {
	cs, ok := cases.([]*dummy.Case)
	if !ok {
		return nil, errors.Errorf("invalid switch case list type; expected []*dummy.Case, got %T", cases)
	}
	c, ok := switchCase.(*dummy.Case)
	if !ok {
		return nil, errors.Errorf("invalid switch case type; expected *dummy.Case, got %T", switchCase)
	}
	return append(cs, c), nil
}

// NewCase returns a new switch case based on the given case comparand and
// target branch.
func NewCase(xTyp, xVal, target interface{}) (*dummy.Case, error) {
	xValue, err := NewValue(xTyp, xVal)
	if err != nil {
		return nil, errors.WithStack(err)
	}
	x, ok := xValue.(*constant.Int)
	if !ok {
		return nil, errors.Errorf("invalid case comparand type; expected *constant.Int, got %T", xValue)
	}
	t, ok := target.(*LocalIdent)
	if !ok {
		return nil, errors.Errorf("invalid target branch type; expected *irx.LocalIdent, got %T", target)
	}
	return dummy.NewCase(x, t.name), nil
}
