package instruction

import "github.com/llir/llvm/ir/types"

// References:
//    http://llvm.org/docs/LangRef.html#vector-operations

// TODO: Add support for the remaining vector operations:
//    http://llvm.org/docs/LangRef.html#extractelement-instruction
//    http://llvm.org/docs/LangRef.html#insertelement-instruction
//    http://llvm.org/docs/LangRef.html#shufflevector-instruction

type ExtractElement struct{}

// RetType returns the type of the value produced by the instruction.
func (*ExtractElement) RetType() types.Type { panic("ExtractElement.RetType: not yet implemented") }
func (*ExtractElement) String() string      { panic("ExtractElement.String: not yet implemented") }

type InsertElement struct{}

// RetType returns the type of the value produced by the instruction.
func (*InsertElement) RetType() types.Type { panic("InsertElement.RetType: not yet implemented") }
func (*InsertElement) String() string      { panic("InsertElement.String: not yet implemented") }

type ShuffleVector struct{}

// RetType returns the type of the value produced by the instruction.
func (*ShuffleVector) RetType() types.Type { panic("ShuffleVector.RetType: not yet implemented") }
func (*ShuffleVector) String() string      { panic("ShuffleVector.String: not yet implemented") }
