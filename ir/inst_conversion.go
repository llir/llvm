package ir

import (
	"fmt"

	"github.com/llir/l/internal/enc"
	"github.com/llir/l/ir/types"
	"github.com/llir/l/ir/value"
)

// --- [ Conversion instructions ] ---------------------------------------------

// ~~~ [ trunc ] ~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~

// InstTrunc is an LLVM IR trunc instruction.
type InstTrunc struct {
	// Name of local variable associated with the result.
	LocalName string
	// Value before conversion.
	From value.Value
	// Type after conversion.
	To types.Type

	// extra.

	// (optional) Metadata.
	// TODO: add metadata.
}

// NewTrunc returns a new trunc instruction based on the given source value and
// target type.
func NewTrunc(from value.Value, to types.Type) *InstTrunc {
	return &InstTrunc{From: from, To: to}
}

// String returns the LLVM syntax representation of the instruction as a
// type-value pair.
func (inst *InstTrunc) String() string {
	return fmt.Sprintf("%v %v", inst.Type(), inst.Ident())
}

// Type returns the type of the instruction.
func (inst *InstTrunc) Type() types.Type {
	return inst.To
}

// Ident returns the identifier associated with the instruction.
func (inst *InstTrunc) Ident() string {
	return enc.Local(inst.LocalName)
}

// Name returns the name of the instruction.
func (inst *InstTrunc) Name() string {
	return inst.LocalName
}

// SetName sets the name of the instruction.
func (inst *InstTrunc) SetName(name string) {
	inst.LocalName = name
}

// ~~~ [ zext ] ~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~

// InstZExt is an LLVM IR zext instruction.
type InstZExt struct {
	// Name of local variable associated with the result.
	LocalName string
	// Value before conversion.
	From value.Value
	// Type after conversion.
	To types.Type

	// extra.

	// (optional) Metadata.
	// TODO: add metadata.
}

// NewZExt returns a new zext instruction based on the given source value and
// target type.
func NewZExt(from value.Value, to types.Type) *InstZExt {
	return &InstZExt{From: from, To: to}
}

// String returns the LLVM syntax representation of the instruction as a
// type-value pair.
func (inst *InstZExt) String() string {
	return fmt.Sprintf("%v %v", inst.Type(), inst.Ident())
}

// Type returns the type of the instruction.
func (inst *InstZExt) Type() types.Type {
	return inst.To
}

// Ident returns the identifier associated with the instruction.
func (inst *InstZExt) Ident() string {
	return enc.Local(inst.LocalName)
}

// Name returns the name of the instruction.
func (inst *InstZExt) Name() string {
	return inst.LocalName
}

// SetName sets the name of the instruction.
func (inst *InstZExt) SetName(name string) {
	inst.LocalName = name
}

// ~~~ [ sext ] ~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~

// InstSExt is an LLVM IR sext instruction.
type InstSExt struct {
	// Name of local variable associated with the result.
	LocalName string
	// Value before conversion.
	From value.Value
	// Type after conversion.
	To types.Type

	// extra.

	// (optional) Metadata.
	// TODO: add metadata.
}

// NewSExt returns a new sext instruction based on the given source value and
// target type.
func NewSExt(from value.Value, to types.Type) *InstSExt {
	return &InstSExt{From: from, To: to}
}

// String returns the LLVM syntax representation of the instruction as a
// type-value pair.
func (inst *InstSExt) String() string {
	return fmt.Sprintf("%v %v", inst.Type(), inst.Ident())
}

// Type returns the type of the instruction.
func (inst *InstSExt) Type() types.Type {
	return inst.To
}

// Ident returns the identifier associated with the instruction.
func (inst *InstSExt) Ident() string {
	return enc.Local(inst.LocalName)
}

// Name returns the name of the instruction.
func (inst *InstSExt) Name() string {
	return inst.LocalName
}

// SetName sets the name of the instruction.
func (inst *InstSExt) SetName(name string) {
	inst.LocalName = name
}

// ~~~ [ fptrunc ] ~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~

// InstFPTrunc is an LLVM IR fptrunc instruction.
type InstFPTrunc struct {
	// Name of local variable associated with the result.
	LocalName string
	// Value before conversion.
	From value.Value
	// Type after conversion.
	To types.Type

	// extra.

	// (optional) Metadata.
	// TODO: add metadata.
}

