// Package astx implements utility functions for generating abstract syntax
// trees of LLVM IR modules.
package astx

import (
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"

	"github.com/llir/llvm/asm/internal/ast"
	"github.com/llir/llvm/asm/internal/token"
	"github.com/llir/llvm/internal/enc"
	"github.com/pkg/errors"
)

// TODO: Remove debug output.

// dbg is a logger which prefixes debug messages with the file name and line
// number of callees.
var dbg = log.New(os.Stdout, "", log.Lshortfile)

// === [ Modules ] =============================================================

// NewModule returns a new module based on the given top-level declarations.
func NewModule(decls interface{}) (*ast.Module, error) {
	var ds []TopLevelDecl
	switch decls := decls.(type) {
	case []TopLevelDecl:
		ds = decls
	case nil:
		// no top-level declarations.
	default:
		return nil, errors.Errorf("invalid top-level declaration list type; expected []astx.TopLevelDecl, got %T", decls)
	}
	m := &ast.Module{}
	for _, d := range ds {
		switch d := d.(type) {
		case *ast.NamedType:
			m.Types = append(m.Types, d)
		case *ast.Global:
			m.Globals = append(m.Globals, d)
		case *ast.Function:
			m.Funcs = append(m.Funcs, d)
		default:
			dbg.Printf("support for %T not yet implemented", d)
		}
	}
	return m, nil
}

// TopLevelDecl represents a top-level declaration.
type TopLevelDecl interface{}

// NewTopLevelDeclList returns a new top-level declaration list based on the
// given top-level declaration.
func NewTopLevelDeclList(decl interface{}) ([]TopLevelDecl, error) {
	d, ok := decl.(TopLevelDecl)
	if !ok {
		return nil, errors.Errorf("invalid top-level declaration type; expected astx.TopLevelDecl, got %T", decl)
	}
	return []TopLevelDecl{d}, nil
}

// AppendTopLevelDecl appends the given top-level declaration to the top-level
// declaration list.
func AppendTopLevelDecl(decls, decl interface{}) ([]TopLevelDecl, error) {
	ds, ok := decls.([]TopLevelDecl)
	if !ok {
		return nil, errors.Errorf("invalid top-level declaration list type; expected []astx.TopLevelDecl, got %T", decls)
	}
	d, ok := decl.(TopLevelDecl)
	if !ok {
		return nil, errors.Errorf("invalid top-level declaration type; expected astx.TopLevelDecl, got %T", decl)
	}
	return append(ds, d), nil
}

// === [ Type definitions ] ====================================================

// NewTypeDef returns a new type definition based on the given type name and
// definition.
func NewTypeDef(name, typ interface{}) (*ast.NamedType, error) {
	n, ok := name.(*LocalIdent)
	if !ok {
		return nil, errors.Errorf("invalid type name type; expected *astx.LocalIdent, got %T", name)
	}
	var t ast.Type
	switch typ := typ.(type) {
	case ast.Type:
		t = typ
	case nil:
		// opaque identified struct type.
	default:
		return nil, errors.Errorf("invalid type; expected ast.Type, got %T", typ)
	}
	return &ast.NamedType{Name: n.name, Def: t}, nil
}

// === [ Global variables ] ====================================================

// NewGlobalDecl returns a new global variable declaration based on the given
// global variable name, immutability and type.
func NewGlobalDecl(name, immutable, typ interface{}) (*ast.Global, error) {
	n, ok := name.(*GlobalIdent)
	if !ok {
		return nil, errors.Errorf("invalid global name type; expected *astx.GlobalIdent, got %T", name)
	}
	imm, ok := immutable.(bool)
	if !ok {
		return nil, errors.Errorf("invalid immutability type; expected bool, got %T", immutable)
	}
	t, ok := typ.(ast.Type)
	if !ok {
		return nil, errors.Errorf("invalid content type; expected ast.Type, got %T", typ)
	}
	global := &ast.Global{Name: n.name, Content: t}
	global.Immutable = imm
	return global, nil
}

// NewGlobalDef returns a new global variable definition based on the given
// global variable name, immutability, type and value.
func NewGlobalDef(name, immutable, typ, val interface{}) (*ast.Global, error) {
	n, ok := name.(*GlobalIdent)
	if !ok {
		return nil, errors.Errorf("invalid global name type; expected *astx.GlobalIdent, got %T", name)
	}
	imm, ok := immutable.(bool)
	if !ok {
		return nil, errors.Errorf("invalid immutability type; expected bool, got %T", immutable)
	}
	t, ok := typ.(ast.Type)
	if !ok {
		return nil, errors.Errorf("invalid type; expected ast.Type, got %T", typ)
	}
	init, err := NewValue(t, val)
	if err != nil {
		return nil, errors.WithStack(err)
	}
	i, ok := init.(ast.Constant)
	if !ok {
		return nil, errors.Errorf("invalid init type; expected ast.Constant, got %T", init)
	}
	global := &ast.Global{Name: n.name, Content: t, Init: i}
	global.Immutable = imm
	return global, nil
}

// === [ Functions ] ===========================================================

// NewFunctionDecl returns a new function declaration based on the given
// return type, function name and parameters.
func NewFunctionDecl(ret, name, params interface{}) (*ast.Function, error) {
	r, ok := ret.(ast.Type)
	if !ok {
		return nil, errors.Errorf("invalid function return type; expected ast.Type, got %T", ret)
	}
	n, ok := name.(*GlobalIdent)
	if !ok {
		return nil, errors.Errorf("invalid function name type; expected *astx.GlobalIdent, got %T", name)
	}
	sig := &ast.FuncType{Ret: r}
	switch ps := params.(type) {
	case *Params:
		for _, param := range ps.params {
			sig.Params = append(sig.Params, param)
		}
		sig.Variadic = ps.variadic
	case nil:
		// no parameters.
	default:
		return nil, errors.Errorf("invalid function parameters type; expected *astx.Params or nil, got %T", params)
	}
	f := &ast.Function{
		Name: n.name,
		Sig:  sig,
	}
	return f, nil
}

// NewFunctionDef returns a new function definition based on the given function
// header and body.
func NewFunctionDef(header, body interface{}) (*ast.Function, error) {
	f, ok := header.(*ast.Function)
	if !ok {
		return nil, errors.Errorf("invalid function header type; expected *ast.Function, got %T", header)
	}
	blocks, ok := body.([]*ast.BasicBlock)
	if !ok {
		return nil, errors.Errorf("invalid function body type; expected []*ast.BasicBlock, got %T", body)
	}
	f.Blocks = blocks
	return f, nil
}

// Params represents a function parameters specifier.
type Params struct {
	// Function parameter types.
	params []*ast.Param
	// Variadicity of the function type.
	variadic bool
}

