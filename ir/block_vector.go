package ir

import (
	"github.com/llir/l/ir/instruction"
	"github.com/llir/l/ir/value"
)

// --- [ Vector instructions ] -------------------------------------------------

// ~~~ [ extractelement ] ~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~

// NewExtractElement returns a new extractelement instruction based on the given
// vector and element index.
func (block *BasicBlock) NewExtractElement(x, index value.Value) *instruction.ExtractElement {
	inst := instruction.NewExtractElement(x, index)
	block.Insts = append(block.Insts)
	return inst
}

// ~~~ [ insertelement ] ~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~

// NewInsertElement returns a new insertelement instruction based on the given
// vector, element and element index.
func (block *BasicBlock) NewInsertElement(x, elem, index value.Value) *instruction.InsertElement {
	inst := instruction.NewInsertElement(x, elem, index)
	block.Insts = append(block.Insts)
	return inst
}

// ~~~ [ shufflevector ] ~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~

// NewShuffleVector returns a new shufflevector instruction based on the given
// vectors and shuffle mask.
func (block *BasicBlock) NewShuffleVector(x, y, mask value.Value) *instruction.ShuffleVector {
	inst := instruction.NewShuffleVector(x, y, mask)
	block.Insts = append(block.Insts)
	return inst
}
