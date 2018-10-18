package ll

//go:generate stringer -linecomment -type CallingConv

// CallingConv is a calling convention.
type CallingConv uint8

// Calling conventions.
const (
	CallingConvNone          CallingConv = iota // none
	CallingConvAmdGPUCS                         // amdgpu_cs
	CallingConvAmdGPUES                         // amdgpu_es
	CallingConvAmdGPUGS                         // amdgpu_gs
	CallingConvAmdGPUHS                         // amdgpu_hs
	CallingConvAmdGPUKernel                     // amdgpu_kernel
	CallingConvAmdGPULS                         // amdgpu_ls
	CallingConvAmdGPUPS                         // amdgpu_ps
	CallingConvAmdGPUVS                         // amdgpu_vs
	CallingConvAnyReg                           // anyregcc
	CallingConvARMAAPCSVFP                      // arm_aapcs_vfpcc
	CallingConvARMAAPCS                         // arm_aapcscc
	CallingConvARMAPCS                          // arm_apcscc
	CallingConvAVRIntr                          // avr_intrcc
	CallingConvAVRSignal                        // avr_signalcc
	CallingConvC                                // ccc
	CallingConvCold                             // coldcc
	CallingConvCXXFastTLS                       // cxx_fast_tlscc
	CallingConvFast                             // fastcc
	CallingConvGHC                              // ghccc
	CallingConvHHVMC                            // hhvm_ccc
	CallingConvHHVM                             // hhvmcc
	CallingConvIntelOCLBI                       // intel_ocl_bicc
	CallingConvMSP430Intr                       // msp430_intrcc
	CallingConvPreserveAll                      // preserve_allcc
	CallingConvPreserveMost                     // preserve_mostcc
	CallingConvPTXDevice                        // ptx_device
	CallingConvPTXKernel                        // ptx_kernel
	CallingConvSPIRFunc                         // spir_func
	CallingConvSPIRKernel                       // spir_kernel
	CallingConvSwift                            // swiftcc
	CallingConvWebKitJS                         // webkit_jscc
	CallingConvWin64                            // win64cc
	CallingConvX86_64SysV                       // x86_64_sysvcc
	CallingConvX86FastCall                      // x86_fastcallcc
	CallingConvX86Intr                          // x86_intrcc
	CallingConvX86RegCall                       // x86_regcallcc
	CallingConvX86StdCall                       // x86_stdcallcc
	CallingConvX86ThisCall                      // x86_thiscallcc
	CallingConvX86VectorCall                    // x86_vectorcallcc
	// Calling conventions defined through cc NNN.
	CallingConvHiPE          // cc 11
	CallingConvAVRBuiltin    // cc 86
	CallingConvAMDGPUVS      // cc 87
	CallingConvAMDGPUGS      // cc 88
	CallingConvAMDGPUPS      // cc 89
	CallingConvAMDGPUCS      // cc 90
	CallingConvAMDGPUKernel  // cc 91
	CallingConvAMDGPUHS      // cc 93
	CallingConvMSP430Builtin // cc 94
	CallingConvAMDGPULS      // cc 95
	CallingConvAMDGPUES      // cc 96
)

//go:generate stringer -linecomment -type DLLStorageClass

// DLLStorageClass specifies the DLL storage class of a global identifier.
type DLLStorageClass uint8

// DLL storage classes.
const (
	DLLStorageClassNone      DLLStorageClass = iota // none
	DLLStorageClassDLLExport                        // dllexport
	DLLStorageClassDLLImport                        // dllimport
)

//go:generate stringer -linecomment -type Linkage

// Linkage specifies the linkage of a global identifier.
type Linkage uint8

// Linkage kinds.
const (
	LinkageNone                Linkage = iota // none
	LinkageAppending                          // appending
	LinkageAvailableExternally                // available_externally
	LinkageCommon                             // common
	LinkageInternal                           // internal
	LinkageLinkOnce                           // linkonce
	LinkageLinkOnceODR                        // linkonce_odr
	LinkagePrivate                            // private
	LinkageWeak                               // weak
	LinkageWeakODR                            // weak_odr
	// External linkage.
	LinkageExternal   // external
	LinkageExternWeak // extern_weak
)

//go:generate stringer -linecomment -type Preemption

// Preemption specifies the preemtion of a global identifier.
type Preemption uint8

// Preemption kinds.
const (
	PreemptionNone           Preemption = iota // none
	PreemptionDSOLocal                         // dso_local
	PreemptionDSOPreemptable                   // dso_preemptable
)

//go:generate stringer -linecomment -type SelectionKind

// SelectionKind is a Comdat selection kind.
type SelectionKind uint8

// Comdat selection kinds.
const (
	SelectionKindAny          SelectionKind = iota // any
	SelectionKindExactMatch                        // exactmatch
	SelectionKindLargest                           // largest
	SelectionKindNoDuplicates                      // noduplicates
	SelectionKindSameSize                          // samesize
)

//go:generate stringer -linecomment -type TLSModel

// TLSModel is a thread local storage model.
type TLSModel uint8

// Thread local storage models.
const (
	TLSModelNone TLSModel = iota // none
	// If no explicit model is given, the "general dynamic" model is used.
	TLSModelGeneric      // thread_local
	TLSModelInitialExec  // thread_local(initialexec)
	TLSModelLocalDynamic // thread_local(localdynamic)
	TLSModelLocalExec    // thread_local(localexec)
)

//go:generate stringer -linecomment -type UnnamedAddr

// UnnamedAddr specifies whether the address is significant.
type UnnamedAddr uint8

// Unnamed address specifiers.
const (
	UnnamedAddrNone             UnnamedAddr = iota // none
	UnnamedAddrLocalUnnamedAddr                    // local_unnamed_addr
	UnnamedAddrUnnamedAddr                         // unnamed_addr
)

//go:generate stringer -linecomment -type Visibility

// Visibility specifies the visibility of a global identifier.
type Visibility uint8

// Visibility kinds.
const (
	VisibilityNone      Visibility = iota // none
	VisibilityDefault                     // default
	VisibilityHidden                      // hidden
	VisibilityProtected                   // protected
)
