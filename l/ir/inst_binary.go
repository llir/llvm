package ir

import (
	"fmt"
	"strings"

	"github.com/llir/l/internal/enc"
	"github.com/llir/l/ir/enum"
	"github.com/llir/l/ir/types"
	"github.com/llir/l/ir/value"
)

// --- [ Binary instructions ] -------------------------------------------------

// ~~~ [ add ] ~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~

// InstAdd is an LLVM IR add instruction.
type InstAdd struct {
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
	Metadata []MetadataAttachment
}

// NewAdd returns a new add instruction based on the given operands.
func NewAdd(x, y value.Value) *InstAdd {
	return &InstAdd{X: x, Y: y}
}

// String returns the LLVM syntax representation of the instruction as a
// type-value pair.
func (inst *InstAdd) String() string {
	return fmt.Sprintf("%v %v", inst.Type(), inst.Ident())
}

// Type returns the type of the instruction.
func (inst *InstAdd) Type() types.Type {
	// Cache type if not present.
	if inst.Typ == nil {
		inst.Typ = inst.X.Type()
	}
	return inst.Typ
}

// Ident returns the identifier associated with the instruction.
func (inst *InstAdd) Ident() string {
	return enc.Local(inst.LocalName)
}

// Name returns the name of the instruction.
func (inst *InstAdd) Name() string {
	return inst.LocalName
}

// SetName sets the name of the instruction.
func (inst *InstAdd) SetName(name string) {
	inst.LocalName = name
}

// Def returns the LLVM syntax representation of the instruction.
func (inst *InstAdd) Def() string {
	// "add" OverflowFlags Type Value "," Value OptCommaSepMetadataAttachmentList
	buf := &strings.Builder{}
	buf.WriteString("add")
	for _, flag := range inst.OverflowFlags {
		fmt.Fprintf(buf, " %v", flag)
	}
	fmt.Fprintf(buf, " %v, %v", inst.X, inst.Y.Ident())
	for _, md := range inst.Metadata {
		fmt.Fprintf(buf, ", %v", md)
	}
	return buf.String()
}

// ~~~ [ fadd ] ~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~

// InstFAdd is an LLVM IR fadd instruction.
type InstFAdd struct {
	// Name of local variable associated with the result.
	LocalName string
	// Operands.
	X, Y value.Value // floating-point scalar or floating-point vector

	// extra.

	// Type of result produced by the instruction.
	Typ types.Type
	// (optional) Fast math flags.
	FastMathFlags []enum.FastMathFlag
	// (optional) Metadata.
	Metadata []MetadataAttachment
}

// NewFAdd returns a new fadd instruction based on the given operands.
func NewFAdd(x, y value.Value) *InstFAdd {
	return &InstFAdd{X: x, Y: y}
}

// String returns the LLVM syntax representation of the instruction as a
// type-value pair.
func (inst *InstFAdd) String() string {
	return fmt.Sprintf("%v %v", inst.Type(), inst.Ident())
}

// Type returns the type of the instruction.
func (inst *InstFAdd) Type() types.Type {
	// Cache type if not present.
	if inst.Typ == nil {
		inst.Typ = inst.X.Type()
	}
	return inst.Typ
}

// Ident returns the identifier associated with the instruction.
func (inst *InstFAdd) Ident() string {
	return enc.Local(inst.LocalName)
}

// Name returns the name of the instruction.
func (inst *InstFAdd) Name() string {
	return inst.LocalName
}

// SetName sets the name of the instruction.
func (inst *InstFAdd) SetName(name string) {
	inst.LocalName = name
}

// Def returns the LLVM syntax representation of the instruction.
func (inst *InstFAdd) Def() string {
	// "fadd" FastMathFlags Type Value "," Value OptCommaSepMetadataAttachmentList
	buf := &strings.Builder{}
	buf.WriteString("fadd")
	for _, flag := range inst.FastMathFlags {
		fmt.Fprintf(buf, " %v", flag)
	}
	fmt.Fprintf(buf, " %v, %v", inst.X, inst.Y.Ident())
	for _, md := range inst.Metadata {
		fmt.Fprintf(buf, ", %v", md)
	}
	return buf.String()
}

