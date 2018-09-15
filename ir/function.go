package ir

import (
	"fmt"

	"github.com/llir/l/internal/enc"
	"github.com/llir/l/ir/types"
)

// === [ Functions ] ===========================================================

// Function is an LLVM IR function.
type Function struct {
	// Return type.
	RetType types.Type
	// Function name.
	FuncName string
	// Function parameters.
	Params []*Param
	// Variable number of arguments.
	Variadic bool
	// Basic blocks.
	Blocks []*BasicBlock
}

// NewFunction returns a new function based on the given function name, return
// type and function parameters.
func NewFunction(name string, retType types.Type, params ...*Param) *Function {
	return &Function{FuncName: name, RetType: retType, Params: params}
}

// String returns the LLVM syntax representation of the function as a type-value
// pair.
func (f *Function) String() string {
	return fmt.Sprintf("%v %v", f.Type(), f.Ident())
}

// Type returns the type of the function.
func (f *Function) Type() types.Type {
	// TODO: cache type?
	sig := types.NewFunc(f.RetType)
	for _, param := range f.Params {
		sig.Params = append(sig.Params, param.Typ)
	}
	sig.Variadic = f.Variadic
	return types.NewPointer(sig)
}

// Ident returns the identifier associated with the function.
func (f *Function) Ident() string {
	return enc.Global(f.FuncName)
}

// Name returns the name of the function.
func (f *Function) Name() string {
	return f.FuncName
}

// SetName sets the name of the function.
func (f *Function) SetName(name string) {
	f.FuncName = name
}

// Def returns the LLVM syntax representation of the function definition or
// declaration.
func (f *Function) Def() string {
	panic("not yet implemented")
}
