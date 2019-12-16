package asm

import (
	"fmt"
	"strconv"

	"github.com/llir/ll/ast"
	asmenum "github.com/llir/llvm/asm/enum"
	"github.com/llir/llvm/internal/gep"
	"github.com/llir/llvm/ir"
	"github.com/llir/llvm/ir/types"
	"github.com/llir/llvm/ir/value"
	"github.com/pkg/errors"
)

// === [ Create IR ] ===========================================================

// newAllocaInst returns a new IR alloca instruction (without body but with
// type) based on the given AST alloca instruction.
func (fgen *funcGen) newAllocaInst(ident ir.LocalIdent, old *ast.AllocaInst) (*ir.InstAlloca, error) {
	elemType, err := fgen.gen.irType(old.ElemType())
	if err != nil {
		return nil, errors.WithStack(err)
	}
	inst := &ir.InstAlloca{LocalIdent: ident, ElemType: elemType}
	// Cache inst.Typ.
	inst.Type()
	return inst, nil
}

// newLoadInst returns a new IR load instruction (without body but with type)
// based on the given AST load instruction.
func (fgen *funcGen) newLoadInst(ident ir.LocalIdent, old *ast.LoadInst) (*ir.InstLoad, error) {
	elemType, err := fgen.gen.irType(old.ElemType())
	if err != nil {
		return nil, errors.WithStack(err)
	}
	return &ir.InstLoad{LocalIdent: ident, ElemType: elemType}, nil
}

// newCmpXchgInst returns a new IR cmpxchg instruction (without body but with
// type) based on the given AST cmpxchg instruction.
func (fgen *funcGen) newCmpXchgInst(ident ir.LocalIdent, old *ast.CmpXchgInst) (*ir.InstCmpXchg, error) {
	oldType, err := fgen.gen.irType(old.New().Typ())
	if err != nil {
		return nil, errors.WithStack(err)
	}
	typ := types.NewStruct(oldType, types.I8)
	return &ir.InstCmpXchg{LocalIdent: ident, Typ: typ}, nil
}

// newAtomicRMWInst returns a new IR atomicrmw instruction (without body but
// with type) based on the given AST atomicrmw instruction.
func (fgen *funcGen) newAtomicRMWInst(ident ir.LocalIdent, old *ast.AtomicRMWInst) (*ir.InstAtomicRMW, error) {
	dstType, err := fgen.gen.irType(old.Dst().Typ())
	if err != nil {
		return nil, errors.WithStack(err)
	}
	dt, ok := dstType.(*types.PointerType)
	if !ok {
		panic(fmt.Errorf("invalid pointer type; expected *types.PointerType, got %T", dstType))
	}
	return &ir.InstAtomicRMW{LocalIdent: ident, Typ: dt.ElemType}, nil
}

// newGetElementPtrInst returns a new IR getelementptr instruction (without body
// but with type) based on the given AST getelementptr instruction.
func (fgen *funcGen) newGetElementPtrInst(ident ir.LocalIdent, old *ast.GetElementPtrInst) (*ir.InstGetElementPtr, error) {
	// TODO: handle address space of Src?
	elemType, err := fgen.gen.irType(old.ElemType())
	if err != nil {
		return nil, errors.WithStack(err)
	}
	srcType, err := fgen.gen.irType(old.Src().Typ())
	if err != nil {
		return nil, errors.WithStack(err)
	}
	typ, err := fgen.gen.gepInstType(elemType, srcType, old.Indices())
	if err != nil {
		return nil, errors.WithStack(err)
	}
	return &ir.InstGetElementPtr{LocalIdent: ident, ElemType: elemType, Typ: typ}, nil
}

// === [ Translate AST to IR ] =================================================

// --- [ alloca ] --------------------------------------------------------------

