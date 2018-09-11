package constant

// --- [ Bitwise expressions ] -------------------------------------------------

// ~~~ [ shl ] ~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~

// Shl is an LLVM IR shl expression.
type Shl struct {
	// Operands.
	X, Y Constant // integer scalars or vectors
}

// NewShl returns a new shl expression based on the given operands.
func NewShl(x, y Constant) *Shl {
	return &Shl{X: x, Y: y}
}

// ~~~ [ lshr ] ~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~

// LShr is an LLVM IR lshr expression.
type LShr struct {
	// Operands.
	X, Y Constant // integer scalars or vectors
}

// NewLShr returns a new lshr expression based on the given operands.
func NewLShr(x, y Constant) *LShr {
	return &LShr{X: x, Y: y}
}

// ~~~ [ ashr ] ~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~

// AShr is an LLVM IR ashr expression.
type AShr struct {
	// Operands.
	X, Y Constant // integer scalars or vectors
}

// NewAShr returns a new ashr expression based on the given operands.
func NewAShr(x, y Constant) *AShr {
	return &AShr{X: x, Y: y}
}

// ~~~ [ and ] ~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~

// And is an LLVM IR and expression.
type And struct {
	// Operands.
	X, Y Constant // integer scalars or vectors
}

// NewAnd returns a new and expression based on the given operands.
func NewAnd(x, y Constant) *And {
	return &And{X: x, Y: y}
}

// ~~~ [ or ] ~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~

// Or is an LLVM IR or expression.
type Or struct {
	// Operands.
	X, Y Constant // integer scalars or vectors
}

// NewOr returns a new or expression based on the given operands.
func NewOr(x, y Constant) *Or {
	return &Or{X: x, Y: y}
}

// ~~~ [ xor ] ~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~

// Xor is an LLVM IR xor expression.
type Xor struct {
	// Operands.
	X, Y Constant // integer scalars or vectors
}

// NewXor returns a new xor expression based on the given operands.
func NewXor(x, y Constant) *Xor {
	return &Xor{X: x, Y: y}
}
