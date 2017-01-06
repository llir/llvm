package floats

// Float128 represents a 128-bit IEEE 754 quadruple-precision floating-point
// value, in binary128 format.
//
// References:
//    https://en.wikipedia.org/wiki/Quadruple-precision_floating-point_format#IEEE_754_quadruple-precision_binary_floating-point_format:_binary128
type Float128 struct {
	// Sign, exponent and fraction.
	//
	//    1 bit:    sign
	//    15 bits:  exponent
	//    112 bits: fraction
	a, b uint64
}

// Bytes returns the IEEE 754 binary representation of f as a byte slice.
func (f Float128) Bytes() []byte {
	panic("not yet implemented")
}
