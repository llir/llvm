package floats

import (
	"math"
	"testing"
)

// === [ float16 ] =============================================================

func TestFloat16Float32(t *testing.T) {
	golden := []struct {
		in   string
		want float32
	}{
		{in: "3C00", want: 1},
		{in: "4000", want: 2},
		{in: "C000", want: -2},
		{in: "7BFE", want: 65472},
		{in: "7BFF", want: 65504},
		{in: "FBFF", want: -65504},
		{in: "0000", want: 0},
		{in: "8000", want: float32(math.Copysign(0, -1))},
		{in: "7C00", want: float32(math.Inf(1))},
		{in: "FC00", want: float32(math.Inf(-1))},
		{in: "5B8F", want: 241.875},
		{in: "48C8", want: 9.5625},
	}
	for _, g := range golden {
		f := NewFloat16FromString(g.in)
		got := f.Float32()
		if got != g.want {
			t.Errorf("float32 mismatch for binary16 0x%s; expected %v, got %v", g.in, g.want, got)
			continue
		}
	}
}

func TestFloat16Float64(t *testing.T) {
	golden := []struct {
		in   uint16
		want float64
	}{
		{in: 0x3C00, want: 1},
		{in: 0x4000, want: 2},
		{in: 0xC000, want: -2},
		{in: 0x7BFE, want: 65472},
		{in: 0x7BFF, want: 65504},
		{in: 0xFBFF, want: -65504},
		{in: 0x0000, want: 0},
		{in: 0x8000, want: math.Copysign(0, -1)},
		{in: 0x7C00, want: math.Inf(1)},
		{in: 0xFC00, want: math.Inf(-1)},
		{in: 0x5B8F, want: 241.875},
		{in: 0x48C8, want: 9.5625},
	}
	for _, g := range golden {
		f := NewFloat16FromBits(g.in)
		got := f.Float64()
		if got != g.want {
			t.Errorf("float64 mismatch for binary16 0x%04X; expected %v, got %v", g.in, g.want, got)
			continue
		}
	}
}

func TestNewFloat16FromFloat32(t *testing.T) {
	golden := []struct {
		want uint16
		in   float32
	}{
		{want: 0x3C00, in: 1},
		{want: 0x4000, in: 2},
		{want: 0xC000, in: -2},
		{want: 0x7BFE, in: 65472},
		{want: 0x7BFF, in: 65504},
		{want: 0xFBFF, in: -65504},
		{want: 0x0000, in: 0},
		{want: 0x8000, in: float32(math.Copysign(0, -1))},
		{want: 0x7C00, in: float32(math.Inf(1))},
		{want: 0xFC00, in: float32(math.Inf(-1))},
		{want: 0x5B8F, in: 241.875},
		{want: 0x48C8, in: 9.5625},
	}
	for _, g := range golden {
		f, exact := NewFloat16FromFloat32(g.in)
		if !exact {
			t.Errorf("unable to represent %v exactly using binary16 format", g.in)
			continue
		}
		got := f.Bits()
		if got != g.want {
			t.Errorf("binary16 mismatch for float32 %v; expected 0x%04X, got 0x%04X", g.in, g.want, got)
			continue
		}
	}
}

func TestNewFloat16FromFloat64(t *testing.T) {
	golden := []struct {
		want uint16
		in   float64
	}{
		{want: 0x3C00, in: 1},
		{want: 0x4000, in: 2},
		{want: 0xC000, in: -2},
		{want: 0x7BFE, in: 65472},
		{want: 0x7BFF, in: 65504},
		{want: 0xFBFF, in: -65504},
		{want: 0x0000, in: 0},
		{want: 0x8000, in: math.Copysign(0, -1)},
		{want: 0x7C00, in: math.Inf(1)},
		{want: 0xFC00, in: math.Inf(-1)},
		{want: 0x5B8F, in: 241.875},
		{want: 0x48C8, in: 9.5625},
	}
	for _, g := range golden {
		f, exact := NewFloat16FromFloat64(g.in)
		if !exact {
			t.Errorf("unable to represent %v exactly using binary16 format", g.in)
			continue
		}
		got := f.Bits()
		if got != g.want {
			t.Errorf("binary16 mismatch for float64 %v; expected 0x%04X, got 0x%04X", g.in, g.want, got)
			continue
		}
	}
}

// === [ float80 ] =============================================================

func TestFloat80Float64(t *testing.T) {
	golden := []struct {
		in   string
		want float64
	}{
		{in: "00000000000000000000", want: 0.0},   // +0
		{in: "80000000000000000000", want: -0.0},  // -0
		{in: "3FFF8000000000000000", want: 1.0},   // 1
		{in: "40008000000000000000", want: 2.0},   // 2
		{in: "3FFC8000000000000000", want: 0.125}, // 0.125
		//{in: "7FFEFFFFFFFFFFFFFFFF", want: 1.18973149535723176505e+4932}, // max normal
		//{in: "00018000000000000000", want: 3.36210314311209350626e-4932}, // min positive normal
		//{in: "00007FFFFFFFFFFFFFFF", want: 3.36210314311209350608e-4932}, // max subnormal
		//{in: "00000000000000000001", want: 3.64519953188247460253e-4951}, // min positive subnormal
		{in: "7FFF8000000000000000", want: math.Inf(1)},  // +inf
		{in: "FFFF8000000000000000", want: math.Inf(-1)}, // -inf
		//{in: "7FFFFFFFFFFFFFFFFFFF", want: math.NaN()},   // QNaN - quiet NaN with greatest fraction
		//{in: "7FFFC000000000000000", want: math.NaN()},   // QNaN - quiet NaN with least fraction
		//{in: "7FFFBFFFFFFFFFFFFFFF", want: math.NaN()},   // SNaN - signaling NaN with greatest fraction
		//{in: "7FFF8000000000000001", want: math.NaN()},   // SNaN - signaling NaN with least fraction
	}
	for _, g := range golden {
		f := NewFloat80FromString(g.in)
		got := f.Float64()
		if got != g.want {
			t.Errorf("float64 mismatch for binary80 0x%s; expected %v, got %v", g.in, g.want, got)
			continue
		}
	}
}
