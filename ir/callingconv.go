package ir

// CallingConv specifies the calling convention of functions, calls and invokes
// [1].
//
//    [1]: http://llvm.org/docs/LangRef.html#calling-conventions
type CallingConv uint

// Calling conventions.
const (
	// CallC represents the C calling convention, which is the default calling
	// convention of LLVM and the only one that supports varargs.
	CallC CallingConv = 0
	// CallFast represents the fast calling convention, which attempts to make
	// calls as fast as possible (e.g. by passing arguments in registers).
	CallFast CallingConv = 8
	// CallCold represents the cold calling convention, which attempts to make
	// code in the caller as efficient as possible under the assumption that the
	// callee is not commonly executed. As such, the callee often preserve all
	// registers for the caller.
	CallCold CallingConv = 9
	// CallGHC represents the calling convention of the Glasgow Haskell Compiler
	// (GHC).
	CallGHC CallingConv = 10
	// CallHiPE represents the calling convention of the High-Performance Erlang
	// (HiPE) compiler.
	CallHiPE CallingConv = 11
	// CallWebkitJS represents the calling convention of the WebKit JavaScript
	// engine.
	CallWebkitJS CallingConv = 12
	// CallAnyReg represents a calling convention which forces the call arguments
	// into registers but allows them to be dynamically allocated.
	CallAnyReg CallingConv = 13
	// CallPreserveMost represents a calling convention for runtime calls that
	// preserves most registers.
	CallPreserveMost CallingConv = 14
	// CallPreserveAll represents a calling convention for runtime calls that
	// preserves (almost) all registers.
	CallPreserveAll CallingConv = 15
)
