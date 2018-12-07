package asm

import (
	"fmt"

	"github.com/llir/ll/ast"
	"github.com/llir/llvm/ir"
	"github.com/pkg/errors"
)

// === [ Create IR ] ===========================================================

// newAddInst returns a new IR add instruction (without body but with type)
// based on the given AST add instruction.
func (fgen *funcGen) newAddInst(ident ir.LocalIdent, old *ast.AddInst) (*ir.InstAdd, error) {
	typ, err := fgen.gen.irType(old.X().Typ())
	if err != nil {
		return nil, errors.WithStack(err)
	}
	return &ir.InstAdd{LocalIdent: ident, Typ: typ}, nil
}

// newFAddInst returns a new IR fadd instruction (without body but with type)
// based on the given AST fadd instruction.
func (fgen *funcGen) newFAddInst(ident ir.LocalIdent, old *ast.FAddInst) (*ir.InstFAdd, error) {
	typ, err := fgen.gen.irType(old.X().Typ())
	if err != nil {
		return nil, errors.WithStack(err)
	}
	return &ir.InstFAdd{LocalIdent: ident, Typ: typ}, nil
}

// newSubInst returns a new IR sub instruction (without body but with type)
// based on the given AST sub instruction.
func (fgen *funcGen) newSubInst(ident ir.LocalIdent, old *ast.SubInst) (*ir.InstSub, error) {
	typ, err := fgen.gen.irType(old.X().Typ())
	if err != nil {
		return nil, errors.WithStack(err)
	}
	return &ir.InstSub{LocalIdent: ident, Typ: typ}, nil
}

// newFSubInst returns a new IR fsub instruction (without body but with type)
// based on the given AST fsub instruction.
func (fgen *funcGen) newFSubInst(ident ir.LocalIdent, old *ast.FSubInst) (*ir.InstFSub, error) {
	typ, err := fgen.gen.irType(old.X().Typ())
	if err != nil {
		return nil, errors.WithStack(err)
	}
	return &ir.InstFSub{LocalIdent: ident, Typ: typ}, nil
}

// newMulInst returns a new IR mul instruction (without body but with type)
// based on the given AST mul instruction.
func (fgen *funcGen) newMulInst(ident ir.LocalIdent, old *ast.MulInst) (*ir.InstMul, error) {
	typ, err := fgen.gen.irType(old.X().Typ())
	if err != nil {
		return nil, errors.WithStack(err)
	}
	return &ir.InstMul{LocalIdent: ident, Typ: typ}, nil
}

// newFMulInst returns a new IR fmul instruction (without body but with type)
// based on the given AST fmul instruction.
func (fgen *funcGen) newFMulInst(ident ir.LocalIdent, old *ast.FMulInst) (*ir.InstFMul, error) {
	typ, err := fgen.gen.irType(old.X().Typ())
	if err != nil {
		return nil, errors.WithStack(err)
	}
	return &ir.InstFMul{LocalIdent: ident, Typ: typ}, nil
}

// newUDivInst returns a new IR udiv instruction (without body but with type)
// based on the given AST udiv instruction.
func (fgen *funcGen) newUDivInst(ident ir.LocalIdent, old *ast.UDivInst) (*ir.InstUDiv, error) {
	typ, err := fgen.gen.irType(old.X().Typ())
	if err != nil {
		return nil, errors.WithStack(err)
	}
	return &ir.InstUDiv{LocalIdent: ident, Typ: typ}, nil
}

// newSDivInst returns a new IR sdiv instruction (without body but with type)
// based on the given AST sdiv instruction.
func (fgen *funcGen) newSDivInst(ident ir.LocalIdent, old *ast.SDivInst) (*ir.InstSDiv, error) {
	typ, err := fgen.gen.irType(old.X().Typ())
	if err != nil {
		return nil, errors.WithStack(err)
	}
	return &ir.InstSDiv{LocalIdent: ident, Typ: typ}, nil
}

// newFDivInst returns a new IR fdiv instruction (without body but with type)
// based on the given AST fdiv instruction.
func (fgen *funcGen) newFDivInst(ident ir.LocalIdent, old *ast.FDivInst) (*ir.InstFDiv, error) {
	typ, err := fgen.gen.irType(old.X().Typ())
	if err != nil {
		return nil, errors.WithStack(err)
	}
	return &ir.InstFDiv{LocalIdent: ident, Typ: typ}, nil
}

// newURemInst returns a new IR urem instruction (without body but with type)
// based on the given AST urem instruction.
func (fgen *funcGen) newURemInst(ident ir.LocalIdent, old *ast.URemInst) (*ir.InstURem, error) {
	typ, err := fgen.gen.irType(old.X().Typ())
	if err != nil {
		return nil, errors.WithStack(err)
	}
	return &ir.InstURem{LocalIdent: ident, Typ: typ}, nil
}

