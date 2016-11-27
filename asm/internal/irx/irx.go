package irx

import (
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"

	"github.com/llir/llvm/asm/internal/token"
	"github.com/llir/llvm/ir"
	"github.com/llir/llvm/ir/constant"
	"github.com/llir/llvm/ir/types"
	"github.com/llir/llvm/ir/value"
	"github.com/pkg/errors"
)

// dbg is a logger which prefixes debug messages with the file name and line
// number of callees.
var dbg = log.New(os.Stdout, "", log.Lshortfile)

// === [ Modules ] =============================================================

// NewModule returns a new module based on the given top-level declarations.
func NewModule(decls interface{}) (*ir.Module, error) {
	ds, ok := decls.([]TopLevelDecl)
	if !ok {
		return nil, errors.Errorf("invalid top-level declaration list type; expected []irx.TopLevelDecl, got %T", decls)
	}
	m := ir.NewModule()
	for _, d := range ds {
		switch d := d.(type) {
		case *ir.Global:
			m.AppendGlobal(d)
		case *ir.Function:
			m.AppendFunction(d)
		default:
			dbg.Printf("support for %T not yet implemented", d)
		}
	}
	// TODO: Replace dummy values with their real values.
	return m, nil
}

// TopLevelDecl represents a top-level declaration.
type TopLevelDecl interface{}

// NewTopLevelDeclList returns a new top-level declaration list based on the
// given top-level declaration.
func NewTopLevelDeclList(decl interface{}) ([]TopLevelDecl, error) {
	d, ok := decl.(TopLevelDecl)
	if !ok {
		return nil, errors.Errorf("invalid top-level declaration type; expected irx.TopLevelDecl, got %T", decl)
	}
	return []TopLevelDecl{d}, nil
}

// AppendTopLevelDecl appends the given top-level declaration to the top-level
// declaration list.
func AppendTopLevelDecl(decls, decl interface{}) ([]TopLevelDecl, error) {
	ds, ok := decls.([]TopLevelDecl)
	if !ok {
		return nil, errors.Errorf("invalid top-level declaration list type; expected []irx.TopLevelDecl, got %T", decls)
	}
	d, ok := decl.(TopLevelDecl)
	if !ok {
		return nil, errors.Errorf("invalid top-level declaration type; expected irx.TopLevelDecl, got %T", decl)
	}
	return append(ds, d), nil
}

// === [ Global variables ] ====================================================

// NewGlobalDecl returns a new global variable declaration based on the given
// global variable name, immutability and type.
func NewGlobalDecl(name, immutable, typ interface{}) (*ir.Global, error) {
	n, ok := name.(*GlobalIdent)
	if !ok {
		return nil, errors.Errorf("invalid global name type; expected *irx.GlobalIdent, got %T", name)
	}
	imm, ok := immutable.(bool)
	if !ok {
		return nil, errors.Errorf("invalid immutability type; expected bool, got %T", immutable)
	}
	t, ok := typ.(types.Type)
	if !ok {
		return nil, errors.Errorf("invalid content type; expected types.Type, got %T", typ)
	}
	global := ir.NewGlobalDecl(n.name, t)
	global.SetImmutable(imm)
	return global, nil
}

// NewGlobalDef returns a new global variable definition based on the given
// global variable name, immutability, type and value.
func NewGlobalDef(name, immutable, typ, val interface{}) (*ir.Global, error) {
	n, ok := name.(*GlobalIdent)
	if !ok {
		return nil, errors.Errorf("invalid global name type; expected *irx.GlobalIdent, got %T", name)
	}
	imm, ok := immutable.(bool)
	if !ok {
		return nil, errors.Errorf("invalid immutability type; expected bool, got %T", immutable)
	}
	init, err := NewValue(typ, val)
	if err != nil {
		return nil, errors.WithStack(err)
	}
	global := ir.NewGlobalDef(n.name, init)
	global.SetImmutable(imm)
	return global, nil
}

// === [ Functions ] ===========================================================

