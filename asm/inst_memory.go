package asm

import (
	"fmt"

	"github.com/llir/ll/ast"
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
	i.InAlloca = old.InAlloca() != nil
	// (optional) Swift error.
	i.SwiftError = old.SwiftError() != nil
	// Element type.
	elemType, err := fgen.gen.irType(old.ElemType())
	if err != nil {
		return nil, errors.WithStack(err)
	}
	i.ElemType = elemType
	// (optional) Number of elements.
	if n := old.NElems(); n != nil {
		nelems, err := fgen.astToIRTypeValue(*n)
		if err != nil {
			return nil, errors.WithStack(err)
		}
		i.NElems = nelems
	}
	// (optional) Alignment.
	if n := old.Alignment(); n != nil {
		i.Alignment = irAlignment(*n)
	}
	// (optional) Address space; stored in i.Typ.
	if n := old.AddrSpace(); n != nil {
		i.Type()
		i.Typ.AddrSpace = irAddrSpace(*n)
	}
	// (optional) Metadata.
	i.Metadata = irMetadataAttachments(old.Metadata())
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
	i.Atomic = old.Atomic() != nil
	// (optional) Volatile.
	i.Volatile = old.Volatile() != nil
	// Source address.
	src, err := fgen.astToIRTypeValue(old.Src())
	if err != nil {
		return nil, errors.WithStack(err)
	}
	i.Src = src
	// (optional) Sync scope.
	if n := old.SyncScope(); n != nil {
		i.SyncScope = n.Scope().Text()
	}
	// (optional) Atomic memory ordering constraints.
	if n := old.Ordering(); n != nil {
		i.Ordering = irAtomicOrdering(*n)
	}
	// (optional) Alignment.
	if n := old.Alignment(); n != nil {
		i.Alignment = irAlignment(*n)
	}
	// (optional) Metadata.
	i.Metadata = irMetadataAttachments(old.Metadata())
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
	i.Atomic = old.Atomic() != nil
	// (optional) Volatile.
	i.Volatile = old.Volatile() != nil
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
	if n := old.SyncScope(); n != nil {
		i.SyncScope = n.Scope().Text()
	}
	// (optional) Atomic memory ordering constraints.
	if n := old.Ordering(); n != nil {
		i.Ordering = irAtomicOrdering(*n)
	}
	// (optional) Alignment.
	if n := old.Alignment(); n != nil {
		i.Alignment = irAlignment(*n)
	}
	// (optional) Metadata.
	i.Metadata = irMetadataAttachments(old.Metadata())
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
	if n := old.SyncScope(); n != nil {
		i.SyncScope = n.Scope().Text()
	}
	// Atomic memory ordering constraints.
	i.Ordering = irAtomicOrdering(old.Ordering())
	// (optional) Metadata.
	i.Metadata = irMetadataAttachments(old.Metadata())
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
	i.Weak = old.Weak() != nil
	// (optional) Volatile.
	i.Volatile = old.Volatile() != nil
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
	if n := old.SyncScope(); n != nil {
		i.SyncScope = n.Scope().Text()
	}
	// Atomic memory ordering constraints on success.
	i.SuccessOrdering = irAtomicOrdering(old.SuccessOrdering())
	// Atomic memory ordering constraints on failure.
	i.FailureOrdering = irAtomicOrdering(old.FailureOrdering())
	// (optional) Metadata.
	i.Metadata = irMetadataAttachments(old.Metadata())
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
	i.Volatile = old.Volatile() != nil
	// Atomic operation.
	i.Op = irAtomicOp(old.Op())
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
	if n := old.SyncScope(); n != nil {
		i.SyncScope = n.Scope().Text()
	}
	// Atomic memory ordering constraints.
	i.Ordering = irAtomicOrdering(old.Ordering())
	// (optional) Metadata.
	i.Metadata = irMetadataAttachments(old.Metadata())
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
	i.InBounds = old.InBounds() != nil
	// Element type.
	elemType, err := fgen.gen.irType(old.ElemType())
	if err != nil {
		return nil, errors.WithStack(err)
	}
	i.ElemType = elemType
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
	i.Metadata = irMetadataAttachments(old.Metadata())
	return i, nil
}
