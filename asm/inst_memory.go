package asm

import (
	"fmt"

	"github.com/llir/ll/ast"
	asmenum "github.com/llir/llvm/asm/enum"
	"github.com/llir/llvm/ir"
	"github.com/pkg/errors"
)

// --- [ Memory instructions ] -------------------------------------------------

// ~~~ [ alloca ] ~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~

// astToIRInstAlloca translates the given AST alloca instruction into an
// equivalent IR instruction.
func (fgen *funcGen) astToIRInstAlloca(inst ir.Instruction, old *ast.AllocaInst) (*ir.InstAlloca, error) {
	i, ok := inst.(*ir.InstAlloca)
	if !ok {
		panic(fmt.Errorf("invalid IR instruction for AST instruction; expected *ir.InstAlloca, got %T", inst))
	}
	// (optional) In-alloca.
	_, inAlloca := old.InAlloca()
	i.InAlloca = inAlloca
	// (optional) Swift error.
	_, swiftError := old.SwiftError()
	i.SwiftError = swiftError
	// Element type.
	elemType, err := fgen.gen.irType(old.ElemType())
	if err != nil {
		return nil, errors.WithStack(err)
	}
	i.ElemType = elemType
	// (optional) Number of elements.
	if n, ok := old.NElems(); ok {
		nelems, err := fgen.astToIRTypeValue(n)
		if err != nil {
			return nil, errors.WithStack(err)
		}
		i.NElems = nelems
	}
	// (optional) Alignment.
	if n, ok := old.Align(); ok {
		i.Align = irAlign(n)
	}
	// (optional) Address space; stored in i.Typ.
	if n, ok := old.AddrSpace(); ok {
		i.Type()
		i.Typ.AddrSpace = irAddrSpace(n)
	}
	// (optional) Metadata.
	md, err := fgen.gen.irMetadataAttachments(old.Metadata())
	if err != nil {
		return nil, errors.WithStack(err)
	}
	i.Metadata = md
	return i, nil
}

// ~~~ [ load ] ~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~

// astToIRInstLoad translates the given AST load instruction into an equivalent
// IR instruction.
func (fgen *funcGen) astToIRInstLoad(inst ir.Instruction, old *ast.LoadInst) (*ir.InstLoad, error) {
	i, ok := inst.(*ir.InstLoad)
	if !ok {
		panic(fmt.Errorf("invalid IR instruction for AST instruction; expected *ir.InstLoad, got %T", inst))
	}
	// (optional) Atomic.
	_, atomic := old.Atomic()
	i.Atomic = atomic
	// (optional) Volatile.
	_, volatile := old.Volatile()
	i.Volatile = volatile
	// Source address.
	src, err := fgen.astToIRTypeValue(old.Src())
	if err != nil {
		return nil, errors.WithStack(err)
	}
	i.Src = src
	// (optional) Sync scope.
	if n, ok := old.SyncScope(); ok {
		i.SyncScope = stringLit(n.Scope())
	}
	// (optional) Atomic memory ordering constraints.
	if n, ok := old.Ordering(); ok {
		i.Ordering = asmenum.AtomicOrderingFromString(n.Text())
	}
	// (optional) Alignment.
	if n, ok := old.Align(); ok {
		i.Align = irAlign(n)
	}
	// (optional) Metadata.
	md, err := fgen.gen.irMetadataAttachments(old.Metadata())
	if err != nil {
		return nil, errors.WithStack(err)
	}
	i.Metadata = md
	return i, nil
}

// ~~~ [ store ] ~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~

// astToIRInstStore translates the given AST store instruction into an
// equivalent IR instruction.
func (fgen *funcGen) astToIRInstStore(inst ir.Instruction, old *ast.StoreInst) (*ir.InstStore, error) {
	i, ok := inst.(*ir.InstStore)
	if !ok {
		panic(fmt.Errorf("invalid IR instruction for AST instruction; expected *ir.InstStore, got %T", inst))
	}
	// (optional) Atomic.
	_, atomic := old.Atomic()
	i.Atomic = atomic
	// (optional) Volatile.
	_, volatile := old.Volatile()
	i.Volatile = volatile
	// Source value.
	src, err := fgen.astToIRTypeValue(old.Src())
	if err != nil {
		return nil, errors.WithStack(err)
	}
	i.Src = src
	// Destination address.
	dst, err := fgen.astToIRTypeValue(old.Dst())
	if err != nil {
		return nil, errors.WithStack(err)
	}
	i.Dst = dst
	// (optional) Sync scope.
	if n, ok := old.SyncScope(); ok {
		i.SyncScope = stringLit(n.Scope())
	}
	// (optional) Atomic memory ordering constraints.
	if n, ok := old.Ordering(); ok {
		i.Ordering = asmenum.AtomicOrderingFromString(n.Text())
	}
	// (optional) Alignment.
	if n, ok := old.Align(); ok {
		i.Align = irAlign(n)
	}
	// (optional) Metadata.
	md, err := fgen.gen.irMetadataAttachments(old.Metadata())
	if err != nil {
		return nil, errors.WithStack(err)
	}
	i.Metadata = md
	return i, nil
}

// ~~~ [ fence ] ~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~

