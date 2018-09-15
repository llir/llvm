package ir

import (
	"fmt"

	"github.com/llir/l/ir/types"
	"github.com/llir/l/ir/value"
)

// --- [ Bitwise instructions ] ------------------------------------------------

// ~~~ [ shl ] ~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~

// InstShl is an LLVM IR shl instruction.
type InstShl struct {
	// Name of local variable associated with the result.
	LocalName string
	// Operands.
	X, Y value.Value // integer scalars or vectors
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
	panic("not yet implemented")
}

// Ident returns the identifier associated with the instruction.
func (inst *InstShl) Ident() string {
	panic("not yet implemented")
}

// Name returns the name of the instruction.
func (inst *InstShl) Name() string {
	return inst.LocalName
}

// SetName sets the name of the instruction.
func (inst *InstShl) SetName(name string) {
	inst.LocalName = name
}

// ~~~ [ lshr ] ~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~

// InstLShr is an LLVM IR lshr instruction.
type InstLShr struct {
	// Name of local variable associated with the result.
	LocalName string
	// Operands.
	X, Y value.Value // integer scalars or vectors
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
	panic("not yet implemented")
}

// Ident returns the identifier associated with the instruction.
func (inst *InstLShr) Ident() string {
	panic("not yet implemented")
}

// Name returns the name of the instruction.
func (inst *InstLShr) Name() string {
	return inst.LocalName
}

// SetName sets the name of the instruction.
func (inst *InstLShr) SetName(name string) {
	inst.LocalName = name
}

// ~~~ [ ashr ] ~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~

// InstAShr is an LLVM IR ashr instruction.
type InstAShr struct {
	// Name of local variable associated with the result.
	LocalName string
	// Operands.
	X, Y value.Value // integer scalars or vectors
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
	panic("not yet implemented")
}

// Ident returns the identifier associated with the instruction.
func (inst *InstAShr) Ident() string {
	panic("not yet implemented")
}

// Name returns the name of the instruction.
func (inst *InstAShr) Name() string {
	return inst.LocalName
}

// SetName sets the name of the instruction.
func (inst *InstAShr) SetName(name string) {
	inst.LocalName = name
}

// ~~~ [ and ] ~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~

// InstAnd is an LLVM IR and instruction.
type InstAnd struct {
	// Name of local variable associated with the result.
	LocalName string
	// Operands.
	X, Y value.Value // integer scalars or vectors
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
	panic("not yet implemented")
}

// Ident returns the identifier associated with the instruction.
func (inst *InstAnd) Ident() string {
	panic("not yet implemented")
}

// Name returns the name of the instruction.
func (inst *InstAnd) Name() string {
	return inst.LocalName
}

// SetName sets the name of the instruction.
func (inst *InstAnd) SetName(name string) {
	inst.LocalName = name
}

// ~~~ [ or ] ~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~

// InstOr is an LLVM IR or instruction.
type InstOr struct {
	// Name of local variable associated with the result.
	LocalName string
	// Operands.
	X, Y value.Value // integer scalars or vectors
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
	panic("not yet implemented")
}

// Ident returns the identifier associated with the instruction.
func (inst *InstOr) Ident() string {
	panic("not yet implemented")
}

// Name returns the name of the instruction.
func (inst *InstOr) Name() string {
	return inst.LocalName
}

// SetName sets the name of the instruction.
func (inst *InstOr) SetName(name string) {
	inst.LocalName = name
}

// ~~~ [ xor ] ~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~

// InstXor is an LLVM IR xor instruction.
type InstXor struct {
	// Name of local variable associated with the result.
	LocalName string
	// Operands.
	X, Y value.Value // integer scalars or vectors
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
	panic("not yet implemented")
}

// Ident returns the identifier associated with the instruction.
func (inst *InstXor) Ident() string {
	panic("not yet implemented")
}

// Name returns the name of the instruction.
func (inst *InstXor) Name() string {
	return inst.LocalName
}

// SetName sets the name of the instruction.
func (inst *InstXor) SetName(name string) {
	inst.LocalName = name
}