// ~~~ [ sub ] ~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~

// InstSub is an LLVM IR sub instruction.
type InstSub struct {
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
	Metadata []MetadataAttachment
}

// NewSub returns a new sub instruction based on the given operands.
func NewSub(x, y value.Value) *InstSub {
	return &InstSub{X: x, Y: y}
}

// String returns the LLVM syntax representation of the instruction as a
// type-value pair.
func (inst *InstSub) String() string {
	return fmt.Sprintf("%v %v", inst.Type(), inst.Ident())
}

// Type returns the type of the instruction.
func (inst *InstSub) Type() types.Type {
	// Cache type if not present.
	if inst.Typ == nil {
		inst.Typ = inst.X.Type()
	}
	return inst.Typ
}

// Ident returns the identifier associated with the instruction.
func (inst *InstSub) Ident() string {
	return enc.Local(inst.LocalName)
}

// Name returns the name of the instruction.
func (inst *InstSub) Name() string {
	return inst.LocalName
}

// SetName sets the name of the instruction.
func (inst *InstSub) SetName(name string) {
	inst.LocalName = name
}

// Def returns the LLVM syntax representation of the instruction.
func (inst *InstSub) Def() string {
	// "sub" OverflowFlags Type Value "," Value OptCommaSepMetadataAttachmentList
	buf := &strings.Builder{}
	buf.WriteString("sub")
	for _, flag := range inst.OverflowFlags {
		fmt.Fprintf(buf, " %v", flag)
	}
	fmt.Fprintf(buf, " %v, %v", inst.X, inst.Y.Ident())
	for _, md := range inst.Metadata {
		fmt.Fprintf(buf, ", %v", md)
	}
	return buf.String()
}

// ~~~ [ fsub ] ~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~

// InstFSub is an LLVM IR fsub instruction.
type InstFSub struct {
	// Name of local variable associated with the result.
	LocalName string
	// Operands.
	X, Y value.Value // floating-point scalar or floating-point vector

	// extra.

	// Type of result produced by the instruction.
	Typ types.Type
	// (optional) Fast math flags.
	FastMathFlags []enum.FastMathFlag
	// (optional) Metadata.
	Metadata []MetadataAttachment
}

// NewFSub returns a new fsub instruction based on the given operands.
func NewFSub(x, y value.Value) *InstFSub {
	return &InstFSub{X: x, Y: y}
}

// String returns the LLVM syntax representation of the instruction as a
// type-value pair.
func (inst *InstFSub) String() string {
	return fmt.Sprintf("%v %v", inst.Type(), inst.Ident())
}

// Type returns the type of the instruction.
func (inst *InstFSub) Type() types.Type {
	// Cache type if not present.
	if inst.Typ == nil {
		inst.Typ = inst.X.Type()
	}
	return inst.Typ
}

// Ident returns the identifier associated with the instruction.
func (inst *InstFSub) Ident() string {
	return enc.Local(inst.LocalName)
}

// Name returns the name of the instruction.
func (inst *InstFSub) Name() string {
	return inst.LocalName
}

// SetName sets the name of the instruction.
func (inst *InstFSub) SetName(name string) {
	inst.LocalName = name
}

// Def returns the LLVM syntax representation of the instruction.
func (inst *InstFSub) Def() string {
	// "fsub" FastMathFlags Type Value "," Value OptCommaSepMetadataAttachmentList
	buf := &strings.Builder{}
	buf.WriteString("fsub")
	for _, flag := range inst.FastMathFlags {
		fmt.Fprintf(buf, " %v", flag)
	}
	fmt.Fprintf(buf, " %v, %v", inst.X, inst.Y.Ident())
	for _, md := range inst.Metadata {
		fmt.Fprintf(buf, ", %v", md)
	}
	return buf.String()
}

