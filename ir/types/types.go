// Package types declares the data types of LLVM IR.
package types

import (
	"fmt"
	"strings"

	"github.com/llir/llvm/internal/enc"
)

// === [ Types ] ===============================================================

// Convenience types.
var (
	// Basic types.
	Void     = &VoidType{}     // void
	MMX      = &MMXType{}      // x86_mmx
	Label    = &LabelType{}    // label
	Token    = &TokenType{}    // token
	Metadata = &MetadataType{} // metadata
	// Integer types.
	I1   = &IntType{BitSize: 1}   // i1
	I8   = &IntType{BitSize: 8}   // i8
	I16  = &IntType{BitSize: 16}  // i16
	I32  = &IntType{BitSize: 32}  // i32
	I64  = &IntType{BitSize: 64}  // i64
	I128 = &IntType{BitSize: 128} // i128
	// Floating-point types.
	Half     = &FloatType{Kind: FloatKindHalf}     // half
	Float    = &FloatType{Kind: FloatKindFloat}    // float
	Double   = &FloatType{Kind: FloatKindDouble}   // double
	X86FP80  = &FloatType{Kind: FloatKindX86FP80}  // x86_fp80
	FP128    = &FloatType{Kind: FloatKindFP128}    // fp128
	PPCFP128 = &FloatType{Kind: FloatKindPPCFP128} // ppc_fp128
	// Integer pointer types.
	I1Ptr   = &PointerType{ElemType: I1}   // i1*
	I8Ptr   = &PointerType{ElemType: I8}   // i8*
	I16Ptr  = &PointerType{ElemType: I16}  // i16*
	I32Ptr  = &PointerType{ElemType: I32}  // i32*
	I64Ptr  = &PointerType{ElemType: I64}  // i64*
	I128Ptr = &PointerType{ElemType: I128} // i128*
)

// Convenience functions.

// IsPointer reports whether the given type is a pointer type.
func IsPointer(t Type) bool {
	_, ok := t.(*PointerType)
	return ok
}

// Equal reports whether t and u are of equal type.
func Equal(t, u Type) bool {
	return t.Equal(u)
}

// Type is an LLVM IR type.
type Type interface {
	fmt.Stringer
	// Def returns the LLVM syntax representation of the definition of the type.
	Def() string
	// Equal reports whether t and u are of equal type.
	Equal(u Type) bool
	// Name returns the type name of the type.
	Name() string
}

// --- [ Void types ] ----------------------------------------------------------

// VoidType is an LLVM IR void type.
type VoidType struct {
	// Type name; or empty if not present.
	TypeName string
}

// Equal reports whether t and u are of equal type.
func (t *VoidType) Equal(u Type) bool {
	if _, ok := u.(*VoidType); ok {
		return true
	}
	return false
}

// String returns the string representation of the void type.
func (t *VoidType) String() string {
	if len(t.TypeName) > 0 {
		return enc.Local(t.TypeName)
	}
	return t.Def()
}

// Def returns the LLVM syntax representation of the definition of the type.
func (t *VoidType) Def() string {
	// 'void'
	return "void"
}

// Name returns the type name of the type.
func (t *VoidType) Name() string {
	return t.TypeName
}

// --- [ Function types ] ------------------------------------------------------

// FuncType is an LLVM IR function type.
type FuncType struct {
	// Type name; or empty if not present.
	TypeName string
	// Return type.
	RetType Type
	// Function parameters.
	Params []Type
	// Variable number of function arguments.
	Variadic bool
}

// NewFunc returns a new function type based on the given return type and
// function parameter types.
func NewFunc(retType Type, params ...Type) *FuncType {
	return &FuncType{
		RetType: retType,
		Params:  params,
	}
}

// Equal reports whether t and u are of equal type.
func (t *FuncType) Equal(u Type) bool {
	if u, ok := u.(*FuncType); ok {
		if !t.RetType.Equal(u.RetType) {
			return false
		}
		if len(t.Params) != len(u.Params) {
			return false
		}
		for i := range t.Params {
			if !t.Params[i].Equal(u.Params[i]) {
				return false
			}
		}
		return t.Variadic == u.Variadic
	}
	return false
}

// String returns the string representation of the function type.
func (t *FuncType) String() string {
	if len(t.TypeName) > 0 {
		return enc.Local(t.TypeName)
	}
	return t.Def()
}