// newSRemInst returns a new IR srem instruction (without body but with type)
// based on the given AST srem instruction.
func (fgen *funcGen) newSRemInst(ident ir.LocalIdent, old *ast.SRemInst) (*ir.InstSRem, error) {
	typ, err := fgen.gen.irType(old.X().Typ())
	if err != nil {
		return nil, errors.WithStack(err)
	}
	return &ir.InstSRem{LocalIdent: ident, Typ: typ}, nil
}

// newFRemInst returns a new IR frem instruction (without body but with type)
// based on the given AST frem instruction.
func (fgen *funcGen) newFRemInst(ident ir.LocalIdent, old *ast.FRemInst) (*ir.InstFRem, error) {
	typ, err := fgen.gen.irType(old.X().Typ())
	if err != nil {
		return nil, errors.WithStack(err)
	}
	return &ir.InstFRem{LocalIdent: ident, Typ: typ}, nil
}

// === [ Translate AST to IR ] =================================================

// --- [ add ] -----------------------------------------------------------------

// irAddInst translates the given AST add instruction into an equivalent IR
// instruction.
func (fgen *funcGen) irAddInst(new ir.Instruction, old *ast.AddInst) error {
	inst, ok := new.(*ir.InstAdd)
	if !ok {
		panic(fmt.Errorf("invalid IR instruction for AST instruction; expected *ir.InstAdd, got %T", new))
	}
	// X operand.
	x, err := fgen.irTypeValue(old.X())
	if err != nil {
		return errors.WithStack(err)
	}
	inst.X = x
	// Y operand.
	y, err := fgen.irValue(x.Type(), old.Y())
	if err != nil {
		return errors.WithStack(err)
	}
	inst.Y = y
	// (optional) Overflow flags.
	inst.OverflowFlags = irOverflowFlags(old.OverflowFlags())
	// (optional) Metadata.
	md, err := fgen.gen.irMetadataAttachments(old.Metadata())
	if err != nil {
		return errors.WithStack(err)
	}
	inst.Metadata = md
	return nil
}

// --- [ fadd ] ----------------------------------------------------------------

// irFAddInst translates the given AST fadd instruction into an equivalent IR
// instruction.
func (fgen *funcGen) irFAddInst(new ir.Instruction, old *ast.FAddInst) error {
	inst, ok := new.(*ir.InstFAdd)
	if !ok {
		panic(fmt.Errorf("invalid IR instruction for AST instruction; expected *ir.InstFAdd, got %T", new))
	}
	// X operand.
	x, err := fgen.irTypeValue(old.X())
	if err != nil {
		return errors.WithStack(err)
	}
	inst.X = x
	// Y operand.
	y, err := fgen.irValue(x.Type(), old.Y())
	if err != nil {
		return errors.WithStack(err)
	}
	inst.Y = y
	// (optional) Fast math flags.
	inst.FastMathFlags = irFastMathFlags(old.FastMathFlags())
	// (optional) Metadata.
	md, err := fgen.gen.irMetadataAttachments(old.Metadata())
	if err != nil {
		return errors.WithStack(err)
	}
	inst.Metadata = md
	return nil
}

// --- [ sub ] -----------------------------------------------------------------

// irSubInst translates the given AST sub instruction into an equivalent IR
// instruction.
func (fgen *funcGen) irSubInst(new ir.Instruction, old *ast.SubInst) error {
	inst, ok := new.(*ir.InstSub)
	if !ok {
		panic(fmt.Errorf("invalid IR instruction for AST instruction; expected *ir.InstSub, got %T", new))
	}
	// X operand.
	x, err := fgen.irTypeValue(old.X())
	if err != nil {
		return errors.WithStack(err)
	}
	inst.X = x
	// Y operand.
	y, err := fgen.irValue(x.Type(), old.Y())
	if err != nil {
		return errors.WithStack(err)
	}
	inst.Y = y
	// (optional) Overflow flags.
	inst.OverflowFlags = irOverflowFlags(old.OverflowFlags())
	// (optional) Metadata.
	md, err := fgen.gen.irMetadataAttachments(old.Metadata())
	if err != nil {
		return errors.WithStack(err)
	}
	inst.Metadata = md
	return nil
}

// --- [ fsub ] ----------------------------------------------------------------

