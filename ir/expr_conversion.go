package ir

import (
	"github.com/llir/l/ir/types"
)

// --- [ Conversion expressions ] ----------------------------------------------

// ~~~ [ trunc ] ~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~

// TruncExpr is an LLVM IR trunc expression.
type TruncExpr struct {
	// Value before conversion.
	From Constant
	// Type after conversion.
	To types.Type
}

// NewTruncExpr returns a new trunc expression based on the given source value
// and target type.
func NewTruncExpr(from Constant, to types.Type) *TruncExpr {
	return &TruncExpr{From: from, To: to}
}

// Type returns the type of the constant expression.
func (e *TruncExpr) Type() types.Type {
	return e.To
}

// Ident returns the identifier associated with the constant expression.
func (e *TruncExpr) Ident() string {
	panic("not yet implemented")
}

// Simplify returns an equivalent (and potentially simplified) constant of the
// constant expression.
func (e *TruncExpr) Simplify() Constant {
	panic("not yet implemented")
}

// ~~~ [ zext ] ~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~

// ZExtExpr is an LLVM IR zext expression.
type ZExtExpr struct {
	// Value before conversion.
	From Constant
	// Type after conversion.
	To types.Type
}

// NewZExtExpr returns a new zext expression based on the given source value and
// target type.
func NewZExtExpr(from Constant, to types.Type) *ZExtExpr {
	return &ZExtExpr{From: from, To: to}
}

// Type returns the type of the constant expression.
func (e *ZExtExpr) Type() types.Type {
	return e.To
}

// Ident returns the identifier associated with the constant expression.
func (e *ZExtExpr) Ident() string {
	panic("not yet implemented")
}

// Simplify returns an equivalent (and potentially simplified) constant of the
// constant expression.
func (e *ZExtExpr) Simplify() Constant {
	panic("not yet implemented")
}

// ~~~ [ sext ] ~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~

// SExtExpr is an LLVM IR sext expression.
type SExtExpr struct {
	// Value before conversion.
	From Constant
	// Type after conversion.
	To types.Type
}

// NewSExtExpr returns a new sext expression based on the given source value and
// target type.
func NewSExtExpr(from Constant, to types.Type) *SExtExpr {
	return &SExtExpr{From: from, To: to}
}

// Type returns the type of the constant expression.
func (e *SExtExpr) Type() types.Type {
	return e.To
}

// Ident returns the identifier associated with the constant expression.
func (e *SExtExpr) Ident() string {
	panic("not yet implemented")
}

// Simplify returns an equivalent (and potentially simplified) constant of the
// constant expression.
func (e *SExtExpr) Simplify() Constant {
	panic("not yet implemented")
}

// ~~~ [ fptrunc ] ~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~

// FPTruncExpr is an LLVM IR fptrunc expression.
type FPTruncExpr struct {
	// Value before conversion.
	From Constant
	// Type after conversion.
	To types.Type
}

// NewFPTruncExpr returns a new fptrunc expression based on the given source
// value and target type.
func NewFPTruncExpr(from Constant, to types.Type) *FPTruncExpr {
	return &FPTruncExpr{From: from, To: to}
}

// Type returns the type of the constant expression.
func (e *FPTruncExpr) Type() types.Type {
	return e.To
}

// Ident returns the identifier associated with the constant expression.
func (e *FPTruncExpr) Ident() string {
	panic("not yet implemented")
}

// Simplify returns an equivalent (and potentially simplified) constant of the
// constant expression.
func (e *FPTruncExpr) Simplify() Constant {
	panic("not yet implemented")
}

// ~~~ [ fpext ] ~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~

// FPExtExpr is an LLVM IR fpext expression.
type FPExtExpr struct {
	// Value before conversion.
	From Constant
	// Type after conversion.
	To types.Type
}

