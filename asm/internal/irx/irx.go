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
	return fixModule(m), nil
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
	global.SetConst(imm)
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
	i, ok := init.(constant.Constant)
	if !ok {
		panic(fmt.Sprintf("invalid init type; expected constant.Constant, got %T", init))
	}
	global := ir.NewGlobalDef(n.name, i)
	global.SetConst(imm)
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
func NewPointerType(elem, space interface{}) (*types.PointerType, error) {
	e, ok := elem.(types.Type)
	if !ok {
		return nil, errors.Errorf("invalid element type; expected types.Type, got %T", elem)
	}
	t := types.NewPointer(e)
	switch space := space.(type) {
	case *addrSpace:
		t.SetAddrSpace(space.space)
	case nil:
		// no address space.
	default:
		return nil, errors.Errorf("invalid address space type; expected *irx.addrSpace or nil, got %T", space)
	}
	return t, nil
}

// addrSpace represents the address space of a pointer type.
type addrSpace struct {
	// Address space.
	space int64
}

// NewAddrSpace returns a new address space pointer based on the given space.
func NewAddrSpace(space interface{}) (*addrSpace, error) {
	s, err := getInt64(space)
	if err != nil {
		return nil, errors.WithStack(err)
	}
	return &addrSpace{space: s}, nil
}

// NewVectorType returns a new vector type based on the given vector length and
// element type.
func NewVectorType(len, elem interface{}) (*types.VectorType, error) {
	e, ok := elem.(types.Type)
	if !ok {
		return nil, errors.Errorf("invalid element type; expected types.Type, got %T", elem)
	}
	l, err := getInt64(len)
	if err != nil {
		return nil, errors.WithStack(err)
	}
	return types.NewVector(e, l), nil
}

