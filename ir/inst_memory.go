package ir

import (
	"fmt"
	"strings"

	"github.com/llir/llvm/internal/enc"
	"github.com/llir/llvm/ir/constant"
	"github.com/llir/llvm/ir/enum"
	"github.com/llir/llvm/ir/metadata"
	"github.com/llir/llvm/ir/types"
	"github.com/llir/llvm/ir/value"
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

	// extra.

	// Type of result produced by the instruction, including an optional address
	// space.
	Typ *types.PointerType
	// (optional) In-alloca.
	InAlloca bool
	// (optional) Swift error.
	SwiftError bool
	// (optional) Alignment; zero if not present.
	Alignment int
	// (optional) Metadata.
	Metadata []*metadata.MetadataAttachment
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
	// Cache type if not present.
	if inst.Typ == nil {
		inst.Typ = types.NewPointer(inst.ElemType)
	}
	return inst.Typ
}

// Ident returns the identifier associated with the instruction.
func (inst *InstAlloca) Ident() string {
	return enc.Local(inst.LocalName)
}

// Name returns the name of the instruction.
func (inst *InstAlloca) Name() string {
	return inst.LocalName
}

// SetName sets the name of the instruction.
func (inst *InstAlloca) SetName(name string) {
	inst.LocalName = name
}

// Def returns the LLVM syntax representation of the instruction.
func (inst *InstAlloca) Def() string {
	// "alloca" OptInAlloca OptSwiftError Type OptCommaTypeValue OptCommaAlignment OptCommaAddrSpace OptCommaSepMetadataAttachmentList
	buf := &strings.Builder{}
	fmt.Fprintf(buf, "%v = ", inst.Ident())
	buf.WriteString("alloca")
	if inst.InAlloca {
		buf.WriteString(" inalloca")
	}
	if inst.SwiftError {
		buf.WriteString(" swifterror")
	}
	fmt.Fprintf(buf, " %v", inst.ElemType)
	if inst.NElems != nil {
		fmt.Fprintf(buf, ", %v", inst.NElems)
	}
	if inst.Alignment != 0 {
		fmt.Fprintf(buf, ", align %v", inst.Alignment)
	}
	if inst.Typ.AddrSpace != 0 {
		fmt.Fprintf(buf, ", %v", inst.Typ.AddrSpace)
	}
	for _, md := range inst.Metadata {
		fmt.Fprintf(buf, ", %v", md)
	}
	return buf.String()
}

// ~~~ [ load ] ~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~

// InstLoad is an LLVM IR load instruction.
type InstLoad struct {
	// Name of local variable associated with the result.
	LocalName string
	// Source address.
	Src value.Value

	// extra.

	// Type of result produced by the instruction.
	Typ types.Type
	// (optional) Atomic.
	Atomic bool
	// (optional) Volatile.
	Volatile bool
	// (optional) Sync scope; empty if not present.
	SyncScope string
	// (optional) Atomic memory ordering constraints; zero if not present.
	Ordering enum.AtomicOrdering
	// (optional) Alignment; zero if not present.
	Alignment int
	// (optional) Metadata.
	Metadata []*metadata.MetadataAttachment
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
	// Cache type if not present.
	if inst.Typ == nil {
		t, ok := inst.Src.Type().(*types.PointerType)
		if !ok {
			panic(fmt.Errorf("invalid source type; expected *types.PointerType, got %T", inst.Src.Type()))
		}
		inst.Typ = t.ElemType
	}
	return inst.Typ
}

// Ident returns the identifier associated with the instruction.
func (inst *InstLoad) Ident() string {
	return enc.Local(inst.LocalName)
}

// Name returns the name of the instruction.
func (inst *InstLoad) Name() string {
	return inst.LocalName
}

// SetName sets the name of the instruction.
func (inst *InstLoad) SetName(name string) {
	inst.LocalName = name
}

