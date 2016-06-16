// References:
//    http://llvm.org/docs/LangRef.html#memory-access-and-addressing-operations

package instruction

import (
	"bytes"
	"fmt"

	"github.com/llir/llvm/ir/types"
	"github.com/llir/llvm/ir/value"
	"github.com/mewkiz/pkg/errutil"
)

// TODO: Add support for the remaining memory access and addressing operations:
//    http://llvm.org/docs/LangRef.html#fence-instruction
//    http://llvm.org/docs/LangRef.html#cmpxchg-instruction
//    http://llvm.org/docs/LangRef.html#atomicrmw-instruction

// Alloca represents an alloca instruction.
type Alloca struct {
	// Result type.
	typ *types.Pointer
	// Element type.
	elem types.Type
	// Number of elements.
	nelems int
}

// NewAlloca returns a new alloca instruction based on the given element type
// and number of elments.
func NewAlloca(elem types.Type, nelems int) (*Alloca, error) {
	// TODO: Add sanity check for nelems?
	typ, err := types.NewPointer(elem)
	if err != nil {
		return nil, errutil.Err(err)
	}
	return &Alloca{typ: typ, elem: elem, nelems: nelems}, nil
}

// RetType returns the type of the value produced by the instruction.
func (inst *Alloca) RetType() types.Type {
	return inst.typ
}

// Elem returns the element type of allocated value.
func (inst *Alloca) Elem() types.Type {
	return inst.elem
}

// NElems returns the number of elements allocated.
func (inst *Alloca) NElems() int {
	return inst.nelems
}

// String returns the string representation of the instruction.
func (inst *Alloca) String() string {
	if inst.nelems != 1 {
		return fmt.Sprintf("alloca %s, i32 %d", inst.Elem(), inst.NElems())
	}
	return fmt.Sprintf("alloca %s", inst.Elem())
}

// Load represents a load instruction.
type Load struct {
	// The type of the value to load from memory.
	typ types.Type
	// Memory address from which to load.
	addr value.Value
}

// NewLoad returns a new load instruction based on the given type and address.
func NewLoad(typ types.Type, addr value.Value) (*Load, error) {
	// Sanity checks.
	switch addrType := addr.Type().(type) {
	case *types.Pointer:
		if elem := addrType.Elem(); !types.Equal(typ, elem) {
			return nil, errutil.Newf("type mismatch between %v and %v", typ, elem)
		}
	default:
		return nil, errutil.Newf("invalid pointer type; expected *types.Pointer, got %T", addrType)
	}
	return &Load{typ: typ, addr: addr}, nil
}

// RetType returns the type of the value produced by the instruction.
func (inst *Load) RetType() types.Type {
	return inst.typ
}

// String returns the string representation of the instruction.
func (inst *Load) String() string {
	return fmt.Sprintf("load %s, %s %s", inst.typ, inst.addr.Type(), inst.addr.ValueString())
}

// Store represents a store instruction.
type Store struct {
	// The value to store in memory.
	val value.Value
	// Memory address at which to store.
	addr value.Value
}

// NewStore returns a new store instruction based on the given value and
// address.
func NewStore(val, addr value.Value) (*Store, error) {
	// Sanity checks.
	switch addrType := addr.Type().(type) {
	case *types.Pointer:
		if elem := addrType.Elem(); !types.Equal(val.Type(), elem) {
			return nil, errutil.Newf("type mismatch between %v and %v", val.Type(), elem)
		}
	default:
		return nil, errutil.Newf("invalid pointer type; expected *types.Pointer, got %T", addrType)
	}
	return &Store{val: val, addr: addr}, nil
}

// String returns the string representation of the instruction.
func (inst *Store) String() string {
	return fmt.Sprintf("store %s %s, %s %s", inst.val.Type(), inst.val.ValueString(), inst.addr.Type(), inst.addr.ValueString())
}

type Fence struct{}

func (*Fence) String() string { panic("Fence.String: not yet implemented") }

type CmpXchg struct{}

// RetType returns the type of the value produced by the instruction.
func (*CmpXchg) RetType() types.Type { panic("CmpXchg.RetType: not yet implemented") }
func (*CmpXchg) String() string      { panic("CmpXchg.String: not yet implemented") }

type AtomicRMW struct{}

// RetType returns the type of the value produced by the instruction.
func (*AtomicRMW) RetType() types.Type { panic("AtomicRMW.RetType: not yet implemented") }
func (*AtomicRMW) String() string      { panic("AtomicRMW.String: not yet implemented") }

// GetElementPtr represents a getelementptr instruction.
type GetElementPtr struct {
	// Result type.
	typ *types.Pointer
	// Element type.
	elem types.Type
	// Memory address of the element.
	addr value.Value
	// Element indices.
	indices []value.Value
}

// NewGetElementPtr returns a new getelementptr instruction based on the given
// element type, address and element indices.
func NewGetElementPtr(elem types.Type, addr value.Value, indices []value.Value) (*GetElementPtr, error) {
	// Sanity checks.
	switch addrType := addr.Type().(type) {
	case *types.Pointer:
		if !types.Equal(elem, addrType.Elem()) {
			return nil, errutil.Newf("type mismatch between %v and %v", elem, addrType.Elem())
		}
	default:
		return nil, errutil.Newf("invalid pointer type; expected *types.Pointer, got %T", addrType)
	}
	var typ *types.Pointer
	switch elem := elem.(type) {
	case *types.Array:
		var err error
		typ, err = types.NewPointer(elem.Elem())
		if err != nil {
			return nil, errutil.Err(err)
		}
	case *types.Pointer:
		var err error
		typ, err = types.NewPointer(elem.Elem())
		if err != nil {
			return nil, errutil.Err(err)
		}
	default:
		panic(fmt.Sprintf("instruction.NewGetElementPtr: support for type %T not yet implemented", elem))
	}
	return &GetElementPtr{typ: typ, elem: elem, addr: addr, indices: indices}, nil
}

// RetType returns the type of the value produced by the instruction.
func (inst *GetElementPtr) RetType() types.Type {
	return inst.typ
}

// String returns the string representation of the instruction.
func (inst *GetElementPtr) String() string {
	indicesBuf := new(bytes.Buffer)
	for _, index := range inst.indices {
		fmt.Fprintf(indicesBuf, ", %s %s", index.Type(), index.ValueString())
	}
	return fmt.Sprintf("getelementptr %s, %s %s%s", inst.elem, inst.addr.Type(), inst.addr.ValueString(), indicesBuf)
}

// isInst ensures that only non-branching instructions can be assigned to the
// Instruction interface.
func (*Store) isInst() {}
func (*Fence) isInst() {}
