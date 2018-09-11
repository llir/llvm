package ir

import "github.com/llir/l/ir/value"

// --- [ Binary instructions ] -------------------------------------------------

// ~~~ [ add ] ~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~

// Add is an LLVM IR add instruction.
type Add struct {
	// Name of local variable associated with the result.
	LocalName string
	// Operands.
	X, Y value.Value
}

// NewAdd returns a new add instruction based on the given operands.
func NewAdd(x, y value.Value) *Add {
	return &Add{X: x, Y: y}
}

// ~~~ [ fadd ] ~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~

// FAdd is an LLVM IR fadd instruction.
type FAdd struct {
	// Name of local variable associated with the result.
	LocalName string
	// Operands.
	X, Y value.Value
}

// NewFAdd returns a new fadd instruction based on the given operands.
func NewFAdd(x, y value.Value) *FAdd {
	return &FAdd{X: x, Y: y}
}

// ~~~ [ sub ] ~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~

// Sub is an LLVM IR sub instruction.
type Sub struct {
	// Name of local variable associated with the result.
	LocalName string
	// Operands.
	X, Y value.Value
}

// NewSub returns a new sub instruction based on the given operands.
func NewSub(x, y value.Value) *Sub {
	return &Sub{X: x, Y: y}
}

// ~~~ [ fsub ] ~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~

// FSub is an LLVM IR fsub instruction.
type FSub struct {
	// Name of local variable associated with the result.
	LocalName string
	// Operands.
	X, Y value.Value
}

// NewFSub returns a new fsub instruction based on the given operands.
func NewFSub(x, y value.Value) *FSub {
	return &FSub{X: x, Y: y}
}

// ~~~ [ mul ] ~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~

// Mul is an LLVM IR mul instruction.
type Mul struct {
	// Name of local variable associated with the result.
	LocalName string
	// Operands.
	X, Y value.Value
}

// NewMul returns a new mul instruction based on the given operands.
func NewMul(x, y value.Value) *Mul {
	return &Mul{X: x, Y: y}
}

// ~~~ [ fmul ] ~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~

// FMul is an LLVM IR fmul instruction.
type FMul struct {
	// Name of local variable associated with the result.
	LocalName string
	// Operands.
	X, Y value.Value
}

// NewFMul returns a new fmul instruction based on the given operands.
func NewFMul(x, y value.Value) *FMul {
	return &FMul{X: x, Y: y}
}

// ~~~ [ udiv ] ~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~

// UDiv is an LLVM IR udiv instruction.
type UDiv struct {
	// Name of local variable associated with the result.
	LocalName string
	// Operands.
	X, Y value.Value
}

// NewUDiv returns a new udiv instruction based on the given operands.
func NewUDiv(x, y value.Value) *UDiv {
	return &UDiv{X: x, Y: y}
}

// ~~~ [ sdiv ] ~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~

// SDiv is an LLVM IR sdiv instruction.
type SDiv struct {
	// Name of local variable associated with the result.
	LocalName string
	// Operands.
	X, Y value.Value
}

// NewSDiv returns a new sdiv instruction based on the given operands.
func NewSDiv(x, y value.Value) *SDiv {
	return &SDiv{X: x, Y: y}
}

// ~~~ [ fdiv ] ~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~

// FDiv is an LLVM IR fdiv instruction.
type FDiv struct {
	// Name of local variable associated with the result.
	LocalName string
	// Operands.
	X, Y value.Value
}

// NewFDiv returns a new fdiv instruction based on the given operands.
func NewFDiv(x, y value.Value) *FDiv {
	return &FDiv{X: x, Y: y}
}

// ~~~ [ urem ] ~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~

// URem is an LLVM IR urem instruction.
type URem struct {
	// Name of local variable associated with the result.
	LocalName string
	// Operands.
	X, Y value.Value
}

// NewURem returns a new urem instruction based on the given operands.
func NewURem(x, y value.Value) *URem {
	return &URem{X: x, Y: y}
}

// ~~~ [ srem ] ~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~

// SRem is an LLVM IR srem instruction.
type SRem struct {
	// Name of local variable associated with the result.
	LocalName string
	// Operands.
	X, Y value.Value
}

// NewSRem returns a new srem instruction based on the given operands.
func NewSRem(x, y value.Value) *SRem {
	return &SRem{X: x, Y: y}
}

// ~~~ [ frem ] ~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~

// FRem is an LLVM IR frem instruction.
type FRem struct {
	// Name of local variable associated with the result.
	LocalName string
	// Operands.
	X, Y value.Value
}

// NewFRem returns a new frem instruction based on the given operands.
func NewFRem(x, y value.Value) *FRem {
	return &FRem{X: x, Y: y}
}
