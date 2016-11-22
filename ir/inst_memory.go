package ir

import (
	"fmt"

	"github.com/llir/llvm/ir/types"
	"github.com/llir/llvm/ir/value"
)

// --- [ alloca ] --------------------------------------------------------------

// --- [ load ] ----------------------------------------------------------------

// InstLoad represents a load instruction.
//
// References:
//    http://llvm.org/docs/LangRef.html#load-instruction
type InstLoad struct {
	// Parent basic block.
	parent *BasicBlock
	// Identifier associated with the instruction.
	id string
	// Source address.
	src value.Value
	// Type of the instruction.
	typ types.Type
}

// NewLoad returns a new load instruction based on the given source address.
func NewLoad(src value.Value) *InstLoad {
	if typ, ok := src.Type().(*types.PointerType); ok {
		return &InstLoad{src: src, typ: typ.Elem()}
	}
	panic(fmt.Sprintf("invalid source address type; expected *types.PointerType, got %T", src.Type()))
}

// Type returns the type of the instruction.
func (i *InstLoad) Type() types.Type {
	return i.typ
}

// Ident returns the identifier associated with the instruction.
func (i *InstLoad) Ident() string {
	return local(i.id)
}

// SetIdent sets the identifier associated with the instruction.
func (i *InstLoad) SetIdent(id string) {
	i.id = id
}

// LLVMString returns the LLVM syntax representation of the instruction.
func (i *InstLoad) LLVMString() string {
	src := i.Src()
	return fmt.Sprintf("%s = load %s, %s %s",
		i.Ident(),
		i.Type().LLVMString(),
		src.Type().LLVMString(),
		src.Ident())
}

// Parent returns the parent basic block of the instruction.
func (i *InstLoad) Parent() *BasicBlock {
	return i.parent
}

// SetParent sets the parent basic block of the instruction.
func (i *InstLoad) SetParent(parent *BasicBlock) {
	i.parent = parent
}

// Src returns the source address of the load instruction.
func (i *InstLoad) Src() value.Value {
	return i.src
}

// --- [ store ] ---------------------------------------------------------------

// --- [ fence ] ---------------------------------------------------------------

// --- [ cmpxchg ] -------------------------------------------------------------

// --- [ atomicrmw ] -----------------------------------------------------------

// --- [ getelementptr ] -------------------------------------------------------
