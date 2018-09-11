package ir

import (
	"github.com/llir/l/ir/instruction"
	"github.com/llir/l/ir/ll"
	"github.com/llir/l/ir/types"
	"github.com/llir/l/ir/value"
)

// --- [ Memory instructions ] -------------------------------------------------

// ~~~ [ alloca ] ~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~

// NewAlloca returns a new alloca instruction based on the given element type.
func (block *BasicBlock) NewAlloca(elemType types.Type) *instruction.Alloca {
	inst := instruction.NewAlloca(elemType)
	block.Insts = append(block.Insts, inst)
	return inst
}

// ~~~ [ load ] ~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~

// NewLoad returns a new load instruction based on the given source address.
func (block *BasicBlock) NewLoad(src value.Value) *instruction.Load {
	inst := instruction.NewLoad(src)
	block.Insts = append(block.Insts, inst)
	return inst
}

// ~~~ [ store ] ~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~

// NewStore returns a new store instruction based on the given source value and
// destination address.
func (block *BasicBlock) NewStore(src, dst value.Value) *instruction.Store {
	inst := instruction.NewStore(src, dst)
	block.Insts = append(block.Insts, inst)
	return inst
}

// ~~~ [ fence ] ~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~

// NewFence returns a new fence instruction based on the given atomic ordering.
func (block *BasicBlock) NewFence(ordering ll.AtomicOrdering) *instruction.Fence {
	inst := instruction.NewFence(ordering)
	block.Insts = append(block.Insts, inst)
	return inst
}

// ~~~ [ cmpxchg ] ~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~

// NewCmpXchg returns a new cmpxchg instruction based on the given address,
// value to compare against, new value to store, and atomic orderings for
// success and failure.
func (block *BasicBlock) NewCmpXchg(ptr, cmp, new value.Value, success, failure ll.AtomicOrdering) *instruction.CmpXchg {
	inst := instruction.NewCmpXchg(ptr, cmp, new, success, failure)
	block.Insts = append(block.Insts, inst)
	return inst
}

// ~~~ [ atomicrmw ] ~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~

// NewAtomicRMW returns a new atomicrmw instruction based on the given atomic
// operation, destination address, operand and atomic ordering.
func (block *BasicBlock) NewAtomicRMW(op ll.AtomicOp, dst, x value.Value, ordering ll.AtomicOrdering) *instruction.AtomicRMW {
	inst := instruction.NewAtomicRMW(op, dst, x, ordering)
	block.Insts = append(block.Insts, inst)
	return inst
}

// ~~~ [ getelementptr ] ~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~

// NewGetElementPtr returns a new getelementptr instruction based on the given
// source address and element indices.
func (block *BasicBlock) NewGetElementPtr(src value.Value, indices ...value.Value) *instruction.GetElementPtr {
	inst := instruction.NewGetElementPtr(src, indices...)
	block.Insts = append(block.Insts, inst)
	return inst
}
