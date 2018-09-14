package ir

import "github.com/llir/l/ir/types"

// --- [ Token constants ] -----------------------------------------------------

// ConstNone is a none token constant.
type ConstNone struct {
}

// Type returns the type of the constant.
func (*ConstNone) Type() types.Type {
	return types.Token
}

// Ident returns the identifier associated with the constant.
func (*ConstNone) Ident() string {
	return "none"
}
