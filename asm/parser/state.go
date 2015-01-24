package parser

// decimal_digit = "0" … "9" .

import (
	"errors"
	"fmt"
	"io"

	"github.com/mewlang/llvm/asm/token"
	"github.com/mewlang/llvm/ir"
)

// TODO: Complete TopLevelEntity EBNF definition.

// parseTopLevelEntity parses a top-level entity and stores it in module.
//
//    TopLevelEntity = FuncDecl | FuncDef .
func (p *parser) parseTopLevelEntity(module *ir.Module) error {
	tok := p.next()
	switch tok.Kind {
	case token.Error:
		return errors.New(tok.Val)
	case token.EOF:
		// Terminate the parser at EOF.
		return io.EOF
	case token.KwDeclare:
		f, err := p.parseDeclare()
		if err != nil {
			return err
		}
		module.Funcs = append(module.Funcs, f)
	case token.KwDefine:
		f, err := p.parseDefine()
		if err != nil {
			return err
		}
		module.Funcs = append(module.Funcs, f)
	}
	return fmt.Errorf("invalid token type %v; expected top-level entity", tok.Kind)
}

// parseDeclare parses a function declaration. A "declare" token has already
// been consumed.
//
//    FuncDecl = "declare" FuncHeader .
func (p *parser) parseDeclare() (*ir.Function, error) {
	panic("not yet implemented.")
}

// parseDefine parses a function definition. A "define" token has already been
// consumed.
//
//    FuncDef = "define" FuncHeader FuncBody .
func (p *parser) parseDefine() (*ir.Function, error) {
	panic("not yet implemented.")
}

// parseFuncHeader parses a function header consisting of a return argument, a
// function name and zero or more function arguments.
//
//    FuncHeader = RetType FuncName "(" ArgList ")" .
//
//    RetType  = Type .
//    FuncName = GlobalName .
//    ArgList  = [ Arg { "," Arg } ] .
//    Arg      = Type [ LocalName ] .
func (p *parser) parseFuncHeader() (header *ir.Function, err error) {
	panic("not yet implemented.")
}

// parseFuncBody parses a function body consisting of one or more basic blocks.
//
//    FuncBody = "{" BasicBlock { BasicBlock } "}" .
//
//    BasicBlock  = [ LabelDecl ] { Instruction } Terminator .
//    LabelDecl   = Label ":" .
//    Instruction = AddInst | FaddInst | SubInst | FsubInst | MulInst |
//                  FmulInst | UdivInst | SdivInst | FdivInst | UremInst |
//                  SremInst | FremInst | ShlInst | LshrInst | AshrInst |
//                  AndInst | OrInst | XorInst | ExtractelementInst |
//                  InsertelementInst | ShufflevectorInst | ExtractvalueInst |
//                  InsertvalueInst | AllocaInst | LoadInst | StoreInst |
//                  FenceInst | CmpxchgInst | AtomicrmwInst |
//                  GetelementptrInst | TruncInst | ZextInst | SextInst |
//                  FptruncInst | FpextInst | FptouiInst | FptosiInst |
//                  UitofpInst | SitofpInst | PtrtointInst | InttoptrInst |
//                  BitcastInst | AddrspacecastInst | IcmpInst | FcmpInst |
//                  PhiInst | SelectInst | CallInst | VaArgInst |
//                  LandingpadInst .
//    Terminator  = RetInst | BrInst | SwitchInst | IndirectbrInst |
//                  InvokeInst | ResumeInst | UnreachableInst .

func (p *parser) parseFuncBody() (body []*ir.BasicBlock, err error) {
	panic("not yet implemented.")
}

// parseType parses a type.
//
//    Type = IntType | VoidType | "half" | "float" | "double" | "x86_fp80" |
//           "fp128" | "ppc_fp128" | "x86_mmx" | LabelType | "metadata" .
//
//    IntType   = "i" decimal_digit { decimal_digit } .
//    VoidType  = "void" .
//    LabelType = "label" .
func (p *parser) parseType() (ir.Type, error) {
	panic("not yet implemented.")
}

// TODO: Complete Value EBNF definition.

// parseValue parses a value.
//
//    Value = TODO .
func (p *parser) parseValue() (ir.Value, error) {
	panic("not yet implemented.")
}

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
//    Target      = LocalName .
//    Cond        = Value .
//    TargetTrue  = LocalName .
//    TargetFalse = LocalName .
func (p *parser) parseBrInst() (*ir.BranchInst, error) {
	panic("not yet implemented.")
}

// parseSwitchInst parses a switch instruction. A "switch" token has already
// been comsumed.
//
//    SwitchInst = "switch" IntType Value "," LabelType TargetDefault "[" { IntType Value "," LabelType TargetCase } "]" .
//
//    TargetDefault = LocalName .
//    TargetCase    = LocalName .
func (p *parser) parseSwitchInst() (*ir.SwitchInst, error) {
	panic("not yet implemented.")
}

// TODO: Add parsing of IndirectbrInst, InvokeInst, ResumeInst, UnreachableInst.

// =============================================================================
// Binary Operations
//
//    ref: http://llvm.org/docs/LangRef.html#binaryops
// =============================================================================

// TODO: Add parsing of AddInst, FaddInst, SubInst, FsubInst, MulInst, FmulInst,
// UdivInst, SdivInst, FdivInst, UremInst, SremInst, FremInst.

// =============================================================================
// Bitwise Binary Operations
//
//    ref: http://llvm.org/docs/LangRef.html#bitwiseops
// =============================================================================

// TODO: Add parsing of ShlInst, LshrInst, AshrInst, AndInst, OrInst, XorInst.

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

// TODO: Add parsing of AllocaInst, LoadInst, StoreInst, FenceInst, CmpxchgInst,
// AtomicrmwInst, GetelementptrInst.

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
