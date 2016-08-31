//go:generate go run gen.go

// TODO: Figure out how to type-check global variable declarations and function
// declarations against their uses.

// TODO: Figure out how to handle forward references of user-defined types.

// NOTE: Both of these issues are related to the fact that the syntax directed
// translation rules do not have access to a context of the module being parsed.
//
// Potential solutions:
//
// * HACK: Global variable used as context.
//
// * Define an AST for the LLVM IR, which may later be traversed by a walker,
// * which has access to a context of the module, to produce the final ir.

// Package irx implements utility functions for generating LLVM IR from Gocc
// parsers.
package irx

import (
	"fmt"
	"log"
	"strconv"
	"strings"

	"github.com/kr/pretty"
	"github.com/llir/llvm/asm/internal/token"
	"github.com/llir/llvm/ir"
	"github.com/llir/llvm/ir/constant"
	"github.com/llir/llvm/ir/instruction"
	"github.com/llir/llvm/ir/types"
	"github.com/llir/llvm/ir/value"
	"github.com/mewkiz/pkg/errutil"
)

// === [ Modules ] =============================================================

// NewModule returns a new module based on the given top-level declarations.
func NewModule(decls interface{}) (*ir.Module, error) {
	if decls, ok := decls.([]TopLevelDecl); ok {
		module := new(ir.Module)
		for _, decl := range decls {
			switch decl := decl.(type) {
			case *TargetLayout:
				module.Layout = decl.layout
			case *TargetTriple:
				module.Triple = decl.triple
			case *ir.Function:
				module.Funcs = append(module.Funcs, decl)
			case *ir.GlobalDecl:
				module.Globals = append(module.Globals, decl)
			default:
				log.Printf("support for top-level declaration type %T not yet implemented", decl)
				//panic(fmt.Sprintf("support for top-level declaration type %T not yet implemented", decl))
			}
		}
		// Replace dummy values with their corresponding local variables.
		return fixModule(module), nil
	}
	return nil, errutil.Newf("invalid top-level declarations type; expected []TopLevelDecl, got %T", decls)
}

// === [ Declarations ] ========================================================

// TODO: Consider refining the TopLevelDecl interface to identify unintentional
// uses; i.e. add isTopLevelDecl and wrapper types for the various top-level
// declaration types?

// TopLevelDecl represents a top-level declaration.
type TopLevelDecl interface{}

// NewTopLevelDeclList returns a new top-level declaration list.
func NewTopLevelDeclList(decl interface{}) ([]TopLevelDecl, error) {
	if decl, ok := decl.(TopLevelDecl); ok {
		return []TopLevelDecl{decl}, nil
	}
	return nil, errutil.Newf("invalid top-level declaration list top-level declaration type; expected TopLevelDecl, got %T", decl)
}

// AppendTopLevelDecl appends decl to the top-level declaration list.
func AppendTopLevelDecl(list, decl interface{}) ([]TopLevelDecl, error) {
	if list, ok := list.([]TopLevelDecl); ok {
		if decl, ok := decl.(TopLevelDecl); ok {
			return append(list, decl), nil
		}
		return nil, errutil.Newf("invalid top-level declaration list top-level declaration type; expected TopLevelDecl, got %T", decl)
	}
	return nil, errutil.Newf("invalid top-level declaration list type; expected []TopLevelDecl, got %T", list)
}

// --- [ Target specifiers ] ---------------------------------------------------

// TargetLayout specifies the data layout of the target.
type TargetLayout struct {
	// Data layout.
	layout string
}

// NewTargetLayout returns a new target data layout specifier based on the given
// target data layout string.
func NewTargetLayout(layoutToken interface{}) (*TargetLayout, error) {
	if layoutToken, ok := layoutToken.(*token.Token); ok {
		layout, err := strconv.Unquote(string(layoutToken.Lit))
		if err != nil {
			return nil, errutil.Err(err)
		}
		return &TargetLayout{layout: layout}, nil
	}
	return nil, errutil.Newf("invalid target data layout type; expected *token.Token, got %T", layoutToken)
}

// TargetTriple specifies the host architecture, operating system and vendor as
// a target triple.
type TargetTriple struct {
	// Target triple.
	triple string
}

// NewTargetTriple returns a new target triple specifier based on the given
// target triple string.
func NewTargetTriple(tripleToken interface{}) (*TargetTriple, error) {
	if tripleToken, ok := tripleToken.(*token.Token); ok {
		triple, err := strconv.Unquote(string(tripleToken.Lit))
		if err != nil {
			return nil, errutil.Err(err)
		}
		return &TargetTriple{triple: triple}, nil
	}
	return nil, errutil.Newf("invalid target triple type; expected *token.Token, got %T", tripleToken)
}

// --- [ Type definitions ] ----------------------------------------------------

// TODO: Add support for user-defined type definitions.

// --- [ Global variable declarations ] ----------------------------------------

// NewGlobalVar returns a new global variable declaration based on the given
// name and global variable.
func NewGlobalVarDecl(gname, globalVar interface{}) (*ir.GlobalDecl, error) {
	if globalVar, ok := globalVar.(*GlobalVar); ok {
		name, err := getGlobalName(gname)
		if err != nil {
			return nil, errutil.Err(err)
		}

		// Global variable declaration.
		if globalVar.val == nil {
			return ir.NewGlobalDecl(name, globalVar.typ, globalVar.immutable)
		}

		// Global variable definition.
		return ir.NewGlobalDef(name, globalVar.val, globalVar.immutable)
	}
	return nil, errutil.Newf("invalid global variable type; expected *GlobalVar, got %T", globalVar)
}

