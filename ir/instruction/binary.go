package instruction

import (
	"fmt"

	"github.com/llir/llvm/ir/types"
	"github.com/llir/llvm/ir/value"
	"github.com/mewkiz/pkg/errutil"
)

// === [ Binary instructions ] =================================================
//
// References:
//    http://llvm.org/docs/LangRef.html#binary-operations

// --- [ `add` ] ---------------------------------------------------------------

// Add represents an addition instruction.
//
// References:
//    http://llvm.org/docs/LangRef.html#add-instruction
type Add struct {
	// Operands.
	x, y value.Value
}

// NewAdd returns a new add instruction based on the given operands.
func NewAdd(x, y value.Value) (*Add, error) {
	if !types.Equal(x.Type(), y.Type()) {
		return nil, errutil.Newf("type mismatch between x (%v) and y (%v)", x.Type(), y.Type())
	}
	return &Add{x: x, y: y}, nil
}

// Type returns the type of the value produced by the instruction.
func (inst *Add) Type() types.Type {
	return inst.x.Type()
}

// String returns the string representation of the instruction.
func (inst *Add) String() string {
	return fmt.Sprintf("add %v %v, %v", inst.x.Type(), inst.x, inst.y)
}

// --- [ `fadd` ] --------------------------------------------------------------

// FAdd represents a floating-point addition instruction.
//
// References:
//    http://llvm.org/docs/LangRef.html#fadd-instruction
type FAdd struct {
	// Operands.
	x, y value.Value
}

// NewFAdd returns a new fadd instruction based on the given operands.
func NewFAdd(x, y value.Value) (*FAdd, error) {
	if !types.Equal(x.Type(), y.Type()) {
		return nil, errutil.Newf("type mismatch between x (%v) and y (%v)", x.Type(), y.Type())
	}
	return &FAdd{x: x, y: y}, nil
}

// Type returns the type of the value produced by the instruction.
func (inst *FAdd) Type() types.Type {
	return inst.x.Type()
}

// String returns the string representation of the instruction.
func (inst *FAdd) String() string {
	return fmt.Sprintf("fadd %v %v, %v", inst.x.Type(), inst.x, inst.y)
}

// --- [ `sub` ] ---------------------------------------------------------------

// Sub represents a subtraction instruction.
//
// References:
//    http://llvm.org/docs/LangRef.html#sub-instruction
type Sub struct {
	// Operands.
	x, y value.Value
}

// NewSub returns a new sub instruction based on the given operands.
func NewSub(x, y value.Value) (*Sub, error) {
	if !types.Equal(x.Type(), y.Type()) {
		return nil, errutil.Newf("type mismatch between x (%v) and y (%v)", x.Type(), y.Type())
	}
	return &Sub{x: x, y: y}, nil
}

// Type returns the type of the value produced by the instruction.
func (inst *Sub) Type() types.Type {
	return inst.x.Type()
}

// String returns the string representation of the instruction.
func (inst *Sub) String() string {
	return fmt.Sprintf("sub %v %v, %v", inst.x.Type(), inst.x, inst.y)
}

// --- [ `fsub` ] --------------------------------------------------------------

// FSub represents a floating-point subtraction instruction.
//
// References:
//    http://llvm.org/docs/LangRef.html#fsub-instruction
type FSub struct {
	// Operands.
	x, y value.Value
}

// NewFSub returns a new fsub instruction based on the given operands.
func NewFSub(x, y value.Value) (*FSub, error) {
	if !types.Equal(x.Type(), y.Type()) {
		return nil, errutil.Newf("type mismatch between x (%v) and y (%v)", x.Type(), y.Type())
	}
	return &FSub{x: x, y: y}, nil
}

// Type returns the type of the value produced by the instruction.
func (inst *FSub) Type() types.Type {
	return inst.x.Type()
}

// String returns the string representation of the instruction.
func (inst *FSub) String() string {
	return fmt.Sprintf("fsub %v %v, %v", inst.x.Type(), inst.x, inst.y)
}

