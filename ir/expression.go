package ir

// An Expression is a constant expression.
type Expression interface {
	Constant
	// Simplify returns an equivalent (and potentially simplified) constant of
	// the constant expression.
	Simplify() Constant
}
