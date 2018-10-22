package ir

import (
	"fmt"

	"github.com/llir/l/internal/enc"
	"github.com/llir/l/ir/enum"
	"github.com/llir/l/ir/types"
	"github.com/llir/l/ir/value"
)

// --- [ Other instructions ] --------------------------------------------------

// ~~~ [ icmp ] ~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~

// InstICmp is an LLVM IR icmp instruction.
type InstICmp struct {
	// Name of local variable associated with the result.
	LocalName string
	// Integer comparison predicate.
	Pred enum.IPred
	// Integer scalar or vector operands.
	X, Y value.Value // integer scalar, pointer, integer vector or pointer vector.

	// extra.

	// Type of result produced by the instruction.
	Typ types.Type // boolean or boolean vector
	// (optional) Metadata.
	// TODO: add metadata.
}

// NewICmp returns a new icmp instruction based on the given integer comparison
// predicate and integer scalar or vector operands.
func NewICmp(pred enum.IPred, x, y value.Value) *InstICmp {
	return &InstICmp{Pred: pred, X: x, Y: y}
}

// String returns the LLVM syntax representation of the instruction as a
// type-value pair.
func (inst *InstICmp) String() string {
	return fmt.Sprintf("%v %v", inst.Type(), inst.Ident())
}

// Type returns the type of the instruction.
func (inst *InstICmp) Type() types.Type {
	// Cache type if not present.
	if inst.Typ == nil {
		switch xType := inst.X.Type().(type) {
		case *types.IntType, *types.PointerType:
			inst.Typ = types.I1
		case *types.VectorType:
			inst.Typ = types.NewVector(xType.Len, types.I1)
		default:
			panic(fmt.Errorf("invalid icmp operand type; expected *types.IntType, *types.PointerType or *types.VectorType, got %T", xType))
		}
	}
	return inst.Typ
}

// Ident returns the identifier associated with the instruction.
func (inst *InstICmp) Ident() string {
	return enc.Local(inst.LocalName)
}

// Name returns the name of the instruction.
func (inst *InstICmp) Name() string {
	return inst.LocalName
}

// SetName sets the name of the instruction.
func (inst *InstICmp) SetName(name string) {
	inst.LocalName = name
}

// ~~~ [ fcmp ] ~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~

// InstFCmp is an LLVM IR fcmp instruction.
type InstFCmp struct {
	// Name of local variable associated with the result.
	LocalName string
	// Floating-point comparison predicate.
	Pred enum.FPred
	// Floating-point scalar or vector operands.
	X, Y value.Value // floating-point scalar or floating-point vector

	// extra.

	// Type of result produced by the instruction.
	Typ types.Type // boolean or boolean vector
	// (optional) Fast math flags.
	FastMathFlags []enum.FastMathFlag
	// Metadata.
	// TODO: add metadata.
}

// NewFCmp returns a new fcmp instruction based on the given floating-point
// comparison predicate and floating-point scalar or vector operands.
func NewFCmp(pred enum.FPred, x, y value.Value) *InstFCmp {
	return &InstFCmp{Pred: pred, X: x, Y: y}
}

// String returns the LLVM syntax representation of the instruction as a
// type-value pair.
func (inst *InstFCmp) String() string {
	return fmt.Sprintf("%v %v", inst.Type(), inst.Ident())
}

// Type returns the type of the instruction.
func (inst *InstFCmp) Type() types.Type {
	// Cache type if not present.
	if inst.Typ == nil {
		switch xType := inst.X.Type().(type) {
		case *types.FloatType:
			inst.Typ = types.I1
		case *types.VectorType:
			inst.Typ = types.NewVector(xType.Len, types.I1)
		default:
			panic(fmt.Errorf("invalid fcmp operand type; expected *types.FloatType or *types.VectorType, got %T", xType))
		}
	}
	return inst.Typ
}

// Ident returns the identifier associated with the instruction.
func (inst *InstFCmp) Ident() string {
	return enc.Local(inst.LocalName)
}

// Name returns the name of the instruction.
func (inst *InstFCmp) Name() string {
	return inst.LocalName
}

// SetName sets the name of the instruction.
func (inst *InstFCmp) SetName(name string) {
	inst.LocalName = name
}

// ~~~ [ phi ] ~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~

// InstPhi is an LLVM IR phi instruction.
type InstPhi struct {
	// Name of local variable associated with the result.
	LocalName string
	// Incoming values.
	Incs []*Incoming

	// extra.

	// Type of result produced by the instruction.
	Typ types.Type // type of incoming value
	// (optional) Metadata.
	// TODO: add metadata.
}

// NewPhi returns a new phi instruction based on the given incoming values.
func NewPhi(incs ...*Incoming) *InstPhi {
	return &InstPhi{Incs: incs}
}

// String returns the LLVM syntax representation of the instruction as a
// type-value pair.
func (inst *InstPhi) String() string {
	return fmt.Sprintf("%v %v", inst.Type(), inst.Ident())
}

