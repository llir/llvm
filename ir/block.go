package ir

import (
	"bytes"
	"fmt"

	"github.com/llir/llvm/ir/internal/enc"
	"github.com/llir/llvm/ir/types"
	"github.com/llir/llvm/ir/value"
)

// A BasicBlock represents an LLVM IR basic block, which consists of a sequence
// of non-branching instructions, terminated by a control flow instruction (e.g.
// br or ret).
//
// Basic blocks may be referenced from terminators (e.g. br), and are thus
// considered LLVM IR values of label type.
type BasicBlock struct {
	// Parent function of the basic block.
	parent *Function
	// Label name of the basic block; or empty if unnamed basic block.
	name string
	// Non-branching instructions of the basic block.
	insts []Instruction
	// Terminator of the basic block.
	term Terminator
}

// NewBlock returns a new basic block based on the given label name. An empty
// label name indicates an unnamed basic block.
func NewBlock(name string) *BasicBlock {
	return &BasicBlock{name: name}
}

// Type returns the type of the basic block.
func (block *BasicBlock) Type() types.Type {
	return types.Label
}

// Ident returns the identifier associated with the basic block.
func (block *BasicBlock) Ident() string {
	return enc.Local(block.name)
}

// SetIdent sets the identifier associated with the basic block.
func (block *BasicBlock) SetIdent(ident string) {
	block.name = ident
}

// LLVMString returns the LLVM syntax representation of the basic block.
func (block *BasicBlock) LLVMString() string {
	buf := &bytes.Buffer{}
	fmt.Fprintf(buf, "%s:\n", enc.Escape(block.name))
	for _, i := range block.Insts() {
		fmt.Fprintf(buf, "\t%s\n", i.LLVMString())
	}
	fmt.Fprintf(buf, "\t%s", block.Term().LLVMString())
	return buf.String()
}

// Parent returns the parent function of the basic block.
func (block *BasicBlock) Parent() *Function {
	return block.parent
}

// SetParent sets the parent function of the basic block.
func (block *BasicBlock) SetParent(parent *Function) {
	block.parent = parent
}

// Insts returns the non-branching instructions of the basic block.
func (block *BasicBlock) Insts() []Instruction {
	return block.insts
}

// Term returns the terminator of the basic block.
func (block *BasicBlock) Term() Terminator {
	return block.term
}

// SetTerm sets the terminator of the basic block.
func (block *BasicBlock) SetTerm(t Terminator) {
	if t, ok := t.(parentSetter); ok {
		t.SetParent(block)
	}
	block.term = t
}

// AppendInst appends the given instruction to the basic block.
func (block *BasicBlock) AppendInst(i Instruction) {
	if i, ok := i.(parentSetter); ok {
		i.SetParent(block)
	}
	block.insts = append(block.insts, i)
}

// --- [ Binary instructions ] -------------------------------------------------

// NewAdd appends a new add instruction to the basic block based on the given
// operands.
func (block *BasicBlock) NewAdd(x, y value.Value) *InstAdd {
	i := NewAdd(x, y)
	block.AppendInst(i)
	return i
}

// NewFAdd appends a new fadd instruction to the basic block based on the given
// operands.
func (block *BasicBlock) NewFAdd(x, y value.Value) *InstFAdd {
	i := NewFAdd(x, y)
	block.AppendInst(i)
	return i
}

// NewSub appends a new sub instruction to the basic block based on the given
// operands.
func (block *BasicBlock) NewSub(x, y value.Value) *InstSub {
	i := NewSub(x, y)
	block.AppendInst(i)
	return i
}

// NewFSub appends a new fsub instruction to the basic block based on the given
// operands.
func (block *BasicBlock) NewFSub(x, y value.Value) *InstFSub {
	i := NewFSub(x, y)
	block.AppendInst(i)
	return i
}

// NewMul appends a new mul instruction to the basic block based on the given
// operands.
func (block *BasicBlock) NewMul(x, y value.Value) *InstMul {
	i := NewMul(x, y)
	block.AppendInst(i)
	return i
}

// NewFMul appends a new fmul instruction to the basic block based on the given
// operands.
func (block *BasicBlock) NewFMul(x, y value.Value) *InstFMul {
	i := NewFMul(x, y)
	block.AppendInst(i)
	return i
}

// NewUDiv appends a new udiv instruction to the basic block based on the given
// operands.
func (block *BasicBlock) NewUDiv(x, y value.Value) *InstUDiv {
	i := NewUDiv(x, y)
	block.AppendInst(i)
	return i
}

// NewSDiv appends a new sdiv instruction to the basic block based on the given
// operands.
func (block *BasicBlock) NewSDiv(x, y value.Value) *InstSDiv {
	i := NewSDiv(x, y)
	block.AppendInst(i)
	return i
}

// NewFDiv appends a new fdiv instruction to the basic block based on the given
// operands.
func (block *BasicBlock) NewFDiv(x, y value.Value) *InstFDiv {
	i := NewFDiv(x, y)
	block.AppendInst(i)
	return i
}

