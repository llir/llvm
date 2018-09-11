package constant

// --- [ Memory expressions ] --------------------------------------------------

// ~~~ [ getelementptr ] ~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~

// GetElementPtr is an LLVM IR getelementptr expression.
type GetElementPtr struct {
	// Source address.
	Src Constant
	// Element indicies.
	Indices []Constant
}

// NewGetElementPtr returns a new getelementptr expression based on the given
// source address and element indices.
func NewGetElementPtr(src Constant, indices ...Constant) *GetElementPtr {
	return &GetElementPtr{Src: src, Indices: indices}
}
