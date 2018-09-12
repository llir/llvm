package ir

import (
	"github.com/llir/l/ir/types"
	"github.com/llir/l/ir/value"
)

// --- [ Vector instructions ] -------------------------------------------------

// ~~~ [ extractelement ] ~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~

// InstExtractElement is an LLVM IR extractelement instruction.
type InstExtractElement struct {
	// Name of local variable associated with the result.
	LocalName string
	// Vector.
	X value.Value
	// Element index.
	Index value.Value
}

// NewExtractElement returns a new extractelement instruction based on the given
// vector and element index.
func NewExtractElement(x, index value.Value) *InstExtractElement {
	return &InstExtractElement{X: x, Index: index}
}

// Type returns the type of the instruction.
func (inst *InstExtractElement) Type() types.Type {
	panic("not yet implemented")
}

// Ident returns the identifier associated with the instruction.
func (inst *InstExtractElement) Ident() string {
	panic("not yet implemented")
}

// Name returns the name of the instruction.
func (inst *InstExtractElement) Name() string {
	return inst.LocalName
}

// SetName sets the name of the instruction.
func (inst *InstExtractElement) SetName(name string) {
	inst.LocalName = name
}

// ~~~ [ insertelement ] ~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~

// InstInsertElement is an LLVM IR insertelement instruction.
type InstInsertElement struct {
	// Name of local variable associated with the result.
	LocalName string
	// Vector.
	X value.Value
	// Element to insert.
	Elem value.Value
	// Element index.
	Index value.Value
}

// NewInsertElement returns a new insertelement instruction based on the given
// vector, element and element index.
func NewInsertElement(x, elem, index value.Value) *InstInsertElement {
	return &InstInsertElement{X: x, Elem: elem, Index: index}
}

// Type returns the type of the instruction.
func (inst *InstInsertElement) Type() types.Type {
	panic("not yet implemented")
}

// Ident returns the identifier associated with the instruction.
func (inst *InstInsertElement) Ident() string {
	panic("not yet implemented")
}

// Name returns the name of the instruction.
func (inst *InstInsertElement) Name() string {
	return inst.LocalName
}

// SetName sets the name of the instruction.
func (inst *InstInsertElement) SetName(name string) {
	inst.LocalName = name
}

// ~~~ [ shufflevector ] ~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~

// InstShuffleVector is an LLVM IR shufflevector instruction.
type InstShuffleVector struct {
	// Name of local variable associated with the result.
	LocalName string
	// Vectors.
	X, Y value.Value
	// Shuffle mask.
	Mask value.Value
}

// NewShuffleVector returns a new shufflevector instruction based on the given
// vectors and shuffle mask.
func NewShuffleVector(x, y, mask value.Value) *InstShuffleVector {
	return &InstShuffleVector{X: x, Y: y, Mask: mask}
}

// Type returns the type of the instruction.
func (inst *InstShuffleVector) Type() types.Type {
	panic("not yet implemented")
}

// Ident returns the identifier associated with the instruction.
func (inst *InstShuffleVector) Ident() string {
	panic("not yet implemented")
}

// Name returns the name of the instruction.
func (inst *InstShuffleVector) Name() string {
	return inst.LocalName
}

// SetName sets the name of the instruction.
func (inst *InstShuffleVector) SetName(name string) {
	inst.LocalName = name
}
