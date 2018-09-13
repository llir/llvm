// Package value provides a definition of LLVM IR values.
package value

import "github.com/llir/l/ir/types"

// Value is an LLVM IR value.
//
// A Value has one of the following underlying types.
//
//    ir.Constant   // https://godoc.org/github.com/llir/l/ir#Constant
//    value.Named   // https://godoc.org/github.com/llir/l/ir/value#Named
//    TODO: add literal metadata value?
type Value interface {
	// Type returns the type of the value.
	Type() types.Type
	// Ident returns the identifier associated with the value.
	Ident() string
}

// Named is a named LLVM IR value.
//
// A Named value has one of the following underlying types.
//
//    *ir.Global            // https://godoc.org/github.com/llir/l/ir#Global
//    *ir.Function          // https://godoc.org/github.com/llir/l/ir#Function
//    *ir.Param             // https://godoc.org/github.com/llir/l/ir#Param
//    *ir.BasicBlock        // https://godoc.org/github.com/llir/l/ir#BasicBlock
//    TODO: add named metadata value?
//    ir.Instruction        // https://godoc.org/github.com/llir/l/ir#Instruction (except store and fence)
//    *ir.TermInvoke        // https://godoc.org/github.com/llir/l/ir#TermInvoke
//    *ir.TermCatchSwitch   // https://godoc.org/github.com/llir/l/ir#TermCatchSwitch (token result used by catchpad)
type Named interface {
	Value
	// Name returns the name of the value.
	Name() string
	// SetName sets the name of the value.
	SetName(name string)
}
