package lexer

// keywords is the set of valid keywords in LLVM IR, grouped by length.
var keywords = [][]string{
	0:  nil,
	1:  {"c", "x"},
	2:  {"br", "cc", "eq", "gc", "ne", "or", "to"},
	3:  {"add", "and", "any", "asm", "ccc", "max", "min", "mul", "nsw", "nsz", "nuw", "oeq", "oge", "ogt", "ole", "olt", "one", "ord", "phi", "ret", "sge", "sgt", "shl", "sle", "slt", "ssp", "sub", "ueq", "uge", "ugt", "ule", "ult", "une", "uno", "xor"},
	4:  {"arcp", "ashr", "call", "cold", "fadd", "fast", "fcmp", "fdiv", "fmul", "frem", "fsub", "half", "icmp", "load", "lshr", "nand", "nest", "ninf", "nnan", "null", "sdiv", "sext", "srem", "sret", "tail", "true", "type", "udiv", "umax", "umin", "urem", "void", "weak", "xchg", "zext"},
	5:  {"alias", "align", "byval", "catch", "exact", "false", "fence", "float", "fp128", "fpext", "ghccc", "inreg", "label", "naked", "store", "trunc", "undef"},
	6:  {"alloca", "atomic", "coldcc", "comdat", "common", "define", "double", "fastcc", "filter", "fptosi", "fptoui", "global", "hidden", "invoke", "module", "opaque", "prefix", "resume", "select", "sitofp", "sspreq", "switch", "target", "triple", "uitofp", "unwind", "va_arg"},
	7:  {"acq_rel", "acquire", "bitcast", "builtin", "cleanup", "cmpxchg", "declare", "default", "fptrunc", "largest", "minsize", "noalias", "nonnull", "optnone", "optsize", "private", "release", "section", "seq_cst", "signext", "uwtable", "x86_mmx", "zeroext"},
	8:  {"anyregcc", "constant", "external", "inalloca", "inbounds", "internal", "inttoptr", "linkonce", "metadata", "musttail", "noinline", "noreturn", "nounwind", "prologue", "ptrtoint", "readnone", "readonly", "returned", "samesize", "volatile", "weak_odr", "x86_fp80"},
	9:  {"addrspace", "appending", "atomicrmw", "dllexport", "dllimport", "jumptable", "localexec", "monotonic", "nobuiltin", "nocapture", "noredzone", "ppc_fp128", "protected", "spir_func", "sspstrong", "unordered"},
	10: {"alignstack", "arm_apcscc", "attributes", "datalayout", "exactmatch", "indirectbr", "inlinehint", "landingpad", "ptx_device", "ptx_kernel", "sideeffect"},
	11: {"arm_aapcscc", "extern_weak", "initialexec", "insertvalue", "noduplicate", "nonlazybind", "personality", "spir_kernel", "unreachable", "webkit_jscc"},
	12: {"alwaysinline", "blockaddress", "extractvalue", "inteldialect", "linkonce_odr", "localdynamic", "noduplicates", "singlethread", "thread_local", "unnamed_addr", "uselistorder"},
	13: {"addrspacecast", "getelementptr", "insertelement", "msp430_intrcc", "returns_twice", "shufflevector", "x86_64_sysvcc", "x86_stdcallcc"},
	14: {"extractelement", "intel_ocl_bicc", "preserve_allcc", "x86_64_win64cc", "x86_fastcallcc", "x86_thiscallcc"},
	15: {"arm_aapcs_vfpcc", "dereferenceable", "noimplicitfloat", "preserve_mostcc", "sanitize_memory", "sanitize_thread", "uselistorder_bb", "zeroinitializer"},
	16: {"sanitize_address", "x86_vectorcallcc"},
	17: nil,
	18: nil,
	19: nil,
	20: {"available_externally"},
	21: nil,
	22: {"externally_initialized"},
}
