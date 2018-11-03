package ir

import (
	"github.com/llir/llvm/ir/value"
)

// --- [ Aggregate instructions ] ----------------------------------------------

// ~~~ [ extractvalue ] ~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~

// NewExtractValue appends a new extractvalue instruction to the basic block
// based on the given aggregate value and indicies.
func (block *BasicBlock) NewExtractValue(x value.Value, indices ...int64) *InstExtractValue {
	inst := NewExtractValue(x, indices...)
	block.Insts = append(block.Insts)
	return inst
}

// ~~~ [ insertvalue ] ~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~

// NewInsertValue appends a new insertvalue instruction to the basic block based
// on the given aggregate value, element and indicies.
func (block *BasicBlock) NewInsertValue(x, elem value.Value, indices ...int64) *InstInsertValue {
	inst := NewInsertValue(x, elem, indices...)
	block.Insts = append(block.Insts)
	return inst
}