// A GlobalVar represents a global variable.
type GlobalVar struct {
	// Variable type.
	typ types.Type
	// Initial value, or nil if defined externally.
	val value.Value
	// Specifies whether the global variable is immutable.
	immutable bool
}

// NewGlobalVar returns a new global variable based on the given type and
// initial value.
func NewGlobalVar(immutable, typ, val interface{}) (*GlobalVar, error) {
	if immutable, ok := immutable.(bool); ok {
		if typ, ok := typ.(types.Type); ok {
			if val == nil {
				// Global variable declaration.
				return &GlobalVar{typ: typ, immutable: immutable}, nil
			}
			// Global variable definition.
			val, err := NewValue(typ, val)
			if err != nil {
				return nil, errutil.Err(err)
			}
			return &GlobalVar{typ: typ, val: val, immutable: immutable}, nil
		}
		return nil, errutil.Newf("invalid global variable type; expected types.Type, got %T", typ)
	}
	return nil, errutil.Newf("invalid global variable immutability type; expected bool, got %T", immutable)
}

// --- [ Function declarations ] -----------------------------------------------

// NewFuncDecl returns a new function declaration based on the given function
// header.
func NewFuncDecl(fn interface{}) (*ir.Function, error) {
	if fn, ok := fn.(*ir.Function); ok {
		return fn, nil
	}
	return nil, errutil.Newf("invalid function header type; expected *ir.Function, got %T", fn)
}

// NewFuncDef returns a new function definition based on the given function
// header and body.
func NewFuncDef(fn, blocks interface{}) (*ir.Function, error) {
	if fn, ok := fn.(*ir.Function); ok {
		if blocks, ok := blocks.([]*ir.BasicBlock); ok {
			if err := fn.SetBlocks(blocks); err != nil {
				return nil, errutil.Err(err)
			}
			return fn, nil
		}
		return nil, errutil.Newf("invalid function body type; expected []*ir.BasicBlock, got %T", blocks)
	}
	return nil, errutil.Newf("invalid function header type; expected *ir.Function, got %T", fn)
}

// NewFunc returns a new function based on the given result type, function name,
// and function parameters.
func NewFunc(result, gname, params interface{}) (*ir.Function, error) {
	if result, ok := result.(types.Type); ok {
		name, err := getGlobalName(gname)
		if err != nil {
			return nil, errutil.Err(err)
		}
		var sig *types.Func
		switch params := params.(type) {
		case *Params:
			sig, err = types.NewFunc(result, params.params, params.variadic)
			if err != nil {
				return nil, errutil.Err(err)
			}
		case nil:
			sig, err = types.NewFunc(result, nil, false)
			if err != nil {
				return nil, errutil.Err(err)
			}
		default:
			return nil, errutil.Newf("invalid function parameters specifier type; expected *Params, got %T", params)
		}
		return ir.NewFunction(name, sig), nil
	}
	return nil, errutil.Newf("invalid function result type; expected types.Type, got %T", result)
}

// Params represents a function parameters specifier.
type Params struct {
	// Function parameter types.
	params []*types.Param
	// A variadic function takes a variable number of arguments.
	variadic bool
}

// NewParams returns a new function parameters specifier. A variadic function
// takes a variable number of arguments.
func NewParams(params interface{}, variadic bool) (*Params, error) {
	switch params := params.(type) {
	case []*types.Param:
		return &Params{params: params, variadic: variadic}, nil
	case nil:
		return &Params{variadic: variadic}, nil
	default:
		return nil, errutil.Newf("invalid function parameter list; expected []*types.Param, got %T", params)
	}
}

// NewParamList returns a new function parameter list.
func NewParamList(param interface{}) ([]*types.Param, error) {
	if param, ok := param.(*types.Param); ok {
		return []*types.Param{param}, nil
	}
	return nil, errutil.Newf("invalid function parameter list parameter type; expected *types.Param, got %T", param)
}

// AppendParam appends param to the function parameter list.
func AppendParam(list, param interface{}) ([]*types.Param, error) {
	if list, ok := list.([]*types.Param); ok {
		if param, ok := param.(*types.Param); ok {
			return append(list, param), nil
		}
		return nil, errutil.Newf("invalid function parameter list parameter type; expected *types.Param, got %T", param)
	}
	return nil, errutil.Newf("invalid function parameter list type; expected []*types.Param, got %T", list)
}

// NewParam returns a new function parameter based on the given type and
// parameter name.
func NewParam(typ, lname interface{}) (*types.Param, error) {
	if typ, ok := typ.(types.Type); ok {
		name, err := getLocalName(lname)
		if err != nil {
			return nil, errutil.Err(err)
		}
		return types.NewParam(typ, name), nil
	}
	return nil, errutil.Newf("invalid function parameter type; expected types.Type, got %T", typ)
}

// --- [ Attribute group definitions ] -----------------------------------------

// TODO: Add support for function attributes.

// --- [ Metadata definitions ] ------------------------------------------------

// TODO: Add support for metadata.

// === [ Types ] ===============================================================

// --- [ Void type ] -----------------------------------------------------------

// --- [ Label type ] ----------------------------------------------------------

// --- [ Integer type ] --------------------------------------------------------

