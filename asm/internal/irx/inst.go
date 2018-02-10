package irx

import (
	"fmt"

	"github.com/llir/llvm/asm/internal/ast"
	"github.com/llir/llvm/internal/enc"
	"github.com/llir/llvm/ir"
	"github.com/llir/llvm/ir/constant"
	"github.com/llir/llvm/ir/types"
	"github.com/llir/llvm/ir/value"
	"github.com/pkg/errors"
)

// --- [ Binary instructions ] -------------------------------------------------

// instAdd resolves the given add instruction, by recursively resolving its
// operands. The boolean return value indicates success.
func (m *Module) instAdd(old *ast.InstAdd, resolved, unresolved map[ast.NamedValue]value.Named) bool {
	if isUnresolved(unresolved, old.X, old.Y) {
		return false
	}
	v := m.getLocal(old.Name)
	inst, ok := v.(*ir.InstAdd)
	if !ok {
		panic(fmt.Errorf("invalid instruction type for instruction %s; expected *ir.InstAdd, got %T", enc.Local(old.Name), v))
	}
	inst.X = m.irValue(old.X)
	inst.Y = m.irValue(old.Y)
	inst.Metadata = m.irMetadata(old.Metadata)
	return true
}

// instFAdd resolves the given fadd instruction, by recursively resolving its
// operands. The boolean return value indicates success.
func (m *Module) instFAdd(old *ast.InstFAdd, resolved, unresolved map[ast.NamedValue]value.Named) bool {
	if isUnresolved(unresolved, old.X, old.Y) {
		return false
	}
	v := m.getLocal(old.Name)
	inst, ok := v.(*ir.InstFAdd)
	if !ok {
		panic(fmt.Errorf("invalid instruction type for instruction %s; expected *ir.InstFAdd, got %T", enc.Local(old.Name), v))
	}
	inst.X = m.irValue(old.X)
	inst.Y = m.irValue(old.Y)
	inst.Metadata = m.irMetadata(old.Metadata)
	return true
}

// instSub resolves the given sub instruction, by recursively resolving its
// operands. The boolean return value indicates success.
func (m *Module) instSub(old *ast.InstSub, resolved, unresolved map[ast.NamedValue]value.Named) bool {
	if isUnresolved(unresolved, old.X, old.Y) {
		return false
	}
	v := m.getLocal(old.Name)
	inst, ok := v.(*ir.InstSub)
	if !ok {
		panic(fmt.Errorf("invalid instruction type for instruction %s; expected *ir.InstSub, got %T", enc.Local(old.Name), v))
	}
	inst.X = m.irValue(old.X)
	inst.Y = m.irValue(old.Y)
	inst.Metadata = m.irMetadata(old.Metadata)
	return true
}

// instFSub resolves the given fsub instruction, by recursively resolving its
// operands. The boolean return value indicates success.
func (m *Module) instFSub(old *ast.InstFSub, resolved, unresolved map[ast.NamedValue]value.Named) bool {
	if isUnresolved(unresolved, old.X, old.Y) {
		return false
	}
	v := m.getLocal(old.Name)
	inst, ok := v.(*ir.InstFSub)
	if !ok {
		panic(fmt.Errorf("invalid instruction type for instruction %s; expected *ir.InstFSub, got %T", enc.Local(old.Name), v))
	}
	inst.X = m.irValue(old.X)
	inst.Y = m.irValue(old.Y)
	inst.Metadata = m.irMetadata(old.Metadata)
	return true
}

// instMul resolves the given mul instruction, by recursively resolving its
// operands. The boolean return value indicates success.
func (m *Module) instMul(old *ast.InstMul, resolved, unresolved map[ast.NamedValue]value.Named) bool {
	if isUnresolved(unresolved, old.X, old.Y) {
		return false
	}
	v := m.getLocal(old.Name)
	inst, ok := v.(*ir.InstMul)
	if !ok {
		panic(fmt.Errorf("invalid instruction type for instruction %s; expected *ir.InstMul, got %T", enc.Local(old.Name), v))
	}
	inst.X = m.irValue(old.X)
	inst.Y = m.irValue(old.Y)
	inst.Metadata = m.irMetadata(old.Metadata)
	return true
}

// instFMul resolves the given fmul instruction, by recursively resolving its
// operands. The boolean return value indicates success.
func (m *Module) instFMul(old *ast.InstFMul, resolved, unresolved map[ast.NamedValue]value.Named) bool {
	if isUnresolved(unresolved, old.X, old.Y) {
		return false
	}
	v := m.getLocal(old.Name)
	inst, ok := v.(*ir.InstFMul)
	if !ok {
		panic(fmt.Errorf("invalid instruction type for instruction %s; expected *ir.InstFMul, got %T", enc.Local(old.Name), v))
	}
	inst.X = m.irValue(old.X)
	inst.Y = m.irValue(old.Y)
	inst.Metadata = m.irMetadata(old.Metadata)
	return true
}

