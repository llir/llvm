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
	"github.com/llir/llvm/ir/metadata"
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
	Parent *BasicBlock
	// Name of the local variable associated with the instruction.
	Name string
	// Type of the instruction.
	Typ *types.PointerType
	// Element type.
	Elem types.Type
	// Number of elements; or nil if one element.
	NElems value.Value
	// Map from metadata identifier (e.g. !dbg) to metadata associated with the
	// instruction.
	Metadata map[string]*metadata.Metadata
}

// NewAlloca returns a new alloca instruction based on the given element type.
func NewAlloca(elem types.Type) *InstAlloca {
	typ := types.NewPointer(elem)
	return &InstAlloca{
		Typ:      typ,
		Elem:     elem,
		Metadata: make(map[string]*metadata.Metadata),
	}
}

// Type returns the type of the instruction.
func (inst *InstAlloca) Type() types.Type {
	return inst.Typ
}

// Ident returns the identifier associated with the instruction.
func (inst *InstAlloca) Ident() string {
	return enc.Local(inst.Name)
}

// GetName returns the name of the local variable associated with the
// instruction.
func (inst *InstAlloca) GetName() string {
	return inst.Name
}

// SetName sets the name of the local variable associated with the instruction.
func (inst *InstAlloca) SetName(name string) {
	inst.Name = name
}

// String returns the LLVM syntax representation of the instruction.
func (inst *InstAlloca) String() string {
	md := metadataString(inst.Metadata, ",")
	if inst.NElems != nil {
		return fmt.Sprintf("%s = alloca %s, %s %s%s",
			inst.Ident(),
			inst.Elem,
			inst.NElems.Type(),
			inst.NElems.Ident(),
			md)
	}
	return fmt.Sprintf("%s = alloca %s%s",
		inst.Ident(),
		inst.Elem,
		md)
}

// GetParent returns the parent basic block of the instruction.
func (inst *InstAlloca) GetParent() *BasicBlock {
	return inst.Parent
}

// SetParent sets the parent basic block of the instruction.
func (inst *InstAlloca) SetParent(parent *BasicBlock) {
	inst.Parent = parent
}

// --- [ load ] ----------------------------------------------------------------

// InstLoad represents a load instruction.
//
// References:
//    http://llvm.org/docs/LangRef.html#load-instruction
type InstLoad struct {
	// Parent basic block.
	Parent *BasicBlock
	// Name of the local variable associated with the instruction.
	Name string
	// Type of the instruction.
	Typ types.Type
	// Source address.
	Src value.Value
	// Map from metadata identifier (e.g. !dbg) to metadata associated with the
	// instruction.
	Metadata map[string]*metadata.Metadata
}

// NewLoad returns a new load instruction based on the given source address.
func NewLoad(src value.Value) *InstLoad {
	t, ok := src.Type().(*types.PointerType)
	if !ok {
		panic(fmt.Errorf("invalid source address type; expected *types.PointerType, got %T", src.Type()))
	}
	return &InstLoad{
		Typ:      t.Elem,
		Src:      src,
		Metadata: make(map[string]*metadata.Metadata),
	}
}

// Type returns the type of the instruction.
func (inst *InstLoad) Type() types.Type {
	return inst.Typ
}

// Ident returns the identifier associated with the instruction.
func (inst *InstLoad) Ident() string {
	return enc.Local(inst.Name)
}

// GetName returns the name of the local variable associated with the
// instruction.
func (inst *InstLoad) GetName() string {
	return inst.Name
}

// SetName sets the name of the local variable associated with the instruction.
func (inst *InstLoad) SetName(name string) {
	inst.Name = name
}

// String returns the LLVM syntax representation of the instruction.
func (inst *InstLoad) String() string {
	md := metadataString(inst.Metadata, ",")
	return fmt.Sprintf("%s = load %s, %s %s%s",
		inst.Ident(),
		inst.Type(),
		inst.Src.Type(),
		inst.Src.Ident(),
		md)
}

// GetParent returns the parent basic block of the instruction.
func (inst *InstLoad) GetParent() *BasicBlock {
	return inst.Parent
}

// SetParent sets the parent basic block of the instruction.
func (inst *InstLoad) SetParent(parent *BasicBlock) {
	inst.Parent = parent
}