// NewURem appends a new urem instruction to the basic block based on the given
// operands.
func (block *BasicBlock) NewURem(x, y value.Value) *InstURem {
	i := NewURem(x, y)
	block.AppendInst(i)
	return i
}

// NewSRem appends a new srem instruction to the basic block based on the given
// operands.
func (block *BasicBlock) NewSRem(x, y value.Value) *InstSRem {
	i := NewSRem(x, y)
	block.AppendInst(i)
	return i
}

// NewFRem appends a new frem instruction to the basic block based on the given
// operands.
func (block *BasicBlock) NewFRem(x, y value.Value) *InstFRem {
	i := NewFRem(x, y)
	block.AppendInst(i)
	return i
}

// --- [ Bitwise instructions ] ------------------------------------------------

// NewShL appends a new shl instruction to the basic block based on the given
// operands.
func (block *BasicBlock) NewShL(x, y value.Value) *InstShL {
	i := NewShL(x, y)
	block.AppendInst(i)
	return i
}

// NewLShR appends a new lshr instruction to the basic block based on the given
// operands.
func (block *BasicBlock) NewLShR(x, y value.Value) *InstLShR {
	i := NewLShR(x, y)
	block.AppendInst(i)
	return i
}

// NewAShR appends a new ashr instruction to the basic block based on the given
// operands.
func (block *BasicBlock) NewAShR(x, y value.Value) *InstAShR {
	i := NewAShR(x, y)
	block.AppendInst(i)
	return i
}

// NewAnd appends a new and instruction to the basic block based on the given
// operands.
func (block *BasicBlock) NewAnd(x, y value.Value) *InstAnd {
	i := NewAnd(x, y)
	block.AppendInst(i)
	return i
}

// NewOr appends a new or instruction to the basic block based on the given
// operands.
func (block *BasicBlock) NewOr(x, y value.Value) *InstOr {
	i := NewOr(x, y)
	block.AppendInst(i)
	return i
}

// NewXor appends a new xor instruction to the basic block based on the given
// operands.
func (block *BasicBlock) NewXor(x, y value.Value) *InstXor {
	i := NewXor(x, y)
	block.AppendInst(i)
	return i
}

// --- [ Vector instructions ] -------------------------------------------------

// --- [ Aggregate instructions ] ----------------------------------------------

// --- [ Memory instructions ] -------------------------------------------------

// NewAlloca appends a new alloca instruction to the basic block based on the
// given element type.
func (block *BasicBlock) NewAlloca(elem types.Type) *InstAlloca {
	i := NewAlloca(elem)
	block.AppendInst(i)
	return i
}

// NewLoad appends a new load instruction to the basic block based on the given
// source address.
func (block *BasicBlock) NewLoad(src value.Value) *InstLoad {
	i := NewLoad(src)
	block.AppendInst(i)
	return i
}

// NewStore appends a new store instruction to the basic block based on the
// given source value and destination address.
func (block *BasicBlock) NewStore(src, dst value.Value) *InstStore {
	i := NewStore(src, dst)
	block.AppendInst(i)
	return i
}

// NewGetElementPtr appends a new getelementptr instruction to the basic block
// based on the given source address and element indices.
func (block *BasicBlock) NewGetElementPtr(src value.Value, indices ...value.Value) *InstGetElementPtr {
	i := NewGetElementPtr(src, indices...)
	block.AppendInst(i)
	return i
}

// --- [ Conversion instructions ] ---------------------------------------------

// NewTrunc appends a new trunc instruction to the basic block based on the
// given source value and target type.
func (block *BasicBlock) NewTrunc(from value.Value, to types.Type) *InstTrunc {
	i := NewTrunc(from, to)
	block.AppendInst(i)
	return i
}

// NewZExt appends a new zext instruction to the basic block based on the given
// source value and target type.
func (block *BasicBlock) NewZExt(from value.Value, to types.Type) *InstZExt {
	i := NewZExt(from, to)
	block.AppendInst(i)
	return i
}

// NewSExt appends a new sext instruction to the basic block based on the given
// source value and target type.
func (block *BasicBlock) NewSExt(from value.Value, to types.Type) *InstSExt {
	i := NewSExt(from, to)
	block.AppendInst(i)
	return i
}

// NewFPTrunc appends a new fptrunc instruction to the basic block based on the
// given source value and target type.
func (block *BasicBlock) NewFPTrunc(from value.Value, to types.Type) *InstFPTrunc {
	i := NewFPTrunc(from, to)
	block.AppendInst(i)
	return i
}

// NewFPExt appends a new fpext instruction to the basic block based on the
// given source value and target type.
func (block *BasicBlock) NewFPExt(from value.Value, to types.Type) *InstFPExt {
	i := NewFPExt(from, to)
	block.AppendInst(i)
	return i
}

// NewFPToUI appends a new fptoui instruction to the basic block based on the
// given source value and target type.
func (block *BasicBlock) NewFPToUI(from value.Value, to types.Type) *InstFPToUI {
	i := NewFPToUI(from, to)
	block.AppendInst(i)
	return i
}

