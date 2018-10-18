// TODO: consider moving ir/ll to ll, as it's shared between package asm and ir.

// TODO: decide where to merge ir/ll with ir.

package ll

type Arg interface {
	isArg()
}

type AtomicOp uint

type AtomicOrdering uint

type Clause struct {
}

type ExceptionScope interface {
	isExceptionScope()
}

type FCond uint

type ICond uint

// --- [ Overflow flags ] ------------------------------------------------------

//go:generate stringer -linecomment -type OverflowFlag

// OverflowFlag is an integer overflow flag.
type OverflowFlag uint8

// Overflow flags.
const (
	OverflowFlagNSW OverflowFlag = iota // nsw
	OverflowFlagNUW                     // nuw
)

// --- [ Unwind targets ] ------------------------------------------------------

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

type ReturnAttribute interface {
	isReturnAttribute()
}
