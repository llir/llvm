// === [ Memory instructions ] =================================================
//
// References:
//    http://llvm.org/docs/LangRef.html#memory-access-and-addressing-operations

package ir

import (
	"bytes"
	"fmt"

	"github.com/llir/llvm/internal/enc"
	"github.com/llir/llvm/ir/constant"
	"github.com/llir/llvm/ir/types"
	"github.com/llir/llvm/ir/value"
)

// --- [ alloca ] --------------------------------------------------------------

// InstAlloca represents an alloca instruction.
//
// References:
//    http://llvm.org/docs/LangRef.html#alloca-instruction
type InstAlloca struct {
	// Parent basic block.
	parent *BasicBlock
	// Identifier associated with the instruction.
	ident string
	// Element type.
	elem types.Type
	// Type of the instruction.
	typ *types.PointerType
	// Number of elements; or nil if one element.
	nelems value.Value
}

// NewAlloca returns a new alloca instruction based on the given element type.
func NewAlloca(elem types.Type) *InstAlloca {
	typ := types.NewPointer(elem)
	return &InstAlloca{elem: elem, typ: typ}
}

// Type returns the type of the instruction.
func (inst *InstAlloca) Type() types.Type {
	return inst.typ
}

// Ident returns the identifier associated with the instruction.
func (inst *InstAlloca) Ident() string {
	return enc.Local(inst.ident)
}

// SetIdent sets the identifier associated with the instruction.
func (inst *InstAlloca) SetIdent(ident string) {
	inst.ident = ident
}

// String returns the LLVM syntax representation of the instruction.
func (inst *InstAlloca) String() string {
	if nelems, ok := inst.NElems(); ok {
		return fmt.Sprintf("%s = alloca %s, %s %s",
			inst.Ident(),
			inst.ElemType(),
			nelems.Type(),
			nelems.Ident())
	}
	return fmt.Sprintf("%s = alloca %s",
		inst.Ident(),
		inst.ElemType())
}

// Parent returns the parent basic block of the instruction.
func (inst *InstAlloca) Parent() *BasicBlock {
	return inst.parent
}

// SetParent sets the parent basic block of the instruction.
func (inst *InstAlloca) SetParent(parent *BasicBlock) {
	inst.parent = parent
}

// ElemType returns the element type of the alloca instruction.
func (inst *InstAlloca) ElemType() types.Type {
	return inst.elem
}

// NElems returns the number of elements of the alloca instruction and a boolean
// indicating if the number of elements were present.
func (inst *InstAlloca) NElems() (value.Value, bool) {
	if inst.nelems != nil {
		return inst.nelems, true
	}
	return nil, false
}

// SetNElems sets the number of elements of the alloca instruction.
func (inst *InstAlloca) SetNElems(nelems value.Value) {
	inst.nelems = nelems
}

// --- [ load ] ----------------------------------------------------------------

// InstLoad represents a load instruction.
//
// References:
//    http://llvm.org/docs/LangRef.html#load-instruction
type InstLoad struct {
	// Parent basic block.
	parent *BasicBlock
	// Identifier associated with the instruction.
	ident string
	// Source address.
	src value.Value
	// Type of the instruction.
	typ types.Type
}

// NewLoad returns a new load instruction based on the given source address.
func NewLoad(src value.Value) *InstLoad {
	if typ, ok := src.Type().(*types.PointerType); ok {
		return &InstLoad{src: src, typ: typ.Elem()}
	}
	panic(fmt.Sprintf("invalid source address type; expected *types.PointerType, got %T", src.Type()))
}

// Type returns the type of the instruction.
func (inst *InstLoad) Type() types.Type {
	return inst.typ
}

// Ident returns the identifier associated with the instruction.
func (inst *InstLoad) Ident() string {
	return enc.Local(inst.ident)
}

