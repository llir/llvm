package ir

import (
	"github.com/llir/llvm/internal/enc"
	"github.com/llir/llvm/ir/types"
)

// A Param represents an LLVM IR function parameter.
//
// Function parameters may be referenced from instructions (e.g. add), and are
// thus considered LLVM IR values.
type Param struct {
	// Underlying type.
	*types.Param
	// Track uses of the value.
	used
}

// NewParam returns a new function parameter based on the given parameter name
// and type.
func NewParam(name string, typ types.Type) *Param {
	return &Param{Param: types.NewParam(name, typ)}
}

// Type returns the type of the function parameter.
func (param *Param) Type() types.Type {
	return param.Typ
}

// Ident returns the identifier associated with the function parameter.
func (param *Param) Ident() string {
	return enc.Local(param.Name)
}

// GetName returns the name of the function parameter.
func (param *Param) GetName() string {
	return param.Name
}

// SetName sets the name of the function parameter.
func (param *Param) SetName(name string) {
	param.Name = name
}
