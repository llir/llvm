package ir

import (
	"fmt"
	"strings"

	"github.com/llir/llvm/internal/enc"
	"github.com/llir/llvm/ir/types"
)

// === [ Basic blocks ] ========================================================

// BasicBlock is an LLVM IR basic block; a sequence of non-branching
// instructions terminated by a control flow instruction.
type BasicBlock struct {
	// Name of local variable associated with the basic block.
	LocalIdent
	// Instructions of the basic block.
	Insts []Instruction
	// Terminator of the basic block.
	Term Terminator
}

// NewBlock returns a new basic block based on the given label name. An empty
// label name indicates an unnamed basic block.
func NewBlock(name string) *BasicBlock {
	return &BasicBlock{LocalIdent: LocalIdent{LocalName: name}}
}

// String returns the LLVM syntax representation of the basic block as a
// type-value pair.
func (block *BasicBlock) String() string {
	return fmt.Sprintf("%s %s", block.Type(), block.Ident())
}

// Type returns the type of the basic block.
func (block *BasicBlock) Type() types.Type {
	return types.Label
}

// Def returns the LLVM syntax representation of the basic block definition.
func (block *BasicBlock) Def() string {
	// Name=LabelIdentopt Insts=Instruction* Term=Terminator
	buf := &strings.Builder{}
	if block.IsUnnamed() {
		fmt.Fprintf(buf, "; <label>:%d\n", block.LocalID)
	} else {
		fmt.Fprintf(buf, "%s\n", enc.Label(block.LocalName))
	}
	for _, inst := range block.Insts {
		fmt.Fprintf(buf, "\t%s\n", inst.Def())
	}
	fmt.Fprintf(buf, "\t%s", block.Term.Def())
	return buf.String()
}