// NewIntType returns a new integer type based on the given integer type lexeme.
func NewIntType(typToken interface{}) (*types.Int, error) {
	if typToken, ok := typToken.(*token.Token); ok {
		// Skip "i" prefix in integer type (e.g. i32).
		s := string(typToken.Lit)
		if !strings.HasPrefix(s, "i") {
			return nil, errutil.Newf("invalid prefix of integer type %q; expected 'i'", s)
		}
		s = s[1:]
		size, err := strconv.Atoi(s)
		if err != nil {
			return nil, errutil.Err(err)
		}
		return types.NewInt(size)
	}
	return nil, errutil.Newf("invalid integer type; expected *token.Token, got %T", typToken)
}

// --- [ Pointer type ] --------------------------------------------------------

// NewPointerType returns a new pointer type based on the given element type.
func NewPointerType(elem interface{}) (*types.Pointer, error) {
	if elem, ok := elem.(types.Type); ok {
		return types.NewPointer(elem)
	}
	return nil, errutil.Newf("invalid pointer element type; expected types.Type, got %T", elem)
}

// --- [ Array type ] ----------------------------------------------------------

// NewArrayType returns a new array type based on the given length and element
// type.
func NewArrayType(lenToken, elem interface{}) (*types.Array, error) {
	if lenToken, ok := lenToken.(*token.Token); ok {
		n, err := strconv.Atoi(string(lenToken.Lit))
		if err != nil {
			return nil, errutil.Err(err)
		}
		if elem, ok := elem.(types.Type); ok {
			return types.NewArray(elem, n)
		}
		return nil, errutil.Newf("invalid array element type; expected types.Type, got %T", elem)
	}
	return nil, errutil.Newf("invalid array length type; expected *token.Token, got %T", lenToken)
}

// --- [ Structure type ] ------------------------------------------------------

// NewStructType returns a new structure type based on the given field types. A
// packed structure is 1 byte aligned.
func NewStructType(fields interface{}, packed bool) (*types.Struct, error) {
	if fields, ok := fields.([]types.Type); ok {
		return types.NewStruct(fields, packed)
	}
	return nil, errutil.Newf("invalid structure fields type; expected []types.Type, got %T", fields)
}

// NewFieldList returns a new structure field list.
func NewFieldList(typ interface{}) ([]types.Type, error) {
	if typ, ok := typ.(types.Type); ok {
		return []types.Type{typ}, nil
	}
	return nil, errutil.Newf("invalid structure field list field type; expected types.Type, got %T", typ)
}

// AppendField appends typ to the structure field list.
func AppendField(list, typ interface{}) ([]types.Type, error) {
	if list, ok := list.([]types.Type); ok {
		if typ, ok := typ.(types.Type); ok {
			return append(list, typ), nil
		}
		return nil, errutil.Newf("invalid structure field list field type; expected types.Type, got %T", typ)
	}
	return nil, errutil.Newf("invalid structure field list type; expected []types.Type, got %T", list)
}

// --- [ Function type ] -------------------------------------------------------

// NewFuncType returns a new function type based on the given result and
// parameter types.
func NewFuncType(result, params interface{}) (*types.Func, error) {
	if result, ok := result.(types.Type); ok {
		switch params := params.(type) {
		case *Params:
			return types.NewFunc(result, params.params, params.variadic)
		case nil:
			return types.NewFunc(result, nil, false)
		default:
			return nil, errutil.Newf("invalid function parameters specifier type; expected *Params, got %T", params)
		}
	}
	return nil, errutil.Newf("invalid function result type; expected types.Type, got %T", result)
}

// --- [ User-defined type ] ---------------------------------------------------

// TODO: Add support for user-defined types.

// === [ Local identifiers ] ===================================================

// A LocalDummy represents a dummy value of a local variable. Dummy values for
// local variables are used during syntax directed translation, and later
// replaced by their corresponding values in a later stage of parsing. The
// reason for this is that there is no clear way to keep a context for functions
// around during parsing within the Gocc generated parsers.
type LocalDummy struct {
	// Local identifier name.
	name string
}

// Dummy implementation of value.NamedValue

func (l *LocalDummy) String() string {
	panic("dummy implementation")
}

func (l *LocalDummy) ValueString() string {
	panic("dummy implementation")
}

func (l *LocalDummy) Type() types.Type {
	panic("dummy implementation")
}

func (l *LocalDummy) Name() string {
	panic("dummy implementation")
}

// NewLocalDummy returns a new dummy value for the given local identifier name.
func NewLocalDummy(nameToken interface{}) (*LocalDummy, error) {
	if nameToken, ok := nameToken.(*token.Token); ok {
		// Strip "%" prefix.
		name, err := stripLocalPrefix(nameToken.Lit)
		if err != nil {
			return nil, errutil.Err(err)
		}
		return &LocalDummy{name: name}, nil
	}
	return nil, errutil.Newf("invalid local identifier name type; expected *token.Token, got %T", nameToken)
}

// stripLocalPrefix strips the "%" prefix from local identifiers.
func stripLocalPrefix(name []byte) (string, error) {
	s := string(name)
	if !strings.HasPrefix(s, "%") {
		return "", errutil.Newf("invalid prefix of local identifier %q; expected '%'", s)
	}
	// Strip "%" prefix.
	return s[1:], nil
}

// === [ Values ] ==============================================================

