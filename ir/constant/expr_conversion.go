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

// Type returns the type of the constant expression.
func (e *Trunc) Type() types.Type {
	return e.To
}

// Ident returns the identifier associated with the constant expression.
func (e *Trunc) Ident() string {
	panic("not yet implemented")
}

// Simplify returns an equivalent (and potentially simplified) constant of the
// constant expression.
func (e *Trunc) Simplify() Constant {
	panic("not yet implemented")
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

// Type returns the type of the constant expression.
func (e *ZExt) Type() types.Type {
	return e.To
}

// Ident returns the identifier associated with the constant expression.
func (e *ZExt) Ident() string {
	panic("not yet implemented")
}

// Simplify returns an equivalent (and potentially simplified) constant of the
// constant expression.
func (e *ZExt) Simplify() Constant {
	panic("not yet implemented")
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

// Type returns the type of the constant expression.
func (e *SExt) Type() types.Type {
	return e.To
}

// Ident returns the identifier associated with the constant expression.
func (e *SExt) Ident() string {
	panic("not yet implemented")
}

// Simplify returns an equivalent (and potentially simplified) constant of the
// constant expression.
func (e *SExt) Simplify() Constant {
	panic("not yet implemented")
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

// Type returns the type of the constant expression.
func (e *FPTrunc) Type() types.Type {
	return e.To
}

// Ident returns the identifier associated with the constant expression.
func (e *FPTrunc) Ident() string {
	panic("not yet implemented")
}

// Simplify returns an equivalent (and potentially simplified) constant of the
// constant expression.
func (e *FPTrunc) Simplify() Constant {
	panic("not yet implemented")
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

// Type returns the type of the constant expression.
func (e *FPExt) Type() types.Type {
	return e.To
}

// Ident returns the identifier associated with the constant expression.
func (e *FPExt) Ident() string {
	panic("not yet implemented")
}

// Simplify returns an equivalent (and potentially simplified) constant of the
// constant expression.
func (e *FPExt) Simplify() Constant {
	panic("not yet implemented")
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

// Type returns the type of the constant expression.
func (e *FPToUI) Type() types.Type {
	return e.To
}

// Ident returns the identifier associated with the constant expression.
func (e *FPToUI) Ident() string {
	panic("not yet implemented")
}

// Simplify returns an equivalent (and potentially simplified) constant of the
// constant expression.
func (e *FPToUI) Simplify() Constant {
	panic("not yet implemented")
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

// Type returns the type of the constant expression.
func (e *FPToSI) Type() types.Type {
	return e.To
}

// Ident returns the identifier associated with the constant expression.
func (e *FPToSI) Ident() string {
	panic("not yet implemented")
}

// Simplify returns an equivalent (and potentially simplified) constant of the
// constant expression.
func (e *FPToSI) Simplify() Constant {
	panic("not yet implemented")
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

// Type returns the type of the constant expression.
func (e *UIToFP) Type() types.Type {
	return e.To
}

// Ident returns the identifier associated with the constant expression.
func (e *UIToFP) Ident() string {
	panic("not yet implemented")
}

// Simplify returns an equivalent (and potentially simplified) constant of the
// constant expression.
func (e *UIToFP) Simplify() Constant {
	panic("not yet implemented")
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

// Type returns the type of the constant expression.
func (e *SIToFP) Type() types.Type {
	return e.To
}

// Ident returns the identifier associated with the constant expression.
func (e *SIToFP) Ident() string {
	panic("not yet implemented")
}

// Simplify returns an equivalent (and potentially simplified) constant of the
// constant expression.
func (e *SIToFP) Simplify() Constant {
	panic("not yet implemented")
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

// Type returns the type of the constant expression.
func (e *PtrToInt) Type() types.Type {
	return e.To
}

// Ident returns the identifier associated with the constant expression.
func (e *PtrToInt) Ident() string {
	panic("not yet implemented")
}

// Simplify returns an equivalent (and potentially simplified) constant of the
// constant expression.
func (e *PtrToInt) Simplify() Constant {
	panic("not yet implemented")
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

// Type returns the type of the constant expression.
func (e *IntToPtr) Type() types.Type {
	return e.To
}

// Ident returns the identifier associated with the constant expression.
func (e *IntToPtr) Ident() string {
	panic("not yet implemented")
}

// Simplify returns an equivalent (and potentially simplified) constant of the
// constant expression.
func (e *IntToPtr) Simplify() Constant {
	panic("not yet implemented")
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

// Type returns the type of the constant expression.
func (e *BitCast) Type() types.Type {
	return e.To
}

// Ident returns the identifier associated with the constant expression.
func (e *BitCast) Ident() string {
	panic("not yet implemented")
}

// Simplify returns an equivalent (and potentially simplified) constant of the
// constant expression.
func (e *BitCast) Simplify() Constant {
	panic("not yet implemented")
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

// Type returns the type of the constant expression.
func (e *AddrSpaceCast) Type() types.Type {
	return e.To
}

// Ident returns the identifier associated with the constant expression.
func (e *AddrSpaceCast) Ident() string {
	panic("not yet implemented")
}

// Simplify returns an equivalent (and potentially simplified) constant of the
// constant expression.
func (e *AddrSpaceCast) Simplify() Constant {
	panic("not yet implemented")
}
