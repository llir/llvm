package ir

import (
	"fmt"
	"strings"

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
	LocalIdent
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
	Align Align
	// (optional) Metadata.
	Metadata []*metadata.MetadataAttachment
}

// NewAlloca returns a new alloca instruction based on the given element type.
func NewAlloca(elemType types.Type) *InstAlloca {
	inst := &InstAlloca{ElemType: elemType}
	// Compute type.
	inst.Type()
	return inst
}

// String returns the LLVM syntax representation of the instruction as a
// type-value pair.
func (inst *InstAlloca) String() string {
	return fmt.Sprintf("%s %s", inst.Type(), inst.Ident())
}

// Type returns the type of the instruction.
func (inst *InstAlloca) Type() types.Type {
	// Cache type if not present.
	if inst.Typ == nil {
		inst.Typ = types.NewPointer(inst.ElemType)
	}
	return inst.Typ
}

// Def returns the LLVM syntax representation of the instruction.
func (inst *InstAlloca) Def() string {
	// 'alloca' InAllocaopt SwiftErroropt ElemType=Type NElems=(',' TypeValue)?
	// (',' Alignment)? (',' AddrSpace)? Metadata=(',' MetadataAttachment)+?
	buf := &strings.Builder{}
	fmt.Fprintf(buf, "%s = ", inst.Ident())
	buf.WriteString("alloca")
	if inst.InAlloca {
		buf.WriteString(" inalloca")
	}
	if inst.SwiftError {
		buf.WriteString(" swifterror")
	}
	fmt.Fprintf(buf, " %s", inst.ElemType)
	if inst.NElems != nil {
		fmt.Fprintf(buf, ", %s", inst.NElems)
	}
	if inst.Align != 0 {
		fmt.Fprintf(buf, ", %s", inst.Align)
	}
	if inst.Typ.AddrSpace != 0 {
		fmt.Fprintf(buf, ", %s", inst.Typ.AddrSpace)
	}
	for _, md := range inst.Metadata {
		fmt.Fprintf(buf, ", %s", md)
	}
	return buf.String()
}

// ~~~ [ load ] ~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~

// InstLoad is an LLVM IR load instruction.
type InstLoad struct {
	// Name of local variable associated with the result.
	LocalIdent
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
	Align Align
	// (optional) Metadata.
	Metadata []*metadata.MetadataAttachment
}

// NewLoad returns a new load instruction based on the given source address.
func NewLoad(src value.Value) *InstLoad {
	inst := &InstLoad{Src: src}
	// Compute type.
	inst.Type()
	return inst
}

