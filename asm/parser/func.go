package parser

import (
	"errors"

	"github.com/llir/llvm/asm/token"
	"github.com/llir/llvm/ir"
	"github.com/mewkiz/pkg/errutil"
)

// parseFuncDecl parses an external function declaration. A "declare" token has
// already been consumed.
//
// Syntax:
//    FuncDecl = "declare" FuncHeader .
//
// Examples:
//    declare i32 @printf(i8*, ...)
//
// References:
//    http://llvm.org/docs/LangRef.html#functions
func (p *parser) parseFuncDecl() error {
	f, err := p.parseFuncHeader()
	if err != nil {
		return errutil.Err(err)
	}
	p.m.Funcs = append(p.m.Funcs, f)
	return nil
}

// parseFuncDef parses a function definition. A "define" token has already been
// consumed.
//
//    FuncDef = "define" FuncHeader FuncBody .
func (p *parser) parseFuncDef() error {
	f, err := p.parseFuncHeader()
	if err != nil {
		return errutil.Err(err)
	}
	f.Blocks, err = p.parseFuncBody()
	if err != nil {
		return errutil.Err(err)
	}
	p.m.Funcs = append(p.m.Funcs, f)
	return nil
}

// parseFuncHeader parses a function header consisting of a return argument, a
// function name and zero or more function arguments.
//
//    FuncHeader = FuncResult FuncName "(" FuncParams ")" .
//
//    FuncName = Global .
func (p *parser) parseFuncHeader() (header *ir.Function, err error) {
	result, err := p.parseType()
	if err != nil {
		return nil, err
	}
	name, ok := p.tryGlobal()
	if !ok {
		return nil, errutil.New("expected function name")
	}
	header = &ir.Function{
		Name: name,
	}
	if !p.accept(token.Lparen) {
		return nil, errors.New("expected '(' in function argument list")
	}
	// TODO: Don't use parseFuncType as it is specifically for function types
	// which do not include parameter names. Write a custom implementation for
	// function declarations and function definitions.
	_ = result
	//header.Sig, err = p.parseFunc(result)
	return header, err
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
	panic("parser.parseFuncBody: not yet implemented.")
}
