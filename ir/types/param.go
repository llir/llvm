package types

import "github.com/llir/llvm/ir/internal/enc"

// A Param represents an LLVM IR function parameter.
//
// Function parameters may be referenced from instructions (e.g. add), and are
// thus considered LLVM IR values.
type Param struct {
	// Parameter name.
	name string
	// Parameter type.
	typ Type
}

// NewParam returns a new function parameter based on the given parameter name
// and type.
func NewParam(name string, typ Type) *Param {
	return &Param{name: name, typ: typ}
}

// Type returns the type of the function parameter.
func (p *Param) Type() Type {
	return p.typ
}

// Ident returns the identifier associated with the function parameter.
func (p *Param) Ident() string {
	return enc.Local(p.name)
}
