package ir

import (
	"fmt"
	"strings"

	"github.com/llir/llvm/internal/enc"
	"github.com/llir/llvm/ir/constant"
	"github.com/llir/llvm/ir/enum"
	"github.com/llir/llvm/ir/metadata"
	"github.com/llir/llvm/ir/types"
	"github.com/llir/llvm/ir/value"
)

// === [ Terminators ] =========================================================

// Terminator is an LLVM IR terminator instruction (a control flow instruction).
//
// A Terminator has one of the following underlying types.
//
// Terminators
//
// https://llvm.org/docs/LangRef.html#terminator-instructions
//
//    *ir.TermRet           // https://godoc.org/github.com/llir/llvm/ir#TermRet
//    *ir.TermBr            // https://godoc.org/github.com/llir/llvm/ir#TermBr
//    *ir.TermCondBr        // https://godoc.org/github.com/llir/llvm/ir#TermCondBr
//    *ir.TermSwitch        // https://godoc.org/github.com/llir/llvm/ir#TermSwitch
//    *ir.TermIndirectBr    // https://godoc.org/github.com/llir/llvm/ir#TermIndirectBr
//    *ir.TermInvoke        // https://godoc.org/github.com/llir/llvm/ir#TermInvoke
//    *ir.TermResume        // https://godoc.org/github.com/llir/llvm/ir#TermResume
//    *ir.TermCatchSwitch   // https://godoc.org/github.com/llir/llvm/ir#TermCatchSwitch
//    *ir.TermCatchRet      // https://godoc.org/github.com/llir/llvm/ir#TermCatchRet
//    *ir.TermCleanupRet    // https://godoc.org/github.com/llir/llvm/ir#TermCleanupRet
//    *ir.TermUnreachable   // https://godoc.org/github.com/llir/llvm/ir#TermUnreachable
type Terminator interface {
	// Def returns the LLVM syntax representation of the terminator.
	Def() string
	// Succs returns the successor basic blocks of the terminator.
	Succs() []*BasicBlock
}

// --- [ ret ] -----------------------------------------------------------------

// TermRet is an LLVM IR ret terminator.
type TermRet struct {
	// Return value; or nil if void return.
	X value.Value

	// extra.

	// (optional) Metadata.
	Metadata []*metadata.MetadataAttachment
}

// NewRet returns a new ret terminator based on the given return value. A nil
// return value indicates a void return.
func NewRet(x value.Value) *TermRet {
	return &TermRet{X: x}
}

// Succs returns the successor basic blocks of the terminator.
func (*TermRet) Succs() []*BasicBlock {
	// no successors.
	return nil
}

// Def returns the LLVM syntax representation of the terminator.
func (term *TermRet) Def() string {
	// Void return instruction.
	//
	// 'ret' XTyp=VoidType Metadata=(',' MetadataAttachment)+?
	//
	// Value return instruction.
	//
	// 'ret' XTyp=ConcreteType X=Value Metadata=(',' MetadataAttachment)+?
	buf := &strings.Builder{}
	if term.X == nil {
		buf.WriteString("ret void")
	} else {
		fmt.Fprintf(buf, "ret %s", term.X)
	}
	for _, md := range term.Metadata {
		fmt.Fprintf(buf, ", %s", md)
	}
	return buf.String()
}

// --- [ br ] ------------------------------------------------------------------

// TermBr is an unconditional LLVM IR br terminator.
type TermBr struct {
	// Target branch.
	Target *BasicBlock

	// extra.

	// Successor basic blocks of the terminator.
	Successors []*BasicBlock
	// (optional) Metadata.
	Metadata []*metadata.MetadataAttachment
}

// NewBr returns a new unconditional br terminator based on the given target
// basic block.
func NewBr(target *BasicBlock) *TermBr {
	return &TermBr{Target: target}
}

// Succs returns the successor basic blocks of the terminator.
func (term *TermBr) Succs() []*BasicBlock {
	// Cache successors if not present.
	if term.Successors == nil {
		term.Successors = []*BasicBlock{term.Target}
	}
	return term.Successors
}