// astToIRInstFence translates the given AST fence instruction into an
// equivalent IR instruction.
func (fgen *funcGen) astToIRInstFence(inst ir.Instruction, old *ast.FenceInst) (*ir.InstFence, error) {
	i, ok := inst.(*ir.InstFence)
	if !ok {
		panic(fmt.Errorf("invalid IR instruction for AST instruction; expected *ir.InstFence, got %T", inst))
	}
	// (optional) Sync scope.
	if n, ok := old.SyncScope(); ok {
		i.SyncScope = stringLit(n.Scope())
	}
	// Atomic memory ordering constraints.
	i.Ordering = asmenum.AtomicOrderingFromString(old.Ordering().Text())
	// (optional) Metadata.
	md, err := fgen.gen.irMetadataAttachments(old.Metadata())
	if err != nil {
		return nil, errors.WithStack(err)
	}
	i.Metadata = md
	return i, nil
}

// ~~~ [ cmpxchg ] ~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~

// astToIRInstCmpXchg translates the given AST cmpxchg instruction into an
// equivalent IR instruction.
func (fgen *funcGen) astToIRInstCmpXchg(inst ir.Instruction, old *ast.CmpXchgInst) (*ir.InstCmpXchg, error) {
	i, ok := inst.(*ir.InstCmpXchg)
	if !ok {
		panic(fmt.Errorf("invalid IR instruction for AST instruction; expected *ir.InstCmpXchg, got %T", inst))
	}
	// (optional) Weak.
	_, weak := old.Weak()
	i.Weak = weak
	// (optional) Volatile.
	_, volatile := old.Volatile()
	i.Volatile = volatile
	// Address to read from, compare against and store to.
	ptr, err := fgen.astToIRTypeValue(old.Ptr())
	if err != nil {
		return nil, errors.WithStack(err)
	}
	i.Ptr = ptr
	// Value to compare against.
	cmp, err := fgen.astToIRTypeValue(old.Cmp())
	if err != nil {
		return nil, errors.WithStack(err)
	}
	i.Cmp = cmp
	// New value to store.
	new, err := fgen.astToIRTypeValue(old.New())
	if err != nil {
		return nil, errors.WithStack(err)
	}
	i.New = new
	// (optional) Sync scope.
	if n, ok := old.SyncScope(); ok {
		i.SyncScope = stringLit(n.Scope())
	}
	// Atomic memory ordering constraints on success.
	i.SuccessOrdering = asmenum.AtomicOrderingFromString(old.SuccessOrdering().Text())
	// Atomic memory ordering constraints on failure.
	i.FailureOrdering = asmenum.AtomicOrderingFromString(old.FailureOrdering().Text())
	// (optional) Metadata.
	md, err := fgen.gen.irMetadataAttachments(old.Metadata())
	if err != nil {
		return nil, errors.WithStack(err)
	}
	i.Metadata = md
	return i, nil
}

// ~~~ [ atomicrmw ] ~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~

// astToIRInstAtomicRMW translates the given AST atomicrmw instruction into an
// equivalent IR instruction.
func (fgen *funcGen) astToIRInstAtomicRMW(inst ir.Instruction, old *ast.AtomicRMWInst) (*ir.InstAtomicRMW, error) {
	i, ok := inst.(*ir.InstAtomicRMW)
	if !ok {
		panic(fmt.Errorf("invalid IR instruction for AST instruction; expected *ir.InstAtomicRMW, got %T", inst))
	}
	// (optional) Volatile.
	_, volatile := old.Volatile()
	i.Volatile = volatile
	// Atomic operation.
	i.Op = asmenum.AtomicOpFromString(old.Op().Text())
	// Destination address.
	dst, err := fgen.astToIRTypeValue(old.Dst())
	if err != nil {
		return nil, errors.WithStack(err)
	}
	i.Dst = dst
	// Operand.
	x, err := fgen.astToIRTypeValue(old.X())
	if err != nil {
		return nil, errors.WithStack(err)
	}
	i.X = x
	// (optional) Sync scope.
	if n, ok := old.SyncScope(); ok {
		i.SyncScope = stringLit(n.Scope())
	}
	// Atomic memory ordering constraints.
	i.Ordering = asmenum.AtomicOrderingFromString(old.Ordering().Text())
	// (optional) Metadata.
	md, err := fgen.gen.irMetadataAttachments(old.Metadata())
	if err != nil {
		return nil, errors.WithStack(err)
	}
	i.Metadata = md
	return i, nil
}

// ~~~ [ getelementptr ] ~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~

// astToIRInstGetElementPtr translates the given AST getelementptr instruction
// into an equivalent IR instruction.
func (fgen *funcGen) astToIRInstGetElementPtr(inst ir.Instruction, old *ast.GetElementPtrInst) (*ir.InstGetElementPtr, error) {
	i, ok := inst.(*ir.InstGetElementPtr)
	if !ok {
		panic(fmt.Errorf("invalid IR instruction for AST instruction; expected *ir.InstGetElementPtr, got %T", inst))
	}
	// (optional) In-bounds.
	_, inBounds := old.InBounds()
	i.InBounds = inBounds
	// Element type; already handled in newIRValueInst.
	// Source address.
	src, err := fgen.astToIRTypeValue(old.Src())
	if err != nil {
		return nil, errors.WithStack(err)
	}
	i.Src = src
	// Element indicies.
	for _, oldIndex := range old.Indices() {
		index, err := fgen.astToIRTypeValue(oldIndex)
		if err != nil {
			return nil, errors.WithStack(err)
		}
		i.Indices = append(i.Indices, index)
	}
	// (optional) Metadata.
	md, err := fgen.gen.irMetadataAttachments(old.Metadata())
	if err != nil {
		return nil, errors.WithStack(err)
	}
	i.Metadata = md
	return i, nil
}