// NewFunctionDecl returns a new function declaration based on the given
// function header and body.
func NewFunctionDecl(result, name, params interface{}) (*ir.Function, error) {
	r, ok := result.(types.Type)
	if !ok {
		return nil, errors.Errorf("invalid function return type; expected types.Type, got %T", result)
	}
	n, ok := name.(*GlobalIdent)
	if !ok {
		return nil, errors.Errorf("invalid function name type; expected *irx.GlobalIdent, got %T", name)
	}
	f := ir.NewFunction(n.name, r)
	switch ps := params.(type) {
	case *Params:
		for _, param := range ps.params {
			f.AppendParam(param)
		}
		f.SetVariadic(ps.variadic)
	case nil:
		// no parameters.
	default:
		return nil, errors.Errorf("invalid function parameters type; expected *irx.Params or nil, got %T", params)
	}
	return f, nil
}

// NewFunctionDef returns a new function definition based on the given function
// header and body.
func NewFunctionDef(header, body interface{}) (*ir.Function, error) {
	f, ok := header.(*ir.Function)
	if !ok {
		return nil, errors.Errorf("invalid function header type; expected *ir.Function, got %T", header)
	}
	blocks, ok := body.([]*ir.BasicBlock)
	if !ok {
		return nil, errors.Errorf("invalid function body type; expected []*ir.BasicBlock, got %T", body)
	}
	for _, block := range blocks {
		f.AppendBlock(block)
	}
	return f, nil
}

// Params represents a function parameters specifier.
type Params struct {
	// Function parameter types.
	params []*types.Param
	// Variadicity of the function type.
	variadic bool
}

// NewParams returns a new function parameters specifier, based on the given
// function parameters and variadicity.
func NewParams(params interface{}, variadic bool) (*Params, error) {
	switch params := params.(type) {
	case []*types.Param:
		return &Params{params: params, variadic: variadic}, nil
	case nil:
		return &Params{variadic: variadic}, nil
	default:
		return nil, errors.Errorf("invalid function parameter list; expected []*types.Param or nil, got %T", params)
	}
}

// NewParamList returns a new function parameter list based on the given
// function parameter.
func NewParamList(param interface{}) ([]*types.Param, error) {
	p, ok := param.(*types.Param)
	if !ok {
		return nil, errors.Errorf("invalid function parameter type; expected *types.Param, got %T", param)
	}
	return []*types.Param{p}, nil
}

// AppendParam appends the given parameter to the function parameter list.
func AppendParam(params, param interface{}) ([]*types.Param, error) {
	ps, ok := params.([]*types.Param)
	if !ok {
		return nil, errors.Errorf("invalid function parameter list type; expected []*types.Param, got %T", params)
	}
	p, ok := param.(*types.Param)
	if !ok {
		return nil, errors.Errorf("invalid function parameter type; expected *types.Param, got %T", param)
	}
	return append(ps, p), nil
}

// NewParam returns a new function parameter based on the given parameter type
// and name.
func NewParam(typ, name interface{}) (*types.Param, error) {
	t, ok := typ.(types.Type)
	if !ok {
		return nil, errors.Errorf("invalid type; expected types.Type, got %T", typ)
	}
	var n string
	switch name := name.(type) {
	case *LocalIdent:
		n = name.name
	case nil:
		// unnamed function parameter.
	default:
		return nil, errors.Errorf("invalid local name type; expected *irx.LocalIdent or nil, got %T", name)
	}
	return types.NewParam(n, t), nil
}

// === [ Identifiers ] =========================================================

// GlobalIdent represents a global identifier.
type GlobalIdent struct {
	// Global identifier name the without "@" prefix.
	name string
}

// NewGlobalIdent returns a new global identifier based on the given global
// identifier token.
func NewGlobalIdent(ident interface{}) (*GlobalIdent, error) {
	s, err := getTokenString(ident)
	if err != nil {
		return nil, errors.WithStack(err)
	}
	if !strings.HasPrefix(s, "@") {
		return nil, errors.Errorf(`invalid global identifier %q; missing "@" prefix`, s)
	}
	s = s[1:]
	return &GlobalIdent{name: s}, nil
}

