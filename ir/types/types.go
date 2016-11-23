// === [ Types ] ===============================================================
//
// References:
//    http://llvm.org/docs/LangRef.html#type-system

// Package types declares the data types of LLVM IR.
package types

import "fmt"

// A Type represents an LLVM IR type.
//
// Type may have one of the following underlying types.
//
//    *types.VoidType
//    *types.LabelType
//    *types.IntType
//    *types.FloatType
//    *types.FuncType
//    *types.PointerType
//    *types.VectorType
//    *types.ArrayType
//    *types.StructType
type Type interface {
	fmt.Stringer
	// Equal reports whether t and u are of equal type.
	Equal(u Type) bool
}

// Equal reports whether t and u are of equal type.
func Equal(t, u Type) bool {
	return t.Equal(u)
}

// Convenience types.
var (
	// Void represents the `void` type.
	Void = &VoidType{}
	// Label represents the `label` type.
	Label = &LabelType{}
	// I1 represents the `i1` integer type.
	I1 = NewInt(1)
	// I8 represents the `i8` integer type.
	I8 = NewInt(8)
	// I16 represents the `i16` integer type.
	I16 = NewInt(16)
	// I32 represents the `i32` integer type.
	I32 = NewInt(32)
	// I64 represents the `i64` integer type.
	I64 = NewInt(64)
	// I128 represents the `i128` integer type.
	I128 = NewInt(128)
	// Half represents the `half` floating-point type.
	Half = &FloatType{kind: FloatKindIEEE_16}
	// Float represents the `float` floating-point type.
	Float = &FloatType{kind: FloatKindIEEE_32}
	// Double represents the `double` floating-point type.
	Double = &FloatType{kind: FloatKindIEEE_64}
	// FP128 represents the `fp128` floating-point type.
	FP128 = &FloatType{kind: FloatKindIEEE_128}
	// X86_FP80 represents the `x86_fp80` floating-point type.
	X86_FP80 = &FloatType{kind: FloatKindDoubleExtended_80}
	// PPC_FP128 represents the `ppc_fp128` floating-point type.
	PPC_FP128 = &FloatType{kind: FloatKindDoubleDouble_128}
)