// instUDiv resolves the given udiv instruction, by recursively resolving its
// operands. The boolean return value indicates success.
func (m *Module) instUDiv(old *ast.InstUDiv, resolved, unresolved map[ast.NamedValue]value.Named) bool {
	if isUnresolved(unresolved, old.X, old.Y) {
		return false
	}
	v := m.getLocal(old.Name)
	inst, ok := v.(*ir.InstUDiv)
	if !ok {
		panic(fmt.Errorf("invalid instruction type for instruction %s; expected *ir.InstUDiv, got %T", enc.Local(old.Name), v))
	}
	inst.X = m.irValue(old.X)
	inst.Y = m.irValue(old.Y)
	inst.Metadata = m.irMetadata(old.Metadata)
	return true
}

// instSDiv resolves the given sdiv instruction, by recursively resolving its
// operands. The boolean return value indicates success.
func (m *Module) instSDiv(old *ast.InstSDiv, resolved, unresolved map[ast.NamedValue]value.Named) bool {
	if isUnresolved(unresolved, old.X, old.Y) {
		return false
	}
	v := m.getLocal(old.Name)
	inst, ok := v.(*ir.InstSDiv)
	if !ok {
		panic(fmt.Errorf("invalid instruction type for instruction %s; expected *ir.InstSDiv, got %T", enc.Local(old.Name), v))
	}
	inst.X = m.irValue(old.X)
	inst.Y = m.irValue(old.Y)
	inst.Metadata = m.irMetadata(old.Metadata)
	return true
}

// instFDiv resolves the given fdiv instruction, by recursively resolving its
// operands. The boolean return value indicates success.
func (m *Module) instFDiv(old *ast.InstFDiv, resolved, unresolved map[ast.NamedValue]value.Named) bool {
	if isUnresolved(unresolved, old.X, old.Y) {
		return false
	}
	v := m.getLocal(old.Name)
	inst, ok := v.(*ir.InstFDiv)
	if !ok {
		panic(fmt.Errorf("invalid instruction type for instruction %s; expected *ir.InstFDiv, got %T", enc.Local(old.Name), v))
	}
	inst.X = m.irValue(old.X)
	inst.Y = m.irValue(old.Y)
	inst.Metadata = m.irMetadata(old.Metadata)
	return true
}

// instURem resolves the given urem instruction, by recursively resolving its
// operands. The boolean return value indicates success.
func (m *Module) instURem(old *ast.InstURem, resolved, unresolved map[ast.NamedValue]value.Named) bool {
	if isUnresolved(unresolved, old.X, old.Y) {
		return false
	}
	v := m.getLocal(old.Name)
	inst, ok := v.(*ir.InstURem)
	if !ok {
		panic(fmt.Errorf("invalid instruction type for instruction %s; expected *ir.InstURem, got %T", enc.Local(old.Name), v))
	}
	inst.X = m.irValue(old.X)
	inst.Y = m.irValue(old.Y)
	inst.Metadata = m.irMetadata(old.Metadata)
	return true
}

// instSRem resolves the given srem instruction, by recursively resolving its
// operands. The boolean return value indicates success.
func (m *Module) instSRem(old *ast.InstSRem, resolved, unresolved map[ast.NamedValue]value.Named) bool {
	if isUnresolved(unresolved, old.X, old.Y) {
		return false
	}
	v := m.getLocal(old.Name)
	inst, ok := v.(*ir.InstSRem)
	if !ok {
		panic(fmt.Errorf("invalid instruction type for instruction %s; expected *ir.InstSRem, got %T", enc.Local(old.Name), v))
	}
	inst.X = m.irValue(old.X)
	inst.Y = m.irValue(old.Y)
	inst.Metadata = m.irMetadata(old.Metadata)
	return true
}

// instFRem resolves the given frem instruction, by recursively resolving its
// operands. The boolean return value indicates success.
func (m *Module) instFRem(old *ast.InstFRem, resolved, unresolved map[ast.NamedValue]value.Named) bool {
	if isUnresolved(unresolved, old.X, old.Y) {
		return false
	}
	v := m.getLocal(old.Name)
	inst, ok := v.(*ir.InstFRem)
	if !ok {
		panic(fmt.Errorf("invalid instruction type for instruction %s; expected *ir.InstFRem, got %T", enc.Local(old.Name), v))
	}
	inst.X = m.irValue(old.X)
	inst.Y = m.irValue(old.Y)
	inst.Metadata = m.irMetadata(old.Metadata)
	return true
}

// --- [ Bitwise instructions ] ------------------------------------------------

