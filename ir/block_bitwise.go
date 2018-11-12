package ir

import (
	"github.com/llir/llvm/ir/value"
)

// --- [ Bitwise instructions ] ------------------------------------------------

// ~~~ [ shl ] ~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~

// NewShl appends a new shl instruction to the basic block based on the given
// operands.
func (block *BasicBlock) NewShl(x, y value.Value) *InstShl {
	inst := NewShl(x, y)
	block.Insts = append(block.Insts, inst)
	return inst
}

// ~~~ [ lshr ] ~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~

// NewLShr appends a new lshr instruction to the basic block based on the given
// operands.
func (block *BasicBlock) NewLShr(x, y value.Value) *InstLShr {
	inst := NewLShr(x, y)
	block.Insts = append(block.Insts, inst)
	return inst
}

// ~~~ [ ashr ] ~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~

// NewAShr appends a new ashr instruction to the basic block based on the given
// operands.
func (block *BasicBlock) NewAShr(x, y value.Value) *InstAShr {
	inst := NewAShr(x, y)
	block.Insts = append(block.Insts, inst)
	return inst
}

// ~~~ [ and ] ~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~

// NewAnd appends a new and instruction to the basic block based on the given
// operands.
func (block *BasicBlock) NewAnd(x, y value.Value) *InstAnd {
	inst := NewAnd(x, y)
	block.Insts = append(block.Insts, inst)
	return inst
}

// ~~~ [ or ] ~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~

// NewOr appends a new or instruction to the basic block based on the given
// operands.
func (block *BasicBlock) NewOr(x, y value.Value) *InstOr {
	inst := NewOr(x, y)
	block.Insts = append(block.Insts, inst)
	return inst
}

// ~~~ [ xor ] ~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~

// NewXor appends a new xor instruction to the basic block based on the given
// operands.
func (block *BasicBlock) NewXor(x, y value.Value) *InstXor {
	inst := NewXor(x, y)
	block.Insts = append(block.Insts, inst)
	return inst
}
