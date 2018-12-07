package asm

import (
	"fmt"

	"github.com/llir/ll/ast"
	"github.com/llir/llvm/ir"
	"github.com/pkg/errors"
)

// === [ Create IR ] ===========================================================

// newTruncInst returns a new IR trunc instruction (without body but with type)
// based on the given AST trunc instruction.
func (fgen *funcGen) newTruncInst(ident ir.LocalIdent, old *ast.TruncInst) (*ir.InstTrunc, error) {
	to, err := fgen.gen.irType(old.To())
	if err != nil {
		return nil, errors.WithStack(err)
	}
	return &ir.InstTrunc{LocalIdent: ident, To: to}, nil
}

// newZExtInst returns a new IR zext instruction (without body but with type)
// based on the given AST zext instruction.
func (fgen *funcGen) newZExtInst(ident ir.LocalIdent, old *ast.ZExtInst) (*ir.InstZExt, error) {
	to, err := fgen.gen.irType(old.To())
	if err != nil {
		return nil, errors.WithStack(err)
	}
	return &ir.InstZExt{LocalIdent: ident, To: to}, nil
}

// newSExtInst returns a new IR sext instruction (without body but with type)
// based on the given AST sext instruction.
func (fgen *funcGen) newSExtInst(ident ir.LocalIdent, old *ast.SExtInst) (*ir.InstSExt, error) {
	to, err := fgen.gen.irType(old.To())
	if err != nil {
		return nil, errors.WithStack(err)
	}
	return &ir.InstSExt{LocalIdent: ident, To: to}, nil
}

// newFPTruncInst returns a new IR fptrunc instruction (without body but with
// type) based on the given AST fptrunc instruction.
func (fgen *funcGen) newFPTruncInst(ident ir.LocalIdent, old *ast.FPTruncInst) (*ir.InstFPTrunc, error) {
	to, err := fgen.gen.irType(old.To())
	if err != nil {
		return nil, errors.WithStack(err)
	}
	return &ir.InstFPTrunc{LocalIdent: ident, To: to}, nil
}

// newFPExtInst returns a new IR fpext instruction (without body but with type)
// based on the given AST fpext instruction.
func (fgen *funcGen) newFPExtInst(ident ir.LocalIdent, old *ast.FPExtInst) (*ir.InstFPExt, error) {
	to, err := fgen.gen.irType(old.To())
	if err != nil {
		return nil, errors.WithStack(err)
	}
	return &ir.InstFPExt{LocalIdent: ident, To: to}, nil
}

// newFPToUIInst returns a new IR fptoui instruction (without body but with
// type) based on the given AST fptoui instruction.
func (fgen *funcGen) newFPToUIInst(ident ir.LocalIdent, old *ast.FPToUIInst) (*ir.InstFPToUI, error) {
	to, err := fgen.gen.irType(old.To())
	if err != nil {
		return nil, errors.WithStack(err)
	}
	return &ir.InstFPToUI{LocalIdent: ident, To: to}, nil
}

// newFPToSIInst returns a new IR fptosi instruction (without body but with
// type) based on the given AST fptosi instruction.
func (fgen *funcGen) newFPToSIInst(ident ir.LocalIdent, old *ast.FPToSIInst) (*ir.InstFPToSI, error) {
	to, err := fgen.gen.irType(old.To())
	if err != nil {
		return nil, errors.WithStack(err)
	}
	return &ir.InstFPToSI{LocalIdent: ident, To: to}, nil
}

// newUIToFPInst returns a new IR uitofp instruction (without body but with
// type) based on the given AST uitofp instruction.
func (fgen *funcGen) newUIToFPInst(ident ir.LocalIdent, old *ast.UIToFPInst) (*ir.InstUIToFP, error) {
	to, err := fgen.gen.irType(old.To())
	if err != nil {
		return nil, errors.WithStack(err)
	}
	return &ir.InstUIToFP{LocalIdent: ident, To: to}, nil
}