// Type returns the type of the instruction.
func (inst *InstPhi) Type() types.Type {
	// Cache type if not present.
	if inst.Typ == nil {
		inst.Typ = inst.Incs[0].X.Type()
	}
	return inst.Typ
}

// Ident returns the identifier associated with the instruction.
func (inst *InstPhi) Ident() string {
	return enc.Local(inst.LocalName)
}

// Name returns the name of the instruction.
func (inst *InstPhi) Name() string {
	return inst.LocalName
}

// SetName sets the name of the instruction.
func (inst *InstPhi) SetName(name string) {
	inst.LocalName = name
}

// ___ [ Incoming value ] ______________________________________________________

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

// InstSelect is an LLVM IR select instruction.
type InstSelect struct {
	// Name of local variable associated with the result.
	LocalName string
	// Selection condition.
	Cond value.Value // boolean or boolean vector
	// Operands.
	X, Y value.Value

	// extra.

	// Type of result produced by the instruction.
	Typ types.Type
	// (optional) Metadata.
	// TODO: add metadata.
}

// NewSelect returns a new select instruction based on the given selection
// condition and operands.
func NewSelect(cond, x, y value.Value) *InstSelect {
	return &InstSelect{Cond: cond, X: x, Y: x}
}

// String returns the LLVM syntax representation of the instruction as a
// type-value pair.
func (inst *InstSelect) String() string {
	return fmt.Sprintf("%v %v", inst.Type(), inst.Ident())
}

// Type returns the type of the instruction.
func (inst *InstSelect) Type() types.Type {
	if inst.Typ == nil {
		inst.Typ = inst.X.Type()
	}
	return inst.Typ
}

// Ident returns the identifier associated with the instruction.
func (inst *InstSelect) Ident() string {
	return enc.Local(inst.LocalName)
}

// Name returns the name of the instruction.
func (inst *InstSelect) Name() string {
	return inst.LocalName
}

// SetName sets the name of the instruction.
func (inst *InstSelect) SetName(name string) {
	inst.LocalName = name
}

// ~~~ [ call ] ~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~

// InstCall is an LLVM IR call instruction.
type InstCall struct {
	// Name of local variable associated with the result.
	LocalName string
	// Callee.
	// TODO: specify the set of underlying types of Callee.
	Callee value.Value
	// Function arguments.
	Args []enum.Arg // TODO: move enum.Arg to ir.Arg.

	// extra.

	// Type of result produced by the instruction, or function signature of the
	// callee (as used when callee is variadic).
	Typ types.Type
	// (optional) Tail; zero if not present.
	Tail enum.Tail
	// (optional) Fast math flags.
	FastMathFlags []enum.FastMathFlag
	// (optional) Calling convention; zero if not present.
	CallingConv enum.CallingConv
	// (optional) Return attributes.
	ReturnAttrs []enum.ReturnAttribute
	// (optional) Address space; zero if not present.
	AddrSpace types.AddrSpace
	// (optional) Function attributes.
	FuncAttrs []enum.FuncAttribute
	// (optional) Operand bundles.
	OperandBundles []enum.OperandBundle
	// Metadata.
	// TODO: add metadata.
}

// NewCall returns a new call instruction based on the given callee and function
// arguments.
//
// TODO: specify the set of underlying types of callee.
func NewCall(callee value.Value, args ...enum.Arg) *InstCall {
	return &InstCall{Callee: callee, Args: args}
}

// String returns the LLVM syntax representation of the instruction as a
// type-value pair.
func (inst *InstCall) String() string {
	return fmt.Sprintf("%v %v", inst.Type(), inst.Ident())
}

// Type returns the type of the instruction.
func (inst *InstCall) Type() types.Type {
	// TODO: cache Typ from Callee if nil?
	if t, ok := inst.Typ.(*types.FuncType); ok {
		return t.RetType
	}
	return inst.Typ
}

// Ident returns the identifier associated with the instruction.
func (inst *InstCall) Ident() string {
	return enc.Local(inst.LocalName)
}

// Name returns the name of the instruction.
func (inst *InstCall) Name() string {
	return inst.LocalName
}

// SetName sets the name of the instruction.
func (inst *InstCall) SetName(name string) {
	inst.LocalName = name
}

// ~~~ [ va_arg ] ~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~

// InstVAArg is an LLVM IR va_arg instruction.
type InstVAArg struct {
	// Name of local variable associated with the result.
	LocalName string
	// Variable argument list.
	VAList value.Value
	// Argument type.
	ArgType types.Type

	// extra.

	// (optional) Metadata.
	// TODO: add metadata.
}

// NewVAArg returns a new va_arg instruction based on the given variable
// argument list and argument type.
func NewVAArg(vaList value.Value, argType types.Type) *InstVAArg {
	return &InstVAArg{VAList: vaList, ArgType: argType}
}

