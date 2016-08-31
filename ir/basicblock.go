package ir

import (
	"bytes"
	"fmt"
	"strings"

	"github.com/llir/llvm/ir/instruction"
	"github.com/llir/llvm/ir/internal/enc"
	"github.com/llir/llvm/ir/types"
	"github.com/mewkiz/pkg/errutil"
)

// A BasicBlock is a sequence of non-branching instructions, terminated by a
// control flow instruction (such as br or ret).
//
// Basic blocks are values since they can be referenced from instructions (such
// as br and switch). The type of a basic block is label.
//
// References:
//    http://llvm.org/docs/LangRef.html#terminators
type BasicBlock struct {
	// Basic block label name.
	name string
	// Parent function of the basic block.
	parent *Function
	// Non-terminator instructions of the basic block.
	insts []instruction.Instruction
	// Terminator instruction of the basic block.
	term instruction.Terminator
}

// NewBasicBlock returns a new basic block based on the given name.
func NewBasicBlock(name string) *BasicBlock {
	// TODO: Verify that name is not a local ID. Unnamed basic blocks should be
	// assigned a local ID implicitly by the internal localID counter of the
	// given function rather than explicitly assigned.
	//
	// Another option (which is currently implemented) is to allow local IDs to
	// be explicitly assigned, as they are validated during the ID assignment
	// stage (see BasicBlock.assignIDs).
	return &BasicBlock{name: name}
}

// Type returns the type of the value.
func (block *BasicBlock) Type() types.Type {
	return types.NewLabel()
}

// Name returns the name of the basic block.
func (block *BasicBlock) Name() string {
	return block.name
}

// TODO: Add note to SetName not set local IDs explicitly, as these are assigned
// implicitly by the internal localID counter.

// SetName sets the name of the basic block.
func (block *BasicBlock) SetName(name string) {
	block.name = name
}

// Insts returns the non-terminating instructions of the basic block.
func (block *BasicBlock) Insts() []instruction.Instruction {
	return block.insts
}

// SetInsts sets the instructions of the basic block.
func (block *BasicBlock) SetInsts(insts []instruction.Instruction) {
	block.insts = insts
}

// AppendInst appends the given instruction to the basic block.
func (block *BasicBlock) AppendInst(inst instruction.Instruction) {
	block.insts = append(block.insts, inst)
}

// Term returns the terminator of the basic block.
func (block *BasicBlock) Term() instruction.Terminator {
	return block.term
}

// SetTerm sets the terminator of the basic block.
func (block *BasicBlock) SetTerm(term instruction.Terminator) {
	block.term = term
}

// Parent returns the parent function of the basic block.
func (block *BasicBlock) Parent() *Function {
	return block.parent
}

// SetParent sets the parent function of the basic block.
func (block *BasicBlock) SetParent(parent *Function) {
	block.parent = parent
}

// String returns the string representation of the basic block.
func (block *BasicBlock) String() string {
	buf := new(bytes.Buffer)
	if len(block.Name()) > 0 {
		if isLocalID(block.Name()) {
			fmt.Fprintf(buf, "; <label>:%s\n", block.Name())
		} else {
			fmt.Fprintf(buf, "%s:\n", block.Name())
		}
	}
	for _, inst := range block.Insts() {
		fmt.Fprintf(buf, "\t%s\n", inst)
	}
	fmt.Fprintf(buf, "\t%s", block.Term())
	return buf.String()
}

// ValueString returns a string representation of the value.
func (block *BasicBlock) ValueString() string {
	return enc.Local(block.Name())
}

// assignIDs assigns unique IDs to unnamed basic blocks and local variable
// definitions.
func (block *BasicBlock) assignIDs() error {
	f := block.Parent()

	// Named represents a named basic block or local variable definition.
	type Named interface {
		Name() string
		SetName(name string)
	}

	// setName assigns unique local IDs to unnamed basic blocks and local
	// variable definitions.
	setName := func(n Named) error {
		// TODO: Ensure that global variables cannot be mixed up with local
		// variables. This should be easy, as global variables may not be unnamed.
		// Check that global variables are always given a name during creation.
		if name := n.Name(); len(name) == 0 {
			n.SetName(f.nextID())
		} else if isLocalID(name) {
			// Validate that explicitly named local IDs conform to the localID
			// counter and update the localID counter to keep explicitly and
			// implicitly named local IDs in sync.
			if want := f.nextID(); name != want {
				return errutil.Newf("invalid local ID; expected %s, got %s", enc.Local(want), enc.Local(name))
			}
		}
		return nil
	}

	// Assign unique local IDs to unnamed basic blocks.
	if err := setName(block); err != nil {
		return errutil.Err(err)
	}

	// Assign unique local IDs to unnamed local variable definitions.
	for _, inst := range block.Insts() {
		if def, ok := inst.(*instruction.LocalVarDef); ok {
			if !types.IsVoid(def.Inst().RetType()) {
				if err := setName(def); err != nil {
					return errutil.Err(err)
				}
			}
		}
	}

	return nil
}

// isLocalID reports whether the given name is a local ID (e.g. "%42").
func isLocalID(name string) bool {
	if len(name) == 0 {
		panic("invalid name length; expected > 0")
	}
	for _, r := range name {
		if strings.IndexRune("0123456789", r) == -1 {
			return false
		}
	}
	return true
}
