package ir

import "github.com/llir/l/ir/value"

// --- [ Bitwise instructions ] ------------------------------------------------

// ~~~ [ shl ] ~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~

// Shl is an LLVM IR shl instruction.
type Shl struct {
	// Name of local variable associated with the result.
	LocalName string
	// Operands.
	X, Y value.Value // integer scalars or vectors
}

// NewShl returns a new shl instruction based on the given operands.
func NewShl(x, y value.Value) *Shl {
	return &Shl{X: x, Y: y}
}

// ~~~ [ lshr ] ~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~

// LShr is an LLVM IR lshr instruction.
type LShr struct {
	// Name of local variable associated with the result.
	LocalName string
	// Operands.
	X, Y value.Value // integer scalars or vectors
}

// NewLShr returns a new lshr instruction based on the given operands.
func NewLShr(x, y value.Value) *LShr {
	return &LShr{X: x, Y: y}
}

// ~~~ [ ashr ] ~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~

// AShr is an LLVM IR ashr instruction.
type AShr struct {
	// Name of local variable associated with the result.
	LocalName string
	// Operands.
	X, Y value.Value // integer scalars or vectors
}

// NewAShr returns a new ashr instruction based on the given operands.
func NewAShr(x, y value.Value) *AShr {
	return &AShr{X: x, Y: y}
}

// ~~~ [ and ] ~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~

// And is an LLVM IR and instruction.
type And struct {
	// Name of local variable associated with the result.
	LocalName string
	// Operands.
	X, Y value.Value // integer scalars or vectors
}

// NewAnd returns a new and instruction based on the given operands.
func NewAnd(x, y value.Value) *And {
	return &And{X: x, Y: y}
}

// ~~~ [ or ] ~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~

// Or is an LLVM IR or instruction.
type Or struct {
	// Name of local variable associated with the result.
	LocalName string
	// Operands.
	X, Y value.Value // integer scalars or vectors
}

// NewOr returns a new or instruction based on the given operands.
func NewOr(x, y value.Value) *Or {
	return &Or{X: x, Y: y}
}

// ~~~ [ xor ] ~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~

// Xor is an LLVM IR xor instruction.
type Xor struct {
	// Name of local variable associated with the result.
	LocalName string
	// Operands.
	X, Y value.Value // integer scalars or vectors
}

// NewXor returns a new xor instruction based on the given operands.
func NewXor(x, y value.Value) *Xor {
	return &Xor{X: x, Y: y}
}
