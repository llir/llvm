package parser

import "github.com/llir/llvm/instruction"

// =============================================================================
// Terminator Instructions
//
//    ref: http://llvm.org/docs/LangRef.html#terminator-instructions
// =============================================================================

// parseRetInst parses a return instruction. A "ret" token has already been
// comsumed.
//
//    RetInst = "ret" VoidType |
//              "ret" Type Value .
func (p *parser) parseRetInst() (*instruction.Ret, error) {
	panic("parser.parseRetInst: not yet implemented.")
}

// parseBrInst parses a branch instruction. A "br" token has already been
// comsumed.
//
//    BrInst = "br" LabelType Target |
//             "br" "i1" Cond "," LabelType TargetTrue "," LabelType TargetFalse .
//
//    Target      = Local .
//    Cond        = Value .
//    TargetTrue  = Local .
//    TargetFalse = Local .
func (p *parser) parseBrInst() (*instruction.Br, error) {
	panic("parser.parseBrInst: not yet implemented.")
}

// parseSwitchInst parses a switch instruction. A "switch" token has already
// been comsumed.
//
//    SwitchInst = "switch" IntType Value "," LabelType TargetDefault "[" { IntType Value "," LabelType TargetCase } "]" .
//
//    TargetDefault = Local .
//    TargetCase    = Local .
func (p *parser) parseSwitchInst() (*instruction.Switch, error) {
	panic("parser.parseSwitchInst: not yet implemented.")
}

// TODO: Add parsing of IndirectbrInst, InvokeInst, ResumeInst.

// parseUnreachableInst parses an unreachable instruction. An "unreachable"
// token has already been comsumed.
//
//    UnreachableInst = "unreachable" .
func (p *parser) parseUnreachableInst() (*instruction.Unreachable, error) {
	panic("parser.parseUnreachableInst: not yet implemented.")
}

// =============================================================================
// Binary Operations
//
//    ref: http://llvm.org/docs/LangRef.html#binaryops
// =============================================================================

// parseAddInst parses an addition instruction. An "add" token has already been
// comsumed.
//
//    AddInst = Result "=" "add" IntsType Op1 "," Op2 .
//
//    Result = Local
//    Op1    = Value
//    Op2    = Value
func (p *parser) parseAddInst() (*instruction.Add, error) {
	panic("parser.parseAddInst: not yet implemented.")
}

// parseFAddInst parses a floating-point addition instruction. A "fadd" token
// has already been comsumed.
//
//    FAddInst = Result "=" "fadd" FloatsType Op1 "," Op2 .
//
//    Result = Local
//    Op1    = Value
//    Op2    = Value
func (p *parser) parseFAddInst() (*instruction.FAdd, error) {
	panic("parser.parseFAddInst: not yet implemented.")
}

// parseSubInst parses a subtraction instruction. A "sub" token has already been
// comsumed.
//
//    SubInst = Result "=" "sub" IntsType Op1 "," Op2 .
//
//    Result = Local
//    Op1    = Value
//    Op2    = Value
func (p *parser) parseSubInst() (*instruction.Sub, error) {
	panic("parser.parseSubInst: not yet implemented.")
}

// parseFSubInst parses a floating-point subtraction instruction. A "fsub" token
// has already been comsumed.
//
//    FSubInst = Result "=" "fsub" FloatsType Op1 "," Op2 .
//
//    Result = Local
//    Op1    = Value
//    Op2    = Value
func (p *parser) parseFSubInst() (*instruction.FSub, error) {
	panic("parser.parseFSubInst: not yet implemented.")
}

// parseMulInst parses a multiplication instruction. A "mul" token has already
// been comsumed.
//
//    MulInst = Result "=" "mul" IntsType Op1 "," Op2 .
//
//    Result = Local
//    Op1    = Value
//    Op2    = Value
func (p *parser) parseMulInst() (*instruction.Mul, error) {
	panic("parser.parseMulInst: not yet implemented.")
}