// String returns the LLVM syntax representation of the instruction as a
// type-value pair.
func (inst *InstVAArg) String() string {
	return fmt.Sprintf("%v %v", inst.Type(), inst.Ident())
}

// Type returns the type of the instruction.
func (inst *InstVAArg) Type() types.Type {
	return inst.ArgType
}

// Ident returns the identifier associated with the instruction.
func (inst *InstVAArg) Ident() string {
	return enc.Local(inst.LocalName)
}

// Name returns the name of the instruction.
func (inst *InstVAArg) Name() string {
	return inst.LocalName
}

// SetName sets the name of the instruction.
func (inst *InstVAArg) SetName(name string) {
	inst.LocalName = name
}

// ~~~ [ landingpad ] ~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~

// InstLandingPad is an LLVM IR landingpad instruction.
type InstLandingPad struct {
	// Name of local variable associated with the result.
	LocalName string
	// Result type.
	ResultType types.Type
	// Cleanup landing pad.
	Cleanup bool
	// Filter and catch clauses; zero or more if Cleanup is true, otherwise one
	// or more.
	Clauses []*enum.Clause

	// extra.

	// (optional) Metadata.
	// TODO: add metadata.
}

// NewLandingPad returns a new landingpad instruction based on the given result
// type and filter/catch clauses.
func NewLandingPad(resultType types.Type, clauses ...*enum.Clause) *InstLandingPad {
	return &InstLandingPad{ResultType: resultType, Clauses: clauses}
}

// String returns the LLVM syntax representation of the instruction as a
// type-value pair.
func (inst *InstLandingPad) String() string {
	return fmt.Sprintf("%v %v", inst.Type(), inst.Ident())
}

// Type returns the type of the instruction.
func (inst *InstLandingPad) Type() types.Type {
	return inst.ResultType
}

// Ident returns the identifier associated with the instruction.
func (inst *InstLandingPad) Ident() string {
	return enc.Local(inst.LocalName)
}

// Name returns the name of the instruction.
func (inst *InstLandingPad) Name() string {
	return inst.LocalName
}

// SetName sets the name of the instruction.
func (inst *InstLandingPad) SetName(name string) {
	inst.LocalName = name
}

// ~~~ [ catchpad ] ~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~

// InstCatchPad is an LLVM IR catchpad instruction.
type InstCatchPad struct {
	// Name of local variable associated with the result.
	LocalName string
	// Exception scope.
	Scope *TermCatchSwitch // TODO: rename to From? rename to Within?
	// Exception arguments.
	Args []enum.Arg

	// extra.

	// (optional) Metadata.
	// TODO: add metadata.
}

// NewCatchPad returns a new catchpad instruction based on the given exception
// scope and exception arguments.
func NewCatchPad(scope *TermCatchSwitch, args ...enum.Arg) *InstCatchPad {
	return &InstCatchPad{Scope: scope, Args: args}
}

// String returns the LLVM syntax representation of the instruction as a
// type-value pair.
func (inst *InstCatchPad) String() string {
	return fmt.Sprintf("%v %v", inst.Type(), inst.Ident())
}

// Type returns the type of the instruction.
func (inst *InstCatchPad) Type() types.Type {
	return types.Token
}

// Ident returns the identifier associated with the instruction.
func (inst *InstCatchPad) Ident() string {
	return enc.Local(inst.LocalName)
}

// Name returns the name of the instruction.
func (inst *InstCatchPad) Name() string {
	return inst.LocalName
}

// SetName sets the name of the instruction.
func (inst *InstCatchPad) SetName(name string) {
	inst.LocalName = name
}

// ~~~ [ cleanuppad ] ~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~

// InstCleanupPad is an LLVM IR cleanuppad instruction.
type InstCleanupPad struct {
	// Name of local variable associated with the result.
	LocalName string
	// Exception scope.
	Scope enum.ExceptionScope // TODO: rename to Parent? rename to From?
	// Exception arguments.
	Args []enum.Arg

	// extra.

	// (optional) Metadata.
	// TODO: add metadata.
}

// NewCleanupPad returns a new cleanuppad instruction based on the given
// exception scope and exception arguments.
func NewCleanupPad(scope enum.ExceptionScope, args ...enum.Arg) *InstCleanupPad {
	return &InstCleanupPad{Scope: scope, Args: args}
}

// String returns the LLVM syntax representation of the instruction as a
// type-value pair.
func (inst *InstCleanupPad) String() string {
	return fmt.Sprintf("%v %v", inst.Type(), inst.Ident())
}

// Type returns the type of the instruction.
func (inst *InstCleanupPad) Type() types.Type {
	return types.Token
}

// Ident returns the identifier associated with the instruction.
func (inst *InstCleanupPad) Ident() string {
	return enc.Local(inst.LocalName)
}

// Name returns the name of the instruction.
func (inst *InstCleanupPad) Name() string {
	return inst.LocalName
}

// SetName sets the name of the instruction.
func (inst *InstCleanupPad) SetName(name string) {
	inst.LocalName = name
}
