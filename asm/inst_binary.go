package asm

import (
	"fmt"

	"github.com/llir/ll/ast"
	"github.com/llir/llvm/ir"
	"github.com/pkg/errors"
)

// --- [ Binary instructions ] -------------------------------------------------

// ~~~ [ add ] ~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~

// astToIRInstAdd translates the given AST add instruction into an equivalent IR
// instruction.
func (fgen *funcGen) astToIRInstAdd(inst ir.Instruction, old *ast.AddInst) (*ir.InstAdd, error) {
	i, ok := inst.(*ir.InstAdd)
	if !ok {
		panic(fmt.Errorf("invalid IR instruction for AST instruction; expected *ir.InstAdd, got %T", inst))
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

// ~~~ [ fadd ] ~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~

// astToIRInstFAdd translates the given AST fadd instruction into an equivalent
// IR instruction.
func (fgen *funcGen) astToIRInstFAdd(inst ir.Instruction, old *ast.FAddInst) (*ir.InstFAdd, error) {
	i, ok := inst.(*ir.InstFAdd)
	if !ok {
		panic(fmt.Errorf("invalid IR instruction for AST instruction; expected *ir.InstFAdd, got %T", inst))
	}
	// (optional) Fast math flags.
	i.FastMathFlags = irFastMathFlags(old.FastMathFlags())
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

// ~~~ [ sub ] ~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~

// astToIRInstSub translates the given AST sub instruction into an equivalent IR
// instruction.
func (fgen *funcGen) astToIRInstSub(inst ir.Instruction, old *ast.SubInst) (*ir.InstSub, error) {
	i, ok := inst.(*ir.InstSub)
	if !ok {
		panic(fmt.Errorf("invalid IR instruction for AST instruction; expected *ir.InstSub, got %T", inst))
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

// ~~~ [ fsub ] ~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~

// astToIRInstFSub translates the given AST fsub instruction into an equivalent
// IR instruction.
func (fgen *funcGen) astToIRInstFSub(inst ir.Instruction, old *ast.FSubInst) (*ir.InstFSub, error) {
	i, ok := inst.(*ir.InstFSub)
	if !ok {
		panic(fmt.Errorf("invalid IR instruction for AST instruction; expected *ir.InstFSub, got %T", inst))
	}
	// (optional) Fast math flags.
	i.FastMathFlags = irFastMathFlags(old.FastMathFlags())
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

// ~~~ [ mul ] ~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~

// astToIRInstMul translates the given AST mul instruction into an equivalent IR
// instruction.
func (fgen *funcGen) astToIRInstMul(inst ir.Instruction, old *ast.MulInst) (*ir.InstMul, error) {
	i, ok := inst.(*ir.InstMul)
	if !ok {
		panic(fmt.Errorf("invalid IR instruction for AST instruction; expected *ir.InstMul, got %T", inst))
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

// ~~~ [ fmul ] ~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~

// astToIRInstFMul translates the given AST fmul instruction into an equivalent
// IR instruction.
func (fgen *funcGen) astToIRInstFMul(inst ir.Instruction, old *ast.FMulInst) (*ir.InstFMul, error) {
	i, ok := inst.(*ir.InstFMul)
	if !ok {
		panic(fmt.Errorf("invalid IR instruction for AST instruction; expected *ir.InstFMul, got %T", inst))
	}
	// (optional) Fast math flags.
	i.FastMathFlags = irFastMathFlags(old.FastMathFlags())
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

// ~~~ [ udiv ] ~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~

// astToIRInstUDiv translates the given AST udiv instruction into an equivalent
// IR instruction.
func (fgen *funcGen) astToIRInstUDiv(inst ir.Instruction, old *ast.UDivInst) (*ir.InstUDiv, error) {
	i, ok := inst.(*ir.InstUDiv)
	if !ok {
		panic(fmt.Errorf("invalid IR instruction for AST instruction; expected *ir.InstUDiv, got %T", inst))
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

// ~~~ [ sdiv ] ~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~

// astToIRInstSDiv translates the given AST sdiv instruction into an equivalent
// IR instruction.
func (fgen *funcGen) astToIRInstSDiv(inst ir.Instruction, old *ast.SDivInst) (*ir.InstSDiv, error) {
	i, ok := inst.(*ir.InstSDiv)
	if !ok {
		panic(fmt.Errorf("invalid IR instruction for AST instruction; expected *ir.InstSDiv, got %T", inst))
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

// ~~~ [ fdiv ] ~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~

// astToIRInstFDiv translates the given AST fdiv instruction into an equivalent
// IR instruction.
func (fgen *funcGen) astToIRInstFDiv(inst ir.Instruction, old *ast.FDivInst) (*ir.InstFDiv, error) {
	i, ok := inst.(*ir.InstFDiv)
	if !ok {
		panic(fmt.Errorf("invalid IR instruction for AST instruction; expected *ir.InstFDiv, got %T", inst))
	}
	// (optional) Fast math flags.
	i.FastMathFlags = irFastMathFlags(old.FastMathFlags())
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

// ~~~ [ urem ] ~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~

// astToIRInstURem translates the given AST urem instruction into an equivalent
// IR instruction.
func (fgen *funcGen) astToIRInstURem(inst ir.Instruction, old *ast.URemInst) (*ir.InstURem, error) {
	i, ok := inst.(*ir.InstURem)
	if !ok {
		panic(fmt.Errorf("invalid IR instruction for AST instruction; expected *ir.InstURem, got %T", inst))
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

// ~~~ [ srem ] ~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~

// astToIRInstSRem translates the given AST srem instruction into an equivalent
// IR instruction.
func (fgen *funcGen) astToIRInstSRem(inst ir.Instruction, old *ast.SRemInst) (*ir.InstSRem, error) {
	i, ok := inst.(*ir.InstSRem)
	if !ok {
		panic(fmt.Errorf("invalid IR instruction for AST instruction; expected *ir.InstSRem, got %T", inst))
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

// ~~~ [ frem ] ~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~

// astToIRInstFRem translates the given AST frem instruction into an equivalent
// IR instruction.
func (fgen *funcGen) astToIRInstFRem(inst ir.Instruction, old *ast.FRemInst) (*ir.InstFRem, error) {
	i, ok := inst.(*ir.InstFRem)
	if !ok {
		panic(fmt.Errorf("invalid IR instruction for AST instruction; expected *ir.InstFRem, got %T", inst))
	}
	// (optional) Fast math flags.
	i.FastMathFlags = irFastMathFlags(old.FastMathFlags())
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