// newSIToFPInst returns a new IR sitofp instruction (without body but with
// type) based on the given AST sitofp instruction.
func (fgen *funcGen) newSIToFPInst(ident ir.LocalIdent, old *ast.SIToFPInst) (*ir.InstSIToFP, error) {
	to, err := fgen.gen.irType(old.To())
	if err != nil {
		return nil, errors.WithStack(err)
	}
	return &ir.InstSIToFP{LocalIdent: ident, To: to}, nil
}

// newPtrToIntInst returns a new IR ptrtoint instruction (without body but with
// type) based on the given AST ptrtoint instruction.
func (fgen *funcGen) newPtrToIntInst(ident ir.LocalIdent, old *ast.PtrToIntInst) (*ir.InstPtrToInt, error) {
	to, err := fgen.gen.irType(old.To())
	if err != nil {
		return nil, errors.WithStack(err)
	}
	return &ir.InstPtrToInt{LocalIdent: ident, To: to}, nil
}

// newIntToPtrInst returns a new IR inttoptr instruction (without body but with
// type) based on the given AST inttoptr instruction.
func (fgen *funcGen) newIntToPtrInst(ident ir.LocalIdent, old *ast.IntToPtrInst) (*ir.InstIntToPtr, error) {
	to, err := fgen.gen.irType(old.To())
	if err != nil {
		return nil, errors.WithStack(err)
	}
	return &ir.InstIntToPtr{LocalIdent: ident, To: to}, nil
}

// newBitCastInst returns a new IR bitcast instruction (without body but with
// type) based on the given AST bitcast instruction.
func (fgen *funcGen) newBitCastInst(ident ir.LocalIdent, old *ast.BitCastInst) (*ir.InstBitCast, error) {
	to, err := fgen.gen.irType(old.To())
	if err != nil {
		return nil, errors.WithStack(err)
	}
	return &ir.InstBitCast{LocalIdent: ident, To: to}, nil
}

// newAddrSpaceCastInst returns a new IR addrspacecast instruction (without body
// but with type) based on the given AST addrspacecast instruction.
func (fgen *funcGen) newAddrSpaceCastInst(ident ir.LocalIdent, old *ast.AddrSpaceCastInst) (*ir.InstAddrSpaceCast, error) {
	to, err := fgen.gen.irType(old.To())
	if err != nil {
		return nil, errors.WithStack(err)
	}
	return &ir.InstAddrSpaceCast{LocalIdent: ident, To: to}, nil
}

// === [ Translate AST to IR ] =================================================

// --- [ trunc ] ---------------------------------------------------------------

// irTruncInst translates the given AST trunc instruction into an equivalent IR
// instruction.
func (fgen *funcGen) irTruncInst(new ir.Instruction, old *ast.TruncInst) error {
	inst, ok := new.(*ir.InstTrunc)
	if !ok {
		panic(fmt.Errorf("invalid IR instruction for AST instruction; expected *ir.InstTrunc, got %T", new))
	}
	// Value before conversion.
	from, err := fgen.irTypeValue(old.From())
	if err != nil {
		return errors.WithStack(err)
	}
	inst.From = from
	// Type after conversion.
	to, err := fgen.gen.irType(old.To())
	if err != nil {
		return errors.WithStack(err)
	}
	inst.To = to
	// (optional) Metadata.
	md, err := fgen.gen.irMetadataAttachments(old.Metadata())
	if err != nil {
		return errors.WithStack(err)
	}
	inst.Metadata = md
	return nil
}

// --- [ zext ] ----------------------------------------------------------------

