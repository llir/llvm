package ir

import (
	"github.com/llir/l/ir/ll"
	"github.com/llir/l/ir/types"
	"github.com/llir/l/ir/value"
)

// --- [ Memory instructions ] -------------------------------------------------

// ~~~ [ alloca ] ~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~

// Alloca is an LLVM IR alloca instruction.
type Alloca struct {
	// Name of local variable associated with the result.
	LocalName string
	// Element type.
	ElemType types.Type
	// (optional) Number of elements; nil if not present.
	NElems value.Value
}

// NewAlloca returns a new alloca instruction based on the given element type.
func NewAlloca(elemType types.Type) *Alloca {
	return &Alloca{ElemType: elemType}
}

// ~~~ [ load ] ~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~

// Load is an LLVM IR load instruction.
type Load struct {
	// Name of local variable associated with the result.
	LocalName string
	// Source address.
	Src value.Value
}

// NewLoad returns a new load instruction based on the given source address.
func NewLoad(src value.Value) *Load {
	return &Load{Src: src}
}

// ~~~ [ store ] ~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~

// Store is an LLVM IR store instruction.
type Store struct {
	// Source value.
	Src value.Value
	// Destination address.
	Dst value.Value
}

// NewStore returns a new store instruction based on the given source value and
// destination address.
func NewStore(src, dst value.Value) *Store {
	return &Store{Src: src, Dst: dst}
}

// ~~~ [ fence ] ~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~

// Fence is an LLVM IR fence instruction.
type Fence struct {
	// Atomic memory ordering constraints.
	Ordering ll.AtomicOrdering
}

// NewFence returns a new fence instruction based on the given atomic ordering.
func NewFence(ordering ll.AtomicOrdering) *Fence {
	return &Fence{Ordering: ordering}
}

// ~~~ [ cmpxchg ] ~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~

// CmpXchg is an LLVM IR cmpxchg instruction.
type CmpXchg struct {
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
func NewCmpXchg(ptr, cmp, new value.Value, success, failure ll.AtomicOrdering) *CmpXchg {
	return &CmpXchg{Ptr: ptr, Cmp: cmp, New: new, Success: success, Failure: failure}
}

// ~~~ [ atomicrmw ] ~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~

// AtomicRMW is an LLVM IR atomicrmw instruction.
type AtomicRMW struct {
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
func NewAtomicRMW(op ll.AtomicOp, dst, x value.Value, ordering ll.AtomicOrdering) *AtomicRMW {
	return &AtomicRMW{Op: op, Dst: dst, X: x, Ordering: ordering}
}

// ~~~ [ getelementptr ] ~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~

// GetElementPtr is an LLVM IR getelementptr instruction.
type GetElementPtr struct {
	// Name of local variable associated with the result.
	LocalName string
	// Source address.
	Src value.Value
	// Element indicies.
	Indices []value.Value
}

// NewGetElementPtr returns a new getelementptr instruction based on the given
// source address and element indices.
func NewGetElementPtr(src value.Value, indices ...value.Value) *GetElementPtr {
	return &GetElementPtr{Src: src, Indices: indices}
}
