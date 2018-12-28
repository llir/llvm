package asm

import (
	"fmt"

	"github.com/llir/ll/ast"
	"github.com/llir/llvm/ir"
	"github.com/pkg/errors"
)

// === [ Create IR ] ===========================================================

// newShlInst returns a new IR shl instruction (without body but with type)
// based on the given AST shl instruction.
func (fgen *funcGen) newShlInst(ident ir.LocalIdent, old *ast.ShlInst) (*ir.InstShl, error) {
	typ, err := fgen.gen.irType(old.X().Typ())
	if err != nil {
		return nil, errors.WithStack(err)
	}
	return &ir.InstShl{LocalIdent: ident, Typ: typ}, nil
}

// newLShrInst returns a new IR lshr instruction (without body but with type)
// based on the given AST lshr instruction.
func (fgen *funcGen) newLShrInst(ident ir.LocalIdent, old *ast.LShrInst) (*ir.InstLShr, error) {
	typ, err := fgen.gen.irType(old.X().Typ())
	if err != nil {
		return nil, errors.WithStack(err)
	}
	return &ir.InstLShr{LocalIdent: ident, Typ: typ}, nil
}

// newAShrInst returns a new IR ashr instruction (without body but with type)
// based on the given AST ashr instruction.
func (fgen *funcGen) newAShrInst(ident ir.LocalIdent, old *ast.AShrInst) (*ir.InstAShr, error) {
	typ, err := fgen.gen.irType(old.X().Typ())
	if err != nil {
		return nil, errors.WithStack(err)
	}
	return &ir.InstAShr{LocalIdent: ident, Typ: typ}, nil
}

// newAndInst returns a new IR and instruction (without body but with type)
// based on the given AST and instruction.
func (fgen *funcGen) newAndInst(ident ir.LocalIdent, old *ast.AndInst) (*ir.InstAnd, error) {
	typ, err := fgen.gen.irType(old.X().Typ())
	if err != nil {
		return nil, errors.WithStack(err)
	}
	return &ir.InstAnd{LocalIdent: ident, Typ: typ}, nil
}

// newOrInst returns a new IR or instruction (without body but with type) based
// on the given AST or instruction.
func (fgen *funcGen) newOrInst(ident ir.LocalIdent, old *ast.OrInst) (*ir.InstOr, error) {
	typ, err := fgen.gen.irType(old.X().Typ())
	if err != nil {
		return nil, errors.WithStack(err)
	}
	return &ir.InstOr{LocalIdent: ident, Typ: typ}, nil
}

// newXorInst returns a new IR xor instruction (without body but with type)
// based on the given AST xor instruction.
func (fgen *funcGen) newXorInst(ident ir.LocalIdent, old *ast.XorInst) (*ir.InstXor, error) {
	typ, err := fgen.gen.irType(old.X().Typ())
	if err != nil {
		return nil, errors.WithStack(err)
	}
	return &ir.InstXor{LocalIdent: ident, Typ: typ}, nil
}

// === [ Translate AST to IR ] =================================================

// --- [ shl ] -----------------------------------------------------------------

// irShlInst translates the given AST shl instruction into an equivalent IR
// instruction.
func (fgen *funcGen) irShlInst(new ir.Instruction, old *ast.ShlInst) error {
	inst, ok := new.(*ir.InstShl)
	if !ok {
		panic(fmt.Errorf("invalid IR instruction for AST instruction; expected *ir.InstShl, got %T", new))
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

// --- [ lshr ] ----------------------------------------------------------------

// irLShrInst translates the given AST lshr instruction into an equivalent IR
// instruction.
func (fgen *funcGen) irLShrInst(new ir.Instruction, old *ast.LShrInst) error {
	inst, ok := new.(*ir.InstLShr)
	if !ok {
		panic(fmt.Errorf("invalid IR instruction for AST instruction; expected *ir.InstLShr, got %T", new))
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

// --- [ ashr ] ----------------------------------------------------------------

// irAShrInst translates the given AST ashr instruction into an equivalent IR
// instruction.
func (fgen *funcGen) irAShrInst(new ir.Instruction, old *ast.AShrInst) error {
	inst, ok := new.(*ir.InstAShr)
	if !ok {
		panic(fmt.Errorf("invalid IR instruction for AST instruction; expected *ir.InstAShr, got %T", new))
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

// --- [ and ] -----------------------------------------------------------------

// irAndInst translates the given AST and instruction into an equivalent IR
// instruction.
func (fgen *funcGen) irAndInst(new ir.Instruction, old *ast.AndInst) error {
	inst, ok := new.(*ir.InstAnd)
	if !ok {
		panic(fmt.Errorf("invalid IR instruction for AST instruction; expected *ir.InstAnd, got %T", new))
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

// --- [ or ] ------------------------------------------------------------------

// irOrInst translates the given AST or instruction into an equivalent IR
// instruction.
func (fgen *funcGen) irOrInst(new ir.Instruction, old *ast.OrInst) error {
	inst, ok := new.(*ir.InstOr)
	if !ok {
		panic(fmt.Errorf("invalid IR instruction for AST instruction; expected *ir.InstOr, got %T", new))
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

// --- [ xor ] -----------------------------------------------------------------

// irXorInst translates the given AST xor instruction into an equivalent IR
// instruction.
func (fgen *funcGen) irXorInst(new ir.Instruction, old *ast.XorInst) error {
	inst, ok := new.(*ir.InstXor)
	if !ok {
		panic(fmt.Errorf("invalid IR instruction for AST instruction; expected *ir.InstXor, got %T", new))
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