// ~~~ [ mul ] ~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~

// InstMul is an LLVM IR mul instruction.
type InstMul struct {
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
	Metadata []MetadataAttachment
}

// NewMul returns a new mul instruction based on the given operands.
func NewMul(x, y value.Value) *InstMul {
	return &InstMul{X: x, Y: y}
}

// String returns the LLVM syntax representation of the instruction as a
// type-value pair.
func (inst *InstMul) String() string {
	return fmt.Sprintf("%v %v", inst.Type(), inst.Ident())
}

// Type returns the type of the instruction.
func (inst *InstMul) Type() types.Type {
	// Cache type if not present.
	if inst.Typ == nil {
		inst.Typ = inst.X.Type()
	}
	return inst.Typ
}

// Ident returns the identifier associated with the instruction.
func (inst *InstMul) Ident() string {
	return enc.Local(inst.LocalName)
}

// Name returns the name of the instruction.
func (inst *InstMul) Name() string {
	return inst.LocalName
}

// SetName sets the name of the instruction.
func (inst *InstMul) SetName(name string) {
	inst.LocalName = name
}

// Def returns the LLVM syntax representation of the instruction.
func (inst *InstMul) Def() string {
	// "mul" OverflowFlags Type Value "," Value OptCommaSepMetadataAttachmentList
	buf := &strings.Builder{}
	buf.WriteString("mul")
	for _, flag := range inst.OverflowFlags {
		fmt.Fprintf(buf, " %v", flag)
	}
	fmt.Fprintf(buf, " %v, %v", inst.X, inst.Y.Ident())
	for _, md := range inst.Metadata {
		fmt.Fprintf(buf, ", %v", md)
	}
	return buf.String()
}

// ~~~ [ fmul ] ~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~

// InstFMul is an LLVM IR fmul instruction.
type InstFMul struct {
	// Name of local variable associated with the result.
	LocalName string
	// Operands.
	X, Y value.Value // floating-point scalar or floating-point vector

	// extra.

	// Type of result produced by the instruction.
	Typ types.Type
	// (optional) Fast math flags.
	FastMathFlags []enum.FastMathFlag
	// (optional) Metadata.
	Metadata []MetadataAttachment
}

// NewFMul returns a new fmul instruction based on the given operands.
func NewFMul(x, y value.Value) *InstFMul {
	return &InstFMul{X: x, Y: y}
}

// String returns the LLVM syntax representation of the instruction as a
// type-value pair.
func (inst *InstFMul) String() string {
	return fmt.Sprintf("%v %v", inst.Type(), inst.Ident())
}

// Type returns the type of the instruction.
func (inst *InstFMul) Type() types.Type {
	// Cache type if not present.
	if inst.Typ == nil {
		inst.Typ = inst.X.Type()
	}
	return inst.Typ
}

// Ident returns the identifier associated with the instruction.
func (inst *InstFMul) Ident() string {
	return enc.Local(inst.LocalName)
}

// Name returns the name of the instruction.
func (inst *InstFMul) Name() string {
	return inst.LocalName
}

// SetName sets the name of the instruction.
func (inst *InstFMul) SetName(name string) {
	inst.LocalName = name
}

// Def returns the LLVM syntax representation of the instruction.
func (inst *InstFMul) Def() string {
	// "fmul" FastMathFlags Type Value "," Value OptCommaSepMetadataAttachmentList
	buf := &strings.Builder{}
	buf.WriteString("fmul")
	for _, flag := range inst.FastMathFlags {
		fmt.Fprintf(buf, " %v", flag)
	}
	fmt.Fprintf(buf, " %v, %v", inst.X, inst.Y.Ident())
	for _, md := range inst.Metadata {
		fmt.Fprintf(buf, ", %v", md)
	}
	return buf.String()
}