// instShl resolves the given shl instruction, by recursively resolving its
// operands. The boolean return value indicates success.
func (m *Module) instShl(old *ast.InstShl, resolved, unresolved map[ast.NamedValue]value.Named) bool {
	if isUnresolved(unresolved, old.X, old.Y) {
		return false
	}
	v := m.getLocal(old.Name)
	inst, ok := v.(*ir.InstShl)
	if !ok {
		panic(fmt.Errorf("invalid instruction type for instruction %s; expected *ir.InstShl, got %T", enc.Local(old.Name), v))
	}
	inst.X = m.irValue(old.X)
	inst.Y = m.irValue(old.Y)
	inst.Metadata = m.irMetadata(old.Metadata)
	return true
}

// instLShr resolves the given lshr instruction, by recursively resolving its
// operands. The boolean return value indicates success.
func (m *Module) instLShr(old *ast.InstLShr, resolved, unresolved map[ast.NamedValue]value.Named) bool {
	if isUnresolved(unresolved, old.X, old.Y) {
		return false
	}
	v := m.getLocal(old.Name)
	inst, ok := v.(*ir.InstLShr)
	if !ok {
		panic(fmt.Errorf("invalid instruction type for instruction %s; expected *ir.InstLShr, got %T", enc.Local(old.Name), v))
	}
	inst.X = m.irValue(old.X)
	inst.Y = m.irValue(old.Y)
	inst.Metadata = m.irMetadata(old.Metadata)
	return true
}

// instAShr resolves the given ashr instruction, by recursively resolving its
// operands. The boolean return value indicates success.
func (m *Module) instAShr(old *ast.InstAShr, resolved, unresolved map[ast.NamedValue]value.Named) bool {
	if isUnresolved(unresolved, old.X, old.Y) {
		return false
	}
	v := m.getLocal(old.Name)
	inst, ok := v.(*ir.InstAShr)
	if !ok {
		panic(fmt.Errorf("invalid instruction type for instruction %s; expected *ir.InstAShr, got %T", enc.Local(old.Name), v))
	}
	inst.X = m.irValue(old.X)
	inst.Y = m.irValue(old.Y)
	inst.Metadata = m.irMetadata(old.Metadata)
	return true
}

// instAnd resolves the given and instruction, by recursively resolving its
// operands. The boolean return value indicates success.
func (m *Module) instAnd(old *ast.InstAnd, resolved, unresolved map[ast.NamedValue]value.Named) bool {
	if isUnresolved(unresolved, old.X, old.Y) {
		return false
	}
	v := m.getLocal(old.Name)
	inst, ok := v.(*ir.InstAnd)
	if !ok {
		panic(fmt.Errorf("invalid instruction type for instruction %s; expected *ir.InstAnd, got %T", enc.Local(old.Name), v))
	}
	inst.X = m.irValue(old.X)
	inst.Y = m.irValue(old.Y)
	inst.Metadata = m.irMetadata(old.Metadata)
	return true
}

// instOr resolves the given or instruction, by recursively resolving its
// operands. The boolean return value indicates success.
func (m *Module) instOr(old *ast.InstOr, resolved, unresolved map[ast.NamedValue]value.Named) bool {
	if isUnresolved(unresolved, old.X, old.Y) {
		return false
	}
	v := m.getLocal(old.Name)
	inst, ok := v.(*ir.InstOr)
	if !ok {
		panic(fmt.Errorf("invalid instruction type for instruction %s; expected *ir.InstOr, got %T", enc.Local(old.Name), v))
	}
	inst.X = m.irValue(old.X)
	inst.Y = m.irValue(old.Y)
	inst.Metadata = m.irMetadata(old.Metadata)
	return true
}

// instXor resolves the given xor instruction, by recursively resolving its
// operands. The boolean return value indicates success.
func (m *Module) instXor(old *ast.InstXor, resolved, unresolved map[ast.NamedValue]value.Named) bool {
	if isUnresolved(unresolved, old.X, old.Y) {
		return false
	}
	v := m.getLocal(old.Name)
	inst, ok := v.(*ir.InstXor)
	if !ok {
		panic(fmt.Errorf("invalid instruction type for instruction %s; expected *ir.InstXor, got %T", enc.Local(old.Name), v))
	}
	inst.X = m.irValue(old.X)
	inst.Y = m.irValue(old.Y)
	inst.Metadata = m.irMetadata(old.Metadata)
	return true
}

// --- [ Vector instructions ] -------------------------------------------------

// instExtractElement resolves the given extractelement instruction, by
// recursively resolving its operands. The boolean return value indicates
// success.
func (m *Module) instExtractElement(old *ast.InstExtractElement, resolved, unresolved map[ast.NamedValue]value.Named) bool {
	if isUnresolved(unresolved, old.X, old.Index) {
		return false
	}
	v := m.getLocal(old.Name)
	inst, ok := v.(*ir.InstExtractElement)
	if !ok {
		panic(fmt.Errorf("invalid instruction type for instruction %s; expected *ir.InstExtractElement, got %T", enc.Local(old.Name), v))
	}
	x := m.irValue(old.X)
	t, ok := x.Type().(*types.VectorType)
	if !ok {
		panic(fmt.Errorf("invalid vector type; expected *types.VectorType, got %T", x.Type()))
	}
	inst.Typ = t.Elem
	inst.X = x
	inst.Index = m.irValue(old.Index)
	inst.Metadata = m.irMetadata(old.Metadata)
	return true
}

