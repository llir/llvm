package asm

import (
	"fmt"

	"github.com/llir/ll/ast"
	"github.com/llir/llvm/ir"
	"github.com/pkg/errors"
)

// --- [ Conversion instructions ] ---------------------------------------------

// ~~~ [ trunc ] ~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~

// astToIRInstTrunc translates the given AST trunc instruction into an
// equivalent IR instruction.
func (fgen *funcGen) astToIRInstTrunc(inst ir.Instruction, old *ast.TruncInst) (*ir.InstTrunc, error) {
	i, ok := inst.(*ir.InstTrunc)
	if !ok {
		panic(fmt.Errorf("invalid IR instruction for AST instruction; expected *ir.InstTrunc, got %T", inst))
	}
	// Value before conversion.
	from, err := fgen.astToIRTypeValue(old.From())
	if err != nil {
		return nil, errors.WithStack(err)
	}
	i.From = from
	// Type after conversion.
	to, err := fgen.gen.irType(old.To())
	if err != nil {
		return nil, errors.WithStack(err)
	}
	i.To = to
	// (optional) Metadata.
	md, err := fgen.gen.irMetadataAttachments(old.Metadata())
	if err != nil {
		return nil, errors.WithStack(err)
	}
	i.Metadata = md
	return i, nil
}

// ~~~ [ zext ] ~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~

// astToIRInstZExt translates the given AST zext instruction into an equivalent
// IR instruction.
func (fgen *funcGen) astToIRInstZExt(inst ir.Instruction, old *ast.ZExtInst) (*ir.InstZExt, error) {
	i, ok := inst.(*ir.InstZExt)
	if !ok {
		panic(fmt.Errorf("invalid IR instruction for AST instruction; expected *ir.InstZExt, got %T", inst))
	}
	// Value before conversion.
	from, err := fgen.astToIRTypeValue(old.From())
	if err != nil {
		return nil, errors.WithStack(err)
	}
	i.From = from
	// Type after conversion.
	to, err := fgen.gen.irType(old.To())
	if err != nil {
		return nil, errors.WithStack(err)
	}
	i.To = to
	// (optional) Metadata.
	md, err := fgen.gen.irMetadataAttachments(old.Metadata())
	if err != nil {
		return nil, errors.WithStack(err)
	}
	i.Metadata = md
	return i, nil
}

// ~~~ [ sext ] ~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~

// astToIRInstSExt translates the given AST sext instruction into an equivalent
// IR instruction.
func (fgen *funcGen) astToIRInstSExt(inst ir.Instruction, old *ast.SExtInst) (*ir.InstSExt, error) {
	i, ok := inst.(*ir.InstSExt)
	if !ok {
		panic(fmt.Errorf("invalid IR instruction for AST instruction; expected *ir.InstSExt, got %T", inst))
	}
	// Value before conversion.
	from, err := fgen.astToIRTypeValue(old.From())
	if err != nil {
		return nil, errors.WithStack(err)
	}
	i.From = from
	// Type after conversion.
	to, err := fgen.gen.irType(old.To())
	if err != nil {
		return nil, errors.WithStack(err)
	}
	i.To = to
	// (optional) Metadata.
	md, err := fgen.gen.irMetadataAttachments(old.Metadata())
	if err != nil {
		return nil, errors.WithStack(err)
	}
	i.Metadata = md
	return i, nil
}

// ~~~ [ fptrunc ] ~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~

// astToIRInstFPTrunc translates the given AST fptrunc instruction into an
// equivalent IR instruction.
func (fgen *funcGen) astToIRInstFPTrunc(inst ir.Instruction, old *ast.FPTruncInst) (*ir.InstFPTrunc, error) {
	i, ok := inst.(*ir.InstFPTrunc)
	if !ok {
		panic(fmt.Errorf("invalid IR instruction for AST instruction; expected *ir.InstFPTrunc, got %T", inst))
	}
	// Value before conversion.
	from, err := fgen.astToIRTypeValue(old.From())
	if err != nil {
		return nil, errors.WithStack(err)
	}
	i.From = from
	// Type after conversion.
	to, err := fgen.gen.irType(old.To())
	if err != nil {
		return nil, errors.WithStack(err)
	}
	i.To = to
	// (optional) Metadata.
	md, err := fgen.gen.irMetadataAttachments(old.Metadata())
	if err != nil {
		return nil, errors.WithStack(err)
	}
	i.Metadata = md
	return i, nil
}

// ~~~ [ fpext ] ~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~

