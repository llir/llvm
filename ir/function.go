package ir

import "github.com/llir/l/ir/types"

// Function is an LLVM IR function.
type Function struct {
	// Function name.
	FuncName string
}

// Type returns the type of the function.
func (f *Function) Type() types.Type {
	panic("not yet implemented")
}

// Ident returns the identifier associated with the function.
func (f *Function) Ident() string {
	panic("not yet implemented")
}

// Name returns the name of the function.
func (f *Function) Name() string {
	return f.FuncName
}

// SetName sets the name of the function.
func (f *Function) SetName(name string) {
	f.FuncName = name
}
