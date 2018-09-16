package ir

import (
	"fmt"

	"github.com/llir/l/ir/ll"
	"github.com/llir/l/ir/types"
	"github.com/llir/l/ir/value"
)

// --- [ Memory instructions ] -------------------------------------------------

// ~~~ [ alloca ] ~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~

// InstAlloca is an LLVM IR alloca instruction.
type InstAlloca struct {
	// Name of local variable associated with the result.
	LocalName string
	// Element type.
	ElemType types.Type
	// (optional) Number of elements; nil if not present.
	NElems value.Value
}

// NewAlloca returns a new alloca instruction based on the given element type.
func NewAlloca(elemType types.Type) *InstAlloca {
	return &InstAlloca{ElemType: elemType}
}

// String returns the LLVM syntax representation of the instruction as a
// type-value pair.
func (inst *InstAlloca) String() string {
	return fmt.Sprintf("%v %v", inst.Type(), inst.Ident())
}

// Type returns the type of the instruction.
func (inst *InstAlloca) Type() types.Type {
	panic("not yet implemented")
}

// Ident returns the identifier associated with the instruction.
func (inst *InstAlloca) Ident() string {
	panic("not yet implemented")
}

// Name returns the name of the instruction.
func (inst *InstAlloca) Name() string {
	return inst.LocalName
}

// SetName sets the name of the instruction.
func (inst *InstAlloca) SetName(name string) {
	inst.LocalName = name
}

// ~~~ [ load ] ~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~

// InstLoad is an LLVM IR load instruction.
type InstLoad struct {
	// Name of local variable associated with the result.
	LocalName string
	// Source address.
	Src value.Value
}

// NewLoad returns a new load instruction based on the given source address.
func NewLoad(src value.Value) *InstLoad {
	return &InstLoad{Src: src}
}

// String returns the LLVM syntax representation of the instruction as a
// type-value pair.
func (inst *InstLoad) String() string {
	return fmt.Sprintf("%v %v", inst.Type(), inst.Ident())
}

// Type returns the type of the instruction.
func (inst *InstLoad) Type() types.Type {
	panic("not yet implemented")
}

// Ident returns the identifier associated with the instruction.
func (inst *InstLoad) Ident() string {
	panic("not yet implemented")
}

// Name returns the name of the instruction.
func (inst *InstLoad) Name() string {
	return inst.LocalName
}

// SetName sets the name of the instruction.
func (inst *InstLoad) SetName(name string) {
	inst.LocalName = name
}

// ~~~ [ store ] ~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~

// InstStore is an LLVM IR store instruction.
type InstStore struct {
	// Source value.
	Src value.Value
	// Destination address.
	Dst value.Value
}

// NewStore returns a new store instruction based on the given source value and
// destination address.
func NewStore(src, dst value.Value) *InstStore {
	return &InstStore{Src: src, Dst: dst}
}

// String returns the LLVM syntax representation of the instruction as a
// type-value pair.
func (inst *InstStore) String() string {
	return fmt.Sprintf("%v %v", inst.Type(), inst.Ident())
}

// Type returns the type of the instruction.
func (inst *InstStore) Type() types.Type {
	panic("not yet implemented")
}

// Ident returns the identifier associated with the instruction.
func (inst *InstStore) Ident() string {
	panic("not yet implemented")
}

// ~~~ [ fence ] ~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~

// InstFence is an LLVM IR fence instruction.
type InstFence struct {
	// Atomic memory ordering constraints.
	Ordering ll.AtomicOrdering
}

// NewFence returns a new fence instruction based on the given atomic ordering.
func NewFence(ordering ll.AtomicOrdering) *InstFence {
	return &InstFence{Ordering: ordering}
}

// String returns the LLVM syntax representation of the instruction as a
// type-value pair.
func (inst *InstFence) String() string {
	return fmt.Sprintf("%v %v", inst.Type(), inst.Ident())
}

// Type returns the type of the instruction.
func (inst *InstFence) Type() types.Type {
	panic("not yet implemented")
}

// Ident returns the identifier associated with the instruction.
func (inst *InstFence) Ident() string {
	panic("not yet implemented")
}

// ~~~ [ cmpxchg ] ~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~

