package floats

import (
	"fmt"
	"math"
)

// Float80 represents an 80-bit IEEE 754 extended precision floating-point
// value, in x86 extended precision format.
//
// References:
//    https://en.wikipedia.org/wiki/Extended_precision#x86_extended_precision_format
type Float80 struct {
	// Sign and exponent.
	//
	//    1 bit:   sign
	//    15 bits: exponent
	se uint16
	// Integer part and fraction.
	//
	//    1 bit:   integer part
	//    63 bits: fraction
	m uint64
}

// Bytes returns the x86 extended precision binary representation of f as a byte
// slice.
func (f Float80) Bytes() []byte {
	return []byte(f.String())
}

// String returns the IEEE 754 binary representation of f as a string,
// containing 10 bytes in hexadecimal format.
func (f Float80) String() string {
	return fmt.Sprintf("%04X%016X", f.se, f.m)
}

// Float64 returns the float64 representation of f.
func (f Float80) Float64() float64 {
	se := uint64(f.se)
	m := f.m
	// 1 bit: sign
	sign := se >> 15
	// 15 bits: exponent
	exp := se & 0x7FFF
	// Adjust for exponent bias.
	//
	// === [ binary64 ] =========================================================
	//
	// Exponent bias 1023.
	//
	//    +===========================+=======================+
	//    | Exponent (in binary)      | Notes                 |
	//    +===========================+=======================+
	//    | 00000000000               | zero/subnormal number |
	//    +---------------------------+-----------------------+
	//    | 00000000001 - 11111111110 | normalized value      |
	//    +---------------------------+-----------------------+
	//    | 11111111111               | infinity/NaN          |
	//    +---------------------------+-----------------------+
	//
	// References:
	//    https://en.wikipedia.org/wiki/Double-precision_floating-point_format#Exponent_encoding
	exp64 := exp - 16383 + 1023
	switch {
	case exp == 0:
		// exponent is all zeroes.
		exp64 = 0
	case exp == 0x7FFF:
		// exponent is all ones.
		exp64 = 0x7FF
	default:
	}
	// 63 bits: fraction
	frac := m & 0x7FFFFFFFFFFFFFFF
	// Sign, exponent and fraction of binary64.
	//
	//    1 bit:   sign
	//    11 bits: exponent
	//    52 bits: fraction
	//
	// References:
	//    https://en.wikipedia.org/wiki/Double-precision_floating-point_format#IEEE_754_double-precision_binary_floating-point_format:_binary64
	bits := sign<<63 | exp64<<52 | frac<<42
	return math.Float64frombits(bits)
}

// NewFloat80FromString returns a new 80-bit floating-point value based on s,
// which contains 20 bytes in hexadecimal format.
func NewFloat80FromString(s string) Float80 {
	return NewFloat80FromBytes([]byte(s))
}

// NewFloat80FromBytes returns a new 80-bit floating-point value based on b,
// which contains 20 bytes in hexadecimal format.
func NewFloat80FromBytes(b []byte) Float80 {
	var f Float80
	if len(b) != 20 {
		panic(fmt.Errorf("invalid length of float80 hexadecimal representation, expected 20, got %d", len(b)))
	}
	f.se = uint16(unhex(b[0])<<12 | unhex(b[1])<<8 | unhex(b[2])<<4 | unhex(b[3])<<0)
	f.m = uint64(unhex(b[4])<<60 | unhex(b[5])<<56 | unhex(b[6])<<52 | unhex(b[7])<<48 | unhex(b[8])<<44 | unhex(b[9])<<40 | unhex(b[10])<<36 | unhex(b[11])<<32 | unhex(b[12])<<28 | unhex(b[13])<<24 | unhex(b[14])<<20 | unhex(b[15])<<16 | unhex(b[16])<<12 | unhex(b[17])<<8 | unhex(b[18])<<4 | unhex(b[19])<<0)
	return f
}
