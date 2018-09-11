package constant

// --- [ Binary expressions ] --------------------------------------------------

// ~~~ [ add ] ~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~

// Add is an LLVM IR add expression.
type Add struct {
	// Operands.
	X, Y Constant // integer scalar or vector constants
}

// NewAdd returns a new add expression based on the given operands.
func NewAdd(x, y Constant) *Add {
	return &Add{X: x, Y: y}
}

// ~~~ [ fadd ] ~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~

// FAdd is an LLVM IR fadd expression.
type FAdd struct {
	// Operands.
	X, Y Constant // floating-point scalar or vector constants
}

// NewFAdd returns a new fadd expression based on the given operands.
func NewFAdd(x, y Constant) *FAdd {
	return &FAdd{X: x, Y: y}
}

// ~~~ [ sub ] ~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~

// Sub is an LLVM IR sub expression.
type Sub struct {
	// Operands.
	X, Y Constant // integer scalar or vector constants
}

// NewSub returns a new sub expression based on the given operands.
func NewSub(x, y Constant) *Sub {
	return &Sub{X: x, Y: y}
}

// ~~~ [ fsub ] ~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~

// FSub is an LLVM IR fsub expression.
type FSub struct {
	// Operands.
	X, Y Constant // floating-point scalar or vector constants
}

// NewFSub returns a new fsub expression based on the given operands.
func NewFSub(x, y Constant) *FSub {
	return &FSub{X: x, Y: y}
}

// ~~~ [ mul ] ~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~

// Mul is an LLVM IR mul expression.
type Mul struct {
	// Operands.
	X, Y Constant // integer scalar or vector constants
}

// NewMul returns a new mul expression based on the given operands.
func NewMul(x, y Constant) *Mul {
	return &Mul{X: x, Y: y}
}

// ~~~ [ fmul ] ~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~

// FMul is an LLVM IR fmul expression.
type FMul struct {
	// Operands.
	X, Y Constant // floating-point scalar or vector constants
}

// NewFMul returns a new fmul expression based on the given operands.
func NewFMul(x, y Constant) *FMul {
	return &FMul{X: x, Y: y}
}

// ~~~ [ udiv ] ~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~

// UDiv is an LLVM IR udiv expression.
type UDiv struct {
	// Operands.
	X, Y Constant // integer scalar or vector constants
}

// NewUDiv returns a new udiv expression based on the given operands.
func NewUDiv(x, y Constant) *UDiv {
	return &UDiv{X: x, Y: y}
}

// ~~~ [ sdiv ] ~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~

// SDiv is an LLVM IR sdiv expression.
type SDiv struct {
	// Operands.
	X, Y Constant // integer scalar or vector constants
}

// NewSDiv returns a new sdiv expression based on the given operands.
func NewSDiv(x, y Constant) *SDiv {
	return &SDiv{X: x, Y: y}
}

// ~~~ [ fdiv ] ~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~

// FDiv is an LLVM IR fdiv expression.
type FDiv struct {
	// Operands.
	X, Y Constant // floating-point scalar or vector constants
}

// NewFDiv returns a new fdiv expression based on the given operands.
func NewFDiv(x, y Constant) *FDiv {
	return &FDiv{X: x, Y: y}
}

// ~~~ [ urem ] ~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~

// URem is an LLVM IR urem expression.
type URem struct {
	// Operands.
	X, Y Constant // integer scalar or vector constants
}

// NewURem returns a new urem expression based on the given operands.
func NewURem(x, y Constant) *URem {
	return &URem{X: x, Y: y}
}

// ~~~ [ srem ] ~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~

// SRem is an LLVM IR srem expression.
type SRem struct {
	// Operands.
	X, Y Constant // integer scalar or vector constants
}

// NewSRem returns a new srem expression based on the given operands.
func NewSRem(x, y Constant) *SRem {
	return &SRem{X: x, Y: y}
}

// ~~~ [ frem ] ~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~

// FRem is an LLVM IR frem expression.
type FRem struct {
	// Operands.
	X, Y Constant // floating-point scalar or vector constants
}

// NewFRem returns a new frem expression based on the given operands.
func NewFRem(x, y Constant) *FRem {
	return &FRem{X: x, Y: y}
}
