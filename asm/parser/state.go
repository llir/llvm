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
//                  GetelementptrInst | ToInst | TruncInst | ZextInst |
//                  SextInst | FptruncInst | FpextInst | FptouiInst |
//                  FptosiInst | UitofpInst | SitofpInst | PtrtointInst |
//                  InttoptrInst | BitcastInst | AddrspacecastInst | IcmpInst |
//                  FcmpInst | PhiInst | SelectInst | CallInst | VaArgInst |
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
