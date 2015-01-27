package parser

// decimal_digit = "0" … "9" .
// int_lit       =  decimal_digit { decimal_digit } .

// Global = GlobalID | GlobalVar .
// Local = LocalID | LocalVar .

import (
	"errors"
	"io"
	"strconv"
	"strings"

	"github.com/mewkiz/pkg/errutil"
	"github.com/mewlang/llvm/asm/token"
	"github.com/mewlang/llvm/ir"
	"github.com/mewlang/llvm/types"
)

// TODO: Complete TopLevelEntity EBNF definition.

// parseTopLevelEntity parses a top-level entity and stores it in module.
//
//    TopLevelEntity = FuncDecl | FuncDef .
func (p *parser) parseTopLevelEntity(module *ir.Module) error {
	tok := p.next()
	switch tok.Kind {
	case token.Error:
		return errutil.New(tok.Val)
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
	default:
		return errutil.Newf("invalid token type %v; expected top-level entity", tok.Kind)
	}
	return nil
}

// parseDeclare parses a function declaration. A "declare" token has already
// been consumed.
//
//    FuncDecl = "declare" FuncHeader .
func (p *parser) parseDeclare() (*ir.Function, error) {
	return p.parseFuncHeader()
}

// parseDefine parses a function definition. A "define" token has already been
// consumed.
//
//    FuncDef = "define" FuncHeader FuncBody .
func (p *parser) parseDefine() (*ir.Function, error) {
	f, err := p.parseFuncHeader()
	if err != nil {
		return nil, err
	}
	f.Blocks, err = p.parseFuncBody()
	return f, err
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
	header.Sig, err = p.parseFunc(result)
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
	panic("not yet implemented.")
}

// parseType parses a type.
//
//    Type = VoidType | IntType | FloatType | MMXType | LabelType |
//           MetadataType | FuncType | PointerType | VectorType | ArrayType |
//           StructType .
//
//    VoidType        = "void" .
//    IntType         = "i" int_lit .
//    FloatType       = "half" | "float" | "double" | "fp128" | x86_fp80 |
//                      "ppc_fp128" .
//    MMXType         = "x86_mmx" .
//    LabelType       = "label" .
//    MetadataType    = "metadata" .
//    FuncType        = FuncResultType "(" ( FuncParamType { "," FuncParamType } ] [ "," "..." ]) | [ "..." ] ")" .
//    FuncResultType  = VoidType | IntType | FloatType | MMXType | PointerType |
//                      VectorType | ArrayType | StructType .
//    FuncParamType   = IntType | FloatType | MMXType | LabelType |
//                      MetadataType | PointerType | VectorType | ArrayType |
//                      StructType .
//    PointerType     = (IntType | FloatType | MMXType | FuncType |
//                      PointerType | VectorType | ArrayType | StructType) "*" .
//
//    IntsType   = ( IntType | IntVectorType ) .
//    FloatsType = ( FloatType | FloatVectorType ) .
func (p *parser) parseType() (typ types.Type, err error) {
	switch tok := p.next(); tok.Kind {
	// Basic type
	//    i32
	//    float
	case token.Type:
		typ, err = basicTypeFromString(tok.Val)
		if err != nil {
			return nil, err
		}

	case token.Less:
		if p.accept(token.Lbrace) {
			// Packed array type
			//    <{i32, i8}>
			typ, err = p.parseStruct(true)
			if err != nil {
				return nil, err
			}
		} else {
			// Vector type
			//    <2 x i32>
			typ, err = p.parseVector()
			if err != nil {
				return nil, err
			}
		}

	// Array type
	//    [2 x float]
	case token.Lbrack:
		typ, err = p.parseArray()
		if err != nil {
			return nil, err
		}

	// Structure type
	//    {i32, float}
	case token.Lbrace:
		typ, err = p.parseStruct(false)
		if err != nil {
			return nil, err
		}

	default:
		return nil, errutil.New("expected type")
	}

	for {
		// Pointer type
		//    i32*
		//    [2 x float]*
		//    i8****
		for p.accept(token.Star) {
			elem := typ
			typ, err = types.NewPointer(elem)
			if err != nil {
				return nil, err
			}
		}

		// Function type
		//    i32 (i8*, ...)
		//    [2 x float]* (i32)
		//    i32 (i32)* (i32)
		if p.accept(token.Lparen) {
			result := typ
			typ, err = p.parseFunc(result)
			if err != nil {
				return nil, err
			}
		} else {
			break
		}
	}

	if typ == nil {
		return nil, errutil.New("expected type")
	}
	return typ, nil
}