// NewValue returns a new value based on the given type and value.
func NewValue(typ, val interface{}) (value.Value, error) {
	// TODO: Verify type equality between typ and val.Type().
	//    * handled by constant.NewInt
	//    * ...
	if typ, ok := typ.(types.Type); ok {
		switch val := val.(type) {
		case *LocalDummy:
			return val, nil
		case *Global:
			// TODO: Retreive type from val.typ once support for forward reference
			// of global variable declarations have been added.
			// TODO: Add type-check between ptrType and val.typ.
			if typ, ok := typ.(*types.Pointer); ok {
				return constant.NewPointer(typ, val.name)
			}
			return nil, errutil.Newf("invalid global type; expected *types.Pointer, got %T", typ)
		case *IntConst:
			return constant.NewInt(typ, val.val)
		case *NullPointerConst:
			if typ, ok := typ.(*types.Pointer); ok {
				return constant.NewNullPointer(typ)
			}
			return nil, errutil.Newf("invalid null pointer type; expected *types.Pointer, got %T", typ)
		case *CharArrayConst:
			return constant.NewCharArray(typ, val.val)
		case *ZeroInitializer:
			return constant.NewZeroInitializer(typ), nil
		case *GetElementPtrExpr:
			return constant.NewGetElementPtr(val.elem, val.addr, val.indices)
		default:
			pretty.Println(val)
			panic(fmt.Sprintf("support for value type %T not yet implemented", val))
		}
	}
	return nil, errutil.Newf("invalid value type; expected types.Type, got %T", typ)
}

// NewValueList returns a new value list.
func NewValueList(val interface{}) ([]value.Value, error) {
	if val, ok := val.(value.Value); ok {
		return []value.Value{val}, nil
	}
	return nil, errutil.Newf("invalid value list value type; expected value.Value, got %T", val)
}

// AppendValue appends val to the value list.
func AppendValue(list, val interface{}) ([]value.Value, error) {
	if list, ok := list.([]value.Value); ok {
		if val, ok := val.(value.Value); ok {
			return append(list, val), nil
		}
		return nil, errutil.Newf("invalid value list value type; expected value.Value, got %T", val)
	}
	return nil, errutil.Newf("invalid value list type; expected []value.Value, got %T", list)
}

// === [ Constants ] ===========================================================

// --- [ Boolean constants ] ---------------------------------------------------

// --- [ Integer constants ] ---------------------------------------------------

// An IntConst represents an integer constant.
type IntConst struct {
	// Integer value.
	val string
}

// NewIntConst returns a new integer constant based on the given value.
func NewIntConst(valToken interface{}) (*IntConst, error) {
	if valToken, ok := valToken.(*token.Token); ok {
		return &IntConst{val: string(valToken.Lit)}, nil
	}
	return nil, errutil.Newf("invalid integer constant type; expected *token.Token, got %T", valToken)
}

// --- [ Null pointer constants ] ----------------------------------------------

// A NullPointerConst represents a null pointer constant.
type NullPointerConst struct{}

// --- [ Array constants ] -----------------------------------------------------

// A CharArrayConst represents a character array constant.
type CharArrayConst struct {
	// Character array value.
	val string
}

// NewCharArrayConst returns a new character array constant based on the given
// value.
func NewCharArrayConst(valToken interface{}) (*CharArrayConst, error) {
	if valToken, ok := valToken.(*token.Token); ok {
		return &CharArrayConst{val: string(valToken.Lit)}, nil
	}
	return nil, errutil.Newf("invalid character array constant type; expected *token.Token, got %T", valToken)
}

// --- [ Zero initializer ] ----------------------------------------------------

// A ZeroInitializer represents a zero initializer.
type ZeroInitializer struct{}

// --- [ Global identifiers ] --------------------------------------------------

// A Global represents a global identifier.
type Global struct {
	// Global identifier name.
	name string
}

// NewGlobal returns a new global identifier based on the given global
// identifier name.
func NewGlobal(nameToken interface{}) (*Global, error) {
	if nameToken, ok := nameToken.(*token.Token); ok {
		// Strip "@" prefix.
		name, err := stripGlobalPrefix(nameToken.Lit)
		if err != nil {
			return nil, errutil.Err(err)
		}
		return &Global{name: name}, nil
	}
	return nil, errutil.Newf("invalid global identifier name type; expected *token.Token, got %T", nameToken)
}

// stripGlobalPrefix strips the "@" prefix from global identifiers.
func stripGlobalPrefix(name []byte) (string, error) {
	s := string(name)
	if !strings.HasPrefix(s, "@") {
		return "", errutil.Newf("invalid prefix of global identifier %q; expected '@'", s)
	}
	// Strip "@" prefix.
	return s[1:], nil
}

// --- [ Constant expressions ] ------------------------------------------------

// ___ [ `getelementptr` expression ] __________________________________________

// A GetElementPtrExpr represents a getelementptr constant expression.
type GetElementPtrExpr struct {
	// Element type.
	elem types.Type
	// Memory address of the element.
	addr value.Value
	// Element indices.
	indices []constant.Constant
}

// NewGetElementPtrExpr returns a new getelementptr constant expression based on
// the given type, address and element indices.
func NewGetElementPtrExpr(elem, addrType, addr, elemIndices interface{}) (*GetElementPtrExpr, error) {
	if elem, ok := elem.(types.Type); ok {
		addr, err := NewValue(addrType, addr)
		if err != nil {
			return nil, errutil.Err(err)
		}
		var indices []constant.Constant
		switch elemIndices := elemIndices.(type) {
		case []constant.Constant:
			indices = elemIndices
		case nil:
		default:
			panic(fmt.Sprintf("support for element indices type %T not yet implemented", elemIndices))
		}
		return &GetElementPtrExpr{elem: elem, addr: addr, indices: indices}, nil
	}
	return nil, errutil.Newf("invalid operand type; expected types.Type, got %T", elem)
}