// irAllocaInst translates the given AST alloca instruction into an equivalent
// IR instruction.
func (fgen *funcGen) irAllocaInst(new ir.Instruction, old *ast.AllocaInst) error {
	inst, ok := new.(*ir.InstAlloca)
	if !ok {
		panic(fmt.Errorf("invalid IR instruction for AST instruction; expected *ir.InstAlloca, got %T", new))
	}
	// Element type.
	elemType, err := fgen.gen.irType(old.ElemType())
	if err != nil {
		return errors.WithStack(err)
	}
	inst.ElemType = elemType
	// (optional) Number of elements.
	if n, ok := old.NElems(); ok {
		nelems, err := fgen.irTypeValue(n)
		if err != nil {
			return errors.WithStack(err)
		}
		inst.NElems = nelems
	}
	// (optional) In-alloca.
	_, inst.InAlloca = old.InAlloca()
	// (optional) Swift error.
	_, inst.SwiftError = old.SwiftError()
	// (optional) Alignment.
	if n, ok := old.Align(); ok {
		inst.Align = irAlign(n)
	}
	// (optional) Address space; stored in i.Typ.
	if n, ok := old.AddrSpace(); ok {
		inst.Typ.AddrSpace = irAddrSpace(n)
	}
	// (optional) Metadata.
	md, err := fgen.gen.irMetadataAttachments(old.Metadata())
	if err != nil {
		return errors.WithStack(err)
	}
	inst.Metadata = md
	return nil
}

// --- [ load ] ----------------------------------------------------------------

// irLoadInst translates the given AST load instruction into an equivalent IR
// instruction.
func (fgen *funcGen) irLoadInst(new ir.Instruction, old *ast.LoadInst) error {
	inst, ok := new.(*ir.InstLoad)
	if !ok {
		panic(fmt.Errorf("invalid IR instruction for AST instruction; expected *ir.InstLoad, got %T", new))
	}
	// Source address.
	src, err := fgen.irTypeValue(old.Src())
	if err != nil {
		return errors.WithStack(err)
	}
	inst.Src = src
	// (optional) Atomic.
	_, inst.Atomic = old.Atomic()
	// (optional) Volatile.
	_, inst.Volatile = old.Volatile()
	// (optional) Sync scope.
	if n, ok := old.SyncScope(); ok {
		inst.SyncScope = stringLit(n.Scope())
	}
	// (optional) Atomic memory ordering constraints.
	if n, ok := old.Ordering(); ok {
		inst.Ordering = asmenum.AtomicOrderingFromString(n.Text())
	}
	// (optional) Alignment.
	if n, ok := old.Align(); ok {
		inst.Align = irAlign(n)
	}
	// (optional) Metadata.
	md, err := fgen.gen.irMetadataAttachments(old.Metadata())
	if err != nil {
		return errors.WithStack(err)
	}
	inst.Metadata = md
	return nil
}

// --- [ store ] ---------------------------------------------------------------

// irStoreInst translates the given AST store instruction into an equivalent IR
// instruction.
func (fgen *funcGen) irStoreInst(new ir.Instruction, old *ast.StoreInst) error {
	inst, ok := new.(*ir.InstStore)
	if !ok {
		panic(fmt.Errorf("invalid IR instruction for AST instruction; expected *ir.InstStore, got %T", new))
	}
	// Source value.
	src, err := fgen.irTypeValue(old.Src())
	if err != nil {
		return errors.WithStack(err)
	}
	inst.Src = src
	// Destination address.
	dst, err := fgen.irTypeValue(old.Dst())
	if err != nil {
		return errors.WithStack(err)
	}
	inst.Dst = dst
	// (optional) Atomic.
	_, inst.Atomic = old.Atomic()
	// (optional) Volatile.
	_, inst.Volatile = old.Volatile()
	// (optional) Sync scope.
	if n, ok := old.SyncScope(); ok {
		inst.SyncScope = stringLit(n.Scope())
	}
	// (optional) Atomic memory ordering constraints.
	if n, ok := old.Ordering(); ok {
		inst.Ordering = asmenum.AtomicOrderingFromString(n.Text())
	}
	// (optional) Alignment.
	if n, ok := old.Align(); ok {
		inst.Align = irAlign(n)
	}
	// (optional) Metadata.
	md, err := fgen.gen.irMetadataAttachments(old.Metadata())
	if err != nil {
		return errors.WithStack(err)
	}
	inst.Metadata = md
	return nil
}

// --- [ fence ] ---------------------------------------------------------------

// irFenceInst translates the given AST fence instruction into an equivalent IR
// instruction.
func (fgen *funcGen) irFenceInst(new ir.Instruction, old *ast.FenceInst) error {
	inst, ok := new.(*ir.InstFence)
	if !ok {
		panic(fmt.Errorf("invalid IR instruction for AST instruction; expected *ir.InstFence, got %T", new))
	}
	// Atomic memory ordering constraints.
	inst.Ordering = asmenum.AtomicOrderingFromString(old.Ordering().Text())
	// (optional) Sync scope.
	if n, ok := old.SyncScope(); ok {
		inst.SyncScope = stringLit(n.Scope())
	}
	// (optional) Metadata.
	md, err := fgen.gen.irMetadataAttachments(old.Metadata())
	if err != nil {
		return errors.WithStack(err)
	}
	inst.Metadata = md
	return nil
}