// InstCmpXchg is an LLVM IR cmpxchg instruction.
type InstCmpXchg struct {
	// Name of local variable associated with the result.
	LocalName string
	// Address to read from, compare against and store to.
	Ptr value.Value
	// Value to compare against.
	Cmp value.Value
	// New value to store.
	New value.Value
	// Atomic memory ordering constraints on success.
	Success ll.AtomicOrdering
	// Atomic memory ordering constraints on failure.
	Failure ll.AtomicOrdering
}

// NewCmpXchg returns a new cmpxchg instruction based on the given address,
// value to compare against, new value to store, and atomic orderings for
// success and failure.
func NewCmpXchg(ptr, cmp, new value.Value, success, failure ll.AtomicOrdering) *InstCmpXchg {
	return &InstCmpXchg{Ptr: ptr, Cmp: cmp, New: new, Success: success, Failure: failure}
}

// String returns the LLVM syntax representation of the instruction as a
// type-value pair.
func (inst *InstCmpXchg) String() string {
	return fmt.Sprintf("%v %v", inst.Type(), inst.Ident())
}

// Type returns the type of the instruction.
func (inst *InstCmpXchg) Type() types.Type {
	panic("not yet implemented")
}

// Ident returns the identifier associated with the instruction.
func (inst *InstCmpXchg) Ident() string {
	panic("not yet implemented")
}

// Name returns the name of the instruction.
func (inst *InstCmpXchg) Name() string {
	return inst.LocalName
}

// SetName sets the name of the instruction.
func (inst *InstCmpXchg) SetName(name string) {
	inst.LocalName = name
}

// ~~~ [ atomicrmw ] ~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~

// InstAtomicRMW is an LLVM IR atomicrmw instruction.
type InstAtomicRMW struct {
	// Name of local variable associated with the result.
	LocalName string
	// Atomic operation.
	Op ll.AtomicOp
	// Destination address.
	Dst value.Value
	// Operand.
	X value.Value
	// Atomic memory ordering constraints.
	Ordering ll.AtomicOrdering
}

// NewAtomicRMW returns a new atomicrmw instruction based on the given atomic
// operation, destination address, operand and atomic ordering.
func NewAtomicRMW(op ll.AtomicOp, dst, x value.Value, ordering ll.AtomicOrdering) *InstAtomicRMW {
	return &InstAtomicRMW{Op: op, Dst: dst, X: x, Ordering: ordering}
}

// String returns the LLVM syntax representation of the instruction as a
// type-value pair.
func (inst *InstAtomicRMW) String() string {
	return fmt.Sprintf("%v %v", inst.Type(), inst.Ident())
}

// Type returns the type of the instruction.
func (inst *InstAtomicRMW) Type() types.Type {
	panic("not yet implemented")
}

// Ident returns the identifier associated with the instruction.
func (inst *InstAtomicRMW) Ident() string {
	panic("not yet implemented")
}

// Name returns the name of the instruction.
func (inst *InstAtomicRMW) Name() string {
	return inst.LocalName
}

// SetName sets the name of the instruction.
func (inst *InstAtomicRMW) SetName(name string) {
	inst.LocalName = name
}

// ~~~ [ getelementptr ] ~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~

// InstGetElementPtr is an LLVM IR getelementptr instruction.
type InstGetElementPtr struct {
	// Name of local variable associated with the result.
	LocalName string
	// Element type.
	ElemType types.Type
	// Source address.
	Src value.Value
	// Element indicies.
	Indices []value.Value
}

// NewGetElementPtr returns a new getelementptr instruction based on the given
// element type, source address and element indices.
func NewGetElementPtr(elemType types.Type, src value.Value, indices ...value.Value) *InstGetElementPtr {
	return &InstGetElementPtr{ElemType: elemType, Src: src, Indices: indices}
}

// String returns the LLVM syntax representation of the instruction as a
// type-value pair.
func (inst *InstGetElementPtr) String() string {
	return fmt.Sprintf("%v %v", inst.Type(), inst.Ident())
}

// Type returns the type of the instruction.
func (inst *InstGetElementPtr) Type() types.Type {
	// TODO: cache type?
	return types.NewPointer(inst.ElemType)
}

// Ident returns the identifier associated with the instruction.
func (inst *InstGetElementPtr) Ident() string {
	panic("not yet implemented")
}

// Name returns the name of the instruction.
func (inst *InstGetElementPtr) Name() string {
	return inst.LocalName
}

// SetName sets the name of the instruction.
func (inst *InstGetElementPtr) SetName(name string) {
	inst.LocalName = name
}
