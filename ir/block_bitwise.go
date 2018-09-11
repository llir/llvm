package ir

import (
	"github.com/llir/l/ir/value"
)

// --- [ Bitwise instructions ] ------------------------------------------------

// ~~~ [ shl ] ~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~

// NewShl returns a new shl instruction based on the given operands.
func (block *BasicBlock) NewShl(x, y value.Value) *Shl {
	inst := NewShl(x, y)
	block.Insts = append(block.Insts, inst)
	return inst
}

// ~~~ [ lshr ] ~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~

// NewLShr returns a new lshr instruction based on the given operands.
func (block *BasicBlock) NewLShr(x, y value.Value) *LShr {
	inst := NewLShr(x, y)
	block.Insts = append(block.Insts, inst)
	return inst
}

// ~~~ [ ashr ] ~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~

// NewAShr returns a new ashr instruction based on the given operands.
func (block *BasicBlock) NewAShr(x, y value.Value) *AShr {
	inst := NewAShr(x, y)
	block.Insts = append(block.Insts, inst)
	return inst
}

// ~~~ [ and ] ~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~

// NewAnd returns a new and instruction based on the given operands.
func (block *BasicBlock) NewAnd(x, y value.Value) *And {
	inst := NewAnd(x, y)
	block.Insts = append(block.Insts, inst)
	return inst
}

// ~~~ [ or ] ~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~

// NewOr returns a new or instruction based on the given operands.
func (block *BasicBlock) NewOr(x, y value.Value) *Or {
	inst := NewOr(x, y)
	block.Insts = append(block.Insts, inst)
	return inst
}

// ~~~ [ xor ] ~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~

// NewXor returns a new xor instruction based on the given operands.
func (block *BasicBlock) NewXor(x, y value.Value) *Xor {
	inst := NewXor(x, y)
	block.Insts = append(block.Insts, inst)
	return inst
}
