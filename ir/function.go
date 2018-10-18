package ir

import (
	"fmt"

	"github.com/llir/l/internal/enc"
	"github.com/llir/l/ir/ll"
	"github.com/llir/l/ir/types"
)

// === [ Functions ] ===========================================================

// Function is an LLVM IR function.
type Function struct {
	// Function signature.
	Sig *types.FuncType
	// Function name.
	FuncName string
	// Function parameters.
	Params []*Param
	// Basic blocks.
	Blocks []*BasicBlock

	// extra.

	// Pointer type to function, including an optional address space. If Typ is
	// nil, the first invocation of Type stores a pointer type with Sig as
	// element.
	Typ *types.PointerType
	// (optional) Linkage.
	Linkage ll.Linkage
	// (optional) Preemption; zero value if not present.
	Preemption ll.Preemption
	// (optional) Visibility; zero value if not present.
	Visibility ll.Visibility
	// (optional) DLL storage class; zero value if not present.
	DLLStorageClass ll.DLLStorageClass
	// (optional) Calling convention; zero value if not present.
	CallingConv ll.CallingConv
	// (optional) Return attributes.
	ReturnAttrs []ll.ReturnAttribute
	// (optional) Unnamed address.
	UnnamedAddr ll.UnnamedAddr
	// (optional) Function attributes.
	FuncAttrs []ll.FuncAttribute
	// (optional) Section; nil if not present.
	Section string
	// (optional) Comdat; nil if not present.
	Comdat *ComdatDef
	// (optional) Garbage collection; empty if not present.
	GC string
	// (optional) Prefix; nil if not present.
	Prefix Constant
	// (optional) Prologue; nil if not present.
	Prologue Constant
	// (optional) Personality; nil if not present.
	Personality Constant
	// (optional) Use list order.
	// TODO: add support for UseListOrder.
	//UseListOrders []*UseListOrder
	// (optional) Metadata attachments.
	// TODO: add support for metadata.
	//Metadata []*metadata.MetadataAttachment
}

// TODO: decide whether to have the function name parameter be the first
// argument (to be consistent with NewGlobal) or after retType (to be consistent
// with order of occurence in LLVM IR syntax).

// NewFunction returns a new function based on the given function name, return
// type and function parameters.
func NewFunction(name string, retType types.Type, params ...*Param) *Function {
	panic("not yet implemented")
	/*
		paramTypes := make([]types.Type, len(params))
		for i, param := range f.Params {
			paramType[i] = param.Type()
		}
		sig := types.NewFunc(f.RetType, paramTypes...)
		return &Function{Sig: sig, FuncName: name, Params: params}
	*/
}

// String returns the LLVM syntax representation of the function as a type-value
// pair.
func (f *Function) String() string {
	return fmt.Sprintf("%v %v", f.Type(), f.Ident())
}

// Type returns the type of the function.
func (f *Function) Type() types.Type {
	// Cache type if not present.
	if f.Typ == nil {
		f.Typ = types.NewPointer(f.Sig)
	}
	return f.Typ
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
