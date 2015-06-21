package parser

import (
	"errors"

	"github.com/llir/llvm/asm/token"
	"github.com/llir/llvm/ir"
	"github.com/llir/llvm/ir/types"
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
//    FuncHeader = FuncResult FuncName "(" FuncArgs ")" .
//
//    FuncName = Global .
//    FuncArgs = [ FuncArg { "," FuncArg } [ "," "..." ] ] | "..." .
//    FuncArg  = [ Local ] Type .
func (p *parser) parseFuncHeader() (header *ir.Function, err error) {
	// Parse return type.
	result, err := p.parseType()
	if err != nil {
		return nil, err
	}

	// Parse function name.
	name, ok := p.tryGlobal()
	if !ok {
		return nil, errutil.New("expected function name")
	}
	header = &ir.Function{
		Name: name,
	}

	// Parse function argument list.
	if !p.accept(token.Lparen) {
		return nil, errors.New("expected '(' in function argument list")
	}
	var params []types.Type
	variadic := false
	for i := 0; ; i++ {
		if i > 0 && !p.accept(token.Comma) {
			break
		}
		if param, ok := p.tryType(); ok {
			params = append(params, param)
			// Argument name.
			if arg, ok := p.tryLocal(); ok {
				header.Args = append(header.Args, arg)
			} else {
				// Unnamed function argument.
				header.Args = append(header.Args, "")
			}
		} else if p.accept(token.Ellipsis) {
			variadic = true
			break
		} else {
			return nil, errutil.New("expected type")
		}
	}
	if !p.accept(token.Rparen) {
		return nil, errutil.New("expected ')' at end of argument list")
	}

	// Create function signature.
	header.Sig, err = types.NewFunc(result, params, variadic)
	if err != nil {
		return nil, errutil.Err(err)
	}
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
