package dummy

import (
	"bytes"
	"fmt"

	"github.com/llir/llvm/internal/enc"
	"github.com/llir/llvm/ir"
	"github.com/llir/llvm/ir/types"
	"github.com/llir/llvm/ir/value"
)

// --- [ load ] ----------------------------------------------------------------

// InstLoad represents a dummy load instruction.
type InstLoad struct {
	// Parent basic block.
	parent *ir.BasicBlock
	// Name of the local variable associated with the instruction.
	name string
	// Type of the instruction.
	typ types.Type
	// Source address element type.
	elem types.Type
	// Source address.
	src value.Value
	// Track uses of the value.
	used
}

// NewLoad returns a new dummy load instruction based on the given element type
// and source address.
func NewLoad(elem types.Type, src value.Value) *InstLoad {
	t, ok := src.Type().(*types.PointerType)
	if !ok {
		panic(fmt.Sprintf("invalid source address type; expected *types.PointerType, got %T", src.Type()))
	}
	return &InstLoad{typ: t.Elem(), elem: elem, src: src}
}

// Type returns the type of the instruction.
func (inst *InstLoad) Type() types.Type {
	return inst.typ
}

// Ident returns the identifier associated with the instruction.
func (inst *InstLoad) Ident() string {
	return enc.Local(inst.name)
}

// Name returns the name of the local variable associated with the instruction.
func (inst *InstLoad) Name() string {
	return inst.name
}

// SetName sets the name of the local variable associated with the instruction.
func (inst *InstLoad) SetName(name string) {
	inst.name = name
}

// String returns the LLVM syntax representation of the instruction.
func (inst *InstLoad) String() string {
	src := inst.Src()
	return fmt.Sprintf("%s = load %s, %s %s",
		inst.Ident(),
		inst.Type(),
		src.Type(),
		src.Ident())
}

// Parent returns the parent basic block of the instruction.
func (inst *InstLoad) Parent() *ir.BasicBlock {
	return inst.parent
}

// SetParent sets the parent basic block of the instruction.
func (inst *InstLoad) SetParent(parent *ir.BasicBlock) {
	inst.parent = parent
}

// ElemType returns the source address element type of the load instruction.
func (inst *InstLoad) ElemType() types.Type {
	return inst.elem
}

// Src returns the source address of the load instruction.
func (inst *InstLoad) Src() value.Value {
	return inst.src
}

// SetSrc sets the source address of the load instruction.
func (inst *InstLoad) SetSrc(src value.Value) {
	inst.src = src
}

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
	// Track uses of the value.
	used
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
		inst.ElemType(),
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