// Def returns the LLVM syntax representation of the definition of the type.
func (t *FuncType) Def() string {
	// RetType=Type '(' Params ')'
	buf := &strings.Builder{}
	fmt.Fprintf(buf, "%v (", t.RetType)
	for i, param := range t.Params {
		if i != 0 {
			buf.WriteString(", ")
		}
		buf.WriteString(param.String())
	}
	if t.Variadic {
		if len(t.Params) > 0 {
			buf.WriteString(", ")
		}
		buf.WriteString("...")
	}
	buf.WriteString(")")
	return buf.String()
}

// Name returns the type name of the type.
func (t *FuncType) Name() string {
	return t.TypeName
}

// --- [ Integer types ] -------------------------------------------------------

// IntType is an LLVM IR integer type.
type IntType struct {
	// Type name; or empty if not present.
	TypeName string
	// Integer size in number of bits.
	BitSize int64
}

// NewInt returns a new integer type based on the given integer bit size.
func NewInt(bitSize int64) *IntType {
	return &IntType{
		BitSize: bitSize,
	}
}

// Equal reports whether t and u are of equal type.
func (t *IntType) Equal(u Type) bool {
	if u, ok := u.(*IntType); ok {
		return t.BitSize == u.BitSize
	}
	return false
}

// String returns the string representation of the integer type.
func (t *IntType) String() string {
	if len(t.TypeName) > 0 {
		return enc.Local(t.TypeName)
	}
	return t.Def()
}

// Def returns the LLVM syntax representation of the definition of the type.
func (t *IntType) Def() string {
	// int_type_tok
	return fmt.Sprintf("i%d", t.BitSize)
}

// Name returns the type name of the type.
func (t *IntType) Name() string {
	return t.TypeName
}

// --- [ Floating-point types ] ------------------------------------------------

// FloatType is an LLVM IR floating-point type.
type FloatType struct {
	// Type name; or empty if not present.
	TypeName string
	// Floating-point kind.
	Kind FloatKind
}

// Equal reports whether t and u are of equal type.
func (t *FloatType) Equal(u Type) bool {
	if u, ok := u.(*FloatType); ok {
		return t.Kind == u.Kind
	}
	return false
}

// String returns the string representation of the floating-point type.
func (t *FloatType) String() string {
	if len(t.TypeName) > 0 {
		return enc.Local(t.TypeName)
	}
	return t.Def()
}

// Def returns the LLVM syntax representation of the definition of the type.
func (t *FloatType) Def() string {
	// FloatKind
	return t.Kind.String()
}

// Name returns the type name of the type.
func (t *FloatType) Name() string {
	return t.TypeName
}

//go:generate stringer -linecomment -type FloatKind

// FloatKind represents the set of floating-point kinds.
type FloatKind uint8

// Floating-point kinds.
const (
	FloatKindHalf     FloatKind = iota // half
	FloatKindFloat                     // float
	FloatKindDouble                    // double
	FloatKindX86FP80                   // x86_fp80
	FloatKindFP128                     // fp128
	FloatKindPPCFP128                  // ppc_fp128
)

// --- [ MMX types ] -----------------------------------------------------------

// MMXType is an LLVM IR MMX type.
type MMXType struct {
	// Type name; or empty if not present.
	TypeName string
}

// Equal reports whether t and u are of equal type.
func (t *MMXType) Equal(u Type) bool {
	if _, ok := u.(*MMXType); ok {
		return true
	}
	return false
}

// String returns the string representation of the MMX type.
func (t *MMXType) String() string {
	if len(t.TypeName) > 0 {
		return enc.Local(t.TypeName)
	}
	return t.Def()
}

// Def returns the LLVM syntax representation of the definition of the type.
func (t *MMXType) Def() string {
	// 'x86_mmx'
	return "x86_mmx"
}

// Name returns the type name of the type.
func (t *MMXType) Name() string {
	return t.TypeName
}

// --- [ Pointer types ] -------------------------------------------------------

// PointerType is an LLVM IR pointer type.
type PointerType struct {
	// Type name; or empty if not present.
	TypeName string
	// Element type.
	ElemType Type
	// Address space; or zero value for default address space.
	AddrSpace AddrSpace
}

// NewPointer returns a new pointer type based on the given element type.
func NewPointer(elemType Type) *PointerType {
	return &PointerType{
		ElemType: elemType,
	}
}

