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

// NewVoid returns a new void type.
func NewVoid() *Void {
	return &Void{}
}

func (*Void) String() string {
	return "void"
}

// Int represents an integer type of arbitrary size.
//
// Examples:
//    i1, i8, i32
//
// References:
//    http://llvm.org/docs/LangRef.html#integer-type
type Int struct {
	// Size in number of bits.
	size int
}

// NewInt returns a new integer type of the specified size in number of bits.
func NewInt(size int) (*Int, error) {
	// Validate size (from 1 bit to 2^23-1 bits)
	if size <= 0 || size >= 1<<23 {
		return nil, fmt.Errorf("invalid integer size (%d)", size)
	}

	return &Int{size: size}, nil
}

// Size returns the size of typ in number of bits.
func (typ *Int) Size() int {
	return typ.size
}

func (typ *Int) String() string {
	return fmt.Sprintf("i%d", typ.size)
}

// Float represents a floating point type.
//
// Examples:
//    float, double, fp128
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
func (typ *Float) Kind() FloatKind {
	return typ.kind
}

func (typ *Float) String() string {
	switch typ.kind {
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

// Floating point types.
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

func (*Metadata) String() string {
	return "metadata"
}

// isType ensures that only types can be assigned to the Type interface.
func (*Void) isType()     {}
func (*Int) isType()      {}
func (*Float) isType()    {}
func (*MMX) isType()      {}
func (*Label) isType()    {}
func (*Metadata) isType() {}
