// Dummy values are used to allow for forward references, and are replaced by
// their real values in later stages of parsing.

package irx

import (
	"bytes"
	"fmt"

	"github.com/llir/llvm/internal/enc"
	"github.com/llir/llvm/ir"
	"github.com/llir/llvm/ir/constant"
	"github.com/llir/llvm/ir/types"
	"github.com/llir/llvm/ir/value"
)

// globalDummy represents a dummy value for a given global identifier name and
// type.
type globalDummy struct {
	// Global name.
	name string
	// Type associated with the global identifier.
	typ types.Type
}

// newGlobalDummy returns a new dummy value for the given global identifier name
// and type.
func newGlobalDummy(name string, typ types.Type) *globalDummy {
	return &globalDummy{name: name, typ: typ}
}

// Type returns the type associated with the global identifier.
func (global *globalDummy) Type() types.Type {
	return global.typ
}

// Ident returns the identifier associated with the global.
func (global *globalDummy) Ident() string {
	return enc.Global(global.name)
}

// Name returns the name of the global.
func (global *globalDummy) Name() string {
	return global.name
}

// SetName sets the name of the global.
func (global *globalDummy) SetName(name string) {
	global.name = name
}

// Immutable ensures that only constants can be assigned to the Constant
// interface.
func (*globalDummy) Immutable() {}

// localDummy represents a dummy value for a given local identifier name and
// type.
type localDummy struct {
	// Local name.
	name string
	// Type associated with the local identifier.
	typ types.Type
}

// newLocalDummy returns a new dummy value for the given local identifier name
// and type.
func newLocalDummy(name string, typ types.Type) *localDummy {
	return &localDummy{name: name, typ: typ}
}

// Type returns the type associated with the local identifier.
func (local *localDummy) Type() types.Type {
	return local.typ
}

// Ident returns the identifier associated with the local.
func (local *localDummy) Ident() string {
	return enc.Local(local.name)
}

// Name returns the name of the local.
func (local *localDummy) Name() string {
	return local.name
}

// SetName sets the name of the local.
func (local *localDummy) SetName(name string) {
	local.name = name
}

// namedTypeDummy represents a dummy named type.
type namedTypeDummy struct {
	// Type name.
	name string
}

// String returns the LLVM syntax representation of the type.
func (t *namedTypeDummy) String() string {
	return enc.Local(t.name)
}

// Equal reports whether t and u are of equal type.
func (t *namedTypeDummy) Equal(u types.Type) bool {
	panic("dummy implementation")
}

// newNamedTypeDummy returns a new dummy named type based on the given type
// name.
func newNamedTypeDummy(name string) *namedTypeDummy {
	return &namedTypeDummy{name: name}
}

// instGetElementPtrDummy represents a dummy getelementptr instruction.
type instGetElementPtrDummy struct {
	// Parent basic block.
	parent *ir.BasicBlock
	// Name of the local variable associated with the instruction.
	name string
	// Source address.
	src value.Value
	// Element indices.
	indices []value.Value
	// Source address element type.
	elem types.Type
}

// newGetElementPtrDummy returns a new dummy getelementptr instruction based on
// the given element type, source address and element indices.
func newGetElementPtrDummy(elem types.Type, src value.Value, indices ...value.Value) *instGetElementPtrDummy {
	return &instGetElementPtrDummy{src: src, indices: indices, elem: elem}
}

// Type returns the type of the instruction.
func (inst *instGetElementPtrDummy) Type() types.Type {
	panic("dummy implementation")
}

// Ident returns the identifier associated with the instruction.
func (inst *instGetElementPtrDummy) Ident() string {
	return enc.Local(inst.name)
}

// Name returns the name of the local variable associated with the instruction.
func (inst *instGetElementPtrDummy) Name() string {
	return inst.name
}

// SetName sets the name of the local variable associated with the instruction.
func (inst *instGetElementPtrDummy) SetName(name string) {
	inst.name = name
}

// String returns the LLVM syntax representation of the instruction.
func (inst *instGetElementPtrDummy) String() string {
	buf := &bytes.Buffer{}
	src := inst.src
	fmt.Fprintf(buf, "%s = getelementptr %s, %s %s",
		inst.Ident(),
		inst.elem,
		src.Type(),
		src.Ident())
	for _, index := range inst.indices {
		fmt.Fprintf(buf, ", %s %s",
			index.Type(),
			index.Ident())
	}
	return buf.String()
}

// Parent returns the parent basic block of the instruction.
func (inst *instGetElementPtrDummy) Parent() *ir.BasicBlock {
	return inst.parent
}

