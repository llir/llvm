package ir

import "github.com/llir/l/ir/value"

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
}

// NewExtractValue returns a new extractvalue instruction based on the given
// aggregate value and indicies.
func NewExtractValue(x value.Value, indices ...int64) *InstExtractValue {
	return &InstExtractValue{X: x, Indices: indices}
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
}

// NewInsertValue returns a new insertvalue instruction based on the given
// aggregate value, element and indicies.
func NewInsertValue(x, elem value.Value, indices ...int64) *InstInsertValue {
	return &InstInsertValue{X: x, Elem: elem, Indices: indices}
}