// String returns the LLVM syntax representation of the instruction as a
// type-value pair.
func (inst *InstLoad) String() string {
	return fmt.Sprintf("%s %s", inst.Type(), inst.Ident())
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

// Def returns the LLVM syntax representation of the instruction.
func (inst *InstLoad) Def() string {
	// Load instruction.
	//
	//    'load' Volatileopt ElemType=Type ',' Src=TypeValue (',' Alignment)?
	//    Metadata=(',' MetadataAttachment)+?
	//
	// Atomic load instruction.
	//
	//    'load' Atomic Volatileopt ElemType=Type ',' Src=TypeValue SyncScopeopt
	//    Ordering=AtomicOrdering (',' Alignment)? Metadata=(','
	//    MetadataAttachment)+?
	buf := &strings.Builder{}
	fmt.Fprintf(buf, "%s = ", inst.Ident())
	buf.WriteString("load")
	if inst.Atomic {
		buf.WriteString(" atomic")
	}
	if inst.Volatile {
		buf.WriteString(" volatile")
	}
	fmt.Fprintf(buf, " %s, %s", inst.Typ, inst.Src)
	if len(inst.SyncScope) > 0 {
		fmt.Fprintf(buf, " syncscope(%s)", quote(inst.SyncScope))
	}
	if inst.Ordering != enum.AtomicOrderingNone {
		fmt.Fprintf(buf, " %s", inst.Ordering)
	}
	if inst.Align != 0 {
		fmt.Fprintf(buf, ", %s", inst.Align)
	}
	for _, md := range inst.Metadata {
		fmt.Fprintf(buf, ", %s", md)
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
	Align Align
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
	// Store instruction.
	//
	//    'store' Volatileopt Src=TypeValue ',' Dst=TypeValue (',' Alignment)?
	//    Metadata=(',' MetadataAttachment)+?
	//
	// Atomic store instruction.
	//
	//    'store' Atomic Volatileopt Src=TypeValue ',' Dst=TypeValue SyncScopeopt
	//    Ordering=AtomicOrdering (',' Alignment)? Metadata=(','
	//    MetadataAttachment)+?
	buf := &strings.Builder{}
	buf.WriteString("store")
	if inst.Atomic {
		buf.WriteString(" atomic")
	}
	if inst.Volatile {
		buf.WriteString(" volatile")
	}
	fmt.Fprintf(buf, " %s, %s", inst.Src, inst.Dst)
	if len(inst.SyncScope) > 0 {
		fmt.Fprintf(buf, " syncscope(%s)", quote(inst.SyncScope))
	}
	if inst.Ordering != enum.AtomicOrderingNone {
		fmt.Fprintf(buf, " %s", inst.Ordering)
	}
	if inst.Align != 0 {
		fmt.Fprintf(buf, ", %s", inst.Align)
	}
	for _, md := range inst.Metadata {
		fmt.Fprintf(buf, ", %s", md)
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
	// 'fence' SyncScopeopt Ordering=AtomicOrdering Metadata=(','
	// MetadataAttachment)+?
	buf := &strings.Builder{}
	buf.WriteString("fence")
	if len(inst.SyncScope) > 0 {
		fmt.Fprintf(buf, " syncscope(%s)", quote(inst.SyncScope))
	}
	fmt.Fprintf(buf, " %s", inst.Ordering)
	for _, md := range inst.Metadata {
		fmt.Fprintf(buf, ", %s", md)
	}
	return buf.String()
}

// ~~~ [ cmpxchg ] ~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~

// InstCmpXchg is an LLVM IR cmpxchg instruction.
type InstCmpXchg struct {
	// Name of local variable associated with the result.
	LocalIdent
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
	inst := &InstCmpXchg{Ptr: ptr, Cmp: cmp, New: new, SuccessOrdering: successOrdering, FailureOrdering: failureOrdering}
	// Compute type.
	inst.Type()
	return inst
}

// String returns the LLVM syntax representation of the instruction as a
// type-value pair.
func (inst *InstCmpXchg) String() string {
	return fmt.Sprintf("%s %s", inst.Type(), inst.Ident())
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

// Def returns the LLVM syntax representation of the instruction.
func (inst *InstCmpXchg) Def() string {
	// 'cmpxchg' Weakopt Volatileopt Ptr=TypeValue ',' Cmp=TypeValue ','
	// New=TypeValue SyncScopeopt SuccessOrdering=AtomicOrdering
	// FailureOrdering=AtomicOrdering Metadata=(',' MetadataAttachment)+?
	buf := &strings.Builder{}
	fmt.Fprintf(buf, "%s = ", inst.Ident())
	buf.WriteString("cmpxchg")
	if inst.Weak {
		buf.WriteString(" weak")
	}
	if inst.Volatile {
		buf.WriteString(" volatile")
	}
	fmt.Fprintf(buf, " %s, %s, %s", inst.Ptr, inst.Cmp, inst.New)
	if len(inst.SyncScope) > 0 {
		fmt.Fprintf(buf, " syncscope(%s)", quote(inst.SyncScope))
	}
	fmt.Fprintf(buf, " %s", inst.SuccessOrdering)
	fmt.Fprintf(buf, " %s", inst.FailureOrdering)
	for _, md := range inst.Metadata {
		fmt.Fprintf(buf, ", %s", md)
	}
	return buf.String()
}

// ~~~ [ atomicrmw ] ~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~

// InstAtomicRMW is an LLVM IR atomicrmw instruction.
type InstAtomicRMW struct {
	// Name of local variable associated with the result.
	LocalIdent
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
	inst := &InstAtomicRMW{Op: op, Dst: dst, X: x, Ordering: ordering}
	// Compute type.
	inst.Type()
	return inst
}

// String returns the LLVM syntax representation of the instruction as a
// type-value pair.
func (inst *InstAtomicRMW) String() string {
	return fmt.Sprintf("%s %s", inst.Type(), inst.Ident())
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

// Def returns the LLVM syntax representation of the instruction.
func (inst *InstAtomicRMW) Def() string {
	// 'atomicrmw' Volatileopt Op=AtomicOp Dst=TypeValue ',' X=TypeValue
	// SyncScopeopt Ordering=AtomicOrdering Metadata=(',' MetadataAttachment)+?
	buf := &strings.Builder{}
	fmt.Fprintf(buf, "%s = ", inst.Ident())
	buf.WriteString("atomicrmw")
	if inst.Volatile {
		buf.WriteString(" volatile")
	}
	fmt.Fprintf(buf, " %s %s, %s", inst.Op, inst.Dst, inst.X)
	if len(inst.SyncScope) > 0 {
		fmt.Fprintf(buf, " syncscope(%s)", quote(inst.SyncScope))
	}
	fmt.Fprintf(buf, " %s", inst.Ordering)
	for _, md := range inst.Metadata {
		fmt.Fprintf(buf, ", %s", md)
	}
	return buf.String()
}

// ~~~ [ getelementptr ] ~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~

// InstGetElementPtr is an LLVM IR getelementptr instruction.
type InstGetElementPtr struct {
	// Name of local variable associated with the result.
	LocalIdent
	// Element type.
	ElemType types.Type
	// Source address.
	Src value.Value
	// Element indicies.
	Indices []value.Value

	// extra.

	// Type of result produced by the instruction.
	Typ types.Type // *types.PointerType or *types.VectorType (with elements of pointer type)
	// (optional) In-bounds.
	InBounds bool
	// (optional) Metadata.
	Metadata []*metadata.MetadataAttachment
}

// NewGetElementPtr returns a new getelementptr instruction based on the given
// source address and element indices.
func NewGetElementPtr(src value.Value, indices ...value.Value) *InstGetElementPtr {
	inst := &InstGetElementPtr{Src: src, Indices: indices}
	// Compute type.
	inst.Type()
	return inst
}

// String returns the LLVM syntax representation of the instruction as a
// type-value pair.
func (inst *InstGetElementPtr) String() string {
	return fmt.Sprintf("%s %s", inst.Type(), inst.Ident())
}

// Type returns the type of the instruction.
func (inst *InstGetElementPtr) Type() types.Type {
	// Cache element type if not present.
	if inst.ElemType == nil {
		switch typ := inst.Src.Type().(type) {
		case *types.PointerType:
			inst.ElemType = typ.ElemType
		case *types.VectorType:
			t, ok := typ.ElemType.(*types.PointerType)
			if !ok {
				panic(fmt.Errorf("invalid vector element type; expected *types.Pointer, got %T", typ.ElemType))
			}
			inst.ElemType = t.ElemType
		default:
			panic(fmt.Errorf("support for souce type %T not yet implemented", typ))
		}
	}
	// Cache type if not present.
	if inst.Typ == nil {
		inst.Typ = gepType(inst.ElemType, inst.Indices)
	}
	return inst.Typ
}

// Def returns the LLVM syntax representation of the instruction.
func (inst *InstGetElementPtr) Def() string {
	// 'getelementptr' InBoundsopt ElemType=Type ',' Src=TypeValue Indices=(','
	// TypeValue)* Metadata=(',' MetadataAttachment)+?
	buf := &strings.Builder{}
	fmt.Fprintf(buf, "%s = ", inst.Ident())
	buf.WriteString("getelementptr")
	if inst.InBounds {
		buf.WriteString(" inbounds")
	}
	fmt.Fprintf(buf, " %s, %s", inst.ElemType, inst.Src)
	for _, index := range inst.Indices {
		fmt.Fprintf(buf, ", %s", index)
	}
	for _, md := range inst.Metadata {
		fmt.Fprintf(buf, ", %s", md)
	}
	return buf.String()
}

// ### [ Helper functions ] ####################################################

// gepType returns the pointer type or vector of pointers type to the element at
// the position in the type specified by the given indices, as calculated by the
// getelementptr instruction.
func gepType(elemType types.Type, indices []value.Value) types.Type {
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
			panic(fmt.Errorf("unable to index into element of pointer type `%s`; for more information, see http://llvm.org/docs/GetElementPtr.html#what-is-dereferenced-by-gep", elemType))
		case *types.VectorType:
			e = t.ElemType
		case *types.ArrayType:
			e = t.ElemType
		case *types.StructType:
			switch index := index.(type) {
			case *constant.Int:
				e = t.Fields[index.X.Int64()]
			case *constant.Vector:
				// TODO: Validate how index vectors in gep are supposed to work.
				idx, ok := index.Elems[0].(*constant.Int)
				if !ok {
					panic(fmt.Errorf("invalid index type for structure element; expected *constant.Int, got %T", index.Elems[0]))
				}
				e = t.Fields[idx.X.Int64()]
			case *constant.ZeroInitializer:
				e = t.Fields[0]
			default:
				panic(fmt.Errorf("invalid index type for structure element; expected *constant.Int, *constant.Vector or *constant.ZeroInitializer, got %T", index))
			}
		default:
			panic(fmt.Errorf("support for indexing element type %T not yet implemented", e))
		}
	}
	// TODO: Validate how index vectors in gep are supposed to work.
	//
	// Example from dir.ll:
	//    %113 = getelementptr inbounds %struct.fileinfo, %struct.fileinfo* %96, <2 x i64> %110, !dbg !4736
	//    %116 = bitcast i8** %115 to <2 x %struct.fileinfo*>*, !dbg !4738
	//    store <2 x %struct.fileinfo*> %113, <2 x %struct.fileinfo*>* %116, align 8, !dbg !4738, !tbaa !1793
	if len(indices) > 0 {
		if t, ok := indices[0].Type().(*types.VectorType); ok {
			return types.NewVector(t.Len, types.NewPointer(e))
		}
	}
	return types.NewPointer(e)
}