// irFSubInst translates the given AST fsub instruction into an equivalent IR
// instruction.
func (fgen *funcGen) irFSubInst(new ir.Instruction, old *ast.FSubInst) error {
	inst, ok := new.(*ir.InstFSub)
	if !ok {
		panic(fmt.Errorf("invalid IR instruction for AST instruction; expected *ir.InstFSub, got %T", new))
	}
	// X operand.
	x, err := fgen.irTypeValue(old.X())
	if err != nil {
		return errors.WithStack(err)
	}
	inst.X = x
	// Y operand.
	y, err := fgen.irValue(x.Type(), old.Y())
	if err != nil {
		return errors.WithStack(err)
	}
	inst.Y = y
	// (optional) Fast math flags.
	inst.FastMathFlags = irFastMathFlags(old.FastMathFlags())
	// (optional) Metadata.
	md, err := fgen.gen.irMetadataAttachments(old.Metadata())
	if err != nil {
		return errors.WithStack(err)
	}
	inst.Metadata = md
	return nil
}

// --- [ mul ] -----------------------------------------------------------------

// irMulInst translates the given AST mul instruction into an equivalent IR
// instruction.
func (fgen *funcGen) irMulInst(new ir.Instruction, old *ast.MulInst) error {
	inst, ok := new.(*ir.InstMul)
	if !ok {
		panic(fmt.Errorf("invalid IR instruction for AST instruction; expected *ir.InstMul, got %T", new))
	}
	// X operand.
	x, err := fgen.irTypeValue(old.X())
	if err != nil {
		return errors.WithStack(err)
	}
	inst.X = x
	// Y operand.
	y, err := fgen.irValue(x.Type(), old.Y())
	if err != nil {
		return errors.WithStack(err)
	}
	inst.Y = y
	// (optional) Overflow flags.
	inst.OverflowFlags = irOverflowFlags(old.OverflowFlags())
	// (optional) Metadata.
	md, err := fgen.gen.irMetadataAttachments(old.Metadata())
	if err != nil {
		return errors.WithStack(err)
	}
	inst.Metadata = md
	return nil
}

// --- [ fmul ] ----------------------------------------------------------------

// irFMulInst translates the given AST fmul instruction into an equivalent IR
// instruction.
func (fgen *funcGen) irFMulInst(new ir.Instruction, old *ast.FMulInst) error {
	inst, ok := new.(*ir.InstFMul)
	if !ok {
		panic(fmt.Errorf("invalid IR instruction for AST instruction; expected *ir.InstFMul, got %T", new))
	}
	// X operand.
	x, err := fgen.irTypeValue(old.X())
	if err != nil {
		return errors.WithStack(err)
	}
	inst.X = x
	// Y operand.
	y, err := fgen.irValue(x.Type(), old.Y())
	if err != nil {
		return errors.WithStack(err)
	}
	inst.Y = y
	// (optional) Fast math flags.
	inst.FastMathFlags = irFastMathFlags(old.FastMathFlags())
	// (optional) Metadata.
	md, err := fgen.gen.irMetadataAttachments(old.Metadata())
	if err != nil {
		return errors.WithStack(err)
	}
	inst.Metadata = md
	return nil
}

// --- [ udiv ] ----------------------------------------------------------------

// irUDivInst translates the given AST udiv instruction into an equivalent IR
// instruction.
func (fgen *funcGen) irUDivInst(new ir.Instruction, old *ast.UDivInst) error {
	inst, ok := new.(*ir.InstUDiv)
	if !ok {
		panic(fmt.Errorf("invalid IR instruction for AST instruction; expected *ir.InstUDiv, got %T", new))
	}
	// X operand.
	x, err := fgen.irTypeValue(old.X())
	if err != nil {
		return errors.WithStack(err)
	}
	inst.X = x
	// Y operand.
	y, err := fgen.irValue(x.Type(), old.Y())
	if err != nil {
		return errors.WithStack(err)
	}
	inst.Y = y
	// (optional) Exact.
	_, inst.Exact = old.Exact()
	// (optional) Metadata.
	md, err := fgen.gen.irMetadataAttachments(old.Metadata())
	if err != nil {
		return errors.WithStack(err)
	}
	inst.Metadata = md
	return nil
}

// --- [ sdiv ] ----------------------------------------------------------------

