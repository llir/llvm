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

// TODO: consider getting rid of UnwindTarget, and let unwind targets be of type
// *ir.BasicBlock, where a nil value indicates the caller, and a non-nil value
// is the unwind target basic block?
type UnwindTarget interface {
	IsUnwindTarget()
}