// Def returns the LLVM syntax representation of the instruction.
func (inst *InstLoad) Def() string {
	// "load" "atomic" OptVolatile Type "," Type Value OptSyncScope AtomicOrdering OptCommaAlignment OptCommaSepMetadataAttachmentList
	// "load" OptVolatile Type "," Type Value OptCommaAlignment OptCommaSepMetadataAttachmentList
	buf := &strings.Builder{}
	fmt.Fprintf(buf, "%v = ", inst.Ident())
	buf.WriteString("load")
	if inst.Atomic {
		buf.WriteString(" atomic")
	}
	if inst.Volatile {
		buf.WriteString(" volatile")
	}
	fmt.Fprintf(buf, " %v, %v", inst.Typ, inst.Src)
	if len(inst.SyncScope) > 0 {
		fmt.Fprintf(buf, " syncscope(%v)", enc.Quote([]byte(inst.SyncScope)))
	}
	if inst.Ordering != enum.AtomicOrderingNone {
		fmt.Fprintf(buf, " %v", inst.Ordering)
	}
	if inst.Alignment != 0 {
		fmt.Fprintf(buf, ", align %v", inst.Alignment)
	}
	for _, md := range inst.Metadata {
		fmt.Fprintf(buf, ", %v", md)
	}
	return buf.String()
}

// ~~~ [ store ] ~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~

// InstStore is an LLVM IR store instruction.
type InstStore struct {
	// Source value.
	Src value.Value
	// Destination address.
	Dst value.Value

	// extra.

	// (optional) Atomic.
	Atomic bool
	// (optional) Volatile.
	Volatile bool
	// (optional) Sync scope; empty if not present.
	SyncScope string
	// (optional) Atomic memory ordering constraints; zero if not present.
	Ordering enum.AtomicOrdering
	// (optional) Alignment; zero if not present.
	Alignment int
	// (optional) Metadata.
	Metadata []*metadata.MetadataAttachment
}

// NewStore returns a new store instruction based on the given source value and
// destination address.
func NewStore(src, dst value.Value) *InstStore {
	return &InstStore{Src: src, Dst: dst}
}

// Def returns the LLVM syntax representation of the instruction.
func (inst *InstStore) Def() string {
	// "store" "atomic" OptVolatile Type Value "," Type Value OptSyncScope AtomicOrdering OptCommaAlignment OptCommaSepMetadataAttachmentList
	// "store" OptVolatile Type Value "," Type Value OptCommaAlignment OptCommaSepMetadataAttachmentList
	buf := &strings.Builder{}
	buf.WriteString("store")
	if inst.Atomic {
		buf.WriteString(" atomic")
	}
	if inst.Volatile {
		buf.WriteString(" volatile")
	}
	fmt.Fprintf(buf, " %v, %v", inst.Src, inst.Dst)
	if len(inst.SyncScope) > 0 {
		fmt.Fprintf(buf, " syncscope(%v)", enc.Quote([]byte(inst.SyncScope)))
	}
	if inst.Ordering != enum.AtomicOrderingNone {
		fmt.Fprintf(buf, " %v", inst.Ordering)
	}
	if inst.Alignment != 0 {
		fmt.Fprintf(buf, ", align %v", inst.Alignment)
	}
	for _, md := range inst.Metadata {
		fmt.Fprintf(buf, ", %v", md)
	}
	return buf.String()
}

// ~~~ [ fence ] ~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~

// InstFence is an LLVM IR fence instruction.
type InstFence struct {
	// Atomic memory ordering constraints.
	Ordering enum.AtomicOrdering

	// extra.

	// (optional) Sync scope; empty if not present.
	SyncScope string
	// (optional) Metadata.
	Metadata []*metadata.MetadataAttachment
}

// NewFence returns a new fence instruction based on the given atomic ordering.
func NewFence(ordering enum.AtomicOrdering) *InstFence {
	return &InstFence{Ordering: ordering}
}

// Def returns the LLVM syntax representation of the instruction.
func (inst *InstFence) Def() string {
	// "fence" OptSyncScope AtomicOrdering OptCommaSepMetadataAttachmentList
	buf := &strings.Builder{}
	buf.WriteString("fence")
	if len(inst.SyncScope) > 0 {
		fmt.Fprintf(buf, " syncscope(%v)", enc.Quote([]byte(inst.SyncScope)))
	}
	fmt.Fprintf(buf, " %v", inst.Ordering)
	for _, md := range inst.Metadata {
		fmt.Fprintf(buf, ", %v", md)
	}
	return buf.String()
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
	SuccessOrdering enum.AtomicOrdering
	// Atomic memory ordering constraints on failure.
	FailureOrdering enum.AtomicOrdering

	// extra.

	// Type of result produced by the instruction; the first field of the struct
	// holds the old value, and the second field indicates success.
	Typ *types.StructType
	// (optional) Weak.
	Weak bool
	// (optional) Volatile.
	Volatile bool
	// (optional) Sync scope; empty if not present.
	SyncScope string
	// (optional) Metadata.
	Metadata []*metadata.MetadataAttachment
}