// === [ Basic blocks ] ========================================================

// NewBasicBlockList returns a new basic block list.
func NewBasicBlockList(block interface{}) ([]*ir.BasicBlock, error) {
	if block, ok := block.(*ir.BasicBlock); ok {
		return []*ir.BasicBlock{block}, nil
	}
	return nil, errutil.Newf("invalid basic block list basic block type; expected *ir.BasicBlock, got %T", block)
}

// AppendBasicBlock appends block to the basic block list.
func AppendBasicBlock(list, block interface{}) ([]*ir.BasicBlock, error) {
	if list, ok := list.([]*ir.BasicBlock); ok {
		if block, ok := block.(*ir.BasicBlock); ok {
			return append(list, block), nil
		}
		return nil, errutil.Newf("invalid basic block list basic block type; expected *ir.BasicBlock, got %T", block)
	}
	return nil, errutil.Newf("invalid basic block list type; expected []*ir.BasicBlock, got %T", list)
}

// NewBasicBlock returns a new basic block based on the given name, non-
// terminating instructions and terminator.
func NewBasicBlock(nameToken, insts, term interface{}) (*ir.BasicBlock, error) {
	// Get label name.
	var name string
	switch nameToken := nameToken.(type) {
	case *token.Token:
		// Strip ":" suffix.
		s, err := stripLabelSuffix(nameToken.Lit)
		if err != nil {
			return nil, errutil.Err(err)
		}
		name = s
	case nil:
		// Unnamed basic block.
	default:
		return nil, errutil.Newf("invalid basic block name type; expected *token.Token, got %T", nameToken)
	}

	if term, ok := term.(instruction.Terminator); ok {
		switch insts := insts.(type) {
		case []instruction.Instruction:
			block := ir.NewBasicBlock(name)
			block.SetInsts(insts)
			block.SetTerm(term)
			return block, nil
		case nil:
			block := ir.NewBasicBlock(name)
			block.SetTerm(term)
			return block, nil
		default:
			return nil, errutil.Newf("invalid non-terminating instructions type; expected []instruction.Instruction, got %T", insts)
		}
	}
	return nil, errutil.Newf("invalid terminator type; expected instruction.Terminator, got %T", term)
}

// stripLabelSuffix strips the ":" suffix from labels.
func stripLabelSuffix(name []byte) (string, error) {
	s := string(name)
	if !strings.HasSuffix(s, ":") {
		return "", errutil.Newf("invalid suffix of label %q; expected ':'", s)
	}
	// Strip ":" suffix.
	return s[:len(s)-1], nil
}

// NewInstructionList returns a new instruction list.
func NewInstructionList(inst interface{}) ([]instruction.Instruction, error) {
	if inst, ok := inst.(instruction.Instruction); ok {
		return []instruction.Instruction{inst}, nil
	}
	return nil, errutil.Newf("invalid instruction list instruction type; expected instruction.Instruction, got %T", inst)
}

// AppendInstruction appends inst to the instruction list.
func AppendInstruction(list, inst interface{}) ([]instruction.Instruction, error) {
	if list, ok := list.([]instruction.Instruction); ok {
		if inst, ok := inst.(instruction.Instruction); ok {
			return append(list, inst), nil
		}
		return nil, errutil.Newf("invalid instruction list instruction type; expected instruction.Instruction, got %T", inst)
	}
	return nil, errutil.Newf("invalid instruction list type; expected []instruction.Instruction, got %T", list)
}

// === [ Instructions ] ========================================================

// --- [ Local variable definition ] -------------------------------------------

// NewLocalVarDef returns a new local variable definition based on the given
// local identifier name and value instruction.
func NewLocalVarDef(lname, valInst interface{}) (*instruction.LocalVarDef, error) {
	if valInst, ok := valInst.(instruction.ValueInst); ok {
		name, err := getLocalName(lname)
		if err != nil {
			return nil, errutil.Err(err)
		}
		return instruction.NewLocalVarDef(name, valInst)
	}
	return nil, errutil.Newf("invalid value instruction type; expected instruction.ValueInst, got %T", valInst)
}

// --- [ Terminator instructions ] ---------------------------------------------

// ___ [ `ret` instruction ] ___________________________________________________

// NewRetInst returns a new return instruction based on the given result type
// and value.
func NewRetInst(typ, val interface{}) (*instruction.Ret, error) {
	if typ, ok := typ.(types.Type); ok {
		val, err := NewValue(typ, val)
		if err != nil {
			return nil, errutil.Err(err)
		}
		return instruction.NewRet(typ, val)
	}
	return nil, errutil.Newf("invalid result type; expected types.Type, got %T", typ)
}

// ___ [ `br` instruction ] ____________________________________________________

// NewJmpInst returns a new unconditional branch instruction based on the given
// target branch.
func NewJmpInst(ltarget interface{}) (*instruction.Jmp, error) {
	target, ok := ltarget.(*LocalDummy)
	if !ok {
		return nil, errutil.Newf("invalid target type; expected *LocalDummy, got %T", ltarget)
	}
	return instruction.NewJmp(target)
}