// LocalIdent represents a local identifier.
type LocalIdent struct {
	// Local identifier name the without "%" prefix.
	name string
}

// NewLocalIdent returns a new local identifier based on the given local
// identifier token.
func NewLocalIdent(ident interface{}) (*LocalIdent, error) {
	s, err := getTokenString(ident)
	if err != nil {
		return nil, errors.WithStack(err)
	}
	if !strings.HasPrefix(s, "%") {
		return nil, errors.Errorf(`invalid local identifier %q; missing "%%" prefix`, s)
	}
	s = s[1:]
	return &LocalIdent{name: s}, nil
}

// LabelIdent represents a label identifier.
type LabelIdent struct {
	// Label identifier name the without ":" suffix.
	name string
}

// NewLabelIdent returns a new label identifier based on the given label
// identifier token.
func NewLabelIdent(ident interface{}) (*LabelIdent, error) {
	s, err := getTokenString(ident)
	if err != nil {
		return nil, errors.WithStack(err)
	}
	if !strings.HasSuffix(s, ":") {
		return nil, errors.Errorf(`invalid label identifier %q; missing ":" suffix`, s)
	}
	s = s[:len(s)-1]
	return &LabelIdent{name: s}, nil
}

// === [ Types ] ===============================================================

// NewIntType returns a new integer type based on the given integer type token.
func NewIntType(typeTok interface{}) (*types.IntType, error) {
	s, err := getTokenString(typeTok)
	if err != nil {
		return nil, errors.WithStack(err)
	}
	if !strings.HasPrefix(s, "i") {
		return nil, errors.Errorf(`invalid integer type %q; missing "i" prefix`, s)
	}
	s = s[1:]
	size, err := strconv.Atoi(s)
	if err != nil {
		return nil, errors.WithStack(err)
	}
	return types.NewInt(size), nil
}

// NewPointerType returns a new pointer type based on the given element type.
func NewPointerType(elem interface{}) (*types.PointerType, error) {
	e, ok := elem.(types.Type)
	if !ok {
		return nil, errors.Errorf("invalid element type; expected types.Type, got %T", elem)
	}
	return types.NewPointer(e), nil
}

// === [ Values ] ==============================================================

// NewValueList returns a new value list based on the given
// value.
func NewValueList(val interface{}) ([]value.Value, error) {
	v, ok := val.(value.Value)
	if !ok {
		return nil, errors.Errorf("invalid value type; expected value.Value, got %T", val)
	}
	return []value.Value{v}, nil
}

// AppendValue appends the given value to the value list.
func AppendValue(vals, val interface{}) ([]value.Value, error) {
	vs, ok := vals.([]value.Value)
	if !ok {
		return nil, errors.Errorf("invalid value list type; expected []value.Value, got %T", vals)
	}
	v, ok := val.(value.Value)
	if !ok {
		return nil, errors.Errorf("invalid value type; expected value.Value, got %T", val)
	}
	return append(vs, v), nil
}

// NewValue returns a value based on the given type and value.
func NewValue(typ, val interface{}) (value.Value, error) {
	t, ok := typ.(types.Type)
	if !ok {
		return nil, errors.Errorf("invalid value type; expected types.Type, got %T", typ)
	}
	switch val := val.(type) {
	case *LocalIdent:
		return newLocalDummy(val.name, t), nil
	case *GlobalIdent:
		return newGlobalDummy(val.name, t), nil
	}
	switch t := t.(type) {
	case *types.IntType:
		switch val := val.(type) {
		case *IntLit:
			return constant.NewIntFromString(val.lit, t), nil
		default:
			panic(fmt.Sprintf("support for value type %T not yet implemented", val))
		}
	default:
		panic(fmt.Sprintf("support for type %T not yet implemented", t))
	}
}