// Def returns the LLVM syntax representation of the terminator.
func (term *TermBr) Def() string {
	// 'br' Target=Label Metadata=(',' MetadataAttachment)+?
	buf := &strings.Builder{}
	fmt.Fprintf(buf, "br %s", term.Target)
	for _, md := range term.Metadata {
		fmt.Fprintf(buf, ", %s", md)
	}
	return buf.String()
}

// --- [ conditional br ] ------------------------------------------------------

// TermCondBr is a conditional LLVM IR br terminator.
type TermCondBr struct {
	// Branching condition.
	Cond value.Value
	// True condition target branch.
	TargetTrue *BasicBlock
	// False condition target branch.
	TargetFalse *BasicBlock

	// extra.

	// Successor basic blocks of the terminator.
	Successors []*BasicBlock
	// (optional) Metadata.
	Metadata []*metadata.MetadataAttachment
}

// NewCondBr returns a new conditional br terminator based on the given
// branching condition and conditional target basic blocks.
func NewCondBr(cond value.Value, targetTrue, targetFalse *BasicBlock) *TermCondBr {
	return &TermCondBr{Cond: cond, TargetTrue: targetTrue, TargetFalse: targetFalse}
}

// Succs returns the successor basic blocks of the terminator.
func (term *TermCondBr) Succs() []*BasicBlock {
	// Cache successors if not present.
	if term.Successors == nil {
		term.Successors = []*BasicBlock{term.TargetTrue, term.TargetFalse}
	}
	return term.Successors
}

// Def returns the LLVM syntax representation of the terminator.
func (term *TermCondBr) Def() string {
	// 'br' CondTyp=IntType Cond=Value ',' TargetTrue=Label ',' TargetFalse=Label
	// Metadata=(',' MetadataAttachment)+?
	buf := &strings.Builder{}
	fmt.Fprintf(buf, "br %s, %s, %s", term.Cond, term.TargetTrue, term.TargetFalse)
	for _, md := range term.Metadata {
		fmt.Fprintf(buf, ", %s", md)
	}
	return buf.String()
}

// --- [ switch ] --------------------------------------------------------------

// TermSwitch is an LLVM IR switch terminator.
type TermSwitch struct {
	// Control variable.
	X value.Value
	// Default target branch.
	TargetDefault *BasicBlock
	// Switch cases.
	Cases []*Case

	// extra.

	// Successor basic blocks of the terminator.
	Successors []*BasicBlock
	// (optional) Metadata.
	Metadata []*metadata.MetadataAttachment
}

// NewSwitch returns a new switch terminator based on the given control
// variable, default target basic block and switch cases.
func NewSwitch(x value.Value, targetDefault *BasicBlock, cases ...*Case) *TermSwitch {
	return &TermSwitch{X: x, TargetDefault: targetDefault, Cases: cases}
}

// Succs returns the successor basic blocks of the terminator.
func (term *TermSwitch) Succs() []*BasicBlock {
	// Cache successors if not present.
	if term.Successors == nil {
		succs := make([]*BasicBlock, 0, 1+len(term.Cases))
		succs = append(succs, term.TargetDefault)
		for _, c := range term.Cases {
			succs = append(succs, c.Target)
		}
		term.Successors = succs
	}
	return term.Successors
}

// Def returns the LLVM syntax representation of the terminator.
func (term *TermSwitch) Def() string {
	// 'switch' X=TypeValue ',' Default=Label '[' Cases=Case* ']' Metadata=(','
	// MetadataAttachment)+?
	buf := &strings.Builder{}
	fmt.Fprintf(buf, "switch %s, %s [\n", term.X, term.TargetDefault)
	for _, c := range term.Cases {
		fmt.Fprintf(buf, "\t\t%s\n", c)
	}
	buf.WriteString("\t]")
	for _, md := range term.Metadata {
		fmt.Fprintf(buf, ", %s", md)
	}
	return buf.String()
}

