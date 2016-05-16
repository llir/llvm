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

// isValueInst ensures that only instructions which return values can be
// assigned to the Value interface.
func (*ExtractValue) isValueInst() {}
func (*InsertValue) isValueInst()  {}
