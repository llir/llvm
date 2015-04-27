package lexer

import "sort"

// keywords is the set of valid keywords in LLVM IR
var keywords = []string{
	"c", "x",
	"br", "cc", "eq", "gc", "ne", "or", "to",
	"add", "and", "any", "asm", "ccc", "max", "min", "mul", "nsw", "nsz", "nuw", "oeq", "oge", "ogt", "ole", "olt", "one", "ord", "phi", "ret", "sge", "sgt", "shl", "sle", "slt", "ssp", "sub", "ueq", "uge", "ugt", "ule", "ult", "une", "uno", "xor",
	"arcp", "ashr", "call", "cold", "fadd", "fast", "fcmp", "fdiv", "fmul", "frem", "fsub", "half", "icmp", "load", "lshr", "nand", "nest", "ninf", "nnan", "null", "sdiv", "sext", "srem", "sret", "tail", "true", "type", "udiv", "umax", "umin", "urem", "void", "weak", "xchg", "zext",
	"alias", "align", "byval", "catch", "exact", "false", "fence", "float", "fp128", "fpext", "ghccc", "inreg", "label", "naked", "store", "trunc", "undef",
	"alloca", "atomic", "coldcc", "comdat", "common", "define", "double", "fastcc", "filter", "fptosi", "fptoui", "global", "hidden", "invoke", "module", "opaque", "prefix", "resume", "select", "sitofp", "sspreq", "switch", "target", "triple", "uitofp", "unwind", "va_arg",
	"acq_rel", "acquire", "bitcast", "builtin", "cleanup", "cmpxchg", "declare", "default", "fptrunc", "largest", "minsize", "noalias", "nonnull", "optnone", "optsize", "private", "release", "section", "seq_cst", "signext", "uwtable", "x86_mmx", "zeroext",
	"anyregcc", "constant", "external", "inalloca", "inbounds", "internal", "inttoptr", "linkonce", "metadata", "musttail", "noinline", "noreturn", "nounwind", "prologue", "ptrtoint", "readnone", "readonly", "returned", "samesize", "volatile", "weak_odr", "x86_fp80",
	"addrspace", "appending", "atomicrmw", "dllexport", "dllimport", "jumptable", "localexec", "monotonic", "nobuiltin", "nocapture", "noredzone", "ppc_fp128", "protected", "spir_func", "sspstrong", "unordered",
	"alignstack", "arm_apcscc", "attributes", "datalayout", "exactmatch", "indirectbr", "inlinehint", "landingpad", "ptx_device", "ptx_kernel", "sideeffect",
	"arm_aapcscc", "extern_weak", "initialexec", "insertvalue", "noduplicate", "nonlazybind", "personality", "spir_kernel", "unreachable", "webkit_jscc",
	"alwaysinline", "blockaddress", "extractvalue", "inteldialect", "linkonce_odr", "localdynamic", "noduplicates", "singlethread", "thread_local", "unnamed_addr", "uselistorder",
	"addrspacecast", "getelementptr", "insertelement", "msp430_intrcc", "returns_twice", "shufflevector", "x86_64_sysvcc", "x86_stdcallcc",
	"extractelement", "intel_ocl_bicc", "preserve_allcc", "x86_64_win64cc", "x86_fastcallcc", "x86_thiscallcc",
	"arm_aapcs_vfpcc", "dereferenceable", "noimplicitfloat", "preserve_mostcc", "sanitize_memory", "sanitize_thread", "uselistorder_bb", "zeroinitializer",
	"sanitize_address", "x86_vectorcallcc",
	"available_externally",
	"externally_initialized",
}

func init() {
	sort.Strings(keywords)
}
