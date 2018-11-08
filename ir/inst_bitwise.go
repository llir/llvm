package ir

import (
	"fmt"
	"strings"

	"github.com/llir/llvm/internal/enc"
	"github.com/llir/llvm/ir/enum"
	"github.com/llir/llvm/ir/types"
	"github.com/llir/llvm/ir/value"
)

// --- [ Bitwise instructions ] ------------------------------------------------

// ~~~ [ shl ] ~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~

// InstShl is an LLVM IR shl instruction.
type InstShl struct {
	// Name of local variable associated with the result.
	LocalName string
	// Operands.
	X, Y value.Value // integer scalar or integer vector

	// extra.

	// Type of result produced by the instruction.
	Typ types.Type
	// (optional) Overflow flags.
	OverflowFlags []enum.OverflowFlag
	// (optional) Metadata.
	Metadata []*MetadataAttachment
}

// NewShl returns a new shl instruction based on the given operands.
func NewShl(x, y value.Value) *InstShl {
	return &InstShl{X: x, Y: y}
}

// String returns the LLVM syntax representation of the instruction as a
// type-value pair.
func (inst *InstShl) String() string {
	return fmt.Sprintf("%v %v", inst.Type(), inst.Ident())
}

// Type returns the type of the instruction.
func (inst *InstShl) Type() types.Type {
	// Cache type if not present.
	if inst.Typ == nil {
		inst.Typ = inst.X.Type()
	}
	return inst.Typ
}

// Ident returns the identifier associated with the instruction.
func (inst *InstShl) Ident() string {
	return enc.Local(inst.LocalName)
}

// Name returns the name of the instruction.
func (inst *InstShl) Name() string {
	return inst.LocalName
}

// SetName sets the name of the instruction.
func (inst *InstShl) SetName(name string) {
	inst.LocalName = name
}

// Def returns the LLVM syntax representation of the instruction.
func (inst *InstShl) Def() string {
	// "shl" OverflowFlags Type Value "," Value OptCommaSepMetadataAttachmentList
	buf := &strings.Builder{}
	fmt.Fprintf(buf, "%v = ", inst.Ident())
	buf.WriteString("shl")
	for _, flag := range inst.OverflowFlags {
		fmt.Fprintf(buf, " %v", flag)
	}
	fmt.Fprintf(buf, " %v, %v", inst.X, inst.Y.Ident())
	for _, md := range inst.Metadata {
		fmt.Fprintf(buf, ", %v", md)
	}
	return buf.String()
}

// ~~~ [ lshr ] ~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~

// InstLShr is an LLVM IR lshr instruction.
type InstLShr struct {
	// Name of local variable associated with the result.
	LocalName string
	// Operands.
	X, Y value.Value // integer scalars or vectors

	// extra.

	// Type of result produced by the instruction.
	Typ types.Type
	// (optional) Exact.
	Exact bool
	// (optional) Metadata.
	Metadata []*MetadataAttachment
}

// NewLShr returns a new lshr instruction based on the given operands.
func NewLShr(x, y value.Value) *InstLShr {
	return &InstLShr{X: x, Y: y}
}

// String returns the LLVM syntax representation of the instruction as a
// type-value pair.
func (inst *InstLShr) String() string {
	return fmt.Sprintf("%v %v", inst.Type(), inst.Ident())
}

// Type returns the type of the instruction.
func (inst *InstLShr) Type() types.Type {
	// Cache type if not present.
	if inst.Typ == nil {
		inst.Typ = inst.X.Type()
	}
	return inst.Typ
}

// Ident returns the identifier associated with the instruction.
func (inst *InstLShr) Ident() string {
	return enc.Local(inst.LocalName)
}

// Name returns the name of the instruction.
func (inst *InstLShr) Name() string {
	return inst.LocalName
}

// SetName sets the name of the instruction.
func (inst *InstLShr) SetName(name string) {
	inst.LocalName = name
}