// NewParams returns a new function parameters specifier, based on the given
// function parameters and variadicity.
func NewParams(params interface{}, variadic bool) (*Params, error) {
	switch params := params.(type) {
	case []*ast.Param:
		return &Params{params: params, variadic: variadic}, nil
	case nil:
		return &Params{variadic: variadic}, nil
	default:
		return nil, errors.Errorf("invalid function parameter list; expected []*ast.Param or nil, got %T", params)
	}
}

// NewParamList returns a new function parameter list based on the given
// function parameter.
func NewParamList(param interface{}) ([]*ast.Param, error) {
	p, ok := param.(*ast.Param)
	if !ok {
		return nil, errors.Errorf("invalid function parameter type; expected *ast.Param, got %T", param)
	}
	return []*ast.Param{p}, nil
}

// AppendParam appends the given parameter to the function parameter list.
func AppendParam(params, param interface{}) ([]*ast.Param, error) {
	ps, ok := params.([]*ast.Param)
	if !ok {
		return nil, errors.Errorf("invalid function parameter list type; expected []*ast.Param, got %T", params)
	}
	p, ok := param.(*ast.Param)
	if !ok {
		return nil, errors.Errorf("invalid function parameter type; expected *ast.Param, got %T", param)
	}
	return append(ps, p), nil
}

