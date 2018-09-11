package constant

import (
	"github.com/llir/l/ir/types"
)

// --- [ Conversion expressions ] ----------------------------------------------

// ~~~ [ trunc ] ~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~

// Trunc is an LLVM IR trunc expression.
type Trunc struct {
	// Value before conversion.
	From Constant
	// Type after conversion.
	To types.Type
}

// NewTrunc returns a new trunc expression based on the given source value and
// target type.
func NewTrunc(from Constant, to types.Type) *Trunc {
	return &Trunc{From: from, To: to}
}

// ~~~ [ zext ] ~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~

// ZExt is an LLVM IR zext expression.
type ZExt struct {
	// Value before conversion.
	From Constant
	// Type after conversion.
	To types.Type
}

// NewZExt returns a new zext expression based on the given source value and
// target type.
func NewZExt(from Constant, to types.Type) *ZExt {
	return &ZExt{From: from, To: to}
}

// ~~~ [ sext ] ~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~

// SExt is an LLVM IR sext expression.
type SExt struct {
	// Value before conversion.
	From Constant
	// Type after conversion.
	To types.Type
}

// NewSExt returns a new sext expression based on the given source value and
// target type.
func NewSExt(from Constant, to types.Type) *SExt {
	return &SExt{From: from, To: to}
}

// ~~~ [ fptrunc ] ~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~

// FPTrunc is an LLVM IR fptrunc expression.
type FPTrunc struct {
	// Value before conversion.
	From Constant
	// Type after conversion.
	To types.Type
}

// NewFPTrunc returns a new fptrunc expression based on the given source value
// and target type.
func NewFPTrunc(from Constant, to types.Type) *FPTrunc {
	return &FPTrunc{From: from, To: to}
}

// ~~~ [ fpext ] ~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~

// FPExt is an LLVM IR fpext expression.
type FPExt struct {
	// Value before conversion.
	From Constant
	// Type after conversion.
	To types.Type
}

// NewFPExt returns a new fpext expression based on the given source value and
// target type.
func NewFPExt(from Constant, to types.Type) *FPExt {
	return &FPExt{From: from, To: to}
}

// ~~~ [ fptoui ] ~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~

// FPToUI is an LLVM IR fptoui expression.
type FPToUI struct {
	// Value before conversion.
	From Constant
	// Type after conversion.
	To types.Type
}

// NewFPToUI returns a new fptoui expression based on the given source value and
// target type.
func NewFPToUI(from Constant, to types.Type) *FPToUI {
	return &FPToUI{From: from, To: to}
}

// ~~~ [ fptosi ] ~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~

// FPToSI is an LLVM IR fptosi expression.
type FPToSI struct {
	// Value before conversion.
	From Constant
	// Type after conversion.
	To types.Type
}

// NewFPToSI returns a new fptosi expression based on the given source value and
// target type.
func NewFPToSI(from Constant, to types.Type) *FPToSI {
	return &FPToSI{From: from, To: to}
}

// ~~~ [ uitofp ] ~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~

// UIToFP is an LLVM IR uitofp expression.
type UIToFP struct {
	// Value before conversion.
	From Constant
	// Type after conversion.
	To types.Type
}

// NewUIToFP returns a new uitofp expression based on the given source value and
// target type.
func NewUIToFP(from Constant, to types.Type) *UIToFP {
	return &UIToFP{From: from, To: to}
}

// ~~~ [ sitofp ] ~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~

// SIToFP is an LLVM IR sitofp expression.
type SIToFP struct {
	// Value before conversion.
	From Constant
	// Type after conversion.
	To types.Type
}

// NewSIToFP returns a new sitofp expression based on the given source value and
// target type.
func NewSIToFP(from Constant, to types.Type) *SIToFP {
	return &SIToFP{From: from, To: to}
}

// ~~~ [ ptrtoint ] ~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~

// PtrToInt is an LLVM IR ptrtoint expression.
type PtrToInt struct {
	// Value before conversion.
	From Constant
	// Type after conversion.
	To types.Type
}

// NewPtrToInt returns a new ptrtoint expression based on the given source value
// and target type.
func NewPtrToInt(from Constant, to types.Type) *PtrToInt {
	return &PtrToInt{From: from, To: to}
}

// ~~~ [ inttoptr ] ~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~

// IntToPtr is an LLVM IR inttoptr expression.
type IntToPtr struct {
	// Value before conversion.
	From Constant
	// Type after conversion.
	To types.Type
}

// NewIntToPtr returns a new inttoptr expression based on the given source value
// and target type.
func NewIntToPtr(from Constant, to types.Type) *IntToPtr {
	return &IntToPtr{From: from, To: to}
}

// ~~~ [ bitcast ] ~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~

// BitCast is an LLVM IR bitcast expression.
type BitCast struct {
	// Value before conversion.
	From Constant
	// Type after conversion.
	To types.Type
}

// NewBitCast returns a new bitcast expression based on the given source value
// and target type.
func NewBitCast(from Constant, to types.Type) *BitCast {
	return &BitCast{From: from, To: to}
}

// ~~~ [ addrspacecast ] ~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~

// AddrSpaceCast is an LLVM IR addrspacecast expression.
type AddrSpaceCast struct {
	// Value before conversion.
	From Constant
	// Type after conversion.
	To types.Type
}

// NewAddrSpaceCast returns a new addrspacecast expression based on the given
// source value and target type.
func NewAddrSpaceCast(from Constant, to types.Type) *AddrSpaceCast {
	return &AddrSpaceCast{From: from, To: to}
}
