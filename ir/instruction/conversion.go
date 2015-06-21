package instruction

import "github.com/llir/llvm/ir/types"

// References:
//    http://llvm.org/docs/LangRef.html#conversion-operations

// TODO: Add support for the remaining conversion operations:
//    http://llvm.org/docs/LangRef.html#trunc-to-instruction
//    http://llvm.org/docs/LangRef.html#zext-to-instruction
//    http://llvm.org/docs/LangRef.html#sext-to-instruction
//    http://llvm.org/docs/LangRef.html#fptrunc-to-instruction
//    http://llvm.org/docs/LangRef.html#fpext-to-instruction
//    http://llvm.org/docs/LangRef.html#fptoui-to-instruction
//    http://llvm.org/docs/LangRef.html#fptosi-to-instruction
//    http://llvm.org/docs/LangRef.html#uitofp-to-instruction
//    http://llvm.org/docs/LangRef.html#sitofp-to-instruction
//    http://llvm.org/docs/LangRef.html#ptrtoint-to-instruction
//    http://llvm.org/docs/LangRef.html#inttoptr-to-instruction
//    http://llvm.org/docs/LangRef.html#bitcast-to-instruction
//    http://llvm.org/docs/LangRef.html#addrspacecast-to-instruction

type Trunc struct{}

func (*Trunc) Type() types.Type { panic("Trunc.Type: not yet implemented") }
func (*Trunc) String() string   { panic("Trunc.String: not yet implemented") }

type ZExt struct{}

func (*ZExt) Type() types.Type { panic("ZExt.Type: not yet implemented") }
func (*ZExt) String() string   { panic("ZExt.String: not yet implemented") }

type SExt struct{}

func (*SExt) Type() types.Type { panic("SExt.Type: not yet implemented") }
func (*SExt) String() string   { panic("SExt.String: not yet implemented") }

type FPTrunc struct{}

func (*FPTrunc) Type() types.Type { panic("FPTrunc.Type: not yet implemented") }
func (*FPTrunc) String() string   { panic("FPTrunc.String: not yet implemented") }

type FPExt struct{}

func (*FPExt) Type() types.Type { panic("FPExt.Type: not yet implemented") }
func (*FPExt) String() string   { panic("FPExt.String: not yet implemented") }

type FPToUI struct{}

func (*FPToUI) Type() types.Type { panic("FPToUI.Type: not yet implemented") }
func (*FPToUI) String() string   { panic("FPToUI.String: not yet implemented") }

type FPToSI struct{}

func (*FPToSI) Type() types.Type { panic("FPToSI.Type: not yet implemented") }
func (*FPToSI) String() string   { panic("FPToSI.String: not yet implemented") }

type UIToFP struct{}

func (*UIToFP) Type() types.Type { panic("UIToFP.Type: not yet implemented") }
func (*UIToFP) String() string   { panic("UIToFP.String: not yet implemented") }

type SIToFP struct{}

func (*SIToFP) Type() types.Type { panic("SIToFP.Type: not yet implemented") }
func (*SIToFP) String() string   { panic("SIToFP.String: not yet implemented") }

type PtrToInt struct{}

func (*PtrToInt) Type() types.Type { panic("PtrToInt.Type: not yet implemented") }
func (*PtrToInt) String() string   { panic("PtrToInt.String: not yet implemented") }

type IntToPtr struct{}

func (*IntToPtr) Type() types.Type { panic("IntToPtr.Type: not yet implemented") }
func (*IntToPtr) String() string   { panic("IntToPtr.String: not yet implemented") }

type BitCast struct{}

func (*BitCast) Type() types.Type { panic("BitCast.Type: not yet implemented") }
func (*BitCast) String() string   { panic("BitCast.String: not yet implemented") }

type AddrSpaceCast struct{}

func (*AddrSpaceCast) Type() types.Type { panic("AddrSpaceCast.Type: not yet implemented") }
func (*AddrSpaceCast) String() string   { panic("AddrSpaceCast.String: not yet implemented") }

// isInst ensures that only non-branching instructions can be assigned to the
// Instruction interface.
func (*Trunc) isInst()         {}
func (*ZExt) isInst()          {}
func (*SExt) isInst()          {}
func (*FPTrunc) isInst()       {}
func (*FPExt) isInst()         {}
func (*FPToUI) isInst()        {}
func (*FPToSI) isInst()        {}
func (*UIToFP) isInst()        {}
func (*SIToFP) isInst()        {}
func (*PtrToInt) isInst()      {}
func (*IntToPtr) isInst()      {}
func (*BitCast) isInst()       {}
func (*AddrSpaceCast) isInst() {}