// SetParent sets the parent basic block of the instruction.
func (inst *instGetElementPtrDummy) SetParent(parent *ir.BasicBlock) {
	inst.parent = parent
}

// instPhiDummy represents a dummy phi instruction.
type instPhiDummy struct {
	// Parent basic block.
	parent *ir.BasicBlock
	// Name of the local variable associated with the instruction.
	name string
	// Incoming values.
	incs []*incomingDummy
	// Type of the instruction.
	typ types.Type
}

// newPhiDummy returns a new dummy phi instruction based on the given incoming
// values.
func newPhiDummy(typ types.Type, incs ...*incomingDummy) *instPhiDummy {
	return &instPhiDummy{incs: incs, typ: typ}
}

// Type returns the type of the instruction.
func (inst *instPhiDummy) Type() types.Type {
	return inst.typ
}

// Ident returns the identifier associated with the instruction.
func (inst *instPhiDummy) Ident() string {
	return enc.Local(inst.name)
}

// Name returns the name of the local variable associated with the instruction.
func (inst *instPhiDummy) Name() string {
	return inst.name
}

// SetName sets the name of the local variable associated with the instruction.
func (inst *instPhiDummy) SetName(name string) {
	inst.name = name
}

// String returns the LLVM syntax representation of the instruction.
func (inst *instPhiDummy) String() string {
	buf := &bytes.Buffer{}
	fmt.Fprintf(buf, "%s = phi %s ",
		inst.Ident(),
		inst.Type())
	for j, inc := range inst.incs {
		if j != 0 {
			buf.WriteString(", ")
		}
		x, ok := inc.x.(value.Value)
		if !ok {
			panic(fmt.Sprintf("invalid x type; expected value.Value, got %T", inc.x))
		}
		fmt.Fprintf(buf, "[ %s, %s ]",
			x.Ident(),
			inc.pred)
	}
	return buf.String()
}

// Parent returns the parent basic block of the instruction.
func (inst *instPhiDummy) Parent() *ir.BasicBlock {
	return inst.parent
}

// SetParent sets the parent basic block of the instruction.
func (inst *instPhiDummy) SetParent(parent *ir.BasicBlock) {
	inst.parent = parent
}

// incomingDummy represents a dummy incoming value of a phi instruction.
type incomingDummy struct {
	// Incoming value; holds *irx.IntLit, *irx.LocalIdent, ... initially, when
	// created from using irx.NewIncoming since type is not known. irx.NewPhiInst
	// later replaces with dummy values (e.g. *localDummy, ...).
	x interface{}
	// Predecessor basic block of the incoming value.
	pred string
}

// newIncomingDummy returns a new dummy incoming value based on the given value
// and predecessor basic block label name.
func newIncomingDummy(x interface{}, pred string) *incomingDummy {
	return &incomingDummy{x: x, pred: pred}
}

// instCallDummy represents a dummy call instruction.
type instCallDummy struct {
	// Parent basic block.
	parent *ir.BasicBlock
	// Name of the local variable associated with the instruction.
	name string
	// Callee.
	callee string
	// Function arguments.
	args []value.Value
	// Return type.
	ret types.Type
}

// newCallDummy returns a new dummy call instruction based on the given callee
// and function arguments.
func newCallDummy(ret types.Type, callee string, args ...value.Value) *instCallDummy {
	return &instCallDummy{callee: callee, args: args, ret: ret}
}

// Type returns the type of the instruction.
func (inst *instCallDummy) Type() types.Type {
	return inst.ret
}

// Ident returns the identifier associated with the instruction.
func (inst *instCallDummy) Ident() string {
	return enc.Local(inst.name)
}

// Name returns the name of the local variable associated with the instruction.
func (inst *instCallDummy) Name() string {
	return inst.name
}

// SetName sets the name of the local variable associated with the instruction.
func (inst *instCallDummy) SetName(name string) {
	inst.name = name
}

// String returns the LLVM syntax representation of the instruction.
func (inst *instCallDummy) String() string {
	buf := &bytes.Buffer{}
	typ := inst.Type()
	if !typ.Equal(types.Void) {
		fmt.Fprintf(buf, "%s = ", inst.Ident())
	}
	fmt.Fprintf(buf, "call %s %s(",
		typ,
		inst.callee)
	for i, arg := range inst.args {
		if i != 0 {
			buf.WriteString(", ")
		}
		fmt.Fprintf(buf, "%s %s",
			arg.Type(),
			arg.Ident())
	}
	buf.WriteString(")")
	return buf.String()
}

