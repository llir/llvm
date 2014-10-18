package ir

// Linkage specifies the linkage type of a global variable or function [1].
//
//    [1]: http://llvm.org/docs/LangRef.html#linkage-types
type Linkage uint16

// Linkage types.
const (
	// LinkPrivate is only allowed on global definitions, and specifies that the
	// definition is only accessible from within the current module.
	//
	// Data but no symbol is emitted to the object file.
	LinkPrivate Linkage = 1 << iota
	// LinkInternal is functionally equivalent to private linkage, except that a
	// local symbol is emitted to the object file.
	//
	// Data and a local symbol (STB_LOCAL in the case of ELF) are emitted to the
	// object file.
	LinkInternal
	// LinkAvailableExternally is only allowed on global definitions, and
	// specifies that the definition is known to exist outside the module. This
	// linkage type exists to allow inlining and other optimizations to take
	// place given knowledge of the definition of the global.
	//
	// Nothing is emitted into the object file.
	LinkAvailableExternally
	// LinkLinkonce is only allowed on global definitions, and specifies that the
	// definition is externally visible and merged with other symbols of the same
	// name during linkage. Regular globals have higher precedence than linkonce
	// globals during linkage.
	//
	// This linkage type can be used to implement some forms of inline functions,
	// templates, or other code which must be generated in each translation unit
	// that uses it, but where the body may be overridden with a more definitive
	// definition later. Unreferenced linkonce globals are allowed to be
	// discarded.
	//
	// Note that linkonce linkage doesn't allow the optimizer to inline the
	// function body, as its definition may be overridden by a stronger
	// definition. Use "linkonce_odr" for this purpose.
	//
	// Data and a weak symbol (STB_WEAK in the case of ELF) are emitted to the
	// object file.
	LinkLinkonce
	// LinkWeak is functionally equivalent to linkonce linkage, except that
	// unreference globals may not be discarded.
	//
	// Data and a weak symbol (STB_WEAK in the case of ELF) are emitted to the
	// object file.
	LinkWeak
	// LinkCommon is similar to weak linkage, except that common symbols have
	// higher precedence, may not have an explicit section, must have a zero
	// initializer, and may not be marked as constant. Only global variables may
	// have common linkage.
	//
	// No data but a common symbol (SHN_COMMON in the case of ELF) is emitted to
	// the object file.
	LinkCommon
	// LinkAppending is only allowed on global variable definitions of pointer to
	// array type. When two global variables with appending linkage are linked
	// together, the two global arrays are appended together in the resulting
	// bitcode.
	LinkAppending
	// LinkExternWeak is functionally equivalent to external linkage, except that
	// the undefined symbol is regarded as weak.
	//
	// No data but a weak, undefined symbol (STB_WEAK and SHN_UNDEF in the case
	// of ELF) is emitted to the object file.
	LinkExternWeak
	// LinkLinkonceODR is functionally equivalent to linkonce linkage, except
	// that it uses the "One Definition Rule" to indicates that only equivalent
	// globals are ever merged; thus allowing the optimizer to inline function
	// bodies.
	//
	// Data and a weak symbol (STB_WEAK in the case of ELF) are emitted to the
	// object file.
	LinkLinkonceODR
	// LinkWeakODR is functionally equivalent to weak linkage, except that it
	// uses the "One Definition Rule" to indicates that only equivalent globals
	// are ever merged; thus allowing the optimizer to inline function bodies.
	//
	// Data and a weak symbol (STB_WEAK in the case of ELF) are emitted to the
	// object file.
	LinkWeakODR
	// LinkExternal is only allowed on global declarations, and specifies that
	// the definition is externally visible.
	//
	// No data but an undefined symbol (SHN_UNDEF in the case of ELF) is emitted
	// to the object file.
	LinkExternal
)