// SetIdent sets the identifier associated with the instruction.
func (inst *InstLoad) SetIdent(ident string) {
	inst.ident = ident
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
func (inst *InstLoad) Parent() *BasicBlock {
	return inst.parent
}

// SetParent sets the parent basic block of the instruction.
func (inst *InstLoad) SetParent(parent *BasicBlock) {
	inst.parent = parent
}

// Src returns the source address of the load instruction.
func (inst *InstLoad) Src() value.Value {
	return inst.src
}

// --- [ store ] ---------------------------------------------------------------

// InstStore represents a store instruction.
//
// References:
//    http://llvm.org/docs/LangRef.html#store-instruction
type InstStore struct {
	// Parent basic block.
	parent *BasicBlock
	// Source value.
	src value.Value
	// Destination address.
	dst value.Value
}

// NewStore returns a new store instruction based on the given source value and
// destination address.
func NewStore(src, dst value.Value) *InstStore {
	return &InstStore{src: src, dst: dst}
}

// String returns the LLVM syntax representation of the instruction.
func (inst *InstStore) String() string {
	src, dst := inst.Src(), inst.Dst()
	return fmt.Sprintf("store %s %s, %s %s",
		src.Type(),
		src.Ident(),
		dst.Type(),
		dst.Ident())
}

// Parent returns the parent basic block of the instruction.
func (inst *InstStore) Parent() *BasicBlock {
	return inst.parent
}

// SetParent sets the parent basic block of the instruction.
func (inst *InstStore) SetParent(parent *BasicBlock) {
	inst.parent = parent
}

// Src returns the source value of the store instruction.
func (inst *InstStore) Src() value.Value {
	return inst.src
}

// Dst returns the destination address of the store instruction.
func (inst *InstStore) Dst() value.Value {
	return inst.dst
}

// --- [ fence ] ---------------------------------------------------------------

// --- [ cmpxchg ] -------------------------------------------------------------

// --- [ atomicrmw ] -----------------------------------------------------------

// --- [ getelementptr ] -------------------------------------------------------

// InstGetElementPtr represents a getelementptr instruction.
//
// References:
//    http://llvm.org/docs/LangRef.html#getelementptr-instruction
type InstGetElementPtr struct {
	// Parent basic block.
	parent *BasicBlock
	// Identifier associated with the instruction.
	ident string
	// Source address.
	src value.Value
	// Element indices.
	indices []value.Value
	// Type of the instruction.
	typ types.Type
	// Source address element type.
	elem types.Type
}

// NewGetElementPtr returns a new getelementptr instruction based on the given
// source address and element indices.
func NewGetElementPtr(src value.Value, indices ...value.Value) *InstGetElementPtr {
	srcType, ok := src.Type().(*types.PointerType)
	if !ok {
		panic(fmt.Sprintf("invalid source address type; expected *types.PointerType, got %T", src.Type()))
	}
	elem := srcType.Elem()
	e := elem
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
			panic("unable to index into element of pointer type; for more information, see http://llvm.org/docs/GetElementPtr.html#what-is-dereferenced-by-gep")
		case *types.ArrayType:
			e = t.Elem()
		case *types.StructType:
			idx, ok := index.(*constant.Int)
			if !ok {
				panic(fmt.Sprintf("invalid index type for structure element; expected *constant.Int, got %T", index))
			}
			e = t.Fields()[idx.Int64()]
		default:
			panic(fmt.Sprintf("support for indexing element type %T not yet implemented", e))
		}
	}
	typ := types.NewPointer(e)
	return &InstGetElementPtr{src: src, indices: indices, typ: typ, elem: elem}
}

// Type returns the type of the instruction.
func (inst *InstGetElementPtr) Type() types.Type {
	return inst.typ
}

// Ident returns the identifier associated with the instruction.
func (inst *InstGetElementPtr) Ident() string {
	return enc.Local(inst.ident)
}

// SetIdent sets the identifier associated with the instruction.
func (inst *InstGetElementPtr) SetIdent(ident string) {
	inst.ident = ident
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
func (inst *InstGetElementPtr) Parent() *BasicBlock {
	return inst.parent
}

// SetParent sets the parent basic block of the instruction.
func (inst *InstGetElementPtr) SetParent(parent *BasicBlock) {
	inst.parent = parent
}

// Src returns the source address of the getelementptr instruction.
func (inst *InstGetElementPtr) Src() value.Value {
	return inst.src
}

// Indices returns the element indices of the getelementptr instruction.
func (inst *InstGetElementPtr) Indices() []value.Value {
	return inst.indices
}
