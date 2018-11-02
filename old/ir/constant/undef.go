// === [ Undefined value constants ] ===========================================
//
// References:
//    http://llvm.org/docs/LangRef.html#undefined-values

package constant

import (
	"github.com/llir/llvm/ir/types"
)

// --- [ undef ] -----------------------------------------------------

// Undef represents a undef constant.
type Undef struct {
	// Constant type.
	Typ types.Type
}

// NewUndef returns a new undefined value constant based on the given type.
func NewUndef(typ types.Type) *Undef {
	return &Undef{Typ: typ}
}

// Type returns the type of the constant.
func (c *Undef) Type() types.Type {
	return c.Typ
}

// Ident returns the string representation of the constant.
func (c *Undef) Ident() string {
	return "undef"
}

// Immutable ensures that only constants can be assigned to the
// constant.Constant interface.
func (*Undef) Immutable() {}

// MetadataNode ensures that only metadata nodes can be assigned to the
// ir.MetadataNode interface.
func (*Undef) MetadataNode() {}
