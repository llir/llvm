package parser

import "github.com/llir/llvm/ir"

// =============================================================================
// Terminator Instructions
//
//    ref: http://llvm.org/docs/LangRef.html#terminators
// =============================================================================

// parseRetInst parses a return instruction. A "ret" token has already been
// comsumed.
//
//    RetInst = "ret" VoidType |
//              "ret" Type Value .
func (p *parser) parseRetInst() (*ir.ReturnInst, error) {
	panic("not yet implemented.")
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
func (p *parser) parseBrInst() (*ir.BranchInst, error) {
	panic("not yet implemented.")
}

// parseSwitchInst parses a switch instruction. A "switch" token has already
// been comsumed.
//
//    SwitchInst = "switch" IntType Value "," LabelType TargetDefault "[" { IntType Value "," LabelType TargetCase } "]" .
//
//    TargetDefault = Local .
//    TargetCase    = Local .
func (p *parser) parseSwitchInst() (*ir.SwitchInst, error) {
	panic("not yet implemented.")
}

// TODO: Add parsing of IndirectbrInst, InvokeInst, ResumeInst.

// parseUnreachableInst parses an unreachable instruction. An "unreachable"
// token has already been comsumed.
//
//    UnreachableInst = "unreachable" .
func (p *parser) parseUnreachableInst() (*ir.UnreachableInst, error) {
	panic("not yet implemented.")
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
func (p *parser) parseAddInst() (*ir.AddInst, error) {
	panic("not yet implemented.")
}

// parseFaddInst parses a floating-point addition instruction. A "fadd" token
// has already been comsumed.
//
//    FaddInst = Result "=" "fadd" FloatsType Op1 "," Op2 .
//
//    Result = Local
//    Op1    = Value
//    Op2    = Value
func (p *parser) parseFaddInst() (*ir.FaddInst, error) {
	panic("not yet implemented.")
}

// parseSubInst parses a subtraction instruction. A "sub" token has already been
// comsumed.
//
//    SubInst = Result "=" "sub" IntsType Op1 "," Op2 .
//
//    Result = Local
//    Op1    = Value
//    Op2    = Value
func (p *parser) parseSubInst() (*ir.SubInst, error) {
	panic("not yet implemented.")
}

// parseFsubInst parses a floating-point subtraction instruction. A "fsub" token
// has already been comsumed.
//
//    FsubInst = Result "=" "fsub" FloatsType Op1 "," Op2 .
//
//    Result = Local
//    Op1    = Value
//    Op2    = Value
func (p *parser) parseFsubInst() (*ir.FsubInst, error) {
	panic("not yet implemented.")
}

// parseMulInst parses a multiplication instruction. A "mul" token has already
// been comsumed.
//
//    MulInst = Result "=" "mul" IntsType Op1 "," Op2 .
//
//    Result = Local
//    Op1    = Value
//    Op2    = Value
func (p *parser) parseMulInst() (*ir.MulInst, error) {
	panic("not yet implemented.")
}

// parseFmulInst parses a floating-point multiplication instruction. A "fmul"
// token has already been comsumed.
//
//    FmulInst = Result "=" "fmul" FloatsType Op1 "," Op2 .
//
//    Result = Local
//    Op1    = Value
//    Op2    = Value
func (p *parser) parseFmulInst() (*ir.FmulInst, error) {
	panic("not yet implemented.")
}

// parseUdivInst parses a unsigned division instruction. An "udiv" token has
// already been comsumed.
//
//    UdivInst = Result "=" "udiv" IntsType Op1 "," Op2 .
//
//    Result = Local
//    Op1    = Value
//    Op2    = Value
func (p *parser) parseUdivInst() (*ir.UdivInst, error) {
	panic("not yet implemented.")
}

// parseSdivInst parses a signed division instruction. A "sdiv" token has
// already been comsumed.
//
//    SdivInst = Result "=" "sdiv" IntsType Op1 "," Op2 .
//
//    Result = Local
//    Op1    = Value
//    Op2    = Value
func (p *parser) parseSdivInst() (*ir.SdivInst, error) {
	panic("not yet implemented.")
}

// parseFdivInst parses a floating-point division instruction. A "fdiv" token
// has already been comsumed.
//
//    FdivInst = Result "=" "fdiv" FloatsType Op1 "," Op2 .
//
//    Result = Local
//    Op1    = Value
//    Op2    = Value
func (p *parser) parseFdivInst() (*ir.FdivInst, error) {
	panic("not yet implemented.")
}

// parseUremInst parses a unsigned modulo instruction. An "urem" token has
// already been comsumed.
//
//    UremInst = Result "=" "urem" IntsType Op1 "," Op2 .
//
//    Result = Local
//    Op1    = Value
//    Op2    = Value
func (p *parser) parseUremInst() (*ir.UremInst, error) {
	panic("not yet implemented.")
}

// parseSremInst parses a signed modulo instruction. An "srem" token has already
// been comsumed.
//
//    SremInst = Result "=" "srem" IntsType Op1 "," Op2 .
//
//    Result = Local
//    Op1    = Value
//    Op2    = Value
func (p *parser) parseSremInst() (*ir.SremInst, error) {
	panic("not yet implemented.")
}

// parseFremInst parses a floating-point modulo instruction. A "frem" token
// has already been comsumed.
//
//    FremInst = Result "=" "frem" FloatsType Op1 "," Op2 .
//
//    Result = Local
//    Op1    = Value
//    Op2    = Value
func (p *parser) parseFremInst() (*ir.FremInst, error) {
	panic("not yet implemented.")
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
func (p *parser) parseShlInst() (*ir.ShlInst, error) {
	panic("not yet implemented.")
}

// parseLshrInst parses a logical shift right instruction. A "lshr" token has
// already been comsumed.
//
//    LshrInst = Result "=" "lshr" IntsType Op1 "," Op2 .
//
//    Result = Local
//    Op1    = Value
//    Op2    = Value
func (p *parser) parseLshrInst() (*ir.LshrInst, error) {
	panic("not yet implemented.")
}

// parseAshrInst parses an arithmetic shift right instruction. An "ashr" token
// has already been comsumed.
//
//    AshrInst = Result "=" "ashr" IntsType Op1 "," Op2 .
//
//    Result = Local
//    Op1    = Value
//    Op2    = Value
func (p *parser) parseAshrInst() (*ir.AshrInst, error) {
	panic("not yet implemented.")
}

// parseAndInst parses a bitwise logical AND instruction. An "and" token has
// already been comsumed.
//
//    AndInst = Result "=" "and" IntsType Op1 "," Op2 .
//
//    Result = Local
//    Op1    = Value
//    Op2    = Value
func (p *parser) parseAndInst() (*ir.AndInst, error) {
	panic("not yet implemented.")
}

// parseOrInst parses a bitwise logical OR instruction. A "or" token has already
// been comsumed.
//
//    OrInst = Result "=" "or" IntsType Op1 "," Op2 .
//
//    Result = Local
//    Op1    = Value
//    Op2    = Value
func (p *parser) parseOrInst() (*ir.OrInst, error) {
	panic("not yet implemented.")
}

// parseXorInst parses a bitwise logical XOR instruction. A "xor" token has
// already been comsumed.
//
//    XorInst = Result "=" "xor" IntsType Op1 "," Op2 .
//
//    Result = Local
//    Op1    = Value
//    Op2    = Value
func (p *parser) parseXorInst() (*ir.XorInst, error) {
	panic("not yet implemented.")
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
func (p *parser) parseAllocaInst() (*ir.AllocaInst, error) {
	panic("not yet implemented.")
}

// parseLoadInst parses a memory load instruction. A "load" token has already
// been comsumed.
//
//    LoadInst = Result "=" "load" Type "*" Addr [ "," "align" Align ] .
//
//    Result = Local
//    Addr   = Global | Local
//    Align  = int_lit
func (p *parser) parseLoadInst() (*ir.LoadInst, error) {
	panic("not yet implemented.")
}

// parseStoreInst parses a memory store instruction. A "store" token has already
// been comsumed.
//
//    StoreInst = "store" Type Value "," Type "*" Addr [ "," "align" Align ] .
//
//    Addr   = Global | Local
//    Align  = int_lit
func (p *parser) parseStoreInst() (*ir.StoreInst, error) {
	panic("not yet implemented.")
}

// TODO: Add parsing of FenceInst, CmpxchgInst, AtomicrmwInst.

// parseGetelementptrInst parses a memory address calculation instruction. A
// "getelementptr" token has already been comsumed.
//
//    GetelementptrInst = Result "=" "getelementptr" Type "*" Addr { "," IntType Idx } .
//
//    Result = Local
//    Addr   = Global | Local
//    Idx    = Value
func (p *parser) parseGetelementptrInst() (*ir.GetelementptrInst, error) {
	panic("not yet implemented.")
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
