// Package astx implements utility functions for generating abstract syntax
// trees of LLVM IR modules.
package astx

import (
	"log"
	"os"
	"strconv"
	"strings"

	"github.com/llir/llvm/asm/internal/ast"
	"github.com/llir/llvm/asm/internal/token"
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

// NewNamedType returns a new named type based on the given local identifier.
func NewNamedType(name interface{}) (*ast.NamedType, error) {
	n, ok := name.(*LocalIdent)
	if !ok {
		return nil, errors.Errorf("invalid type name type; expected *astx.LocalIdent, got %T", name)
	}
	// NOTE: The named type does not yet have a type definition body.
	return &ast.NamedType{Name: n.name}, nil
}

// === [ Values ] ==============================================================

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