// NewBrInst returns a new br instruction based on the given branching
// condition, and the true and false target branches.
func NewBrInst(condType, condVal, ltrueBranch, lfalseBranch interface{}) (*instruction.Br, error) {
	cond, err := NewValue(condType, condVal)
	if err != nil {
		return nil, errutil.Err(err)
	}
	falseBranch, ok := lfalseBranch.(*LocalDummy)
	if !ok {
		return nil, errutil.Newf("invalid false branch type; expected *LocalDummy, got %T", lfalseBranch)
	}
	trueBranch, ok := ltrueBranch.(*LocalDummy)
	if !ok {
		return nil, errutil.Newf("invalid true branch type; expected *LocalDummy, got %T", ltrueBranch)
	}
	return instruction.NewBr(cond, trueBranch, falseBranch)
}

// --- [ Memory instructions ] -------------------------------------------------

// ___ [ `alloca` instruction ] ________________________________________________

// NewAllocaInst returns a new allocation instruction based on the given element
// type and number of elements.
func NewAllocaInst(typ, nelems interface{}) (*instruction.Alloca, error) {
	if typ, ok := typ.(types.Type); ok {
		switch nelems := nelems.(type) {
		case *token.Token:
			n, err := strconv.Atoi(string(nelems.Lit))
			if err != nil {
				return nil, errutil.Err(err)
			}
			return instruction.NewAlloca(typ, n)
		case int:
			return instruction.NewAlloca(typ, nelems)
		default:
			panic(fmt.Sprintf("support for number of elements type %T not yet implemented", nelems))
		}
	}
	return nil, errutil.Newf("invalid element type; expected types.Type, got %T", typ)
}

// ___ [ `load` instruction ] __________________________________________________

// NewLoadInst returns a new load instruction based on the given type and
// address.
func NewLoadInst(typ, addrType, addr interface{}) (*instruction.Load, error) {
	if typ, ok := typ.(types.Type); ok {
		addr, err := NewValue(addrType, addr)
		if err != nil {
			return nil, errutil.Err(err)
		}
		return instruction.NewLoad(typ, addr)
	}
	return nil, errutil.Newf("invalid operand type; expected types.Type, got %T", typ)
}

// ___ [ `store` instruction ] _________________________________________________

// NewStoreInst returns a new store instruction based on the given value and
// address.
func NewStoreInst(typ, val, addrType, addr interface{}) (*instruction.Store, error) {
	{
		val, err := NewValue(typ, val)
		if err != nil {
			return nil, errutil.Err(err)
		}
		if addrType, ok := addrType.(types.Type); ok {
			addr, err := NewValue(addrType, addr)
			if err != nil {
				return nil, errutil.Err(err)
			}
			return instruction.NewStore(val, addr)
		}
	}
	return nil, errutil.Newf("invalid pointer type; expected types.Type, got %T", addrType)
}

// ___ [ `getelementptr` instruction ] _________________________________________

// NewGetElementPtrInst returns a new getelementptr instruction based on the
// given element type, address and element indices.
func NewGetElementPtrInst(elem, addrType, addr, elemIndices interface{}) (*instruction.GetElementPtr, error) {
	if elem, ok := elem.(types.Type); ok {
		addr, err := NewValue(addrType, addr)
		if err != nil {
			return nil, errutil.Err(err)
		}
		var indices []value.Value
		switch elemIndices := elemIndices.(type) {
		case []value.Value:
			indices = elemIndices
		case nil:
		default:
			panic(fmt.Sprintf("support for element indices type %T not yet implemented", elemIndices))
		}
		return instruction.NewGetElementPtr(elem, addr, indices)
	}
	return nil, errutil.Newf("invalid operand type; expected types.Type, got %T", elem)
}

// --- [ Conversion instructions ] ---------------------------------------------

// ___ [ `trunc` instruction ] _________________________________________________

// NewTruncInst returns a new trunc instruction based on the given value and
// target type.
func NewTruncInst(fromType, fromVal, to interface{}) (*instruction.Trunc, error) {
	from, err := NewValue(fromType, fromVal)
	if err != nil {
		return nil, errutil.Err(err)
	}
	if to, ok := to.(types.Type); ok {
		return instruction.NewTrunc(from, to)
	}
	return nil, errutil.Newf("invalid target type; expected types.Type, got %T", to)
}

// ___ [ `zext` instruction ] __________________________________________________

// NewZExtInst returns a new zext instruction based on the given value and
// target type.
func NewZExtInst(fromType, fromVal, to interface{}) (*instruction.ZExt, error) {
	from, err := NewValue(fromType, fromVal)
	if err != nil {
		return nil, errutil.Err(err)
	}
	if to, ok := to.(types.Type); ok {
		return instruction.NewZExt(from, to)
	}
	return nil, errutil.Newf("invalid target type; expected types.Type, got %T", to)
}

// ___ [ `sext` instruction ] __________________________________________________

// NewSExtInst returns a new sext instruction based on the given value and
// target type.
func NewSExtInst(fromType, fromVal, to interface{}) (*instruction.SExt, error) {
	from, err := NewValue(fromType, fromVal)
	if err != nil {
		return nil, errutil.Err(err)
	}
	if to, ok := to.(types.Type); ok {
		return instruction.NewSExt(from, to)
	}
	return nil, errutil.Newf("invalid target type; expected types.Type, got %T", to)
}