// instInsertElement resolves the given insertelement instruction, by
// recursively resolving its operands. The boolean return value indicates
// success.
func (m *Module) instInsertElement(old *ast.InstInsertElement, resolved, unresolved map[ast.NamedValue]value.Named) bool {
	if isUnresolved(unresolved, old.X, old.Elem, old.Index) {
		return false
	}
	v := m.getLocal(old.Name)
	inst, ok := v.(*ir.InstInsertElement)
	if !ok {
		panic(fmt.Errorf("invalid instruction type for instruction %s; expected *ir.InstInsertElement, got %T", enc.Local(old.Name), v))
	}
	inst.X = m.irValue(old.X)
	inst.Elem = m.irValue(old.Elem)
	inst.Index = m.irValue(old.Index)
	inst.Metadata = m.irMetadata(old.Metadata)
	return true
}

// instShuffleVector resolves the given shufflevector instruction, by
// recursively resolving its operands. The boolean return value indicates
// success.
func (m *Module) instShuffleVector(old *ast.InstShuffleVector, resolved, unresolved map[ast.NamedValue]value.Named) bool {
	if isUnresolved(unresolved, old.X, old.Y) {
		return false
	}
	v := m.getLocal(old.Name)
	inst, ok := v.(*ir.InstShuffleVector)
	if !ok {
		panic(fmt.Errorf("invalid instruction type for instruction %s; expected *ir.InstShuffleVector, got %T", enc.Local(old.Name), v))
	}
	inst.X = m.irValue(old.X)
	inst.Y = m.irValue(old.Y)
	inst.Mask = m.irValue(old.Mask)
	inst.Metadata = m.irMetadata(old.Metadata)
	return true
}

// --- [ Aggregate instructions ] ----------------------------------------------

// instExtractValue resolves the given extractvalue instruction, by recursively
// resolving its operands. The boolean return value indicates success.
func (m *Module) instExtractValue(old *ast.InstExtractValue, resolved, unresolved map[ast.NamedValue]value.Named) bool {
	if isUnresolved(unresolved, old.X) {
		return false
	}
	v := m.getLocal(old.Name)
	inst, ok := v.(*ir.InstExtractValue)
	if !ok {
		panic(fmt.Errorf("invalid instruction type for instruction %s; expected *ir.InstExtractValue, got %T", enc.Local(old.Name), v))
	}
	x := m.irValue(old.X)
	typ := aggregateElemType(x.Type(), old.Indices)
	inst.Typ = typ
	inst.X = x
	inst.Indices = old.Indices
	inst.Metadata = m.irMetadata(old.Metadata)
	return true
}

// instInsertValue resolves the given insertvalue instruction, by recursively
// resolving its operands. The boolean return value indicates success.
func (m *Module) instInsertValue(old *ast.InstInsertValue, resolved, unresolved map[ast.NamedValue]value.Named) bool {
	if isUnresolved(unresolved, old.X, old.Elem) {
		return false
	}
	v := m.getLocal(old.Name)
	inst, ok := v.(*ir.InstInsertValue)
	if !ok {
		panic(fmt.Errorf("invalid instruction type for instruction %s; expected *ir.InstInsertValue, got %T", enc.Local(old.Name), v))
	}
	inst.X = m.irValue(old.X)
	inst.Elem = m.irValue(old.Elem)
	inst.Indices = old.Indices
	inst.Metadata = m.irMetadata(old.Metadata)
	return true
}

// --- [ Memory instructions ] -------------------------------------------------

// instAlloca resolves the given alloca instruction, by recursively resolving
// its operands. The boolean return value indicates success.
func (m *Module) instAlloca(old *ast.InstAlloca, resolved, unresolved map[ast.NamedValue]value.Named) bool {
	// TODO: Validate if a non-nil check is needed for nelems.
	if isUnresolved(unresolved, old.NElems) {
		return false
	}
	v := m.getLocal(old.Name)
	inst, ok := v.(*ir.InstAlloca)
	if !ok {
		panic(fmt.Errorf("invalid instruction type for instruction %s; expected *ir.InstAlloca, got %T", enc.Local(old.Name), v))
	}
	elem := m.irType(old.Elem)
	typ := types.NewPointer(elem)
	inst.Typ = typ
	inst.Elem = elem
	if old.NElems != nil {
		inst.NElems = m.irValue(old.NElems)
	}
	inst.Metadata = m.irMetadata(old.Metadata)
	return true
}

