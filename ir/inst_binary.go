package ir

import (
	"fmt"

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
	X, Y value.Value // integer scalars or vectors
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
	return inst.X.Type()
}

// Ident returns the identifier associated with the instruction.
func (inst *InstAdd) Ident() string {
	panic("not yet implemented")
}

// Name returns the name of the instruction.
func (inst *InstAdd) Name() string {
	return inst.LocalName
}

// SetName sets the name of the instruction.
func (inst *InstAdd) SetName(name string) {
	inst.LocalName = name
}

// ~~~ [ fadd ] ~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~

// InstFAdd is an LLVM IR fadd instruction.
type InstFAdd struct {
	// Name of local variable associated with the result.
	LocalName string
	// Operands.
	X, Y value.Value // floating-point scalars or vectors
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
	panic("not yet implemented")
}

// Ident returns the identifier associated with the instruction.
func (inst *InstFAdd) Ident() string {
	panic("not yet implemented")
}

// Name returns the name of the instruction.
func (inst *InstFAdd) Name() string {
	return inst.LocalName
}

// SetName sets the name of the instruction.
func (inst *InstFAdd) SetName(name string) {
	inst.LocalName = name
}

// ~~~ [ sub ] ~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~

// InstSub is an LLVM IR sub instruction.
type InstSub struct {
	// Name of local variable associated with the result.
	LocalName string
	// Operands.
	X, Y value.Value // integer scalars or vectors
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
	panic("not yet implemented")
}

// Ident returns the identifier associated with the instruction.
func (inst *InstSub) Ident() string {
	panic("not yet implemented")
}

// Name returns the name of the instruction.
func (inst *InstSub) Name() string {
	return inst.LocalName
}

// SetName sets the name of the instruction.
func (inst *InstSub) SetName(name string) {
	inst.LocalName = name
}

// ~~~ [ fsub ] ~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~

// InstFSub is an LLVM IR fsub instruction.
type InstFSub struct {
	// Name of local variable associated with the result.
	LocalName string
	// Operands.
	X, Y value.Value // floating-point scalars or vectors
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
	panic("not yet implemented")
}

// Ident returns the identifier associated with the instruction.
func (inst *InstFSub) Ident() string {
	panic("not yet implemented")
}

// Name returns the name of the instruction.
func (inst *InstFSub) Name() string {
	return inst.LocalName
}

// SetName sets the name of the instruction.
func (inst *InstFSub) SetName(name string) {
	inst.LocalName = name
}

// ~~~ [ mul ] ~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~

// InstMul is an LLVM IR mul instruction.
type InstMul struct {
	// Name of local variable associated with the result.
	LocalName string
	// Operands.
	X, Y value.Value // integer scalars or vectors
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
	panic("not yet implemented")
}

// Ident returns the identifier associated with the instruction.
func (inst *InstMul) Ident() string {
	panic("not yet implemented")
}

// Name returns the name of the instruction.
func (inst *InstMul) Name() string {
	return inst.LocalName
}

// SetName sets the name of the instruction.
func (inst *InstMul) SetName(name string) {
	inst.LocalName = name
}

// ~~~ [ fmul ] ~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~

// InstFMul is an LLVM IR fmul instruction.
type InstFMul struct {
	// Name of local variable associated with the result.
	LocalName string
	// Operands.
	X, Y value.Value // floating-point scalars or vectors
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
	panic("not yet implemented")
}

// Ident returns the identifier associated with the instruction.
func (inst *InstFMul) Ident() string {
	panic("not yet implemented")
}

// Name returns the name of the instruction.
func (inst *InstFMul) Name() string {
	return inst.LocalName
}

// SetName sets the name of the instruction.
func (inst *InstFMul) SetName(name string) {
	inst.LocalName = name
}

// ~~~ [ udiv ] ~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~

// InstUDiv is an LLVM IR udiv instruction.
type InstUDiv struct {
	// Name of local variable associated with the result.
	LocalName string
	// Operands.
	X, Y value.Value // integer scalars or vectors
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
	panic("not yet implemented")
}

// Ident returns the identifier associated with the instruction.
func (inst *InstUDiv) Ident() string {
	panic("not yet implemented")
}

// Name returns the name of the instruction.
func (inst *InstUDiv) Name() string {
	return inst.LocalName
}

// SetName sets the name of the instruction.
func (inst *InstUDiv) SetName(name string) {
	inst.LocalName = name
}

// ~~~ [ sdiv ] ~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~

// InstSDiv is an LLVM IR sdiv instruction.
type InstSDiv struct {
	// Name of local variable associated with the result.
	LocalName string
	// Operands.
	X, Y value.Value // integer scalars or vectors
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
	panic("not yet implemented")
}

// Ident returns the identifier associated with the instruction.
func (inst *InstSDiv) Ident() string {
	panic("not yet implemented")
}

// Name returns the name of the instruction.
func (inst *InstSDiv) Name() string {
	return inst.LocalName
}

// SetName sets the name of the instruction.
func (inst *InstSDiv) SetName(name string) {
	inst.LocalName = name
}

// ~~~ [ fdiv ] ~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~

// InstFDiv is an LLVM IR fdiv instruction.
type InstFDiv struct {
	// Name of local variable associated with the result.
	LocalName string
	// Operands.
	X, Y value.Value // floating-point scalars or vectors
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
	panic("not yet implemented")
}

// Ident returns the identifier associated with the instruction.
func (inst *InstFDiv) Ident() string {
	panic("not yet implemented")
}

// Name returns the name of the instruction.
func (inst *InstFDiv) Name() string {
	return inst.LocalName
}

// SetName sets the name of the instruction.
func (inst *InstFDiv) SetName(name string) {
	inst.LocalName = name
}

// ~~~ [ urem ] ~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~

// InstURem is an LLVM IR urem instruction.
type InstURem struct {
	// Name of local variable associated with the result.
	LocalName string
	// Operands.
	X, Y value.Value // integer scalars or vectors
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
	panic("not yet implemented")
}

// Ident returns the identifier associated with the instruction.
func (inst *InstURem) Ident() string {
	panic("not yet implemented")
}

// Name returns the name of the instruction.
func (inst *InstURem) Name() string {
	return inst.LocalName
}

// SetName sets the name of the instruction.
func (inst *InstURem) SetName(name string) {
	inst.LocalName = name
}

// ~~~ [ srem ] ~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~

// InstSRem is an LLVM IR srem instruction.
type InstSRem struct {
	// Name of local variable associated with the result.
	LocalName string
	// Operands.
	X, Y value.Value // integer scalars or vectors
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
	panic("not yet implemented")
}

// Ident returns the identifier associated with the instruction.
func (inst *InstSRem) Ident() string {
	panic("not yet implemented")
}

// Name returns the name of the instruction.
func (inst *InstSRem) Name() string {
	return inst.LocalName
}

// SetName sets the name of the instruction.
func (inst *InstSRem) SetName(name string) {
	inst.LocalName = name
}

// ~~~ [ frem ] ~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~

// InstFRem is an LLVM IR frem instruction.
type InstFRem struct {
	// Name of local variable associated with the result.
	LocalName string
	// Operands.
	X, Y value.Value // floating-point scalars or vectors
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
	panic("not yet implemented")
}

// Ident returns the identifier associated with the instruction.
func (inst *InstFRem) Ident() string {
	panic("not yet implemented")
}

// Name returns the name of the instruction.
func (inst *InstFRem) Name() string {
	return inst.LocalName
}

// SetName sets the name of the instruction.
func (inst *InstFRem) SetName(name string) {
	inst.LocalName = name
}