// astToIRInstFPExt translates the given AST fpext instruction into an
// equivalent IR instruction.
func (fgen *funcGen) astToIRInstFPExt(inst ir.Instruction, old *ast.FPExtInst) (*ir.InstFPExt, error) {
	i, ok := inst.(*ir.InstFPExt)
	if !ok {
		panic(fmt.Errorf("invalid IR instruction for AST instruction; expected *ir.InstFPExt, got %T", inst))
	}
	// Value before conversion.
	from, err := fgen.astToIRTypeValue(old.From())
	if err != nil {
		return nil, errors.WithStack(err)
	}
	i.From = from
	// Type after conversion.
	to, err := fgen.gen.irType(old.To())
	if err != nil {
		return nil, errors.WithStack(err)
	}
	i.To = to
	// (optional) Metadata.
	md, err := fgen.gen.irMetadataAttachments(old.Metadata())
	if err != nil {
		return nil, errors.WithStack(err)
	}
	i.Metadata = md
	return i, nil
}

// ~~~ [ fptoui ] ~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~

// astToIRInstFPToUI translates the given AST fptoui instruction into an
// equivalent IR instruction.
func (fgen *funcGen) astToIRInstFPToUI(inst ir.Instruction, old *ast.FPToUIInst) (*ir.InstFPToUI, error) {
	i, ok := inst.(*ir.InstFPToUI)
	if !ok {
		panic(fmt.Errorf("invalid IR instruction for AST instruction; expected *ir.InstFPToUI, got %T", inst))
	}
	// Value before conversion.
	from, err := fgen.astToIRTypeValue(old.From())
	if err != nil {
		return nil, errors.WithStack(err)
	}
	i.From = from
	// Type after conversion.
	to, err := fgen.gen.irType(old.To())
	if err != nil {
		return nil, errors.WithStack(err)
	}
	i.To = to
	// (optional) Metadata.
	md, err := fgen.gen.irMetadataAttachments(old.Metadata())
	if err != nil {
		return nil, errors.WithStack(err)
	}
	i.Metadata = md
	return i, nil
}

// ~~~ [ fptosi ] ~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~

// astToIRInstFPToSI translates the given AST fptosi instruction into an
// equivalent IR instruction.
func (fgen *funcGen) astToIRInstFPToSI(inst ir.Instruction, old *ast.FPToSIInst) (*ir.InstFPToSI, error) {
	i, ok := inst.(*ir.InstFPToSI)
	if !ok {
		panic(fmt.Errorf("invalid IR instruction for AST instruction; expected *ir.InstFPToSI, got %T", inst))
	}
	// Value before conversion.
	from, err := fgen.astToIRTypeValue(old.From())
	if err != nil {
		return nil, errors.WithStack(err)
	}
	i.From = from
	// Type after conversion.
	to, err := fgen.gen.irType(old.To())
	if err != nil {
		return nil, errors.WithStack(err)
	}
	i.To = to
	// (optional) Metadata.
	md, err := fgen.gen.irMetadataAttachments(old.Metadata())
	if err != nil {
		return nil, errors.WithStack(err)
	}
	i.Metadata = md
	return i, nil
}

// ~~~ [ uitofp ] ~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~

// astToIRInstUIToFP translates the given AST uitofp instruction into an
// equivalent IR instruction.
func (fgen *funcGen) astToIRInstUIToFP(inst ir.Instruction, old *ast.UIToFPInst) (*ir.InstUIToFP, error) {
	i, ok := inst.(*ir.InstUIToFP)
	if !ok {
		panic(fmt.Errorf("invalid IR instruction for AST instruction; expected *ir.InstUIToFP, got %T", inst))
	}
	// Value before conversion.
	from, err := fgen.astToIRTypeValue(old.From())
	if err != nil {
		return nil, errors.WithStack(err)
	}
	i.From = from
	// Type after conversion.
	to, err := fgen.gen.irType(old.To())
	if err != nil {
		return nil, errors.WithStack(err)
	}
	i.To = to
	// (optional) Metadata.
	md, err := fgen.gen.irMetadataAttachments(old.Metadata())
	if err != nil {
		return nil, errors.WithStack(err)
	}
	i.Metadata = md
	return i, nil
}

// ~~~ [ sitofp ] ~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~

// astToIRInstSIToFP translates the given AST sitofp instruction into an
// equivalent IR instruction.
func (fgen *funcGen) astToIRInstSIToFP(inst ir.Instruction, old *ast.SIToFPInst) (*ir.InstSIToFP, error) {
	i, ok := inst.(*ir.InstSIToFP)
	if !ok {
		panic(fmt.Errorf("invalid IR instruction for AST instruction; expected *ir.InstSIToFP, got %T", inst))
	}
	// Value before conversion.
	from, err := fgen.astToIRTypeValue(old.From())
	if err != nil {
		return nil, errors.WithStack(err)
	}
	i.From = from
	// Type after conversion.
	to, err := fgen.gen.irType(old.To())
	if err != nil {
		return nil, errors.WithStack(err)
	}
	i.To = to
	// (optional) Metadata.
	md, err := fgen.gen.irMetadataAttachments(old.Metadata())
	if err != nil {
		return nil, errors.WithStack(err)
	}
	i.Metadata = md
	return i, nil
}