// irZExtInst translates the given AST zext instruction into an equivalent IR
// instruction.
func (fgen *funcGen) irZExtInst(new ir.Instruction, old *ast.ZExtInst) error {
	inst, ok := new.(*ir.InstZExt)
	if !ok {
		panic(fmt.Errorf("invalid IR instruction for AST instruction; expected *ir.InstZExt, got %T", new))
	}
	// Value before conversion.
	from, err := fgen.irTypeValue(old.From())
	if err != nil {
		return errors.WithStack(err)
	}
	inst.From = from
	// Type after conversion.
	to, err := fgen.gen.irType(old.To())
	if err != nil {
		return errors.WithStack(err)
	}
	inst.To = to
	// (optional) Metadata.
	md, err := fgen.gen.irMetadataAttachments(old.Metadata())
	if err != nil {
		return errors.WithStack(err)
	}
	inst.Metadata = md
	return nil
}

// --- [ sext ] ----------------------------------------------------------------

// irSExtInst translates the given AST sext instruction into an equivalent IR
// instruction.
func (fgen *funcGen) irSExtInst(new ir.Instruction, old *ast.SExtInst) error {
	inst, ok := new.(*ir.InstSExt)
	if !ok {
		panic(fmt.Errorf("invalid IR instruction for AST instruction; expected *ir.InstSExt, got %T", new))
	}
	// Value before conversion.
	from, err := fgen.irTypeValue(old.From())
	if err != nil {
		return errors.WithStack(err)
	}
	inst.From = from
	// Type after conversion.
	to, err := fgen.gen.irType(old.To())
	if err != nil {
		return errors.WithStack(err)
	}
	inst.To = to
	// (optional) Metadata.
	md, err := fgen.gen.irMetadataAttachments(old.Metadata())
	if err != nil {
		return errors.WithStack(err)
	}
	inst.Metadata = md
	return nil
}

// --- [ fptrunc ] -------------------------------------------------------------

// irFPTruncInst translates the given AST fptrunc instruction into an equivalent
// IR instruction.
func (fgen *funcGen) irFPTruncInst(new ir.Instruction, old *ast.FPTruncInst) error {
	inst, ok := new.(*ir.InstFPTrunc)
	if !ok {
		panic(fmt.Errorf("invalid IR instruction for AST instruction; expected *ir.InstFPTrunc, got %T", new))
	}
	// Value before conversion.
	from, err := fgen.irTypeValue(old.From())
	if err != nil {
		return errors.WithStack(err)
	}
	inst.From = from
	// Type after conversion.
	to, err := fgen.gen.irType(old.To())
	if err != nil {
		return errors.WithStack(err)
	}
	inst.To = to
	// (optional) Metadata.
	md, err := fgen.gen.irMetadataAttachments(old.Metadata())
	if err != nil {
		return errors.WithStack(err)
	}
	inst.Metadata = md
	return nil
}

// --- [ fpext ] ---------------------------------------------------------------

// irFPExtInst translates the given AST fpext instruction into an equivalent IR
// instruction.
func (fgen *funcGen) irFPExtInst(new ir.Instruction, old *ast.FPExtInst) error {
	inst, ok := new.(*ir.InstFPExt)
	if !ok {
		panic(fmt.Errorf("invalid IR instruction for AST instruction; expected *ir.InstFPExt, got %T", new))
	}
	// Value before conversion.
	from, err := fgen.irTypeValue(old.From())
	if err != nil {
		return errors.WithStack(err)
	}
	inst.From = from
	// Type after conversion.
	to, err := fgen.gen.irType(old.To())
	if err != nil {
		return errors.WithStack(err)
	}
	inst.To = to
	// (optional) Metadata.
	md, err := fgen.gen.irMetadataAttachments(old.Metadata())
	if err != nil {
		return errors.WithStack(err)
	}
	inst.Metadata = md
	return nil
}

// --- [ fptoui ] --------------------------------------------------------------