// NewFPTrunc returns a new fptrunc instruction based on the given source value
// and target type.
func NewFPTrunc(from value.Value, to types.Type) *InstFPTrunc {
	return &InstFPTrunc{From: from, To: to}
}

// String returns the LLVM syntax representation of the instruction as a
// type-value pair.
func (inst *InstFPTrunc) String() string {
	return fmt.Sprintf("%v %v", inst.Type(), inst.Ident())
}

// Type returns the type of the instruction.
func (inst *InstFPTrunc) Type() types.Type {
	return inst.To
}

// Ident returns the identifier associated with the instruction.
func (inst *InstFPTrunc) Ident() string {
	return enc.Local(inst.LocalName)
}

// Name returns the name of the instruction.
func (inst *InstFPTrunc) Name() string {
	return inst.LocalName
}

// SetName sets the name of the instruction.
func (inst *InstFPTrunc) SetName(name string) {
	inst.LocalName = name
}

// ~~~ [ fpext ] ~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~

// InstFPExt is an LLVM IR fpext instruction.
type InstFPExt struct {
	// Name of local variable associated with the result.
	LocalName string
	// Value before conversion.
	From value.Value
	// Type after conversion.
	To types.Type

	// extra.

	// (optional) Metadata.
	// TODO: add metadata.
}

// NewFPExt returns a new fpext instruction based on the given source value and
// target type.
func NewFPExt(from value.Value, to types.Type) *InstFPExt {
	return &InstFPExt{From: from, To: to}
}

// String returns the LLVM syntax representation of the instruction as a
// type-value pair.
func (inst *InstFPExt) String() string {
	return fmt.Sprintf("%v %v", inst.Type(), inst.Ident())
}

// Type returns the type of the instruction.
func (inst *InstFPExt) Type() types.Type {
	return inst.To
}

// Ident returns the identifier associated with the instruction.
func (inst *InstFPExt) Ident() string {
	return enc.Local(inst.LocalName)
}

// Name returns the name of the instruction.
func (inst *InstFPExt) Name() string {
	return inst.LocalName
}

// SetName sets the name of the instruction.
func (inst *InstFPExt) SetName(name string) {
	inst.LocalName = name
}

// ~~~ [ fptoui ] ~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~

// InstFPToUI is an LLVM IR fptoui instruction.
type InstFPToUI struct {
	// Name of local variable associated with the result.
	LocalName string
	// Value before conversion.
	From value.Value
	// Type after conversion.
	To types.Type

	// extra.

	// (optional) Metadata.
	// TODO: add metadata.
}

// NewFPToUI returns a new fptoui instruction based on the given source value
// and target type.
func NewFPToUI(from value.Value, to types.Type) *InstFPToUI {
	return &InstFPToUI{From: from, To: to}
}

// String returns the LLVM syntax representation of the instruction as a
// type-value pair.
func (inst *InstFPToUI) String() string {
	return fmt.Sprintf("%v %v", inst.Type(), inst.Ident())
}

// Type returns the type of the instruction.
func (inst *InstFPToUI) Type() types.Type {
	return inst.To
}

// Ident returns the identifier associated with the instruction.
func (inst *InstFPToUI) Ident() string {
	return enc.Local(inst.LocalName)
}

// Name returns the name of the instruction.
func (inst *InstFPToUI) Name() string {
	return inst.LocalName
}

// SetName sets the name of the instruction.
func (inst *InstFPToUI) SetName(name string) {
	inst.LocalName = name
}

// ~~~ [ fptosi ] ~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~

// InstFPToSI is an LLVM IR fptosi instruction.
type InstFPToSI struct {
	// Name of local variable associated with the result.
	LocalName string
	// Value before conversion.
	From value.Value
	// Type after conversion.
	To types.Type

	// extra.

	// (optional) Metadata.
	// TODO: add metadata.
}

// NewFPToSI returns a new fptosi instruction based on the given source value
// and target type.
func NewFPToSI(from value.Value, to types.Type) *InstFPToSI {
	return &InstFPToSI{From: from, To: to}
}

