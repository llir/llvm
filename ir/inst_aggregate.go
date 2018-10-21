package ir

import (
	"fmt"

	"github.com/llir/l/ir/types"
	"github.com/llir/l/ir/value"
)

// --- [ Aggregate instructions ] ----------------------------------------------

// ~~~ [ extractvalue ] ~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~

// InstExtractValue is an LLVM IR extractvalue instruction.
type InstExtractValue struct {
	// Name of local variable associated with the result.
	LocalName string
	// Aggregate value.
	X value.Value
	// Element indices.
	Indices []int64

	// extra.

	// Type of result produced by the instruction.
	Typ types.Type
}

// NewExtractValue returns a new extractvalue instruction based on the given
// aggregate value and indicies.
func NewExtractValue(x value.Value, indices ...int64) *InstExtractValue {
	return &InstExtractValue{X: x, Indices: indices}
}

// String returns the LLVM syntax representation of the instruction as a
// type-value pair.
func (inst *InstExtractValue) String() string {
	return fmt.Sprintf("%v %v", inst.Type(), inst.Ident())
}

// Type returns the type of the instruction.
func (inst *InstExtractValue) Type() types.Type {
	panic("not yet implemented")
}

// Ident returns the identifier associated with the instruction.
func (inst *InstExtractValue) Ident() string {
	panic("not yet implemented")
}

// Name returns the name of the instruction.
func (inst *InstExtractValue) Name() string {
	return inst.LocalName
}

// SetName sets the name of the instruction.
func (inst *InstExtractValue) SetName(name string) {
	inst.LocalName = name
}

// ~~~ [ insertvalue ] ~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~

// InstInsertValue is an LLVM IR insertvalue instruction.
type InstInsertValue struct {
	// Name of local variable associated with the result.
	LocalName string
	// Aggregate value.
	X value.Value
	// Element to insert.
	Elem value.Value
	// Element indices.
	Indices []int64

	// extra.

	// Type of result produced by the instruction.
	Typ types.Type
}

// NewInsertValue returns a new insertvalue instruction based on the given
// aggregate value, element and indicies.
func NewInsertValue(x, elem value.Value, indices ...int64) *InstInsertValue {
	return &InstInsertValue{X: x, Elem: elem, Indices: indices}
}

// String returns the LLVM syntax representation of the instruction as a
// type-value pair.
func (inst *InstInsertValue) String() string {
	return fmt.Sprintf("%v %v", inst.Type(), inst.Ident())
}

// Type returns the type of the instruction.
func (inst *InstInsertValue) Type() types.Type {
	panic("not yet implemented")
}

// Ident returns the identifier associated with the instruction.
func (inst *InstInsertValue) Ident() string {
	panic("not yet implemented")
}

// Name returns the name of the instruction.
func (inst *InstInsertValue) Name() string {
	return inst.LocalName
}

// SetName sets the name of the instruction.
func (inst *InstInsertValue) SetName(name string) {
	inst.LocalName = name
}
