package ir

import (
	"fmt"
	"strings"

	"github.com/llir/l/internal/enc"
	"github.com/llir/l/ir/ll"
	"github.com/llir/l/ir/types"
)

func (*BasicBlock) IsUnwindTarget() {}

// --- [ Function parameters ] -------------------------------------------------

// Param is an LLVM IR function parameter.
type Param struct {
	// Parameter type.
	Typ types.Type
	// (optional) Parameter name.
	ParamName string

	// extra.

	// (optional) Parameter attributes.
	Attrs []ll.ParamAttribute
}

// NewParam returns a new function parameter based on the given type and name.
func NewParam(typ types.Type, name string) *Param {
	return &Param{Typ: typ, ParamName: name}
}

// String returns the LLVM syntax representation of the function parameter as a
// type-value pair.
func (p *Param) String() string {
	return fmt.Sprintf("%v %v", p.Type(), p.Ident())
}

// Type returns the type of the function parameter.
func (p *Param) Type() types.Type {
	return p.Typ
}

// Ident returns the identifier associated with the function parameter.
func (p *Param) Ident() string {
	return enc.Local(p.ParamName)
}

// Name returns the name of the function parameter.
func (p *Param) Name() string {
	return p.ParamName
}

// SetName sets the name of the function parameter.
func (p *Param) SetName(name string) {
	p.ParamName = name
}

// ### [ Helper functions ] ####################################################

// isUnnamed reports whether the given identifier is unnamed.
func isUnnamed(name string) bool {
	return len(name) == 0
}

// isLocalID reports whether the given identifier is a local ID (e.g. "%42").
func isLocalID(name string) bool {
	for _, r := range name {
		if !strings.ContainsRune("0123456789", r) {
			return false
		}
	}
	return len(name) > 0
}

// quote returns s as a double-quoted string literal.
func quote(s string) string {
	return enc.Quote([]byte(s))
}

// unquote interprets s as a double-quoted string literal, returning the string
// value that s quotes.
func unquote(s string) string {
	return string(enc.Unquote(s))
}
