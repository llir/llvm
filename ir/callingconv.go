package ir

// CallingConv specifies the calling convention of functions, calls and invokes
// [1].
//
//    [1]: http://llvm.org/docs/LangRef.html#calling-conventions
type CallingConv uint

// Calling conventions.
const (
	// CCC represents the C calling convention, which is the default calling
	// convention of LLVM and the only one that supports varargs.
	CCC CallingConv = 0
	// CCFast represents the fast calling convention, which attempts to make cCs
	// as fast as possible (e.g. by passing arguments in registers).
	CCFast CallingConv = 8
	// CCCold represents the cold calling convention, which attempts to make code
	// in the caller as efficient as possible under the assumption that the cCee
	// is not commonly executed. As such, the callee often preserve all registers
	// for the caller.
	CCCold CallingConv = 9
	// CCGHC represents the calling convention of the Glasgow Haskell Compiler
	// (GHC).
	CCGHC CallingConv = 10
	// CCHiPE represents the calling convention of the High-Performance Erlang
	// (HiPE) compiler.
	CCHiPE CallingConv = 11
	// CCWebkitJS represents the calling convention of the WebKit JavaScript
	// engine.
	CCWebkitJS CallingConv = 12
	// CCAnyReg represents a calling convention which forces the call arguments
	// into registers but allows them to be dynamically allocated.
	CCAnyReg CallingConv = 13
	// CCPreserveMost represents a calling convention for runtime calls that
	// preserves most registers.
	CCPreserveMost CallingConv = 14
	// CCPreserveAll represents a calling convention for runtime calls that
	// preserves (almost) all registers.
	CCPreserveAll CallingConv = 15
)