// Equal reports whether t and u are of equal type.
func (t *PointerType) Equal(u Type) bool {
	// HACK: to prevent infinite loops (e.g. struct foo containing field of type
	// pointer to foo).
	return t.String() == u.String()
}

// String returns the string representation of the pointer type.
func (t *PointerType) String() string {
	if len(t.TypeName) > 0 {
		return enc.Local(t.TypeName)
	}
	return t.Def()
}

// Def returns the LLVM syntax representation of the definition of the type.
func (t *PointerType) Def() string {
	// Elem=Type AddrSpaceopt '*'
	buf := &strings.Builder{}
	buf.WriteString(t.ElemType.String())
	if t.AddrSpace != 0 {
		fmt.Fprintf(buf, " %v", t.AddrSpace)
	}
	buf.WriteString("*")
	return buf.String()
}

// Name returns the type name of the type.
func (t *PointerType) Name() string {
	return t.TypeName
}

// AddrSpace is an LLVM IR pointer type address space.
type AddrSpace int64

// String returns the string representation of the pointer type address space.
func (a AddrSpace) String() string {
	// 'addrspace' '(' N=UintLit ')'
	return fmt.Sprintf("addrspace(%d)", int64(a))
}

// --- [ Vector types ] --------------------------------------------------------

// VectorType is an LLVM IR vector type.
type VectorType struct {
	// Type name; or empty if not present.
	TypeName string
	// Vector length.
	Len int64
	// Element type.
	ElemType Type
}

// NewVector returns a new vector type based on the given vector length and
// element type.
func NewVector(len int64, elemType Type) *VectorType {
	return &VectorType{
		Len:      len,
		ElemType: elemType,
	}
}

// Equal reports whether t and u are of equal type.
func (t *VectorType) Equal(u Type) bool {
	if u, ok := u.(*VectorType); ok {
		if t.Len != u.Len {
			return false
		}
		return t.ElemType.Equal(u.ElemType)
	}
	return false
}

// String returns the string representation of the vector type.
func (t *VectorType) String() string {
	if len(t.TypeName) > 0 {
		return enc.Local(t.TypeName)
	}
	return t.Def()
}

// Def returns the LLVM syntax representation of the definition of the type.
func (t *VectorType) Def() string {
	// '<' Len=UintLit 'x' Elem=Type '>'
	return fmt.Sprintf("<%d x %v>", t.Len, t.ElemType)
}

// Name returns the type name of the type.
func (t *VectorType) Name() string {
	return t.TypeName
}

// --- [ Label types ] ---------------------------------------------------------

// LabelType is an LLVM IR label type.
type LabelType struct {
	// Type name; or empty if not present.
	TypeName string
}

// Equal reports whether t and u are of equal type.
func (t *LabelType) Equal(u Type) bool {
	if _, ok := u.(*LabelType); ok {
		return true
	}
	return false
}

// String returns the string representation of the label type.
func (t *LabelType) String() string {
	if len(t.TypeName) > 0 {
		return enc.Local(t.TypeName)
	}
	return t.Def()
}

// Def returns the LLVM syntax representation of the definition of the type.
func (t *LabelType) Def() string {
	// 'label'
	return "label"
}

// Name returns the type name of the type.
func (t *LabelType) Name() string {
	return t.TypeName
}

// --- [ Token types ] ---------------------------------------------------------

// TokenType is an LLVM IR token type.
type TokenType struct {
	// Type name; or empty if not present.
	TypeName string
}

// Equal reports whether t and u are of equal type.
func (t *TokenType) Equal(u Type) bool {
	if _, ok := u.(*TokenType); ok {
		return true
	}
	return false
}

// String returns the string representation of the token type.
func (t *TokenType) String() string {
	if len(t.TypeName) > 0 {
		return enc.Local(t.TypeName)
	}
	return t.Def()
}

// Def returns the LLVM syntax representation of the definition of the type.
func (t *TokenType) Def() string {
	// 'token'
	return "token"
}

// Name returns the type name of the type.
func (t *TokenType) Name() string {
	return t.TypeName
}

// --- [ Metadata types ] ------------------------------------------------------

// MetadataType is an LLVM IR metadata type.
type MetadataType struct {
	// Type name; or empty if not present.
	TypeName string
}