// ~~~ [ udiv ] ~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~

// InstUDiv is an LLVM IR udiv instruction.
type InstUDiv struct {
	// Name of local variable associated with the result.
	LocalName string
	// Operands.
	X, Y value.Value // integer scalar or integer vector

	// extra.

	// Type of result produced by the instruction.
	Typ types.Type
	// (optional) Exact.
	Exact bool
	// (optional) Metadata.
	Metadata []MetadataAttachment
}

// NewUDiv returns a new udiv instruction based on the given operands.
func NewUDiv(x, y value.Value) *InstUDiv {
	return &InstUDiv{X: x, Y: y}
}

// String returns the LLVM syntax representation of the instruction as a
// type-value pair.
func (inst *InstUDiv) String() string {
	return fmt.Sprintf("%v %v", inst.Type(), inst.Ident())
}

// Type returns the type of the instruction.
func (inst *InstUDiv) Type() types.Type {
	// Cache type if not present.
	if inst.Typ == nil {
		inst.Typ = inst.X.Type()
	}
	return inst.Typ
}

// Ident returns the identifier associated with the instruction.
func (inst *InstUDiv) Ident() string {
	return enc.Local(inst.LocalName)
}

// Name returns the name of the instruction.
func (inst *InstUDiv) Name() string {
	return inst.LocalName
}

// SetName sets the name of the instruction.
func (inst *InstUDiv) SetName(name string) {
	inst.LocalName = name
}

// Def returns the LLVM syntax representation of the instruction.
func (inst *InstUDiv) Def() string {
	// "udiv" OptExact Type Value "," Value OptCommaSepMetadataAttachmentList
	buf := &strings.Builder{}
	buf.WriteString("udiv")
	if inst.Exact {
		buf.WriteString(" exact")
	}
	fmt.Fprintf(buf, " %v, %v", inst.X, inst.Y.Ident())
	for _, md := range inst.Metadata {
		fmt.Fprintf(buf, ", %v", md)
	}
	return buf.String()
}

// ~~~ [ sdiv ] ~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~

// InstSDiv is an LLVM IR sdiv instruction.
type InstSDiv struct {
	// Name of local variable associated with the result.
	LocalName string
	// Operands.
	X, Y value.Value // integer scalar or integer vector

	// extra.

	// Type of result produced by the instruction.
	Typ types.Type
	// (optional) Exact.
	Exact bool
	// (optional) Metadata.
	Metadata []MetadataAttachment
}

// NewSDiv returns a new sdiv instruction based on the given operands.
func NewSDiv(x, y value.Value) *InstSDiv {
	return &InstSDiv{X: x, Y: y}
}

// String returns the LLVM syntax representation of the instruction as a
// type-value pair.
func (inst *InstSDiv) String() string {
	return fmt.Sprintf("%v %v", inst.Type(), inst.Ident())
}

// Type returns the type of the instruction.
func (inst *InstSDiv) Type() types.Type {
	// Cache type if not present.
	if inst.Typ == nil {
		inst.Typ = inst.X.Type()
	}
	return inst.Typ
}

// Ident returns the identifier associated with the instruction.
func (inst *InstSDiv) Ident() string {
	return enc.Local(inst.LocalName)
}

// Name returns the name of the instruction.
func (inst *InstSDiv) Name() string {
	return inst.LocalName
}

// SetName sets the name of the instruction.
func (inst *InstSDiv) SetName(name string) {
	inst.LocalName = name
}

// Def returns the LLVM syntax representation of the instruction.
func (inst *InstSDiv) Def() string {
	// "sdiv" OptExact Type Value "," Value OptCommaSepMetadataAttachmentList
	buf := &strings.Builder{}
	buf.WriteString("sdiv")
	if inst.Exact {
		buf.WriteString(" exact")
	}
	fmt.Fprintf(buf, " %v, %v", inst.X, inst.Y.Ident())
	for _, md := range inst.Metadata {
		fmt.Fprintf(buf, ", %v", md)
	}
	return buf.String()
}