// NewCmpXchg returns a new cmpxchg instruction based on the given address,
// value to compare against, new value to store, and atomic orderings for
// success and failure.
func NewCmpXchg(ptr, cmp, new value.Value, successOrdering, failureOrdering enum.AtomicOrdering) *InstCmpXchg {
	return &InstCmpXchg{Ptr: ptr, Cmp: cmp, New: new, SuccessOrdering: successOrdering, FailureOrdering: failureOrdering}
}

// String returns the LLVM syntax representation of the instruction as a
// type-value pair.
func (inst *InstCmpXchg) String() string {
	return fmt.Sprintf("%v %v", inst.Type(), inst.Ident())
}

// Type returns the type of the instruction.
func (inst *InstCmpXchg) Type() types.Type {
	// Cache type if not present.
	if inst.Typ == nil {
		oldType := inst.New.Type()
		inst.Typ = types.NewStruct(oldType, types.I1)
	}
	return inst.Typ
}

// Ident returns the identifier associated with the instruction.
func (inst *InstCmpXchg) Ident() string {
	return enc.Local(inst.LocalName)
}

// Name returns the name of the instruction.
func (inst *InstCmpXchg) Name() string {
	return inst.LocalName
}

// SetName sets the name of the instruction.
func (inst *InstCmpXchg) SetName(name string) {
	inst.LocalName = name
}

// Def returns the LLVM syntax representation of the instruction.
func (inst *InstCmpXchg) Def() string {
	// "cmpxchg" OptWeak OptVolatile Type Value "," Type Value "," Type Value OptSyncScope AtomicOrdering AtomicOrdering OptCommaSepMetadataAttachmentList
	buf := &strings.Builder{}
	fmt.Fprintf(buf, "%v = ", inst.Ident())
	buf.WriteString("cmpxchg")
	if inst.Weak {
		buf.WriteString(" weak")
	}
	if inst.Volatile {
		buf.WriteString(" volatile")
	}
	fmt.Fprintf(buf, " %v, %v, %v", inst.Ptr, inst.Cmp, inst.New)
	if len(inst.SyncScope) > 0 {
		fmt.Fprintf(buf, " syncscope(%v)", enc.Quote([]byte(inst.SyncScope)))
	}
	fmt.Fprintf(buf, " %v", inst.SuccessOrdering)
	fmt.Fprintf(buf, " %v", inst.FailureOrdering)
	for _, md := range inst.Metadata {
		fmt.Fprintf(buf, ", %v", md)
	}
	return buf.String()
}

// ~~~ [ atomicrmw ] ~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~

// InstAtomicRMW is an LLVM IR atomicrmw instruction.
type InstAtomicRMW struct {
	// Name of local variable associated with the result.
	LocalName string
	// Atomic operation.
	Op enum.AtomicOp
	// Destination address.
	Dst value.Value
	// Operand.
	X value.Value
	// Atomic memory ordering constraints.
	Ordering enum.AtomicOrdering

	// extra.

	// Type of result produced by the instruction.
	Typ types.Type
	// (optional) Volatile.
	Volatile bool
	// (optional) Sync scope; empty if not present.
	SyncScope string
	// (optional) Metadata.
	Metadata []*metadata.MetadataAttachment
}

// NewAtomicRMW returns a new atomicrmw instruction based on the given atomic
// operation, destination address, operand and atomic ordering.
func NewAtomicRMW(op enum.AtomicOp, dst, x value.Value, ordering enum.AtomicOrdering) *InstAtomicRMW {
	return &InstAtomicRMW{Op: op, Dst: dst, X: x, Ordering: ordering}
}

// String returns the LLVM syntax representation of the instruction as a
// type-value pair.
func (inst *InstAtomicRMW) String() string {
	return fmt.Sprintf("%v %v", inst.Type(), inst.Ident())
}

// Type returns the type of the instruction.
func (inst *InstAtomicRMW) Type() types.Type {
	// Cache type if not present.
	if inst.Typ == nil {
		t, ok := inst.Dst.Type().(*types.PointerType)
		if !ok {
			panic(fmt.Errorf("invalid destination type; expected *types.PointerType, got %T", inst.Dst.Type()))
		}
		inst.Typ = t.ElemType
	}
	return inst.Typ
}