// ~~~ [ Switch case ] ~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~

// Case is a switch case.
type Case struct {
	// Case comparand.
	X constant.Constant // integer constant or interger constant expression
	// Case target branch.
	Target *BasicBlock
}

// NewCase returns a new switch case based on the given case comparand and
// target basic block.
func NewCase(x constant.Constant, target *BasicBlock) *Case {
	return &Case{X: x, Target: target}
}

// String returns the string representation of the switch case.
func (c *Case) String() string {
	// X=TypeConst ',' Target=Label
	return fmt.Sprintf("%s, %s", c.X, c.Target)
}

// --- [ indirectbr ] ----------------------------------------------------------

// TermIndirectBr is an LLVM IR indirectbr terminator.
type TermIndirectBr struct {
	// Target address.
	Addr value.Value // blockaddress
	// Set of valid target basic blocks.
	ValidTargets []*BasicBlock

	// extra.

	// (optional) Metadata.
	Metadata []*metadata.MetadataAttachment
}

// NewIndirectBr returns a new indirectbr terminator based on the given target
// address (derived from a blockaddress constant) and set of valid target basic
// blocks.
func NewIndirectBr(addr constant.Constant, validTargets ...*BasicBlock) *TermIndirectBr {
	return &TermIndirectBr{Addr: addr, ValidTargets: validTargets}
}

// Succs returns the successor basic blocks of the terminator.
func (term *TermIndirectBr) Succs() []*BasicBlock {
	return term.ValidTargets
}

// Def returns the LLVM syntax representation of the terminator.
func (term *TermIndirectBr) Def() string {
	// 'indirectbr' Addr=TypeValue ',' '[' ValidTargets=(Label separator ',')+
	// ']' Metadata=(',' MetadataAttachment)+?
	buf := &strings.Builder{}
	fmt.Fprintf(buf, "indirectbr %s, [", term.Addr)
	for i, target := range term.ValidTargets {
		if i != 0 {
			buf.WriteString(", ")
		}
		buf.WriteString(target.String())
	}
	buf.WriteString("]")
	for _, md := range term.Metadata {
		fmt.Fprintf(buf, ", %s", md)
	}
	return buf.String()
}

// --- [ invoke ] --------------------------------------------------------------

// TermInvoke is an LLVM IR invoke terminator.
type TermInvoke struct {
	// Name of local variable associated with the result.
	LocalName string
	// Invokee (callee function).
	// TODO: specify the set of underlying types of Invokee.
	Invokee value.Value
	// Function arguments.
	//
	// Arg has one of the following underlying types:
	//    value.Value
	//    TODO: add metadata value?
	Args []value.Value
	// Normal control flow return point.
	Normal *BasicBlock
	// Exception control flow return point.
	Exception *BasicBlock

	// extra.

	// Type of result produced by the terminator, or function signature of the
	// invokee (as used when invokee is variadic).
	Typ types.Type
	// Successor basic blocks of the terminator.
	Successors []*BasicBlock
	// (optional) Calling convention; zero if not present.
	CallingConv enum.CallingConv
	// (optional) Return attributes.
	ReturnAttrs []ReturnAttribute
	// (optional) Address space; zero if not present.
	AddrSpace types.AddrSpace
	// (optional) Function attributes.
	FuncAttrs []FuncAttribute
	// (optional) Operand bundles.
	OperandBundles []*OperandBundle
	// (optional) Metadata.
	Metadata []*metadata.MetadataAttachment
}

// NewInvoke returns a new invoke terminator based on the given invokee, function
// arguments and control flow return points for normal and exceptional
// execution.
//
// TODO: specify the set of underlying types of invokee.
func NewInvoke(invokee value.Value, args []value.Value, normal, exception *BasicBlock) *TermInvoke {
	term := &TermInvoke{Invokee: invokee, Args: args, Normal: normal, Exception: exception}
	// Compute type.
	term.Type()
	return term
}

