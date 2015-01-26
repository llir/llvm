package types

import "fmt"

// The Void type does not represent any value and has no size.
//
// Example:
//    void
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
// Example:
//    i32
type Int struct {
	// Number of bits.
	n int
}

// NewInt returns a new integer type of the specified bit size.
func NewInt(n int) (*Int, error) {
	// Validate bit width (from 1 bit to 2^23-1 bits)
	if n <= 0 || n >= 1<<23 {
		return nil, fmt.Errorf("invalid integer bit width (%d)", n)
	}

	return &Int{n: n}, nil
}

func (typ *Int) String() string {
	return fmt.Sprintf("i%d", typ.n)
}

// Float represents a floating point type.
//
// Example:
//    double
type Float struct {
	kind FloatKind
}

// NewFloat returns a floating point type of the given kind.
func NewFloat(kind FloatKind) *Float {
	return &Float{kind: kind}
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
// Example:
//    x86_mmx
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
// Example:
//    label
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
// Example:
//    metadata
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
