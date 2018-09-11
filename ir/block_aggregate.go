package ir

import (
	"github.com/llir/l/ir/instruction"
	"github.com/llir/l/ir/value"
)

// --- [ Aggregate instructions ] ----------------------------------------------

// ~~~ [ extractvalue ] ~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~

// NewExtractValue returns a new extractvalue instruction based on the given
// aggregate value and indicies.
func (block *BasicBlock) NewExtractValue(x value.Value, indices ...int64) *instruction.ExtractValue {
	inst := instruction.NewExtractValue(x, indices...)
	block.Insts = append(block.Insts)
	return inst
}

// ~~~ [ insertvalue ] ~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~

// NewInsertValue returns a new insertvalue instruction based on the given
// aggregate value, element and indicies.
func (block *BasicBlock) NewInsertValue(x, elem value.Value, indices ...int64) *instruction.InsertValue {
	inst := instruction.NewInsertValue(x, elem, indices...)
	block.Insts = append(block.Insts)
	return inst
}