// --- [ `mul` ] ---------------------------------------------------------------

// Mul represents a multiplication instruction.
//
// References:
//    http://llvm.org/docs/LangRef.html#mul-instruction
type Mul struct {
	// Operands.
	x, y value.Value
}

// NewMul returns a new mul instruction based on the given operands.
func NewMul(x, y value.Value) (*Mul, error) {
	if !types.Equal(x.Type(), y.Type()) {
		return nil, errutil.Newf("type mismatch between x (%v) and y (%v)", x.Type(), y.Type())
	}
	return &Mul{x: x, y: y}, nil
}

// Type returns the type of the value produced by the instruction.
func (inst *Mul) Type() types.Type {
	return inst.x.Type()
}

// String returns the string representation of the instruction.
func (inst *Mul) String() string {
	return fmt.Sprintf("mul %v %v, %v", inst.x.Type(), inst.x, inst.y)
}

// --- [ `fmul` ] --------------------------------------------------------------

// FMul represents a floating-point multiplication instruction.
//
// References:
//    http://llvm.org/docs/LangRef.html#fmul-instruction
type FMul struct {
	// Operands.
	x, y value.Value
}

// NewFMul returns a new fmul instruction based on the given operands.
func NewFMul(x, y value.Value) (*FMul, error) {
	if !types.Equal(x.Type(), y.Type()) {
		return nil, errutil.Newf("type mismatch between x (%v) and y (%v)", x.Type(), y.Type())
	}
	return &FMul{x: x, y: y}, nil
}

// Type returns the type of the value produced by the instruction.
func (inst *FMul) Type() types.Type {
	return inst.x.Type()
}

// String returns the string representation of the instruction.
func (inst *FMul) String() string {
	return fmt.Sprintf("fmul %v %v, %v", inst.x.Type(), inst.x, inst.y)
}

// --- [ `udiv` ] --------------------------------------------------------------

// UDiv represents an unsigned division instruction.
//
// References:
//    http://llvm.org/docs/LangRef.html#udiv-instruction
type UDiv struct {
	// Operands.
	x, y value.Value
}

// NewUDiv returns a new udiv instruction based on the given operands.
func NewUDiv(x, y value.Value) (*UDiv, error) {
	if !types.Equal(x.Type(), y.Type()) {
		return nil, errutil.Newf("type mismatch between x (%v) and y (%v)", x.Type(), y.Type())
	}
	return &UDiv{x: x, y: y}, nil
}

// Type returns the type of the value produced by the instruction.
func (inst *UDiv) Type() types.Type {
	return inst.x.Type()
}

// String returns the string representation of the instruction.
func (inst *UDiv) String() string {
	return fmt.Sprintf("udiv %v %v, %v", inst.x.Type(), inst.x, inst.y)
}

// --- [ `sdiv` ] --------------------------------------------------------------

// SDiv represents a signed division instruction.
//
// References:
//    http://llvm.org/docs/LangRef.html#sdiv-instruction
type SDiv struct {
	// Operands.
	x, y value.Value
}

// NewSDiv returns a new sdiv instruction based on the given operands.
func NewSDiv(x, y value.Value) (*SDiv, error) {
	if !types.Equal(x.Type(), y.Type()) {
		return nil, errutil.Newf("type mismatch between x (%v) and y (%v)", x.Type(), y.Type())
	}
	return &SDiv{x: x, y: y}, nil
}

// Type returns the type of the value produced by the instruction.
func (inst *SDiv) Type() types.Type {
	return inst.x.Type()
}

// String returns the string representation of the instruction.
func (inst *SDiv) String() string {
	return fmt.Sprintf("sdiv %v %v, %v", inst.x.Type(), inst.x, inst.y)
}

// --- [ `fdiv` ] --------------------------------------------------------------

// FDiv represents a floating-point division instruction.
//
// References:
//    http://llvm.org/docs/LangRef.html#fdiv-instruction
type FDiv struct {
	// Operands.
	x, y value.Value
}