// NewFPExtExpr returns a new fpext expression based on the given source value
// and target type.
func NewFPExtExpr(from Constant, to types.Type) *FPExtExpr {
	return &FPExtExpr{From: from, To: to}
}

// Type returns the type of the constant expression.
func (e *FPExtExpr) Type() types.Type {
	return e.To
}

// Ident returns the identifier associated with the constant expression.
func (e *FPExtExpr) Ident() string {
	panic("not yet implemented")
}

// Simplify returns an equivalent (and potentially simplified) constant of the
// constant expression.
func (e *FPExtExpr) Simplify() Constant {
	panic("not yet implemented")
}

// ~~~ [ fptoui ] ~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~

// FPToUIExpr is an LLVM IR fptoui expression.
type FPToUIExpr struct {
	// Value before conversion.
	From Constant
	// Type after conversion.
	To types.Type
}

// NewFPToUIExpr returns a new fptoui expression based on the given source value
// and target type.
func NewFPToUIExpr(from Constant, to types.Type) *FPToUIExpr {
	return &FPToUIExpr{From: from, To: to}
}

// Type returns the type of the constant expression.
func (e *FPToUIExpr) Type() types.Type {
	return e.To
}

// Ident returns the identifier associated with the constant expression.
func (e *FPToUIExpr) Ident() string {
	panic("not yet implemented")
}

// Simplify returns an equivalent (and potentially simplified) constant of the
// constant expression.
func (e *FPToUIExpr) Simplify() Constant {
	panic("not yet implemented")
}

// ~~~ [ fptosi ] ~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~

// FPToSIExpr is an LLVM IR fptosi expression.
type FPToSIExpr struct {
	// Value before conversion.
	From Constant
	// Type after conversion.
	To types.Type
}

// NewFPToSIExpr returns a new fptosi expression based on the given source value
// and target type.
func NewFPToSIExpr(from Constant, to types.Type) *FPToSIExpr {
	return &FPToSIExpr{From: from, To: to}
}

// Type returns the type of the constant expression.
func (e *FPToSIExpr) Type() types.Type {
	return e.To
}

// Ident returns the identifier associated with the constant expression.
func (e *FPToSIExpr) Ident() string {
	panic("not yet implemented")
}

// Simplify returns an equivalent (and potentially simplified) constant of the
// constant expression.
func (e *FPToSIExpr) Simplify() Constant {
	panic("not yet implemented")
}

// ~~~ [ uitofp ] ~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~

// UIToFPExpr is an LLVM IR uitofp expression.
type UIToFPExpr struct {
	// Value before conversion.
	From Constant
	// Type after conversion.
	To types.Type
}

// NewUIToFPExpr returns a new uitofp expression based on the given source value
// and target type.
func NewUIToFPExpr(from Constant, to types.Type) *UIToFPExpr {
	return &UIToFPExpr{From: from, To: to}
}

// Type returns the type of the constant expression.
func (e *UIToFPExpr) Type() types.Type {
	return e.To
}

// Ident returns the identifier associated with the constant expression.
func (e *UIToFPExpr) Ident() string {
	panic("not yet implemented")
}

// Simplify returns an equivalent (and potentially simplified) constant of the
// constant expression.
func (e *UIToFPExpr) Simplify() Constant {
	panic("not yet implemented")
}

// ~~~ [ sitofp ] ~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~

// SIToFPExpr is an LLVM IR sitofp expression.
type SIToFPExpr struct {
	// Value before conversion.
	From Constant
	// Type after conversion.
	To types.Type
}

// NewSIToFPExpr returns a new sitofp expression based on the given source value
// and target type.
func NewSIToFPExpr(from Constant, to types.Type) *SIToFPExpr {
	return &SIToFPExpr{From: from, To: to}
}

// Type returns the type of the constant expression.
func (e *SIToFPExpr) Type() types.Type {
	return e.To
}

// Ident returns the identifier associated with the constant expression.
func (e *SIToFPExpr) Ident() string {
	panic("not yet implemented")
}