// Def returns the LLVM syntax representation of the instruction.
func (inst *InstLShr) Def() string {
	// "lshr" OptExact Type Value "," Value OptCommaSepMetadataAttachmentList
	buf := &strings.Builder{}
	fmt.Fprintf(buf, "%v = ", inst.Ident())
	buf.WriteString("lshr")
	if inst.Exact {
		buf.WriteString(" exact")
	}
	fmt.Fprintf(buf, " %v, %v", inst.X, inst.Y.Ident())
	for _, md := range inst.Metadata {
		fmt.Fprintf(buf, ", %v", md)
	}
	return buf.String()
}

// ~~~ [ ashr ] ~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~

// InstAShr is an LLVM IR ashr instruction.
type InstAShr struct {
	// Name of local variable associated with the result.
	LocalName string
	// Operands.
	X, Y value.Value // integer scalars or vectors

	// extra.

	// Type of result produced by the instruction.
	Typ types.Type
	// (optional) Exact.
	Exact bool
	// (optional) Metadata.
	Metadata []*MetadataAttachment
}

// NewAShr returns a new ashr instruction based on the given operands.
func NewAShr(x, y value.Value) *InstAShr {
	return &InstAShr{X: x, Y: y}
}

// String returns the LLVM syntax representation of the instruction as a
// type-value pair.
func (inst *InstAShr) String() string {
	return fmt.Sprintf("%v %v", inst.Type(), inst.Ident())
}

// Type returns the type of the instruction.
func (inst *InstAShr) Type() types.Type {
	// Cache type if not present.
	if inst.Typ == nil {
		inst.Typ = inst.X.Type()
	}
	return inst.Typ
}

// Ident returns the identifier associated with the instruction.
func (inst *InstAShr) Ident() string {
	return enc.Local(inst.LocalName)
}

// Name returns the name of the instruction.
func (inst *InstAShr) Name() string {
	return inst.LocalName
}

// SetName sets the name of the instruction.
func (inst *InstAShr) SetName(name string) {
	inst.LocalName = name
}

// Def returns the LLVM syntax representation of the instruction.
func (inst *InstAShr) Def() string {
	// "ashr" OptExact Type Value "," Value OptCommaSepMetadataAttachmentList
	buf := &strings.Builder{}
	fmt.Fprintf(buf, "%v = ", inst.Ident())
	buf.WriteString("ashr")
	if inst.Exact {
		buf.WriteString(" exact")
	}
	fmt.Fprintf(buf, " %v, %v", inst.X, inst.Y.Ident())
	for _, md := range inst.Metadata {
		fmt.Fprintf(buf, ", %v", md)
	}
	return buf.String()
}

// ~~~ [ and ] ~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~

// InstAnd is an LLVM IR and instruction.
type InstAnd struct {
	// Name of local variable associated with the result.
	LocalName string
	// Operands.
	X, Y value.Value // integer scalars or vectors

	// extra.

	// Type of result produced by the instruction.
	Typ types.Type
	// (optional) Metadata.
	Metadata []*MetadataAttachment
}

// NewAnd returns a new and instruction based on the given operands.
func NewAnd(x, y value.Value) *InstAnd {
	return &InstAnd{X: x, Y: y}
}

// String returns the LLVM syntax representation of the instruction as a
// type-value pair.
func (inst *InstAnd) String() string {
	return fmt.Sprintf("%v %v", inst.Type(), inst.Ident())
}

// Type returns the type of the instruction.
func (inst *InstAnd) Type() types.Type {
	// Cache type if not present.
	if inst.Typ == nil {
		inst.Typ = inst.X.Type()
	}
	return inst.Typ
}

// Ident returns the identifier associated with the instruction.
func (inst *InstAnd) Ident() string {
	return enc.Local(inst.LocalName)
}

// Name returns the name of the instruction.
func (inst *InstAnd) Name() string {
	return inst.LocalName
}

// SetName sets the name of the instruction.
func (inst *InstAnd) SetName(name string) {
	inst.LocalName = name
}

