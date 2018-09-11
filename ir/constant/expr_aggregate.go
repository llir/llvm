package constant

// --- [ Aggregate expressions ] -----------------------------------------------

// ~~~ [ extractvalue ] ~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~

// ExtractValue is an LLVM IR extractvalue expression.
type ExtractValue struct {
	// Aggregate value.
	X Constant
	// Element indices.
	Indices []int64
}

// NewExtractValue returns a new extractvalue expression based on the given
// aggregate value and indicies.
func NewExtractValue(x Constant, indices ...int64) *ExtractValue {
	return &ExtractValue{X: x, Indices: indices}
}

// ~~~ [ insertvalue ] ~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~

// InsertValue is an LLVM IR insertvalue expression.
type InsertValue struct {
	// Aggregate value.
	X Constant
	// Element to insert.
	Elem Constant
	// Element indices.
	Indices []int64
}

// NewInsertValue returns a new insertvalue expression based on the given
// aggregate value, element and indicies.
func NewInsertValue(x, elem Constant, indices ...int64) *InsertValue {
	return &InsertValue{X: x, Elem: elem, Indices: indices}
}
