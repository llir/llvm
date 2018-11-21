package ir

import (
	"fmt"
	"strings"

	"github.com/llir/llvm/ir/metadata"
	"github.com/llir/llvm/ir/types"
	"github.com/llir/llvm/ir/value"
)

// --- [ Conversion instructions ] ---------------------------------------------

// ~~~ [ trunc ] ~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~

// InstTrunc is an LLVM IR trunc instruction.
type InstTrunc struct {
	// Name of local variable associated with the result.
	LocalIdent
	// Value before conversion.
	From value.Value
	// Type after conversion.
	To types.Type

	// extra.

	// (optional) Metadata.
	Metadata []*metadata.MetadataAttachment
}

// NewTrunc returns a new trunc instruction based on the given source value and
// target type.
func NewTrunc(from value.Value, to types.Type) *InstTrunc {
	return &InstTrunc{From: from, To: to}
}

// String returns the LLVM syntax representation of the instruction as a
// type-value pair.
func (inst *InstTrunc) String() string {
	return fmt.Sprintf("%s %s", inst.Type(), inst.Ident())
}

// Type returns the type of the instruction.
func (inst *InstTrunc) Type() types.Type {
	return inst.To
}

// Def returns the LLVM syntax representation of the instruction.
func (inst *InstTrunc) Def() string {
	// 'trunc' From=TypeValue 'to' To=Type Metadata=(',' MetadataAttachment)+?
	buf := &strings.Builder{}
	fmt.Fprintf(buf, "%s = ", inst.Ident())
	fmt.Fprintf(buf, "trunc %s to %s", inst.From, inst.To)
	for _, md := range inst.Metadata {
		fmt.Fprintf(buf, ", %s", md)
	}
	return buf.String()
}

// ~~~ [ zext ] ~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~

// InstZExt is an LLVM IR zext instruction.
type InstZExt struct {
	// Name of local variable associated with the result.
	LocalIdent
	// Value before conversion.
	From value.Value
	// Type after conversion.
	To types.Type

	// extra.

	// (optional) Metadata.
	Metadata []*metadata.MetadataAttachment
}

// NewZExt returns a new zext instruction based on the given source value and
// target type.
func NewZExt(from value.Value, to types.Type) *InstZExt {
	return &InstZExt{From: from, To: to}
}

// String returns the LLVM syntax representation of the instruction as a
// type-value pair.
func (inst *InstZExt) String() string {
	return fmt.Sprintf("%s %s", inst.Type(), inst.Ident())
}

// Type returns the type of the instruction.
func (inst *InstZExt) Type() types.Type {
	return inst.To
}

// Def returns the LLVM syntax representation of the instruction.
func (inst *InstZExt) Def() string {
	// 'zext' From=TypeValue 'to' To=Type Metadata=(',' MetadataAttachment)+?
	buf := &strings.Builder{}
	fmt.Fprintf(buf, "%s = ", inst.Ident())
	fmt.Fprintf(buf, "zext %s to %s", inst.From, inst.To)
	for _, md := range inst.Metadata {
		fmt.Fprintf(buf, ", %s", md)
	}
	return buf.String()
}

// ~~~ [ sext ] ~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~

// InstSExt is an LLVM IR sext instruction.
type InstSExt struct {
	// Name of local variable associated with the result.
	LocalIdent
	// Value before conversion.
	From value.Value
	// Type after conversion.
	To types.Type

	// extra.

	// (optional) Metadata.
	Metadata []*metadata.MetadataAttachment
}

// NewSExt returns a new sext instruction based on the given source value and
// target type.
func NewSExt(from value.Value, to types.Type) *InstSExt {
	return &InstSExt{From: from, To: to}
}

// String returns the LLVM syntax representation of the instruction as a
// type-value pair.
func (inst *InstSExt) String() string {
	return fmt.Sprintf("%s %s", inst.Type(), inst.Ident())
}

// Type returns the type of the instruction.
func (inst *InstSExt) Type() types.Type {
	return inst.To
}

// Def returns the LLVM syntax representation of the instruction.
func (inst *InstSExt) Def() string {
	// 'sext' From=TypeValue 'to' To=Type Metadata=(',' MetadataAttachment)+?
	buf := &strings.Builder{}
	fmt.Fprintf(buf, "%s = ", inst.Ident())
	fmt.Fprintf(buf, "sext %s to %s", inst.From, inst.To)
	for _, md := range inst.Metadata {
		fmt.Fprintf(buf, ", %s", md)
	}
	return buf.String()
}

// ~~~ [ fptrunc ] ~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~

// InstFPTrunc is an LLVM IR fptrunc instruction.
type InstFPTrunc struct {
	// Name of local variable associated with the result.
	LocalIdent
	// Value before conversion.
	From value.Value
	// Type after conversion.
	To types.Type

	// extra.

	// (optional) Metadata.
	Metadata []*metadata.MetadataAttachment
}

