package ir

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/llir/llvm/internal/enc"
	"github.com/llir/llvm/ir/enum"
	"github.com/llir/llvm/ir/metadata"
	"github.com/llir/llvm/ir/types"
	"github.com/llir/llvm/ir/value"
)

// Align is a memory alignment attribute.
type Align uint64

// String returns the string representation of the alignment attribute.
func (align Align) String() string {
	// Note, alignment is printed as `align = 8` in attribute groups.
	return fmt.Sprintf("align %d", uint64(align))
}

// AlignStack is a stack alignment attribute.
type AlignStack uint64

// String returns the string representation of the stack alignment attribute.
func (align AlignStack) String() string {
	// Note, stack alignment is printed as `alignstack = 8` in attribute groups.
	return fmt.Sprintf("alignstack(%d)", uint64(align))
}

// AllocSize is an attribute for functions like malloc. If the second parameter
// is omitted, NElemsIndex will be -1.
type AllocSize struct {
	// Element size parameter index.
	ElemSizeIndex int
	// Number of elements parameter index; -1 if not present.
	NElemsIndex int
}

// String returns the string representation of the allocsize attribute.
func (a AllocSize) String() string {
	if a.NElemsIndex == -1 {
		return fmt.Sprintf("allocsize(%d)", a.ElemSizeIndex)
	}
	return fmt.Sprintf("allocsize(%d, %d)", a.ElemSizeIndex, a.NElemsIndex)
}

// Arg is a function argument with optional parameter attributes.
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
		fmt.Fprintf(buf, " %s", attr)
	}
	fmt.Fprintf(buf, " %s", arg.Ident())
	return buf.String()
}

// AttrPair is an attribute key-value pair (used in function, parameter and
// return attributes).
type AttrPair struct {
	Key   string
	Value string
}

// String returns the string representation of the attribute key-value pair.
func (a AttrPair) String() string {
	// Key=StringLit '=' Val=StringLit
	return fmt.Sprintf("%s=%s", quote(a.Key), quote(a.Value))
}

// AttrString is an attribute string (used in function, parameter and return
// attributes).
type AttrString string

// String returns the string representation of the attribute string.
func (a AttrString) String() string {
	return quote(string(a))
}

// Dereferenceable is a dereferenceable memory attribute.
type Dereferenceable struct {
	// Number of bytes known to be dereferenceable.
	N uint64
	// (optional) Either dereferenceable or null if set.
	DerefOrNull bool
}

// String returns the string representation of the dereferenceable memory
// attribute.
func (d Dereferenceable) String() string {
	// 'dereferenceable' '(' N=UintLit ')'
	//
	// 'dereferenceable_or_null' '(' N=UintLit ')'
	if d.DerefOrNull {
		return fmt.Sprintf("dereferenceable_or_null(%d)", d.N)
	}
	return fmt.Sprintf("dereferenceable(%d)", d.N)
}

// TODO: figure out definition of ExceptionScope.

// ExceptionScope is an exception scope.
type ExceptionScope interface {
	value.Value
	//isExceptionScope()
}

// FuncAttribute is a function attribute.
//
// A FuncAttribute has one of the following underlying types.
//
//    ir.AttrString
//    ir.AttrPair
//    *ir.AttrGroupDef
//    ir.Align
//    ir.AlignStack
//    ir.AllocSize
//    enum.FuncAttr
type FuncAttribute interface {
	fmt.Stringer
	// IsFuncAttribute ensures that only function attributes can be assigned to
	// the ir.FuncAttribute interface.
	IsFuncAttribute()
}

// GlobalIdent is a global identifier.
type GlobalIdent struct {
	GlobalName string
	GlobalID   int64
}

// Ident returns the identifier associated with the global identifier.
func (i GlobalIdent) Ident() string {
	if i.IsUnnamed() {
		id := strconv.FormatInt(i.GlobalID, 10)
		return enc.Global(id)
	}
	name := i.GlobalName
	if x, err := strconv.ParseInt(name, 10, 64); err == nil {
		// Print GlobalName with quotes if it is a number; e.g. @"42".
		return fmt.Sprintf(`@"%d"`, x)
	}
	return enc.Global(name)
}

// Name returns the name of the global identifier.
//
// If unnamed, the global ID is returned. To distinguish numeric names from
// unnamed IDs, numeric names are quoted.
func (i GlobalIdent) Name() string {
	if i.IsUnnamed() {
		return strconv.FormatInt(i.GlobalID, 10)
	}
	if x, err := strconv.ParseInt(i.GlobalName, 10, 64); err == nil {
		// Print GlobalName with quotes if it is a number; e.g. "42".
		return fmt.Sprintf(`"%d"`, x)
	}
	return i.GlobalName
}

// SetName sets the name of the global identifier.
func (i *GlobalIdent) SetName(name string) {
	i.GlobalName = name
	i.GlobalID = 0
}

// ID returns the ID of the global identifier.
func (i GlobalIdent) ID() int64 {
	return i.GlobalID
}

// SetID sets the ID of the global identifier.
func (i *GlobalIdent) SetID(id int64) {
	i.GlobalID = id
}

// IsUnnamed reports whether the global identifier is unnamed.
func (i GlobalIdent) IsUnnamed() bool {
	return len(i.GlobalName) == 0
}

// LocalIdent is a local identifier.
type LocalIdent struct {
	LocalName string
	LocalID   int64
}

