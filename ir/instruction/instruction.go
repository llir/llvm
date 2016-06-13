//go:generate go run gen.go

// Package instruction declares the instructions of LLVM IR.
package instruction

import (
	"fmt"

	"github.com/llir/llvm/ir/types"
)

// An Instruction represents a non-branching instruction.
//
// Instruction may have one of the following underlying types.
//
//    *LocalVarDef
//    *Store
//    *Fence
type Instruction interface {
	fmt.Stringer
	// isInst ensures that only non-branching instructions can be assigned to the
	// Instruction interface.
	isInst()
}

// A ValueInst instruction is a non-branching instruction which returns a value.
//
// ValueInst may have one of the following underlying types, from one of the
// specified categories.
//
// 1) Binary Operations
//
//    *Add
//    *FAdd
//    *Sub
//    *FSub
//    *Mul
//    *FMul
//    *UDiv
//    *SDiv
//    *FDiv
//    *URem
//    *SRem
//    *FRem
//
// 2) Bitwise Binary Operations
//
//    *ShL
//    *LShR
//    *AShR
//    *And
//    *Or
//    *Xor
//
// 3) Vector Operations
//
//    *ExtractElement
//    *InsertElement
//    *ShuffleVector
//
// 4) Aggregate Operations
//
//    *ExtractValue
//    *InsertValue
//
// 5) Memory Access and Addressing Operations
//
//    *Alloca
//    *Load
//    *CmpXchg
//    *AtomicRMW
//    *GetElementPtr
//
// 6) Conversion Operations
//
//    *Trunc
//    *ZExt
//    *SExt
//    *FPTrunc
//    *FPExt
//    *FPToUI
//    *FPToSI
//    *UIToFP
//    *SIToFP
//    *PtrToInt
//    *IntToPtr
//    *BitCast
//    *AddrSpaceCast
//
// 7) Other Operations
//
//    *ICmp
//    *FCmp
//    *PHI
//    *Select
//    *Call
//    *VAArg
//    *LandingPad
type ValueInst interface {
	fmt.Stringer
	// RetType returns the type of the value produced by the instruction.
	RetType() types.Type
}

// A Terminator is a control flow instruction (e.g. br, ret) which terminates a
// basic block.
//
// Terminator may have one of the following underlying types.
//
//    *Ret
//    *Jmp
//    *Br
//    *Switch
//    *IndirectBr
//    *Invoke
//    *Resume
//    *Unreachable
//
// References:
//    http://llvm.org/docs/LangRef.html#terminator-instructions
type Terminator interface {
	fmt.Stringer
	// isTerm ensures that only terminator instructions can be assigned to the
	// Terminator interface.
	isTerm()
}
