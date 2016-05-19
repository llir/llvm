// References:
//    http://llvm.org/docs/LangRef.html#conversion-operations

package instruction

import (
	"fmt"

	"github.com/llir/llvm/ir/types"
	"github.com/llir/llvm/ir/value"
)

// Trunc represents a truncation instruction.
type Trunc struct {
	// Value to truncate.
	from value.Value
	// Type after truncation.
	to types.Type
}

// NewTrunc returns a new trunc instruction based on the given value and target
// type.
func NewTrunc(from value.Value, to types.Type) (*Trunc, error) {
	return &Trunc{from: from, to: to}, nil
}

// Type returns the type of the value produced by the instruction.
func (inst *Trunc) Type() types.Type {
	return inst.to
}

// String returns the string representation of the instruction.
func (inst *Trunc) String() string {
	return fmt.Sprintf("trunc %s %s to %s", inst.from.Type(), inst.from, inst.to)
}

// ZExt represents a zero extension instruction.
type ZExt struct {
	// Value to zero extend.
	from value.Value
	// Type after zero extension.
	to types.Type
}

// NewZExt returns a new zext instruction based on the given value and target
// type.
func NewZExt(from value.Value, to types.Type) (*ZExt, error) {
	return &ZExt{from: from, to: to}, nil
}

// Type returns the type of the value produced by the instruction.
func (inst *ZExt) Type() types.Type {
	return inst.to
}

// String returns the string representation of the instruction.
func (inst *ZExt) String() string {
	return fmt.Sprintf("zext %s %s to %s", inst.from.Type(), inst.from, inst.to)
}

// SExt represents a sign extension instruction.
type SExt struct {
	// Value to sign extend.
	from value.Value
	// Type after sign extension.
	to types.Type
}

// NewSExt returns a new sext instruction based on the given value and target
// type.
func NewSExt(from value.Value, to types.Type) (*SExt, error) {
	return &SExt{from: from, to: to}, nil
}

// Type returns the type of the value produced by the instruction.
func (inst *SExt) Type() types.Type {
	return inst.to
}

// String returns the string representation of the instruction.
func (inst *SExt) String() string {
	return fmt.Sprintf("sext %s %s to %s", inst.from.Type(), inst.from, inst.to)
}

// FPTrunc represents a floating-point truncation instruction.
type FPTrunc struct {
	// Value to truncate.
	from value.Value
	// Type after truncation.
	to types.Type
}

// NewFPTrunc returns a new fptrunc instruction based on the given value and
// target type.
func NewFPTrunc(from value.Value, to types.Type) (*FPTrunc, error) {
	return &FPTrunc{from: from, to: to}, nil
}

// Type returns the type of the value produced by the instruction.
func (inst *FPTrunc) Type() types.Type {
	return inst.to
}

// String returns the string representation of the instruction.
func (inst *FPTrunc) String() string {
	return fmt.Sprintf("fptrunc %s %s to %s", inst.from.Type(), inst.from, inst.to)
}

// FPExt represents a floating-point extension instruction.
type FPExt struct {
	// Value to extend.
	from value.Value
	// Type after extension.
	to types.Type
}

// NewFPExt returns a new fpext instruction based on the given value and target
// type.
func NewFPExt(from value.Value, to types.Type) (*FPExt, error) {
	return &FPExt{from: from, to: to}, nil
}

// Type returns the type of the value produced by the instruction.
func (inst *FPExt) Type() types.Type {
	return inst.to
}

// String returns the string representation of the instruction.
func (inst *FPExt) String() string {
	return fmt.Sprintf("fpext %s %s to %s", inst.from.Type(), inst.from, inst.to)
}

// FPToUI represents a floating-point to unsigned integer conversion
// instruction.
type FPToUI struct {
	// Value to convert.
	from value.Value
	// Type after conversion.
	to types.Type
}

// NewFPToUI returns a new fptoui instruction based on the given value and
// target type.
func NewFPToUI(from value.Value, to types.Type) (*FPToUI, error) {
	return &FPToUI{from: from, to: to}, nil
}

// Type returns the type of the value produced by the instruction.
func (inst *FPToUI) Type() types.Type {
	return inst.to
}

// String returns the string representation of the instruction.
func (inst *FPToUI) String() string {
	return fmt.Sprintf("fptoui %s %s to %s", inst.from.Type(), inst.from, inst.to)
}

// FPToSI represents a floating-point to signed integer conversion instruction.
type FPToSI struct {
	// Value to convert.
	from value.Value
	// Type after conversion.
	to types.Type
}

// NewFPToSI returns a new fstosi instruction based on the given value and
// target type.
func NewFPToSI(from value.Value, to types.Type) (*FPToSI, error) {
	return &FPToSI{from: from, to: to}, nil
}

// Type returns the type of the value produced by the instruction.
func (inst *FPToSI) Type() types.Type {
	return inst.to
}

// String returns the string representation of the instruction.
func (inst *FPToSI) String() string {
	return fmt.Sprintf("fstosi %s %s to %s", inst.from.Type(), inst.from, inst.to)
}