// === [ Constants ] ===========================================================

// IntLit represents an integer literal.
type IntLit struct {
	// Integer literal.
	lit string
}

// NewIntLit returns a new integer literal based on the given integer token.
func NewIntLit(tok interface{}) (*IntLit, error) {
	s, err := getTokenString(tok)
	if err != nil {
		return nil, errors.WithStack(err)
	}
	return &IntLit{lit: s}, nil
}

// === [ Basic blocks ] ========================================================

// NewBasicBlockList returns a new basic block list based on the given basic
// block.
func NewBasicBlockList(block interface{}) ([]*ir.BasicBlock, error) {
	b, ok := block.(*ir.BasicBlock)
	if !ok {
		return nil, errors.Errorf("invalid basic block type; expected *ir.BasicBlock, got %T", block)
	}
	return []*ir.BasicBlock{b}, nil
}

// AppendBasicBlock appends the given basic block to the basic block list.
func AppendBasicBlock(blocks, block interface{}) ([]*ir.BasicBlock, error) {
	bs, ok := blocks.([]*ir.BasicBlock)
	if !ok {
		return nil, errors.Errorf("invalid basic block list type; expected []*ir.BasicBlock, got %T", blocks)
	}
	b, ok := block.(*ir.BasicBlock)
	if !ok {
		return nil, errors.Errorf("invalid basic block type; expected *ir.BasicBlock, got %T", block)
	}
	return append(bs, b), nil
}

// NewBasicBlock returns a new basic block based on the given label name, non-
// branching instructions and terminator.
func NewBasicBlock(name, insts, term interface{}) (*ir.BasicBlock, error) {
	block := ir.NewBlock("")
	switch name := name.(type) {
	case *LabelIdent:
		block.SetIdent(name.name)
	case nil:
		// unnamed basic block.
	default:
		return nil, errors.Errorf("invalid label name type; expected *irx.LabelIdent or nil, got %T", name)
	}
	is, ok := insts.([]ir.Instruction)
	if !ok {
		return nil, errors.Errorf("invalid instruction list type; expected []ir.Instruction, got %T", insts)
	}
	t, ok := term.(ir.Terminator)
	if !ok {
		return nil, errors.Errorf("invalid terminator type; expected ir.Terminator, got %T", term)
	}
	for _, inst := range is {
		block.AppendInst(inst)
	}
	block.SetTerm(t)
	return block, nil
}

// === [ Instructions ] ========================================================

// NewInstructionList returns a new instruction list based on the given
// instruction.
func NewInstructionList(inst interface{}) ([]ir.Instruction, error) {
	i, ok := inst.(ir.Instruction)
	if !ok {
		return nil, errors.Errorf("invalid instruction type; expected ir.Instruction, got %T", inst)
	}
	return []ir.Instruction{i}, nil
}

// AppendInstruction appends the given instruction to the instruction list.
func AppendInstruction(insts, inst interface{}) ([]ir.Instruction, error) {
	is, ok := insts.([]ir.Instruction)
	if !ok {
		return nil, errors.Errorf("invalid instruction list type; expected []ir.Instruction, got %T", insts)
	}
	i, ok := inst.(ir.Instruction)
	if !ok {
		return nil, errors.Errorf("invalid instruction type; expected ir.Instruction, got %T", inst)
	}
	return append(is, i), nil
}

// NewNamedInstruction returns a named instruction based on the given local
// variable name and instruction.
func NewNamedInstruction(name, inst interface{}) (ir.Instruction, error) {
	// namedInstruction represents a namedInstruction instruction.
	type namedInstruction interface {
		ir.Instruction
		// SetIdent sets the identifier associated with the value.
		SetIdent(ident string)
	}
	n, ok := name.(*LocalIdent)
	if !ok {
		return nil, errors.Errorf("invalid local variable name type; expected *irx.LocalIdent, got %T", name)
	}
	i, ok := inst.(namedInstruction)
	if !ok {
		return nil, errors.Errorf("invalid instruction type; expected namedInstruction, got %T", inst)
	}
	i.SetIdent(n.name)
	return i, nil
}

