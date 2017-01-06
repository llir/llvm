// The implementation of Float16 is heavily inspired by
// https://github.com/h2so5/half which is released into the public domain.

// Package floats implements encoding and decoding of IEEE 754 floating-point
// values.
package floats

import "math"

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
	panic("not yet implemented")
}

// Bytes returns the IEEE 754 binary representation of f as a byte slice.
func (f Float16) Bytes() []byte {
	panic("not yet implemented")
}

// Float32 returns the float32 representation of f.
func (f Float16) Float32() float32 {
	a := uint32(f.a)
	// 1 bit: sign
	sign := a >> 15
	// 5 bits: exponent
	exp := (a >> 10) & 0x1F
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
	if exp == 0 {
		exp32 = 0
	} else if exp == 0x1F {
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
	exp := (a >> 10) & 0x1F
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
	if exp == 0 {
		exp64 = 0
	} else if exp == 0x1F {
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

// NewFloat16FromFloat32 returns a new 16-bit floating-point value based on f.
func NewFloat16FromFloat32(f float32) Float16 {
	panic("not yet implemented")
}

// NewFloat16FromFloat64 returns a new 16-bit floating-point value based on f.
func NewFloat16FromFloat64(f float64) Float16 {
	panic("not yet implemented")
}

// NewFloat16FromString returns a new 16-bit floating-point value based on s,
// which is in hexadecimal format.
func NewFloat16FromString(s string) Float16 {
	panic("not yet implemented")
}

// NewFloat16FromBits returns a new 16-bit floating-point value based on bits.
func NewFloat16FromBits(bits uint16) Float16 {
	panic("not yet implemented")
}