// Equal reports whether t and u are of equal type.
func (t *MetadataType) Equal(u Type) bool {
	if _, ok := u.(*MetadataType); ok {
		return true
	}
	return false
}

// String returns the string representation of the metadata type.
func (t *MetadataType) String() string {
	if len(t.TypeName) > 0 {
		return enc.Local(t.TypeName)
	}
	return t.Def()
}

// Def returns the LLVM syntax representation of the definition of the type.
func (t *MetadataType) Def() string {
	// 'metadata'
	return "metadata"
}

// Name returns the type name of the type.
func (t *MetadataType) Name() string {
	return t.TypeName
}

// --- [ Array types ] ---------------------------------------------------------

// ArrayType is an LLVM IR array type.
type ArrayType struct {
	// Type name; or empty if not present.
	TypeName string
	// Array length.
	Len int64
	// Element type.
	ElemType Type
}

// NewArray returns a new array type based on the given array length and element
// type.
func NewArray(len int64, elemType Type) *ArrayType {
	return &ArrayType{
		Len:      len,
		ElemType: elemType,
	}
}

// Equal reports whether t and u are of equal type.
func (t *ArrayType) Equal(u Type) bool {
	if u, ok := u.(*ArrayType); ok {
		if t.Len != u.Len {
			return false
		}
		return t.ElemType.Equal(u.ElemType)
	}
	return false
}

// String returns the string representation of the array type.
func (t *ArrayType) String() string {
	if len(t.TypeName) > 0 {
		return enc.Local(t.TypeName)
	}
	return t.Def()
}

// Def returns the LLVM syntax representation of the definition of the type.
func (t *ArrayType) Def() string {
	// '[' Len=UintLit 'x' Elem=Type ']'
	return fmt.Sprintf("[%d x %v]", t.Len, t.ElemType)
}

// Name returns the type name of the type.
func (t *ArrayType) Name() string {
	return t.TypeName
}

// --- [ Structure types ] -----------------------------------------------------

// StructType is an LLVM IR structure type.
type StructType struct {
	// Type name; or empty if not present.
	TypeName string
	// Packed memory layout.
	Packed bool
	// Struct fields.
	Fields []Type
	// Opaque struct type.
	Opaque bool
}

// NewStruct returns a new struct type based on the given field types.
func NewStruct(fields ...Type) *StructType {
	return &StructType{
		Fields: fields,
	}
}

// Equal reports whether t and u are of equal type.
func (t *StructType) Equal(u Type) bool {
	if u, ok := u.(*StructType); ok {
		if len(t.TypeName) > 0 || len(u.TypeName) > 0 {
			// Identified struct types are uniqued by type names, not by structural
			// identity.
			//
			// t or u is an identified struct type.
			return t.TypeName == u.TypeName
		}
		// Literal struct types are uniqued by structural identity.
		if t.Packed != u.Packed {
			return false
		}
		if len(t.Fields) != len(u.Fields) {
			return false
		}
		for i := range t.Fields {
			if !t.Fields[i].Equal(u.Fields[i]) {
				return false
			}
		}
		return true
	}
	return false
}

// String returns the string representation of the structure type.
func (t *StructType) String() string {
	if len(t.TypeName) > 0 {
		return enc.Local(t.TypeName)
	}
	return t.Def()
}

// Def returns the LLVM syntax representation of the definition of the type.
func (t *StructType) Def() string {
	// Opaque struct type.
	//
	//    'opaque'
	//
	// Struct type.
	//
	//    '{' Fields=(Type separator ',')+? '}'
	//
	// Packed struct type.
	//
	//    '<' '{' Fields=(Type separator ',')+? '}' '>'   -> PackedStructType
	if t.Opaque {
		return "opaque"
	}
	if len(t.Fields) == 0 {
		if t.Packed {
			return "<{}>"
		}
		return "{}"
	}
	buf := &strings.Builder{}
	if t.Packed {
		buf.WriteString("<")
	}
	buf.WriteString("{ ")
	for i, field := range t.Fields {
		if i != 0 {
			buf.WriteString(", ")
		}
		buf.WriteString(field.String())
	}
	buf.WriteString(" }")
	if t.Packed {
		buf.WriteString(">")
	}
	return buf.String()
}

// Name returns the type name of the type.
func (t *StructType) Name() string {
	return t.TypeName
}
