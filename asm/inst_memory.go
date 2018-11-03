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
	// TODO: implement
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
	// TODO: implement
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
	// TODO: implement
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
	// TODO: implement
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
	// TODO: implement
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
	// TODO: implement
	// (optional) Metadata.
	i.Metadata = irMetadataAttachments(old.Metadata())
	return i, nil
}
