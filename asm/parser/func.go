package parser

import (
	"errors"

	"github.com/llir/llvm/asm"
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

	// TODO: The function argument list parsing is almost identical to
	// parseFuncType (except for argument names); see if it would be possible to
	// merge these.

	// Parse function argument list.
	if !p.accept(token.Lparen) {
		return nil, errors.New("expected '(' in function argument list")
	}
	// Early return for empty parameter list.
	if p.accept(token.Rparen) {
		header.Sig, err = types.NewFunc(result, nil, false)
		if err != nil {
			return nil, errutil.Err(err)
		}
		return header, nil
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
				if contains(header.Args, arg) {
					return nil, errutil.Newf("redefinition of argument %q", asm.EncLocal(arg))
				}
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
	return header, nil
}

// parseFuncBody parses a function body consisting of one or more basic blocks.
//
//    FuncBody = "{" BasicBlock { BasicBlock } "}" .
func (p *parser) parseFuncBody() (blocks []*ir.BasicBlock, err error) {
	// Parse left brace "{".
	if !p.accept(token.Lbrace) {
		return nil, errutil.Newf("expected '{' in function body")
	}

	// Parse basic blocks.
	for !p.accept(token.Rbrace) {
		block, err := p.parseBlock()
		if err != nil {
			return nil, errutil.Err(err)
		}
		blocks = append(blocks, block)
	}

	if len(blocks) < 1 {
		return nil, errutil.New("function body requires at least one basic block")
	}
	return blocks, nil
}

// parseBlock parses and returns a basic block.
//
//    BasicBlock  = [ LabelDecl ] { Instruction } Terminator .
//    LabelDecl   = Label ":" .
//    Label       = Local .
func (p *parser) parseBlock() (block *ir.BasicBlock, err error) {
	// Parse optional label declaration.
	block = new(ir.BasicBlock)
	if label, ok := p.try(token.Label); ok {
		block.Name = label
	}

	// TODO: Figure out a cleaner way to implement parsing of basic blocks.

	// Parse instructions.
	for {
		// Try to parse terminator instruction
		term, err := p.parseTerm()
		if err == nil {
			block.Term = term
			break
		}

		// Try to parse instruction.
		inst, e2 := p.parseInst()
		if e2 != nil {
			// If unable to locate neither terminator nor instruction, return the
			// initial error; i.e. the one from terminator parsing.
			return nil, errutil.Err(err)
		}
		block.Insts = append(block.Insts, inst)
	}

	return block, nil
}

// contains returns true if xs contains y.
func contains(xs []string, y string) bool {
	for _, x := range xs {
		if x == y {
			return true
		}
	}
	return false
}