// Parent returns the parent basic block of the instruction.
func (inst *instCallDummy) Parent() *ir.BasicBlock {
	return inst.parent
}

// SetParent sets the parent basic block of the instruction.
func (inst *instCallDummy) SetParent(parent *ir.BasicBlock) {
	inst.parent = parent
}

// termBrDummy represents a dummy unconditional br terminator.
type termBrDummy struct {
	// Parent basic block.
	parent *ir.BasicBlock
	// Target branch.
	target string
}

// newBrDummy returns a new dummy unconditional br terminator based on the given
// target branch.
func newBrDummy(target string) *termBrDummy {
	return &termBrDummy{target: target}
}

// String returns the LLVM syntax representation of the terminator.
func (term *termBrDummy) String() string {
	return fmt.Sprintf("br label %s", term.target)
}

// Parent returns the parent basic block of the terminator.
func (term *termBrDummy) Parent() *ir.BasicBlock {
	return term.parent
}

// SetParent sets the parent basic block of the terminator.
func (term *termBrDummy) SetParent(parent *ir.BasicBlock) {
	term.parent = parent
}

// Successors returns the successor basic blocks of the terminator.
func (term *termBrDummy) Successors() []*ir.BasicBlock {
	panic("dummy implementation")
}

// termCondBrDummy represents a dummy conditional br terminator.
type termCondBrDummy struct {
	// Parent basic block.
	parent *ir.BasicBlock
	// Branching condition.
	cond value.Value
	// Target branch when condition is true.
	targetTrue string
	// Target branch when condition is false.
	targetFalse string
}

// NewCondBr returns a new dummy conditional br terminator based on the given
// branching condition and conditional target branches.
func newCondBrDummy(cond value.Value, targetTrue, targetFalse string) *termCondBrDummy {
	return &termCondBrDummy{cond: cond, targetTrue: targetTrue, targetFalse: targetFalse}
}

// String returns the LLVM syntax representation of the terminator.
func (term *termCondBrDummy) String() string {
	return fmt.Sprintf("br i1 %s, label %s, label %s",
		term.cond.Ident(),
		term.targetTrue,
		term.targetFalse)
}

// Parent returns the parent basic block of the terminator.
func (term *termCondBrDummy) Parent() *ir.BasicBlock {
	return term.parent
}

// SetParent sets the parent basic block of the terminator.
func (term *termCondBrDummy) SetParent(parent *ir.BasicBlock) {
	term.parent = parent
}

// Successors returns the successor basic blocks of the terminator.
func (term *termCondBrDummy) Successors() []*ir.BasicBlock {
	panic("dummy implementation")
}

// termSwitchDummy represents a dummy switch terminator.
type termSwitchDummy struct {
	// Parent basic block.
	parent *ir.BasicBlock
	// Control variable.
	x value.Value
	// Default target branch.
	targetDefault string
	// Switch cases.
	cases []*caseDummy
}

// newSwitchDummy returns a new dummy switch terminator based on the given
// control variable, default target branch and switch cases.
func newSwitchDummy(x value.Value, targetDefault string, cases ...*caseDummy) *termSwitchDummy {
	return &termSwitchDummy{x: x, targetDefault: targetDefault, cases: cases}
}

// String returns the LLVM syntax representation of the terminator.
func (term *termSwitchDummy) String() string {
	buf := &bytes.Buffer{}
	x := term.x
	fmt.Fprintf(buf, "switch %s %s, label %s [\n",
		x.Type(),
		x.Ident(),
		term.targetDefault)
	for _, c := range term.cases {
		x := c.x
		fmt.Fprintf(buf, "\t\t%s %s, label %s\n",
			x.Type(),
			x.Ident(),
			c.target)
	}
	buf.WriteString("\t]")
	return buf.String()
}

// Parent returns the parent basic block of the terminator.
func (term *termSwitchDummy) Parent() *ir.BasicBlock {
	return term.parent
}

// SetParent sets the parent basic block of the terminator.
func (term *termSwitchDummy) SetParent(parent *ir.BasicBlock) {
	term.parent = parent
}

// Successors returns the successor basic blocks of the terminator.
func (term *termSwitchDummy) Successors() []*ir.BasicBlock {
	panic("dummy implementation")
}

// caseDummy represents a dummy case value of a switch terminator.
type caseDummy struct {
	// Case comparand.
	x *constant.Int
	// Case target branch.
	target string
}

// newCaseDummy returns a new dummy switch case based on the given case
// comparand and target branch.
func newCaseDummy(x *constant.Int, target string) *caseDummy {
	return &caseDummy{x: x, target: target}
}