// NewFPToSI appends a new fptosi instruction to the basic block based on the
// given source value and target type.
func (block *BasicBlock) NewFPToSI(from value.Value, to types.Type) *InstFPToSI {
	i := NewFPToSI(from, to)
	block.AppendInst(i)
	return i
}

// NewUIToFP appends a new uitofp instruction to the basic block based on the
// given source value and target type.
func (block *BasicBlock) NewUIToFP(from value.Value, to types.Type) *InstUIToFP {
	i := NewUIToFP(from, to)
	block.AppendInst(i)
	return i
}

// NewSIToFP appends a new sitofp instruction to the basic block based on the
// given source value and target type.
func (block *BasicBlock) NewSIToFP(from value.Value, to types.Type) *InstSIToFP {
	i := NewSIToFP(from, to)
	block.AppendInst(i)
	return i
}

// NewPtrToInt appends a new ptrtoint instruction to the basic block based on
// the given source value and target type.
func (block *BasicBlock) NewPtrToInt(from value.Value, to types.Type) *InstPtrToInt {
	i := NewPtrToInt(from, to)
	block.AppendInst(i)
	return i
}

// NewIntToPtr appends a new inttoptr instruction to the basic block based on
// the given source value and target type.
func (block *BasicBlock) NewIntToPtr(from value.Value, to types.Type) *InstIntToPtr {
	i := NewIntToPtr(from, to)
	block.AppendInst(i)
	return i
}

// NewBitCast appends a new bitcast instruction to the basic block based on the
// given source value and target type.
func (block *BasicBlock) NewBitCast(from value.Value, to types.Type) *InstBitCast {
	i := NewBitCast(from, to)
	block.AppendInst(i)
	return i
}

// NewAddrSpaceCast appends a new addrspacecast instruction to the basic block
// based on the given source value and target type.
func (block *BasicBlock) NewAddrSpaceCast(from value.Value, to types.Type) *InstAddrSpaceCast {
	i := NewAddrSpaceCast(from, to)
	block.AppendInst(i)
	return i
}

// --- [ Other instructions ] --------------------------------------------------

// NewICmp appends a new icmp instruction to the basic block based on the given
// integer condition code and operands.
func (block *BasicBlock) NewICmp(cond IntPred, x, y value.Value) *InstICmp {
	i := NewICmp(cond, x, y)
	block.AppendInst(i)
	return i
}

// NewFCmp appends a new fcmp instruction to the basic block based on the given
// floating-point condition code and operands.
func (block *BasicBlock) NewFCmp(cond FloatPred, x, y value.Value) *InstFCmp {
	i := NewFCmp(cond, x, y)
	block.AppendInst(i)
	return i
}

// NewPHI appends a new phi instruction to the basic block based on the given
// incoming values.
func (block *BasicBlock) NewPHI(incs ...*Incoming) *InstPHI {
	i := NewPHI(incs...)
	block.AppendInst(i)
	return i
}

// NewSelect appends a new select instruction to the basic block based on the
// given selection condition and operands.
func (block *BasicBlock) NewSelect(cond, x, y value.Value) *InstSelect {
	i := NewSelect(cond, x, y)
	block.AppendInst(i)
	return i
}

// NewCall appends a new call instruction to the basic block based on the given
// callee and function arguments.
func (block *BasicBlock) NewCall(callee *Function, args ...value.Value) *InstCall {
	i := NewCall(callee, args...)
	block.AppendInst(i)
	return i
}

// --- [ Terminators ] ---------------------------------------------------------

// NewRet sets the terminator of the basic block to a new ret terminator based
// on the given return value. A nil return value indicates a "void" return.
func (block *BasicBlock) NewRet(x value.Value) *TermRet {
	t := NewRet(x)
	block.SetTerm(t)
	return t
}

// NewBr sets the terminator of the basic block to a new unconditional br
// terminator based on the given target branch.
func (block *BasicBlock) NewBr(target *BasicBlock) *TermBr {
	t := NewBr(target)
	block.SetTerm(t)
	return t
}

// NewCondBr sets the terminator of the basic block to a new conditional br
// terminator based on the given branching condition and conditional target
// branches.
func (block *BasicBlock) NewCondBr(cond value.Value, targetTrue, targetFalse *BasicBlock) *TermCondBr {
	t := NewCondBr(cond, targetTrue, targetFalse)
	block.SetTerm(t)
	return t
}

// NewSwitch sets the terminator of the basic block to a new switch terminator
// based on the given control variable, default target branch and switch cases.
func (block *BasicBlock) NewSwitch(x value.Value, targetDefault *BasicBlock, cases ...*Case) *TermSwitch {
	t := NewSwitch(x, targetDefault, cases...)
	block.SetTerm(t)
	return t
}

// NewUnreachable sets the terminator of the basic block to a new unreachable
// terminator.
func (block *BasicBlock) NewUnreachable() *TermUnreachable {
	t := NewUnreachable()
	block.SetTerm(t)
	return t
}
