// Dummy values are used to allow for forward references, and are replaced by
// their real values in later stages of parsing.

package irx

import (
	"bytes"
	"fmt"

	"github.com/llir/llvm/internal/enc"
	"github.com/llir/llvm/ir"
	"github.com/llir/llvm/ir/types"
	"github.com/llir/llvm/ir/value"
)

// globalDummy represents a dummy value for a given global identifier name and
// type.
type globalDummy struct {
	// Global name.
	name string
	// Type associated with the global identifier.
	typ types.Type
}

// newGlobalDummy returns a new dummy value for the given global identifier name
// and type.
func newGlobalDummy(name string, typ types.Type) *globalDummy {
	return &globalDummy{name: name, typ: typ}
}

// Ident returns the identifier associated with the global.
func (g *globalDummy) Ident() string {
	return enc.Global(g.name)
}

// Type returns the type associated with the global identifier.
func (g *globalDummy) Type() types.Type {
	return g.typ
}

// localDummy represents a dummy value for a given local identifier name and
// type.
type localDummy struct {
	// Local name.
	name string
	// Type associated with the local identifier.
	typ types.Type
}

// newLocalDummy returns a new dummy value for the given local identifier name
// and type.
func newLocalDummy(name string, typ types.Type) *localDummy {
	return &localDummy{name: name, typ: typ}
}

// Ident returns the identifier associated with the local.
func (g *localDummy) Ident() string {
	return enc.Local(g.name)
}

// Type returns the type associated with the local identifier.
func (g *localDummy) Type() types.Type {
	return g.typ
}

// instCallDummy represents a dummy call instruction.
type instCallDummy struct {
	// Parent basic block.
	parent *ir.BasicBlock
	// Identifier associated with the instruction.
	ident string
	// Callee.
	callee string
	// Function arguments.
	args []value.Value
	// Return type.
	ret types.Type
}

// newCallDummy returns a new dummy call instruction based on the given callee
// and function arguments.
func newCallDummy(ret types.Type, callee string, args ...value.Value) *instCallDummy {
	return &instCallDummy{callee: callee, args: args, ret: ret}
}

// Type returns the type of the instruction.
func (inst *instCallDummy) Type() types.Type {
	return inst.ret
}

// Ident returns the identifier associated with the instruction.
func (inst *instCallDummy) Ident() string {
	return enc.Local(inst.ident)
}

// SetIdent sets the identifier associated with the instruction.
func (inst *instCallDummy) SetIdent(ident string) {
	inst.ident = ident
}

// String returns the LLVM syntax representation of the instruction.
func (inst *instCallDummy) String() string {
	buf := &bytes.Buffer{}
	typ := inst.Type()
	if !typ.Equal(types.Void) {
		fmt.Fprintf(buf, "%s = ", inst.Ident())
	}
	fmt.Fprintf(buf, "call %s %s(",
		typ,
		inst.callee)
	for i, arg := range inst.args {
		if i != 0 {
			buf.WriteString(", ")
		}
		fmt.Fprintf(buf, "%s %s",
			arg.Type(),
			arg.Ident())
	}
	buf.WriteString(")")
	return buf.String()
}

// Parent returns the parent basic block of the instruction.
func (inst *instCallDummy) Parent() *ir.BasicBlock {
	return inst.parent
}

// SetParent sets the parent basic block of the instruction.
func (inst *instCallDummy) SetParent(parent *ir.BasicBlock) {
	inst.parent = parent
}

// termBrDummy represents a dummy br terminator.
type termBrDummy struct {
	// Parent basic block.
	parent *ir.BasicBlock
	// Target branch.
	target string
	// Successors basic blocks.
	successors []*ir.BasicBlock
}

// newBrDummy returns a new dummy unconditional br terminator based on the given
// target branch.
func newBrDummy(target string) *termBrDummy {
	return &termBrDummy{target: target}
}

// String returns the LLVM syntax representation of the terminator.
func (term *termBrDummy) String() string {
	return fmt.Sprintf("br label %s", term.target)
}

// Parent returns the parent basic block of the terminator.
func (term *termBrDummy) Parent() *ir.BasicBlock {
	return term.parent
}

// SetParent sets the parent basic block of the terminator.
func (term *termBrDummy) SetParent(parent *ir.BasicBlock) {
	term.parent = parent
}

// Successors returns the successor basic blocks of the terminator.
func (term *termBrDummy) Successors() []*ir.BasicBlock {
	panic("dummy implementation")
}
