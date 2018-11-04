package ir

import (
	"fmt"
	"strings"

	"github.com/llir/llvm/internal/enc"
	"github.com/llir/llvm/ir/types"
)

// TODO: move to the right place.

// TODO: figure out definition of arg.
type Arg interface {
	String() string
	isArg()
}

type ExceptionScope interface {
	isExceptionScope()
}

// TODO: add proper implementations.
type FuncAttribute interface {
	isFuncAttribute()
}

type ParamAttribute interface {
	isParamAttribute()
}

type ReturnAttribute interface {
	isReturnAttribute()
}

type OperandBundle struct {
	// TODO: implement body.
}

type Clause struct {
}

// TODO: consider getting rid of UnwindTarget, and let unwind targets be of type
// *ir.BasicBlock, where a nil value indicates the caller, and a non-nil value
// is the unwind target basic block?
type UnwindTarget interface {
	isUnwindTarget()
}

// TODO: remove isUnwindTarget? or unexport.
func (*BasicBlock) isUnwindTarget() {}

// --- [ Function parameters ] -------------------------------------------------

// Param is an LLVM IR function parameter.
type Param struct {
	// Parameter type.
	Typ types.Type
	// (optional) Parameter name (without '%' prefix).
	LocalName string

	// extra.

	// (optional) Parameter attributes.
	Attrs []ParamAttribute
}

// NewParam returns a new function parameter based on the given type and name.
func NewParam(typ types.Type, name string) *Param {
	return &Param{Typ: typ, LocalName: name}
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
	return enc.Local(p.LocalName)
}

// Name returns the name of the function parameter.
func (p *Param) Name() string {
	return p.LocalName
}

// SetName sets the name of the function parameter.
func (p *Param) SetName(name string) {
	p.LocalName = name
}

// Def returns the LLVM syntax representation of the function parameter.
func (p *Param) Def() string {
	// Type ParamAttrs OptLocalIdent
	buf := &strings.Builder{}
	buf.WriteString(p.Typ.String())
	for _, attr := range p.Attrs {
		fmt.Fprintf(buf, " %v", attr)
	}
	if !isUnnamed(p.LocalName) && !isLocalID(p.LocalName) {
		fmt.Fprintf(buf, " %v", enc.Local(p.LocalName))
	}
	return buf.String()
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