// String returns the LLVM syntax representation of the instruction as a
// type-value pair.
func (inst *InstFPToSI) String() string {
	return fmt.Sprintf("%v %v", inst.Type(), inst.Ident())
}

// Type returns the type of the instruction.
func (inst *InstFPToSI) Type() types.Type {
	return inst.To
}

// Ident returns the identifier associated with the instruction.
func (inst *InstFPToSI) Ident() string {
	return enc.Local(inst.LocalName)
}

// Name returns the name of the instruction.
func (inst *InstFPToSI) Name() string {
	return inst.LocalName
}

// SetName sets the name of the instruction.
func (inst *InstFPToSI) SetName(name string) {
	inst.LocalName = name
}

// ~~~ [ uitofp ] ~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~

// InstUIToFP is an LLVM IR uitofp instruction.
type InstUIToFP struct {
	// Name of local variable associated with the result.
	LocalName string
	// Value before conversion.
	From value.Value
	// Type after conversion.
	To types.Type

	// extra.

	// (optional) Metadata.
	// TODO: add metadata.
}

// NewUIToFP returns a new uitofp instruction based on the given source value
// and target type.
func NewUIToFP(from value.Value, to types.Type) *InstUIToFP {
	return &InstUIToFP{From: from, To: to}
}

// String returns the LLVM syntax representation of the instruction as a
// type-value pair.
func (inst *InstUIToFP) String() string {
	return fmt.Sprintf("%v %v", inst.Type(), inst.Ident())
}

// Type returns the type of the instruction.
func (inst *InstUIToFP) Type() types.Type {
	return inst.To
}

// Ident returns the identifier associated with the instruction.
func (inst *InstUIToFP) Ident() string {
	return enc.Local(inst.LocalName)
}

// Name returns the name of the instruction.
func (inst *InstUIToFP) Name() string {
	return inst.LocalName
}

// SetName sets the name of the instruction.
func (inst *InstUIToFP) SetName(name string) {
	inst.LocalName = name
}

// ~~~ [ sitofp ] ~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~

// InstSIToFP is an LLVM IR sitofp instruction.
type InstSIToFP struct {
	// Name of local variable associated with the result.
	LocalName string
	// Value before conversion.
	From value.Value
	// Type after conversion.
	To types.Type

	// extra.

	// (optional) Metadata.
	// TODO: add metadata.
}

// NewSIToFP returns a new sitofp instruction based on the given source value
// and target type.
func NewSIToFP(from value.Value, to types.Type) *InstSIToFP {
	return &InstSIToFP{From: from, To: to}
}

// String returns the LLVM syntax representation of the instruction as a
// type-value pair.
func (inst *InstSIToFP) String() string {
	return fmt.Sprintf("%v %v", inst.Type(), inst.Ident())
}

// Type returns the type of the instruction.
func (inst *InstSIToFP) Type() types.Type {
	return inst.To
}

// Ident returns the identifier associated with the instruction.
func (inst *InstSIToFP) Ident() string {
	return enc.Local(inst.LocalName)
}

// Name returns the name of the instruction.
func (inst *InstSIToFP) Name() string {
	return inst.LocalName
}

// SetName sets the name of the instruction.
func (inst *InstSIToFP) SetName(name string) {
	inst.LocalName = name
}

// ~~~ [ ptrtoint ] ~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~

// InstPtrToInt is an LLVM IR ptrtoint instruction.
type InstPtrToInt struct {
	// Name of local variable associated with the result.
	LocalName string
	// Value before conversion.
	From value.Value
	// Type after conversion.
	To types.Type

	// extra.

	// (optional) Metadata.
	// TODO: add metadata.
}

// NewPtrToInt returns a new ptrtoint instruction based on the given source
// value and target type.
func NewPtrToInt(from value.Value, to types.Type) *InstPtrToInt {
	return &InstPtrToInt{From: from, To: to}
}

// String returns the LLVM syntax representation of the instruction as a
// type-value pair.
func (inst *InstPtrToInt) String() string {
	return fmt.Sprintf("%v %v", inst.Type(), inst.Ident())
}

// Type returns the type of the instruction.
func (inst *InstPtrToInt) Type() types.Type {
	return inst.To
}

// Ident returns the identifier associated with the instruction.
func (inst *InstPtrToInt) Ident() string {
	return enc.Local(inst.LocalName)
}

