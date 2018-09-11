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

type UnwindTarget interface {
	IsUnwindTarget()
}