// NewParam returns a new function parameter based on the given parameter type
// and name.
func NewParam(typ, name interface{}) (*ast.Param, error) {
	t, ok := typ.(ast.Type)
	if !ok {
		return nil, errors.Errorf("invalid type; expected ast.Type, got %T", typ)
	}
	var n string
	switch name := name.(type) {
	case *LocalIdent:
		n = name.name
	case nil:
		// unnamed function parameter.
	default:
		return nil, errors.Errorf("invalid local name type; expected *astx.LocalIdent or nil, got %T", name)
	}
	return &ast.Param{Name: n, Type: t}, nil
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

// NewTypeList returns a new type list based on the given type.
func NewTypeList(typ interface{}) ([]ast.Type, error) {
	t, ok := typ.(ast.Type)
	if !ok {
		return nil, errors.Errorf("invalid type; expected ast.Type, got %T", typ)
	}
	return []ast.Type{t}, nil
}

// AppendType appends the given type to the type list.
func AppendType(typs, typ interface{}) ([]ast.Type, error) {
	ts, ok := typs.([]ast.Type)
	if !ok {
		return nil, errors.Errorf("invalid type list type; expected []ast.Type, got %T", typs)
	}
	t, ok := typ.(ast.Type)
	if !ok {
		return nil, errors.Errorf("invalid type; expected ast.Type, got %T", typ)
	}
	return append(ts, t), nil
}

// NewIntType returns a new integer type based on the given integer type token.
func NewIntType(typeTok interface{}) (*ast.IntType, error) {
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
	return &ast.IntType{Size: size}, nil
}

// NewFuncType returns a new function type based on the given return type and
// function parameters.
func NewFuncType(ret, params interface{}) (*ast.FuncType, error) {
	r, ok := ret.(ast.Type)
	if !ok {
		return nil, errors.Errorf("invalid function return type; expected ast.Type, got %T", ret)
	}
	sig := &ast.FuncType{Ret: r}
	switch ps := params.(type) {
	case *Params:
		for _, param := range ps.params {
			sig.Params = append(sig.Params, param)
		}
		sig.Variadic = ps.variadic
	case nil:
		// no parameters.
	default:
		return nil, errors.Errorf("invalid function parameters type; expected *astx.Params or nil, got %T", params)
	}
	return sig, nil
}

// NewPointerType returns a new pointer type based on the given element type.
func NewPointerType(elem, space interface{}) (*ast.PointerType, error) {
	e, ok := elem.(ast.Type)
	if !ok {
		return nil, errors.Errorf("invalid element type; expected ast.Type, got %T", elem)
	}
	t := &ast.PointerType{Elem: e}
	switch space := space.(type) {
	case *AddrSpace:
		t.Space = space.space
	case nil:
		// no address space.
	default:
		return nil, errors.Errorf("invalid address space type; expected *astx.AddrSpace or nil, got %T", space)
	}
	return t, nil
}

// AddrSpace represents the address space of a pointer type.
type AddrSpace struct {
	// Address space.
	space int64
}

// NewAddrSpace returns a new address space pointer based on the given space.
func NewAddrSpace(space interface{}) (*AddrSpace, error) {
	s, err := getInt64(space)
	if err != nil {
		return nil, errors.WithStack(err)
	}
	return &AddrSpace{space: s}, nil
}

// NewVectorType returns a new vector type based on the given vector length and
// element type.
func NewVectorType(len, elem interface{}) (*ast.VectorType, error) {
	e, ok := elem.(ast.Type)
	if !ok {
		return nil, errors.Errorf("invalid element type; expected ast.Type, got %T", elem)
	}
	l, err := getInt64(len)
	if err != nil {
		return nil, errors.WithStack(err)
	}
	return &ast.VectorType{Elem: e, Len: l}, nil
}

// NewArrayType returns a new array type based on the given array length and
// element type.
func NewArrayType(len, elem interface{}) (*ast.ArrayType, error) {
	e, ok := elem.(ast.Type)
	if !ok {
		return nil, errors.Errorf("invalid element type; expected ast.Type, got %T", elem)
	}
	l, err := getInt64(len)
	if err != nil {
		return nil, errors.WithStack(err)
	}
	return &ast.ArrayType{Elem: e, Len: l}, nil
}

// NewStructType returns a new struct type based on the given struct fields.
func NewStructType(fields interface{}) (*ast.StructType, error) {
	var fs []ast.Type
	switch fields := fields.(type) {
	case []ast.Type:
		fs = fields
	case nil:
		// no struct fields.
	default:
		return nil, errors.Errorf("invalid struct fields type; expected []ast.Type, got %T", fields)
	}
	return &ast.StructType{Fields: fs}, nil
}

// NewTypeIdent returns a new type identifier based on the given local
// identifier.
func NewTypeIdent(name interface{}) (*ast.NamedTypeDummy, error) {
	n, ok := name.(*LocalIdent)
	if !ok {
		return nil, errors.Errorf("invalid type name type; expected *astx.LocalIdent, got %T", name)
	}
	return &ast.NamedTypeDummy{Name: n.name}, nil
}

// === [ Values ] ==============================================================

// NewValueList returns a new value list based on the given
// value.
func NewValueList(val interface{}) ([]ast.Value, error) {
	v, ok := val.(ast.Value)
	if !ok {
		return nil, errors.Errorf("invalid value type; expected ast.Value, got %T", val)
	}
	return []ast.Value{v}, nil
}

// AppendValue appends the given value to the value list.
func AppendValue(vals, val interface{}) ([]ast.Value, error) {
	vs, ok := vals.([]ast.Value)
	if !ok {
		return nil, errors.Errorf("invalid value list type; expected []ast.Value, got %T", vals)
	}
	v, ok := val.(ast.Value)
	if !ok {
		return nil, errors.Errorf("invalid value type; expected ast.Value, got %T", val)
	}
	return append(vs, v), nil
}

// NewValue returns a value based on the given type and value.
func NewValue(typ, val interface{}) (ast.Value, error) {
	t, ok := typ.(ast.Type)
	if !ok {
		return nil, errors.Errorf("invalid value type; expected ast.Type, got %T", typ)
	}
	switch val := val.(type) {
	case *LocalIdent:
		return &ast.LocalDummy{Name: val.name, Type: t}, nil
	case *GlobalIdent:
		return &ast.GlobalDummy{Name: val.name, Type: t}, nil
	case *IntLit:
		return &ast.IntConst{Type: t, Lit: val.lit}, nil
	case *FloatLit:
		return &ast.FloatConst{Type: t, Lit: val.lit}, nil
	case *NullLit:
		return &ast.NullConst{Type: t}, nil
	case *ZeroInitializerLit:
		return &ast.ZeroInitializerConst{Type: t}, nil

	// Store type of vector, array and struct constants and constant expressions,
	// so that it may be evaluated after type resolution.
	case *ast.VectorConst:
		// Vector constant type should not be known at this stage of parsing, as
		// they've been constructed from VectorConst literals.
		if val.Type != nil {
			return nil, errors.Errorf("invalid vector constant type, expected nil, got %T", val.Type)
		}
		val.Type = t
		return val, nil
	case *ast.ArrayConst:
		// Array constant type should not be known at this stage of parsing, as
		// they've been constructed from ArrayConst literals.
		if val.Type != nil {
			return nil, errors.Errorf("invalid array constant type, expected nil, got %T", val.Type)
		}
		val.Type = t
		return val, nil
	case *ast.StructConst:
		// Struct constant type should not be known at this stage of parsing, as
		// they've been constructed from StructConst literals.
		if val.Type != nil {
			return nil, errors.Errorf("invalid struct constant type, expected nil, got %T", val.Type)
		}
		val.Type = t
		return val, nil

	// Binary instructions
	case *ast.ExprAdd:
		// Constant expression type should not be known at this stage of parsing,
		// as they've been constructed from ExprAdd production rules.
		if val.Type != nil {
			return nil, errors.Errorf("invalid add constant expression type, expected nil, got %T", val.Type)
		}
		val.Type = t
		return val, nil
	case *ast.ExprFAdd:
		// Constant expression type should not be known at this stage of parsing,
		// as they've been constructed from ExprFAdd production rules.
		if val.Type != nil {
			return nil, errors.Errorf("invalid fadd constant expression type, expected nil, got %T", val.Type)
		}
		val.Type = t
		return val, nil
	case *ast.ExprSub:
		// Constant expression type should not be known at this stage of parsing,
		// as they've been constructed from ExprSub production rules.
		if val.Type != nil {
			return nil, errors.Errorf("invalid sub constant expression type, expected nil, got %T", val.Type)
		}
		val.Type = t
		return val, nil
	case *ast.ExprFSub:
		// Constant expression type should not be known at this stage of parsing,
		// as they've been constructed from ExprFSub production rules.
		if val.Type != nil {
			return nil, errors.Errorf("invalid fsub constant expression type, expected nil, got %T", val.Type)
		}
		val.Type = t
		return val, nil
	case *ast.ExprMul:
		// Constant expression type should not be known at this stage of parsing,
		// as they've been constructed from ExprMul production rules.
		if val.Type != nil {
			return nil, errors.Errorf("invalid mul constant expression type, expected nil, got %T", val.Type)
		}
		val.Type = t
		return val, nil
	case *ast.ExprFMul:
		// Constant expression type should not be known at this stage of parsing,
		// as they've been constructed from ExprFMul production rules.
		if val.Type != nil {
			return nil, errors.Errorf("invalid fmul constant expression type, expected nil, got %T", val.Type)
		}
		val.Type = t
		return val, nil
	case *ast.ExprUDiv:
		// Constant expression type should not be known at this stage of parsing,
		// as they've been constructed from ExprUDiv production rules.
		if val.Type != nil {
			return nil, errors.Errorf("invalid udiv constant expression type, expected nil, got %T", val.Type)
		}
		val.Type = t
		return val, nil
	case *ast.ExprSDiv:
		// Constant expression type should not be known at this stage of parsing,
		// as they've been constructed from ExprSDiv production rules.
		if val.Type != nil {
			return nil, errors.Errorf("invalid sdiv constant expression type, expected nil, got %T", val.Type)
		}
		val.Type = t
		return val, nil
	case *ast.ExprFDiv:
		// Constant expression type should not be known at this stage of parsing,
		// as they've been constructed from ExprFDiv production rules.
		if val.Type != nil {
			return nil, errors.Errorf("invalid fdiv constant expression type, expected nil, got %T", val.Type)
		}
		val.Type = t
		return val, nil
	case *ast.ExprURem:
		// Constant expression type should not be known at this stage of parsing,
		// as they've been constructed from ExprURem production rules.
		if val.Type != nil {
			return nil, errors.Errorf("invalid urem constant expression type, expected nil, got %T", val.Type)
		}
		val.Type = t
		return val, nil
	case *ast.ExprSRem:
		// Constant expression type should not be known at this stage of parsing,
		// as they've been constructed from ExprSRem production rules.
		if val.Type != nil {
			return nil, errors.Errorf("invalid srem constant expression type, expected nil, got %T", val.Type)
		}
		val.Type = t
		return val, nil
	case *ast.ExprFRem:
		// Constant expression type should not be known at this stage of parsing,
		// as they've been constructed from ExprFRem production rules.
		if val.Type != nil {
			return nil, errors.Errorf("invalid frem constant expression type, expected nil, got %T", val.Type)
		}
		val.Type = t
		return val, nil

	// Bitwise instructions
	case *ast.ExprShl:
		// Constant expression type should not be known at this stage of parsing,
		// as they've been constructed from ExprShl production rules.
		if val.Type != nil {
			return nil, errors.Errorf("invalid shl constant expression type, expected nil, got %T", val.Type)
		}
		val.Type = t
		return val, nil
	case *ast.ExprLShr:
		// Constant expression type should not be known at this stage of parsing,
		// as they've been constructed from ExprLShr production rules.
		if val.Type != nil {
			return nil, errors.Errorf("invalid lshr constant expression type, expected nil, got %T", val.Type)
		}
		val.Type = t
		return val, nil
	case *ast.ExprAShr:
		// Constant expression type should not be known at this stage of parsing,
		// as they've been constructed from ExprAShr production rules.
		if val.Type != nil {
			return nil, errors.Errorf("invalid ashr constant expression type, expected nil, got %T", val.Type)
		}
		val.Type = t
		return val, nil
	case *ast.ExprAnd:
		// Constant expression type should not be known at this stage of parsing,
		// as they've been constructed from ExprAnd production rules.
		if val.Type != nil {
			return nil, errors.Errorf("invalid and constant expression type, expected nil, got %T", val.Type)
		}
		val.Type = t
		return val, nil
	case *ast.ExprOr:
		// Constant expression type should not be known at this stage of parsing,
		// as they've been constructed from ExprOr production rules.
		if val.Type != nil {
			return nil, errors.Errorf("invalid or constant expression type, expected nil, got %T", val.Type)
		}
		val.Type = t
		return val, nil
	case *ast.ExprXor:
		// Constant expression type should not be known at this stage of parsing,
		// as they've been constructed from ExprXor production rules.
		if val.Type != nil {
			return nil, errors.Errorf("invalid xor constant expression type, expected nil, got %T", val.Type)
		}
		val.Type = t
		return val, nil

	// Memory instructions
	case *ast.ExprGetElementPtr:
		// Constant expression type should not be known at this stage of parsing,
		// as they've been constructed from ExprGetElementPtr production rules.
		if val.Type != nil {
			return nil, errors.Errorf("invalid getelementptr constant expression type, expected nil, got %T", val.Type)
		}
		val.Type = t
		return val, nil

	// Conversion instructions
	case *ast.ExprTrunc:
		// Constant expression type should not be known at this stage of parsing,
		// as they've been constructed from ExprTrunc production rules.
		if val.Type != nil {
			return nil, errors.Errorf("invalid trunc constant expression type, expected nil, got %T", val.Type)
		}
		val.Type = t
		return val, nil
	case *ast.ExprZExt:
		// Constant expression type should not be known at this stage of parsing,
		// as they've been constructed from ExprZExt production rules.
		if val.Type != nil {
			return nil, errors.Errorf("invalid zext constant expression type, expected nil, got %T", val.Type)
		}
		val.Type = t
		return val, nil
	case *ast.ExprSExt:
		// Constant expression type should not be known at this stage of parsing,
		// as they've been constructed from ExprSExt production rules.
		if val.Type != nil {
			return nil, errors.Errorf("invalid sext constant expression type, expected nil, got %T", val.Type)
		}
		val.Type = t
		return val, nil
	case *ast.ExprFPTrunc:
		// Constant expression type should not be known at this stage of parsing,
		// as they've been constructed from ExprFPTrunc production rules.
		if val.Type != nil {
			return nil, errors.Errorf("invalid fptrunc constant expression type, expected nil, got %T", val.Type)
		}
		val.Type = t
		return val, nil
	case *ast.ExprFPExt:
		// Constant expression type should not be known at this stage of parsing,
		// as they've been constructed from ExprFPExt production rules.
		if val.Type != nil {
			return nil, errors.Errorf("invalid fpext constant expression type, expected nil, got %T", val.Type)
		}
		val.Type = t
		return val, nil
	case *ast.ExprFPToUI:
		// Constant expression type should not be known at this stage of parsing,
		// as they've been constructed from ExprFPToUI production rules.
		if val.Type != nil {
			return nil, errors.Errorf("invalid fptoui constant expression type, expected nil, got %T", val.Type)
		}
		val.Type = t
		return val, nil
	case *ast.ExprFPToSI:
		// Constant expression type should not be known at this stage of parsing,
		// as they've been constructed from ExprFPToSI production rules.
		if val.Type != nil {
			return nil, errors.Errorf("invalid fptosi constant expression type, expected nil, got %T", val.Type)
		}
		val.Type = t
		return val, nil
	case *ast.ExprUIToFP:
		// Constant expression type should not be known at this stage of parsing,
		// as they've been constructed from ExprUIToFP production rules.
		if val.Type != nil {
			return nil, errors.Errorf("invalid uitofp constant expression type, expected nil, got %T", val.Type)
		}
		val.Type = t
		return val, nil
	case *ast.ExprSIToFP:
		// Constant expression type should not be known at this stage of parsing,
		// as they've been constructed from ExprSIToFP production rules.
		if val.Type != nil {
			return nil, errors.Errorf("invalid sitofp constant expression type, expected nil, got %T", val.Type)
		}
		val.Type = t
		return val, nil
	case *ast.ExprPtrToInt:
		// Constant expression type should not be known at this stage of parsing,
		// as they've been constructed from ExprPtrToInt production rules.
		if val.Type != nil {
			return nil, errors.Errorf("invalid ptrtoint constant expression type, expected nil, got %T", val.Type)
		}
		val.Type = t
		return val, nil
	case *ast.ExprIntToPtr:
		// Constant expression type should not be known at this stage of parsing,
		// as they've been constructed from ExprIntToPtr production rules.
		if val.Type != nil {
			return nil, errors.Errorf("invalid inttoptr constant expression type, expected nil, got %T", val.Type)
		}
		val.Type = t
		return val, nil
	case *ast.ExprBitCast:
		// Constant expression type should not be known at this stage of parsing,
		// as they've been constructed from ExprBitCast production rules.
		if val.Type != nil {
			return nil, errors.Errorf("invalid bitcast constant expression type, expected nil, got %T", val.Type)
		}
		val.Type = t
		return val, nil
	case *ast.ExprAddrSpaceCast:
		// Constant expression type should not be known at this stage of parsing,
		// as they've been constructed from ExprAddrSpaceCast production rules.
		if val.Type != nil {
			return nil, errors.Errorf("invalid addrspacecast constant expression type, expected nil, got %T", val.Type)
		}
		val.Type = t
		return val, nil

	// Other instructions
	case *ast.ExprICmp:
		// Constant expression type should not be known at this stage of parsing,
		// as they've been constructed from ExprICmp production rules.
		if val.Type != nil {
			return nil, errors.Errorf("invalid icmp constant expression type, expected nil, got %T", val.Type)
		}
		val.Type = t
		return val, nil
	case *ast.ExprFCmp:
		// Constant expression type should not be known at this stage of parsing,
		// as they've been constructed from ExprFCmp production rules.
		if val.Type != nil {
			return nil, errors.Errorf("invalid fcmp constant expression type, expected nil, got %T", val.Type)
		}
		val.Type = t
		return val, nil
	case *ast.ExprSelect:
		// Constant expression type should not be known at this stage of parsing,
		// as they've been constructed from ExprSelect production rules.
		if val.Type != nil {
			return nil, errors.Errorf("invalid select constant expression type, expected nil, got %T", val.Type)
		}
		val.Type = t
		return val, nil
	default:
		panic(fmt.Sprintf("support for value type %T not yet implemented", val))
	}
}

// === [ Constants ] ===========================================================

// NewConstantList returns a new constant list based on the given constant.
func NewConstantList(x interface{}) ([]ast.Constant, error) {
	c, ok := x.(ast.Constant)
	if !ok {
		return nil, errors.Errorf("invalid constant type; expected ast.Constant, got %T", x)
	}
	return []ast.Constant{c}, nil
}

// AppendConstant appends the given constant to the constant list.
func AppendConstant(xs, x interface{}) ([]ast.Constant, error) {
	cs, ok := xs.([]ast.Constant)
	if !ok {
		return nil, errors.Errorf("invalid constant list type; expected []ast.Constant, got %T", xs)
	}
	c, ok := x.(ast.Constant)
	if !ok {
		return nil, errors.Errorf("invalid constant type; expected ast.Constant, got %T", x)
	}
	return append(cs, c), nil
}

// NewConstant returns a constant based on the given type and value.
func NewConstant(typ, val interface{}) (ast.Constant, error) {
	v, err := NewValue(typ, val)
	if err != nil {
		return nil, errors.WithStack(err)
	}
	c, ok := v.(ast.Constant)
	if !ok {
		return nil, errors.Errorf("invalid constant type; expected ast.Constant, got %T", v)
	}
	return c, nil
}

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

// NewVectorConst returns a new vector constant based on the given elements.
func NewVectorConst(elems interface{}) (*ast.VectorConst, error) {
	es, ok := elems.([]ast.Constant)
	if !ok {
		return nil, errors.Errorf("invalid vector elements type; expected []ast.Constant, got %T", elems)
	}
	return &ast.VectorConst{Elems: es}, nil
}

// NewArrayConst returns a new array constant based on the given elements.
func NewArrayConst(elems interface{}) (*ast.ArrayConst, error) {
	es, ok := elems.([]ast.Constant)
	if !ok {
		return nil, errors.Errorf("invalid array elements type; expected []ast.Constant, got %T", elems)
	}
	return &ast.ArrayConst{Elems: es}, nil
}

// NewCharArrayConst returns a new character array constant based on the given
// string.
func NewCharArrayConst(str interface{}) (*ast.ArrayConst, error) {
	s, err := getTokenString(str)
	if err != nil {
		return nil, errors.WithStack(err)
	}
	// Skip double-quotes.
	s = s[1 : len(s)-1]
	s = enc.Unescape(s)
	var elems []ast.Constant
	for i := 0; i < len(s); i++ {
		// TODO: Validate that string(s[i]) works for the entire byte range 0-255.
		// Otherwise, use *big.Int to implement integer constants in package ast.
		elem := &ast.IntConst{Type: &ast.IntType{Size: 8}, Lit: string(s[i])}
		elems = append(elems, elem)
	}
	c := &ast.ArrayConst{Elems: elems}
	c.CharArray = true
	return c, nil
}

// NewStructConst returns a new struct constant based on the given fields.
func NewStructConst(fields interface{}) (*ast.StructConst, error) {
	fs, ok := fields.([]ast.Constant)
	if !ok {
		return nil, errors.Errorf("invalid struct fields type; expected []ast.Constant, got %T", fields)
	}
	return &ast.StructConst{Fields: fs}, nil
}

// ZeroInitializerLit represents a zeroinitializer literal.
type ZeroInitializerLit struct {
}

// --- [ Binary expressions ] --------------------------------------------------

// NewAddExpr returns a new add expression based on the given type and operands.
func NewAddExpr(xTyp, xVal, yTyp, yVal interface{}) (*ast.ExprAdd, error) {
	x, err := NewConstant(xTyp, xVal)
	if err != nil {
		return nil, errors.WithStack(err)
	}
	y, err := NewConstant(yTyp, yVal)
	if err != nil {
		return nil, errors.WithStack(err)
	}
	return &ast.ExprAdd{X: x, Y: y}, nil
}

// NewFAddExpr returns a new fadd expression based on the given type and
// operands.
func NewFAddExpr(xTyp, xVal, yTyp, yVal interface{}) (*ast.ExprFAdd, error) {
	x, err := NewConstant(xTyp, xVal)
	if err != nil {
		return nil, errors.WithStack(err)
	}
	y, err := NewConstant(yTyp, yVal)
	if err != nil {
		return nil, errors.WithStack(err)
	}
	return &ast.ExprFAdd{X: x, Y: y}, nil
}

// NewSubExpr returns a new sub expression based on the given type and operands.
func NewSubExpr(xTyp, xVal, yTyp, yVal interface{}) (*ast.ExprSub, error) {
	x, err := NewConstant(xTyp, xVal)
	if err != nil {
		return nil, errors.WithStack(err)
	}
	y, err := NewConstant(yTyp, yVal)
	if err != nil {
		return nil, errors.WithStack(err)
	}
	return &ast.ExprSub{X: x, Y: y}, nil
}

// NewFSubExpr returns a new fsub expression based on the given type and
// operands.
func NewFSubExpr(xTyp, xVal, yTyp, yVal interface{}) (*ast.ExprFSub, error) {
	x, err := NewConstant(xTyp, xVal)
	if err != nil {
		return nil, errors.WithStack(err)
	}
	y, err := NewConstant(yTyp, yVal)
	if err != nil {
		return nil, errors.WithStack(err)
	}
	return &ast.ExprFSub{X: x, Y: y}, nil
}

// NewMulExpr returns a new mul expression based on the given type and operands.
func NewMulExpr(xTyp, xVal, yTyp, yVal interface{}) (*ast.ExprMul, error) {
	x, err := NewConstant(xTyp, xVal)
	if err != nil {
		return nil, errors.WithStack(err)
	}
	y, err := NewConstant(yTyp, yVal)
	if err != nil {
		return nil, errors.WithStack(err)
	}
	return &ast.ExprMul{X: x, Y: y}, nil
}

// NewFMulExpr returns a new fmul expression based on the given type and
// operands.
func NewFMulExpr(xTyp, xVal, yTyp, yVal interface{}) (*ast.ExprFMul, error) {
	x, err := NewConstant(xTyp, xVal)
	if err != nil {
		return nil, errors.WithStack(err)
	}
	y, err := NewConstant(yTyp, yVal)
	if err != nil {
		return nil, errors.WithStack(err)
	}
	return &ast.ExprFMul{X: x, Y: y}, nil
}

// NewUDivExpr returns a new udiv expression based on the given type and
// operands.
func NewUDivExpr(xTyp, xVal, yTyp, yVal interface{}) (*ast.ExprUDiv, error) {
	x, err := NewConstant(xTyp, xVal)
	if err != nil {
		return nil, errors.WithStack(err)
	}
	y, err := NewConstant(yTyp, yVal)
	if err != nil {
		return nil, errors.WithStack(err)
	}
	return &ast.ExprUDiv{X: x, Y: y}, nil
}

// NewSDivExpr returns a new sdiv expression based on the given type and
// operands.
func NewSDivExpr(xTyp, xVal, yTyp, yVal interface{}) (*ast.ExprSDiv, error) {
	x, err := NewConstant(xTyp, xVal)
	if err != nil {
		return nil, errors.WithStack(err)
	}
	y, err := NewConstant(yTyp, yVal)
	if err != nil {
		return nil, errors.WithStack(err)
	}
	return &ast.ExprSDiv{X: x, Y: y}, nil
}

// NewFDivExpr returns a new fdiv expression based on the given type and
// operands.
func NewFDivExpr(xTyp, xVal, yTyp, yVal interface{}) (*ast.ExprFDiv, error) {
	x, err := NewConstant(xTyp, xVal)
	if err != nil {
		return nil, errors.WithStack(err)
	}
	y, err := NewConstant(yTyp, yVal)
	if err != nil {
		return nil, errors.WithStack(err)
	}
	return &ast.ExprFDiv{X: x, Y: y}, nil
}

// NewURemExpr returns a new urem expression based on the given type and
// operands.
func NewURemExpr(xTyp, xVal, yTyp, yVal interface{}) (*ast.ExprURem, error) {
	x, err := NewConstant(xTyp, xVal)
	if err != nil {
		return nil, errors.WithStack(err)
	}
	y, err := NewConstant(yTyp, yVal)
	if err != nil {
		return nil, errors.WithStack(err)
	}
	return &ast.ExprURem{X: x, Y: y}, nil
}

// NewSRemExpr returns a new srem expression based on the given type and
// operands.
func NewSRemExpr(xTyp, xVal, yTyp, yVal interface{}) (*ast.ExprSRem, error) {
	x, err := NewConstant(xTyp, xVal)
	if err != nil {
		return nil, errors.WithStack(err)
	}
	y, err := NewConstant(yTyp, yVal)
	if err != nil {
		return nil, errors.WithStack(err)
	}
	return &ast.ExprSRem{X: x, Y: y}, nil
}

// NewFRemExpr returns a new frem expression based on the given type and
// operands.
func NewFRemExpr(xTyp, xVal, yTyp, yVal interface{}) (*ast.ExprFRem, error) {
	x, err := NewConstant(xTyp, xVal)
	if err != nil {
		return nil, errors.WithStack(err)
	}
	y, err := NewConstant(yTyp, yVal)
	if err != nil {
		return nil, errors.WithStack(err)
	}
	return &ast.ExprFRem{X: x, Y: y}, nil
}

// --- [ Bitwise expressions ] -------------------------------------------------

// NewShlExpr returns a new shl expression based on the given type and operands.
func NewShlExpr(xTyp, xVal, yTyp, yVal interface{}) (*ast.ExprShl, error) {
	x, err := NewConstant(xTyp, xVal)
	if err != nil {
		return nil, errors.WithStack(err)
	}
	y, err := NewConstant(yTyp, yVal)
	if err != nil {
		return nil, errors.WithStack(err)
	}
	return &ast.ExprShl{X: x, Y: y}, nil
}

// NewLShrExpr returns a new lshr expression based on the given type and
// operands.
func NewLShrExpr(xTyp, xVal, yTyp, yVal interface{}) (*ast.ExprLShr, error) {
	x, err := NewConstant(xTyp, xVal)
	if err != nil {
		return nil, errors.WithStack(err)
	}
	y, err := NewConstant(yTyp, yVal)
	if err != nil {
		return nil, errors.WithStack(err)
	}
	return &ast.ExprLShr{X: x, Y: y}, nil
}

// NewAShrExpr returns a new ashr expression based on the given type and
// operands.
func NewAShrExpr(xTyp, xVal, yTyp, yVal interface{}) (*ast.ExprAShr, error) {
	x, err := NewConstant(xTyp, xVal)
	if err != nil {
		return nil, errors.WithStack(err)
	}
	y, err := NewConstant(yTyp, yVal)
	if err != nil {
		return nil, errors.WithStack(err)
	}
	return &ast.ExprAShr{X: x, Y: y}, nil
}

// NewAndExpr returns a new and expression based on the given type and operands.
func NewAndExpr(xTyp, xVal, yTyp, yVal interface{}) (*ast.ExprAnd, error) {
	x, err := NewConstant(xTyp, xVal)
	if err != nil {
		return nil, errors.WithStack(err)
	}
	y, err := NewConstant(yTyp, yVal)
	if err != nil {
		return nil, errors.WithStack(err)
	}
	return &ast.ExprAnd{X: x, Y: y}, nil
}

// NewOrExpr returns a new or expression based on the given type and operands.
func NewOrExpr(xTyp, xVal, yTyp, yVal interface{}) (*ast.ExprOr, error) {
	x, err := NewConstant(xTyp, xVal)
	if err != nil {
		return nil, errors.WithStack(err)
	}
	y, err := NewConstant(yTyp, yVal)
	if err != nil {
		return nil, errors.WithStack(err)
	}
	return &ast.ExprOr{X: x, Y: y}, nil
}

// NewXorExpr returns a new xor expression based on the given type and operands.
func NewXorExpr(xTyp, xVal, yTyp, yVal interface{}) (*ast.ExprXor, error) {
	x, err := NewConstant(xTyp, xVal)
	if err != nil {
		return nil, errors.WithStack(err)
	}
	y, err := NewConstant(yTyp, yVal)
	if err != nil {
		return nil, errors.WithStack(err)
	}
	return &ast.ExprXor{X: x, Y: y}, nil
}

// --- [ Memory expressions ] --------------------------------------------------

// NewGetElementPtrExpr returns a new getelementptr expression based on the
// given element type, source address type and value, and element indices.
func NewGetElementPtrExpr(elem, srcTyp, srcVal, indices interface{}) (*ast.ExprGetElementPtr, error) {
	e, ok := elem.(ast.Type)
	if !ok {
		return nil, errors.Errorf("invalid element type; expected ast.Type, got %T", elem)
	}
	src, err := NewConstant(srcTyp, srcVal)
	if err != nil {
		return nil, errors.WithStack(err)
	}
	var is []ast.Constant
	switch indices := indices.(type) {
	case []ast.Constant:
		is = indices
	case nil:
		// no indices.
	default:
		return nil, errors.Errorf("invalid indices type; expected []ast.Constant or nil, got %T", indices)
	}
	return &ast.ExprGetElementPtr{Elem: e, Src: src, Indices: is}, nil
}

// --- [ Conversion expressions ] ----------------------------------------------

// NewTruncExpr returns a new trunc expression based on the given source value
// and target type.
func NewTruncExpr(fromTyp, fromVal, to interface{}) (*ast.ExprTrunc, error) {
	from, err := NewConstant(fromTyp, fromVal)
	if err != nil {
		return nil, errors.WithStack(err)
	}
	t, ok := to.(ast.Type)
	if !ok {
		return nil, errors.Errorf("invalid type; expected ast.Type, got %T", to)
	}
	return &ast.ExprTrunc{From: from, To: t}, nil
}

// NewZExtExpr returns a new zext expression based on the given source value and
// target type.
func NewZExtExpr(fromTyp, fromVal, to interface{}) (*ast.ExprZExt, error) {
	from, err := NewConstant(fromTyp, fromVal)
	if err != nil {
		return nil, errors.WithStack(err)
	}
	t, ok := to.(ast.Type)
	if !ok {
		return nil, errors.Errorf("invalid type; expected ast.Type, got %T", to)
	}
	return &ast.ExprZExt{From: from, To: t}, nil
}

// NewSExtExpr returns a new sext expression based on the given source value and
// target type.
func NewSExtExpr(fromTyp, fromVal, to interface{}) (*ast.ExprSExt, error) {
	from, err := NewConstant(fromTyp, fromVal)
	if err != nil {
		return nil, errors.WithStack(err)
	}
	t, ok := to.(ast.Type)
	if !ok {
		return nil, errors.Errorf("invalid type; expected ast.Type, got %T", to)
	}
	return &ast.ExprSExt{From: from, To: t}, nil
}

// NewFPTruncExpr returns a new fptrunc expression based on the given source
// value and target type.
func NewFPTruncExpr(fromTyp, fromVal, to interface{}) (*ast.ExprFPTrunc, error) {
	from, err := NewConstant(fromTyp, fromVal)
	if err != nil {
		return nil, errors.WithStack(err)
	}
	t, ok := to.(ast.Type)
	if !ok {
		return nil, errors.Errorf("invalid type; expected ast.Type, got %T", to)
	}
	return &ast.ExprFPTrunc{From: from, To: t}, nil
}

// NewFPExtExpr returns a new fpext expression based on the given source value
// and target type.
func NewFPExtExpr(fromTyp, fromVal, to interface{}) (*ast.ExprFPExt, error) {
	from, err := NewConstant(fromTyp, fromVal)
	if err != nil {
		return nil, errors.WithStack(err)
	}
	t, ok := to.(ast.Type)
	if !ok {
		return nil, errors.Errorf("invalid type; expected ast.Type, got %T", to)
	}
	return &ast.ExprFPExt{From: from, To: t}, nil
}

// NewFPToUIExpr returns a new fptoui expression based on the given source value
// and target type.
func NewFPToUIExpr(fromTyp, fromVal, to interface{}) (*ast.ExprFPToUI, error) {
	from, err := NewConstant(fromTyp, fromVal)
	if err != nil {
		return nil, errors.WithStack(err)
	}
	t, ok := to.(ast.Type)
	if !ok {
		return nil, errors.Errorf("invalid type; expected ast.Type, got %T", to)
	}
	return &ast.ExprFPToUI{From: from, To: t}, nil
}

// NewFPToSIExpr returns a new fptosi expression based on the given source value
// and target type.
func NewFPToSIExpr(fromTyp, fromVal, to interface{}) (*ast.ExprFPToSI, error) {
	from, err := NewConstant(fromTyp, fromVal)
	if err != nil {
		return nil, errors.WithStack(err)
	}
	t, ok := to.(ast.Type)
	if !ok {
		return nil, errors.Errorf("invalid type; expected ast.Type, got %T", to)
	}
	return &ast.ExprFPToSI{From: from, To: t}, nil
}

// NewUIToFPExpr returns a new uitofp expression based on the given source value
// and target type.
func NewUIToFPExpr(fromTyp, fromVal, to interface{}) (*ast.ExprUIToFP, error) {
	from, err := NewConstant(fromTyp, fromVal)
	if err != nil {
		return nil, errors.WithStack(err)
	}
	t, ok := to.(ast.Type)
	if !ok {
		return nil, errors.Errorf("invalid type; expected ast.Type, got %T", to)
	}
	return &ast.ExprUIToFP{From: from, To: t}, nil
}

// NewSIToFPExpr returns a new sitofp expression based on the given source value
// and target type.
func NewSIToFPExpr(fromTyp, fromVal, to interface{}) (*ast.ExprSIToFP, error) {
	from, err := NewConstant(fromTyp, fromVal)
	if err != nil {
		return nil, errors.WithStack(err)
	}
	t, ok := to.(ast.Type)
	if !ok {
		return nil, errors.Errorf("invalid type; expected ast.Type, got %T", to)
	}
	return &ast.ExprSIToFP{From: from, To: t}, nil
}

// NewPtrToIntExpr returns a new ptrtoint expression based on the given source
// value and target type.
func NewPtrToIntExpr(fromTyp, fromVal, to interface{}) (*ast.ExprPtrToInt, error) {
	from, err := NewConstant(fromTyp, fromVal)
	if err != nil {
		return nil, errors.WithStack(err)
	}
	t, ok := to.(ast.Type)
	if !ok {
		return nil, errors.Errorf("invalid type; expected ast.Type, got %T", to)
	}
	return &ast.ExprPtrToInt{From: from, To: t}, nil
}

// NewIntToPtrExpr returns a new inttoptr expression based on the given source
// value and target type.
func NewIntToPtrExpr(fromTyp, fromVal, to interface{}) (*ast.ExprIntToPtr, error) {
	from, err := NewConstant(fromTyp, fromVal)
	if err != nil {
		return nil, errors.WithStack(err)
	}
	t, ok := to.(ast.Type)
	if !ok {
		return nil, errors.Errorf("invalid type; expected ast.Type, got %T", to)
	}
	return &ast.ExprIntToPtr{From: from, To: t}, nil
}

// NewBitCastExpr returns a new bitcast expression based on the given source
// value and target type.
func NewBitCastExpr(fromTyp, fromVal, to interface{}) (*ast.ExprBitCast, error) {
	from, err := NewConstant(fromTyp, fromVal)
	if err != nil {
		return nil, errors.WithStack(err)
	}
	t, ok := to.(ast.Type)
	if !ok {
		return nil, errors.Errorf("invalid type; expected ast.Type, got %T", to)
	}
	return &ast.ExprBitCast{From: from, To: t}, nil
}

// NewAddrSpaceCastExpr returns a new addrspacecast expression based on the
// given source value and target type.
func NewAddrSpaceCastExpr(fromTyp, fromVal, to interface{}) (*ast.ExprAddrSpaceCast, error) {
	from, err := NewConstant(fromTyp, fromVal)
	if err != nil {
		return nil, errors.WithStack(err)
	}
	t, ok := to.(ast.Type)
	if !ok {
		return nil, errors.Errorf("invalid type; expected ast.Type, got %T", to)
	}
	return &ast.ExprAddrSpaceCast{From: from, To: t}, nil
}

// --- [ Other expressions ] ---------------------------------------------------

// NewICmpExpr returns a new icmp expression based on the given integer
// condition code, type and operands.
func NewICmpExpr(cond, xTyp, xVal, yTyp, yVal interface{}) (*ast.ExprICmp, error) {
	c, ok := cond.(ast.IntPred)
	if !ok {
		return nil, errors.Errorf("invalid integer predicate type; expected ast.IntPred, got %T", cond)
	}
	x, err := NewConstant(xTyp, xVal)
	if err != nil {
		return nil, errors.WithStack(err)
	}
	y, err := NewConstant(yTyp, yVal)
	if err != nil {
		return nil, errors.WithStack(err)
	}
	return &ast.ExprICmp{Cond: c, X: x, Y: y}, nil
}

// NewFCmpExpr returns a new fcmp expression based on the given floating-point
// condition code, type and operands.
func NewFCmpExpr(cond, xTyp, xVal, yTyp, yVal interface{}) (*ast.ExprFCmp, error) {
	c, ok := cond.(ast.FloatPred)
	if !ok {
		return nil, errors.Errorf("invalid floating-point predicate type; expected ast.FloatPred, got %T", cond)
	}
	x, err := NewConstant(xTyp, xVal)
	if err != nil {
		return nil, errors.WithStack(err)
	}
	y, err := NewConstant(yTyp, yVal)
	if err != nil {
		return nil, errors.WithStack(err)
	}
	return &ast.ExprFCmp{Cond: c, X: x, Y: y}, nil
}

// NewSeExpr returns a new select expression based on the given selection
// condition type and value, and operands.
func NewSelectExpr(condTyp, condVal, xTyp, xVal, yTyp, yVal interface{}) (*ast.ExprSelect, error) {
	cond, err := NewConstant(condTyp, condVal)
	if err != nil {
		return nil, errors.WithStack(err)
	}
	x, err := NewConstant(xTyp, xVal)
	if err != nil {
		return nil, errors.WithStack(err)
	}
	y, err := NewConstant(yTyp, yVal)
	if err != nil {
		return nil, errors.WithStack(err)
	}
	return &ast.ExprSelect{Cond: cond, X: x, Y: y}, nil
}

// === [ Basic blocks ] ========================================================

// NewBasicBlockList returns a new basic block list based on the given basic
// block.
func NewBasicBlockList(block interface{}) ([]*ast.BasicBlock, error) {
	b, ok := block.(*ast.BasicBlock)
	if !ok {
		return nil, errors.Errorf("invalid basic block type; expected *ast.BasicBlock, got %T", block)
	}
	return []*ast.BasicBlock{b}, nil
}

// AppendBasicBlock appends the given basic block to the basic block list.
func AppendBasicBlock(blocks, block interface{}) ([]*ast.BasicBlock, error) {
	bs, ok := blocks.([]*ast.BasicBlock)
	if !ok {
		return nil, errors.Errorf("invalid basic block list type; expected []*ast.BasicBlock, got %T", blocks)
	}
	b, ok := block.(*ast.BasicBlock)
	if !ok {
		return nil, errors.Errorf("invalid basic block type; expected *ast.BasicBlock, got %T", block)
	}
	return append(bs, b), nil
}

// NewBasicBlock returns a new basic block based on the given label name, non-
// branching instructions and terminator.
func NewBasicBlock(name, insts, term interface{}) (*ast.BasicBlock, error) {
	block := &ast.BasicBlock{}
	switch name := name.(type) {
	case *LabelIdent:
		block.Name = name.name
	case nil:
		// unnamed basic block.
	default:
		return nil, errors.Errorf("invalid label name type; expected *astx.LabelIdent or nil, got %T", name)
	}
	var is []ast.Instruction
	switch insts := insts.(type) {
	case []ast.Instruction:
		is = insts
	case nil:
		// no instructions.
	default:
		return nil, errors.Errorf("invalid instruction list type; expected []ast.Instruction, got %T", insts)
	}
	t, ok := term.(ast.Terminator)
	if !ok {
		return nil, errors.Errorf("invalid terminator type; expected ast.Terminator, got %T", term)
	}
	block.Insts = is
	block.Term = t
	return block, nil
}

// === [ Instructions ] ========================================================

// NewInstructionList returns a new instruction list based on the given
// instruction.
func NewInstructionList(inst interface{}) ([]ast.Instruction, error) {
	i, ok := inst.(ast.Instruction)
	if !ok {
		return nil, errors.Errorf("invalid instruction type; expected ast.Instruction, got %T", inst)
	}
	return []ast.Instruction{i}, nil
}

// AppendInstruction appends the given instruction to the instruction list.
func AppendInstruction(insts, inst interface{}) ([]ast.Instruction, error) {
	is, ok := insts.([]ast.Instruction)
	if !ok {
		return nil, errors.Errorf("invalid instruction list type; expected []ast.Instruction, got %T", insts)
	}
	i, ok := inst.(ast.Instruction)
	if !ok {
		return nil, errors.Errorf("invalid instruction type; expected ast.Instruction, got %T", inst)
	}
	return append(is, i), nil
}

// === [ Terminators ] =========================================================

// --- [ ret ] -----------------------------------------------------------------

// NewRetTerm returns a new ret terminator based on the given return type and
// value.
func NewRetTerm(xTyp, xVal interface{}) (*ast.TermRet, error) {
	x, err := NewValue(xTyp, xVal)
	if err != nil {
		return nil, errors.WithStack(err)
	}
	return &ast.TermRet{X: x}, nil
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
		return 0, errors.Errorf("invalid array length type; expected *astx.IntLit, got %T", lit)
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
