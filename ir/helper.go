package ir

import (
	"fmt"
	"strings"

	"github.com/llir/llvm/internal/enc"
	"github.com/llir/llvm/ir/enum"
	"github.com/llir/llvm/ir/types"
	"github.com/llir/llvm/ir/value"
)

// TODO: move to the right place.

// Alignment is a memory alignment attribute.
type Alignment int64

// IsParamAttribute ensures that only parameter attributes can be assigned to
// the ir.ParamAttribute interface.
func (Alignment) IsParamAttribute() {}

// IsReturnAttribute ensures that only return attributes can be assigned to the
// ir.ReturnAttribute interface.
func (Alignment) IsReturnAttribute() {}

// Arg is a function argument.
type Arg struct {
	// Argument value.
	value.Value
	// (optional) Parameter attributes.
	Attrs []ParamAttribute
}

// String returns a string representation of the function argument.
func (arg *Arg) String() string {
	// ConcreteType ParamAttrs Value
	buf := &strings.Builder{}
	buf.WriteString(arg.Type().String())
	for _, attr := range arg.Attrs {
		fmt.Fprintf(buf, " %v", attr)
	}
	fmt.Fprintf(buf, " %v", arg.Ident())
	return buf.String()
}

// TODO: figure out definition of ExceptionScope.
type ExceptionScope interface {
	value.Value
	//isExceptionScope()
}

// TODO: consider getting rid of UnwindTarget, and let unwind targets be of type
// *ir.BasicBlock, where a nil value indicates the caller, and a non-nil value
// is the unwind target basic block?

// TODO: figure out definition of UnwindTarget.
type UnwindTarget interface {
	//value.Value
	isUnwindTarget()
}

type UnwindToCaller struct{}

// String returns the string representation of the unwind to caller target.
func (*UnwindToCaller) String() string {
	return "to caller"
}

func (*UnwindToCaller) isUnwindTarget() {}

// TODO: remove isUnwindTarget? or unexport.
func (*BasicBlock) isUnwindTarget() {}

// FuncAttribute is a function attribute.
type FuncAttribute interface {
	fmt.Stringer
	// IsFuncAttribute ensures that only function attributes can be assigned to
	// the ir.FuncAttribute interface.
	IsFuncAttribute()
}

// ParamAttribute is a parameter attribute.
type ParamAttribute interface {
	// IsParamAttribute ensures that only parameter attributes can be assigned to
	// the ir.ParamAttribute interface.
	IsParamAttribute()
}

// ReturnAttribute is a return attribute.
type ReturnAttribute interface {
	// IsReturnAttribute ensures that only return attributes can be assigned to
	// the ir.ReturnAttribute interface.
	IsReturnAttribute()
}

// AttrString is an attribute string (used in function, parameter and return
// attributes).
type AttrString string

// String returns the string representation of the attribute string.
func (a AttrString) String() string {
	return enc.Quote([]byte(a))
}

// IsFuncAttribute ensures that only function attributes can be assigned to
// the ir.FuncAttribute interface.
func (AttrString) IsFuncAttribute() {}

// IsParamAttribute ensures that only parameter attributes can be assigned to
// the ir.ParamAttribute interface.
func (AttrString) IsParamAttribute() {}

// IsReturnAttribute ensures that only return attributes can be assigned to
// the ir.ReturnAttribute interface.
func (AttrString) IsReturnAttribute() {}

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

// callingConvString returns the string representation of the given calling
// convention.
func callingConvString(callingConv enum.CallingConv) string {
	if callingConv > enum.CallingConvNNN {
		cc := callingConv - enum.CallingConvNNN
		return fmt.Sprintf("cc %d", cc)
	}
	return callingConv.String()
}
