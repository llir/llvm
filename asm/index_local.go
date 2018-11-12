package asm

import (
	"github.com/llir/ll/ast"
	"github.com/llir/llvm/internal/enc"
	"github.com/llir/llvm/ir"
	"github.com/llir/llvm/ir/types"
	"github.com/llir/llvm/ir/value"
	"github.com/pkg/errors"
)

// indexLocals indexes the function parameters, basic blocks and local variables
// (produced by instructions and terminators) of the given function.
//
// Post-condition: fgen.ls maps from local identifier (without '%' prefix) to
// corresponding skeleton IR value.
func (fgen *funcGen) indexLocals(oldBlocks []ast.BasicBlock) error {
	// Create local variable skeletons (with type and without body).
	if err := fgen.newLocals(oldBlocks); err != nil {
		return errors.WithStack(err)
	}
	// Assign local IDs.
	//
	// Note: We need to store the type of call instructions and invoke
	// terminators before assigning local IDs, since they may be values or
	// non-values based on return type. This is done in newLocals.
	f := fgen.f
	if err := f.AssignIDs(); err != nil {
		return errors.WithStack(err)
	}
	// Index local identifiers.
	for _, param := range f.Params {
		if err := fgen.addLocal(param.LocalName, param); err != nil {
			return errors.WithStack(err)
		}
	}
	for _, block := range f.Blocks {
		if err := fgen.addLocal(block.LocalName, block); err != nil {
			return errors.WithStack(err)
		}
		for _, inst := range block.Insts {
			if n, ok := inst.(value.Named); ok {
				if isVoidValue(n) {
					continue
				}
				if err := fgen.addLocal(n.Name(), n); err != nil {
					return errors.WithStack(err)
				}
			}
		}
		if n, ok := block.Term.(value.Named); ok {
			if isVoidValue(n) {
				continue
			}
			if err := fgen.addLocal(n.Name(), n); err != nil {
				return errors.WithStack(err)
			}
		}
	}
	return nil
}

// newLocals creates IR skeletons (with type and without body) for corresponding
// AST basic blocks, instructions and terminators of the given function.
//
// Post-condition: fgen.f.Blocks is populated with IR skeletons.
func (fgen *funcGen) newLocals(oldBlocks []ast.BasicBlock) error {
	// Note: Function parameters are already translated in astToIRFuncHeader.
	f := fgen.f
	for _, oldBlock := range oldBlocks {
		blockName := optLabelIdent(oldBlock.Name())
		block := ir.NewBlock(blockName)
		for _, oldInst := range oldBlock.Insts() {
			inst, err := fgen.newIRInst(oldInst)
			if err != nil {
				return errors.WithStack(err)
			}
			block.Insts = append(block.Insts, inst)
		}
		term, err := fgen.newIRTerm(oldBlock.Term())
		if err != nil {
			return errors.WithStack(err)
		}
		block.Term = term
		f.Blocks = append(f.Blocks, block)
	}
	return nil
}

// ### [ Helper functions ] ####################################################

// addLocal adds the local variable with the given name to the map of local
// variables of the function.
func (fgen *funcGen) addLocal(name string, v value.Value) error {
	if prev, ok := fgen.ls[name]; ok {
		return errors.Errorf("local identifier %q already present; prev `%s`, new `%s`", enc.Local(name), prev, v)
	}
	fgen.ls[name] = v
	return nil
}

// isVoidValue reports whether the given named value is a non-value (i.e. a call
// instruction or invoke terminator with void-return type).
func isVoidValue(n value.Named) bool {
	switch n.(type) {
	case *ir.InstCall, *ir.TermInvoke:
		return n.Type().Equal(types.Void)
	}
	return false
}