// ___ [ `fptrunc` instruction ] _______________________________________________

// NewFPTruncInst returns a new fptrunc instruction based on the given value and
// target type.
func NewFPTruncInst(fromType, fromVal, to interface{}) (*instruction.FPTrunc, error) {
	from, err := NewValue(fromType, fromVal)
	if err != nil {
		return nil, errutil.Err(err)
	}
	if to, ok := to.(types.Type); ok {
		return instruction.NewFPTrunc(from, to)
	}
	return nil, errutil.Newf("invalid target type; expected types.Type, got %T", to)
}

// ___ [ `fpext` instruction ] _________________________________________________

// NewFPExtInst returns a new fpext instruction based on the given value and
// target type.
func NewFPExtInst(fromType, fromVal, to interface{}) (*instruction.FPExt, error) {
	from, err := NewValue(fromType, fromVal)
	if err != nil {
		return nil, errutil.Err(err)
	}
	if to, ok := to.(types.Type); ok {
		return instruction.NewFPExt(from, to)
	}
	return nil, errutil.Newf("invalid target type; expected types.Type, got %T", to)
}

// ___ [ `fptoui` instruction ] ________________________________________________

// NewFPToUIInst returns a new fptoui instruction based on the given value and
// target type.
func NewFPToUIInst(fromType, fromVal, to interface{}) (*instruction.FPToUI, error) {
	from, err := NewValue(fromType, fromVal)
	if err != nil {
		return nil, errutil.Err(err)
	}
	if to, ok := to.(types.Type); ok {
		return instruction.NewFPToUI(from, to)
	}
	return nil, errutil.Newf("invalid target type; expected types.Type, got %T", to)
}

// ___ [ `fptosi` instruction ] ________________________________________________

// NewFPToSIInst returns a new fptosi instruction based on the given value and
// target type.
func NewFPToSIInst(fromType, fromVal, to interface{}) (*instruction.FPToSI, error) {
	from, err := NewValue(fromType, fromVal)
	if err != nil {
		return nil, errutil.Err(err)
	}
	if to, ok := to.(types.Type); ok {
		return instruction.NewFPToSI(from, to)
	}
	return nil, errutil.Newf("invalid target type; expected types.Type, got %T", to)
}

// ___ [ `uitofp` instruction ] ________________________________________________

// NewUIToFPInst returns a new uitofp instruction based on the given value and
// target type.
func NewUIToFPInst(fromType, fromVal, to interface{}) (*instruction.UIToFP, error) {
	from, err := NewValue(fromType, fromVal)
	if err != nil {
		return nil, errutil.Err(err)
	}
	if to, ok := to.(types.Type); ok {
		return instruction.NewUIToFP(from, to)
	}
	return nil, errutil.Newf("invalid target type; expected types.Type, got %T", to)
}

// ___ [ `sitofp` instruction ] ________________________________________________

// NewSIToFPInst returns a new sitofp instruction based on the given value and
// target type.
func NewSIToFPInst(fromType, fromVal, to interface{}) (*instruction.SIToFP, error) {
	from, err := NewValue(fromType, fromVal)
	if err != nil {
		return nil, errutil.Err(err)
	}
	if to, ok := to.(types.Type); ok {
		return instruction.NewSIToFP(from, to)
	}
	return nil, errutil.Newf("invalid target type; expected types.Type, got %T", to)
}

// ___ [ `ptrtoint` instruction ] ______________________________________________

// NewPtrToIntInst returns a new ptrtoint instruction based on the given value
// and target type.
func NewPtrToIntInst(fromType, fromVal, to interface{}) (*instruction.PtrToInt, error) {
	from, err := NewValue(fromType, fromVal)
	if err != nil {
		return nil, errutil.Err(err)
	}
	if to, ok := to.(types.Type); ok {
		return instruction.NewPtrToInt(from, to)
	}
	return nil, errutil.Newf("invalid target type; expected types.Type, got %T", to)
}

// ___ [ `inttoptr` instruction ] ______________________________________________

// NewIntToPtrInst returns a new inttoptr instruction based on the given value
// and target type.
func NewIntToPtrInst(fromType, fromVal, to interface{}) (*instruction.IntToPtr, error) {
	from, err := NewValue(fromType, fromVal)
	if err != nil {
		return nil, errutil.Err(err)
	}
	if to, ok := to.(types.Type); ok {
		return instruction.NewIntToPtr(from, to)
	}
	return nil, errutil.Newf("invalid target type; expected types.Type, got %T", to)
}

// ___ [ `bitcast` instruction ] _______________________________________________

// NewBitCastInst returns a new bitcast instruction based on the given value and
// target type.
func NewBitCastInst(fromType, fromVal, to interface{}) (*instruction.BitCast, error) {
	from, err := NewValue(fromType, fromVal)
	if err != nil {
		return nil, errutil.Err(err)
	}
	if to, ok := to.(types.Type); ok {
		return instruction.NewBitCast(from, to)
	}
	return nil, errutil.Newf("invalid target type; expected types.Type, got %T", to)
}

// ___ [ `addrspacecast` instruction ] _________________________________________

// NewAddrSpaceCastInst returns a new addrspacecast instruction based on the
// given value and target type.
func NewAddrSpaceCastInst(fromType, fromVal, to interface{}) (*instruction.AddrSpaceCast, error) {
	from, err := NewValue(fromType, fromVal)
	if err != nil {
		return nil, errutil.Err(err)
	}
	if to, ok := to.(types.Type); ok {
		return instruction.NewAddrSpaceCast(from, to)
	}
	return nil, errutil.Newf("invalid target type; expected types.Type, got %T", to)
}

