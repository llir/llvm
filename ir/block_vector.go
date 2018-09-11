package ir

import (
	"github.com/llir/l/ir/value"
)

// --- [ Vector instructions ] -------------------------------------------------

// ~~~ [ extractelement ] ~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~

// NewExtractElement returns a new extractelement instruction based on the given
// vector and element index.
func (block *BasicBlock) NewExtractElement(x, index value.Value) *ExtractElement {
	inst := NewExtractElement(x, index)
	block.Insts = append(block.Insts)
	return inst
}

// ~~~ [ insertelement ] ~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~

// NewInsertElement returns a new insertelement instruction based on the given
// vector, element and element index.
func (block *BasicBlock) NewInsertElement(x, elem, index value.Value) *InsertElement {
	inst := NewInsertElement(x, elem, index)
	block.Insts = append(block.Insts)
	return inst
}

// ~~~ [ shufflevector ] ~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~

// NewShuffleVector returns a new shufflevector instruction based on the given
// vectors and shuffle mask.
func (block *BasicBlock) NewShuffleVector(x, y, mask value.Value) *ShuffleVector {
	inst := NewShuffleVector(x, y, mask)
	block.Insts = append(block.Insts)
	return inst
}
