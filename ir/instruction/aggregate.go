package instruction

import "github.com/llir/llvm/ir/types"

// References:
//    http://llvm.org/docs/LangRef.html#aggregate-operations

// TODO: Add support for the remaining aggregate operations:
//    http://llvm.org/docs/LangRef.html#extractvalue-instruction
//    http://llvm.org/docs/LangRef.html#insertvalue-instruction

type ExtractValue struct{}

func (*ExtractValue) Type() types.Type { panic("ExtractValue.Type: not yet implemented") }
func (*ExtractValue) String() string   { panic("ExtractValue.String: not yet implemented") }

type InsertValue struct{}

func (*InsertValue) Type() types.Type { panic("InsertValue.Type: not yet implemented") }
func (*InsertValue) String() string   { panic("InsertValue.String: not yet implemented") }

// isInst ensures that only non-branching instructions can be assigned to the
// Instruction interface.
func (*ExtractValue) isInst() {}
func (*InsertValue) isInst()  {}