// UIToFP represents a unsigned integer to floating-point conversion
// instruction.
type UIToFP struct {
	// Value to convert.
	from value.Value
	// Type after conversion.
	to types.Type
}

// NewUIToFP returns a new uitofp instruction based on the given value and
// target type.
func NewUIToFP(from value.Value, to types.Type) (*UIToFP, error) {
	return &UIToFP{from: from, to: to}, nil
}

// Type returns the type of the value produced by the instruction.
func (inst *UIToFP) Type() types.Type {
	return inst.to
}

// String returns the string representation of the instruction.
func (inst *UIToFP) String() string {
	return fmt.Sprintf("uitofp %s %s to %s", inst.from.Type(), inst.from, inst.to)
}

// SIToFP represents a signed integer to floating-point conversion instruction.
type SIToFP struct {
	// Value to convert.
	from value.Value
	// Type after conversion.
	to types.Type
}

// NewSIToFP returns a new sitofp instruction based on the given value and
// target type.
func NewSIToFP(from value.Value, to types.Type) (*SIToFP, error) {
	return &SIToFP{from: from, to: to}, nil
}

// Type returns the type of the value produced by the instruction.
func (inst *SIToFP) Type() types.Type {
	return inst.to
}

// String returns the string representation of the instruction.
func (inst *SIToFP) String() string {
	return fmt.Sprintf("sitofp %s %s to %s", inst.from.Type(), inst.from, inst.to)
}

// PtrToInt represents a pointer to integer conversion instruction.
type PtrToInt struct {
	// Value to convert.
	from value.Value
	// Type after conversion.
	to types.Type
}

// NewPtrToInt returns a new ptrtoint instruction based on the given value and
// target type.
func NewPtrToInt(from value.Value, to types.Type) (*PtrToInt, error) {
	return &PtrToInt{from: from, to: to}, nil
}

// Type returns the type of the value produced by the instruction.
func (inst *PtrToInt) Type() types.Type {
	return inst.to
}

// String returns the string representation of the instruction.
func (inst *PtrToInt) String() string {
	return fmt.Sprintf("ptrtoint %s %s to %s", inst.from.Type(), inst.from, inst.to)
}

// IntToPtr represents an integer to pointer conversion instruction.
type IntToPtr struct {
	// Value to convert.
	from value.Value
	// Type after conversion.
	to types.Type
}

// NewIntToPtr returns a new inttoptr instruction based on the given value and
// target type.
func NewIntToPtr(from value.Value, to types.Type) (*IntToPtr, error) {
	return &IntToPtr{from: from, to: to}, nil
}

// Type returns the type of the value produced by the instruction.
func (inst *IntToPtr) Type() types.Type {
	return inst.to
}

// String returns the string representation of the instruction.
func (inst *IntToPtr) String() string {
	return fmt.Sprintf("inttoptr %s %s to %s", inst.from.Type(), inst.from, inst.to)
}

// BitCast represents a bitcast instruction.
type BitCast struct {
	// Value to cast.
	from value.Value
	// Type after conversion.
	to types.Type
}

// NewBitCast returns a new bitcast instruction based on the given value and
// target type.
func NewBitCast(from value.Value, to types.Type) (*BitCast, error) {
	return &BitCast{from: from, to: to}, nil
}

// Type returns the type of the value produced by the instruction.
func (inst *BitCast) Type() types.Type {
	return inst.to
}

// String returns the string representation of the instruction.
func (inst *BitCast) String() string {
	return fmt.Sprintf("bitcast %s %s to %s", inst.from.Type(), inst.from, inst.to)
}

// AddrSpaceCast represents an address space cast instruction.
type AddrSpaceCast struct {
	// Value to cast.
	from value.Value
	// Type after conversion.
	to types.Type
}

// NewAddrSpaceCast returns a new addrspacecast instruction based on the given
// value and target type.
func NewAddrSpaceCast(from value.Value, to types.Type) (*AddrSpaceCast, error) {
	return &AddrSpaceCast{from: from, to: to}, nil
}

// Type returns the type of the value produced by the instruction.
func (inst *AddrSpaceCast) Type() types.Type {
	return inst.to
}

// String returns the string representation of the instruction.
func (inst *AddrSpaceCast) String() string {
	return fmt.Sprintf("addrspacecast %s %s to %s", inst.from.Type(), inst.from, inst.to)
}

// isValueInst ensures that only instructions which return values can be
// assigned to the Value interface.
func (*Trunc) isValueInst()         {}
func (*ZExt) isValueInst()          {}
func (*SExt) isValueInst()          {}
func (*FPTrunc) isValueInst()       {}
func (*FPExt) isValueInst()         {}
func (*FPToUI) isValueInst()        {}
func (*FPToSI) isValueInst()        {}
func (*UIToFP) isValueInst()        {}
func (*SIToFP) isValueInst()        {}
func (*PtrToInt) isValueInst()      {}
func (*IntToPtr) isValueInst()      {}
func (*BitCast) isValueInst()       {}
func (*AddrSpaceCast) isValueInst() {}
