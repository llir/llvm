// The implementation of Float16 is heavily inspired by
// https://github.com/h2so5/half which is released into the public domain.

// Package floats implements encoding and decoding of IEEE 754 floating-point
// values.
package floats

import (
	"fmt"
	"math"
)

// Float16 represents a 16-bit IEEE 754 half-precision floating-point value, in
// binary16 format.
//
// References:
//    https://en.wikipedia.org/wiki/Half-precision_floating-point_format#IEEE_754_half-precision_binary_floating-point_format:_binary16
type Float16 struct {
	// Sign, exponent and fraction.
	//
	//    1 bit:   sign
	//    5 bits:  exponent
	//    10 bits: fraction
	a uint16
}

// Bits returns the IEEE 754 binary representation of f.
func (f Float16) Bits() uint16 {
	return f.a
}

// Bytes returns the IEEE 754 binary representation of f as a byte slice,
// containing 4 bytes in hexadecimal format.
func (f Float16) Bytes() []byte {
	return []byte(f.String())
}

// String returns the IEEE 754 binary representation of f as a string,
// containing 4 bytes in hexadecimal format.
func (f Float16) String() string {
	return fmt.Sprintf("%04X", f.Bits())
}

// Float32 returns the float32 representation of f.
func (f Float16) Float32() float32 {
	a := uint32(f.a)
	// 1 bit: sign
	sign := a >> 15
	// 5 bits: exponent
	exp := a >> 10 & 0x1F
	// Adjust for exponent bias.
	//
	// === [ binary16 ] =========================================================
	//
	// Exponent bias 15.
	//
	//    +======================+=======================+
	//    | Exponent (in binary) | Notes                 |
	//    +======================+=======================+
	//    | 00000                | zero/subnormal number |
	//    +----------------------+-----------------------+
	//    | 00001 - 11110        | normalized value      |
	//    +----------------------+-----------------------+
	//    | 11111                | infinity/NaN          |
	//    +----------------------+-----------------------+
	//
	// References:
	//    https://en.wikipedia.org/wiki/Half-precision_floating-point_format#Exponent_encoding
	//
	// === [ binary32 ] =========================================================
	//
	// Exponent bias 127.
	//
	//    +===================+=======================+
	//    | Exponent (in hex) | Notes                 |
	//    +===================+=======================+
	//    | 00                | zero/subnormal number |
	//    +-------------------+-----------------------+
	//    | 01 - FE           | normalized value      |
	//    +-------------------+-----------------------+
	//    | FF                | infinity/NaN          |
	//    +-------------------+-----------------------+
	//
	// References:
	//    https://en.wikipedia.org/wiki/Single-precision_floating-point_format#Exponent_encoding
	exp32 := exp - 15 + 127
	switch {
	case exp == 0:
		exp32 = 0
	case exp == 0x1F:
		exp32 = 0xFF
	}
	// 10 bits: fraction
	frac := a & 0x3FF
	// Sign, exponent and fraction of binary32.
	//
	//    1 bit:   sign
	//    8 bits:  exponent
	//    23 bits: fraction
	//
	// References:
	//    https://en.wikipedia.org/wiki/Single-precision_floating-point_format#IEEE_754_single-precision_binary_floating-point_format:_binary32
	bits := sign<<31 | exp32<<23 | frac<<13
	return math.Float32frombits(bits)
}

