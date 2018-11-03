package ir

import (
	"github.com/llir/llvm/ir/value"
)

// --- [ Binary instructions ] -------------------------------------------------

// ~~~ [ add ] ~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~

// NewAdd appends a new add instruction to the basic block based on the given
// operands.
func (block *BasicBlock) NewAdd(x, y value.Value) *InstAdd {
	inst := NewAdd(x, y)
	block.Insts = append(block.Insts, inst)
	return inst
}

// ~~~ [ fadd ] ~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~

// NewFAdd appends a new fadd instruction to the basic block based on the given
// operands.
func (block *BasicBlock) NewFAdd(x, y value.Value) *InstFAdd {
	inst := NewFAdd(x, y)
	block.Insts = append(block.Insts, inst)
	return inst
}

// ~~~ [ sub ] ~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~

// NewSub appends a new sub instruction to the basic block based on the given
// operands.
func (block *BasicBlock) NewSub(x, y value.Value) *InstSub {
	inst := NewSub(x, y)
	block.Insts = append(block.Insts, inst)
	return inst
}

// ~~~ [ fsub ] ~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~

// NewFSub appends a new fsub instruction to the basic block based on the given
// operands.
func (block *BasicBlock) NewFSub(x, y value.Value) *InstFSub {
	inst := NewFSub(x, y)
	block.Insts = append(block.Insts, inst)
	return inst
}

// ~~~ [ mul ] ~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~

// NewMul appends a new mul instruction to the basic block based on the given
// operands.
func (block *BasicBlock) NewMul(x, y value.Value) *InstMul {
	inst := NewMul(x, y)
	block.Insts = append(block.Insts, inst)
	return inst
}

// ~~~ [ fmul ] ~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~

// NewFMul appends a new fmul instruction to the basic block based on the given
// operands.
func (block *BasicBlock) NewFMul(x, y value.Value) *InstFMul {
	inst := NewFMul(x, y)
	block.Insts = append(block.Insts, inst)
	return inst
}

// ~~~ [ udiv ] ~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~

// NewUDiv appends a new udiv instruction to the basic block based on the given
// operands.
func (block *BasicBlock) NewUDiv(x, y value.Value) *InstUDiv {
	inst := NewUDiv(x, y)
	block.Insts = append(block.Insts, inst)
	return inst
}

// ~~~ [ sdiv ] ~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~

// NewSDiv appends a new sdiv instruction to the basic block based on the given
// operands.
func (block *BasicBlock) NewSDiv(x, y value.Value) *InstSDiv {
	inst := NewSDiv(x, y)
	block.Insts = append(block.Insts, inst)
	return inst
}

// ~~~ [ fdiv ] ~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~

// NewFDiv appends a new fdiv instruction to the basic block based on the given
// operands.
func (block *BasicBlock) NewFDiv(x, y value.Value) *InstFDiv {
	inst := NewFDiv(x, y)
	block.Insts = append(block.Insts, inst)
	return inst
}

// ~~~ [ urem ] ~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~

// NewURem appends a new urem instruction to the basic block based on the given
// operands.
func (block *BasicBlock) NewURem(x, y value.Value) *InstURem {
	inst := NewURem(x, y)
	block.Insts = append(block.Insts, inst)
	return inst
}

// ~~~ [ srem ] ~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~

// NewSRem appends a new srem instruction to the basic block based on the given
// operands.
func (block *BasicBlock) NewSRem(x, y value.Value) *InstSRem {
	inst := NewSRem(x, y)
	block.Insts = append(block.Insts, inst)
	return inst
}

// ~~~ [ frem ] ~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~

// NewFRem appends a new frem instruction to the basic block based on the given
// operands.
func (block *BasicBlock) NewFRem(x, y value.Value) *InstFRem {
	inst := NewFRem(x, y)
	block.Insts = append(block.Insts, inst)
	return inst
}
