// === [ Memory instructions ] =================================================
//
// References:
//    http://llvm.org/docs/LangRef.html#memory-access-and-addressing-operations

package instruction

import (
	"bytes"
	"fmt"

	"github.com/llir/llvm/ir/constant"
	"github.com/llir/llvm/ir/types"
	"github.com/llir/llvm/ir/value"
	"github.com/mewkiz/pkg/errutil"
)

// --- [ alloca ] --------------------------------------------------------------

// Alloca represents an alloca instruction, which is used to allocate memory on
// the stack frame of a function.
//
// References:
//    http://llvm.org/docs/LangRef.html#alloca-instruction
type Alloca struct {
	// Element type.
	elem types.Type
	// Number of elements.
	nelems int
	// Result type.
	typ *types.Pointer
}

// NewAlloca returns a new alloca instruction based on the given element type
// and number of elments.
//
// Pre-condition:
//    1. elem is a valid pointer element type (i.e. any type except void, label
//       and metadata)
func NewAlloca(elem types.Type, nelems int) (*Alloca, error) {
	// Determine result type.
	typ, err := types.NewPointer(elem)
	if err != nil {
		return nil, errutil.Err(err)
	}
	return &Alloca{elem: elem, nelems: nelems, typ: typ}, nil
}

// Elem returns the element type of the memory allocated by the alloca
// instruction.
func (inst *Alloca) Elem() types.Type {
	return inst.elem
}

// NElems returns the number of elements allocated by the alloca instruction.
func (inst *Alloca) NElems() int {
	return inst.nelems
}

// RetType returns the type of the value produced by the instruction.
func (inst *Alloca) RetType() types.Type {
	return inst.typ
}

// String returns the string representation of the instruction.
func (inst *Alloca) String() string {
	if inst.nelems != 1 {
		return fmt.Sprintf("alloca %s, i32 %d", inst.Elem(), inst.NElems())
	}
	return fmt.Sprintf("alloca %s", inst.Elem())
}

// --- [ load ] ----------------------------------------------------------------

// Load represents a load instruction, which is used to read from memory.
//
// References:
//    http://llvm.org/docs/LangRef.html#load-instruction
type Load struct {
	// Source address.
	srcAddr value.Value
	// Result type.
	typ types.Type
}

// NewLoad returns a new load instruction based on the given source address.
//
// Pre-conditions:
//    1. srcAddr is of pointer type
func NewLoad(srcAddr value.Value) (*Load, error) {
	// Validate that srcAddr is of pointer type.
	srcAddrType, ok := srcAddr.Type().(*types.Pointer)
	if !ok {
		return nil, errutil.Newf("invalid source address pointer type; expected *types.Pointer, got %T", srcAddr.Type())
	}
	// Determine result type.
	typ := srcAddrType.Elem()
	return &Load{srcAddr: srcAddr, typ: typ}, nil
}

// SrcAddr returns the source address of the load instruction.
func (inst *Load) SrcAddr() value.Value {
	return inst.srcAddr
}

// RetType returns the type of the value produced by the instruction.
func (inst *Load) RetType() types.Type {
	return inst.typ
}

// String returns the string representation of the instruction.
func (inst *Load) String() string {
	srcAddr := inst.SrcAddr()
	return fmt.Sprintf("load %s, %s %s", inst.RetType(), srcAddr.Type(), srcAddr.ValueString())
}

// --- [ store ] ---------------------------------------------------------------

// Store represents a store instruction, which is used to write to memory.
//
// References:
//    http://llvm.org/docs/LangRef.html#store-instruction
type Store struct {
	// Source value.
	src value.Value
	// Destination address.
	dstAddr value.Value
}

// NewStore returns a new store instruction based on the given source value and
// destination address.
//
// Pre-condition:
//    1. dstAddr is of pointer type
//    2. src is of identical type as the element type of dstAddr
func NewStore(src, dstAddr value.Value) (*Store, error) {
	// Validate that dstAddr is of pointer type.
	dstAddrType, ok := dstAddr.Type().(*types.Pointer)
	if !ok {
		return nil, errutil.Newf("invalid destination address pointer type; expected *types.Pointer, got %T", dstAddr.Type())
	}
	// Validate that src is of identical type as the element type of dstAddr.
	if !types.Equal(src.Type(), dstAddrType.Elem()) {
		return nil, errutil.Newf("type mismatch between source value (%v) and destination address element type (%v)", src.Type(), dstAddrType.Elem())
	}
	return &Store{src: src, dstAddr: dstAddr}, nil
}

// Src returns the source value of the store instruction.
func (inst *Store) Src() value.Value {
	return inst.src
}

// DstAddr returns the destination address of the store instruction.
func (inst *Store) DstAddr() value.Value {
	return inst.dstAddr
}

// String returns the string representation of the instruction.
func (inst *Store) String() string {
	src, dstAddr := inst.Src(), inst.DstAddr()
	return fmt.Sprintf("store %s %s, %s %s", src.Type(), src.ValueString(), dstAddr.Type(), dstAddr.ValueString())
}

// --- [ fence ] ---------------------------------------------------------------

// TODO: Implement Fence.

