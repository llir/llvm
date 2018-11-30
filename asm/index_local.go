package asm

import (
	"github.com/llir/ll/ast"
	"github.com/llir/llvm/ir"
	"github.com/llir/llvm/ir/types"
	"github.com/llir/llvm/ir/value"
	"github.com/pkg/errors"
)

// local is a local variable.
type local interface {
	value.Named
	// ID returns the ID of the local identifier.
	ID() int64
	// SetID sets the ID of the local identifier.
	SetID(id int64)
	// IsUnnamed reports whether the local identifier is unnamed.
	IsUnnamed() bool
}

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
		if err := fgen.addLocal(param.LocalIdent, param); err != nil {
			return errors.WithStack(err)
		}
	}
	for _, block := range f.Blocks {
		if err := fgen.addLocal(block.LocalIdent, block); err != nil {
			return errors.WithStack(err)
		}
		for _, inst := range block.Insts {
			if n, ok := inst.(local); ok {
				if isVoidValue(n) {
					continue
				}
				ident := ir.LocalIdent{LocalName: n.Name(), LocalID: n.ID()}
				if err := fgen.addLocal(ident, n); err != nil {
					return errors.WithStack(err)
				}
			}
		}
		if n, ok := block.Term.(local); ok {
			if isVoidValue(n) {
				continue
			}
			ident := ir.LocalIdent{LocalName: n.Name(), LocalID: n.ID()}
			if err := fgen.addLocal(ident, n); err != nil {
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
		block := &ir.BasicBlock{}
		if n, ok := oldBlock.Name(); ok {
			block.LocalIdent = labelIdent(n)
		}
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
		block.Parent = f
		f.Blocks = append(f.Blocks, block)
	}
	return nil
}

// ### [ Helper functions ] ####################################################

// addLocal adds the local variable with the given name to the map of local
// variables of the function.
func (fgen *funcGen) addLocal(ident ir.LocalIdent, v value.Value) error {
	if prev, ok := fgen.ls[ident]; ok {
		return errors.Errorf("local identifier %q already present; prev `%s`, new `%s`", ident.Ident(), prev, v)
	}
	fgen.ls[ident] = v
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
