// TODO: decide where to merge ir/ll with ir.

package ll

import (
	"github.com/llir/l/ir/constant"
	"github.com/llir/l/ir/value"
)

type Arg interface {
	isArg()
}

type AtomicOp uint

type AtomicOrdering uint

type Case struct {
	X      constant.Constant // integer constant or interger constant expression
	Target value.Value       // *ir.BasicBlock
}

type Clause struct {
}

type ExceptionScope interface {
	isExceptionScope()
}

type FCond uint

type ICond uint

// Incoming is an incoming value of a phi instruction.
type Incoming struct {
	// Incoming value.
	X value.Value
	// Predecessor basic block of the incoming value.
	Pred value.Value // *ir.BasicBlock
}

type UnwindTarget interface {
	isUnwindTarget()
}
