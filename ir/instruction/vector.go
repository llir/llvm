package instruction

import "github.com/llir/llvm/ir/types"

// References:
//    http://llvm.org/docs/LangRef.html#vector-operations

// TODO: Add support for the remaining vector operations:
//    http://llvm.org/docs/LangRef.html#extractelement-instruction
//    http://llvm.org/docs/LangRef.html#insertelement-instruction
//    http://llvm.org/docs/LangRef.html#shufflevector-instruction

type ExtractElement struct{}

func (*ExtractElement) Type() types.Type { panic("ExtractElement.Type: not yet implemented") }
func (*ExtractElement) String() string   { panic("ExtractElement.String: not yet implemented") }

type InsertElement struct{}

func (*InsertElement) Type() types.Type { panic("InsertElement.Type: not yet implemented") }
func (*InsertElement) String() string   { panic("InsertElement.String: not yet implemented") }

type ShuffleVector struct{}

func (*ShuffleVector) Type() types.Type { panic("ShuffleVector.Type: not yet implemented") }
func (*ShuffleVector) String() string   { panic("ShuffleVector.String: not yet implemented") }

// isInst ensures that only non-branching instructions can be assigned to the
// Instruction interface.
func (*ExtractElement) isInst() {}
func (*InsertElement) isInst()  {}
func (*ShuffleVector) isInst()  {}