// Ident returns the identifier associated with the local identifier.
func (i LocalIdent) Ident() string {
	if i.IsUnnamed() {
		id := strconv.FormatInt(i.LocalID, 10)
		return enc.Local(id)
	}
	name := i.LocalName
	if x, err := strconv.ParseInt(name, 10, 64); err == nil {
		// Print LocalName with quotes if it is a number; e.g. %"42".
		return fmt.Sprintf(`%%"%d"`, x)
	}
	return enc.Local(name)
}

// Name returns the name of the local identifier.
//
// If unnamed, the local ID is returned. To distinguish numeric names from
// unnamed IDs, numeric names are quoted.
func (i LocalIdent) Name() string {
	if i.IsUnnamed() {
		return strconv.FormatInt(i.LocalID, 10)
	}
	if x, err := strconv.ParseInt(i.LocalName, 10, 64); err == nil {
		// Print LocalName with quotes if it is a number; e.g. "42".
		return fmt.Sprintf(`"%d"`, x)
	}
	return i.LocalName
}

// SetName sets the name of the local identifier.
func (i *LocalIdent) SetName(name string) {
	i.LocalName = name
	i.LocalID = 0
}

// ID returns the ID of the local identifier.
func (i LocalIdent) ID() int64 {
	return i.LocalID
}

// SetID sets the ID of the local identifier.
func (i *LocalIdent) SetID(id int64) {
	i.LocalID = id
}

// IsUnnamed reports whether the local identifier is unnamed.
func (i LocalIdent) IsUnnamed() bool {
	return len(i.LocalName) == 0
}

// Metadata is a list of metadata attachments.
type Metadata []*metadata.Attachment

// MDAttachments returns the metadata attachments of the value.
func (mds Metadata) MDAttachments() []*metadata.Attachment {
	return mds
}

// OperandBundle is an operand bundle.
type OperandBundle struct {
	Tag    string
	Inputs []value.Value
}

// NewOperandBundle returns a new operand bundle based on the given tag and
// input values.
func NewOperandBundle(tag string, inputs ...value.Value) *OperandBundle {
	return &OperandBundle{Tag: tag, Inputs: inputs}
}

// String returns a string representation of the operand bundle.
func (o *OperandBundle) String() string {
	// Tag=StringLit '(' Inputs=(TypeValue separator ',')* ')'
	buf := &strings.Builder{}
	fmt.Fprintf(buf, "%s(", quote(o.Tag))
	for i, input := range o.Inputs {
		if i != 0 {
			buf.WriteString(", ")
		}
		buf.WriteString(input.String())
	}
	buf.WriteString(")")
	return buf.String()
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
//
// An UnwindTarget has one of the following underlying types.
//
//    *ir.Block
//    ir.UnwindToCaller
type UnwindTarget interface {
	// isUnwindTarget ensures that only unwind targets can be assigned to the
	// ir.UnwindTarget interface.
	isUnwindTarget()
}

// UnwindToCaller specifies the caller as an unwind target.
type UnwindToCaller struct{}

// String returns the string representation of the unwind target.
func (UnwindToCaller) String() string {
	return "to caller"
}

// ___ [ Function parameter ] __________________________________________________

// Param is an LLVM IR function parameter.
type Param struct {
	// (optional) Parameter name (without '%' prefix).
	LocalIdent
	// Parameter type.
	Typ types.Type

	// extra.

	// (optional) Parameter attributes.
	Attrs []ParamAttribute

	// for implement ignore attr behavior
	ignoreAttr bool
}

// NewParam returns a new function parameter based on the given name and type.
func NewParam(name string, typ types.Type) *Param {
	return &Param{
		LocalIdent: LocalIdent{LocalName: name},
		Typ:        typ,
	}
}

// String returns the LLVM syntax representation of the function parameter as a
// type-value pair.
func (p *Param) String() string {
	buf := &strings.Builder{}
	buf.WriteString(p.Type().String())
	if p.Type().Equal(types.Metadata) || !p.ignoreAttr {
		for _, attr := range p.Attrs {
			buf.WriteRune(' ')
			buf.WriteString(attr.String())
		}
	} else {
		p.ignoreAttr = false // immediately set back after use
	}
	buf.WriteRune(' ')
	buf.WriteString(p.Ident())
	return buf.String()
}

func (p *Param) IgnoreAttr() {
	p.ignoreAttr = true
}

// Type returns the type of the function parameter.
func (p *Param) Type() types.Type {
	return p.Typ
}

// LLString returns the LLVM syntax representation of the function parameter.
func (p *Param) LLString() string {
	// Typ=Type Attrs=ParamAttribute* Name=LocalIdent?
	buf := &strings.Builder{}
	buf.WriteString(p.Typ.String())
	for _, attr := range p.Attrs {
		fmt.Fprintf(buf, " %s", attr)
	}
	if !p.IsUnnamed() {
		fmt.Fprintf(buf, " %s", p.Ident())
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
	s := callingConv.String()
	cc := uint(callingConv)
	if unknown := fmt.Sprintf("CallingConv(%d)", cc); s == unknown {
		return fmt.Sprintf("cc %d", cc)
	}
	return s
}

// tlsModelString returns the string representation of the given thread local
// storage model.
func tlsModelString(model enum.TLSModel) string {
	if model == enum.TLSModelGeneric {
		return "thread_local"
	}
	return fmt.Sprintf("thread_local(%s)", model)
}