// ~~~ [ ptrtoint ] ~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~

// astToIRInstPtrToInt translates the given AST ptrtoint instruction into an
// equivalent IR instruction.
func (fgen *funcGen) astToIRInstPtrToInt(inst ir.Instruction, old *ast.PtrToIntInst) (*ir.InstPtrToInt, error) {
	i, ok := inst.(*ir.InstPtrToInt)
	if !ok {
		panic(fmt.Errorf("invalid IR instruction for AST instruction; expected *ir.InstPtrToInt, got %T", inst))
	}
	// Value before conversion.
	from, err := fgen.astToIRTypeValue(old.From())
	if err != nil {
		return nil, errors.WithStack(err)
	}
	i.From = from
	// Type after conversion.
	to, err := fgen.gen.irType(old.To())
	if err != nil {
		return nil, errors.WithStack(err)
	}
	i.To = to
	// (optional) Metadata.
	md, err := fgen.gen.irMetadataAttachments(old.Metadata())
	if err != nil {
		return nil, errors.WithStack(err)
	}
	i.Metadata = md
	return i, nil
}

// ~~~ [ inttoptr ] ~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~

// astToIRInstIntToPtr translates the given AST inttoptr instruction into an
// equivalent IR instruction.
func (fgen *funcGen) astToIRInstIntToPtr(inst ir.Instruction, old *ast.IntToPtrInst) (*ir.InstIntToPtr, error) {
	i, ok := inst.(*ir.InstIntToPtr)
	if !ok {
		panic(fmt.Errorf("invalid IR instruction for AST instruction; expected *ir.InstIntToPtr, got %T", inst))
	}
	// Value before conversion.
	from, err := fgen.astToIRTypeValue(old.From())
	if err != nil {
		return nil, errors.WithStack(err)
	}
	i.From = from
	// Type after conversion.
	to, err := fgen.gen.irType(old.To())
	if err != nil {
		return nil, errors.WithStack(err)
	}
	i.To = to
	// (optional) Metadata.
	md, err := fgen.gen.irMetadataAttachments(old.Metadata())
	if err != nil {
		return nil, errors.WithStack(err)
	}
	i.Metadata = md
	return i, nil
}

// ~~~ [ bitcast ] ~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~

// astToIRInstBitCast translates the given AST bitcast instruction into an
// equivalent IR instruction.
func (fgen *funcGen) astToIRInstBitCast(inst ir.Instruction, old *ast.BitCastInst) (*ir.InstBitCast, error) {
	i, ok := inst.(*ir.InstBitCast)
	if !ok {
		panic(fmt.Errorf("invalid IR instruction for AST instruction; expected *ir.InstBitCast, got %T", inst))
	}
	// Value before conversion.
	from, err := fgen.astToIRTypeValue(old.From())
	if err != nil {
		return nil, errors.WithStack(err)
	}
	i.From = from
	// Type after conversion.
	to, err := fgen.gen.irType(old.To())
	if err != nil {
		return nil, errors.WithStack(err)
	}
	i.To = to
	// (optional) Metadata.
	md, err := fgen.gen.irMetadataAttachments(old.Metadata())
	if err != nil {
		return nil, errors.WithStack(err)
	}
	i.Metadata = md
	return i, nil
}

// ~~~ [ addrspacecast ] ~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~

// astToIRInstAddrSpaceCast translates the given AST addrspacecast instruction
// into an equivalent IR instruction.
func (fgen *funcGen) astToIRInstAddrSpaceCast(inst ir.Instruction, old *ast.AddrSpaceCastInst) (*ir.InstAddrSpaceCast, error) {
	i, ok := inst.(*ir.InstAddrSpaceCast)
	if !ok {
		panic(fmt.Errorf("invalid IR instruction for AST instruction; expected *ir.InstAddrSpaceCast, got %T", inst))
	}
	// Value before conversion.
	from, err := fgen.astToIRTypeValue(old.From())
	if err != nil {
		return nil, errors.WithStack(err)
	}
	i.From = from
	// Type after conversion.
	to, err := fgen.gen.irType(old.To())
	if err != nil {
		return nil, errors.WithStack(err)
	}
	i.To = to
	// (optional) Metadata.
	md, err := fgen.gen.irMetadataAttachments(old.Metadata())
	if err != nil {
		return nil, errors.WithStack(err)
	}
	i.Metadata = md
	return i, nil
}
