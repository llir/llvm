package ir

import (
	"github.com/llir/l/ir/types"
)

// --- [ blockaddress constants ] ----------------------------------------------

// ConstBlockAddress is an LLVM IR blockaddress constant.
type ConstBlockAddress struct {
	// Parent function.
	Func *Function
	// Basic block to take address of.
	Block *BasicBlock
}

// NewBlockAddress returns a new blockaddress constant based on the given parent
// function and basic block.
func NewBlockAddress(f *Function, block *BasicBlock) *ConstBlockAddress {
	return &ConstBlockAddress{Func: f, Block: block}
}

// Type returns the type of the constant.
func (c *ConstBlockAddress) Type() types.Type {
	panic("not yet implemented")
}

// Ident returns the identifier associated with the constant.
func (c *ConstBlockAddress) Ident() string {
	panic("not yet implemented")
}