// String returns the LLVM syntax representation of the terminator as a type-
// value pair.
func (term *TermInvoke) String() string {
	return fmt.Sprintf("%s %s", term.Type(), term.Ident())
}

// Type returns the type of the terminator.
func (term *TermInvoke) Type() types.Type {
	// Cache type if not present.
	if term.Typ == nil {
		t, ok := term.Invokee.Type().(*types.PointerType)
		if !ok {
			panic(fmt.Errorf("invalid invokee type; expected *types.PointerType, got %T", term.Invokee.Type()))
		}
		sig, ok := t.ElemType.(*types.FuncType)
		if !ok {
			panic(fmt.Errorf("invalid invokee type; expected *types.FuncType, got %T", t.ElemType))
		}
		if sig.Variadic {
			term.Typ = sig
		} else {
			term.Typ = sig.RetType
		}
	}
	if t, ok := term.Typ.(*types.FuncType); ok {
		return t.RetType
	}
	return term.Typ
}

// Ident returns the identifier associated with the terminator.
func (term *TermInvoke) Ident() string {
	return enc.Local(term.LocalName)
}

// Name returns the name of the terminator.
func (term *TermInvoke) Name() string {
	return term.LocalName
}

// SetName sets the name of the terminator.
func (term *TermInvoke) SetName(name string) {
	term.LocalName = name
}

// Succs returns the successor basic blocks of the terminator.
func (term *TermInvoke) Succs() []*BasicBlock {
	// Cache successors if not present.
	if term.Successors == nil {
		term.Successors = []*BasicBlock{term.Normal, term.Exception}
	}
	return term.Successors
}

// Def returns the LLVM syntax representation of the terminator.
func (term *TermInvoke) Def() string {
	// 'invoke' CallingConvopt ReturnAttrs=ReturnAttribute* AddrSpaceopt Typ=Type
	// Invokee=Value '(' Args ')' FuncAttrs=FuncAttribute* OperandBundles=('['
	// (OperandBundle separator ',')+ ']')? 'to' Normal=Label 'unwind'
	// Exception=Label Metadata=(',' MetadataAttachment)+?
	buf := &strings.Builder{}
	if !term.Type().Equal(types.Void) {
		fmt.Fprintf(buf, "%s = ", term.Ident())
	}
	buf.WriteString("invoke")
	if term.CallingConv != enum.CallingConvNone {
		fmt.Fprintf(buf, " %s", callingConvString(term.CallingConv))
	}
	for _, attr := range term.ReturnAttrs {
		fmt.Fprintf(buf, " %s", attr)
	}
	// Use function signature instead of return type for variadic functions.
	typ := term.Type()
	if t, ok := term.Typ.(*types.FuncType); ok {
		if t.Variadic {
			typ = t
		}
	}
	fmt.Fprintf(buf, " %s %s(", typ, term.Invokee.Ident())
	for i, arg := range term.Args {
		if i != 0 {
			buf.WriteString(", ")
		}
		buf.WriteString(arg.String())
	}
	buf.WriteString(")")
	for _, attr := range term.FuncAttrs {
		fmt.Fprintf(buf, " %s", attr)
	}
	if len(term.OperandBundles) > 0 {
		buf.WriteString(" [ ")
		for i, operandBundle := range term.OperandBundles {
			if i != 0 {
				buf.WriteString(", ")
			}
			buf.WriteString(operandBundle.String())
		}
		buf.WriteString(" ]")
	}
	fmt.Fprintf(buf, "\n\t\tto %s unwind %s", term.Normal, term.Exception)
	for _, md := range term.Metadata {
		fmt.Fprintf(buf, ", %s", md)
	}
	return buf.String()
}

// --- [ resume ] --------------------------------------------------------------

// TermResume is an LLVM IR resume terminator.
type TermResume struct {
	// Exception argument to propagate.
	X value.Value

	// extra.

	// (optional) Metadata.
	Metadata []*metadata.MetadataAttachment
}

