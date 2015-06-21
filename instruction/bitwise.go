package instruction

import "github.com/llir/llvm/types"

// References:
//    http://llvm.org/docs/LangRef.html#bitwise-binary-operations

// TODO: Add support for the remaining bitwise binary operations:
//    http://llvm.org/docs/LangRef.html#shl-instruction
//    http://llvm.org/docs/LangRef.html#lshr-instruction
//    http://llvm.org/docs/LangRef.html#ashr-instruction
//    http://llvm.org/docs/LangRef.html#and-instruction
//    http://llvm.org/docs/LangRef.html#or-instruction
//    http://llvm.org/docs/LangRef.html#xor-instruction

type Shl struct{}

func (*Shl) Type() types.Type { panic("Shl.Type: not yet implemented") }
func (*Shl) String() string   { panic("Shl.String: not yet implemented") }

type LShr struct{}

func (*LShr) Type() types.Type { panic("LShr.Type: not yet implemented") }
func (*LShr) String() string   { panic("LShr.String: not yet implemented") }

type AShr struct{}

func (*AShr) Type() types.Type { panic("AShr.Type: not yet implemented") }
func (*AShr) String() string   { panic("AShr.String: not yet implemented") }

type And struct{}

func (*And) Type() types.Type { panic("And.Type: not yet implemented") }
func (*And) String() string   { panic("And.String: not yet implemented") }

type Or struct{}

func (*Or) Type() types.Type { panic("Or.Type: not yet implemented") }
func (*Or) String() string   { panic("Or.String: not yet implemented") }

type Xor struct{}

func (*Xor) Type() types.Type { panic("Xor.Type: not yet implemented") }
func (*Xor) String() string   { panic("Xor.String: not yet implemented") }

// isInst ensures that only non-branching instructions can be assigned to the
// Instruction interface.
func (*Shl) isInst()  {}
func (*LShr) isInst() {}
func (*AShr) isInst() {}
func (*And) isInst()  {}
func (*Or) isInst()   {}
func (*Xor) isInst()  {}