// parseVector parses a vector type. A "<" token has already been consumed.
//
// Syntax:
//    VectorType      = IntVectorType | FloatVectorType |
//                      "<" VectorLen "x" PointerType ">" .
//    IntVectorType   = "<" VectorLen "x" IntType ">" .
//    FloatVectorType = "<" VectorLen "x" FloatType ">" .
//    VectorLen       = int_lit .
//
// Example:
//    <2 x i32>
func (p *parser) parseVector() (*types.Vector, error) {
	// Vector length.
	s, ok := p.try(token.Int)
	if !ok {
		return nil, errutil.New("expected vector length")
	}
	n, err := strconv.Atoi(s)
	if err != nil {
		return nil, errutil.Newf("invalid vector length (%v); %v", s, err)
	}
	if n < 1 {
		return nil, errutil.Newf("invalid vector length (%d); expected >= 1", n)
	}

	// x
	if !p.accept(token.KwX) {
		return nil, errutil.New("expected 'x' after vector length")
	}

	// Element type.
	elem, err := p.parseType()
	if err != nil {
		return nil, err
	}

	return types.NewVector(elem, n)
}

// parseArray parses a array type. A "[" token has already been consumed.
//
// Syntax:
//    ArrayType = "[" ArrayLen "x" ElemType "]" .
//    ElemType  = IntType | FloatType | MMXType | PointerType | VectorType |
//                ArrayType | StructType .
//    ArrayLen  = int_lit .
//
// Example:
//    [5 x float]
func (p *parser) parseArray() (*types.Array, error) {
	// Array length.
	s, ok := p.try(token.Int)
	if !ok {
		return nil, errutil.New("expected array length")
	}
	n, err := strconv.Atoi(s)
	if err != nil {
		return nil, errutil.Newf("invalid array length (%v); %v", s, err)
	}
	if n < 0 {
		return nil, errutil.Newf("invalid array length (%d); expected >= 0", n)
	}

	// x
	if !p.accept(token.KwX) {
		return nil, errutil.New("expected 'x' after array length")
	}

	// Element type.
	elem, err := p.parseType()
	if err != nil {
		return nil, err
	}

	return types.NewArray(elem, n)
}

// parseStruct parses a structure type. The structure is 1 byte aligned if
// packed is true. A "{" token has already been consumed, unless the structure
// is packed in which case a "<" and a "{" token have already been consumed.
//
//    StructType = "{" [ ElemType { "," ElemType } ] "}" | "<" "{" [ ElemType { "," ElemType } ] "}" ">" .
func (p *parser) parseStruct(packed bool) (*types.Struct, error) {
	panic("not yet implemented.")
}

// parseFunc parses a function type. A result type, an optional function name
// and a "(" token has already been consumed.
//
//    FuncType   = FuncResult "(" FuncParams ")" .
//    FuncResult = Type .
//    FuncParams = [ FuncParam { "," FuncParam } [ "," "..." ] ] | "..." .
//    FuncParam  = Type [ Local ] .
func (p *parser) parseFunc(result types.Type) (typ *types.Func, err error) {
	// Early return for empty parameter list.
	if p.accept(token.Rparen) {
		return types.NewFunc(result, nil, false)
	}

	// Function parameters.
	var params []types.Type
	variadic := false
	for i := 0; ; i++ {
		if i != 0 && !p.accept(token.Comma) {
			break
		}
		// TODO: Figure out how to handle and where to store store function
		// parameter names.
		//    i32 (i32 %foo, i32 %bar)
		if param, ok := p.tryType(); ok {
			params = append(params, param)
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

	return types.NewFunc(result, params, variadic)
}

// basicTypeFromString returns the basic type corresponding to s.
func basicTypeFromString(s string) (types.Type, error) {
	switch s {
	case "void":
		return types.NewVoid(), nil
	case "half":
		return types.NewFloat(types.Float16), nil
	case "float":
		return types.NewFloat(types.Float32), nil
	case "double":
		return types.NewFloat(types.Float64), nil
	case "fp128":
		return types.NewFloat(types.Float128), nil
	case "x86_fp80":
		return types.NewFloat(types.X86Float80), nil
	case "ppc_fp128":
		return types.NewFloat(types.PPCFloat128), nil
	case "x86_mmx":
		return types.NewMMX(), nil
	case "label":
		return types.NewLabel(), nil
	case "metadata":
		return types.NewMetadata(), nil
	}

	// Integer type (e.g. i32).
	if !strings.HasPrefix(s, "i") {
		return nil, errutil.Newf("unknown basic type %q", s)
	}
	n, err := strconv.Atoi(s[1:]) // skip leading "i".
	if err != nil {
		return nil, errutil.Newf("unknown basic type %q", s)
	}
	return types.NewInt(n)
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