// NewArrayType returns a new array type based on the given array length and
// element type.
func NewArrayType(len, elem interface{}) (*types.ArrayType, error) {
	e, ok := elem.(types.Type)
	if !ok {
		return nil, errors.Errorf("invalid element type; expected types.Type, got %T", elem)
	}
	l, err := getInt64(len)
	if err != nil {
		return nil, errors.WithStack(err)
	}
	return types.NewArray(e, l), nil
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
	case *types.FloatType:
		switch val := val.(type) {
		case *FloatLit:
			return constant.NewFloatFromString(val.lit, t), nil
		default:
			panic(fmt.Sprintf("support for value type %T not yet implemented", val))
		}
	case *types.PointerType:
		switch val := val.(type) {
		case *NullLit:
			return constant.NewNull(t), nil
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

// FloatLit represents an floating-point literal.
type FloatLit struct {
	// Floating-point literal.
	lit string
}

// NewFloatLit returns a new floating-point literal based on the given floating-point  token.
func NewFloatLit(tok interface{}) (*FloatLit, error) {
	s, err := getTokenString(tok)
	if err != nil {
		return nil, errors.WithStack(err)
	}
	return &FloatLit{lit: s}, nil
}

// NullLit represents a null literal.
type NullLit struct {
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
	var is []ir.Instruction
	switch insts := insts.(type) {
	case []ir.Instruction:
		is = insts
	case nil:
		// no instructions.
	default:
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

// NewFAddInst returns a new fadd instruction based on the given type and
// operands.
func NewFAddInst(typ, xVal, yVal interface{}) (*ir.InstFAdd, error) {
	x, err := NewValue(typ, xVal)
	if err != nil {
		return nil, errors.WithStack(err)
	}
	y, err := NewValue(typ, yVal)
	if err != nil {
		return nil, errors.WithStack(err)
	}
	return ir.NewFAdd(x, y), nil
}

// NewSubInst returns a new sub instruction based on the given type and
// operands.
func NewSubInst(typ, xVal, yVal interface{}) (*ir.InstSub, error) {
	x, err := NewValue(typ, xVal)
	if err != nil {
		return nil, errors.WithStack(err)
	}
	y, err := NewValue(typ, yVal)
	if err != nil {
		return nil, errors.WithStack(err)
	}
	return ir.NewSub(x, y), nil
}

// NewFSubInst returns a new fsub instruction based on the given type and
// operands.
func NewFSubInst(typ, xVal, yVal interface{}) (*ir.InstFSub, error) {
	x, err := NewValue(typ, xVal)
	if err != nil {
		return nil, errors.WithStack(err)
	}
	y, err := NewValue(typ, yVal)
	if err != nil {
		return nil, errors.WithStack(err)
	}
	return ir.NewFSub(x, y), nil
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

// NewFMulInst returns a new fmul instruction based on the given type and
// operands.
func NewFMulInst(typ, xVal, yVal interface{}) (*ir.InstFMul, error) {
	x, err := NewValue(typ, xVal)
	if err != nil {
		return nil, errors.WithStack(err)
	}
	y, err := NewValue(typ, yVal)
	if err != nil {
		return nil, errors.WithStack(err)
	}
	return ir.NewFMul(x, y), nil
}

// NewUDivInst returns a new udiv instruction based on the given type and
// operands.
func NewUDivInst(typ, xVal, yVal interface{}) (*ir.InstUDiv, error) {
	x, err := NewValue(typ, xVal)
	if err != nil {
		return nil, errors.WithStack(err)
	}
	y, err := NewValue(typ, yVal)
	if err != nil {
		return nil, errors.WithStack(err)
	}
	return ir.NewUDiv(x, y), nil
}

// NewSDivInst returns a new sdiv instruction based on the given type and
// operands.
func NewSDivInst(typ, xVal, yVal interface{}) (*ir.InstSDiv, error) {
	x, err := NewValue(typ, xVal)
	if err != nil {
		return nil, errors.WithStack(err)
	}
	y, err := NewValue(typ, yVal)
	if err != nil {
		return nil, errors.WithStack(err)
	}
	return ir.NewSDiv(x, y), nil
}

// NewFDivInst returns a new fdiv instruction based on the given type and
// operands.
func NewFDivInst(typ, xVal, yVal interface{}) (*ir.InstFDiv, error) {
	x, err := NewValue(typ, xVal)
	if err != nil {
		return nil, errors.WithStack(err)
	}
	y, err := NewValue(typ, yVal)
	if err != nil {
		return nil, errors.WithStack(err)
	}
	return ir.NewFDiv(x, y), nil
}

// NewURemInst returns a new urem instruction based on the given type and
// operands.
func NewURemInst(typ, xVal, yVal interface{}) (*ir.InstURem, error) {
	x, err := NewValue(typ, xVal)
	if err != nil {
		return nil, errors.WithStack(err)
	}
	y, err := NewValue(typ, yVal)
	if err != nil {
		return nil, errors.WithStack(err)
	}
	return ir.NewURem(x, y), nil
}

// NewSRemInst returns a new srem instruction based on the given type and
// operands.
func NewSRemInst(typ, xVal, yVal interface{}) (*ir.InstSRem, error) {
	x, err := NewValue(typ, xVal)
	if err != nil {
		return nil, errors.WithStack(err)
	}
	y, err := NewValue(typ, yVal)
	if err != nil {
		return nil, errors.WithStack(err)
	}
	return ir.NewSRem(x, y), nil
}

// NewFRemInst returns a new frem instruction based on the given type and
// operands.
func NewFRemInst(typ, xVal, yVal interface{}) (*ir.InstFRem, error) {
	x, err := NewValue(typ, xVal)
	if err != nil {
		return nil, errors.WithStack(err)
	}
	y, err := NewValue(typ, yVal)
	if err != nil {
		return nil, errors.WithStack(err)
	}
	return ir.NewFRem(x, y), nil
}

// --- [ Bitwise instructions ] ------------------------------------------------

// NewShlInst returns a new shl instruction based on the given type and
// operands.
func NewShlInst(typ, xVal, yVal interface{}) (*ir.InstShl, error) {
	x, err := NewValue(typ, xVal)
	if err != nil {
		return nil, errors.WithStack(err)
	}
	y, err := NewValue(typ, yVal)
	if err != nil {
		return nil, errors.WithStack(err)
	}
	return ir.NewShl(x, y), nil
}

// NewLShrInst returns a new lshr instruction based on the given type and
// operands.
func NewLShrInst(typ, xVal, yVal interface{}) (*ir.InstLShr, error) {
	x, err := NewValue(typ, xVal)
	if err != nil {
		return nil, errors.WithStack(err)
	}
	y, err := NewValue(typ, yVal)
	if err != nil {
		return nil, errors.WithStack(err)
	}
	return ir.NewLShr(x, y), nil
}

// NewAShrInst returns a new ashr instruction based on the given type and
// operands.
func NewAShrInst(typ, xVal, yVal interface{}) (*ir.InstAShr, error) {
	x, err := NewValue(typ, xVal)
	if err != nil {
		return nil, errors.WithStack(err)
	}
	y, err := NewValue(typ, yVal)
	if err != nil {
		return nil, errors.WithStack(err)
	}
	return ir.NewAShr(x, y), nil
}

// NewAndInst returns a new and instruction based on the given type and
// operands.
func NewAndInst(typ, xVal, yVal interface{}) (*ir.InstAnd, error) {
	x, err := NewValue(typ, xVal)
	if err != nil {
		return nil, errors.WithStack(err)
	}
	y, err := NewValue(typ, yVal)
	if err != nil {
		return nil, errors.WithStack(err)
	}
	return ir.NewAnd(x, y), nil
}

// NewOrInst returns a new or instruction based on the given type and
// operands.
func NewOrInst(typ, xVal, yVal interface{}) (*ir.InstOr, error) {
	x, err := NewValue(typ, xVal)
	if err != nil {
		return nil, errors.WithStack(err)
	}
	y, err := NewValue(typ, yVal)
	if err != nil {
		return nil, errors.WithStack(err)
	}
	return ir.NewOr(x, y), nil
}

// NewXorInst returns a new xor instruction based on the given type and
// operands.
func NewXorInst(typ, xVal, yVal interface{}) (*ir.InstXor, error) {
	x, err := NewValue(typ, xVal)
	if err != nil {
		return nil, errors.WithStack(err)
	}
	y, err := NewValue(typ, yVal)
	if err != nil {
		return nil, errors.WithStack(err)
	}
	return ir.NewXor(x, y), nil
}

// --- [ Memory instructions ] -------------------------------------------------

// NewAllocaInst returns a new alloca instruction based on the given element
// type and number of elements.
func NewAllocaInst(elem, nelems interface{}) (*ir.InstAlloca, error) {
	e, ok := elem.(types.Type)
	if !ok {
		return nil, errors.Errorf("invalid element type; expected types.Type, got %T", elem)
	}
	inst := ir.NewAlloca(e)
	switch nelems := nelems.(type) {
	case value.Value:
		inst.SetNElems(nelems)
	case nil:
		// no nelems.
	default:
		return nil, errors.Errorf("invalid number of elements type; expected value.Value or nil, got %T", nelems)
	}
	return inst, nil
}

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

// NewGetElementPtrInst returns a new getelementptr instruction based on the
// given element type, source address type and value, and element indices.
func NewGetElementPtrInst(elem, srcTyp, srcVal, indices interface{}) (*ir.InstGetElementPtr, error) {
	e, ok := elem.(types.Type)
	if !ok {
		return nil, errors.Errorf("invalid element type; expected types.Type, got %T", elem)
	}
	st, ok := srcTyp.(*types.PointerType)
	if !ok {
		return nil, errors.Errorf("invalid source type; expected *types.Pointer, got %T", srcTyp)
	}
	if !e.Equal(st.Elem()) {
		return nil, errors.Errorf("type mismatch between element type `%v` and source element type `%v`", e, st)
	}
	src, err := NewValue(srcTyp, srcVal)
	if err != nil {
		return nil, errors.WithStack(err)
	}
	var is []value.Value
	switch indices := indices.(type) {
	case []value.Value:
		is = indices
	case nil:
		// no indices.
	default:
		return nil, errors.Errorf("invalid indices type; expected []value.Value or nil, got %T", indices)
	}
	return ir.NewGetElementPtr(src, is...), nil
}

// --- [ Conversion instructions ] ---------------------------------------------

// NewTruncInst returns a new trunc instruction based on the given source value
// and target type.
func NewTruncInst(fromTyp, fromVal, to interface{}) (*ir.InstTrunc, error) {
	from, err := NewValue(fromTyp, fromVal)
	if err != nil {
		return nil, errors.WithStack(err)
	}
	t, ok := to.(types.Type)
	if !ok {
		return nil, errors.Errorf("invalid type; expected types.Type, got %T", to)
	}
	return ir.NewTrunc(from, t), nil
}

// NewZExtInst returns a new zext instruction based on the given source value
// and target type.
func NewZExtInst(fromTyp, fromVal, to interface{}) (*ir.InstZExt, error) {
	from, err := NewValue(fromTyp, fromVal)
	if err != nil {
		return nil, errors.WithStack(err)
	}
	t, ok := to.(types.Type)
	if !ok {
		return nil, errors.Errorf("invalid type; expected types.Type, got %T", to)
	}
	return ir.NewZExt(from, t), nil
}

// NewSExtInst returns a new sext instruction based on the given source value
// and target type.
func NewSExtInst(fromTyp, fromVal, to interface{}) (*ir.InstSExt, error) {
	from, err := NewValue(fromTyp, fromVal)
	if err != nil {
		return nil, errors.WithStack(err)
	}
	t, ok := to.(types.Type)
	if !ok {
		return nil, errors.Errorf("invalid type; expected types.Type, got %T", to)
	}
	return ir.NewSExt(from, t), nil
}

// NewFPTruncInst returns a new fptrunc instruction based on the given source value
// and target type.
func NewFPTruncInst(fromTyp, fromVal, to interface{}) (*ir.InstFPTrunc, error) {
	from, err := NewValue(fromTyp, fromVal)
	if err != nil {
		return nil, errors.WithStack(err)
	}
	t, ok := to.(types.Type)
	if !ok {
		return nil, errors.Errorf("invalid type; expected types.Type, got %T", to)
	}
	return ir.NewFPTrunc(from, t), nil
}

// NewFPExtInst returns a new fpext instruction based on the given source value
// and target type.
func NewFPExtInst(fromTyp, fromVal, to interface{}) (*ir.InstFPExt, error) {
	from, err := NewValue(fromTyp, fromVal)
	if err != nil {
		return nil, errors.WithStack(err)
	}
	t, ok := to.(types.Type)
	if !ok {
		return nil, errors.Errorf("invalid type; expected types.Type, got %T", to)
	}
	return ir.NewFPExt(from, t), nil
}

// NewFPToUIInst returns a new fptoui instruction based on the given source value
// and target type.
func NewFPToUIInst(fromTyp, fromVal, to interface{}) (*ir.InstFPToUI, error) {
	from, err := NewValue(fromTyp, fromVal)
	if err != nil {
		return nil, errors.WithStack(err)
	}
	t, ok := to.(types.Type)
	if !ok {
		return nil, errors.Errorf("invalid type; expected types.Type, got %T", to)
	}
	return ir.NewFPToUI(from, t), nil
}

// NewFPToSIInst returns a new fptosi instruction based on the given source value
// and target type.
func NewFPToSIInst(fromTyp, fromVal, to interface{}) (*ir.InstFPToSI, error) {
	from, err := NewValue(fromTyp, fromVal)
	if err != nil {
		return nil, errors.WithStack(err)
	}
	t, ok := to.(types.Type)
	if !ok {
		return nil, errors.Errorf("invalid type; expected types.Type, got %T", to)
	}
	return ir.NewFPToSI(from, t), nil
}

// NewUIToFPInst returns a new uitofp instruction based on the given source value
// and target type.
func NewUIToFPInst(fromTyp, fromVal, to interface{}) (*ir.InstUIToFP, error) {
	from, err := NewValue(fromTyp, fromVal)
	if err != nil {
		return nil, errors.WithStack(err)
	}
	t, ok := to.(types.Type)
	if !ok {
		return nil, errors.Errorf("invalid type; expected types.Type, got %T", to)
	}
	return ir.NewUIToFP(from, t), nil
}

// NewSIToFPInst returns a new sitofp instruction based on the given source value
// and target type.
func NewSIToFPInst(fromTyp, fromVal, to interface{}) (*ir.InstSIToFP, error) {
	from, err := NewValue(fromTyp, fromVal)
	if err != nil {
		return nil, errors.WithStack(err)
	}
	t, ok := to.(types.Type)
	if !ok {
		return nil, errors.Errorf("invalid type; expected types.Type, got %T", to)
	}
	return ir.NewSIToFP(from, t), nil
}

// NewPtrToIntInst returns a new ptrtoint instruction based on the given source value
// and target type.
func NewPtrToIntInst(fromTyp, fromVal, to interface{}) (*ir.InstPtrToInt, error) {
	from, err := NewValue(fromTyp, fromVal)
	if err != nil {
		return nil, errors.WithStack(err)
	}
	t, ok := to.(types.Type)
	if !ok {
		return nil, errors.Errorf("invalid type; expected types.Type, got %T", to)
	}
	return ir.NewPtrToInt(from, t), nil
}

// NewIntToPtrInst returns a new inttoptr instruction based on the given source value
// and target type.
func NewIntToPtrInst(fromTyp, fromVal, to interface{}) (*ir.InstIntToPtr, error) {
	from, err := NewValue(fromTyp, fromVal)
	if err != nil {
		return nil, errors.WithStack(err)
	}
	t, ok := to.(types.Type)
	if !ok {
		return nil, errors.Errorf("invalid type; expected types.Type, got %T", to)
	}
	return ir.NewIntToPtr(from, t), nil
}

// NewBitCastInst returns a new bitcast instruction based on the given source value
// and target type.
func NewBitCastInst(fromTyp, fromVal, to interface{}) (*ir.InstBitCast, error) {
	from, err := NewValue(fromTyp, fromVal)
	if err != nil {
		return nil, errors.WithStack(err)
	}
	t, ok := to.(types.Type)
	if !ok {
		return nil, errors.Errorf("invalid type; expected types.Type, got %T", to)
	}
	return ir.NewBitCast(from, t), nil
}

// NewAddrSpaceCastInst returns a new addrspacecast instruction based on the given source value
// and target type.
func NewAddrSpaceCastInst(fromTyp, fromVal, to interface{}) (*ir.InstAddrSpaceCast, error) {
	from, err := NewValue(fromTyp, fromVal)
	if err != nil {
		return nil, errors.WithStack(err)
	}
	t, ok := to.(types.Type)
	if !ok {
		return nil, errors.Errorf("invalid type; expected types.Type, got %T", to)
	}
	return ir.NewAddrSpaceCast(from, t), nil
}

// --- [ Other instructions ] --------------------------------------------------

// NewPhiInst returns a new phi instruction based on the given incoming values.
func NewPhiInst(typ, incs interface{}) (*instPhiDummy, error) {
	t, ok := typ.(types.Type)
	if !ok {
		return nil, errors.Errorf("invalid type; expected types.Type, got %T", typ)
	}
	is, ok := incs.([]*incomingDummy)
	if !ok {
		return nil, errors.Errorf("invalid incoming value list type; expected []*irx.incomingDummy, got %T", incs)
	}
	for _, inc := range is {
		x, err := NewValue(typ, inc.x)
		if err != nil {
			return nil, errors.WithStack(err)
		}
		inc.x = x
	}
	return newPhiDummy(t, is...), nil
}

// NewIncomingList returns a new incoming value list based on the given incoming
// value.
func NewIncomingList(inc interface{}) ([]*incomingDummy, error) {
	i, ok := inc.(*incomingDummy)
	if !ok {
		return nil, errors.Errorf("invalid incoming value type; expected *irx.incomingDummy, got %T", inc)
	}
	return []*incomingDummy{i}, nil
}

// AppendIncoming appends the given incoming value to the incoming value list.
func AppendIncoming(incs, inc interface{}) ([]*incomingDummy, error) {
	is, ok := incs.([]*incomingDummy)
	if !ok {
		return nil, errors.Errorf("invalid incoming value list type; expected []*irx.incomingDummy, got %T", incs)
	}
	i, ok := inc.(*incomingDummy)
	if !ok {
		return nil, errors.Errorf("invalid incoming value type; expected *irx.incomingDummy, got %T", inc)
	}
	return append(is, i), nil
}

// NewIncoming returns a new incoming value based on the given value and
// predecessor basic block.
func NewIncoming(x, pred interface{}) (*incomingDummy, error) {
	p, ok := pred.(*LocalIdent)
	if !ok {
		return nil, errors.Errorf("invalid predecessor type; expected *irx.LocalIdent, got %T", pred)
	}
	return newIncomingDummy(x, p.name), nil
}

// NewICmpInst returns a new icmp instruction based on the given integer
// condition code, type and operands.
func NewICmpInst(cond, typ, xVal, yVal interface{}) (*ir.InstICmp, error) {
	c, ok := cond.(ir.IntPred)
	if !ok {
		return nil, errors.Errorf("invalid integer predicate type; expected ir.IntPred, got %T", cond)
	}
	x, err := NewValue(typ, xVal)
	if err != nil {
		return nil, errors.WithStack(err)
	}
	y, err := NewValue(typ, yVal)
	if err != nil {
		return nil, errors.WithStack(err)
	}
	return ir.NewICmp(c, x, y), nil
}

// NewFCmpInst returns a new fcmp instruction based on the given floating-point
// condition code, type and operands.
func NewFCmpInst(cond, typ, xVal, yVal interface{}) (*ir.InstFCmp, error) {
	c, ok := cond.(ir.FloatPred)
	if !ok {
		return nil, errors.Errorf("invalid floating-point predicate type; expected ir.FloatPred, got %T", cond)
	}
	x, err := NewValue(typ, xVal)
	if err != nil {
		return nil, errors.WithStack(err)
	}
	y, err := NewValue(typ, yVal)
	if err != nil {
		return nil, errors.WithStack(err)
	}
	return ir.NewFCmp(c, x, y), nil
}

// NewSelect returns a new select instruction based on the given selection
// condition type and value, and operands.
func NewSelectInst(condTyp, condVal, xTyp, xVal, yTyp, yVal interface{}) (*ir.InstSelect, error) {
	cond, err := NewValue(condTyp, condVal)
	if err != nil {
		return nil, errors.WithStack(err)
	}
	x, err := NewValue(xTyp, xVal)
	if err != nil {
		return nil, errors.WithStack(err)
	}
	y, err := NewValue(yTyp, yVal)
	if err != nil {
		return nil, errors.WithStack(err)
	}
	return ir.NewSelect(cond, x, y), nil
}

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

// NewBrTerm returns a new unconditional br terminator based on the given target
// branch.
func NewBrTerm(target interface{}) (*termBrDummy, error) {
	t, ok := target.(*LocalIdent)
	if !ok {
		return nil, errors.Errorf("invalid target branch type; expected *irx.LocalIdent, got %T", target)
	}
	return newBrDummy(t.name), nil
}

// NewCondBrTerm returns a new conditional br terminator based on the given
// branching condition type and value, and conditional target branches.
func NewCondBrTerm(condTyp, condVal, targetTrue, targetFalse interface{}) (*termCondBrDummy, error) {
	cond, err := NewValue(condTyp, condVal)
	if err != nil {
		return nil, errors.WithStack(err)
	}
	tTrue, ok := targetTrue.(*LocalIdent)
	if !ok {
		return nil, errors.Errorf("invalid true target branch type; expected *irx.LocalIdent, got %T", targetTrue)
	}
	tFalse, ok := targetFalse.(*LocalIdent)
	if !ok {
		return nil, errors.Errorf("invalid true target branch type; expected *irx.LocalIdent, got %T", targetFalse)
	}
	return newCondBrDummy(cond, tTrue.name, tFalse.name), nil
}

// NewSwitchTerm returns a new switch terminator based on the given control
// variable type and value, default target branch and switch cases.
func NewSwitchTerm(xTyp, xVal, targetDefault, cases interface{}) (*termSwitchDummy, error) {
	x, err := NewValue(xTyp, xVal)
	if err != nil {
		return nil, errors.WithStack(err)
	}
	tDefault, ok := targetDefault.(*LocalIdent)
	if !ok {
		return nil, errors.Errorf("invalid default target branch type; expected *irx.LocalIdent, got %T", targetDefault)
	}
	var cs []*caseDummy
	switch cases := cases.(type) {
	case []*caseDummy:
		cs = cases
	case nil:
		// no cases.
	default:
		return nil, errors.Errorf("invalid switch cases type; expected []*irx.caseDummy or nil, got %T", cases)
	}
	return newSwitchDummy(x, tDefault.name, cs...), nil
}

// NewCaseList returns a new switch case list based on the given case.
func NewCaseList(switchCase interface{}) ([]*caseDummy, error) {
	c, ok := switchCase.(*caseDummy)
	if !ok {
		return nil, errors.Errorf("invalid switch case type; expected *irx.caseDummy, got %T", switchCase)
	}
	return []*caseDummy{c}, nil
}

// AppendCase appends the given case to the switch case list.
func AppendCase(cases, switchCase interface{}) ([]*caseDummy, error) {
	cs, ok := cases.([]*caseDummy)
	if !ok {
		return nil, errors.Errorf("invalid switch case list type; expected []*caseDummy, got %T", cases)
	}
	c, ok := switchCase.(*caseDummy)
	if !ok {
		return nil, errors.Errorf("invalid switch case type; expected *irx.caseDummy, got %T", switchCase)
	}
	return append(cs, c), nil
}

// NewCase returns a new switch case based on the given case comparand and
// target branch.
func NewCase(xTyp, xVal, target interface{}) (*caseDummy, error) {
	xValue, err := NewValue(xTyp, xVal)
	if err != nil {
		return nil, errors.WithStack(err)
	}
	x, ok := xValue.(*constant.Int)
	if !ok {
		return nil, errors.Errorf("invalid case comparand type; expected *constant.Int, got %T", xValue)
	}
	t, ok := target.(*LocalIdent)
	if !ok {
		return nil, errors.Errorf("invalid target branch type; expected *irx.LocalIdent, got %T", target)
	}
	return newCaseDummy(x, t.name), nil
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

// getInt64 returns the int64 representation of the given integer literal.
func getInt64(lit interface{}) (int64, error) {
	l, ok := lit.(*IntLit)
	if !ok {
		return 0, errors.Errorf("invalid array length type; expected *IntLit, got %T", lit)
	}
	switch l.lit {
	case "true":
		return 1, nil
	case "false":
		return 0, nil
	}
	n, err := strconv.ParseInt(l.lit, 10, 64)
	if err != nil {
		return 0, errors.WithStack(err)
	}
	return n, nil
}