// --- [ Binary instructions ] -------------------------------------------------

// NewAddInst returns a new add instruction based on the given type and
// operands.
func NewAddInst(typ, xVal, yVal interface{}) (*ir.InstAdd, error) {
	x, err := NewValue(typ, xVal)
	if err != nil {
		return nil, errors.WithStack(err)
	}
	y, err := NewValue(typ, yVal)
	if err != nil {
		return nil, errors.WithStack(err)
	}
	return ir.NewAdd(x, y), nil
}

// NewMulInst returns a new mul instruction based on the given type and
// operands.
func NewMulInst(typ, xVal, yVal interface{}) (*ir.InstMul, error) {
	x, err := NewValue(typ, xVal)
	if err != nil {
		return nil, errors.WithStack(err)
	}
	y, err := NewValue(typ, yVal)
	if err != nil {
		return nil, errors.WithStack(err)
	}
	return ir.NewMul(x, y), nil
}

// --- [ Bitwise instructions ] ------------------------------------------------

// --- [ Memory instructions ] -------------------------------------------------

// NewLoadInst returns a new load instruction based on the given element type,
// source address type and value.
func NewLoadInst(elem, srcTyp, src interface{}) (*ir.InstLoad, error) {
	e, ok := elem.(types.Type)
	if !ok {
		return nil, errors.Errorf("invalid element type; expected types.Type, got %T", elem)
	}
	s, err := NewValue(srcTyp, src)
	if err != nil {
		return nil, errors.WithStack(err)
	}
	inst := ir.NewLoad(s)
	if !types.Equal(inst.Type(), e) {
		return nil, errors.Errorf("element type mismatch; expected %v, got %v", inst.Type(), e)
	}
	return inst, nil
}

// NewStoreInst returns a new store instruction based on the given element type,
// source address type and value.
func NewStoreInst(srcTyp, srcVal, dstTyp, dstVal interface{}) (*ir.InstStore, error) {
	src, err := NewValue(srcTyp, srcVal)
	if err != nil {
		return nil, errors.WithStack(err)
	}
	dst, err := NewValue(dstTyp, dstVal)
	if err != nil {
		return nil, errors.WithStack(err)
	}
	return ir.NewStore(src, dst), nil
}

// --- [ Conversion instructions ] ---------------------------------------------

// --- [ Other instructions ] --------------------------------------------------

// NewCallInst returns a new call instruction based on the given return type,
// callee name, and function arguments.
func NewCallInst(retTyp, callee, args interface{}) (*instCallDummy, error) {
	r, ok := retTyp.(types.Type)
	if !ok {
		return nil, errors.Errorf("invalid return type; expected types.Type, got %T", retTyp)
	}
	c, ok := callee.(*GlobalIdent)
	if !ok {
		return nil, errors.Errorf("invalid callee type; expected *irx.GlobalIdent, got %T", callee)
	}
	var as []value.Value
	switch args := args.(type) {
	case []value.Value:
		as = args
	case nil:
		// no arguments.
	default:
		return nil, errors.Errorf("invalid function arguments type; expected []value.Value or nil, got %T", args)
	}
	return newCallDummy(r, c.name, as...), nil
}

// === [ Terminators ] =========================================================

// NewRetTerm returns a new ret terminator based on the given return type and
// value.
func NewRetTerm(typ, val interface{}) (*ir.TermRet, error) {
	v, err := NewValue(typ, val)
	if err != nil {
		return nil, errors.WithStack(err)
	}
	return ir.NewRet(v), nil
}

// ### [ Helper functions ] ####################################################

// getTokenString returns the string literal of the given token.
func getTokenString(tok interface{}) (string, error) {
	t, ok := tok.(*token.Token)
	if !ok {
		return "", errors.Errorf("invalid token type; expected *token.Token, got %T", tok)
	}
	return string(t.Lit), nil
}