// Name returns the name of the instruction.
func (inst *InstPtrToInt) Name() string {
	return inst.LocalName
}

// SetName sets the name of the instruction.
func (inst *InstPtrToInt) SetName(name string) {
	inst.LocalName = name
}

// ~~~ [ inttoptr ] ~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~

// InstIntToPtr is an LLVM IR inttoptr instruction.
type InstIntToPtr struct {
	// Name of local variable associated with the result.
	LocalName string
	// Value before conversion.
	From value.Value
	// Type after conversion.
	To types.Type

	// extra.

	// (optional) Metadata.
	// TODO: add metadata.
}

// NewIntToPtr returns a new inttoptr instruction based on the given source
// value and target type.
func NewIntToPtr(from value.Value, to types.Type) *InstIntToPtr {
	return &InstIntToPtr{From: from, To: to}
}

// String returns the LLVM syntax representation of the instruction as a
// type-value pair.
func (inst *InstIntToPtr) String() string {
	return fmt.Sprintf("%v %v", inst.Type(), inst.Ident())
}

// Type returns the type of the instruction.
func (inst *InstIntToPtr) Type() types.Type {
	return inst.To
}

// Ident returns the identifier associated with the instruction.
func (inst *InstIntToPtr) Ident() string {
	return enc.Local(inst.LocalName)
}

// Name returns the name of the instruction.
func (inst *InstIntToPtr) Name() string {
	return inst.LocalName
}

// SetName sets the name of the instruction.
func (inst *InstIntToPtr) SetName(name string) {
	inst.LocalName = name
}

// ~~~ [ bitcast ] ~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~

// InstBitCast is an LLVM IR bitcast instruction.
type InstBitCast struct {
	// Name of local variable associated with the result.
	LocalName string
	// Value before conversion.
	From value.Value
	// Type after conversion.
	To types.Type

	// extra.

	// (optional) Metadata.
	// TODO: add metadata.
}

// NewBitCast returns a new bitcast instruction based on the given source value
// and target type.
func NewBitCast(from value.Value, to types.Type) *InstBitCast {
	return &InstBitCast{From: from, To: to}
}

// String returns the LLVM syntax representation of the instruction as a
// type-value pair.
func (inst *InstBitCast) String() string {
	return fmt.Sprintf("%v %v", inst.Type(), inst.Ident())
}

// Type returns the type of the instruction.
func (inst *InstBitCast) Type() types.Type {
	return inst.To
}

// Ident returns the identifier associated with the instruction.
func (inst *InstBitCast) Ident() string {
	return enc.Local(inst.LocalName)
}

// Name returns the name of the instruction.
func (inst *InstBitCast) Name() string {
	return inst.LocalName
}

// SetName sets the name of the instruction.
func (inst *InstBitCast) SetName(name string) {
	inst.LocalName = name
}

// ~~~ [ addrspacecast ] ~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~

// InstAddrSpaceCast is an LLVM IR addrspacecast instruction.
type InstAddrSpaceCast struct {
	// Name of local variable associated with the result.
	LocalName string
	// Value before conversion.
	From value.Value
	// Type after conversion.
	To types.Type

	// extra.

	// (optional) Metadata.
	// TODO: add metadata.
}

// NewAddrSpaceCast returns a new addrspacecast instruction based on the given
// source value and target type.
func NewAddrSpaceCast(from value.Value, to types.Type) *InstAddrSpaceCast {
	return &InstAddrSpaceCast{From: from, To: to}
}

// String returns the LLVM syntax representation of the instruction as a
// type-value pair.
func (inst *InstAddrSpaceCast) String() string {
	return fmt.Sprintf("%v %v", inst.Type(), inst.Ident())
}

// Type returns the type of the instruction.
func (inst *InstAddrSpaceCast) Type() types.Type {
	return inst.To
}

// Ident returns the identifier associated with the instruction.
func (inst *InstAddrSpaceCast) Ident() string {
	return enc.Local(inst.LocalName)
}

// Name returns the name of the instruction.
func (inst *InstAddrSpaceCast) Name() string {
	return inst.LocalName
}

// SetName sets the name of the instruction.
func (inst *InstAddrSpaceCast) SetName(name string) {
	inst.LocalName = name
}