// NewFPTrunc returns a new fptrunc instruction based on the given source value
// and target type.
func NewFPTrunc(from value.Value, to types.Type) *InstFPTrunc {
	return &InstFPTrunc{From: from, To: to}
}

// String returns the LLVM syntax representation of the instruction as a
// type-value pair.
func (inst *InstFPTrunc) String() string {
	return fmt.Sprintf("%s %s", inst.Type(), inst.Ident())
}

// Type returns the type of the instruction.
func (inst *InstFPTrunc) Type() types.Type {
	return inst.To
}

// Def returns the LLVM syntax representation of the instruction.
func (inst *InstFPTrunc) Def() string {
	// 'fptrunc' From=TypeValue 'to' To=Type Metadata=(',' MetadataAttachment)+?
	buf := &strings.Builder{}
	fmt.Fprintf(buf, "%s = ", inst.Ident())
	fmt.Fprintf(buf, "fptrunc %s to %s", inst.From, inst.To)
	for _, md := range inst.Metadata {
		fmt.Fprintf(buf, ", %s", md)
	}
	return buf.String()
}

// ~~~ [ fpext ] ~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~

// InstFPExt is an LLVM IR fpext instruction.
type InstFPExt struct {
	// Name of local variable associated with the result.
	LocalIdent
	// Value before conversion.
	From value.Value
	// Type after conversion.
	To types.Type

	// extra.

	// (optional) Metadata.
	Metadata []*metadata.MetadataAttachment
}

// NewFPExt returns a new fpext instruction based on the given source value and
// target type.
func NewFPExt(from value.Value, to types.Type) *InstFPExt {
	return &InstFPExt{From: from, To: to}
}

// String returns the LLVM syntax representation of the instruction as a
// type-value pair.
func (inst *InstFPExt) String() string {
	return fmt.Sprintf("%s %s", inst.Type(), inst.Ident())
}

// Type returns the type of the instruction.
func (inst *InstFPExt) Type() types.Type {
	return inst.To
}

// Def returns the LLVM syntax representation of the instruction.
func (inst *InstFPExt) Def() string {
	// 'fpext' From=TypeValue 'to' To=Type Metadata=(',' MetadataAttachment)+?
	buf := &strings.Builder{}
	fmt.Fprintf(buf, "%s = ", inst.Ident())
	fmt.Fprintf(buf, "fpext %s to %s", inst.From, inst.To)
	for _, md := range inst.Metadata {
		fmt.Fprintf(buf, ", %s", md)
	}
	return buf.String()
}

// ~~~ [ fptoui ] ~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~

// InstFPToUI is an LLVM IR fptoui instruction.
type InstFPToUI struct {
	// Name of local variable associated with the result.
	LocalIdent
	// Value before conversion.
	From value.Value
	// Type after conversion.
	To types.Type

	// extra.

	// (optional) Metadata.
	Metadata []*metadata.MetadataAttachment
}

// NewFPToUI returns a new fptoui instruction based on the given source value
// and target type.
func NewFPToUI(from value.Value, to types.Type) *InstFPToUI {
	return &InstFPToUI{From: from, To: to}
}

// String returns the LLVM syntax representation of the instruction as a
// type-value pair.
func (inst *InstFPToUI) String() string {
	return fmt.Sprintf("%s %s", inst.Type(), inst.Ident())
}

// Type returns the type of the instruction.
func (inst *InstFPToUI) Type() types.Type {
	return inst.To
}

// Def returns the LLVM syntax representation of the instruction.
func (inst *InstFPToUI) Def() string {
	// 'fptoui' From=TypeValue 'to' To=Type Metadata=(',' MetadataAttachment)+?
	buf := &strings.Builder{}
	fmt.Fprintf(buf, "%s = ", inst.Ident())
	fmt.Fprintf(buf, "fptoui %s to %s", inst.From, inst.To)
	for _, md := range inst.Metadata {
		fmt.Fprintf(buf, ", %s", md)
	}
	return buf.String()
}

// ~~~ [ fptosi ] ~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~

// InstFPToSI is an LLVM IR fptosi instruction.
type InstFPToSI struct {
	// Name of local variable associated with the result.
	LocalIdent
	// Value before conversion.
	From value.Value
	// Type after conversion.
	To types.Type

	// extra.

	// (optional) Metadata.
	Metadata []*metadata.MetadataAttachment
}

// NewFPToSI returns a new fptosi instruction based on the given source value
// and target type.
func NewFPToSI(from value.Value, to types.Type) *InstFPToSI {
	return &InstFPToSI{From: from, To: to}
}

// String returns the LLVM syntax representation of the instruction as a
// type-value pair.
func (inst *InstFPToSI) String() string {
	return fmt.Sprintf("%s %s", inst.Type(), inst.Ident())
}

