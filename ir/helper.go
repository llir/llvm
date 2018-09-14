package ir

import (
	"github.com/llir/l/internal/enc"
	"github.com/llir/l/ir/types"
)

func (*BasicBlock) IsUnwindTarget() {}

// --- [ Function parameters ] -------------------------------------------------

// Param is an LLVM IR function parameter.
type Param struct {
	// Parameter type.
	Typ types.Type
	// Parameter name.
	ParamName string
}

// Type returns the type of the function parameter.
func (p *Param) Type() types.Type {
	return p.Typ
}

// Ident returns the identifier associated with the function parameter.
func (p *Param) Ident() string {
	return enc.Local(p.ParamName)
}

// Name returns the name of the function parameter.
func (p *Param) Name() string {
	return p.ParamName
}

// SetName sets the name of the function parameter.
func (p *Param) SetName(name string) {
	p.ParamName = name
}
