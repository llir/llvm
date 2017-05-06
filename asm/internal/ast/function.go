package ast

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/llir/llvm/internal/enc"
)

// A Function represents an LLVM IR function definition or external function
// declaration. The body of a function definition consists of a set of basic
// blocks, interconnected by control flow instructions.
type Function struct {
	// Function name.
	Name string
	// Function signature.
	Sig *FuncType
	// Calling convention.
	CallConv CallConv
	// Basic blocks of the function; or nil if defined externally.
	Blocks []*BasicBlock
	// Metadata attached to the function.
	Metadata []*AttachedMD
}

// GetName returns the name of the value.
func (f *Function) GetName() string {
	return f.Name
}

// SetName sets the name of the value.
func (f *Function) SetName(name string) {
	f.Name = name
}

// isValue ensures that only values can be assigned to the ast.Value interface.
func (*Function) isValue() {}

// isConstant ensures that only constants can be assigned to the ast.Constant
// interface.
func (*Function) isConstant() {}

// AssignIDs assigns unique local IDs to unnamed basic blocks and local
// variables of the function.
func (f *Function) AssignIDs() {
	id := 0
	setName := func(n NamedValue) {
		name := n.GetName()
		switch {
		case isUnnamed(name):
			n.SetName(strconv.Itoa(id))
			id++
		case isID(name):
			want := strconv.Itoa(id)
			if name != want {
				panic(fmt.Errorf("invalid local ID in function %s; expected %s, got %s", enc.Global(f.Name), enc.Local(want), enc.Local(name)))
			}
			id++
		}
	}
	for _, param := range f.Sig.Params {
		// Assign local IDs to unnamed parameters of function definitions.
		if len(f.Blocks) > 0 {
			setName(param)
		}
	}
	for _, block := range f.Blocks {
		// Assign local IDs to unnamed basic blocks.
		setName(block)
		for _, inst := range block.Insts {
			n, ok := inst.(NamedValue)
			if !ok {
				continue
			}
			if inst, ok := inst.(*InstCall); ok {
				if _, ok := inst.Type.(*VoidType); ok {
					continue
				}
				if sig, ok := inst.Type.(*FuncType); ok {
					if _, ok := sig.Ret.(*VoidType); ok {
						continue
					}
				}
			}
			// Assign local IDs to unnamed local variables.
			setName(n)
		}
	}
}

// isUnnamed reports whether the given identifier is unnamed.
func isUnnamed(name string) bool {
	return len(name) == 0
}

// isID reports whether the given identifier is an ID (e.g. "%42").
func isID(name string) bool {
	for _, r := range name {
		if strings.IndexRune("0123456789", r) == -1 {
			return false
		}
	}
	return len(name) > 0
}

// CallConv represents the set of calling conventions.
type CallConv uint

// TODO: Change calling convention enums to match the Haskell LLVM library.

// Calling conventions.
const (
	CallConvNone CallConv = iota // no calling convention specified.

	// amdgpu_cs
	// amdgpu_gs
	// amdgpu_kernel
	// amdgpu_ps
	// amdgpu_vs
	// anyregcc
	// arm_aapcs_vfpcc
	// arm_aapcscc
	// arm_apcscc
	// avr_intrcc
	// avr_signalcc
	// cc int_lit
	// ccc
	// coldcc
	// cxx_fast_tlscc
	// fastcc
	// ghccc
	// hhvm_ccc
	// hhvmcc
	// intel_ocl_bicc
	// msp430_intrcc
	// preserve_allcc
	// preserve_mostcc
	// ptx_device
	// ptx_kernel
	// spir_func
	// spir_kernel
	// swiftcc
	// webkit_jscc
	CallConvX86_64SysV // x86_64_sysvcc
	// x86_64_win64cc

	CallConvX86FastCall // x86_fastcallcc
	// x86_intrcc
	// x86_regcallcc

	CallConvX86StdCall // x86_stdcallcc

	// x86_thiscallcc
	// x86_vectorcallcc
)
