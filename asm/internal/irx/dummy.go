// Dummy values are used to allow for forward references, and are replaced by
// their real values in later stages of parsing.

package irx

import (
	"bytes"
	"fmt"

	"github.com/llir/llvm/internal/enc"
	"github.com/llir/llvm/ir"
	"github.com/llir/llvm/ir/constant"
	"github.com/llir/llvm/ir/types"
	"github.com/llir/llvm/ir/value"
)

// namedTypeDummy represents a dummy named type.
type namedTypeDummy struct {
	// Type name.
	name string
}

// String returns the LLVM syntax representation of the type.
func (t *namedTypeDummy) String() string {
	return enc.Local(t.name)
}

// Equal reports whether t and u are of equal type.
func (t *namedTypeDummy) Equal(u types.Type) bool {
	panic("dummy implementation")
}

// newNamedTypeDummy returns a new dummy named type based on the given type
// name.
func newNamedTypeDummy(name string) *namedTypeDummy {
	return &namedTypeDummy{name: name}
}

// instGetElementPtrDummy represents a dummy getelementptr instruction.
type instGetElementPtrDummy struct {
	// Parent basic block.
	parent *ir.BasicBlock
	// Name of the local variable associated with the instruction.
	name string
	// Source address.
	src value.Value
	// Element indices.
	indices []value.Value
	// Source address element type.
	elem types.Type
}

// newGetElementPtrDummy returns a new dummy getelementptr instruction based on
// the given element type, source address and element indices.
func newGetElementPtrDummy(elem types.Type, src value.Value, indices ...value.Value) *instGetElementPtrDummy {
	return &instGetElementPtrDummy{src: src, indices: indices, elem: elem}
}

// Type returns the type of the instruction.
func (inst *instGetElementPtrDummy) Type() types.Type {
	panic("dummy implementation")
}

// Ident returns the identifier associated with the instruction.
func (inst *instGetElementPtrDummy) Ident() string {
	return enc.Local(inst.name)
}

// Name returns the name of the local variable associated with the instruction.
func (inst *instGetElementPtrDummy) Name() string {
	return inst.name
}

// SetName sets the name of the local variable associated with the instruction.
func (inst *instGetElementPtrDummy) SetName(name string) {
	inst.name = name
}

// String returns the LLVM syntax representation of the instruction.
func (inst *instGetElementPtrDummy) String() string {
	buf := &bytes.Buffer{}
	src := inst.src
	fmt.Fprintf(buf, "%s = getelementptr %s, %s %s",
		inst.Ident(),
		inst.elem,
		src.Type(),
		src.Ident())
	for _, index := range inst.indices {
		fmt.Fprintf(buf, ", %s %s",
			index.Type(),
			index.Ident())
	}
	return buf.String()
}

// Parent returns the parent basic block of the instruction.
func (inst *instGetElementPtrDummy) Parent() *ir.BasicBlock {
	return inst.parent
}

// SetParent sets the parent basic block of the instruction.
func (inst *instGetElementPtrDummy) SetParent(parent *ir.BasicBlock) {
	inst.parent = parent
}

// termSwitchDummy represents a dummy switch terminator.
type termSwitchDummy struct {
	// Parent basic block.
	parent *ir.BasicBlock
	// Control variable.
	x value.Value
	// Default target branch.
	targetDefault string
	// Switch cases.
	cases []*caseDummy
}

// newSwitchDummy returns a new dummy switch terminator based on the given
// control variable, default target branch and switch cases.
func newSwitchDummy(x value.Value, targetDefault string, cases ...*caseDummy) *termSwitchDummy {
	return &termSwitchDummy{x: x, targetDefault: targetDefault, cases: cases}
}

// String returns the LLVM syntax representation of the terminator.
func (term *termSwitchDummy) String() string {
	buf := &bytes.Buffer{}
	x := term.x
	fmt.Fprintf(buf, "switch %s %s, label %s [\n",
		x.Type(),
		x.Ident(),
		term.targetDefault)
	for _, c := range term.cases {
		x := c.x
		fmt.Fprintf(buf, "\t\t%s %s, label %s\n",
			x.Type(),
			x.Ident(),
			c.target)
	}
	buf.WriteString("\t]")
	return buf.String()
}

// Parent returns the parent basic block of the terminator.
func (term *termSwitchDummy) Parent() *ir.BasicBlock {
	return term.parent
}

// SetParent sets the parent basic block of the terminator.
func (term *termSwitchDummy) SetParent(parent *ir.BasicBlock) {
	term.parent = parent
}

// Successors returns the successor basic blocks of the terminator.
func (term *termSwitchDummy) Successors() []*ir.BasicBlock {
	panic("dummy implementation")
}

// caseDummy represents a dummy case value of a switch terminator.
type caseDummy struct {
	// Case comparand.
	x *constant.Int
	// Case target branch.
	target string
}

// newCaseDummy returns a new dummy switch case based on the given case
// comparand and target branch.
func newCaseDummy(x *constant.Int, target string) *caseDummy {
	return &caseDummy{x: x, target: target}
}