// irFPToUIInst translates the given AST fptoui instruction into an equivalent
// IR instruction.
func (fgen *funcGen) irFPToUIInst(new ir.Instruction, old *ast.FPToUIInst) error {
	inst, ok := new.(*ir.InstFPToUI)
	if !ok {
		panic(fmt.Errorf("invalid IR instruction for AST instruction; expected *ir.InstFPToUI, got %T", new))
	}
	// Value before conversion.
	from, err := fgen.irTypeValue(old.From())
	if err != nil {
		return errors.WithStack(err)
	}
	inst.From = from
	// Type after conversion.
	to, err := fgen.gen.irType(old.To())
	if err != nil {
		return errors.WithStack(err)
	}
	inst.To = to
	// (optional) Metadata.
	md, err := fgen.gen.irMetadataAttachments(old.Metadata())
	if err != nil {
		return errors.WithStack(err)
	}
	inst.Metadata = md
	return nil
}

// --- [ fptosi ] --------------------------------------------------------------

// irFPToSIInst translates the given AST fptosi instruction into an equivalent
// IR instruction.
func (fgen *funcGen) irFPToSIInst(new ir.Instruction, old *ast.FPToSIInst) error {
	inst, ok := new.(*ir.InstFPToSI)
	if !ok {
		panic(fmt.Errorf("invalid IR instruction for AST instruction; expected *ir.InstFPToSI, got %T", new))
	}
	// Value before conversion.
	from, err := fgen.irTypeValue(old.From())
	if err != nil {
		return errors.WithStack(err)
	}
	inst.From = from
	// Type after conversion.
	to, err := fgen.gen.irType(old.To())
	if err != nil {
		return errors.WithStack(err)
	}
	inst.To = to
	// (optional) Metadata.
	md, err := fgen.gen.irMetadataAttachments(old.Metadata())
	if err != nil {
		return errors.WithStack(err)
	}
	inst.Metadata = md
	return nil
}

// --- [ uitofp ] --------------------------------------------------------------

// irUIToFPInst translates the given AST uitofp instruction into an equivalent
// IR instruction.
func (fgen *funcGen) irUIToFPInst(new ir.Instruction, old *ast.UIToFPInst) error {
	inst, ok := new.(*ir.InstUIToFP)
	if !ok {
		panic(fmt.Errorf("invalid IR instruction for AST instruction; expected *ir.InstUIToFP, got %T", new))
	}
	// Value before conversion.
	from, err := fgen.irTypeValue(old.From())
	if err != nil {
		return errors.WithStack(err)
	}
	inst.From = from
	// Type after conversion.
	to, err := fgen.gen.irType(old.To())
	if err != nil {
		return errors.WithStack(err)
	}
	inst.To = to
	// (optional) Metadata.
	md, err := fgen.gen.irMetadataAttachments(old.Metadata())
	if err != nil {
		return errors.WithStack(err)
	}
	inst.Metadata = md
	return nil
}

// --- [ sitofp ] --------------------------------------------------------------

// irSIToFPInst translates the given AST sitofp instruction into an equivalent
// IR instruction.
func (fgen *funcGen) irSIToFPInst(new ir.Instruction, old *ast.SIToFPInst) error {
	inst, ok := new.(*ir.InstSIToFP)
	if !ok {
		panic(fmt.Errorf("invalid IR instruction for AST instruction; expected *ir.InstSIToFP, got %T", new))
	}
	// Value before conversion.
	from, err := fgen.irTypeValue(old.From())
	if err != nil {
		return errors.WithStack(err)
	}
	inst.From = from
	// Type after conversion.
	to, err := fgen.gen.irType(old.To())
	if err != nil {
		return errors.WithStack(err)
	}
	inst.To = to
	// (optional) Metadata.
	md, err := fgen.gen.irMetadataAttachments(old.Metadata())
	if err != nil {
		return errors.WithStack(err)
	}
	inst.Metadata = md
	return nil
}

// --- [ ptrtoint ] ------------------------------------------------------------