// NewFDiv returns a new fdiv instruction based on the given operands.
func NewFDiv(x, y value.Value) (*FDiv, error) {
	if !types.Equal(x.Type(), y.Type()) {
		return nil, errutil.Newf("type mismatch between x (%v) and y (%v)", x.Type(), y.Type())
	}
	return &FDiv{x: x, y: y}, nil
}

// Type returns the type of the value produced by the instruction.
func (inst *FDiv) Type() types.Type {
	return inst.x.Type()
}

// String returns the string representation of the instruction.
func (inst *FDiv) String() string {
	return fmt.Sprintf("fdiv %v %v, %v", inst.x.Type(), inst.x, inst.y)
}

// --- [ `urem` ] --------------------------------------------------------------

// URem represents an unsigned remainder instruction.
//
// References:
//    http://llvm.org/docs/LangRef.html#urem-instruction
type URem struct {
	// Operands.
	x, y value.Value
}

// NewURem returns a new urem instruction based on the given operands.
func NewURem(x, y value.Value) (*URem, error) {
	if !types.Equal(x.Type(), y.Type()) {
		return nil, errutil.Newf("type mismatch between x (%v) and y (%v)", x.Type(), y.Type())
	}
	return &URem{x: x, y: y}, nil
}

// Type returns the type of the value produced by the instruction.
func (inst *URem) Type() types.Type {
	return inst.x.Type()
}

// String returns the string representation of the instruction.
func (inst *URem) String() string {
	return fmt.Sprintf("urem %v %v, %v", inst.x.Type(), inst.x, inst.y)
}

// --- [ `srem` ] --------------------------------------------------------------

// SRem represents a signed remainder instruction.
//
// References:
//    http://llvm.org/docs/LangRef.html#srem-instruction
type SRem struct {
	// Operands.
	x, y value.Value
}

// NewSRem returns a new srem instruction based on the given operands.
func NewSRem(x, y value.Value) (*SRem, error) {
	if !types.Equal(x.Type(), y.Type()) {
		return nil, errutil.Newf("type mismatch between x (%v) and y (%v)", x.Type(), y.Type())
	}
	return &SRem{x: x, y: y}, nil
}

// Type returns the type of the value produced by the instruction.
func (inst *SRem) Type() types.Type {
	return inst.x.Type()
}

// String returns the string representation of the instruction.
func (inst *SRem) String() string {
	return fmt.Sprintf("srem %v %v, %v", inst.x.Type(), inst.x, inst.y)
}

// --- [ `frem` ] --------------------------------------------------------------

// FRem represents a floating-point remainder instruction.
//
// References:
//    http://llvm.org/docs/LangRef.html#frem-instruction
type FRem struct {
	// Operands.
	x, y value.Value
}

// NewFRem returns a new frem instruction based on the given operands.
func NewFRem(x, y value.Value) (*FRem, error) {
	if !types.Equal(x.Type(), y.Type()) {
		return nil, errutil.Newf("type mismatch between x (%v) and y (%v)", x.Type(), y.Type())
	}
	return &FRem{x: x, y: y}, nil
}

// Type returns the type of the value produced by the instruction.
func (inst *FRem) Type() types.Type {
	return inst.x.Type()
}

// String returns the string representation of the instruction.
func (inst *FRem) String() string {
	return fmt.Sprintf("frem %v %v, %v", inst.x.Type(), inst.x, inst.y)
}

// isValueInst ensures that only instructions which return values can be
// assigned to the Value interface.
func (*Add) isValueInst()  {}
func (*FAdd) isValueInst() {}
func (*Sub) isValueInst()  {}
func (*FSub) isValueInst() {}
func (*Mul) isValueInst()  {}
func (*FMul) isValueInst() {}
func (*UDiv) isValueInst() {}
func (*SDiv) isValueInst() {}
func (*FDiv) isValueInst() {}
func (*URem) isValueInst() {}
func (*SRem) isValueInst() {}
func (*FRem) isValueInst() {}
