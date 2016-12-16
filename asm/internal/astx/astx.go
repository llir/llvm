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
		t, ok := t.(*ast.IntType)
		if !ok {
			return nil, errors.Errorf("invalid integer constant type; expected *ast.IntType, got %T", t)
		}
		return &ast.IntConst{Type: t, Lit: val.lit}, nil
	case *FloatLit:
		t, ok := t.(*ast.FloatType)
		if !ok {
			return nil, errors.Errorf("invalid floating-point constant type; expected *ast.FloatType, got %T", t)
		}
		return &ast.FloatConst{Type: t, Lit: val.lit}, nil
	case *NullLit:
		t, ok := t.(*ast.PointerType)
		if !ok {
			return nil, errors.Errorf("invalid null constant type; expected *ast.PointerType, got %T", t)
		}
		return &ast.NullConst{Type: t}, nil
	// TODO: Add dummy/dummyconstant package for Vector, Array and Struct,
	// ExprGetElementPtr, so that typ may be stored and evaluated after type
	// resolution.
	case *ZeroInitializerLit:
		return &ast.ZeroInitializerConst{Type: t}, nil
	case ast.Constant:
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