// NewResume returns a new resume terminator based on the given exception
// argument to propagate.
func NewResume(x value.Value) *TermResume {
	return &TermResume{X: x}
}

// Succs returns the successor basic blocks of the terminator.
func (term *TermResume) Succs() []*BasicBlock {
	// no successors.
	return nil
}

// Def returns the LLVM syntax representation of the terminator.
func (term *TermResume) Def() string {
	// 'resume' X=TypeValue Metadata=(',' MetadataAttachment)+?
	buf := &strings.Builder{}
	fmt.Fprintf(buf, "resume %s", term.X)
	for _, md := range term.Metadata {
		fmt.Fprintf(buf, ", %s", md)
	}
	return buf.String()
}

// --- [ catchswitch ] ---------------------------------------------------------

// TermCatchSwitch is an LLVM IR catchswitch terminator.
type TermCatchSwitch struct {
	// Name of local variable associated with the result.
	LocalName string
	// Exception scope.
	Scope ExceptionScope // TODO: rename to Parent? rename to From?
	// Exception handlers.
	Handlers []*BasicBlock
	// Unwind target; basic block or caller function.
	UnwindTarget UnwindTarget // TODO: rename to To? rename to DefaultTarget?

	// extra.

	// Successor basic blocks of the terminator.
	Successors []*BasicBlock
	// (optional) Metadata.
	Metadata []*metadata.MetadataAttachment
}

// NewCatchSwitch returns a new catchswitch terminator based on the given
// exception scope, exception handlers and unwind target.
func NewCatchSwitch(scope ExceptionScope, handlers []*BasicBlock, unwindTarget UnwindTarget) *TermCatchSwitch {
	return &TermCatchSwitch{Scope: scope, Handlers: handlers, UnwindTarget: unwindTarget}
}

// String returns the LLVM syntax representation of the terminator as a type-
// value pair.
func (term *TermCatchSwitch) String() string {
	return fmt.Sprintf("%s %s", term.Type(), term.Ident())
}

// Type returns the type of the terminator.
func (term *TermCatchSwitch) Type() types.Type {
	return types.Token
}

// Ident returns the identifier associated with the terminator.
func (term *TermCatchSwitch) Ident() string {
	return enc.Local(term.LocalName)
}

// Name returns the name of the terminator.
func (term *TermCatchSwitch) Name() string {
	return term.LocalName
}

// SetName sets the name of the terminator.
func (term *TermCatchSwitch) SetName(name string) {
	term.LocalName = name
}

// Succs returns the successor basic blocks of the terminator.
func (term *TermCatchSwitch) Succs() []*BasicBlock {
	// Cache successors if not present.
	if term.Successors == nil {
		if unwindTarget, ok := term.UnwindTarget.(*BasicBlock); ok {
			term.Successors = append(term.Handlers, unwindTarget)
		} else {
			term.Successors = term.Handlers
		}
	}
	return term.Successors
}

// Def returns the LLVM syntax representation of the terminator.
func (term *TermCatchSwitch) Def() string {
	// 'catchswitch' 'within' Scope=ExceptionScope '[' Handlers=(Label separator
	// ',')+ ']' 'unwind' UnwindTarget=UnwindTarget Metadata=(','
	// MetadataAttachment)+?
	buf := &strings.Builder{}
	fmt.Fprintf(buf, "%s = ", term.Ident())
	fmt.Fprintf(buf, "catchswitch within %s [", term.Scope.Ident())
	for i, handler := range term.Handlers {
		if i != 0 {
			buf.WriteString(", ")
		}
		buf.WriteString(handler.String())
	}
	fmt.Fprintf(buf, "] unwind %s", term.UnwindTarget)
	for _, md := range term.Metadata {
		fmt.Fprintf(buf, ", %s", md)
	}
	return buf.String()
}

// --- [ catchret ] ------------------------------------------------------------

