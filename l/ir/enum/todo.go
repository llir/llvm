// TODO: consider moving ir/ll to ll, as it's shared between package asm and ir.

// TODO: decide where to merge ir/ll with ir.

package enum

type AtomicOp uint

type Clause struct {
}

type ExceptionScope interface {
	isExceptionScope()
}

type OperandBundle struct {
	// TODO: implement body.
}

// TODO: consider getting rid of UnwindTarget, and let unwind targets be of type
// *ir.BasicBlock, where a nil value indicates the caller, and a non-nil value
// is the unwind target basic block?
type UnwindTarget interface {
	IsUnwindTarget()
}

// TODO: add proper implementations.
type FuncAttribute interface {
	isFuncAttribute()
}

type ParamAttribute interface {
	isParamAttribute()
}

type ReturnAttribute interface {
	isReturnAttribute()
}