// Ident returns the identifier associated with the instruction.
func (inst *InstAtomicRMW) Ident() string {
	return enc.Local(inst.LocalName)
}

// Name returns the name of the instruction.
func (inst *InstAtomicRMW) Name() string {
	return inst.LocalName
}

// SetName sets the name of the instruction.
func (inst *InstAtomicRMW) SetName(name string) {
	inst.LocalName = name
}

// Def returns the LLVM syntax representation of the instruction.
func (inst *InstAtomicRMW) Def() string {
	// "atomicrmw" OptVolatile BinOp Type Value "," Type Value OptSyncScope AtomicOrdering OptCommaSepMetadataAttachmentList
	buf := &strings.Builder{}
	fmt.Fprintf(buf, "%v = ", inst.Ident())
	buf.WriteString("atomicrmw")
	if inst.Volatile {
		buf.WriteString(" volatile")
	}
	fmt.Fprintf(buf, " %v %v, %v", inst.Op, inst.Dst, inst.X)
	if len(inst.SyncScope) > 0 {
		fmt.Fprintf(buf, " syncscope(%v)", enc.Quote([]byte(inst.SyncScope)))
	}
	fmt.Fprintf(buf, " %v", inst.Ordering)
	for _, md := range inst.Metadata {
		fmt.Fprintf(buf, ", %v", md)
	}
	return buf.String()
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

	// extra.

	// Type of result produced by the instruction.
	Typ types.Type
	// (optional) In-bounds.
	InBounds bool
	// (optional) Metadata.
	Metadata []*metadata.MetadataAttachment
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
	// Cache type if not present.
	if inst.Typ == nil {
		inst.Typ = gepType(inst.ElemType, inst.Indices)
	}
	return inst.Typ
}

// Ident returns the identifier associated with the instruction.
func (inst *InstGetElementPtr) Ident() string {
	return enc.Local(inst.LocalName)
}

// Name returns the name of the instruction.
func (inst *InstGetElementPtr) Name() string {
	return inst.LocalName
}

// SetName sets the name of the instruction.
func (inst *InstGetElementPtr) SetName(name string) {
	inst.LocalName = name
}

// Def returns the LLVM syntax representation of the instruction.
func (inst *InstGetElementPtr) Def() string {
	// "getelementptr" OptInBounds Type "," Type Value GEPIndices OptCommaSepMetadataAttachmentList
	buf := &strings.Builder{}
	fmt.Fprintf(buf, "%v = ", inst.Ident())
	buf.WriteString("getelementptr")
	if inst.InBounds {
		buf.WriteString(" inbounds")
	}
	fmt.Fprintf(buf, " %v, %v", inst.ElemType, inst.Src)
	for _, index := range inst.Indices {
		fmt.Fprintf(buf, ", %v", index)
	}
	for _, md := range inst.Metadata {
		fmt.Fprintf(buf, ", %v", md)
	}
	return buf.String()
}

// ### [ Helper functions ] ####################################################

// gepType returns the pointer type to the element at the position in the type
// specified by the given indices, as calculated by the getelementptr
// instruction.
func gepType(elemType types.Type, indices []value.Value) *types.PointerType {
	e := elemType
	for i, index := range indices {
		if i == 0 {
			// Ignore checking the 0th index as it simply follows the pointer of
			// src.
			//
			// ref: http://llvm.org/docs/GetElementPtr.html#why-is-the-extra-0-index-required
			continue
		}
		switch t := e.(type) {
		case *types.PointerType:
			// ref: http://llvm.org/docs/GetElementPtr.html#what-is-dereferenced-by-gep
			panic(fmt.Errorf("unable to index into element of pointer type `%v`; for more information, see http://llvm.org/docs/GetElementPtr.html#what-is-dereferenced-by-gep", elemType))
		case *types.VectorType:
			e = t.ElemType
		case *types.ArrayType:
			e = t.ElemType
		case *types.StructType:
			idx, ok := index.(*constant.Int)
			if !ok {
				panic(fmt.Errorf("invalid index type for structure element; expected *constant.Int, got %T", index))
			}
			e = t.Fields[idx.X.Int64()]
		default:
			panic(fmt.Errorf("support for indexing element type %T not yet implemented", e))
		}
	}
	return types.NewPointer(e)
}
