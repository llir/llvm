package instruction

import (
	"fmt"

	"github.com/llir/llvm/ir/types"
	"github.com/llir/llvm/ir/value"
	"github.com/mewkiz/pkg/errutil"
)

// TODO: Consider removing the type from bitwise and binary instructions, as it
// may be validated against the type of the operands when created by the
// constructed, and later the type of the operand may simply be the type of the
// first operand (which is equal to the type of the second operand).

// References:
//    http://llvm.org/docs/LangRef.html#bitwise-binary-operations

// ShL represents an shift left instruction.
//
// References:
//    http://llvm.org/docs/LangRef.html#shl-instruction
type ShL struct {
	// Operand type.
	typ types.Type
	// Operands.
	x, y value.Value
}

// NewShL returns a new shl instruction based on the given operand type and
// values.
func NewShL(typ types.Type, x, y value.Value) (*ShL, error) {
	// Sanity check.
	if !types.Equal(typ, x.Type()) {
		return nil, errutil.Newf("type mismatch between operand type %q and type of x %q", typ, x.Type())
	}
	if !types.Equal(typ, y.Type()) {
		return nil, errutil.Newf("type mismatch between operand type %q and type of y %q", typ, y.Type())
	}
	return &ShL{typ: typ, x: x, y: y}, nil
}

// Type returns the type of the value produced by the instruction.
func (inst *ShL) Type() types.Type {
	return inst.typ
}

// String returns the string representation of the instruction.
func (inst *ShL) String() string {
	return fmt.Sprintf("shl %v %v, %v", inst.typ, inst.x, inst.y)
}

// LShR represents a logical shift right instruction.
//
// References:
//    http://llvm.org/docs/LangRef.html#lshr-instruction
type LShR struct {
	// Operand type.
	typ types.Type
	// Operands.
	x, y value.Value
}

// NewLShR returns a new lshr instruction based on the given operand type and
// values.
func NewLShR(typ types.Type, x, y value.Value) (*LShR, error) {
	// Sanity check.
	if !types.Equal(typ, x.Type()) {
		return nil, errutil.Newf("type mismatch between operand type %q and type of x %q", typ, x.Type())
	}
	if !types.Equal(typ, y.Type()) {
		return nil, errutil.Newf("type mismatch between operand type %q and type of y %q", typ, y.Type())
	}
	return &LShR{typ: typ, x: x, y: y}, nil
}

// Type returns the type of the value produced by the instruction.
func (inst *LShR) Type() types.Type {
	return inst.typ
}

// String returns the string representation of the instruction.
func (inst *LShR) String() string {
	return fmt.Sprintf("lshr %v %v, %v", inst.typ, inst.x, inst.y)
}

// AShR represents an arithmetic shift right instruction.
//
// References:
//    http://llvm.org/docs/LangRef.html#ashr-instruction
type AShR struct {
	// Operand type.
	typ types.Type
	// Operands.
	x, y value.Value
}

// NewAShR returns a new ashr instruction based on the given operand type and
// values.
func NewAShR(typ types.Type, x, y value.Value) (*AShR, error) {
	// Sanity check.
	if !types.Equal(typ, x.Type()) {
		return nil, errutil.Newf("type mismatch between operand type %q and type of x %q", typ, x.Type())
	}
	if !types.Equal(typ, y.Type()) {
		return nil, errutil.Newf("type mismatch between operand type %q and type of y %q", typ, y.Type())
	}
	return &AShR{typ: typ, x: x, y: y}, nil
}

// Type returns the type of the value produced by the instruction.
func (inst *AShR) Type() types.Type {
	return inst.typ
}

// String returns the string representation of the instruction.
func (inst *AShR) String() string {
	return fmt.Sprintf("ashr %v %v, %v", inst.typ, inst.x, inst.y)
}

// And represents an AND instruction.
//
// References:
//    http://llvm.org/docs/LangRef.html#and-instruction
type And struct {
	// Operand type.
	typ types.Type
	// Operands.
	x, y value.Value
}

// NewAnd returns a new and instruction based on the given operand type and
// values.
func NewAnd(typ types.Type, x, y value.Value) (*And, error) {
	// Sanity check.
	if !types.Equal(typ, x.Type()) {
		return nil, errutil.Newf("type mismatch between operand type %q and type of x %q", typ, x.Type())
	}
	if !types.Equal(typ, y.Type()) {
		return nil, errutil.Newf("type mismatch between operand type %q and type of y %q", typ, y.Type())
	}
	return &And{typ: typ, x: x, y: y}, nil
}

// Type returns the type of the value produced by the instruction.
func (inst *And) Type() types.Type {
	return inst.typ
}

// String returns the string representation of the instruction.
func (inst *And) String() string {
	return fmt.Sprintf("and %v %v, %v", inst.typ, inst.x, inst.y)
}

// Or represents an OR instruction.
//
// References:
//    http://llvm.org/docs/LangRef.html#or-instruction
type Or struct {
	// Operand type.
	typ types.Type
	// Operands.
	x, y value.Value
}

// NewOr returns a new or instruction based on the given operand type and
// values.
func NewOr(typ types.Type, x, y value.Value) (*Or, error) {
	// Sanity check.
	if !types.Equal(typ, x.Type()) {
		return nil, errutil.Newf("type mismatch between operand type %q and type of x %q", typ, x.Type())
	}
	if !types.Equal(typ, y.Type()) {
		return nil, errutil.Newf("type mismatch between operand type %q and type of y %q", typ, y.Type())
	}
	return &Or{typ: typ, x: x, y: y}, nil
}

// Type returns the type of the value produced by the instruction.
func (inst *Or) Type() types.Type {
	return inst.typ
}

// String returns the string representation of the instruction.
func (inst *Or) String() string {
	return fmt.Sprintf("or %v %v, %v", inst.typ, inst.x, inst.y)
}

// Xor represents an exclusive-OR instruction.
//
// References:
//    http://llvm.org/docs/LangRef.html#xor-instruction
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

// isValueInst ensures that only instructions which return values can be
// assigned to the Value interface.
func (*ShL) isValueInst()  {}
func (*LShR) isValueInst() {}
func (*AShR) isValueInst() {}
func (*And) isValueInst()  {}
func (*Or) isValueInst()   {}
func (*Xor) isValueInst()  {}