// parseFMulInst parses a floating-point multiplication instruction. A "fmul"
// token has already been comsumed.
//
//    FMulInst = Result "=" "fmul" FloatsType Op1 "," Op2 .
//
//    Result = Local
//    Op1    = Value
//    Op2    = Value
func (p *parser) parseFMulInst() (*instruction.FMul, error) {
	panic("parser.parseFMulInst: not yet implemented.")
}

// parseUDivInst parses a unsigned division instruction. An "udiv" token has
// already been comsumed.
//
//    UDivInst = Result "=" "udiv" IntsType Op1 "," Op2 .
//
//    Result = Local
//    Op1    = Value
//    Op2    = Value
func (p *parser) parseUDivInst() (*instruction.UDiv, error) {
	panic("parser.parseUDivInst: not yet implemented.")
}

// parseSDivInst parses a signed division instruction. A "sdiv" token has
// already been comsumed.
//
//    SDivInst = Result "=" "sdiv" IntsType Op1 "," Op2 .
//
//    Result = Local
//    Op1    = Value
//    Op2    = Value
func (p *parser) parseSDivInst() (*instruction.SDiv, error) {
	panic("parser.parseSDivInst: not yet implemented.")
}

// parseFDivInst parses a floating-point division instruction. A "fdiv" token
// has already been comsumed.
//
//    FDivInst = Result "=" "fdiv" FloatsType Op1 "," Op2 .
//
//    Result = Local
//    Op1    = Value
//    Op2    = Value
func (p *parser) parseFDivInst() (*instruction.FDiv, error) {
	panic("parser.parseFDivInst: not yet implemented.")
}

// parseURemInst parses a unsigned modulo instruction. An "urem" token has
// already been comsumed.
//
//    URemInst = Result "=" "urem" IntsType Op1 "," Op2 .
//
//    Result = Local
//    Op1    = Value
//    Op2    = Value
func (p *parser) parseURemInst() (*instruction.URem, error) {
	panic("parser.parseURemInst: not yet implemented.")
}

// parseSRemInst parses a signed modulo instruction. An "srem" token has already
// been comsumed.
//
//    SRemInst = Result "=" "srem" IntsType Op1 "," Op2 .
//
//    Result = Local
//    Op1    = Value
//    Op2    = Value
func (p *parser) parseSRemInst() (*instruction.SRem, error) {
	panic("parser.parseSRemInst: not yet implemented.")
}

// parseFRemInst parses a floating-point modulo instruction. A "frem" token
// has already been comsumed.
//
//    FRemInst = Result "=" "frem" FloatsType Op1 "," Op2 .
//
//    Result = Local
//    Op1    = Value
//    Op2    = Value
func (p *parser) parseFRemInst() (*instruction.FRem, error) {
	panic("parser.parseFRemInst: not yet implemented.")
}

// =============================================================================
// Bitwise Binary Operations
//
//    ref: http://llvm.org/docs/LangRef.html#bitwiseops
// =============================================================================

// parseShlInst parses a shift left instruction. A "shl" token has already been
// comsumed.
//
//    ShlInst = Result "=" "shl" IntsType Op1 "," Op2 .
//
//    Result = Local
//    Op1    = Value
//    Op2    = Value
func (p *parser) parseShlInst() (*instruction.Shl, error) {
	panic("parser.parseShlInst: not yet implemented.")
}

// parseLShrInst parses a logical shift right instruction. A "lshr" token has
// already been comsumed.
//
//    LShrInst = Result "=" "lshr" IntsType Op1 "," Op2 .
//
//    Result = Local
//    Op1    = Value
//    Op2    = Value
func (p *parser) parseLShrInst() (*instruction.LShr, error) {
	panic("parser.parseLShrInst: not yet implemented.")
}

// parseAShrInst parses an arithmetic shift right instruction. An "ashr" token
// has already been comsumed.
//
//    AShrInst = Result "=" "ashr" IntsType Op1 "," Op2 .
//
//    Result = Local
//    Op1    = Value
//    Op2    = Value
func (p *parser) parseAShrInst() (*instruction.AShr, error) {
	panic("parser.parseAShrInst: not yet implemented.")
}

