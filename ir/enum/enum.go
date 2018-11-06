// Package enum defines enumerate types of LLVM IR.
package enum

//go:generate stringer -linecomment -type AtomicOp

// AtomicOp is an AtomicRMW binary operation.
type AtomicOp uint8

// AtomicRMW binary operations.
const (
	AtomicOpAdd  AtomicOp = iota + 1 // add
	AtomicOpAnd                      // and
	AtomicOpMax                      // max
	AtomicOpMin                      // min
	AtomicOpNAnd                     // nand
	AtomicOpOr                       // or
	AtomicOpSub                      // sub
	AtomicOpUMax                     // umax
	AtomicOpUMin                     // umin
	AtomicOpXChg                     // xchg
	AtomicOpXor                      // xor
)

//go:generate stringer -linecomment -type AtomicOrdering

// AtomicOrdering is an atomic ordering attribute.
type AtomicOrdering uint8

// Atomic ordering attributes.
const (
	AtomicOrderingNone      AtomicOrdering = iota // none
	AtomicOrderingAcqRel                          // acq_rel
	AtomicOrderingAcquire                         // acquire
	AtomicOrderingMonotonic                       // monotonic
	AtomicOrderingRelease                         // release
	AtomicOrderingSeqCst                          // seq_cst
	AtomicOrderingUnordered                       // unordered
)

//go:generate stringer -linecomment -type CallingConv

// CallingConv is a calling convention.
type CallingConv uint8

// TODO: Check if there are any calling conventions defined in LLVM 7.0 that are
// missing from this list.

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

//go:generate stringer -linecomment -type ClauseType

// ClauseType specifies the clause type of a landingpad clause.
type ClauseType uint8

// Clause types.
const (
	ClauseTypeCatch  ClauseType = iota + 1 // catch
	ClauseTypeFilter                       // filter
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

//go:generate stringer -linecomment -type FastMathFlag

// FastMathFlag is a fast-math flag.
type FastMathFlag uint8

// Fast-math flags.
const (
	FastMathFlagAFn      FastMathFlag = iota // afn
	FastMathFlagARcp                         // arcp
	FastMathFlagContract                     // contract
	FastMathFlagFast                         // fast
	FastMathFlagNInf                         // ninf
	FastMathFlagNNaN                         // nnan
	FastMathFlagNSZ                          // nsz
	FastMathFlagReassoc                      // reassoc
)

//go:generate stringer -linecomment -type FPred

// FPred is a floating-point comparison predicate.
type FPred uint8

// Floating-point predicates.
const (
	FPredFalse FPred = iota // false
	FPredOEQ                // oeq
	FPredOGE                // oge
	FPredOGT                // ogt
	FPredOLE                // ole
	FPredOLT                // olt
	FPredONE                // one
	FPredORD                // ord
	FPredTrue               // true
	FPredUEQ                // ueq
	FPredUGE                // uge
	FPredUGT                // ugt
	FPredULE                // ule
	FPredULT                // ult
	FPredUNE                // une
	FPredUNO                // uno
)

//go:generate stringer -linecomment -type FuncAttr

// FuncAttr is a function attribute.
type FuncAttr uint8

// Function attributes.
const (
	FuncAttrAlwaysInline                FuncAttr = iota // alwaysinline
	FuncAttrArgMemOnly                                  // argmemonly
	FuncAttrBuiltin                                     // builtin
	FuncAttrCold                                        // cold
	FuncAttrConvergent                                  // convergent
	FuncAttrInaccessibleMemOrArgMemOnly                 // inaccessiblemem_or_argmemonly
	FuncAttrInaccessibleMemOnly                         // inaccessiblememonly
	FuncAttrInlineHint                                  // inlinehint
	FuncAttrJumpTable                                   // jumptable
	FuncAttrMinSize                                     // minsize
	FuncAttrNaked                                       // naked
	FuncAttrNoBuiltin                                   // nobuiltin
	FuncAttrNoDuplicate                                 // noduplicate
	FuncAttrNoImplicitFloat                             // noimplicitfloat
	FuncAttrNoInline                                    // noinline
	FuncAttrNonLazyBind                                 // nonlazybind
	FuncAttrNoRecurse                                   // norecurse
	FuncAttrNoRedZone                                   // noredzone
	FuncAttrNoReturn                                    // noreturn
	FuncAttrNoUnwind                                    // nounwind
	FuncAttrOptNone                                     // optnone
	FuncAttrOptSize                                     // optsize
	FuncAttrReadNone                                    // readnone
	FuncAttrReadOnly                                    // readonly
	FuncAttrReturnsTwice                                // returns_twice
	FuncAttrSafeStack                                   // safestack
	FuncAttrSanitizeAddress                             // sanitize_address
	FuncAttrSanitizeHWAddress                           // sanitize_hwaddress
	FuncAttrSanitizeMemory                              // sanitize_memory
	FuncAttrSanitizeThread                              // sanitize_thread
	FuncAttrSpeculatable                                // speculatable
	FuncAttrSSP                                         // ssp
	FuncAttrSSPReq                                      // sspreq
	FuncAttrSSPStrong                                   // sspstrong
	FuncAttrStrictFP                                    // strictfp
	FuncAttrUwtable                                     // uwtable
	FuncAttrWriteOnly                                   // writeonly
)

//go:generate stringer -linecomment -type IPred

// IPred is an integer comparison predicate.
type IPred uint8

// Integer predicates.
const (
	IPredEQ  IPred = iota // eq
	IPredNE               // ne
	IPredSGE              // sge
	IPredSGT              // sgt
	IPredSLE              // sle
	IPredSLT              // slt
	IPredUGE              // uge
	IPredUGT              // ugt
	IPredULE              // ule
	IPredULT              // ult
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

//go:generate stringer -linecomment -type OverflowFlag

// OverflowFlag is an integer overflow flag.
type OverflowFlag uint8

// Overflow flags.
const (
	OverflowFlagNSW OverflowFlag = iota // nsw
	OverflowFlagNUW                     // nuw
)

//go:generate stringer -linecomment -type ParamAttr

// ParamAttr is a parameter attribute.
type ParamAttr uint8

// Parameter attributes.
const (
	ParamAttrByval      ParamAttr = iota // byval
	ParamAttrInAlloca                    // inalloca
	ParamAttrInReg                       // inreg
	ParamAttrNest                        // nest
	ParamAttrNoAlias                     // noalias
	ParamAttrNoCapture                   // nocapture
	ParamAttrNonNull                     // nonnull
	ParamAttrReadNone                    // readnone
	ParamAttrReadOnly                    // readonly
	ParamAttrReturned                    // returned
	ParamAttrSignExt                     // signext
	ParamAttrSRet                        // sret
	ParamAttrSwiftError                  // swifterror
	ParamAttrSwiftSelf                   // swiftself
	ParamAttrWriteOnly                   // writeonly
	ParamAttrZeroExt                     // zeroext
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

//go:generate stringer -linecomment -type ReturnAttr

// ReturnAttr is a return argument attribute.
type ReturnAttr uint8

// Return argument attributes.
const (
	ReturnAttrInReg   ReturnAttr = iota // inreg
	ReturnAttrNoAlias                   // noalias
	ReturnAttrNonNull                   // nonnull
	ReturnAttrSignExt                   // signext
	ReturnAttrZeroExt                   // zeroext
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

//go:generate stringer -linecomment -type Tail

// Tail is a tail call attribute.
type Tail uint8

// Tail call attributes.
const (
	TailNone     Tail = iota // none
	TailMustTail             // musttail
	TailNoTail               // notail
	TailTail                 // tail
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
