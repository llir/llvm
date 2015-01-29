package types

import "fmt"

// The Void type does not represent any value and has no size.
//
// Examples:
//    void
//
// References:
//    http://llvm.org/docs/LangRef.html#void-type
type Void struct{}

// NewVoid returns a void type.
func NewVoid() *Void {
	return &Void{}
}

// Equal returns true if the given types are equal, and false otherwise.
func (*Void) Equal(u Type) bool {
	_, ok := u.(*Void)
	return ok
}

func (*Void) String() string {
	return "void"
}

// Int represents an integer type of arbitrary size.
//
// Examples:
//    i1, i8, i32   ; iN, where 0 < N < 2^23
//
// References:
//    http://llvm.org/docs/LangRef.html#integer-type
type Int struct {
	// Size in number of bits.
	size int
}

// NewInt returns an integer type with the specified number of bits.
func NewInt(size int) (*Int, error) {
	// Validate size (from 1 bit to 2^23-1 bits)
	if size <= 0 || size >= 1<<23 {
		return nil, fmt.Errorf("invalid integer size (%d)", size)
	}

	return &Int{size: size}, nil
}

// Size returns the size of t in number of bits.
func (t *Int) Size() int {
	return t.size
}

// Equal returns true if the given types are equal, and false otherwise.
func (t *Int) Equal(u Type) bool {
	switch u := u.(type) {
	case *Int:
		return t.size == u.size
	}
	return false
}

func (t *Int) String() string {
	return fmt.Sprintf("i%d", t.size)
}

// Float represents a floating point type.
//
// Examples:
//    half, float, double, fp128, x86_fp80, ppc_fp128
//
// References:
//    http://llvm.org/docs/LangRef.html#floating-point-types
type Float struct {
	// Specifies the kind of the floating point type.
	kind FloatKind
}

// NewFloat returns a floating point type of the given kind.
func NewFloat(kind FloatKind) *Float {
	return &Float{kind: kind}
}

// Kind returns the kind of the floating point type.
func (t *Float) Kind() FloatKind {
	return t.kind
}

// Equal returns true if the given types are equal, and false otherwise.
func (t *Float) Equal(u Type) bool {
	switch u := u.(type) {
	case *Float:
		return t.kind == u.kind
	}
	return false
}

func (t *Float) String() string {
	switch t.kind {
	case Float16:
		return "half"
	case Float32:
		return "float"
	case Float64:
		return "double"
	case Float128:
		return "fp128"
	case X86Float80:
		return "x86_fp80"
	case PPCFloat128:
		return "ppc_fp128"
	}
	return "<unknown float type>"
}

// FloatKind specifies the kind of a floating point type.
type FloatKind int

// Floating point kinds.
const (
	Float16     FloatKind = iota // half:      16-bit floating point type
	Float32                      // float:     32-bit floating point type
	Float64                      // double:    64-bit floating point type
	Float128                     // fp128:     128-bit floating point type (112-bit mantissa)
	X86Float80                   // x86_fp80:  80-bit floating point type (x87)
	PPCFloat128                  // ppc_fp128: 128-bit floating point type (two 64-bits, PowerPC)
)

// MMX represents an MMX vector type (64 bits, x86 specific).
//
// Examples:
//    x86_mmx
//
// References:
//    http://llvm.org/docs/LangRef.html#x86-mmx-type
type MMX struct{}

// NewMMX returns an MMX vector type (64 bits, x86 specific).
func NewMMX() *MMX {
	return &MMX{}
}

// Equal returns true if the given types are equal, and false otherwise.
func (*MMX) Equal(u Type) bool {
	_, ok := u.(*MMX)
	return ok
}

func (*MMX) String() string {
	return "x86_mmx"
}

// Label represents a label type.
//
// Examples:
//    label
//
// References:
//    http://llvm.org/docs/LangRef.html#label-type
type Label struct{}

// NewLabel returns a label type.
func NewLabel() *Label {
	return &Label{}
}

// Equal returns true if the given types are equal, and false otherwise.
func (*Label) Equal(u Type) bool {
	_, ok := u.(*Label)
	return ok
}

func (*Label) String() string {
	return "label"
}

// Metadata represents a metadata type.
//
// Examples:
//    metadata
//
// References:
//    http://llvm.org/docs/LangRef.html#metadata-type
type Metadata struct{}

// NewMetadata returns a metadata type.
func NewMetadata() *Metadata {
	return &Metadata{}
}

// Equal returns true if the given types are equal, and false otherwise.
func (*Metadata) Equal(u Type) bool {
	_, ok := u.(*Metadata)
	return ok
}

func (*Metadata) String() string {
	return "metadata"
}
