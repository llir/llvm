package constant

import "github.com/llir/l/ir/types"

// None is a none token constant.
type None struct {
}

// Type returns the type of the constant.
func (*None) Type() types.Type {
	return types.Token
}

// Ident returns the identifier associated with the constant.
func (*None) Ident() string {
	return "none"
}
