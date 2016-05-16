package instruction

import (
	"fmt"

	"github.com/llir/llvm/ir/types"
	"github.com/llir/llvm/ir/value"
	"github.com/mewkiz/pkg/errutil"
)

// References:
//    http://llvm.org/docs/LangRef.html#bitwise-binary-operations

// TODO: Add support for the remaining bitwise binary operations:
//    http://llvm.org/docs/LangRef.html#shl-instruction
//    http://llvm.org/docs/LangRef.html#lshr-instruction
//    http://llvm.org/docs/LangRef.html#ashr-instruction
//    http://llvm.org/docs/LangRef.html#and-instruction
//    http://llvm.org/docs/LangRef.html#or-instruction
//    http://llvm.org/docs/LangRef.html#xor-instruction

type Shl struct{}

func (*Shl) Type() types.Type { panic("Shl.Type: not yet implemented") }
func (*Shl) String() string   { panic("Shl.String: not yet implemented") }

type LShr struct{}

func (*LShr) Type() types.Type { panic("LShr.Type: not yet implemented") }
func (*LShr) String() string   { panic("LShr.String: not yet implemented") }

type AShr struct{}

func (*AShr) Type() types.Type { panic("AShr.Type: not yet implemented") }
func (*AShr) String() string   { panic("AShr.String: not yet implemented") }

type And struct{}

func (*And) Type() types.Type { panic("And.Type: not yet implemented") }
func (*And) String() string   { panic("And.String: not yet implemented") }

type Or struct{}

func (*Or) Type() types.Type { panic("Or.Type: not yet implemented") }
func (*Or) String() string   { panic("Or.String: not yet implemented") }

// Xor represents an exclusive-OR instruction.
type Xor struct {
	// Operand type.
	typ types.Type
	// Operands.
	x, y value.Value
}

// NewXor returns a new xor instruction based on the given operand type and
// values.
func NewXor(typ types.Type, x, y value.Value) (*Xor, error) {
	// Sanity check.
	if !types.Equal(typ, x.Type()) {
		return nil, errutil.Newf("type mismatch between operand type %q and type of x %q", typ, x.Type())
	}
	if !types.Equal(typ, y.Type()) {
		return nil, errutil.Newf("type mismatch between operand type %q and type of y %q", typ, y.Type())
	}
	return &Xor{typ: typ, x: x, y: y}, nil
}

// Type returns the type of the value produced by the instruction.
func (inst *Xor) Type() types.Type {
	return inst.typ
}

// String returns the string representation of the instruction.
func (inst *Xor) String() string {
	return fmt.Sprintf("xor %v %v, %v", inst.typ, inst.x, inst.y)
}

// isInst ensures that only non-branching instructions can be assigned to the
// Instruction interface.
func (*Shl) isInst()  {}
func (*LShr) isInst() {}
func (*AShr) isInst() {}
func (*And) isInst()  {}
func (*Or) isInst()   {}
func (*Xor) isInst()  {}
