package ir

import "github.com/llir/l/ir/types"

// Function is an LLVM IR function.
type Function struct {
}

// Type returns the type of the value.
func (f *Function) Type() types.Type {
	panic("not yet implemented")
}

// Ident returns the identifier associated with the value.
func (f *Function) Ident() string {
	panic("not yet implemented")
}
