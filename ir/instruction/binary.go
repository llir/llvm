package instruction

import (
	"fmt"

	"github.com/llir/llvm/ir/types"
	"github.com/llir/llvm/ir/value"
	"github.com/mewkiz/pkg/errutil"
)

// References:
//    http://llvm.org/docs/LangRef.html#binary-operations

// TODO: Add support for the remaining binary operations:
//    http://llvm.org/docs/LangRef.html#add-instruction
//    http://llvm.org/docs/LangRef.html#fadd-instruction
//    http://llvm.org/docs/LangRef.html#sub-instruction
//    http://llvm.org/docs/LangRef.html#fsub-instruction
//    http://llvm.org/docs/LangRef.html#mul-instruction
//    http://llvm.org/docs/LangRef.html#fmul-instruction
//    http://llvm.org/docs/LangRef.html#udiv-instruction
//    http://llvm.org/docs/LangRef.html#sdiv-instruction
//    http://llvm.org/docs/LangRef.html#fdiv-instruction
//    http://llvm.org/docs/LangRef.html#urem-instruction
//    http://llvm.org/docs/LangRef.html#srem-instruction
//    http://llvm.org/docs/LangRef.html#frem-instruction

// Add represents an addition instruction.
type Add struct {
	// Operand type.
	typ types.Type
	// Operands.
	x, y value.Value
}

// NewAdd returns a new add instruction based on the given operand type and
// values.
func NewAdd(typ types.Type, x, y value.Value) (*Add, error) {
	// Sanity check.
	if !types.Equal(typ, x.Type()) {
		return nil, errutil.Newf("type mismatch between operand type %q and type of x %q", typ, x.Type())
	}
	if !types.Equal(typ, y.Type()) {
		return nil, errutil.Newf("type mismatch between operand type %q and type of y %q", typ, y.Type())
	}
	return &Add{typ: typ, x: x, y: y}, nil
}

// Type returns the type of the value produced by the instruction.
func (inst *Add) Type() types.Type {
	return inst.typ
}

// String returns the string representation of the instruction.
func (inst *Add) String() string {
	return fmt.Sprintf("add %v %v, %v", inst.typ, inst.x, inst.y)
}

type FAdd struct{}

func (*FAdd) Type() types.Type { panic("FAdd.Type: not yet implemented") }
func (*FAdd) String() string   { panic("FAdd.String: not yet implemented") }

// Sub represents a subtraction instruction.
type Sub struct {
	// Operand type.
	typ types.Type
	// Operands.
	x, y value.Value
}

// NewSub returns a new sub instruction based on the given operand type and
// values.
func NewSub(typ types.Type, x, y value.Value) (*Sub, error) {
	// Sanity check.
	if !types.Equal(typ, x.Type()) {
		return nil, errutil.Newf("type mismatch between operand type %q and type of x %q", typ, x.Type())
	}
	if !types.Equal(typ, y.Type()) {
		return nil, errutil.Newf("type mismatch between operand type %q and type of y %q", typ, y.Type())
	}
	return &Sub{typ: typ, x: x, y: y}, nil
}

// Type returns the type of the value produced by the instruction.
func (inst *Sub) Type() types.Type {
	return inst.typ
}

// String returns the string representation of the instruction.
func (inst *Sub) String() string {
	return fmt.Sprintf("sub %v %v, %v", inst.typ, inst.x, inst.y)
}

type FSub struct{}

func (*FSub) Type() types.Type { panic("FSub.Type: not yet implemented") }
func (*FSub) String() string   { panic("FSub.String: not yet implemented") }

// Mul represents a multiplication instruction.
type Mul struct {
	// Operand type.
	typ types.Type
	// Operands.
	x, y value.Value
}

// NewMul returns a new mul instruction based on the given operand type and
// values.
func NewMul(typ types.Type, x, y value.Value) (*Mul, error) {
	// Sanity check.
	if !types.Equal(typ, x.Type()) {
		return nil, errutil.Newf("type mismatch between operand type %q and type of x %q", typ, x.Type())
	}
	if !types.Equal(typ, y.Type()) {
		return nil, errutil.Newf("type mismatch between operand type %q and type of y %q", typ, y.Type())
	}
	return &Mul{typ: typ, x: x, y: y}, nil
}

// Type returns the type of the value produced by the instruction.
func (inst *Mul) Type() types.Type {
	return inst.typ
}

// String returns the string representation of the instruction.
func (inst *Mul) String() string {
	return fmt.Sprintf("mul %v %v, %v", inst.typ, inst.x, inst.y)
}

type FMul struct{}

func (*FMul) Type() types.Type { panic("FMul.Type: not yet implemented") }
func (*FMul) String() string   { panic("FMul.String: not yet implemented") }

type UDiv struct{}

func (*UDiv) Type() types.Type { panic("UDiv.Type: not yet implemented") }
func (*UDiv) String() string   { panic("UDiv.String: not yet implemented") }

// SDiv represents a signed division instruction.
type SDiv struct {
	// Operand type.
	typ types.Type
	// Operands.
	x, y value.Value
}

// NewSDiv returns a new sdiv instruction based on the given operand type and
// values.
func NewSDiv(typ types.Type, x, y value.Value) (*SDiv, error) {
	// Sanity check.
	if !types.Equal(typ, x.Type()) {
		return nil, errutil.Newf("type mismatch between operand type %q and type of x %q", typ, x.Type())
	}
	if !types.Equal(typ, y.Type()) {
		return nil, errutil.Newf("type mismatch between operand type %q and type of y %q", typ, y.Type())
	}
	return &SDiv{typ: typ, x: x, y: y}, nil
}

// Type returns the type of the value produced by the instruction.
func (inst *SDiv) Type() types.Type {
	return inst.typ
}

// String returns the string representation of the instruction.
func (inst *SDiv) String() string {
	return fmt.Sprintf("sdiv %v %v, %v", inst.typ, inst.x, inst.y)
}

type FDiv struct{}

func (*FDiv) Type() types.Type { panic("FDiv.Type: not yet implemented") }
func (*FDiv) String() string   { panic("FDiv.String: not yet implemented") }

type URem struct{}

func (*URem) Type() types.Type { panic("URem.Type: not yet implemented") }
func (*URem) String() string   { panic("URem.String: not yet implemented") }

type SRem struct{}

func (*SRem) Type() types.Type { panic("SRem.Type: not yet implemented") }
func (*SRem) String() string   { panic("SRem.String: not yet implemented") }

type FRem struct{}

func (*FRem) Type() types.Type { panic("FRem.Type: not yet implemented") }
func (*FRem) String() string   { panic("FRem.String: not yet implemented") }

// isInst ensures that only non-branching instructions can be assigned to the
// Instruction interface.
func (*Add) isInst()  {}
func (*FAdd) isInst() {}
func (*Sub) isInst()  {}
func (*FSub) isInst() {}
func (*Mul) isInst()  {}
func (*FMul) isInst() {}
func (*UDiv) isInst() {}
func (*SDiv) isInst() {}
func (*FDiv) isInst() {}
func (*URem) isInst() {}
func (*SRem) isInst() {}
func (*FRem) isInst() {}
