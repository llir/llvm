package asm

import (
	"fmt"

	"github.com/llir/ll/ast"
	"github.com/llir/llvm/ir"
	"github.com/pkg/errors"
)

// --- [ Bitwise instructions ] ------------------------------------------------

// ~~~ [ shl ] ~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~

// astToIRInstShl translates the given AST shl instruction into an equivalent IR
// instruction.
func (fgen *funcGen) astToIRInstShl(inst ir.Instruction, old *ast.ShlInst) (*ir.InstShl, error) {
	i, ok := inst.(*ir.InstShl)
	if !ok {
		panic(fmt.Errorf("invalid IR instruction for AST instruction; expected *ir.InstShl, got %T", inst))
	}
	// (optional) Overflow flags.
	i.OverflowFlags = irOverflowFlags(old.OverflowFlags())
	// X operand.
	x, err := fgen.astToIRTypeValue(old.X())
	if err != nil {
		return nil, errors.WithStack(err)
	}
	i.X = x
	// Y operand.
	y, err := fgen.astToIRValue(x.Type(), old.Y())
	if err != nil {
		return nil, errors.WithStack(err)
	}
	i.Y = y
	// (optional) Metadata.
	md, err := fgen.gen.irMetadataAttachments(old.Metadata())
	if err != nil {
		return nil, errors.WithStack(err)
	}
	i.Metadata = md
	return i, nil
}

// ~~~ [ lshr ] ~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~

// astToIRInstLShr translates the given AST lshr instruction into an equivalent
// IR instruction.
func (fgen *funcGen) astToIRInstLShr(inst ir.Instruction, old *ast.LShrInst) (*ir.InstLShr, error) {
	i, ok := inst.(*ir.InstLShr)
	if !ok {
		panic(fmt.Errorf("invalid IR instruction for AST instruction; expected *ir.InstLShr, got %T", inst))
	}
	// (optional) Exact.
	_, exact := old.Exact()
	i.Exact = exact
	// X operand.
	x, err := fgen.astToIRTypeValue(old.X())
	if err != nil {
		return nil, errors.WithStack(err)
	}
	i.X = x
	// Y operand.
	y, err := fgen.astToIRValue(x.Type(), old.Y())
	if err != nil {
		return nil, errors.WithStack(err)
	}
	i.Y = y
	// (optional) Metadata.
	md, err := fgen.gen.irMetadataAttachments(old.Metadata())
	if err != nil {
		return nil, errors.WithStack(err)
	}
	i.Metadata = md
	return i, nil
}

// ~~~ [ ashr ] ~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~

// astToIRInstAShr translates the given AST ashr instruction into an equivalent
// IR instruction.
func (fgen *funcGen) astToIRInstAShr(inst ir.Instruction, old *ast.AShrInst) (*ir.InstAShr, error) {
	i, ok := inst.(*ir.InstAShr)
	if !ok {
		panic(fmt.Errorf("invalid IR instruction for AST instruction; expected *ir.InstAShr, got %T", inst))
	}
	// (optional) Exact.
	_, exact := old.Exact()
	i.Exact = exact
	// X operand.
	x, err := fgen.astToIRTypeValue(old.X())
	if err != nil {
		return nil, errors.WithStack(err)
	}
	i.X = x
	// Y operand.
	y, err := fgen.astToIRValue(x.Type(), old.Y())
	if err != nil {
		return nil, errors.WithStack(err)
	}
	i.Y = y
	// (optional) Metadata.
	md, err := fgen.gen.irMetadataAttachments(old.Metadata())
	if err != nil {
		return nil, errors.WithStack(err)
	}
	i.Metadata = md
	return i, nil
}

// ~~~ [ and ] ~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~

// astToIRInstAnd translates the given AST and instruction into an equivalent IR
// instruction.
func (fgen *funcGen) astToIRInstAnd(inst ir.Instruction, old *ast.AndInst) (*ir.InstAnd, error) {
	i, ok := inst.(*ir.InstAnd)
	if !ok {
		panic(fmt.Errorf("invalid IR instruction for AST instruction; expected *ir.InstAnd, got %T", inst))
	}
	// X operand.
	x, err := fgen.astToIRTypeValue(old.X())
	if err != nil {
		return nil, errors.WithStack(err)
	}
	i.X = x
	// Y operand.
	y, err := fgen.astToIRValue(x.Type(), old.Y())
	if err != nil {
		return nil, errors.WithStack(err)
	}
	i.Y = y
	// (optional) Metadata.
	md, err := fgen.gen.irMetadataAttachments(old.Metadata())
	if err != nil {
		return nil, errors.WithStack(err)
	}
	i.Metadata = md
	return i, nil
}

// ~~~ [ or ] ~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~

// astToIRInstOr translates the given AST or instruction into an equivalent IR
// instruction.
func (fgen *funcGen) astToIRInstOr(inst ir.Instruction, old *ast.OrInst) (*ir.InstOr, error) {
	i, ok := inst.(*ir.InstOr)
	if !ok {
		panic(fmt.Errorf("invalid IR instruction for AST instruction; expected *ir.InstOr, got %T", inst))
	}
	// X operand.
	x, err := fgen.astToIRTypeValue(old.X())
	if err != nil {
		return nil, errors.WithStack(err)
	}
	i.X = x
	// Y operand.
	y, err := fgen.astToIRValue(x.Type(), old.Y())
	if err != nil {
		return nil, errors.WithStack(err)
	}
	i.Y = y
	// (optional) Metadata.
	md, err := fgen.gen.irMetadataAttachments(old.Metadata())
	if err != nil {
		return nil, errors.WithStack(err)
	}
	i.Metadata = md
	return i, nil
}

// ~~~ [ xor ] ~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~

// astToIRInstXor translates the given AST xor instruction into an equivalent IR
// instruction.
func (fgen *funcGen) astToIRInstXor(inst ir.Instruction, old *ast.XorInst) (*ir.InstXor, error) {
	i, ok := inst.(*ir.InstXor)
	if !ok {
		panic(fmt.Errorf("invalid IR instruction for AST instruction; expected *ir.InstXor, got %T", inst))
	}
	// X operand.
	x, err := fgen.astToIRTypeValue(old.X())
	if err != nil {
		return nil, errors.WithStack(err)
	}
	i.X = x
	// Y operand.
	y, err := fgen.astToIRValue(x.Type(), old.Y())
	if err != nil {
		return nil, errors.WithStack(err)
	}
	i.Y = y
	// (optional) Metadata.
	md, err := fgen.gen.irMetadataAttachments(old.Metadata())
	if err != nil {
		return nil, errors.WithStack(err)
	}
	i.Metadata = md
	return i, nil
}
