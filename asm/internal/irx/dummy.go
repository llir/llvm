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
