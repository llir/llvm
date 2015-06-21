package instruction

import "github.com/llir/llvm/types"

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

type Alloca struct{}

func (*Alloca) Type() types.Type { panic("Alloca.Type: not yet implemented") }
func (*Alloca) String() string   { panic("Alloca.String: not yet implemented") }

type Load struct{}

func (*Load) Type() types.Type { panic("Load.Type: not yet implemented") }
func (*Load) String() string   { panic("Load.String: not yet implemented") }

type Store struct{}

func (*Store) Type() types.Type { panic("Store.Type: not yet implemented") }
func (*Store) String() string   { panic("Store.String: not yet implemented") }

type Fence struct{}

func (*Fence) Type() types.Type { panic("Fence.Type: not yet implemented") }
func (*Fence) String() string   { panic("Fence.String: not yet implemented") }

type CmpXchg struct{}

func (*CmpXchg) Type() types.Type { panic("CmpXchg.Type: not yet implemented") }
func (*CmpXchg) String() string   { panic("CmpXchg.String: not yet implemented") }

type AtomicRMW struct{}

func (*AtomicRMW) Type() types.Type { panic("AtomicRMW.Type: not yet implemented") }
func (*AtomicRMW) String() string   { panic("AtomicRMW.String: not yet implemented") }

type GetElementPtr struct{}

func (*GetElementPtr) Type() types.Type { panic("GetElementPtr.Type: not yet implemented") }
func (*GetElementPtr) String() string   { panic("GetElementPtr.String: not yet implemented") }

// isInst ensures that only non-branching instructions can be assigned to the
// Instruction interface.
func (*Alloca) isInst()        {}
func (*Load) isInst()          {}
func (*Store) isInst()         {}
func (*Fence) isInst()         {}
func (*CmpXchg) isInst()       {}
func (*AtomicRMW) isInst()     {}
func (*GetElementPtr) isInst() {}
