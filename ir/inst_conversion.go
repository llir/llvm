package ir

import (
	"github.com/llir/l/ir/types"
	"github.com/llir/l/ir/value"
)

// --- [ Conversion instructions ] ---------------------------------------------

// ~~~ [ trunc ] ~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~

// Trunc is an LLVM IR trunc instruction.
type Trunc struct {
	// Name of local variable associated with the result.
	LocalName string
	// Value before conversion.
	From value.Value
	// Type after conversion.
	To types.Type
}

// NewTrunc returns a new trunc instruction based on the given source value and
// target type.
func NewTrunc(from value.Value, to types.Type) *Trunc {
	return &Trunc{From: from, To: to}
}

// ~~~ [ zext ] ~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~

// ZExt is an LLVM IR zext instruction.
type ZExt struct {
	// Name of local variable associated with the result.
	LocalName string
	// Value before conversion.
	From value.Value
	// Type after conversion.
	To types.Type
}

// NewZExt returns a new zext instruction based on the given source value and
// target type.
func NewZExt(from value.Value, to types.Type) *ZExt {
	return &ZExt{From: from, To: to}
}

// ~~~ [ sext ] ~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~

// SExt is an LLVM IR sext instruction.
type SExt struct {
	// Name of local variable associated with the result.
	LocalName string
	// Value before conversion.
	From value.Value
	// Type after conversion.
	To types.Type
}

// NewSExt returns a new sext instruction based on the given source value and
// target type.
func NewSExt(from value.Value, to types.Type) *SExt {
	return &SExt{From: from, To: to}
}

// ~~~ [ fptrunc ] ~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~

// FPTrunc is an LLVM IR fptrunc instruction.
type FPTrunc struct {
	// Name of local variable associated with the result.
	LocalName string
	// Value before conversion.
	From value.Value
	// Type after conversion.
	To types.Type
}

// NewFPTrunc returns a new fptrunc instruction based on the given source value
// and target type.
func NewFPTrunc(from value.Value, to types.Type) *FPTrunc {
	return &FPTrunc{From: from, To: to}
}

// ~~~ [ fpext ] ~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~

// FPExt is an LLVM IR fpext instruction.
type FPExt struct {
	// Name of local variable associated with the result.
	LocalName string
	// Value before conversion.
	From value.Value
	// Type after conversion.
	To types.Type
}

// NewFPExt returns a new fpext instruction based on the given source value and
// target type.
func NewFPExt(from value.Value, to types.Type) *FPExt {
	return &FPExt{From: from, To: to}
}

// ~~~ [ fptoui ] ~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~

// FPToUI is an LLVM IR fptoui instruction.
type FPToUI struct {
	// Name of local variable associated with the result.
	LocalName string
	// Value before conversion.
	From value.Value
	// Type after conversion.
	To types.Type
}

// NewFPToUI returns a new fptoui instruction based on the given source value
// and target type.
func NewFPToUI(from value.Value, to types.Type) *FPToUI {
	return &FPToUI{From: from, To: to}
}

// ~~~ [ fptosi ] ~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~

// FPToSI is an LLVM IR fptosi instruction.
type FPToSI struct {
	// Name of local variable associated with the result.
	LocalName string
	// Value before conversion.
	From value.Value
	// Type after conversion.
	To types.Type
}

// NewFPToSI returns a new fptosi instruction based on the given source value
// and target type.
func NewFPToSI(from value.Value, to types.Type) *FPToSI {
	return &FPToSI{From: from, To: to}
}

// ~~~ [ uitofp ] ~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~

// UIToFP is an LLVM IR uitofp instruction.
type UIToFP struct {
	// Name of local variable associated with the result.
	LocalName string
	// Value before conversion.
	From value.Value
	// Type after conversion.
	To types.Type
}

// NewUIToFP returns a new uitofp instruction based on the given source value
// and target type.
func NewUIToFP(from value.Value, to types.Type) *UIToFP {
	return &UIToFP{From: from, To: to}
}

// ~~~ [ sitofp ] ~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~

// SIToFP is an LLVM IR sitofp instruction.
type SIToFP struct {
	// Name of local variable associated with the result.
	LocalName string
	// Value before conversion.
	From value.Value
	// Type after conversion.
	To types.Type
}

// NewSIToFP returns a new sitofp instruction based on the given source value
// and target type.
func NewSIToFP(from value.Value, to types.Type) *SIToFP {
	return &SIToFP{From: from, To: to}
}

// ~~~ [ ptrtoint ] ~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~

// PtrToInt is an LLVM IR ptrtoint instruction.
type PtrToInt struct {
	// Name of local variable associated with the result.
	LocalName string
	// Value before conversion.
	From value.Value
	// Type after conversion.
	To types.Type
}

// NewPtrToInt returns a new ptrtoint instruction based on the given source
// value and target type.
func NewPtrToInt(from value.Value, to types.Type) *PtrToInt {
	return &PtrToInt{From: from, To: to}
}

// ~~~ [ inttoptr ] ~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~

// IntToPtr is an LLVM IR inttoptr instruction.
type IntToPtr struct {
	// Name of local variable associated with the result.
	LocalName string
	// Value before conversion.
	From value.Value
	// Type after conversion.
	To types.Type
}

// NewIntToPtr returns a new inttoptr instruction based on the given source
// value and target type.
func NewIntToPtr(from value.Value, to types.Type) *IntToPtr {
	return &IntToPtr{From: from, To: to}
}

// ~~~ [ bitcast ] ~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~

// BitCast is an LLVM IR bitcast instruction.
type BitCast struct {
	// Name of local variable associated with the result.
	LocalName string
	// Value before conversion.
	From value.Value
	// Type after conversion.
	To types.Type
}

// NewBitCast returns a new bitcast instruction based on the given source value
// and target type.
func NewBitCast(from value.Value, to types.Type) *BitCast {
	return &BitCast{From: from, To: to}
}

// ~~~ [ addrspacecast ] ~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~

// AddrSpaceCast is an LLVM IR addrspacecast instruction.
type AddrSpaceCast struct {
	// Name of local variable associated with the result.
	LocalName string
	// Value before conversion.
	From value.Value
	// Type after conversion.
	To types.Type
}

// NewAddrSpaceCast returns a new addrspacecast instruction based on the given
// source value and target type.
func NewAddrSpaceCast(from value.Value, to types.Type) *AddrSpaceCast {
	return &AddrSpaceCast{From: from, To: to}
}