// instLoad resolves the given load instruction, by recursively resolving its
// operands. The boolean return value indicates success.
func (m *Module) instLoad(old *ast.InstLoad, resolved, unresolved map[ast.NamedValue]value.Named) bool {
	if isUnresolved(unresolved, old.Src) {
		return false
	}
	v := m.getLocal(old.Name)
	inst, ok := v.(*ir.InstLoad)
	if !ok {
		panic(fmt.Errorf("invalid instruction type for instruction %s; expected *ir.InstLoad, got %T", enc.Local(old.Name), v))
	}
	src := m.irValue(old.Src)
	srcType, ok := src.Type().(*types.PointerType)
	if !ok {
		panic(fmt.Errorf("invalid source type; expected *types.PointerType, got %T", src.Type()))
	}
	typ := srcType.Elem
	if got, want := typ, m.irType(old.Elem); !got.Equal(want) {
		m.errs = append(m.errs, errors.Errorf("source element type mismatch; expected `%v`, got `%v`", want, got))
	}
	inst.Typ = typ
	inst.Src = src
	inst.Metadata = m.irMetadata(old.Metadata)
	return true
}

// instGetElementPtr resolves the given getelementptr instruction, by
// recursively resolving its operands. The boolean return value indicates
// success.
func (m *Module) instGetElementPtr(old *ast.InstGetElementPtr, resolved, unresolved map[ast.NamedValue]value.Named) bool {
	if isUnresolved(unresolved, old.Src) || isUnresolved(unresolved, old.Indices...) {
		return false
	}
	v := m.getLocal(old.Name)
	inst, ok := v.(*ir.InstGetElementPtr)
	if !ok {
		panic(fmt.Errorf("invalid instruction type for instruction %s; expected *ir.InstGetElementPtr, got %T", enc.Local(old.Name), v))
	}
	src := m.irValue(old.Src)
	srcType, ok := src.Type().(*types.PointerType)
	if !ok {
		m.errs = append(m.errs, errors.Errorf("invalid source type; expected *types.PointerType, got %T", src.Type()))
	}
	elem := srcType.Elem
	if got, want := elem, m.irType(old.Elem); !got.Equal(want) {
		m.errs = append(m.errs, errors.Errorf("source element type mismatch; expected `%v`, got `%v`", want, got))
	}
	var indices []value.Value
	for _, oldIndex := range old.Indices {
		index := m.irValue(oldIndex)
		indices = append(indices, index)
	}
	e := elem
	for i, index := range indices {
		if i == 0 {
			// Ignore checking the 0th index as it simply follows the pointer of
			// src.
			//
			// ref: http://llvm.org/docs/GetElementPtr.html#why-is-the-extra-0-index-required
			continue
		}
		switch t := e.(type) {
		case *types.PointerType:
			// ref: http://llvm.org/docs/GetElementPtr.html#what-is-dereferenced-by-gep
			panic("unable to index into element of pointer type; for more information, see http://llvm.org/docs/GetElementPtr.html#what-is-dereferenced-by-gep")
		case *types.ArrayType:
			e = t.Elem
		case *types.StructType:
			idx, ok := index.(*constant.Int)
			if !ok {
				panic(fmt.Errorf("invalid index type for structure element; expected *constant.Int, got %T", index))
			}
			e = t.Fields[idx.Int64()]
		default:
			panic(fmt.Errorf("support for indexing element type %T not yet implemented", e))
		}
	}
	typ := types.NewPointer(e)
	inst.Typ = typ
	inst.Elem = elem
	inst.Src = src
	inst.Indices = indices
	inst.Metadata = m.irMetadata(old.Metadata)
	return true
}

// --- [ Conversion instructions ] ---------------------------------------------

// instTrunc resolves the given trunc instruction, by recursively resolving its
// operands. The boolean return value indicates success.
func (m *Module) instTrunc(old *ast.InstTrunc, resolved, unresolved map[ast.NamedValue]value.Named) bool {
	if isUnresolved(unresolved, old.From) {
		return false
	}
	v := m.getLocal(old.Name)
	inst, ok := v.(*ir.InstTrunc)
	if !ok {
		panic(fmt.Errorf("invalid instruction type for instruction %s; expected *ir.InstTrunc, got %T", enc.Local(old.Name), v))
	}
	inst.From = m.irValue(old.From)
	inst.To = m.irType(old.To)
	inst.Metadata = m.irMetadata(old.Metadata)
	return true
}

// instZExt resolves the given zext instruction, by recursively resolving its
// operands. The boolean return value indicates success.
func (m *Module) instZExt(old *ast.InstZExt, resolved, unresolved map[ast.NamedValue]value.Named) bool {
	if isUnresolved(unresolved, old.From) {
		return false
	}
	v := m.getLocal(old.Name)
	inst, ok := v.(*ir.InstZExt)
	if !ok {
		panic(fmt.Errorf("invalid instruction type for instruction %s; expected *ir.InstZExt, got %T", enc.Local(old.Name), v))
	}
	inst.From = m.irValue(old.From)
	inst.To = m.irType(old.To)
	inst.Metadata = m.irMetadata(old.Metadata)
	return true
}

