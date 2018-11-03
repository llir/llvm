package ir

import (
	"fmt"

	"github.com/llir/llvm/ir/types"
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

// String returns the LLVM syntax representation of the constant as a type-value
// pair.
func (c *ConstBlockAddress) String() string {
	return fmt.Sprintf("%v %v", c.Type(), c.Ident())
}

// Type returns the type of the constant.
func (c *ConstBlockAddress) Type() types.Type {
	return types.I8Ptr
}

// Ident returns the identifier associated with the constant.
func (c *ConstBlockAddress) Ident() string {
	// "blockaddress" "(" GlobalIdent "," LocalIdent ")"
	return fmt.Sprintf("blockaddress(%v, %v)", c.Func.Ident(), c.Block.Ident())
}
