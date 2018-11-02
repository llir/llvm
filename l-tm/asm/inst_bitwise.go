package asm

import (
	"fmt"

	"github.com/llir/l/ir"
	"github.com/mewmew/l-tm/asm/ll/ast"
	"github.com/pkg/errors"
)

// --- [ Bitwise instructions ] ------------------------------------------------

// ~~~ [ shl ] ~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~

func (fgen *funcGen) astToIRInstShl(inst ir.Instruction, old *ast.ShlInst) (*ir.InstShl, error) {
	i, ok := inst.(*ir.InstShl)
	if !ok {
		// NOTE: panic since this would indicate a bug in the implementation.
		panic(fmt.Errorf("invalid IR instruction for AST instruction; expected *ir.InstShl, got %T", inst))
	}
	// Overflow flags.
	i.OverflowFlags = irOverflowFlags(old.OverflowFlags())
	// X operand.
	xType, err := fgen.gen.irType(old.X().Typ())
	if err != nil {
		return nil, errors.WithStack(err)
	}
	x, err := fgen.astToIRTypeValue(old.X())
	if err != nil {
		return nil, errors.WithStack(err)
	}
	i.X = x
	// Y operand.
	y, err := fgen.astToIRValue(xType, old.Y())
	if err != nil {
		return nil, errors.WithStack(err)
	}
	i.Y = y
	return i, nil
}

// ~~~ [ lshr ] ~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~

func (fgen *funcGen) astToIRInstLShr(inst ir.Instruction, old *ast.LShrInst) (*ir.InstLShr, error) {
	i, ok := inst.(*ir.InstLShr)
	if !ok {
		// NOTE: panic since this would indicate a bug in the implementation.
		panic(fmt.Errorf("invalid IR instruction for AST instruction; expected *ir.InstLShr, got %T", inst))
	}
	// X operand.
	xType, err := fgen.gen.irType(old.X().Typ())
	if err != nil {
		return nil, errors.WithStack(err)
	}
	x, err := fgen.astToIRTypeValue(old.X())
	if err != nil {
		return nil, errors.WithStack(err)
	}
	i.X = x
	// Y operand.
	y, err := fgen.astToIRValue(xType, old.Y())
	if err != nil {
		return nil, errors.WithStack(err)
	}
	i.Y = y
	return i, nil
}

// ~~~ [ ashr ] ~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~

func (fgen *funcGen) astToIRInstAShr(inst ir.Instruction, old *ast.AShrInst) (*ir.InstAShr, error) {
	i, ok := inst.(*ir.InstAShr)
	if !ok {
		// NOTE: panic since this would indicate a bug in the implementation.
		panic(fmt.Errorf("invalid IR instruction for AST instruction; expected *ir.InstAShr, got %T", inst))
	}
	// X operand.
	xType, err := fgen.gen.irType(old.X().Typ())
	if err != nil {
		return nil, errors.WithStack(err)
	}
	x, err := fgen.astToIRTypeValue(old.X())
	if err != nil {
		return nil, errors.WithStack(err)
	}
	i.X = x
	// Y operand.
	y, err := fgen.astToIRValue(xType, old.Y())
	if err != nil {
		return nil, errors.WithStack(err)
	}
	i.Y = y
	return i, nil
}

// ~~~ [ and ] ~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~

func (fgen *funcGen) astToIRInstAnd(inst ir.Instruction, old *ast.AndInst) (*ir.InstAnd, error) {
	i, ok := inst.(*ir.InstAnd)
	if !ok {
		// NOTE: panic since this would indicate a bug in the implementation.
		panic(fmt.Errorf("invalid IR instruction for AST instruction; expected *ir.InstAnd, got %T", inst))
	}
	// X operand.
	xType, err := fgen.gen.irType(old.X().Typ())
	if err != nil {
		return nil, errors.WithStack(err)
	}
	x, err := fgen.astToIRTypeValue(old.X())
	if err != nil {
		return nil, errors.WithStack(err)
	}
	i.X = x
	// Y operand.
	y, err := fgen.astToIRValue(xType, old.Y())
	if err != nil {
		return nil, errors.WithStack(err)
	}
	i.Y = y
	return i, nil
}

// ~~~ [ or ] ~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~

func (fgen *funcGen) astToIRInstOr(inst ir.Instruction, old *ast.OrInst) (*ir.InstOr, error) {
	i, ok := inst.(*ir.InstOr)
	if !ok {
		// NOTE: panic since this would indicate a bug in the implementation.
		panic(fmt.Errorf("invalid IR instruction for AST instruction; expected *ir.InstOr, got %T", inst))
	}
	// X operand.
	xType, err := fgen.gen.irType(old.X().Typ())
	if err != nil {
		return nil, errors.WithStack(err)
	}
	x, err := fgen.astToIRTypeValue(old.X())
	if err != nil {
		return nil, errors.WithStack(err)
	}
	i.X = x
	// Y operand.
	y, err := fgen.astToIRValue(xType, old.Y())
	if err != nil {
		return nil, errors.WithStack(err)
	}
	i.Y = y
	return i, nil
}

// ~~~ [ xor ] ~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~

func (fgen *funcGen) astToIRInstXor(inst ir.Instruction, old *ast.XorInst) (*ir.InstXor, error) {
	i, ok := inst.(*ir.InstXor)
	if !ok {
		// NOTE: panic since this would indicate a bug in the implementation.
		panic(fmt.Errorf("invalid IR instruction for AST instruction; expected *ir.InstXor, got %T", inst))
	}
	// X operand.
	xType, err := fgen.gen.irType(old.X().Typ())
	if err != nil {
		return nil, errors.WithStack(err)
	}
	x, err := fgen.astToIRTypeValue(old.X())
	if err != nil {
		return nil, errors.WithStack(err)
	}
	i.X = x
	// Y operand.
	y, err := fgen.astToIRValue(xType, old.Y())
	if err != nil {
		return nil, errors.WithStack(err)
	}
	i.Y = y
	return i, nil
}