// ~~~ [ fdiv ] ~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~

// InstFDiv is an LLVM IR fdiv instruction.
type InstFDiv struct {
	// Name of local variable associated with the result.
	LocalName string
	// Operands.
	X, Y value.Value // floating-point scalar or floating-point vector

	// extra.

	// Type of result produced by the instruction.
	Typ types.Type
	// (optional) Fast math flags.
	FastMathFlags []enum.FastMathFlag
	// (optional) Metadata.
	Metadata []MetadataAttachment
}

// NewFDiv returns a new fdiv instruction based on the given operands.
func NewFDiv(x, y value.Value) *InstFDiv {
	return &InstFDiv{X: x, Y: y}
}

// String returns the LLVM syntax representation of the instruction as a
// type-value pair.
func (inst *InstFDiv) String() string {
	return fmt.Sprintf("%v %v", inst.Type(), inst.Ident())
}

// Type returns the type of the instruction.
func (inst *InstFDiv) Type() types.Type {
	// Cache type if not present.
	if inst.Typ == nil {
		inst.Typ = inst.X.Type()
	}
	return inst.Typ
}

// Ident returns the identifier associated with the instruction.
func (inst *InstFDiv) Ident() string {
	return enc.Local(inst.LocalName)
}

// Name returns the name of the instruction.
func (inst *InstFDiv) Name() string {
	return inst.LocalName
}

// SetName sets the name of the instruction.
func (inst *InstFDiv) SetName(name string) {
	inst.LocalName = name
}

// Def returns the LLVM syntax representation of the instruction.
func (inst *InstFDiv) Def() string {
	// "fdiv" FastMathFlags Type Value "," Value OptCommaSepMetadataAttachmentList
	buf := &strings.Builder{}
	buf.WriteString("fdiv")
	for _, flag := range inst.FastMathFlags {
		fmt.Fprintf(buf, " %v", flag)
	}
	fmt.Fprintf(buf, " %v, %v", inst.X, inst.Y.Ident())
	for _, md := range inst.Metadata {
		fmt.Fprintf(buf, ", %v", md)
	}
	return buf.String()
}

// ~~~ [ urem ] ~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~

// InstURem is an LLVM IR urem instruction.
type InstURem struct {
	// Name of local variable associated with the result.
	LocalName string
	// Operands.
	X, Y value.Value // integer scalar or integer vector

	// extra.

	// Type of result produced by the instruction.
	Typ types.Type
	// (optional) Metadata.
	Metadata []MetadataAttachment
}

// NewURem returns a new urem instruction based on the given operands.
func NewURem(x, y value.Value) *InstURem {
	return &InstURem{X: x, Y: y}
}

// String returns the LLVM syntax representation of the instruction as a
// type-value pair.
func (inst *InstURem) String() string {
	return fmt.Sprintf("%v %v", inst.Type(), inst.Ident())
}

// Type returns the type of the instruction.
func (inst *InstURem) Type() types.Type {
	// Cache type if not present.
	if inst.Typ == nil {
		inst.Typ = inst.X.Type()
	}
	return inst.Typ
}

// Ident returns the identifier associated with the instruction.
func (inst *InstURem) Ident() string {
	return enc.Local(inst.LocalName)
}

// Name returns the name of the instruction.
func (inst *InstURem) Name() string {
	return inst.LocalName
}

// SetName sets the name of the instruction.
func (inst *InstURem) SetName(name string) {
	inst.LocalName = name
}

// Def returns the LLVM syntax representation of the instruction.
func (inst *InstURem) Def() string {
	// "urem" Type Value "," Value OptCommaSepMetadataAttachmentList
	buf := &strings.Builder{}
	fmt.Fprintf(buf, "urem %v, %v", inst.X, inst.Y.Ident())
	for _, md := range inst.Metadata {
		fmt.Fprintf(buf, ", %v", md)
	}
	return buf.String()
}