// Simplify returns an equivalent (and potentially simplified) constant of the
// constant expression.
func (e *SIToFPExpr) Simplify() Constant {
	panic("not yet implemented")
}

// ~~~ [ ptrtoint ] ~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~

// PtrToIntExpr is an LLVM IR ptrtoint expression.
type PtrToIntExpr struct {
	// Value before conversion.
	From Constant
	// Type after conversion.
	To types.Type
}

// NewPtrToIntExpr returns a new ptrtoint expression based on the given source
// value and target type.
func NewPtrToIntExpr(from Constant, to types.Type) *PtrToIntExpr {
	return &PtrToIntExpr{From: from, To: to}
}

// Type returns the type of the constant expression.
func (e *PtrToIntExpr) Type() types.Type {
	return e.To
}

// Ident returns the identifier associated with the constant expression.
func (e *PtrToIntExpr) Ident() string {
	panic("not yet implemented")
}

// Simplify returns an equivalent (and potentially simplified) constant of the
// constant expression.
func (e *PtrToIntExpr) Simplify() Constant {
	panic("not yet implemented")
}

// ~~~ [ inttoptr ] ~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~

// IntToPtrExpr is an LLVM IR inttoptr expression.
type IntToPtrExpr struct {
	// Value before conversion.
	From Constant
	// Type after conversion.
	To types.Type
}

// NewIntToPtrExpr returns a new inttoptr expression based on the given source
// value and target type.
func NewIntToPtrExpr(from Constant, to types.Type) *IntToPtrExpr {
	return &IntToPtrExpr{From: from, To: to}
}

// Type returns the type of the constant expression.
func (e *IntToPtrExpr) Type() types.Type {
	return e.To
}

// Ident returns the identifier associated with the constant expression.
func (e *IntToPtrExpr) Ident() string {
	panic("not yet implemented")
}

// Simplify returns an equivalent (and potentially simplified) constant of the
// constant expression.
func (e *IntToPtrExpr) Simplify() Constant {
	panic("not yet implemented")
}

// ~~~ [ bitcast ] ~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~

// BitCastExpr is an LLVM IR bitcast expression.
type BitCastExpr struct {
	// Value before conversion.
	From Constant
	// Type after conversion.
	To types.Type
}

// NewBitCastExpr returns a new bitcast expression based on the given source
// value and target type.
func NewBitCastExpr(from Constant, to types.Type) *BitCastExpr {
	return &BitCastExpr{From: from, To: to}
}

// Type returns the type of the constant expression.
func (e *BitCastExpr) Type() types.Type {
	return e.To
}

// Ident returns the identifier associated with the constant expression.
func (e *BitCastExpr) Ident() string {
	panic("not yet implemented")
}

// Simplify returns an equivalent (and potentially simplified) constant of the
// constant expression.
func (e *BitCastExpr) Simplify() Constant {
	panic("not yet implemented")
}

// ~~~ [ addrspacecast ] ~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~

// AddrSpaceCastExpr is an LLVM IR addrspacecast expression.
type AddrSpaceCastExpr struct {
	// Value before conversion.
	From Constant
	// Type after conversion.
	To types.Type
}

// NewAddrSpaceCastExpr returns a new addrspacecast expression based on the
// given source value and target type.
func NewAddrSpaceCastExpr(from Constant, to types.Type) *AddrSpaceCastExpr {
	return &AddrSpaceCastExpr{From: from, To: to}
}

// Type returns the type of the constant expression.
func (e *AddrSpaceCastExpr) Type() types.Type {
	return e.To
}

// Ident returns the identifier associated with the constant expression.
func (e *AddrSpaceCastExpr) Ident() string {
	panic("not yet implemented")
}

// Simplify returns an equivalent (and potentially simplified) constant of the
// constant expression.
func (e *AddrSpaceCastExpr) Simplify() Constant {
	panic("not yet implemented")
}