// instSExt resolves the given sext instruction, by recursively resolving its
// operands. The boolean return value indicates success.
func (m *Module) instSExt(old *ast.InstSExt, resolved, unresolved map[ast.NamedValue]value.Named) bool {
	if isUnresolved(unresolved, old.From) {
		return false
	}
	v := m.getLocal(old.Name)
	inst, ok := v.(*ir.InstSExt)
	if !ok {
		panic(fmt.Errorf("invalid instruction type for instruction %s; expected *ir.InstSExt, got %T", enc.Local(old.Name), v))
	}
	inst.From = m.irValue(old.From)
	inst.To = m.irType(old.To)
	inst.Metadata = m.irMetadata(old.Metadata)
	return true
}

// instFPTrunc resolves the given fptrunc instruction, by recursively resolving
// its operands. The boolean return value indicates success.
func (m *Module) instFPTrunc(old *ast.InstFPTrunc, resolved, unresolved map[ast.NamedValue]value.Named) bool {
	if isUnresolved(unresolved, old.From) {
		return false
	}
	v := m.getLocal(old.Name)
	inst, ok := v.(*ir.InstFPTrunc)
	if !ok {
		panic(fmt.Errorf("invalid instruction type for instruction %s; expected *ir.InstFPTrunc, got %T", enc.Local(old.Name), v))
	}
	inst.From = m.irValue(old.From)
	inst.To = m.irType(old.To)
	inst.Metadata = m.irMetadata(old.Metadata)
	return true
}

// instFPExt resolves the given fpext instruction, by recursively resolving its
// operands. The boolean return value indicates success.
func (m *Module) instFPExt(old *ast.InstFPExt, resolved, unresolved map[ast.NamedValue]value.Named) bool {
	if isUnresolved(unresolved, old.From) {
		return false
	}
	v := m.getLocal(old.Name)
	inst, ok := v.(*ir.InstFPExt)
	if !ok {
		panic(fmt.Errorf("invalid instruction type for instruction %s; expected *ir.InstFPExt, got %T", enc.Local(old.Name), v))
	}
	inst.From = m.irValue(old.From)
	inst.To = m.irType(old.To)
	inst.Metadata = m.irMetadata(old.Metadata)
	return true
}

// instFPToUI resolves the given fptoui instruction, by recursively resolving
// its operands. The boolean return value indicates success.
func (m *Module) instFPToUI(old *ast.InstFPToUI, resolved, unresolved map[ast.NamedValue]value.Named) bool {
	if isUnresolved(unresolved, old.From) {
		return false
	}
	v := m.getLocal(old.Name)
	inst, ok := v.(*ir.InstFPToUI)
	if !ok {
		panic(fmt.Errorf("invalid instruction type for instruction %s; expected *ir.InstFPToUI, got %T", enc.Local(old.Name), v))
	}
	inst.From = m.irValue(old.From)
	inst.To = m.irType(old.To)
	inst.Metadata = m.irMetadata(old.Metadata)
	return true
}

// instFPToSI resolves the given fptosi instruction, by recursively resolving
// its operands. The boolean return value indicates success.
func (m *Module) instFPToSI(old *ast.InstFPToSI, resolved, unresolved map[ast.NamedValue]value.Named) bool {
	if isUnresolved(unresolved, old.From) {
		return false
	}
	v := m.getLocal(old.Name)
	inst, ok := v.(*ir.InstFPToSI)
	if !ok {
		panic(fmt.Errorf("invalid instruction type for instruction %s; expected *ir.InstFPToSI, got %T", enc.Local(old.Name), v))
	}
	inst.From = m.irValue(old.From)
	inst.To = m.irType(old.To)
	inst.Metadata = m.irMetadata(old.Metadata)
	return true
}

// instUIToFP resolves the given uitofp instruction, by recursively resolving
// its operands. The boolean return value indicates success.
func (m *Module) instUIToFP(old *ast.InstUIToFP, resolved, unresolved map[ast.NamedValue]value.Named) bool {
	if isUnresolved(unresolved, old.From) {
		return false
	}
	v := m.getLocal(old.Name)
	inst, ok := v.(*ir.InstUIToFP)
	if !ok {
		panic(fmt.Errorf("invalid instruction type for instruction %s; expected *ir.InstUIToFP, got %T", enc.Local(old.Name), v))
	}
	inst.From = m.irValue(old.From)
	inst.To = m.irType(old.To)
	inst.Metadata = m.irMetadata(old.Metadata)
	return true
}