// ~~~ [ srem ] ~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~

// InstSRem is an LLVM IR srem instruction.
type InstSRem struct {
	// Name of local variable associated with the result.
	LocalName string
	// Operands.
	X, Y value.Value // integer scalar or integer vector

	// extra.

	// Type of result produced by the instruction.
	Typ types.Type
	// (optional) Metadata.
	Metadata []MetadataAttachment
}

// NewSRem returns a new srem instruction based on the given operands.
func NewSRem(x, y value.Value) *InstSRem {
	return &InstSRem{X: x, Y: y}
}

// String returns the LLVM syntax representation of the instruction as a
// type-value pair.
func (inst *InstSRem) String() string {
	return fmt.Sprintf("%v %v", inst.Type(), inst.Ident())
}

// Type returns the type of the instruction.
func (inst *InstSRem) Type() types.Type {
	// Cache type if not present.
	if inst.Typ == nil {
		inst.Typ = inst.X.Type()
	}
	return inst.Typ
}

// Ident returns the identifier associated with the instruction.
func (inst *InstSRem) Ident() string {
	return enc.Local(inst.LocalName)
}

// Name returns the name of the instruction.
func (inst *InstSRem) Name() string {
	return inst.LocalName
}

// SetName sets the name of the instruction.
func (inst *InstSRem) SetName(name string) {
	inst.LocalName = name
}

// Def returns the LLVM syntax representation of the instruction.
func (inst *InstSRem) Def() string {
	// "srem" Type Value "," Value OptCommaSepMetadataAttachmentList
	buf := &strings.Builder{}
	fmt.Fprintf(buf, "srem %v, %v", inst.X, inst.Y.Ident())
	for _, md := range inst.Metadata {
		fmt.Fprintf(buf, ", %v", md)
	}
	return buf.String()
}

// ~~~ [ frem ] ~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~

// InstFRem is an LLVM IR frem instruction.
type InstFRem struct {
	// Name of local variable associated with the result.
	LocalName string
	// Operands.
	X, Y value.Value // floating-point scalar or floating-point vector

	// extra.

	// Type of result produced by the instruction.
	Typ types.Type
	// (optional) Fast math flags.
	FastMathFlags []enum.FastMathFlag
	// (optional) Metadata.
	Metadata []MetadataAttachment
}

// NewFRem returns a new frem instruction based on the given operands.
func NewFRem(x, y value.Value) *InstFRem {
	return &InstFRem{X: x, Y: y}
}

// String returns the LLVM syntax representation of the instruction as a
// type-value pair.
func (inst *InstFRem) String() string {
	return fmt.Sprintf("%v %v", inst.Type(), inst.Ident())
}

// Type returns the type of the instruction.
func (inst *InstFRem) Type() types.Type {
	// Cache type if not present.
	if inst.Typ == nil {
		inst.Typ = inst.X.Type()
	}
	return inst.Typ
}

// Ident returns the identifier associated with the instruction.
func (inst *InstFRem) Ident() string {
	return enc.Local(inst.LocalName)
}

// Name returns the name of the instruction.
func (inst *InstFRem) Name() string {
	return inst.LocalName
}

// SetName sets the name of the instruction.
func (inst *InstFRem) SetName(name string) {
	inst.LocalName = name
}

// Def returns the LLVM syntax representation of the instruction.
func (inst *InstFRem) Def() string {
	// "frem" FastMathFlags Type Value "," Value OptCommaSepMetadataAttachmentList
	buf := &strings.Builder{}
	buf.WriteString("frem")
	for _, flag := range inst.FastMathFlags {
		fmt.Fprintf(buf, " %v", flag)
	}
	fmt.Fprintf(buf, " %v, %v", inst.X, inst.Y.Ident())
	for _, md := range inst.Metadata {
		fmt.Fprintf(buf, ", %v", md)
	}
	return buf.String()
}
