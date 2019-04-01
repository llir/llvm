package asm

import (
	"fmt"

	"github.com/llir/ll/ast"
	"github.com/llir/llvm/ir"
	"github.com/pkg/errors"
)

// === [ Create IR ] ===========================================================

// newFNegInst returns a new IR fneg instruction (without body but with type)
// based on the given AST fneg instruction.
func (fgen *funcGen) newFNegInst(ident ir.LocalIdent, old *ast.FNegInst) (*ir.InstFNeg, error) {
	typ, err := fgen.gen.irType(old.X().Typ())
	if err != nil {
		return nil, errors.WithStack(err)
	}
	return &ir.InstFNeg{LocalIdent: ident, Typ: typ}, nil
}

// === [ Translate AST to IR ] =================================================

// --- [ fneg ] ----------------------------------------------------------------

// irFNegInst translates the given AST fneg instruction into an equivalent IR
// instruction.
func (fgen *funcGen) irFNegInst(new ir.Instruction, old *ast.FNegInst) error {
	inst, ok := new.(*ir.InstFNeg)
	if !ok {
		panic(fmt.Errorf("invalid IR instruction for AST instruction; expected *ir.InstFNeg, got %T", new))
	}
	// X operand.
	x, err := fgen.irTypeValue(old.X())
	if err != nil {
		return errors.WithStack(err)
	}
	inst.X = x
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