// instSIToFP resolves the given sitofp instruction, by recursively resolving
// its operands. The boolean return value indicates success.
func (m *Module) instSIToFP(old *ast.InstSIToFP, resolved, unresolved map[ast.NamedValue]value.Named) bool {
	if isUnresolved(unresolved, old.From) {
		return false
	}
	v := m.getLocal(old.Name)
	inst, ok := v.(*ir.InstSIToFP)
	if !ok {
		panic(fmt.Errorf("invalid instruction type for instruction %s; expected *ir.InstSIToFP, got %T", enc.Local(old.Name), v))
	}
	inst.From = m.irValue(old.From)
	inst.To = m.irType(old.To)
	inst.Metadata = m.irMetadata(old.Metadata)
	return true
}

// instPtrToInt resolves the given ptrtoint instruction, by recursively
// resolving its operands. The boolean return value indicates success.
func (m *Module) instPtrToInt(old *ast.InstPtrToInt, resolved, unresolved map[ast.NamedValue]value.Named) bool {
	if isUnresolved(unresolved, old.From) {
		return false
	}
	v := m.getLocal(old.Name)
	inst, ok := v.(*ir.InstPtrToInt)
	if !ok {
		panic(fmt.Errorf("invalid instruction type for instruction %s; expected *ir.InstPtrToInt, got %T", enc.Local(old.Name), v))
	}
	inst.From = m.irValue(old.From)
	inst.To = m.irType(old.To)
	inst.Metadata = m.irMetadata(old.Metadata)
	return true
}

// instIntToPtr resolves the given inttoptr instruction, by recursively
// resolving its operands. The boolean return value indicates success.
func (m *Module) instIntToPtr(old *ast.InstIntToPtr, resolved, unresolved map[ast.NamedValue]value.Named) bool {
	if isUnresolved(unresolved, old.From) {
		return false
	}
	v := m.getLocal(old.Name)
	inst, ok := v.(*ir.InstIntToPtr)
	if !ok {
		panic(fmt.Errorf("invalid instruction type for instruction %s; expected *ir.InstIntToPtr, got %T", enc.Local(old.Name), v))
	}
	inst.From = m.irValue(old.From)
	inst.To = m.irType(old.To)
	inst.Metadata = m.irMetadata(old.Metadata)
	return true
}

// instBitCast resolves the given bitcast instruction, by recursively resolving
// its operands. The boolean return value indicates success.
func (m *Module) instBitCast(old *ast.InstBitCast, resolved, unresolved map[ast.NamedValue]value.Named) bool {
	if isUnresolved(unresolved, old.From) {
		return false
	}
	v := m.getLocal(old.Name)
	inst, ok := v.(*ir.InstBitCast)
	if !ok {
		panic(fmt.Errorf("invalid instruction type for instruction %s; expected *ir.InstBitCast, got %T", enc.Local(old.Name), v))
	}
	inst.From = m.irValue(old.From)
	inst.To = m.irType(old.To)
	inst.Metadata = m.irMetadata(old.Metadata)
	return true
}

// instAddrSpaceCast resolves the given addrspacecast instruction, by
// recursively resolving its operands. The boolean return value indicates
// success.
func (m *Module) instAddrSpaceCast(old *ast.InstAddrSpaceCast, resolved, unresolved map[ast.NamedValue]value.Named) bool {
	if isUnresolved(unresolved, old.From) {
		return false
	}
	v := m.getLocal(old.Name)
	inst, ok := v.(*ir.InstAddrSpaceCast)
	if !ok {
		panic(fmt.Errorf("invalid instruction type for instruction %s; expected *ir.InstAddrSpaceCast, got %T", enc.Local(old.Name), v))
	}
	inst.From = m.irValue(old.From)
	inst.To = m.irType(old.To)
	inst.Metadata = m.irMetadata(old.Metadata)
	return true
}

// --- [ Other instructions ] --------------------------------------------------

// instICmp resolves the given icmp instruction, by recursively resolving its
// operands. The boolean return value indicates success.
func (m *Module) instICmp(old *ast.InstICmp, resolved, unresolved map[ast.NamedValue]value.Named) bool {
	if isUnresolved(unresolved, old.X, old.Y) {
		return false
	}
	v := m.getLocal(old.Name)
	inst, ok := v.(*ir.InstICmp)
	if !ok {
		panic(fmt.Errorf("invalid instruction type for instruction %s; expected *ir.InstICmp, got %T", enc.Local(old.Name), v))
	}
	pred := irIntPred(old.Pred)
	x := m.irValue(old.X)
	y := m.irValue(old.Y)
	var typ types.Type = types.I1
	if t, ok := x.Type().(*types.VectorType); ok {
		typ = types.NewVector(types.I1, t.Len)
	}
	inst.Typ = typ
	inst.Pred = pred
	inst.X = x
	inst.Y = y
	inst.Metadata = m.irMetadata(old.Metadata)
	return true
}

