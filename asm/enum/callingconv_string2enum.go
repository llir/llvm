// Code generated by "string2enum -linecomment -type CallingConv ../../ir/enum"; DO NOT EDIT.

package enum

import (
	"fmt"

	"github.com/llir/llvm/ir/enum"
)

func _() {
	// An "invalid array index" compiler error signifies that the constant values have changed.
	// Re-run the string2enum command to generate them again.
	var x [1]struct{}
	_ = x[enum.CallingConvNone-0]
	_ = x[enum.CallingConvC-1]
	_ = x[enum.CallingConvFast-8]
	_ = x[enum.CallingConvCold-9]
	_ = x[enum.CallingConvGHC-10]
	_ = x[enum.CallingConvHiPE-11]
	_ = x[enum.CallingConvWebKitJS-12]
	_ = x[enum.CallingConvAnyReg-13]
	_ = x[enum.CallingConvPreserveMost-14]
	_ = x[enum.CallingConvPreserveAll-15]
	_ = x[enum.CallingConvSwift-16]
	_ = x[enum.CallingConvCXXFastTLS-17]
	_ = x[enum.CallingConvTail-18]
	_ = x[enum.CallingConvCFGuardCheck-19]
	_ = x[enum.CallingConvX86StdCall-64]
	_ = x[enum.CallingConvX86FastCall-65]
	_ = x[enum.CallingConvARM_APCS-66]
	_ = x[enum.CallingConvARM_AAPCS-67]
	_ = x[enum.CallingConvARM_AAPCS_VFP-68]
	_ = x[enum.CallingConvMSP430Intr-69]
	_ = x[enum.CallingConvX86ThisCall-70]
	_ = x[enum.CallingConvPTXKernel-71]
	_ = x[enum.CallingConvPTXDevice-72]
	_ = x[enum.CallingConvSPIRFunc-75]
	_ = x[enum.CallingConvSPIRKernel-76]
	_ = x[enum.CallingConvIntelOCL_BI-77]
	_ = x[enum.CallingConvX86_64SysV-78]
	_ = x[enum.CallingConvWin64-79]
	_ = x[enum.CallingConvX86VectorCall-80]
	_ = x[enum.CallingConvHHVM-81]
	_ = x[enum.CallingConvHHVM_C-82]
	_ = x[enum.CallingConvX86Intr-83]
	_ = x[enum.CallingConvAVRIntr-84]
	_ = x[enum.CallingConvAVRSignal-85]
	_ = x[enum.CallingConvAVRBuiltin-86]
	_ = x[enum.CallingConvAMDGPU_VS-87]
	_ = x[enum.CallingConvAMDGPU_GS-88]
	_ = x[enum.CallingConvAMDGPU_PS-89]
	_ = x[enum.CallingConvAMDGPU_CS-90]
	_ = x[enum.CallingConvAMDGPUKernel-91]
	_ = x[enum.CallingConvX86RegCall-92]
	_ = x[enum.CallingConvAMDGPU_HS-93]
	_ = x[enum.CallingConvMSP430Builtin-94]
	_ = x[enum.CallingConvAMDGPU_LS-95]
	_ = x[enum.CallingConvAMDGPU_ES-96]
	_ = x[enum.CallingConvAArch64VectorCall-97]
	_ = x[enum.CallingConvAArch64SVEVectorCall-98]
	_ = x[enum.CallingConvAMDGPUGfx-100]
}

const (
	_CallingConv_name_0 = "noneccc"
	_CallingConv_name_1 = "fastcccoldccghccccc 11webkit_jsccanyregccpreserve_mostccpreserve_allccswiftcccxx_fast_tlscctailcccfguard_checkcc"
	_CallingConv_name_2 = "x86_stdcallccx86_fastcallccarm_apcsccarm_aapcsccarm_aapcs_vfpccmsp430_intrccx86_thiscallccptx_kernelptx_device"
	_CallingConv_name_3 = "spir_funcspir_kernelintel_ocl_biccx86_64_sysvccwin64ccx86_vectorcallcchhvmcchhvm_cccx86_intrccavr_intrccavr_signalcccc 86amdgpu_vsamdgpu_gsamdgpu_psamdgpu_csamdgpu_kernelx86_regcallccamdgpu_hscc 94amdgpu_lsamdgpu_esaarch64_vector_pcsaarch64_sve_vector_pcs"
	_CallingConv_name_4 = "amdgpu_gfx"
)

var (
	_CallingConv_index_0 = [...]uint8{0, 4, 7}
	_CallingConv_index_1 = [...]uint8{0, 6, 12, 17, 22, 33, 41, 56, 70, 77, 91, 97, 112}
	_CallingConv_index_2 = [...]uint8{0, 13, 27, 37, 48, 63, 76, 90, 100, 110}
	_CallingConv_index_3 = [...]uint8{0, 9, 20, 34, 47, 54, 70, 76, 84, 94, 104, 116, 121, 130, 139, 148, 157, 170, 183, 192, 197, 206, 215, 233, 255}
)

// CallingConvFromString returns the CallingConv enum corresponding to s.
func CallingConvFromString(s string) enum.CallingConv {
	if len(s) == 0 {
		return 0
	}
	for i := range _CallingConv_index_0[:len(_CallingConv_index_0)-1] {
		if s == _CallingConv_name_0[_CallingConv_index_0[i]:_CallingConv_index_0[i+1]] {
			return enum.CallingConv(i + 0)
		}
	}
	for i := range _CallingConv_index_1[:len(_CallingConv_index_1)-1] {
		if s == _CallingConv_name_1[_CallingConv_index_1[i]:_CallingConv_index_1[i+1]] {
			return enum.CallingConv(i + 8)
		}
	}
	for i := range _CallingConv_index_2[:len(_CallingConv_index_2)-1] {
		if s == _CallingConv_name_2[_CallingConv_index_2[i]:_CallingConv_index_2[i+1]] {
			return enum.CallingConv(i + 64)
		}
	}
	for i := range _CallingConv_index_3[:len(_CallingConv_index_3)-1] {
		if s == _CallingConv_name_3[_CallingConv_index_3[i]:_CallingConv_index_3[i+1]] {
			return enum.CallingConv(i + 75)
		}
	}
	if s == _CallingConv_name_4 {
		return enum.CallingConv(100)
	}
	panic(fmt.Errorf("unable to locate CallingConv enum corresponding to %q", s))
}
