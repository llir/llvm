// Package token defines constants representing the lexical tokens of the LLVM
// IR assembly language.
package token

// A Token represents a lexical token of the LLVM IR assembly language.
type Token struct {
	// The token type.
	Kind
	// The string value of the token.
	Val string
	// Start position in the input string.
	Pos int
}

func (tok Token) String() string {
	if tok.Kind == EOF {
		return "EOF"
	}
	return tok.Val
}

//go:generate stringer -type Kind

// Kind is the set of lexical token types of the LLVM IR assembly language.
type Kind uint16

// NOTE: The token kinds are based on lib/AsmParser/LLToken.h (rev 224917) and
// docs/LangRef.rst (rev 223189) of LLVM.

// Token types.
const (
	// Special tokens.
	EOF     Kind = iota // End of file
	Error               // Token value holds error message (e.g. unterminated string)
	Comment             // ; line comment

	// Identifiers.
	Type        // iN, void, half, float, double, x86_fp80, fp128, ppc_fp128, x86_mmx, label, metadata
	Label       // foo:, "fo\6F":, .42$foo:
	GlobalVar   // @foo, @"fo\6F"
	LocalVar    // %foo, %"fo\6F"
	MetadataVar // !foo, !fo\6F
	ComdatVar   // $foo, $"fo\6F"
	GlobalID    // @42
	LocalID     // %42
	AttrID      // #42

	operatorStart

	// Operators and delimiters.
	Ellipsis // ...
	Equal    // =
	Comma    // ,
	Star     // *
	Lbrack   // [
	Rbrack   // ]
	Lbrace   // {
	Rbrace   // }
	Lparen   // (
	Rparen   // )
	Less     // <
	Greater  // >
	Exclaim  // !

	operatorEnd

	// Constants.
	Int    // 12345, [us]0x[0-9A-Fa-f]+
	Float  // 123.45, 1.2345e+2, 0x[0-9A-Fa-f]+
	String // "foo"

	keywordStart

	// Keywords.
	KwNull     // null
	KwTrue     // true
	KwFalse    // false
	KwX        // x
	KwDeclare  // declare
	KwDefine   // define
	KwGlobal   // global
	KwConstant // constant

	// Linkage types.
	KwPrivate             // private
	KwInternal            // internal
	KwAvailableExternally // available_externally
	KwLinkonce            // linkonce
	KwWeak                // weak; used as a linkage, and a modifier for "cmpxchg".
	KwCommon              // common
	KwAppending           // appending
	KwExternWeak          // extern_weak
	KwLinkonceOdr         // linkonce_odr
	KwWeakOdr             // weak_odr
	KwExternal            // external

	// Calling conventions.
	KwCcc             // ccc
	KwFastcc          // fastcc
	KwColdcc          // coldcc
	KwWebkitJscc      // webkit_jscc
	KwAnyregcc        // anyregcc
	KwPreserveMostcc  // preserve_mostcc
	KwPreserveAllcc   // preserve_allcc
	KwCc              // cc
	KwIntelOclBicc    // intel_ocl_bicc
	KwX86Stdcallcc    // x86_stdcallcc
	KwX86Fastcallcc   // x86_fastcallcc
	KwX86Thiscallcc   // x86_thiscallcc
	KwX86Vectorcallcc // x86_vectorcallcc
	KwArmApcscc       // arm_apcscc
	KwArmAapcscc      // arm_aapcscc
	KwArmAapcsVfpcc   // arm_aapcs_vfpcc
	KwMsp430Intrcc    // msp430_intrcc
	KwPtxKernel       // ptx_kernel
	KwPtxDevice       // ptx_device
	KwSpirKernel      // spir_kernel
	KwSpirFunc        // spir_func
	KwX86_64Sysvcc    // x86_64_sysvcc
	KwX86_64Win64cc   // x86_64_win64cc
	KwGhccc           // ghccc

	// Visibility styles.
	KwDefault   // default
	KwHidden    // hidden
	KwProtected // protected

	// DLL storage classes.
	KwDllimport // dllimport
	KwDllexport // dllexport

	// Thread local storage models.
	KwLocaldynamic // localdynamic
	KwInitialexec  // initialexec
	KwLocalexec    // localexec

	// Global variables.
	KwThreadLocal           // thread_local
	KwUnnamedAddr           // unnamed_addr
	KwAddrspace             // addrspace
	KwExternallyInitialized // externally_initialized
	KwSection               // section

	// Functions.
	KwPrologue // prologue

	// Aliases.
	KwAlias // alias

	// Comdats.
	KwComdat       // comdat
	KwAny          // any
	KwExactmatch   // exactmatch
	KwLargest      // largest
	KwNoduplicates // noduplicates
	KwSamesize     // samesize

	// Parameter attributes.
	KwZeroext         // zeroext
	KwSignext         // signext
	KwInreg           // inreg
	KwByval           // byval
	KwInalloca        // inalloca
	KwSret            // sret
	KwAlign           // align
	KwNoalias         // noalias
	KwNocapture       // nocapture
	KwNest            // nest
	KwReturned        // returned
	KwNonnull         // nonnull
	KwDereferenceable // dereferenceable

	// Garbage collector names.
	KwGc // gc

	// Prefix data.
	KwPrefix // prefix

	// Attribute groups and function attributes.
	KwAttributes      // attributes
	KwAlignstack      // alignstack
	KwAlwaysinline    // alwaysinline
	KwBuiltin         // builtin
	KwCold            // cold
	KwInlinehint      // inlinehint
	KwJumptable       // jumptable
	KwMinsize         // minsize
	KwNaked           // naked
	KwNobuiltin       // nobuiltin
	KwNoduplicate     // noduplicate
	KwNoimplicitfloat // noimplicitfloat
	KwNoinline        // noinline
	KwNonlazybind     // nonlazybind
	KwNoredzone       // noredzone
	KwNoreturn        // noreturn
	KwNounwind        // nounwind
	KwOptnone         // optnone
	KwOptsize         // optsize
	KwReadnone        // readnone
	KwReadonly        // readonly
	KwReturnsTwice    // returns_twice
	KwSanitizeAddress // sanitize_address
	KwSanitizeMemory  // sanitize_memory
	KwSanitizeThread  // sanitize_thread
	KwSsp             // ssp
	KwSspreq          // sspreq
	KwSspstrong       // sspstrong
	KwUwtable         // uwtable

	// Inline assembly.
	KwModule       // module
	KwAsm          // asm
	KwSideeffect   // sideeffect
	KwInteldialect // inteldialect

	// Data layout and target triple.
	KwTarget     // target
	KwDatalayout // datalayout
	KwTriple     // triple

	// Volatile memory accesses.
	KwVolatile // volatile

	// Complex constants.
	KwC               // c
	KwZeroinitializer // zeroinitializer

	// Atomic memory ordering constraints.
	KwUnordered // unordered
	KwMonotonic // monotonic
	KwAcquire   // acquire
	KwRelease   // release
	KwAcqRel    // acq_rel (acquire+release)
	KwSeqCst    // seq_cst (sequentially consistent)

	// atomicrmw operations.
	KwXchg // xchg
	KwNand // nand
	KwMax  // max
	KwMin  // min
	KwUmax // umax
	KwUmin // umin

	// Fast-math flags.
	KwNnan // nnan: No NaNs
	KwNinf // ninf: No Infs
	KwNsz  // nsz: No Signed Zeros
	KwArcp // arcp: Allow Reciprocal
	KwFast // fast: Fast

	// Use-list order directives.
	KwUselistorder   // uselistorder
	KwUselistorderBb // uselistorder_bb

	// Structure types.
	KwType   // type
	KwOpaque // opaque

	// Undefined values.
	KwUndef // undef

	// Addresses of basic blocks.
	KwBlockaddress // blockaddress

	instructionStart

	// Instructions.
	// Terminator instructions.
	KwRet         // ret
	KwBr          // br
	KwSwitch      // switch
	KwIndirectbr  // indirectbr
	KwInvoke      // invoke
	KwResume      // resume
	KwUnreachable // unreachable

	// Binary operations.
	KwAdd  // add
	KwFadd // fadd
	KwSub  // sub
	KwFsub // fsub
	KwMul  // mul
	KwFmul // fmul
	KwUdiv // udiv
	KwSdiv // sdiv
	KwFdiv // fdiv
	KwUrem // urem
	KwSrem // srem
	KwFrem // frem

	// Bitwise binary operations.
	KwShl  // shl
	KwLshr // lshr
	KwAshr // ashr
	KwAnd  // and
	KwOr   // or
	KwXor  // xor

	// Vector operations.
	KwExtractelement // extractelement
	KwInsertelement  // insertelement
	KwShufflevector  // shufflevector

	// Aggregate operations.
	KwExtractvalue // extractvalue
	KwInsertvalue  // insertvalue

	// Memory access and addressing operations.
	KwAlloca        // alloca
	KwLoad          // load
	KwStore         // store
	KwFence         // fence
	KwCmpxchg       // cmpxchg
	KwAtomicrmw     // atomicrmw
	KwGetelementptr // getelementptr

	// Conversion operations.
	KwTo            // to
	KwTrunc         // trunc
	KwZext          // zext
	KwSext          // sext
	KwFptrunc       // fptrunc
	KwFpext         // fpext
	KwFptoui        // fptoui
	KwFptosi        // fptosi
	KwUitofp        // uitofp
	KwSitofp        // sitofp
	KwPtrtoint      // ptrtoint
	KwInttoptr      // inttoptr
	KwBitcast       // bitcast
	KwAddrspacecast // addrspacecast

	// Other operations.
	KwIcmp       // icmp
	KwFcmp       // fcmp
	KwPhi        // phi
	KwSelect     // select
	KwCall       // call
	KwVaArg      // va_arg
	KwLandingpad // landingpad

	instructionEnd

	// Poison behaviour.
	KwNuw      // nuw: No Unsigned Wrap
	KwNsw      // nsw: No Signed Wrap
	KwExect    // exact
	KwInbounds // inbounds

	// Concurrency behaviour.
	KwAtomic       // atomic
	KwSinglethread // singlethread

	// icmp conditional codes.
	KwEq  // eq: equal
	KwNe  // ne: not equal
	KwUgt // ugt: unsigned greater than
	KwUge // uge: unsigned greater or equal
	KwUlt // ult: unsigned less than
	KwUle // ule: unsigned less or equal
	KwSgt // sgt: signed greater than
	KwSge // sge: signed greater or equal
	KwSlt // slt: signed less than
	KwSle // sle: signed less or equal

	// fcmp conditional codes.
	KwOeq // oeq: ordered and equal
	KwOgt // ogt: ordered and greater than
	KwOge // oge: ordered and greater than or equal
	KwOlt // olt: ordered and less than
	KwOle // ole: ordered and less than or equal
	KwOne // one: ordered and not equal
	KwOrd // ord: ordered (no nans)
	KwUeq // ueq: unordered or equal
	KwUne // une: unordered or not equal
	KwUno // uno: unordered (either nans)

	// Tail calls.
	KwTail     // tail
	KwMusttail // musttail

	// Exception handling.
	KwPersonality // personality
	KwCleanup     // cleanup
	KwCatch       // catch
	KwFilter      // filter
	KwUnwind      // unwind

	keywordEnd
)

// IsKeyword returns true if kind is a keyword, and false otherwise.
func (kind Kind) IsKeyword() bool {
	return keywordStart < kind && kind < keywordEnd
}

// IsOperator returns true if kind is an operator or a delimiter, and false
// otherwise.
func (kind Kind) IsOperator() bool {
	return operatorStart < kind && kind < operatorEnd
}

// IsInstruction returns true if kind is an instruction, and false otherwise.
func (kind Kind) IsInstruction() bool {
	return instructionStart < kind && kind < instructionEnd
}

// TODO: Decide if the IsLiteral methods should be included, removed or renamed
// (to IsConst).

/*
// IsLiteral returns true if kind is an identifier or a basic literal, and false
// otherwise.
func (kind Kind) IsLiteral() bool {
	return literalStart < kind && kind < literalEnd
}
*/