// instFCmp resolves the given fcmp instruction, by recursively resolving its
// operands. The boolean return value indicates success.
func (m *Module) instFCmp(old *ast.InstFCmp, resolved, unresolved map[ast.NamedValue]value.Named) bool {
	if isUnresolved(unresolved, old.X, old.Y) {
		return false
	}
	v := m.getLocal(old.Name)
	inst, ok := v.(*ir.InstFCmp)
	if !ok {
		panic(fmt.Errorf("invalid instruction type for instruction %s; expected *ir.InstFCmp, got %T", enc.Local(old.Name), v))
	}
	pred := irFloatPred(old.Pred)
	x := m.irValue(old.X)
	y := m.irValue(old.Y)
	var typ types.Type = types.I1
	if t, ok := x.Type().(*types.VectorType); ok {
		typ = types.NewVector(types.I1, t.Len)
	}
	inst.Typ = typ
	inst.Pred = pred
	inst.X = x
	inst.Y = y
	inst.Metadata = m.irMetadata(old.Metadata)
	return true
}

// instPhi resolves the given phi instruction, by recursively resolving its
// operands. The boolean return value indicates success.
func (m *Module) instPhi(old *ast.InstPhi, resolved, unresolved map[ast.NamedValue]value.Named) bool {
	// Note, to prevent infinite loops in the type-checker, PHI-instructions are
	// not checked to see if they contain unresolved values. This is fine as the
	// PHI type is computed from old.Typ anyways, so it does not need to be
	// inferred from the operands of the PHI-instruction.
	v := m.getLocal(old.Name)
	inst, ok := v.(*ir.InstPhi)
	if !ok {
		panic(fmt.Errorf("invalid instruction type for instruction %s; expected *ir.InstPhi, got %T", enc.Local(old.Name), v))
	}
	inst.Typ = m.irType(old.Type)
	for _, oldInc := range old.Incs {
		x := m.irValue(oldInc.X)
		v := m.getLocal(oldInc.Pred.GetName())
		pred, ok := v.(*ir.BasicBlock)
		if !ok {
			panic(fmt.Errorf("invalid basic block type; expected *ir.BasicBlock, got %T", v))
		}
		inc := &ir.Incoming{
			X:    x,
			Pred: pred,
		}
		inst.Incs = append(inst.Incs, inc)
	}
	inst.Metadata = m.irMetadata(old.Metadata)
	return true
}

// instSelect resolves the given select instruction, by recursively resolving
// its operands. The boolean return value indicates success.
func (m *Module) instSelect(old *ast.InstSelect, resolved, unresolved map[ast.NamedValue]value.Named) bool {
	if isUnresolved(unresolved, old.Cond, old.X, old.Y) {
		return false
	}
	v := m.getLocal(old.Name)
	inst, ok := v.(*ir.InstSelect)
	if !ok {
		panic(fmt.Errorf("invalid instruction type for instruction %s; expected *ir.InstSelect, got %T", enc.Local(old.Name), v))
	}
	inst.Cond = m.irValue(old.Cond)
	inst.X = m.irValue(old.X)
	inst.Y = m.irValue(old.Y)
	inst.Metadata = m.irMetadata(old.Metadata)
	return true
}

// instCall resolves the given call instruction, by recursively resolving its
// operands. The boolean return value indicates success.
func (m *Module) instCall(old *ast.InstCall, resolved, unresolved map[ast.NamedValue]value.Named) bool {
	if isUnresolved(unresolved, old.Callee) || isUnresolved(unresolved, old.Args...) {
		return false
	}
	v := m.getLocal(old.Name)
	inst, ok := v.(*ir.InstCall)
	if !ok {
		panic(fmt.Errorf("invalid instruction type for instruction %s; expected *ir.InstCall, got %T", enc.Local(old.Name), v))
	}
	callee := m.irValue(old.Callee)
	typ, ok := callee.Type().(*types.PointerType)
	if !ok {
		panic(fmt.Errorf("invalid callee type, expected *types.PointerType, got %T", callee.Type()))
	}
	sig, ok := typ.Elem.(*types.FuncType)
	if !ok {
		panic(fmt.Errorf("invalid callee signature type, expected *types.FuncType, got %T", typ.Elem))
	}
	inst.Callee = callee
	inst.Sig = sig
	// TODO: Validate old.Type against inst.Sig.
	for _, oldArg := range old.Args {
		arg := m.irValue(oldArg)
		inst.Args = append(inst.Args, arg)
	}
	inst.CallConv = ir.CallConv(old.CallConv)
	inst.Metadata = m.irMetadata(old.Metadata)
	return true
}

// ### [ Helper functions ] ####################################################

// isUnresolved reports whether the given value is unresolved.
func isUnresolved(unresolved map[ast.NamedValue]value.Named, vs ...ast.Value) bool {
	for _, v := range vs {
		if v, ok := v.(ast.NamedValue); ok {
			if _, ok := unresolved[v]; ok {
				return true
			}
		}
	}
	return false
}