// --- [ store ] ---------------------------------------------------------------

// InstStore represents a store instruction.
//
// References:
//    http://llvm.org/docs/LangRef.html#store-instruction
type InstStore struct {
	// Parent basic block.
	Parent *BasicBlock
	// Source value.
	Src value.Value
	// Destination address.
	Dst value.Value
	// Map from metadata identifier (e.g. !dbg) to metadata associated with the
	// instruction.
	Metadata map[string]*metadata.Metadata
}

// NewStore returns a new store instruction based on the given source value and
// destination address.
func NewStore(src, dst value.Value) *InstStore {
	return &InstStore{
		Src:      src,
		Dst:      dst,
		Metadata: make(map[string]*metadata.Metadata),
	}
}

// String returns the LLVM syntax representation of the instruction.
func (inst *InstStore) String() string {
	md := metadataString(inst.Metadata, ",")
	return fmt.Sprintf("store %s %s, %s %s%s",
		inst.Src.Type(),
		inst.Src.Ident(),
		inst.Dst.Type(),
		inst.Dst.Ident(),
		md)
}

// GetParent returns the parent basic block of the instruction.
func (inst *InstStore) GetParent() *BasicBlock {
	return inst.Parent
}

// SetParent sets the parent basic block of the instruction.
func (inst *InstStore) SetParent(parent *BasicBlock) {
	inst.Parent = parent
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
	Parent *BasicBlock
	// Name of the local variable associated with the instruction.
	Name string
	// Type of the instruction.
	Typ *types.PointerType
	// Source address element type.
	Elem types.Type
	// Source address.
	Src value.Value
	// Element indices.
	Indices []value.Value
	// Map from metadata identifier (e.g. !dbg) to metadata associated with the
	// instruction.
	Metadata map[string]*metadata.Metadata
}

// NewGetElementPtr returns a new getelementptr instruction based on the given
// source address and element indices.
func NewGetElementPtr(src value.Value, indices ...value.Value) *InstGetElementPtr {
	srcType, ok := src.Type().(*types.PointerType)
	if !ok {
		panic(fmt.Errorf("invalid source address type; expected *types.PointerType, got %T", src.Type()))
	}
	elem := srcType.Elem
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
			e = t.Elem
		case *types.StructType:
			idx, ok := index.(*constant.Int)
			if !ok {
				panic(fmt.Errorf("invalid index type for structure element; expected *constant.Int, got %T", index))
			}
			e = t.Fields[idx.Int64()]
		default:
			panic(fmt.Errorf("support for indexing element type %T not yet implemented", e))
		}
	}
	typ := types.NewPointer(e)
	return &InstGetElementPtr{
		Typ:      typ,
		Elem:     elem,
		Src:      src,
		Indices:  indices,
		Metadata: make(map[string]*metadata.Metadata),
	}
}

// Type returns the type of the instruction.
func (inst *InstGetElementPtr) Type() types.Type {
	return inst.Typ
}

// Ident returns the identifier associated with the instruction.
func (inst *InstGetElementPtr) Ident() string {
	return enc.Local(inst.Name)
}

// GetName returns the name of the local variable associated with the
// instruction.
func (inst *InstGetElementPtr) GetName() string {
	return inst.Name
}

// SetName sets the name of the local variable associated with the instruction.
func (inst *InstGetElementPtr) SetName(name string) {
	inst.Name = name
}

// String returns the LLVM syntax representation of the instruction.
func (inst *InstGetElementPtr) String() string {
	indices := &bytes.Buffer{}
	for _, index := range inst.Indices {
		fmt.Fprintf(indices, ", %s %s",
			index.Type(),
			index.Ident())
	}
	md := metadataString(inst.Metadata, ",")
	return fmt.Sprintf("%s = getelementptr %s, %s %s%s%s",
		inst.Ident(),
		inst.Elem,
		inst.Src.Type(),
		inst.Src.Ident(),
		indices,
		md)
}

// GetParent returns the parent basic block of the instruction.
func (inst *InstGetElementPtr) GetParent() *BasicBlock {
	return inst.Parent
}

// SetParent sets the parent basic block of the instruction.
func (inst *InstGetElementPtr) SetParent(parent *BasicBlock) {
	inst.Parent = parent
}