// Float64 returns the float64 representation of f.
func (f Float16) Float64() float64 {
	a := uint64(f.a)
	// 1 bit: sign
	sign := a >> 15
	// 5 bits: exponent
	exp := a >> 10 & 0x1F
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
	exp64 := exp - 15 + 1023
	switch {
	case exp == 0:
		exp64 = 0
	case exp == 0x1F:
		exp64 = 0x7FF
	}
	// 10 bits: fraction
	frac := a & 0x3FF
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

// NewFloat16FromFloat32 returns the nearest 16-bit floating-point value for x
// and a bool indicating whether f represents x exactly.
func NewFloat16FromFloat32(x float32) (f Float16, exact bool) {
	exact = true
	// Sign, exponent and fraction of binary32.
	//
	//    1 bit:   sign
	//    8 bits:  exponent
	//    23 bits: fraction
	bits := math.Float32bits(x)
	// 1 bit: sign
	sign := uint16(bits >> 31)
	// 8 bits: exponent
	exp := bits >> 23 & 0xFF
	// 23 bits: fraction
	frac := bits & 0x7FFFFF

	// Sign, exponent and fraction of binary16.
	//
	//    1 bit:   sign
	//    5 bits:  exponent
	//    10 bits: fraction

	// 5 bits: exponent.
	//
	// Exponent bias 127 (binary32)
	// Exponent bias 15  (binary16)
	exp16 := int16(exp) - 127 + 15
	// 10 bits: fraction.
	//
	// Truncate 13 bits of the binary32 fraction.
	if frac&0x1FFF != 0 {
		exact = false
	}
	frac16 := uint16(frac >> 13)
	switch {
	case exp == 0:
		exp16 = 0
	case exp == 0xFF:
		exp16 = 0x1F
	default:
		if exp16 < 0x1 {
			// set float16 to zero if exp is too low.
			exp16 = 0
			frac16 = 0
			exact = false
		} else if exp16 > 0x1E {
			// set float16 to infinity if exp is too high.
			exp16 = 0x1F
			frac16 = 0
			exact = false
		}
	}
	a := sign<<15 | uint16(exp16<<10) | frac16
	return NewFloat16FromBits(a), exact
}

// NewFloat16FromFloat64 returns the nearest 16-bit floating-point value for x
// and a bool indicating whether f represents x exactly.
func NewFloat16FromFloat64(x float64) (f Float16, exact bool) {
	exact = true
	// Sign, exponent and fraction of binary64.
	//
	//    1 bit:   sign
	//    11 bits: exponent
	//    52 bits: fraction
	bits := math.Float64bits(x)
	// 1 bit: sign
	sign := uint16(bits >> 63)
	// 11 bits: exponent
	exp := bits >> 52 & 0x7FF
	// 52 bits: fraction
	frac := bits & 0xFFFFFFFFFFFFF

	// Sign, exponent and fraction of binary16.
	//
	//    1 bit:   sign
	//    5 bits:  exponent
	//    10 bits: fraction

	// 5 bits: exponent.
	//
	// Exponent bias 1023 (binary64)
	// Exponent bias 15   (binary16)
	exp16 := int16(exp) - 1023 + 15
	// 10 bits: fraction.
	//
	// Truncate 42 bits of the binary64 fraction.
	if frac&0x3FFFFFFFFFF != 0 {
		exact = false
	}
	frac16 := uint16(frac >> 42)
	switch {
	case exp == 0:
		exp16 = 0
	case exp == 0x7FF:
		exp16 = 0x1F
	default:
		if exp16 < 0x1 {
			// set float16 to zero if exp is too low.
			exp16 = 0
			frac16 = 0
			exact = false
		} else if exp16 > 0x1E {
			// set float16 to infinity if exp is too high.
			exp16 = 0x1F
			frac16 = 0
			exact = false
		}
	}
	a := sign<<15 | uint16(exp16<<10) | frac16
	return NewFloat16FromBits(a), exact
}

// NewFloat16FromString returns a new 16-bit floating-point value based on s,
// which contains 4 bytes in hexadecimal format.
func NewFloat16FromString(s string) Float16 {
	return NewFloat16FromBytes([]byte(s))
}

// NewFloat16FromBytes returns a new 16-bit floating-point value based on b,
// which contains 4 bytes in hexadecimal format.
func NewFloat16FromBytes(b []byte) Float16 {
	if len(b) != 4 {
		panic(fmt.Errorf("invalid length of float16 hexadecimal representation, expected 4, got %d", len(b)))
	}
	bits := uint16(unhex(b[0])<<12 | unhex(b[1])<<8 | unhex(b[2])<<4 | unhex(b[3])<<0)
	return NewFloat16FromBits(bits)
}

// NewFloat16FromBits returns a new 16-bit floating-point value based on bits.
func NewFloat16FromBits(bits uint16) Float16 {
	return Float16{a: bits}
}

// ### [ helper functions ] ####################################################

// unhex returns the numeric value represented by the hexadecimal digit b. It
// panics if b is not a hexadecimal digit.
func unhex(b byte) uint64 {
	// This is an adapted copy of the unhex function from the strconv package,
	// which is goverend by a BSD-style license.
	switch {
	case '0' <= b && b <= '9':
		return uint64(b - '0')
	case 'a' <= b && b <= 'f':
		return uint64(b - 'a' + 10)
	case 'A' <= b && b <= 'F':
		return uint64(b - 'A' + 10)
	}
	panic(fmt.Errorf("invalid byte; expected hexadecimal, got %q", b))
}
