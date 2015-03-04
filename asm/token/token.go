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
	Type        // iN, void, half, float, double, fp128, x86_fp80, ppc_fp128, x86_mmx, label, metadata
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
	Float  // 123.45, 1.2345e+2, 0x[KLMH]?[0-9A-Fa-f]+
	String // "foo"

	keywordStart

	instructionStart

	// Keywords.
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

// Keywords is the set of valid keywords in LLVM IR assembly.
var Keywords = map[string]Kind{
	// Types.
	"void":      Type,
	"half":      Type,
	"float":     Type,
	"double":    Type,
	"fp128":     Type,
	"x86_fp80":  Type,
	"ppc_fp128": Type,
	"x86_mmx":   Type,
	"label":     Type,
	"metadata":  Type,

	// Instructions.
	"ret":            KwRet,
	"br":             KwBr,
	"switch":         KwSwitch,
	"indirectbr":     KwIndirectbr,
	"invoke":         KwInvoke,
	"resume":         KwResume,
	"unreachable":    KwUnreachable,
	"add":            KwAdd,
	"fadd":           KwFadd,
	"sub":            KwSub,
	"fsub":           KwFsub,
	"mul":            KwMul,
	"fmul":           KwFmul,
	"udiv":           KwUdiv,
	"sdiv":           KwSdiv,
	"fdiv":           KwFdiv,
	"urem":           KwUrem,
	"srem":           KwSrem,
	"frem":           KwFrem,
	"shl":            KwShl,
	"lshr":           KwLshr,
	"ashr":           KwAshr,
	"and":            KwAnd,
	"or":             KwOr,
	"xor":            KwXor,
	"extractelement": KwExtractelement,
	"insertelement":  KwInsertelement,
	"shufflevector":  KwShufflevector,
	"extractvalue":   KwExtractvalue,
	"insertvalue":    KwInsertvalue,
	"alloca":         KwAlloca,
	"load":           KwLoad,
	"store":          KwStore,
	"fence":          KwFence,
	"cmpxchg":        KwCmpxchg,
	"atomicrmw":      KwAtomicrmw,
	"getelementptr":  KwGetelementptr,
	"to":             KwTo,
	"trunc":          KwTrunc,
	"zext":           KwZext,
	"sext":           KwSext,
	"fptrunc":        KwFptrunc,
	"fpext":          KwFpext,
	"fptoui":         KwFptoui,
	"fptosi":         KwFptosi,
	"uitofp":         KwUitofp,
	"sitofp":         KwSitofp,
	"ptrtoint":       KwPtrtoint,
	"inttoptr":       KwInttoptr,
	"bitcast":        KwBitcast,
	"addrspacecast":  KwAddrspacecast,
	"icmp":           KwIcmp,
	"fcmp":           KwFcmp,
	"phi":            KwPhi,
	"select":         KwSelect,
	"call":           KwCall,
	"va_arg":         KwVaArg,
	"landingpad":     KwLandingpad,

	// Keywords.
	"null":                   KwNull,
	"true":                   KwTrue,
	"false":                  KwFalse,
	"x":                      KwX,
	"declare":                KwDeclare,
	"define":                 KwDefine,
	"global":                 KwGlobal,
	"constant":               KwConstant,
	"private":                KwPrivate,
	"internal":               KwInternal,
	"available_externally":   KwAvailableExternally,
	"linkonce":               KwLinkonce,
	"weak":                   KwWeak,
	"common":                 KwCommon,
	"appending":              KwAppending,
	"extern_weak":            KwExternWeak,
	"linkonce_odr":           KwLinkonceOdr,
	"weak_odr":               KwWeakOdr,
	"external":               KwExternal,
	"ccc":                    KwCcc,
	"fastcc":                 KwFastcc,
	"coldcc":                 KwColdcc,
	"webkit_jscc":            KwWebkitJscc,
	"anyregcc":               KwAnyregcc,
	"preserve_mostcc":        KwPreserveMostcc,
	"preserve_allcc":         KwPreserveAllcc,
	"cc":                     KwCc,
	"intel_ocl_bicc":         KwIntelOclBicc,
	"x86_stdcallcc":          KwX86Stdcallcc,
	"x86_fastcallcc":         KwX86Fastcallcc,
	"x86_thiscallcc":         KwX86Thiscallcc,
	"x86_vectorcallcc":       KwX86Vectorcallcc,
	"arm_apcscc":             KwArmApcscc,
	"arm_aapcscc":            KwArmAapcscc,
	"arm_aapcs_vfpcc":        KwArmAapcsVfpcc,
	"msp430_intrcc":          KwMsp430Intrcc,
	"ptx_kernel":             KwPtxKernel,
	"ptx_device":             KwPtxDevice,
	"spir_kernel":            KwSpirKernel,
	"spir_func":              KwSpirFunc,
	"x86_64_sysvcc":          KwX86_64Sysvcc,
	"x86_64_win64cc":         KwX86_64Win64cc,
	"ghccc":                  KwGhccc,
	"default":                KwDefault,
	"hidden":                 KwHidden,
	"protected":              KwProtected,
	"dllimport":              KwDllimport,
	"dllexport":              KwDllexport,
	"localdynamic":           KwLocaldynamic,
	"initialexec":            KwInitialexec,
	"localexec":              KwLocalexec,
	"thread_local":           KwThreadLocal,
	"unnamed_addr":           KwUnnamedAddr,
	"addrspace":              KwAddrspace,
	"externally_initialized": KwExternallyInitialized,
	"section":                KwSection,
	"prologue":               KwPrologue,
	"alias":                  KwAlias,
	"comdat":                 KwComdat,
	"any":                    KwAny,
	"exactmatch":             KwExactmatch,
	"largest":                KwLargest,
	"noduplicates":           KwNoduplicates,
	"samesize":               KwSamesize,
	"zeroext":                KwZeroext,
	"signext":                KwSignext,
	"inreg":                  KwInreg,
	"byval":                  KwByval,
	"inalloca":               KwInalloca,
	"sret":                   KwSret,
	"align":                  KwAlign,
	"noalias":                KwNoalias,
	"nocapture":              KwNocapture,
	"nest":                   KwNest,
	"returned":               KwReturned,
	"nonnull":                KwNonnull,
	"dereferenceable":        KwDereferenceable,
	"gc":                     KwGc,
	"prefix":                 KwPrefix,
	"attributes":             KwAttributes,
	"alignstack":             KwAlignstack,
	"alwaysinline":           KwAlwaysinline,
	"builtin":                KwBuiltin,
	"cold":                   KwCold,
	"inlinehint":             KwInlinehint,
	"jumptable":              KwJumptable,
	"minsize":                KwMinsize,
	"naked":                  KwNaked,
	"nobuiltin":              KwNobuiltin,
	"noduplicate":            KwNoduplicate,
	"noimplicitfloat":        KwNoimplicitfloat,
	"noinline":               KwNoinline,
	"nonlazybind":            KwNonlazybind,
	"noredzone":              KwNoredzone,
	"noreturn":               KwNoreturn,
	"nounwind":               KwNounwind,
	"optnone":                KwOptnone,
	"optsize":                KwOptsize,
	"readnone":               KwReadnone,
	"readonly":               KwReadonly,
	"returns_twice":          KwReturnsTwice,
	"sanitize_address":       KwSanitizeAddress,
	"sanitize_memory":        KwSanitizeMemory,
	"sanitize_thread":        KwSanitizeThread,
	"ssp":                    KwSsp,
	"sspreq":                 KwSspreq,
	"sspstrong":              KwSspstrong,
	"uwtable":                KwUwtable,
	"module":                 KwModule,
	"asm":                    KwAsm,
	"sideeffect":             KwSideeffect,
	"inteldialect":           KwInteldialect,
	"target":                 KwTarget,
	"datalayout":             KwDatalayout,
	"triple":                 KwTriple,
	"volatile":               KwVolatile,
	"c":                      KwC,
	"zeroinitializer":        KwZeroinitializer,
	"unordered":              KwUnordered,
	"monotonic":              KwMonotonic,
	"acquire":                KwAcquire,
	"release":                KwRelease,
	"acq_rel":                KwAcqRel,
	"seq_cst":                KwSeqCst,
	"xchg":                   KwXchg,
	"nand":                   KwNand,
	"max":                    KwMax,
	"min":                    KwMin,
	"umax":                   KwUmax,
	"umin":                   KwUmin,
	"nnan":                   KwNnan,
	"ninf":                   KwNinf,
	"nsz":                    KwNsz,
	"arcp":                   KwArcp,
	"fast":                   KwFast,
	"uselistorder":           KwUselistorder,
	"uselistorder_bb":        KwUselistorderBb,
	"type":                   KwType,
	"opaque":                 KwOpaque,
	"undef":                  KwUndef,
	"blockaddress":           KwBlockaddress,
	"nuw":                    KwNuw,
	"nsw":                    KwNsw,
	"exact":                  KwExect,
	"inbounds":               KwInbounds,
	"atomic":                 KwAtomic,
	"singlethread":           KwSinglethread,
	"eq":                     KwEq,
	"ne":                     KwNe,
	"ugt":                    KwUgt,
	"uge":                    KwUge,
	"ult":                    KwUlt,
	"ule":                    KwUle,
	"sgt":                    KwSgt,
	"sge":                    KwSge,
	"slt":                    KwSlt,
	"sle":                    KwSle,
	"oeq":                    KwOeq,
	"ogt":                    KwOgt,
	"oge":                    KwOge,
	"olt":                    KwOlt,
	"ole":                    KwOle,
	"one":                    KwOne,
	"ord":                    KwOrd,
	"ueq":                    KwUeq,
	"une":                    KwUne,
	"uno":                    KwUno,
	"tail":                   KwTail,
	"musttail":               KwMusttail,
	"personality":            KwPersonality,
	"cleanup":                KwCleanup,
	"catch":                  KwCatch,
	"filter":                 KwFilter,
	"unwind":                 KwUnwind,
}