// irSDivInst translates the given AST sdiv instruction into an equivalent IR
// instruction.
func (fgen *funcGen) irSDivInst(new ir.Instruction, old *ast.SDivInst) error {
	inst, ok := new.(*ir.InstSDiv)
	if !ok {
		panic(fmt.Errorf("invalid IR instruction for AST instruction; expected *ir.InstSDiv, got %T", new))
	}
	// X operand.
	x, err := fgen.irTypeValue(old.X())
	if err != nil {
		return errors.WithStack(err)
	}
	inst.X = x
	// Y operand.
	y, err := fgen.irValue(x.Type(), old.Y())
	if err != nil {
		return errors.WithStack(err)
	}
	inst.Y = y
	// (optional) Exact.
	_, inst.Exact = old.Exact()
	// (optional) Metadata.
	md, err := fgen.gen.irMetadataAttachments(old.Metadata())
	if err != nil {
		return errors.WithStack(err)
	}
	inst.Metadata = md
	return nil
}

// --- [ fdiv ] ----------------------------------------------------------------

// irFDivInst translates the given AST fdiv instruction into an equivalent IR
// instruction.
func (fgen *funcGen) irFDivInst(new ir.Instruction, old *ast.FDivInst) error {
	inst, ok := new.(*ir.InstFDiv)
	if !ok {
		panic(fmt.Errorf("invalid IR instruction for AST instruction; expected *ir.InstFDiv, got %T", new))
	}
	// X operand.
	x, err := fgen.irTypeValue(old.X())
	if err != nil {
		return errors.WithStack(err)
	}
	inst.X = x
	// Y operand.
	y, err := fgen.irValue(x.Type(), old.Y())
	if err != nil {
		return errors.WithStack(err)
	}
	inst.Y = y
	// (optional) Fast math flags.
	inst.FastMathFlags = irFastMathFlags(old.FastMathFlags())
	// (optional) Metadata.
	md, err := fgen.gen.irMetadataAttachments(old.Metadata())
	if err != nil {
		return errors.WithStack(err)
	}
	inst.Metadata = md
	return nil
}

// --- [ urem ] ----------------------------------------------------------------

// irURemInst translates the given AST urem instruction into an equivalent IR
// instruction.
func (fgen *funcGen) irURemInst(new ir.Instruction, old *ast.URemInst) error {
	inst, ok := new.(*ir.InstURem)
	if !ok {
		panic(fmt.Errorf("invalid IR instruction for AST instruction; expected *ir.InstURem, got %T", new))
	}
	// X operand.
	x, err := fgen.irTypeValue(old.X())
	if err != nil {
		return errors.WithStack(err)
	}
	inst.X = x
	// Y operand.
	y, err := fgen.irValue(x.Type(), old.Y())
	if err != nil {
		return errors.WithStack(err)
	}
	inst.Y = y
	// (optional) Metadata.
	md, err := fgen.gen.irMetadataAttachments(old.Metadata())
	if err != nil {
		return errors.WithStack(err)
	}
	inst.Metadata = md
	return nil
}

// --- [ srem ] ----------------------------------------------------------------

// irSRemInst translates the given AST srem instruction into an equivalent IR
// instruction.
func (fgen *funcGen) irSRemInst(new ir.Instruction, old *ast.SRemInst) error {
	inst, ok := new.(*ir.InstSRem)
	if !ok {
		panic(fmt.Errorf("invalid IR instruction for AST instruction; expected *ir.InstSRem, got %T", new))
	}
	// X operand.
	x, err := fgen.irTypeValue(old.X())
	if err != nil {
		return errors.WithStack(err)
	}
	inst.X = x
	// Y operand.
	y, err := fgen.irValue(x.Type(), old.Y())
	if err != nil {
		return errors.WithStack(err)
	}
	inst.Y = y
	// (optional) Metadata.
	md, err := fgen.gen.irMetadataAttachments(old.Metadata())
	if err != nil {
		return errors.WithStack(err)
	}
	inst.Metadata = md
	return nil
}

// --- [ frem ] ----------------------------------------------------------------

// irFRemInst translates the given AST frem instruction into an equivalent IR
// instruction.
func (fgen *funcGen) irFRemInst(new ir.Instruction, old *ast.FRemInst) error {
	inst, ok := new.(*ir.InstFRem)
	if !ok {
		panic(fmt.Errorf("invalid IR instruction for AST instruction; expected *ir.InstFRem, got %T", new))
	}
	// X operand.
	x, err := fgen.irTypeValue(old.X())
	if err != nil {
		return errors.WithStack(err)
	}
	inst.X = x
	// Y operand.
	y, err := fgen.irValue(x.Type(), old.Y())
	if err != nil {
		return errors.WithStack(err)
	}
	inst.Y = y
	// (optional) Fast math flags.
	inst.FastMathFlags = irFastMathFlags(old.FastMathFlags())
	// (optional) Metadata.
	md, err := fgen.gen.irMetadataAttachments(old.Metadata())
	if err != nil {
		return errors.WithStack(err)
	}
	inst.Metadata = md
	return nil
}
