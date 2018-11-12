package ir

import (
	"fmt"
	"strings"

	"github.com/llir/llvm/internal/enc"
	"github.com/llir/llvm/ir/enum"
	"github.com/llir/llvm/ir/types"
	"github.com/llir/llvm/ir/value"
)

// Alignment is a memory alignment attribute.
type Align int64

// String returns the string representation of the alignment attribute.
func (align Align) String() string {
	// Note, alignment is printed as `align = 8` in attribute groups.
	return fmt.Sprintf("align %d", int64(align))
}

// Arg is a function argument.
type Arg struct {
	// Argument value.
	value.Value
	// (optional) Parameter attributes.
	Attrs []ParamAttribute
}

// NewArg returns a new function argument based on the given value and parameter
// attributes.
func NewArg(x value.Value, attrs ...ParamAttribute) *Arg {
	return &Arg{Value: x, Attrs: attrs}
}

// String returns a string representation of the function argument.
func (arg *Arg) String() string {
	// Typ=ConcreteType Attrs=ParamAttribute* Val=Value
	buf := &strings.Builder{}
	buf.WriteString(arg.Type().String())
	for _, attr := range arg.Attrs {
		fmt.Fprintf(buf, " %v", attr)
	}
	fmt.Fprintf(buf, " %v", arg.Ident())
	return buf.String()
}

// AttrString is an attribute string (used in function, parameter and return
// attributes).
type AttrString string

// String returns the string representation of the attribute string.
func (a AttrString) String() string {
	return quote(string(a))
}

// TODO: figure out definition of ExceptionScope.

// ExceptionScope is an exception scope.
type ExceptionScope interface {
	value.Value
	//isExceptionScope()
}

// FuncAttribute is a function attribute.
type FuncAttribute interface {
	fmt.Stringer
	// IsFuncAttribute ensures that only function attributes can be assigned to
	// the ir.FuncAttribute interface.
	IsFuncAttribute()
}

// ParamAttribute is a parameter attribute.
//
// A ParamAttribute has one of the following underlying types.
//
//    ir.AttrString
//    ir.AttrPair
//    ir.Align
//    ir.Dereferenceable
//    enum.ParamAttr
type ParamAttribute interface {
	fmt.Stringer
	// IsParamAttribute ensures that only parameter attributes can be assigned to
	// the ir.ParamAttribute interface.
	IsParamAttribute()
}

// ReturnAttribute is a return attribute.
//
// A ReturnAttribute has one of the following underlying types.
//
//    ir.AttrString
//    ir.AttrPair
//    ir.Align
//    ir.Dereferenceable
//    enum.ReturnAttr
type ReturnAttribute interface {
	fmt.Stringer
	// IsReturnAttribute ensures that only return attributes can be assigned to
	// the ir.ReturnAttribute interface.
	IsReturnAttribute()
}

// UnwindTarget is an unwind target.
type UnwindTarget interface {
	isUnwindTarget()
}

// UnwindToCaller specifies the caller as an unwind target.
type UnwindToCaller struct{}

// String returns the string representation of the unwind target.
func (*UnwindToCaller) String() string {
	return "to caller"
}

// AttrPair is an attribute key-value pair (used in function, parameter and
// return attributes).
type AttrPair struct {
	Key   string
	Value string
}

// String returns the string representation of the attribute key-value pair.
func (a AttrPair) String() string {
	return fmt.Sprintf("%s=%s", enc.Quote([]byte(a.Key)), enc.Quote([]byte(a.Value)))
}

// IsFuncAttribute ensures that only function attributes can be assigned to
// the ir.FuncAttribute interface.
func (AttrPair) IsFuncAttribute() {}

// IsParamAttribute ensures that only parameter attributes can be assigned to
// the ir.ParamAttribute interface.
func (AttrPair) IsParamAttribute() {}

// IsReturnAttribute ensures that only return attributes can be assigned to
// the ir.ReturnAttribute interface.
func (AttrPair) IsReturnAttribute() {}

type OperandBundle struct {
	// TODO: implement body.
}

// --- [ Function parameters ] -------------------------------------------------

// Param is an LLVM IR function parameter.
type Param struct {
	// (optional) Parameter name (without '%' prefix).
	LocalName string
	// Parameter type.
	Typ types.Type

	// extra.

	// (optional) Parameter attributes.
	Attrs []ParamAttribute
}

// NewParam returns a new function parameter based on the given name and type.
func NewParam(name string, typ types.Type) *Param {
	return &Param{LocalName: name, Typ: typ}
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

// callingConvString returns the string representation of the given calling
// convention.
func callingConvString(callingConv enum.CallingConv) string {
	if callingConv > enum.CallingConvNNN {
		cc := callingConv - enum.CallingConvNNN
		return fmt.Sprintf("cc %d", cc)
	}
	return callingConv.String()
}
