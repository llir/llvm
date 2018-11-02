// === [ Single value types ] ==================================================
//
// References:
//    http://llvm.org/docs/LangRef.html#single-value-types

package types

import (
	"fmt"

	"github.com/llir/llvm/internal/enc"
)

// --- [ integer ] -------------------------------------------------------------

// IntType represents an integer type.
//
// References:
//    http://llvm.org/docs/LangRef.html#integer-type
type IntType struct {
	// Type name alias.
	Name string
	// Bit size.
	Size int
}

// NewInt returns a new integer type based on the given bit size.
func NewInt(size int) *IntType {
	return &IntType{Size: size}
}

// String returns the LLVM syntax representation of the type.
func (t *IntType) String() string {
	if len(t.Name) > 0 {
		return enc.Local(t.Name)
	}
	return t.Def()
}

// Def returns the LLVM syntax representation of the definition of the type.
func (t *IntType) Def() string {
	return fmt.Sprintf("i%d", t.Size)
}

// Equal reports whether t and u are of equal type.
func (t *IntType) Equal(u Type) bool {
	if u, ok := u.(*IntType); ok {
		return t.Size == u.Size
	}
	return false
}

// GetName returns the name of the type.
func (t *IntType) GetName() string {
	return t.Name
}

// SetName sets the name of the type.
func (t *IntType) SetName(name string) {
	t.Name = name
}

// --- [ floating-point ] ------------------------------------------------------

// FloatType represents a floating-point type.
//
// References:
//    http://llvm.org/docs/LangRef.html#floating-point-types
type FloatType struct {
	// Type name alias.
	Name string
	// Floating-point kind.
	Kind FloatKind
}

// String returns the LLVM syntax representation of the type.
func (t *FloatType) String() string {
	if len(t.Name) > 0 {
		return enc.Local(t.Name)
	}
	return t.Def()
}

// Def returns the LLVM syntax representation of the definition of the type.
func (t *FloatType) Def() string {
	return t.Kind.String()
}

// Equal reports whether t and u are of equal type.
func (t *FloatType) Equal(u Type) bool {
	if u, ok := u.(*FloatType); ok {
		return t.Kind == u.Kind
	}
	return false
}

// GetName returns the name of the type.
func (t *FloatType) GetName() string {
	return t.Name
}

// SetName sets the name of the type.
func (t *FloatType) SetName(name string) {
	t.Name = name
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
	// Type name alias.
	Name string
	// Element type.
	Elem Type
	// Address space; or 0 for default address space.
	AddrSpace int
}

// NewPointer returns a new pointer type based on the given element type.
func NewPointer(elem Type) *PointerType {
	return &PointerType{Elem: elem}
}

// String returns the LLVM syntax representation of the type.
func (t *PointerType) String() string {
	if len(t.Name) > 0 {
		return enc.Local(t.Name)
	}
	return t.Def()
}

// Def returns the LLVM syntax representation of the definition of the type.
func (t *PointerType) Def() string {
	if t.AddrSpace != 0 {
		return fmt.Sprintf("%s addrspace(%d)*", t.Elem, t.AddrSpace)
	}
	return fmt.Sprintf("%s*", t.Elem)
}

// Equal reports whether t and u are of equal type.
func (t *PointerType) Equal(u Type) bool {
	if u, ok := u.(*PointerType); ok {
		return t.Elem.Equal(u.Elem)
	}
	return false
}

// GetName returns the name of the type.
func (t *PointerType) GetName() string {
	return t.Name
}

// SetName sets the name of the type.
func (t *PointerType) SetName(name string) {
	t.Name = name
}

// --- [ vector ] --------------------------------------------------------------

// VectorType represents a vector type.
//
// References:
//    http://llvm.org/docs/LangRef.html#vector-type
type VectorType struct {
	// Type name alias.
	Name string
	// Element type.
	Elem Type
	// Vector length.
	Len int64
}

// NewVector returns a new vector type based on the given element type and
// vector length.
func NewVector(elem Type, len int64) *VectorType {
	return &VectorType{Elem: elem, Len: len}
}

// String returns the LLVM syntax representation of the type.
func (t *VectorType) String() string {
	if len(t.Name) > 0 {
		return enc.Local(t.Name)
	}
	return t.Def()
}

// Def returns the LLVM syntax representation of the definition of the type.
func (t *VectorType) Def() string {
	return fmt.Sprintf("<%d x %s>",
		t.Len,
		t.Elem)
}

// Equal reports whether t and u are of equal type.
func (t *VectorType) Equal(u Type) bool {
	if u, ok := u.(*VectorType); ok {
		return t.Elem.Equal(u.Elem) && t.Len == u.Len
	}
	return false
}

// GetName returns the name of the type.
func (t *VectorType) GetName() string {
	return t.Name
}

// SetName sets the name of the type.
func (t *VectorType) SetName(name string) {
	t.Name = name
}