// Type returns the type of the instruction.
func (inst *InstFPToSI) Type() types.Type {
	return inst.To
}

// Def returns the LLVM syntax representation of the instruction.
func (inst *InstFPToSI) Def() string {
	// 'fptosi' From=TypeValue 'to' To=Type Metadata=(',' MetadataAttachment)+?
	buf := &strings.Builder{}
	fmt.Fprintf(buf, "%s = ", inst.Ident())
	fmt.Fprintf(buf, "fptosi %s to %s", inst.From, inst.To)
	for _, md := range inst.Metadata {
		fmt.Fprintf(buf, ", %s", md)
	}
	return buf.String()
}

// ~~~ [ uitofp ] ~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~

// InstUIToFP is an LLVM IR uitofp instruction.
type InstUIToFP struct {
	// Name of local variable associated with the result.
	LocalIdent
	// Value before conversion.
	From value.Value
	// Type after conversion.
	To types.Type

	// extra.

	// (optional) Metadata.
	Metadata []*metadata.MetadataAttachment
}

// NewUIToFP returns a new uitofp instruction based on the given source value
// and target type.
func NewUIToFP(from value.Value, to types.Type) *InstUIToFP {
	return &InstUIToFP{From: from, To: to}
}

// String returns the LLVM syntax representation of the instruction as a
// type-value pair.
func (inst *InstUIToFP) String() string {
	return fmt.Sprintf("%s %s", inst.Type(), inst.Ident())
}

// Type returns the type of the instruction.
func (inst *InstUIToFP) Type() types.Type {
	return inst.To
}

// Def returns the LLVM syntax representation of the instruction.
func (inst *InstUIToFP) Def() string {
	// 'uitofp' From=TypeValue 'to' To=Type Metadata=(',' MetadataAttachment)+?
	buf := &strings.Builder{}
	fmt.Fprintf(buf, "%s = ", inst.Ident())
	fmt.Fprintf(buf, "uitofp %s to %s", inst.From, inst.To)
	for _, md := range inst.Metadata {
		fmt.Fprintf(buf, ", %s", md)
	}
	return buf.String()
}

// ~~~ [ sitofp ] ~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~

// InstSIToFP is an LLVM IR sitofp instruction.
type InstSIToFP struct {
	// Name of local variable associated with the result.
	LocalIdent
	// Value before conversion.
	From value.Value
	// Type after conversion.
	To types.Type

	// extra.

	// (optional) Metadata.
	Metadata []*metadata.MetadataAttachment
}

// NewSIToFP returns a new sitofp instruction based on the given source value
// and target type.
func NewSIToFP(from value.Value, to types.Type) *InstSIToFP {
	return &InstSIToFP{From: from, To: to}
}

// String returns the LLVM syntax representation of the instruction as a
// type-value pair.
func (inst *InstSIToFP) String() string {
	return fmt.Sprintf("%s %s", inst.Type(), inst.Ident())
}

// Type returns the type of the instruction.
func (inst *InstSIToFP) Type() types.Type {
	return inst.To
}

// Def returns the LLVM syntax representation of the instruction.
func (inst *InstSIToFP) Def() string {
	// 'sitofp' From=TypeValue 'to' To=Type Metadata=(',' MetadataAttachment)+?
	buf := &strings.Builder{}
	fmt.Fprintf(buf, "%s = ", inst.Ident())
	fmt.Fprintf(buf, "sitofp %s to %s", inst.From, inst.To)
	for _, md := range inst.Metadata {
		fmt.Fprintf(buf, ", %s", md)
	}
	return buf.String()
}

// ~~~ [ ptrtoint ] ~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~

// InstPtrToInt is an LLVM IR ptrtoint instruction.
type InstPtrToInt struct {
	// Name of local variable associated with the result.
	LocalIdent
	// Value before conversion.
	From value.Value
	// Type after conversion.
	To types.Type

	// extra.

	// (optional) Metadata.
	Metadata []*metadata.MetadataAttachment
}

// NewPtrToInt returns a new ptrtoint instruction based on the given source
// value and target type.
func NewPtrToInt(from value.Value, to types.Type) *InstPtrToInt {
	return &InstPtrToInt{From: from, To: to}
}

// String returns the LLVM syntax representation of the instruction as a
// type-value pair.
func (inst *InstPtrToInt) String() string {
	return fmt.Sprintf("%s %s", inst.Type(), inst.Ident())
}

// Type returns the type of the instruction.
func (inst *InstPtrToInt) Type() types.Type {
	return inst.To
}