// Def returns the LLVM syntax representation of the instruction.
func (inst *InstAnd) Def() string {
	// "and" Type Value "," Value OptCommaSepMetadataAttachmentList
	buf := &strings.Builder{}
	fmt.Fprintf(buf, "%v = ", inst.Ident())
	fmt.Fprintf(buf, "and %v, %v", inst.X, inst.Y.Ident())
	for _, md := range inst.Metadata {
		fmt.Fprintf(buf, ", %v", md)
	}
	return buf.String()
}

// ~~~ [ or ] ~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~

// InstOr is an LLVM IR or instruction.
type InstOr struct {
	// Name of local variable associated with the result.
	LocalName string
	// Operands.
	X, Y value.Value // integer scalars or vectors

	// extra.

	// Type of result produced by the instruction.
	Typ types.Type
	// (optional) Metadata.
	Metadata []*MetadataAttachment
}

// NewOr returns a new or instruction based on the given operands.
func NewOr(x, y value.Value) *InstOr {
	return &InstOr{X: x, Y: y}
}

// String returns the LLVM syntax representation of the instruction as a
// type-value pair.
func (inst *InstOr) String() string {
	return fmt.Sprintf("%v %v", inst.Type(), inst.Ident())
}

// Type returns the type of the instruction.
func (inst *InstOr) Type() types.Type {
	// Cache type if not present.
	if inst.Typ == nil {
		inst.Typ = inst.X.Type()
	}
	return inst.Typ
}

// Ident returns the identifier associated with the instruction.
func (inst *InstOr) Ident() string {
	return enc.Local(inst.LocalName)
}

// Name returns the name of the instruction.
func (inst *InstOr) Name() string {
	return inst.LocalName
}

// SetName sets the name of the instruction.
func (inst *InstOr) SetName(name string) {
	inst.LocalName = name
}

// Def returns the LLVM syntax representation of the instruction.
func (inst *InstOr) Def() string {
	// "or" Type Value "," Value OptCommaSepMetadataAttachmentList
	buf := &strings.Builder{}
	fmt.Fprintf(buf, "%v = ", inst.Ident())
	fmt.Fprintf(buf, "or %v, %v", inst.X, inst.Y.Ident())
	for _, md := range inst.Metadata {
		fmt.Fprintf(buf, ", %v", md)
	}
	return buf.String()
}

// ~~~ [ xor ] ~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~

// InstXor is an LLVM IR xor instruction.
type InstXor struct {
	// Name of local variable associated with the result.
	LocalName string
	// Operands.
	X, Y value.Value // integer scalars or vectors

	// extra.

	// Type of result produced by the instruction.
	Typ types.Type
	// (optional) Metadata.
	Metadata []*MetadataAttachment
}

// NewXor returns a new xor instruction based on the given operands.
func NewXor(x, y value.Value) *InstXor {
	return &InstXor{X: x, Y: y}
}

// String returns the LLVM syntax representation of the instruction as a
// type-value pair.
func (inst *InstXor) String() string {
	return fmt.Sprintf("%v %v", inst.Type(), inst.Ident())
}

// Type returns the type of the instruction.
func (inst *InstXor) Type() types.Type {
	// Cache type if not present.
	if inst.Typ == nil {
		inst.Typ = inst.X.Type()
	}
	return inst.Typ
}

// Ident returns the identifier associated with the instruction.
func (inst *InstXor) Ident() string {
	return enc.Local(inst.LocalName)
}

// Name returns the name of the instruction.
func (inst *InstXor) Name() string {
	return inst.LocalName
}

// SetName sets the name of the instruction.
func (inst *InstXor) SetName(name string) {
	inst.LocalName = name
}

// Def returns the LLVM syntax representation of the instruction.
func (inst *InstXor) Def() string {
	// "xor" Type Value "," Value OptCommaSepMetadataAttachmentList
	buf := &strings.Builder{}
	fmt.Fprintf(buf, "%v = ", inst.Ident())
	fmt.Fprintf(buf, "xor %v, %v", inst.X, inst.Y.Ident())
	for _, md := range inst.Metadata {
		fmt.Fprintf(buf, ", %v", md)
	}
	return buf.String()
}