// --- [ Other instructions ] --------------------------------------------------

// ___ [ `icmp` instruction ] __________________________________________________

// NewICmpInst returns a new icmp instruction based on the given condition and
// operands.
func NewICmpInst(cond, typ, x, y interface{}) (*instruction.ICmp, error) {
	if cond, ok := cond.(instruction.ICond); ok {
		x, err := NewValue(typ, x)
		if err != nil {
			return nil, errutil.Err(err)
		}
		y, err := NewValue(typ, y)
		if err != nil {
			return nil, errutil.Err(err)
		}
		return instruction.NewICmp(cond, x, y)
	}
	return nil, errutil.Newf("invalid condition type; expected instruction.ICond, got %T", cond)
}

// ___ [ `phi` instruction ] ___________________________________________________

// NewPHIInst returns a new phi instruction based on the given type and incoming
// values.
func NewPHIInst(typ, incList interface{}) (*instruction.PHI, error) {
	if typ, ok := typ.(types.Type); ok {
		if incList, ok := incList.([]*Incoming); ok {
			var incs []*instruction.Incoming
			for _, inc := range incList {
				val, err := NewValue(typ, inc.val)
				if err != nil {
					return nil, errutil.Err(err)
				}
				inc, err := instruction.NewIncoming(val, inc.pred)
				if err != nil {
					return nil, errutil.Err(err)
				}
				incs = append(incs, inc)
			}
			return instruction.NewPHI(typ, incs)
		}
		return nil, errutil.Newf("invalid incoming values type; expected []*Incoming, got %T", incList)
	}
	return nil, errutil.Newf("invalid operand type; expected types.Type, got %T", typ)
}

// NewIncomingList returns a new incoming values list.
func NewIncomingList(inc interface{}) ([]*Incoming, error) {
	if inc, ok := inc.(*Incoming); ok {
		return []*Incoming{inc}, nil
	}
	return nil, errutil.Newf("invalid incoming values list incoming value type; expected *Incoming, got %T", inc)
}

// AppendIncoming appends inc to the incoming values list.
func AppendIncoming(list, inc interface{}) ([]*Incoming, error) {
	if list, ok := list.([]*Incoming); ok {
		if inc, ok := inc.(*Incoming); ok {
			return append(list, inc), nil
		}
		return nil, errutil.Newf("invalid incoming values list incoming value type; expected *Incoming, got %T", inc)
	}
	return nil, errutil.Newf("invalid incoming values list type; expected []*Incoming, got %T", list)
}

// Incoming represents an incoming value from a predecessor basic block, as
// specified by PHI instructions.
type Incoming struct {
	// Incoming value.
	val interface{}
	// Predecessor basic block of the incoming value.
	pred *LocalDummy
}

// NewIncoming returns a new incoming value based on the given value and
// predecessor basic block label name.
func NewIncoming(val, lpred interface{}) (*Incoming, error) {
	pred, ok := lpred.(*LocalDummy)
	if !ok {
		return nil, errutil.Newf("invalid predecessor type; expected *LocalDummy, got %T", lpred)
	}
	return &Incoming{val: val, pred: pred}, nil
}

// ___ [ `select` instruction ] ________________________________________________

// NewSelectInst returns a new select instruction based on the given selection
// condition, and operands.
func NewSelectInst(condType, condVal, xType, xVal, yType, yVal interface{}) (*instruction.Select, error) {
	cond, err := NewValue(condType, condVal)
	if err != nil {
		return nil, errutil.Err(err)
	}
	x, err := NewValue(xType, xVal)
	if err != nil {
		return nil, errutil.Err(err)
	}
	y, err := NewValue(yType, yVal)
	if err != nil {
		return nil, errutil.Err(err)
	}
	return instruction.NewSelect(cond, x, y)
}

// ___ [ `call` instruction ] __________________________________________________

// NewCallInst returns a new call instruction based on the given function return
// type, function name and function arguments.
func NewCallInst(result, gname, args interface{}) (*instruction.Call, error) {
	name, err := getGlobalName(gname)
	if err != nil {
		return nil, errutil.Err(err)
	}
	if result, ok := result.(types.Type); ok {
		switch args := args.(type) {
		case []value.Value:
			return instruction.NewCall(result, name, args)
		case nil:
			return instruction.NewCall(result, name, nil)
		default:
			return nil, errutil.Newf("invalid function arguments type; expected []value.Value, got %T", args)
		}
	}
	return nil, errutil.Newf("invalid function return type; expected types.Type, got %T", result)
}

// getGlobalName returns the name of the given global identifier.
func getGlobalName(gname interface{}) (string, error) {
	switch gname := gname.(type) {
	case *Global:
		return gname.name, nil
	default:
		return "", errutil.Newf("invalid global identifier name type; expected *Global, got %T", gname)
	}
}

// getLocalName returns the name of the given local identifier, or empty string
// if nil is provided.
func getLocalName(lname interface{}) (string, error) {
	switch lname := lname.(type) {
	case *LocalDummy:
		return lname.name, nil
	case nil:
		return "", nil
	default:
		return "", errutil.Newf("invalid local identifier name type; expected *LocalDummy, got %T", lname)
	}
}