// irPtrToIntInst translates the given AST ptrtoint instruction into an
// equivalent IR instruction.
func (fgen *funcGen) irPtrToIntInst(new ir.Instruction, old *ast.PtrToIntInst) error {
	inst, ok := new.(*ir.InstPtrToInt)
	if !ok {
		panic(fmt.Errorf("invalid IR instruction for AST instruction; expected *ir.InstPtrToInt, got %T", new))
	}
	// Value before conversion.
	from, err := fgen.irTypeValue(old.From())
	if err != nil {
		return errors.WithStack(err)
	}
	inst.From = from
	// Type after conversion.
	to, err := fgen.gen.irType(old.To())
	if err != nil {
		return errors.WithStack(err)
	}
	inst.To = to
	// (optional) Metadata.
	md, err := fgen.gen.irMetadataAttachments(old.Metadata())
	if err != nil {
		return errors.WithStack(err)
	}
	inst.Metadata = md
	return nil
}

// --- [ inttoptr ] ------------------------------------------------------------

// irIntToPtrInst translates the given AST inttoptr instruction into an
// equivalent IR instruction.
func (fgen *funcGen) irIntToPtrInst(new ir.Instruction, old *ast.IntToPtrInst) error {
	inst, ok := new.(*ir.InstIntToPtr)
	if !ok {
		panic(fmt.Errorf("invalid IR instruction for AST instruction; expected *ir.InstIntToPtr, got %T", new))
	}
	// Value before conversion.
	from, err := fgen.irTypeValue(old.From())
	if err != nil {
		return errors.WithStack(err)
	}
	inst.From = from
	// Type after conversion.
	to, err := fgen.gen.irType(old.To())
	if err != nil {
		return errors.WithStack(err)
	}
	inst.To = to
	// (optional) Metadata.
	md, err := fgen.gen.irMetadataAttachments(old.Metadata())
	if err != nil {
		return errors.WithStack(err)
	}
	inst.Metadata = md
	return nil
}

// --- [ bitcast ] -------------------------------------------------------------

// irBitCastInst translates the given AST bitcast instruction into an equivalent
// IR instruction.
func (fgen *funcGen) irBitCastInst(new ir.Instruction, old *ast.BitCastInst) error {
	inst, ok := new.(*ir.InstBitCast)
	if !ok {
		panic(fmt.Errorf("invalid IR instruction for AST instruction; expected *ir.InstBitCast, got %T", new))
	}
	// Value before conversion.
	from, err := fgen.irTypeValue(old.From())
	if err != nil {
		return errors.WithStack(err)
	}
	inst.From = from
	// Type after conversion.
	to, err := fgen.gen.irType(old.To())
	if err != nil {
		return errors.WithStack(err)
	}
	inst.To = to
	// (optional) Metadata.
	md, err := fgen.gen.irMetadataAttachments(old.Metadata())
	if err != nil {
		return errors.WithStack(err)
	}
	inst.Metadata = md
	return nil
}

// --- [ addrspacecast ] -------------------------------------------------------

// irAddrSpaceCastInst translates the given AST addrspacecast instruction into
// an equivalent IR instruction.
func (fgen *funcGen) irAddrSpaceCastInst(new ir.Instruction, old *ast.AddrSpaceCastInst) error {
	inst, ok := new.(*ir.InstAddrSpaceCast)
	if !ok {
		panic(fmt.Errorf("invalid IR instruction for AST instruction; expected *ir.InstAddrSpaceCast, got %T", new))
	}
	// Value before conversion.
	from, err := fgen.irTypeValue(old.From())
	if err != nil {
		return errors.WithStack(err)
	}
	inst.From = from
	// Type after conversion.
	to, err := fgen.gen.irType(old.To())
	if err != nil {
		return errors.WithStack(err)
	}
	inst.To = to
	// (optional) Metadata.
	md, err := fgen.gen.irMetadataAttachments(old.Metadata())
	if err != nil {
		return errors.WithStack(err)
	}
	inst.Metadata = md
	return nil
}
