package instruction

import (
	"bytes"
	"fmt"

	"github.com/llir/llvm/ir/types"
	"github.com/llir/llvm/ir/value"
	"github.com/mewkiz/pkg/errutil"
)

// References:
//    http://llvm.org/docs/LangRef.html#memory-access-and-addressing-operations

// TODO: Add support for the remaining memory access and addressing operations:
//    http://llvm.org/docs/LangRef.html#alloca-instruction
//    http://llvm.org/docs/LangRef.html#load-instruction
//    http://llvm.org/docs/LangRef.html#store-instruction
//    http://llvm.org/docs/LangRef.html#fence-instruction
//    http://llvm.org/docs/LangRef.html#cmpxchg-instruction
//    http://llvm.org/docs/LangRef.html#atomicrmw-instruction
//    http://llvm.org/docs/LangRef.html#getelementptr-instruction

// Alloca represents an alloca instruction.
type Alloca struct {
	// Element type.
	typ types.Type
	// Number of elements.
	nelems int
}

// NewAlloca returns a new alloca instruction based on the given element type
// and number of elments.
func NewAlloca(typ types.Type, nelems int) (*Alloca, error) {
	// TODO: Add sanity check for nelems?
	return &Alloca{typ: typ, nelems: nelems}, nil
}

// Type returns the type of the value produced by the instruction.
func (inst *Alloca) Type() types.Type {
	return inst.typ
}

// String returns the string representation of the instruction.
func (inst *Alloca) String() string {
	if inst.nelems != 1 {
		return fmt.Sprintf("alloca %s, i32 %d", inst.typ, inst.nelems)
	}
	return fmt.Sprintf("alloca %s", inst.typ)
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
		if elem := addrType.Elem(); types.Equal(typ, elem) {
			return nil, errutil.Newf("type mismatch between %v and %v", typ, elem)
		}
	default:
		return nil, errutil.Newf("invalid pointer type; expected *types.Pointer, got %T", addrType)
	}
	return &Load{typ: typ, addr: addr}, nil
}

// Type returns the type of the value produced by the instruction.
func (inst *Load) Type() types.Type {
	return inst.typ
}

// String returns the string representation of the instruction.
func (inst *Load) String() string {
	return fmt.Sprintf("load %s, %s %s", inst.typ, inst.addr.Type(), inst.addr)
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

// Type returns the type of the value produced by the instruction.
func (inst *Store) Type() types.Type {
	return types.NewVoid()
}

// String returns the string representation of the instruction.
func (inst *Store) String() string {
	return fmt.Sprintf("store %s %s, %s %s", inst.val.Type(), inst.val, inst.addr.Type(), inst.addr)
}

type Fence struct{}

func (*Fence) Type() types.Type { panic("Fence.Type: not yet implemented") }
func (*Fence) String() string   { panic("Fence.String: not yet implemented") }

type CmpXchg struct{}

func (*CmpXchg) Type() types.Type { panic("CmpXchg.Type: not yet implemented") }
func (*CmpXchg) String() string   { panic("CmpXchg.String: not yet implemented") }

type AtomicRMW struct{}

func (*AtomicRMW) Type() types.Type { panic("AtomicRMW.Type: not yet implemented") }
func (*AtomicRMW) String() string   { panic("AtomicRMW.String: not yet implemented") }

// GetElementPtr represents a getelementptr instruction.
type GetElementPtr struct {
	// Element type.
	typ types.Type
	// Memory address of the element.
	addr value.Value
	// Element indices.
	indices []value.Value
}

// NewGetElementPtr returns a new getelementptr instruction based on the given
// type, address and element indices.
func NewGetElementPtr(typ types.Type, addr value.Value, indices []value.Value) (*GetElementPtr, error) {
	// Sanity checks.
	switch addrType := addr.Type().(type) {
	case *types.Pointer:
		if elem := addrType.Elem(); types.Equal(typ, elem) {
			return nil, errutil.Newf("type mismatch between %v and %v", typ, elem)
		}
	default:
		return nil, errutil.Newf("invalid pointer type; expected *types.Pointer, got %T", addrType)
	}
	return &GetElementPtr{typ: typ, addr: addr, indices: indices}, nil
}

// Type returns the type of the value produced by the instruction.
func (inst *GetElementPtr) Type() types.Type {
	return inst.typ
}

// String returns the string representation of the instruction.
func (inst *GetElementPtr) String() string {
	if len(inst.indices) > 0 {
		indicesBuf := new(bytes.Buffer)
		for _, index := range inst.indices {
			fmt.Fprintf(indicesBuf, ", %s %s", index.Type(), index)
		}
		return fmt.Sprintf("getelementptr %s, %s %s %s", inst.typ, inst.addr.Type(), inst.addr, indicesBuf)
	}
	return fmt.Sprintf("getelementptr %s, %s %s", inst.typ, inst.addr.Type(), inst.addr)
}

// isValueInst ensures that only instructions which return values can be
// assigned to the Value interface.
func (*Alloca) isValueInst()        {}
func (*Load) isValueInst()          {}
func (*CmpXchg) isValueInst()       {}
func (*AtomicRMW) isValueInst()     {}
func (*GetElementPtr) isValueInst() {}

// isInst ensures that only non-branching instructions can be assigned to the
// Instruction interface.
func (*Store) isInst() {}
func (*Fence) isInst() {}
