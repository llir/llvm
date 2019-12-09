// The basic block API of the ir package was heavily inspired by
// https://github.com/mrbenshef/goory.

package ir

import (
	"fmt"
	"strings"

	"github.com/llir/llvm/internal/enc"
	"github.com/llir/llvm/ir/types"
)

// === [ Basic blocks ] ========================================================

// Block is an LLVM IR basic block; a sequence of non-branching instructions
// terminated by a control flow instruction (e.g. br or ret).
type Block struct {
	// Name of local variable associated with the basic block.
	LocalIdent
	// Instructions of the basic block.
	Insts []Instruction
	// Terminator of the basic block.
	Term Terminator

	// extra.

	// Parent function; field set by ir.Func.NewBlock.
	Parent *Func `json:"-"`
}

// NewBlock returns a new basic block based on the given label name. An empty
// label name indicates an unnamed basic block.
func NewBlock(name string) *Block {
	block := &Block{}
	block.SetName(name)
	return block
}

// String returns the LLVM syntax representation of the basic block as a
// type-value pair.
func (block *Block) String() string {
	return fmt.Sprintf("%s %s", block.Type(), block.Ident())
}

// Type returns the type of the basic block.
func (block *Block) Type() types.Type {
	return types.Label
}

// LLString returns the LLVM syntax representation of the basic block
// definition.
func (block *Block) LLString() string {
	// Name=LabelIdentopt Insts=Instruction* Term=Terminator
	buf := &strings.Builder{}
	if block.IsUnnamed() {
		fmt.Fprintf(buf, "; <label>:%d\n", block.LocalID)
		// TODO: explicitly print basic block label to conform with Clang 9.0, and
		// because it's the sane thing to do.
		//fmt.Fprintf(buf, "%s\n", enc.LabelID(block.LocalID))
	} else {
		fmt.Fprintf(buf, "%s\n", enc.LabelName(block.LocalName))
	}
	for _, inst := range block.Insts {
		fmt.Fprintf(buf, "\t%s\n", inst.LLString())
	}
	if block.Term == nil {
		panic(fmt.Sprintf("missing terminator in basic block %q.\ncurrent instructions:\n%s", block.Name(), buf.String()))
	}
	fmt.Fprintf(buf, "\t%s", block.Term.LLString())
	return buf.String()
}
