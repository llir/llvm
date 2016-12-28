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
//    *types.VoidType       (https://godoc.org/github.com/llir/llvm/ir/types#VoidType)
//    *types.FuncType       (https://godoc.org/github.com/llir/llvm/ir/types#FuncType)
//    *types.IntType        (https://godoc.org/github.com/llir/llvm/ir/types#IntType)
//    *types.FloatType      (https://godoc.org/github.com/llir/llvm/ir/types#FloatType)
//    *types.PointerType    (https://godoc.org/github.com/llir/llvm/ir/types#PointerType)
//    *types.VectorType     (https://godoc.org/github.com/llir/llvm/ir/types#VectorType)
//    *types.LabelType      (https://godoc.org/github.com/llir/llvm/ir/types#LabelType)
//    *types.MetadataType   (https://godoc.org/github.com/llir/llvm/ir/types#MetadataType)
//    *types.ArrayType      (https://godoc.org/github.com/llir/llvm/ir/types#ArrayType)
//    *types.StructType     (https://godoc.org/github.com/llir/llvm/ir/types#StructType)
//    *types.NamedType      (https://godoc.org/github.com/llir/llvm/ir/types#NamedType)
type Type interface {
	fmt.Stringer
	// Equal reports whether t and u are of equal type.
	Equal(u Type) bool
}

// Convenience types.
var (
	// Void represents the `void` type.
	Void = &VoidType{}
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
	Half = &FloatType{Kind: FloatKindIEEE_16}
	// Float represents the `float` floating-point type.
	Float = &FloatType{Kind: FloatKindIEEE_32}
	// Double represents the `double` floating-point type.
	Double = &FloatType{Kind: FloatKindIEEE_64}
	// FP128 represents the `fp128` floating-point type.
	FP128 = &FloatType{Kind: FloatKindIEEE_128}
	// X86_FP80 represents the `x86_fp80` floating-point type.
	X86_FP80 = &FloatType{Kind: FloatKindDoubleExtended_80}
	// PPC_FP128 represents the `ppc_fp128` floating-point type.
	PPC_FP128 = &FloatType{Kind: FloatKindDoubleDouble_128}
	// Label represents the `label` type.
	Label = &LabelType{}
	// Metadata represents the `metadata` type.
	Metadata = &MetadataType{}
)

// Equal reports whether t and u are of equal type.
func Equal(t, u Type) bool {
	return t.Equal(u)
}

// IsVoid reports whether the given type is a void type.
func IsVoid(t Type) bool {
	switch t := t.(type) {
	case *VoidType:
		return true
	case *NamedType:
		return IsVoid(t.Def)
	}
	return false
}

// IsFunc reports whether the given type is a function type.
func IsFunc(t Type) bool {
	switch t := t.(type) {
	case *FuncType:
		return true
	case *NamedType:
		return IsFunc(t.Def)
	}
	return false
}

// IsBool reports whether the given type is a boolean type (i.e. an integer type
// with bit size 1).
func IsBool(t Type) bool {
	switch t := t.(type) {
	case *IntType:
		return t.Size == 1
	case *NamedType:
		return IsBool(t.Def)
	}
	return false
}

// IsInt reports whether the given type is an integer type.
func IsInt(t Type) bool {
	switch t := t.(type) {
	case *IntType:
		return true
	case *NamedType:
		return IsInt(t.Def)
	}
	return false
}

// IsFloat reports whether the given type is a floating-point type.
func IsFloat(t Type) bool {
	switch t := t.(type) {
	case *FloatType:
		return true
	case *NamedType:
		return IsFloat(t.Def)
	}
	return false
}

// IsPointer reports whether the given type is a pointer type.
func IsPointer(t Type) bool {
	switch t := t.(type) {
	case *PointerType:
		return true
	case *NamedType:
		return IsPointer(t.Def)
	}
	return false
}

// IsVector reports whether the given type is a vector type.
func IsVector(t Type) bool {
	switch t := t.(type) {
	case *VectorType:
		return true
	case *NamedType:
		return IsVector(t.Def)
	}
	return false
}

// IsLabel reports whether the given type is a label type.
func IsLabel(t Type) bool {
	switch t := t.(type) {
	case *LabelType:
		return true
	case *NamedType:
		return IsLabel(t.Def)
	}
	return false
}

// IsMetadata reports whether the given type is a metadata type.
func IsMetadata(t Type) bool {
	switch t := t.(type) {
	case *MetadataType:
		return true
	case *NamedType:
		return IsMetadata(t.Def)
	}
	return false
}

// IsArray reports whether the given type is an array type.
func IsArray(t Type) bool {
	switch t := t.(type) {
	case *ArrayType:
		return true
	case *NamedType:
		return IsArray(t.Def)
	}
	return false
}

// IsStruct reports whether the given type is a struct type.
func IsStruct(t Type) bool {
	switch t := t.(type) {
	case *StructType:
		return true
	case *NamedType:
		return IsStruct(t.Def)
	}
	return false
}