// Fence represents a fence instruction, which is used to introduce happens-
// before edges between operations.
//
// References:
//    http://llvm.org/docs/LangRef.html#fence-instruction
type Fence struct{}

// String returns the string representation of the instruction.
func (*Fence) String() string { panic("Fence.String: not yet implemented") }

// --- [ cmpxchg ] -------------------------------------------------------------

// TODO: Implement CmpXchg.

// CmpXchg represents a cmpxchg instruction, which is used to atomically modify
// memory by a compare and exchange operation.
//
// References:
//    http://llvm.org/docs/LangRef.html#cmpxchg-instruction
type CmpXchg struct{}

// RetType returns the type of the value produced by the instruction.
func (*CmpXchg) RetType() types.Type { panic("CmpXchg.RetType: not yet implemented") }

// String returns the string representation of the instruction.
func (*CmpXchg) String() string { panic("CmpXchg.String: not yet implemented") }

// --- [ atomicrmw ] -----------------------------------------------------------

// TODO: Implement AtomicRMW.

// AtomicRMW represents an atomicrmw instruction, which is used to atomically
// modify memory by a read, modify, write operation.
//
// References:
//    http://llvm.org/docs/LangRef.html#atomicrmw-instruction
type AtomicRMW struct{}

// RetType returns the type of the value produced by the instruction.
func (*AtomicRMW) RetType() types.Type { panic("AtomicRMW.RetType: not yet implemented") }

// String returns the string representation of the instruction.
func (*AtomicRMW) String() string { panic("AtomicRMW.String: not yet implemented") }

// --- [ getelementptr ] -------------------------------------------------------

// GetElementPtr represents a getelementptr instruction, which is used to get a
// subelement of an aggregate data structure.
//
// References:
//    http://llvm.org/docs/LangRef.html#getelementptr-instruction
type GetElementPtr struct {
	// Source address.
	srcAddr value.Value
	// Element indices.
	indices []value.Value
	// Result type.
	typ *types.Pointer
	// Element type.
	elem types.Type
}

// NewGetElementPtr returns a new getelementptr instruction based on the given
// source address and element indices.
//
// Pre-condition:
//    1. srcAddr is of pointer type
//    2. indices used to index structure fields are integer constants
func NewGetElementPtr(srcAddr value.Value, indices []value.Value) (*GetElementPtr, error) {
	// Validate that srcAddr is of pointer type.
	srcAddrType, ok := srcAddr.Type().(*types.Pointer)
	if !ok {
		return nil, errutil.Newf("invalid source address pointer type; expected *types.Pointer, got %T", srcAddr.Type())
	}
	// Validate that indices used to index structure fields are integer
	// constants.
	e := srcAddrType.Elem()
	for i, index := range indices {
		if i == 0 {
			// Ignore checking the 0th index as it simply follows the pointer of
			// srcAddr.
			//
			// ref: http://llvm.org/docs/GetElementPtr.html#why-is-the-extra-0-index-required
			continue
		}
		switch ee := e.(type) {
		case *types.Pointer:
			// ref: http://llvm.org/docs/GetElementPtr.html#what-is-dereferenced-by-gep
			return nil, errutil.Newf(`unable to index into element of pointer type; for more information, see http://llvm.org/docs/GetElementPtr.html#what-is-dereferenced-by-gep`)
		case *types.Array:
			e = ee.Elem()
		case *types.Struct:
			idx, ok := index.(*constant.Int)
			if !ok {
				return nil, errutil.Newf("invalid index type for structure element; expected *constant.Int, got %T", index)
			}
			e = ee.Fields()[idx.Value().Int64()]
		default:
			panic(fmt.Sprintf("instruction.NewGetElementPtr: support for indexing element type %T not yet implemented", e))
		}
	}
	// Determine result type.
	typ, err := types.NewPointer(e)
	if err != nil {
		return nil, errutil.Err(err)
	}
	// Determine element type.
	elem := srcAddrType.Elem()
	return &GetElementPtr{srcAddr: srcAddr, indices: indices, typ: typ, elem: elem}, nil
}

// SrcAddr returns the source address of the getelementptr instruction.
func (inst *GetElementPtr) SrcAddr() value.Value {
	return inst.srcAddr
}

// Indices returns the element indices of the getelementptr instruction.
func (inst *GetElementPtr) Indices() []value.Value {
	return inst.indices
}

// RetType returns the type of the value produced by the instruction.
func (inst *GetElementPtr) RetType() types.Type {
	return inst.typ
}

// String returns the string representation of the instruction.
func (inst *GetElementPtr) String() string {
	indicesBuf := new(bytes.Buffer)
	for _, index := range inst.Indices() {
		fmt.Fprintf(indicesBuf, ", %s %s", index.Type(), index.ValueString())
	}
	srcAddr := inst.SrcAddr()
	return fmt.Sprintf("getelementptr %s, %s %s%s", inst.elem, srcAddr.Type(), srcAddr.ValueString(), indicesBuf)
}

// isInst ensures that only non-branching instructions can be assigned to the
// Instruction interface.
func (*Store) isInst() {}
func (*Fence) isInst() {}
