package asm

import (
	"fmt"

	"github.com/llir/ll/ast"
	"github.com/llir/llvm/ir"
	"github.com/llir/llvm/ir/types"
	"github.com/pkg/errors"
)

// === [ Create IR ] ===========================================================

// newExtractElementInst returns a new IR extractelement instruction (without
// body but with type) based on the given AST extractelement instruction.
func (fgen *funcGen) newExtractElementInst(ident ir.LocalIdent, old *ast.ExtractElementInst) (*ir.InstExtractElement, error) {
	xType, err := fgen.gen.irType(old.X().Typ())
	if err != nil {
		return nil, errors.WithStack(err)
	}
	xt, ok := xType.(*types.VectorType)
	if !ok {
		panic(fmt.Errorf("invalid vector type; expected *types.VectorType, got %T", xType))
	}
	return &ir.InstExtractElement{LocalIdent: ident, Typ: xt.ElemType}, nil
}

// newInsertElementInst returns a new IR insertelement instruction (without body
// but with type) based on the given AST insertelement instruction.
func (fgen *funcGen) newInsertElementInst(ident ir.LocalIdent, old *ast.InsertElementInst) (*ir.InstInsertElement, error) {
	xType, err := fgen.gen.irType(old.X().Typ())
	if err != nil {
		return nil, errors.WithStack(err)
	}
	xt, ok := xType.(*types.VectorType)
	if !ok {
		panic(fmt.Errorf("invalid vector type; expected *types.VectorType, got %T", xType))
	}
	return &ir.InstInsertElement{LocalIdent: ident, Typ: xt}, nil
}

// newShuffleVectorInst returns a new IR shufflevector instruction (without body
// but with type) based on the given AST shufflevector instruction.
func (fgen *funcGen) newShuffleVectorInst(ident ir.LocalIdent, old *ast.ShuffleVectorInst) (*ir.InstShuffleVector, error) {
	xType, err := fgen.gen.irType(old.X().Typ())
	if err != nil {
		return nil, errors.WithStack(err)
	}
	xt, ok := xType.(*types.VectorType)
	if !ok {
		panic(fmt.Errorf("invalid vector type; expected *types.VectorType, got %T", xType))
	}
	maskType, err := fgen.gen.irType(old.Mask().Typ())
	if err != nil {
		return nil, errors.WithStack(err)
	}
	mt, ok := maskType.(*types.VectorType)
	if !ok {
		panic(fmt.Errorf("invalid vector type; expected *types.VectorType, got %T", maskType))
	}
	typ := types.NewVector(mt.Len, xt.ElemType)
	return &ir.InstShuffleVector{LocalIdent: ident, Typ: typ}, nil
}

// === [ Translate AST to IR ] =================================================

// --- [ extractelement ] ------------------------------------------------------

// irExtractElementInst translates the given AST extractelement instruction into
// an equivalent IR instruction.
func (fgen *funcGen) irExtractElementInst(new ir.Instruction, old *ast.ExtractElementInst) error {
	inst, ok := new.(*ir.InstExtractElement)
	if !ok {
		panic(fmt.Errorf("invalid IR instruction for AST instruction; expected *ir.InstExtractElement, got %T", new))
	}
	// Vector.
	x, err := fgen.irTypeValue(old.X())
	if err != nil {
		return errors.WithStack(err)
	}
	inst.X = x
	// Element index.
	index, err := fgen.irTypeValue(old.Index())
	if err != nil {
		return errors.WithStack(err)
	}
	inst.Index = index
	// (optional) Metadata.
	md, err := fgen.gen.irMetadataAttachments(old.Metadata())
	if err != nil {
		return errors.WithStack(err)
	}
	inst.Metadata = md
	return nil
}

// --- [ insertelement ] -------------------------------------------------------

// irInsertElementInst translates the given AST insertelement instruction into
// an equivalent IR instruction.
func (fgen *funcGen) irInsertElementInst(new ir.Instruction, old *ast.InsertElementInst) error {
	inst, ok := new.(*ir.InstInsertElement)
	if !ok {
		panic(fmt.Errorf("invalid IR instruction for AST instruction; expected *ir.InstInsertElement, got %T", new))
	}
	// Vector.
	x, err := fgen.irTypeValue(old.X())
	if err != nil {
		return errors.WithStack(err)
	}
	inst.X = x
	// Element to insert.
	elem, err := fgen.irTypeValue(old.Elem())
	if err != nil {
		return errors.WithStack(err)
	}
	inst.Elem = elem
	// Element index.
	index, err := fgen.irTypeValue(old.Index())
	if err != nil {
		return errors.WithStack(err)
	}
	inst.Index = index
	// (optional) Metadata.
	md, err := fgen.gen.irMetadataAttachments(old.Metadata())
	if err != nil {
		return errors.WithStack(err)
	}
	inst.Metadata = md
	return nil
}

// --- [ shufflevector ] -------------------------------------------------------

// irShuffleVectorInst translates the given AST shufflevector instruction into
// an equivalent IR instruction.
func (fgen *funcGen) irShuffleVectorInst(new ir.Instruction, old *ast.ShuffleVectorInst) error {
	inst, ok := new.(*ir.InstShuffleVector)
	if !ok {
		panic(fmt.Errorf("invalid IR instruction for AST instruction; expected *ir.InstShuffleVector, got %T", new))
	}
	// X vector.
	x, err := fgen.irTypeValue(old.X())
	if err != nil {
		return errors.WithStack(err)
	}
	inst.X = x
	// Y vector.
	y, err := fgen.irTypeValue(old.Y())
	if err != nil {
		return errors.WithStack(err)
	}
	inst.Y = y
	// Shuffle mask.
	mask, err := fgen.irTypeValue(old.Mask())
	if err != nil {
		return errors.WithStack(err)
	}
	inst.Mask = mask
	// (optional) Metadata.
	md, err := fgen.gen.irMetadataAttachments(old.Metadata())
	if err != nil {
		return errors.WithStack(err)
	}
	inst.Metadata = md
	return nil
}
