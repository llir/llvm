package constant

import (
	"fmt"

	"github.com/llir/llvm/ir/types"
	"github.com/llir/llvm/ir/value"
)

// --- [ blockaddress constants ] ----------------------------------------------

// BlockAddress is an LLVM IR blockaddress constant.
type BlockAddress struct {
	// Parent function.
	Func Constant // *ir.Function
	// Basic block to take address of.
	Block value.Named // *ir.BasicBlock
}

// NewBlockAddress returns a new blockaddress constant based on the given parent
// function and basic block.
func NewBlockAddress(f Constant, block value.Named) *BlockAddress {
	return &BlockAddress{Func: f, Block: block}
}

// String returns the LLVM syntax representation of the constant as a type-value
// pair.
func (c *BlockAddress) String() string {
	return fmt.Sprintf("%v %v", c.Type(), c.Ident())
}

// Type returns the type of the constant.
func (c *BlockAddress) Type() types.Type {
	return types.I8Ptr
}

// Ident returns the identifier associated with the constant.
func (c *BlockAddress) Ident() string {
	// "blockaddress" "(" GlobalIdent "," LocalIdent ")"
	return fmt.Sprintf("blockaddress(%v, %v)", c.Func.Ident(), c.Block.Ident())
}
