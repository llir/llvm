package dummy

import (
	"bytes"
	"fmt"

	"github.com/llir/llvm/internal/enc"
	"github.com/llir/llvm/ir"
	"github.com/llir/llvm/ir/types"
	"github.com/llir/llvm/ir/value"
)

// --- [ getelementptr ] -------------------------------------------------------

// InstGetElementPtr represents a dummy getelementptr instruction.
type InstGetElementPtr struct {
	// Parent basic block.
	parent *ir.BasicBlock
	// Name of the local variable associated with the instruction.
	name string
	// Source address element type.
	elem types.Type
	// Source address.
	src value.Value
	// Element indices.
	indices []value.Value
}

// NewGetElementPtr returns a new dummy getelementptr instruction based on the
// given source address element type, source address and element indices.
func NewGetElementPtr(elem types.Type, src value.Value, indices ...value.Value) *InstGetElementPtr {
	return &InstGetElementPtr{elem: elem, src: src, indices: indices}
}

// Type returns the type of the instruction.
func (inst *InstGetElementPtr) Type() types.Type {
	panic("dummy implementation")
}

// Ident returns the identifier associated with the instruction.
func (inst *InstGetElementPtr) Ident() string {
	return enc.Local(inst.name)
}

// Name returns the name of the local variable associated with the instruction.
func (inst *InstGetElementPtr) Name() string {
	return inst.name
}

// SetName sets the name of the local variable associated with the instruction.
func (inst *InstGetElementPtr) SetName(name string) {
	inst.name = name
}

// String returns the LLVM syntax representation of the instruction.
func (inst *InstGetElementPtr) String() string {
	buf := &bytes.Buffer{}
	src := inst.Src()
	fmt.Fprintf(buf, "%s = getelementptr %s, %s %s",
		inst.Ident(),
		inst.elem,
		src.Type(),
		src.Ident())
	for _, index := range inst.Indices() {
		fmt.Fprintf(buf, ", %s %s",
			index.Type(),
			index.Ident())
	}
	return buf.String()
}

// Parent returns the parent basic block of the instruction.
func (inst *InstGetElementPtr) Parent() *ir.BasicBlock {
	return inst.parent
}

// SetParent sets the parent basic block of the instruction.
func (inst *InstGetElementPtr) SetParent(parent *ir.BasicBlock) {
	inst.parent = parent
}

// ElemType returns the source address element type of the getelementptr
// instruction.
func (inst *InstGetElementPtr) ElemType() types.Type {
	return inst.elem
}

// Src returns the source address of the getelementptr instruction.
func (inst *InstGetElementPtr) Src() value.Value {
	return inst.src
}

// Indices returns the element indices of the getelementptr instruction.
func (inst *InstGetElementPtr) Indices() []value.Value {
	return inst.indices
}