// Def returns the LLVM syntax representation of the instruction.
func (inst *InstPtrToInt) Def() string {
	// 'ptrtoint' From=TypeValue 'to' To=Type Metadata=(',' MetadataAttachment)+?
	buf := &strings.Builder{}
	fmt.Fprintf(buf, "%s = ", inst.Ident())
	fmt.Fprintf(buf, "ptrtoint %s to %s", inst.From, inst.To)
	for _, md := range inst.Metadata {
		fmt.Fprintf(buf, ", %s", md)
	}
	return buf.String()
}

// ~~~ [ inttoptr ] ~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~

// InstIntToPtr is an LLVM IR inttoptr instruction.
type InstIntToPtr struct {
	// Name of local variable associated with the result.
	LocalIdent
	// Value before conversion.
	From value.Value
	// Type after conversion.
	To types.Type

	// extra.

	// (optional) Metadata.
	Metadata []*metadata.MetadataAttachment
}

// NewIntToPtr returns a new inttoptr instruction based on the given source
// value and target type.
func NewIntToPtr(from value.Value, to types.Type) *InstIntToPtr {
	return &InstIntToPtr{From: from, To: to}
}

// String returns the LLVM syntax representation of the instruction as a
// type-value pair.
func (inst *InstIntToPtr) String() string {
	return fmt.Sprintf("%s %s", inst.Type(), inst.Ident())
}

// Type returns the type of the instruction.
func (inst *InstIntToPtr) Type() types.Type {
	return inst.To
}

// Def returns the LLVM syntax representation of the instruction.
func (inst *InstIntToPtr) Def() string {
	// 'inttoptr' From=TypeValue 'to' To=Type Metadata=(',' MetadataAttachment)+?
	buf := &strings.Builder{}
	fmt.Fprintf(buf, "%s = ", inst.Ident())
	fmt.Fprintf(buf, "inttoptr %s to %s", inst.From, inst.To)
	for _, md := range inst.Metadata {
		fmt.Fprintf(buf, ", %s", md)
	}
	return buf.String()
}

// ~~~ [ bitcast ] ~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~

// InstBitCast is an LLVM IR bitcast instruction.
type InstBitCast struct {
	// Name of local variable associated with the result.
	LocalIdent
	// Value before conversion.
	From value.Value
	// Type after conversion.
	To types.Type

	// extra.

	// (optional) Metadata.
	Metadata []*metadata.MetadataAttachment
}

// NewBitCast returns a new bitcast instruction based on the given source value
// and target type.
func NewBitCast(from value.Value, to types.Type) *InstBitCast {
	return &InstBitCast{From: from, To: to}
}

// String returns the LLVM syntax representation of the instruction as a
// type-value pair.
func (inst *InstBitCast) String() string {
	return fmt.Sprintf("%s %s", inst.Type(), inst.Ident())
}

// Type returns the type of the instruction.
func (inst *InstBitCast) Type() types.Type {
	return inst.To
}

// Def returns the LLVM syntax representation of the instruction.
func (inst *InstBitCast) Def() string {
	// 'bitcast' From=TypeValue 'to' To=Type Metadata=(',' MetadataAttachment)+?
	buf := &strings.Builder{}
	fmt.Fprintf(buf, "%s = ", inst.Ident())
	fmt.Fprintf(buf, "bitcast %s to %s", inst.From, inst.To)
	for _, md := range inst.Metadata {
		fmt.Fprintf(buf, ", %s", md)
	}
	return buf.String()
}

// ~~~ [ addrspacecast ] ~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~

// InstAddrSpaceCast is an LLVM IR addrspacecast instruction.
type InstAddrSpaceCast struct {
	// Name of local variable associated with the result.
	LocalIdent
	// Value before conversion.
	From value.Value
	// Type after conversion.
	To types.Type

	// extra.

	// (optional) Metadata.
	Metadata []*metadata.MetadataAttachment
}

// NewAddrSpaceCast returns a new addrspacecast instruction based on the given
// source value and target type.
func NewAddrSpaceCast(from value.Value, to types.Type) *InstAddrSpaceCast {
	return &InstAddrSpaceCast{From: from, To: to}
}

// String returns the LLVM syntax representation of the instruction as a
// type-value pair.
func (inst *InstAddrSpaceCast) String() string {
	return fmt.Sprintf("%s %s", inst.Type(), inst.Ident())
}

// Type returns the type of the instruction.
func (inst *InstAddrSpaceCast) Type() types.Type {
	return inst.To
}

// Def returns the LLVM syntax representation of the instruction.
func (inst *InstAddrSpaceCast) Def() string {
	// 'addrspacecast' From=TypeValue 'to' To=Type Metadata=(','
	// MetadataAttachment)+?
	buf := &strings.Builder{}
	fmt.Fprintf(buf, "%s = ", inst.Ident())
	fmt.Fprintf(buf, "addrspacecast %s to %s", inst.From, inst.To)
	for _, md := range inst.Metadata {
		fmt.Fprintf(buf, ", %s", md)
	}
	return buf.String()
}
