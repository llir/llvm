package instruction

import "github.com/llir/llvm/ir/value"

// Ret represents an addition instruction.
type Ret struct {
	// Return value; or nil if "void" return.
	x value.Value
}

// NewRet returns a new ret terminator based on the given return value. A nil
// return value indicates a "void" return.
func NewRet(x value.Value) *Ret {
	return &Ret{x: x}
}