// --- [ cmpxchg ] -------------------------------------------------------------

// irCmpXchgInst translates the given AST cmpxchg instruction into an equivalent
// IR instruction.
func (fgen *funcGen) irCmpXchgInst(new ir.Instruction, old *ast.CmpXchgInst) error {
	inst, ok := new.(*ir.InstCmpXchg)
	if !ok {
		panic(fmt.Errorf("invalid IR instruction for AST instruction; expected *ir.InstCmpXchg, got %T", new))
	}
	// Address to read from, compare against and store to.
	ptr, err := fgen.irTypeValue(old.Ptr())
	if err != nil {
		return errors.WithStack(err)
	}
	inst.Ptr = ptr
	// Value to compare against.
	cmp, err := fgen.irTypeValue(old.Cmp())
	if err != nil {
		return errors.WithStack(err)
	}
	inst.Cmp = cmp
	// New value to store.
	newValue, err := fgen.irTypeValue(old.New())
	if err != nil {
		return errors.WithStack(err)
	}
	inst.New = newValue
	// Atomic memory ordering constraints on success.
	inst.SuccessOrdering = asmenum.AtomicOrderingFromString(old.SuccessOrdering().Text())
	// Atomic memory ordering constraints on failure.
	inst.FailureOrdering = asmenum.AtomicOrderingFromString(old.FailureOrdering().Text())
	// (optional) Weak.
	_, inst.Weak = old.Weak()
	// (optional) Volatile.
	_, inst.Volatile = old.Volatile()
	// (optional) Sync scope.
	if n, ok := old.SyncScope(); ok {
		inst.SyncScope = stringLit(n.Scope())
	}
	// (optional) Metadata.
	md, err := fgen.gen.irMetadataAttachments(old.Metadata())
	if err != nil {
		return errors.WithStack(err)
	}
	inst.Metadata = md
	return nil
}

// --- [ atomicrmw ] -----------------------------------------------------------

// irAtomicRMWInst translates the given AST atomicrmw instruction into an
// equivalent IR instruction.
func (fgen *funcGen) irAtomicRMWInst(new ir.Instruction, old *ast.AtomicRMWInst) error {
	inst, ok := new.(*ir.InstAtomicRMW)
	if !ok {
		panic(fmt.Errorf("invalid IR instruction for AST instruction; expected *ir.InstAtomicRMW, got %T", new))
	}
	// Atomic operation.
	inst.Op = asmenum.AtomicOpFromString(old.Op().Text())
	// Destination address.
	dst, err := fgen.irTypeValue(old.Dst())
	if err != nil {
		return errors.WithStack(err)
	}
	inst.Dst = dst
	// Operand.
	x, err := fgen.irTypeValue(old.X())
	if err != nil {
		return errors.WithStack(err)
	}
	inst.X = x
	// Atomic memory ordering constraints.
	inst.Ordering = asmenum.AtomicOrderingFromString(old.Ordering().Text())
	// (optional) Volatile.
	_, inst.Volatile = old.Volatile()
	// (optional) Sync scope.
	if n, ok := old.SyncScope(); ok {
		inst.SyncScope = stringLit(n.Scope())
	}
	// (optional) Metadata.
	md, err := fgen.gen.irMetadataAttachments(old.Metadata())
	if err != nil {
		return errors.WithStack(err)
	}
	inst.Metadata = md
	return nil
}

// --- [ getelementptr ] -------------------------------------------------------

