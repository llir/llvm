// === [ Global identifiers ] ==================================================

package dummy

import (
	"github.com/llir/llvm/internal/enc"
	"github.com/llir/llvm/ir/types"
)

// Global represents a dummy global identifier.
type Global struct {
	// Global name.
	name string
	// Type associated with the global.
	typ types.Type
}

// NewGlobal returns a new dummy global identifier based on the given global name
// and type.
func NewGlobal(name string, typ types.Type) *Global {
	return &Global{name: name, typ: typ}
}

// Type returns the type associated with the global.
func (global *Global) Type() types.Type {
	return global.typ
}

// Ident returns the identifier associated with the global.
func (global *Global) Ident() string {
	return enc.Global(global.name)
}

// Name returns the name of the global.
func (global *Global) Name() string {
	return global.name
}

// SetName sets the name of the global.
func (global *Global) SetName(name string) {
	global.name = name
}

// Immutable ensures that only constants can be assigned to the
// constant.Constant interface.
func (*Global) Immutable() {}