// TermCatchRet is an LLVM IR catchret terminator.
type TermCatchRet struct {
	// Exit catchpad.
	From *InstCatchPad
	// Target basic block to transfer control flow to.
	To *BasicBlock

	// extra.

	// Successor basic blocks of the terminator.
	Successors []*BasicBlock
	// (optional) Metadata.
	Metadata []*metadata.MetadataAttachment
}

// NewCatchRet returns a new catchret terminator based on the given exit
// catchpad and target basic block.
func NewCatchRet(from *InstCatchPad, to *BasicBlock) *TermCatchRet {
	return &TermCatchRet{From: from, To: to}
}

// Succs returns the successor basic blocks of the terminator.
func (term *TermCatchRet) Succs() []*BasicBlock {
	// Cache successors if not present.
	if term.Successors == nil {
		term.Successors = []*BasicBlock{term.To}
	}
	return term.Successors
}

// Def returns the LLVM syntax representation of the terminator.
func (term *TermCatchRet) Def() string {
	// 'catchret' 'from' From=Value 'to' To=Label Metadata=(','
	// MetadataAttachment)+?
	buf := &strings.Builder{}
	fmt.Fprintf(buf, "catchret from %s to %s", term.From.Ident(), term.To)
	for _, md := range term.Metadata {
		fmt.Fprintf(buf, ", %s", md)
	}
	return buf.String()
}

// --- [ cleanupret ] ----------------------------------------------------------

// TermCleanupRet is an LLVM IR cleanupret terminator.
type TermCleanupRet struct {
	// Exit cleanuppad.
	From *InstCleanupPad
	// Unwind target; basic block or caller function.
	UnwindTarget UnwindTarget

	// extra.

	// Successor basic blocks of the terminator.
	Successors []*BasicBlock
	// (optional) Metadata.
	Metadata []*metadata.MetadataAttachment
}

// NewCleanupRet returns a new cleanupret terminator based on the given exit
// cleanuppad and unwind target.
func NewCleanupRet(from *InstCleanupPad, unwindTarget UnwindTarget) *TermCleanupRet {
	return &TermCleanupRet{From: from, UnwindTarget: unwindTarget}
}

// Succs returns the successor basic blocks of the terminator.
func (term *TermCleanupRet) Succs() []*BasicBlock {
	// Cache successors if not present.
	if term.Successors == nil {
		if unwindTarget, ok := term.UnwindTarget.(*BasicBlock); ok {
			term.Successors = []*BasicBlock{unwindTarget}
		} else {
			term.Successors = []*BasicBlock{}
		}
	}
	return term.Successors
}

// Def returns the LLVM syntax representation of the terminator.
func (term *TermCleanupRet) Def() string {
	// 'cleanupret' 'from' From=Value 'unwind' UnwindTarget Metadata=(','
	// MetadataAttachment)+?
	buf := &strings.Builder{}
	fmt.Fprintf(buf, "cleanupret from %s unwind %s", term.From.Ident(), term.UnwindTarget)
	for _, md := range term.Metadata {
		fmt.Fprintf(buf, ", %s", md)
	}
	return buf.String()
}

// --- [ unreachable ] ---------------------------------------------------------

// TermUnreachable is an LLVM IR unreachable terminator.
type TermUnreachable struct {
	// extra.

	// (optional) Metadata.
	Metadata []*metadata.MetadataAttachment
}

// NewUnreachable returns a new unreachable terminator.
func NewUnreachable() *TermUnreachable {
	return &TermUnreachable{}
}

// Succs returns the successor basic blocks of the terminator.
func (term *TermUnreachable) Succs() []*BasicBlock {
	// no successors.
	return nil
}

// Def returns the LLVM syntax representation of the terminator.
func (term *TermUnreachable) Def() string {
	// 'unreachable' Metadata=(',' MetadataAttachment)+?
	buf := &strings.Builder{}
	buf.WriteString("unreachable")
	for _, md := range term.Metadata {
		fmt.Fprintf(buf, ", %s", md)
	}
	return buf.String()
}
