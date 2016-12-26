package ast

import "fmt"

// --- [ integer ] -------------------------------------------------------------

// IntType represents an integer type.
//
// References:
//    http://llvm.org/docs/LangRef.html#integer-type
type IntType struct {
	// Bit size.
	Size int
}

// --- [ floating-point ] ------------------------------------------------------

// FloatType represents a floating-point type.
//
// References:
//    http://llvm.org/docs/LangRef.html#floating-point-types
type FloatType struct {
	// Floating-point kind.
	Kind FloatKind
}

// FloatKind represents the set of floating-point kinds.
type FloatKind int

// Floating point kinds.
const (
	FloatKindIEEE_16           FloatKind = iota // half:      16-bit floating point type
	FloatKindIEEE_32                            // float:     32-bit floating point type
	FloatKindIEEE_64                            // double:    64-bit floating point type
	FloatKindIEEE_128                           // fp128:     128-bit floating point type (112-bit mantissa)
	FloatKindDoubleExtended_80                  // x86_fp80:  80-bit floating point type (x87)
	FloatKindDoubleDouble_128                   // ppc_fp128: 128-bit floating point type (two 64-bits, PowerPC)
)

// String returns the LLVM syntax representation of the floating-point kind.
func (kind FloatKind) String() string {
	switch kind {
	case FloatKindIEEE_16:
		return "half"
	case FloatKindIEEE_32:
		return "float"
	case FloatKindIEEE_64:
		return "double"
	case FloatKindIEEE_128:
		return "fp128"
	case FloatKindDoubleExtended_80:
		return "x86_fp80"
	case FloatKindDoubleDouble_128:
		return "ppc_fp128"
	}
	return fmt.Sprintf("<unknown floating-point kind %d>", int(kind))
}

// --- [ pointer ] -------------------------------------------------------------

// PointerType represents a pointer type.
//
// References:
//    http://llvm.org/docs/LangRef.html#pointer-type
type PointerType struct {
	// Element type.
	Elem Type
	// Address space.
	AddrSpace int64
}

// --- [ vector ] --------------------------------------------------------------

// VectorType represents a vector type.
//
// References:
//    http://llvm.org/docs/LangRef.html#vector-type
type VectorType struct {
	// Element type.
	Elem Type
	// Vector length.
	Len int64
}

// isType ensures that only types can be assigned to the ast.Type interface.
func (*IntType) isType()     {}
func (*FloatType) isType()   {}
func (*PointerType) isType() {}
func (*VectorType) isType()  {}