// parseAndInst parses a bitwise logical AND instruction. An "and" token has
// already been comsumed.
//
//    AndInst = Result "=" "and" IntsType Op1 "," Op2 .
//
//    Result = Local
//    Op1    = Value
//    Op2    = Value
func (p *parser) parseAndInst() (*instruction.And, error) {
	panic("parser.parseAndInst: not yet implemented.")
}

// parseOrInst parses a bitwise logical OR instruction. A "or" token has already
// been comsumed.
//
//    OrInst = Result "=" "or" IntsType Op1 "," Op2 .
//
//    Result = Local
//    Op1    = Value
//    Op2    = Value
func (p *parser) parseOrInst() (*instruction.Or, error) {
	panic("parser.parseOrInst: not yet implemented.")
}

// parseXorInst parses a bitwise logical XOR instruction. A "xor" token has
// already been comsumed.
//
//    XorInst = Result "=" "xor" IntsType Op1 "," Op2 .
//
//    Result = Local
//    Op1    = Value
//    Op2    = Value
func (p *parser) parseXorInst() (*instruction.Xor, error) {
	panic("parser.parseXorInst: not yet implemented.")
}

// =============================================================================
// Vector Operations
//
//    ref: http://llvm.org/docs/LangRef.html#vector-operations
// =============================================================================

// TODO: Add parsing of ExtractelementInst, InsertelementInst,
// ShufflevectorInst.

// =============================================================================
// Aggregate Operations
//
//    ref: http://llvm.org/docs/LangRef.html#aggregate-operations
// =============================================================================

// TODO: Add parsing of ExtractvalueInst, InsertvalueInst.

// =============================================================================
// Memory Access and Addressing Operations
//
//    ref: http://llvm.org/docs/LangRef.html#memoryops
// =============================================================================

// parseAllocaInst parses a stack allocation instruction. An "alloca" token has
// already been comsumed.
//
//    AllocaInst = Result "=" "alloca" Type [ "," IntType NumElems ] [ "," "align" Align ] .
//
//    Result   = Local
//    NumElems = Value
//    Align    = int_lit
func (p *parser) parseAllocaInst() (*instruction.Alloca, error) {
	panic("parser.parseAllocaInst: not yet implemented.")
}

// parseLoadInst parses a memory load instruction. A "load" token has already
// been comsumed.
//
//    LoadInst = Result "=" "load" Type "*" Addr [ "," "align" Align ] .
//
//    Result = Local
//    Addr   = Global | Local
//    Align  = int_lit
func (p *parser) parseLoadInst() (*instruction.Load, error) {
	panic("parser.parseLoadInst: not yet implemented.")
}

// parseStoreInst parses a memory store instruction. A "store" token has already
// been comsumed.
//
//    StoreInst = "store" Type Value "," Type "*" Addr [ "," "align" Align ] .
//
//    Addr   = Global | Local
//    Align  = int_lit
func (p *parser) parseStoreInst() (*instruction.Store, error) {
	panic("parser.parseStoreInst: not yet implemented.")
}

// TODO: Add parsing of FenceInst, CmpxchgInst, AtomicrmwInst.

// parseGetElementPtrInst parses a memory address calculation instruction. A
// "getelementptr" token has already been comsumed.
//
//    GetElementPtrInst = Result "=" "getelementptr" Type "*" Addr { "," IntType Idx } .
//
//    Result = Local
//    Addr   = Global | Local
//    Idx    = Value
func (p *parser) parseGetElementPtrInst() (*instruction.GetElementPtr, error) {
	panic("parser.parseGetElementPtrInst: not yet implemented.")
}

// =============================================================================
// Conversion Operations
//
//    ref: http://llvm.org/docs/LangRef.html#conversion-operations
// =============================================================================

// TODO: Add parsing of TruncInst, ZextInst, SextInst, FptruncInst, FpextInst,
// FptouiInst, FptosiInst, UitofpInst, SitofpInst, PtrtointInst, InttoptrInst,
// BitcastInst, AddrspacecastInst.

// =============================================================================
// Other Operations
//
//    ref: http://llvm.org/docs/LangRef.html#other-operations
// =============================================================================

// TODO: Add parsing of IcmpInst, FcmpInst, PhiInst, SelectInst, CallInst,
// VaArgInst, LandingpadInst.
