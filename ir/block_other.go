package ir

import (
	"github.com/llir/l/ir/ll"
	"github.com/llir/l/ir/types"
	"github.com/llir/l/ir/value"
)

// --- [ Other instructions ] --------------------------------------------------

// ~~~ [ icmp ] ~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~

// NewICmp returns a new icmp instruction based on the given integer comparison
// condition and integer scalar or vector operands.
func (block *BasicBlock) NewICmp(cond ll.ICond, x, y value.Value) *InstICmp {
	inst := NewICmp(cond, x, y)
	block.Insts = append(block.Insts, inst)
	return inst
}

// ~~~ [ fcmp ] ~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~

// NewFCmp returns a new fcmp instruction based on the given floating-point
// comparison condition and floating-point scalar or vector operands.
func (block *BasicBlock) NewFCmp(cond ll.FCond, x, y value.Value) *InstFCmp {
	inst := NewFCmp(cond, x, y)
	block.Insts = append(block.Insts, inst)
	return inst
}

// ~~~ [ phi ] ~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~

// NewPhi returns a new phi instruction based on the given incoming values.
func (block *BasicBlock) NewPhi(incs ...*Incoming) *InstPhi {
	inst := NewPhi(incs...)
	block.Insts = append(block.Insts, inst)
	return inst
}

// ~~~ [ select ] ~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~

// NewSelect returns a new select instruction based on the given selection
// condition and operands.
func (block *BasicBlock) NewSelect(cond, x, y value.Value) *InstSelect {
	inst := NewSelect(cond, x, x)
	block.Insts = append(block.Insts, inst)
	return inst
}

// ~~~ [ call ] ~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~

// NewCall returns a new call instruction based on the given callee and function
// arguments.
func (block *BasicBlock) NewCall(callee value.Value, args ...ll.Arg) *InstCall {
	inst := NewCall(callee, args...)
	block.Insts = append(block.Insts, inst)
	return inst
}

// ~~~ [ va_arg ] ~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~

// NewVAArg returns a new va_arg instruction based on the given variable
// argument list and argument type.
func (block *BasicBlock) NewVAArg(list value.Value, argType types.Type) *InstVAArg {
	inst := NewVAArg(list, argType)
	block.Insts = append(block.Insts, inst)
	return inst
}

// ~~~ [ landingpad ] ~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~

// NewLandingPad returns a new landingpad instruction based on the given filter
// and catch clauses.
func (block *BasicBlock) NewLandingPad(clauses ...*ll.Clause) *InstLandingPad {
	inst := NewLandingPad(clauses...)
	block.Insts = append(block.Insts, inst)
	return inst
}

// ~~~ [ catchpad ] ~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~

// NewCatchPad returns a new catchpad instruction based on the given exception
// scope and exception arguments.
func (block *BasicBlock) NewCatchPad(scope value.Value, args ...ll.Arg) *InstCatchPad {
	inst := NewCatchPad(scope, args...)
	block.Insts = append(block.Insts, inst)
	return inst
}

// ~~~ [ cleanuppad ] ~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~

// NewCleanupPad returns a new cleanuppad instruction based on the given
// exception scope and exception arguments.
func (block *BasicBlock) NewCleanupPad(scope ll.ExceptionScope, args ...ll.Arg) *InstCleanupPad {
	inst := NewCleanupPad(scope, args...)
	block.Insts = append(block.Insts, inst)
	return inst
}
