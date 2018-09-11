package ir

import (
	"github.com/llir/l/ir/ll"
	"github.com/llir/l/ir/types"
	"github.com/llir/l/ir/value"
)

// --- [ Other instructions ] --------------------------------------------------

// ~~~ [ icmp ] ~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~

// ICmp is an LLVM IR icmp instruction.
type ICmp struct {
	// Name of local variable associated with the result.
	LocalName string
	// Integer comparison condition.
	Cond ll.ICond
	// Integer scalar or vector operands.
	X, Y value.Value
}

// NewICmp returns a new icmp instruction based on the given integer comparison
// condition and integer scalar or vector operands.
func NewICmp(cond ll.ICond, x, y value.Value) *ICmp {
	return &ICmp{Cond: cond, X: x, Y: y}
}

// ~~~ [ fcmp ] ~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~

// FCmp is an LLVM IR fcmp instruction.
type FCmp struct {
	// Name of local variable associated with the result.
	LocalName string
	// Floating-point comparison condition.
	Cond ll.FCond
	// Floating-point scalar or vector operands.
	X, Y value.Value
}

// NewFCmp returns a new fcmp instruction based on the given floating-point
// comparison condition and floating-point scalar or vector operands.
func NewFCmp(cond ll.FCond, x, y value.Value) *FCmp {
	return &FCmp{Cond: cond, X: x, Y: y}
}

// ~~~ [ phi ] ~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~

// Phi is an LLVM IR phi instruction.
type Phi struct {
	// Name of local variable associated with the result.
	LocalName string
	// Incoming values.
	Incs []*Incoming
}

// NewPhi returns a new phi instruction based on the given incoming values.
func NewPhi(incs ...*Incoming) *Phi {
	return &Phi{Incs: incs}
}

// Incoming is an incoming value of a phi instruction.
type Incoming struct {
	// Incoming value.
	X value.Value
	// Predecessor basic block of the incoming value.
	Pred *BasicBlock
}

// NewIncoming returns a new incoming value based on the given value and
// predecessor basic block.
func NewIncoming(x value.Value, pred *BasicBlock) *Incoming {
	return &Incoming{X: x, Pred: pred}
}

// ~~~ [ select ] ~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~

// Select is an LLVM IR select instruction.
type Select struct {
	// Name of local variable associated with the result.
	LocalName string
	// Selection condition.
	Cond value.Value
	// Operands.
	X, Y value.Value
}

// NewSelect returns a new select instruction based on the given selection
// condition and operands.
func NewSelect(cond, x, y value.Value) *Select {
	return &Select{Cond: cond, X: x, Y: x}
}

// ~~~ [ call ] ~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~

// Call is an LLVM IR call instruction.
type Call struct {
	// Name of local variable associated with the result.
	LocalName string
	// Callee.
	// TODO: specify the set of underlying types of Callee.
	Callee value.Value
	// Function arguments.
	Args []ll.Arg
}

// NewCall returns a new call instruction based on the given callee and function
// arguments.
func NewCall(callee value.Value, args ...ll.Arg) *Call {
	return &Call{Callee: callee, Args: args}
}

// ~~~ [ va_arg ] ~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~

// VAArg is an LLVM IR va_arg instruction.
type VAArg struct {
	// Name of local variable associated with the result.
	LocalName string
	// Variable argument list.
	List value.Value
	// Argument type.
	ArgType types.Type
}

// NewVAArg returns a new va_arg instruction based on the given variable
// argument list and argument type.
func NewVAArg(list value.Value, argType types.Type) *VAArg {
	return &VAArg{List: list, ArgType: argType}
}

// ~~~ [ landingpad ] ~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~

// LandingPad is an LLVM IR landingpad instruction.
type LandingPad struct {
	// Name of local variable associated with the result.
	LocalName string
	// Result type.
	ResultType types.Type
	// Cleanup landing pad.
	Cleanup bool
	// Filter and catch clauses; zero or more if Cleanup is true, otherwise one
	// or more.
	Clauses []*ll.Clause
}

// NewLandingPad returns a new landingpad instruction based on the given filter
// and catch clauses.
func NewLandingPad(clauses ...*ll.Clause) *LandingPad {
	return &LandingPad{Clauses: clauses}
}

// ~~~ [ catchpad ] ~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~

// CatchPad is an LLVM IR catchpad instruction.
type CatchPad struct {
	// Name of local variable associated with the result.
	LocalName string
	// Exception scope.
	Scope value.Value // TODO: assess if Scope is a good name. figure out the specific set of underlying types. rename to From?
	// Exception arguments.
	Args []ll.Arg
}

// NewCatchPad returns a new catchpad instruction based on the given exception
// scope and exception arguments.
func NewCatchPad(scope value.Value, args ...ll.Arg) *CatchPad {
	return &CatchPad{Scope: scope, Args: args}
}

// ~~~ [ cleanuppad ] ~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~

// CleanupPad is an LLVM IR cleanuppad instruction.
type CleanupPad struct {
	// Name of local variable associated with the result.
	LocalName string
	// Exception scope.
	Scope ll.ExceptionScope // TODO: rename to Parent? rename to From?
	// Exception arguments.
	Args []ll.Arg
}

// NewCleanupPad returns a new cleanuppad instruction based on the given
// exception scope and exception arguments.
func NewCleanupPad(scope ll.ExceptionScope, args ...ll.Arg) *CleanupPad {
	return &CleanupPad{Scope: scope, Args: args}
}
