// === [ Local identifiers ] ===================================================

package dummy

import (
	"github.com/llir/llvm/internal/enc"
	"github.com/llir/llvm/ir/types"
)

// Local represents a dummy local identifier.
type Local struct {
	// Local name.
	name string
	// Type associated with the local.
	typ types.Type
}

// NewLocal returns a new dummy local identifier based on the given local name
// and type.
func NewLocal(name string, typ types.Type) *Local {
	return &Local{name: name, typ: typ}
}

// Type returns the type associated with the local.
func (local *Local) Type() types.Type {
	return local.typ
}

// Ident returns the identifier associated with the local.
func (local *Local) Ident() string {
	return enc.Local(local.name)
}

// Name returns the name of the local.
func (local *Local) Name() string {
	return local.name
}

// SetName sets the name of the local.
func (local *Local) SetName(name string) {
	local.name = name
}