// irGetElementPtrInst translates the given AST getelementptr instruction into
// an equivalent IR instruction.
func (fgen *funcGen) irGetElementPtrInst(new ir.Instruction, old *ast.GetElementPtrInst) error {
	inst, ok := new.(*ir.InstGetElementPtr)
	if !ok {
		panic(fmt.Errorf("invalid IR instruction for AST instruction; expected *ir.InstGetElementPtr, got %T", new))
	}
	// Element type; already handled in fgen.newValueInst.
	// Source address.
	src, err := fgen.irTypeValue(old.Src())
	if err != nil {
		return errors.WithStack(err)
	}
	inst.Src = src
	// Element indicies.
	if oldIndices := old.Indices(); len(oldIndices) > 0 {
		inst.Indices = make([]value.Value, len(oldIndices))
		for i, oldIndex := range oldIndices {
			index, err := fgen.irTypeValue(oldIndex)
			if err != nil {
				return errors.WithStack(err)
			}
			inst.Indices[i] = index
		}
	}
	// (optional) In-bounds.
	_, inst.InBounds = old.InBounds()
	// (optional) Metadata.
	md, err := fgen.gen.irMetadataAttachments(old.Metadata())
	if err != nil {
		return errors.WithStack(err)
	}
	inst.Metadata = md
	return nil
}

// ### [ Helper functions ] ####################################################

// gepInstType computes the result type of a getelementptr instruction.
//
//    getelementptr ElemType, Src, Indices
func (gen *generator) gepInstType(elemType, src types.Type, indices []ast.TypeValue) (types.Type, error) {
	var idxs []gep.Index
	for _, index := range indices {
		var idx gep.Index
		if indexVal, ok := index.Val().(ast.Constant); ok {
			idx = getIndex(indexVal)
		} else {
			idx = gep.Index{HasVal: false}
			// Check if index is of vector type.
			indexType, err := gen.irType(index.Typ())
			if err != nil {
				return nil, errors.WithStack(err)
			}
			if indexType, ok := indexType.(*types.VectorType); ok {
				idx.VectorLen = indexType.Len
			}
		}
		idxs = append(idxs, idx)
	}
	return gep.ResultType(elemType, src, idxs), nil
}

// NOTE: keep getIndex in sync with getIndex in:
//
//    * ast/inst_memory.go
//    * ir/inst_memory.go
//    * ir/constant/expr_memory.go
//
// The reference point and source of truth is in ir/constant/expr_memory.go.

// getIndex returns the gep index corresponding to the given constant index.
func getIndex(index ast.Constant) gep.Index {
	switch index := index.(type) {
	case *ast.IntConst:
		val, err := strconv.ParseInt(index.Text(), 10, 64)
		if err != nil {
			panic(fmt.Errorf("unable to parse integer %q; %v", index.Text(), err))
		}
		return gep.NewIndex(val)
	case *ast.BoolConst:
		if boolLit(index.BoolLit()) {
			return gep.NewIndex(1)
		}
		return gep.NewIndex(0)
	case *ast.ZeroInitializerConst:
		return gep.NewIndex(0)
	case *ast.VectorConst:
		// ref: https://llvm.org/docs/LangRef.html#getelementptr-instruction
		//
		// > The getelementptr returns a vector of pointers, instead of a single
		// > address, when one or more of its arguments is a vector. In such
		// > cases, all vector arguments should have the same number of elements,
		// > and every scalar argument will be effectively broadcast into a vector
		// > during address calculation.
		elems := index.Elems()
		if len(elems) == 0 {
			return gep.Index{HasVal: false}
		}
		// Sanity check. All vector elements must be integers, and must have the
		// same value.
		var val int64
		for i, elem := range elems {
			switch elem := elem.Val().(type) {
			case *ast.IntConst:
				x, err := strconv.ParseInt(elem.Text(), 10, 64)
				if err != nil {
					panic(fmt.Errorf("unable to parse integer %q; %v", index.Text(), err))
				}
				if i == 0 {
					val = x
				} else if x != val {
					// since all elements were not identical, we must conclude that
					// the index vector does not have a concrete value.
					return gep.Index{
						HasVal:    false,
						VectorLen: uint64(len(elems)),
					}
				}
			default:
				// TODO: remove debug output.
				panic(fmt.Errorf("support for gep index vector element type %T not yet implemented", elem))
				return gep.Index{HasVal: false}
			}
		}
		return gep.Index{
			HasVal:    true,
			Val:       val,
			VectorLen: uint64(len(elems)),
		}
	case *ast.PtrToIntExpr:
		return gep.Index{HasVal: false}
	case *ast.UndefConst:
		return gep.Index{HasVal: false}
	default:
		// TODO: add support for more constant expressions.
		// TODO: remove debug output.
		panic(fmt.Errorf("support for gep index type %T not yet implemented", index))
		return gep.Index{HasVal: false}
	}
}
