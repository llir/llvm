// Package types declares the data types of LLVM IR.
package types

// Type represents a type.
type Type interface {
	// isType ensures that only types can be assigned to the Type interface.
	isType()
}

// TODO: Add FirstClass? Every type except void and function is a first class
// type.

// The Void type does not represent any value and has no size.
//
// Example:
//    void
type Void struct{}

// NewVoid returns a new void type.
func NewVoid() *Void {
	return &Void{}
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
func NewInt(n int) *Int {
	return &Int{n: n}
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

// Label represents a label type.
type Label struct{}

// NewLabel returns a label type.
func NewLabel() *Label {
	return &Label{}
}

// Metadata represents a metadata type.
type Metadata struct{}

// NewMetadata returns a metadata type.
func NewMetadata() *Metadata {
	return &Metadata{}
}

// IsInt returns true if typ is an integer type.
func IsInt(typ Type) bool {
	_, ok := typ.(*Int)
	return ok
}

// IsInts returns true if typ is an integer type or a vector of integers type.
func IsInts(typ Type) bool {
	if t, ok := typ.(*Vector); ok {
		return IsInts(t.Elem())
	}
	return IsInt(typ)
}

// IsFloat returns true if typ is a floating point type.
func IsFloat(typ Type) bool {
	_, ok := typ.(*Float)
	return ok
}

// IsFloats returns true if typ is a floating point type or a vector of floating
// points type.
func IsFloats(typ Type) bool {
	if t, ok := typ.(*Vector); ok {
		return IsFloats(t.Elem())
	}
	return IsFloat(typ)
}

// isType ensures that only types can be assigned to the Type interface.
func (*Void) isType()     {}
func (*Int) isType()      {}
func (*Float) isType()    {}
func (*MMX) isType()      {}
func (*Label) isType()    {}
func (*Metadata) isType() {}
