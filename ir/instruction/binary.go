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

type Add struct {
	typ  types.Type
	x, y value.Value
}

func (inst *Add) Type() types.Type {
	return inst.typ
}

// String returns the string representation of the addition instruction.
func (inst *Add) String() string {
	return fmt.Sprintf("add %v %v, %v", inst.typ, inst.x, inst.y)
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

type FAdd struct{}

func (*FAdd) Type() types.Type { panic("FAdd.Type: not yet implemented") }
func (*FAdd) String() string   { panic("FAdd.String: not yet implemented") }

type Sub struct{}

func (*Sub) Type() types.Type { panic("Sub.Type: not yet implemented") }
func (*Sub) String() string   { panic("Sub.String: not yet implemented") }

type FSub struct{}

func (*FSub) Type() types.Type { panic("FSub.Type: not yet implemented") }
func (*FSub) String() string   { panic("FSub.String: not yet implemented") }

type Mul struct{}

func (*Mul) Type() types.Type { panic("Mul.Type: not yet implemented") }
func (*Mul) String() string   { panic("Mul.String: not yet implemented") }

type FMul struct{}

func (*FMul) Type() types.Type { panic("FMul.Type: not yet implemented") }
func (*FMul) String() string   { panic("FMul.String: not yet implemented") }

type UDiv struct{}

func (*UDiv) Type() types.Type { panic("UDiv.Type: not yet implemented") }
func (*UDiv) String() string   { panic("UDiv.String: not yet implemented") }

type SDiv struct{}

func (*SDiv) Type() types.Type { panic("SDiv.Type: not yet implemented") }
func (*SDiv) String() string   { panic("SDiv.String: not yet implemented") }

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
