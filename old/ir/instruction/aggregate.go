package instruction

import "github.com/llir/llvm/ir/types"

// References:
//    http://llvm.org/docs/LangRef.html#aggregate-operations

// TODO: Add support for the remaining aggregate operations:
//    http://llvm.org/docs/LangRef.html#extractvalue-instruction
//    http://llvm.org/docs/LangRef.html#insertvalue-instruction

type ExtractValue struct{}

// RetType returns the type of the value produced by the instruction.
func (*ExtractValue) RetType() types.Type { panic("ExtractValue.RetType: not yet implemented") }
func (*ExtractValue) String() string      { panic("ExtractValue.String: not yet implemented") }

type InsertValue struct{}

// RetType returns the type of the value produced by the instruction.
func (*InsertValue) RetType() types.Type { panic("InsertValue.RetType: not yet implemented") }
func (*InsertValue) String() string      { panic("InsertValue.String: not yet implemented") }
