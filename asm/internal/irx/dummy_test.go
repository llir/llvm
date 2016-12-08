package irx

import "github.com/llir/llvm/ir"

// Valutate that the relevant types satisfy the ir.Terminator interface.
var (
	// Terminators
	_ ir.Terminator = &termSwitchDummy{}
)
